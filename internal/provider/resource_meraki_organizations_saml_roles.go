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
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

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
	_ resource.Resource              = &OrganizationsSamlRolesResource{}
	_ resource.ResourceWithConfigure = &OrganizationsSamlRolesResource{}
)

func NewOrganizationsSamlRolesResource() resource.Resource {
	return &OrganizationsSamlRolesResource{}
}

type OrganizationsSamlRolesResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsSamlRolesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsSamlRolesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_saml_roles"
}

func (r *OrganizationsSamlRolesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"camera": schema.SetNestedAttribute{
				MarkdownDescription: `The list of camera access privileges for SAML administrator`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"access": schema.StringAttribute{
							MarkdownDescription: `Camera access ability`,
							Computed:            true,
						},
						"org_wide": schema.BoolAttribute{
							MarkdownDescription: `Whether or not SAML administrator has org-wide access`,
							Computed:            true,
						},
					},
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `ID associated with the SAML role`,
				Computed:            true,
			},
			"networks": schema.SetNestedAttribute{
				MarkdownDescription: `The list of networks that the SAML administrator has privileges on`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"access": schema.StringAttribute{
							MarkdownDescription: `The privilege of the SAML administrator on the network
                                        Allowed values: [full,guest-ambassador,monitor-only,port-tags,read-only,ssid-admin]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"full",
									"guest-ambassador",
									"monitor-only",
									"port-tags",
									"read-only",
									"ssid-admin",
								),
							},
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `The network ID`,
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
				MarkdownDescription: `The privilege of the SAML administrator on the organization
                                  Allowed values: [enterprise,full,none,read-only]`,
				Computed: true,
				Optional: true,
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
				Required:            true,
			},
			"role": schema.StringAttribute{
				MarkdownDescription: `The role of the SAML administrator`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"saml_role_id": schema.StringAttribute{
				MarkdownDescription: `samlRoleId path parameter. Saml role ID`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"tags": schema.SetNestedAttribute{
				MarkdownDescription: `The list of tags that the SAML administrator has privleges on`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"access": schema.StringAttribute{
							MarkdownDescription: `The privilege of the SAML administrator on the tag
                                        Allowed values: [full,guest-ambassador,monitor-only,read-only]`,
							Computed: true,
							Optional: true,
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
							MarkdownDescription: `The name of the tag`,
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
	}
}

// path params to set ['samlRoleId']
func (r *OrganizationsSamlRolesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsSamlRolesRs

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
	vvRole := data.Role.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationSamlRoles(vvOrganizationID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationSamlRoles",
					err.Error(),
				)
				return
			}
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Role", vvRole, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvSamlRoleID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter SamlRoleID",
					"Fail Parsing SamlRoleID",
				)
				return
			}
			r.client.Organizations.UpdateOrganizationSamlRole(vvOrganizationID, vvSamlRoleID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Organizations.GetOrganizationSamlRole(vvOrganizationID, vvSamlRoleID)
			if responseVerifyItem2 != nil {
				data = ResponseOrganizationsGetOrganizationSamlRoleItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Organizations.CreateOrganizationSamlRole(vvOrganizationID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationSamlRole",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationSamlRole",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationSamlRoles(vvOrganizationID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSamlRoles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationSamlRoles",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Role", vvRole, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvSamlRoleID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter SamlRoleID",
				"Error",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Organizations.GetOrganizationSamlRole(vvOrganizationID, vvSamlRoleID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseOrganizationsGetOrganizationSamlRoleItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationSamlRole",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSamlRole",
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

func (r *OrganizationsSamlRolesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsSamlRolesRs

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

	vvOrganizationID := data.OrganizationID.ValueString()
	vvSamlRoleID := data.SamlRoleID.ValueString()
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationSamlRole(vvOrganizationID, vvSamlRoleID)
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
				"Failure when executing GetOrganizationSamlRole",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationSamlRole",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationSamlRoleItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsSamlRolesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("saml_role_id"), idParts[1])...)
}

func (r *OrganizationsSamlRolesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsSamlRolesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvSamlRoleID := data.SamlRoleID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationSamlRole(vvOrganizationID, vvSamlRoleID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationSamlRole",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationSamlRole",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsSamlRolesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsSamlRolesRs
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
	vvSamlRoleID := state.SamlRoleID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationSamlRole(vvOrganizationID, vvSamlRoleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationSamlRole", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsSamlRolesRs struct {
	OrganizationID types.String                                              `tfsdk:"organization_id"`
	SamlRoleID     types.String                                              `tfsdk:"saml_role_id"`
	Camera         *[]ResponseOrganizationsGetOrganizationSamlRoleCameraRs   `tfsdk:"camera"`
	ID             types.String                                              `tfsdk:"id"`
	Networks       *[]ResponseOrganizationsGetOrganizationSamlRoleNetworksRs `tfsdk:"networks"`
	OrgAccess      types.String                                              `tfsdk:"org_access"`
	Role           types.String                                              `tfsdk:"role"`
	Tags           *[]ResponseOrganizationsGetOrganizationSamlRoleTagsRs     `tfsdk:"tags"`
}

type ResponseOrganizationsGetOrganizationSamlRoleCameraRs struct {
	Access  types.String `tfsdk:"access"`
	OrgWide types.Bool   `tfsdk:"org_wide"`
}

type ResponseOrganizationsGetOrganizationSamlRoleNetworksRs struct {
	Access types.String `tfsdk:"access"`
	ID     types.String `tfsdk:"id"`
}

type ResponseOrganizationsGetOrganizationSamlRoleTagsRs struct {
	Access types.String `tfsdk:"access"`
	Tag    types.String `tfsdk:"tag"`
}

// FromBody
func (r *OrganizationsSamlRolesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationSamlRole {
	emptyString := ""
	var requestOrganizationsCreateOrganizationSamlRoleNetworks []merakigosdk.RequestOrganizationsCreateOrganizationSamlRoleNetworks
	if r.Networks != nil {
		for _, rItem1 := range *r.Networks {
			access := rItem1.Access.ValueString()
			iD := rItem1.ID.ValueString()
			requestOrganizationsCreateOrganizationSamlRoleNetworks = append(requestOrganizationsCreateOrganizationSamlRoleNetworks, merakigosdk.RequestOrganizationsCreateOrganizationSamlRoleNetworks{
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
	role := new(string)
	if !r.Role.IsUnknown() && !r.Role.IsNull() {
		*role = r.Role.ValueString()
	} else {
		role = &emptyString
	}
	var requestOrganizationsCreateOrganizationSamlRoleTags []merakigosdk.RequestOrganizationsCreateOrganizationSamlRoleTags
	if r.Tags != nil {
		for _, rItem1 := range *r.Tags {
			access := rItem1.Access.ValueString()
			tag := rItem1.Tag.ValueString()
			requestOrganizationsCreateOrganizationSamlRoleTags = append(requestOrganizationsCreateOrganizationSamlRoleTags, merakigosdk.RequestOrganizationsCreateOrganizationSamlRoleTags{
				Access: access,
				Tag:    tag,
			})
		}
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationSamlRole{
		Networks: func() *[]merakigosdk.RequestOrganizationsCreateOrganizationSamlRoleNetworks {
			if len(requestOrganizationsCreateOrganizationSamlRoleNetworks) > 0 {
				return &requestOrganizationsCreateOrganizationSamlRoleNetworks
			}
			return nil
		}(),
		OrgAccess: *orgAccess,
		Role:      *role,
		Tags: func() *[]merakigosdk.RequestOrganizationsCreateOrganizationSamlRoleTags {
			if len(requestOrganizationsCreateOrganizationSamlRoleTags) > 0 {
				return &requestOrganizationsCreateOrganizationSamlRoleTags
			}
			return nil
		}(),
	}
	return &out
}
func (r *OrganizationsSamlRolesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationSamlRole {
	emptyString := ""
	var requestOrganizationsUpdateOrganizationSamlRoleNetworks []merakigosdk.RequestOrganizationsUpdateOrganizationSamlRoleNetworks
	if r.Networks != nil {
		for _, rItem1 := range *r.Networks {
			access := rItem1.Access.ValueString()
			iD := rItem1.ID.ValueString()
			requestOrganizationsUpdateOrganizationSamlRoleNetworks = append(requestOrganizationsUpdateOrganizationSamlRoleNetworks, merakigosdk.RequestOrganizationsUpdateOrganizationSamlRoleNetworks{
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
	role := new(string)
	if !r.Role.IsUnknown() && !r.Role.IsNull() {
		*role = r.Role.ValueString()
	} else {
		role = &emptyString
	}
	var requestOrganizationsUpdateOrganizationSamlRoleTags []merakigosdk.RequestOrganizationsUpdateOrganizationSamlRoleTags
	if r.Tags != nil {
		for _, rItem1 := range *r.Tags {
			access := rItem1.Access.ValueString()
			tag := rItem1.Tag.ValueString()
			requestOrganizationsUpdateOrganizationSamlRoleTags = append(requestOrganizationsUpdateOrganizationSamlRoleTags, merakigosdk.RequestOrganizationsUpdateOrganizationSamlRoleTags{
				Access: access,
				Tag:    tag,
			})
		}
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationSamlRole{
		Networks: func() *[]merakigosdk.RequestOrganizationsUpdateOrganizationSamlRoleNetworks {
			if len(requestOrganizationsUpdateOrganizationSamlRoleNetworks) > 0 {
				return &requestOrganizationsUpdateOrganizationSamlRoleNetworks
			}
			return nil
		}(),
		OrgAccess: *orgAccess,
		Role:      *role,
		Tags: func() *[]merakigosdk.RequestOrganizationsUpdateOrganizationSamlRoleTags {
			if len(requestOrganizationsUpdateOrganizationSamlRoleTags) > 0 {
				return &requestOrganizationsUpdateOrganizationSamlRoleTags
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationSamlRoleItemToBodyRs(state OrganizationsSamlRolesRs, response *merakigosdk.ResponseOrganizationsGetOrganizationSamlRole, is_read bool) OrganizationsSamlRolesRs {
	itemState := OrganizationsSamlRolesRs{
		Camera: func() *[]ResponseOrganizationsGetOrganizationSamlRoleCameraRs {
			if response.Camera != nil {
				result := make([]ResponseOrganizationsGetOrganizationSamlRoleCameraRs, len(*response.Camera))
				for i, camera := range *response.Camera {
					result[i] = ResponseOrganizationsGetOrganizationSamlRoleCameraRs{
						Access: types.StringValue(camera.Access),
						OrgWide: func() types.Bool {
							if camera.OrgWide != nil {
								return types.BoolValue(*camera.OrgWide)
							}
							return types.Bool{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		ID: types.StringValue(response.ID),
		Networks: func() *[]ResponseOrganizationsGetOrganizationSamlRoleNetworksRs {
			if response.Networks != nil {
				result := make([]ResponseOrganizationsGetOrganizationSamlRoleNetworksRs, len(*response.Networks))
				for i, networks := range *response.Networks {
					result[i] = ResponseOrganizationsGetOrganizationSamlRoleNetworksRs{
						Access: types.StringValue(networks.Access),
						ID:     types.StringValue(networks.ID),
					}
				}
				return &result
			}
			return nil
		}(),
		OrgAccess:  types.StringValue(response.OrgAccess),
		Role:       types.StringValue(response.Role),
		SamlRoleID: types.StringValue(response.ID),
		Tags: func() *[]ResponseOrganizationsGetOrganizationSamlRoleTagsRs {
			if response.Tags != nil {
				result := make([]ResponseOrganizationsGetOrganizationSamlRoleTagsRs, len(*response.Tags))
				for i, tags := range *response.Tags {
					result[i] = ResponseOrganizationsGetOrganizationSamlRoleTagsRs{
						Access: types.StringValue(tags.Access),
						Tag:    types.StringValue(tags.Tag),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsSamlRolesRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsSamlRolesRs)
}
