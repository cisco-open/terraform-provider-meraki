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
	_ resource.Resource              = &NetworksSwitchStacksRoutingInterfacesResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchStacksRoutingInterfacesResource{}
)

func NewNetworksSwitchStacksRoutingInterfacesResource() resource.Resource {
	return &NetworksSwitchStacksRoutingInterfacesResource{}
}

type NetworksSwitchStacksRoutingInterfacesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchStacksRoutingInterfacesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchStacksRoutingInterfacesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stacks_routing_interfaces"
}

func (r *NetworksSwitchStacksRoutingInterfacesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"default_gateway": schema.StringAttribute{
				MarkdownDescription: `The next hop for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: `interfaceId path parameter. Interface ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"interface_ip": schema.StringAttribute{
				MarkdownDescription: `The IP address this switch stack will use for layer 3 routing on this VLAN or subnet. This cannot be the same as the switch's management IP.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ipv6": schema.SingleNestedAttribute{
				MarkdownDescription: `The IPv6 settings of the interface.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"address": schema.StringAttribute{
						MarkdownDescription: `The IPv6 address of the interface. Required if assignmentMode is 'static'. Must not be included if assignmentMode is 'eui-64'.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"assignment_mode": schema.StringAttribute{
						MarkdownDescription: `The IPv6 assignment mode for the interface. Can be either 'eui-64' or 'static'.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"gateway": schema.StringAttribute{
						MarkdownDescription: `The IPv6 default gateway of the interface. Required if prefix is defined and this is the first interface with IPv6 configured for the stack.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"prefix": schema.StringAttribute{
						MarkdownDescription: `The IPv6 prefix of the interface. Required if IPv6 object is included.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"multicast_routing": schema.StringAttribute{
				MarkdownDescription: `Enable multicast support if, multicast routing between VLANs is required. Options are, 'disabled', 'enabled' or 'IGMP snooping querier'. Default is 'disabled'.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `A friendly name or description for the interface or VLAN.`,
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
			"ospf_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `The OSPF routing settings of the interface.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"area": schema.StringAttribute{
						MarkdownDescription: `The OSPF area to which this interface should belong. Can be either 'disabled' or the identifier of an existing OSPF area. Defaults to 'disabled'.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"cost": schema.Int64Attribute{
						MarkdownDescription: `The path cost for this interface. Defaults to 1, but can be increased up to 65535 to give lower priority.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"is_passive_enabled": schema.BoolAttribute{
						MarkdownDescription: `When enabled, OSPF will not run on the interface, but the subnet will still be advertised.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"ospf_v3": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"area": schema.StringAttribute{
						Computed: true,
					},
					"cost": schema.Int64Attribute{
						Computed: true,
					},
					"is_passive_enabled": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
			"subnet": schema.StringAttribute{
				MarkdownDescription: `The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: `switchStackId path parameter. Switch stack ID`,
				Required:            true,
			},
			"vlan_id": schema.Int64Attribute{
				MarkdownDescription: `The VLAN this routed interface is on. VLAN must be between 1 and 4094.`,
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

func (r *NetworksSwitchStacksRoutingInterfacesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchStacksRoutingInterfacesRs

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
	vvSwitchStackID := data.SwitchStackID.ValueString()
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchStackRoutingInterfaces(vvNetworkID, vvSwitchStackID)
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStackRoutingInterfaces",
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
			r.client.Switch.UpdateNetworkSwitchStackRoutingInterface(vvNetworkID, vvSwitchStackID, vvInterfaceID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Switch.GetNetworkSwitchStackRoutingInterface(vvNetworkID, vvSwitchStackID, vvInterfaceID)
			if responseVerifyItem2 != nil {
				data = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp2, err := r.client.Switch.CreateNetworkSwitchStackRoutingInterface(vvNetworkID, vvSwitchStackID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkSwitchStackRoutingInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkSwitchStackRoutingInterface",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchStackRoutingInterfaces(vvNetworkID, vvSwitchStackID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStackRoutingInterfaces",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchStackRoutingInterfaces",
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
		responseVerifyItem2, restyRespGet, err := r.client.Switch.GetNetworkSwitchStackRoutingInterface(vvNetworkID, vvSwitchStackID, vvInterfaceID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkSwitchStackRoutingInterface",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStackRoutingInterface",
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

func (r *NetworksSwitchStacksRoutingInterfacesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchStacksRoutingInterfacesRs

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
	vvSwitchStackID := data.SwitchStackID.ValueString()
	// switch_stack_id
	vvInterfaceID := data.InterfaceID.ValueString()
	// interface_id
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchStackRoutingInterface(vvNetworkID, vvSwitchStackID, vvInterfaceID)
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
				"Failure when executing GetNetworkSwitchStackRoutingInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchStackRoutingInterface",
			err.Error(),
		)
		return
	}

	data = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchStacksRoutingInterfacesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 3 || idParts[0] == "" || idParts[1] == "" || idParts[2] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("switch_stack_id"), idParts[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("interface_id"), idParts[2])...)
}

func (r *NetworksSwitchStacksRoutingInterfacesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchStacksRoutingInterfacesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvSwitchStackID := data.SwitchStackID.ValueString()
	vvInterfaceID := data.InterfaceID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Switch.UpdateNetworkSwitchStackRoutingInterface(vvNetworkID, vvSwitchStackID, vvInterfaceID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchStackRoutingInterface",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchStackRoutingInterface",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchStacksRoutingInterfacesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksSwitchStacksRoutingInterfacesRs
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
	vvSwitchStackID := state.SwitchStackID.ValueString()
	vvInterfaceID := state.InterfaceID.ValueString()
	_, err := r.client.Switch.DeleteNetworkSwitchStackRoutingInterface(vvNetworkID, vvSwitchStackID, vvInterfaceID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkSwitchStackRoutingInterface", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksSwitchStacksRoutingInterfacesRs struct {
	NetworkID        types.String                                                       `tfsdk:"network_id"`
	SwitchStackID    types.String                                                       `tfsdk:"switch_stack_id"`
	InterfaceID      types.String                                                       `tfsdk:"interface_id"`
	DefaultGateway   types.String                                                       `tfsdk:"default_gateway"`
	InterfaceIP      types.String                                                       `tfsdk:"interface_ip"`
	IPv6             *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIpv6Rs         `tfsdk:"ipv6"`
	MulticastRouting types.String                                                       `tfsdk:"multicast_routing"`
	Name             types.String                                                       `tfsdk:"name"`
	OspfSettings     *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettingsRs `tfsdk:"ospf_settings"`
	OspfV3           *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3Rs       `tfsdk:"ospf_v3"`
	Subnet           types.String                                                       `tfsdk:"subnet"`
	VLANID           types.Int64                                                        `tfsdk:"vlan_id"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIpv6Rs struct {
	Address        types.String `tfsdk:"address"`
	AssignmentMode types.String `tfsdk:"assignment_mode"`
	Gateway        types.String `tfsdk:"gateway"`
	Prefix         types.String `tfsdk:"prefix"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettingsRs struct {
	Area             types.String `tfsdk:"area"`
	Cost             types.Int64  `tfsdk:"cost"`
	IsPassiveEnabled types.Bool   `tfsdk:"is_passive_enabled"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3Rs struct {
	Area             types.String `tfsdk:"area"`
	Cost             types.Int64  `tfsdk:"cost"`
	IsPassiveEnabled types.Bool   `tfsdk:"is_passive_enabled"`
}

// FromBody
func (r *NetworksSwitchStacksRoutingInterfacesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCreateNetworkSwitchStackRoutingInterface {
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
	var requestSwitchCreateNetworkSwitchStackRoutingInterfaceIPv6 *merakigosdk.RequestSwitchCreateNetworkSwitchStackRoutingInterfaceIPv6
	if r.IPv6 != nil {
		address := r.IPv6.Address.ValueString()
		assignmentMode := r.IPv6.AssignmentMode.ValueString()
		gateway := r.IPv6.Gateway.ValueString()
		prefix := r.IPv6.Prefix.ValueString()
		requestSwitchCreateNetworkSwitchStackRoutingInterfaceIPv6 = &merakigosdk.RequestSwitchCreateNetworkSwitchStackRoutingInterfaceIPv6{
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
	var requestSwitchCreateNetworkSwitchStackRoutingInterfaceOspfSettings *merakigosdk.RequestSwitchCreateNetworkSwitchStackRoutingInterfaceOspfSettings
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
		requestSwitchCreateNetworkSwitchStackRoutingInterfaceOspfSettings = &merakigosdk.RequestSwitchCreateNetworkSwitchStackRoutingInterfaceOspfSettings{
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
	out := merakigosdk.RequestSwitchCreateNetworkSwitchStackRoutingInterface{
		DefaultGateway:   *defaultGateway,
		InterfaceIP:      *interfaceIP,
		IPv6:             requestSwitchCreateNetworkSwitchStackRoutingInterfaceIPv6,
		MulticastRouting: *multicastRouting,
		Name:             *name,
		OspfSettings:     requestSwitchCreateNetworkSwitchStackRoutingInterfaceOspfSettings,
		Subnet:           *subnet,
		VLANID:           int64ToIntPointer(vLANID),
	}
	return &out
}
func (r *NetworksSwitchStacksRoutingInterfacesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterface {
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
	var requestSwitchUpdateNetworkSwitchStackRoutingInterfaceIPv6 *merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceIPv6
	if r.IPv6 != nil {
		address := r.IPv6.Address.ValueString()
		assignmentMode := r.IPv6.AssignmentMode.ValueString()
		gateway := r.IPv6.Gateway.ValueString()
		prefix := r.IPv6.Prefix.ValueString()
		requestSwitchUpdateNetworkSwitchStackRoutingInterfaceIPv6 = &merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceIPv6{
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
	var requestSwitchUpdateNetworkSwitchStackRoutingInterfaceOspfSettings *merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceOspfSettings
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
		requestSwitchUpdateNetworkSwitchStackRoutingInterfaceOspfSettings = &merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterfaceOspfSettings{
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
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchStackRoutingInterface{
		DefaultGateway:   *defaultGateway,
		InterfaceIP:      *interfaceIP,
		IPv6:             requestSwitchUpdateNetworkSwitchStackRoutingInterfaceIPv6,
		MulticastRouting: *multicastRouting,
		Name:             *name,
		OspfSettings:     requestSwitchUpdateNetworkSwitchStackRoutingInterfaceOspfSettings,
		Subnet:           *subnet,
		VLANID:           int64ToIntPointer(vLANID),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchStackRoutingInterfaceItemToBodyRs(state NetworksSwitchStacksRoutingInterfacesRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchStackRoutingInterface, is_read bool) NetworksSwitchStacksRoutingInterfacesRs {
	itemState := NetworksSwitchStacksRoutingInterfacesRs{
		DefaultGateway: types.StringValue(response.DefaultGateway),
		InterfaceID:    types.StringValue(response.InterfaceID),
		InterfaceIP:    types.StringValue(response.InterfaceIP),
		IPv6: func() *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIpv6Rs {
			if response.IPv6 != nil {
				return &ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIpv6Rs{
					Address:        types.StringValue(response.IPv6.Address),
					AssignmentMode: types.StringValue(response.IPv6.AssignmentMode),
					Gateway:        types.StringValue(response.IPv6.Gateway),
					Prefix:         types.StringValue(response.IPv6.Prefix),
				}
			}
			return &ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIpv6Rs{}
		}(),
		MulticastRouting: types.StringValue(response.MulticastRouting),
		Name:             types.StringValue(response.Name),
		OspfSettings: func() *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettingsRs {
			if response.OspfSettings != nil {
				return &ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettingsRs{
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
			return &ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettingsRs{}
		}(),
		OspfV3: func() *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3Rs {
			if response.OspfV3 != nil {
				return &ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3Rs{
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
			return &ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3Rs{}
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
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchStacksRoutingInterfacesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchStacksRoutingInterfacesRs)
}
