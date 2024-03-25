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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchMtuResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchMtuResource{}
)

func NewNetworksSwitchMtuResource() resource.Resource {
	return &NetworksSwitchMtuResource{}
}

type NetworksSwitchMtuResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchMtuResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchMtuResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_mtu"
}

func (r *NetworksSwitchMtuResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"default_mtu_size": schema.Int64Attribute{
				MarkdownDescription: `MTU size for the entire network. Default value is 9578.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"overrides": schema.SetNestedAttribute{
				MarkdownDescription: `Override MTU size for individual switches or switch profiles.
      An empty array will clear overrides.`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"mtu_size": schema.Int64Attribute{
							MarkdownDescription: `MTU size for the switches or switch profiles.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"switch_profiles": schema.SetAttribute{
							MarkdownDescription: `List of switch profile IDs. Applicable only for template network.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"switches": schema.SetAttribute{
							MarkdownDescription: `List of switch serials. Applicable only for switch network.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
					},
				},
			},
		},
	}
}

func (r *NetworksSwitchMtuResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchMtuRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchMtu(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchMtu only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchMtu only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Switch.UpdateNetworkSwitchMtu(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchMtu",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchMtu",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchMtu(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchMtu",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchMtu",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchMtuItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchMtuResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchMtuRs

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
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchMtu(vvNetworkID)
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
				"Failure when executing GetNetworkSwitchMtu",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchMtu",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchMtuItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSwitchMtuResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchMtuResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchMtuRs
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
	restyResp2, err := r.client.Switch.UpdateNetworkSwitchMtu(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchMtu",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchMtu",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchMtuResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchMtuRs struct {
	NetworkID      types.String                                    `tfsdk:"network_id"`
	DefaultMtuSize types.Int64                                     `tfsdk:"default_mtu_size"`
	Overrides      *[]ResponseSwitchGetNetworkSwitchMtuOverridesRs `tfsdk:"overrides"`
}

type ResponseSwitchGetNetworkSwitchMtuOverridesRs struct {
	MtuSize        types.Int64 `tfsdk:"mtu_size"`
	SwitchProfiles types.Set   `tfsdk:"switch_profiles"`
	Switches       types.Set   `tfsdk:"switches"`
}

// FromBody
func (r *NetworksSwitchMtuRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchMtu {
	defaultMtuSize := new(int64)
	if !r.DefaultMtuSize.IsUnknown() && !r.DefaultMtuSize.IsNull() {
		*defaultMtuSize = r.DefaultMtuSize.ValueInt64()
	} else {
		defaultMtuSize = nil
	}
	var requestSwitchUpdateNetworkSwitchMtuOverrides []merakigosdk.RequestSwitchUpdateNetworkSwitchMtuOverrides
	if r.Overrides != nil {
		for _, rItem1 := range *r.Overrides {
			mtuSize := func() *int64 {
				if !rItem1.MtuSize.IsUnknown() && !rItem1.MtuSize.IsNull() {
					return rItem1.MtuSize.ValueInt64Pointer()
				}
				return nil
			}()
			var switchProfiles []string = nil

			rItem1.SwitchProfiles.ElementsAs(ctx, &switchProfiles, false)
			var switches []string = nil

			rItem1.Switches.ElementsAs(ctx, &switches, false)
			requestSwitchUpdateNetworkSwitchMtuOverrides = append(requestSwitchUpdateNetworkSwitchMtuOverrides, merakigosdk.RequestSwitchUpdateNetworkSwitchMtuOverrides{
				MtuSize:        int64ToIntPointer(mtuSize),
				SwitchProfiles: switchProfiles,
				Switches:       switches,
			})
		}
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchMtu{
		DefaultMtuSize: int64ToIntPointer(defaultMtuSize),
		Overrides: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchMtuOverrides {
			if len(requestSwitchUpdateNetworkSwitchMtuOverrides) > 0 {
				return &requestSwitchUpdateNetworkSwitchMtuOverrides
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchMtuItemToBodyRs(state NetworksSwitchMtuRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchMtu, is_read bool) NetworksSwitchMtuRs {
	itemState := NetworksSwitchMtuRs{
		DefaultMtuSize: func() types.Int64 {
			if response.DefaultMtuSize != nil {
				return types.Int64Value(int64(*response.DefaultMtuSize))
			}
			return types.Int64{}
		}(),
		Overrides: func() *[]ResponseSwitchGetNetworkSwitchMtuOverridesRs {
			if response.Overrides != nil {
				result := make([]ResponseSwitchGetNetworkSwitchMtuOverridesRs, len(*response.Overrides))
				for i, overrides := range *response.Overrides {
					result[i] = ResponseSwitchGetNetworkSwitchMtuOverridesRs{
						MtuSize: func() types.Int64 {
							if overrides.MtuSize != nil {
								return types.Int64Value(int64(*overrides.MtuSize))
							}
							return types.Int64{}
						}(),
						SwitchProfiles: StringSliceToSet(overrides.SwitchProfiles),
						Switches:       StringSliceToSet(overrides.Switches),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchMtuOverridesRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchMtuRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchMtuRs)
}
