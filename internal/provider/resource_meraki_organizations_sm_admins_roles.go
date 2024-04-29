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

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
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
	_ resource.Resource              = &OrganizationsSmAdminsRolesResource{}
	_ resource.ResourceWithConfigure = &OrganizationsSmAdminsRolesResource{}
)

func NewOrganizationsSmAdminsRolesResource() resource.Resource {
	return &OrganizationsSmAdminsRolesResource{}
}

type OrganizationsSmAdminsRolesResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsSmAdminsRolesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsSmAdminsRolesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_sm_admins_roles"
}

func (r *OrganizationsSmAdminsRolesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the limited access role`,
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
				MarkdownDescription: `The Id of the limited access role`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"scope": schema.StringAttribute{
				MarkdownDescription: `The scope of the limited access role`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"all_tags",
						"some",
						"without_all_tags",
						"without_some",
					),
				},
			},
			"tags": schema.SetAttribute{
				MarkdownDescription: `The tags of the limited access role`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
		},
	}
}

func (r *OrganizationsSmAdminsRolesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsSmAdminsRolesRs

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
	vvName := data.Name.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Sm.GetOrganizationSmAdminsRoles(vvOrganizationID, nil)
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmAdminsRoles",
				err.Error(),
			)
			return
		}
	}
	if responseVerifyItem != nil {

		responseStruct := structToMap(responseVerifyItem.Items)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvRoleID, ok := result2["RoleID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter SamlRoleID",
					"Error",
				)
				return
			}
			r.client.Sm.UpdateOrganizationSmAdminsRole(vvOrganizationID, vvRoleID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Sm.GetOrganizationSmAdminsRole(vvOrganizationID, vvRoleID)
			if responseVerifyItem2 != nil {
				data = ResponseSmGetOrganizationSmAdminsRoleItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Sm.CreateOrganizationSmAdminsRole(vvOrganizationID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationSmAdminsRole",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationSmAdminsRole",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Sm.GetOrganizationSmAdminsRoles(vvOrganizationID, nil)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmAdminsRoles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationSmAdminsRoles",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvRoleID, ok := result2["RoleID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter PolicyObjectID",
				"Error",
			)
			return
		}
		responseVerifyItem2, restyRespGet, _ := r.client.Sm.GetOrganizationSmAdminsRole(vvOrganizationID, vvRoleID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseSmGetOrganizationSmAdminsRoleItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationSmAdminsRole",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmAdminsRole",
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

func (r *OrganizationsSmAdminsRolesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsSmAdminsRolesRs

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
	responseGet, restyRespGet, err := r.client.Sm.GetOrganizationSmAdminsRole(vvOrganizationID, vvRoleID)
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
				"Failure when executing GetOrganizationSmAdminsRoles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationSmAdminsRoles",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSmGetOrganizationSmAdminsRoleItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsSmAdminsRolesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsSmAdminsRolesRs
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
	response, restyResp2, err := r.client.Sm.UpdateOrganizationSmAdminsRole(vvOrganizationID, vvRoleID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationSmAdminsRole",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationSmAdminsRole",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsSmAdminsRolesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsSmAdminsRolesRs
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
	_, err := r.client.Sm.DeleteOrganizationSmAdminsRole(vvOrganizationID, vvRoleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationSmAdminsRole", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsSmAdminsRolesRs struct {
	OrganizationID types.String `tfsdk:"organization_id"`
	RoleID         types.String `tfsdk:"role_id"`
	Name           types.String `tfsdk:"name"`
	Scope          types.String `tfsdk:"scope"`
	Tags           types.Set    `tfsdk:"tags"`
}

// FromBody
func (r *OrganizationsSmAdminsRolesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSmCreateOrganizationSmAdminsRole {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	scope := new(string)
	if !r.Scope.IsUnknown() && !r.Scope.IsNull() {
		*scope = r.Scope.ValueString()
	} else {
		scope = &emptyString
	}
	var tags []string = nil
	r.Tags.ElementsAs(ctx, &tags, false)
	out := merakigosdk.RequestSmCreateOrganizationSmAdminsRole{
		Name:  *name,
		Scope: *scope,
		Tags:  tags,
	}
	return &out
}
func (r *OrganizationsSmAdminsRolesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSmUpdateOrganizationSmAdminsRole {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	scope := new(string)
	if !r.Scope.IsUnknown() && !r.Scope.IsNull() {
		*scope = r.Scope.ValueString()
	} else {
		scope = &emptyString
	}
	var tags []string = nil
	r.Tags.ElementsAs(ctx, &tags, false)
	out := merakigosdk.RequestSmUpdateOrganizationSmAdminsRole{
		Name:  *name,
		Scope: *scope,
		Tags:  tags,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSmGetOrganizationSmAdminsRoleItemToBodyRs(state OrganizationsSmAdminsRolesRs, response *merakigosdk.ResponseSmGetOrganizationSmAdminsRole, is_read bool) OrganizationsSmAdminsRolesRs {
	itemState := OrganizationsSmAdminsRolesRs{
		Name:   types.StringValue(response.Name),
		RoleID: types.StringValue(response.RoleID),
		Scope:  types.StringValue(response.Scope),
		Tags:   StringSliceToSet(response.Tags),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsSmAdminsRolesRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsSmAdminsRolesRs)
}
