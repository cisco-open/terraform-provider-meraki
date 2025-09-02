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
	_ datasource.DataSource              = &NetworksApplianceFirewallPortForwardingRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceFirewallPortForwardingRulesDataSource{}
)

func NewNetworksApplianceFirewallPortForwardingRulesDataSource() datasource.DataSource {
	return &NetworksApplianceFirewallPortForwardingRulesDataSource{}
}

type NetworksApplianceFirewallPortForwardingRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceFirewallPortForwardingRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceFirewallPortForwardingRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_port_forwarding_rules"
}

func (d *NetworksApplianceFirewallPortForwardingRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `An array of port forwarding rules`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"allowed_ips": schema.ListAttribute{
									MarkdownDescription: `An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges (or any)`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"lan_ip": schema.StringAttribute{
									MarkdownDescription: `IP address of the device subject to port forwarding`,
									Computed:            true,
								},
								"local_port": schema.StringAttribute{
									MarkdownDescription: `The port or port range that receives forwarded traffic from the WAN`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the rule`,
									Computed:            true,
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `Protocol the rule applies to`,
									Computed:            true,
								},
								"public_port": schema.StringAttribute{
									MarkdownDescription: `The port or port range forwarded to the host on the LAN`,
									Computed:            true,
								},
								"uplink": schema.StringAttribute{
									MarkdownDescription: `The physical WAN interface on which the traffic arrives; allowed values vary by appliance model and configuration`,
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

func (d *NetworksApplianceFirewallPortForwardingRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceFirewallPortForwardingRules NetworksApplianceFirewallPortForwardingRules
	diags := req.Config.Get(ctx, &networksApplianceFirewallPortForwardingRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceFirewallPortForwardingRules")
		vvNetworkID := networksApplianceFirewallPortForwardingRules.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceFirewallPortForwardingRules(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallPortForwardingRules",
				err.Error(),
			)
			return
		}

		networksApplianceFirewallPortForwardingRules = ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesItemToBody(networksApplianceFirewallPortForwardingRules, response1)
		diags = resp.State.Set(ctx, &networksApplianceFirewallPortForwardingRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceFirewallPortForwardingRules struct {
	NetworkID types.String                                                     `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules struct {
	AllowedIPs types.List   `tfsdk:"allowed_ips"`
	LanIP      types.String `tfsdk:"lan_ip"`
	LocalPort  types.String `tfsdk:"local_port"`
	Name       types.String `tfsdk:"name"`
	Protocol   types.String `tfsdk:"protocol"`
	PublicPort types.String `tfsdk:"public_port"`
	Uplink     types.String `tfsdk:"uplink"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesItemToBody(state NetworksApplianceFirewallPortForwardingRules, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules) NetworksApplianceFirewallPortForwardingRules {
	itemState := ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRules{
						AllowedIPs: StringSliceToList(rules.AllowedIPs),
						LanIP: func() types.String {
							if rules.LanIP != "" {
								return types.StringValue(rules.LanIP)
							}
							return types.String{}
						}(),
						LocalPort: func() types.String {
							if rules.LocalPort != "" {
								return types.StringValue(rules.LocalPort)
							}
							return types.String{}
						}(),
						Name: func() types.String {
							if rules.Name != "" {
								return types.StringValue(rules.Name)
							}
							return types.String{}
						}(),
						Protocol: func() types.String {
							if rules.Protocol != "" {
								return types.StringValue(rules.Protocol)
							}
							return types.String{}
						}(),
						PublicPort: func() types.String {
							if rules.PublicPort != "" {
								return types.StringValue(rules.PublicPort)
							}
							return types.String{}
						}(),
						Uplink: func() types.String {
							if rules.Uplink != "" {
								return types.StringValue(rules.Uplink)
							}
							return types.String{}
						}(),
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
