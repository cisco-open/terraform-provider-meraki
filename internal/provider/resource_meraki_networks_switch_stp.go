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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchStpResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchStpResource{}
)

func NewNetworksSwitchStpResource() resource.Resource {
	return &NetworksSwitchStpResource{}
}

type NetworksSwitchStpResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchStpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchStpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stp"
}

func (r *NetworksSwitchStpResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"rstp_enabled": schema.BoolAttribute{
				MarkdownDescription: `The spanning tree protocol status in network`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"stp_bridge_priority": schema.ListNestedAttribute{
				MarkdownDescription: `STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings.`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"stacks": schema.ListAttribute{
							MarkdownDescription: `List of stack IDs`,
							Optional:            true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"stp_priority": schema.Int64Attribute{
							MarkdownDescription: `STP priority for switch, stacks, or switch templates`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"switch_profiles": schema.ListAttribute{
							MarkdownDescription: `List of switch template IDs`,
							Optional:            true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"switches": schema.ListAttribute{
							MarkdownDescription: `List of switch serial numbers`,
							Optional:            true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

func (r *NetworksSwitchStpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchStpRs

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
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchStp(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchStp",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchStp",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksSwitchStpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchStpRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchStp(vvNetworkID)
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
				"Failure when executing GetNetworkSwitchStp",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchStp",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchStpItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksSwitchStpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchStpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksSwitchStpRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchStp(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchStp",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchStp",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksSwitchStpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSwitchStp", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchStpRs struct {
	NetworkID         types.String                                            `tfsdk:"network_id"`
	RstpEnabled       types.Bool                                              `tfsdk:"rstp_enabled"`
	StpBridgePriority *[]ResponseSwitchGetNetworkSwitchStpStpBridgePriorityRs `tfsdk:"stp_bridge_priority"`
}

type ResponseSwitchGetNetworkSwitchStpStpBridgePriorityRs struct {
	Stacks         types.List  `tfsdk:"stacks"`
	StpPriority    types.Int64 `tfsdk:"stp_priority"`
	SwitchProfiles types.List  `tfsdk:"switch_profiles"`
	Switches       types.List  `tfsdk:"switches"`
}

// FromBody
func (r *NetworksSwitchStpRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchStp {
	rstpEnabled := new(bool)
	if !r.RstpEnabled.IsUnknown() && !r.RstpEnabled.IsNull() {
		*rstpEnabled = r.RstpEnabled.ValueBool()
	} else {
		rstpEnabled = nil
	}
	var requestSwitchUpdateNetworkSwitchStpStpBridgePriority []merakigosdk.RequestSwitchUpdateNetworkSwitchStpStpBridgePriority

	if r.StpBridgePriority != nil {
		for _, rItem1 := range *r.StpBridgePriority {

			var stacks []string = nil
			rItem1.Stacks.ElementsAs(ctx, &stacks, false)
			stpPriority := func() *int64 {
				if !rItem1.StpPriority.IsUnknown() && !rItem1.StpPriority.IsNull() {
					return rItem1.StpPriority.ValueInt64Pointer()
				}
				return nil
			}()

			var switchProfiles []string = nil
			rItem1.SwitchProfiles.ElementsAs(ctx, &switchProfiles, false)

			var switches []string = nil
			rItem1.Switches.ElementsAs(ctx, &switches, false)
			requestSwitchUpdateNetworkSwitchStpStpBridgePriority = append(requestSwitchUpdateNetworkSwitchStpStpBridgePriority, merakigosdk.RequestSwitchUpdateNetworkSwitchStpStpBridgePriority{
				Stacks:         stacks,
				StpPriority:    int64ToIntPointer(stpPriority),
				SwitchProfiles: switchProfiles,
				Switches:       switches,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchStp{
		RstpEnabled: rstpEnabled,
		StpBridgePriority: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchStpStpBridgePriority {
			if len(requestSwitchUpdateNetworkSwitchStpStpBridgePriority) > 0 {
				return &requestSwitchUpdateNetworkSwitchStpStpBridgePriority
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchStpItemToBodyRs(state NetworksSwitchStpRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchStp, is_read bool) NetworksSwitchStpRs {
	itemState := NetworksSwitchStpRs{
		RstpEnabled: func() types.Bool {
			if response.RstpEnabled != nil {
				return types.BoolValue(*response.RstpEnabled)
			}
			return types.Bool{}
		}(),
		StpBridgePriority: func() *[]ResponseSwitchGetNetworkSwitchStpStpBridgePriorityRs {
			if response.StpBridgePriority != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStpStpBridgePriorityRs, len(*response.StpBridgePriority))
				for i, stpBridgePriority := range *response.StpBridgePriority {
					result[i] = ResponseSwitchGetNetworkSwitchStpStpBridgePriorityRs{
						Stacks: StringSliceToList(stpBridgePriority.Stacks),
						StpPriority: func() types.Int64 {
							if stpBridgePriority.StpPriority != nil {
								return types.Int64Value(int64(*stpBridgePriority.StpPriority))
							}
							return types.Int64{}
						}(),
						SwitchProfiles: StringSliceToList(stpBridgePriority.SwitchProfiles),
						Switches:       StringSliceToList(stpBridgePriority.Switches),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchStpRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchStpRs)
}
