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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DevicesSwitchPortsDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesSwitchPortsDataSource{}
)

func NewDevicesSwitchPortsDataSource() datasource.DataSource {
	return &DevicesSwitchPortsDataSource{}
}

type DevicesSwitchPortsDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesSwitchPortsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesSwitchPortsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_switch_ports"
}

func (d *DevicesSwitchPortsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"port_id": schema.StringAttribute{
				MarkdownDescription: `portId path parameter. Port ID`,
				Optional:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"access_policy_number": schema.Int64Attribute{
						MarkdownDescription: `The number of a custom access policy to configure on the switch port. Only applicable when 'accessPolicyType' is 'Custom access policy'.`,
						Computed:            true,
					},
					"access_policy_type": schema.StringAttribute{
						MarkdownDescription: `The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.`,
						Computed:            true,
					},
					"adaptive_policy_group_id": schema.StringAttribute{
						MarkdownDescription: `The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile.`,
						Computed:            true,
					},
					"allowed_vlans": schema.StringAttribute{
						MarkdownDescription: `The VLANs allowed on the switch port. Only applicable to trunk ports.`,
						Computed:            true,
					},
					"dai_trusted": schema.BoolAttribute{
						MarkdownDescription: `If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.`,
						Computed:            true,
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `The status of the switch port.`,
						Computed:            true,
					},
					"flexible_stacking_enabled": schema.BoolAttribute{
						MarkdownDescription: `For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.`,
						Computed:            true,
					},
					"isolation_enabled": schema.BoolAttribute{
						MarkdownDescription: `The isolation status of the switch port.`,
						Computed:            true,
					},
					"link_negotiation": schema.StringAttribute{
						MarkdownDescription: `The link speed for the switch port.`,
						Computed:            true,
					},
					"link_negotiation_capabilities": schema.ListAttribute{
						MarkdownDescription: `Available link speeds for the switch port.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"mac_allow_list": schema.ListAttribute{
						MarkdownDescription: `Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the switch port.`,
						Computed:            true,
					},
					"peer_sgt_capable": schema.BoolAttribute{
						MarkdownDescription: `If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.`,
						Computed:            true,
					},
					"poe_enabled": schema.BoolAttribute{
						MarkdownDescription: `The PoE status of the switch port.`,
						Computed:            true,
					},
					"port_id": schema.StringAttribute{
						MarkdownDescription: `The identifier of the switch port.`,
						Computed:            true,
					},
					"port_schedule_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the port schedule. A value of null will clear the port schedule.`,
						Computed:            true,
					},
					"profile": schema.SingleNestedAttribute{
						MarkdownDescription: `Profile attributes`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `When enabled, override this port's configuration with a port profile.`,
								Computed:            true,
							},
							"id": schema.StringAttribute{
								MarkdownDescription: `When enabled, the ID of the port profile used to override the port's configuration.`,
								Computed:            true,
							},
							"iname": schema.StringAttribute{
								MarkdownDescription: `When enabled, the IName of the profile.`,
								Computed:            true,
							},
						},
					},
					"rstp_enabled": schema.BoolAttribute{
						MarkdownDescription: `The rapid spanning tree protocol status.`,
						Computed:            true,
					},
					"sticky_mac_allow_list": schema.ListAttribute{
						MarkdownDescription: `The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"sticky_mac_allow_list_limit": schema.Int64Attribute{
						MarkdownDescription: `The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.`,
						Computed:            true,
					},
					"storm_control_enabled": schema.BoolAttribute{
						MarkdownDescription: `The storm control status of the switch port.`,
						Computed:            true,
					},
					"stp_guard": schema.StringAttribute{
						MarkdownDescription: `The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').`,
						Computed:            true,
					},
					"tags": schema.ListAttribute{
						MarkdownDescription: `The list of tags of the switch port.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"type": schema.StringAttribute{
						MarkdownDescription: `The type of the switch port ('trunk' or 'access').`,
						Computed:            true,
					},
					"udld": schema.StringAttribute{
						MarkdownDescription: `The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.`,
						Computed:            true,
					},
					"vlan": schema.Int64Attribute{
						MarkdownDescription: `The VLAN of the switch port. A null value will clear the value set for trunk ports.`,
						Computed:            true,
					},
					"voice_vlan": schema.Int64Attribute{
						MarkdownDescription: `The voice VLAN of the switch port. Only applicable to access ports.`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetDeviceSwitchPorts`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"access_policy_number": schema.Int64Attribute{
							MarkdownDescription: `The number of a custom access policy to configure on the switch port. Only applicable when 'accessPolicyType' is 'Custom access policy'.`,
							Computed:            true,
						},
						"access_policy_type": schema.StringAttribute{
							MarkdownDescription: `The type of the access policy of the switch port. Only applicable to access ports. Can be one of 'Open', 'Custom access policy', 'MAC allow list' or 'Sticky MAC allow list'.`,
							Computed:            true,
						},
						"adaptive_policy_group_id": schema.StringAttribute{
							MarkdownDescription: `The adaptive policy group ID that will be used to tag traffic through this switch port. This ID must pre-exist during the configuration, else needs to be created using adaptivePolicy/groups API. Cannot be applied to a port on a switch bound to profile.`,
							Computed:            true,
						},
						"allowed_vlans": schema.StringAttribute{
							MarkdownDescription: `The VLANs allowed on the switch port. Only applicable to trunk ports.`,
							Computed:            true,
						},
						"dai_trusted": schema.BoolAttribute{
							MarkdownDescription: `If true, ARP packets for this port will be considered trusted, and Dynamic ARP Inspection will allow the traffic.`,
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: `The status of the switch port.`,
							Computed:            true,
						},
						"flexible_stacking_enabled": schema.BoolAttribute{
							MarkdownDescription: `For supported switches (e.g. MS420/MS425), whether or not the port has flexible stacking enabled.`,
							Computed:            true,
						},
						"isolation_enabled": schema.BoolAttribute{
							MarkdownDescription: `The isolation status of the switch port.`,
							Computed:            true,
						},
						"link_negotiation": schema.StringAttribute{
							MarkdownDescription: `The link speed for the switch port.`,
							Computed:            true,
						},
						"link_negotiation_capabilities": schema.ListAttribute{
							MarkdownDescription: `Available link speeds for the switch port.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"mac_allow_list": schema.ListAttribute{
							MarkdownDescription: `Only devices with MAC addresses specified in this list will have access to this port. Up to 20 MAC addresses can be defined. Only applicable when 'accessPolicyType' is 'MAC allow list'.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the switch port.`,
							Computed:            true,
						},
						"peer_sgt_capable": schema.BoolAttribute{
							MarkdownDescription: `If true, Peer SGT is enabled for traffic through this switch port. Applicable to trunk port only, not access port. Cannot be applied to a port on a switch bound to profile.`,
							Computed:            true,
						},
						"poe_enabled": schema.BoolAttribute{
							MarkdownDescription: `The PoE status of the switch port.`,
							Computed:            true,
						},
						"port_id": schema.StringAttribute{
							MarkdownDescription: `The identifier of the switch port.`,
							Computed:            true,
						},
						"port_schedule_id": schema.StringAttribute{
							MarkdownDescription: `The ID of the port schedule. A value of null will clear the port schedule.`,
							Computed:            true,
						},
						"profile": schema.SingleNestedAttribute{
							MarkdownDescription: `Profile attributes`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `When enabled, override this port's configuration with a port profile.`,
									Computed:            true,
								},
								"id": schema.StringAttribute{
									MarkdownDescription: `When enabled, the ID of the port profile used to override the port's configuration.`,
									Computed:            true,
								},
								"iname": schema.StringAttribute{
									MarkdownDescription: `When enabled, the IName of the profile.`,
									Computed:            true,
								},
							},
						},
						"rstp_enabled": schema.BoolAttribute{
							MarkdownDescription: `The rapid spanning tree protocol status.`,
							Computed:            true,
						},
						"sticky_mac_allow_list": schema.ListAttribute{
							MarkdownDescription: `The initial list of MAC addresses for sticky Mac allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"sticky_mac_allow_list_limit": schema.Int64Attribute{
							MarkdownDescription: `The maximum number of MAC addresses for sticky MAC allow list. Only applicable when 'accessPolicyType' is 'Sticky MAC allow list'.`,
							Computed:            true,
						},
						"storm_control_enabled": schema.BoolAttribute{
							MarkdownDescription: `The storm control status of the switch port.`,
							Computed:            true,
						},
						"stp_guard": schema.StringAttribute{
							MarkdownDescription: `The state of the STP guard ('disabled', 'root guard', 'bpdu guard' or 'loop guard').`,
							Computed:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `The list of tags of the switch port.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The type of the switch port ('trunk' or 'access').`,
							Computed:            true,
						},
						"udld": schema.StringAttribute{
							MarkdownDescription: `The action to take when Unidirectional Link is detected (Alert only, Enforce). Default configuration is Alert only.`,
							Computed:            true,
						},
						"vlan": schema.Int64Attribute{
							MarkdownDescription: `The VLAN of the switch port. A null value will clear the value set for trunk ports.`,
							Computed:            true,
						},
						"voice_vlan": schema.Int64Attribute{
							MarkdownDescription: `The voice VLAN of the switch port. Only applicable to access ports.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *DevicesSwitchPortsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesSwitchPorts DevicesSwitchPorts
	diags := req.Config.Get(ctx, &devicesSwitchPorts)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!devicesSwitchPorts.Serial.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!devicesSwitchPorts.Serial.IsNull(), !devicesSwitchPorts.PortID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceSwitchPorts")
		vvSerial := devicesSwitchPorts.Serial.ValueString()

		response1, restyResp1, err := d.client.Switch.GetDeviceSwitchPorts(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchPorts",
				err.Error(),
			)
			return
		}

		devicesSwitchPorts = ResponseSwitchGetDeviceSwitchPortsItemsToBody(devicesSwitchPorts, response1)
		diags = resp.State.Set(ctx, &devicesSwitchPorts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDeviceSwitchPort")
		vvSerial := devicesSwitchPorts.Serial.ValueString()
		vvPortID := devicesSwitchPorts.PortID.ValueString()

		response2, restyResp2, err := d.client.Switch.GetDeviceSwitchPort(vvSerial, vvPortID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchPort",
				err.Error(),
			)
			return
		}

		devicesSwitchPorts = ResponseSwitchGetDeviceSwitchPortItemToBody(devicesSwitchPorts, response2)
		diags = resp.State.Set(ctx, &devicesSwitchPorts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesSwitchPorts struct {
	Serial types.String                              `tfsdk:"serial"`
	PortID types.String                              `tfsdk:"port_id"`
	Items  *[]ResponseItemSwitchGetDeviceSwitchPorts `tfsdk:"items"`
	Item   *ResponseSwitchGetDeviceSwitchPort        `tfsdk:"item"`
}

type ResponseItemSwitchGetDeviceSwitchPorts struct {
	AccessPolicyNumber          types.Int64                                    `tfsdk:"access_policy_number"`
	AccessPolicyType            types.String                                   `tfsdk:"access_policy_type"`
	AdaptivePolicyGroupID       types.String                                   `tfsdk:"adaptive_policy_group_id"`
	AllowedVLANs                types.String                                   `tfsdk:"allowed_vlans"`
	DaiTrusted                  types.Bool                                     `tfsdk:"dai_trusted"`
	Enabled                     types.Bool                                     `tfsdk:"enabled"`
	FlexibleStackingEnabled     types.Bool                                     `tfsdk:"flexible_stacking_enabled"`
	IsolationEnabled            types.Bool                                     `tfsdk:"isolation_enabled"`
	LinkNegotiation             types.String                                   `tfsdk:"link_negotiation"`
	LinkNegotiationCapabilities types.List                                     `tfsdk:"link_negotiation_capabilities"`
	MacAllowList                types.List                                     `tfsdk:"mac_allow_list"`
	Name                        types.String                                   `tfsdk:"name"`
	PeerSgtCapable              types.Bool                                     `tfsdk:"peer_sgt_capable"`
	PoeEnabled                  types.Bool                                     `tfsdk:"poe_enabled"`
	PortID                      types.String                                   `tfsdk:"port_id"`
	PortScheduleID              types.String                                   `tfsdk:"port_schedule_id"`
	Profile                     *ResponseItemSwitchGetDeviceSwitchPortsProfile `tfsdk:"profile"`
	RstpEnabled                 types.Bool                                     `tfsdk:"rstp_enabled"`
	StickyMacAllowList          types.List                                     `tfsdk:"sticky_mac_allow_list"`
	StickyMacAllowListLimit     types.Int64                                    `tfsdk:"sticky_mac_allow_list_limit"`
	StormControlEnabled         types.Bool                                     `tfsdk:"storm_control_enabled"`
	StpGuard                    types.String                                   `tfsdk:"stp_guard"`
	Tags                        types.List                                     `tfsdk:"tags"`
	Type                        types.String                                   `tfsdk:"type"`
	Udld                        types.String                                   `tfsdk:"udld"`
	VLAN                        types.Int64                                    `tfsdk:"vlan"`
	VoiceVLAN                   types.Int64                                    `tfsdk:"voice_vlan"`
}

type ResponseItemSwitchGetDeviceSwitchPortsProfile struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	ID      types.String `tfsdk:"id"`
	Iname   types.String `tfsdk:"iname"`
}

type ResponseSwitchGetDeviceSwitchPort struct {
	AccessPolicyNumber          types.Int64                               `tfsdk:"access_policy_number"`
	AccessPolicyType            types.String                              `tfsdk:"access_policy_type"`
	AdaptivePolicyGroupID       types.String                              `tfsdk:"adaptive_policy_group_id"`
	AllowedVLANs                types.String                              `tfsdk:"allowed_vlans"`
	DaiTrusted                  types.Bool                                `tfsdk:"dai_trusted"`
	Enabled                     types.Bool                                `tfsdk:"enabled"`
	FlexibleStackingEnabled     types.Bool                                `tfsdk:"flexible_stacking_enabled"`
	IsolationEnabled            types.Bool                                `tfsdk:"isolation_enabled"`
	LinkNegotiation             types.String                              `tfsdk:"link_negotiation"`
	LinkNegotiationCapabilities types.List                                `tfsdk:"link_negotiation_capabilities"`
	MacAllowList                types.List                                `tfsdk:"mac_allow_list"`
	Name                        types.String                              `tfsdk:"name"`
	PeerSgtCapable              types.Bool                                `tfsdk:"peer_sgt_capable"`
	PoeEnabled                  types.Bool                                `tfsdk:"poe_enabled"`
	PortID                      types.String                              `tfsdk:"port_id"`
	PortScheduleID              types.String                              `tfsdk:"port_schedule_id"`
	Profile                     *ResponseSwitchGetDeviceSwitchPortProfile `tfsdk:"profile"`
	RstpEnabled                 types.Bool                                `tfsdk:"rstp_enabled"`
	StickyMacAllowList          types.List                                `tfsdk:"sticky_mac_allow_list"`
	StickyMacAllowListLimit     types.Int64                               `tfsdk:"sticky_mac_allow_list_limit"`
	StormControlEnabled         types.Bool                                `tfsdk:"storm_control_enabled"`
	StpGuard                    types.String                              `tfsdk:"stp_guard"`
	Tags                        types.List                                `tfsdk:"tags"`
	Type                        types.String                              `tfsdk:"type"`
	Udld                        types.String                              `tfsdk:"udld"`
	VLAN                        types.Int64                               `tfsdk:"vlan"`
	VoiceVLAN                   types.Int64                               `tfsdk:"voice_vlan"`
}

type ResponseSwitchGetDeviceSwitchPortProfile struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	ID      types.String `tfsdk:"id"`
	Iname   types.String `tfsdk:"iname"`
}

// ToBody
func ResponseSwitchGetDeviceSwitchPortsItemsToBody(state DevicesSwitchPorts, response *merakigosdk.ResponseSwitchGetDeviceSwitchPorts) DevicesSwitchPorts {
	var items []ResponseItemSwitchGetDeviceSwitchPorts
	for _, item := range *response {
		itemState := ResponseItemSwitchGetDeviceSwitchPorts{
			AccessPolicyNumber: func() types.Int64 {
				if item.AccessPolicyNumber != nil {
					return types.Int64Value(int64(*item.AccessPolicyNumber))
				}
				return types.Int64{}
			}(),
			AccessPolicyType:      types.StringValue(item.AccessPolicyType),
			AdaptivePolicyGroupID: types.StringValue(item.AdaptivePolicyGroupID),
			AllowedVLANs:          types.StringValue(item.AllowedVLANs),
			DaiTrusted: func() types.Bool {
				if item.DaiTrusted != nil {
					return types.BoolValue(*item.DaiTrusted)
				}
				return types.Bool{}
			}(),
			Enabled: func() types.Bool {
				if item.Enabled != nil {
					return types.BoolValue(*item.Enabled)
				}
				return types.Bool{}
			}(),
			FlexibleStackingEnabled: func() types.Bool {
				if item.FlexibleStackingEnabled != nil {
					return types.BoolValue(*item.FlexibleStackingEnabled)
				}
				return types.Bool{}
			}(),
			IsolationEnabled: func() types.Bool {
				if item.IsolationEnabled != nil {
					return types.BoolValue(*item.IsolationEnabled)
				}
				return types.Bool{}
			}(),
			LinkNegotiation:             types.StringValue(item.LinkNegotiation),
			LinkNegotiationCapabilities: StringSliceToList(item.LinkNegotiationCapabilities),
			MacAllowList:                StringSliceToList(item.MacAllowList),
			Name:                        types.StringValue(item.Name),
			PeerSgtCapable: func() types.Bool {
				if item.PeerSgtCapable != nil {
					return types.BoolValue(*item.PeerSgtCapable)
				}
				return types.Bool{}
			}(),
			PoeEnabled: func() types.Bool {
				if item.PoeEnabled != nil {
					return types.BoolValue(*item.PoeEnabled)
				}
				return types.Bool{}
			}(),
			PortID:         types.StringValue(item.PortID),
			PortScheduleID: types.StringValue(item.PortScheduleID),
			Profile: func() *ResponseItemSwitchGetDeviceSwitchPortsProfile {
				if item.Profile != nil {
					return &ResponseItemSwitchGetDeviceSwitchPortsProfile{
						Enabled: func() types.Bool {
							if item.Profile.Enabled != nil {
								return types.BoolValue(*item.Profile.Enabled)
							}
							return types.Bool{}
						}(),
						ID:    types.StringValue(item.Profile.ID),
						Iname: types.StringValue(item.Profile.Iname),
					}
				}
				return &ResponseItemSwitchGetDeviceSwitchPortsProfile{}
			}(),
			RstpEnabled: func() types.Bool {
				if item.RstpEnabled != nil {
					return types.BoolValue(*item.RstpEnabled)
				}
				return types.Bool{}
			}(),
			StickyMacAllowList: StringSliceToList(item.StickyMacAllowList),
			StickyMacAllowListLimit: func() types.Int64 {
				if item.StickyMacAllowListLimit != nil {
					return types.Int64Value(int64(*item.StickyMacAllowListLimit))
				}
				return types.Int64{}
			}(),
			StormControlEnabled: func() types.Bool {
				if item.StormControlEnabled != nil {
					return types.BoolValue(*item.StormControlEnabled)
				}
				return types.Bool{}
			}(),
			StpGuard: types.StringValue(item.StpGuard),
			Tags:     StringSliceToList(item.Tags),
			Type:     types.StringValue(item.Type),
			Udld:     types.StringValue(item.Udld),
			VLAN: func() types.Int64 {
				if item.VLAN != nil {
					return types.Int64Value(int64(*item.VLAN))
				}
				return types.Int64{}
			}(),
			VoiceVLAN: func() types.Int64 {
				if item.VoiceVLAN != nil {
					return types.Int64Value(int64(*item.VoiceVLAN))
				}
				return types.Int64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSwitchGetDeviceSwitchPortItemToBody(state DevicesSwitchPorts, response *merakigosdk.ResponseSwitchGetDeviceSwitchPort) DevicesSwitchPorts {
	itemState := ResponseSwitchGetDeviceSwitchPort{
		AccessPolicyNumber: func() types.Int64 {
			if response.AccessPolicyNumber != nil {
				return types.Int64Value(int64(*response.AccessPolicyNumber))
			}
			return types.Int64{}
		}(),
		AccessPolicyType:      types.StringValue(response.AccessPolicyType),
		AdaptivePolicyGroupID: types.StringValue(response.AdaptivePolicyGroupID),
		AllowedVLANs:          types.StringValue(response.AllowedVLANs),
		DaiTrusted: func() types.Bool {
			if response.DaiTrusted != nil {
				return types.BoolValue(*response.DaiTrusted)
			}
			return types.Bool{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		FlexibleStackingEnabled: func() types.Bool {
			if response.FlexibleStackingEnabled != nil {
				return types.BoolValue(*response.FlexibleStackingEnabled)
			}
			return types.Bool{}
		}(),
		IsolationEnabled: func() types.Bool {
			if response.IsolationEnabled != nil {
				return types.BoolValue(*response.IsolationEnabled)
			}
			return types.Bool{}
		}(),
		LinkNegotiation:             types.StringValue(response.LinkNegotiation),
		LinkNegotiationCapabilities: StringSliceToList(response.LinkNegotiationCapabilities),
		MacAllowList:                StringSliceToList(response.MacAllowList),
		Name:                        types.StringValue(response.Name),
		PeerSgtCapable: func() types.Bool {
			if response.PeerSgtCapable != nil {
				return types.BoolValue(*response.PeerSgtCapable)
			}
			return types.Bool{}
		}(),
		PoeEnabled: func() types.Bool {
			if response.PoeEnabled != nil {
				return types.BoolValue(*response.PoeEnabled)
			}
			return types.Bool{}
		}(),
		PortID:         types.StringValue(response.PortID),
		PortScheduleID: types.StringValue(response.PortScheduleID),
		Profile: func() *ResponseSwitchGetDeviceSwitchPortProfile {
			if response.Profile != nil {
				return &ResponseSwitchGetDeviceSwitchPortProfile{
					Enabled: func() types.Bool {
						if response.Profile.Enabled != nil {
							return types.BoolValue(*response.Profile.Enabled)
						}
						return types.Bool{}
					}(),
					ID:    types.StringValue(response.Profile.ID),
					Iname: types.StringValue(response.Profile.Iname),
				}
			}
			return &ResponseSwitchGetDeviceSwitchPortProfile{}
		}(),
		RstpEnabled: func() types.Bool {
			if response.RstpEnabled != nil {
				return types.BoolValue(*response.RstpEnabled)
			}
			return types.Bool{}
		}(),
		StickyMacAllowList: StringSliceToList(response.StickyMacAllowList),
		StickyMacAllowListLimit: func() types.Int64 {
			if response.StickyMacAllowListLimit != nil {
				return types.Int64Value(int64(*response.StickyMacAllowListLimit))
			}
			return types.Int64{}
		}(),
		StormControlEnabled: func() types.Bool {
			if response.StormControlEnabled != nil {
				return types.BoolValue(*response.StormControlEnabled)
			}
			return types.Bool{}
		}(),
		StpGuard: types.StringValue(response.StpGuard),
		Tags:     StringSliceToList(response.Tags),
		Type:     types.StringValue(response.Type),
		Udld:     types.StringValue(response.Udld),
		VLAN: func() types.Int64 {
			if response.VLAN != nil {
				return types.Int64Value(int64(*response.VLAN))
			}
			return types.Int64{}
		}(),
		VoiceVLAN: func() types.Int64 {
			if response.VoiceVLAN != nil {
				return types.Int64Value(int64(*response.VoiceVLAN))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
