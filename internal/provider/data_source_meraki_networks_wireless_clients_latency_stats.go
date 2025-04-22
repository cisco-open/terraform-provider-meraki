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
	_ datasource.DataSource              = &NetworksWirelessClientsLatencyStatsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessClientsLatencyStatsDataSource{}
)

func NewNetworksWirelessClientsLatencyStatsDataSource() datasource.DataSource {
	return &NetworksWirelessClientsLatencyStatsDataSource{}
}

type NetworksWirelessClientsLatencyStatsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessClientsLatencyStatsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessClientsLatencyStatsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_clients_latency_stats"
}

func (d *NetworksWirelessClientsLatencyStatsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"fields": schema.StringAttribute{
				MarkdownDescription: `fields query parameter. Partial selection: If present, this call will return only the selected fields of ["rawDistribution", "avg"]. All fields will be returned by default. Selected fields must be entered as a comma separated string.`,
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
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"latency_stats": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"background_traffic": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"avg": schema.Float64Attribute{
										Computed: true,
									},
									"raw_distribution": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{

											"status_0": schema.Int64Attribute{
												Computed: true,
											},
											"status_1": schema.Int64Attribute{
												Computed: true,
											},
											"status_1024": schema.Int64Attribute{
												Computed: true,
											},
											"status_128": schema.Int64Attribute{
												Computed: true,
											},
											"status_16": schema.Int64Attribute{
												Computed: true,
											},
											"status_2": schema.Int64Attribute{
												Computed: true,
											},
											"status_2048": schema.Int64Attribute{
												Computed: true,
											},
											"status_256": schema.Int64Attribute{
												Computed: true,
											},
											"status_32": schema.Int64Attribute{
												Computed: true,
											},
											"status_4": schema.Int64Attribute{
												Computed: true,
											},
											"status_512": schema.Int64Attribute{
												Computed: true,
											},
											"status_64": schema.Int64Attribute{
												Computed: true,
											},
											"status_8": schema.Int64Attribute{
												Computed: true,
											},
										},
									},
								},
							},
							"best_effort_traffic": schema.StringAttribute{
								Computed: true,
							},
							"video_traffic": schema.StringAttribute{
								Computed: true,
							},
							"voice_traffic": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"mac": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessClientsLatencyStatsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessClientsLatencyStats NetworksWirelessClientsLatencyStats
	diags := req.Config.Get(ctx, &networksWirelessClientsLatencyStats)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessClientLatencyStats")
		vvNetworkID := networksWirelessClientsLatencyStats.NetworkID.ValueString()
		vvClientID := networksWirelessClientsLatencyStats.ClientID.ValueString()
		queryParams1 := merakigosdk.GetNetworkWirelessClientLatencyStatsQueryParams{}

		queryParams1.T0 = networksWirelessClientsLatencyStats.T0.ValueString()
		queryParams1.T1 = networksWirelessClientsLatencyStats.T1.ValueString()
		queryParams1.Timespan = networksWirelessClientsLatencyStats.Timespan.ValueFloat64()
		queryParams1.Band = networksWirelessClientsLatencyStats.Band.ValueString()
		queryParams1.SSID = int(networksWirelessClientsLatencyStats.SSID.ValueInt64())
		queryParams1.VLAN = int(networksWirelessClientsLatencyStats.VLAN.ValueInt64())
		queryParams1.ApTag = networksWirelessClientsLatencyStats.ApTag.ValueString()
		queryParams1.Fields = networksWirelessClientsLatencyStats.Fields.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessClientLatencyStats(vvNetworkID, vvClientID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessClientLatencyStats",
				err.Error(),
			)
			return
		}

		networksWirelessClientsLatencyStats = ResponseWirelessGetNetworkWirelessClientLatencyStatsItemToBody(networksWirelessClientsLatencyStats, response1)
		diags = resp.State.Set(ctx, &networksWirelessClientsLatencyStats)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessClientsLatencyStats struct {
	NetworkID types.String                                          `tfsdk:"network_id"`
	ClientID  types.String                                          `tfsdk:"client_id"`
	T0        types.String                                          `tfsdk:"t0"`
	T1        types.String                                          `tfsdk:"t1"`
	Timespan  types.Float64                                         `tfsdk:"timespan"`
	Band      types.String                                          `tfsdk:"band"`
	SSID      types.Int64                                           `tfsdk:"ssid"`
	VLAN      types.Int64                                           `tfsdk:"vlan"`
	ApTag     types.String                                          `tfsdk:"ap_tag"`
	Fields    types.String                                          `tfsdk:"fields"`
	Item      *ResponseWirelessGetNetworkWirelessClientLatencyStats `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessClientLatencyStats struct {
	LatencyStats *ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStats `tfsdk:"latency_stats"`
	Mac          types.String                                                      `tfsdk:"mac"`
}

type ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStats struct {
	BackgroundTraffic *ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStatsBackgroundTraffic `tfsdk:"background_traffic"`
	BestEffortTraffic types.String                                                                       `tfsdk:"best_effort_traffic"`
	VideoTraffic      types.String                                                                       `tfsdk:"video_traffic"`
	VoiceTraffic      types.String                                                                       `tfsdk:"voice_traffic"`
}

type ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStatsBackgroundTraffic struct {
	Avg             types.Float64                                                                                     `tfsdk:"avg"`
	RawDistribution *ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStatsBackgroundTrafficRawDistribution `tfsdk:"raw_distribution"`
}

type ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStatsBackgroundTrafficRawDistribution struct {
	Status0    types.Int64 `tfsdk:"status_0"`
	Status1    types.Int64 `tfsdk:"status_1"`
	Status1024 types.Int64 `tfsdk:"status_1024"`
	Status128  types.Int64 `tfsdk:"status_128"`
	Status16   types.Int64 `tfsdk:"status_16"`
	Status2    types.Int64 `tfsdk:"status_2"`
	Status2048 types.Int64 `tfsdk:"status_2048"`
	Status256  types.Int64 `tfsdk:"status_256"`
	Status32   types.Int64 `tfsdk:"status_32"`
	Status4    types.Int64 `tfsdk:"status_4"`
	Status512  types.Int64 `tfsdk:"status_512"`
	Status64   types.Int64 `tfsdk:"status_64"`
	Status8    types.Int64 `tfsdk:"status_8"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessClientLatencyStatsItemToBody(state NetworksWirelessClientsLatencyStats, response *merakigosdk.ResponseWirelessGetNetworkWirelessClientLatencyStats) NetworksWirelessClientsLatencyStats {
	itemState := ResponseWirelessGetNetworkWirelessClientLatencyStats{
		LatencyStats: func() *ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStats {
			if response.LatencyStats != nil {
				return &ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStats{
					BackgroundTraffic: func() *ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStatsBackgroundTraffic {
						if response.LatencyStats.BackgroundTraffic != nil {
							return &ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStatsBackgroundTraffic{
								Avg: func() types.Float64 {
									if response.LatencyStats.BackgroundTraffic.Avg != nil {
										return types.Float64Value(float64(*response.LatencyStats.BackgroundTraffic.Avg))
									}
									return types.Float64{}
								}(),
								RawDistribution: func() *ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStatsBackgroundTrafficRawDistribution {
									if response.LatencyStats.BackgroundTraffic.RawDistribution != nil {
										return &ResponseWirelessGetNetworkWirelessClientLatencyStatsLatencyStatsBackgroundTrafficRawDistribution{
											Status0: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status0 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status0))
												}
												return types.Int64{}
											}(),
											Status1: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status1 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status1))
												}
												return types.Int64{}
											}(),
											Status1024: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status1024 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status1024))
												}
												return types.Int64{}
											}(),
											Status128: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status128 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status128))
												}
												return types.Int64{}
											}(),
											Status16: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status16 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status16))
												}
												return types.Int64{}
											}(),
											Status2: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status2 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status2))
												}
												return types.Int64{}
											}(),
											Status2048: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status2048 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status2048))
												}
												return types.Int64{}
											}(),
											Status256: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status256 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status256))
												}
												return types.Int64{}
											}(),
											Status32: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status32 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status32))
												}
												return types.Int64{}
											}(),
											Status4: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status4 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status4))
												}
												return types.Int64{}
											}(),
											Status512: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status512 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status512))
												}
												return types.Int64{}
											}(),
											Status64: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status64 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status64))
												}
												return types.Int64{}
											}(),
											Status8: func() types.Int64 {
												if response.LatencyStats.BackgroundTraffic.RawDistribution.Status8 != nil {
													return types.Int64Value(int64(*response.LatencyStats.BackgroundTraffic.RawDistribution.Status8))
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
					BestEffortTraffic: types.StringValue(response.LatencyStats.BestEffortTraffic),
					VideoTraffic:      types.StringValue(response.LatencyStats.VideoTraffic),
					VoiceTraffic:      types.StringValue(response.LatencyStats.VoiceTraffic),
				}
			}
			return nil
		}(),
		Mac: types.StringValue(response.Mac),
	}
	state.Item = &itemState
	return state
}
