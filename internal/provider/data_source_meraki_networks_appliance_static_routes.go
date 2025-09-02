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
	_ datasource.DataSource              = &NetworksApplianceStaticRoutesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceStaticRoutesDataSource{}
)

func NewNetworksApplianceStaticRoutesDataSource() datasource.DataSource {
	return &NetworksApplianceStaticRoutesDataSource{}
}

type NetworksApplianceStaticRoutesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceStaticRoutesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceStaticRoutesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_static_routes"
}

func (d *NetworksApplianceStaticRoutesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether the route is enabled or not`,
						Computed:            true,
					},
					"fixed_ip_assignments": schema.StringAttribute{
						//Entro en string ds
						//TODO interface
						MarkdownDescription: `Fixed DHCP IP assignments on the route`,
						Computed:            true,
					},
					"gateway_ip": schema.StringAttribute{
						MarkdownDescription: `Gateway IP address (next hop)`,
						Computed:            true,
					},
					"gateway_vlan_id": schema.Int64Attribute{
						MarkdownDescription: `Gateway VLAN ID`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `Route ID`,
						Computed:            true,
					},
					"ip_version": schema.Int64Attribute{
						MarkdownDescription: `IP protocol version`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the route`,
						Computed:            true,
					},
					"network_id": schema.StringAttribute{
						MarkdownDescription: `Network ID`,
						Computed:            true,
					},
					"reserved_ip_ranges": schema.SetNestedAttribute{
						MarkdownDescription: `DHCP reserved IP ranges`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"comment": schema.StringAttribute{
									MarkdownDescription: `Description of the range`,
									Computed:            true,
								},
								"end": schema.StringAttribute{
									MarkdownDescription: `Last address in the reserved range`,
									Computed:            true,
								},
								"start": schema.StringAttribute{
									MarkdownDescription: `First address in the reserved range`,
									Computed:            true,
								},
							},
						},
					},
					"subnet": schema.StringAttribute{
						MarkdownDescription: `Subnet of the route`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseApplianceGetNetworkApplianceStaticRoutes`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"enabled": schema.BoolAttribute{
							MarkdownDescription: `Whether the route is enabled or not`,
							Computed:            true,
						},
						"fixed_ip_assignments": schema.StringAttribute{
							//Entro en string ds
							//TODO interface
							MarkdownDescription: `Fixed DHCP IP assignments on the route`,
							Computed:            true,
						},
						"gateway_ip": schema.StringAttribute{
							MarkdownDescription: `Gateway IP address (next hop)`,
							Computed:            true,
						},
						"gateway_vlan_id": schema.Int64Attribute{
							MarkdownDescription: `Gateway VLAN ID`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `Route ID`,
							Computed:            true,
						},
						"ip_version": schema.Int64Attribute{
							MarkdownDescription: `IP protocol version`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the route`,
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: `Network ID`,
							Computed:            true,
						},
						"reserved_ip_ranges": schema.SetNestedAttribute{
							MarkdownDescription: `DHCP reserved IP ranges`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"comment": schema.StringAttribute{
										MarkdownDescription: `Description of the range`,
										Computed:            true,
									},
									"end": schema.StringAttribute{
										MarkdownDescription: `Last address in the reserved range`,
										Computed:            true,
									},
									"start": schema.StringAttribute{
										MarkdownDescription: `First address in the reserved range`,
										Computed:            true,
									},
								},
							},
						},
						"subnet": schema.StringAttribute{
							MarkdownDescription: `Subnet of the route`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceStaticRoutesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceStaticRoutes NetworksApplianceStaticRoutes
	diags := req.Config.Get(ctx, &networksApplianceStaticRoutes)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksApplianceStaticRoutes.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksApplianceStaticRoutes.NetworkID.IsNull(), !networksApplianceStaticRoutes.StaticRouteID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceStaticRoutes")
		vvNetworkID := networksApplianceStaticRoutes.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceStaticRoutes(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceStaticRoutes",
				err.Error(),
			)
			return
		}

		networksApplianceStaticRoutes = ResponseApplianceGetNetworkApplianceStaticRoutesItemsToBody(networksApplianceStaticRoutes, response1)
		diags = resp.State.Set(ctx, &networksApplianceStaticRoutes)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceStaticRoute")
		vvNetworkID := networksApplianceStaticRoutes.NetworkID.ValueString()
		vvStaticRouteID := networksApplianceStaticRoutes.StaticRouteID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Appliance.GetNetworkApplianceStaticRoute(vvNetworkID, vvStaticRouteID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceStaticRoute",
				err.Error(),
			)
			return
		}

		networksApplianceStaticRoutes = ResponseApplianceGetNetworkApplianceStaticRouteItemToBody(networksApplianceStaticRoutes, response2)
		diags = resp.State.Set(ctx, &networksApplianceStaticRoutes)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceStaticRoutes struct {
	NetworkID     types.String                                            `tfsdk:"network_id"`
	StaticRouteID types.String                                            `tfsdk:"static_route_id"`
	Items         *[]ResponseItemApplianceGetNetworkApplianceStaticRoutes `tfsdk:"items"`
	Item          *ResponseApplianceGetNetworkApplianceStaticRoute        `tfsdk:"item"`
}

type ResponseItemApplianceGetNetworkApplianceStaticRoutes struct {
	Enabled            types.Bool                                                              `tfsdk:"enabled"`
	FixedIPAssignments *ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	GatewayIP          types.String                                                            `tfsdk:"gateway_ip"`
	GatewayVLANID      types.Int64                                                             `tfsdk:"gateway_vlan_id"`
	ID                 types.String                                                            `tfsdk:"id"`
	IPVersion          types.Int64                                                             `tfsdk:"ip_version"`
	Name               types.String                                                            `tfsdk:"name"`
	NetworkID          types.String                                                            `tfsdk:"network_id"`
	ReservedIPRanges   *[]ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIpRanges `tfsdk:"reserved_ip_ranges"`
	Subnet             types.String                                                            `tfsdk:"subnet"`
}

type ResponseItemApplianceGetNetworkApplianceStaticRoutesFixedIpAssignments interface{}

type ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

type ResponseApplianceGetNetworkApplianceStaticRoute struct {
	Enabled            types.Bool                                                         `tfsdk:"enabled"`
	FixedIPAssignments *ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments `tfsdk:"fixed_ip_assignments"`
	GatewayIP          types.String                                                       `tfsdk:"gateway_ip"`
	GatewayVLANID      types.Int64                                                        `tfsdk:"gateway_vlan_id"`
	ID                 types.String                                                       `tfsdk:"id"`
	IPVersion          types.Int64                                                        `tfsdk:"ip_version"`
	Name               types.String                                                       `tfsdk:"name"`
	NetworkID          types.String                                                       `tfsdk:"network_id"`
	ReservedIPRanges   *[]ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRanges `tfsdk:"reserved_ip_ranges"`
	Subnet             types.String                                                       `tfsdk:"subnet"`
}

type ResponseApplianceGetNetworkApplianceStaticRouteFixedIpAssignments interface{}

type ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRanges struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceStaticRoutesItemsToBody(state NetworksApplianceStaticRoutes, response *merakigosdk.ResponseApplianceGetNetworkApplianceStaticRoutes) NetworksApplianceStaticRoutes {
	var items []ResponseItemApplianceGetNetworkApplianceStaticRoutes
	for _, item := range *response {
		itemState := ResponseItemApplianceGetNetworkApplianceStaticRoutes{
			Enabled: func() types.Bool {
				if item.Enabled != nil {
					return types.BoolValue(*item.Enabled)
				}
				return types.Bool{}
			}(),
			// FixedIPAssignments: func() types.String {
			GatewayIP: func() types.String {
				if item.GatewayIP != "" {
					return types.StringValue(item.GatewayIP)
				}
				return types.String{}
			}(),
			GatewayVLANID: func() types.Int64 {
				if item.GatewayVLANID != nil {
					return types.Int64Value(int64(*item.GatewayVLANID))
				}
				return types.Int64{}
			}(),
			ID: func() types.String {
				if item.ID != "" {
					return types.StringValue(item.ID)
				}
				return types.String{}
			}(),
			IPVersion: func() types.Int64 {
				if item.IPVersion != nil {
					return types.Int64Value(int64(*item.IPVersion))
				}
				return types.Int64{}
			}(),
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
			NetworkID: func() types.String {
				if item.NetworkID != "" {
					return types.StringValue(item.NetworkID)
				}
				return types.String{}
			}(),
			ReservedIPRanges: func() *[]ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIpRanges {
				if item.ReservedIPRanges != nil {
					result := make([]ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIpRanges, len(*item.ReservedIPRanges))
					for i, reservedIPRanges := range *item.ReservedIPRanges {
						result[i] = ResponseItemApplianceGetNetworkApplianceStaticRoutesReservedIpRanges{
							Comment: func() types.String {
								if reservedIPRanges.Comment != "" {
									return types.StringValue(reservedIPRanges.Comment)
								}
								return types.String{}
							}(),
							End: func() types.String {
								if reservedIPRanges.End != "" {
									return types.StringValue(reservedIPRanges.End)
								}
								return types.String{}
							}(),
							Start: func() types.String {
								if reservedIPRanges.Start != "" {
									return types.StringValue(reservedIPRanges.Start)
								}
								return types.String{}
							}(),
						}
					}
					return &result
				}
				return nil
			}(),
			Subnet: func() types.String {
				if item.Subnet != "" {
					return types.StringValue(item.Subnet)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseApplianceGetNetworkApplianceStaticRouteItemToBody(state NetworksApplianceStaticRoutes, response *merakigosdk.ResponseApplianceGetNetworkApplianceStaticRoute) NetworksApplianceStaticRoutes {
	itemState := ResponseApplianceGetNetworkApplianceStaticRoute{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		// FixedIPAssignments: func() types.String {
		GatewayIP: func() types.String {
			if response.GatewayIP != "" {
				return types.StringValue(response.GatewayIP)
			}
			return types.String{}
		}(),
		GatewayVLANID: func() types.Int64 {
			if response.GatewayVLANID != nil {
				return types.Int64Value(int64(*response.GatewayVLANID))
			}
			return types.Int64{}
		}(),
		ID: func() types.String {
			if response.ID != "" {
				return types.StringValue(response.ID)
			}
			return types.String{}
		}(),
		IPVersion: func() types.Int64 {
			if response.IPVersion != nil {
				return types.Int64Value(int64(*response.IPVersion))
			}
			return types.Int64{}
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
		NetworkID: func() types.String {
			if response.NetworkID != "" {
				return types.StringValue(response.NetworkID)
			}
			return types.String{}
		}(),
		ReservedIPRanges: func() *[]ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRanges {
			if response.ReservedIPRanges != nil {
				result := make([]ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRanges, len(*response.ReservedIPRanges))
				for i, reservedIPRanges := range *response.ReservedIPRanges {
					result[i] = ResponseApplianceGetNetworkApplianceStaticRouteReservedIpRanges{
						Comment: func() types.String {
							if reservedIPRanges.Comment != "" {
								return types.StringValue(reservedIPRanges.Comment)
							}
							return types.String{}
						}(),
						End: func() types.String {
							if reservedIPRanges.End != "" {
								return types.StringValue(reservedIPRanges.End)
							}
							return types.String{}
						}(),
						Start: func() types.String {
							if reservedIPRanges.Start != "" {
								return types.StringValue(reservedIPRanges.Start)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Subnet: func() types.String {
			if response.Subnet != "" {
				return types.StringValue(response.Subnet)
			}
			return types.String{}
		}(),
	}
	state.Item = &itemState
	return state
}
