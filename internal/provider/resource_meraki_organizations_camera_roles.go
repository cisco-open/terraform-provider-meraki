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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsCameraRolesResource{}
	_ resource.ResourceWithConfigure = &OrganizationsCameraRolesResource{}
)

func NewOrganizationsCameraRolesResource() resource.Resource {
	return &OrganizationsCameraRolesResource{}
}

type OrganizationsCameraRolesResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsCameraRolesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsCameraRolesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_camera_roles"
}

func (r *OrganizationsCameraRolesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"applied_on_devices": schema.SetNestedAttribute{
				MarkdownDescription: `Device tag on which this specified permission is applied.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `Device id.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"in_networks_with_id": schema.StringAttribute{
							MarkdownDescription: `Network id scope`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"in_networks_with_tag": schema.StringAttribute{
							MarkdownDescription: `Network tag scope`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"permission_level": schema.StringAttribute{
							Computed: true,
						},
						"permission_scope": schema.StringAttribute{
							Computed: true,
						},
						"permission_scope_id": schema.StringAttribute{
							MarkdownDescription: `Permission scope id`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"tag": schema.StringAttribute{
							MarkdownDescription: `Device tag.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"applied_on_networks": schema.SetNestedAttribute{
				MarkdownDescription: `Network tag on which this specified permission is applied.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `Network id.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"permission_level": schema.StringAttribute{
							Computed: true,
						},
						"permission_scope": schema.StringAttribute{
							Computed: true,
						},
						"permission_scope_id": schema.StringAttribute{
							MarkdownDescription: `Permission scope id`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"tag": schema.StringAttribute{
							MarkdownDescription: `Network tag`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"applied_org_wide": schema.SetNestedAttribute{
				MarkdownDescription: `Permissions to be applied org wide.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"permission_level": schema.StringAttribute{
							Computed: true,
						},
						"permission_scope": schema.StringAttribute{
							Computed: true,
						},
						"permission_scope_id": schema.StringAttribute{
							MarkdownDescription: `Permission scope id`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"tag": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the new role. Must be unique. This parameter is required.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"role_id": schema.StringAttribute{
				MarkdownDescription: `roleId path parameter. Role ID`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `ID of the camera role`,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['roleId']

func (r *OrganizationsCameraRolesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsCameraRolesRs

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
	//Has Item and has items and post

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Camera.GetOrganizationCameraRoles(vvOrganizationID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationCameraRoles",
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
			vvRoleID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter RoleID",
					"Fail Parsing RoleID",
				)
				return
			}
			r.client.Camera.UpdateOrganizationCameraRole(vvOrganizationID, vvRoleID, data.toSdkApiRequestUpdate(ctx))

			responseVerifyItem2, _, _ := r.client.Camera.GetOrganizationCameraRole(vvOrganizationID, vvRoleID)
			if responseVerifyItem2 != nil {
				data = ResponseCameraGetOrganizationCameraRoleItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp2, err := r.client.Camera.CreateOrganizationCameraRole(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationCameraRole",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationCameraRole",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Camera.GetOrganizationCameraRoles(vvOrganizationID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraRoles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationCameraRoles",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvRoleID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter RoleID",
				"Fail Parsing RoleID",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Camera.GetOrganizationCameraRole(vvOrganizationID, vvRoleID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseCameraGetOrganizationCameraRoleItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationCameraRole",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraRole",
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

func (r *OrganizationsCameraRolesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsCameraRolesRs

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
	vvRoleID := data.RoleID.ValueString()
	responseGet, restyRespGet, err := r.client.Camera.GetOrganizationCameraRole(vvOrganizationID, vvRoleID)
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
				"Failure when executing GetOrganizationCameraRole",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationCameraRole",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCameraGetOrganizationCameraRoleItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsCameraRolesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("role_id"), idParts[1])...)
}

func (r *OrganizationsCameraRolesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsCameraRolesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvRoleID := data.RoleID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Camera.UpdateOrganizationCameraRole(vvOrganizationID, vvRoleID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationCameraRole",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationCameraRole",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsCameraRolesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsCameraRolesRs
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
	vvRoleID := state.RoleID.ValueString()
	_, err := r.client.Camera.DeleteOrganizationCameraRole(vvOrganizationID, vvRoleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationCameraRole", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsCameraRolesRs struct {
	OrganizationID    types.String                                                  `tfsdk:"organization_id"`
	RoleID            types.String                                                  `tfsdk:"role_id"`
	AppliedOnDevices  *[]ResponseCameraGetOrganizationCameraRoleAppliedOnDevicesRs  `tfsdk:"applied_on_devices"`
	AppliedOnNetworks *[]ResponseCameraGetOrganizationCameraRoleAppliedOnNetworksRs `tfsdk:"applied_on_networks"`
	AppliedOrgWide    *[]ResponseCameraGetOrganizationCameraRoleAppliedOrgWideRs    `tfsdk:"applied_org_wide"`
	Name              types.String                                                  `tfsdk:"name"`
	ID                types.String                                                  `tfsdk:"id"`
}

type ResponseCameraGetOrganizationCameraRoleAppliedOnDevicesRs struct {
	ID                types.String `tfsdk:"id"`
	PermissionLevel   types.String `tfsdk:"permission_level"`
	PermissionScope   types.String `tfsdk:"permission_scope"`
	PermissionScopeID types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
	InNetworksWithID  types.String `tfsdk:"in_networks_with_id"`
	InNetworksWithTag types.String `tfsdk:"in_networks_with_tag"`
}

type ResponseCameraGetOrganizationCameraRoleAppliedOnNetworksRs struct {
	ID                types.String `tfsdk:"id"`
	PermissionLevel   types.String `tfsdk:"permission_level"`
	PermissionScope   types.String `tfsdk:"permission_scope"`
	PermissionScopeID types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
}

type ResponseCameraGetOrganizationCameraRoleAppliedOrgWideRs struct {
	PermissionLevel   types.String `tfsdk:"permission_level"`
	PermissionScope   types.String `tfsdk:"permission_scope"`
	PermissionScopeID types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
}

// FromBody
func (r *OrganizationsCameraRolesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestCameraCreateOrganizationCameraRole {
	emptyString := ""
	var requestCameraCreateOrganizationCameraRoleAppliedOnDevices []merakigosdk.RequestCameraCreateOrganizationCameraRoleAppliedOnDevices

	if r.AppliedOnDevices != nil {
		for _, rItem1 := range *r.AppliedOnDevices {
			id := rItem1.ID.ValueString()
			inNetworksWithID := rItem1.InNetworksWithID.ValueString()
			inNetworksWithTag := rItem1.InNetworksWithTag.ValueString()
			permissionScopeID := rItem1.PermissionScopeID.ValueString()
			tag := rItem1.Tag.ValueString()
			requestCameraCreateOrganizationCameraRoleAppliedOnDevices = append(requestCameraCreateOrganizationCameraRoleAppliedOnDevices, merakigosdk.RequestCameraCreateOrganizationCameraRoleAppliedOnDevices{
				ID:                id,
				InNetworksWithID:  inNetworksWithID,
				InNetworksWithTag: inNetworksWithTag,
				PermissionScopeID: permissionScopeID,
				Tag:               tag,
			})
			//[debug] Is Array: True
		}
	}
	var requestCameraCreateOrganizationCameraRoleAppliedOnNetworks []merakigosdk.RequestCameraCreateOrganizationCameraRoleAppliedOnNetworks

	if r.AppliedOnNetworks != nil {
		for _, rItem1 := range *r.AppliedOnNetworks {
			id := rItem1.ID.ValueString()
			permissionScopeID := rItem1.PermissionScopeID.ValueString()
			tag := rItem1.Tag.ValueString()
			requestCameraCreateOrganizationCameraRoleAppliedOnNetworks = append(requestCameraCreateOrganizationCameraRoleAppliedOnNetworks, merakigosdk.RequestCameraCreateOrganizationCameraRoleAppliedOnNetworks{
				ID:                id,
				PermissionScopeID: permissionScopeID,
				Tag:               tag,
			})
			//[debug] Is Array: True
		}
	}
	var requestCameraCreateOrganizationCameraRoleAppliedOrgWide []merakigosdk.RequestCameraCreateOrganizationCameraRoleAppliedOrgWide

	if r.AppliedOrgWide != nil {
		for _, rItem1 := range *r.AppliedOrgWide {
			permissionScopeID := rItem1.PermissionScopeID.ValueString()
			requestCameraCreateOrganizationCameraRoleAppliedOrgWide = append(requestCameraCreateOrganizationCameraRoleAppliedOrgWide, merakigosdk.RequestCameraCreateOrganizationCameraRoleAppliedOrgWide{
				PermissionScopeID: permissionScopeID,
			})
			//[debug] Is Array: True
		}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestCameraCreateOrganizationCameraRole{
		AppliedOnDevices: func() *[]merakigosdk.RequestCameraCreateOrganizationCameraRoleAppliedOnDevices {
			if len(requestCameraCreateOrganizationCameraRoleAppliedOnDevices) > 0 {
				return &requestCameraCreateOrganizationCameraRoleAppliedOnDevices
			}
			return nil
		}(),
		AppliedOnNetworks: func() *[]merakigosdk.RequestCameraCreateOrganizationCameraRoleAppliedOnNetworks {
			if len(requestCameraCreateOrganizationCameraRoleAppliedOnNetworks) > 0 {
				return &requestCameraCreateOrganizationCameraRoleAppliedOnNetworks
			}
			return nil
		}(),
		AppliedOrgWide: func() *[]merakigosdk.RequestCameraCreateOrganizationCameraRoleAppliedOrgWide {
			if len(requestCameraCreateOrganizationCameraRoleAppliedOrgWide) > 0 {
				return &requestCameraCreateOrganizationCameraRoleAppliedOrgWide
			}
			return nil
		}(),
		Name: *name,
	}
	return &out
}
func (r *OrganizationsCameraRolesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCameraUpdateOrganizationCameraRole {
	emptyString := ""
	var requestCameraUpdateOrganizationCameraRoleAppliedOnDevices []merakigosdk.RequestCameraUpdateOrganizationCameraRoleAppliedOnDevices

	if r.AppliedOnDevices != nil {
		for _, rItem1 := range *r.AppliedOnDevices {
			id := rItem1.ID.ValueString()
			inNetworksWithID := rItem1.InNetworksWithID.ValueString()
			inNetworksWithTag := rItem1.InNetworksWithTag.ValueString()
			permissionScopeID := rItem1.PermissionScopeID.ValueString()
			tag := rItem1.Tag.ValueString()
			requestCameraUpdateOrganizationCameraRoleAppliedOnDevices = append(requestCameraUpdateOrganizationCameraRoleAppliedOnDevices, merakigosdk.RequestCameraUpdateOrganizationCameraRoleAppliedOnDevices{
				ID:                id,
				InNetworksWithID:  inNetworksWithID,
				InNetworksWithTag: inNetworksWithTag,
				PermissionScopeID: permissionScopeID,
				Tag:               tag,
			})
			//[debug] Is Array: True
		}
	}
	var requestCameraUpdateOrganizationCameraRoleAppliedOnNetworks []merakigosdk.RequestCameraUpdateOrganizationCameraRoleAppliedOnNetworks

	if r.AppliedOnNetworks != nil {
		for _, rItem1 := range *r.AppliedOnNetworks {
			id := rItem1.ID.ValueString()
			permissionScopeID := rItem1.PermissionScopeID.ValueString()
			tag := rItem1.Tag.ValueString()
			requestCameraUpdateOrganizationCameraRoleAppliedOnNetworks = append(requestCameraUpdateOrganizationCameraRoleAppliedOnNetworks, merakigosdk.RequestCameraUpdateOrganizationCameraRoleAppliedOnNetworks{
				ID:                id,
				PermissionScopeID: permissionScopeID,
				Tag:               tag,
			})
			//[debug] Is Array: True
		}
	}
	var requestCameraUpdateOrganizationCameraRoleAppliedOrgWide []merakigosdk.RequestCameraUpdateOrganizationCameraRoleAppliedOrgWide

	if r.AppliedOrgWide != nil {
		for _, rItem1 := range *r.AppliedOrgWide {
			permissionScopeID := rItem1.PermissionScopeID.ValueString()
			requestCameraUpdateOrganizationCameraRoleAppliedOrgWide = append(requestCameraUpdateOrganizationCameraRoleAppliedOrgWide, merakigosdk.RequestCameraUpdateOrganizationCameraRoleAppliedOrgWide{
				PermissionScopeID: permissionScopeID,
			})
			//[debug] Is Array: True
		}
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestCameraUpdateOrganizationCameraRole{
		AppliedOnDevices: func() *[]merakigosdk.RequestCameraUpdateOrganizationCameraRoleAppliedOnDevices {
			if len(requestCameraUpdateOrganizationCameraRoleAppliedOnDevices) > 0 {
				return &requestCameraUpdateOrganizationCameraRoleAppliedOnDevices
			}
			return nil
		}(),
		AppliedOnNetworks: func() *[]merakigosdk.RequestCameraUpdateOrganizationCameraRoleAppliedOnNetworks {
			if len(requestCameraUpdateOrganizationCameraRoleAppliedOnNetworks) > 0 {
				return &requestCameraUpdateOrganizationCameraRoleAppliedOnNetworks
			}
			return nil
		}(),
		AppliedOrgWide: func() *[]merakigosdk.RequestCameraUpdateOrganizationCameraRoleAppliedOrgWide {
			if len(requestCameraUpdateOrganizationCameraRoleAppliedOrgWide) > 0 {
				return &requestCameraUpdateOrganizationCameraRoleAppliedOrgWide
			}
			return nil
		}(),
		Name: *name,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCameraGetOrganizationCameraRoleItemToBodyRs(state OrganizationsCameraRolesRs, response *merakigosdk.ResponseCameraGetOrganizationCameraRole, is_read bool) OrganizationsCameraRolesRs {
	itemState := OrganizationsCameraRolesRs{
		AppliedOnDevices: func() *[]ResponseCameraGetOrganizationCameraRoleAppliedOnDevicesRs {
			if response.AppliedOnDevices != nil {
				result := make([]ResponseCameraGetOrganizationCameraRoleAppliedOnDevicesRs, len(*response.AppliedOnDevices))
				for i, appliedOnDevices := range *response.AppliedOnDevices {
					result[i] = ResponseCameraGetOrganizationCameraRoleAppliedOnDevicesRs{
						ID:                types.StringValue(appliedOnDevices.ID),
						PermissionLevel:   types.StringValue(appliedOnDevices.PermissionLevel),
						PermissionScope:   types.StringValue(appliedOnDevices.PermissionScope),
						PermissionScopeID: types.StringValue(appliedOnDevices.PermissionScopeID),
						Tag:               types.StringValue(appliedOnDevices.Tag),
					}
				}
				return &result
			}
			return nil
		}(),
		AppliedOnNetworks: func() *[]ResponseCameraGetOrganizationCameraRoleAppliedOnNetworksRs {
			if response.AppliedOnNetworks != nil {
				result := make([]ResponseCameraGetOrganizationCameraRoleAppliedOnNetworksRs, len(*response.AppliedOnNetworks))
				for i, appliedOnNetworks := range *response.AppliedOnNetworks {
					result[i] = ResponseCameraGetOrganizationCameraRoleAppliedOnNetworksRs{
						ID:                types.StringValue(appliedOnNetworks.ID),
						PermissionLevel:   types.StringValue(appliedOnNetworks.PermissionLevel),
						PermissionScope:   types.StringValue(appliedOnNetworks.PermissionScope),
						PermissionScopeID: types.StringValue(appliedOnNetworks.PermissionScopeID),
						Tag:               types.StringValue(appliedOnNetworks.Tag),
					}
				}
				return &result
			}
			return nil
		}(),
		AppliedOrgWide: func() *[]ResponseCameraGetOrganizationCameraRoleAppliedOrgWideRs {
			if response.AppliedOrgWide != nil {
				result := make([]ResponseCameraGetOrganizationCameraRoleAppliedOrgWideRs, len(*response.AppliedOrgWide))
				for i, appliedOrgWide := range *response.AppliedOrgWide {
					result[i] = ResponseCameraGetOrganizationCameraRoleAppliedOrgWideRs{
						PermissionLevel:   types.StringValue(appliedOrgWide.PermissionLevel),
						PermissionScope:   types.StringValue(appliedOrgWide.PermissionScope),
						PermissionScopeID: types.StringValue(appliedOrgWide.PermissionScopeID),
						Tag:               types.StringValue(appliedOrgWide.Tag),
					}
				}
				return &result
			}
			return nil
		}(),
		Name:   types.StringValue(response.Name),
		RoleID: types.StringValue(response.ID),
		ID:     types.StringValue(response.ID),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsCameraRolesRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsCameraRolesRs)
}
