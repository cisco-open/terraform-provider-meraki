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

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DevicesManagementInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesManagementInterfaceDataSource{}
)

func NewDevicesManagementInterfaceDataSource() datasource.DataSource {
	return &DevicesManagementInterfaceDataSource{}
}

type DevicesManagementInterfaceDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesManagementInterfaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesManagementInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_management_interface"
}

func (d *DevicesManagementInterfaceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
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
					"wan1": schema.SingleNestedAttribute{
						MarkdownDescription: `WAN 1 settings`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"static_dns": schema.ListAttribute{
								MarkdownDescription: `Up to two DNS IPs.`,
								Computed:            true,
								ElementType:         types.StringType,
							},
							"static_gateway_ip": schema.StringAttribute{
								MarkdownDescription: `The IP of the gateway on the WAN.`,
								Computed:            true,
							},
							"static_ip": schema.StringAttribute{
								MarkdownDescription: `The IP the device should use on the WAN.`,
								Computed:            true,
							},
							"static_subnet_mask": schema.StringAttribute{
								MarkdownDescription: `The subnet mask for the WAN.`,
								Computed:            true,
							},
							"using_static_ip": schema.BoolAttribute{
								MarkdownDescription: `Configure the interface to have static IP settings or use DHCP.`,
								Computed:            true,
							},
							"vlan": schema.Int64Attribute{
								MarkdownDescription: `The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.`,
								Computed:            true,
							},
							"wan_enabled": schema.StringAttribute{
								MarkdownDescription: `Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.`,
								Computed:            true,
							},
						},
					},
					"wan2": schema.SingleNestedAttribute{
						MarkdownDescription: `WAN 2 settings (only for MX devices)`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"static_dns": schema.ListAttribute{
								MarkdownDescription: `Up to two DNS IPs.`,
								Computed:            true,
								ElementType:         types.StringType,
							},
							"static_gateway_ip": schema.StringAttribute{
								MarkdownDescription: `The IP of the gateway on the WAN.`,
								Computed:            true,
							},
							"static_ip": schema.StringAttribute{
								MarkdownDescription: `The IP the device should use on the WAN.`,
								Computed:            true,
							},
							"static_subnet_mask": schema.StringAttribute{
								MarkdownDescription: `The subnet mask for the WAN.`,
								Computed:            true,
							},
							"using_static_ip": schema.BoolAttribute{
								MarkdownDescription: `Configure the interface to have static IP settings or use DHCP.`,
								Computed:            true,
							},
							"vlan": schema.Int64Attribute{
								MarkdownDescription: `The VLAN that management traffic should be tagged with. Applies whether usingStaticIp is true or false.`,
								Computed:            true,
							},
							"wan_enabled": schema.StringAttribute{
								MarkdownDescription: `Enable or disable the interface (only for MX devices). Valid values are 'enabled', 'disabled', and 'not configured'.`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *DevicesManagementInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesManagementInterface DevicesManagementInterface
	diags := req.Config.Get(ctx, &devicesManagementInterface)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceManagementInterface")
		vvSerial := devicesManagementInterface.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Devices.GetDeviceManagementInterface(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceManagementInterface",
				err.Error(),
			)
			return
		}

		devicesManagementInterface = ResponseDevicesGetDeviceManagementInterfaceItemToBody(devicesManagementInterface, response1)
		diags = resp.State.Set(ctx, &devicesManagementInterface)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesManagementInterface struct {
	Serial types.String                                 `tfsdk:"serial"`
	Item   *ResponseDevicesGetDeviceManagementInterface `tfsdk:"item"`
}

type ResponseDevicesGetDeviceManagementInterface struct {
	DdnsHostnames *ResponseDevicesGetDeviceManagementInterfaceDdnsHostnames `tfsdk:"ddns_hostnames"`
	Wan1          *ResponseDevicesGetDeviceManagementInterfaceWan1          `tfsdk:"wan1"`
	Wan2          *ResponseDevicesGetDeviceManagementInterfaceWan2          `tfsdk:"wan2"`
}

type ResponseDevicesGetDeviceManagementInterfaceDdnsHostnames struct {
	ActiveDdnsHostname types.String `tfsdk:"active_ddns_hostname"`
	DdnsHostnameWan1   types.String `tfsdk:"ddns_hostname_wan1"`
	DdnsHostnameWan2   types.String `tfsdk:"ddns_hostname_wan2"`
}

type ResponseDevicesGetDeviceManagementInterfaceWan1 struct {
	StaticDNS        types.List   `tfsdk:"static_dns"`
	StaticGatewayIP  types.String `tfsdk:"static_gateway_ip"`
	StaticIP         types.String `tfsdk:"static_ip"`
	StaticSubnetMask types.String `tfsdk:"static_subnet_mask"`
	UsingStaticIP    types.Bool   `tfsdk:"using_static_ip"`
	VLAN             types.Int64  `tfsdk:"vlan"`
	WanEnabled       types.String `tfsdk:"wan_enabled"`
}

type ResponseDevicesGetDeviceManagementInterfaceWan2 struct {
	StaticDNS        types.List   `tfsdk:"static_dns"`
	StaticGatewayIP  types.String `tfsdk:"static_gateway_ip"`
	StaticIP         types.String `tfsdk:"static_ip"`
	StaticSubnetMask types.String `tfsdk:"static_subnet_mask"`
	UsingStaticIP    types.Bool   `tfsdk:"using_static_ip"`
	VLAN             types.Int64  `tfsdk:"vlan"`
	WanEnabled       types.String `tfsdk:"wan_enabled"`
}

// ToBody
func ResponseDevicesGetDeviceManagementInterfaceItemToBody(state DevicesManagementInterface, response *merakigosdk.ResponseDevicesGetDeviceManagementInterface) DevicesManagementInterface {
	itemState := ResponseDevicesGetDeviceManagementInterface{
		DdnsHostnames: func() *ResponseDevicesGetDeviceManagementInterfaceDdnsHostnames {
			if response.DdnsHostnames != nil {
				return &ResponseDevicesGetDeviceManagementInterfaceDdnsHostnames{
					ActiveDdnsHostname: types.StringValue(response.DdnsHostnames.ActiveDdnsHostname),
					DdnsHostnameWan1:   types.StringValue(response.DdnsHostnames.DdnsHostnameWan1),
					DdnsHostnameWan2:   types.StringValue(response.DdnsHostnames.DdnsHostnameWan2),
				}
			}
			return nil
		}(),
		Wan1: func() *ResponseDevicesGetDeviceManagementInterfaceWan1 {
			if response.Wan1 != nil {
				return &ResponseDevicesGetDeviceManagementInterfaceWan1{
					StaticDNS:        StringSliceToList(response.Wan1.StaticDNS),
					StaticGatewayIP:  types.StringValue(response.Wan1.StaticGatewayIP),
					StaticIP:         types.StringValue(response.Wan1.StaticIP),
					StaticSubnetMask: types.StringValue(response.Wan1.StaticSubnetMask),
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
					WanEnabled: types.StringValue(response.Wan1.WanEnabled),
				}
			}
			return nil
		}(),
		Wan2: func() *ResponseDevicesGetDeviceManagementInterfaceWan2 {
			if response.Wan2 != nil {
				return &ResponseDevicesGetDeviceManagementInterfaceWan2{
					StaticDNS:        StringSliceToList(response.Wan2.StaticDNS),
					StaticGatewayIP:  types.StringValue(response.Wan2.StaticGatewayIP),
					StaticIP:         types.StringValue(response.Wan2.StaticIP),
					StaticSubnetMask: types.StringValue(response.Wan2.StaticSubnetMask),
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
					WanEnabled: types.StringValue(response.Wan2.WanEnabled),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
