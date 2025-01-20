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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsDevicesStatusesOverviewDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesStatusesOverviewDataSource{}
)

func NewOrganizationsDevicesStatusesOverviewDataSource() datasource.DataSource {
	return &OrganizationsDevicesStatusesOverviewDataSource{}
}

type OrganizationsDevicesStatusesOverviewDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesStatusesOverviewDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesStatusesOverviewDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_statuses_overview"
}

func (d *OrganizationsDevicesStatusesOverviewDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. An optional parameter to filter device statuses by network.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. An optional parameter to filter device statuses by product type. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, wirelessController, and secureConnect.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"counts": schema.SingleNestedAttribute{
						MarkdownDescription: `counts`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"by_status": schema.SingleNestedAttribute{
								MarkdownDescription: `byStatus`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"alerting": schema.Int64Attribute{
										MarkdownDescription: `alerting count`,
										Computed:            true,
									},
									"dormant": schema.Int64Attribute{
										MarkdownDescription: `dormant count`,
										Computed:            true,
									},
									"offline": schema.Int64Attribute{
										MarkdownDescription: `offline count`,
										Computed:            true,
									},
									"online": schema.Int64Attribute{
										MarkdownDescription: `online count`,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsDevicesStatusesOverviewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevicesStatusesOverview OrganizationsDevicesStatusesOverview
	diags := req.Config.Get(ctx, &organizationsDevicesStatusesOverview)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevicesStatusesOverview")
		vvOrganizationID := organizationsDevicesStatusesOverview.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesStatusesOverviewQueryParams{}

		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsDevicesStatusesOverview.ProductTypes)
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsDevicesStatusesOverview.NetworkIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevicesStatusesOverview(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevicesStatusesOverview",
				err.Error(),
			)
			return
		}

		organizationsDevicesStatusesOverview = ResponseOrganizationsGetOrganizationDevicesStatusesOverviewItemToBody(organizationsDevicesStatusesOverview, response1)
		diags = resp.State.Set(ctx, &organizationsDevicesStatusesOverview)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevicesStatusesOverview struct {
	OrganizationID types.String                                                 `tfsdk:"organization_id"`
	ProductTypes   types.List                                                   `tfsdk:"product_types"`
	NetworkIDs     types.List                                                   `tfsdk:"network_ids"`
	Item           *ResponseOrganizationsGetOrganizationDevicesStatusesOverview `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationDevicesStatusesOverview struct {
	Counts *ResponseOrganizationsGetOrganizationDevicesStatusesOverviewCounts `tfsdk:"counts"`
}

type ResponseOrganizationsGetOrganizationDevicesStatusesOverviewCounts struct {
	ByStatus *ResponseOrganizationsGetOrganizationDevicesStatusesOverviewCountsByStatus `tfsdk:"by_status"`
}

type ResponseOrganizationsGetOrganizationDevicesStatusesOverviewCountsByStatus struct {
	Alerting types.Int64 `tfsdk:"alerting"`
	Dormant  types.Int64 `tfsdk:"dormant"`
	Offline  types.Int64 `tfsdk:"offline"`
	Online   types.Int64 `tfsdk:"online"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesStatusesOverviewItemToBody(state OrganizationsDevicesStatusesOverview, response *merakigosdk.ResponseOrganizationsGetOrganizationDevicesStatusesOverview) OrganizationsDevicesStatusesOverview {
	itemState := ResponseOrganizationsGetOrganizationDevicesStatusesOverview{
		Counts: func() *ResponseOrganizationsGetOrganizationDevicesStatusesOverviewCounts {
			if response.Counts != nil {
				return &ResponseOrganizationsGetOrganizationDevicesStatusesOverviewCounts{
					ByStatus: func() *ResponseOrganizationsGetOrganizationDevicesStatusesOverviewCountsByStatus {
						if response.Counts.ByStatus != nil {
							return &ResponseOrganizationsGetOrganizationDevicesStatusesOverviewCountsByStatus{
								Alerting: func() types.Int64 {
									if response.Counts.ByStatus.Alerting != nil {
										return types.Int64Value(int64(*response.Counts.ByStatus.Alerting))
									}
									return types.Int64{}
								}(),
								Dormant: func() types.Int64 {
									if response.Counts.ByStatus.Dormant != nil {
										return types.Int64Value(int64(*response.Counts.ByStatus.Dormant))
									}
									return types.Int64{}
								}(),
								Offline: func() types.Int64 {
									if response.Counts.ByStatus.Offline != nil {
										return types.Int64Value(int64(*response.Counts.ByStatus.Offline))
									}
									return types.Int64{}
								}(),
								Online: func() types.Int64 {
									if response.Counts.ByStatus.Online != nil {
										return types.Int64Value(int64(*response.Counts.ByStatus.Online))
									}
									return types.Int64{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
