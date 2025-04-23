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
	"log"
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
	_ resource.Resource              = &OrganizationsPolicyObjectsGroupsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsPolicyObjectsGroupsResource{}
)

func NewOrganizationsPolicyObjectsGroupsResource() resource.Resource {
	return &OrganizationsPolicyObjectsGroupsResource{}
}

type OrganizationsPolicyObjectsGroupsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsPolicyObjectsGroupsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsPolicyObjectsGroupsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_policy_objects_groups"
}

func (r *OrganizationsPolicyObjectsGroupsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"category": schema.StringAttribute{
				MarkdownDescription: `Type of object groups. (NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
			},
			"created_at": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the Policy object group.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_ids": schema.SetAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				ElementType: types.StringType,
			},
			"object_ids": schema.SetAttribute{
				MarkdownDescription: `Policy objects associated with Network Object Group or Port Object Group`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"policy_object_group_id": schema.StringAttribute{
				MarkdownDescription: `policyObjectGroupId path parameter. Policy object group ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: `Time Stamp of policy object updation.`,
				Computed:            true,
			},
		},
	}
}

//path params to assign NOT EDITABLE ['category']

func (r *OrganizationsPolicyObjectsGroupsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsPolicyObjectsGroupsRs

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
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationPolicyObjectsGroups(vvOrganizationID, &merakigosdk.GetOrganizationPolicyObjectsGroupsQueryParams{
		PerPage: -1,
	})
	log.Println("responseVerifyItem", responseVerifyItem)
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationPolicyObjectsGroups",
				err.Error(),
			)
			return
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvPolicyObjectGroupID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter PolicyObjectGroupID",
					"Error",
				)
				return
			}
			r.client.Organizations.UpdateOrganizationPolicyObjectsGroup(vvOrganizationID, vvPolicyObjectGroupID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Organizations.GetOrganizationPolicyObjectsGroup(vvOrganizationID, vvPolicyObjectGroupID)
			if responseVerifyItem2 != nil {
				data = ResponseOrganizationsGetOrganizationPolicyObjectsGroupItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	_, restyResp2, err := r.client.Organizations.CreateOrganizationPolicyObjectsGroup(vvOrganizationID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationPolicyObjectsGroup",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationPolicyObjectsGroup",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationPolicyObjectsGroups(vvOrganizationID, &merakigosdk.GetOrganizationPolicyObjectsGroupsQueryParams{
		PerPage: -1,
	})
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationPolicyObjectsGroups",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationPolicyObjectsGroups",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvPolicyObjectGroupID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter PolicyObjectGroupID",
				"Error",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Organizations.GetOrganizationPolicyObjectsGroup(vvOrganizationID, vvPolicyObjectGroupID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseOrganizationsGetOrganizationPolicyObjectsGroupItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationPolicyObjectsGroup",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationPolicyObjectsGroup",
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

func (r *OrganizationsPolicyObjectsGroupsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsPolicyObjectsGroupsRs

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
	// organization_id
	vvPolicyObjectGroupID := data.PolicyObjectGroupID.ValueString()
	// policy_object_group_id
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationPolicyObjectsGroup(vvOrganizationID, vvPolicyObjectGroupID)
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
				"Failure when executing GetOrganizationPolicyObjectsGroups",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationPolicyObjectsGroups",
			err.Error(),
		)
		return
	}

	data = ResponseOrganizationsGetOrganizationPolicyObjectsGroupItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsPolicyObjectsGroupsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("policy_object_group_id"), idParts[1])...)
}

func (r *OrganizationsPolicyObjectsGroupsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsPolicyObjectsGroupsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvPolicyObjectGroupID := data.PolicyObjectGroupID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationPolicyObjectsGroup(vvOrganizationID, vvPolicyObjectGroupID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationPolicyObjectsGroup",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationPolicyObjectsGroup",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsPolicyObjectsGroupsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsPolicyObjectsGroupsRs
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
	vvPolicyObjectGroupID := state.PolicyObjectGroupID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationPolicyObjectsGroup(vvOrganizationID, vvPolicyObjectGroupID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationPolicyObjectsGroup", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsPolicyObjectsGroupsRs struct {
	OrganizationID      types.String `tfsdk:"organization_id"`
	PolicyObjectGroupID types.String `tfsdk:"policy_object_group_id"`
	Category            types.String `tfsdk:"category"`
	CreatedAt           types.String `tfsdk:"created_at"`
	ID                  types.String `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	NetworkIDs          types.Set    `tfsdk:"network_ids"`
	ObjectIDs           types.Set    `tfsdk:"object_ids"`
	UpdatedAt           types.String `tfsdk:"updated_at"`
}

// FromBody
func (r *OrganizationsPolicyObjectsGroupsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationPolicyObjectsGroup {
	emptyString := ""
	category := new(string)
	if !r.Category.IsUnknown() && !r.Category.IsNull() {
		*category = r.Category.ValueString()
	} else {
		category = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var objectIDs *[]string = nil
	r.ObjectIDs.ElementsAs(ctx, &objectIDs, false)
	out := merakigosdk.RequestOrganizationsCreateOrganizationPolicyObjectsGroup{
		Category:  *category,
		Name:      *name,
		ObjectIDs: objectIDs,
	}
	return &out
}
func (r *OrganizationsPolicyObjectsGroupsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationPolicyObjectsGroup {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var objectIDs *[]string = nil
	r.ObjectIDs.ElementsAs(ctx, &objectIDs, false)
	out := merakigosdk.RequestOrganizationsUpdateOrganizationPolicyObjectsGroup{
		Name:      *name,
		ObjectIDs: objectIDs,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationPolicyObjectsGroupItemToBodyRs(state OrganizationsPolicyObjectsGroupsRs, response *merakigosdk.ResponseOrganizationsGetOrganizationPolicyObjectsGroup, is_read bool) OrganizationsPolicyObjectsGroupsRs {
	itemState := OrganizationsPolicyObjectsGroupsRs{
		Category:   types.StringValue(response.Category),
		CreatedAt:  types.StringValue(response.CreatedAt),
		ID:         types.StringValue(response.ID),
		Name:       types.StringValue(response.Name),
		NetworkIDs: StringSliceToSet(response.NetworkIDs),
		ObjectIDs: func() types.Set {
			if response.ObjectIDs == nil {
				return types.SetNull(types.StringType)
			}
			return StringSliceToSet(*response.ObjectIDs)
		}(),
		UpdatedAt: types.StringValue(response.UpdatedAt),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsPolicyObjectsGroupsRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsPolicyObjectsGroupsRs)
}
