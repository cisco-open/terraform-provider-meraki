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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesSwitchRoutingInterfacesResource{}
	_ resource.ResourceWithConfigure = &DevicesSwitchRoutingInterfacesResource{}
)

func NewDevicesSwitchRoutingInterfacesResource() resource.Resource {
	return &DevicesSwitchRoutingInterfacesResource{}
}

type DevicesSwitchRoutingInterfacesResource struct {
	client *merakigosdk.Client
}

func (r *DevicesSwitchRoutingInterfacesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesSwitchRoutingInterfacesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_switch_routing_interfaces"
}

func (r *DevicesSwitchRoutingInterfacesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"default_gateway": schema.StringAttribute{
				MarkdownDescription: `IPv4 default gateway`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: `The id`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"interface_ip": schema.StringAttribute{
				MarkdownDescription: `IPv4 address`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ipv6": schema.SingleNestedAttribute{
				MarkdownDescription: `IPv6 addressing`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"address": schema.StringAttribute{
						MarkdownDescription: `IPv6 address`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"assignment_mode": schema.StringAttribute{
						MarkdownDescription: `Assignment mode`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"gateway": schema.StringAttribute{
						MarkdownDescription: `IPv6 gateway`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"prefix": schema.StringAttribute{
						MarkdownDescription: `IPv6 subnet`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"multicast_routing": schema.StringAttribute{
				MarkdownDescription: `Multicast routing status`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ospf_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `IPv4 OSPF Settings`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"area": schema.StringAttribute{
						MarkdownDescription: `Area id`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"cost": schema.Int64Attribute{
						MarkdownDescription: `OSPF Cost`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"is_passive_enabled": schema.BoolAttribute{
						MarkdownDescription: `Disable sending Hello packets on this interface's IPv4 area`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"ospf_v3": schema.SingleNestedAttribute{
				MarkdownDescription: `IPv6 OSPF Settings`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"area": schema.StringAttribute{
						MarkdownDescription: `Area id`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"cost": schema.Int64Attribute{
						MarkdownDescription: `OSPF Cost`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"is_passive_enabled": schema.BoolAttribute{
						MarkdownDescription: `Disable sending Hello packets on this interface's IPv6 area`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"subnet": schema.StringAttribute{
				MarkdownDescription: `IPv4 subnet`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: `VLAN id`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['interfaceId']

func (r *DevicesSwitchRoutingInterfacesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesSwitchRoutingInterfacesRs

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
	vvSerial := data.Serial.ValueString()
	// serial
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Switch.GetDeviceSwitchRoutingInterfaces(vvSerial)
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchRoutingInterfaces",
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
			vvInterfaceID, ok := result2["InterfaceID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter InterfaceID",
					"Error",
				)
				return
			}
			r.client.Switch.UpdateDeviceSwitchRoutingInterface(vvSerial, vvInterfaceID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Switch.GetDeviceSwitchRoutingInterface(vvSerial, vvInterfaceID)
			if responseVerifyItem2 != nil {
				data = ResponseSwitchGetDeviceSwitchRoutingInterfaceItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Switch.CreateDeviceSwitchRoutingInterface(vvSerial, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateDeviceSwitchRoutingInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateDeviceSwitchRoutingInterface",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Switch.GetDeviceSwitchRoutingInterfaces(vvSerial)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchRoutingInterfaces",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceSwitchRoutingInterfaces",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvInterfaceID, ok := result2["InterfaceID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter InterfaceID",
				"Error",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Switch.GetDeviceSwitchRoutingInterface(vvSerial, vvInterfaceID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseSwitchGetDeviceSwitchRoutingInterfaceItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetDeviceSwitchRoutingInterface",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchRoutingInterface",
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

func (r *DevicesSwitchRoutingInterfacesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesSwitchRoutingInterfacesRs

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

	vvSerial := data.Serial.ValueString()
	// serial
	vvInterfaceID := data.InterfaceID.ValueString()
	// interface_id
	responseGet, restyRespGet, err := r.client.Switch.GetDeviceSwitchRoutingInterface(vvSerial, vvInterfaceID)
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
				"Failure when executing GetDeviceSwitchRoutingInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceSwitchRoutingInterface",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetDeviceSwitchRoutingInterfaceItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesSwitchRoutingInterfacesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("interface_id"), idParts[1])...)
}

func (r *DevicesSwitchRoutingInterfacesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesSwitchRoutingInterfacesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	// serial
	vvInterfaceID := data.InterfaceID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Switch.UpdateDeviceSwitchRoutingInterface(vvSerial, vvInterfaceID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceSwitchRoutingInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceSwitchRoutingInterface",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesSwitchRoutingInterfacesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DevicesSwitchRoutingInterfacesRs
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

	vvSerial := state.Serial.ValueString()
	vvInterfaceID := state.InterfaceID.ValueString()
	_, err := r.client.Switch.DeleteDeviceSwitchRoutingInterface(vvSerial, vvInterfaceID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteDeviceSwitchRoutingInterface", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type DevicesSwitchRoutingInterfacesRs struct {
	Serial           types.String                                                 `tfsdk:"serial"`
	InterfaceID      types.String                                                 `tfsdk:"interface_id"`
	DefaultGateway   types.String                                                 `tfsdk:"default_gateway"`
	InterfaceIP      types.String                                                 `tfsdk:"interface_ip"`
	IPv6             *ResponseSwitchGetDeviceSwitchRoutingInterfaceIpv6Rs         `tfsdk:"ipv6"`
	MulticastRouting types.String                                                 `tfsdk:"multicast_routing"`
	Name             types.String                                                 `tfsdk:"name"`
	OspfSettings     *ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfSettingsRs `tfsdk:"ospf_settings"`
	OspfV3           *ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfV3Rs       `tfsdk:"ospf_v3"`
	Subnet           types.String                                                 `tfsdk:"subnet"`
	VLANID           types.Int64                                                  `tfsdk:"vlan_id"`
}

type ResponseSwitchGetDeviceSwitchRoutingInterfaceIpv6Rs struct {
	Address        types.String `tfsdk:"address"`
	AssignmentMode types.String `tfsdk:"assignment_mode"`
	Gateway        types.String `tfsdk:"gateway"`
	Prefix         types.String `tfsdk:"prefix"`
}

type ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfSettingsRs struct {
	Area             types.String `tfsdk:"area"`
	Cost             types.Int64  `tfsdk:"cost"`
	IsPassiveEnabled types.Bool   `tfsdk:"is_passive_enabled"`
}

type ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfV3Rs struct {
	Area             types.String `tfsdk:"area"`
	Cost             types.Int64  `tfsdk:"cost"`
	IsPassiveEnabled types.Bool   `tfsdk:"is_passive_enabled"`
}

// FromBody
func (r *DevicesSwitchRoutingInterfacesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCreateDeviceSwitchRoutingInterface {
	emptyString := ""
	defaultGateway := new(string)
	if !r.DefaultGateway.IsUnknown() && !r.DefaultGateway.IsNull() {
		*defaultGateway = r.DefaultGateway.ValueString()
	} else {
		defaultGateway = &emptyString
	}
	interfaceIP := new(string)
	if !r.InterfaceIP.IsUnknown() && !r.InterfaceIP.IsNull() {
		*interfaceIP = r.InterfaceIP.ValueString()
	} else {
		interfaceIP = &emptyString
	}
	var requestSwitchCreateDeviceSwitchRoutingInterfaceIPv6 *merakigosdk.RequestSwitchCreateDeviceSwitchRoutingInterfaceIPv6
	if r.IPv6 != nil {
		address := r.IPv6.Address.ValueString()
		assignmentMode := r.IPv6.AssignmentMode.ValueString()
		gateway := r.IPv6.Gateway.ValueString()
		prefix := r.IPv6.Prefix.ValueString()
		requestSwitchCreateDeviceSwitchRoutingInterfaceIPv6 = &merakigosdk.RequestSwitchCreateDeviceSwitchRoutingInterfaceIPv6{
			Address:        address,
			AssignmentMode: assignmentMode,
			Gateway:        gateway,
			Prefix:         prefix,
		}
	}
	multicastRouting := new(string)
	if !r.MulticastRouting.IsUnknown() && !r.MulticastRouting.IsNull() {
		*multicastRouting = r.MulticastRouting.ValueString()
	} else {
		multicastRouting = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestSwitchCreateDeviceSwitchRoutingInterfaceOspfSettings *merakigosdk.RequestSwitchCreateDeviceSwitchRoutingInterfaceOspfSettings
	if r.OspfSettings != nil {
		area := r.OspfSettings.Area.ValueString()
		cost := func() *int64 {
			if !r.OspfSettings.Cost.IsUnknown() && !r.OspfSettings.Cost.IsNull() {
				return r.OspfSettings.Cost.ValueInt64Pointer()
			}
			return nil
		}()
		isPassiveEnabled := func() *bool {
			if !r.OspfSettings.IsPassiveEnabled.IsUnknown() && !r.OspfSettings.IsPassiveEnabled.IsNull() {
				return r.OspfSettings.IsPassiveEnabled.ValueBoolPointer()
			}
			return nil
		}()
		requestSwitchCreateDeviceSwitchRoutingInterfaceOspfSettings = &merakigosdk.RequestSwitchCreateDeviceSwitchRoutingInterfaceOspfSettings{
			Area:             area,
			Cost:             int64ToIntPointer(cost),
			IsPassiveEnabled: isPassiveEnabled,
		}
	}
	var requestSwitchCreateDeviceSwitchRoutingInterfaceOspfV3 *merakigosdk.RequestSwitchCreateDeviceSwitchRoutingInterfaceOspfV3
	if r.OspfV3 != nil {
		area := r.OspfV3.Area.ValueString()
		cost := func() *int64 {
			if !r.OspfV3.Cost.IsUnknown() && !r.OspfV3.Cost.IsNull() {
				return r.OspfV3.Cost.ValueInt64Pointer()
			}
			return nil
		}()
		isPassiveEnabled := func() *bool {
			if !r.OspfV3.IsPassiveEnabled.IsUnknown() && !r.OspfV3.IsPassiveEnabled.IsNull() {
				return r.OspfV3.IsPassiveEnabled.ValueBoolPointer()
			}
			return nil
		}()
		requestSwitchCreateDeviceSwitchRoutingInterfaceOspfV3 = &merakigosdk.RequestSwitchCreateDeviceSwitchRoutingInterfaceOspfV3{
			Area:             area,
			Cost:             int64ToIntPointer(cost),
			IsPassiveEnabled: isPassiveEnabled,
		}
	}
	subnet := new(string)
	if !r.Subnet.IsUnknown() && !r.Subnet.IsNull() {
		*subnet = r.Subnet.ValueString()
	} else {
		subnet = &emptyString
	}
	vLANID := new(int64)
	if !r.VLANID.IsUnknown() && !r.VLANID.IsNull() {
		*vLANID = r.VLANID.ValueInt64()
	} else {
		vLANID = nil
	}
	out := merakigosdk.RequestSwitchCreateDeviceSwitchRoutingInterface{
		DefaultGateway:   *defaultGateway,
		InterfaceIP:      *interfaceIP,
		IPv6:             requestSwitchCreateDeviceSwitchRoutingInterfaceIPv6,
		MulticastRouting: *multicastRouting,
		Name:             *name,
		OspfSettings:     requestSwitchCreateDeviceSwitchRoutingInterfaceOspfSettings,
		OspfV3:           requestSwitchCreateDeviceSwitchRoutingInterfaceOspfV3,
		Subnet:           *subnet,
		VLANID:           int64ToIntPointer(vLANID),
	}
	return &out
}
func (r *DevicesSwitchRoutingInterfacesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterface {
	emptyString := ""
	defaultGateway := new(string)
	if !r.DefaultGateway.IsUnknown() && !r.DefaultGateway.IsNull() {
		*defaultGateway = r.DefaultGateway.ValueString()
	} else {
		defaultGateway = &emptyString
	}
	interfaceIP := new(string)
	if !r.InterfaceIP.IsUnknown() && !r.InterfaceIP.IsNull() {
		*interfaceIP = r.InterfaceIP.ValueString()
	} else {
		interfaceIP = &emptyString
	}
	var requestSwitchUpdateDeviceSwitchRoutingInterfaceIPv6 *merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceIPv6
	if r.IPv6 != nil {
		address := r.IPv6.Address.ValueString()
		assignmentMode := r.IPv6.AssignmentMode.ValueString()
		gateway := r.IPv6.Gateway.ValueString()
		prefix := r.IPv6.Prefix.ValueString()
		requestSwitchUpdateDeviceSwitchRoutingInterfaceIPv6 = &merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceIPv6{
			Address:        address,
			AssignmentMode: assignmentMode,
			Gateway:        gateway,
			Prefix:         prefix,
		}
	}
	multicastRouting := new(string)
	if !r.MulticastRouting.IsUnknown() && !r.MulticastRouting.IsNull() {
		*multicastRouting = r.MulticastRouting.ValueString()
	} else {
		multicastRouting = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestSwitchUpdateDeviceSwitchRoutingInterfaceOspfSettings *merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceOspfSettings
	if r.OspfSettings != nil {
		area := r.OspfSettings.Area.ValueString()
		cost := func() *int64 {
			if !r.OspfSettings.Cost.IsUnknown() && !r.OspfSettings.Cost.IsNull() {
				return r.OspfSettings.Cost.ValueInt64Pointer()
			}
			return nil
		}()
		isPassiveEnabled := func() *bool {
			if !r.OspfSettings.IsPassiveEnabled.IsUnknown() && !r.OspfSettings.IsPassiveEnabled.IsNull() {
				return r.OspfSettings.IsPassiveEnabled.ValueBoolPointer()
			}
			return nil
		}()
		requestSwitchUpdateDeviceSwitchRoutingInterfaceOspfSettings = &merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceOspfSettings{
			Area:             area,
			Cost:             int64ToIntPointer(cost),
			IsPassiveEnabled: isPassiveEnabled,
		}
	}
	var requestSwitchUpdateDeviceSwitchRoutingInterfaceOspfV3 *merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceOspfV3
	if r.OspfV3 != nil {
		area := r.OspfV3.Area.ValueString()
		cost := func() *int64 {
			if !r.OspfV3.Cost.IsUnknown() && !r.OspfV3.Cost.IsNull() {
				return r.OspfV3.Cost.ValueInt64Pointer()
			}
			return nil
		}()
		isPassiveEnabled := func() *bool {
			if !r.OspfV3.IsPassiveEnabled.IsUnknown() && !r.OspfV3.IsPassiveEnabled.IsNull() {
				return r.OspfV3.IsPassiveEnabled.ValueBoolPointer()
			}
			return nil
		}()
		requestSwitchUpdateDeviceSwitchRoutingInterfaceOspfV3 = &merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterfaceOspfV3{
			Area:             area,
			Cost:             int64ToIntPointer(cost),
			IsPassiveEnabled: isPassiveEnabled,
		}
	}
	subnet := new(string)
	if !r.Subnet.IsUnknown() && !r.Subnet.IsNull() {
		*subnet = r.Subnet.ValueString()
	} else {
		subnet = &emptyString
	}
	vLANID := new(int64)
	if !r.VLANID.IsUnknown() && !r.VLANID.IsNull() {
		*vLANID = r.VLANID.ValueInt64()
	} else {
		vLANID = nil
	}
	out := merakigosdk.RequestSwitchUpdateDeviceSwitchRoutingInterface{
		DefaultGateway:   *defaultGateway,
		InterfaceIP:      *interfaceIP,
		IPv6:             requestSwitchUpdateDeviceSwitchRoutingInterfaceIPv6,
		MulticastRouting: *multicastRouting,
		Name:             *name,
		OspfSettings:     requestSwitchUpdateDeviceSwitchRoutingInterfaceOspfSettings,
		OspfV3:           requestSwitchUpdateDeviceSwitchRoutingInterfaceOspfV3,
		Subnet:           *subnet,
		VLANID:           int64ToIntPointer(vLANID),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetDeviceSwitchRoutingInterfaceItemToBodyRs(state DevicesSwitchRoutingInterfacesRs, response *merakigosdk.ResponseSwitchGetDeviceSwitchRoutingInterface, is_read bool) DevicesSwitchRoutingInterfacesRs {
	itemState := DevicesSwitchRoutingInterfacesRs{
		DefaultGateway: types.StringValue(response.DefaultGateway),
		InterfaceID:    types.StringValue(response.InterfaceID),
		InterfaceIP:    types.StringValue(response.InterfaceIP),
		IPv6: func() *ResponseSwitchGetDeviceSwitchRoutingInterfaceIpv6Rs {
			if response.IPv6 != nil {
				return &ResponseSwitchGetDeviceSwitchRoutingInterfaceIpv6Rs{
					Address:        types.StringValue(response.IPv6.Address),
					AssignmentMode: types.StringValue(response.IPv6.AssignmentMode),
					Gateway:        types.StringValue(response.IPv6.Gateway),
					Prefix:         types.StringValue(response.IPv6.Prefix),
				}
			}
			return &ResponseSwitchGetDeviceSwitchRoutingInterfaceIpv6Rs{}
		}(),
		MulticastRouting: types.StringValue(response.MulticastRouting),
		Name:             types.StringValue(response.Name),
		OspfSettings: func() *ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfSettingsRs {
			if response.OspfSettings != nil {
				return &ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfSettingsRs{
					Area: types.StringValue(response.OspfSettings.Area),
					Cost: func() types.Int64 {
						if response.OspfSettings.Cost != nil {
							return types.Int64Value(int64(*response.OspfSettings.Cost))
						}
						return types.Int64{}
					}(),
					IsPassiveEnabled: func() types.Bool {
						if response.OspfSettings.IsPassiveEnabled != nil {
							return types.BoolValue(*response.OspfSettings.IsPassiveEnabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfSettingsRs{}
		}(),
		OspfV3: func() *ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfV3Rs {
			if response.OspfV3 != nil {
				return &ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfV3Rs{
					Area: types.StringValue(response.OspfV3.Area),
					Cost: func() types.Int64 {
						if response.OspfV3.Cost != nil {
							return types.Int64Value(int64(*response.OspfV3.Cost))
						}
						return types.Int64{}
					}(),
					IsPassiveEnabled: func() types.Bool {
						if response.OspfV3.IsPassiveEnabled != nil {
							return types.BoolValue(*response.OspfV3.IsPassiveEnabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseSwitchGetDeviceSwitchRoutingInterfaceOspfV3Rs{}
		}(),
		Subnet: types.StringValue(response.Subnet),
		VLANID: func() types.Int64 {
			if response.VLANID != nil {
				return types.Int64Value(int64(*response.VLANID))
			}
			return types.Int64{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesSwitchRoutingInterfacesRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesSwitchRoutingInterfacesRs)
}
