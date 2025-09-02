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
	"strconv"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchStormControlResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchStormControlResource{}
)

func NewNetworksSwitchStormControlResource() resource.Resource {
	return &NetworksSwitchStormControlResource{}
}

type NetworksSwitchStormControlResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchStormControlResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchStormControlResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_storm_control"
}

func (r *NetworksSwitchStormControlResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"broadcast_threshold": schema.Int64Attribute{
				MarkdownDescription: `Broadcast threshold.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"multicast_threshold": schema.Int64Attribute{
				MarkdownDescription: `Multicast threshold.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"treat_these_traffic_types_as_one_threshold": schema.ListAttribute{
				MarkdownDescription: `Grouped traffic types`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
				Default:     listdefault.StaticValue(types.ListNull(types.StringType)),
			},
			"unknown_unicast_threshold": schema.Int64Attribute{
				MarkdownDescription: `Unknown Unicast threshold.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *NetworksSwitchStormControlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchStormControlRs

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
	//Has Item and not has items

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchStormControl(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchStormControl",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchStormControl",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksSwitchStormControlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchStormControlRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchStormControl(vvNetworkID)
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
				"Failure when executing GetNetworkSwitchStormControl",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchStormControl",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchStormControlItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksSwitchStormControlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchStormControlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksSwitchStormControlRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchStormControl(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchStormControl",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchStormControl",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksSwitchStormControlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSwitchStormControl", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchStormControlRs struct {
	NetworkID                            types.String `tfsdk:"network_id"`
	BroadcastThreshold                   types.Int64  `tfsdk:"broadcast_threshold"`
	MulticastThreshold                   types.Int64  `tfsdk:"multicast_threshold"`
	TreatTheseTrafficTypesAsOneThreshold types.List   `tfsdk:"treat_these_traffic_types_as_one_threshold"`
	UnknownUnicastThreshold              types.Int64  `tfsdk:"unknown_unicast_threshold"`
}

// FromBody
func (r *NetworksSwitchStormControlRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchStormControl {
	broadcastThreshold := new(int64)
	if !r.BroadcastThreshold.IsUnknown() && !r.BroadcastThreshold.IsNull() {
		*broadcastThreshold = r.BroadcastThreshold.ValueInt64()
	} else {
		broadcastThreshold = nil
	}
	multicastThreshold := new(int64)
	if !r.MulticastThreshold.IsUnknown() && !r.MulticastThreshold.IsNull() {
		*multicastThreshold = r.MulticastThreshold.ValueInt64()
	} else {
		multicastThreshold = nil
	}
	var treatTheseTrafficTypesAsOneThreshold []string = nil
	r.TreatTheseTrafficTypesAsOneThreshold.ElementsAs(ctx, &treatTheseTrafficTypesAsOneThreshold, false)
	unknownUnicastThreshold := new(int64)
	if !r.UnknownUnicastThreshold.IsUnknown() && !r.UnknownUnicastThreshold.IsNull() {
		*unknownUnicastThreshold = r.UnknownUnicastThreshold.ValueInt64()
	} else {
		unknownUnicastThreshold = nil
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchStormControl{
		BroadcastThreshold:                   int64ToIntPointer(broadcastThreshold),
		MulticastThreshold:                   int64ToIntPointer(multicastThreshold),
		TreatTheseTrafficTypesAsOneThreshold: treatTheseTrafficTypesAsOneThreshold,
		UnknownUnicastThreshold:              int64ToIntPointer(unknownUnicastThreshold),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchStormControlItemToBodyRs(state NetworksSwitchStormControlRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchStormControl, is_read bool) NetworksSwitchStormControlRs {
	itemState := NetworksSwitchStormControlRs{
		BroadcastThreshold: func() types.Int64 {
			if response.BroadcastThreshold != nil {
				return types.Int64Value(int64(*response.BroadcastThreshold))
			}
			return types.Int64{}
		}(),
		MulticastThreshold: func() types.Int64 {
			if response.MulticastThreshold != nil {
				return types.Int64Value(int64(*response.MulticastThreshold))
			}
			return types.Int64{}
		}(),
		TreatTheseTrafficTypesAsOneThreshold: StringSliceToList(response.TreatTheseTrafficTypesAsOneThreshold),
		UnknownUnicastThreshold: func() types.Int64 {
			if response.UnknownUnicastThreshold != nil {
				return types.Int64Value(int64(*response.UnknownUnicastThreshold))
			}
			return types.Int64{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchStormControlRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchStormControlRs)
}
