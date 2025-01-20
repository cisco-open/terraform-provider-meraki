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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksMerakiAuthUsersDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksMerakiAuthUsersDataSource{}
)

func NewNetworksMerakiAuthUsersDataSource() datasource.DataSource {
	return &NetworksMerakiAuthUsersDataSource{}
}

type NetworksMerakiAuthUsersDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksMerakiAuthUsersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksMerakiAuthUsersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_meraki_auth_users"
}

func (d *NetworksMerakiAuthUsersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"meraki_auth_user_id": schema.StringAttribute{
				MarkdownDescription: `merakiAuthUserId path parameter. Meraki auth user ID`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"account_type": schema.StringAttribute{
						MarkdownDescription: `Authorization type for user.`,
						Computed:            true,
					},
					"authorizations": schema.SetNestedAttribute{
						MarkdownDescription: `User authorization info`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"authorized_by_email": schema.StringAttribute{
									MarkdownDescription: `User is authorized by the account email address`,
									Computed:            true,
								},
								"authorized_by_name": schema.StringAttribute{
									MarkdownDescription: `User is authorized by the account name`,
									Computed:            true,
								},
								"authorized_zone": schema.StringAttribute{
									MarkdownDescription: `Authorized zone of the user`,
									Computed:            true,
								},
								"expires_at": schema.StringAttribute{
									MarkdownDescription: `Authorization expiration time`,
									Computed:            true,
								},
								"ssid_number": schema.Int64Attribute{
									MarkdownDescription: `SSID number`,
									Computed:            true,
								},
							},
						},
					},
					"created_at": schema.StringAttribute{
						MarkdownDescription: `Creation time of the user`,
						Computed:            true,
					},
					"email": schema.StringAttribute{
						MarkdownDescription: `Email address of the user`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `Meraki auth user id`,
						Computed:            true,
					},
					"is_admin": schema.BoolAttribute{
						MarkdownDescription: `Whether or not the user is a Dashboard administrator`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the user`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkMerakiAuthUsers`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"account_type": schema.StringAttribute{
							MarkdownDescription: `Authorization type for user.`,
							Computed:            true,
						},
						"authorizations": schema.SetNestedAttribute{
							MarkdownDescription: `User authorization info`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"authorized_by_email": schema.StringAttribute{
										MarkdownDescription: `User is authorized by the account email address`,
										Computed:            true,
									},
									"authorized_by_name": schema.StringAttribute{
										MarkdownDescription: `User is authorized by the account name`,
										Computed:            true,
									},
									"authorized_zone": schema.StringAttribute{
										MarkdownDescription: `Authorized zone of the user`,
										Computed:            true,
									},
									"expires_at": schema.StringAttribute{
										MarkdownDescription: `Authorization expiration time`,
										Computed:            true,
									},
									"ssid_number": schema.Int64Attribute{
										MarkdownDescription: `SSID number`,
										Computed:            true,
									},
								},
							},
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: `Creation time of the user`,
							Computed:            true,
						},
						"email": schema.StringAttribute{
							MarkdownDescription: `Email address of the user`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `Meraki auth user id`,
							Computed:            true,
						},
						"is_admin": schema.BoolAttribute{
							MarkdownDescription: `Whether or not the user is a Dashboard administrator`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the user`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksMerakiAuthUsersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksMerakiAuthUsers NetworksMerakiAuthUsers
	diags := req.Config.Get(ctx, &networksMerakiAuthUsers)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksMerakiAuthUsers.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksMerakiAuthUsers.NetworkID.IsNull(), !networksMerakiAuthUsers.MerakiAuthUserID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkMerakiAuthUsers")
		vvNetworkID := networksMerakiAuthUsers.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkMerakiAuthUsers(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkMerakiAuthUsers",
				err.Error(),
			)
			return
		}

		networksMerakiAuthUsers = ResponseNetworksGetNetworkMerakiAuthUsersItemsToBody(networksMerakiAuthUsers, response1)
		diags = resp.State.Set(ctx, &networksMerakiAuthUsers)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkMerakiAuthUser")
		vvNetworkID := networksMerakiAuthUsers.NetworkID.ValueString()
		vvMerakiAuthUserID := networksMerakiAuthUsers.MerakiAuthUserID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Networks.GetNetworkMerakiAuthUser(vvNetworkID, vvMerakiAuthUserID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkMerakiAuthUser",
				err.Error(),
			)
			return
		}

		networksMerakiAuthUsers = ResponseNetworksGetNetworkMerakiAuthUserItemToBody(networksMerakiAuthUsers, response2)
		diags = resp.State.Set(ctx, &networksMerakiAuthUsers)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksMerakiAuthUsers struct {
	NetworkID        types.String                                     `tfsdk:"network_id"`
	MerakiAuthUserID types.String                                     `tfsdk:"meraki_auth_user_id"`
	Items            *[]ResponseItemNetworksGetNetworkMerakiAuthUsers `tfsdk:"items"`
	Item             *ResponseNetworksGetNetworkMerakiAuthUser        `tfsdk:"item"`
}

type ResponseItemNetworksGetNetworkMerakiAuthUsers struct {
	AccountType    types.String                                                   `tfsdk:"account_type"`
	Authorizations *[]ResponseItemNetworksGetNetworkMerakiAuthUsersAuthorizations `tfsdk:"authorizations"`
	CreatedAt      types.String                                                   `tfsdk:"created_at"`
	Email          types.String                                                   `tfsdk:"email"`
	ID             types.String                                                   `tfsdk:"id"`
	IsAdmin        types.Bool                                                     `tfsdk:"is_admin"`
	Name           types.String                                                   `tfsdk:"name"`
}

type ResponseItemNetworksGetNetworkMerakiAuthUsersAuthorizations struct {
	AuthorizedByEmail types.String `tfsdk:"authorized_by_email"`
	AuthorizedByName  types.String `tfsdk:"authorized_by_name"`
	AuthorizedZone    types.String `tfsdk:"authorized_zone"`
	ExpiresAt         types.String `tfsdk:"expires_at"`
	SSIDNumber        types.Int64  `tfsdk:"ssid_number"`
}

type ResponseNetworksGetNetworkMerakiAuthUser struct {
	AccountType    types.String                                              `tfsdk:"account_type"`
	Authorizations *[]ResponseNetworksGetNetworkMerakiAuthUserAuthorizations `tfsdk:"authorizations"`
	CreatedAt      types.String                                              `tfsdk:"created_at"`
	Email          types.String                                              `tfsdk:"email"`
	ID             types.String                                              `tfsdk:"id"`
	IsAdmin        types.Bool                                                `tfsdk:"is_admin"`
	Name           types.String                                              `tfsdk:"name"`
}

type ResponseNetworksGetNetworkMerakiAuthUserAuthorizations struct {
	AuthorizedByEmail types.String `tfsdk:"authorized_by_email"`
	AuthorizedByName  types.String `tfsdk:"authorized_by_name"`
	AuthorizedZone    types.String `tfsdk:"authorized_zone"`
	ExpiresAt         types.String `tfsdk:"expires_at"`
	SSIDNumber        types.Int64  `tfsdk:"ssid_number"`
}

// ToBody
func ResponseNetworksGetNetworkMerakiAuthUsersItemsToBody(state NetworksMerakiAuthUsers, response *merakigosdk.ResponseNetworksGetNetworkMerakiAuthUsers) NetworksMerakiAuthUsers {
	var items []ResponseItemNetworksGetNetworkMerakiAuthUsers
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkMerakiAuthUsers{
			AccountType: types.StringValue(item.AccountType),
			Authorizations: func() *[]ResponseItemNetworksGetNetworkMerakiAuthUsersAuthorizations {
				if item.Authorizations != nil {
					result := make([]ResponseItemNetworksGetNetworkMerakiAuthUsersAuthorizations, len(*item.Authorizations))
					for i, authorizations := range *item.Authorizations {
						result[i] = ResponseItemNetworksGetNetworkMerakiAuthUsersAuthorizations{
							AuthorizedByEmail: types.StringValue(authorizations.AuthorizedByEmail),
							AuthorizedByName:  types.StringValue(authorizations.AuthorizedByName),
							AuthorizedZone:    types.StringValue(authorizations.AuthorizedZone),
							ExpiresAt:         types.StringValue(authorizations.ExpiresAt),
							SSIDNumber: func() types.Int64 {
								if authorizations.SSIDNumber != nil {
									return types.Int64Value(int64(*authorizations.SSIDNumber))
								}
								return types.Int64{}
							}(),
						}
					}
					return &result
				}
				return nil
			}(),
			CreatedAt: types.StringValue(item.CreatedAt),
			Email:     types.StringValue(item.Email),
			ID:        types.StringValue(item.ID),
			IsAdmin: func() types.Bool {
				if item.IsAdmin != nil {
					return types.BoolValue(*item.IsAdmin)
				}
				return types.Bool{}
			}(),
			Name: types.StringValue(item.Name),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseNetworksGetNetworkMerakiAuthUserItemToBody(state NetworksMerakiAuthUsers, response *merakigosdk.ResponseNetworksGetNetworkMerakiAuthUser) NetworksMerakiAuthUsers {
	itemState := ResponseNetworksGetNetworkMerakiAuthUser{
		AccountType: types.StringValue(response.AccountType),
		Authorizations: func() *[]ResponseNetworksGetNetworkMerakiAuthUserAuthorizations {
			if response.Authorizations != nil {
				result := make([]ResponseNetworksGetNetworkMerakiAuthUserAuthorizations, len(*response.Authorizations))
				for i, authorizations := range *response.Authorizations {
					result[i] = ResponseNetworksGetNetworkMerakiAuthUserAuthorizations{
						AuthorizedByEmail: types.StringValue(authorizations.AuthorizedByEmail),
						AuthorizedByName:  types.StringValue(authorizations.AuthorizedByName),
						AuthorizedZone:    types.StringValue(authorizations.AuthorizedZone),
						ExpiresAt:         types.StringValue(authorizations.ExpiresAt),
						SSIDNumber: func() types.Int64 {
							if authorizations.SSIDNumber != nil {
								return types.Int64Value(int64(*authorizations.SSIDNumber))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		CreatedAt: types.StringValue(response.CreatedAt),
		Email:     types.StringValue(response.Email),
		ID:        types.StringValue(response.ID),
		IsAdmin: func() types.Bool {
			if response.IsAdmin != nil {
				return types.BoolValue(*response.IsAdmin)
			}
			return types.Bool{}
		}(),
		Name: types.StringValue(response.Name),
	}
	state.Item = &itemState
	return state
}
