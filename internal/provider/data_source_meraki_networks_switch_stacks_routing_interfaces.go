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
	_ datasource.DataSource              = &NetworksSwitchStacksRoutingInterfacesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchStacksRoutingInterfacesDataSource{}
)

func NewNetworksSwitchStacksRoutingInterfacesDataSource() datasource.DataSource {
	return &NetworksSwitchStacksRoutingInterfacesDataSource{}
}

type NetworksSwitchStacksRoutingInterfacesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchStacksRoutingInterfacesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchStacksRoutingInterfacesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stacks_routing_interfaces"
}

func (d *NetworksSwitchStacksRoutingInterfacesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"interface_id": schema.StringAttribute{
				MarkdownDescription: `interfaceId path parameter. Interface ID`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"protocol": schema.StringAttribute{
				MarkdownDescription: `protocol query parameter. Optional parameter to filter L3 interfaces by protocol.`,
				Optional:            true,
			},
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: `switchStackId path parameter. Switch stack ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"default_gateway": schema.StringAttribute{
						MarkdownDescription: `IPv4 default gateway`,
						Computed:            true,
					},
					"interface_id": schema.StringAttribute{
						MarkdownDescription: `The ID`,
						Computed:            true,
					},
					"interface_ip": schema.StringAttribute{
						MarkdownDescription: `IPv4 address`,
						Computed:            true,
					},
					"ipv6": schema.SingleNestedAttribute{
						MarkdownDescription: `IPv6 addressing`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"address": schema.StringAttribute{
								MarkdownDescription: `IPv6 address`,
								Computed:            true,
							},
							"assignment_mode": schema.StringAttribute{
								MarkdownDescription: `Assignment mode`,
								Computed:            true,
							},
							"gateway": schema.StringAttribute{
								MarkdownDescription: `IPv6 gateway`,
								Computed:            true,
							},
							"prefix": schema.StringAttribute{
								MarkdownDescription: `IPv6 subnet`,
								Computed:            true,
							},
						},
					},
					"multicast_routing": schema.StringAttribute{
						MarkdownDescription: `Multicast routing status`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name`,
						Computed:            true,
					},
					"ospf_settings": schema.SingleNestedAttribute{
						MarkdownDescription: `IPv4 OSPF Settings`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"area": schema.StringAttribute{
								MarkdownDescription: `Area ID`,
								Computed:            true,
							},
							"cost": schema.Int64Attribute{
								MarkdownDescription: `OSPF Cost`,
								Computed:            true,
							},
							"is_passive_enabled": schema.BoolAttribute{
								MarkdownDescription: `Disable sending Hello packets on this interface's IPv4 area`,
								Computed:            true,
							},
						},
					},
					"ospf_v3": schema.SingleNestedAttribute{
						MarkdownDescription: `IPv6 OSPF Settings`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"area": schema.StringAttribute{
								MarkdownDescription: `Area ID`,
								Computed:            true,
							},
							"cost": schema.Int64Attribute{
								MarkdownDescription: `OSPF Cost`,
								Computed:            true,
							},
							"is_passive_enabled": schema.BoolAttribute{
								MarkdownDescription: `Disable sending Hello packets on this interface's IPv6 area`,
								Computed:            true,
							},
						},
					},
					"subnet": schema.StringAttribute{
						MarkdownDescription: `IPv4 subnet`,
						Computed:            true,
					},
					"uplink_v4": schema.BoolAttribute{
						MarkdownDescription: `When true, this interface is used as static IPv4 uplink`,
						Computed:            true,
					},
					"uplink_v6": schema.BoolAttribute{
						MarkdownDescription: `When true, this interface is used as static IPv6 uplink`,
						Computed:            true,
					},
					"vlan_id": schema.Int64Attribute{
						MarkdownDescription: `VLAN ID`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchStackRoutingInterfaces`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"default_gateway": schema.StringAttribute{
							MarkdownDescription: `IPv4 default gateway`,
							Computed:            true,
						},
						"interface_id": schema.StringAttribute{
							MarkdownDescription: `The ID`,
							Computed:            true,
						},
						"interface_ip": schema.StringAttribute{
							MarkdownDescription: `IPv4 address`,
							Computed:            true,
						},
						"ipv6": schema.SingleNestedAttribute{
							MarkdownDescription: `IPv6 addressing`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"address": schema.StringAttribute{
									MarkdownDescription: `IPv6 address`,
									Computed:            true,
								},
								"assignment_mode": schema.StringAttribute{
									MarkdownDescription: `Assignment mode`,
									Computed:            true,
								},
								"gateway": schema.StringAttribute{
									MarkdownDescription: `IPv6 gateway`,
									Computed:            true,
								},
								"prefix": schema.StringAttribute{
									MarkdownDescription: `IPv6 subnet`,
									Computed:            true,
								},
							},
						},
						"multicast_routing": schema.StringAttribute{
							MarkdownDescription: `Multicast routing status`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name`,
							Computed:            true,
						},
						"ospf_settings": schema.SingleNestedAttribute{
							MarkdownDescription: `IPv4 OSPF Settings`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"area": schema.StringAttribute{
									MarkdownDescription: `Area ID`,
									Computed:            true,
								},
								"cost": schema.Int64Attribute{
									MarkdownDescription: `OSPF Cost`,
									Computed:            true,
								},
								"is_passive_enabled": schema.BoolAttribute{
									MarkdownDescription: `Disable sending Hello packets on this interface's IPv4 area`,
									Computed:            true,
								},
							},
						},
						"ospf_v3": schema.SingleNestedAttribute{
							MarkdownDescription: `IPv6 OSPF Settings`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"area": schema.StringAttribute{
									MarkdownDescription: `Area ID`,
									Computed:            true,
								},
								"cost": schema.Int64Attribute{
									MarkdownDescription: `OSPF Cost`,
									Computed:            true,
								},
								"is_passive_enabled": schema.BoolAttribute{
									MarkdownDescription: `Disable sending Hello packets on this interface's IPv6 area`,
									Computed:            true,
								},
							},
						},
						"subnet": schema.StringAttribute{
							MarkdownDescription: `IPv4 subnet`,
							Computed:            true,
						},
						"uplink_v4": schema.BoolAttribute{
							MarkdownDescription: `When true, this interface is used as static IPv4 uplink`,
							Computed:            true,
						},
						"uplink_v6": schema.BoolAttribute{
							MarkdownDescription: `When true, this interface is used as static IPv6 uplink`,
							Computed:            true,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: `VLAN ID`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchStacksRoutingInterfacesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchStacksRoutingInterfaces NetworksSwitchStacksRoutingInterfaces
	diags := req.Config.Get(ctx, &networksSwitchStacksRoutingInterfaces)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksSwitchStacksRoutingInterfaces.NetworkID.IsNull(), !networksSwitchStacksRoutingInterfaces.SwitchStackID.IsNull(), !networksSwitchStacksRoutingInterfaces.Protocol.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksSwitchStacksRoutingInterfaces.NetworkID.IsNull(), !networksSwitchStacksRoutingInterfaces.SwitchStackID.IsNull(), !networksSwitchStacksRoutingInterfaces.InterfaceID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStackRoutingInterfaces")
		vvNetworkID := networksSwitchStacksRoutingInterfaces.NetworkID.ValueString()
		vvSwitchStackID := networksSwitchStacksRoutingInterfaces.SwitchStackID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSwitchStackRoutingInterfacesQueryParams{}

		queryParams1.Protocol = networksSwitchStacksRoutingInterfaces.Protocol.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchStackRoutingInterfaces(vvNetworkID, vvSwitchStackID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStackRoutingInterfaces",
				err.Error(),
			)
			return
		}

		networksSwitchStacksRoutingInterfaces = ResponseSwitchGetNetworkSwitchStackRoutingInterfacesItemsToBody(networksSwitchStacksRoutingInterfaces, response1)
		diags = resp.State.Set(ctx, &networksSwitchStacksRoutingInterfaces)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStackRoutingInterface")
		vvNetworkID := networksSwitchStacksRoutingInterfaces.NetworkID.ValueString()
		vvSwitchStackID := networksSwitchStacksRoutingInterfaces.SwitchStackID.ValueString()
		vvInterfaceID := networksSwitchStacksRoutingInterfaces.InterfaceID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Switch.GetNetworkSwitchStackRoutingInterface(vvNetworkID, vvSwitchStackID, vvInterfaceID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStackRoutingInterface",
				err.Error(),
			)
			return
		}

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStackRoutingInterface",
				err.Error(),
			)
			return
		}

		networksSwitchStacksRoutingInterfaces = ResponseSwitchGetNetworkSwitchStackRoutingInterfaceItemToBody(networksSwitchStacksRoutingInterfaces, response2)
		diags = resp.State.Set(ctx, &networksSwitchStacksRoutingInterfaces)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchStacksRoutingInterfaces struct {
	NetworkID     types.String                                                `tfsdk:"network_id"`
	SwitchStackID types.String                                                `tfsdk:"switch_stack_id"`
	Protocol      types.String                                                `tfsdk:"protocol"`
	InterfaceID   types.String                                                `tfsdk:"interface_id"`
	Items         *[]ResponseItemSwitchGetNetworkSwitchStackRoutingInterfaces `tfsdk:"items"`
	Item          *ResponseSwitchGetNetworkSwitchStackRoutingInterface        `tfsdk:"item"`
}

type ResponseItemSwitchGetNetworkSwitchStackRoutingInterfaces struct {
	DefaultGateway   types.String                                                          `tfsdk:"default_gateway"`
	InterfaceID      types.String                                                          `tfsdk:"interface_id"`
	InterfaceIP      types.String                                                          `tfsdk:"interface_ip"`
	IPv6             *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesIpv6         `tfsdk:"ipv6"`
	MulticastRouting types.String                                                          `tfsdk:"multicast_routing"`
	Name             types.String                                                          `tfsdk:"name"`
	OspfSettings     *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfSettings `tfsdk:"ospf_settings"`
	OspfV3           *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfV3       `tfsdk:"ospf_v3"`
	Subnet           types.String                                                          `tfsdk:"subnet"`
	UplinkV4         types.Bool                                                            `tfsdk:"uplink_v4"`
	UplinkV6         types.Bool                                                            `tfsdk:"uplink_v6"`
	VLANID           types.Int64                                                           `tfsdk:"vlan_id"`
}

type ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesIpv6 struct {
	Address        types.String `tfsdk:"address"`
	AssignmentMode types.String `tfsdk:"assignment_mode"`
	Gateway        types.String `tfsdk:"gateway"`
	Prefix         types.String `tfsdk:"prefix"`
}

type ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfSettings struct {
	Area             types.String `tfsdk:"area"`
	Cost             types.Int64  `tfsdk:"cost"`
	IsPassiveEnabled types.Bool   `tfsdk:"is_passive_enabled"`
}

type ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfV3 struct {
	Area             types.String `tfsdk:"area"`
	Cost             types.Int64  `tfsdk:"cost"`
	IsPassiveEnabled types.Bool   `tfsdk:"is_passive_enabled"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterface struct {
	DefaultGateway   types.String                                                     `tfsdk:"default_gateway"`
	InterfaceID      types.String                                                     `tfsdk:"interface_id"`
	InterfaceIP      types.String                                                     `tfsdk:"interface_ip"`
	IPv6             *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIpv6         `tfsdk:"ipv6"`
	MulticastRouting types.String                                                     `tfsdk:"multicast_routing"`
	Name             types.String                                                     `tfsdk:"name"`
	OspfSettings     *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettings `tfsdk:"ospf_settings"`
	OspfV3           *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3       `tfsdk:"ospf_v3"`
	Subnet           types.String                                                     `tfsdk:"subnet"`
	UplinkV4         types.Bool                                                       `tfsdk:"uplink_v4"`
	UplinkV6         types.Bool                                                       `tfsdk:"uplink_v6"`
	VLANID           types.Int64                                                      `tfsdk:"vlan_id"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIpv6 struct {
	Address        types.String `tfsdk:"address"`
	AssignmentMode types.String `tfsdk:"assignment_mode"`
	Gateway        types.String `tfsdk:"gateway"`
	Prefix         types.String `tfsdk:"prefix"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettings struct {
	Area             types.String `tfsdk:"area"`
	Cost             types.Int64  `tfsdk:"cost"`
	IsPassiveEnabled types.Bool   `tfsdk:"is_passive_enabled"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3 struct {
	Area             types.String `tfsdk:"area"`
	Cost             types.Int64  `tfsdk:"cost"`
	IsPassiveEnabled types.Bool   `tfsdk:"is_passive_enabled"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchStackRoutingInterfacesItemsToBody(state NetworksSwitchStacksRoutingInterfaces, response *merakigosdk.ResponseSwitchGetNetworkSwitchStackRoutingInterfaces) NetworksSwitchStacksRoutingInterfaces {
	var items []ResponseItemSwitchGetNetworkSwitchStackRoutingInterfaces
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchStackRoutingInterfaces{
			DefaultGateway: func() types.String {
				if item.DefaultGateway != "" {
					return types.StringValue(item.DefaultGateway)
				}
				return types.String{}
			}(),
			InterfaceID: func() types.String {
				if item.InterfaceID != "" {
					return types.StringValue(item.InterfaceID)
				}
				return types.String{}
			}(),
			InterfaceIP: func() types.String {
				if item.InterfaceIP != "" {
					return types.StringValue(item.InterfaceIP)
				}
				return types.String{}
			}(),
			IPv6: func() *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesIpv6 {
				if item.IPv6 != nil {
					return &ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesIpv6{
						Address: func() types.String {
							if item.IPv6.Address != "" {
								return types.StringValue(item.IPv6.Address)
							}
							return types.String{}
						}(),
						AssignmentMode: func() types.String {
							if item.IPv6.AssignmentMode != "" {
								return types.StringValue(item.IPv6.AssignmentMode)
							}
							return types.String{}
						}(),
						Gateway: func() types.String {
							if item.IPv6.Gateway != "" {
								return types.StringValue(item.IPv6.Gateway)
							}
							return types.String{}
						}(),
						Prefix: func() types.String {
							if item.IPv6.Prefix != "" {
								return types.StringValue(item.IPv6.Prefix)
							}
							return types.String{}
						}(),
					}
				}
				return nil
			}(),
			MulticastRouting: func() types.String {
				if item.MulticastRouting != "" {
					return types.StringValue(item.MulticastRouting)
				}
				return types.String{}
			}(),
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
			OspfSettings: func() *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfSettings {
				if item.OspfSettings != nil {
					return &ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfSettings{
						Area: func() types.String {
							if item.OspfSettings.Area != "" {
								return types.StringValue(item.OspfSettings.Area)
							}
							return types.String{}
						}(),
						Cost: func() types.Int64 {
							if item.OspfSettings.Cost != nil {
								return types.Int64Value(int64(*item.OspfSettings.Cost))
							}
							return types.Int64{}
						}(),
						IsPassiveEnabled: func() types.Bool {
							if item.OspfSettings.IsPassiveEnabled != nil {
								return types.BoolValue(*item.OspfSettings.IsPassiveEnabled)
							}
							return types.Bool{}
						}(),
					}
				}
				return nil
			}(),
			OspfV3: func() *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfV3 {
				if item.OspfV3 != nil {
					return &ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfV3{
						Area: func() types.String {
							if item.OspfV3.Area != "" {
								return types.StringValue(item.OspfV3.Area)
							}
							return types.String{}
						}(),
						Cost: func() types.Int64 {
							if item.OspfV3.Cost != nil {
								return types.Int64Value(int64(*item.OspfV3.Cost))
							}
							return types.Int64{}
						}(),
						IsPassiveEnabled: func() types.Bool {
							if item.OspfV3.IsPassiveEnabled != nil {
								return types.BoolValue(*item.OspfV3.IsPassiveEnabled)
							}
							return types.Bool{}
						}(),
					}
				}
				return nil
			}(),
			Subnet: func() types.String {
				if item.Subnet != "" {
					return types.StringValue(item.Subnet)
				}
				return types.String{}
			}(),
			UplinkV4: func() types.Bool {
				if item.UplinkV4 != nil {
					return types.BoolValue(*item.UplinkV4)
				}
				return types.Bool{}
			}(),
			UplinkV6: func() types.Bool {
				if item.UplinkV6 != nil {
					return types.BoolValue(*item.UplinkV6)
				}
				return types.Bool{}
			}(),
			VLANID: func() types.Int64 {
				if item.VLANID != nil {
					return types.Int64Value(int64(*item.VLANID))
				}
				return types.Int64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSwitchGetNetworkSwitchStackRoutingInterfaceItemToBody(state NetworksSwitchStacksRoutingInterfaces, response *merakigosdk.ResponseSwitchGetNetworkSwitchStackRoutingInterface) NetworksSwitchStacksRoutingInterfaces {
	itemState := ResponseSwitchGetNetworkSwitchStackRoutingInterface{
		DefaultGateway: func() types.String {
			if response.DefaultGateway != "" {
				return types.StringValue(response.DefaultGateway)
			}
			return types.String{}
		}(),
		InterfaceID: func() types.String {
			if response.InterfaceID != "" {
				return types.StringValue(response.InterfaceID)
			}
			return types.String{}
		}(),
		InterfaceIP: func() types.String {
			if response.InterfaceIP != "" {
				return types.StringValue(response.InterfaceIP)
			}
			return types.String{}
		}(),
		IPv6: func() *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIpv6 {
			if response.IPv6 != nil {
				return &ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIpv6{
					Address: func() types.String {
						if response.IPv6.Address != "" {
							return types.StringValue(response.IPv6.Address)
						}
						return types.String{}
					}(),
					AssignmentMode: func() types.String {
						if response.IPv6.AssignmentMode != "" {
							return types.StringValue(response.IPv6.AssignmentMode)
						}
						return types.String{}
					}(),
					Gateway: func() types.String {
						if response.IPv6.Gateway != "" {
							return types.StringValue(response.IPv6.Gateway)
						}
						return types.String{}
					}(),
					Prefix: func() types.String {
						if response.IPv6.Prefix != "" {
							return types.StringValue(response.IPv6.Prefix)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		MulticastRouting: func() types.String {
			if response.MulticastRouting != "" {
				return types.StringValue(response.MulticastRouting)
			}
			return types.String{}
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
		OspfSettings: func() *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettings {
			if response.OspfSettings != nil {
				return &ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettings{
					Area: func() types.String {
						if response.OspfSettings.Area != "" {
							return types.StringValue(response.OspfSettings.Area)
						}
						return types.String{}
					}(),
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
			return nil
		}(),
		OspfV3: func() *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3 {
			if response.OspfV3 != nil {
				return &ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3{
					Area: func() types.String {
						if response.OspfV3.Area != "" {
							return types.StringValue(response.OspfV3.Area)
						}
						return types.String{}
					}(),
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
			return nil
		}(),
		Subnet: func() types.String {
			if response.Subnet != "" {
				return types.StringValue(response.Subnet)
			}
			return types.String{}
		}(),
		UplinkV4: func() types.Bool {
			if response.UplinkV4 != nil {
				return types.BoolValue(*response.UplinkV4)
			}
			return types.Bool{}
		}(),
		UplinkV6: func() types.Bool {
			if response.UplinkV6 != nil {
				return types.BoolValue(*response.UplinkV6)
			}
			return types.Bool{}
		}(),
		VLANID: func() types.Int64 {
			if response.VLANID != nil {
				return types.Int64Value(int64(*response.VLANID))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
