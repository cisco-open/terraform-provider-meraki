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
						MarkdownDescription: `The id`,
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
								MarkdownDescription: `Area id`,
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
								MarkdownDescription: `Area id`,
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
						MarkdownDescription: `Whether this is the switch's IPv4 uplink`,
						Computed:            true,
					},
					"uplink_v6": schema.BoolAttribute{
						MarkdownDescription: `Whether this is the switch's IPv6 uplink`,
						Computed:            true,
					},
					"vlan_id": schema.Int64Attribute{
						MarkdownDescription: `VLAN id`,
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
							MarkdownDescription: `The id`,
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
									MarkdownDescription: `Area id`,
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
									MarkdownDescription: `Area id`,
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
							MarkdownDescription: `Whether this is the switch's IPv4 uplink`,
							Computed:            true,
						},
						"uplink_v6": schema.BoolAttribute{
							MarkdownDescription: `Whether this is the switch's IPv6 uplink`,
							Computed:            true,
						},
						"vlan_id": schema.Int64Attribute{
							MarkdownDescription: `VLAN id`,
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

	method1 := []bool{!networksSwitchStacksRoutingInterfaces.NetworkID.IsNull(), !networksSwitchStacksRoutingInterfaces.SwitchStackID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksSwitchStacksRoutingInterfaces.NetworkID.IsNull(), !networksSwitchStacksRoutingInterfaces.SwitchStackID.IsNull(), !networksSwitchStacksRoutingInterfaces.InterfaceID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStackRoutingInterfaces")
		vvNetworkID := networksSwitchStacksRoutingInterfaces.NetworkID.ValueString()
		vvSwitchStackID := networksSwitchStacksRoutingInterfaces.SwitchStackID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchStackRoutingInterfaces(vvNetworkID, vvSwitchStackID)

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
			DefaultGateway: types.StringValue(item.DefaultGateway),
			InterfaceID:    types.StringValue(item.InterfaceID),
			InterfaceIP:    types.StringValue(item.InterfaceIP),
			IPv6: func() *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesIpv6 {
				if item.IPv6 != nil {
					return &ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesIpv6{
						Address:        types.StringValue(item.IPv6.Address),
						AssignmentMode: types.StringValue(item.IPv6.AssignmentMode),
						Gateway:        types.StringValue(item.IPv6.Gateway),
						Prefix:         types.StringValue(item.IPv6.Prefix),
					}
				}
				return nil
			}(),
			MulticastRouting: types.StringValue(item.MulticastRouting),
			Name:             types.StringValue(item.Name),
			OspfSettings: func() *ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfSettings {
				if item.OspfSettings != nil {
					return &ResponseItemSwitchGetNetworkSwitchStackRoutingInterfacesOspfSettings{
						Area: types.StringValue(item.OspfSettings.Area),
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
						Area: types.StringValue(item.OspfV3.Area),
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
			Subnet: types.StringValue(item.Subnet),
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
		DefaultGateway: types.StringValue(response.DefaultGateway),
		InterfaceID:    types.StringValue(response.InterfaceID),
		InterfaceIP:    types.StringValue(response.InterfaceIP),
		IPv6: func() *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIpv6 {
			if response.IPv6 != nil {
				return &ResponseSwitchGetNetworkSwitchStackRoutingInterfaceIpv6{
					Address:        types.StringValue(response.IPv6.Address),
					AssignmentMode: types.StringValue(response.IPv6.AssignmentMode),
					Gateway:        types.StringValue(response.IPv6.Gateway),
					Prefix:         types.StringValue(response.IPv6.Prefix),
				}
			}
			return nil
		}(),
		MulticastRouting: types.StringValue(response.MulticastRouting),
		Name:             types.StringValue(response.Name),
		OspfSettings: func() *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettings {
			if response.OspfSettings != nil {
				return &ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfSettings{
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
			return nil
		}(),
		OspfV3: func() *ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3 {
			if response.OspfV3 != nil {
				return &ResponseSwitchGetNetworkSwitchStackRoutingInterfaceOspfV3{
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
			return nil
		}(),
		Subnet: types.StringValue(response.Subnet),
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
