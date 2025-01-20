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
	_ datasource.DataSource              = &NetworksWirelessClientsConnectionStatsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessClientsConnectionStatsDataSource{}
)

func NewNetworksWirelessClientsConnectionStatsDataSource() datasource.DataSource {
	return &NetworksWirelessClientsConnectionStatsDataSource{}
}

type NetworksWirelessClientsConnectionStatsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessClientsConnectionStatsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessClientsConnectionStatsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_clients_connection_stats"
}

func (d *NetworksWirelessClientsConnectionStatsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"client_id": schema.StringAttribute{
				MarkdownDescription: `clientId path parameter. Client ID`,
				Required:            true,
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
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"connection_stats": schema.SingleNestedAttribute{
						MarkdownDescription: `Connection stats`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"assoc": schema.Int64Attribute{
								MarkdownDescription: `Association count`,
								Computed:            true,
							},
							"auth": schema.Int64Attribute{
								MarkdownDescription: `Authorization count`,
								Computed:            true,
							},
							"dhcp": schema.Int64Attribute{
								MarkdownDescription: `DHCP count`,
								Computed:            true,
							},
							"success": schema.Int64Attribute{
								MarkdownDescription: `successful count`,
								Computed:            true,
							},
						},
					},
					"mac": schema.StringAttribute{
						MarkdownDescription: `MAC address of the client`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessClientsConnectionStatsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessClientsConnectionStats NetworksWirelessClientsConnectionStats
	diags := req.Config.Get(ctx, &networksWirelessClientsConnectionStats)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessClientConnectionStats")
		vvNetworkID := networksWirelessClientsConnectionStats.NetworkID.ValueString()
		vvClientID := networksWirelessClientsConnectionStats.ClientID.ValueString()
		queryParams1 := merakigosdk.GetNetworkWirelessClientConnectionStatsQueryParams{}

		queryParams1.T0 = networksWirelessClientsConnectionStats.T0.ValueString()
		queryParams1.T1 = networksWirelessClientsConnectionStats.T1.ValueString()
		queryParams1.Timespan = networksWirelessClientsConnectionStats.Timespan.ValueFloat64()
		queryParams1.Band = networksWirelessClientsConnectionStats.Band.ValueString()
		queryParams1.SSID = int(networksWirelessClientsConnectionStats.SSID.ValueInt64())
		queryParams1.VLAN = int(networksWirelessClientsConnectionStats.VLAN.ValueInt64())
		queryParams1.ApTag = networksWirelessClientsConnectionStats.ApTag.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessClientConnectionStats(vvNetworkID, vvClientID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessClientConnectionStats",
				err.Error(),
			)
			return
		}

		networksWirelessClientsConnectionStats = ResponseWirelessGetNetworkWirelessClientConnectionStatsItemToBody(networksWirelessClientsConnectionStats, response1)
		diags = resp.State.Set(ctx, &networksWirelessClientsConnectionStats)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessClientsConnectionStats struct {
	NetworkID types.String                                             `tfsdk:"network_id"`
	ClientID  types.String                                             `tfsdk:"client_id"`
	T0        types.String                                             `tfsdk:"t0"`
	T1        types.String                                             `tfsdk:"t1"`
	Timespan  types.Float64                                            `tfsdk:"timespan"`
	Band      types.String                                             `tfsdk:"band"`
	SSID      types.Int64                                              `tfsdk:"ssid"`
	VLAN      types.Int64                                              `tfsdk:"vlan"`
	ApTag     types.String                                             `tfsdk:"ap_tag"`
	Item      *ResponseWirelessGetNetworkWirelessClientConnectionStats `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessClientConnectionStats struct {
	ConnectionStats *ResponseWirelessGetNetworkWirelessClientConnectionStatsConnectionStats `tfsdk:"connection_stats"`
	Mac             types.String                                                            `tfsdk:"mac"`
}

type ResponseWirelessGetNetworkWirelessClientConnectionStatsConnectionStats struct {
	Assoc   types.Int64 `tfsdk:"assoc"`
	Auth    types.Int64 `tfsdk:"auth"`
	Dhcp    types.Int64 `tfsdk:"dhcp"`
	Success types.Int64 `tfsdk:"success"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessClientConnectionStatsItemToBody(state NetworksWirelessClientsConnectionStats, response *merakigosdk.ResponseWirelessGetNetworkWirelessClientConnectionStats) NetworksWirelessClientsConnectionStats {
	itemState := ResponseWirelessGetNetworkWirelessClientConnectionStats{
		ConnectionStats: func() *ResponseWirelessGetNetworkWirelessClientConnectionStatsConnectionStats {
			if response.ConnectionStats != nil {
				return &ResponseWirelessGetNetworkWirelessClientConnectionStatsConnectionStats{
					Assoc: func() types.Int64 {
						if response.ConnectionStats.Assoc != nil {
							return types.Int64Value(int64(*response.ConnectionStats.Assoc))
						}
						return types.Int64{}
					}(),
					Auth: func() types.Int64 {
						if response.ConnectionStats.Auth != nil {
							return types.Int64Value(int64(*response.ConnectionStats.Auth))
						}
						return types.Int64{}
					}(),
					Dhcp: func() types.Int64 {
						if response.ConnectionStats.Dhcp != nil {
							return types.Int64Value(int64(*response.ConnectionStats.Dhcp))
						}
						return types.Int64{}
					}(),
					Success: func() types.Int64 {
						if response.ConnectionStats.Success != nil {
							return types.Int64Value(int64(*response.ConnectionStats.Success))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		Mac: types.StringValue(response.Mac),
	}
	state.Item = &itemState
	return state
}
