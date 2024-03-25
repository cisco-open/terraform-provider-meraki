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
	_ datasource.DataSource              = &NetworksApplianceFirewallOneToManyNatRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceFirewallOneToManyNatRulesDataSource{}
)

func NewNetworksApplianceFirewallOneToManyNatRulesDataSource() datasource.DataSource {
	return &NetworksApplianceFirewallOneToManyNatRulesDataSource{}
}

type NetworksApplianceFirewallOneToManyNatRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceFirewallOneToManyNatRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceFirewallOneToManyNatRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_one_to_many_nat_rules"
}

func (d *NetworksApplianceFirewallOneToManyNatRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"port_rules": schema.SetNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"allowed_ips": schema.ListAttribute{
												Computed:    true,
												ElementType: types.StringType,
											},
											"local_ip": schema.StringAttribute{
												Computed: true,
											},
											"local_port": schema.StringAttribute{
												Computed: true,
											},
											"name": schema.StringAttribute{
												Computed: true,
											},
											"protocol": schema.StringAttribute{
												Computed: true,
											},
											"public_port": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"public_ip": schema.StringAttribute{
									Computed: true,
								},
								"uplink": schema.StringAttribute{
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

func (d *NetworksApplianceFirewallOneToManyNatRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceFirewallOneToManyNatRules NetworksApplianceFirewallOneToManyNatRules
	diags := req.Config.Get(ctx, &networksApplianceFirewallOneToManyNatRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceFirewallOneToManyNatRules")
		vvNetworkID := networksApplianceFirewallOneToManyNatRules.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceFirewallOneToManyNatRules(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallOneToManyNatRules",
				err.Error(),
			)
			return
		}

		networksApplianceFirewallOneToManyNatRules = ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesItemToBody(networksApplianceFirewallOneToManyNatRules, response1)
		diags = resp.State.Set(ctx, &networksApplianceFirewallOneToManyNatRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceFirewallOneToManyNatRules struct {
	NetworkID types.String                                                   `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRules `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRules struct {
	Rules *[]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRules `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRules struct {
	PortRules *[]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRules `tfsdk:"port_rules"`
	PublicIP  types.String                                                                   `tfsdk:"public_ip"`
	Uplink    types.String                                                                   `tfsdk:"uplink"`
}

type ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRules struct {
	AllowedIPs types.List   `tfsdk:"allowed_ips"`
	LocalIP    types.String `tfsdk:"local_ip"`
	LocalPort  types.String `tfsdk:"local_port"`
	Name       types.String `tfsdk:"name"`
	Protocol   types.String `tfsdk:"protocol"`
	PublicPort types.String `tfsdk:"public_port"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesItemToBody(state NetworksApplianceFirewallOneToManyNatRules, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRules) NetworksApplianceFirewallOneToManyNatRules {
	itemState := ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRules{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRules {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRules{
						PortRules: func() *[]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRules {
							if rules.PortRules != nil {
								result := make([]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRules, len(*rules.PortRules))
								for i, portRules := range *rules.PortRules {
									result[i] = ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRules{
										AllowedIPs: StringSliceToList(portRules.AllowedIPs),
										LocalIP:    types.StringValue(portRules.LocalIP),
										LocalPort:  types.StringValue(portRules.LocalPort),
										Name:       types.StringValue(portRules.Name),
										Protocol:   types.StringValue(portRules.Protocol),
										PublicPort: types.StringValue(portRules.PublicPort),
									}
								}
								return &result
							}
							return &[]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRulesPortRules{}
						}(),
						PublicIP: types.StringValue(rules.PublicIP),
						Uplink:   types.StringValue(rules.Uplink),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceFirewallOneToManyNatRulesRules{}
		}(),
	}
	state.Item = &itemState
	return state
}
