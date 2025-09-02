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
	_ datasource.DataSource              = &OrganizationsSummaryTopSSIDsByUsageDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSummaryTopSSIDsByUsageDataSource{}
)

func NewOrganizationsSummaryTopSSIDsByUsageDataSource() datasource.DataSource {
	return &OrganizationsSummaryTopSSIDsByUsageDataSource{}
}

type OrganizationsSummaryTopSSIDsByUsageDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSummaryTopSSIDsByUsageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSummaryTopSSIDsByUsageDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_summary_top_ssids_by_usage"
}

func (d *OrganizationsSummaryTopSSIDsByUsageDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationSummaryTopSsidsByUsage`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"clients": schema.SingleNestedAttribute{
							MarkdownDescription: `Clients info of the SSID`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"counts": schema.SingleNestedAttribute{
									MarkdownDescription: `Counts of the clients`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"total": schema.Int64Attribute{
											MarkdownDescription: `Total counts of the clients`,
											Computed:            true,
										},
									},
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the SSID`,
							Computed:            true,
						},
						"usage": schema.SingleNestedAttribute{
							MarkdownDescription: `Date usage of the SSID, in megabytes`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"downstream": schema.Float64Attribute{
									MarkdownDescription: `Downstream usage of the SSID`,
									Computed:            true,
								},
								"percentage": schema.Float64Attribute{
									MarkdownDescription: `Percentage usage of the SSID`,
									Computed:            true,
								},
								"total": schema.Float64Attribute{
									MarkdownDescription: `Total usage of the SSID`,
									Computed:            true,
								},
								"upstream": schema.Float64Attribute{
									MarkdownDescription: `Upstream usage of the SSID`,
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

func (d *OrganizationsSummaryTopSSIDsByUsageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSummaryTopSSIDsByUsage OrganizationsSummaryTopSSIDsByUsage
	diags := req.Config.Get(ctx, &organizationsSummaryTopSSIDsByUsage)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSummaryTopSSIDsByUsage")
		vvOrganizationID := organizationsSummaryTopSSIDsByUsage.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSummaryTopSSIDsByUsageQueryParams{}

		queryParams1.NetworkTag = organizationsSummaryTopSSIDsByUsage.NetworkTag.ValueString()
		queryParams1.DeviceTag = organizationsSummaryTopSSIDsByUsage.DeviceTag.ValueString()
		queryParams1.Quantity = int(organizationsSummaryTopSSIDsByUsage.Quantity.ValueInt64())
		queryParams1.SSIDName = organizationsSummaryTopSSIDsByUsage.SSIDName.ValueString()
		queryParams1.UsageUplink = organizationsSummaryTopSSIDsByUsage.UsageUplink.ValueString()
		queryParams1.T0 = organizationsSummaryTopSSIDsByUsage.T0.ValueString()
		queryParams1.T1 = organizationsSummaryTopSSIDsByUsage.T1.ValueString()
		queryParams1.Timespan = organizationsSummaryTopSSIDsByUsage.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSummaryTopSSIDsByUsage(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSummaryTopSSIDsByUsage",
				err.Error(),
			)
			return
		}

		organizationsSummaryTopSSIDsByUsage = ResponseOrganizationsGetOrganizationSummaryTopSSIDsByUsageItemsToBody(organizationsSummaryTopSSIDsByUsage, response1)
		diags = resp.State.Set(ctx, &organizationsSummaryTopSSIDsByUsage)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSummaryTopSSIDsByUsage struct {
	OrganizationID types.String                                                      `tfsdk:"organization_id"`
	NetworkTag     types.String                                                      `tfsdk:"network_tag"`
	DeviceTag      types.String                                                      `tfsdk:"device_tag"`
	Quantity       types.Int64                                                       `tfsdk:"quantity"`
	SSIDName       types.String                                                      `tfsdk:"ssid_name"`
	UsageUplink    types.String                                                      `tfsdk:"usage_uplink"`
	T0             types.String                                                      `tfsdk:"t0"`
	T1             types.String                                                      `tfsdk:"t1"`
	Timespan       types.Float64                                                     `tfsdk:"timespan"`
	Items          *[]ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsage `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsage struct {
	Clients *ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsageClients `tfsdk:"clients"`
	Name    types.String                                                           `tfsdk:"name"`
	Usage   *ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsageUsage   `tfsdk:"usage"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsageClients struct {
	Counts *ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsageClientsCounts `tfsdk:"counts"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsageClientsCounts struct {
	Total types.Int64 `tfsdk:"total"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsageUsage struct {
	Downstream types.Float64 `tfsdk:"downstream"`
	Percentage types.Float64 `tfsdk:"percentage"`
	Total      types.Float64 `tfsdk:"total"`
	Upstream   types.Float64 `tfsdk:"upstream"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSummaryTopSSIDsByUsageItemsToBody(state OrganizationsSummaryTopSSIDsByUsage, response *merakigosdk.ResponseOrganizationsGetOrganizationSummaryTopSSIDsByUsage) OrganizationsSummaryTopSSIDsByUsage {
	var items []ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsage
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsage{
			Clients: func() *ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsageClients {
				if item.Clients != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsageClients{
						Counts: func() *ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsageClientsCounts {
							if item.Clients.Counts != nil {
								return &ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsageClientsCounts{
									Total: func() types.Int64 {
										if item.Clients.Counts.Total != nil {
											return types.Int64Value(int64(*item.Clients.Counts.Total))
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
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
			Usage: func() *ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsageUsage {
				if item.Usage != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopSsidsByUsageUsage{
						Downstream: func() types.Float64 {
							if item.Usage.Downstream != nil {
								return types.Float64Value(float64(*item.Usage.Downstream))
							}
							return types.Float64{}
						}(),
						Percentage: func() types.Float64 {
							if item.Usage.Percentage != nil {
								return types.Float64Value(float64(*item.Usage.Percentage))
							}
							return types.Float64{}
						}(),
						Total: func() types.Float64 {
							if item.Usage.Total != nil {
								return types.Float64Value(float64(*item.Usage.Total))
							}
							return types.Float64{}
						}(),
						Upstream: func() types.Float64 {
							if item.Usage.Upstream != nil {
								return types.Float64Value(float64(*item.Usage.Upstream))
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
