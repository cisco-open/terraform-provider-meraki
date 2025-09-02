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
	_ datasource.DataSource              = &NetworksWirelessDevicesConnectionStatsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessDevicesConnectionStatsDataSource{}
)

func NewNetworksWirelessDevicesConnectionStatsDataSource() datasource.DataSource {
	return &NetworksWirelessDevicesConnectionStatsDataSource{}
}

type NetworksWirelessDevicesConnectionStatsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessDevicesConnectionStatsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessDevicesConnectionStatsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_devices_connection_stats"
}

func (d *NetworksWirelessDevicesConnectionStatsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ap_tag": schema.StringAttribute{
				MarkdownDescription: `apTag query parameter. Filter results by AP Tag`,
				Optional:            true,
			},
			"band": schema.StringAttribute{
				MarkdownDescription: `band query parameter. Filter results by band (either '2.4', '5' or '6'). Note that data prior to February 2020 will not have band information.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"ssid": schema.Int64Attribute{
				MarkdownDescription: `ssid query parameter. Filter results by SSID`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 180 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 7 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 7 days.`,
				Optional:            true,
			},
			"vlan": schema.Int64Attribute{
				MarkdownDescription: `vlan query parameter. Filter results by VLAN`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetNetworkWirelessDevicesConnectionStats`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"connection_stats": schema.SingleNestedAttribute{
							MarkdownDescription: `The connection stats of the device`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"assoc": schema.Int64Attribute{
									MarkdownDescription: `The number of failed association attempts`,
									Computed:            true,
								},
								"auth": schema.Int64Attribute{
									MarkdownDescription: `The number of failed authentication attempts`,
									Computed:            true,
								},
								"dhcp": schema.Int64Attribute{
									MarkdownDescription: `The number of failed DHCP attempts`,
									Computed:            true,
								},
								"dns": schema.Int64Attribute{
									MarkdownDescription: `The number of failed DNS attempts`,
									Computed:            true,
								},
								"success": schema.Int64Attribute{
									MarkdownDescription: `The number of successful connection attempts`,
									Computed:            true,
								},
							},
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `The serial number for the device`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessDevicesConnectionStatsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessDevicesConnectionStats NetworksWirelessDevicesConnectionStats
	diags := req.Config.Get(ctx, &networksWirelessDevicesConnectionStats)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessDevicesConnectionStats")
		vvNetworkID := networksWirelessDevicesConnectionStats.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkWirelessDevicesConnectionStatsQueryParams{}

		queryParams1.T0 = networksWirelessDevicesConnectionStats.T0.ValueString()
		queryParams1.T1 = networksWirelessDevicesConnectionStats.T1.ValueString()
		queryParams1.Timespan = networksWirelessDevicesConnectionStats.Timespan.ValueFloat64()
		queryParams1.Band = networksWirelessDevicesConnectionStats.Band.ValueString()
		queryParams1.SSID = int(networksWirelessDevicesConnectionStats.SSID.ValueInt64())
		queryParams1.VLAN = int(networksWirelessDevicesConnectionStats.VLAN.ValueInt64())
		queryParams1.ApTag = networksWirelessDevicesConnectionStats.ApTag.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessDevicesConnectionStats(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessDevicesConnectionStats",
				err.Error(),
			)
			return
		}

		networksWirelessDevicesConnectionStats = ResponseWirelessGetNetworkWirelessDevicesConnectionStatsItemsToBody(networksWirelessDevicesConnectionStats, response1)
		diags = resp.State.Set(ctx, &networksWirelessDevicesConnectionStats)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessDevicesConnectionStats struct {
	NetworkID types.String                                                    `tfsdk:"network_id"`
	T0        types.String                                                    `tfsdk:"t0"`
	T1        types.String                                                    `tfsdk:"t1"`
	Timespan  types.Float64                                                   `tfsdk:"timespan"`
	Band      types.String                                                    `tfsdk:"band"`
	SSID      types.Int64                                                     `tfsdk:"ssid"`
	VLAN      types.Int64                                                     `tfsdk:"vlan"`
	ApTag     types.String                                                    `tfsdk:"ap_tag"`
	Items     *[]ResponseItemWirelessGetNetworkWirelessDevicesConnectionStats `tfsdk:"items"`
}

type ResponseItemWirelessGetNetworkWirelessDevicesConnectionStats struct {
	ConnectionStats *ResponseItemWirelessGetNetworkWirelessDevicesConnectionStatsConnectionStats `tfsdk:"connection_stats"`
	Serial          types.String                                                                 `tfsdk:"serial"`
}

type ResponseItemWirelessGetNetworkWirelessDevicesConnectionStatsConnectionStats struct {
	Assoc   types.Int64 `tfsdk:"assoc"`
	Auth    types.Int64 `tfsdk:"auth"`
	Dhcp    types.Int64 `tfsdk:"dhcp"`
	DNS     types.Int64 `tfsdk:"dns"`
	Success types.Int64 `tfsdk:"success"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessDevicesConnectionStatsItemsToBody(state NetworksWirelessDevicesConnectionStats, response *merakigosdk.ResponseWirelessGetNetworkWirelessDevicesConnectionStats) NetworksWirelessDevicesConnectionStats {
	var items []ResponseItemWirelessGetNetworkWirelessDevicesConnectionStats
	for _, item := range *response {
		itemState := ResponseItemWirelessGetNetworkWirelessDevicesConnectionStats{
			ConnectionStats: func() *ResponseItemWirelessGetNetworkWirelessDevicesConnectionStatsConnectionStats {
				if item.ConnectionStats != nil {
					return &ResponseItemWirelessGetNetworkWirelessDevicesConnectionStatsConnectionStats{
						Assoc: func() types.Int64 {
							if item.ConnectionStats.Assoc != nil {
								return types.Int64Value(int64(*item.ConnectionStats.Assoc))
							}
							return types.Int64{}
						}(),
						Auth: func() types.Int64 {
							if item.ConnectionStats.Auth != nil {
								return types.Int64Value(int64(*item.ConnectionStats.Auth))
							}
							return types.Int64{}
						}(),
						Dhcp: func() types.Int64 {
							if item.ConnectionStats.Dhcp != nil {
								return types.Int64Value(int64(*item.ConnectionStats.Dhcp))
							}
							return types.Int64{}
						}(),
						DNS: func() types.Int64 {
							if item.ConnectionStats.DNS != nil {
								return types.Int64Value(int64(*item.ConnectionStats.DNS))
							}
							return types.Int64{}
						}(),
						Success: func() types.Int64 {
							if item.ConnectionStats.Success != nil {
								return types.Int64Value(int64(*item.ConnectionStats.Success))
							}
							return types.Int64{}
						}(),
					}
				}
				return nil
			}(),
			Serial: func() types.String {
				if item.Serial != "" {
					return types.StringValue(item.Serial)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
