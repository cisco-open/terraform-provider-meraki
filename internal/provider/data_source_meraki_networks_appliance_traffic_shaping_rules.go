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
	_ datasource.DataSource              = &NetworksApplianceTrafficShapingRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceTrafficShapingRulesDataSource{}
)

func NewNetworksApplianceTrafficShapingRulesDataSource() datasource.DataSource {
	return &NetworksApplianceTrafficShapingRulesDataSource{}
}

type NetworksApplianceTrafficShapingRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceTrafficShapingRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceTrafficShapingRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_traffic_shaping_rules"
}

func (d *NetworksApplianceTrafficShapingRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
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
								"priority": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceTrafficShapingRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceTrafficShapingRules NetworksApplianceTrafficShapingRules
	diags := req.Config.Get(ctx, &networksApplianceTrafficShapingRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceTrafficShapingRules")
		vvNetworkID := networksApplianceTrafficShapingRules.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceTrafficShapingRules(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceTrafficShapingRules",
				err.Error(),
			)
			return
		}

		networksApplianceTrafficShapingRules = ResponseApplianceGetNetworkApplianceTrafficShapingRulesItemToBody(networksApplianceTrafficShapingRules, response1)
		diags = resp.State.Set(ctx, &networksApplianceTrafficShapingRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceTrafficShapingRules struct {
	NetworkID types.String                                             `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceTrafficShapingRules `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingRules struct {
	DefaultRulesEnabled types.Bool                                                      `tfsdk:"default_rules_enabled"`
	Rules               *[]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRules `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingRulesRules struct {
	Definitions              *[]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitions            `tfsdk:"definitions"`
	DscpTagValue             types.Int64                                                                           `tfsdk:"dscp_tag_value"`
	PerClientBandwidthLimits *ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits `tfsdk:"per_client_bandwidth_limits"`
	Priority                 types.String                                                                          `tfsdk:"priority"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitions struct {
	Type      types.String                                                              `tfsdk:"type"`
	Value     types.String                                                              `tfsdk:"value"`
	ValueList types.Set                                                                 `tfsdk:"value_list"`
	ValueObj  *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj `tfsdk:"value_obj"`
}

type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits struct {
	BandwidthLimits *ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits `tfsdk:"bandwidth_limits"`
	Settings        types.String                                                                                         `tfsdk:"settings"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceTrafficShapingRulesItemToBody(state NetworksApplianceTrafficShapingRules, response *merakigosdk.ResponseApplianceGetNetworkApplianceTrafficShapingRules) NetworksApplianceTrafficShapingRules {
	itemState := ResponseApplianceGetNetworkApplianceTrafficShapingRules{
		DefaultRulesEnabled: func() types.Bool {
			if response.DefaultRulesEnabled != nil {
				return types.BoolValue(*response.DefaultRulesEnabled)
			}
			return types.Bool{}
		}(),
		Rules: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRules {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingRulesRules{
						Definitions: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitions {
							if rules.Definitions != nil {
								result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitions, len(*rules.Definitions))
								for i, definitions := range *rules.Definitions {
									result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitions{
										Type: types.StringValue(definitions.Type),
										Value: func() types.String {
											if definitions.Value == nil {
												return types.StringNull()
											}
											return types.StringValue(*definitions.Value)
										}(),
										ValueList: func() types.Set {
											if definitions.ValueList == nil {
												return types.SetNull(types.StringType)
											}
											return StringSliceToSet(*definitions.ValueList)
										}(),
										ValueObj: func() *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj {
											if definitions.ValueObj == nil {
												return nil
											}
											return &ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj{
												ID:   types.StringValue(definitions.ValueObj.ID),
												Name: types.StringValue(definitions.ValueObj.Name),
											}
										}(),
									}
								}
								return &result
							}
							return nil
						}(),
						DscpTagValue: func() types.Int64 {
							if rules.DscpTagValue != nil {
								return types.Int64Value(int64(*rules.DscpTagValue))
							}
							return types.Int64{}
						}(),
						PerClientBandwidthLimits: func() *ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits {
							if rules.PerClientBandwidthLimits != nil {
								return &ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits{
									BandwidthLimits: func() *ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits {
										if rules.PerClientBandwidthLimits.BandwidthLimits != nil {
											return &ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits{
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
										return nil
									}(),
									Settings: types.StringValue(rules.PerClientBandwidthLimits.Settings),
								}
							}
							return nil
						}(),
						Priority: types.StringValue(rules.Priority),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
