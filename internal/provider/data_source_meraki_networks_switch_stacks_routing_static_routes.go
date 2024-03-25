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
	_ datasource.DataSource              = &NetworksSwitchStacksRoutingStaticRoutesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchStacksRoutingStaticRoutesDataSource{}
)

func NewNetworksSwitchStacksRoutingStaticRoutesDataSource() datasource.DataSource {
	return &NetworksSwitchStacksRoutingStaticRoutesDataSource{}
}

type NetworksSwitchStacksRoutingStaticRoutesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchStacksRoutingStaticRoutesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchStacksRoutingStaticRoutesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stacks_routing_static_routes"
}

func (d *NetworksSwitchStacksRoutingStaticRoutesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"static_route_id": schema.StringAttribute{
				MarkdownDescription: `staticRouteId path parameter. Static route ID`,
				Optional:            true,
			},
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: `switchStackId path parameter. Switch stack ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"advertise_via_ospf_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"next_hop_ip": schema.StringAttribute{
						Computed: true,
					},
					"prefer_over_ospf_routes_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"static_route_id": schema.StringAttribute{
						Computed: true,
					},
					"subnet": schema.StringAttribute{
						Computed: true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchStackRoutingStaticRoutes`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"advertise_via_ospf_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"next_hop_ip": schema.StringAttribute{
							Computed: true,
						},
						"prefer_over_ospf_routes_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"static_route_id": schema.StringAttribute{
							Computed: true,
						},
						"subnet": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchStacksRoutingStaticRoutesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchStacksRoutingStaticRoutes NetworksSwitchStacksRoutingStaticRoutes
	diags := req.Config.Get(ctx, &networksSwitchStacksRoutingStaticRoutes)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksSwitchStacksRoutingStaticRoutes.NetworkID.IsNull(), !networksSwitchStacksRoutingStaticRoutes.SwitchStackID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksSwitchStacksRoutingStaticRoutes.NetworkID.IsNull(), !networksSwitchStacksRoutingStaticRoutes.SwitchStackID.IsNull(), !networksSwitchStacksRoutingStaticRoutes.StaticRouteID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStackRoutingStaticRoutes")
		vvNetworkID := networksSwitchStacksRoutingStaticRoutes.NetworkID.ValueString()
		vvSwitchStackID := networksSwitchStacksRoutingStaticRoutes.SwitchStackID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchStackRoutingStaticRoutes(vvNetworkID, vvSwitchStackID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStackRoutingStaticRoutes",
				err.Error(),
			)
			return
		}

		networksSwitchStacksRoutingStaticRoutes = ResponseSwitchGetNetworkSwitchStackRoutingStaticRoutesItemsToBody(networksSwitchStacksRoutingStaticRoutes, response1)
		diags = resp.State.Set(ctx, &networksSwitchStacksRoutingStaticRoutes)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStackRoutingStaticRoute")
		vvNetworkID := networksSwitchStacksRoutingStaticRoutes.NetworkID.ValueString()
		vvSwitchStackID := networksSwitchStacksRoutingStaticRoutes.SwitchStackID.ValueString()
		vvStaticRouteID := networksSwitchStacksRoutingStaticRoutes.StaticRouteID.ValueString()

		response2, restyResp2, err := d.client.Switch.GetNetworkSwitchStackRoutingStaticRoute(vvNetworkID, vvSwitchStackID, vvStaticRouteID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStackRoutingStaticRoute",
				err.Error(),
			)
			return
		}

		networksSwitchStacksRoutingStaticRoutes = ResponseSwitchGetNetworkSwitchStackRoutingStaticRouteItemToBody(networksSwitchStacksRoutingStaticRoutes, response2)
		diags = resp.State.Set(ctx, &networksSwitchStacksRoutingStaticRoutes)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchStacksRoutingStaticRoutes struct {
	NetworkID     types.String                                                  `tfsdk:"network_id"`
	SwitchStackID types.String                                                  `tfsdk:"switch_stack_id"`
	StaticRouteID types.String                                                  `tfsdk:"static_route_id"`
	Items         *[]ResponseItemSwitchGetNetworkSwitchStackRoutingStaticRoutes `tfsdk:"items"`
	Item          *ResponseSwitchGetNetworkSwitchStackRoutingStaticRoute        `tfsdk:"item"`
}

type ResponseItemSwitchGetNetworkSwitchStackRoutingStaticRoutes struct {
	AdvertiseViaOspfEnabled     types.Bool   `tfsdk:"advertise_via_ospf_enabled"`
	Name                        types.String `tfsdk:"name"`
	NextHopIP                   types.String `tfsdk:"next_hop_ip"`
	PreferOverOspfRoutesEnabled types.Bool   `tfsdk:"prefer_over_ospf_routes_enabled"`
	StaticRouteID               types.String `tfsdk:"static_route_id"`
	Subnet                      types.String `tfsdk:"subnet"`
}

type ResponseSwitchGetNetworkSwitchStackRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     types.Bool   `tfsdk:"advertise_via_ospf_enabled"`
	Name                        types.String `tfsdk:"name"`
	NextHopIP                   types.String `tfsdk:"next_hop_ip"`
	PreferOverOspfRoutesEnabled types.Bool   `tfsdk:"prefer_over_ospf_routes_enabled"`
	StaticRouteID               types.String `tfsdk:"static_route_id"`
	Subnet                      types.String `tfsdk:"subnet"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchStackRoutingStaticRoutesItemsToBody(state NetworksSwitchStacksRoutingStaticRoutes, response *merakigosdk.ResponseSwitchGetNetworkSwitchStackRoutingStaticRoutes) NetworksSwitchStacksRoutingStaticRoutes {
	var items []ResponseItemSwitchGetNetworkSwitchStackRoutingStaticRoutes
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchStackRoutingStaticRoutes{
			AdvertiseViaOspfEnabled: func() types.Bool {
				if item.AdvertiseViaOspfEnabled != nil {
					return types.BoolValue(*item.AdvertiseViaOspfEnabled)
				}
				return types.Bool{}
			}(),
			Name:      types.StringValue(item.Name),
			NextHopIP: types.StringValue(item.NextHopIP),
			PreferOverOspfRoutesEnabled: func() types.Bool {
				if item.PreferOverOspfRoutesEnabled != nil {
					return types.BoolValue(*item.PreferOverOspfRoutesEnabled)
				}
				return types.Bool{}
			}(),
			StaticRouteID: types.StringValue(item.StaticRouteID),
			Subnet:        types.StringValue(item.Subnet),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSwitchGetNetworkSwitchStackRoutingStaticRouteItemToBody(state NetworksSwitchStacksRoutingStaticRoutes, response *merakigosdk.ResponseSwitchGetNetworkSwitchStackRoutingStaticRoute) NetworksSwitchStacksRoutingStaticRoutes {
	itemState := ResponseSwitchGetNetworkSwitchStackRoutingStaticRoute{
		AdvertiseViaOspfEnabled: func() types.Bool {
			if response.AdvertiseViaOspfEnabled != nil {
				return types.BoolValue(*response.AdvertiseViaOspfEnabled)
			}
			return types.Bool{}
		}(),
		Name:      types.StringValue(response.Name),
		NextHopIP: types.StringValue(response.NextHopIP),
		PreferOverOspfRoutesEnabled: func() types.Bool {
			if response.PreferOverOspfRoutesEnabled != nil {
				return types.BoolValue(*response.PreferOverOspfRoutesEnabled)
			}
			return types.Bool{}
		}(),
		StaticRouteID: types.StringValue(response.StaticRouteID),
		Subnet:        types.StringValue(response.Subnet),
	}
	state.Item = &itemState
	return state
}
