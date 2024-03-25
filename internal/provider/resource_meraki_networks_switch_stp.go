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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
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
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"stp_bridge_priority_response": schema.SetNestedAttribute{
				MarkdownDescription: `STP bridge priority for switches/stacks or switch profiles. An empty array will clear the STP bridge priority settings.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"stacks": schema.SetAttribute{
							MarkdownDescription: `List of stack IDs`,
							Computed:            true,

							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
							Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
						},
						"stp_priority": schema.Int64Attribute{
							MarkdownDescription: `STP priority for switch, stacks, or switch profiles`,
							Computed:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"switch_profiles": schema.SetAttribute{
							MarkdownDescription: `List of switch profile IDs`,
							Computed:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
							ElementType: types.StringType,
						},
						"switches": schema.SetAttribute{
							MarkdownDescription: `List of switch serial numbers`,
							Computed:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
							ElementType: types.StringType,
						},
					},
				},
			},
			"stp_bridge_priority": schema.SetNestedAttribute{
				MarkdownDescription: `STP bridge priority for switches/stacks or switch profiles. An empty array will clear the STP bridge priority settings.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"stacks": schema.SetAttribute{
							MarkdownDescription: `List of stack IDs`,
							Optional:            true,
							Computed:            true,

							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
							Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
						},
						"stp_priority": schema.Int64Attribute{
							MarkdownDescription: `STP priority for switch, stacks, or switch profiles`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"switch_profiles": schema.SetAttribute{
							MarkdownDescription: `List of switch profile IDs`,
							Optional:            true,
							Computed:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
							ElementType: types.StringType,
						},
						"switches": schema.SetAttribute{
							MarkdownDescription: `List of switch serial numbers`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
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
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchStp Error setting item.",
			"Setting item issue",
		)
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchStp Error setting item.",
			"Setting item issue",
		)
		return
	}

	//Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	//Item
	responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchStp(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchStp only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchStp only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Switch.UpdateNetworkSwitchStp(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchStp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchStp",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchStp(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchStp",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchStpItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchStpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchStpRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchStp Error setting item.",
			"Setting item issue",
		)
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchStp Error setting item.",
			"Setting item issue",
		)
		return
	}
	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	// network_id
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
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchStp",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchStpItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSwitchStpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchStpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchStpRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Switch.UpdateNetworkSwitchStp(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchStp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchStp",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchStpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchStpRs struct {
	NetworkID                 types.String                                            `tfsdk:"network_id"`
	RstpEnabled               types.Bool                                              `tfsdk:"rstp_enabled"`
	StpBridgePriority         *[]ResponseSwitchGetNetworkSwitchStpStpBridgePriorityRs `tfsdk:"stp_bridge_priority"`
	StpBridgePriorityResponse *[]ResponseSwitchGetNetworkSwitchStpStpBridgePriorityRs `tfsdk:"stp_bridge_priority_response"`
}

type ResponseSwitchGetNetworkSwitchStpStpBridgePriorityRs struct {
	StpPriority    types.Int64 `tfsdk:"stp_priority"`
	Switches       types.Set   `tfsdk:"switches"`
	Stacks         types.Set   `tfsdk:"stacks"`
	SwitchProfiles types.Set   `tfsdk:"switch_profiles"`
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
		StpBridgePriorityResponse: func() *[]ResponseSwitchGetNetworkSwitchStpStpBridgePriorityRs {
			if response.StpBridgePriority != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStpStpBridgePriorityRs, len(*response.StpBridgePriority))
				var localStacks basetypes.SetValue
				var localSwitchProfiles basetypes.SetValue
				b := *state.StpBridgePriority
				for i, stpBridgePriority := range *response.StpBridgePriority {
					if len(b) > i {
						localStacks = b[i].Stacks
						localSwitchProfiles = b[i].SwitchProfiles
					} else {
						localStacks = types.SetNull(types.StringType)
						localSwitchProfiles = types.SetNull(types.StringType)
					}
					result[i] = ResponseSwitchGetNetworkSwitchStpStpBridgePriorityRs{
						StpPriority: func() types.Int64 {
							if stpBridgePriority.StpPriority != nil {
								return types.Int64Value(int64(*stpBridgePriority.StpPriority))
							}
							return types.Int64{}
						}(),
						Switches:       StringSliceToSet(stpBridgePriority.Switches),
						Stacks:         localStacks,
						SwitchProfiles: localSwitchProfiles,
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchStpStpBridgePriorityRs{}
		}(),
	}

	// a := *itemState.StpBridgePriority
	// a[0].Stacks = types.SetNull(types.StringType)
	// a[0].SwitchProfiles = types.SetNull(types.StringType)
	// a[0].Switches = types.SetNull(types.StringType)
	itemState.StpBridgePriority = state.StpBridgePriority
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchStpRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchStpRs)
}
