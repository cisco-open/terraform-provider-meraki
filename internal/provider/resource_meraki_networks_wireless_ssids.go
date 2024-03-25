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

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsResource{}
)

func NewNetworksWirelessSSIDsResource() resource.Resource {
	return &NetworksWirelessSSIDsResource{}
}

type NetworksWirelessSSIDsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids"
}

func (r *NetworksWirelessSSIDsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"active_directory": schema.SingleNestedAttribute{
				MarkdownDescription: `The current setting for Active Directory. Only valid if splashPage is 'Password-protected with Active Directory'`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"credentials": schema.SingleNestedAttribute{
						MarkdownDescription: `(Optional) The credentials of the user account to be used by the AP to bind to your Active Directory server. The Active Directory account should have permissions on all your Active Directory servers. Only valid if the splashPage is 'Password-protected with Active Directory'.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"logon_name": schema.StringAttribute{
								MarkdownDescription: `The logon name of the Active Directory account.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"password": schema.StringAttribute{
								MarkdownDescription: `The password to the Active Directory user account.`,
								Sensitive:           true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"servers": schema.SetNestedAttribute{
						MarkdownDescription: `The Active Directory servers to be used for authentication.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"host": schema.StringAttribute{
									MarkdownDescription: `IP address of your Active Directory server.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"port": schema.Int64Attribute{
									MarkdownDescription: `(Optional) UDP port the Active Directory server listens on. By default, uses port 3268.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
				},
			},
			"admin_splash_url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"adult_content_filtering_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether or not adult content will be blocked`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"ap_tags_and_vlan_ids": schema.SetNestedAttribute{
				MarkdownDescription: `The list of tags and VLAN IDs used for VLAN tagging. This param is only valid when the ipAssignmentMode is 'Bridge mode' or 'Layer 3 roaming'`,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"tags": schema.SetAttribute{
							MarkdownDescription: `Array of AP tags`,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: `Numerical identifier that is assigned to the VLAN`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"auth_mode": schema.StringAttribute{
				MarkdownDescription: `The association control method for the SSID ('open', 'open-enhanced', 'psk', 'open-with-radius', 'open-with-nac', '8021x-meraki', '8021x-nac', '8021x-radius', '8021x-google', '8021x-localradius', 'ipsk-with-radius' or 'ipsk-without-radius')`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"availability_tags": schema.SetAttribute{
				MarkdownDescription: `Accepts a list of tags for this SSID. If availableOnAllAps is false, then the SSID will only be broadcast by APs with tags matching any of the tags in this list.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"available_on_all_aps": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether all APs should broadcast the SSID or if it should be restricted to APs matching any availability tags. Can only be false if the SSID has availability tags.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"band_selection": schema.StringAttribute{
				MarkdownDescription: `The client-serving radio frequencies of this SSID in the default indoor RF profile. ('Dual band operation', '5 GHz band only' or 'Dual band operation with Band Steering')`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"concentrator_network_id": schema.StringAttribute{
				MarkdownDescription: `The concentrator to use when the ipAssignmentMode is 'Layer 3 roaming with a concentrator' or 'VPN'.`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"default_vlan_id": schema.Int64Attribute{
				MarkdownDescription: `The default VLAN ID used for 'all other APs'. This param is only valid when the ipAssignmentMode is 'Bridge mode' or 'Layer 3 roaming'`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"disassociate_clients_on_vpn_failover": schema.BoolAttribute{
				MarkdownDescription: `Disassociate clients when 'VPN' concentrator failover occurs in order to trigger clients to re-associate and generate new DHCP requests. This param is only valid if ipAssignmentMode is 'VPN'.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"dns_rewrite": schema.SingleNestedAttribute{
				MarkdownDescription: `DNS servers rewrite settings`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"dns_custom_nameservers": schema.SetAttribute{
						MarkdownDescription: `User specified DNS servers (up to two servers)`,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether or not DNS server rewrite is enabled. If disabled, upstream DNS will be used`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"dot11r": schema.SingleNestedAttribute{
				MarkdownDescription: `The current setting for 802.11r`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"adaptive": schema.BoolAttribute{
						MarkdownDescription: `(Optional) Whether 802.11r is adaptive or not.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether 802.11r is enabled or not.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"dot11w": schema.SingleNestedAttribute{
				MarkdownDescription: `The current setting for Protected Management Frames (802.11w).`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether 802.11w is enabled or not.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"required": schema.BoolAttribute{
						MarkdownDescription: `(Optional) Whether 802.11w is required or not.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether or not the SSID is enabled`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"encryption_mode": schema.StringAttribute{
				MarkdownDescription: `The psk encryption mode for the SSID ('wep' or 'wpa'). This param is only valid if the authMode is 'psk'`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enterprise_admin_access": schema.StringAttribute{
				MarkdownDescription: `Whether or not an SSID is accessible by 'enterprise' administrators ('access disabled' or 'access enabled')`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"gre": schema.SingleNestedAttribute{
				MarkdownDescription: `Ethernet over GRE settings`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"concentrator": schema.SingleNestedAttribute{
						MarkdownDescription: `The EoGRE concentrator's settings`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"host": schema.StringAttribute{
								MarkdownDescription: `The EoGRE concentrator's IP or FQDN. This param is required when ipAssignmentMode is 'Ethernet over GRE'.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"key": schema.Int64Attribute{
						MarkdownDescription: `Optional numerical identifier that will add the GRE key field to the GRE header. Used to identify an individual traffic flow within a tunnel.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"ip_assignment_mode": schema.StringAttribute{
				MarkdownDescription: `The client IP assignment mode ('NAT mode', 'Bridge mode', 'Layer 3 roaming', 'Ethernet over GRE', 'Layer 3 roaming with a concentrator' or 'VPN')`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"lan_isolation_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether Layer 2 LAN isolation should be enabled or disabled. Only configurable when ipAssignmentMode is 'Bridge mode'.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"ldap": schema.SingleNestedAttribute{
				MarkdownDescription: `The current setting for LDAP. Only valid if splashPage is 'Password-protected with LDAP'.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"base_distinguished_name": schema.StringAttribute{
						MarkdownDescription: `The base distinguished name of users on the LDAP server.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"credentials": schema.SingleNestedAttribute{
						MarkdownDescription: `(Optional) The credentials of the user account to be used by the AP to bind to your LDAP server. The LDAP account should have permissions on all your LDAP servers.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"distinguished_name": schema.StringAttribute{
								MarkdownDescription: `The distinguished name of the LDAP user account (example: cn=user,dc=meraki,dc=com).`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"password": schema.StringAttribute{
								MarkdownDescription: `The password of the LDAP user account.`,
								Sensitive:           true,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"server_ca_certificate": schema.SingleNestedAttribute{
						MarkdownDescription: `The CA certificate used to sign the LDAP server's key.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"contents": schema.StringAttribute{
								MarkdownDescription: `The contents of the CA certificate. Must be in PEM or DER format.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"servers": schema.SetNestedAttribute{
						MarkdownDescription: `The LDAP servers to be used for authentication.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"host": schema.StringAttribute{
									MarkdownDescription: `IP address of your LDAP server.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"port": schema.Int64Attribute{
									MarkdownDescription: `UDP port the LDAP server listens on.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
				},
			},
			"local_radius": schema.SingleNestedAttribute{
				MarkdownDescription: `The current setting for Local Authentication, a built-in RADIUS server on the access point. Only valid if authMode is '8021x-localradius'.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"cache_timeout": schema.Int64Attribute{
						MarkdownDescription: `The duration (in seconds) for which LDAP and OCSP lookups are cached.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"certificate_authentication": schema.SingleNestedAttribute{
						MarkdownDescription: `The current setting for certificate verification.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"client_root_ca_certificate": schema.SingleNestedAttribute{
								MarkdownDescription: `The Client CA Certificate used to sign the client certificate.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"contents": schema.StringAttribute{
										MarkdownDescription: `The contents of the Client CA Certificate. Must be in PEM or DER format.`,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
								},
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether or not to use EAP-TLS certificate-based authentication to validate wireless clients.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"ocsp_responder_url": schema.StringAttribute{
								MarkdownDescription: `(Optional) The URL of the OCSP responder to verify client certificate status.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"use_ldap": schema.BoolAttribute{
								MarkdownDescription: `Whether or not to verify the certificate with LDAP.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"use_ocsp": schema.BoolAttribute{
								MarkdownDescription: `Whether or not to verify the certificate with OCSP.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"password_authentication": schema.SingleNestedAttribute{
						MarkdownDescription: `The current setting for password-based authentication.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether or not to use EAP-TTLS/PAP or PEAP-GTC password-based authentication via LDAP lookup.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
				},
			},
			"mandatory_dhcp_enabled": schema.BoolAttribute{
				MarkdownDescription: `If true, Mandatory DHCP will enforce that clients connecting to this SSID must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"min_bitrate": schema.Int64Attribute{
				MarkdownDescription: `The minimum bitrate in Mbps of this SSID in the default indoor RF profile. ('1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54')`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				//            Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the SSID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.Int64Attribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
				//            Differents_types: `   parameter: schema.TypeString, item: schema.TypeInt`,
			},
			"oauth": schema.SingleNestedAttribute{
				MarkdownDescription: `The OAuth settings of this SSID. Only valid if splashPage is 'Google OAuth'.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"allowed_domains": schema.SetAttribute{
						MarkdownDescription: `(Optional) The list of domains allowed access to the network.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
				},
			},
			"per_client_bandwidth_limit_down": schema.Int64Attribute{
				MarkdownDescription: `The download bandwidth limit in Kbps. (0 represents no limit.)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"per_client_bandwidth_limit_up": schema.Int64Attribute{
				MarkdownDescription: `The upload bandwidth limit in Kbps. (0 represents no limit.)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"per_ssid_bandwidth_limit_down": schema.Int64Attribute{
				MarkdownDescription: `The total download bandwidth limit in Kbps. (0 represents no limit.)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"per_ssid_bandwidth_limit_up": schema.Int64Attribute{
				MarkdownDescription: `The total upload bandwidth limit in Kbps. (0 represents no limit.)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"psk": schema.StringAttribute{
				MarkdownDescription: `The passkey for the SSID. This param is only valid if the authMode is 'psk'`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_accounting_enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether or not RADIUS accounting is enabled. This param is only valid if the authMode is 'open-with-radius', '8021x-radius' or 'ipsk-with-radius'`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_accounting_interim_interval": schema.Int64Attribute{
				MarkdownDescription: `The interval (in seconds) in which accounting information is updated and sent to the RADIUS accounting server.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"radius_accounting_servers": schema.SetNestedAttribute{
				MarkdownDescription: `The RADIUS accounting 802.1X servers to be used for authentication. This param is only valid if the authMode is 'open-with-radius', '8021x-radius' or 'ipsk-with-radius' and radiusAccountingEnabled is 'true'`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"ca_certificate": schema.StringAttribute{
							MarkdownDescription: `Certificate used for authorization for the RADSEC Server`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"host": schema.StringAttribute{
							MarkdownDescription: `IP address to which the APs will send RADIUS accounting messages`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"open_roaming_certificate_id": schema.Int64Attribute{
							Computed: true,
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: `Port on the RADIUS server that is listening for accounting messages`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"radsec_enabled": schema.BoolAttribute{
							MarkdownDescription: `Use RADSEC (TLS over TCP) to connect to this RADIUS accounting server. Requires radiusProxyEnabled.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `Shared key used to authenticate messages between the APs and RADIUS server`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"radius_attribute_for_group_policies": schema.StringAttribute{
				MarkdownDescription: `Specify the RADIUS attribute used to look up group policies ('Filter-Id', 'Reply-Message', 'Airespace-ACL-Name' or 'Aruba-User-Role'). Access points must receive this attribute in the RADIUS Access-Accept message`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_authentication_nas_id": schema.StringAttribute{
				MarkdownDescription: `The template of the NAS identifier to be used for RADIUS authentication (ex. $NODE_MAC$:$VAP_NUM$).`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_called_station_id": schema.StringAttribute{
				MarkdownDescription: `The template of the called station identifier to be used for RADIUS (ex. $NODE_MAC$:$VAP_NUM$).`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_coa_enabled": schema.BoolAttribute{
				MarkdownDescription: `If true, Meraki devices will act as a RADIUS Dynamic Authorization Server and will respond to RADIUS Change-of-Authorization and Disconnect messages sent by the RADIUS server.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_enabled": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_failover_policy": schema.StringAttribute{
				MarkdownDescription: `This policy determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable ('Deny access' or 'Allow access')`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_fallback_enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether or not higher priority RADIUS servers should be retried after 60 seconds.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_guest_vlan_enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether or not RADIUS Guest VLAN is enabled. This param is only valid if the authMode is 'open-with-radius' and addressing mode is not set to 'isolated' or 'nat' mode`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_guest_vlan_id": schema.Int64Attribute{
				MarkdownDescription: `VLAN ID of the RADIUS Guest VLAN. This param is only valid if the authMode is 'open-with-radius' and addressing mode is not set to 'isolated' or 'nat' mode`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"radius_load_balancing_policy": schema.StringAttribute{
				MarkdownDescription: `This policy determines which RADIUS server will be contacted first in an authentication attempt and the ordering of any necessary retry attempts ('Strict priority order' or 'Round robin')`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_override": schema.BoolAttribute{
				MarkdownDescription: `If true, the RADIUS response can override VLAN tag. This is not valid when ipAssignmentMode is 'NAT mode'.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_proxy_enabled": schema.BoolAttribute{
				MarkdownDescription: `If true, Meraki devices will proxy RADIUS messages through the Meraki cloud to the configured RADIUS auth and accounting servers.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"radius_server_attempts_limit": schema.Int64Attribute{
				MarkdownDescription: `The maximum number of transmit attempts after which a RADIUS server is failed over (must be between 1-5).`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"radius_server_timeout": schema.Int64Attribute{
				MarkdownDescription: `The amount of time for which a RADIUS client waits for a reply from the RADIUS server (must be between 1-10 seconds).`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"radius_servers": schema.SetNestedAttribute{
				MarkdownDescription: `The RADIUS 802.1X servers to be used for authentication. This param is only valid if the authMode is 'open-with-radius', '8021x-radius' or 'ipsk-with-radius'`,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"ca_certificate": schema.StringAttribute{
							MarkdownDescription: `Certificate used for authorization for the RADSEC Server`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"host": schema.StringAttribute{
							MarkdownDescription: `IP address of your RADIUS server`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"open_roaming_certificate_id": schema.Int64Attribute{
							MarkdownDescription: `The ID of the Openroaming Certificate attached to radius server.`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: `UDP port the RADIUS server listens on for Access-requests`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"radsec_enabled": schema.BoolAttribute{
							MarkdownDescription: `Use RADSEC (TLS over TCP) to connect to this RADIUS server. Requires radiusProxyEnabled.`,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `RADIUS client shared secret`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"radius_servers_response": schema.SetNestedAttribute{
				MarkdownDescription: `The RADIUS 802.1X servers to be used for authentication. This param is only valid if the authMode is 'open-with-radius', '8021x-radius' or 'ipsk-with-radius'`,
				Computed:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"ca_certificate": schema.StringAttribute{
							MarkdownDescription: `Certificate used for authorization for the RADSEC Server`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"host": schema.StringAttribute{
							MarkdownDescription: `IP address of your RADIUS server`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"open_roaming_certificate_id": schema.Int64Attribute{
							MarkdownDescription: `The ID of the Openroaming Certificate attached to radius server.`,
							Computed:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: `UDP port the RADIUS server listens on for Access-requests`,
							Computed:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"radsec_enabled": schema.BoolAttribute{
							MarkdownDescription: `Use RADSEC (TLS over TCP) to connect to this RADIUS server. Requires radiusProxyEnabled.`,
							Computed:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `RADIUS client shared secret`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"radius_testing_enabled": schema.BoolAttribute{
				MarkdownDescription: `If true, Meraki devices will periodically send Access-Request messages to configured RADIUS servers using identity 'meraki_8021x_test' to ensure that the RADIUS servers are reachable.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"secondary_concentrator_network_id": schema.StringAttribute{
				MarkdownDescription: `The secondary concentrator to use when the ipAssignmentMode is 'VPN'. If configured, the APs will switch to using this concentrator if the primary concentrator is unreachable. This param is optional. ('disabled' represents no secondary concentrator.)`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"speed_burst": schema.SingleNestedAttribute{
				MarkdownDescription: `The SpeedBurst setting for this SSID'`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether or not to allow users to temporarily exceed the bandwidth limit for short periods while still keeping them under the bandwidth limit over time.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"splash_guest_sponsor_domains": schema.SetAttribute{
				MarkdownDescription: `Array of valid sponsor email domains for sponsored guest splash type.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"splash_page": schema.StringAttribute{
				MarkdownDescription: `The type of splash page for the SSID ('None', 'Click-through splash page', 'Billing', 'Password-protected with Meraki RADIUS', 'Password-protected with custom RADIUS', 'Password-protected with Active Directory', 'Password-protected with LDAP', 'SMS authentication', 'Systems Manager Sentry', 'Facebook Wi-Fi', 'Google OAuth', 'Sponsored guest', 'Cisco ISE' or 'Google Apps domain'). This attribute is not supported for template children.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"splash_timeout": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ssid_admin_accessible": schema.BoolAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"use_vlan_tagging": schema.BoolAttribute{
				MarkdownDescription: `Whether or not traffic should be directed to use specific VLANs. This param is only valid if the ipAssignmentMode is 'Bridge mode' or 'Layer 3 roaming'`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"visible": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether APs should advertise or hide this SSID. APs will only broadcast this SSID if set to true`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: `The VLAN ID used for VLAN tagging. This param is only valid when the ipAssignmentMode is 'Layer 3 roaming with a concentrator' or 'VPN'`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"walled_garden_enabled": schema.BoolAttribute{
				MarkdownDescription: `Allow access to a configurable list of IP ranges, which users may access prior to sign-on.`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"walled_garden_ranges": schema.SetAttribute{
				MarkdownDescription: `Specify your walled garden by entering an array of addresses, ranges using CIDR notation, domain names, and domain wildcards (e.g. '192.168.1.1/24', '192.168.37.10/32', 'www.yahoo.com', '*.google.com']). Meraki's splash page is automatically included in your walled garden.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"wpa_encryption_mode": schema.StringAttribute{
				MarkdownDescription: `The types of WPA encryption. ('WPA1 only', 'WPA1 and WPA2', 'WPA2 only', 'WPA3 Transition Mode', 'WPA3 only' or 'WPA3 192-bit Security')`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['number']

func (r *NetworksWirelessSSIDsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsRs

	var item types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	a := int(data.Number.ValueInt64())
	vvNumber := strconv.Itoa(a)
	//Item
	responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSID(vvNetworkID, vvNumber)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDs only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDs only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSID(vvNetworkID, vvNumber, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSID",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSID",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSID(vvNetworkID, vvNumber)
	// Has only items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDs",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDs",
			err.Error(),
		)
		return
	}
	data = ResponseWirelessGetNetworkWirelessSSIDItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	a := int(data.Number.ValueInt64())
	vvNumber := strconv.Itoa(a)
	// number
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSID(vvNetworkID, vvNumber)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSID",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSID",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessSSIDItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessSSIDsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
}

func (r *NetworksWirelessSSIDsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessSSIDsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	a := int(data.Number.ValueInt64())
	vvNumber := strconv.Itoa(a)
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSID(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSID",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSID",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsRs struct {
	NetworkID                        types.String                                                       `tfsdk:"network_id"`
	Number                           types.Int64                                                        `tfsdk:"number"`
	AdminSplashURL                   types.String                                                       `tfsdk:"admin_splash_url"`
	AuthMode                         types.String                                                       `tfsdk:"auth_mode"`
	AvailabilityTags                 types.Set                                                          `tfsdk:"availability_tags"`
	AvailableOnAllAps                types.Bool                                                         `tfsdk:"available_on_all_aps"`
	BandSelection                    types.String                                                       `tfsdk:"band_selection"`
	Enabled                          types.Bool                                                         `tfsdk:"enabled"`
	EncryptionMode                   types.String                                                       `tfsdk:"encryption_mode"`
	IPAssignmentMode                 types.String                                                       `tfsdk:"ip_assignment_mode"`
	MandatoryDhcpEnabled             types.Bool                                                         `tfsdk:"mandatory_dhcp_enabled"`
	MinBitrate                       types.Int64                                                        `tfsdk:"min_bitrate"`
	Name                             types.String                                                       `tfsdk:"name"`
	PerClientBandwidthLimitDown      types.Int64                                                        `tfsdk:"per_client_bandwidth_limit_down"`
	PerClientBandwidthLimitUp        types.Int64                                                        `tfsdk:"per_client_bandwidth_limit_up"`
	PerSSIDBandwidthLimitDown        types.Int64                                                        `tfsdk:"per_ssid_bandwidth_limit_down"`
	PerSSIDBandwidthLimitUp          types.Int64                                                        `tfsdk:"per_ssid_bandwidth_limit_up"`
	RadiusAccountingEnabled          types.Bool                                                         `tfsdk:"radius_accounting_enabled"`
	RadiusAccountingServers          *[]ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServersRs `tfsdk:"radius_accounting_servers"`
	RadiusAttributeForGroupPolicies  types.String                                                       `tfsdk:"radius_attribute_for_group_policies"`
	RadiusEnabled                    types.Bool                                                         `tfsdk:"radius_enabled"`
	RadiusFailoverPolicy             types.String                                                       `tfsdk:"radius_failover_policy"`
	RadiusLoadBalancingPolicy        types.String                                                       `tfsdk:"radius_load_balancing_policy"`
	RadiusServers                    *[]ResponseWirelessGetNetworkWirelessSsidRadiusServersRs           `tfsdk:"radius_servers"`
	RadiusServersResponse            *[]ResponseWirelessGetNetworkWirelessSsidRadiusServersRs           `tfsdk:"radius_servers_response"`
	SplashPage                       types.String                                                       `tfsdk:"splash_page"`
	SplashTimeout                    types.String                                                       `tfsdk:"splash_timeout"`
	SSIDAdminAccessible              types.Bool                                                         `tfsdk:"ssid_admin_accessible"`
	Visible                          types.Bool                                                         `tfsdk:"visible"`
	WalledGardenEnabled              types.Bool                                                         `tfsdk:"walled_garden_enabled"`
	WalledGardenRanges               types.Set                                                          `tfsdk:"walled_garden_ranges"`
	WpaEncryptionMode                types.String                                                       `tfsdk:"wpa_encryption_mode"`
	ActiveDirectory                  *RequestWirelessUpdateNetworkWirelessSsidActiveDirectoryRs         `tfsdk:"active_directory"`
	AdultContentFilteringEnabled     types.Bool                                                         `tfsdk:"adult_content_filtering_enabled"`
	ApTagsAndVLANIDs                 *[]RequestWirelessUpdateNetworkWirelessSsidApTagsAndVlanIdsRs      `tfsdk:"ap_tags_and_vlan_ids"`
	ConcentratorNetworkID            types.String                                                       `tfsdk:"concentrator_network_id"`
	DefaultVLANID                    types.Int64                                                        `tfsdk:"default_vlan_id"`
	DisassociateClientsOnVpnFailover types.Bool                                                         `tfsdk:"disassociate_clients_on_vpn_failover"`
	DNSRewrite                       *RequestWirelessUpdateNetworkWirelessSsidDnsRewriteRs              `tfsdk:"dns_rewrite"`
	Dot11R                           *RequestWirelessUpdateNetworkWirelessSsidDot11RRs                  `tfsdk:"dot11r"`
	Dot11W                           *RequestWirelessUpdateNetworkWirelessSsidDot11WRs                  `tfsdk:"dot11w"`
	EnterpriseAdminAccess            types.String                                                       `tfsdk:"enterprise_admin_access"`
	Gre                              *RequestWirelessUpdateNetworkWirelessSsidGreRs                     `tfsdk:"gre"`
	LanIsolationEnabled              types.Bool                                                         `tfsdk:"lan_isolation_enabled"`
	Ldap                             *RequestWirelessUpdateNetworkWirelessSsidLdapRs                    `tfsdk:"ldap"`
	LocalRadius                      *RequestWirelessUpdateNetworkWirelessSsidLocalRadiusRs             `tfsdk:"local_radius"`
	Oauth                            *RequestWirelessUpdateNetworkWirelessSsidOauthRs                   `tfsdk:"oauth"`
	Psk                              types.String                                                       `tfsdk:"psk"`
	RadiusAccountingInterimInterval  types.Int64                                                        `tfsdk:"radius_accounting_interim_interval"`
	RadiusAuthenticationNasID        types.String                                                       `tfsdk:"radius_authentication_nas_id"`
	RadiusCalledStationID            types.String                                                       `tfsdk:"radius_called_station_id"`
	RadiusCoaEnabled                 types.Bool                                                         `tfsdk:"radius_coa_enabled"`
	RadiusFallbackEnabled            types.Bool                                                         `tfsdk:"radius_fallback_enabled"`
	RadiusGuestVLANEnabled           types.Bool                                                         `tfsdk:"radius_guest_vlan_enabled"`
	RadiusGuestVLANID                types.Int64                                                        `tfsdk:"radius_guest_vlan_id"`
	RadiusOverride                   types.Bool                                                         `tfsdk:"radius_override"`
	RadiusProxyEnabled               types.Bool                                                         `tfsdk:"radius_proxy_enabled"`
	RadiusServerAttemptsLimit        types.Int64                                                        `tfsdk:"radius_server_attempts_limit"`
	RadiusServerTimeout              types.Int64                                                        `tfsdk:"radius_server_timeout"`
	RadiusTestingEnabled             types.Bool                                                         `tfsdk:"radius_testing_enabled"`
	SecondaryConcentratorNetworkID   types.String                                                       `tfsdk:"secondary_concentrator_network_id"`
	SpeedBurst                       *RequestWirelessUpdateNetworkWirelessSsidSpeedBurstRs              `tfsdk:"speed_burst"`
	SplashGuestSponsorDomains        types.Set                                                          `tfsdk:"splash_guest_sponsor_domains"`
	UseVLANTagging                   types.Bool                                                         `tfsdk:"use_vlan_tagging"`
	VLANID                           types.Int64                                                        `tfsdk:"vlan_id"`
}

type ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServersRs struct {
	CaCertificate            types.String `tfsdk:"ca_certificate"`
	Host                     types.String `tfsdk:"host"`
	OpenRoamingCertificateID types.Int64  `tfsdk:"open_roaming_certificate_id"`
	Port                     types.Int64  `tfsdk:"port"`
	RadsecEnabled            types.Bool   `tfsdk:"radsec_enabled"`
	Secret                   types.String `tfsdk:"secret"`
}

type ResponseWirelessGetNetworkWirelessSsidRadiusServersRs struct {
	CaCertificate            types.String `tfsdk:"ca_certificate"`
	Host                     types.String `tfsdk:"host"`
	OpenRoamingCertificateID types.Int64  `tfsdk:"open_roaming_certificate_id"`
	Port                     types.Int64  `tfsdk:"port"`
	RadsecEnabled            types.Bool   `tfsdk:"radsec_enabled"`
	Secret                   types.String `tfsdk:"secret"`
}

type RequestWirelessUpdateNetworkWirelessSsidActiveDirectoryRs struct {
	Credentials *RequestWirelessUpdateNetworkWirelessSsidActiveDirectoryCredentialsRs `tfsdk:"credentials"`
	Servers     *[]RequestWirelessUpdateNetworkWirelessSsidActiveDirectoryServersRs   `tfsdk:"servers"`
}

type RequestWirelessUpdateNetworkWirelessSsidActiveDirectoryCredentialsRs struct {
	LogonName types.String `tfsdk:"logon_name"`
	Password  types.String `tfsdk:"password"`
}

type RequestWirelessUpdateNetworkWirelessSsidActiveDirectoryServersRs struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

type RequestWirelessUpdateNetworkWirelessSsidApTagsAndVlanIdsRs struct {
	Tags   types.Set   `tfsdk:"tags"`
	VLANID types.Int64 `tfsdk:"vlan_id"`
}

type RequestWirelessUpdateNetworkWirelessSsidDnsRewriteRs struct {
	DNSCustomNameservers types.Set  `tfsdk:"dns_custom_nameservers"`
	Enabled              types.Bool `tfsdk:"enabled"`
}

type RequestWirelessUpdateNetworkWirelessSsidDot11RRs struct {
	Adaptive types.Bool `tfsdk:"adaptive"`
	Enabled  types.Bool `tfsdk:"enabled"`
}

type RequestWirelessUpdateNetworkWirelessSsidDot11WRs struct {
	Enabled  types.Bool `tfsdk:"enabled"`
	Required types.Bool `tfsdk:"required"`
}

type RequestWirelessUpdateNetworkWirelessSsidGreRs struct {
	Concentrator *RequestWirelessUpdateNetworkWirelessSsidGreConcentratorRs `tfsdk:"concentrator"`
	Key          types.Int64                                                `tfsdk:"key"`
}

type RequestWirelessUpdateNetworkWirelessSsidGreConcentratorRs struct {
	Host types.String `tfsdk:"host"`
}

type RequestWirelessUpdateNetworkWirelessSsidLdapRs struct {
	BaseDistinguishedName types.String                                                       `tfsdk:"base_distinguished_name"`
	Credentials           *RequestWirelessUpdateNetworkWirelessSsidLdapCredentialsRs         `tfsdk:"credentials"`
	ServerCaCertificate   *RequestWirelessUpdateNetworkWirelessSsidLdapServerCaCertificateRs `tfsdk:"server_ca_certificate"`
	Servers               *[]RequestWirelessUpdateNetworkWirelessSsidLdapServersRs           `tfsdk:"servers"`
}

type RequestWirelessUpdateNetworkWirelessSsidLdapCredentialsRs struct {
	DistinguishedName types.String `tfsdk:"distinguished_name"`
	Password          types.String `tfsdk:"password"`
}

type RequestWirelessUpdateNetworkWirelessSsidLdapServerCaCertificateRs struct {
	Contents types.String `tfsdk:"contents"`
}

type RequestWirelessUpdateNetworkWirelessSsidLdapServersRs struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

type RequestWirelessUpdateNetworkWirelessSsidLocalRadiusRs struct {
	CacheTimeout              types.Int64                                                                     `tfsdk:"cache_timeout"`
	CertificateAuthentication *RequestWirelessUpdateNetworkWirelessSsidLocalRadiusCertificateAuthenticationRs `tfsdk:"certificate_authentication"`
	PasswordAuthentication    *RequestWirelessUpdateNetworkWirelessSsidLocalRadiusPasswordAuthenticationRs    `tfsdk:"password_authentication"`
}

type RequestWirelessUpdateNetworkWirelessSsidLocalRadiusCertificateAuthenticationRs struct {
	ClientRootCaCertificate *RequestWirelessUpdateNetworkWirelessSsidLocalRadiusCertificateAuthenticationClientRootCaCertificateRs `tfsdk:"client_root_ca_certificate"`
	Enabled                 types.Bool                                                                                             `tfsdk:"enabled"`
	OcspResponderURL        types.String                                                                                           `tfsdk:"ocsp_responder_url"`
	UseLdap                 types.Bool                                                                                             `tfsdk:"use_ldap"`
	UseOcsp                 types.Bool                                                                                             `tfsdk:"use_ocsp"`
}

type RequestWirelessUpdateNetworkWirelessSsidLocalRadiusCertificateAuthenticationClientRootCaCertificateRs struct {
	Contents types.String `tfsdk:"contents"`
}

type RequestWirelessUpdateNetworkWirelessSsidLocalRadiusPasswordAuthenticationRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type RequestWirelessUpdateNetworkWirelessSsidOauthRs struct {
	AllowedDomains types.Set `tfsdk:"allowed_domains"`
}

type RequestWirelessUpdateNetworkWirelessSsidSpeedBurstRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// FromBody
func (r *NetworksWirelessSSIDsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSID {
	emptyString := ""
	var requestWirelessUpdateNetworkWirelessSSIDActiveDirectory *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDActiveDirectory
	if r.ActiveDirectory != nil {
		var requestWirelessUpdateNetworkWirelessSSIDActiveDirectoryCredentials *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDActiveDirectoryCredentials
		if r.ActiveDirectory.Credentials != nil {
			logonName := r.ActiveDirectory.Credentials.LogonName.ValueString()
			password := r.ActiveDirectory.Credentials.Password.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDActiveDirectoryCredentials = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDActiveDirectoryCredentials{
				LogonName: logonName,
				Password:  password,
			}
		}
		var requestWirelessUpdateNetworkWirelessSSIDActiveDirectoryServers []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDActiveDirectoryServers
		if r.ActiveDirectory.Servers != nil {
			for _, rItem1 := range *r.ActiveDirectory.Servers { //ActiveDirectory.Servers// name: servers
				host := rItem1.Host.ValueString()
				port := func() *int64 {
					if !rItem1.Port.IsUnknown() && !rItem1.Port.IsNull() {
						return rItem1.Port.ValueInt64Pointer()
					}
					return nil
				}()
				requestWirelessUpdateNetworkWirelessSSIDActiveDirectoryServers = append(requestWirelessUpdateNetworkWirelessSSIDActiveDirectoryServers, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDActiveDirectoryServers{
					Host: host,
					Port: int64ToIntPointer(port),
				})
			}
		}
		requestWirelessUpdateNetworkWirelessSSIDActiveDirectory = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDActiveDirectory{
			Credentials: requestWirelessUpdateNetworkWirelessSSIDActiveDirectoryCredentials,
			Servers: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDActiveDirectoryServers {
				if len(requestWirelessUpdateNetworkWirelessSSIDActiveDirectoryServers) > 0 {
					return &requestWirelessUpdateNetworkWirelessSSIDActiveDirectoryServers
				}
				return nil
			}(),
		}
	}
	adultContentFilteringEnabled := new(bool)
	if !r.AdultContentFilteringEnabled.IsUnknown() && !r.AdultContentFilteringEnabled.IsNull() {
		*adultContentFilteringEnabled = r.AdultContentFilteringEnabled.ValueBool()
	} else {
		adultContentFilteringEnabled = nil
	}
	var requestWirelessUpdateNetworkWirelessSSIDApTagsAndVLANIDs []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDApTagsAndVLANIDs
	if r.ApTagsAndVLANIDs != nil {
		for _, rItem1 := range *r.ApTagsAndVLANIDs {
			var tags []string = nil

			rItem1.Tags.ElementsAs(ctx, &tags, false)
			vLANID := func() *int64 {
				if !rItem1.VLANID.IsUnknown() && !rItem1.VLANID.IsNull() {
					return rItem1.VLANID.ValueInt64Pointer()
				}
				return nil
			}()
			requestWirelessUpdateNetworkWirelessSSIDApTagsAndVLANIDs = append(requestWirelessUpdateNetworkWirelessSSIDApTagsAndVLANIDs, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDApTagsAndVLANIDs{
				Tags:   tags,
				VLANID: int64ToIntPointer(vLANID),
			})
		}
	}
	authMode := new(string)
	if !r.AuthMode.IsUnknown() && !r.AuthMode.IsNull() {
		*authMode = r.AuthMode.ValueString()
	} else {
		authMode = &emptyString
	}
	var availabilityTags []string = nil
	r.AvailabilityTags.ElementsAs(ctx, &availabilityTags, false)
	availableOnAllAps := new(bool)
	if !r.AvailableOnAllAps.IsUnknown() && !r.AvailableOnAllAps.IsNull() {
		*availableOnAllAps = r.AvailableOnAllAps.ValueBool()
	} else {
		availableOnAllAps = nil
	}
	bandSelection := new(string)
	if !r.BandSelection.IsUnknown() && !r.BandSelection.IsNull() {
		*bandSelection = r.BandSelection.ValueString()
	} else {
		bandSelection = &emptyString
	}
	concentratorNetworkID := new(string)
	if !r.ConcentratorNetworkID.IsUnknown() && !r.ConcentratorNetworkID.IsNull() {
		*concentratorNetworkID = r.ConcentratorNetworkID.ValueString()
	} else {
		concentratorNetworkID = &emptyString
	}
	defaultVLANID := new(int64)
	if !r.DefaultVLANID.IsUnknown() && !r.DefaultVLANID.IsNull() {
		*defaultVLANID = r.DefaultVLANID.ValueInt64()
	} else {
		defaultVLANID = nil
	}
	disassociateClientsOnVpnFailover := new(bool)
	if !r.DisassociateClientsOnVpnFailover.IsUnknown() && !r.DisassociateClientsOnVpnFailover.IsNull() {
		*disassociateClientsOnVpnFailover = r.DisassociateClientsOnVpnFailover.ValueBool()
	} else {
		disassociateClientsOnVpnFailover = nil
	}
	var requestWirelessUpdateNetworkWirelessSSIDDNSRewrite *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDDNSRewrite
	if r.DNSRewrite != nil {
		var dnsCustomNameservers []string = nil

		r.DNSRewrite.DNSCustomNameservers.ElementsAs(ctx, &dnsCustomNameservers, false)
		enabled := func() *bool {
			if !r.DNSRewrite.Enabled.IsUnknown() && !r.DNSRewrite.Enabled.IsNull() {
				return r.DNSRewrite.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestWirelessUpdateNetworkWirelessSSIDDNSRewrite = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDDNSRewrite{
			DNSCustomNameservers: dnsCustomNameservers,
			Enabled:              enabled,
		}
	}
	var requestWirelessUpdateNetworkWirelessSSIDDot11R *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDDot11R
	if r.Dot11R != nil {
		adaptive := func() *bool {
			if !r.Dot11R.Adaptive.IsUnknown() && !r.Dot11R.Adaptive.IsNull() {
				return r.Dot11R.Adaptive.ValueBoolPointer()
			}
			return nil
		}()
		enabled := func() *bool {
			if !r.Dot11R.Enabled.IsUnknown() && !r.Dot11R.Enabled.IsNull() {
				return r.Dot11R.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestWirelessUpdateNetworkWirelessSSIDDot11R = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDDot11R{
			Adaptive: adaptive,
			Enabled:  enabled,
		}
	}
	var requestWirelessUpdateNetworkWirelessSSIDDot11W *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDDot11W
	if r.Dot11W != nil {
		enabled := func() *bool {
			if !r.Dot11W.Enabled.IsUnknown() && !r.Dot11W.Enabled.IsNull() {
				return r.Dot11W.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		required := func() *bool {
			if !r.Dot11W.Required.IsUnknown() && !r.Dot11W.Required.IsNull() {
				return r.Dot11W.Required.ValueBoolPointer()
			}
			return nil
		}()
		requestWirelessUpdateNetworkWirelessSSIDDot11W = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDDot11W{
			Enabled:  enabled,
			Required: required,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	encryptionMode := new(string)
	if !r.EncryptionMode.IsUnknown() && !r.EncryptionMode.IsNull() {
		*encryptionMode = r.EncryptionMode.ValueString()
		log.Printf("EncryptionMode: %v", *encryptionMode)
		log.Printf("Condition: %v", strings.Contains(*encryptionMode, "-"))

		if strings.Contains(*encryptionMode, "-") {
			encryptionMode = &emptyString
		}
	} else {
		encryptionMode = &emptyString
	}
	enterpriseAdminAccess := new(string)
	if !r.EnterpriseAdminAccess.IsUnknown() && !r.EnterpriseAdminAccess.IsNull() {
		*enterpriseAdminAccess = r.EnterpriseAdminAccess.ValueString()
	} else {
		enterpriseAdminAccess = &emptyString
	}
	var requestWirelessUpdateNetworkWirelessSSIDGre *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDGre
	if r.Gre != nil {
		var requestWirelessUpdateNetworkWirelessSSIDGreConcentrator *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDGreConcentrator
		if r.Gre.Concentrator != nil {
			host := r.Gre.Concentrator.Host.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDGreConcentrator = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDGreConcentrator{
				Host: host,
			}
		}
		key := func() *int64 {
			if !r.Gre.Key.IsUnknown() && !r.Gre.Key.IsNull() {
				return r.Gre.Key.ValueInt64Pointer()
			}
			return nil
		}()
		requestWirelessUpdateNetworkWirelessSSIDGre = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDGre{
			Concentrator: requestWirelessUpdateNetworkWirelessSSIDGreConcentrator,
			Key:          int64ToIntPointer(key),
		}
	}
	iPAssignmentMode := new(string)
	if !r.IPAssignmentMode.IsUnknown() && !r.IPAssignmentMode.IsNull() {
		*iPAssignmentMode = r.IPAssignmentMode.ValueString()
	} else {
		iPAssignmentMode = &emptyString
	}
	lanIsolationEnabled := new(bool)
	if !r.LanIsolationEnabled.IsUnknown() && !r.LanIsolationEnabled.IsNull() {
		*lanIsolationEnabled = r.LanIsolationEnabled.ValueBool()
	} else {
		lanIsolationEnabled = nil
	}
	var requestWirelessUpdateNetworkWirelessSSIDLdap *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLdap
	if r.Ldap != nil {
		baseDistinguishedName := r.Ldap.BaseDistinguishedName.ValueString()
		var requestWirelessUpdateNetworkWirelessSSIDLdapCredentials *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLdapCredentials
		if r.Ldap.Credentials != nil {
			distinguishedName := r.Ldap.Credentials.DistinguishedName.ValueString()
			password := r.Ldap.Credentials.Password.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDLdapCredentials = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLdapCredentials{
				DistinguishedName: distinguishedName,
				Password:          password,
			}
		}
		var requestWirelessUpdateNetworkWirelessSSIDLdapServerCaCertificate *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLdapServerCaCertificate
		if r.Ldap.ServerCaCertificate != nil {
			contents := r.Ldap.ServerCaCertificate.Contents.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDLdapServerCaCertificate = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLdapServerCaCertificate{
				Contents: contents,
			}
		}
		var requestWirelessUpdateNetworkWirelessSSIDLdapServers []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLdapServers
		if r.Ldap.Servers != nil {
			for _, rItem1 := range *r.Ldap.Servers { //Ldap.Servers// name: servers
				host := rItem1.Host.ValueString()
				port := func() *int64 {
					if !rItem1.Port.IsUnknown() && !rItem1.Port.IsNull() {
						return rItem1.Port.ValueInt64Pointer()
					}
					return nil
				}()
				requestWirelessUpdateNetworkWirelessSSIDLdapServers = append(requestWirelessUpdateNetworkWirelessSSIDLdapServers, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLdapServers{
					Host: host,
					Port: int64ToIntPointer(port),
				})
			}
		}
		requestWirelessUpdateNetworkWirelessSSIDLdap = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLdap{
			BaseDistinguishedName: baseDistinguishedName,
			Credentials:           requestWirelessUpdateNetworkWirelessSSIDLdapCredentials,
			ServerCaCertificate:   requestWirelessUpdateNetworkWirelessSSIDLdapServerCaCertificate,
			Servers: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLdapServers {
				if len(requestWirelessUpdateNetworkWirelessSSIDLdapServers) > 0 {
					return &requestWirelessUpdateNetworkWirelessSSIDLdapServers
				}
				return nil
			}(),
		}
	}
	var requestWirelessUpdateNetworkWirelessSSIDLocalRadius *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLocalRadius
	if r.LocalRadius != nil {
		cacheTimeout := func() *int64 {
			if !r.LocalRadius.CacheTimeout.IsUnknown() && !r.LocalRadius.CacheTimeout.IsNull() {
				return r.LocalRadius.CacheTimeout.ValueInt64Pointer()
			}
			return nil
		}()
		var requestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthentication *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthentication
		if r.LocalRadius.CertificateAuthentication != nil {
			var requestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthenticationClientRootCaCertificate *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthenticationClientRootCaCertificate
			if r.LocalRadius.CertificateAuthentication.ClientRootCaCertificate != nil {
				contents := r.LocalRadius.CertificateAuthentication.ClientRootCaCertificate.Contents.ValueString()
				requestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthenticationClientRootCaCertificate = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthenticationClientRootCaCertificate{
					Contents: contents,
				}
			}
			enabled := func() *bool {
				if !r.LocalRadius.CertificateAuthentication.Enabled.IsUnknown() && !r.LocalRadius.CertificateAuthentication.Enabled.IsNull() {
					return r.LocalRadius.CertificateAuthentication.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			ocspResponderURL := r.LocalRadius.CertificateAuthentication.OcspResponderURL.ValueString()
			useLdap := func() *bool {
				if !r.LocalRadius.CertificateAuthentication.UseLdap.IsUnknown() && !r.LocalRadius.CertificateAuthentication.UseLdap.IsNull() {
					return r.LocalRadius.CertificateAuthentication.UseLdap.ValueBoolPointer()
				}
				return nil
			}()
			useOcsp := func() *bool {
				if !r.LocalRadius.CertificateAuthentication.UseOcsp.IsUnknown() && !r.LocalRadius.CertificateAuthentication.UseOcsp.IsNull() {
					return r.LocalRadius.CertificateAuthentication.UseOcsp.ValueBoolPointer()
				}
				return nil
			}()
			requestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthentication = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthentication{
				ClientRootCaCertificate: requestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthenticationClientRootCaCertificate,
				Enabled:                 enabled,
				OcspResponderURL:        ocspResponderURL,
				UseLdap:                 useLdap,
				UseOcsp:                 useOcsp,
			}
		}
		var requestWirelessUpdateNetworkWirelessSSIDLocalRadiusPasswordAuthentication *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLocalRadiusPasswordAuthentication
		if r.LocalRadius.PasswordAuthentication != nil {
			enabled := func() *bool {
				if !r.LocalRadius.PasswordAuthentication.Enabled.IsUnknown() && !r.LocalRadius.PasswordAuthentication.Enabled.IsNull() {
					return r.LocalRadius.PasswordAuthentication.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			requestWirelessUpdateNetworkWirelessSSIDLocalRadiusPasswordAuthentication = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLocalRadiusPasswordAuthentication{
				Enabled: enabled,
			}
		}
		requestWirelessUpdateNetworkWirelessSSIDLocalRadius = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDLocalRadius{
			CacheTimeout:              int64ToIntPointer(cacheTimeout),
			CertificateAuthentication: requestWirelessUpdateNetworkWirelessSSIDLocalRadiusCertificateAuthentication,
			PasswordAuthentication:    requestWirelessUpdateNetworkWirelessSSIDLocalRadiusPasswordAuthentication,
		}
	}
	mandatoryDhcpEnabled := new(bool)
	if !r.MandatoryDhcpEnabled.IsUnknown() && !r.MandatoryDhcpEnabled.IsNull() {
		*mandatoryDhcpEnabled = r.MandatoryDhcpEnabled.ValueBool()
	} else {
		mandatoryDhcpEnabled = nil
	}
	minBitrate := new(float64)
	if !r.MinBitrate.IsUnknown() && !r.MinBitrate.IsNull() {
		*minBitrate = float64(r.MinBitrate.ValueInt64())
	} else {
		minBitrate = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestWirelessUpdateNetworkWirelessSSIDOauth *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDOauth
	if r.Oauth != nil {
		var allowedDomains []string = nil

		r.Oauth.AllowedDomains.ElementsAs(ctx, &allowedDomains, false)
		requestWirelessUpdateNetworkWirelessSSIDOauth = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDOauth{
			AllowedDomains: allowedDomains,
		}
	}
	perClientBandwidthLimitDown := new(int64)
	if !r.PerClientBandwidthLimitDown.IsUnknown() && !r.PerClientBandwidthLimitDown.IsNull() {
		*perClientBandwidthLimitDown = r.PerClientBandwidthLimitDown.ValueInt64()
	} else {
		perClientBandwidthLimitDown = nil
	}
	perClientBandwidthLimitUp := new(int64)
	if !r.PerClientBandwidthLimitUp.IsUnknown() && !r.PerClientBandwidthLimitUp.IsNull() {
		*perClientBandwidthLimitUp = r.PerClientBandwidthLimitUp.ValueInt64()
	} else {
		perClientBandwidthLimitUp = nil
	}
	perSSIDBandwidthLimitDown := new(int64)
	if !r.PerSSIDBandwidthLimitDown.IsUnknown() && !r.PerSSIDBandwidthLimitDown.IsNull() {
		*perSSIDBandwidthLimitDown = r.PerSSIDBandwidthLimitDown.ValueInt64()
	} else {
		perSSIDBandwidthLimitDown = nil
	}
	perSSIDBandwidthLimitUp := new(int64)
	if !r.PerSSIDBandwidthLimitUp.IsUnknown() && !r.PerSSIDBandwidthLimitUp.IsNull() {
		*perSSIDBandwidthLimitUp = r.PerSSIDBandwidthLimitUp.ValueInt64()
	} else {
		perSSIDBandwidthLimitUp = nil
	}
	psk := new(string)
	if !r.Psk.IsUnknown() && !r.Psk.IsNull() {
		*psk = r.Psk.ValueString()
	} else {
		psk = &emptyString
	}
	radiusAccountingEnabled := new(bool)
	if !r.RadiusAccountingEnabled.IsUnknown() && !r.RadiusAccountingEnabled.IsNull() {
		*radiusAccountingEnabled = r.RadiusAccountingEnabled.ValueBool()
	} else {
		radiusAccountingEnabled = nil
	}
	radiusAccountingInterimInterval := new(int64)
	if !r.RadiusAccountingInterimInterval.IsUnknown() && !r.RadiusAccountingInterimInterval.IsNull() {
		*radiusAccountingInterimInterval = r.RadiusAccountingInterimInterval.ValueInt64()
	} else {
		radiusAccountingInterimInterval = nil
	}
	var requestWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers
	if r.RadiusAccountingServers != nil {
		for _, rItem1 := range *r.RadiusAccountingServers {
			caCertificate := rItem1.CaCertificate.ValueString()
			host := rItem1.Host.ValueString()
			port := func() *int64 {
				if !rItem1.Port.IsUnknown() && !rItem1.Port.IsNull() {
					return rItem1.Port.ValueInt64Pointer()
				}
				return nil
			}()
			radsecEnabled := func() *bool {
				if !rItem1.RadsecEnabled.IsUnknown() && !rItem1.RadsecEnabled.IsNull() {
					return rItem1.RadsecEnabled.ValueBoolPointer()
				}
				return nil
			}()
			secret := rItem1.Secret.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers = append(requestWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers{
				CaCertificate: caCertificate,
				Host:          host,
				Port:          int64ToIntPointer(port),
				RadsecEnabled: radsecEnabled,
				Secret:        secret,
			})
		}
	}
	radiusAttributeForGroupPolicies := new(string)
	if !r.RadiusAttributeForGroupPolicies.IsUnknown() && !r.RadiusAttributeForGroupPolicies.IsNull() {
		*radiusAttributeForGroupPolicies = r.RadiusAttributeForGroupPolicies.ValueString()
	} else {
		radiusAttributeForGroupPolicies = &emptyString
	}
	radiusAuthenticationNasID := new(string)
	if !r.RadiusAuthenticationNasID.IsUnknown() && !r.RadiusAuthenticationNasID.IsNull() {
		*radiusAuthenticationNasID = r.RadiusAuthenticationNasID.ValueString()
	} else {
		radiusAuthenticationNasID = &emptyString
	}
	radiusCalledStationID := new(string)
	if !r.RadiusCalledStationID.IsUnknown() && !r.RadiusCalledStationID.IsNull() {
		*radiusCalledStationID = r.RadiusCalledStationID.ValueString()
	} else {
		radiusCalledStationID = &emptyString
	}
	radiusCoaEnabled := new(bool)
	if !r.RadiusCoaEnabled.IsUnknown() && !r.RadiusCoaEnabled.IsNull() {
		*radiusCoaEnabled = r.RadiusCoaEnabled.ValueBool()
	} else {
		radiusCoaEnabled = nil
	}
	radiusFailoverPolicy := new(string)
	if !r.RadiusFailoverPolicy.IsUnknown() && !r.RadiusFailoverPolicy.IsNull() {
		*radiusFailoverPolicy = r.RadiusFailoverPolicy.ValueString()
	} else {
		radiusFailoverPolicy = &emptyString
	}
	radiusFallbackEnabled := new(bool)
	if !r.RadiusFallbackEnabled.IsUnknown() && !r.RadiusFallbackEnabled.IsNull() {
		*radiusFallbackEnabled = r.RadiusFallbackEnabled.ValueBool()
	} else {
		radiusFallbackEnabled = nil
	}
	radiusGuestVLANEnabled := new(bool)
	if !r.RadiusGuestVLANEnabled.IsUnknown() && !r.RadiusGuestVLANEnabled.IsNull() {
		*radiusGuestVLANEnabled = r.RadiusGuestVLANEnabled.ValueBool()
	} else {
		radiusGuestVLANEnabled = nil
	}
	radiusGuestVLANID := new(int64)
	if !r.RadiusGuestVLANID.IsUnknown() && !r.RadiusGuestVLANID.IsNull() {
		*radiusGuestVLANID = r.RadiusGuestVLANID.ValueInt64()
	} else {
		radiusGuestVLANID = nil
	}
	radiusLoadBalancingPolicy := new(string)
	if !r.RadiusLoadBalancingPolicy.IsUnknown() && !r.RadiusLoadBalancingPolicy.IsNull() {
		*radiusLoadBalancingPolicy = r.RadiusLoadBalancingPolicy.ValueString()
	} else {
		radiusLoadBalancingPolicy = &emptyString
	}
	radiusOverride := new(bool)
	if !r.RadiusOverride.IsUnknown() && !r.RadiusOverride.IsNull() {
		*radiusOverride = r.RadiusOverride.ValueBool()
	} else {
		radiusOverride = nil
	}
	radiusProxyEnabled := new(bool)
	if !r.RadiusProxyEnabled.IsUnknown() && !r.RadiusProxyEnabled.IsNull() {
		*radiusProxyEnabled = r.RadiusProxyEnabled.ValueBool()
	} else {
		radiusProxyEnabled = nil
	}
	radiusServerAttemptsLimit := new(int64)
	if !r.RadiusServerAttemptsLimit.IsUnknown() && !r.RadiusServerAttemptsLimit.IsNull() {
		*radiusServerAttemptsLimit = r.RadiusServerAttemptsLimit.ValueInt64()
	} else {
		radiusServerAttemptsLimit = nil
	}
	radiusServerTimeout := new(int64)
	if !r.RadiusServerTimeout.IsUnknown() && !r.RadiusServerTimeout.IsNull() {
		*radiusServerTimeout = r.RadiusServerTimeout.ValueInt64()
	} else {
		radiusServerTimeout = nil
	}
	var requestWirelessUpdateNetworkWirelessSSIDRadiusServers []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDRadiusServers
	if r.RadiusServers != nil {
		for _, rItem1 := range *r.RadiusServers {
			caCertificate := rItem1.CaCertificate.ValueString()
			host := rItem1.Host.ValueString()
			openRoamingCertificateID := func() *int64 {
				if !rItem1.OpenRoamingCertificateID.IsUnknown() && !rItem1.OpenRoamingCertificateID.IsNull() {
					return rItem1.OpenRoamingCertificateID.ValueInt64Pointer()
				}
				return nil
			}()
			port := func() *int64 {
				if !rItem1.Port.IsUnknown() && !rItem1.Port.IsNull() {
					return rItem1.Port.ValueInt64Pointer()
				}
				return nil
			}()
			radsecEnabled := func() *bool {
				if !rItem1.RadsecEnabled.IsUnknown() && !rItem1.RadsecEnabled.IsNull() {
					return rItem1.RadsecEnabled.ValueBoolPointer()
				}
				return nil
			}()
			secret := rItem1.Secret.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDRadiusServers = append(requestWirelessUpdateNetworkWirelessSSIDRadiusServers, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDRadiusServers{
				CaCertificate:            caCertificate,
				Host:                     host,
				OpenRoamingCertificateID: int64ToIntPointer(openRoamingCertificateID),
				Port:                     int64ToIntPointer(port),
				RadsecEnabled:            radsecEnabled,
				Secret:                   secret,
			})
		}
	}
	radiusTestingEnabled := new(bool)
	if !r.RadiusTestingEnabled.IsUnknown() && !r.RadiusTestingEnabled.IsNull() {
		*radiusTestingEnabled = r.RadiusTestingEnabled.ValueBool()
	} else {
		radiusTestingEnabled = nil
	}
	secondaryConcentratorNetworkID := new(string)
	if !r.SecondaryConcentratorNetworkID.IsUnknown() && !r.SecondaryConcentratorNetworkID.IsNull() {
		*secondaryConcentratorNetworkID = r.SecondaryConcentratorNetworkID.ValueString()
	} else {
		secondaryConcentratorNetworkID = &emptyString
	}
	var requestWirelessUpdateNetworkWirelessSSIDSpeedBurst *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSpeedBurst
	if r.SpeedBurst != nil {
		enabled := func() *bool {
			if !r.SpeedBurst.Enabled.IsUnknown() && !r.SpeedBurst.Enabled.IsNull() {
				return r.SpeedBurst.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestWirelessUpdateNetworkWirelessSSIDSpeedBurst = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSpeedBurst{
			Enabled: enabled,
		}
	}
	var splashGuestSponsorDomains []string = nil
	r.SplashGuestSponsorDomains.ElementsAs(ctx, &splashGuestSponsorDomains, false)
	splashPage := new(string)
	if !r.SplashPage.IsUnknown() && !r.SplashPage.IsNull() {
		*splashPage = r.SplashPage.ValueString()
	} else {
		splashPage = &emptyString
	}
	useVLANTagging := new(bool)
	if !r.UseVLANTagging.IsUnknown() && !r.UseVLANTagging.IsNull() {
		*useVLANTagging = r.UseVLANTagging.ValueBool()
	} else {
		useVLANTagging = nil
	}
	visible := new(bool)
	if !r.Visible.IsUnknown() && !r.Visible.IsNull() {
		*visible = r.Visible.ValueBool()
	} else {
		visible = nil
	}
	vLANID := new(int64)
	if !r.VLANID.IsUnknown() && !r.VLANID.IsNull() {
		*vLANID = r.VLANID.ValueInt64()
	} else {
		vLANID = nil
	}
	walledGardenEnabled := new(bool)
	if !r.WalledGardenEnabled.IsUnknown() && !r.WalledGardenEnabled.IsNull() {
		*walledGardenEnabled = r.WalledGardenEnabled.ValueBool()
	} else {
		walledGardenEnabled = nil
	}
	var walledGardenRanges []string = nil
	r.WalledGardenRanges.ElementsAs(ctx, &walledGardenRanges, false)
	wpaEncryptionMode := new(string)
	if !r.WpaEncryptionMode.IsUnknown() && !r.WpaEncryptionMode.IsNull() {
		*wpaEncryptionMode = r.WpaEncryptionMode.ValueString()
	} else {
		wpaEncryptionMode = &emptyString
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSID{
		ActiveDirectory:              requestWirelessUpdateNetworkWirelessSSIDActiveDirectory,
		AdultContentFilteringEnabled: adultContentFilteringEnabled,
		ApTagsAndVLANIDs: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDApTagsAndVLANIDs {
			if len(requestWirelessUpdateNetworkWirelessSSIDApTagsAndVLANIDs) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDApTagsAndVLANIDs
			}
			return nil
		}(),
		AuthMode:                         *authMode,
		AvailabilityTags:                 availabilityTags,
		AvailableOnAllAps:                availableOnAllAps,
		BandSelection:                    *bandSelection,
		ConcentratorNetworkID:            *concentratorNetworkID,
		DefaultVLANID:                    int64ToIntPointer(defaultVLANID),
		DisassociateClientsOnVpnFailover: disassociateClientsOnVpnFailover,
		DNSRewrite:                       requestWirelessUpdateNetworkWirelessSSIDDNSRewrite,
		Dot11R:                           requestWirelessUpdateNetworkWirelessSSIDDot11R,
		Dot11W:                           requestWirelessUpdateNetworkWirelessSSIDDot11W,
		Enabled:                          enabled,
		EncryptionMode:                   *encryptionMode,
		EnterpriseAdminAccess:            *enterpriseAdminAccess,
		Gre:                              requestWirelessUpdateNetworkWirelessSSIDGre,
		IPAssignmentMode:                 *iPAssignmentMode,
		LanIsolationEnabled:              lanIsolationEnabled,
		Ldap:                             requestWirelessUpdateNetworkWirelessSSIDLdap,
		LocalRadius:                      requestWirelessUpdateNetworkWirelessSSIDLocalRadius,
		MandatoryDhcpEnabled:             mandatoryDhcpEnabled,
		MinBitrate:                       minBitrate,
		Name:                             *name,
		Oauth:                            requestWirelessUpdateNetworkWirelessSSIDOauth,
		PerClientBandwidthLimitDown:      int64ToIntPointer(perClientBandwidthLimitDown),
		PerClientBandwidthLimitUp:        int64ToIntPointer(perClientBandwidthLimitUp),
		PerSSIDBandwidthLimitDown:        int64ToIntPointer(perSSIDBandwidthLimitDown),
		PerSSIDBandwidthLimitUp:          int64ToIntPointer(perSSIDBandwidthLimitUp),
		Psk:                              *psk,
		RadiusAccountingEnabled:          radiusAccountingEnabled,
		RadiusAccountingInterimInterval:  int64ToIntPointer(radiusAccountingInterimInterval),
		RadiusAccountingServers: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers {
			if len(requestWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDRadiusAccountingServers
			}
			return nil
		}(),
		RadiusAttributeForGroupPolicies: *radiusAttributeForGroupPolicies,
		RadiusAuthenticationNasID:       *radiusAuthenticationNasID,
		RadiusCalledStationID:           *radiusCalledStationID,
		RadiusCoaEnabled:                radiusCoaEnabled,
		RadiusFailoverPolicy:            *radiusFailoverPolicy,
		RadiusFallbackEnabled:           radiusFallbackEnabled,
		RadiusGuestVLANEnabled:          radiusGuestVLANEnabled,
		RadiusGuestVLANID:               int64ToIntPointer(radiusGuestVLANID),
		RadiusLoadBalancingPolicy:       *radiusLoadBalancingPolicy,
		RadiusOverride:                  radiusOverride,
		RadiusProxyEnabled:              radiusProxyEnabled,
		RadiusServerAttemptsLimit:       int64ToIntPointer(radiusServerAttemptsLimit),
		RadiusServerTimeout:             int64ToIntPointer(radiusServerTimeout),
		RadiusServers: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDRadiusServers {
			if len(requestWirelessUpdateNetworkWirelessSSIDRadiusServers) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDRadiusServers
			}
			return nil
		}(),
		RadiusTestingEnabled:           radiusTestingEnabled,
		SecondaryConcentratorNetworkID: *secondaryConcentratorNetworkID,
		SpeedBurst:                     requestWirelessUpdateNetworkWirelessSSIDSpeedBurst,
		SplashGuestSponsorDomains:      splashGuestSponsorDomains,
		SplashPage:                     *splashPage,
		UseVLANTagging:                 useVLANTagging,
		Visible:                        visible,
		VLANID:                         int64ToIntPointer(vLANID),
		WalledGardenEnabled:            walledGardenEnabled,
		WalledGardenRanges:             walledGardenRanges,
		WpaEncryptionMode:              *wpaEncryptionMode,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDItemToBodyRs(state NetworksWirelessSSIDsRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSID, is_read bool) NetworksWirelessSSIDsRs {
	itemState := NetworksWirelessSSIDsRs{
		AdminSplashURL:   types.StringValue(response.AdminSplashURL),
		AuthMode:         types.StringValue(response.AuthMode),
		AvailabilityTags: StringSliceToSet(response.AvailabilityTags),
		AvailableOnAllAps: func() types.Bool {
			if response.AvailableOnAllAps != nil {
				return types.BoolValue(*response.AvailableOnAllAps)
			}
			return types.Bool{}
		}(),
		BandSelection: types.StringValue(response.BandSelection),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		EncryptionMode:   types.StringValue(response.EncryptionMode),
		IPAssignmentMode: types.StringValue(response.IPAssignmentMode),
		MandatoryDhcpEnabled: func() types.Bool {
			if response.MandatoryDhcpEnabled != nil {
				return types.BoolValue(*response.MandatoryDhcpEnabled)
			}
			return types.Bool{}
		}(),
		MinBitrate: func() types.Int64 {
			if response.MinBitrate != nil {
				return types.Int64Value(int64(*response.MinBitrate))
			}
			return types.Int64{}
		}(),
		Name: types.StringValue(response.Name),
		Number: func() types.Int64 {
			if response.Number != nil {
				return types.Int64Value(int64(*response.Number))
			}
			return types.Int64{}
		}(),
		PerClientBandwidthLimitDown: func() types.Int64 {
			if response.PerClientBandwidthLimitDown != nil {
				return types.Int64Value(int64(*response.PerClientBandwidthLimitDown))
			}
			return types.Int64{}
		}(),
		PerClientBandwidthLimitUp: func() types.Int64 {
			if response.PerClientBandwidthLimitUp != nil {
				return types.Int64Value(int64(*response.PerClientBandwidthLimitUp))
			}
			return types.Int64{}
		}(),
		PerSSIDBandwidthLimitDown: func() types.Int64 {
			if response.PerSSIDBandwidthLimitDown != nil {
				return types.Int64Value(int64(*response.PerSSIDBandwidthLimitDown))
			}
			return types.Int64{}
		}(),
		PerSSIDBandwidthLimitUp: func() types.Int64 {
			if response.PerSSIDBandwidthLimitUp != nil {
				return types.Int64Value(int64(*response.PerSSIDBandwidthLimitUp))
			}
			return types.Int64{}
		}(),
		RadiusAccountingEnabled: func() types.Bool {
			if response.RadiusAccountingEnabled != nil {
				return types.BoolValue(*response.RadiusAccountingEnabled)
			}
			return types.Bool{}
		}(),
		RadiusAccountingServers: func() *[]ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServersRs {
			if response.RadiusAccountingServers != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServersRs, len(*response.RadiusAccountingServers))
				for i, radiusAccountingServers := range *response.RadiusAccountingServers {
					result[i] = ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServersRs{
						CaCertificate: types.StringValue(radiusAccountingServers.CaCertificate),
						Host:          types.StringValue(radiusAccountingServers.Host),
						OpenRoamingCertificateID: func() types.Int64 {
							if radiusAccountingServers.OpenRoamingCertificateID != nil {
								return types.Int64Value(int64(*radiusAccountingServers.OpenRoamingCertificateID))
							}
							return types.Int64{}
						}(),
						Port: func() types.Int64 {
							if radiusAccountingServers.Port != nil {
								return types.Int64Value(int64(*radiusAccountingServers.Port))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServersRs{}
		}(),
		RadiusAttributeForGroupPolicies: types.StringValue(response.RadiusAttributeForGroupPolicies),
		RadiusEnabled: func() types.Bool {
			if response.RadiusEnabled != nil {
				return types.BoolValue(*response.RadiusEnabled)
			}
			return types.Bool{}
		}(),
		RadiusFailoverPolicy:      types.StringValue(response.RadiusFailoverPolicy),
		RadiusLoadBalancingPolicy: types.StringValue(response.RadiusLoadBalancingPolicy),
		RadiusServersResponse: func() *[]ResponseWirelessGetNetworkWirelessSsidRadiusServersRs {
			if response.RadiusServers != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidRadiusServersRs, len(*response.RadiusServers))
				for i, radiusServers := range *response.RadiusServers {
					result[i] = ResponseWirelessGetNetworkWirelessSsidRadiusServersRs{
						CaCertificate: types.StringValue(radiusServers.CaCertificate),
						Host:          types.StringValue(radiusServers.Host),
						OpenRoamingCertificateID: func() types.Int64 {
							if radiusServers.OpenRoamingCertificateID != nil {
								return types.Int64Value(int64(*radiusServers.OpenRoamingCertificateID))
							}
							return types.Int64{}
						}(),
						Port: func() types.Int64 {
							if radiusServers.Port != nil {
								return types.Int64Value(int64(*radiusServers.Port))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidRadiusServersRs{}
		}(),
		SplashPage:    types.StringValue(response.SplashPage),
		SplashTimeout: types.StringValue(response.SplashTimeout),
		SSIDAdminAccessible: func() types.Bool {
			if response.SSIDAdminAccessible != nil {
				return types.BoolValue(*response.SSIDAdminAccessible)
			}
			return types.Bool{}
		}(),
		Visible: func() types.Bool {
			if response.Visible != nil {
				return types.BoolValue(*response.Visible)
			}
			return types.Bool{}
		}(),
		WalledGardenEnabled: func() types.Bool {
			if response.WalledGardenEnabled != nil {
				return types.BoolValue(*response.WalledGardenEnabled)
			}
			return types.Bool{}
		}(),
		WalledGardenRanges: StringSliceToSet(response.WalledGardenRanges),
		WpaEncryptionMode:  types.StringValue(response.WpaEncryptionMode),
	}
	// state.SplashGuestSponsorDomains = types.SetNull(types.StringType)
	itemState.SplashGuestSponsorDomains = state.SplashGuestSponsorDomains
	itemState.DefaultVLANID = state.DefaultVLANID
	itemState.Psk = state.Psk
	itemState.RadiusServers = state.RadiusServers
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsRs)
}
