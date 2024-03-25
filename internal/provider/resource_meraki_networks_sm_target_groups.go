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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSmTargetGroupsResource{}
	_ resource.ResourceWithConfigure = &NetworksSmTargetGroupsResource{}
)

func NewNetworksSmTargetGroupsResource() resource.Resource {
	return &NetworksSmTargetGroupsResource{}
}

type NetworksSmTargetGroupsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSmTargetGroupsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSmTargetGroupsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_target_groups"
}

func (r *NetworksSmTargetGroupsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of this target group`,
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
			"scope": schema.StringAttribute{
				MarkdownDescription: `The scope and tag options of the target group. Comma separated values beginning with one of withAny, withAll, withoutAny, withoutAll, all, none, followed by tags. Default to none if empty.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"tags": schema.StringAttribute{
				Computed: true,
			},
			"target_group_id": schema.StringAttribute{
				MarkdownDescription: `targetGroupId path parameter. Target group ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"type": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

//path params to set ['targetGroupId']

func (r *NetworksSmTargetGroupsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSmTargetGroupsRs

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
	// network_id
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Sm.GetNetworkSmTargetGroups(vvNetworkID, nil)
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmTargetGroups",
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
			vvTargetGroupID, ok := result2["TargetGroupID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter TargetGroupID",
					"Error",
				)
				return
			}
			r.client.Sm.UpdateNetworkSmTargetGroup(vvNetworkID, vvTargetGroupID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Sm.GetNetworkSmTargetGroup(vvNetworkID, vvTargetGroupID, nil)
			if responseVerifyItem2 != nil {
				data = ResponseSmGetNetworkSmTargetGroupItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp2, err := r.client.Sm.CreateNetworkSmTargetGroup(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkSmTargetGroup",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkSmTargetGroup",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Sm.GetNetworkSmTargetGroups(vvNetworkID, nil)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmTargetGroups",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSmTargetGroups",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvTargetGroupID, ok := result2["TargetGroupID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter TargetGroupID",
				"Error",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Sm.GetNetworkSmTargetGroup(vvNetworkID, vvTargetGroupID, nil)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseSmGetNetworkSmTargetGroupItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkSmTargetGroup",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmTargetGroup",
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

func (r *NetworksSmTargetGroupsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSmTargetGroupsRs

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
	// network_id
	vvTargetGroupID := data.TargetGroupID.ValueString()
	// target_group_id
	responseGet, restyRespGet, err := r.client.Sm.GetNetworkSmTargetGroup(vvNetworkID, vvTargetGroupID, nil)
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
				"Failure when executing GetNetworkSmTargetGroup",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSmTargetGroup",
			err.Error(),
		)
		return
	}

	data = ResponseSmGetNetworkSmTargetGroupItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSmTargetGroupsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("target_group_id"), idParts[1])...)
}

func (r *NetworksSmTargetGroupsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSmTargetGroupsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvTargetGroupID := data.TargetGroupID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Sm.UpdateNetworkSmTargetGroup(vvNetworkID, vvTargetGroupID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSmTargetGroup",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSmTargetGroup",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSmTargetGroupsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksSmTargetGroupsRs
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
	vvTargetGroupID := state.TargetGroupID.ValueString()
	_, err := r.client.Sm.DeleteNetworkSmTargetGroup(vvNetworkID, vvTargetGroupID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkSmTargetGroup", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksSmTargetGroupsRs struct {
	NetworkID     types.String `tfsdk:"network_id"`
	TargetGroupID types.String `tfsdk:"target_group_id"`
	Name          types.String `tfsdk:"name"`
	Scope         types.String `tfsdk:"scope"`
	Tags          types.String `tfsdk:"tags"`
	Type          types.String `tfsdk:"type"`
}

// FromBody
func (r *NetworksSmTargetGroupsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSmCreateNetworkSmTargetGroup {
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
	out := merakigosdk.RequestSmCreateNetworkSmTargetGroup{
		Name:  *name,
		Scope: *scope,
	}
	return &out
}
func (r *NetworksSmTargetGroupsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSmUpdateNetworkSmTargetGroup {
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
	out := merakigosdk.RequestSmUpdateNetworkSmTargetGroup{
		Name:  *name,
		Scope: *scope,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSmGetNetworkSmTargetGroupItemToBodyRs(state NetworksSmTargetGroupsRs, response *merakigosdk.ResponseSmGetNetworkSmTargetGroup, is_read bool) NetworksSmTargetGroupsRs {
	itemState := NetworksSmTargetGroupsRs{
		Name:  types.StringValue(response.Name),
		Scope: types.StringValue(response.Scope),
		Tags:  types.StringValue(response.Tags),
		Type:  types.StringValue(response.Type),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSmTargetGroupsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSmTargetGroupsRs)
}
