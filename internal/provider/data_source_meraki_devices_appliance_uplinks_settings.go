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
	_ datasource.DataSource              = &DevicesApplianceUplinksSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesApplianceUplinksSettingsDataSource{}
)

func NewDevicesApplianceUplinksSettingsDataSource() datasource.DataSource {
	return &DevicesApplianceUplinksSettingsDataSource{}
}

type DevicesApplianceUplinksSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesApplianceUplinksSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesApplianceUplinksSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_appliance_uplinks_settings"
}

func (d *DevicesApplianceUplinksSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"interfaces": schema.SingleNestedAttribute{
						MarkdownDescription: `Interface settings.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"wan1": schema.SingleNestedAttribute{
								MarkdownDescription: `WAN 1 settings.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Enable or disable the interface.`,
										Computed:            true,
									},
									"pppoe": schema.SingleNestedAttribute{
										MarkdownDescription: `Configuration options for PPPoE.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"authentication": schema.SingleNestedAttribute{
												MarkdownDescription: `Settings for PPPoE Authentication.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"enabled": schema.BoolAttribute{
														MarkdownDescription: `Whether PPPoE authentication is enabled.`,
														Computed:            true,
													},
													"username": schema.StringAttribute{
														MarkdownDescription: `Username for PPPoE authentication.`,
														Computed:            true,
													},
												},
											},
											"enabled": schema.BoolAttribute{
												MarkdownDescription: `Whether PPPoE is enabled.`,
												Computed:            true,
											},
										},
									},
									"svis": schema.SingleNestedAttribute{
										MarkdownDescription: `SVI settings by protocol.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"ipv4": schema.SingleNestedAttribute{
												MarkdownDescription: `IPv4 settings for static/dynamic mode.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"address": schema.StringAttribute{
														MarkdownDescription: `IP address and subnet mask when in static mode.`,
														Computed:            true,
													},
													"assignment_mode": schema.StringAttribute{
														MarkdownDescription: `The assignment mode for this SVI. Applies only when PPPoE is disabled.`,
														Computed:            true,
													},
													"gateway": schema.StringAttribute{
														MarkdownDescription: `Gateway IP address when in static mode.`,
														Computed:            true,
													},
													"nameservers": schema.SingleNestedAttribute{
														MarkdownDescription: `The nameserver settings for this SVI.`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"addresses": schema.ListAttribute{
																MarkdownDescription: `Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.`,
																Computed:            true,
																ElementType:         types.StringType,
															},
														},
													},
												},
											},
											"ipv6": schema.SingleNestedAttribute{
												MarkdownDescription: `IPv6 settings for static/dynamic mode.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"address": schema.StringAttribute{
														MarkdownDescription: `Static address that will override the one(s) received by SLAAC.`,
														Computed:            true,
													},
													"assignment_mode": schema.StringAttribute{
														MarkdownDescription: `The assignment mode for this SVI. Applies only when PPPoE is disabled.`,
														Computed:            true,
													},
													"gateway": schema.StringAttribute{
														MarkdownDescription: `Static gateway that will override the one received by autoconf.`,
														Computed:            true,
													},
													"nameservers": schema.SingleNestedAttribute{
														MarkdownDescription: `The nameserver settings for this SVI.`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"addresses": schema.ListAttribute{
																MarkdownDescription: `Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.`,
																Computed:            true,
																ElementType:         types.StringType,
															},
														},
													},
												},
											},
										},
									},
									"vlan_tagging": schema.SingleNestedAttribute{
										MarkdownDescription: `VLAN tagging settings.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.BoolAttribute{
												MarkdownDescription: `Whether VLAN tagging is enabled.`,
												Computed:            true,
											},
											"vlan_id": schema.Int64Attribute{
												MarkdownDescription: `The ID of the VLAN to use for VLAN tagging.`,
												Computed:            true,
											},
										},
									},
								},
							},
							"wan2": schema.SingleNestedAttribute{
								MarkdownDescription: `WAN 2 settings.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Enable or disable the interface.`,
										Computed:            true,
									},
									"pppoe": schema.SingleNestedAttribute{
										MarkdownDescription: `Configuration options for PPPoE.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"authentication": schema.SingleNestedAttribute{
												MarkdownDescription: `Settings for PPPoE Authentication.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"enabled": schema.BoolAttribute{
														MarkdownDescription: `Whether PPPoE authentication is enabled.`,
														Computed:            true,
													},
													"username": schema.StringAttribute{
														MarkdownDescription: `Username for PPPoE authentication.`,
														Computed:            true,
													},
												},
											},
											"enabled": schema.BoolAttribute{
												MarkdownDescription: `Whether PPPoE is enabled.`,
												Computed:            true,
											},
										},
									},
									"svis": schema.SingleNestedAttribute{
										MarkdownDescription: `SVI settings by protocol.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"ipv4": schema.SingleNestedAttribute{
												MarkdownDescription: `IPv4 settings for static/dynamic mode.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"address": schema.StringAttribute{
														MarkdownDescription: `IP address and subnet mask when in static mode.`,
														Computed:            true,
													},
													"assignment_mode": schema.StringAttribute{
														MarkdownDescription: `The assignment mode for this SVI. Applies only when PPPoE is disabled.`,
														Computed:            true,
													},
													"gateway": schema.StringAttribute{
														MarkdownDescription: `Gateway IP address when in static mode.`,
														Computed:            true,
													},
													"nameservers": schema.SingleNestedAttribute{
														MarkdownDescription: `The nameserver settings for this SVI.`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"addresses": schema.ListAttribute{
																MarkdownDescription: `Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.`,
																Computed:            true,
																ElementType:         types.StringType,
															},
														},
													},
												},
											},
											"ipv6": schema.SingleNestedAttribute{
												MarkdownDescription: `IPv6 settings for static/dynamic mode.`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"address": schema.StringAttribute{
														MarkdownDescription: `Static address that will override the one(s) received by SLAAC.`,
														Computed:            true,
													},
													"assignment_mode": schema.StringAttribute{
														MarkdownDescription: `The assignment mode for this SVI. Applies only when PPPoE is disabled.`,
														Computed:            true,
													},
													"gateway": schema.StringAttribute{
														MarkdownDescription: `Static gateway that will override the one received by autoconf.`,
														Computed:            true,
													},
													"nameservers": schema.SingleNestedAttribute{
														MarkdownDescription: `The nameserver settings for this SVI.`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"addresses": schema.ListAttribute{
																MarkdownDescription: `Up to 2 nameserver addresses to use, ordered in priority from highest to lowest priority.`,
																Computed:            true,
																ElementType:         types.StringType,
															},
														},
													},
												},
											},
										},
									},
									"vlan_tagging": schema.SingleNestedAttribute{
										MarkdownDescription: `VLAN tagging settings.`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"enabled": schema.BoolAttribute{
												MarkdownDescription: `Whether VLAN tagging is enabled.`,
												Computed:            true,
											},
											"vlan_id": schema.Int64Attribute{
												MarkdownDescription: `The ID of the VLAN to use for VLAN tagging.`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *DevicesApplianceUplinksSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesApplianceUplinksSettings DevicesApplianceUplinksSettings
	diags := req.Config.Get(ctx, &devicesApplianceUplinksSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceApplianceUplinksSettings")
		vvSerial := devicesApplianceUplinksSettings.Serial.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetDeviceApplianceUplinksSettings(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceApplianceUplinksSettings",
				err.Error(),
			)
			return
		}

		devicesApplianceUplinksSettings = ResponseApplianceGetDeviceApplianceUplinksSettingsItemToBody(devicesApplianceUplinksSettings, response1)
		diags = resp.State.Set(ctx, &devicesApplianceUplinksSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesApplianceUplinksSettings struct {
	Serial types.String                                        `tfsdk:"serial"`
	Item   *ResponseApplianceGetDeviceApplianceUplinksSettings `tfsdk:"item"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettings struct {
	Interfaces *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfaces `tfsdk:"interfaces"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfaces struct {
	Wan1 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1 `tfsdk:"wan1"`
	Wan2 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2 `tfsdk:"wan2"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1 struct {
	Enabled     types.Bool                                                                   `tfsdk:"enabled"`
	Pppoe       *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Pppoe       `tfsdk:"pppoe"`
	Svis        *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Svis        `tfsdk:"svis"`
	VLANTagging *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1VlanTagging `tfsdk:"vlan_tagging"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Pppoe struct {
	Authentication *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication `tfsdk:"authentication"`
	Enabled        types.Bool                                                                           `tfsdk:"enabled"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication struct {
	Enabled  types.Bool   `tfsdk:"enabled"`
	Username types.String `tfsdk:"username"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Svis struct {
	IPv4 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4 `tfsdk:"ipv4"`
	IPv6 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6 `tfsdk:"ipv6"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4 struct {
	Address        types.String                                                                         `tfsdk:"address"`
	AssignmentMode types.String                                                                         `tfsdk:"assignment_mode"`
	Gateway        types.String                                                                         `tfsdk:"gateway"`
	Nameservers    *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4Nameservers `tfsdk:"nameservers"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4Nameservers struct {
	Addresses types.List `tfsdk:"addresses"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6 struct {
	Address        types.String                                                                         `tfsdk:"address"`
	AssignmentMode types.String                                                                         `tfsdk:"assignment_mode"`
	Gateway        types.String                                                                         `tfsdk:"gateway"`
	Nameservers    *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6Nameservers `tfsdk:"nameservers"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6Nameservers struct {
	Addresses types.List `tfsdk:"addresses"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1VlanTagging struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	VLANID  types.Int64 `tfsdk:"vlan_id"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2 struct {
	Enabled     types.Bool                                                                   `tfsdk:"enabled"`
	Pppoe       *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Pppoe       `tfsdk:"pppoe"`
	Svis        *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Svis        `tfsdk:"svis"`
	VLANTagging *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2VlanTagging `tfsdk:"vlan_tagging"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Pppoe struct {
	Authentication *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication `tfsdk:"authentication"`
	Enabled        types.Bool                                                                           `tfsdk:"enabled"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication struct {
	Enabled  types.Bool   `tfsdk:"enabled"`
	Username types.String `tfsdk:"username"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Svis struct {
	IPv4 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4 `tfsdk:"ipv4"`
	IPv6 *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6 `tfsdk:"ipv6"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4 struct {
	Address        types.String                                                                         `tfsdk:"address"`
	AssignmentMode types.String                                                                         `tfsdk:"assignment_mode"`
	Gateway        types.String                                                                         `tfsdk:"gateway"`
	Nameservers    *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4Nameservers `tfsdk:"nameservers"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4Nameservers struct {
	Addresses types.List `tfsdk:"addresses"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6 struct {
	Address        types.String                                                                         `tfsdk:"address"`
	AssignmentMode types.String                                                                         `tfsdk:"assignment_mode"`
	Gateway        types.String                                                                         `tfsdk:"gateway"`
	Nameservers    *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6Nameservers `tfsdk:"nameservers"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6Nameservers struct {
	Addresses types.List `tfsdk:"addresses"`
}

type ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2VlanTagging struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	VLANID  types.Int64 `tfsdk:"vlan_id"`
}

// ToBody
func ResponseApplianceGetDeviceApplianceUplinksSettingsItemToBody(state DevicesApplianceUplinksSettings, response *merakigosdk.ResponseApplianceGetDeviceApplianceUplinksSettings) DevicesApplianceUplinksSettings {
	itemState := ResponseApplianceGetDeviceApplianceUplinksSettings{
		Interfaces: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfaces {
			if response.Interfaces != nil {
				return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfaces{
					Wan1: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1 {
						if response.Interfaces.Wan1 != nil {
							return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1{
								Enabled: func() types.Bool {
									if response.Interfaces.Wan1.Enabled != nil {
										return types.BoolValue(*response.Interfaces.Wan1.Enabled)
									}
									return types.Bool{}
								}(),
								Pppoe: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Pppoe {
									if response.Interfaces.Wan1.Pppoe != nil {
										return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Pppoe{
											Authentication: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication {
												if response.Interfaces.Wan1.Pppoe.Authentication != nil {
													return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication{
														Enabled: func() types.Bool {
															if response.Interfaces.Wan1.Pppoe.Authentication.Enabled != nil {
																return types.BoolValue(*response.Interfaces.Wan1.Pppoe.Authentication.Enabled)
															}
															return types.Bool{}
														}(),
														Username: types.StringValue(response.Interfaces.Wan1.Pppoe.Authentication.Username),
													}
												}
												return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1PppoeAuthentication{}
											}(),
											Enabled: func() types.Bool {
												if response.Interfaces.Wan1.Pppoe.Enabled != nil {
													return types.BoolValue(*response.Interfaces.Wan1.Pppoe.Enabled)
												}
												return types.Bool{}
											}(),
										}
									}
									return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Pppoe{}
								}(),
								Svis: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Svis {
									if response.Interfaces.Wan1.Svis != nil {
										return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Svis{
											IPv4: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4 {
												if response.Interfaces.Wan1.Svis.IPv4 != nil {
													return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4{
														Address:        types.StringValue(response.Interfaces.Wan1.Svis.IPv4.Address),
														AssignmentMode: types.StringValue(response.Interfaces.Wan1.Svis.IPv4.AssignmentMode),
														Gateway:        types.StringValue(response.Interfaces.Wan1.Svis.IPv4.Gateway),
														Nameservers: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4Nameservers {
															if response.Interfaces.Wan1.Svis.IPv4.Nameservers != nil {
																return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4Nameservers{
																	Addresses: StringSliceToList(response.Interfaces.Wan1.Svis.IPv4.Nameservers.Addresses),
																}
															}
															return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4Nameservers{}
														}(),
													}
												}
												return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv4{}
											}(),
											IPv6: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6 {
												if response.Interfaces.Wan1.Svis.IPv6 != nil {
													return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6{
														Address:        types.StringValue(response.Interfaces.Wan1.Svis.IPv6.Address),
														AssignmentMode: types.StringValue(response.Interfaces.Wan1.Svis.IPv6.AssignmentMode),
														Gateway:        types.StringValue(response.Interfaces.Wan1.Svis.IPv6.Gateway),
														Nameservers: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6Nameservers {
															if response.Interfaces.Wan1.Svis.IPv6.Nameservers != nil {
																return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6Nameservers{
																	Addresses: StringSliceToList(response.Interfaces.Wan1.Svis.IPv6.Nameservers.Addresses),
																}
															}
															return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6Nameservers{}
														}(),
													}
												}
												return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1SvisIpv6{}
											}(),
										}
									}
									return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1Svis{}
								}(),
								VLANTagging: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1VlanTagging {
									if response.Interfaces.Wan1.VLANTagging != nil {
										return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1VlanTagging{
											Enabled: func() types.Bool {
												if response.Interfaces.Wan1.VLANTagging.Enabled != nil {
													return types.BoolValue(*response.Interfaces.Wan1.VLANTagging.Enabled)
												}
												return types.Bool{}
											}(),
											VLANID: func() types.Int64 {
												if response.Interfaces.Wan1.VLANTagging.VLANID != nil {
													return types.Int64Value(int64(*response.Interfaces.Wan1.VLANTagging.VLANID))
												}
												return types.Int64{}
											}(),
										}
									}
									return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1VlanTagging{}
								}(),
							}
						}
						return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan1{}
					}(),
					Wan2: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2 {
						if response.Interfaces.Wan2 != nil {
							return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2{
								Enabled: func() types.Bool {
									if response.Interfaces.Wan2.Enabled != nil {
										return types.BoolValue(*response.Interfaces.Wan2.Enabled)
									}
									return types.Bool{}
								}(),
								Pppoe: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Pppoe {
									if response.Interfaces.Wan2.Pppoe != nil {
										return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Pppoe{
											Authentication: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication {
												if response.Interfaces.Wan2.Pppoe.Authentication != nil {
													return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication{
														Enabled: func() types.Bool {
															if response.Interfaces.Wan2.Pppoe.Authentication.Enabled != nil {
																return types.BoolValue(*response.Interfaces.Wan2.Pppoe.Authentication.Enabled)
															}
															return types.Bool{}
														}(),
														Username: types.StringValue(response.Interfaces.Wan2.Pppoe.Authentication.Username),
													}
												}
												return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2PppoeAuthentication{}
											}(),
											Enabled: func() types.Bool {
												if response.Interfaces.Wan2.Pppoe.Enabled != nil {
													return types.BoolValue(*response.Interfaces.Wan2.Pppoe.Enabled)
												}
												return types.Bool{}
											}(),
										}
									}
									return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Pppoe{}
								}(),
								Svis: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Svis {
									if response.Interfaces.Wan2.Svis != nil {
										return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Svis{
											IPv4: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4 {
												if response.Interfaces.Wan2.Svis.IPv4 != nil {
													return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4{
														Address:        types.StringValue(response.Interfaces.Wan2.Svis.IPv4.Address),
														AssignmentMode: types.StringValue(response.Interfaces.Wan2.Svis.IPv4.AssignmentMode),
														Gateway:        types.StringValue(response.Interfaces.Wan2.Svis.IPv4.Gateway),
														Nameservers: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4Nameservers {
															if response.Interfaces.Wan2.Svis.IPv4.Nameservers != nil {
																return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4Nameservers{
																	Addresses: StringSliceToList(response.Interfaces.Wan2.Svis.IPv4.Nameservers.Addresses),
																}
															}
															return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4Nameservers{}
														}(),
													}
												}
												return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv4{}
											}(),
											IPv6: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6 {
												if response.Interfaces.Wan2.Svis.IPv6 != nil {
													return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6{
														Address:        types.StringValue(response.Interfaces.Wan2.Svis.IPv6.Address),
														AssignmentMode: types.StringValue(response.Interfaces.Wan2.Svis.IPv6.AssignmentMode),
														Gateway:        types.StringValue(response.Interfaces.Wan2.Svis.IPv6.Gateway),
														Nameservers: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6Nameservers {
															if response.Interfaces.Wan2.Svis.IPv6.Nameservers != nil {
																return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6Nameservers{
																	Addresses: StringSliceToList(response.Interfaces.Wan2.Svis.IPv6.Nameservers.Addresses),
																}
															}
															return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6Nameservers{}
														}(),
													}
												}
												return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2SvisIpv6{}
											}(),
										}
									}
									return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2Svis{}
								}(),
								VLANTagging: func() *ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2VlanTagging {
									if response.Interfaces.Wan2.VLANTagging != nil {
										return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2VlanTagging{
											Enabled: func() types.Bool {
												if response.Interfaces.Wan2.VLANTagging.Enabled != nil {
													return types.BoolValue(*response.Interfaces.Wan2.VLANTagging.Enabled)
												}
												return types.Bool{}
											}(),
											VLANID: func() types.Int64 {
												if response.Interfaces.Wan2.VLANTagging.VLANID != nil {
													return types.Int64Value(int64(*response.Interfaces.Wan2.VLANTagging.VLANID))
												}
												return types.Int64{}
											}(),
										}
									}
									return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2VlanTagging{}
								}(),
							}
						}
						return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfacesWan2{}
					}(),
				}
			}
			return &ResponseApplianceGetDeviceApplianceUplinksSettingsInterfaces{}
		}(),
	}
	state.Item = &itemState
	return state
}
