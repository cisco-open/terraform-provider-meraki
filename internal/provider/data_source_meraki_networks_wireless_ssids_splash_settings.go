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
	_ datasource.DataSource              = &NetworksWirelessSSIDsSplashSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsSplashSettingsDataSource{}
)

func NewNetworksWirelessSSIDsSplashSettingsDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsSplashSettingsDataSource{}
}

type NetworksWirelessSSIDsSplashSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsSplashSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsSplashSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_splash_settings"
}

func (d *NetworksWirelessSSIDsSplashSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

					"allow_simultaneous_logins": schema.BoolAttribute{
						MarkdownDescription: `Whether or not to allow simultaneous logins from different devices.`,
						Computed:            true,
					},
					"billing": schema.SingleNestedAttribute{
						MarkdownDescription: `Details associated with billing splash`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"free_access": schema.SingleNestedAttribute{
								MarkdownDescription: `Details associated with a free access plan with limits`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"duration_in_minutes": schema.Int64Attribute{
										MarkdownDescription: `How long a device can use a network for free.`,
										Computed:            true,
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Whether or not free access is enabled.`,
										Computed:            true,
									},
								},
							},
							"prepaid_access_fast_login_enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether or not billing uses the fast login prepaid access option.`,
								Computed:            true,
							},
							"reply_to_email_address": schema.StringAttribute{
								MarkdownDescription: `The email address that reeceives replies from clients`,
								Computed:            true,
							},
						},
					},
					"block_all_traffic_before_sign_on": schema.BoolAttribute{
						MarkdownDescription: `How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.`,
						Computed:            true,
					},
					"controller_disconnection_behavior": schema.StringAttribute{
						MarkdownDescription: `How login attempts should be handled when the controller is unreachable.`,
						Computed:            true,
					},
					"guest_sponsorship": schema.SingleNestedAttribute{
						MarkdownDescription: `Details associated with guest sponsored splash`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"duration_in_minutes": schema.Int64Attribute{
								MarkdownDescription: `Duration in minutes of sponsored guest authorization.`,
								Computed:            true,
							},
							"guest_can_request_timeframe": schema.BoolAttribute{
								MarkdownDescription: `Whether or not guests can specify how much time they are requesting.`,
								Computed:            true,
							},
						},
					},
					"redirect_url": schema.StringAttribute{
						MarkdownDescription: `The custom redirect URL where the users will go after the splash page.`,
						Computed:            true,
					},
					"self_registration": schema.SingleNestedAttribute{
						MarkdownDescription: `Self-registration for splash with Meraki authentication.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"authorization_type": schema.StringAttribute{
								MarkdownDescription: `How created user accounts should be authorized.`,
								Computed:            true,
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether or not to allow users to create their own account on the network.`,
								Computed:            true,
							},
						},
					},
					"sentry_enrollment": schema.SingleNestedAttribute{
						MarkdownDescription: `Systems Manager sentry enrollment splash settings.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enforced_systems": schema.ListAttribute{
								MarkdownDescription: `The system types that the Sentry enforces.`,
								Computed:            true,
								ElementType:         types.StringType,
							},
							"strength": schema.StringAttribute{
								MarkdownDescription: `The strength of the enforcement of selected system types.`,
								Computed:            true,
							},
							"systems_manager_network": schema.SingleNestedAttribute{
								MarkdownDescription: `Systems Manager network targeted for sentry enrollment.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"id": schema.StringAttribute{
										MarkdownDescription: `The network ID of the Systems Manager network.`,
										Computed:            true,
									},
								},
							},
						},
					},
					"splash_image": schema.SingleNestedAttribute{
						MarkdownDescription: `The image used in the splash page.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"extension": schema.StringAttribute{
								MarkdownDescription: `The extension of the image file.`,
								Computed:            true,
							},
							"md5": schema.StringAttribute{
								MarkdownDescription: `The MD5 value of the image file.`,
								Computed:            true,
							},
						},
					},
					"splash_logo": schema.SingleNestedAttribute{
						MarkdownDescription: `The logo used in the splash page.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"extension": schema.StringAttribute{
								MarkdownDescription: `The extension of the logo file.`,
								Computed:            true,
							},
							"md5": schema.StringAttribute{
								MarkdownDescription: `The MD5 value of the logo file.`,
								Computed:            true,
							},
						},
					},
					"splash_page": schema.StringAttribute{
						MarkdownDescription: `The type of splash page for this SSID`,
						Computed:            true,
					},
					"splash_prepaid_front": schema.SingleNestedAttribute{
						MarkdownDescription: `The prepaid front image used in the splash page.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"extension": schema.StringAttribute{
								MarkdownDescription: `The extension of the prepaid front image file.`,
								Computed:            true,
							},
							"md5": schema.StringAttribute{
								MarkdownDescription: `The MD5 value of the prepaid front image file.`,
								Computed:            true,
							},
						},
					},
					"splash_timeout": schema.Int64Attribute{
						MarkdownDescription: `Splash timeout in minutes.`,
						Computed:            true,
					},
					"splash_url": schema.StringAttribute{
						MarkdownDescription: `The custom splash URL of the click-through splash page.`,
						Computed:            true,
					},
					"ssid_number": schema.Int64Attribute{
						MarkdownDescription: `SSID number`,
						Computed:            true,
					},
					"theme_id": schema.StringAttribute{
						MarkdownDescription: `The id of the selected splash theme.`,
						Computed:            true,
					},
					"use_redirect_url": schema.BoolAttribute{
						MarkdownDescription: `The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page.`,
						Computed:            true,
					},
					"use_splash_url": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether the users will be redirected to the custom splash url`,
						Computed:            true,
					},
					"welcome_message": schema.StringAttribute{
						MarkdownDescription: `The welcome message for the users on the splash page.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessSSIDsSplashSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsSplashSettings NetworksWirelessSSIDsSplashSettings
	diags := req.Config.Get(ctx, &networksWirelessSSIDsSplashSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDSplashSettings")
		vvNetworkID := networksWirelessSSIDsSplashSettings.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsSplashSettings.Number.ValueString()

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDSplashSettings(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDSplashSettings",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsSplashSettings = ResponseWirelessGetNetworkWirelessSSIDSplashSettingsItemToBody(networksWirelessSSIDsSplashSettings, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsSplashSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsSplashSettings struct {
	NetworkID types.String                                          `tfsdk:"network_id"`
	Number    types.String                                          `tfsdk:"number"`
	Item      *ResponseWirelessGetNetworkWirelessSsidSplashSettings `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettings struct {
	AllowSimultaneousLogins         types.Bool                                                              `tfsdk:"allow_simultaneous_logins"`
	Billing                         *ResponseWirelessGetNetworkWirelessSsidSplashSettingsBilling            `tfsdk:"billing"`
	BlockAllTrafficBeforeSignOn     types.Bool                                                              `tfsdk:"block_all_traffic_before_sign_on"`
	ControllerDisconnectionBehavior types.String                                                            `tfsdk:"controller_disconnection_behavior"`
	GuestSponsorship                *ResponseWirelessGetNetworkWirelessSsidSplashSettingsGuestSponsorship   `tfsdk:"guest_sponsorship"`
	RedirectURL                     types.String                                                            `tfsdk:"redirect_url"`
	SelfRegistration                *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSelfRegistration   `tfsdk:"self_registration"`
	SentryEnrollment                *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollment   `tfsdk:"sentry_enrollment"`
	SplashImage                     *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashImage        `tfsdk:"splash_image"`
	SplashLogo                      *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashLogo         `tfsdk:"splash_logo"`
	SplashPage                      types.String                                                            `tfsdk:"splash_page"`
	SplashPrepaidFront              *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashPrepaidFront `tfsdk:"splash_prepaid_front"`
	SplashTimeout                   types.Int64                                                             `tfsdk:"splash_timeout"`
	SplashURL                       types.String                                                            `tfsdk:"splash_url"`
	SSIDNumber                      types.Int64                                                             `tfsdk:"ssid_number"`
	ThemeID                         types.String                                                            `tfsdk:"theme_id"`
	UseRedirectURL                  types.Bool                                                              `tfsdk:"use_redirect_url"`
	UseSplashURL                    types.Bool                                                              `tfsdk:"use_splash_url"`
	WelcomeMessage                  types.String                                                            `tfsdk:"welcome_message"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsBilling struct {
	FreeAccess                    *ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingFreeAccess `tfsdk:"free_access"`
	PrepaidAccessFastLoginEnabled types.Bool                                                             `tfsdk:"prepaid_access_fast_login_enabled"`
	ReplyToEmailAddress           types.String                                                           `tfsdk:"reply_to_email_address"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingFreeAccess struct {
	DurationInMinutes types.Int64 `tfsdk:"duration_in_minutes"`
	Enabled           types.Bool  `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsGuestSponsorship struct {
	DurationInMinutes        types.Int64 `tfsdk:"duration_in_minutes"`
	GuestCanRequestTimeframe types.Bool  `tfsdk:"guest_can_request_timeframe"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsSelfRegistration struct {
	AuthorizationType types.String `tfsdk:"authorization_type"`
	Enabled           types.Bool   `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollment struct {
	EnforcedSystems       types.List                                                                                 `tfsdk:"enforced_systems"`
	Strength              types.String                                                                               `tfsdk:"strength"`
	SystemsManagerNetwork *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentSystemsManagerNetwork `tfsdk:"systems_manager_network"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentSystemsManagerNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashImage struct {
	Extension types.String `tfsdk:"extension"`
	Md5       types.String `tfsdk:"md5"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashLogo struct {
	Extension types.String `tfsdk:"extension"`
	Md5       types.String `tfsdk:"md5"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashPrepaidFront struct {
	Extension types.String `tfsdk:"extension"`
	Md5       types.String `tfsdk:"md5"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDSplashSettingsItemToBody(state NetworksWirelessSSIDsSplashSettings, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDSplashSettings) NetworksWirelessSSIDsSplashSettings {
	itemState := ResponseWirelessGetNetworkWirelessSsidSplashSettings{
		AllowSimultaneousLogins: func() types.Bool {
			if response.AllowSimultaneousLogins != nil {
				return types.BoolValue(*response.AllowSimultaneousLogins)
			}
			return types.Bool{}
		}(),
		Billing: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsBilling {
			if response.Billing != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsBilling{
					FreeAccess: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingFreeAccess {
						if response.Billing.FreeAccess != nil {
							return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingFreeAccess{
								DurationInMinutes: func() types.Int64 {
									if response.Billing.FreeAccess.DurationInMinutes != nil {
										return types.Int64Value(int64(*response.Billing.FreeAccess.DurationInMinutes))
									}
									return types.Int64{}
								}(),
								Enabled: func() types.Bool {
									if response.Billing.FreeAccess.Enabled != nil {
										return types.BoolValue(*response.Billing.FreeAccess.Enabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingFreeAccess{}
					}(),
					PrepaidAccessFastLoginEnabled: func() types.Bool {
						if response.Billing.PrepaidAccessFastLoginEnabled != nil {
							return types.BoolValue(*response.Billing.PrepaidAccessFastLoginEnabled)
						}
						return types.Bool{}
					}(),
					ReplyToEmailAddress: types.StringValue(response.Billing.ReplyToEmailAddress),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsBilling{}
		}(),
		BlockAllTrafficBeforeSignOn: func() types.Bool {
			if response.BlockAllTrafficBeforeSignOn != nil {
				return types.BoolValue(*response.BlockAllTrafficBeforeSignOn)
			}
			return types.Bool{}
		}(),
		ControllerDisconnectionBehavior: types.StringValue(response.ControllerDisconnectionBehavior),
		GuestSponsorship: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsGuestSponsorship {
			if response.GuestSponsorship != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsGuestSponsorship{
					DurationInMinutes: func() types.Int64 {
						if response.GuestSponsorship.DurationInMinutes != nil {
							return types.Int64Value(int64(*response.GuestSponsorship.DurationInMinutes))
						}
						return types.Int64{}
					}(),
					GuestCanRequestTimeframe: func() types.Bool {
						if response.GuestSponsorship.GuestCanRequestTimeframe != nil {
							return types.BoolValue(*response.GuestSponsorship.GuestCanRequestTimeframe)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsGuestSponsorship{}
		}(),
		RedirectURL: types.StringValue(response.RedirectURL),
		SelfRegistration: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSelfRegistration {
			if response.SelfRegistration != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSelfRegistration{
					AuthorizationType: types.StringValue(response.SelfRegistration.AuthorizationType),
					Enabled: func() types.Bool {
						if response.SelfRegistration.Enabled != nil {
							return types.BoolValue(*response.SelfRegistration.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSelfRegistration{}
		}(),
		SentryEnrollment: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollment {
			if response.SentryEnrollment != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollment{
					EnforcedSystems: StringSliceToList(response.SentryEnrollment.EnforcedSystems),
					Strength:        types.StringValue(response.SentryEnrollment.Strength),
					SystemsManagerNetwork: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentSystemsManagerNetwork {
						if response.SentryEnrollment.SystemsManagerNetwork != nil {
							return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentSystemsManagerNetwork{
								ID: types.StringValue(response.SentryEnrollment.SystemsManagerNetwork.ID),
							}
						}
						return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentSystemsManagerNetwork{}
					}(),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollment{}
		}(),
		SplashImage: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashImage {
			if response.SplashImage != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashImage{
					Extension: types.StringValue(response.SplashImage.Extension),
					Md5:       types.StringValue(response.SplashImage.Md5),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashImage{}
		}(),
		SplashLogo: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashLogo {
			if response.SplashLogo != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashLogo{
					Extension: types.StringValue(response.SplashLogo.Extension),
					Md5:       types.StringValue(response.SplashLogo.Md5),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashLogo{}
		}(),
		SplashPage: types.StringValue(response.SplashPage),
		SplashPrepaidFront: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashPrepaidFront {
			if response.SplashPrepaidFront != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashPrepaidFront{
					Extension: types.StringValue(response.SplashPrepaidFront.Extension),
					Md5:       types.StringValue(response.SplashPrepaidFront.Md5),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashPrepaidFront{}
		}(),
		SplashTimeout: func() types.Int64 {
			if response.SplashTimeout != nil {
				return types.Int64Value(int64(*response.SplashTimeout))
			}
			return types.Int64{}
		}(),
		SplashURL: types.StringValue(response.SplashURL),
		SSIDNumber: func() types.Int64 {
			if response.SSIDNumber != nil {
				return types.Int64Value(int64(*response.SSIDNumber))
			}
			return types.Int64{}
		}(),
		ThemeID: types.StringValue(response.ThemeID),
		UseRedirectURL: func() types.Bool {
			if response.UseRedirectURL != nil {
				return types.BoolValue(*response.UseRedirectURL)
			}
			return types.Bool{}
		}(),
		UseSplashURL: func() types.Bool {
			if response.UseSplashURL != nil {
				return types.BoolValue(*response.UseSplashURL)
			}
			return types.Bool{}
		}(),
		WelcomeMessage: types.StringValue(response.WelcomeMessage),
	}
	state.Item = &itemState
	return state
}
