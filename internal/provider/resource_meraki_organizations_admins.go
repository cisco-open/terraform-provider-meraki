package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"log"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsAdminsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsAdminsResource{}
)

func NewOrganizationsAdminsResource() resource.Resource {
	return &OrganizationsAdminsResource{}
}

type OrganizationsAdminsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsAdminsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsAdminsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_admins"
}

func (r *OrganizationsAdminsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_status": schema.StringAttribute{
				MarkdownDescription: `Status of the admin's account`,
				Computed:            true,
			},
			"admin_id": schema.StringAttribute{
				MarkdownDescription: `adminId path parameter. Admin ID`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"authentication_method": schema.StringAttribute{
				MarkdownDescription: `Admin's authentication method`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Cisco SecureX Sign-On",
						"Email",
					),
				},
			},
			"email": schema.StringAttribute{
				MarkdownDescription: `Admin's email address`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
			},
			"has_api_key": schema.BoolAttribute{
				MarkdownDescription: `Indicates whether the admin has an API key`,
				Computed:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `Admin's ID`,
				Computed:            true,
			},
			"last_active": schema.StringAttribute{
				MarkdownDescription: `Time when the admin was last active`,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Admin's username`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"networks": schema.SetNestedAttribute{
				MarkdownDescription: `Admin network access information`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"access": schema.StringAttribute{
							MarkdownDescription: `Admin's level of access to the network`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `Network ID`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"org_access": schema.StringAttribute{
				MarkdownDescription: `Admin's level of access to the organization`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"enterprise",
						"full",
						"none",
						"read-only",
					),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Required: true,
			},
			"tags": schema.SetNestedAttribute{
				MarkdownDescription: `Admin tag information`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"access": schema.StringAttribute{
							MarkdownDescription: `Access level for the tag`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"full",
									"guest-ambassador",
									"monitor-only",
									"read-only",
								),
							},
						},
						"tag": schema.StringAttribute{
							MarkdownDescription: `Tag value`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"two_factor_auth_enabled": schema.BoolAttribute{
				MarkdownDescription: `Indicates whether two-factor authentication is enabled`,
				Computed:            true,
			},
		},
	}
}

//path params to set ['adminId']
//path params to assign NOT EDITABLE ['authenticationMethod', 'email']

func (r *OrganizationsAdminsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsAdminsRs

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
	vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	vvName := data.Email.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationAdmins(vvOrganizationID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationAdmins",
					err.Error(),
				)
				return
			}
		}
	}
	//TODO HAS ONLY ITEMS
	// Create

	responseStruct := structToMap(*responseVerifyItem)
	log.Printf("[DEBUG] responseStruct %v", responseStruct)
	log.Printf("[DEBUG] vvName %v", vvName)
	result := getDictResult(responseStruct, "Email", vvName, simpleCmp)
	log.Printf("[DEBUG] result %v", result)
	var responseVerifyItem2 merakigosdk.ResponseItemOrganizationsGetOrganizationAdmins
	if result != nil {
		err := mapToStruct(result.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		r.client.Organizations.UpdateOrganizationAdmin(vvOrganizationID, responseVerifyItem2.ID, data.toSdkApiRequestUpdate(ctx))

		responseVerifyItem2.Email = data.Email.ValueString()

		responseVerifyItem2.AuthenticationMethod = data.AuthenticationMethod.ValueString()

		data = ResponseOrganizationsGetOrganizationAdminsItemToBodyRs(data, &responseVerifyItem2, false)
		log.Printf("[DEBUG] data2 %v", data)

		// ['authenticationMethod', 'email']

		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Organizations.CreateOrganizationAdmin(vvOrganizationID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationAdmin",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationAdmin",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationAdmins(vvOrganizationID)
	// Has only items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdmins",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationAdmins",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet)
	result2 := getDictResult(responseStruct2, "Email", vvName, simpleCmp)
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		data = ResponseOrganizationsGetOrganizationAdminsItemToBodyRs(data, &responseVerifyItem2, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchLinkAggregations Result",
			"Not Found",
		)
		return
	}
}

func (r *OrganizationsAdminsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsAdminsRs

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
	// Not has Item

	vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	vvName := data.ID.ValueString()
	// name

	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationAdmins(vvOrganizationID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdmins",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationAdmins",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet)
	result2 := getDictResult(responseStruct2, "ID", vvName, simpleCmp)
	var responseVerifyItem2 merakigosdk.ResponseItemOrganizationsGetOrganizationAdmins
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		//entro aqui
		data = ResponseOrganizationsGetOrganizationAdminsItemToBodyRs(data, &responseVerifyItem2, true)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationAdmins Result",
			"Not found",
		)
		return
	}
}

func (r *OrganizationsAdminsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

func (r *OrganizationsAdminsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsAdminsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	vvAdminID := data.ID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)

	response, restyResp2, err := r.client.Organizations.UpdateOrganizationAdmin(vvOrganizationID, vvAdminID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationAdmin",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationAdmin",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsAdminsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsAdminsRs
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

	vvOrganizationID := state.OrganizationID.ValueString()
	vvAdminID := state.ID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationAdmin(vvOrganizationID, vvAdminID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationAdmin", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsAdminsRs struct {
	OrganizationID types.String `tfsdk:"organization_id"`
	AdminID        types.String `tfsdk:"admin_id"`
	//TIENE ITEMS
	AccountStatus        types.String                                                `tfsdk:"account_status"`
	AuthenticationMethod types.String                                                `tfsdk:"authentication_method"`
	Email                types.String                                                `tfsdk:"email"`
	HasAPIKey            types.Bool                                                  `tfsdk:"has_api_key"`
	ID                   types.String                                                `tfsdk:"id"`
	LastActive           types.String                                                `tfsdk:"last_active"`
	Name                 types.String                                                `tfsdk:"name"`
	Networks             *[]ResponseItemOrganizationsGetOrganizationAdminsNetworksRs `tfsdk:"networks"`
	OrgAccess            types.String                                                `tfsdk:"org_access"`
	Tags                 *[]ResponseItemOrganizationsGetOrganizationAdminsTagsRs     `tfsdk:"tags"`
	TwoFactorAuthEnabled types.Bool                                                  `tfsdk:"two_factor_auth_enabled"`
}

type ResponseItemOrganizationsGetOrganizationAdminsNetworksRs struct {
	Access types.String `tfsdk:"access"`
	ID     types.String `tfsdk:"id"`
}

type ResponseItemOrganizationsGetOrganizationAdminsTagsRs struct {
	Access types.String `tfsdk:"access"`
	Tag    types.String `tfsdk:"tag"`
}

// FromBody
func (r *OrganizationsAdminsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationAdmin {
	emptyString := ""
	authenticationMethod := new(string)
	if !r.AuthenticationMethod.IsUnknown() && !r.AuthenticationMethod.IsNull() {
		*authenticationMethod = r.AuthenticationMethod.ValueString()
	} else {
		authenticationMethod = &emptyString
	}
	email := new(string)
	if !r.Email.IsUnknown() && !r.Email.IsNull() {
		*email = r.Email.ValueString()
	} else {
		email = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestOrganizationsCreateOrganizationAdminNetworks []merakigosdk.RequestOrganizationsCreateOrganizationAdminNetworks
	if r.Networks != nil {
		for _, rItem1 := range *r.Networks {
			access := rItem1.Access.ValueString()
			iD := rItem1.ID.ValueString()
			requestOrganizationsCreateOrganizationAdminNetworks = append(requestOrganizationsCreateOrganizationAdminNetworks, merakigosdk.RequestOrganizationsCreateOrganizationAdminNetworks{
				Access: access,
				ID:     iD,
			})
		}
	}
	orgAccess := new(string)
	if !r.OrgAccess.IsUnknown() && !r.OrgAccess.IsNull() {
		*orgAccess = r.OrgAccess.ValueString()
	} else {
		orgAccess = &emptyString
	}
	var requestOrganizationsCreateOrganizationAdminTags []merakigosdk.RequestOrganizationsCreateOrganizationAdminTags
	if r.Tags != nil {
		for _, rItem1 := range *r.Tags {
			access := rItem1.Access.ValueString()
			tag := rItem1.Tag.ValueString()
			requestOrganizationsCreateOrganizationAdminTags = append(requestOrganizationsCreateOrganizationAdminTags, merakigosdk.RequestOrganizationsCreateOrganizationAdminTags{
				Access: access,
				Tag:    tag,
			})
		}
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationAdmin{
		AuthenticationMethod: *authenticationMethod,
		Email:                *email,
		Name:                 *name,
		Networks: func() *[]merakigosdk.RequestOrganizationsCreateOrganizationAdminNetworks {
			if len(requestOrganizationsCreateOrganizationAdminNetworks) > 0 {
				return &requestOrganizationsCreateOrganizationAdminNetworks
			}
			return nil
		}(),
		OrgAccess: *orgAccess,
		Tags: func() *[]merakigosdk.RequestOrganizationsCreateOrganizationAdminTags {
			if len(requestOrganizationsCreateOrganizationAdminTags) > 0 {
				return &requestOrganizationsCreateOrganizationAdminTags
			}
			return nil
		}(),
	}
	return &out
}
func (r *OrganizationsAdminsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationAdmin {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestOrganizationsUpdateOrganizationAdminNetworks []merakigosdk.RequestOrganizationsUpdateOrganizationAdminNetworks
	if r.Networks != nil {
		for _, rItem1 := range *r.Networks {
			access := rItem1.Access.ValueString()
			iD := rItem1.ID.ValueString()
			requestOrganizationsUpdateOrganizationAdminNetworks = append(requestOrganizationsUpdateOrganizationAdminNetworks, merakigosdk.RequestOrganizationsUpdateOrganizationAdminNetworks{
				Access: access,
				ID:     iD,
			})
		}
	}
	orgAccess := new(string)
	if !r.OrgAccess.IsUnknown() && !r.OrgAccess.IsNull() {
		*orgAccess = r.OrgAccess.ValueString()
	} else {
		orgAccess = &emptyString
	}
	var requestOrganizationsUpdateOrganizationAdminTags []merakigosdk.RequestOrganizationsUpdateOrganizationAdminTags
	if r.Tags != nil {
		for _, rItem1 := range *r.Tags {
			access := rItem1.Access.ValueString()
			tag := rItem1.Tag.ValueString()
			requestOrganizationsUpdateOrganizationAdminTags = append(requestOrganizationsUpdateOrganizationAdminTags, merakigosdk.RequestOrganizationsUpdateOrganizationAdminTags{
				Access: access,
				Tag:    tag,
			})
		}
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationAdmin{
		Name: *name,
		Networks: func() *[]merakigosdk.RequestOrganizationsUpdateOrganizationAdminNetworks {
			if len(requestOrganizationsUpdateOrganizationAdminNetworks) > 0 {
				return &requestOrganizationsUpdateOrganizationAdminNetworks
			}
			return nil
		}(),
		OrgAccess: *orgAccess,
		Tags: func() *[]merakigosdk.RequestOrganizationsUpdateOrganizationAdminTags {
			if len(requestOrganizationsUpdateOrganizationAdminTags) > 0 {
				return &requestOrganizationsUpdateOrganizationAdminTags
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationAdminsItemToBodyRs(state OrganizationsAdminsRs, response *merakigosdk.ResponseItemOrganizationsGetOrganizationAdmins, is_read bool) OrganizationsAdminsRs {
	itemState := OrganizationsAdminsRs{
		AccountStatus:        types.StringValue(response.AccountStatus),
		OrganizationID:       state.OrganizationID,
		AuthenticationMethod: types.StringValue(response.AuthenticationMethod),
		Email:                types.StringValue(response.Email),
		HasAPIKey: func() types.Bool {
			if response.HasAPIKey != nil {
				return types.BoolValue(*response.HasAPIKey)
			}
			return types.Bool{}
		}(),
		ID:         types.StringValue(response.ID),
		LastActive: types.StringValue(response.LastActive),
		Name:       types.StringValue(response.Name),
		Networks: func() *[]ResponseItemOrganizationsGetOrganizationAdminsNetworksRs {
			if response.Networks != nil {
				result := make([]ResponseItemOrganizationsGetOrganizationAdminsNetworksRs, len(*response.Networks))
				for i, networks := range *response.Networks {
					result[i] = ResponseItemOrganizationsGetOrganizationAdminsNetworksRs{
						Access: types.StringValue(networks.Access),
						ID:     types.StringValue(networks.ID),
					}
				}
				return &result
			}
			return nil
		}(),
		OrgAccess: types.StringValue(response.OrgAccess),
		Tags: func() *[]ResponseItemOrganizationsGetOrganizationAdminsTagsRs {
			if response.Tags != nil {
				result := make([]ResponseItemOrganizationsGetOrganizationAdminsTagsRs, len(*response.Tags))
				for i, tags := range *response.Tags {
					result[i] = ResponseItemOrganizationsGetOrganizationAdminsTagsRs{
						Access: types.StringValue(tags.Access),
						Tag:    types.StringValue(tags.Tag),
					}
				}
				return &result
			}
			return nil
		}(),
		TwoFactorAuthEnabled: func() types.Bool {
			if response.TwoFactorAuthEnabled != nil {
				return types.BoolValue(*response.TwoFactorAuthEnabled)
			}
			return types.Bool{}
		}(),
	}
	state = itemState
	return state
}
