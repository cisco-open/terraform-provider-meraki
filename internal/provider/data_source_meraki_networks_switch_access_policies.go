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
	_ datasource.DataSource              = &NetworksSwitchAccessPoliciesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchAccessPoliciesDataSource{}
)

func NewNetworksSwitchAccessPoliciesDataSource() datasource.DataSource {
	return &NetworksSwitchAccessPoliciesDataSource{}
}

type NetworksSwitchAccessPoliciesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchAccessPoliciesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchAccessPoliciesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_access_policies"
}

func (d *NetworksSwitchAccessPoliciesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"access_policy_number": schema.StringAttribute{
				MarkdownDescription: `accessPolicyNumber path parameter. Access policy number`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"access_policy_type": schema.StringAttribute{
						MarkdownDescription: `Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.`,
						Computed:            true,
					},
					"counts": schema.SingleNestedAttribute{
						MarkdownDescription: `Counts associated with the access policy`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"ports": schema.SingleNestedAttribute{
								MarkdownDescription: `Counts associated with ports`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"with_this_policy": schema.Int64Attribute{
										MarkdownDescription: `Number of ports in the network with this policy. For template networks, this is the number of template ports (not child ports) with this policy.`,
										Computed:            true,
									},
								},
							},
						},
					},
					"dot1x": schema.SingleNestedAttribute{
						MarkdownDescription: `802.1x Settings`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"control_direction": schema.StringAttribute{
								MarkdownDescription: `Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'`,
								Computed:            true,
							},
						},
					},
					"guest_port_bouncing": schema.BoolAttribute{
						MarkdownDescription: `If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers`,
						Computed:            true,
					},
					"guest_vlan_id": schema.Int64Attribute{
						MarkdownDescription: `ID for the guest VLAN allow unauthorized devices access to limited network resources`,
						Computed:            true,
					},
					"host_mode": schema.StringAttribute{
						MarkdownDescription: `Choose the Host Mode for the access policy.`,
						Computed:            true,
					},
					"increase_access_speed": schema.BoolAttribute{
						MarkdownDescription: `Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the access policy`,
						Computed:            true,
					},
					"radius": schema.SingleNestedAttribute{
						MarkdownDescription: `Object for RADIUS Settings`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"cache": schema.SingleNestedAttribute{
								MarkdownDescription: `Object for RADIUS Cache Settings`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Enable to cache authorization and authentication responses on the RADIUS server`,
										Computed:            true,
									},
									"timeout": schema.Int64Attribute{
										MarkdownDescription: `If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication`,
										Computed:            true,
									},
								},
							},
							"critical_auth": schema.SingleNestedAttribute{
								MarkdownDescription: `Critical auth settings for when authentication is rejected by the RADIUS server`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"data_vlan_id": schema.Int64Attribute{
										MarkdownDescription: `VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth`,
										Computed:            true,
									},
									"suspend_port_bounce": schema.BoolAttribute{
										MarkdownDescription: `Enable to suspend port bounce when RADIUS servers are unreachable`,
										Computed:            true,
									},
									"voice_vlan_id": schema.Int64Attribute{
										MarkdownDescription: `VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth`,
										Computed:            true,
									},
								},
							},
							"failed_auth_vlan_id": schema.Int64Attribute{
								MarkdownDescription: `VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth`,
								Computed:            true,
							},
							"re_authentication_interval": schema.Int64Attribute{
								MarkdownDescription: `Re-authentication period in seconds. Will be null if hostMode is Multi-Auth`,
								Computed:            true,
							},
						},
					},
					"radius_accounting_enabled": schema.BoolAttribute{
						MarkdownDescription: `Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients`,
						Computed:            true,
					},
					"radius_accounting_servers": schema.SetNestedAttribute{
						MarkdownDescription: `List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"host": schema.StringAttribute{
									MarkdownDescription: `Public IP address of the RADIUS accounting server`,
									Computed:            true,
								},
								"organization_radius_server_id": schema.StringAttribute{
									MarkdownDescription: `Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server`,
									Computed:            true,
								},
								"port": schema.Int64Attribute{
									MarkdownDescription: `UDP port that the RADIUS Accounting server listens on for access requests`,
									Computed:            true,
								},
								"server_id": schema.StringAttribute{
									MarkdownDescription: `Unique ID of the RADIUS accounting server`,
									Computed:            true,
								},
							},
						},
					},
					"radius_coa_support_enabled": schema.BoolAttribute{
						MarkdownDescription: `Change of authentication for RADIUS re-authentication and disconnection`,
						Computed:            true,
					},
					"radius_group_attribute": schema.StringAttribute{
						MarkdownDescription: `Acceptable values are **""** for None, or **"11"** for Group Policies ACL`,
						Computed:            true,
					},
					"radius_servers": schema.SetNestedAttribute{
						MarkdownDescription: `List of RADIUS servers to require connecting devices to authenticate against before granting network access`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"host": schema.StringAttribute{
									MarkdownDescription: `Public IP address of the RADIUS server`,
									Computed:            true,
								},
								"organization_radius_server_id": schema.StringAttribute{
									MarkdownDescription: `Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server`,
									Computed:            true,
								},
								"port": schema.Int64Attribute{
									MarkdownDescription: `UDP port that the RADIUS server listens on for access requests`,
									Computed:            true,
								},
								"server_id": schema.StringAttribute{
									MarkdownDescription: `Unique ID of the RADIUS server`,
									Computed:            true,
								},
							},
						},
					},
					"radius_testing_enabled": schema.BoolAttribute{
						MarkdownDescription: `If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers`,
						Computed:            true,
					},
					"url_redirect_walled_garden_enabled": schema.BoolAttribute{
						MarkdownDescription: `Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication`,
						Computed:            true,
					},
					"url_redirect_walled_garden_ranges": schema.ListAttribute{
						MarkdownDescription: `IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"voice_vlan_clients": schema.BoolAttribute{
						MarkdownDescription: `CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchAccessPolicies`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"access_policy_type": schema.StringAttribute{
							MarkdownDescription: `Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.`,
							Computed:            true,
						},
						"counts": schema.SingleNestedAttribute{
							MarkdownDescription: `Counts associated with the access policy`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"ports": schema.SingleNestedAttribute{
									MarkdownDescription: `Counts associated with ports`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"with_this_policy": schema.Int64Attribute{
											MarkdownDescription: `Number of ports in the network with this policy. For template networks, this is the number of template ports (not child ports) with this policy.`,
											Computed:            true,
										},
									},
								},
							},
						},
						"dot1x": schema.SingleNestedAttribute{
							MarkdownDescription: `802.1x Settings`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"control_direction": schema.StringAttribute{
									MarkdownDescription: `Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'`,
									Computed:            true,
								},
							},
						},
						"guest_port_bouncing": schema.BoolAttribute{
							MarkdownDescription: `If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers`,
							Computed:            true,
						},
						"guest_vlan_id": schema.Int64Attribute{
							MarkdownDescription: `ID for the guest VLAN allow unauthorized devices access to limited network resources`,
							Computed:            true,
						},
						"host_mode": schema.StringAttribute{
							MarkdownDescription: `Choose the Host Mode for the access policy.`,
							Computed:            true,
						},
						"increase_access_speed": schema.BoolAttribute{
							MarkdownDescription: `Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the access policy`,
							Computed:            true,
						},
						"radius": schema.SingleNestedAttribute{
							MarkdownDescription: `Object for RADIUS Settings`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"cache": schema.SingleNestedAttribute{
									MarkdownDescription: `Object for RADIUS Cache Settings`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"enabled": schema.BoolAttribute{
											MarkdownDescription: `Enable to cache authorization and authentication responses on the RADIUS server`,
											Computed:            true,
										},
										"timeout": schema.Int64Attribute{
											MarkdownDescription: `If RADIUS caching is enabled, this value dictates how long the cache will remain in the RADIUS server, in hours, to allow network access without authentication`,
											Computed:            true,
										},
									},
								},
								"critical_auth": schema.SingleNestedAttribute{
									MarkdownDescription: `Critical auth settings for when authentication is rejected by the RADIUS server`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"data_vlan_id": schema.Int64Attribute{
											MarkdownDescription: `VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth`,
											Computed:            true,
										},
										"suspend_port_bounce": schema.BoolAttribute{
											MarkdownDescription: `Enable to suspend port bounce when RADIUS servers are unreachable`,
											Computed:            true,
										},
										"voice_vlan_id": schema.Int64Attribute{
											MarkdownDescription: `VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth`,
											Computed:            true,
										},
									},
								},
								"failed_auth_vlan_id": schema.Int64Attribute{
									MarkdownDescription: `VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth`,
									Computed:            true,
								},
								"re_authentication_interval": schema.Int64Attribute{
									MarkdownDescription: `Re-authentication period in seconds. Will be null if hostMode is Multi-Auth`,
									Computed:            true,
								},
							},
						},
						"radius_accounting_enabled": schema.BoolAttribute{
							MarkdownDescription: `Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients`,
							Computed:            true,
						},
						"radius_accounting_servers": schema.SetNestedAttribute{
							MarkdownDescription: `List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"host": schema.StringAttribute{
										MarkdownDescription: `Public IP address of the RADIUS accounting server`,
										Computed:            true,
									},
									"organization_radius_server_id": schema.StringAttribute{
										MarkdownDescription: `Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server`,
										Computed:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: `UDP port that the RADIUS Accounting server listens on for access requests`,
										Computed:            true,
									},
									"server_id": schema.StringAttribute{
										MarkdownDescription: `Unique ID of the RADIUS accounting server`,
										Computed:            true,
									},
								},
							},
						},
						"radius_coa_support_enabled": schema.BoolAttribute{
							MarkdownDescription: `Change of authentication for RADIUS re-authentication and disconnection`,
							Computed:            true,
						},
						"radius_group_attribute": schema.StringAttribute{
							MarkdownDescription: `Acceptable values are **""** for None, or **"11"** for Group Policies ACL`,
							Computed:            true,
						},
						"radius_servers": schema.SetNestedAttribute{
							MarkdownDescription: `List of RADIUS servers to require connecting devices to authenticate against before granting network access`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"host": schema.StringAttribute{
										MarkdownDescription: `Public IP address of the RADIUS server`,
										Computed:            true,
									},
									"organization_radius_server_id": schema.StringAttribute{
										MarkdownDescription: `Organization wide RADIUS server ID. This value will be empty if this RADIUS server is not an organization wide RADIUS server`,
										Computed:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: `UDP port that the RADIUS server listens on for access requests`,
										Computed:            true,
									},
									"server_id": schema.StringAttribute{
										MarkdownDescription: `Unique ID of the RADIUS server`,
										Computed:            true,
									},
								},
							},
						},
						"radius_testing_enabled": schema.BoolAttribute{
							MarkdownDescription: `If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers`,
							Computed:            true,
						},
						"url_redirect_walled_garden_enabled": schema.BoolAttribute{
							MarkdownDescription: `Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication`,
							Computed:            true,
						},
						"url_redirect_walled_garden_ranges": schema.ListAttribute{
							MarkdownDescription: `IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"voice_vlan_clients": schema.BoolAttribute{
							MarkdownDescription: `CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchAccessPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchAccessPolicies NetworksSwitchAccessPolicies
	diags := req.Config.Get(ctx, &networksSwitchAccessPolicies)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksSwitchAccessPolicies.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksSwitchAccessPolicies.NetworkID.IsNull(), !networksSwitchAccessPolicies.AccessPolicyNumber.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchAccessPolicies")
		vvNetworkID := networksSwitchAccessPolicies.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchAccessPolicies(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchAccessPolicies",
				err.Error(),
			)
			return
		}

		networksSwitchAccessPolicies = ResponseSwitchGetNetworkSwitchAccessPoliciesItemsToBody(networksSwitchAccessPolicies, response1)
		diags = resp.State.Set(ctx, &networksSwitchAccessPolicies)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchAccessPolicy")
		vvNetworkID := networksSwitchAccessPolicies.NetworkID.ValueString()
		vvAccessPolicyNumber := networksSwitchAccessPolicies.AccessPolicyNumber.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Switch.GetNetworkSwitchAccessPolicy(vvNetworkID, vvAccessPolicyNumber)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchAccessPolicy",
				err.Error(),
			)
			return
		}

		networksSwitchAccessPolicies = ResponseSwitchGetNetworkSwitchAccessPolicyItemToBody(networksSwitchAccessPolicies, response2)
		diags = resp.State.Set(ctx, &networksSwitchAccessPolicies)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchAccessPolicies struct {
	NetworkID          types.String                                        `tfsdk:"network_id"`
	AccessPolicyNumber types.String                                        `tfsdk:"access_policy_number"`
	Items              *[]ResponseItemSwitchGetNetworkSwitchAccessPolicies `tfsdk:"items"`
	Item               *ResponseSwitchGetNetworkSwitchAccessPolicy         `tfsdk:"item"`
}

type ResponseItemSwitchGetNetworkSwitchAccessPolicies struct {
	AccessPolicyType               types.String                                                               `tfsdk:"access_policy_type"`
	Counts                         *ResponseItemSwitchGetNetworkSwitchAccessPoliciesCounts                    `tfsdk:"counts"`
	Dot1X                          *ResponseItemSwitchGetNetworkSwitchAccessPoliciesDot1X                     `tfsdk:"dot1x"`
	GuestPortBouncing              types.Bool                                                                 `tfsdk:"guest_port_bouncing"`
	GuestVLANID                    types.Int64                                                                `tfsdk:"guest_vlan_id"`
	HostMode                       types.String                                                               `tfsdk:"host_mode"`
	IncreaseAccessSpeed            types.Bool                                                                 `tfsdk:"increase_access_speed"`
	Name                           types.String                                                               `tfsdk:"name"`
	Radius                         *ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadius                    `tfsdk:"radius"`
	RadiusAccountingEnabled        types.Bool                                                                 `tfsdk:"radius_accounting_enabled"`
	RadiusAccountingServers        *[]ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusAccountingServers `tfsdk:"radius_accounting_servers"`
	RadiusCoaSupportEnabled        types.Bool                                                                 `tfsdk:"radius_coa_support_enabled"`
	RadiusGroupAttribute           types.String                                                               `tfsdk:"radius_group_attribute"`
	RadiusServers                  *[]ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusServers           `tfsdk:"radius_servers"`
	RadiusTestingEnabled           types.Bool                                                                 `tfsdk:"radius_testing_enabled"`
	URLRedirectWalledGardenEnabled types.Bool                                                                 `tfsdk:"url_redirect_walled_garden_enabled"`
	URLRedirectWalledGardenRanges  types.List                                                                 `tfsdk:"url_redirect_walled_garden_ranges"`
	VoiceVLANClients               types.Bool                                                                 `tfsdk:"voice_vlan_clients"`
}

type ResponseItemSwitchGetNetworkSwitchAccessPoliciesCounts struct {
	Ports *ResponseItemSwitchGetNetworkSwitchAccessPoliciesCountsPorts `tfsdk:"ports"`
}

type ResponseItemSwitchGetNetworkSwitchAccessPoliciesCountsPorts struct {
	WithThisPolicy types.Int64 `tfsdk:"with_this_policy"`
}

type ResponseItemSwitchGetNetworkSwitchAccessPoliciesDot1X struct {
	ControlDirection types.String `tfsdk:"control_direction"`
}

type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadius struct {
	Cache                    *ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCache        `tfsdk:"cache"`
	CriticalAuth             *ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCriticalAuth `tfsdk:"critical_auth"`
	FailedAuthVLANID         types.Int64                                                         `tfsdk:"failed_auth_vlan_id"`
	ReAuthenticationInterval types.Int64                                                         `tfsdk:"re_authentication_interval"`
}

type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCache struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	Timeout types.Int64 `tfsdk:"timeout"`
}

type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCriticalAuth struct {
	DataVLANID        types.Int64 `tfsdk:"data_vlan_id"`
	SuspendPortBounce types.Bool  `tfsdk:"suspend_port_bounce"`
	VoiceVLANID       types.Int64 `tfsdk:"voice_vlan_id"`
}

type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusAccountingServers struct {
	Host                       types.String `tfsdk:"host"`
	OrganizationRadiusServerID types.String `tfsdk:"organization_radius_server_id"`
	Port                       types.Int64  `tfsdk:"port"`
	ServerID                   types.String `tfsdk:"server_id"`
}

type ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusServers struct {
	Host                       types.String `tfsdk:"host"`
	OrganizationRadiusServerID types.String `tfsdk:"organization_radius_server_id"`
	Port                       types.Int64  `tfsdk:"port"`
	ServerID                   types.String `tfsdk:"server_id"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicy struct {
	AccessPolicyType               types.String                                                         `tfsdk:"access_policy_type"`
	Counts                         *ResponseSwitchGetNetworkSwitchAccessPolicyCounts                    `tfsdk:"counts"`
	Dot1X                          *ResponseSwitchGetNetworkSwitchAccessPolicyDot1X                     `tfsdk:"dot1x"`
	GuestPortBouncing              types.Bool                                                           `tfsdk:"guest_port_bouncing"`
	GuestVLANID                    types.Int64                                                          `tfsdk:"guest_vlan_id"`
	HostMode                       types.String                                                         `tfsdk:"host_mode"`
	IncreaseAccessSpeed            types.Bool                                                           `tfsdk:"increase_access_speed"`
	Name                           types.String                                                         `tfsdk:"name"`
	Radius                         *ResponseSwitchGetNetworkSwitchAccessPolicyRadius                    `tfsdk:"radius"`
	RadiusAccountingEnabled        types.Bool                                                           `tfsdk:"radius_accounting_enabled"`
	RadiusAccountingServers        *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServers `tfsdk:"radius_accounting_servers"`
	RadiusCoaSupportEnabled        types.Bool                                                           `tfsdk:"radius_coa_support_enabled"`
	RadiusGroupAttribute           types.String                                                         `tfsdk:"radius_group_attribute"`
	RadiusServers                  *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServers           `tfsdk:"radius_servers"`
	RadiusTestingEnabled           types.Bool                                                           `tfsdk:"radius_testing_enabled"`
	URLRedirectWalledGardenEnabled types.Bool                                                           `tfsdk:"url_redirect_walled_garden_enabled"`
	URLRedirectWalledGardenRanges  types.List                                                           `tfsdk:"url_redirect_walled_garden_ranges"`
	VoiceVLANClients               types.Bool                                                           `tfsdk:"voice_vlan_clients"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyCounts struct {
	Ports *ResponseSwitchGetNetworkSwitchAccessPolicyCountsPorts `tfsdk:"ports"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyCountsPorts struct {
	WithThisPolicy types.Int64 `tfsdk:"with_this_policy"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyDot1X struct {
	ControlDirection types.String `tfsdk:"control_direction"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyRadius struct {
	Cache                    *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCache        `tfsdk:"cache"`
	CriticalAuth             *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCriticalAuth `tfsdk:"critical_auth"`
	FailedAuthVLANID         types.Int64                                                   `tfsdk:"failed_auth_vlan_id"`
	ReAuthenticationInterval types.Int64                                                   `tfsdk:"re_authentication_interval"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCache struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	Timeout types.Int64 `tfsdk:"timeout"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCriticalAuth struct {
	DataVLANID        types.Int64 `tfsdk:"data_vlan_id"`
	SuspendPortBounce types.Bool  `tfsdk:"suspend_port_bounce"`
	VoiceVLANID       types.Int64 `tfsdk:"voice_vlan_id"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServers struct {
	Host                       types.String `tfsdk:"host"`
	OrganizationRadiusServerID types.String `tfsdk:"organization_radius_server_id"`
	Port                       types.Int64  `tfsdk:"port"`
	ServerID                   types.String `tfsdk:"server_id"`
}

type ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServers struct {
	Host                       types.String `tfsdk:"host"`
	OrganizationRadiusServerID types.String `tfsdk:"organization_radius_server_id"`
	Port                       types.Int64  `tfsdk:"port"`
	ServerID                   types.String `tfsdk:"server_id"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchAccessPoliciesItemsToBody(state NetworksSwitchAccessPolicies, response *merakigosdk.ResponseSwitchGetNetworkSwitchAccessPolicies) NetworksSwitchAccessPolicies {
	var items []ResponseItemSwitchGetNetworkSwitchAccessPolicies
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchAccessPolicies{
			AccessPolicyType: types.StringValue(item.AccessPolicyType),
			Counts: func() *ResponseItemSwitchGetNetworkSwitchAccessPoliciesCounts {
				if item.Counts != nil {
					return &ResponseItemSwitchGetNetworkSwitchAccessPoliciesCounts{
						Ports: func() *ResponseItemSwitchGetNetworkSwitchAccessPoliciesCountsPorts {
							if item.Counts.Ports != nil {
								return &ResponseItemSwitchGetNetworkSwitchAccessPoliciesCountsPorts{
									WithThisPolicy: func() types.Int64 {
										if item.Counts.Ports.WithThisPolicy != nil {
											return types.Int64Value(int64(*item.Counts.Ports.WithThisPolicy))
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
			Dot1X: func() *ResponseItemSwitchGetNetworkSwitchAccessPoliciesDot1X {
				if item.Dot1X != nil {
					return &ResponseItemSwitchGetNetworkSwitchAccessPoliciesDot1X{
						ControlDirection: types.StringValue(item.Dot1X.ControlDirection),
					}
				}
				return nil
			}(),
			GuestPortBouncing: func() types.Bool {
				if item.GuestPortBouncing != nil {
					return types.BoolValue(*item.GuestPortBouncing)
				}
				return types.Bool{}
			}(),
			GuestVLANID: func() types.Int64 {
				if item.GuestVLANID != nil {
					return types.Int64Value(int64(*item.GuestVLANID))
				}
				return types.Int64{}
			}(),
			HostMode: types.StringValue(item.HostMode),
			IncreaseAccessSpeed: func() types.Bool {
				if item.IncreaseAccessSpeed != nil {
					return types.BoolValue(*item.IncreaseAccessSpeed)
				}
				return types.Bool{}
			}(),
			Name: types.StringValue(item.Name),
			Radius: func() *ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadius {
				if item.Radius != nil {
					return &ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadius{
						Cache: func() *ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCache {
							if item.Radius.Cache != nil {
								return &ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCache{
									Enabled: func() types.Bool {
										if item.Radius.Cache.Enabled != nil {
											return types.BoolValue(*item.Radius.Cache.Enabled)
										}
										return types.Bool{}
									}(),
									Timeout: func() types.Int64 {
										if item.Radius.Cache.Timeout != nil {
											return types.Int64Value(int64(*item.Radius.Cache.Timeout))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						CriticalAuth: func() *ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCriticalAuth {
							if item.Radius.CriticalAuth != nil {
								return &ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusCriticalAuth{
									DataVLANID: func() types.Int64 {
										if item.Radius.CriticalAuth.DataVLANID != nil {
											return types.Int64Value(int64(*item.Radius.CriticalAuth.DataVLANID))
										}
										return types.Int64{}
									}(),
									SuspendPortBounce: func() types.Bool {
										if item.Radius.CriticalAuth.SuspendPortBounce != nil {
											return types.BoolValue(*item.Radius.CriticalAuth.SuspendPortBounce)
										}
										return types.Bool{}
									}(),
									VoiceVLANID: func() types.Int64 {
										if item.Radius.CriticalAuth.VoiceVLANID != nil {
											return types.Int64Value(int64(*item.Radius.CriticalAuth.VoiceVLANID))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						FailedAuthVLANID: func() types.Int64 {
							if item.Radius.FailedAuthVLANID != nil {
								return types.Int64Value(int64(*item.Radius.FailedAuthVLANID))
							}
							return types.Int64{}
						}(),
						ReAuthenticationInterval: func() types.Int64 {
							if item.Radius.ReAuthenticationInterval != nil {
								return types.Int64Value(int64(*item.Radius.ReAuthenticationInterval))
							}
							return types.Int64{}
						}(),
					}
				}
				return nil
			}(),
			RadiusAccountingEnabled: func() types.Bool {
				if item.RadiusAccountingEnabled != nil {
					return types.BoolValue(*item.RadiusAccountingEnabled)
				}
				return types.Bool{}
			}(),
			RadiusAccountingServers: func() *[]ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusAccountingServers {
				if item.RadiusAccountingServers != nil {
					result := make([]ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusAccountingServers, len(*item.RadiusAccountingServers))
					for i, radiusAccountingServers := range *item.RadiusAccountingServers {
						result[i] = ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusAccountingServers{
							Host:                       types.StringValue(radiusAccountingServers.Host),
							OrganizationRadiusServerID: types.StringValue(radiusAccountingServers.OrganizationRadiusServerID),
							Port: func() types.Int64 {
								if radiusAccountingServers.Port != nil {
									return types.Int64Value(int64(*radiusAccountingServers.Port))
								}
								return types.Int64{}
							}(),
							ServerID: types.StringValue(radiusAccountingServers.ServerID),
						}
					}
					return &result
				}
				return nil
			}(),
			RadiusCoaSupportEnabled: func() types.Bool {
				if item.RadiusCoaSupportEnabled != nil {
					return types.BoolValue(*item.RadiusCoaSupportEnabled)
				}
				return types.Bool{}
			}(),
			RadiusGroupAttribute: types.StringValue(item.RadiusGroupAttribute),
			RadiusServers: func() *[]ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusServers {
				if item.RadiusServers != nil {
					result := make([]ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusServers, len(*item.RadiusServers))
					for i, radiusServers := range *item.RadiusServers {
						result[i] = ResponseItemSwitchGetNetworkSwitchAccessPoliciesRadiusServers{
							Host:                       types.StringValue(radiusServers.Host),
							OrganizationRadiusServerID: types.StringValue(radiusServers.OrganizationRadiusServerID),
							Port: func() types.Int64 {
								if radiusServers.Port != nil {
									return types.Int64Value(int64(*radiusServers.Port))
								}
								return types.Int64{}
							}(),
							ServerID: types.StringValue(radiusServers.ServerID),
						}
					}
					return &result
				}
				return nil
			}(),
			RadiusTestingEnabled: func() types.Bool {
				if item.RadiusTestingEnabled != nil {
					return types.BoolValue(*item.RadiusTestingEnabled)
				}
				return types.Bool{}
			}(),
			URLRedirectWalledGardenEnabled: func() types.Bool {
				if item.URLRedirectWalledGardenEnabled != nil {
					return types.BoolValue(*item.URLRedirectWalledGardenEnabled)
				}
				return types.Bool{}
			}(),
			URLRedirectWalledGardenRanges: StringSliceToList(item.URLRedirectWalledGardenRanges),
			VoiceVLANClients: func() types.Bool {
				if item.VoiceVLANClients != nil {
					return types.BoolValue(*item.VoiceVLANClients)
				}
				return types.Bool{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSwitchGetNetworkSwitchAccessPolicyItemToBody(state NetworksSwitchAccessPolicies, response *merakigosdk.ResponseSwitchGetNetworkSwitchAccessPolicy) NetworksSwitchAccessPolicies {
	itemState := ResponseSwitchGetNetworkSwitchAccessPolicy{
		AccessPolicyType: types.StringValue(response.AccessPolicyType),
		Counts: func() *ResponseSwitchGetNetworkSwitchAccessPolicyCounts {
			if response.Counts != nil {
				return &ResponseSwitchGetNetworkSwitchAccessPolicyCounts{
					Ports: func() *ResponseSwitchGetNetworkSwitchAccessPolicyCountsPorts {
						if response.Counts.Ports != nil {
							return &ResponseSwitchGetNetworkSwitchAccessPolicyCountsPorts{
								WithThisPolicy: func() types.Int64 {
									if response.Counts.Ports.WithThisPolicy != nil {
										return types.Int64Value(int64(*response.Counts.Ports.WithThisPolicy))
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
		Dot1X: func() *ResponseSwitchGetNetworkSwitchAccessPolicyDot1X {
			if response.Dot1X != nil {
				return &ResponseSwitchGetNetworkSwitchAccessPolicyDot1X{
					ControlDirection: types.StringValue(response.Dot1X.ControlDirection),
				}
			}
			return nil
		}(),
		GuestPortBouncing: func() types.Bool {
			if response.GuestPortBouncing != nil {
				return types.BoolValue(*response.GuestPortBouncing)
			}
			return types.Bool{}
		}(),
		GuestVLANID: func() types.Int64 {
			if response.GuestVLANID != nil {
				return types.Int64Value(int64(*response.GuestVLANID))
			}
			return types.Int64{}
		}(),
		HostMode: types.StringValue(response.HostMode),
		IncreaseAccessSpeed: func() types.Bool {
			if response.IncreaseAccessSpeed != nil {
				return types.BoolValue(*response.IncreaseAccessSpeed)
			}
			return types.Bool{}
		}(),
		Name: types.StringValue(response.Name),
		Radius: func() *ResponseSwitchGetNetworkSwitchAccessPolicyRadius {
			if response.Radius != nil {
				return &ResponseSwitchGetNetworkSwitchAccessPolicyRadius{
					Cache: func() *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCache {
						if response.Radius.Cache != nil {
							return &ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCache{
								Enabled: func() types.Bool {
									if response.Radius.Cache.Enabled != nil {
										return types.BoolValue(*response.Radius.Cache.Enabled)
									}
									return types.Bool{}
								}(),
								Timeout: func() types.Int64 {
									if response.Radius.Cache.Timeout != nil {
										return types.Int64Value(int64(*response.Radius.Cache.Timeout))
									}
									return types.Int64{}
								}(),
							}
						}
						return nil
					}(),
					CriticalAuth: func() *ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCriticalAuth {
						if response.Radius.CriticalAuth != nil {
							return &ResponseSwitchGetNetworkSwitchAccessPolicyRadiusCriticalAuth{
								DataVLANID: func() types.Int64 {
									if response.Radius.CriticalAuth.DataVLANID != nil {
										return types.Int64Value(int64(*response.Radius.CriticalAuth.DataVLANID))
									}
									return types.Int64{}
								}(),
								SuspendPortBounce: func() types.Bool {
									if response.Radius.CriticalAuth.SuspendPortBounce != nil {
										return types.BoolValue(*response.Radius.CriticalAuth.SuspendPortBounce)
									}
									return types.Bool{}
								}(),
								VoiceVLANID: func() types.Int64 {
									if response.Radius.CriticalAuth.VoiceVLANID != nil {
										return types.Int64Value(int64(*response.Radius.CriticalAuth.VoiceVLANID))
									}
									return types.Int64{}
								}(),
							}
						}
						return nil
					}(),
					FailedAuthVLANID: func() types.Int64 {
						if response.Radius.FailedAuthVLANID != nil {
							return types.Int64Value(int64(*response.Radius.FailedAuthVLANID))
						}
						return types.Int64{}
					}(),
					ReAuthenticationInterval: func() types.Int64 {
						if response.Radius.ReAuthenticationInterval != nil {
							return types.Int64Value(int64(*response.Radius.ReAuthenticationInterval))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		RadiusAccountingEnabled: func() types.Bool {
			if response.RadiusAccountingEnabled != nil {
				return types.BoolValue(*response.RadiusAccountingEnabled)
			}
			return types.Bool{}
		}(),
		RadiusAccountingServers: func() *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServers {
			if response.RadiusAccountingServers != nil {
				result := make([]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServers, len(*response.RadiusAccountingServers))
				for i, radiusAccountingServers := range *response.RadiusAccountingServers {
					result[i] = ResponseSwitchGetNetworkSwitchAccessPolicyRadiusAccountingServers{
						Host:                       types.StringValue(radiusAccountingServers.Host),
						OrganizationRadiusServerID: types.StringValue(radiusAccountingServers.OrganizationRadiusServerID),
						Port: func() types.Int64 {
							if radiusAccountingServers.Port != nil {
								return types.Int64Value(int64(*radiusAccountingServers.Port))
							}
							return types.Int64{}
						}(),
						ServerID: types.StringValue(radiusAccountingServers.ServerID),
					}
				}
				return &result
			}
			return nil
		}(),
		RadiusCoaSupportEnabled: func() types.Bool {
			if response.RadiusCoaSupportEnabled != nil {
				return types.BoolValue(*response.RadiusCoaSupportEnabled)
			}
			return types.Bool{}
		}(),
		RadiusGroupAttribute: types.StringValue(response.RadiusGroupAttribute),
		RadiusServers: func() *[]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServers {
			if response.RadiusServers != nil {
				result := make([]ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServers, len(*response.RadiusServers))
				for i, radiusServers := range *response.RadiusServers {
					result[i] = ResponseSwitchGetNetworkSwitchAccessPolicyRadiusServers{
						Host:                       types.StringValue(radiusServers.Host),
						OrganizationRadiusServerID: types.StringValue(radiusServers.OrganizationRadiusServerID),
						Port: func() types.Int64 {
							if radiusServers.Port != nil {
								return types.Int64Value(int64(*radiusServers.Port))
							}
							return types.Int64{}
						}(),
						ServerID: types.StringValue(radiusServers.ServerID),
					}
				}
				return &result
			}
			return nil
		}(),
		RadiusTestingEnabled: func() types.Bool {
			if response.RadiusTestingEnabled != nil {
				return types.BoolValue(*response.RadiusTestingEnabled)
			}
			return types.Bool{}
		}(),
		URLRedirectWalledGardenEnabled: func() types.Bool {
			if response.URLRedirectWalledGardenEnabled != nil {
				return types.BoolValue(*response.URLRedirectWalledGardenEnabled)
			}
			return types.Bool{}
		}(),
		URLRedirectWalledGardenRanges: StringSliceToList(response.URLRedirectWalledGardenRanges),
		VoiceVLANClients: func() types.Bool {
			if response.VoiceVLANClients != nil {
				return types.BoolValue(*response.VoiceVLANClients)
			}
			return types.Bool{}
		}(),
	}
	state.Item = &itemState
	return state
}
