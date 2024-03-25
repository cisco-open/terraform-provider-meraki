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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchRoutingMulticastResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchRoutingMulticastResource{}
)

func NewNetworksSwitchRoutingMulticastResource() resource.Resource {
	return &NetworksSwitchRoutingMulticastResource{}
}

type NetworksSwitchRoutingMulticastResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchRoutingMulticastResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchRoutingMulticastResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_routing_multicast"
}

func (r *NetworksSwitchRoutingMulticastResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"default_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Default multicast setting for entire network. IGMP snooping and Flood unknown multicast traffic settings are enabled by default.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"flood_unknown_multicast_traffic_enabled": schema.BoolAttribute{
						MarkdownDescription: `Flood unknown multicast traffic setting for entire network`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"igmp_snooping_enabled": schema.BoolAttribute{
						MarkdownDescription: `IGMP snooping setting for entire network`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"overrides": schema.SetNestedAttribute{
				MarkdownDescription: `Array of paired switches/stacks/profiles and corresponding multicast settings. An empty array will clear the multicast settings.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"flood_unknown_multicast_traffic_enabled": schema.BoolAttribute{
							MarkdownDescription: `Flood unknown multicast traffic setting for switches, switch stacks or switch profiles`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"igmp_snooping_enabled": schema.BoolAttribute{
							MarkdownDescription: `IGMP snooping setting for switches, switch stacks or switch profiles`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"stacks": schema.SetAttribute{
							MarkdownDescription: `List of switch stack ids for non-template network`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"switch_profiles": schema.SetAttribute{
							MarkdownDescription: `List of switch profiles ids for template network`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"switches": schema.SetAttribute{
							MarkdownDescription: `List of switch serials for non-template network`,
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

func (r *NetworksSwitchRoutingMulticastResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchRoutingMulticastRs

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
	responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchRoutingMulticast(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchRoutingMulticast only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchRoutingMulticast only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Switch.UpdateNetworkSwitchRoutingMulticast(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchRoutingMulticast",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchRoutingMulticast",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchRoutingMulticast(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchRoutingMulticast",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchRoutingMulticast",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchRoutingMulticastItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchRoutingMulticastResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchRoutingMulticastRs

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
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchRoutingMulticast(vvNetworkID)
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
				"Failure when executing GetNetworkSwitchRoutingMulticast",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchRoutingMulticast",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchRoutingMulticastItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSwitchRoutingMulticastResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchRoutingMulticastResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchRoutingMulticastRs
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
	restyResp2, err := r.client.Switch.UpdateNetworkSwitchRoutingMulticast(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchRoutingMulticast",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchRoutingMulticast",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchRoutingMulticastResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchRoutingMulticastRs struct {
	NetworkID       types.String                                                     `tfsdk:"network_id"`
	DefaultSettings *ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettingsRs `tfsdk:"default_settings"`
	Overrides       *[]ResponseSwitchGetNetworkSwitchRoutingMulticastOverridesRs     `tfsdk:"overrides"`
}

type ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettingsRs struct {
	FloodUnknownMulticastTrafficEnabled types.Bool `tfsdk:"flood_unknown_multicast_traffic_enabled"`
	IgmpSnoopingEnabled                 types.Bool `tfsdk:"igmp_snooping_enabled"`
}

type ResponseSwitchGetNetworkSwitchRoutingMulticastOverridesRs struct {
	FloodUnknownMulticastTrafficEnabled types.Bool `tfsdk:"flood_unknown_multicast_traffic_enabled"`
	IgmpSnoopingEnabled                 types.Bool `tfsdk:"igmp_snooping_enabled"`
	Switches                            types.Set  `tfsdk:"switches"`
	Stacks                              types.Set  `tfsdk:"stacks"`
	SwitchProfiles                      types.Set  `tfsdk:"switch_profiles"`
}

// FromBody
func (r *NetworksSwitchRoutingMulticastRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingMulticast {
	var requestSwitchUpdateNetworkSwitchRoutingMulticastDefaultSettings *merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingMulticastDefaultSettings
	if r.DefaultSettings != nil {
		floodUnknownMulticastTrafficEnabled := func() *bool {
			if !r.DefaultSettings.FloodUnknownMulticastTrafficEnabled.IsUnknown() && !r.DefaultSettings.FloodUnknownMulticastTrafficEnabled.IsNull() {
				return r.DefaultSettings.FloodUnknownMulticastTrafficEnabled.ValueBoolPointer()
			}
			return nil
		}()
		igmpSnoopingEnabled := func() *bool {
			if !r.DefaultSettings.IgmpSnoopingEnabled.IsUnknown() && !r.DefaultSettings.IgmpSnoopingEnabled.IsNull() {
				return r.DefaultSettings.IgmpSnoopingEnabled.ValueBoolPointer()
			}
			return nil
		}()
		requestSwitchUpdateNetworkSwitchRoutingMulticastDefaultSettings = &merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingMulticastDefaultSettings{
			FloodUnknownMulticastTrafficEnabled: floodUnknownMulticastTrafficEnabled,
			IgmpSnoopingEnabled:                 igmpSnoopingEnabled,
		}
	}
	var requestSwitchUpdateNetworkSwitchRoutingMulticastOverrides []merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingMulticastOverrides
	if r.Overrides != nil {
		for _, rItem1 := range *r.Overrides {
			floodUnknownMulticastTrafficEnabled := func() *bool {
				if !rItem1.FloodUnknownMulticastTrafficEnabled.IsUnknown() && !rItem1.FloodUnknownMulticastTrafficEnabled.IsNull() {
					return rItem1.FloodUnknownMulticastTrafficEnabled.ValueBoolPointer()
				}
				return nil
			}()
			igmpSnoopingEnabled := func() *bool {
				if !rItem1.IgmpSnoopingEnabled.IsUnknown() && !rItem1.IgmpSnoopingEnabled.IsNull() {
					return rItem1.IgmpSnoopingEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var stacks []string = nil

			rItem1.Stacks.ElementsAs(ctx, &stacks, false)
			var switchProfiles []string = nil

			rItem1.SwitchProfiles.ElementsAs(ctx, &switchProfiles, false)
			var switches []string = nil

			rItem1.Switches.ElementsAs(ctx, &switches, false)
			requestSwitchUpdateNetworkSwitchRoutingMulticastOverrides = append(requestSwitchUpdateNetworkSwitchRoutingMulticastOverrides, merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingMulticastOverrides{
				FloodUnknownMulticastTrafficEnabled: floodUnknownMulticastTrafficEnabled,
				IgmpSnoopingEnabled:                 igmpSnoopingEnabled,
				Stacks:                              stacks,
				SwitchProfiles:                      switchProfiles,
				Switches:                            switches,
			})
		}
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingMulticast{
		DefaultSettings: requestSwitchUpdateNetworkSwitchRoutingMulticastDefaultSettings,
		Overrides: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingMulticastOverrides {
			if len(requestSwitchUpdateNetworkSwitchRoutingMulticastOverrides) > 0 {
				return &requestSwitchUpdateNetworkSwitchRoutingMulticastOverrides
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchRoutingMulticastItemToBodyRs(state NetworksSwitchRoutingMulticastRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchRoutingMulticast, is_read bool) NetworksSwitchRoutingMulticastRs {
	itemState := NetworksSwitchRoutingMulticastRs{
		DefaultSettings: func() *ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettingsRs {
			if response.DefaultSettings != nil {
				return &ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettingsRs{
					FloodUnknownMulticastTrafficEnabled: func() types.Bool {
						if response.DefaultSettings.FloodUnknownMulticastTrafficEnabled != nil {
							return types.BoolValue(*response.DefaultSettings.FloodUnknownMulticastTrafficEnabled)
						}
						return types.Bool{}
					}(),
					IgmpSnoopingEnabled: func() types.Bool {
						if response.DefaultSettings.IgmpSnoopingEnabled != nil {
							return types.BoolValue(*response.DefaultSettings.IgmpSnoopingEnabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettingsRs{}
		}(),
		Overrides: func() *[]ResponseSwitchGetNetworkSwitchRoutingMulticastOverridesRs {
			if response.Overrides != nil {
				result := make([]ResponseSwitchGetNetworkSwitchRoutingMulticastOverridesRs, len(*response.Overrides))
				for i, overrides := range *response.Overrides {
					result[i] = ResponseSwitchGetNetworkSwitchRoutingMulticastOverridesRs{
						FloodUnknownMulticastTrafficEnabled: func() types.Bool {
							if overrides.FloodUnknownMulticastTrafficEnabled != nil {
								return types.BoolValue(*overrides.FloodUnknownMulticastTrafficEnabled)
							}
							return types.Bool{}
						}(),
						IgmpSnoopingEnabled: func() types.Bool {
							if overrides.IgmpSnoopingEnabled != nil {
								return types.BoolValue(*overrides.IgmpSnoopingEnabled)
							}
							return types.Bool{}
						}(),
						Switches: StringSliceToSet(overrides.Switches),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchRoutingMulticastOverridesRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchRoutingMulticastRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchRoutingMulticastRs)
}
