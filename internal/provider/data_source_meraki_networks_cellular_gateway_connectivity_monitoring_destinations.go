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
	_ datasource.DataSource              = &NetworksCellularGatewayConnectivityMonitoringDestinationsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksCellularGatewayConnectivityMonitoringDestinationsDataSource{}
)

func NewNetworksCellularGatewayConnectivityMonitoringDestinationsDataSource() datasource.DataSource {
	return &NetworksCellularGatewayConnectivityMonitoringDestinationsDataSource{}
}

type NetworksCellularGatewayConnectivityMonitoringDestinationsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksCellularGatewayConnectivityMonitoringDestinationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksCellularGatewayConnectivityMonitoringDestinationsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_cellular_gateway_connectivity_monitoring_destinations"
}

func (d *NetworksCellularGatewayConnectivityMonitoringDestinationsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"destinations": schema.SetNestedAttribute{
						MarkdownDescription: `The list of connectivity monitoring destinations`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"default": schema.BoolAttribute{
									MarkdownDescription: `Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed`,
									Computed:            true,
								},
								"description": schema.StringAttribute{
									MarkdownDescription: `Description of the testing destination. Optional, defaults to an empty string`,
									Computed:            true,
								},
								"ip": schema.StringAttribute{
									MarkdownDescription: `The IP address to test connectivity with`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksCellularGatewayConnectivityMonitoringDestinationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksCellularGatewayConnectivityMonitoringDestinations NetworksCellularGatewayConnectivityMonitoringDestinations
	diags := req.Config.Get(ctx, &networksCellularGatewayConnectivityMonitoringDestinations)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkCellularGatewayConnectivityMonitoringDestinations")
		vvNetworkID := networksCellularGatewayConnectivityMonitoringDestinations.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.CellularGateway.GetNetworkCellularGatewayConnectivityMonitoringDestinations(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCellularGatewayConnectivityMonitoringDestinations",
				err.Error(),
			)
			return
		}

		networksCellularGatewayConnectivityMonitoringDestinations = ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinationsItemToBody(networksCellularGatewayConnectivityMonitoringDestinations, response1)
		diags = resp.State.Set(ctx, &networksCellularGatewayConnectivityMonitoringDestinations)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksCellularGatewayConnectivityMonitoringDestinations struct {
	NetworkID types.String                                                                        `tfsdk:"network_id"`
	Item      *ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinations `tfsdk:"item"`
}

type ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinations struct {
	Destinations *[]ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations `tfsdk:"destinations"`
}

type ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations struct {
	Default     types.Bool   `tfsdk:"default"`
	Description types.String `tfsdk:"description"`
	IP          types.String `tfsdk:"ip"`
}

// ToBody
func ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinationsItemToBody(state NetworksCellularGatewayConnectivityMonitoringDestinations, response *merakigosdk.ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinations) NetworksCellularGatewayConnectivityMonitoringDestinations {
	itemState := ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinations{
		Destinations: func() *[]ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations {
			if response.Destinations != nil {
				result := make([]ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations, len(*response.Destinations))
				for i, destinations := range *response.Destinations {
					result[i] = ResponseCellularGatewayGetNetworkCellularGatewayConnectivityMonitoringDestinationsDestinations{
						Default: func() types.Bool {
							if destinations.Default != nil {
								return types.BoolValue(*destinations.Default)
							}
							return types.Bool{}
						}(),
						Description: types.StringValue(destinations.Description),
						IP:          types.StringValue(destinations.IP),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
