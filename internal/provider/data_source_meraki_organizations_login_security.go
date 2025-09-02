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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsLoginSecurityDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsLoginSecurityDataSource{}
)

func NewOrganizationsLoginSecurityDataSource() datasource.DataSource {
	return &OrganizationsLoginSecurityDataSource{}
}

type OrganizationsLoginSecurityDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsLoginSecurityDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsLoginSecurityDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_login_security"
}

func (d *OrganizationsLoginSecurityDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"account_lockout_attempts": schema.Int64Attribute{
						MarkdownDescription: `Number of consecutive failed login attempts after which users' accounts will be locked.`,
						Computed:            true,
					},
					"api_authentication": schema.SingleNestedAttribute{
						MarkdownDescription: `Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"ip_restrictions_for_keys": schema.SingleNestedAttribute{
								MarkdownDescription: `Details for API-only IP restrictions.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges.`,
										Computed:            true,
									},
									"ranges": schema.ListAttribute{
										MarkdownDescription: `List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.`,
										Computed:            true,
										ElementType:         types.StringType,
									},
								},
							},
						},
					},
					"enforce_account_lockout": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether users' Dashboard accounts will be locked out after a specified number of consecutive failed login attempts.`,
						Computed:            true,
					},
					"enforce_different_passwords": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether users, when setting a new password, are forced to choose a new password that is different from any past passwords.`,
						Computed:            true,
					},
					"enforce_idle_timeout": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether users will be logged out after being idle for the specified number of minutes.`,
						Computed:            true,
					},
					"enforce_login_ip_ranges": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether organization will restrict access to Dashboard (including the API) from certain IP addresses.`,
						Computed:            true,
					},
					"enforce_password_expiration": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether users are forced to change their password every X number of days.`,
						Computed:            true,
					},
					"enforce_strong_passwords": schema.BoolAttribute{
						MarkdownDescription: `Deprecated. This will always be 'true'.`,
						Computed:            true,
					},
					"enforce_two_factor_auth": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether users in this organization will be required to use an extra verification code when logging in to Dashboard. This code will be sent to their mobile phone via SMS, or can be generated by the authenticator application.`,
						Computed:            true,
					},
					"idle_timeout_minutes": schema.Int64Attribute{
						MarkdownDescription: `Number of minutes users can remain idle before being logged out of their accounts.`,
						Computed:            true,
					},
					"login_ip_ranges": schema.ListAttribute{
						MarkdownDescription: `List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"minimum_password_length": schema.Int64Attribute{
						MarkdownDescription: `The minimum number of characters required in admins' passwords.`,
						Computed:            true,
					},
					"num_different_passwords": schema.Int64Attribute{
						MarkdownDescription: `Number of recent passwords that new password must be distinct from.`,
						Computed:            true,
					},
					"password_expiration_days": schema.Int64Attribute{
						MarkdownDescription: `Number of days after which users will be forced to change their password.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *OrganizationsLoginSecurityDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsLoginSecurity OrganizationsLoginSecurity
	diags := req.Config.Get(ctx, &organizationsLoginSecurity)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationLoginSecurity")
		vvOrganizationID := organizationsLoginSecurity.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationLoginSecurity(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationLoginSecurity",
				err.Error(),
			)
			return
		}

		organizationsLoginSecurity = ResponseOrganizationsGetOrganizationLoginSecurityItemToBody(organizationsLoginSecurity, response1)
		diags = resp.State.Set(ctx, &organizationsLoginSecurity)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsLoginSecurity struct {
	OrganizationID types.String                                       `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsGetOrganizationLoginSecurity `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationLoginSecurity struct {
	AccountLockoutAttempts    types.Int64                                                         `tfsdk:"account_lockout_attempts"`
	APIAuthentication         *ResponseOrganizationsGetOrganizationLoginSecurityApiAuthentication `tfsdk:"api_authentication"`
	EnforceAccountLockout     types.Bool                                                          `tfsdk:"enforce_account_lockout"`
	EnforceDifferentPasswords types.Bool                                                          `tfsdk:"enforce_different_passwords"`
	EnforceIDleTimeout        types.Bool                                                          `tfsdk:"enforce_idle_timeout"`
	EnforceLoginIPRanges      types.Bool                                                          `tfsdk:"enforce_login_ip_ranges"`
	EnforcePasswordExpiration types.Bool                                                          `tfsdk:"enforce_password_expiration"`
	EnforceStrongPasswords    types.Bool                                                          `tfsdk:"enforce_strong_passwords"`
	EnforceTwoFactorAuth      types.Bool                                                          `tfsdk:"enforce_two_factor_auth"`
	IDleTimeoutMinutes        types.Int64                                                         `tfsdk:"idle_timeout_minutes"`
	LoginIPRanges             types.List                                                          `tfsdk:"login_ip_ranges"`
	MinimumPasswordLength     types.Int64                                                         `tfsdk:"minimum_password_length"`
	NumDifferentPasswords     types.Int64                                                         `tfsdk:"num_different_passwords"`
	PasswordExpirationDays    types.Int64                                                         `tfsdk:"password_expiration_days"`
}

type ResponseOrganizationsGetOrganizationLoginSecurityApiAuthentication struct {
	IPRestrictionsForKeys *ResponseOrganizationsGetOrganizationLoginSecurityApiAuthenticationIpRestrictionsForKeys `tfsdk:"ip_restrictions_for_keys"`
}

type ResponseOrganizationsGetOrganizationLoginSecurityApiAuthenticationIpRestrictionsForKeys struct {
	Enabled types.Bool `tfsdk:"enabled"`
	Ranges  types.List `tfsdk:"ranges"`
}

// ToBody
func ResponseOrganizationsGetOrganizationLoginSecurityItemToBody(state OrganizationsLoginSecurity, response *merakigosdk.ResponseOrganizationsGetOrganizationLoginSecurity) OrganizationsLoginSecurity {
	itemState := ResponseOrganizationsGetOrganizationLoginSecurity{
		AccountLockoutAttempts: func() types.Int64 {
			if response.AccountLockoutAttempts != nil {
				return types.Int64Value(int64(*response.AccountLockoutAttempts))
			}
			return types.Int64{}
		}(),
		APIAuthentication: func() *ResponseOrganizationsGetOrganizationLoginSecurityApiAuthentication {
			if response.APIAuthentication != nil {
				return &ResponseOrganizationsGetOrganizationLoginSecurityApiAuthentication{
					IPRestrictionsForKeys: func() *ResponseOrganizationsGetOrganizationLoginSecurityApiAuthenticationIpRestrictionsForKeys {
						if response.APIAuthentication.IPRestrictionsForKeys != nil {
							return &ResponseOrganizationsGetOrganizationLoginSecurityApiAuthenticationIpRestrictionsForKeys{
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
	state.Item = &itemState
	return state
}
