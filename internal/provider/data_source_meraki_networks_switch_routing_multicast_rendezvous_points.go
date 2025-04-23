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
	_ datasource.DataSource              = &NetworksSwitchRoutingMulticastRendezvousPointsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchRoutingMulticastRendezvousPointsDataSource{}
)

func NewNetworksSwitchRoutingMulticastRendezvousPointsDataSource() datasource.DataSource {
	return &NetworksSwitchRoutingMulticastRendezvousPointsDataSource{}
}

type NetworksSwitchRoutingMulticastRendezvousPointsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchRoutingMulticastRendezvousPointsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchRoutingMulticastRendezvousPointsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_routing_multicast_rendezvous_points"
}

func (d *NetworksSwitchRoutingMulticastRendezvousPointsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"rendezvous_point_id": schema.StringAttribute{
				MarkdownDescription: `rendezvousPointId path parameter. Rendezvous point ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"interface_ip": schema.StringAttribute{
						MarkdownDescription: `The IP address of the interface to use.`,
						Computed:            true,
					},
					"interface_name": schema.StringAttribute{
						MarkdownDescription: `The name of the interface to use.`,
						Computed:            true,
					},
					"multicast_group": schema.StringAttribute{
						MarkdownDescription: `'Any', or the IP address of a multicast group.`,
						Computed:            true,
					},
					"rendezvous_point_id": schema.StringAttribute{
						MarkdownDescription: `The id.`,
						Computed:            true,
					},
					"serial": schema.StringAttribute{
						MarkdownDescription: `The serial.`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"interface_ip": schema.StringAttribute{
							MarkdownDescription: `The IP address of the interface to use.`,
							Computed:            true,
						},
						"interface_name": schema.StringAttribute{
							MarkdownDescription: `The name of the interface to use.`,
							Computed:            true,
						},
						"multicast_group": schema.StringAttribute{
							MarkdownDescription: `'Any', or the IP address of a multicast group.`,
							Computed:            true,
						},
						"rendezvous_point_id": schema.StringAttribute{
							MarkdownDescription: `The id.`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `The serial.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchRoutingMulticastRendezvousPointsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchRoutingMulticastRendezvousPoints NetworksSwitchRoutingMulticastRendezvousPoints
	diags := req.Config.Get(ctx, &networksSwitchRoutingMulticastRendezvousPoints)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksSwitchRoutingMulticastRendezvousPoints.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksSwitchRoutingMulticastRendezvousPoints.NetworkID.IsNull(), !networksSwitchRoutingMulticastRendezvousPoints.RendezvousPointID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchRoutingMulticastRendezvousPoints")
		vvNetworkID := networksSwitchRoutingMulticastRendezvousPoints.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchRoutingMulticastRendezvousPoints(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchRoutingMulticastRendezvousPoints",
				err.Error(),
			)
			return
		}

		networksSwitchRoutingMulticastRendezvousPoints = ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPointsItemsToBody(networksSwitchRoutingMulticastRendezvousPoints, response1)
		diags = resp.State.Set(ctx, &networksSwitchRoutingMulticastRendezvousPoints)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchRoutingMulticastRendezvousPoint")
		vvNetworkID := networksSwitchRoutingMulticastRendezvousPoints.NetworkID.ValueString()
		vvRendezvousPointID := networksSwitchRoutingMulticastRendezvousPoints.RendezvousPointID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Switch.GetNetworkSwitchRoutingMulticastRendezvousPoint(vvNetworkID, vvRendezvousPointID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchRoutingMulticastRendezvousPoint",
				err.Error(),
			)
			return
		}

		networksSwitchRoutingMulticastRendezvousPoints = ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPointItemToBody(networksSwitchRoutingMulticastRendezvousPoints, response2)
		diags = resp.State.Set(ctx, &networksSwitchRoutingMulticastRendezvousPoints)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchRoutingMulticastRendezvousPoints struct {
	NetworkID         types.String                                                          `tfsdk:"network_id"`
	RendezvousPointID types.String                                                          `tfsdk:"rendezvous_point_id"`
	Items             *[]ResponseItemSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints `tfsdk:"items"`
	Item              *ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoint        `tfsdk:"item"`
}

type ResponseItemSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints struct {
	InterfaceIP       types.String `tfsdk:"interface_ip"`
	InterfaceName     types.String `tfsdk:"interface_name"`
	MulticastGroup    types.String `tfsdk:"multicast_group"`
	RendezvousPointID types.String `tfsdk:"rendezvous_point_id"`
	Serial            types.String `tfsdk:"serial"`
}

type ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoint struct {
	InterfaceIP       types.String `tfsdk:"interface_ip"`
	InterfaceName     types.String `tfsdk:"interface_name"`
	MulticastGroup    types.String `tfsdk:"multicast_group"`
	RendezvousPointID types.String `tfsdk:"rendezvous_point_id"`
	Serial            types.String `tfsdk:"serial"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPointsItemsToBody(state NetworksSwitchRoutingMulticastRendezvousPoints, response *merakigosdk.ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints) NetworksSwitchRoutingMulticastRendezvousPoints {
	var items []ResponseItemSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchRoutingMulticastRendezvousPoints{
			InterfaceIP:       types.StringValue(item.InterfaceIP),
			InterfaceName:     types.StringValue(item.InterfaceName),
			MulticastGroup:    types.StringValue(item.MulticastGroup),
			RendezvousPointID: types.StringValue(item.RendezvousPointID),
			Serial:            types.StringValue(item.Serial),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPointItemToBody(state NetworksSwitchRoutingMulticastRendezvousPoints, response *merakigosdk.ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoint) NetworksSwitchRoutingMulticastRendezvousPoints {
	itemState := ResponseSwitchGetNetworkSwitchRoutingMulticastRendezvousPoint{
		InterfaceIP:       types.StringValue(response.InterfaceIP),
		InterfaceName:     types.StringValue(response.InterfaceName),
		MulticastGroup:    types.StringValue(response.MulticastGroup),
		RendezvousPointID: types.StringValue(response.RendezvousPointID),
		Serial:            types.StringValue(response.Serial),
	}
	state.Item = &itemState
	return state
}
