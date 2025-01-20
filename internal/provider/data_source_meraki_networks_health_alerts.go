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
	_ datasource.DataSource              = &NetworksHealthAlertsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksHealthAlertsDataSource{}
)

func NewNetworksHealthAlertsDataSource() datasource.DataSource {
	return &NetworksHealthAlertsDataSource{}
}

type NetworksHealthAlertsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksHealthAlertsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksHealthAlertsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_health_alerts"
}

func (d *NetworksHealthAlertsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkHealthAlerts`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"category": schema.StringAttribute{
							MarkdownDescription: `Category of the alert`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `Alert identifier. Value can be empty`,
							Computed:            true,
						},
						"scope": schema.SingleNestedAttribute{
							MarkdownDescription: `The scope of the alert`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"applications": schema.SetNestedAttribute{
									MarkdownDescription: `Applications related to the alert`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"name": schema.StringAttribute{
												MarkdownDescription: `Name of the application`,
												Computed:            true,
											},
											"url": schema.StringAttribute{
												MarkdownDescription: `URL to the application`,
												Computed:            true,
											},
										},
									},
								},
								"devices": schema.SetNestedAttribute{
									MarkdownDescription: `Devices related to the alert`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"clients": schema.SetNestedAttribute{
												MarkdownDescription: `Clients related to the device`,
												Computed:            true,
												NestedObject: schema.NestedAttributeObject{
													Attributes: map[string]schema.Attribute{

														"mac": schema.StringAttribute{
															MarkdownDescription: `Mac address of the client`,
															Computed:            true,
														},
													},
												},
											},
											"lldp": schema.SingleNestedAttribute{
												MarkdownDescription: `Lldp information`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"port_id": schema.StringAttribute{
														MarkdownDescription: `Port Id`,
														Computed:            true,
													},
												},
											},
											"mac": schema.StringAttribute{
												MarkdownDescription: `The mac address of the device`,
												Computed:            true,
											},
											"name": schema.StringAttribute{
												MarkdownDescription: `Name of the device`,
												Computed:            true,
											},
											"product_type": schema.StringAttribute{
												MarkdownDescription: `Product type of the device`,
												Computed:            true,
											},
											"serial": schema.StringAttribute{
												MarkdownDescription: `Serial number of the device`,
												Computed:            true,
											},
											"url": schema.StringAttribute{
												MarkdownDescription: `URL to the device`,
												Computed:            true,
											},
										},
									},
								},
								"peers": schema.SetNestedAttribute{
									MarkdownDescription: `Peers related to the alert`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"network": schema.SingleNestedAttribute{
												MarkdownDescription: `Network of the peer`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"id": schema.StringAttribute{
														MarkdownDescription: `Id of the network`,
														Computed:            true,
													},
													"name": schema.StringAttribute{
														MarkdownDescription: `Name of the network`,
														Computed:            true,
													},
												},
											},
											"url": schema.StringAttribute{
												MarkdownDescription: `URL to the peer`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"severity": schema.StringAttribute{
							MarkdownDescription: `Severity of the alert`,
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `Alert type`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksHealthAlertsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksHealthAlerts NetworksHealthAlerts
	diags := req.Config.Get(ctx, &networksHealthAlerts)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkHealthAlerts")
		vvNetworkID := networksHealthAlerts.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkHealthAlerts(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkHealthAlerts",
				err.Error(),
			)
			return
		}

		networksHealthAlerts = ResponseNetworksGetNetworkHealthAlertsItemsToBody(networksHealthAlerts, response1)
		diags = resp.State.Set(ctx, &networksHealthAlerts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksHealthAlerts struct {
	NetworkID types.String                                  `tfsdk:"network_id"`
	Items     *[]ResponseItemNetworksGetNetworkHealthAlerts `tfsdk:"items"`
}

type ResponseItemNetworksGetNetworkHealthAlerts struct {
	Category types.String                                     `tfsdk:"category"`
	ID       types.String                                     `tfsdk:"id"`
	Scope    *ResponseItemNetworksGetNetworkHealthAlertsScope `tfsdk:"scope"`
	Severity types.String                                     `tfsdk:"severity"`
	Type     types.String                                     `tfsdk:"type"`
}

type ResponseItemNetworksGetNetworkHealthAlertsScope struct {
	Applications *[]ResponseItemNetworksGetNetworkHealthAlertsScopeApplications `tfsdk:"applications"`
	Devices      *[]ResponseItemNetworksGetNetworkHealthAlertsScopeDevices      `tfsdk:"devices"`
	Peers        *[]ResponseItemNetworksGetNetworkHealthAlertsScopePeers        `tfsdk:"peers"`
}

type ResponseItemNetworksGetNetworkHealthAlertsScopeApplications struct {
	Name types.String `tfsdk:"name"`
	URL  types.String `tfsdk:"url"`
}

type ResponseItemNetworksGetNetworkHealthAlertsScopeDevices struct {
	Clients     *[]ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesClients `tfsdk:"clients"`
	Lldp        *ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesLldp      `tfsdk:"lldp"`
	Mac         types.String                                                     `tfsdk:"mac"`
	Name        types.String                                                     `tfsdk:"name"`
	ProductType types.String                                                     `tfsdk:"product_type"`
	Serial      types.String                                                     `tfsdk:"serial"`
	URL         types.String                                                     `tfsdk:"url"`
}

type ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesClients struct {
	Mac types.String `tfsdk:"mac"`
}

type ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesLldp struct {
	PortID types.String `tfsdk:"port_id"`
}

type ResponseItemNetworksGetNetworkHealthAlertsScopePeers struct {
	Network *ResponseItemNetworksGetNetworkHealthAlertsScopePeersNetwork `tfsdk:"network"`
	URL     types.String                                                 `tfsdk:"url"`
}

type ResponseItemNetworksGetNetworkHealthAlertsScopePeersNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseNetworksGetNetworkHealthAlertsItemsToBody(state NetworksHealthAlerts, response *merakigosdk.ResponseNetworksGetNetworkHealthAlerts) NetworksHealthAlerts {
	var items []ResponseItemNetworksGetNetworkHealthAlerts
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkHealthAlerts{
			Category: types.StringValue(item.Category),
			ID:       types.StringValue(item.ID),
			Scope: func() *ResponseItemNetworksGetNetworkHealthAlertsScope {
				if item.Scope != nil {
					return &ResponseItemNetworksGetNetworkHealthAlertsScope{
						Applications: func() *[]ResponseItemNetworksGetNetworkHealthAlertsScopeApplications {
							if item.Scope.Applications != nil {
								result := make([]ResponseItemNetworksGetNetworkHealthAlertsScopeApplications, len(*item.Scope.Applications))
								for i, applications := range *item.Scope.Applications {
									result[i] = ResponseItemNetworksGetNetworkHealthAlertsScopeApplications{
										Name: types.StringValue(applications.Name),
										URL:  types.StringValue(applications.URL),
									}
								}
								return &result
							}
							return nil
						}(),
						Devices: func() *[]ResponseItemNetworksGetNetworkHealthAlertsScopeDevices {
							if item.Scope.Devices != nil {
								result := make([]ResponseItemNetworksGetNetworkHealthAlertsScopeDevices, len(*item.Scope.Devices))
								for i, devices := range *item.Scope.Devices {
									result[i] = ResponseItemNetworksGetNetworkHealthAlertsScopeDevices{
										Clients: func() *[]ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesClients {
											if devices.Clients != nil {
												result := make([]ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesClients, len(*devices.Clients))
												for i, clients := range *devices.Clients {
													result[i] = ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesClients{
														Mac: types.StringValue(clients.Mac),
													}
												}
												return &result
											}
											return nil
										}(),
										Lldp: func() *ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesLldp {
											if devices.Lldp != nil {
												return &ResponseItemNetworksGetNetworkHealthAlertsScopeDevicesLldp{
													PortID: types.StringValue(devices.Lldp.PortID),
												}
											}
											return nil
										}(),
										Mac:         types.StringValue(devices.Mac),
										Name:        types.StringValue(devices.Name),
										ProductType: types.StringValue(devices.ProductType),
										Serial:      types.StringValue(devices.Serial),
										URL:         types.StringValue(devices.URL),
									}
								}
								return &result
							}
							return nil
						}(),
						Peers: func() *[]ResponseItemNetworksGetNetworkHealthAlertsScopePeers {
							if item.Scope.Peers != nil {
								result := make([]ResponseItemNetworksGetNetworkHealthAlertsScopePeers, len(*item.Scope.Peers))
								for i, peers := range *item.Scope.Peers {
									result[i] = ResponseItemNetworksGetNetworkHealthAlertsScopePeers{
										Network: func() *ResponseItemNetworksGetNetworkHealthAlertsScopePeersNetwork {
											if peers.Network != nil {
												return &ResponseItemNetworksGetNetworkHealthAlertsScopePeersNetwork{
													ID:   types.StringValue(peers.Network.ID),
													Name: types.StringValue(peers.Network.Name),
												}
											}
											return nil
										}(),
										URL: types.StringValue(peers.URL),
									}
								}
								return &result
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
			Severity: types.StringValue(item.Severity),
			Type:     types.StringValue(item.Type),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
