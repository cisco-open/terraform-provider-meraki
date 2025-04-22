// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsDevicesOverviewByModelDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesOverviewByModelDataSource{}
)

func NewOrganizationsDevicesOverviewByModelDataSource() datasource.DataSource {
	return &OrganizationsDevicesOverviewByModelDataSource{}
}

type OrganizationsDevicesOverviewByModelDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesOverviewByModelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesOverviewByModelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_overview_by_model"
}

func (d *OrganizationsDevicesOverviewByModelDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"models": schema.ListAttribute{
				MarkdownDescription: `models query parameter. Optional parameter to filter devices by one or more models. All returned devices will have a model that is an exact match.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter devices by networkId.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. Optional parameter to filter device by device product types. This filter uses multiple exact matches.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"counts": schema.SetNestedAttribute{
						MarkdownDescription: `Counts of devices per model`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"model": schema.StringAttribute{
									MarkdownDescription: `Device model`,
									Computed:            true,
								},
								"total": schema.Int64Attribute{
									MarkdownDescription: `Total number of devices for the model`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsDevicesOverviewByModelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevicesOverviewByModel OrganizationsDevicesOverviewByModel
	diags := req.Config.Get(ctx, &organizationsDevicesOverviewByModel)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevicesOverviewByModel")
		vvOrganizationID := organizationsDevicesOverviewByModel.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesOverviewByModelQueryParams{}

		queryParams1.Models = elementsToStrings(ctx, organizationsDevicesOverviewByModel.Models)
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsDevicesOverviewByModel.NetworkIDs)
		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsDevicesOverviewByModel.ProductTypes)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevicesOverviewByModel(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevicesOverviewByModel",
				err.Error(),
			)
			return
		}

		organizationsDevicesOverviewByModel = ResponseOrganizationsGetOrganizationDevicesOverviewByModelItemToBody(organizationsDevicesOverviewByModel, response1)
		diags = resp.State.Set(ctx, &organizationsDevicesOverviewByModel)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevicesOverviewByModel struct {
	OrganizationID types.String                                                `tfsdk:"organization_id"`
	Models         types.List                                                  `tfsdk:"models"`
	NetworkIDs     types.List                                                  `tfsdk:"network_ids"`
	ProductTypes   types.List                                                  `tfsdk:"product_types"`
	Item           *ResponseOrganizationsGetOrganizationDevicesOverviewByModel `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationDevicesOverviewByModel struct {
	Counts *[]ResponseOrganizationsGetOrganizationDevicesOverviewByModelCounts `tfsdk:"counts"`
}

type ResponseOrganizationsGetOrganizationDevicesOverviewByModelCounts struct {
	Model types.String `tfsdk:"model"`
	Total types.Int64  `tfsdk:"total"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesOverviewByModelItemToBody(state OrganizationsDevicesOverviewByModel, response *merakigosdk.ResponseOrganizationsGetOrganizationDevicesOverviewByModel) OrganizationsDevicesOverviewByModel {
	itemState := ResponseOrganizationsGetOrganizationDevicesOverviewByModel{
		Counts: func() *[]ResponseOrganizationsGetOrganizationDevicesOverviewByModelCounts {
			if response.Counts != nil {
				result := make([]ResponseOrganizationsGetOrganizationDevicesOverviewByModelCounts, len(*response.Counts))
				for i, counts := range *response.Counts {
					result[i] = ResponseOrganizationsGetOrganizationDevicesOverviewByModelCounts{
						Model: types.StringValue(counts.Model),
						Total: func() types.Int64 {
							if counts.Total != nil {
								return types.Int64Value(int64(*counts.Total))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
