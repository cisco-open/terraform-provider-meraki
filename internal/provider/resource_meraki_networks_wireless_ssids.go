package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"logon_name": schema.StringAttribute{
								MarkdownDescription: `The logon name of the Active Directory account.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"password": schema.StringAttribute{
								MarkdownDescription: `The password to the Active Directory user account.`,
								Sensitive:           true,
								Computed:            true,
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
									MarkdownDescription: `IP address (or FQDN) of your Active Directory server.`,
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
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: `Numerical identifier that is assigned to the VLAN`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"auth_mode": schema.StringAttribute{
				MarkdownDescription: `The association control method for the SSID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"8021x-google",
						"8021x-localradius",
						"8021x-meraki",
						"8021x-nac",
						"8021x-radius",
						"ipsk-with-nac",
						"ipsk-with-radius",
						"ipsk-without-radius",
						"open",
						"open-enhanced",
						"open-with-nac",
						"open-with-radius",
						"psk",
					),
				},
			},
			"availability_tags": schema.SetAttribute{
				MarkdownDescription: `List of tags for this SSID. If availableOnAllAps is false, then the SSID is only broadcast by APs with tags matching any of the tags in this list`,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"available_on_all_aps": schema.BoolAttribute{
				MarkdownDescription: `Whether all APs broadcast the SSID or if it's restricted to APs matching any availability tags`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"band_selection": schema.StringAttribute{
				MarkdownDescription: `The client-serving radio frequencies of this SSID in the default indoor RF profile`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"5 GHz band only",
						"Dual band operation",
						"Dual band operation with Band Steering",
					),
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
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether or not DNS server rewrite is enabled. If disabled, upstream DNS will be used`,
						Computed:            true,
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
				MarkdownDescription: `The psk encryption mode for the SSID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				// Validators: []validator.String{
				// 	stringvalidator.OneOf(
				// 		"wep",
				// 		"wpa",
				// 	),
				// },
			},
			"enterprise_admin_access": schema.StringAttribute{
				MarkdownDescription: `Whether or not an SSID is accessible by 'enterprise' administrators ('access disabled' or 'access enabled')`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"access disabled",
						"access enabled",
					),
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
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"host": schema.StringAttribute{
								MarkdownDescription: `The EoGRE concentrator's IP or FQDN. This param is required when ipAssignmentMode is 'Ethernet over GRE'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"key": schema.Int64Attribute{
						MarkdownDescription: `Optional numerical identifier that will add the GRE key field to the GRE header. Used to identify an individual traffic flow within a tunnel.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"ip_assignment_mode": schema.StringAttribute{
				MarkdownDescription: `The client IP assignment mode`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Bridge mode",
						"Ethernet over GRE",
						"Layer 3 roaming",
						"Layer 3 roaming with a concentrator",
						"NAT mode",
						"VPN",
					),
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
									MarkdownDescription: `IP address (or FQDN) of your LDAP server.`,
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
			// "local_auth": schema.BoolAttribute{
			// 	MarkdownDescription: `Extended local auth flag for Enterprise NAC`,
			// 	Computed:            true,
			// 	PlanModifiers: []planmodifier.Bool{
			// 		boolplanmodifier.UseStateForUnknown(),
			// 	},
			// },
			"local_radius": schema.SingleNestedAttribute{
				MarkdownDescription: `The current setting for Local Authentication, a built-in RADIUS server on the access point. Only valid if authMode is '8021x-localradius'.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"cache_timeout": schema.Int64Attribute{
						MarkdownDescription: `The duration (in seconds) for which LDAP and OCSP lookups are cached.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"certificate_authentication": schema.SingleNestedAttribute{
						MarkdownDescription: `The current setting for certificate verification.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"client_root_ca_certificate": schema.SingleNestedAttribute{
								MarkdownDescription: `The Client CA Certificate used to sign the client certificate.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"contents": schema.StringAttribute{
										MarkdownDescription: `The contents of the Client CA Certificate. Must be in PEM or DER format.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
								},
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether or not to use EAP-TLS certificate-based authentication to validate wireless clients.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"ocsp_responder_url": schema.StringAttribute{
								MarkdownDescription: `(Optional) The URL of the OCSP responder to verify client certificate status.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"use_ldap": schema.BoolAttribute{
								MarkdownDescription: `Whether or not to verify the certificate with LDAP.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"use_ocsp": schema.BoolAttribute{
								MarkdownDescription: `Whether or not to verify the certificate with OCSP.`,
								Computed:            true,
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
				MarkdownDescription: `Whether clients connecting to this SSID must use the IP address assigned by the DHCP server`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"min_bitrate": schema.Int64Attribute{
				MarkdownDescription: `The minimum bitrate in Mbps of this SSID in the default indoor RF profile`,
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
			"named_vlans": schema.SingleNestedAttribute{
				MarkdownDescription: `Named VLAN settings.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"radius": schema.SingleNestedAttribute{
						MarkdownDescription: `RADIUS settings. This param is only valid when authMode is 'open-with-radius' and ipAssignmentMode is not 'NAT mode'.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"guest_vlan": schema.SingleNestedAttribute{
								MarkdownDescription: `Guest VLAN settings. Used to direct traffic to a guest VLAN when none of the RADIUS servers are reachable or a client receives access-reject from the RADIUS server.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Whether or not RADIUS guest named VLAN is enabled.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.UseStateForUnknown(),
										},
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `RADIUS guest VLAN name.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
								},
							},
						},
					},
					"tagging": schema.SingleNestedAttribute{
						MarkdownDescription: `VLAN tagging settings. This param is only valid when ipAssignmentMode is 'Bridge mode' or 'Layer 3 roaming'.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"by_ap_tags": schema.SetNestedAttribute{
								MarkdownDescription: `The list of AP tags and VLAN names used for named VLAN tagging. If an AP has a tag matching one in the list, then traffic on this SSID will be directed to use the VLAN name associated to the tag.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Set{
									setplanmodifier.UseStateForUnknown(),
								},
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"tags": schema.SetAttribute{
											MarkdownDescription: `List of AP tags.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Set{
												setplanmodifier.UseStateForUnknown(),
											},

											ElementType: types.StringType,
										},
										"vlan_name": schema.StringAttribute{
											MarkdownDescription: `VLAN name that will be used to tag traffic.`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
										},
									},
								},
							},
							"default_vlan_name": schema.StringAttribute{
								MarkdownDescription: `The default VLAN name used to tag traffic in the absence of a matching AP tag.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether or not traffic should be directed to use specific VLAN names.`,
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
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.Int64Attribute{
				MarkdownDescription: `Unique identifier of the SSID`,
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
						Computed:            true,
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
				MarkdownDescription: `The total download bandwidth limit in Kbps (0 represents no limit)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"per_ssid_bandwidth_limit_up": schema.Int64Attribute{
				MarkdownDescription: `The total upload bandwidth limit in Kbps (0 represents no limit)`,
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
				MarkdownDescription: `Whether or not RADIUS accounting is enabled`,
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
				MarkdownDescription: `List of RADIUS accounting 802.1X servers to be used for authentication`,
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
							MarkdownDescription: `IP address (or FQDN) to which the APs will send RADIUS accounting messages`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: `Port on the RADIUS server that is listening for accounting messages`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"open_roaming_certificate_id": schema.Int64Attribute{
							MarkdownDescription: `The ID of the Openroaming Certificate attached to radius server`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"radsec_enabled": schema.BoolAttribute{
							MarkdownDescription: `Use RADSEC (TLS over TCP) to connect to this RADIUS accounting server. Requires radiusProxyEnabled.`,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `Shared key used to authenticate messages between the APs and RADIUS server`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"radius_accounting_servers_response": schema.SetNestedAttribute{
				MarkdownDescription: `List of RADIUS accounting 802.1X servers to be used for authentication`,
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
							MarkdownDescription: `IP address (or FQDN) to which the APs will send RADIUS accounting messages`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"open_roaming_certificate_id": schema.Int64Attribute{
							MarkdownDescription: `The ID of the Openroaming Certificate attached to radius server`,
							Computed:            true,
							Default:             int64default.StaticInt64(types.Int64Null().ValueInt64()),
						},
						"port": schema.Int64Attribute{
							MarkdownDescription: `Port on the RADIUS server that is listening for accounting messages`,
							Computed:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"radsec_enabled": schema.BoolAttribute{
							MarkdownDescription: `Use RADSEC (TLS over TCP) to connect to this RADIUS accounting server. Requires radiusProxyEnabled.`,
							Computed:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"secret": schema.StringAttribute{
							MarkdownDescription: `Shared key used to authenticate messages between the APs and RADIUS server`,
							Computed:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"radius_attribute_for_group_policies": schema.StringAttribute{
				MarkdownDescription: `RADIUS attribute used to look up group policies`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Airespace-ACL-Name",
						"Aruba-User-Role",
						"Filter-Id",
						"Reply-Message",
					),
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
			// "radius_enabled": schema.BoolAttribute{
			// 	MarkdownDescription: `Whether RADIUS authentication is enabled`,
			// 	Computed:            true,
			// 	PlanModifiers: []planmodifier.Bool{
			// 		boolplanmodifier.UseStateForUnknown(),
			// 	},
			// },
			"radius_failover_policy": schema.StringAttribute{
				MarkdownDescription: `Policy which determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Allow access",
						"Deny access",
					),
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
				MarkdownDescription: `Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Round robin",
						"Strict priority order",
					),
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
							Default: int64default.StaticInt64(types.Int64Null().ValueInt64()),
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
				MarkdownDescription: `The type of splash page for the SSID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Billing",
						"Cisco ISE",
						"Click-through splash page",
						"Facebook Wi-Fi",
						"Google Apps domain",
						"Google OAuth",
						"None",
						"Password-protected with Active Directory",
						"Password-protected with LDAP",
						"Password-protected with Meraki RADIUS",
						"Password-protected with custom RADIUS",
						"SMS authentication",
						"Sponsored guest",
						"Systems Manager Sentry",
					),
				},
			},
			"splash_timeout": schema.StringAttribute{
				MarkdownDescription: `Splash page timeout`,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Computed: true,
			},
			"ssid_admin_accessible": schema.BoolAttribute{
				MarkdownDescription: `SSID Administrator access status`,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Computed: true,
			},
			"use_vlan_tagging": schema.BoolAttribute{
				MarkdownDescription: `Whether or not traffic should be directed to use specific VLANs. This param is only valid if the ipAssignmentMode is 'Bridge mode' or 'Layer 3 roaming'`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"visible": schema.BoolAttribute{
				MarkdownDescription: `Whether the SSID is advertised or hidden by the AP`,
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
				MarkdownDescription: `Allow users to access a configurable list of IP ranges prior to sign-on`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"walled_garden_ranges": schema.SetAttribute{
				MarkdownDescription: `Domain names and IP address ranges available in Walled Garden mode`,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"wpa_encryption_mode": schema.StringAttribute{
				MarkdownDescription: `The types of WPA encryption`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"WPA1 and WPA2",
						"WPA1 only",
						"WPA2 only",
						"WPA3 192-bit Security",
						"WPA3 Transition Mode",
						"WPA3 only",
					),
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
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSID(vvNetworkID, vvNumber, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
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
	//entro aqui 2
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
	number, err := strconv.Atoi(idParts[1])
	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Import Type",
			fmt.Sprintf("Expected import type integer: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), number)...)
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
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSID(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
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
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessSSIDs", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsRs struct {
	NetworkID         types.String `tfsdk:"network_id"`
	Number            types.Int64  `tfsdk:"number"`
	AdminSplashURL    types.String `tfsdk:"admin_splash_url"`
	AuthMode          types.String `tfsdk:"auth_mode"`
	AvailabilityTags  types.Set    `tfsdk:"availability_tags"`
	AvailableOnAllAps types.Bool   `tfsdk:"available_on_all_aps"`
	BandSelection     types.String `tfsdk:"band_selection"`
	Enabled           types.Bool   `tfsdk:"enabled"`
	EncryptionMode    types.String `tfsdk:"encryption_mode"`
	IPAssignmentMode  types.String `tfsdk:"ip_assignment_mode"`
	// LocalAuth                       types.Bool                                                         `tfsdk:"local_auth"`
	MandatoryDhcpEnabled            types.Bool                                                         `tfsdk:"mandatory_dhcp_enabled"`
	MinBitrate                      types.Int64                                                        `tfsdk:"min_bitrate"`
	Name                            types.String                                                       `tfsdk:"name"`
	PerClientBandwidthLimitDown     types.Int64                                                        `tfsdk:"per_client_bandwidth_limit_down"`
	PerClientBandwidthLimitUp       types.Int64                                                        `tfsdk:"per_client_bandwidth_limit_up"`
	PerSSIDBandwidthLimitDown       types.Int64                                                        `tfsdk:"per_ssid_bandwidth_limit_down"`
	PerSSIDBandwidthLimitUp         types.Int64                                                        `tfsdk:"per_ssid_bandwidth_limit_up"`
	RadiusAccountingEnabled         types.Bool                                                         `tfsdk:"radius_accounting_enabled"`
	RadiusAccountingServers         *[]ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServersRs `tfsdk:"radius_accounting_servers"`
	RadiusAccountingServersResponse *[]ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServersRs `tfsdk:"radius_accounting_servers_response"`
	RadiusAttributeForGroupPolicies types.String                                                       `tfsdk:"radius_attribute_for_group_policies"`
	// RadiusEnabled                    types.Bool                                                         `tfsdk:"radius_enabled"`
	RadiusFailoverPolicy             types.String                                                  `tfsdk:"radius_failover_policy"`
	RadiusLoadBalancingPolicy        types.String                                                  `tfsdk:"radius_load_balancing_policy"`
	RadiusServers                    *[]ResponseWirelessGetNetworkWirelessSsidRadiusServersRs      `tfsdk:"radius_servers"`
	RadiusServersResponse            *[]ResponseWirelessGetNetworkWirelessSsidRadiusServersRs      `tfsdk:"radius_servers_response"`
	SplashPage                       types.String                                                  `tfsdk:"splash_page"`
	SplashTimeout                    types.String                                                  `tfsdk:"splash_timeout"`
	SSIDAdminAccessible              types.Bool                                                    `tfsdk:"ssid_admin_accessible"`
	Visible                          types.Bool                                                    `tfsdk:"visible"`
	WalledGardenEnabled              types.Bool                                                    `tfsdk:"walled_garden_enabled"`
	WalledGardenRanges               types.Set                                                     `tfsdk:"walled_garden_ranges"`
	WpaEncryptionMode                types.String                                                  `tfsdk:"wpa_encryption_mode"`
	ActiveDirectory                  *RequestWirelessUpdateNetworkWirelessSsidActiveDirectoryRs    `tfsdk:"active_directory"`
	AdultContentFilteringEnabled     types.Bool                                                    `tfsdk:"adult_content_filtering_enabled"`
	ApTagsAndVLANIDs                 *[]RequestWirelessUpdateNetworkWirelessSsidApTagsAndVlanIdsRs `tfsdk:"ap_tags_and_vlan_ids"`
	ConcentratorNetworkID            types.String                                                  `tfsdk:"concentrator_network_id"`
	DefaultVLANID                    types.Int64                                                   `tfsdk:"default_vlan_id"`
	DisassociateClientsOnVpnFailover types.Bool                                                    `tfsdk:"disassociate_clients_on_vpn_failover"`
	DNSRewrite                       *RequestWirelessUpdateNetworkWirelessSsidDnsRewriteRs         `tfsdk:"dns_rewrite"`
	Dot11R                           *RequestWirelessUpdateNetworkWirelessSsidDot11RRs             `tfsdk:"dot11r"`
	Dot11W                           *RequestWirelessUpdateNetworkWirelessSsidDot11WRs             `tfsdk:"dot11w"`
	EnterpriseAdminAccess            types.String                                                  `tfsdk:"enterprise_admin_access"`
	Gre                              *RequestWirelessUpdateNetworkWirelessSsidGreRs                `tfsdk:"gre"`
	LanIsolationEnabled              types.Bool                                                    `tfsdk:"lan_isolation_enabled"`
	Ldap                             *RequestWirelessUpdateNetworkWirelessSsidLdapRs               `tfsdk:"ldap"`
	LocalRadius                      *RequestWirelessUpdateNetworkWirelessSsidLocalRadiusRs        `tfsdk:"local_radius"`
	NamedVLANs                       *RequestWirelessUpdateNetworkWirelessSsidNamedVlansRs         `tfsdk:"named_vlans"`
	Oauth                            *RequestWirelessUpdateNetworkWirelessSsidOauthRs              `tfsdk:"oauth"`
	Psk                              types.String                                                  `tfsdk:"psk"`
	RadiusAccountingInterimInterval  types.Int64                                                   `tfsdk:"radius_accounting_interim_interval"`
	RadiusAuthenticationNasID        types.String                                                  `tfsdk:"radius_authentication_nas_id"`
	RadiusCalledStationID            types.String                                                  `tfsdk:"radius_called_station_id"`
	RadiusCoaEnabled                 types.Bool                                                    `tfsdk:"radius_coa_enabled"`
	RadiusFallbackEnabled            types.Bool                                                    `tfsdk:"radius_fallback_enabled"`
	RadiusGuestVLANEnabled           types.Bool                                                    `tfsdk:"radius_guest_vlan_enabled"`
	RadiusGuestVLANID                types.Int64                                                   `tfsdk:"radius_guest_vlan_id"`
	RadiusOverride                   types.Bool                                                    `tfsdk:"radius_override"`
	RadiusProxyEnabled               types.Bool                                                    `tfsdk:"radius_proxy_enabled"`
	RadiusServerAttemptsLimit        types.Int64                                                   `tfsdk:"radius_server_attempts_limit"`
	RadiusServerTimeout              types.Int64                                                   `tfsdk:"radius_server_timeout"`
	RadiusTestingEnabled             types.Bool                                                    `tfsdk:"radius_testing_enabled"`
	SecondaryConcentratorNetworkID   types.String                                                  `tfsdk:"secondary_concentrator_network_id"`
	SpeedBurst                       *RequestWirelessUpdateNetworkWirelessSsidSpeedBurstRs         `tfsdk:"speed_burst"`
	SplashGuestSponsorDomains        types.Set                                                     `tfsdk:"splash_guest_sponsor_domains"`
	UseVLANTagging                   types.Bool                                                    `tfsdk:"use_vlan_tagging"`
	VLANID                           types.Int64                                                   `tfsdk:"vlan_id"`
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

type RequestWirelessUpdateNetworkWirelessSsidNamedVlansRs struct {
	Radius  *RequestWirelessUpdateNetworkWirelessSsidNamedVlansRadiusRs  `tfsdk:"radius"`
	Tagging *RequestWirelessUpdateNetworkWirelessSsidNamedVlansTaggingRs `tfsdk:"tagging"`
}

type RequestWirelessUpdateNetworkWirelessSsidNamedVlansRadiusRs struct {
	GuestVLAN *RequestWirelessUpdateNetworkWirelessSsidNamedVlansRadiusGuestVlanRs `tfsdk:"guest_vlan"`
}

type RequestWirelessUpdateNetworkWirelessSsidNamedVlansRadiusGuestVlanRs struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Name    types.String `tfsdk:"name"`
}

type RequestWirelessUpdateNetworkWirelessSsidNamedVlansTaggingRs struct {
	ByApTags        *[]RequestWirelessUpdateNetworkWirelessSsidNamedVlansTaggingByApTagsRs `tfsdk:"by_ap_tags"`
	DefaultVLANName types.String                                                           `tfsdk:"default_vlan_name"`
	Enabled         types.Bool                                                             `tfsdk:"enabled"`
}

type RequestWirelessUpdateNetworkWirelessSsidNamedVlansTaggingByApTagsRs struct {
	Tags     types.Set    `tfsdk:"tags"`
	VLANName types.String `tfsdk:"vlan_name"`
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
			//Hoola aqui
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
		//Hoola aqui
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
	var requestWirelessUpdateNetworkWirelessSSIDNamedVLANs *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDNamedVLANs
	if r.NamedVLANs != nil {
		var requestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadius *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadius
		if r.NamedVLANs.Radius != nil {
			var requestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadiusGuestVLAN *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadiusGuestVLAN
			if r.NamedVLANs.Radius.GuestVLAN != nil {
				enabled := func() *bool {
					if !r.NamedVLANs.Radius.GuestVLAN.Enabled.IsUnknown() && !r.NamedVLANs.Radius.GuestVLAN.Enabled.IsNull() {
						return r.NamedVLANs.Radius.GuestVLAN.Enabled.ValueBoolPointer()
					}
					return nil
				}()
				name := r.NamedVLANs.Radius.GuestVLAN.Name.ValueString()
				requestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadiusGuestVLAN = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadiusGuestVLAN{
					Enabled: enabled,
					Name:    name,
				}
			}
			requestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadius = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadius{
				GuestVLAN: requestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadiusGuestVLAN,
			}
		}
		var requestWirelessUpdateNetworkWirelessSSIDNamedVLANsTagging *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsTagging
		if r.NamedVLANs.Tagging != nil {
			var requestWirelessUpdateNetworkWirelessSSIDNamedVLANsTaggingByApTags []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsTaggingByApTags
			if r.NamedVLANs.Tagging.ByApTags != nil {
				for _, rItem1 := range *r.NamedVLANs.Tagging.ByApTags { //NamedVLANs.Tagging.ByApTags// name: byApTags
					var tags []string = nil
					//Hoola aqui
					rItem1.Tags.ElementsAs(ctx, &tags, false)
					vLANName := rItem1.VLANName.ValueString()
					requestWirelessUpdateNetworkWirelessSSIDNamedVLANsTaggingByApTags = append(requestWirelessUpdateNetworkWirelessSSIDNamedVLANsTaggingByApTags, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsTaggingByApTags{
						Tags:     tags,
						VLANName: vLANName,
					})
				}
			}
			defaultVLANName := r.NamedVLANs.Tagging.DefaultVLANName.ValueString()
			enabled := func() *bool {
				if !r.NamedVLANs.Tagging.Enabled.IsUnknown() && !r.NamedVLANs.Tagging.Enabled.IsNull() {
					return r.NamedVLANs.Tagging.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			requestWirelessUpdateNetworkWirelessSSIDNamedVLANsTagging = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsTagging{
				ByApTags: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDNamedVLANsTaggingByApTags {
					if len(requestWirelessUpdateNetworkWirelessSSIDNamedVLANsTaggingByApTags) > 0 {
						return &requestWirelessUpdateNetworkWirelessSSIDNamedVLANsTaggingByApTags
					}
					return nil
				}(),
				DefaultVLANName: defaultVLANName,
				Enabled:         enabled,
			}
		}
		requestWirelessUpdateNetworkWirelessSSIDNamedVLANs = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDNamedVLANs{
			Radius:  requestWirelessUpdateNetworkWirelessSSIDNamedVLANsRadius,
			Tagging: requestWirelessUpdateNetworkWirelessSSIDNamedVLANsTagging,
		}
	}
	var requestWirelessUpdateNetworkWirelessSSIDOauth *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDOauth
	if r.Oauth != nil {
		var allowedDomains []string = nil
		//Hoola aqui
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
		NamedVLANs:                       requestWirelessUpdateNetworkWirelessSSIDNamedVLANs,
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
		AdminSplashURL: types.StringValue(response.AdminSplashURL),
		AuthMode:       types.StringValue(response.AuthMode),
		// AvailabilityTags: StringSliceToSet(response.AvailabilityTags),
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
		// LocalAuth: func() types.Bool {
		// 	if response.LocalAuth != nil {
		// 		return types.BoolValue(*response.LocalAuth)
		// 	}
		// 	return types.BoolNull()
		// }(),
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
		RadiusAccountingServersResponse: func() *[]ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServersRs {
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
		RadiusFailoverPolicy:            types.StringValue(response.RadiusFailoverPolicy),
		RadiusLoadBalancingPolicy:       types.StringValue(response.RadiusLoadBalancingPolicy),
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
		// WalledGardenEnabledResponse: func() types.Bool {
		// 	if response.WalledGardenEnabled != nil {
		// 		return types.BoolValue(*response.WalledGardenEnabled)
		// 	}
		// 	return types.Bool{}
		// }(),
		WalledGardenRanges:              StringSliceToSet(response.WalledGardenRanges),
		WpaEncryptionMode:               types.StringValue(response.WpaEncryptionMode),
		RadiusAccountingServers:         state.RadiusAccountingServers,
		RadiusOverride:                  state.RadiusOverride,
		RadiusAccountingInterimInterval: state.RadiusAccountingInterimInterval,
		RadiusAuthenticationNasID:       state.RadiusAuthenticationNasID,
		RadiusCalledStationID:           state.RadiusCalledStationID,
		RadiusCoaEnabled:                state.RadiusCoaEnabled,
		RadiusFallbackEnabled:           state.RadiusFallbackEnabled,
		RadiusServerTimeout:             state.RadiusServerTimeout,
		RadiusTestingEnabled:            state.RadiusTestingEnabled,
		SpeedBurst:                      state.SpeedBurst,
		UseVLANTagging:                  state.UseVLANTagging,
		RadiusServerAttemptsLimit:       state.RadiusServerAttemptsLimit,
		RadiusProxyEnabled:              state.RadiusProxyEnabled,
		// RadiusEnabled:                   state.RadiusEnabled,
		AvailabilityTags:    state.AvailabilityTags,
		WalledGardenEnabled: state.WalledGardenEnabled,
	}
	// state.SplashGuestSponsorDomains = types.SetNull(types.StringType)
	itemState.SplashGuestSponsorDomains = state.SplashGuestSponsorDomains

	itemState.DefaultVLANID = state.DefaultVLANID
	itemState.Psk = state.Psk
	itemState.RadiusServers = state.RadiusServers
	itemState.AdultContentFilteringEnabled = state.AdultContentFilteringEnabled
	itemState.LanIsolationEnabled = state.LanIsolationEnabled
	itemState.Dot11R = state.Dot11R
	itemState.Dot11W = state.Dot11W
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsRs)
}
