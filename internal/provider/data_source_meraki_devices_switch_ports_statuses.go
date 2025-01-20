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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DevicesSwitchPortsStatusesDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesSwitchPortsStatusesDataSource{}
)

func NewDevicesSwitchPortsStatusesDataSource() datasource.DataSource {
	return &DevicesSwitchPortsStatusesDataSource{}
}

type DevicesSwitchPortsStatusesDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesSwitchPortsStatusesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesSwitchPortsStatusesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_switch_ports_statuses"
}

func (d *DevicesSwitchPortsStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetDeviceSwitchPortsStatuses`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"cdp": schema.SingleNestedAttribute{
							MarkdownDescription: `The Cisco Discovery Protocol (CDP) information of the connected device.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"address": schema.StringAttribute{
									MarkdownDescription: `Contains network addresses of both receiving and sending devices.`,
									Computed:            true,
								},
								"capabilities": schema.StringAttribute{
									MarkdownDescription: `Identifies the device type, which indicates the functional capabilities of the device.`,
									Computed:            true,
								},
								"device_id": schema.StringAttribute{
									MarkdownDescription: `Identifies the device name.`,
									Computed:            true,
								},
								"management_address": schema.StringAttribute{
									MarkdownDescription: `The device's management IP.`,
									Computed:            true,
								},
								"native_vlan": schema.Int64Attribute{
									MarkdownDescription: `Indicates, per interface, the assumed VLAN for untagged packets on the interface.`,
									Computed:            true,
								},
								"platform": schema.StringAttribute{
									MarkdownDescription: `Identifies the hardware platform of the device.`,
									Computed:            true,
								},
								"port_id": schema.StringAttribute{
									MarkdownDescription: `Identifies the port from which the CDP packet was sent.`,
									Computed:            true,
								},
								"system_name": schema.StringAttribute{
									MarkdownDescription: `The system name.`,
									Computed:            true,
								},
								"version": schema.StringAttribute{
									MarkdownDescription: `Contains the device software release information.`,
									Computed:            true,
								},
								"vtp_management_domain": schema.StringAttribute{
									MarkdownDescription: `Advertises the configured VLAN Trunking Protocl (VTP)-management-domain name of the system.`,
									Computed:            true,
								},
							},
						},
						"client_count": schema.Int64Attribute{
							MarkdownDescription: `The number of clients connected through this port.`,
							Computed:            true,
						},
						"duplex": schema.StringAttribute{
							MarkdownDescription: `The current duplex of a connected port.`,
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: `Whether the port is configured to be enabled.`,
							Computed:            true,
						},
						"errors": schema.ListAttribute{
							MarkdownDescription: `All errors present on the port.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"is_uplink": schema.BoolAttribute{
							MarkdownDescription: `Whether the port is the switch's uplink.`,
							Computed:            true,
						},
						"lldp": schema.SingleNestedAttribute{
							MarkdownDescription: `The Link Layer Discovery Protocol (LLDP) information of the connected device.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"chassis_id": schema.StringAttribute{
									MarkdownDescription: `The device's chassis ID.`,
									Computed:            true,
								},
								"management_address": schema.StringAttribute{
									MarkdownDescription: `The device's management IP.`,
									Computed:            true,
								},
								"management_vlan": schema.Int64Attribute{
									MarkdownDescription: `The device's management VLAN.`,
									Computed:            true,
								},
								"port_description": schema.StringAttribute{
									MarkdownDescription: `Description of the port from which the LLDP packet was sent.`,
									Computed:            true,
								},
								"port_id": schema.StringAttribute{
									MarkdownDescription: `Identifies the port from which the LLDP packet was sent`,
									Computed:            true,
								},
								"port_vlan": schema.Int64Attribute{
									MarkdownDescription: `The port's VLAN.`,
									Computed:            true,
								},
								"system_capabilities": schema.StringAttribute{
									MarkdownDescription: `Identifies the device type, which indicates the functional capabilities of the device.`,
									Computed:            true,
								},
								"system_description": schema.StringAttribute{
									MarkdownDescription: `The device's system description.`,
									Computed:            true,
								},
								"system_name": schema.StringAttribute{
									MarkdownDescription: `The device's system name.`,
									Computed:            true,
								},
							},
						},
						"poe": schema.SingleNestedAttribute{
							MarkdownDescription: `PoE status of the port.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"is_allocated": schema.BoolAttribute{
									MarkdownDescription: `Whether the port is drawing power`,
									Computed:            true,
								},
							},
						},
						"port_id": schema.StringAttribute{
							MarkdownDescription: `The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.`,
							Computed:            true,
						},
						"power_usage_in_wh": schema.Float64Attribute{
							MarkdownDescription: `How much power (in watt-hours) has been delivered by this port during the timespan.`,
							Computed:            true,
						},
						"secure_port": schema.SingleNestedAttribute{
							MarkdownDescription: `The Secure Port status of the port.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"active": schema.BoolAttribute{
									MarkdownDescription: `Whether Secure Port is currently active for this port.`,
									Computed:            true,
								},
								"authentication_status": schema.StringAttribute{
									MarkdownDescription: `The current Secure Port status.`,
									Computed:            true,
								},
								"config_overrides": schema.SingleNestedAttribute{
									MarkdownDescription: `The configuration overrides applied to this port when Secure Port is active.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"allowed_vlans": schema.StringAttribute{
											MarkdownDescription: `The VLANs allowed on the . Only applicable to trunk ports.`,
											Computed:            true,
										},
										"type": schema.StringAttribute{
											MarkdownDescription: `The type of the  ('trunk', 'access' or 'stack').`,
											Computed:            true,
										},
										"vlan": schema.Int64Attribute{
											MarkdownDescription: `The VLAN of the . For a trunk port, this is the native VLAN. A null value will clear the value set for trunk ports.`,
											Computed:            true,
										},
										"voice_vlan": schema.Int64Attribute{
											MarkdownDescription: `The voice VLAN of the . Only applicable to access ports.`,
											Computed:            true,
										},
									},
								},
								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Whether Secure Port is turned on for this port.`,
									Computed:            true,
								},
							},
						},
						"spanning_tree": schema.SingleNestedAttribute{
							MarkdownDescription: `The Spanning Tree Protocol (STP) information of the connected device.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"statuses": schema.ListAttribute{
									MarkdownDescription: `The current Spanning Tree Protocol statuses of the port.`,
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
						"speed": schema.StringAttribute{
							MarkdownDescription: `The current data transfer rate which the port is operating at.`,
							Computed:            true,
						},
						"status": schema.StringAttribute{
							MarkdownDescription: `The current connection status of the port.`,
							Computed:            true,
						},
						"traffic_in_kbps": schema.SingleNestedAttribute{
							MarkdownDescription: `A breakdown of the average speed of data that has passed through this port during the timespan.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"recv": schema.Float64Attribute{
									MarkdownDescription: `The average speed of the data received (in kilobits-per-second).`,
									Computed:            true,
								},
								"sent": schema.Float64Attribute{
									MarkdownDescription: `The average speed of the data sent (in kilobits-per-second).`,
									Computed:            true,
								},
								"total": schema.Float64Attribute{
									MarkdownDescription: `The average speed of the data sent and received (in kilobits-per-second).`,
									Computed:            true,
								},
							},
						},
						"usage_in_kb": schema.SingleNestedAttribute{
							MarkdownDescription: `A breakdown of how many kilobytes have passed through this port during the timespan.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"recv": schema.Int64Attribute{
									MarkdownDescription: `The amount of data received (in kilobytes).`,
									Computed:            true,
								},
								"sent": schema.Int64Attribute{
									MarkdownDescription: `The amount of data sent (in kilobytes).`,
									Computed:            true,
								},
								"total": schema.Int64Attribute{
									MarkdownDescription: `The total amount of data sent and received (in kilobytes).`,
									Computed:            true,
								},
							},
						},
						"warnings": schema.ListAttribute{
							MarkdownDescription: `All warnings present on the port.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
					},
				},
			},
		},
	}
}

func (d *DevicesSwitchPortsStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesSwitchPortsStatuses DevicesSwitchPortsStatuses
	diags := req.Config.Get(ctx, &devicesSwitchPortsStatuses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceSwitchPortsStatuses")
		vvSerial := devicesSwitchPortsStatuses.Serial.ValueString()
		queryParams1 := merakigosdk.GetDeviceSwitchPortsStatusesQueryParams{}

		queryParams1.T0 = devicesSwitchPortsStatuses.T0.ValueString()
		queryParams1.Timespan = devicesSwitchPortsStatuses.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetDeviceSwitchPortsStatuses(vvSerial, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchPortsStatuses",
				err.Error(),
			)
			return
		}

		devicesSwitchPortsStatuses = ResponseSwitchGetDeviceSwitchPortsStatusesItemsToBody(devicesSwitchPortsStatuses, response1)
		diags = resp.State.Set(ctx, &devicesSwitchPortsStatuses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesSwitchPortsStatuses struct {
	Serial   types.String                                      `tfsdk:"serial"`
	T0       types.String                                      `tfsdk:"t0"`
	Timespan types.Float64                                     `tfsdk:"timespan"`
	Items    *[]ResponseItemSwitchGetDeviceSwitchPortsStatuses `tfsdk:"items"`
}

type ResponseItemSwitchGetDeviceSwitchPortsStatuses struct {
	Cdp            *ResponseItemSwitchGetDeviceSwitchPortsStatusesCdp           `tfsdk:"cdp"`
	ClientCount    types.Int64                                                  `tfsdk:"client_count"`
	Duplex         types.String                                                 `tfsdk:"duplex"`
	Enabled        types.Bool                                                   `tfsdk:"enabled"`
	Errors         types.List                                                   `tfsdk:"errors"`
	IsUplink       types.Bool                                                   `tfsdk:"is_uplink"`
	Lldp           *ResponseItemSwitchGetDeviceSwitchPortsStatusesLldp          `tfsdk:"lldp"`
	Poe            *ResponseItemSwitchGetDeviceSwitchPortsStatusesPoe           `tfsdk:"poe"`
	PortID         types.String                                                 `tfsdk:"port_id"`
	PowerUsageInWh types.Float64                                                `tfsdk:"power_usage_in_wh"`
	SecurePort     *ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePort    `tfsdk:"secure_port"`
	SpanningTree   *ResponseItemSwitchGetDeviceSwitchPortsStatusesSpanningTree  `tfsdk:"spanning_tree"`
	Speed          types.String                                                 `tfsdk:"speed"`
	Status         types.String                                                 `tfsdk:"status"`
	TrafficInKbps  *ResponseItemSwitchGetDeviceSwitchPortsStatusesTrafficInKbps `tfsdk:"traffic_in_kbps"`
	UsageInKb      *ResponseItemSwitchGetDeviceSwitchPortsStatusesUsageInKb     `tfsdk:"usage_in_kb"`
	Warnings       types.List                                                   `tfsdk:"warnings"`
}

type ResponseItemSwitchGetDeviceSwitchPortsStatusesCdp struct {
	Address             types.String `tfsdk:"address"`
	Capabilities        types.String `tfsdk:"capabilities"`
	DeviceID            types.String `tfsdk:"device_id"`
	ManagementAddress   types.String `tfsdk:"management_address"`
	NativeVLAN          types.Int64  `tfsdk:"native_vlan"`
	Platform            types.String `tfsdk:"platform"`
	PortID              types.String `tfsdk:"port_id"`
	SystemName          types.String `tfsdk:"system_name"`
	Version             types.String `tfsdk:"version"`
	VtpManagementDomain types.String `tfsdk:"vtp_management_domain"`
}

type ResponseItemSwitchGetDeviceSwitchPortsStatusesLldp struct {
	ChassisID          types.String `tfsdk:"chassis_id"`
	ManagementAddress  types.String `tfsdk:"management_address"`
	ManagementVLAN     types.Int64  `tfsdk:"management_vlan"`
	PortDescription    types.String `tfsdk:"port_description"`
	PortID             types.String `tfsdk:"port_id"`
	PortVLAN           types.Int64  `tfsdk:"port_vlan"`
	SystemCapabilities types.String `tfsdk:"system_capabilities"`
	SystemDescription  types.String `tfsdk:"system_description"`
	SystemName         types.String `tfsdk:"system_name"`
}

type ResponseItemSwitchGetDeviceSwitchPortsStatusesPoe struct {
	IsAllocated types.Bool `tfsdk:"is_allocated"`
}

type ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePort struct {
	Active               types.Bool                                                               `tfsdk:"active"`
	AuthenticationStatus types.String                                                             `tfsdk:"authentication_status"`
	ConfigOverrides      *ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePortConfigOverrides `tfsdk:"config_overrides"`
	Enabled              types.Bool                                                               `tfsdk:"enabled"`
}

type ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePortConfigOverrides struct {
	AllowedVLANs types.String `tfsdk:"allowed_vlans"`
	Type         types.String `tfsdk:"type"`
	VLAN         types.Int64  `tfsdk:"vlan"`
	VoiceVLAN    types.Int64  `tfsdk:"voice_vlan"`
}

type ResponseItemSwitchGetDeviceSwitchPortsStatusesSpanningTree struct {
	Statuses types.List `tfsdk:"statuses"`
}

type ResponseItemSwitchGetDeviceSwitchPortsStatusesTrafficInKbps struct {
	Recv  types.Float64 `tfsdk:"recv"`
	Sent  types.Float64 `tfsdk:"sent"`
	Total types.Float64 `tfsdk:"total"`
}

type ResponseItemSwitchGetDeviceSwitchPortsStatusesUsageInKb struct {
	Recv  types.Int64 `tfsdk:"recv"`
	Sent  types.Int64 `tfsdk:"sent"`
	Total types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseSwitchGetDeviceSwitchPortsStatusesItemsToBody(state DevicesSwitchPortsStatuses, response *merakigosdk.ResponseSwitchGetDeviceSwitchPortsStatuses) DevicesSwitchPortsStatuses {
	var items []ResponseItemSwitchGetDeviceSwitchPortsStatuses
	for _, item := range *response {
		itemState := ResponseItemSwitchGetDeviceSwitchPortsStatuses{
			Cdp: func() *ResponseItemSwitchGetDeviceSwitchPortsStatusesCdp {
				if item.Cdp != nil {
					return &ResponseItemSwitchGetDeviceSwitchPortsStatusesCdp{
						Address:           types.StringValue(item.Cdp.Address),
						Capabilities:      types.StringValue(item.Cdp.Capabilities),
						DeviceID:          types.StringValue(item.Cdp.DeviceID),
						ManagementAddress: types.StringValue(item.Cdp.ManagementAddress),
						NativeVLAN: func() types.Int64 {
							if item.Cdp.NativeVLAN != nil {
								return types.Int64Value(int64(*item.Cdp.NativeVLAN))
							}
							return types.Int64{}
						}(),
						Platform:            types.StringValue(item.Cdp.Platform),
						PortID:              types.StringValue(item.Cdp.PortID),
						SystemName:          types.StringValue(item.Cdp.SystemName),
						Version:             types.StringValue(item.Cdp.Version),
						VtpManagementDomain: types.StringValue(item.Cdp.VtpManagementDomain),
					}
				}
				return nil
			}(),
			ClientCount: func() types.Int64 {
				if item.ClientCount != nil {
					return types.Int64Value(int64(*item.ClientCount))
				}
				return types.Int64{}
			}(),
			Duplex: types.StringValue(item.Duplex),
			Enabled: func() types.Bool {
				if item.Enabled != nil {
					return types.BoolValue(*item.Enabled)
				}
				return types.Bool{}
			}(),
			Errors: StringSliceToList(item.Errors),
			IsUplink: func() types.Bool {
				if item.IsUplink != nil {
					return types.BoolValue(*item.IsUplink)
				}
				return types.Bool{}
			}(),
			Lldp: func() *ResponseItemSwitchGetDeviceSwitchPortsStatusesLldp {
				if item.Lldp != nil {
					return &ResponseItemSwitchGetDeviceSwitchPortsStatusesLldp{
						ChassisID:         types.StringValue(item.Lldp.ChassisID),
						ManagementAddress: types.StringValue(item.Lldp.ManagementAddress),
						ManagementVLAN: func() types.Int64 {
							if item.Lldp.ManagementVLAN != nil {
								return types.Int64Value(int64(*item.Lldp.ManagementVLAN))
							}
							return types.Int64{}
						}(),
						PortDescription: types.StringValue(item.Lldp.PortDescription),
						PortID:          types.StringValue(item.Lldp.PortID),
						PortVLAN: func() types.Int64 {
							if item.Lldp.PortVLAN != nil {
								return types.Int64Value(int64(*item.Lldp.PortVLAN))
							}
							return types.Int64{}
						}(),
						SystemCapabilities: types.StringValue(item.Lldp.SystemCapabilities),
						SystemDescription:  types.StringValue(item.Lldp.SystemDescription),
						SystemName:         types.StringValue(item.Lldp.SystemName),
					}
				}
				return nil
			}(),
			Poe: func() *ResponseItemSwitchGetDeviceSwitchPortsStatusesPoe {
				if item.Poe != nil {
					return &ResponseItemSwitchGetDeviceSwitchPortsStatusesPoe{
						IsAllocated: func() types.Bool {
							if item.Poe.IsAllocated != nil {
								return types.BoolValue(*item.Poe.IsAllocated)
							}
							return types.Bool{}
						}(),
					}
				}
				return nil
			}(),
			PortID: types.StringValue(item.PortID),
			PowerUsageInWh: func() types.Float64 {
				if item.PowerUsageInWh != nil {
					return types.Float64Value(float64(*item.PowerUsageInWh))
				}
				return types.Float64{}
			}(),
			SecurePort: func() *ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePort {
				if item.SecurePort != nil {
					return &ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePort{
						Active: func() types.Bool {
							if item.SecurePort.Active != nil {
								return types.BoolValue(*item.SecurePort.Active)
							}
							return types.Bool{}
						}(),
						AuthenticationStatus: types.StringValue(item.SecurePort.AuthenticationStatus),
						ConfigOverrides: func() *ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePortConfigOverrides {
							if item.SecurePort.ConfigOverrides != nil {
								return &ResponseItemSwitchGetDeviceSwitchPortsStatusesSecurePortConfigOverrides{
									AllowedVLANs: types.StringValue(item.SecurePort.ConfigOverrides.AllowedVLANs),
									Type:         types.StringValue(item.SecurePort.ConfigOverrides.Type),
									VLAN: func() types.Int64 {
										if item.SecurePort.ConfigOverrides.VLAN != nil {
											return types.Int64Value(int64(*item.SecurePort.ConfigOverrides.VLAN))
										}
										return types.Int64{}
									}(),
									VoiceVLAN: func() types.Int64 {
										if item.SecurePort.ConfigOverrides.VoiceVLAN != nil {
											return types.Int64Value(int64(*item.SecurePort.ConfigOverrides.VoiceVLAN))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						Enabled: func() types.Bool {
							if item.SecurePort.Enabled != nil {
								return types.BoolValue(*item.SecurePort.Enabled)
							}
							return types.Bool{}
						}(),
					}
				}
				return nil
			}(),
			SpanningTree: func() *ResponseItemSwitchGetDeviceSwitchPortsStatusesSpanningTree {
				if item.SpanningTree != nil {
					return &ResponseItemSwitchGetDeviceSwitchPortsStatusesSpanningTree{
						Statuses: StringSliceToList(item.SpanningTree.Statuses),
					}
				}
				return nil
			}(),
			Speed:  types.StringValue(item.Speed),
			Status: types.StringValue(item.Status),
			TrafficInKbps: func() *ResponseItemSwitchGetDeviceSwitchPortsStatusesTrafficInKbps {
				if item.TrafficInKbps != nil {
					return &ResponseItemSwitchGetDeviceSwitchPortsStatusesTrafficInKbps{
						Recv: func() types.Float64 {
							if item.TrafficInKbps.Recv != nil {
								return types.Float64Value(float64(*item.TrafficInKbps.Recv))
							}
							return types.Float64{}
						}(),
						Sent: func() types.Float64 {
							if item.TrafficInKbps.Sent != nil {
								return types.Float64Value(float64(*item.TrafficInKbps.Sent))
							}
							return types.Float64{}
						}(),
						Total: func() types.Float64 {
							if item.TrafficInKbps.Total != nil {
								return types.Float64Value(float64(*item.TrafficInKbps.Total))
							}
							return types.Float64{}
						}(),
					}
				}
				return nil
			}(),
			UsageInKb: func() *ResponseItemSwitchGetDeviceSwitchPortsStatusesUsageInKb {
				if item.UsageInKb != nil {
					return &ResponseItemSwitchGetDeviceSwitchPortsStatusesUsageInKb{
						Recv: func() types.Int64 {
							if item.UsageInKb.Recv != nil {
								return types.Int64Value(int64(*item.UsageInKb.Recv))
							}
							return types.Int64{}
						}(),
						Sent: func() types.Int64 {
							if item.UsageInKb.Sent != nil {
								return types.Int64Value(int64(*item.UsageInKb.Sent))
							}
							return types.Int64{}
						}(),
						Total: func() types.Int64 {
							if item.UsageInKb.Total != nil {
								return types.Int64Value(int64(*item.UsageInKb.Total))
							}
							return types.Int64{}
						}(),
					}
				}
				return nil
			}(),
			Warnings: StringSliceToList(item.Warnings),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
