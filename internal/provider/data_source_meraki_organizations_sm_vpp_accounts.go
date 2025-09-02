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
	_ datasource.DataSource              = &OrganizationsSmVppAccountsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSmVppAccountsDataSource{}
)

func NewOrganizationsSmVppAccountsDataSource() datasource.DataSource {
	return &OrganizationsSmVppAccountsDataSource{}
}

type OrganizationsSmVppAccountsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSmVppAccountsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSmVppAccountsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_sm_vpp_accounts"
}

func (d *OrganizationsSmVppAccountsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"vpp_account_id": schema.StringAttribute{
				MarkdownDescription: `vppAccountId path parameter. Vpp account ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"allowed_admins": schema.StringAttribute{
						MarkdownDescription: `The allowed admins for the VPP account`,
						Computed:            true,
					},
					"assignable_network_ids": schema.ListAttribute{
						MarkdownDescription: `The network IDs of the assignable networks for the VPP account`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"assignable_networks": schema.StringAttribute{
						MarkdownDescription: `The assignable networks for the VPP account`,
						Computed:            true,
					},
					"content_token": schema.StringAttribute{
						MarkdownDescription: `The VPP service token`,
						Computed:            true,
					},
					"email": schema.StringAttribute{
						MarkdownDescription: `The email address associated with the VPP account`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `The id of the VPP Account`,
						Computed:            true,
					},
					"last_force_synced_at": schema.StringAttribute{
						MarkdownDescription: `The last time the VPP account was force synced`,
						Computed:            true,
					},
					"last_synced_at": schema.StringAttribute{
						MarkdownDescription: `The last time the VPP account was synced`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the VPP account`,
						Computed:            true,
					},
					"network_id_admins": schema.StringAttribute{
						MarkdownDescription: `The network IDs of the admins for the VPP account`,
						Computed:            true,
					},
					"parsed_token": schema.SingleNestedAttribute{
						MarkdownDescription: `The parsed VPP service token`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"expires_at": schema.StringAttribute{
								MarkdownDescription: `The expiration time of the token`,
								Computed:            true,
							},
							"hashed_token": schema.StringAttribute{
								MarkdownDescription: `The hashed token`,
								Computed:            true,
							},
							"org_name": schema.StringAttribute{
								MarkdownDescription: `The organization name`,
								Computed:            true,
							},
						},
					},
					"vpp_account_id": schema.StringAttribute{
						MarkdownDescription: `The id of the VPP Account`,
						Computed:            true,
					},
					"vpp_location_id": schema.StringAttribute{
						MarkdownDescription: `The VPP location ID`,
						Computed:            true,
					},
					"vpp_location_name": schema.StringAttribute{
						MarkdownDescription: `The VPP location name`,
						Computed:            true,
					},
					"vpp_service_token": schema.StringAttribute{
						MarkdownDescription: `The VPP Account's Service Token`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetOrganizationSmVppAccounts`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"allowed_admins": schema.StringAttribute{
							MarkdownDescription: `The allowed admins for the VPP account`,
							Computed:            true,
						},
						"assignable_network_ids": schema.ListAttribute{
							MarkdownDescription: `The network IDs of the assignable networks for the VPP account`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"assignable_networks": schema.StringAttribute{
							MarkdownDescription: `The assignable networks for the VPP account`,
							Computed:            true,
						},
						"content_token": schema.StringAttribute{
							MarkdownDescription: `The VPP service token`,
							Computed:            true,
						},
						"email": schema.StringAttribute{
							MarkdownDescription: `The email address associated with the VPP account`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `The id of the VPP Account`,
							Computed:            true,
						},
						"last_force_synced_at": schema.StringAttribute{
							MarkdownDescription: `The last time the VPP account was force synced`,
							Computed:            true,
						},
						"last_synced_at": schema.StringAttribute{
							MarkdownDescription: `The last time the VPP account was synced`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the VPP account`,
							Computed:            true,
						},
						"network_id_admins": schema.StringAttribute{
							MarkdownDescription: `The network IDs of the admins for the VPP account`,
							Computed:            true,
						},
						"parsed_token": schema.SingleNestedAttribute{
							MarkdownDescription: `The parsed VPP service token`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"expires_at": schema.StringAttribute{
									MarkdownDescription: `The expiration time of the token`,
									Computed:            true,
								},
								"hashed_token": schema.StringAttribute{
									MarkdownDescription: `The hashed token`,
									Computed:            true,
								},
								"org_name": schema.StringAttribute{
									MarkdownDescription: `The organization name`,
									Computed:            true,
								},
							},
						},
						"vpp_account_id": schema.StringAttribute{
							MarkdownDescription: `The id of the VPP Account`,
							Computed:            true,
						},
						"vpp_location_id": schema.StringAttribute{
							MarkdownDescription: `The VPP location ID`,
							Computed:            true,
						},
						"vpp_location_name": schema.StringAttribute{
							MarkdownDescription: `The VPP location name`,
							Computed:            true,
						},
						"vpp_service_token": schema.StringAttribute{
							MarkdownDescription: `The VPP Account's Service Token`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsSmVppAccountsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSmVppAccounts OrganizationsSmVppAccounts
	diags := req.Config.Get(ctx, &organizationsSmVppAccounts)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsSmVppAccounts.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsSmVppAccounts.OrganizationID.IsNull(), !organizationsSmVppAccounts.VppAccountID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSmVppAccounts")
		vvOrganizationID := organizationsSmVppAccounts.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetOrganizationSmVppAccounts(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmVppAccounts",
				err.Error(),
			)
			return
		}

		organizationsSmVppAccounts = ResponseSmGetOrganizationSmVppAccountsItemsToBody(organizationsSmVppAccounts, response1)
		diags = resp.State.Set(ctx, &organizationsSmVppAccounts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSmVppAccount")
		vvOrganizationID := organizationsSmVppAccounts.OrganizationID.ValueString()
		vvVppAccountID := organizationsSmVppAccounts.VppAccountID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Sm.GetOrganizationSmVppAccount(vvOrganizationID, vvVppAccountID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmVppAccount",
				err.Error(),
			)
			return
		}

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmVppAccount",
				err.Error(),
			)
			return
		}

		organizationsSmVppAccounts = ResponseSmGetOrganizationSmVppAccountItemToBody(organizationsSmVppAccounts, response2)
		diags = resp.State.Set(ctx, &organizationsSmVppAccounts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSmVppAccounts struct {
	OrganizationID types.String                                  `tfsdk:"organization_id"`
	VppAccountID   types.String                                  `tfsdk:"vpp_account_id"`
	Items          *[]ResponseItemSmGetOrganizationSmVppAccounts `tfsdk:"items"`
	Item           *ResponseSmGetOrganizationSmVppAccount        `tfsdk:"item"`
}

type ResponseItemSmGetOrganizationSmVppAccounts struct {
	AllowedAdmins        types.String                                           `tfsdk:"allowed_admins"`
	AssignableNetworkIDs types.List                                             `tfsdk:"assignable_network_ids"`
	AssignableNetworks   types.String                                           `tfsdk:"assignable_networks"`
	ContentToken         types.String                                           `tfsdk:"content_token"`
	Email                types.String                                           `tfsdk:"email"`
	ID                   types.String                                           `tfsdk:"id"`
	LastForceSyncedAt    types.String                                           `tfsdk:"last_force_synced_at"`
	LastSyncedAt         types.String                                           `tfsdk:"last_synced_at"`
	Name                 types.String                                           `tfsdk:"name"`
	NetworkIDAdmins      types.String                                           `tfsdk:"network_id_admins"`
	ParsedToken          *ResponseItemSmGetOrganizationSmVppAccountsParsedToken `tfsdk:"parsed_token"`
	VppAccountID         types.String                                           `tfsdk:"vpp_account_id"`
	VppLocationID        types.String                                           `tfsdk:"vpp_location_id"`
	VppLocationName      types.String                                           `tfsdk:"vpp_location_name"`
	VppServiceToken      types.String                                           `tfsdk:"vpp_service_token"`
}

type ResponseItemSmGetOrganizationSmVppAccountsParsedToken struct {
	ExpiresAt   types.String `tfsdk:"expires_at"`
	HashedToken types.String `tfsdk:"hashed_token"`
	OrgName     types.String `tfsdk:"org_name"`
}

type ResponseSmGetOrganizationSmVppAccount struct {
	AllowedAdmins        types.String                                      `tfsdk:"allowed_admins"`
	AssignableNetworkIDs types.List                                        `tfsdk:"assignable_network_ids"`
	AssignableNetworks   types.String                                      `tfsdk:"assignable_networks"`
	ContentToken         types.String                                      `tfsdk:"content_token"`
	Email                types.String                                      `tfsdk:"email"`
	ID                   types.String                                      `tfsdk:"id"`
	LastForceSyncedAt    types.String                                      `tfsdk:"last_force_synced_at"`
	LastSyncedAt         types.String                                      `tfsdk:"last_synced_at"`
	Name                 types.String                                      `tfsdk:"name"`
	NetworkIDAdmins      types.String                                      `tfsdk:"network_id_admins"`
	ParsedToken          *ResponseSmGetOrganizationSmVppAccountParsedToken `tfsdk:"parsed_token"`
	VppAccountID         types.String                                      `tfsdk:"vpp_account_id"`
	VppLocationID        types.String                                      `tfsdk:"vpp_location_id"`
	VppLocationName      types.String                                      `tfsdk:"vpp_location_name"`
	VppServiceToken      types.String                                      `tfsdk:"vpp_service_token"`
}

type ResponseSmGetOrganizationSmVppAccountParsedToken struct {
	ExpiresAt   types.String `tfsdk:"expires_at"`
	HashedToken types.String `tfsdk:"hashed_token"`
	OrgName     types.String `tfsdk:"org_name"`
}

// ToBody
func ResponseSmGetOrganizationSmVppAccountsItemsToBody(state OrganizationsSmVppAccounts, response *merakigosdk.ResponseSmGetOrganizationSmVppAccounts) OrganizationsSmVppAccounts {
	var items []ResponseItemSmGetOrganizationSmVppAccounts
	for _, item := range *response {
		itemState := ResponseItemSmGetOrganizationSmVppAccounts{
			AllowedAdmins: func() types.String {
				if item.AllowedAdmins != "" {
					return types.StringValue(item.AllowedAdmins)
				}
				return types.String{}
			}(),
			AssignableNetworkIDs: StringSliceToList(item.AssignableNetworkIDs),
			AssignableNetworks: func() types.String {
				if item.AssignableNetworks != "" {
					return types.StringValue(item.AssignableNetworks)
				}
				return types.String{}
			}(),
			ContentToken: func() types.String {
				if item.ContentToken != "" {
					return types.StringValue(item.ContentToken)
				}
				return types.String{}
			}(),
			Email: func() types.String {
				if item.Email != "" {
					return types.StringValue(item.Email)
				}
				return types.String{}
			}(),
			ID: func() types.String {
				if item.ID != "" {
					return types.StringValue(item.ID)
				}
				return types.String{}
			}(),
			LastForceSyncedAt: func() types.String {
				if item.LastForceSyncedAt != "" {
					return types.StringValue(item.LastForceSyncedAt)
				}
				return types.String{}
			}(),
			LastSyncedAt: func() types.String {
				if item.LastSyncedAt != "" {
					return types.StringValue(item.LastSyncedAt)
				}
				return types.String{}
			}(),
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
			NetworkIDAdmins: func() types.String {
				if item.NetworkIDAdmins != "" {
					return types.StringValue(item.NetworkIDAdmins)
				}
				return types.String{}
			}(),
			ParsedToken: func() *ResponseItemSmGetOrganizationSmVppAccountsParsedToken {
				if item.ParsedToken != nil {
					return &ResponseItemSmGetOrganizationSmVppAccountsParsedToken{
						ExpiresAt: func() types.String {
							if item.ParsedToken.ExpiresAt != "" {
								return types.StringValue(item.ParsedToken.ExpiresAt)
							}
							return types.String{}
						}(),
						HashedToken: func() types.String {
							if item.ParsedToken.HashedToken != "" {
								return types.StringValue(item.ParsedToken.HashedToken)
							}
							return types.String{}
						}(),
						OrgName: func() types.String {
							if item.ParsedToken.OrgName != "" {
								return types.StringValue(item.ParsedToken.OrgName)
							}
							return types.String{}
						}(),
					}
				}
				return nil
			}(),
			VppAccountID: func() types.String {
				if item.VppAccountID != "" {
					return types.StringValue(item.VppAccountID)
				}
				return types.String{}
			}(),
			VppLocationID: func() types.String {
				if item.VppLocationID != "" {
					return types.StringValue(item.VppLocationID)
				}
				return types.String{}
			}(),
			VppLocationName: func() types.String {
				if item.VppLocationName != "" {
					return types.StringValue(item.VppLocationName)
				}
				return types.String{}
			}(),
			VppServiceToken: func() types.String {
				if item.VppServiceToken != "" {
					return types.StringValue(item.VppServiceToken)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSmGetOrganizationSmVppAccountItemToBody(state OrganizationsSmVppAccounts, response *merakigosdk.ResponseSmGetOrganizationSmVppAccount) OrganizationsSmVppAccounts {
	itemState := ResponseSmGetOrganizationSmVppAccount{
		AllowedAdmins: func() types.String {
			if response.AllowedAdmins != "" {
				return types.StringValue(response.AllowedAdmins)
			}
			return types.String{}
		}(),
		AssignableNetworkIDs: StringSliceToList(response.AssignableNetworkIDs),
		AssignableNetworks: func() types.String {
			if response.AssignableNetworks != "" {
				return types.StringValue(response.AssignableNetworks)
			}
			return types.String{}
		}(),
		ContentToken: func() types.String {
			if response.ContentToken != "" {
				return types.StringValue(response.ContentToken)
			}
			return types.String{}
		}(),
		Email: func() types.String {
			if response.Email != "" {
				return types.StringValue(response.Email)
			}
			return types.String{}
		}(),
		ID: func() types.String {
			if response.ID != "" {
				return types.StringValue(response.ID)
			}
			return types.String{}
		}(),
		LastForceSyncedAt: func() types.String {
			if response.LastForceSyncedAt != "" {
				return types.StringValue(response.LastForceSyncedAt)
			}
			return types.String{}
		}(),
		LastSyncedAt: func() types.String {
			if response.LastSyncedAt != "" {
				return types.StringValue(response.LastSyncedAt)
			}
			return types.String{}
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
		NetworkIDAdmins: func() types.String {
			if response.NetworkIDAdmins != "" {
				return types.StringValue(response.NetworkIDAdmins)
			}
			return types.String{}
		}(),
		ParsedToken: func() *ResponseSmGetOrganizationSmVppAccountParsedToken {
			if response.ParsedToken != nil {
				return &ResponseSmGetOrganizationSmVppAccountParsedToken{
					ExpiresAt: func() types.String {
						if response.ParsedToken.ExpiresAt != "" {
							return types.StringValue(response.ParsedToken.ExpiresAt)
						}
						return types.String{}
					}(),
					HashedToken: func() types.String {
						if response.ParsedToken.HashedToken != "" {
							return types.StringValue(response.ParsedToken.HashedToken)
						}
						return types.String{}
					}(),
					OrgName: func() types.String {
						if response.ParsedToken.OrgName != "" {
							return types.StringValue(response.ParsedToken.OrgName)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		VppAccountID: func() types.String {
			if response.VppAccountID != "" {
				return types.StringValue(response.VppAccountID)
			}
			return types.String{}
		}(),
		VppLocationID: func() types.String {
			if response.VppLocationID != "" {
				return types.StringValue(response.VppLocationID)
			}
			return types.String{}
		}(),
		VppLocationName: func() types.String {
			if response.VppLocationName != "" {
				return types.StringValue(response.VppLocationName)
			}
			return types.String{}
		}(),
		VppServiceToken: func() types.String {
			if response.VppServiceToken != "" {
				return types.StringValue(response.VppServiceToken)
			}
			return types.String{}
		}(),
	}
	state.Item = &itemState
	return state
}
