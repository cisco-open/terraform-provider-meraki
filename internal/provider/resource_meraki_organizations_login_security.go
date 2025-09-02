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
	"strconv"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsLoginSecurityResource{}
	_ resource.ResourceWithConfigure = &OrganizationsLoginSecurityResource{}
)

func NewOrganizationsLoginSecurityResource() resource.Resource {
	return &OrganizationsLoginSecurityResource{}
}

type OrganizationsLoginSecurityResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsLoginSecurityResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsLoginSecurityResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_login_security"
}

func (r *OrganizationsLoginSecurityResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_lockout_attempts": schema.Int64Attribute{
				MarkdownDescription: `Number of consecutive failed login attempts after which users' accounts will be locked.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"api_authentication": schema.SingleNestedAttribute{
				MarkdownDescription: `Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"ip_restrictions_for_keys": schema.SingleNestedAttribute{
						MarkdownDescription: `Details for API-only IP restrictions.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges.`,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"ranges": schema.ListAttribute{
								MarkdownDescription: `List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.`,
								Optional:            true,
								PlanModifiers: []planmodifier.List{
									listplanmodifier.UseStateForUnknown(),
								},

								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"enforce_account_lockout": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether users' Dashboard accounts will be locked out after a specified number of consecutive failed login attempts.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"enforce_different_passwords": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether users, when setting a new password, are forced to choose a new password that is different from any past passwords.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"enforce_idle_timeout": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether users will be logged out after being idle for the specified number of minutes.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"enforce_login_ip_ranges": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether organization will restrict access to Dashboard (including the API) from certain IP addresses.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"enforce_password_expiration": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether users are forced to change their password every X number of days.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"enforce_strong_passwords": schema.BoolAttribute{
				MarkdownDescription: `Deprecated. This will always be 'true'.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"enforce_two_factor_auth": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating whether users in this organization will be required to use an extra verification code when logging in to Dashboard. This code will be sent to their mobile phone via SMS, or can be generated by the authenticator application.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"idle_timeout_minutes": schema.Int64Attribute{
				MarkdownDescription: `Number of minutes users can remain idle before being logged out of their accounts.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"login_ip_ranges": schema.ListAttribute{
				MarkdownDescription: `List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
				Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
			},
			"minimum_password_length": schema.Int64Attribute{
				MarkdownDescription: `The minimum number of characters required in admins' passwords.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"num_different_passwords": schema.Int64Attribute{
				MarkdownDescription: `Number of recent passwords that new password must be distinct from.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"password_expiration_days": schema.Int64Attribute{
				MarkdownDescription: `Number of days after which users will be forced to change their password.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *OrganizationsLoginSecurityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsLoginSecurityRs

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
	vvOrganizationID := data.OrganizationID.ValueString()
	//Has Item and not has items

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationLoginSecurity(vvOrganizationID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationLoginSecurity",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationLoginSecurity",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *OrganizationsLoginSecurityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsLoginSecurityRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvOrganizationID := data.OrganizationID.ValueString()
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationLoginSecurity(vvOrganizationID)
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
				"Failure when executing GetOrganizationLoginSecurity",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationLoginSecurity",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationLoginSecurityItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *OrganizationsLoginSecurityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), req.ID)...)
}

func (r *OrganizationsLoginSecurityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan OrganizationsLoginSecurityRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvOrganizationID := plan.OrganizationID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationLoginSecurity(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationLoginSecurity",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationLoginSecurity",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *OrganizationsLoginSecurityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting OrganizationsLoginSecurity", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsLoginSecurityRs struct {
	OrganizationID            types.String                                                          `tfsdk:"organization_id"`
	AccountLockoutAttempts    types.Int64                                                           `tfsdk:"account_lockout_attempts"`
	APIAuthentication         *ResponseOrganizationsGetOrganizationLoginSecurityApiAuthenticationRs `tfsdk:"api_authentication"`
	EnforceAccountLockout     types.Bool                                                            `tfsdk:"enforce_account_lockout"`
	EnforceDifferentPasswords types.Bool                                                            `tfsdk:"enforce_different_passwords"`
	EnforceIDleTimeout        types.Bool                                                            `tfsdk:"enforce_idle_timeout"`
	EnforceLoginIPRanges      types.Bool                                                            `tfsdk:"enforce_login_ip_ranges"`
	EnforcePasswordExpiration types.Bool                                                            `tfsdk:"enforce_password_expiration"`
	EnforceStrongPasswords    types.Bool                                                            `tfsdk:"enforce_strong_passwords"`
	EnforceTwoFactorAuth      types.Bool                                                            `tfsdk:"enforce_two_factor_auth"`
	IDleTimeoutMinutes        types.Int64                                                           `tfsdk:"idle_timeout_minutes"`
	LoginIPRanges             types.List                                                            `tfsdk:"login_ip_ranges"`
	MinimumPasswordLength     types.Int64                                                           `tfsdk:"minimum_password_length"`
	NumDifferentPasswords     types.Int64                                                           `tfsdk:"num_different_passwords"`
	PasswordExpirationDays    types.Int64                                                           `tfsdk:"password_expiration_days"`
}

type ResponseOrganizationsGetOrganizationLoginSecurityApiAuthenticationRs struct {
	IPRestrictionsForKeys *ResponseOrganizationsGetOrganizationLoginSecurityApiAuthenticationIpRestrictionsForKeysRs `tfsdk:"ip_restrictions_for_keys"`
}

type ResponseOrganizationsGetOrganizationLoginSecurityApiAuthenticationIpRestrictionsForKeysRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
	Ranges  types.List `tfsdk:"ranges"`
}

// FromBody
func (r *OrganizationsLoginSecurityRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationLoginSecurity {
	accountLockoutAttempts := new(int64)
	if !r.AccountLockoutAttempts.IsUnknown() && !r.AccountLockoutAttempts.IsNull() {
		*accountLockoutAttempts = r.AccountLockoutAttempts.ValueInt64()
	} else {
		accountLockoutAttempts = nil
	}
	var requestOrganizationsUpdateOrganizationLoginSecurityAPIAuthentication *merakigosdk.RequestOrganizationsUpdateOrganizationLoginSecurityAPIAuthentication

	if r.APIAuthentication != nil {
		var requestOrganizationsUpdateOrganizationLoginSecurityAPIAuthenticationIPRestrictionsForKeys *merakigosdk.RequestOrganizationsUpdateOrganizationLoginSecurityAPIAuthenticationIPRestrictionsForKeys

		if r.APIAuthentication.IPRestrictionsForKeys != nil {
			enabled := func() *bool {
				if !r.APIAuthentication.IPRestrictionsForKeys.Enabled.IsUnknown() && !r.APIAuthentication.IPRestrictionsForKeys.Enabled.IsNull() {
					return r.APIAuthentication.IPRestrictionsForKeys.Enabled.ValueBoolPointer()
				}
				return nil
			}()

			var ranges []string = nil
			r.APIAuthentication.IPRestrictionsForKeys.Ranges.ElementsAs(ctx, &ranges, false)
			requestOrganizationsUpdateOrganizationLoginSecurityAPIAuthenticationIPRestrictionsForKeys = &merakigosdk.RequestOrganizationsUpdateOrganizationLoginSecurityAPIAuthenticationIPRestrictionsForKeys{
				Enabled: enabled,
				Ranges:  ranges,
			}
			//[debug] Is Array: False
		}
		requestOrganizationsUpdateOrganizationLoginSecurityAPIAuthentication = &merakigosdk.RequestOrganizationsUpdateOrganizationLoginSecurityAPIAuthentication{
			IPRestrictionsForKeys: requestOrganizationsUpdateOrganizationLoginSecurityAPIAuthenticationIPRestrictionsForKeys,
		}
		//[debug] Is Array: False
	}
	enforceAccountLockout := new(bool)
	if !r.EnforceAccountLockout.IsUnknown() && !r.EnforceAccountLockout.IsNull() {
		*enforceAccountLockout = r.EnforceAccountLockout.ValueBool()
	} else {
		enforceAccountLockout = nil
	}
	enforceDifferentPasswords := new(bool)
	if !r.EnforceDifferentPasswords.IsUnknown() && !r.EnforceDifferentPasswords.IsNull() {
		*enforceDifferentPasswords = r.EnforceDifferentPasswords.ValueBool()
	} else {
		enforceDifferentPasswords = nil
	}
	enforceIDleTimeout := new(bool)
	if !r.EnforceIDleTimeout.IsUnknown() && !r.EnforceIDleTimeout.IsNull() {
		*enforceIDleTimeout = r.EnforceIDleTimeout.ValueBool()
	} else {
		enforceIDleTimeout = nil
	}
	enforceLoginIPRanges := new(bool)
	if !r.EnforceLoginIPRanges.IsUnknown() && !r.EnforceLoginIPRanges.IsNull() {
		*enforceLoginIPRanges = r.EnforceLoginIPRanges.ValueBool()
	} else {
		enforceLoginIPRanges = nil
	}
	enforcePasswordExpiration := new(bool)
	if !r.EnforcePasswordExpiration.IsUnknown() && !r.EnforcePasswordExpiration.IsNull() {
		*enforcePasswordExpiration = r.EnforcePasswordExpiration.ValueBool()
	} else {
		enforcePasswordExpiration = nil
	}
	enforceStrongPasswords := new(bool)
	if !r.EnforceStrongPasswords.IsUnknown() && !r.EnforceStrongPasswords.IsNull() {
		*enforceStrongPasswords = r.EnforceStrongPasswords.ValueBool()
	} else {
		enforceStrongPasswords = nil
	}
	enforceTwoFactorAuth := new(bool)
	if !r.EnforceTwoFactorAuth.IsUnknown() && !r.EnforceTwoFactorAuth.IsNull() {
		*enforceTwoFactorAuth = r.EnforceTwoFactorAuth.ValueBool()
	} else {
		enforceTwoFactorAuth = nil
	}
	iDleTimeoutMinutes := new(int64)
	if !r.IDleTimeoutMinutes.IsUnknown() && !r.IDleTimeoutMinutes.IsNull() {
		*iDleTimeoutMinutes = r.IDleTimeoutMinutes.ValueInt64()
	} else {
		iDleTimeoutMinutes = nil
	}
	var loginIPRanges []string = nil
	r.LoginIPRanges.ElementsAs(ctx, &loginIPRanges, false)
	minimumPasswordLength := new(int64)
	if !r.MinimumPasswordLength.IsUnknown() && !r.MinimumPasswordLength.IsNull() {
		*minimumPasswordLength = r.MinimumPasswordLength.ValueInt64()
	} else {
		minimumPasswordLength = nil
	}
	numDifferentPasswords := new(int64)
	if !r.NumDifferentPasswords.IsUnknown() && !r.NumDifferentPasswords.IsNull() {
		*numDifferentPasswords = r.NumDifferentPasswords.ValueInt64()
	} else {
		numDifferentPasswords = nil
	}
	passwordExpirationDays := new(int64)
	if !r.PasswordExpirationDays.IsUnknown() && !r.PasswordExpirationDays.IsNull() {
		*passwordExpirationDays = r.PasswordExpirationDays.ValueInt64()
	} else {
		passwordExpirationDays = nil
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationLoginSecurity{
		AccountLockoutAttempts:    int64ToIntPointer(accountLockoutAttempts),
		APIAuthentication:         requestOrganizationsUpdateOrganizationLoginSecurityAPIAuthentication,
		EnforceAccountLockout:     enforceAccountLockout,
		EnforceDifferentPasswords: enforceDifferentPasswords,
		EnforceIDleTimeout:        enforceIDleTimeout,
		EnforceLoginIPRanges:      enforceLoginIPRanges,
		EnforcePasswordExpiration: enforcePasswordExpiration,
		EnforceStrongPasswords:    enforceStrongPasswords,
		EnforceTwoFactorAuth:      enforceTwoFactorAuth,
		IDleTimeoutMinutes:        int64ToIntPointer(iDleTimeoutMinutes),
		LoginIPRanges:             loginIPRanges,
		MinimumPasswordLength:     int64ToIntPointer(minimumPasswordLength),
		NumDifferentPasswords:     int64ToIntPointer(numDifferentPasswords),
		PasswordExpirationDays:    int64ToIntPointer(passwordExpirationDays),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationLoginSecurityItemToBodyRs(state OrganizationsLoginSecurityRs, response *merakigosdk.ResponseOrganizationsGetOrganizationLoginSecurity, is_read bool) OrganizationsLoginSecurityRs {
	itemState := OrganizationsLoginSecurityRs{
		AccountLockoutAttempts: func() types.Int64 {
			if response.AccountLockoutAttempts != nil {
				return types.Int64Value(int64(*response.AccountLockoutAttempts))
			}
			return types.Int64{}
		}(),
		APIAuthentication: func() *ResponseOrganizationsGetOrganizationLoginSecurityApiAuthenticationRs {
			if response.APIAuthentication != nil {
				return &ResponseOrganizationsGetOrganizationLoginSecurityApiAuthenticationRs{
					IPRestrictionsForKeys: func() *ResponseOrganizationsGetOrganizationLoginSecurityApiAuthenticationIpRestrictionsForKeysRs {
						if response.APIAuthentication.IPRestrictionsForKeys != nil {
							return &ResponseOrganizationsGetOrganizationLoginSecurityApiAuthenticationIpRestrictionsForKeysRs{
								Enabled: func() types.Bool {
									if response.APIAuthentication.IPRestrictionsForKeys.Enabled != nil {
										return types.BoolValue(*response.APIAuthentication.IPRestrictionsForKeys.Enabled)
									}
									return types.Bool{}
								}(),
								Ranges: StringSliceToList(response.APIAuthentication.IPRestrictionsForKeys.Ranges),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		EnforceAccountLockout: func() types.Bool {
			if response.EnforceAccountLockout != nil {
				return types.BoolValue(*response.EnforceAccountLockout)
			}
			return types.Bool{}
		}(),
		EnforceDifferentPasswords: func() types.Bool {
			if response.EnforceDifferentPasswords != nil {
				return types.BoolValue(*response.EnforceDifferentPasswords)
			}
			return types.Bool{}
		}(),
		EnforceIDleTimeout: func() types.Bool {
			if response.EnforceIDleTimeout != nil {
				return types.BoolValue(*response.EnforceIDleTimeout)
			}
			return types.Bool{}
		}(),
		EnforceLoginIPRanges: func() types.Bool {
			if response.EnforceLoginIPRanges != nil {
				return types.BoolValue(*response.EnforceLoginIPRanges)
			}
			return types.Bool{}
		}(),
		EnforcePasswordExpiration: func() types.Bool {
			if response.EnforcePasswordExpiration != nil {
				return types.BoolValue(*response.EnforcePasswordExpiration)
			}
			return types.Bool{}
		}(),
		EnforceStrongPasswords: func() types.Bool {
			if response.EnforceStrongPasswords != nil {
				return types.BoolValue(*response.EnforceStrongPasswords)
			}
			return types.Bool{}
		}(),
		EnforceTwoFactorAuth: func() types.Bool {
			if response.EnforceTwoFactorAuth != nil {
				return types.BoolValue(*response.EnforceTwoFactorAuth)
			}
			return types.Bool{}
		}(),
		IDleTimeoutMinutes: func() types.Int64 {
			if response.IDleTimeoutMinutes != nil {
				return types.Int64Value(int64(*response.IDleTimeoutMinutes))
			}
			return types.Int64{}
		}(),
		LoginIPRanges: StringSliceToList(response.LoginIPRanges),
		MinimumPasswordLength: func() types.Int64 {
			if response.MinimumPasswordLength != nil {
				return types.Int64Value(int64(*response.MinimumPasswordLength))
			}
			return types.Int64{}
		}(),
		NumDifferentPasswords: func() types.Int64 {
			if response.NumDifferentPasswords != nil {
				return types.Int64Value(int64(*response.NumDifferentPasswords))
			}
			return types.Int64{}
		}(),
		PasswordExpirationDays: func() types.Int64 {
			if response.PasswordExpirationDays != nil {
				return types.Int64Value(int64(*response.PasswordExpirationDays))
			}
			return types.Int64{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsLoginSecurityRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsLoginSecurityRs)
}
