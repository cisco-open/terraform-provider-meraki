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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
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
				MarkdownDescription: `Authorization type for user.`,
				Computed:            true,
				Optional:            true,
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
			"authorizations": schema.SetNestedAttribute{
				MarkdownDescription: `User authorization info`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
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
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"ssid_number": schema.Int64Attribute{
							MarkdownDescription: `SSID number`,
							Computed:            true,
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
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
			},
			"email_password_to_user": schema.BoolAttribute{
				MarkdownDescription: `Whether or not Meraki should email the password to user. Default is false.`,
				Computed:            true,
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
				Computed:            true,
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
				Computed:            true,
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
				Computed:            true,
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
	//Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkMerakiAuthUsers(vvNetworkID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkMerakiAuthUsers",
					err.Error(),
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
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkMerakiAuthUser",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkMerakiAuthUser",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Networks.GetNetworkMerakiAuthUsers(vvNetworkID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkMerakiAuthUsers",
				err.Error(),
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
				err.Error(),
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
					err.Error(),
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
				err.Error(),
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
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksMerakiAuthUsersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("meraki_user_id"), idParts[1])...)
}

func (r *NetworksMerakiAuthUsersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksMerakiAuthUsersRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvMerakiAuthUserID := data.MerakiAuthUserID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkMerakiAuthUser(vvNetworkID, vvMerakiAuthUserID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkMerakiAuthUser",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkMerakiAuthUser",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
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
			sSIDNumber := func() *int64 {
				if !rItem1.SSIDNumber.IsUnknown() && !rItem1.SSIDNumber.IsNull() {
					return rItem1.SSIDNumber.ValueInt64Pointer()
				}
				return nil
			}()
			requestNetworksCreateNetworkMerakiAuthUserAuthorizations = append(requestNetworksCreateNetworkMerakiAuthUserAuthorizations, merakigosdk.RequestNetworksCreateNetworkMerakiAuthUserAuthorizations{
				ExpiresAt:  expiresAt,
				SSIDNumber: int64ToIntPointer(sSIDNumber),
			})
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
			sSIDNumber := func() *int64 {
				if !rItem1.SSIDNumber.IsUnknown() && !rItem1.SSIDNumber.IsNull() {
					return rItem1.SSIDNumber.ValueInt64Pointer()
				}
				return nil
			}()
			requestNetworksUpdateNetworkMerakiAuthUserAuthorizations = append(requestNetworksUpdateNetworkMerakiAuthUserAuthorizations, merakigosdk.RequestNetworksUpdateNetworkMerakiAuthUserAuthorizations{
				ExpiresAt:  expiresAt,
				SSIDNumber: int64ToIntPointer(sSIDNumber),
			})
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
		AccountType: types.StringValue(response.AccountType),
		Authorizations: func() *[]ResponseNetworksGetNetworkMerakiAuthUserAuthorizationsRs {
			if response.Authorizations != nil {
				result := make([]ResponseNetworksGetNetworkMerakiAuthUserAuthorizationsRs, len(*response.Authorizations))
				for i, authorizations := range *response.Authorizations {
					result[i] = ResponseNetworksGetNetworkMerakiAuthUserAuthorizationsRs{
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
			return &[]ResponseNetworksGetNetworkMerakiAuthUserAuthorizationsRs{}
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
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksMerakiAuthUsersRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksMerakiAuthUsersRs)
}
