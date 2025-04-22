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
	_ datasource.DataSource              = &NetworksSwitchLinkAggregationsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchLinkAggregationsDataSource{}
)

func NewNetworksSwitchLinkAggregationsDataSource() datasource.DataSource {
	return &NetworksSwitchLinkAggregationsDataSource{}
}

type NetworksSwitchLinkAggregationsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchLinkAggregationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchLinkAggregationsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_link_aggregations"
}

func (d *NetworksSwitchLinkAggregationsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchLinkAggregations`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `The ID for the link aggregation.`,
							Computed:            true,
						},
						"switch_ports": schema.SetNestedAttribute{
							MarkdownDescription: `The ID for the link aggregation.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"port_id": schema.StringAttribute{
										MarkdownDescription: `The ID for the switch port.`,
										Computed:            true,
									},
									"serial": schema.StringAttribute{
										MarkdownDescription: `The serial number for the switch port.`,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchLinkAggregationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchLinkAggregations NetworksSwitchLinkAggregations
	diags := req.Config.Get(ctx, &networksSwitchLinkAggregations)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchLinkAggregations")
		vvNetworkID := networksSwitchLinkAggregations.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchLinkAggregations(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchLinkAggregations",
				err.Error(),
			)
			return
		}

		networksSwitchLinkAggregations = ResponseSwitchGetNetworkSwitchLinkAggregationsItemsToBody(networksSwitchLinkAggregations, response1)
		diags = resp.State.Set(ctx, &networksSwitchLinkAggregations)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchLinkAggregations struct {
	NetworkID types.String                                          `tfsdk:"network_id"`
	Items     *[]ResponseItemSwitchGetNetworkSwitchLinkAggregations `tfsdk:"items"`
}

type ResponseItemSwitchGetNetworkSwitchLinkAggregations struct {
	ID          types.String                                                     `tfsdk:"id"`
	SwitchPorts *[]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts `tfsdk:"switch_ports"`
}

type ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts struct {
	PortID types.String `tfsdk:"port_id"`
	Serial types.String `tfsdk:"serial"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchLinkAggregationsItemsToBody(state NetworksSwitchLinkAggregations, response *merakigosdk.ResponseSwitchGetNetworkSwitchLinkAggregations) NetworksSwitchLinkAggregations {
	var items []ResponseItemSwitchGetNetworkSwitchLinkAggregations
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchLinkAggregations{
			ID: types.StringValue(item.ID),
			SwitchPorts: func() *[]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts {
				if item.SwitchPorts != nil {
					result := make([]ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts, len(*item.SwitchPorts))
					for i, switchPorts := range *item.SwitchPorts {
						result[i] = ResponseItemSwitchGetNetworkSwitchLinkAggregationsSwitchPorts{
							PortID: types.StringValue(switchPorts.PortID),
							Serial: types.StringValue(switchPorts.Serial),
						}
					}
					return &result
				}
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
