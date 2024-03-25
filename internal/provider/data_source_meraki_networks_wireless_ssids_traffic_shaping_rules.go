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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksWirelessSSIDsTrafficShapingRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsTrafficShapingRulesDataSource{}
)

func NewNetworksWirelessSSIDsTrafficShapingRulesDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsTrafficShapingRulesDataSource{}
}

type NetworksWirelessSSIDsTrafficShapingRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsTrafficShapingRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsTrafficShapingRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_traffic_shaping_rules"
}

func (d *NetworksWirelessSSIDsTrafficShapingRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"default_rules_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"rules": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"definitions": schema.SetNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"type": schema.StringAttribute{
												Computed: true,
											},
											"value": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"dscp_tag_value": schema.Int64Attribute{
									Computed: true,
								},
								"pcp_tag_value": schema.Int64Attribute{
									Computed: true,
								},
								"per_client_bandwidth_limits": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"bandwidth_limits": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{

												"limit_down": schema.Int64Attribute{
													Computed: true,
												},
												"limit_up": schema.Int64Attribute{
													Computed: true,
												},
											},
										},
										"settings": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
					},
					"traffic_shaping_enabled": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessSSIDsTrafficShapingRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsTrafficShapingRules NetworksWirelessSSIDsTrafficShapingRules
	diags := req.Config.Get(ctx, &networksWirelessSSIDsTrafficShapingRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDTrafficShapingRules")
		vvNetworkID := networksWirelessSSIDsTrafficShapingRules.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsTrafficShapingRules.Number.ValueString()

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDTrafficShapingRules(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDTrafficShapingRules",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsTrafficShapingRules = ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesItemToBody(networksWirelessSSIDsTrafficShapingRules, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsTrafficShapingRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsTrafficShapingRules struct {
	NetworkID types.String                                               `tfsdk:"network_id"`
	Number    types.String                                               `tfsdk:"number"`
	Item      *ResponseWirelessGetNetworkWirelessSsidTrafficShapingRules `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSsidTrafficShapingRules struct {
	DefaultRulesEnabled   types.Bool                                                        `tfsdk:"default_rules_enabled"`
	Rules                 *[]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRules `tfsdk:"rules"`
	TrafficShapingEnabled types.Bool                                                        `tfsdk:"traffic_shaping_enabled"`
}

type ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRules struct {
	Definitions              *[]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesDefinitions            `tfsdk:"definitions"`
	DscpTagValue             types.Int64                                                                             `tfsdk:"dscp_tag_value"`
	PcpTagValue              types.Int64                                                                             `tfsdk:"pcp_tag_value"`
	PerClientBandwidthLimits *ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimits `tfsdk:"per_client_bandwidth_limits"`
}

type ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesDefinitions struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimits struct {
	BandwidthLimits *ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits `tfsdk:"bandwidth_limits"`
	Settings        types.String                                                                                           `tfsdk:"settings"`
}

type ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesItemToBody(state NetworksWirelessSSIDsTrafficShapingRules, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRules) NetworksWirelessSSIDsTrafficShapingRules {
	itemState := ResponseWirelessGetNetworkWirelessSsidTrafficShapingRules{
		DefaultRulesEnabled: func() types.Bool {
			if response.DefaultRulesEnabled != nil {
				return types.BoolValue(*response.DefaultRulesEnabled)
			}
			return types.Bool{}
		}(),
		Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRules {
			if response.Rules != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRules{
						Definitions: func() *[]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesDefinitions {
							if rules.Definitions != nil {
								result := make([]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesDefinitions, len(*rules.Definitions))
								for i, definitions := range *rules.Definitions {
									result[i] = ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesDefinitions{
										Type:  types.StringValue(definitions.Type),
										Value: types.StringValue(definitions.Value),
									}
								}
								return &result
							}
							return &[]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesDefinitions{}
						}(),
						DscpTagValue: func() types.Int64 {
							if rules.DscpTagValue != nil {
								return types.Int64Value(int64(*rules.DscpTagValue))
							}
							return types.Int64{}
						}(),
						PcpTagValue: func() types.Int64 {
							if rules.PcpTagValue != nil {
								return types.Int64Value(int64(*rules.PcpTagValue))
							}
							return types.Int64{}
						}(),
						PerClientBandwidthLimits: func() *ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimits {
							if rules.PerClientBandwidthLimits != nil {
								return &ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimits{
									BandwidthLimits: func() *ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits {
										if rules.PerClientBandwidthLimits.BandwidthLimits != nil {
											return &ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits{
												LimitDown: func() types.Int64 {
													if rules.PerClientBandwidthLimits.BandwidthLimits.LimitDown != nil {
														return types.Int64Value(int64(*rules.PerClientBandwidthLimits.BandwidthLimits.LimitDown))
													}
													return types.Int64{}
												}(),
												LimitUp: func() types.Int64 {
													if rules.PerClientBandwidthLimits.BandwidthLimits.LimitUp != nil {
														return types.Int64Value(int64(*rules.PerClientBandwidthLimits.BandwidthLimits.LimitUp))
													}
													return types.Int64{}
												}(),
											}
										}
										return &ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits{}
									}(),
									Settings: types.StringValue(rules.PerClientBandwidthLimits.Settings),
								}
							}
							return &ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimits{}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRules{}
		}(),
		TrafficShapingEnabled: func() types.Bool {
			if response.TrafficShapingEnabled != nil {
				return types.BoolValue(*response.TrafficShapingEnabled)
			}
			return types.Bool{}
		}(),
	}
	state.Item = &itemState
	return state
}
