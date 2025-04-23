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
	_ datasource.DataSource              = &OrganizationsSummaryTopDevicesModelsByUsageDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSummaryTopDevicesModelsByUsageDataSource{}
)

func NewOrganizationsSummaryTopDevicesModelsByUsageDataSource() datasource.DataSource {
	return &OrganizationsSummaryTopDevicesModelsByUsageDataSource{}
}

type OrganizationsSummaryTopDevicesModelsByUsageDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSummaryTopDevicesModelsByUsageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSummaryTopDevicesModelsByUsageDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_summary_top_devices_models_by_usage"
}

func (d *OrganizationsSummaryTopDevicesModelsByUsageDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_tag": schema.StringAttribute{
				MarkdownDescription: `deviceTag query parameter. Match result to an exact device tag`,
				Optional:            true,
			},
			"network_tag": schema.StringAttribute{
				MarkdownDescription: `networkTag query parameter. Match result to an exact network tag`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"quantity": schema.Int64Attribute{
				MarkdownDescription: `quantity query parameter. Set number of desired results to return. Default is 10.`,
				Optional:            true,
			},
			"ssid_name": schema.StringAttribute{
				MarkdownDescription: `ssidName query parameter. Filter results by ssid name`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 186 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 8 hours and be less than or equal to 186 days. The default is 1 day.`,
				Optional:            true,
			},
			"usage_uplink": schema.StringAttribute{
				MarkdownDescription: `usageUplink query parameter. Filter results by usage uplink`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"count": schema.Int64Attribute{
							MarkdownDescription: `Total number of devices per model`,
							Computed:            true,
						},
						"model": schema.StringAttribute{
							MarkdownDescription: `The device model`,
							Computed:            true,
						},
						"usage": schema.SingleNestedAttribute{
							MarkdownDescription: `Usage info in megabytes`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"average": schema.Float64Attribute{
									MarkdownDescription: `Average usage in megabytes`,
									Computed:            true,
								},
								"total": schema.Float64Attribute{
									MarkdownDescription: `Total usage in megabytes`,
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

func (d *OrganizationsSummaryTopDevicesModelsByUsageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSummaryTopDevicesModelsByUsage OrganizationsSummaryTopDevicesModelsByUsage
	diags := req.Config.Get(ctx, &organizationsSummaryTopDevicesModelsByUsage)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSummaryTopDevicesModelsByUsage")
		vvOrganizationID := organizationsSummaryTopDevicesModelsByUsage.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSummaryTopDevicesModelsByUsageQueryParams{}

		queryParams1.NetworkTag = organizationsSummaryTopDevicesModelsByUsage.NetworkTag.ValueString()
		queryParams1.DeviceTag = organizationsSummaryTopDevicesModelsByUsage.DeviceTag.ValueString()
		queryParams1.Quantity = int(organizationsSummaryTopDevicesModelsByUsage.Quantity.ValueInt64())
		queryParams1.SSIDName = organizationsSummaryTopDevicesModelsByUsage.SSIDName.ValueString()
		queryParams1.UsageUplink = organizationsSummaryTopDevicesModelsByUsage.UsageUplink.ValueString()
		queryParams1.T0 = organizationsSummaryTopDevicesModelsByUsage.T0.ValueString()
		queryParams1.T1 = organizationsSummaryTopDevicesModelsByUsage.T1.ValueString()
		queryParams1.Timespan = organizationsSummaryTopDevicesModelsByUsage.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSummaryTopDevicesModelsByUsage(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSummaryTopDevicesModelsByUsage",
				err.Error(),
			)
			return
		}

		organizationsSummaryTopDevicesModelsByUsage = ResponseOrganizationsGetOrganizationSummaryTopDevicesModelsByUsageItemsToBody(organizationsSummaryTopDevicesModelsByUsage, response1)
		diags = resp.State.Set(ctx, &organizationsSummaryTopDevicesModelsByUsage)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSummaryTopDevicesModelsByUsage struct {
	OrganizationID types.String                                                              `tfsdk:"organization_id"`
	NetworkTag     types.String                                                              `tfsdk:"network_tag"`
	DeviceTag      types.String                                                              `tfsdk:"device_tag"`
	Quantity       types.Int64                                                               `tfsdk:"quantity"`
	SSIDName       types.String                                                              `tfsdk:"ssid_name"`
	UsageUplink    types.String                                                              `tfsdk:"usage_uplink"`
	T0             types.String                                                              `tfsdk:"t0"`
	T1             types.String                                                              `tfsdk:"t1"`
	Timespan       types.Float64                                                             `tfsdk:"timespan"`
	Items          *[]ResponseItemOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage struct {
	Count types.Int64                                                                  `tfsdk:"count"`
	Model types.String                                                                 `tfsdk:"model"`
	Usage *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesModelsByUsageUsage `tfsdk:"usage"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesModelsByUsageUsage struct {
	Average types.Float64 `tfsdk:"average"`
	Total   types.Float64 `tfsdk:"total"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSummaryTopDevicesModelsByUsageItemsToBody(state OrganizationsSummaryTopDevicesModelsByUsage, response *merakigosdk.ResponseOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage) OrganizationsSummaryTopDevicesModelsByUsage {
	var items []ResponseItemOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationSummaryTopDevicesModelsByUsage{
			Count: func() types.Int64 {
				if item.Count != nil {
					return types.Int64Value(int64(*item.Count))
				}
				return types.Int64{}
			}(),
			Model: types.StringValue(item.Model),
			Usage: func() *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesModelsByUsageUsage {
				if item.Usage != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopDevicesModelsByUsageUsage{
						Average: func() types.Float64 {
							if item.Usage.Average != nil {
								return types.Float64Value(float64(*item.Usage.Average))
							}
							return types.Float64{}
						}(),
						Total: func() types.Float64 {
							if item.Usage.Total != nil {
								return types.Float64Value(float64(*item.Usage.Total))
							}
							return types.Float64{}
						}(),
					}
				}
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
