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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesManagementInterfaceResource{}
	_ resource.ResourceWithConfigure = &DevicesManagementInterfaceResource{}
)

func NewDevicesManagementInterfaceResource() resource.Resource {
	return &DevicesManagementInterfaceResource{}
}

type DevicesManagementInterfaceResource struct {
	client *merakigosdk.Client
}

func (r *DevicesManagementInterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesManagementInterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_management_interface"
}

func (r *DevicesManagementInterfaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ddns_hostnames": schema.SingleNestedAttribute{
				MarkdownDescription: `Dynamic DNS hostnames.`,
				Computed:            true,
				Attributes: map[string]schema.Attribute{

					"active_ddns_hostname": schema.StringAttribute{
						MarkdownDescription: `Active dynamic DNS hostname.`,
						Computed:            true,
					},
					"ddns_hostname_wan1": schema.StringAttribute{
						MarkdownDescription: `WAN 1 dynamic DNS hostname.`,
						Computed:            true,
					},
					"ddns_hostname_wan2": schema.StringAttribute{
						MarkdownDescription: `WAN 2 dynamic DNS hostname.`,
						Computed:            true,
					},
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"wan1": schema.SingleNestedAttribute{
				MarkdownDescription: `WAN 1 settings`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"static_dns": schema.SetAttribute{
						MarkdownDescription: `Up to two DNS IPs.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
						Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
					},
					"static_gateway_ip": schema.StringAttribute{
						MarkdownDescription: `The IP of the gateway on the WAN.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"static_ip": schema.StringAttribute{
						MarkdownDescription: `The IP the device should use on the WAN.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"static_subnet_mask": schema.StringAttribute{
						MarkdownDescription: `The subnet mask for the WAN.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"using_static_ip": schema.BoolAttribute{
						MarkdownDescription: `Configure the interface to have static IP settings or use DHCP.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"vlan": schema.Int64Attribute{
						MarkdownDescription: `The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"wan_enabled": schema.StringAttribute{
						MarkdownDescription: `Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
                                        Allowed values: [disabled,enabled,not configured]`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"disabled",
								"enabled",
								"not configured",
							),
						},
					},
				},
			},
			"wan2": schema.SingleNestedAttribute{
				MarkdownDescription: `WAN 2 settings (only for MX devices)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"static_dns": schema.SetAttribute{
						MarkdownDescription: `Up to two DNS IPs.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
						ElementType: types.StringType,
					},
					"static_gateway_ip": schema.StringAttribute{
						MarkdownDescription: `The IP of the gateway on the WAN.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"static_ip": schema.StringAttribute{
						MarkdownDescription: `The IP the device should use on the WAN.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"static_subnet_mask": schema.StringAttribute{
						MarkdownDescription: `The subnet mask for the WAN.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"using_static_ip": schema.BoolAttribute{
						MarkdownDescription: `Configure the interface to have static IP settings or use DHCP.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"vlan": schema.Int64Attribute{
						MarkdownDescription: `The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"wan_enabled": schema.StringAttribute{
						MarkdownDescription: `Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.
                                        Allowed values: [disabled,enabled,not configured]`,
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"disabled",
								"enabled",
								"not configured",
							),
						},
					},
				},
			},
		},
	}
}

func (r *DevicesManagementInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesManagementInterfaceRs

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
	vvSerial := data.Serial.ValueString()
	//Has Item and not has items

	if vvSerial != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Devices.GetDeviceManagementInterface(vvSerial)
		//Has Post
		if err != nil {
			if restyResp1 != nil {
				if restyResp1.StatusCode() != 404 {
					resp.Diagnostics.AddError(
						"Failure when executing GetDeviceManagementInterface",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data = ResponseDevicesGetDeviceManagementInterfaceItemToBodyRs(data, responseVerifyItem, false)
			//Path params in update assigned
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
	}

	response, restyResp2, err := r.client.Devices.RebootDevice(vvSerial)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing RebootDevice",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing RebootDevice",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Devices.GetDeviceManagementInterface(vvSerial)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceManagementInterface",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceManagementInterface",
			err.Error(),
		)
		return
	}

	data = ResponseDevicesGetDeviceManagementInterfaceItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *DevicesManagementInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesManagementInterfaceRs

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
	responseGet, restyRespGet, err := r.client.Devices.GetDeviceManagementInterface(vvSerial)
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
				"Failure when executing GetDeviceManagementInterface",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceManagementInterface",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseDevicesGetDeviceManagementInterfaceItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesManagementInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesManagementInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesManagementInterfaceRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Devices.UpdateDeviceManagementInterface(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceManagementInterface",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceManagementInterface",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesManagementInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesManagementInterface", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesManagementInterfaceRs struct {
	Serial        types.String                                                `tfsdk:"serial"`
	DdnsHostnames *ResponseDevicesGetDeviceManagementInterfaceDdnsHostnamesRs `tfsdk:"ddns_hostnames"`
	Wan1          *ResponseDevicesGetDeviceManagementInterfaceWan1Rs          `tfsdk:"wan1"`
	Wan2          *ResponseDevicesGetDeviceManagementInterfaceWan2Rs          `tfsdk:"wan2"`
}

type ResponseDevicesGetDeviceManagementInterfaceDdnsHostnamesRs struct {
	ActiveDdnsHostname types.String `tfsdk:"active_ddns_hostname"`
	DdnsHostnameWan1   types.String `tfsdk:"ddns_hostname_wan1"`
	DdnsHostnameWan2   types.String `tfsdk:"ddns_hostname_wan2"`
}

type ResponseDevicesGetDeviceManagementInterfaceWan1Rs struct {
	StaticDNS        types.Set    `tfsdk:"static_dns"`
	StaticGatewayIP  types.String `tfsdk:"static_gateway_ip"`
	StaticIP         types.String `tfsdk:"static_ip"`
	StaticSubnetMask types.String `tfsdk:"static_subnet_mask"`
	UsingStaticIP    types.Bool   `tfsdk:"using_static_ip"`
	VLAN             types.Int64  `tfsdk:"vlan"`
	WanEnabled       types.String `tfsdk:"wan_enabled"`
}

type ResponseDevicesGetDeviceManagementInterfaceWan2Rs struct {
	StaticDNS        types.Set    `tfsdk:"static_dns"`
	StaticGatewayIP  types.String `tfsdk:"static_gateway_ip"`
	StaticIP         types.String `tfsdk:"static_ip"`
	StaticSubnetMask types.String `tfsdk:"static_subnet_mask"`
	UsingStaticIP    types.Bool   `tfsdk:"using_static_ip"`
	VLAN             types.Int64  `tfsdk:"vlan"`
	WanEnabled       types.String `tfsdk:"wan_enabled"`
}

// FromBody
func (r *DevicesManagementInterfaceRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestDevicesUpdateDeviceManagementInterface {
	var requestDevicesUpdateDeviceManagementInterfaceWan1 *merakigosdk.RequestDevicesUpdateDeviceManagementInterfaceWan1

	if r.Wan1 != nil {

		var staticDNS []string = nil
		r.Wan1.StaticDNS.ElementsAs(ctx, &staticDNS, false)
		staticGatewayIP := r.Wan1.StaticGatewayIP.ValueString()
		staticIP := r.Wan1.StaticIP.ValueString()
		staticSubnetMask := r.Wan1.StaticSubnetMask.ValueString()
		usingStaticIP := func() *bool {
			if !r.Wan1.UsingStaticIP.IsUnknown() && !r.Wan1.UsingStaticIP.IsNull() {
				return r.Wan1.UsingStaticIP.ValueBoolPointer()
			}
			return nil
		}()
		vlan := func() *int64 {
			if !r.Wan1.VLAN.IsUnknown() && !r.Wan1.VLAN.IsNull() {
				return r.Wan1.VLAN.ValueInt64Pointer()
			}
			return nil
		}()
		wanEnabled := r.Wan1.WanEnabled.ValueString()
		requestDevicesUpdateDeviceManagementInterfaceWan1 = &merakigosdk.RequestDevicesUpdateDeviceManagementInterfaceWan1{
			StaticDNS:        staticDNS,
			StaticGatewayIP:  staticGatewayIP,
			StaticIP:         staticIP,
			StaticSubnetMask: staticSubnetMask,
			UsingStaticIP:    usingStaticIP,
			VLAN:             int64ToIntPointer(vlan),
			WanEnabled:       wanEnabled,
		}
		//[debug] Is Array: False
	}
	var requestDevicesUpdateDeviceManagementInterfaceWan2 *merakigosdk.RequestDevicesUpdateDeviceManagementInterfaceWan2

	if r.Wan2 != nil {

		var staticDNS []string = nil
		r.Wan2.StaticDNS.ElementsAs(ctx, &staticDNS, false)
		staticGatewayIP := r.Wan2.StaticGatewayIP.ValueString()
		staticIP := r.Wan2.StaticIP.ValueString()
		staticSubnetMask := r.Wan2.StaticSubnetMask.ValueString()
		usingStaticIP := func() *bool {
			if !r.Wan2.UsingStaticIP.IsUnknown() && !r.Wan2.UsingStaticIP.IsNull() {
				return r.Wan2.UsingStaticIP.ValueBoolPointer()
			}
			return nil
		}()
		vlan := func() *int64 {
			if !r.Wan2.VLAN.IsUnknown() && !r.Wan2.VLAN.IsNull() {
				return r.Wan2.VLAN.ValueInt64Pointer()
			}
			return nil
		}()
		wanEnabled := r.Wan2.WanEnabled.ValueString()
		requestDevicesUpdateDeviceManagementInterfaceWan2 = &merakigosdk.RequestDevicesUpdateDeviceManagementInterfaceWan2{
			StaticDNS:        staticDNS,
			StaticGatewayIP:  staticGatewayIP,
			StaticIP:         staticIP,
			StaticSubnetMask: staticSubnetMask,
			UsingStaticIP:    usingStaticIP,
			VLAN:             int64ToIntPointer(vlan),
			WanEnabled:       wanEnabled,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestDevicesUpdateDeviceManagementInterface{
		Wan1: requestDevicesUpdateDeviceManagementInterfaceWan1,
		Wan2: requestDevicesUpdateDeviceManagementInterfaceWan2,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseDevicesGetDeviceManagementInterfaceItemToBodyRs(state DevicesManagementInterfaceRs, response *merakigosdk.ResponseDevicesGetDeviceManagementInterface, is_read bool) DevicesManagementInterfaceRs {
	itemState := DevicesManagementInterfaceRs{
		DdnsHostnames: func() *ResponseDevicesGetDeviceManagementInterfaceDdnsHostnamesRs {
			if response.DdnsHostnames != nil {
				return &ResponseDevicesGetDeviceManagementInterfaceDdnsHostnamesRs{
					ActiveDdnsHostname: func() types.String {
						if response.DdnsHostnames.ActiveDdnsHostname != "" {
							return types.StringValue(response.DdnsHostnames.ActiveDdnsHostname)
						}
						return types.String{}
					}(),
					DdnsHostnameWan1: func() types.String {
						if response.DdnsHostnames.DdnsHostnameWan1 != "" {
							return types.StringValue(response.DdnsHostnames.DdnsHostnameWan1)
						}
						return types.String{}
					}(),
					DdnsHostnameWan2: func() types.String {
						if response.DdnsHostnames.DdnsHostnameWan2 != "" {
							return types.StringValue(response.DdnsHostnames.DdnsHostnameWan2)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		Wan1: func() *ResponseDevicesGetDeviceManagementInterfaceWan1Rs {
			if response.Wan1 != nil {
				return &ResponseDevicesGetDeviceManagementInterfaceWan1Rs{
					StaticDNS: StringSliceToSet(response.Wan1.StaticDNS),
					StaticGatewayIP: func() types.String {
						if response.Wan1.StaticGatewayIP != "" {
							return types.StringValue(response.Wan1.StaticGatewayIP)
						}
						return types.String{}
					}(),
					StaticIP: func() types.String {
						if response.Wan1.StaticIP != "" {
							return types.StringValue(response.Wan1.StaticIP)
						}
						return types.String{}
					}(),
					StaticSubnetMask: func() types.String {
						if response.Wan1.StaticSubnetMask != "" {
							return types.StringValue(response.Wan1.StaticSubnetMask)
						}
						return types.String{}
					}(),
					UsingStaticIP: func() types.Bool {
						if response.Wan1.UsingStaticIP != nil {
							return types.BoolValue(*response.Wan1.UsingStaticIP)
						}
						return types.Bool{}
					}(),
					VLAN: func() types.Int64 {
						if response.Wan1.VLAN != nil {
							return types.Int64Value(int64(*response.Wan1.VLAN))
						}
						return types.Int64{}
					}(),
					WanEnabled: func() types.String {
						if response.Wan1.WanEnabled != "" {
							return types.StringValue(response.Wan1.WanEnabled)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		Wan2: func() *ResponseDevicesGetDeviceManagementInterfaceWan2Rs {
			if response.Wan2 != nil {
				return &ResponseDevicesGetDeviceManagementInterfaceWan2Rs{
					StaticDNS: StringSliceToSet(response.Wan2.StaticDNS),
					StaticGatewayIP: func() types.String {
						if response.Wan2.StaticGatewayIP != "" {
							return types.StringValue(response.Wan2.StaticGatewayIP)
						}
						return types.String{}
					}(),
					StaticIP: func() types.String {
						if response.Wan2.StaticIP != "" {
							return types.StringValue(response.Wan2.StaticIP)
						}
						return types.String{}
					}(),
					StaticSubnetMask: func() types.String {
						if response.Wan2.StaticSubnetMask != "" {
							return types.StringValue(response.Wan2.StaticSubnetMask)
						}
						return types.String{}
					}(),
					UsingStaticIP: func() types.Bool {
						if response.Wan2.UsingStaticIP != nil {
							return types.BoolValue(*response.Wan2.UsingStaticIP)
						}
						return types.Bool{}
					}(),
					VLAN: func() types.Int64 {
						if response.Wan2.VLAN != nil {
							return types.Int64Value(int64(*response.Wan2.VLAN))
						}
						return types.Int64{}
					}(),
					WanEnabled: func() types.String {
						if response.Wan2.WanEnabled != "" {
							return types.StringValue(response.Wan2.WanEnabled)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesManagementInterfaceRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesManagementInterfaceRs)
}
