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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksMerakiAuthUsersResource{}
	_ resource.ResourceWithConfigure = &NetworksMerakiAuthUsersResource{}
)

func NewNetworksMerakiAuthUsersResource() resource.Resource {
	return &NetworksMerakiAuthUsersResource{}
}

type NetworksMerakiAuthUsersResource struct {
	client *merakigosdk.Client
}

func (r *NetworksMerakiAuthUsersResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksMerakiAuthUsersResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_meraki_auth_users"
}

func (r *NetworksMerakiAuthUsersResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_type": schema.StringAttribute{
				MarkdownDescription: `Authorization type for user.
                                  Allowed values: [802.1X,Client VPN,Guest]`,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"802.1X",
						"Client VPN",
						"Guest",
					),
				},
			},
			"authorizations": schema.ListNestedAttribute{
				MarkdownDescription: `User authorization info`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
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
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"ssid_number": schema.Int64Attribute{
							MarkdownDescription: `SSID number`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
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
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
			},
			"email_password_to_user": schema.BoolAttribute{
				MarkdownDescription: `Whether or not Meraki should email the password to user. Default is false.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `Meraki auth user id`,
				Computed:            true,
			},
			"is_admin": schema.BoolAttribute{
				MarkdownDescription: `Whether or not the user is a Dashboard administrator`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
					SuppressDiffBool(),
				},
			},
			"meraki_auth_user_id": schema.StringAttribute{
				MarkdownDescription: `merakiAuthUserId path parameter. Meraki auth user ID`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the user`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"password": schema.StringAttribute{
				MarkdownDescription: `The password for this user account. Only required If the user is not a Dashboard administrator.`,
				Sensitive:           true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['merakiAuthUserId']
//path params to assign NOT EDITABLE ['accountType', 'email', 'isAdmin']

func (r *NetworksMerakiAuthUsersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksMerakiAuthUsersRs

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
	//Has Item and has items and post

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkMerakiAuthUsers(vvNetworkID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkMerakiAuthUsers",
					restyResp1.String(),
				)
				return
			}
		}
	}

	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvMerakiAuthUserID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter MerakiAuthUserID",
					"Fail Parsing MerakiAuthUserID",
				)
				return
			}
			r.client.Networks.UpdateNetworkMerakiAuthUser(vvNetworkID, vvMerakiAuthUserID, data.toSdkApiRequestUpdate(ctx))

			responseVerifyItem2, _, _ := r.client.Networks.GetNetworkMerakiAuthUser(vvNetworkID, vvMerakiAuthUserID)
			if responseVerifyItem2 != nil {
				data = ResponseNetworksGetNetworkMerakiAuthUserItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Networks.CreateNetworkMerakiAuthUser(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkMerakiAuthUser",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkMerakiAuthUser",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Networks.GetNetworkMerakiAuthUsers(vvNetworkID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkMerakiAuthUsers",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkMerakiAuthUsers",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvMerakiAuthUserID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter MerakiAuthUserID",
				"Fail Parsing MerakiAuthUserID",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Networks.GetNetworkMerakiAuthUser(vvNetworkID, vvMerakiAuthUserID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseNetworksGetNetworkMerakiAuthUserItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkMerakiAuthUser",
					restyRespGet.String(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkMerakiAuthUser",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}

}

func (r *NetworksMerakiAuthUsersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksMerakiAuthUsersRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	vvMerakiAuthUserID := data.MerakiAuthUserID.ValueString()
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkMerakiAuthUser(vvNetworkID, vvMerakiAuthUserID)
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
				"Failure when executing GetNetworkMerakiAuthUser",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkMerakiAuthUser",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkMerakiAuthUserItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksMerakiAuthUsersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: networkId,merakiAuthUserId. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("meraki_auth_user_id"), idParts[1])...)
}

func (r *NetworksMerakiAuthUsersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksMerakiAuthUsersRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	vvMerakiAuthUserID := plan.MerakiAuthUserID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkMerakiAuthUser(vvNetworkID, vvMerakiAuthUserID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkMerakiAuthUser",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkMerakiAuthUser",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksMerakiAuthUsersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksMerakiAuthUsersRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvNetworkID := state.NetworkID.ValueString()
	vvMerakiAuthUserID := state.MerakiAuthUserID.ValueString()
	_, err := r.client.Networks.DeleteNetworkMerakiAuthUser(vvNetworkID, vvMerakiAuthUserID, nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkMerakiAuthUser", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksMerakiAuthUsersRs struct {
	NetworkID           types.String                                                `tfsdk:"network_id"`
	MerakiAuthUserID    types.String                                                `tfsdk:"meraki_auth_user_id"`
	AccountType         types.String                                                `tfsdk:"account_type"`
	Authorizations      *[]ResponseNetworksGetNetworkMerakiAuthUserAuthorizationsRs `tfsdk:"authorizations"`
	CreatedAt           types.String                                                `tfsdk:"created_at"`
	Email               types.String                                                `tfsdk:"email"`
	ID                  types.String                                                `tfsdk:"id"`
	IsAdmin             types.Bool                                                  `tfsdk:"is_admin"`
	Name                types.String                                                `tfsdk:"name"`
	EmailPasswordToUser types.Bool                                                  `tfsdk:"email_password_to_user"`
	Password            types.String                                                `tfsdk:"password"`
}

type ResponseNetworksGetNetworkMerakiAuthUserAuthorizationsRs struct {
	AuthorizedByEmail types.String `tfsdk:"authorized_by_email"`
	AuthorizedByName  types.String `tfsdk:"authorized_by_name"`
	AuthorizedZone    types.String `tfsdk:"authorized_zone"`
	ExpiresAt         types.String `tfsdk:"expires_at"`
	SSIDNumber        types.Int64  `tfsdk:"ssid_number"`
}

// FromBody
func (r *NetworksMerakiAuthUsersRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksCreateNetworkMerakiAuthUser {
	emptyString := ""
	accountType := new(string)
	if !r.AccountType.IsUnknown() && !r.AccountType.IsNull() {
		*accountType = r.AccountType.ValueString()
	} else {
		accountType = &emptyString
	}
	var requestNetworksCreateNetworkMerakiAuthUserAuthorizations []merakigosdk.RequestNetworksCreateNetworkMerakiAuthUserAuthorizations

	if r.Authorizations != nil {
		for _, rItem1 := range *r.Authorizations {
			expiresAt := rItem1.ExpiresAt.ValueString()
			ssidNumber := func() *int64 {
				if !rItem1.SSIDNumber.IsUnknown() && !rItem1.SSIDNumber.IsNull() {
					return rItem1.SSIDNumber.ValueInt64Pointer()
				}
				return nil
			}()
			requestNetworksCreateNetworkMerakiAuthUserAuthorizations = append(requestNetworksCreateNetworkMerakiAuthUserAuthorizations, merakigosdk.RequestNetworksCreateNetworkMerakiAuthUserAuthorizations{
				ExpiresAt:  expiresAt,
				SSIDNumber: int64ToIntPointer(ssidNumber),
			})
			//[debug] Is Array: True
		}
	}
	email := new(string)
	if !r.Email.IsUnknown() && !r.Email.IsNull() {
		*email = r.Email.ValueString()
	} else {
		email = &emptyString
	}
	emailPasswordToUser := new(bool)
	if !r.EmailPasswordToUser.IsUnknown() && !r.EmailPasswordToUser.IsNull() {
		*emailPasswordToUser = r.EmailPasswordToUser.ValueBool()
	} else {
		emailPasswordToUser = nil
	}
	isAdmin := new(bool)
	if !r.IsAdmin.IsUnknown() && !r.IsAdmin.IsNull() {
		*isAdmin = r.IsAdmin.ValueBool()
	} else {
		isAdmin = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	password := new(string)
	if !r.Password.IsUnknown() && !r.Password.IsNull() {
		*password = r.Password.ValueString()
	} else {
		password = &emptyString
	}
	out := merakigosdk.RequestNetworksCreateNetworkMerakiAuthUser{
		AccountType: *accountType,
		Authorizations: func() *[]merakigosdk.RequestNetworksCreateNetworkMerakiAuthUserAuthorizations {
			if len(requestNetworksCreateNetworkMerakiAuthUserAuthorizations) > 0 {
				return &requestNetworksCreateNetworkMerakiAuthUserAuthorizations
			}
			return nil
		}(),
		Email:               *email,
		EmailPasswordToUser: emailPasswordToUser,
		IsAdmin:             isAdmin,
		Name:                *name,
		Password:            *password,
	}
	return &out
}
func (r *NetworksMerakiAuthUsersRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkMerakiAuthUser {
	emptyString := ""
	var requestNetworksUpdateNetworkMerakiAuthUserAuthorizations []merakigosdk.RequestNetworksUpdateNetworkMerakiAuthUserAuthorizations

	if r.Authorizations != nil {
		for _, rItem1 := range *r.Authorizations {
			expiresAt := rItem1.ExpiresAt.ValueString()
			ssidNumber := func() *int64 {
				if !rItem1.SSIDNumber.IsUnknown() && !rItem1.SSIDNumber.IsNull() {
					return rItem1.SSIDNumber.ValueInt64Pointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkMerakiAuthUserAuthorizations = append(requestNetworksUpdateNetworkMerakiAuthUserAuthorizations, merakigosdk.RequestNetworksUpdateNetworkMerakiAuthUserAuthorizations{
				ExpiresAt:  expiresAt,
				SSIDNumber: int64ToIntPointer(ssidNumber),
			})
			//[debug] Is Array: True
		}
	}
	emailPasswordToUser := new(bool)
	if !r.EmailPasswordToUser.IsUnknown() && !r.EmailPasswordToUser.IsNull() {
		*emailPasswordToUser = r.EmailPasswordToUser.ValueBool()
	} else {
		emailPasswordToUser = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	password := new(string)
	if !r.Password.IsUnknown() && !r.Password.IsNull() {
		*password = r.Password.ValueString()
	} else {
		password = &emptyString
	}
	out := merakigosdk.RequestNetworksUpdateNetworkMerakiAuthUser{
		Authorizations: func() *[]merakigosdk.RequestNetworksUpdateNetworkMerakiAuthUserAuthorizations {
			if len(requestNetworksUpdateNetworkMerakiAuthUserAuthorizations) > 0 {
				return &requestNetworksUpdateNetworkMerakiAuthUserAuthorizations
			}
			return nil
		}(),
		EmailPasswordToUser: emailPasswordToUser,
		Name:                *name,
		Password:            *password,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkMerakiAuthUserItemToBodyRs(state NetworksMerakiAuthUsersRs, response *merakigosdk.ResponseNetworksGetNetworkMerakiAuthUser, is_read bool) NetworksMerakiAuthUsersRs {
	itemState := NetworksMerakiAuthUsersRs{
		AccountType: func() types.String {
			if response.AccountType != "" {
				return types.StringValue(response.AccountType)
			}
			return types.String{}
		}(),
		Authorizations: func() *[]ResponseNetworksGetNetworkMerakiAuthUserAuthorizationsRs {
			if response.Authorizations != nil {
				result := make([]ResponseNetworksGetNetworkMerakiAuthUserAuthorizationsRs, len(*response.Authorizations))
				for i, authorizations := range *response.Authorizations {
					result[i] = ResponseNetworksGetNetworkMerakiAuthUserAuthorizationsRs{
						AuthorizedByEmail: func() types.String {
							if authorizations.AuthorizedByEmail != "" {
								return types.StringValue(authorizations.AuthorizedByEmail)
							}
							return types.String{}
						}(),
						AuthorizedByName: func() types.String {
							if authorizations.AuthorizedByName != "" {
								return types.StringValue(authorizations.AuthorizedByName)
							}
							return types.String{}
						}(),
						AuthorizedZone: func() types.String {
							if authorizations.AuthorizedZone != "" {
								return types.StringValue(authorizations.AuthorizedZone)
							}
							return types.String{}
						}(),
						ExpiresAt: func() types.String {
							if authorizations.ExpiresAt != "" {
								return types.StringValue(authorizations.ExpiresAt)
							}
							return types.String{}
						}(),
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
		CreatedAt: func() types.String {
			if response.CreatedAt != "" {
				return types.StringValue(response.CreatedAt)
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
		IsAdmin: func() types.Bool {
			if response.IsAdmin != nil {
				return types.BoolValue(*response.IsAdmin)
			}
			return types.Bool{}
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksMerakiAuthUsersRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksMerakiAuthUsersRs)
}
