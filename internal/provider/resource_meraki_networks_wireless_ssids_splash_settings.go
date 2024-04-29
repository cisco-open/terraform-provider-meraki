package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
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
	_ resource.Resource              = &NetworksWirelessSSIDsSplashSettingsResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsSplashSettingsResource{}
)

func NewNetworksWirelessSSIDsSplashSettingsResource() resource.Resource {
	return &NetworksWirelessSSIDsSplashSettingsResource{}
}

type NetworksWirelessSSIDsSplashSettingsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsSplashSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsSplashSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_splash_settings"
}

func (r *NetworksWirelessSSIDsSplashSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"allow_simultaneous_logins": schema.BoolAttribute{
				MarkdownDescription: `Whether or not to allow simultaneous logins from different devices.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"billing": schema.SingleNestedAttribute{
				MarkdownDescription: `Details associated with billing splash`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"free_access": schema.SingleNestedAttribute{
						MarkdownDescription: `Details associated with a free access plan with limits`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"duration_in_minutes": schema.Int64Attribute{
								MarkdownDescription: `How long a device can use a network for free.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether or not free access is enabled.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"prepaid_access_fast_login_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether or not billing uses the fast login prepaid access option.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"reply_to_email_address": schema.StringAttribute{
						MarkdownDescription: `The email address that reeceives replies from clients`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"block_all_traffic_before_sign_on": schema.BoolAttribute{
				MarkdownDescription: `How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"controller_disconnection_behavior": schema.StringAttribute{
				MarkdownDescription: `How login attempts should be handled when the controller is unreachable.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"default",
						"open",
						"restricted",
					),
				},
			},
			"guest_sponsorship": schema.SingleNestedAttribute{
				MarkdownDescription: `Details associated with guest sponsored splash`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"duration_in_minutes": schema.Int64Attribute{
						MarkdownDescription: `Duration in minutes of sponsored guest authorization.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"guest_can_request_timeframe": schema.BoolAttribute{
						MarkdownDescription: `Whether or not guests can specify how much time they are requesting.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"redirect_url": schema.StringAttribute{
				MarkdownDescription: `The custom redirect URL where the users will go after the splash page.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enforced_systems": schema.SetAttribute{
						MarkdownDescription: `The system types that the Sentry enforces.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
					"strength": schema.StringAttribute{
						MarkdownDescription: `The strength of the enforcement of selected system types.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"click-through",
								"focused",
								"strict",
							),
						},
					},
					"systems_manager_network": schema.SingleNestedAttribute{
						MarkdownDescription: `Systems Manager network targeted for sentry enrollment.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `The network ID of the Systems Manager network.`,
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
			"splash_image": schema.SingleNestedAttribute{
				MarkdownDescription: `The image used in the splash page.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"extension": schema.StringAttribute{
						MarkdownDescription: `The extension of the image file.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"image": schema.SingleNestedAttribute{
						MarkdownDescription: `Properties for setting a new image.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"contents": schema.StringAttribute{
								MarkdownDescription: `The file contents (a base 64 encoded string) of your new image.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"format": schema.StringAttribute{
								MarkdownDescription: `The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"gif",
										"jpg",
										"png",
									),
								},
							},
						},
					},
					"md5": schema.StringAttribute{
						MarkdownDescription: `The MD5 value of the image file.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"splash_logo": schema.SingleNestedAttribute{
				MarkdownDescription: `The logo used in the splash page.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"extension": schema.StringAttribute{
						MarkdownDescription: `The extension of the logo file.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"image": schema.SingleNestedAttribute{
						MarkdownDescription: `Properties for setting a new image.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"contents": schema.StringAttribute{
								MarkdownDescription: `The file contents (a base 64 encoded string) of your new logo.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"format": schema.StringAttribute{
								MarkdownDescription: `The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"gif",
										"jpg",
										"png",
									),
								},
							},
						},
					},
					"md5": schema.StringAttribute{
						MarkdownDescription: `The MD5 value of the logo file.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
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
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"extension": schema.StringAttribute{
						MarkdownDescription: `The extension of the prepaid front image file.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"image": schema.SingleNestedAttribute{
						MarkdownDescription: `Properties for setting a new image.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"contents": schema.StringAttribute{
								MarkdownDescription: `The file contents (a base 64 encoded string) of your new prepaid front.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"format": schema.StringAttribute{
								MarkdownDescription: `The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"gif",
										"jpg",
										"png",
									),
								},
							},
						},
					},
					"md5": schema.StringAttribute{
						MarkdownDescription: `The MD5 value of the prepaid front image file.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"splash_timeout": schema.Int64Attribute{
				MarkdownDescription: `Splash timeout in minutes.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"splash_url": schema.StringAttribute{
				MarkdownDescription: `The custom splash URL of the click-through splash page.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ssid_number": schema.Int64Attribute{
				MarkdownDescription: `SSID number`,
				Computed:            true,
			},
			"theme_id": schema.StringAttribute{
				MarkdownDescription: `The id of the selected splash theme.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"use_redirect_url": schema.BoolAttribute{
				MarkdownDescription: `The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"use_splash_url": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether the users will be redirected to the custom splash url`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"welcome_message": schema.StringAttribute{
				MarkdownDescription: `The welcome message for the users on the splash page.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *NetworksWirelessSSIDsSplashSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsSplashSettingsRs

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
	vvNumber := data.Number.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDSplashSettings(vvNetworkID, vvNumber)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDsSplashSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDsSplashSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDSplashSettings(vvNetworkID, vvNumber, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDSplashSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDSplashSettings",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDSplashSettings(vvNetworkID, vvNumber)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDSplashSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDSplashSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDSplashSettingsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsSplashSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsSplashSettingsRs

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
	vvNumber := data.Number.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDSplashSettings(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkWirelessSSIDSplashSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDSplashSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDSplashSettingsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessSSIDsSplashSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *NetworksWirelessSSIDsSplashSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessSSIDsSplashSettingsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvNumber := data.Number.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDSplashSettings(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDSplashSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDSplashSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsSplashSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessSSIDsSplashSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsSplashSettingsRs struct {
	NetworkID                       types.String                                                              `tfsdk:"network_id"`
	Number                          types.String                                                              `tfsdk:"number"`
	AllowSimultaneousLogins         types.Bool                                                                `tfsdk:"allow_simultaneous_logins"`
	Billing                         *ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingRs            `tfsdk:"billing"`
	BlockAllTrafficBeforeSignOn     types.Bool                                                                `tfsdk:"block_all_traffic_before_sign_on"`
	ControllerDisconnectionBehavior types.String                                                              `tfsdk:"controller_disconnection_behavior"`
	GuestSponsorship                *ResponseWirelessGetNetworkWirelessSsidSplashSettingsGuestSponsorshipRs   `tfsdk:"guest_sponsorship"`
	RedirectURL                     types.String                                                              `tfsdk:"redirect_url"`
	SelfRegistration                *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSelfRegistrationRs   `tfsdk:"self_registration"`
	SentryEnrollment                *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentRs   `tfsdk:"sentry_enrollment"`
	SplashImage                     *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashImageRs        `tfsdk:"splash_image"`
	SplashLogo                      *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashLogoRs         `tfsdk:"splash_logo"`
	SplashPage                      types.String                                                              `tfsdk:"splash_page"`
	SplashPrepaidFront              *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashPrepaidFrontRs `tfsdk:"splash_prepaid_front"`
	SplashTimeout                   types.Int64                                                               `tfsdk:"splash_timeout"`
	SplashURL                       types.String                                                              `tfsdk:"splash_url"`
	SSIDNumber                      types.Int64                                                               `tfsdk:"ssid_number"`
	ThemeID                         types.String                                                              `tfsdk:"theme_id"`
	UseRedirectURL                  types.Bool                                                                `tfsdk:"use_redirect_url"`
	UseSplashURL                    types.Bool                                                                `tfsdk:"use_splash_url"`
	WelcomeMessage                  types.String                                                              `tfsdk:"welcome_message"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingRs struct {
	FreeAccess                    *ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingFreeAccessRs `tfsdk:"free_access"`
	PrepaidAccessFastLoginEnabled types.Bool                                                               `tfsdk:"prepaid_access_fast_login_enabled"`
	ReplyToEmailAddress           types.String                                                             `tfsdk:"reply_to_email_address"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingFreeAccessRs struct {
	DurationInMinutes types.Int64 `tfsdk:"duration_in_minutes"`
	Enabled           types.Bool  `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsGuestSponsorshipRs struct {
	DurationInMinutes        types.Int64 `tfsdk:"duration_in_minutes"`
	GuestCanRequestTimeframe types.Bool  `tfsdk:"guest_can_request_timeframe"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsSelfRegistrationRs struct {
	AuthorizationType types.String `tfsdk:"authorization_type"`
	Enabled           types.Bool   `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentRs struct {
	EnforcedSystems       types.Set                                                                                    `tfsdk:"enforced_systems"`
	Strength              types.String                                                                                 `tfsdk:"strength"`
	SystemsManagerNetwork *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentSystemsManagerNetworkRs `tfsdk:"systems_manager_network"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentSystemsManagerNetworkRs struct {
	ID types.String `tfsdk:"id"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashImageRs struct {
	Extension types.String                                                              `tfsdk:"extension"`
	Md5       types.String                                                              `tfsdk:"md5"`
	Image     *RequestWirelessUpdateNetworkWirelessSsidSplashSettingsSplashImageImageRs `tfsdk:"image"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashLogoRs struct {
	Extension types.String                                                             `tfsdk:"extension"`
	Md5       types.String                                                             `tfsdk:"md5"`
	Image     *RequestWirelessUpdateNetworkWirelessSsidSplashSettingsSplashLogoImageRs `tfsdk:"image"`
}

type ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashPrepaidFrontRs struct {
	Extension types.String                                                                     `tfsdk:"extension"`
	Md5       types.String                                                                     `tfsdk:"md5"`
	Image     *RequestWirelessUpdateNetworkWirelessSsidSplashSettingsSplashPrepaidFrontImageRs `tfsdk:"image"`
}

type RequestWirelessUpdateNetworkWirelessSsidSplashSettingsSplashImageImageRs struct {
	Contents types.String `tfsdk:"contents"`
	Format   types.String `tfsdk:"format"`
}

type RequestWirelessUpdateNetworkWirelessSsidSplashSettingsSplashLogoImageRs struct {
	Contents types.String `tfsdk:"contents"`
	Format   types.String `tfsdk:"format"`
}

type RequestWirelessUpdateNetworkWirelessSsidSplashSettingsSplashPrepaidFrontImageRs struct {
	Contents types.String `tfsdk:"contents"`
	Format   types.String `tfsdk:"format"`
}

// FromBody
func (r *NetworksWirelessSSIDsSplashSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettings {
	emptyString := ""
	allowSimultaneousLogins := new(bool)
	if !r.AllowSimultaneousLogins.IsUnknown() && !r.AllowSimultaneousLogins.IsNull() {
		*allowSimultaneousLogins = r.AllowSimultaneousLogins.ValueBool()
	} else {
		allowSimultaneousLogins = nil
	}
	var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsBilling *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsBilling
	if r.Billing != nil {
		var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsBillingFreeAccess *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsBillingFreeAccess
		if r.Billing.FreeAccess != nil {
			durationInMinutes := func() *int64 {
				if !r.Billing.FreeAccess.DurationInMinutes.IsUnknown() && !r.Billing.FreeAccess.DurationInMinutes.IsNull() {
					return r.Billing.FreeAccess.DurationInMinutes.ValueInt64Pointer()
				}
				return nil
			}()
			enabled := func() *bool {
				if !r.Billing.FreeAccess.Enabled.IsUnknown() && !r.Billing.FreeAccess.Enabled.IsNull() {
					return r.Billing.FreeAccess.Enabled.ValueBoolPointer()
				}
				return nil
			}()
			requestWirelessUpdateNetworkWirelessSSIDSplashSettingsBillingFreeAccess = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsBillingFreeAccess{
				DurationInMinutes: int64ToIntPointer(durationInMinutes),
				Enabled:           enabled,
			}
		}
		prepaidAccessFastLoginEnabled := func() *bool {
			if !r.Billing.PrepaidAccessFastLoginEnabled.IsUnknown() && !r.Billing.PrepaidAccessFastLoginEnabled.IsNull() {
				return r.Billing.PrepaidAccessFastLoginEnabled.ValueBoolPointer()
			}
			return nil
		}()
		replyToEmailAddress := r.Billing.ReplyToEmailAddress.ValueString()
		requestWirelessUpdateNetworkWirelessSSIDSplashSettingsBilling = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsBilling{
			FreeAccess:                    requestWirelessUpdateNetworkWirelessSSIDSplashSettingsBillingFreeAccess,
			PrepaidAccessFastLoginEnabled: prepaidAccessFastLoginEnabled,
			ReplyToEmailAddress:           replyToEmailAddress,
		}
	}
	blockAllTrafficBeforeSignOn := new(bool)
	if !r.BlockAllTrafficBeforeSignOn.IsUnknown() && !r.BlockAllTrafficBeforeSignOn.IsNull() {
		*blockAllTrafficBeforeSignOn = r.BlockAllTrafficBeforeSignOn.ValueBool()
	} else {
		blockAllTrafficBeforeSignOn = nil
	}
	controllerDisconnectionBehavior := new(string)
	if !r.ControllerDisconnectionBehavior.IsUnknown() && !r.ControllerDisconnectionBehavior.IsNull() {
		*controllerDisconnectionBehavior = r.ControllerDisconnectionBehavior.ValueString()
	} else {
		controllerDisconnectionBehavior = &emptyString
	}
	var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsGuestSponsorship *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsGuestSponsorship
	if r.GuestSponsorship != nil {
		durationInMinutes := func() *int64 {
			if !r.GuestSponsorship.DurationInMinutes.IsUnknown() && !r.GuestSponsorship.DurationInMinutes.IsNull() {
				return r.GuestSponsorship.DurationInMinutes.ValueInt64Pointer()
			}
			return nil
		}()
		guestCanRequestTimeframe := func() *bool {
			if !r.GuestSponsorship.GuestCanRequestTimeframe.IsUnknown() && !r.GuestSponsorship.GuestCanRequestTimeframe.IsNull() {
				return r.GuestSponsorship.GuestCanRequestTimeframe.ValueBoolPointer()
			}
			return nil
		}()
		requestWirelessUpdateNetworkWirelessSSIDSplashSettingsGuestSponsorship = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsGuestSponsorship{
			DurationInMinutes:        int64ToIntPointer(durationInMinutes),
			GuestCanRequestTimeframe: guestCanRequestTimeframe,
		}
	}
	redirectURL := new(string)
	if !r.RedirectURL.IsUnknown() && !r.RedirectURL.IsNull() {
		*redirectURL = r.RedirectURL.ValueString()
	} else {
		redirectURL = &emptyString
	}
	var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment
	if r.SentryEnrollment != nil {
		var enforcedSystems []string = nil
		//Hoola aqui
		r.SentryEnrollment.EnforcedSystems.ElementsAs(ctx, &enforcedSystems, false)
		strength := r.SentryEnrollment.Strength.ValueString()
		var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork
		if r.SentryEnrollment.SystemsManagerNetwork != nil {
			iD := r.SentryEnrollment.SystemsManagerNetwork.ID.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork{
				ID: iD,
			}
		}
		requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment{
			EnforcedSystems:       enforcedSystems,
			Strength:              strength,
			SystemsManagerNetwork: requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork,
		}
	}
	var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImage *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImage
	if r.SplashImage != nil {
		extension := r.SplashImage.Extension.ValueString()
		var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImageImage *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImageImage
		if r.SplashImage.Image != nil {
			contents := r.SplashImage.Image.Contents.ValueString()
			format := r.SplashImage.Image.Format.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImageImage = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImageImage{
				Contents: contents,
				Format:   format,
			}
		}
		md5 := r.SplashImage.Md5.ValueString()
		requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImage = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImage{
			Extension: extension,
			Image:     requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImageImage,
			Md5:       md5,
		}
	}
	var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogo *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogo
	if r.SplashLogo != nil {
		extension := r.SplashLogo.Extension.ValueString()
		var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogoImage *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogoImage
		if r.SplashLogo.Image != nil {
			contents := r.SplashLogo.Image.Contents.ValueString()
			format := r.SplashLogo.Image.Format.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogoImage = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogoImage{
				Contents: contents,
				Format:   format,
			}
		}
		md5 := r.SplashLogo.Md5.ValueString()
		requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogo = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogo{
			Extension: extension,
			Image:     requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogoImage,
			Md5:       md5,
		}
	}
	var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFront *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFront
	if r.SplashPrepaidFront != nil {
		extension := r.SplashPrepaidFront.Extension.ValueString()
		var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFrontImage *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFrontImage
		if r.SplashPrepaidFront.Image != nil {
			contents := r.SplashPrepaidFront.Image.Contents.ValueString()
			format := r.SplashPrepaidFront.Image.Format.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFrontImage = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFrontImage{
				Contents: contents,
				Format:   format,
			}
		}
		md5 := r.SplashPrepaidFront.Md5.ValueString()
		requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFront = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFront{
			Extension: extension,
			Image:     requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFrontImage,
			Md5:       md5,
		}
	}
	splashTimeout := new(int64)
	if !r.SplashTimeout.IsUnknown() && !r.SplashTimeout.IsNull() {
		*splashTimeout = r.SplashTimeout.ValueInt64()
	} else {
		splashTimeout = nil
	}
	splashURL := new(string)
	if !r.SplashURL.IsUnknown() && !r.SplashURL.IsNull() {
		*splashURL = r.SplashURL.ValueString()
	} else {
		splashURL = &emptyString
	}
	themeID := new(string)
	if !r.ThemeID.IsUnknown() && !r.ThemeID.IsNull() {
		*themeID = r.ThemeID.ValueString()
	} else {
		themeID = &emptyString
	}
	useRedirectURL := new(bool)
	if !r.UseRedirectURL.IsUnknown() && !r.UseRedirectURL.IsNull() {
		*useRedirectURL = r.UseRedirectURL.ValueBool()
	} else {
		useRedirectURL = nil
	}
	useSplashURL := new(bool)
	if !r.UseSplashURL.IsUnknown() && !r.UseSplashURL.IsNull() {
		*useSplashURL = r.UseSplashURL.ValueBool()
	} else {
		useSplashURL = nil
	}
	welcomeMessage := new(string)
	if !r.WelcomeMessage.IsUnknown() && !r.WelcomeMessage.IsNull() {
		*welcomeMessage = r.WelcomeMessage.ValueString()
	} else {
		welcomeMessage = &emptyString
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettings{
		AllowSimultaneousLogins:         allowSimultaneousLogins,
		Billing:                         requestWirelessUpdateNetworkWirelessSSIDSplashSettingsBilling,
		BlockAllTrafficBeforeSignOn:     blockAllTrafficBeforeSignOn,
		ControllerDisconnectionBehavior: *controllerDisconnectionBehavior,
		GuestSponsorship:                requestWirelessUpdateNetworkWirelessSSIDSplashSettingsGuestSponsorship,
		RedirectURL:                     *redirectURL,
		SentryEnrollment:                requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment,
		SplashImage:                     requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImage,
		SplashLogo:                      requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogo,
		SplashPrepaidFront:              requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFront,
		SplashTimeout:                   int64ToIntPointer(splashTimeout),
		SplashURL:                       *splashURL,
		ThemeID:                         *themeID,
		UseRedirectURL:                  useRedirectURL,
		UseSplashURL:                    useSplashURL,
		WelcomeMessage:                  *welcomeMessage,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDSplashSettingsItemToBodyRs(state NetworksWirelessSSIDsSplashSettingsRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDSplashSettings, is_read bool) NetworksWirelessSSIDsSplashSettingsRs {
	itemState := NetworksWirelessSSIDsSplashSettingsRs{
		AllowSimultaneousLogins: func() types.Bool {
			if response.AllowSimultaneousLogins != nil {
				return types.BoolValue(*response.AllowSimultaneousLogins)
			}
			return types.Bool{}
		}(),
		Billing: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingRs {
			if response.Billing != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingRs{
					FreeAccess: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingFreeAccessRs {
						if response.Billing.FreeAccess != nil {
							return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingFreeAccessRs{
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
						return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingFreeAccessRs{}
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
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsBillingRs{}
		}(),
		BlockAllTrafficBeforeSignOn: func() types.Bool {
			if response.BlockAllTrafficBeforeSignOn != nil {
				return types.BoolValue(*response.BlockAllTrafficBeforeSignOn)
			}
			return types.Bool{}
		}(),
		ControllerDisconnectionBehavior: types.StringValue(response.ControllerDisconnectionBehavior),
		GuestSponsorship: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsGuestSponsorshipRs {
			if response.GuestSponsorship != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsGuestSponsorshipRs{
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
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsGuestSponsorshipRs{}
		}(),
		RedirectURL: types.StringValue(response.RedirectURL),
		SelfRegistration: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSelfRegistrationRs {
			if response.SelfRegistration != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSelfRegistrationRs{
					AuthorizationType: types.StringValue(response.SelfRegistration.AuthorizationType),
					Enabled: func() types.Bool {
						if response.SelfRegistration.Enabled != nil {
							return types.BoolValue(*response.SelfRegistration.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSelfRegistrationRs{}
		}(),
		SentryEnrollment: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentRs {
			if response.SentryEnrollment != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentRs{
					EnforcedSystems: StringSliceToSet(response.SentryEnrollment.EnforcedSystems),
					Strength:        types.StringValue(response.SentryEnrollment.Strength),
					SystemsManagerNetwork: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentSystemsManagerNetworkRs {
						if response.SentryEnrollment.SystemsManagerNetwork != nil {
							return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentSystemsManagerNetworkRs{
								ID: types.StringValue(response.SentryEnrollment.SystemsManagerNetwork.ID),
							}
						}
						return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentSystemsManagerNetworkRs{}
					}(),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentRs{}
		}(),
		SplashImage: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashImageRs {
			if response.SplashImage != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashImageRs{
					Extension: types.StringValue(response.SplashImage.Extension),
					Md5:       types.StringValue(response.SplashImage.Md5),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashImageRs{}
		}(),
		SplashLogo: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashLogoRs {
			if response.SplashLogo != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashLogoRs{
					Extension: types.StringValue(response.SplashLogo.Extension),
					Md5:       types.StringValue(response.SplashLogo.Md5),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashLogoRs{}
		}(),
		SplashPage: types.StringValue(response.SplashPage),
		SplashPrepaidFront: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashPrepaidFrontRs {
			if response.SplashPrepaidFront != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashPrepaidFrontRs{
					Extension: types.StringValue(response.SplashPrepaidFront.Extension),
					Md5:       types.StringValue(response.SplashPrepaidFront.Md5),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashPrepaidFrontRs{}
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
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsSplashSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsSplashSettingsRs)
}
