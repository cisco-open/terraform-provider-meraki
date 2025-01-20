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
						MarkdownDescription: `Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.`,
						Computed:            true,
					},
					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `    An array of traffic shaping rules. Rules are applied in the order that
    they are specified in. An empty list (or null) means no rules. Note that
    you are allowed a maximum of 8 rules.
`,
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"definitions": schema.SetNestedAttribute{
									MarkdownDescription: `    A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
`,
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"type": schema.StringAttribute{
												MarkdownDescription: `The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.`,
												Computed:            true,
											},
											"value": schema.StringAttribute{
												MarkdownDescription: `    If "type" is 'host', 'port', 'ipRange' or 'localNet', then "value" must be a string, matching either
    a hostname (e.g. "somesite.com"), a port (e.g. 8080), or an IP range ("192.1.0.0",
    "192.1.0.0/16", or "10.1.0.0/16:80"). 'localNet' also supports CIDR notation, excluding
    custom ports.
     If "type" is 'application' or 'applicationCategory', then "value" must be an object
    with the structure { "id": "meraki:layer7/..." }, where "id" is the application category or
    application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories
    endpoint).
`,
												Computed: true,
											},
										},
									},
								},
								"dscp_tag_value": schema.Int64Attribute{
									MarkdownDescription: `    The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.
    For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
`,
									Computed: true,
								},
								"pcp_tag_value": schema.Int64Attribute{
									MarkdownDescription: `    The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).
    null means 'Do not set PCP tag'.
`,
									Computed: true,
								},
								"per_client_bandwidth_limits": schema.SingleNestedAttribute{
									MarkdownDescription: `    An object describing the bandwidth settings for your rule.
`,
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"bandwidth_limits": schema.SingleNestedAttribute{
											MarkdownDescription: `The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"limit_down": schema.Int64Attribute{
													MarkdownDescription: `The maximum download limit (integer, in Kbps).`,
													Computed:            true,
												},
												"limit_up": schema.Int64Attribute{
													MarkdownDescription: `The maximum upload limit (integer, in Kbps).`,
													Computed:            true,
												},
											},
										},
										"settings": schema.StringAttribute{
											MarkdownDescription: `How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.`,
											Computed:            true,
										},
									},
								},
							},
						},
					},
					"traffic_shaping_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether traffic shaping rules are applied to clients on your SSID.`,
						Computed:            true,
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

		// has_unknown_response: None

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
							return nil
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
										return nil
									}(),
									Settings: types.StringValue(rules.PerClientBandwidthLimits.Settings),
								}
							}
							return nil
						}(),
					}
				}
				return &result
			}
			return nil
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
