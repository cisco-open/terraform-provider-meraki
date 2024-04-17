package provider

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksWirelessSSIDsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsDataSource{}
)

func NewNetworksWirelessSSIDsDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsDataSource{}
}

type NetworksWirelessSSIDsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids"
}

func (d *NetworksWirelessSSIDsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"admin_splash_url": schema.StringAttribute{
						MarkdownDescription: `URL for the admin splash page`,
						Computed:            true,
					},
					"auth_mode": schema.StringAttribute{
						MarkdownDescription: `The association control method for the SSID`,
						Computed:            true,
					},
					"availability_tags": schema.ListAttribute{
						MarkdownDescription: `List of tags for this SSID. If availableOnAllAps is false, then the SSID is only broadcast by APs with tags matching any of the tags in this list`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"available_on_all_aps": schema.BoolAttribute{
						MarkdownDescription: `Whether all APs broadcast the SSID or if it's restricted to APs matching any availability tags`,
						Computed:            true,
					},
					"band_selection": schema.StringAttribute{
						MarkdownDescription: `The client-serving radio frequencies of this SSID in the default indoor RF profile`,
						Computed:            true,
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether or not the SSID is enabled`,
						Computed:            true,
					},
					"encryption_mode": schema.StringAttribute{
						MarkdownDescription: `The psk encryption mode for the SSID`,
						Computed:            true,
					},
					"ip_assignment_mode": schema.StringAttribute{
						MarkdownDescription: `The client IP assignment mode`,
						Computed:            true,
					},
					"local_auth": schema.BoolAttribute{
						MarkdownDescription: `Extended local auth flag for Enterprise NAC`,
						Computed:            true,
					},
					"mandatory_dhcp_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether clients connecting to this SSID must use the IP address assigned by the DHCP server`,
						Computed:            true,
					},
					"min_bitrate": schema.Int64Attribute{
						MarkdownDescription: `The minimum bitrate in Mbps of this SSID in the default indoor RF profile`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the SSID`,
						Computed:            true,
					},
					"number": schema.Int64Attribute{
						MarkdownDescription: `Unique identifier of the SSID`,
						Computed:            true,
					},
					"per_client_bandwidth_limit_down": schema.Int64Attribute{
						MarkdownDescription: `The download bandwidth limit in Kbps. (0 represents no limit.)`,
						Computed:            true,
					},
					"per_client_bandwidth_limit_up": schema.Int64Attribute{
						MarkdownDescription: `The upload bandwidth limit in Kbps. (0 represents no limit.)`,
						Computed:            true,
					},
					"per_ssid_bandwidth_limit_down": schema.Int64Attribute{
						MarkdownDescription: `The total download bandwidth limit in Kbps (0 represents no limit)`,
						Computed:            true,
					},
					"per_ssid_bandwidth_limit_up": schema.Int64Attribute{
						MarkdownDescription: `The total upload bandwidth limit in Kbps (0 represents no limit)`,
						Computed:            true,
					},
					"radius_accounting_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether or not RADIUS accounting is enabled`,
						Computed:            true,
					},
					"radius_accounting_servers": schema.SetNestedAttribute{
						MarkdownDescription: `List of RADIUS accounting 802.1X servers to be used for authentication`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"ca_certificate": schema.StringAttribute{
									MarkdownDescription: `Certificate used for authorization for the RADSEC Server`,
									Computed:            true,
								},
								"host": schema.StringAttribute{
									MarkdownDescription: `IP address (or FQDN) to which the APs will send RADIUS accounting messages`,
									Computed:            true,
								},
								"open_roaming_certificate_id": schema.Int64Attribute{
									MarkdownDescription: `The ID of the Openroaming Certificate attached to radius server`,
									Computed:            true,
								},
								"port": schema.Int64Attribute{
									MarkdownDescription: `Port on the RADIUS server that is listening for accounting messages`,
									Computed:            true,
								},
							},
						},
					},
					"radius_attribute_for_group_policies": schema.StringAttribute{
						MarkdownDescription: `RADIUS attribute used to look up group policies`,
						Computed:            true,
					},
					"radius_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether RADIUS authentication is enabled`,
						Computed:            true,
					},
					"radius_failover_policy": schema.StringAttribute{
						MarkdownDescription: `Policy which determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable`,
						Computed:            true,
					},
					"radius_load_balancing_policy": schema.StringAttribute{
						MarkdownDescription: `Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts`,
						Computed:            true,
					},
					"radius_servers": schema.SetNestedAttribute{
						MarkdownDescription: `List of RADIUS 802.1X servers to be used for authentication`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"ca_certificate": schema.StringAttribute{
									MarkdownDescription: `Certificate used for authorization for the RADSEC Server`,
									Computed:            true,
								},
								"host": schema.StringAttribute{
									MarkdownDescription: `IP address (or FQDN) of your RADIUS server`,
									Computed:            true,
								},
								"open_roaming_certificate_id": schema.Int64Attribute{
									MarkdownDescription: `The ID of the Openroaming Certificate attached to radius server`,
									Computed:            true,
								},
								"port": schema.Int64Attribute{
									MarkdownDescription: `UDP port the RADIUS server listens on for Access-requests`,
									Computed:            true,
								},
							},
						},
					},
					"splash_page": schema.StringAttribute{
						MarkdownDescription: `The type of splash page for the SSID`,
						Computed:            true,
					},
					"splash_timeout": schema.StringAttribute{
						MarkdownDescription: `Splash page timeout`,
						Computed:            true,
					},
					"ssid_admin_accessible": schema.BoolAttribute{
						MarkdownDescription: `SSID Administrator access status`,
						Computed:            true,
					},
					"visible": schema.BoolAttribute{
						MarkdownDescription: `Whether the SSID is advertised or hidden by the AP`,
						Computed:            true,
					},
					"walled_garden_enabled": schema.BoolAttribute{
						MarkdownDescription: `Allow users to access a configurable list of IP ranges prior to sign-on`,
						Computed:            true,
					},
					"walled_garden_ranges": schema.ListAttribute{
						MarkdownDescription: `Domain names and IP address ranges available in Walled Garden mode`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"wpa_encryption_mode": schema.StringAttribute{
						MarkdownDescription: `The types of WPA encryption`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetNetworkWirelessSsids`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"admin_splash_url": schema.StringAttribute{
							MarkdownDescription: `URL for the admin splash page`,
							Computed:            true,
						},
						"auth_mode": schema.StringAttribute{
							MarkdownDescription: `The association control method for the SSID`,
							Computed:            true,
						},
						"availability_tags": schema.ListAttribute{
							MarkdownDescription: `List of tags for this SSID. If availableOnAllAps is false, then the SSID is only broadcast by APs with tags matching any of the tags in this list`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"available_on_all_aps": schema.BoolAttribute{
							MarkdownDescription: `Whether all APs broadcast the SSID or if it's restricted to APs matching any availability tags`,
							Computed:            true,
						},
						"band_selection": schema.StringAttribute{
							MarkdownDescription: `The client-serving radio frequencies of this SSID in the default indoor RF profile`,
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: `Whether or not the SSID is enabled`,
							Computed:            true,
						},
						"encryption_mode": schema.StringAttribute{
							MarkdownDescription: `The psk encryption mode for the SSID`,
							Computed:            true,
						},
						"ip_assignment_mode": schema.StringAttribute{
							MarkdownDescription: `The client IP assignment mode`,
							Computed:            true,
						},
						"local_auth": schema.BoolAttribute{
							MarkdownDescription: `Extended local auth flag for Enterprise NAC`,
							Computed:            true,
						},
						"mandatory_dhcp_enabled": schema.BoolAttribute{
							MarkdownDescription: `Whether clients connecting to this SSID must use the IP address assigned by the DHCP server`,
							Computed:            true,
						},
						"min_bitrate": schema.Int64Attribute{
							MarkdownDescription: `The minimum bitrate in Mbps of this SSID in the default indoor RF profile`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the SSID`,
							Computed:            true,
						},
						"number": schema.Int64Attribute{
							MarkdownDescription: `Unique identifier of the SSID`,
							Computed:            true,
						},
						"per_client_bandwidth_limit_down": schema.Int64Attribute{
							MarkdownDescription: `The download bandwidth limit in Kbps. (0 represents no limit.)`,
							Computed:            true,
						},
						"per_client_bandwidth_limit_up": schema.Int64Attribute{
							MarkdownDescription: `The upload bandwidth limit in Kbps. (0 represents no limit.)`,
							Computed:            true,
						},
						"per_ssid_bandwidth_limit_down": schema.Int64Attribute{
							MarkdownDescription: `The total download bandwidth limit in Kbps (0 represents no limit)`,
							Computed:            true,
						},
						"per_ssid_bandwidth_limit_up": schema.Int64Attribute{
							MarkdownDescription: `The total upload bandwidth limit in Kbps (0 represents no limit)`,
							Computed:            true,
						},
						"radius_accounting_enabled": schema.BoolAttribute{
							MarkdownDescription: `Whether or not RADIUS accounting is enabled`,
							Computed:            true,
						},
						"radius_accounting_servers": schema.SetNestedAttribute{
							MarkdownDescription: `List of RADIUS accounting 802.1X servers to be used for authentication`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"ca_certificate": schema.StringAttribute{
										MarkdownDescription: `Certificate used for authorization for the RADSEC Server`,
										Computed:            true,
									},
									"host": schema.StringAttribute{
										MarkdownDescription: `IP address (or FQDN) to which the APs will send RADIUS accounting messages`,
										Computed:            true,
									},
									"open_roaming_certificate_id": schema.Int64Attribute{
										MarkdownDescription: `The ID of the Openroaming Certificate attached to radius server`,
										Computed:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: `Port on the RADIUS server that is listening for accounting messages`,
										Computed:            true,
									},
								},
							},
						},
						"radius_attribute_for_group_policies": schema.StringAttribute{
							MarkdownDescription: `RADIUS attribute used to look up group policies`,
							Computed:            true,
						},
						"radius_enabled": schema.BoolAttribute{
							MarkdownDescription: `Whether RADIUS authentication is enabled`,
							Computed:            true,
						},
						"radius_failover_policy": schema.StringAttribute{
							MarkdownDescription: `Policy which determines how authentication requests should be handled in the event that all of the configured RADIUS servers are unreachable`,
							Computed:            true,
						},
						"radius_load_balancing_policy": schema.StringAttribute{
							MarkdownDescription: `Policy which determines which RADIUS server will be contacted first in an authentication attempt, and the ordering of any necessary retry attempts`,
							Computed:            true,
						},
						"radius_servers": schema.SetNestedAttribute{
							MarkdownDescription: `List of RADIUS 802.1X servers to be used for authentication`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"ca_certificate": schema.StringAttribute{
										MarkdownDescription: `Certificate used for authorization for the RADSEC Server`,
										Computed:            true,
									},
									"host": schema.StringAttribute{
										MarkdownDescription: `IP address (or FQDN) of your RADIUS server`,
										Computed:            true,
									},
									"open_roaming_certificate_id": schema.Int64Attribute{
										MarkdownDescription: `The ID of the Openroaming Certificate attached to radius server`,
										Computed:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: `UDP port the RADIUS server listens on for Access-requests`,
										Computed:            true,
									},
								},
							},
						},
						"splash_page": schema.StringAttribute{
							MarkdownDescription: `The type of splash page for the SSID`,
							Computed:            true,
						},
						"splash_timeout": schema.StringAttribute{
							MarkdownDescription: `Splash page timeout`,
							Computed:            true,
						},
						"ssid_admin_accessible": schema.BoolAttribute{
							MarkdownDescription: `SSID Administrator access status`,
							Computed:            true,
						},
						"visible": schema.BoolAttribute{
							MarkdownDescription: `Whether the SSID is advertised or hidden by the AP`,
							Computed:            true,
						},
						"walled_garden_enabled": schema.BoolAttribute{
							MarkdownDescription: `Allow users to access a configurable list of IP ranges prior to sign-on`,
							Computed:            true,
						},
						"walled_garden_ranges": schema.ListAttribute{
							MarkdownDescription: `Domain names and IP address ranges available in Walled Garden mode`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"wpa_encryption_mode": schema.StringAttribute{
							MarkdownDescription: `The types of WPA encryption`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessSSIDsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDs NetworksWirelessSSIDs
	diags := req.Config.Get(ctx, &networksWirelessSSIDs)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksWirelessSSIDs.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksWirelessSSIDs.NetworkID.IsNull(), !networksWirelessSSIDs.Number.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDs")
		vvNetworkID := networksWirelessSSIDs.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDs(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDs",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDs = ResponseWirelessGetNetworkWirelessSSIDsItemsToBody(networksWirelessSSIDs, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSID")
		vvNetworkID := networksWirelessSSIDs.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDs.Number.ValueString()

		response2, restyResp2, err := d.client.Wireless.GetNetworkWirelessSSID(vvNetworkID, vvNumber)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSID",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDs = ResponseWirelessGetNetworkWirelessSSIDItemToBody(networksWirelessSSIDs, response2)
		diags = resp.State.Set(ctx, &networksWirelessSSIDs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDs struct {
	NetworkID types.String                                   `tfsdk:"network_id"`
	Number    types.String                                   `tfsdk:"number"`
	Items     *[]ResponseItemWirelessGetNetworkWirelessSsids `tfsdk:"items"`
	Item      *ResponseWirelessGetNetworkWirelessSsid        `tfsdk:"item"`
}

type ResponseItemWirelessGetNetworkWirelessSsids struct {
	AdminSplashURL                  types.String                                                          `tfsdk:"admin_splash_url"`
	AuthMode                        types.String                                                          `tfsdk:"auth_mode"`
	AvailabilityTags                types.List                                                            `tfsdk:"availability_tags"`
	AvailableOnAllAps               types.Bool                                                            `tfsdk:"available_on_all_aps"`
	BandSelection                   types.String                                                          `tfsdk:"band_selection"`
	Enabled                         types.Bool                                                            `tfsdk:"enabled"`
	EncryptionMode                  types.String                                                          `tfsdk:"encryption_mode"`
	IPAssignmentMode                types.String                                                          `tfsdk:"ip_assignment_mode"`
	LocalAuth                       types.Bool                                                            `tfsdk:"local_auth"`
	MandatoryDhcpEnabled            types.Bool                                                            `tfsdk:"mandatory_dhcp_enabled"`
	MinBitrate                      types.Int64                                                           `tfsdk:"min_bitrate"`
	Name                            types.String                                                          `tfsdk:"name"`
	Number                          types.Int64                                                           `tfsdk:"number"`
	PerClientBandwidthLimitDown     types.Int64                                                           `tfsdk:"per_client_bandwidth_limit_down"`
	PerClientBandwidthLimitUp       types.Int64                                                           `tfsdk:"per_client_bandwidth_limit_up"`
	PerSSIDBandwidthLimitDown       types.Int64                                                           `tfsdk:"per_ssid_bandwidth_limit_down"`
	PerSSIDBandwidthLimitUp         types.Int64                                                           `tfsdk:"per_ssid_bandwidth_limit_up"`
	RadiusAccountingEnabled         types.Bool                                                            `tfsdk:"radius_accounting_enabled"`
	RadiusAccountingServers         *[]ResponseItemWirelessGetNetworkWirelessSsidsRadiusAccountingServers `tfsdk:"radius_accounting_servers"`
	RadiusAttributeForGroupPolicies types.String                                                          `tfsdk:"radius_attribute_for_group_policies"`
	RadiusEnabled                   types.Bool                                                            `tfsdk:"radius_enabled"`
	RadiusFailoverPolicy            types.String                                                          `tfsdk:"radius_failover_policy"`
	RadiusLoadBalancingPolicy       types.String                                                          `tfsdk:"radius_load_balancing_policy"`
	RadiusServers                   *[]ResponseItemWirelessGetNetworkWirelessSsidsRadiusServers           `tfsdk:"radius_servers"`
	SplashPage                      types.String                                                          `tfsdk:"splash_page"`
	SplashTimeout                   types.String                                                          `tfsdk:"splash_timeout"`
	SSIDAdminAccessible             types.Bool                                                            `tfsdk:"ssid_admin_accessible"`
	Visible                         types.Bool                                                            `tfsdk:"visible"`
	WalledGardenEnabled             types.Bool                                                            `tfsdk:"walled_garden_enabled"`
	WalledGardenRanges              types.List                                                            `tfsdk:"walled_garden_ranges"`
	WpaEncryptionMode               types.String                                                          `tfsdk:"wpa_encryption_mode"`
}

type ResponseItemWirelessGetNetworkWirelessSsidsRadiusAccountingServers struct {
	CaCertificate            types.String `tfsdk:"ca_certificate"`
	Host                     types.String `tfsdk:"host"`
	OpenRoamingCertificateID types.Int64  `tfsdk:"open_roaming_certificate_id"`
	Port                     types.Int64  `tfsdk:"port"`
}

type ResponseItemWirelessGetNetworkWirelessSsidsRadiusServers struct {
	CaCertificate            types.String `tfsdk:"ca_certificate"`
	Host                     types.String `tfsdk:"host"`
	OpenRoamingCertificateID types.Int64  `tfsdk:"open_roaming_certificate_id"`
	Port                     types.Int64  `tfsdk:"port"`
}

type ResponseWirelessGetNetworkWirelessSsid struct {
	AdminSplashURL                  types.String                                                     `tfsdk:"admin_splash_url"`
	AuthMode                        types.String                                                     `tfsdk:"auth_mode"`
	AvailabilityTags                types.List                                                       `tfsdk:"availability_tags"`
	AvailableOnAllAps               types.Bool                                                       `tfsdk:"available_on_all_aps"`
	BandSelection                   types.String                                                     `tfsdk:"band_selection"`
	Enabled                         types.Bool                                                       `tfsdk:"enabled"`
	EncryptionMode                  types.String                                                     `tfsdk:"encryption_mode"`
	IPAssignmentMode                types.String                                                     `tfsdk:"ip_assignment_mode"`
	LocalAuth                       types.Bool                                                       `tfsdk:"local_auth"`
	MandatoryDhcpEnabled            types.Bool                                                       `tfsdk:"mandatory_dhcp_enabled"`
	MinBitrate                      types.Int64                                                      `tfsdk:"min_bitrate"`
	Name                            types.String                                                     `tfsdk:"name"`
	Number                          types.Int64                                                      `tfsdk:"number"`
	PerClientBandwidthLimitDown     types.Int64                                                      `tfsdk:"per_client_bandwidth_limit_down"`
	PerClientBandwidthLimitUp       types.Int64                                                      `tfsdk:"per_client_bandwidth_limit_up"`
	PerSSIDBandwidthLimitDown       types.Int64                                                      `tfsdk:"per_ssid_bandwidth_limit_down"`
	PerSSIDBandwidthLimitUp         types.Int64                                                      `tfsdk:"per_ssid_bandwidth_limit_up"`
	RadiusAccountingEnabled         types.Bool                                                       `tfsdk:"radius_accounting_enabled"`
	RadiusAccountingServers         *[]ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServers `tfsdk:"radius_accounting_servers"`
	RadiusAttributeForGroupPolicies types.String                                                     `tfsdk:"radius_attribute_for_group_policies"`
	RadiusEnabled                   types.Bool                                                       `tfsdk:"radius_enabled"`
	RadiusFailoverPolicy            types.String                                                     `tfsdk:"radius_failover_policy"`
	RadiusLoadBalancingPolicy       types.String                                                     `tfsdk:"radius_load_balancing_policy"`
	RadiusServers                   *[]ResponseWirelessGetNetworkWirelessSsidRadiusServers           `tfsdk:"radius_servers"`
	SplashPage                      types.String                                                     `tfsdk:"splash_page"`
	SplashTimeout                   types.String                                                     `tfsdk:"splash_timeout"`
	SSIDAdminAccessible             types.Bool                                                       `tfsdk:"ssid_admin_accessible"`
	Visible                         types.Bool                                                       `tfsdk:"visible"`
	WalledGardenEnabled             types.Bool                                                       `tfsdk:"walled_garden_enabled"`
	WalledGardenRanges              types.List                                                       `tfsdk:"walled_garden_ranges"`
	WpaEncryptionMode               types.String                                                     `tfsdk:"wpa_encryption_mode"`
}

type ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServers struct {
	CaCertificate            types.String `tfsdk:"ca_certificate"`
	Host                     types.String `tfsdk:"host"`
	OpenRoamingCertificateID types.Int64  `tfsdk:"open_roaming_certificate_id"`
	Port                     types.Int64  `tfsdk:"port"`
}

type ResponseWirelessGetNetworkWirelessSsidRadiusServers struct {
	CaCertificate            types.String `tfsdk:"ca_certificate"`
	Host                     types.String `tfsdk:"host"`
	OpenRoamingCertificateID types.Int64  `tfsdk:"open_roaming_certificate_id"`
	Port                     types.Int64  `tfsdk:"port"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDsItemsToBody(state NetworksWirelessSSIDs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDs) NetworksWirelessSSIDs {
	var items []ResponseItemWirelessGetNetworkWirelessSsids
	for _, item := range *response {
		itemState := ResponseItemWirelessGetNetworkWirelessSsids{
			AdminSplashURL:   types.StringValue(item.AdminSplashURL),
			AuthMode:         types.StringValue(item.AuthMode),
			AvailabilityTags: StringSliceToList(item.AvailabilityTags),
			AvailableOnAllAps: func() types.Bool {
				if item.AvailableOnAllAps != nil {
					return types.BoolValue(*item.AvailableOnAllAps)
				}
				return types.Bool{}
			}(),
			BandSelection: types.StringValue(item.BandSelection),
			Enabled: func() types.Bool {
				if item.Enabled != nil {
					return types.BoolValue(*item.Enabled)
				}
				return types.Bool{}
			}(),
			EncryptionMode:   types.StringValue(item.EncryptionMode),
			IPAssignmentMode: types.StringValue(item.IPAssignmentMode),
			LocalAuth: func() types.Bool {
				if item.LocalAuth != nil {
					return types.BoolValue(*item.LocalAuth)
				}
				return types.Bool{}
			}(),
			MandatoryDhcpEnabled: func() types.Bool {
				if item.MandatoryDhcpEnabled != nil {
					return types.BoolValue(*item.MandatoryDhcpEnabled)
				}
				return types.Bool{}
			}(),
			MinBitrate: func() types.Int64 {
				if item.MinBitrate != nil {
					return types.Int64Value(int64(*item.MinBitrate))
				}
				return types.Int64{}
			}(),
			Name: types.StringValue(item.Name),
			Number: func() types.Int64 {
				if item.Number != nil {
					return types.Int64Value(int64(*item.Number))
				}
				return types.Int64{}
			}(),
			PerClientBandwidthLimitDown: func() types.Int64 {
				if item.PerClientBandwidthLimitDown != nil {
					return types.Int64Value(int64(*item.PerClientBandwidthLimitDown))
				}
				return types.Int64{}
			}(),
			PerClientBandwidthLimitUp: func() types.Int64 {
				if item.PerClientBandwidthLimitUp != nil {
					return types.Int64Value(int64(*item.PerClientBandwidthLimitUp))
				}
				return types.Int64{}
			}(),
			PerSSIDBandwidthLimitDown: func() types.Int64 {
				if item.PerSSIDBandwidthLimitDown != nil {
					return types.Int64Value(int64(*item.PerSSIDBandwidthLimitDown))
				}
				return types.Int64{}
			}(),
			PerSSIDBandwidthLimitUp: func() types.Int64 {
				if item.PerSSIDBandwidthLimitUp != nil {
					return types.Int64Value(int64(*item.PerSSIDBandwidthLimitUp))
				}
				return types.Int64{}
			}(),
			RadiusAccountingEnabled: func() types.Bool {
				if item.RadiusAccountingEnabled != nil {
					return types.BoolValue(*item.RadiusAccountingEnabled)
				}
				return types.Bool{}
			}(),
			RadiusAccountingServers: func() *[]ResponseItemWirelessGetNetworkWirelessSsidsRadiusAccountingServers {
				if item.RadiusAccountingServers != nil {
					result := make([]ResponseItemWirelessGetNetworkWirelessSsidsRadiusAccountingServers, len(*item.RadiusAccountingServers))
					for i, radiusAccountingServers := range *item.RadiusAccountingServers {
						result[i] = ResponseItemWirelessGetNetworkWirelessSsidsRadiusAccountingServers{
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
				return &[]ResponseItemWirelessGetNetworkWirelessSsidsRadiusAccountingServers{}
			}(),
			RadiusAttributeForGroupPolicies: types.StringValue(item.RadiusAttributeForGroupPolicies),
			RadiusEnabled: func() types.Bool {
				if item.RadiusEnabled != nil {
					return types.BoolValue(*item.RadiusEnabled)
				}
				return types.Bool{}
			}(),
			RadiusFailoverPolicy:      types.StringValue(item.RadiusFailoverPolicy),
			RadiusLoadBalancingPolicy: types.StringValue(item.RadiusLoadBalancingPolicy),
			RadiusServers: func() *[]ResponseItemWirelessGetNetworkWirelessSsidsRadiusServers {
				if item.RadiusServers != nil {
					result := make([]ResponseItemWirelessGetNetworkWirelessSsidsRadiusServers, len(*item.RadiusServers))
					for i, radiusServers := range *item.RadiusServers {
						result[i] = ResponseItemWirelessGetNetworkWirelessSsidsRadiusServers{
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
				return &[]ResponseItemWirelessGetNetworkWirelessSsidsRadiusServers{}
			}(),
			SplashPage:    types.StringValue(item.SplashPage),
			SplashTimeout: types.StringValue(item.SplashTimeout),
			SSIDAdminAccessible: func() types.Bool {
				if item.SSIDAdminAccessible != nil {
					return types.BoolValue(*item.SSIDAdminAccessible)
				}
				return types.Bool{}
			}(),
			Visible: func() types.Bool {
				if item.Visible != nil {
					return types.BoolValue(*item.Visible)
				}
				return types.Bool{}
			}(),
			WalledGardenEnabled: func() types.Bool {
				if item.WalledGardenEnabled != nil {
					return types.BoolValue(*item.WalledGardenEnabled)
				}
				return types.Bool{}
			}(),
			WalledGardenRanges: StringSliceToList(item.WalledGardenRanges),
			WpaEncryptionMode:  types.StringValue(item.WpaEncryptionMode),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseWirelessGetNetworkWirelessSSIDItemToBody(state NetworksWirelessSSIDs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSID) NetworksWirelessSSIDs {
	itemState := ResponseWirelessGetNetworkWirelessSsid{
		AdminSplashURL:   types.StringValue(response.AdminSplashURL),
		AuthMode:         types.StringValue(response.AuthMode),
		AvailabilityTags: StringSliceToList(response.AvailabilityTags),
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
		LocalAuth: func() types.Bool {
			if response.LocalAuth != nil {
				return types.BoolValue(*response.LocalAuth)
			}
			return types.Bool{}
		}(),
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
		RadiusAccountingServers: func() *[]ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServers {
			if response.RadiusAccountingServers != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServers, len(*response.RadiusAccountingServers))
				for i, radiusAccountingServers := range *response.RadiusAccountingServers {
					result[i] = ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServers{
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
			return &[]ResponseWirelessGetNetworkWirelessSsidRadiusAccountingServers{}
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
		RadiusServers: func() *[]ResponseWirelessGetNetworkWirelessSsidRadiusServers {
			if response.RadiusServers != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidRadiusServers, len(*response.RadiusServers))
				for i, radiusServers := range *response.RadiusServers {
					result[i] = ResponseWirelessGetNetworkWirelessSsidRadiusServers{
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
			return &[]ResponseWirelessGetNetworkWirelessSsidRadiusServers{}
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
		WalledGardenRanges: StringSliceToList(response.WalledGardenRanges),
		WpaEncryptionMode:  types.StringValue(response.WpaEncryptionMode),
	}
	state.Item = &itemState
	return state
}
