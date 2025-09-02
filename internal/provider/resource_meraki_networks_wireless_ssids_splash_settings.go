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
	"strconv"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
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
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"billing": schema.SingleNestedAttribute{
				MarkdownDescription: `Details associated with billing splash`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"free_access": schema.SingleNestedAttribute{
						MarkdownDescription: `Details associated with a free access plan with limits`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"duration_in_minutes": schema.Int64Attribute{
								MarkdownDescription: `How long a device can use a network for free.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
							},
							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Whether or not free access is enabled.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
						},
					},
					"prepaid_access_fast_login_enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether or not billing uses the fast login prepaid access option.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"reply_to_email_address": schema.StringAttribute{
						MarkdownDescription: `The email address that reeceives replies from clients`,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"block_all_traffic_before_sign_on": schema.BoolAttribute{
				MarkdownDescription: `How restricted allowing traffic should be. If true, all traffic types are blocked until the splash page is acknowledged. If false, all non-HTTP traffic is allowed before the splash page is acknowledged.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"controller_disconnection_behavior": schema.StringAttribute{
				MarkdownDescription: `How login attempts should be handled when the controller is unreachable.
                                  Allowed values: [default,open,restricted]`,
				Optional: true,
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
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"duration_in_minutes": schema.Int64Attribute{
						MarkdownDescription: `Duration in minutes of sponsored guest authorization.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"guest_can_request_timeframe": schema.BoolAttribute{
						MarkdownDescription: `Whether or not guests can specify how much time they are requesting.`,
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
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"self_registration": schema.SingleNestedAttribute{
				MarkdownDescription: `Self-registration for splash with Meraki authentication.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"authorization_type": schema.StringAttribute{
						MarkdownDescription: `How created user accounts should be authorized.
                                        Allowed values: [admin,auto,self_email]`,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"admin",
								"auto",
								"self_email",
							),
						},
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether or not to allow users to create their own account on the network.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"sentry_enrollment": schema.SingleNestedAttribute{
				MarkdownDescription: `Systems Manager sentry enrollment splash settings.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enforced_systems": schema.ListAttribute{
						MarkdownDescription: `The system types that the Sentry enforces.`,
						Optional:            true,
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
					"strength": schema.StringAttribute{
						MarkdownDescription: `The strength of the enforcement of selected system types.
                                        Allowed values: [click-through,focused,strict]`,
						Optional: true,
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
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `The network ID of the Systems Manager network.`,
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
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"extension": schema.StringAttribute{
						MarkdownDescription: `The extension of the image file.`,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"image": schema.SingleNestedAttribute{
						MarkdownDescription: `Properties for setting a new image.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"contents": schema.StringAttribute{
								MarkdownDescription: `The file contents (a base 64 encoded string) of your new image.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"format": schema.StringAttribute{
								MarkdownDescription: `The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.
                                              Allowed values: [gif,jpg,png]`,
								Optional: true,
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
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"splash_logo": schema.SingleNestedAttribute{
				MarkdownDescription: `The logo used in the splash page.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"extension": schema.StringAttribute{
						MarkdownDescription: `The extension of the logo file.`,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"image": schema.SingleNestedAttribute{
						MarkdownDescription: `Properties for setting a new image.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"contents": schema.StringAttribute{
								MarkdownDescription: `The file contents (a base 64 encoded string) of your new logo.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"format": schema.StringAttribute{
								MarkdownDescription: `The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.
                                              Allowed values: [gif,jpg,png]`,
								Optional: true,
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
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"extension": schema.StringAttribute{
						MarkdownDescription: `The extension of the prepaid front image file.`,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"image": schema.SingleNestedAttribute{
						MarkdownDescription: `Properties for setting a new image.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"contents": schema.StringAttribute{
								MarkdownDescription: `The file contents (a base 64 encoded string) of your new prepaid front.`,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
							},
							"format": schema.StringAttribute{
								MarkdownDescription: `The format of the encoded contents. Supported formats are 'png', 'gif', and jpg'.
                                              Allowed values: [gif,jpg,png]`,
								Optional: true,
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
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"splash_timeout": schema.Int64Attribute{
				MarkdownDescription: `Splash timeout in minutes.
                                  Allowed values: [30,60,120,240,480,720,1080,1440,2880,5760,7200,10080,20160,43200,86400,129600]`,
				Optional: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"splash_url": schema.StringAttribute{
				MarkdownDescription: `The custom splash URL of the click-through splash page.`,
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
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"use_redirect_url": schema.BoolAttribute{
				MarkdownDescription: `The Boolean indicating whether the the user will be redirected to the custom redirect URL after the splash page.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"use_splash_url": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether the users will be redirected to the custom splash url`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"welcome_message": schema.StringAttribute{
				MarkdownDescription: `The welcome message for the users on the splash page.`,
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
	// Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	vvNumber := data.Number.ValueString()
	//Has Item and not has items

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDSplashSettings(vvNetworkID, vvNumber, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDSplashSettings",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDSplashSettings",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksWirelessSSIDsSplashSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsSplashSettingsRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
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
				restyRespGet.String(),
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksWirelessSSIDsSplashSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: networkId,number. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
}

func (r *NetworksWirelessSSIDsSplashSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksWirelessSSIDsSplashSettingsRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	vvNumber := plan.Number.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDSplashSettings(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDSplashSettings",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDSplashSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
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
	EnforcedSystems       types.List                                                                                   `tfsdk:"enforced_systems"`
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
			//[debug] Is Array: False
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
		//[debug] Is Array: False
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
		//[debug] Is Array: False
	}
	redirectURL := new(string)
	if !r.RedirectURL.IsUnknown() && !r.RedirectURL.IsNull() {
		*redirectURL = r.RedirectURL.ValueString()
	} else {
		redirectURL = &emptyString
	}
	var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSelfRegistration *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSelfRegistration

	if r.SelfRegistration != nil {
		authorizationType := r.SelfRegistration.AuthorizationType.ValueString()
		enabled := func() *bool {
			if !r.SelfRegistration.Enabled.IsUnknown() && !r.SelfRegistration.Enabled.IsNull() {
				return r.SelfRegistration.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSelfRegistration = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSelfRegistration{
			AuthorizationType: authorizationType,
			Enabled:           enabled,
		}
		//[debug] Is Array: False
	}
	var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment

	if r.SentryEnrollment != nil {

		var enforcedSystems []string = nil
		r.SentryEnrollment.EnforcedSystems.ElementsAs(ctx, &enforcedSystems, false)
		strength := r.SentryEnrollment.Strength.ValueString()
		var requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork

		if r.SentryEnrollment.SystemsManagerNetwork != nil {
			id := r.SentryEnrollment.SystemsManagerNetwork.ID.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork{
				ID: id,
			}
			//[debug] Is Array: False
		}
		requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollment{
			EnforcedSystems:       enforcedSystems,
			Strength:              strength,
			SystemsManagerNetwork: requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSentryEnrollmentSystemsManagerNetwork,
		}
		//[debug] Is Array: False
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
			//[debug] Is Array: False
		}
		md5 := r.SplashImage.Md5.ValueString()
		requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImage = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImage{
			Extension: extension,
			Image:     requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashImageImage,
			Md5:       md5,
		}
		//[debug] Is Array: False
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
			//[debug] Is Array: False
		}
		md5 := r.SplashLogo.Md5.ValueString()
		requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogo = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogo{
			Extension: extension,
			Image:     requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashLogoImage,
			Md5:       md5,
		}
		//[debug] Is Array: False
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
			//[debug] Is Array: False
		}
		md5 := r.SplashPrepaidFront.Md5.ValueString()
		requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFront = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFront{
			Extension: extension,
			Image:     requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSplashPrepaidFrontImage,
			Md5:       md5,
		}
		//[debug] Is Array: False
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
		SelfRegistration:                requestWirelessUpdateNetworkWirelessSSIDSplashSettingsSelfRegistration,
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
						return nil
					}(),
					PrepaidAccessFastLoginEnabled: func() types.Bool {
						if response.Billing.PrepaidAccessFastLoginEnabled != nil {
							return types.BoolValue(*response.Billing.PrepaidAccessFastLoginEnabled)
						}
						return types.Bool{}
					}(),
					ReplyToEmailAddress: func() types.String {
						if response.Billing.ReplyToEmailAddress != "" {
							return types.StringValue(response.Billing.ReplyToEmailAddress)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		BlockAllTrafficBeforeSignOn: func() types.Bool {
			if response.BlockAllTrafficBeforeSignOn != nil {
				return types.BoolValue(*response.BlockAllTrafficBeforeSignOn)
			}
			return types.Bool{}
		}(),
		ControllerDisconnectionBehavior: func() types.String {
			if response.ControllerDisconnectionBehavior != "" {
				return types.StringValue(response.ControllerDisconnectionBehavior)
			}
			return types.String{}
		}(),
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
			return nil
		}(),
		RedirectURL: func() types.String {
			if response.RedirectURL != "" {
				return types.StringValue(response.RedirectURL)
			}
			return types.String{}
		}(),
		SelfRegistration: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSelfRegistrationRs {
			if response.SelfRegistration != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSelfRegistrationRs{
					AuthorizationType: func() types.String {
						if response.SelfRegistration.AuthorizationType != "" {
							return types.StringValue(response.SelfRegistration.AuthorizationType)
						}
						return types.String{}
					}(),
					Enabled: func() types.Bool {
						if response.SelfRegistration.Enabled != nil {
							return types.BoolValue(*response.SelfRegistration.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		SentryEnrollment: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentRs {
			if response.SentryEnrollment != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentRs{
					EnforcedSystems: StringSliceToList(response.SentryEnrollment.EnforcedSystems),
					Strength: func() types.String {
						if response.SentryEnrollment.Strength != "" {
							return types.StringValue(response.SentryEnrollment.Strength)
						}
						return types.String{}
					}(),
					SystemsManagerNetwork: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentSystemsManagerNetworkRs {
						if response.SentryEnrollment.SystemsManagerNetwork != nil {
							return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSentryEnrollmentSystemsManagerNetworkRs{
								ID: func() types.String {
									if response.SentryEnrollment.SystemsManagerNetwork.ID != "" {
										return types.StringValue(response.SentryEnrollment.SystemsManagerNetwork.ID)
									}
									return types.String{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		SplashImage: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashImageRs {
			if response.SplashImage != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashImageRs{
					Extension: func() types.String {
						if response.SplashImage.Extension != "" {
							return types.StringValue(response.SplashImage.Extension)
						}
						return types.String{}
					}(),
					Md5: func() types.String {
						if response.SplashImage.Md5 != "" {
							return types.StringValue(response.SplashImage.Md5)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		SplashLogo: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashLogoRs {
			if response.SplashLogo != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashLogoRs{
					Extension: func() types.String {
						if response.SplashLogo.Extension != "" {
							return types.StringValue(response.SplashLogo.Extension)
						}
						return types.String{}
					}(),
					Md5: func() types.String {
						if response.SplashLogo.Md5 != "" {
							return types.StringValue(response.SplashLogo.Md5)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		SplashPage: func() types.String {
			if response.SplashPage != "" {
				return types.StringValue(response.SplashPage)
			}
			return types.String{}
		}(),
		SplashPrepaidFront: func() *ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashPrepaidFrontRs {
			if response.SplashPrepaidFront != nil {
				return &ResponseWirelessGetNetworkWirelessSsidSplashSettingsSplashPrepaidFrontRs{
					Extension: func() types.String {
						if response.SplashPrepaidFront.Extension != "" {
							return types.StringValue(response.SplashPrepaidFront.Extension)
						}
						return types.String{}
					}(),
					Md5: func() types.String {
						if response.SplashPrepaidFront.Md5 != "" {
							return types.StringValue(response.SplashPrepaidFront.Md5)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		SplashTimeout: func() types.Int64 {
			if response.SplashTimeout != nil {
				return types.Int64Value(int64(*response.SplashTimeout))
			}
			return types.Int64{}
		}(),
		SplashURL: func() types.String {
			if response.SplashURL != "" {
				return types.StringValue(response.SplashURL)
			}
			return types.String{}
		}(),
		SSIDNumber: func() types.Int64 {
			if response.SSIDNumber != nil {
				return types.Int64Value(int64(*response.SSIDNumber))
			}
			return types.Int64{}
		}(),
		ThemeID: func() types.String {
			if response.ThemeID != "" {
				return types.StringValue(response.ThemeID)
			}
			return types.String{}
		}(),
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
		WelcomeMessage: func() types.String {
			if response.WelcomeMessage != "" {
				return types.StringValue(response.WelcomeMessage)
			}
			return types.String{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsSplashSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsSplashSettingsRs)
}
