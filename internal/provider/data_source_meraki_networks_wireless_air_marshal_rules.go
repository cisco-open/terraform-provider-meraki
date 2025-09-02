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
	_ datasource.DataSource              = &NetworksWirelessAirMarshalRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessAirMarshalRulesDataSource{}
)

func NewNetworksWirelessAirMarshalRulesDataSource() datasource.DataSource {
	return &NetworksWirelessAirMarshalRulesDataSource{}
}

type NetworksWirelessAirMarshalRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessAirMarshalRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessAirMarshalRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_air_marshal_rules"
}

func (d *NetworksWirelessAirMarshalRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 7 days.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetNetworkWirelessAirMarshal`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"bssids": schema.SetNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"bssid": schema.StringAttribute{
										Computed: true,
									},
									"contained": schema.BoolAttribute{
										Computed: true,
									},
									"detected_by": schema.SetNestedAttribute{
										Computed: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{

												"device": schema.StringAttribute{
													Computed: true,
												},
												"rssi": schema.Int64Attribute{
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"channels": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"first_seen": schema.Int64Attribute{
							Computed: true,
						},
						"last_seen": schema.Int64Attribute{
							Computed: true,
						},
						"ssid": schema.StringAttribute{
							Computed: true,
						},
						"wired_last_seen": schema.Int64Attribute{
							Computed: true,
						},
						"wired_macs": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"wired_vlans": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessAirMarshalRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessAirMarshalRules NetworksWirelessAirMarshalRules
	diags := req.Config.Get(ctx, &networksWirelessAirMarshalRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessAirMarshal")
		vvNetworkID := networksWirelessAirMarshalRules.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkWirelessAirMarshalQueryParams{}

		queryParams1.T0 = networksWirelessAirMarshalRules.T0.ValueString()
		queryParams1.Timespan = networksWirelessAirMarshalRules.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessAirMarshal(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessAirMarshal",
				err.Error(),
			)
			return
		}

		networksWirelessAirMarshalRules = ResponseWirelessGetNetworkWirelessAirMarshalItemsToBody(networksWirelessAirMarshalRules, response1)
		diags = resp.State.Set(ctx, &networksWirelessAirMarshalRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessAirMarshalRules struct {
	NetworkID types.String                                        `tfsdk:"network_id"`
	T0        types.String                                        `tfsdk:"t0"`
	Timespan  types.Float64                                       `tfsdk:"timespan"`
	Items     *[]ResponseItemWirelessGetNetworkWirelessAirMarshal `tfsdk:"items"`
}

type ResponseItemWirelessGetNetworkWirelessAirMarshal struct {
	Bssids        *[]ResponseItemWirelessGetNetworkWirelessAirMarshalBssids `tfsdk:"bssids"`
	Channels      types.List                                                `tfsdk:"channels"`
	FirstSeen     types.Int64                                               `tfsdk:"first_seen"`
	LastSeen      types.Int64                                               `tfsdk:"last_seen"`
	SSID          types.String                                              `tfsdk:"ssid"`
	WiredLastSeen types.Int64                                               `tfsdk:"wired_last_seen"`
	WiredMacs     types.List                                                `tfsdk:"wired_macs"`
	WiredVLANs    types.List                                                `tfsdk:"wired_vlans"`
}

type ResponseItemWirelessGetNetworkWirelessAirMarshalBssids struct {
	Bssid      types.String                                                        `tfsdk:"bssid"`
	Contained  types.Bool                                                          `tfsdk:"contained"`
	DetectedBy *[]ResponseItemWirelessGetNetworkWirelessAirMarshalBssidsDetectedBy `tfsdk:"detected_by"`
}

type ResponseItemWirelessGetNetworkWirelessAirMarshalBssidsDetectedBy struct {
	Device types.String `tfsdk:"device"`
	Rssi   types.Int64  `tfsdk:"rssi"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessAirMarshalItemsToBody(state NetworksWirelessAirMarshalRules, response *merakigosdk.ResponseWirelessGetNetworkWirelessAirMarshal) NetworksWirelessAirMarshalRules {
	var items []ResponseItemWirelessGetNetworkWirelessAirMarshal
	for _, item := range *response {
		itemState := ResponseItemWirelessGetNetworkWirelessAirMarshal{
			Bssids: func() *[]ResponseItemWirelessGetNetworkWirelessAirMarshalBssids {
				if item.Bssids != nil {
					result := make([]ResponseItemWirelessGetNetworkWirelessAirMarshalBssids, len(*item.Bssids))
					for i, bssids := range *item.Bssids {
						result[i] = ResponseItemWirelessGetNetworkWirelessAirMarshalBssids{
							Bssid: func() types.String {
								if bssids.Bssid != "" {
									return types.StringValue(bssids.Bssid)
								}
								return types.String{}
							}(),
							Contained: func() types.Bool {
								if bssids.Contained != nil {
									return types.BoolValue(*bssids.Contained)
								}
								return types.Bool{}
							}(),
							DetectedBy: func() *[]ResponseItemWirelessGetNetworkWirelessAirMarshalBssidsDetectedBy {
								if bssids.DetectedBy != nil {
									result := make([]ResponseItemWirelessGetNetworkWirelessAirMarshalBssidsDetectedBy, len(*bssids.DetectedBy))
									for i, detectedBy := range *bssids.DetectedBy {
										result[i] = ResponseItemWirelessGetNetworkWirelessAirMarshalBssidsDetectedBy{
											Device: func() types.String {
												if detectedBy.Device != "" {
													return types.StringValue(detectedBy.Device)
												}
												return types.String{}
											}(),
											Rssi: func() types.Int64 {
												if detectedBy.Rssi != nil {
													return types.Int64Value(int64(*detectedBy.Rssi))
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
					}
					return &result
				}
				return nil
			}(),
			Channels: StringSliceToList(item.Channels),
			FirstSeen: func() types.Int64 {
				if item.FirstSeen != nil {
					return types.Int64Value(int64(*item.FirstSeen))
				}
				return types.Int64{}
			}(),
			LastSeen: func() types.Int64 {
				if item.LastSeen != nil {
					return types.Int64Value(int64(*item.LastSeen))
				}
				return types.Int64{}
			}(),
			SSID: func() types.String {
				if item.SSID != "" {
					return types.StringValue(item.SSID)
				}
				return types.String{}
			}(),
			WiredLastSeen: func() types.Int64 {
				if item.WiredLastSeen != nil {
					return types.Int64Value(int64(*item.WiredLastSeen))
				}
				return types.Int64{}
			}(),
			WiredMacs:  StringSliceToList(item.WiredMacs),
			WiredVLANs: StringSliceToList(item.WiredVLANs),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
