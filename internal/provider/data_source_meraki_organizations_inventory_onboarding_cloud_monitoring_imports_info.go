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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsInventoryOnboardingCloudMonitoringImportsInfoDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsInventoryOnboardingCloudMonitoringImportsInfoDataSource{}
)

func NewOrganizationsInventoryOnboardingCloudMonitoringImportsInfoDataSource() datasource.DataSource {
	return &OrganizationsInventoryOnboardingCloudMonitoringImportsInfoDataSource{}
}

type OrganizationsInventoryOnboardingCloudMonitoringImportsInfoDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsInventoryOnboardingCloudMonitoringImportsInfoDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsInventoryOnboardingCloudMonitoringImportsInfoDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_inventory_onboarding_cloud_monitoring_imports_info"
}

func (d *OrganizationsInventoryOnboardingCloudMonitoringImportsInfoDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"import_ids": schema.ListAttribute{
				MarkdownDescription: `importIds query parameter. import ids from an imports`,
				Required:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"device": schema.SingleNestedAttribute{
							MarkdownDescription: `Represents the details of an imported device.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"created": schema.BoolAttribute{
									MarkdownDescription: `Whether or not the device was successfully created in dashboard.`,
									Computed:            true,
								},
								"status": schema.StringAttribute{
									MarkdownDescription: `Represents the current state of importing the device.`,
									Computed:            true,
								},
								"url": schema.StringAttribute{
									MarkdownDescription: `The url to the device details page within dashboard.`,
									Computed:            true,
								},
							},
						},
						"import_id": schema.StringAttribute{
							MarkdownDescription: `Database ID for the new entity entry.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsInventoryOnboardingCloudMonitoringImportsInfoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsInventoryOnboardingCloudMonitoringImportsInfo OrganizationsInventoryOnboardingCloudMonitoringImportsInfo
	diags := req.Config.Get(ctx, &organizationsInventoryOnboardingCloudMonitoringImportsInfo)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationInventoryOnboardingCloudMonitoringImports")
		vvOrganizationID := organizationsInventoryOnboardingCloudMonitoringImportsInfo.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationInventoryOnboardingCloudMonitoringImportsQueryParams{}

		queryParams1.ImportIDs = elementsToStrings(ctx, organizationsInventoryOnboardingCloudMonitoringImportsInfo.ImportIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationInventoryOnboardingCloudMonitoringImports(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationInventoryOnboardingCloudMonitoringImports",
				err.Error(),
			)
			return
		}

		organizationsInventoryOnboardingCloudMonitoringImportsInfo = ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsItemsToBody(organizationsInventoryOnboardingCloudMonitoringImportsInfo, response1)
		diags = resp.State.Set(ctx, &organizationsInventoryOnboardingCloudMonitoringImportsInfo)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsInventoryOnboardingCloudMonitoringImportsInfo struct {
	OrganizationID types.String                                                                         `tfsdk:"organization_id"`
	ImportIDs      types.List                                                                           `tfsdk:"import_ids"`
	Items          *[]ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports struct {
	Device   *ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsDevice `tfsdk:"device"`
	ImportID types.String                                                                             `tfsdk:"import_id"`
}

type ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsDevice struct {
	Created types.Bool   `tfsdk:"created"`
	Status  types.String `tfsdk:"status"`
	URL     types.String `tfsdk:"url"`
}

// ToBody
func ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsItemsToBody(state OrganizationsInventoryOnboardingCloudMonitoringImportsInfo, response *merakigosdk.ResponseOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports) OrganizationsInventoryOnboardingCloudMonitoringImportsInfo {
	var items []ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImports{
			Device: func() *ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsDevice {
				if item.Device != nil {
					return &ResponseItemOrganizationsGetOrganizationInventoryOnboardingCloudMonitoringImportsDevice{
						Created: func() types.Bool {
							if item.Device.Created != nil {
								return types.BoolValue(*item.Device.Created)
							}
							return types.Bool{}
						}(),
						Status: func() types.String {
							if item.Device.Status != "" {
								return types.StringValue(item.Device.Status)
							}
							return types.String{}
						}(),
						URL: func() types.String {
							if item.Device.URL != "" {
								return types.StringValue(item.Device.URL)
							}
							return types.String{}
						}(),
					}
				}
				return nil
			}(),
			ImportID: func() types.String {
				if item.ImportID != "" {
					return types.StringValue(item.ImportID)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
