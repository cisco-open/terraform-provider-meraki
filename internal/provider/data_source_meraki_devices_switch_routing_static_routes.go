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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DevicesSwitchRoutingStaticRoutesDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesSwitchRoutingStaticRoutesDataSource{}
)

func NewDevicesSwitchRoutingStaticRoutesDataSource() datasource.DataSource {
	return &DevicesSwitchRoutingStaticRoutesDataSource{}
}

type DevicesSwitchRoutingStaticRoutesDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesSwitchRoutingStaticRoutesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesSwitchRoutingStaticRoutesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_switch_routing_static_routes"
}

func (d *DevicesSwitchRoutingStaticRoutesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Optional:            true,
			},
			"static_route_id": schema.StringAttribute{
				MarkdownDescription: `staticRouteId path parameter. Static route ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"advertise_via_ospf_enabled": schema.BoolAttribute{
						MarkdownDescription: `Option to advertise static routes via OSPF`,
						Computed:            true,
					},
					"management_next_hop": schema.StringAttribute{
						MarkdownDescription: `Optional fallback IP address for management traffic`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name or description of the layer 3 static route`,
						Computed:            true,
					},
					"next_hop_ip": schema.StringAttribute{
						MarkdownDescription: `The IP address of the router to which traffic for this destination network should be sent`,
						Computed:            true,
					},
					"prefer_over_ospf_routes_enabled": schema.BoolAttribute{
						MarkdownDescription: `Option to prefer static routes over OSPF routes`,
						Computed:            true,
					},
					"static_route_id": schema.StringAttribute{
						MarkdownDescription: `The identifier of a layer 3 static route`,
						Computed:            true,
					},
					"subnet": schema.StringAttribute{
						MarkdownDescription: `The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetDeviceSwitchRoutingStaticRoutes`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"advertise_via_ospf_enabled": schema.BoolAttribute{
							MarkdownDescription: `Option to advertise static routes via OSPF`,
							Computed:            true,
						},
						"management_next_hop": schema.StringAttribute{
							MarkdownDescription: `Optional fallback IP address for management traffic`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name or description of the layer 3 static route`,
							Computed:            true,
						},
						"next_hop_ip": schema.StringAttribute{
							MarkdownDescription: `The IP address of the router to which traffic for this destination network should be sent`,
							Computed:            true,
						},
						"prefer_over_ospf_routes_enabled": schema.BoolAttribute{
							MarkdownDescription: `Option to prefer static routes over OSPF routes`,
							Computed:            true,
						},
						"static_route_id": schema.StringAttribute{
							MarkdownDescription: `The identifier of a layer 3 static route`,
							Computed:            true,
						},
						"subnet": schema.StringAttribute{
							MarkdownDescription: `The IP address of the subnetwork specified in CIDR notation (ex. 1.2.3.0/24)`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *DevicesSwitchRoutingStaticRoutesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesSwitchRoutingStaticRoutes DevicesSwitchRoutingStaticRoutes
	diags := req.Config.Get(ctx, &devicesSwitchRoutingStaticRoutes)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!devicesSwitchRoutingStaticRoutes.Serial.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!devicesSwitchRoutingStaticRoutes.Serial.IsNull(), !devicesSwitchRoutingStaticRoutes.StaticRouteID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceSwitchRoutingStaticRoutes")
		vvSerial := devicesSwitchRoutingStaticRoutes.Serial.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetDeviceSwitchRoutingStaticRoutes(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchRoutingStaticRoutes",
				err.Error(),
			)
			return
		}

		devicesSwitchRoutingStaticRoutes = ResponseSwitchGetDeviceSwitchRoutingStaticRoutesItemsToBody(devicesSwitchRoutingStaticRoutes, response1)
		diags = resp.State.Set(ctx, &devicesSwitchRoutingStaticRoutes)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDeviceSwitchRoutingStaticRoute")
		vvSerial := devicesSwitchRoutingStaticRoutes.Serial.ValueString()
		vvStaticRouteID := devicesSwitchRoutingStaticRoutes.StaticRouteID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Switch.GetDeviceSwitchRoutingStaticRoute(vvSerial, vvStaticRouteID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSwitchRoutingStaticRoute",
				err.Error(),
			)
			return
		}

		devicesSwitchRoutingStaticRoutes = ResponseSwitchGetDeviceSwitchRoutingStaticRouteItemToBody(devicesSwitchRoutingStaticRoutes, response2)
		diags = resp.State.Set(ctx, &devicesSwitchRoutingStaticRoutes)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesSwitchRoutingStaticRoutes struct {
	Serial        types.String                                            `tfsdk:"serial"`
	StaticRouteID types.String                                            `tfsdk:"static_route_id"`
	Items         *[]ResponseItemSwitchGetDeviceSwitchRoutingStaticRoutes `tfsdk:"items"`
	Item          *ResponseSwitchGetDeviceSwitchRoutingStaticRoute        `tfsdk:"item"`
}

type ResponseItemSwitchGetDeviceSwitchRoutingStaticRoutes struct {
	AdvertiseViaOspfEnabled     types.Bool   `tfsdk:"advertise_via_ospf_enabled"`
	ManagementNextHop           types.String `tfsdk:"management_next_hop"`
	Name                        types.String `tfsdk:"name"`
	NextHopIP                   types.String `tfsdk:"next_hop_ip"`
	PreferOverOspfRoutesEnabled types.Bool   `tfsdk:"prefer_over_ospf_routes_enabled"`
	StaticRouteID               types.String `tfsdk:"static_route_id"`
	Subnet                      types.String `tfsdk:"subnet"`
}

type ResponseSwitchGetDeviceSwitchRoutingStaticRoute struct {
	AdvertiseViaOspfEnabled     types.Bool   `tfsdk:"advertise_via_ospf_enabled"`
	ManagementNextHop           types.String `tfsdk:"management_next_hop"`
	Name                        types.String `tfsdk:"name"`
	NextHopIP                   types.String `tfsdk:"next_hop_ip"`
	PreferOverOspfRoutesEnabled types.Bool   `tfsdk:"prefer_over_ospf_routes_enabled"`
	StaticRouteID               types.String `tfsdk:"static_route_id"`
	Subnet                      types.String `tfsdk:"subnet"`
}

// ToBody
func ResponseSwitchGetDeviceSwitchRoutingStaticRoutesItemsToBody(state DevicesSwitchRoutingStaticRoutes, response *merakigosdk.ResponseSwitchGetDeviceSwitchRoutingStaticRoutes) DevicesSwitchRoutingStaticRoutes {
	var items []ResponseItemSwitchGetDeviceSwitchRoutingStaticRoutes
	for _, item := range *response {
		itemState := ResponseItemSwitchGetDeviceSwitchRoutingStaticRoutes{
			AdvertiseViaOspfEnabled: func() types.Bool {
				if item.AdvertiseViaOspfEnabled != nil {
					return types.BoolValue(*item.AdvertiseViaOspfEnabled)
				}
				return types.Bool{}
			}(),
			ManagementNextHop: types.StringValue(item.ManagementNextHop),
			Name:              types.StringValue(item.Name),
			NextHopIP:         types.StringValue(item.NextHopIP),
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

func ResponseSwitchGetDeviceSwitchRoutingStaticRouteItemToBody(state DevicesSwitchRoutingStaticRoutes, response *merakigosdk.ResponseSwitchGetDeviceSwitchRoutingStaticRoute) DevicesSwitchRoutingStaticRoutes {
	itemState := ResponseSwitchGetDeviceSwitchRoutingStaticRoute{
		AdvertiseViaOspfEnabled: func() types.Bool {
			if response.AdvertiseViaOspfEnabled != nil {
				return types.BoolValue(*response.AdvertiseViaOspfEnabled)
			}
			return types.Bool{}
		}(),
		ManagementNextHop: types.StringValue(response.ManagementNextHop),
		Name:              types.StringValue(response.Name),
		NextHopIP:         types.StringValue(response.NextHopIP),
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
