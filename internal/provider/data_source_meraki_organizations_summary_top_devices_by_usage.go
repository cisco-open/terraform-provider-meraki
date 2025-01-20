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
	_ datasource.DataSource              = &OrganizationsSummaryTopDevicesByUsageDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSummaryTopDevicesByUsageDataSource{}
)

func NewOrganizationsSummaryTopDevicesByUsageDataSource() datasource.DataSource {
	return &OrganizationsSummaryTopDevicesByUsageDataSource{}
}

type OrganizationsSummaryTopDevicesByUsageDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSummaryTopDevicesByUsageDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSummaryTopDevicesByUsageDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_summary_top_devices_by_usage"
}

func (d *OrganizationsSummaryTopDevicesByUsageDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationSummaryTopDevicesByUsage`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"clients": schema.SingleNestedAttribute{
							MarkdownDescription: `Clients`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"counts": schema.SingleNestedAttribute{
									MarkdownDescription: `Counts of clients`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"total": schema.Int64Attribute{
											MarkdownDescription: `Total counts of clients`,
											Computed:            true,
										},
									},
								},
							},
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `Mac address of the device`,
							Computed:            true,
						},
						"model": schema.StringAttribute{
							MarkdownDescription: `Model of the device`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the device`,
							Computed:            true,
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network info`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `Network id`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Network name`,
									Computed:            true,
								},
							},
						},
						"product_type": schema.StringAttribute{
							MarkdownDescription: `Product type of the device`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial number of the device`,
							Computed:            true,
						},
						"usage": schema.SingleNestedAttribute{
							MarkdownDescription: `Data usage of the device`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"percentage": schema.Float64Attribute{
									MarkdownDescription: `Data usage of the device by percentage`,
									Computed:            true,
								},
								"total": schema.Float64Attribute{
									MarkdownDescription: `Total data usage of the device`,
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

func (d *OrganizationsSummaryTopDevicesByUsageDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSummaryTopDevicesByUsage OrganizationsSummaryTopDevicesByUsage
	diags := req.Config.Get(ctx, &organizationsSummaryTopDevicesByUsage)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSummaryTopDevicesByUsage")
		vvOrganizationID := organizationsSummaryTopDevicesByUsage.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSummaryTopDevicesByUsageQueryParams{}

		queryParams1.NetworkTag = organizationsSummaryTopDevicesByUsage.NetworkTag.ValueString()
		queryParams1.DeviceTag = organizationsSummaryTopDevicesByUsage.DeviceTag.ValueString()
		queryParams1.Quantity = int(organizationsSummaryTopDevicesByUsage.Quantity.ValueInt64())
		queryParams1.SSIDName = organizationsSummaryTopDevicesByUsage.SSIDName.ValueString()
		queryParams1.UsageUplink = organizationsSummaryTopDevicesByUsage.UsageUplink.ValueString()
		queryParams1.T0 = organizationsSummaryTopDevicesByUsage.T0.ValueString()
		queryParams1.T1 = organizationsSummaryTopDevicesByUsage.T1.ValueString()
		queryParams1.Timespan = organizationsSummaryTopDevicesByUsage.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSummaryTopDevicesByUsage(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSummaryTopDevicesByUsage",
				err.Error(),
			)
			return
		}

		organizationsSummaryTopDevicesByUsage = ResponseOrganizationsGetOrganizationSummaryTopDevicesByUsageItemsToBody(organizationsSummaryTopDevicesByUsage, response1)
		diags = resp.State.Set(ctx, &organizationsSummaryTopDevicesByUsage)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSummaryTopDevicesByUsage struct {
	OrganizationID types.String                                                        `tfsdk:"organization_id"`
	NetworkTag     types.String                                                        `tfsdk:"network_tag"`
	DeviceTag      types.String                                                        `tfsdk:"device_tag"`
	Quantity       types.Int64                                                         `tfsdk:"quantity"`
	SSIDName       types.String                                                        `tfsdk:"ssid_name"`
	UsageUplink    types.String                                                        `tfsdk:"usage_uplink"`
	T0             types.String                                                        `tfsdk:"t0"`
	T1             types.String                                                        `tfsdk:"t1"`
	Timespan       types.Float64                                                       `tfsdk:"timespan"`
	Items          *[]ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsage `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsage struct {
	Clients     *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageClients `tfsdk:"clients"`
	Mac         types.String                                                             `tfsdk:"mac"`
	Model       types.String                                                             `tfsdk:"model"`
	Name        types.String                                                             `tfsdk:"name"`
	Network     *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageNetwork `tfsdk:"network"`
	ProductType types.String                                                             `tfsdk:"product_type"`
	Serial      types.String                                                             `tfsdk:"serial"`
	Usage       *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageUsage   `tfsdk:"usage"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageClients struct {
	Counts *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageClientsCounts `tfsdk:"counts"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageClientsCounts struct {
	Total types.Int64 `tfsdk:"total"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageUsage struct {
	Percentage types.Float64 `tfsdk:"percentage"`
	Total      types.Float64 `tfsdk:"total"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSummaryTopDevicesByUsageItemsToBody(state OrganizationsSummaryTopDevicesByUsage, response *merakigosdk.ResponseOrganizationsGetOrganizationSummaryTopDevicesByUsage) OrganizationsSummaryTopDevicesByUsage {
	var items []ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsage
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsage{
			Clients: func() *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageClients {
				if item.Clients != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageClients{
						Counts: func() *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageClientsCounts {
							if item.Clients.Counts != nil {
								return &ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageClientsCounts{
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
			Mac:   types.StringValue(item.Mac),
			Model: types.StringValue(item.Model),
			Name:  types.StringValue(item.Name),
			Network: func() *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageNetwork {
				if item.Network != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageNetwork{
						ID:   types.StringValue(item.Network.ID),
						Name: types.StringValue(item.Network.Name),
					}
				}
				return nil
			}(),
			ProductType: types.StringValue(item.ProductType),
			Serial:      types.StringValue(item.Serial),
			Usage: func() *ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageUsage {
				if item.Usage != nil {
					return &ResponseItemOrganizationsGetOrganizationSummaryTopDevicesByUsageUsage{
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
