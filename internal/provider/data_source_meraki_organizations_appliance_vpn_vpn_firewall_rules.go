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
	_ datasource.DataSource              = &OrganizationsApplianceVpnVpnFirewallRulesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsApplianceVpnVpnFirewallRulesDataSource{}
)

func NewOrganizationsApplianceVpnVpnFirewallRulesDataSource() datasource.DataSource {
	return &OrganizationsApplianceVpnVpnFirewallRulesDataSource{}
}

type OrganizationsApplianceVpnVpnFirewallRulesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsApplianceVpnVpnFirewallRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsApplianceVpnVpnFirewallRulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_vpn_vpn_firewall_rules"
}

func (d *OrganizationsApplianceVpnVpnFirewallRulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"rules": schema.SetNestedAttribute{
						MarkdownDescription: `An ordered array of the firewall rules (not including the default rule)`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									MarkdownDescription: `Description of the rule (optional)`,
									Computed:            true,
								},
								"dest_cidr": schema.StringAttribute{
									MarkdownDescription: `Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'`,
									Computed:            true,
								},
								"dest_port": schema.StringAttribute{
									MarkdownDescription: `Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'`,
									Computed:            true,
								},
								"policy": schema.StringAttribute{
									MarkdownDescription: `'allow' or 'deny' traffic specified by this rule`,
									Computed:            true,
								},
								"protocol": schema.StringAttribute{
									MarkdownDescription: `The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')`,
									Computed:            true,
								},
								"src_cidr": schema.StringAttribute{
									MarkdownDescription: `Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)`,
									Computed:            true,
								},
								"src_port": schema.StringAttribute{
									MarkdownDescription: `Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'`,
									Computed:            true,
								},
								"syslog_enabled": schema.BoolAttribute{
									MarkdownDescription: `Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)`,
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

func (d *OrganizationsApplianceVpnVpnFirewallRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsApplianceVpnVpnFirewallRules OrganizationsApplianceVpnVpnFirewallRules
	diags := req.Config.Get(ctx, &organizationsApplianceVpnVpnFirewallRules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationApplianceVpnVpnFirewallRules")
		vvOrganizationID := organizationsApplianceVpnVpnFirewallRules.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetOrganizationApplianceVpnVpnFirewallRules(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceVpnVpnFirewallRules",
				err.Error(),
			)
			return
		}

		organizationsApplianceVpnVpnFirewallRules = ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRulesItemToBody(organizationsApplianceVpnVpnFirewallRules, response1)
		diags = resp.State.Set(ctx, &organizationsApplianceVpnVpnFirewallRules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsApplianceVpnVpnFirewallRules struct {
	OrganizationID types.String                                                  `tfsdk:"organization_id"`
	Item           *ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRules `tfsdk:"item"`
}

type ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRules struct {
	Rules *[]ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRulesRules `tfsdk:"rules"`
}

type ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRulesRules struct {
	Comment       types.String `tfsdk:"comment"`
	DestCidr      types.String `tfsdk:"dest_cidr"`
	DestPort      types.String `tfsdk:"dest_port"`
	Policy        types.String `tfsdk:"policy"`
	Protocol      types.String `tfsdk:"protocol"`
	SrcCidr       types.String `tfsdk:"src_cidr"`
	SrcPort       types.String `tfsdk:"src_port"`
	SyslogEnabled types.Bool   `tfsdk:"syslog_enabled"`
}

// ToBody
func ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRulesItemToBody(state OrganizationsApplianceVpnVpnFirewallRules, response *merakigosdk.ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRules) OrganizationsApplianceVpnVpnFirewallRules {
	itemState := ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRules{
		Rules: func() *[]ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRulesRules {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRulesRules, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetOrganizationApplianceVpnVpnFirewallRulesRules{
						Comment:  types.StringValue(rules.Comment),
						DestCidr: types.StringValue(rules.DestCidr),
						DestPort: types.StringValue(rules.DestPort),
						Policy:   types.StringValue(rules.Policy),
						Protocol: types.StringValue(rules.Protocol),
						SrcCidr:  types.StringValue(rules.SrcCidr),
						SrcPort:  types.StringValue(rules.SrcPort),
						SyslogEnabled: func() types.Bool {
							if rules.SyslogEnabled != nil {
								return types.BoolValue(*rules.SyslogEnabled)
							}
							return types.Bool{}
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
