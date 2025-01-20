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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchAlternateManagementInterfaceResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchAlternateManagementInterfaceResource{}
)

func NewNetworksSwitchAlternateManagementInterfaceResource() resource.Resource {
	return &NetworksSwitchAlternateManagementInterfaceResource{}
}

type NetworksSwitchAlternateManagementInterfaceResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchAlternateManagementInterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchAlternateManagementInterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_alternate_management_interface"
}

func (r *NetworksSwitchAlternateManagementInterfaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"protocols": schema.SetAttribute{
				MarkdownDescription: `Can be one or more of the following values: 'radius', 'snmp' or 'syslog'`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"switches": schema.SetNestedAttribute{
				MarkdownDescription: `Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put 'switches' in the body when updating template networks. Also, an empty 'switches' array will remove all previous assignments`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"alternate_management_ip": schema.StringAttribute{
							MarkdownDescription: `Switch alternative management IP. To remove a prior IP setting, provide an empty string`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"gateway": schema.StringAttribute{
							MarkdownDescription: `Switch gateway must be in IP format. Only and must be specified for Polaris switches`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Switch serial number`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"subnet_mask": schema.StringAttribute{
							MarkdownDescription: `Switch subnet mask must be in IP format. Only and must be specified for Polaris switches`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: `Alternate management VLAN, must be between 1 and 4094`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *NetworksSwitchAlternateManagementInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchAlternateManagementInterfaceRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchAlternateManagementInterface(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchAlternateManagementInterface only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchAlternateManagementInterface only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchAlternateManagementInterface(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchAlternateManagementInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchAlternateManagementInterface",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchAlternateManagementInterface(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchAlternateManagementInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchAlternateManagementInterface",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchAlternateManagementInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchAlternateManagementInterfaceRs

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
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchAlternateManagementInterface(vvNetworkID)
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
				"Failure when executing GetNetworkSwitchAlternateManagementInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchAlternateManagementInterface",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSwitchAlternateManagementInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchAlternateManagementInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchAlternateManagementInterfaceRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateNetworkSwitchAlternateManagementInterface(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchAlternateManagementInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchAlternateManagementInterface",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchAlternateManagementInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSwitchAlternateManagementInterface", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchAlternateManagementInterfaceRs struct {
	NetworkID types.String                                                            `tfsdk:"network_id"`
	Enabled   types.Bool                                                              `tfsdk:"enabled"`
	Protocols types.Set                                                               `tfsdk:"protocols"`
	Switches  *[]ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitchesRs `tfsdk:"switches"`
	VLANID    types.Int64                                                             `tfsdk:"vlan_id"`
}

type ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitchesRs struct {
	AlternateManagementIP types.String `tfsdk:"alternate_management_ip"`
	Gateway               types.String `tfsdk:"gateway"`
	Serial                types.String `tfsdk:"serial"`
	SubnetMask            types.String `tfsdk:"subnet_mask"`
}

// FromBody
func (r *NetworksSwitchAlternateManagementInterfaceRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchAlternateManagementInterface {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var protocols []string = nil
	r.Protocols.ElementsAs(ctx, &protocols, false)
	var requestSwitchUpdateNetworkSwitchAlternateManagementInterfaceSwitches []merakigosdk.RequestSwitchUpdateNetworkSwitchAlternateManagementInterfaceSwitches
	if r.Switches != nil {
		for _, rItem1 := range *r.Switches {
			alternateManagementIP := rItem1.AlternateManagementIP.ValueString()
			gateway := rItem1.Gateway.ValueString()
			serial := rItem1.Serial.ValueString()
			subnetMask := rItem1.SubnetMask.ValueString()
			requestSwitchUpdateNetworkSwitchAlternateManagementInterfaceSwitches = append(requestSwitchUpdateNetworkSwitchAlternateManagementInterfaceSwitches, merakigosdk.RequestSwitchUpdateNetworkSwitchAlternateManagementInterfaceSwitches{
				AlternateManagementIP: alternateManagementIP,
				Gateway:               gateway,
				Serial:                serial,
				SubnetMask:            subnetMask,
			})
		}
	}
	vLANID := new(int64)
	if !r.VLANID.IsUnknown() && !r.VLANID.IsNull() {
		*vLANID = r.VLANID.ValueInt64()
	} else {
		vLANID = nil
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchAlternateManagementInterface{
		Enabled:   enabled,
		Protocols: protocols,
		Switches: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchAlternateManagementInterfaceSwitches {
			if len(requestSwitchUpdateNetworkSwitchAlternateManagementInterfaceSwitches) > 0 {
				return &requestSwitchUpdateNetworkSwitchAlternateManagementInterfaceSwitches
			}
			return nil
		}(),
		VLANID: int64ToIntPointer(vLANID),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceItemToBodyRs(state NetworksSwitchAlternateManagementInterfaceRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchAlternateManagementInterface, is_read bool) NetworksSwitchAlternateManagementInterfaceRs {
	itemState := NetworksSwitchAlternateManagementInterfaceRs{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Protocols: StringSliceToSet(response.Protocols),
		Switches: func() *[]ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitchesRs {
			if response.Switches != nil {
				result := make([]ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitchesRs, len(*response.Switches))
				for i, switches := range *response.Switches {
					result[i] = ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitchesRs{
						AlternateManagementIP: types.StringValue(switches.AlternateManagementIP),
						Gateway:               types.StringValue(switches.Gateway),
						Serial:                types.StringValue(switches.Serial),
						SubnetMask:            types.StringValue(switches.SubnetMask),
					}
				}
				return &result
			}
			return nil
		}(),
		VLANID: func() types.Int64 {
			if response.VLANID != nil {
				return types.Int64Value(int64(*response.VLANID))
			}
			return types.Int64{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchAlternateManagementInterfaceRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchAlternateManagementInterfaceRs)
}
