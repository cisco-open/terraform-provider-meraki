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
	_ datasource.DataSource              = &NetworksTopologyLinkLayerDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksTopologyLinkLayerDataSource{}
)

func NewNetworksTopologyLinkLayerDataSource() datasource.DataSource {
	return &NetworksTopologyLinkLayerDataSource{}
}

type NetworksTopologyLinkLayerDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksTopologyLinkLayerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksTopologyLinkLayerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_topology_link_layer"
}

func (d *NetworksTopologyLinkLayerDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"errors": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"links": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"ends": schema.SetNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"device": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{

													"name": schema.StringAttribute{
														Computed: true,
													},
													"serial": schema.StringAttribute{
														Computed: true,
													},
												},
											},
											"discovered": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{

													"cdp": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{

															"native_vlan": schema.Int64Attribute{
																Computed: true,
															},
															"port_id": schema.StringAttribute{
																Computed: true,
															},
														},
													},
													"lldp": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{

															"port_description": schema.StringAttribute{
																Computed: true,
															},
															"port_id": schema.StringAttribute{
																Computed: true,
															},
														},
													},
												},
											},
											"node": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{

													"derived_id": schema.StringAttribute{
														Computed: true,
													},
													"type": schema.StringAttribute{
														Computed: true,
													},
												},
											},
										},
									},
								},
								"last_reported_at": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"nodes": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"derived_id": schema.StringAttribute{
									Computed: true,
								},
								"discovered": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"cdp": schema.StringAttribute{
											Computed: true,
										},
										"lldp": schema.SingleNestedAttribute{
											Computed: true,
											Attributes: map[string]schema.Attribute{

												"chassis_id": schema.StringAttribute{
													Computed: true,
												},
												"management_address": schema.StringAttribute{
													Computed: true,
												},
												"system_capabilities": schema.ListAttribute{
													Computed:    true,
													ElementType: types.StringType,
												},
												"system_description": schema.StringAttribute{
													Computed: true,
												},
												"system_name": schema.StringAttribute{
													Computed: true,
												},
											},
										},
									},
								},
								"mac": schema.StringAttribute{
									Computed: true,
								},
								"root": schema.BoolAttribute{
									Computed: true,
								},
								"type": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksTopologyLinkLayerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksTopologyLinkLayer NetworksTopologyLinkLayer
	diags := req.Config.Get(ctx, &networksTopologyLinkLayer)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkTopologyLinkLayer")
		vvNetworkID := networksTopologyLinkLayer.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkTopologyLinkLayer(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkTopologyLinkLayer",
				err.Error(),
			)
			return
		}

		networksTopologyLinkLayer = ResponseNetworksGetNetworkTopologyLinkLayerItemToBody(networksTopologyLinkLayer, response1)
		diags = resp.State.Set(ctx, &networksTopologyLinkLayer)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksTopologyLinkLayer struct {
	NetworkID types.String                                 `tfsdk:"network_id"`
	Item      *ResponseNetworksGetNetworkTopologyLinkLayer `tfsdk:"item"`
}

type ResponseNetworksGetNetworkTopologyLinkLayer struct {
	Errors types.List                                          `tfsdk:"errors"`
	Links  *[]ResponseNetworksGetNetworkTopologyLinkLayerLinks `tfsdk:"links"`
	Nodes  *[]ResponseNetworksGetNetworkTopologyLinkLayerNodes `tfsdk:"nodes"`
}

type ResponseNetworksGetNetworkTopologyLinkLayerLinks struct {
	Ends           *[]ResponseNetworksGetNetworkTopologyLinkLayerLinksEnds `tfsdk:"ends"`
	LastReportedAt types.String                                            `tfsdk:"last_reported_at"`
}

type ResponseNetworksGetNetworkTopologyLinkLayerLinksEnds struct {
	Device     *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDevice     `tfsdk:"device"`
	Discovered *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscovered `tfsdk:"discovered"`
	Node       *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsNode       `tfsdk:"node"`
}

type ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDevice struct {
	Name   types.String `tfsdk:"name"`
	Serial types.String `tfsdk:"serial"`
}

type ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscovered struct {
	Cdp  *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscoveredCdp  `tfsdk:"cdp"`
	Lldp *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscoveredLldp `tfsdk:"lldp"`
}

type ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscoveredCdp struct {
	NativeVLAN types.Int64  `tfsdk:"native_vlan"`
	PortID     types.String `tfsdk:"port_id"`
}

type ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscoveredLldp struct {
	PortDescription types.String `tfsdk:"port_description"`
	PortID          types.String `tfsdk:"port_id"`
}

type ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsNode struct {
	DerivedID types.String `tfsdk:"derived_id"`
	Type      types.String `tfsdk:"type"`
}

type ResponseNetworksGetNetworkTopologyLinkLayerNodes struct {
	DerivedID  types.String                                                `tfsdk:"derived_id"`
	Discovered *ResponseNetworksGetNetworkTopologyLinkLayerNodesDiscovered `tfsdk:"discovered"`
	Mac        types.String                                                `tfsdk:"mac"`
	Root       types.Bool                                                  `tfsdk:"root"`
	Type       types.String                                                `tfsdk:"type"`
}

type ResponseNetworksGetNetworkTopologyLinkLayerNodesDiscovered struct {
	Cdp  types.String                                                    `tfsdk:"cdp"`
	Lldp *ResponseNetworksGetNetworkTopologyLinkLayerNodesDiscoveredLldp `tfsdk:"lldp"`
}

type ResponseNetworksGetNetworkTopologyLinkLayerNodesDiscoveredLldp struct {
	ChassisID          types.String `tfsdk:"chassis_id"`
	ManagementAddress  types.String `tfsdk:"management_address"`
	SystemCapabilities types.List   `tfsdk:"system_capabilities"`
	SystemDescription  types.String `tfsdk:"system_description"`
	SystemName         types.String `tfsdk:"system_name"`
}

// ToBody
func ResponseNetworksGetNetworkTopologyLinkLayerItemToBody(state NetworksTopologyLinkLayer, response *merakigosdk.ResponseNetworksGetNetworkTopologyLinkLayer) NetworksTopologyLinkLayer {
	itemState := ResponseNetworksGetNetworkTopologyLinkLayer{
		Errors: StringSliceToList(response.Errors),
		Links: func() *[]ResponseNetworksGetNetworkTopologyLinkLayerLinks {
			if response.Links != nil {
				result := make([]ResponseNetworksGetNetworkTopologyLinkLayerLinks, len(*response.Links))
				for i, links := range *response.Links {
					result[i] = ResponseNetworksGetNetworkTopologyLinkLayerLinks{
						Ends: func() *[]ResponseNetworksGetNetworkTopologyLinkLayerLinksEnds {
							if links.Ends != nil {
								result := make([]ResponseNetworksGetNetworkTopologyLinkLayerLinksEnds, len(*links.Ends))
								for i, ends := range *links.Ends {
									result[i] = ResponseNetworksGetNetworkTopologyLinkLayerLinksEnds{
										Device: func() *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDevice {
											if ends.Device != nil {
												return &ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDevice{
													Name:   types.StringValue(ends.Device.Name),
													Serial: types.StringValue(ends.Device.Serial),
												}
											}
											return nil
										}(),
										Discovered: func() *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscovered {
											if ends.Discovered != nil {
												return &ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscovered{
													Cdp: func() *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscoveredCdp {
														if ends.Discovered.Cdp != nil {
															return &ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscoveredCdp{
																NativeVLAN: func() types.Int64 {
																	if ends.Discovered.Cdp.NativeVLAN != nil {
																		return types.Int64Value(int64(*ends.Discovered.Cdp.NativeVLAN))
																	}
																	return types.Int64{}
																}(),
																PortID: types.StringValue(ends.Discovered.Cdp.PortID),
															}
														}
														return nil
													}(),
													Lldp: func() *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscoveredLldp {
														if ends.Discovered.Lldp != nil {
															return &ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsDiscoveredLldp{
																PortDescription: types.StringValue(ends.Discovered.Lldp.PortDescription),
																PortID:          types.StringValue(ends.Discovered.Lldp.PortID),
															}
														}
														return nil
													}(),
												}
											}
											return nil
										}(),
										Node: func() *ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsNode {
											if ends.Node != nil {
												return &ResponseNetworksGetNetworkTopologyLinkLayerLinksEndsNode{
													DerivedID: types.StringValue(ends.Node.DerivedID),
													Type:      types.StringValue(ends.Node.Type),
												}
											}
											return nil
										}(),
									}
								}
								return &result
							}
							return nil
						}(),
						LastReportedAt: types.StringValue(links.LastReportedAt),
					}
				}
				return &result
			}
			return nil
		}(),
		Nodes: func() *[]ResponseNetworksGetNetworkTopologyLinkLayerNodes {
			if response.Nodes != nil {
				result := make([]ResponseNetworksGetNetworkTopologyLinkLayerNodes, len(*response.Nodes))
				for i, nodes := range *response.Nodes {
					result[i] = ResponseNetworksGetNetworkTopologyLinkLayerNodes{
						DerivedID: types.StringValue(nodes.DerivedID),
						Discovered: func() *ResponseNetworksGetNetworkTopologyLinkLayerNodesDiscovered {
							if nodes.Discovered != nil {
								return &ResponseNetworksGetNetworkTopologyLinkLayerNodesDiscovered{
									Cdp: types.StringValue(nodes.Discovered.Cdp),
									Lldp: func() *ResponseNetworksGetNetworkTopologyLinkLayerNodesDiscoveredLldp {
										if nodes.Discovered.Lldp != nil {
											return &ResponseNetworksGetNetworkTopologyLinkLayerNodesDiscoveredLldp{
												ChassisID:          types.StringValue(nodes.Discovered.Lldp.ChassisID),
												ManagementAddress:  types.StringValue(nodes.Discovered.Lldp.ManagementAddress),
												SystemCapabilities: StringSliceToList(nodes.Discovered.Lldp.SystemCapabilities),
												SystemDescription:  types.StringValue(nodes.Discovered.Lldp.SystemDescription),
												SystemName:         types.StringValue(nodes.Discovered.Lldp.SystemName),
											}
										}
										return nil
									}(),
								}
							}
							return nil
						}(),
						Mac: types.StringValue(nodes.Mac),
						Root: func() types.Bool {
							if nodes.Root != nil {
								return types.BoolValue(*nodes.Root)
							}
							return types.Bool{}
						}(),
						Type: types.StringValue(nodes.Type),
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
