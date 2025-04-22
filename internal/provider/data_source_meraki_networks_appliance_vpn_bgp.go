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
	_ datasource.DataSource              = &NetworksApplianceVpnBgpDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceVpnBgpDataSource{}
)

func NewNetworksApplianceVpnBgpDataSource() datasource.DataSource {
	return &NetworksApplianceVpnBgpDataSource{}
}

type NetworksApplianceVpnBgpDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceVpnBgpDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceVpnBgpDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_vpn_bgp"
}

func (d *NetworksApplianceVpnBgpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"as_number": schema.Int64Attribute{
						MarkdownDescription: `The number of the Autonomous System to which the appliance belongs`,
						Computed:            true,
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether BGP is enabled on the appliance`,
						Computed:            true,
					},
					"ibgp_hold_timer": schema.Int64Attribute{
						MarkdownDescription: `The iBGP hold time in seconds`,
						Computed:            true,
					},
					"neighbors": schema.SetNestedAttribute{
						MarkdownDescription: `List of eBGP neighbor configurations`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"allow_transit": schema.BoolAttribute{
									MarkdownDescription: `Whether the appliance will advertise routes learned from other Autonomous Systems`,
									Computed:            true,
								},
								"authentication": schema.SingleNestedAttribute{
									MarkdownDescription: `Authentication settings between BGP peers`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"password": schema.StringAttribute{
											MarkdownDescription: `Password to configure MD5 authentication between BGP peers`,
											Sensitive:           true,
											Computed:            true,
										},
									},
								},
								"ebgp_hold_timer": schema.Int64Attribute{
									MarkdownDescription: `The eBGP hold time in seconds for the neighbor`,
									Computed:            true,
								},
								"ebgp_multihop": schema.Int64Attribute{
									MarkdownDescription: `The number of hops the appliance must traverse to establish a peering relationship with the neighbor`,
									Computed:            true,
								},
								"ip": schema.StringAttribute{
									MarkdownDescription: `The IPv4 address of the neighbor`,
									Computed:            true,
								},
								"ipv6": schema.SingleNestedAttribute{
									MarkdownDescription: `Information regarding IPv6 address of the neighbor`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"address": schema.StringAttribute{
											MarkdownDescription: `The IPv6 address of the neighbor`,
											Computed:            true,
										},
									},
								},
								"next_hop_ip": schema.StringAttribute{
									MarkdownDescription: `The IPv4 address of the neighbor that will establish a TCP session with the appliance`,
									Computed:            true,
								},
								"receive_limit": schema.Int64Attribute{
									MarkdownDescription: `The maximum number of routes that the appliance can receive from the neighbor`,
									Computed:            true,
								},
								"remote_as_number": schema.Int64Attribute{
									MarkdownDescription: `Remote AS number of the neighbor`,
									Computed:            true,
								},
								"source_interface": schema.StringAttribute{
									MarkdownDescription: `The output interface the appliance uses to establish a peering relationship with the neighbor`,
									Computed:            true,
								},
								"ttl_security": schema.SingleNestedAttribute{
									MarkdownDescription: `Settings for BGP TTL security to protect BGP peering sessions from forged IP attacks`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"enabled": schema.BoolAttribute{
											MarkdownDescription: `Whether BGP TTL security is enabled`,
											Computed:            true,
										},
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

func (d *NetworksApplianceVpnBgpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceVpnBgp NetworksApplianceVpnBgp
	diags := req.Config.Get(ctx, &networksApplianceVpnBgp)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceVpnBgp")
		vvNetworkID := networksApplianceVpnBgp.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceVpnBgp(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceVpnBgp",
				err.Error(),
			)
			return
		}

		networksApplianceVpnBgp = ResponseApplianceGetNetworkApplianceVpnBgpItemToBody(networksApplianceVpnBgp, response1)
		diags = resp.State.Set(ctx, &networksApplianceVpnBgp)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceVpnBgp struct {
	NetworkID types.String                                `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceVpnBgp `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceVpnBgp struct {
	AsNumber      types.Int64                                            `tfsdk:"as_number"`
	Enabled       types.Bool                                             `tfsdk:"enabled"`
	IbgpHoldTimer types.Int64                                            `tfsdk:"ibgp_hold_timer"`
	Neighbors     *[]ResponseApplianceGetNetworkApplianceVpnBgpNeighbors `tfsdk:"neighbors"`
}

type ResponseApplianceGetNetworkApplianceVpnBgpNeighbors struct {
	AllowTransit    types.Bool                                                         `tfsdk:"allow_transit"`
	Authentication  *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsAuthentication `tfsdk:"authentication"`
	EbgpHoldTimer   types.Int64                                                        `tfsdk:"ebgp_hold_timer"`
	EbgpMultihop    types.Int64                                                        `tfsdk:"ebgp_multihop"`
	IP              types.String                                                       `tfsdk:"ip"`
	IPv6            *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsIpv6           `tfsdk:"ipv6"`
	NextHopIP       types.String                                                       `tfsdk:"next_hop_ip"`
	ReceiveLimit    types.Int64                                                        `tfsdk:"receive_limit"`
	RemoteAsNumber  types.Int64                                                        `tfsdk:"remote_as_number"`
	SourceInterface types.String                                                       `tfsdk:"source_interface"`
	TtlSecurity     *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsTtlSecurity    `tfsdk:"ttl_security"`
}

type ResponseApplianceGetNetworkApplianceVpnBgpNeighborsAuthentication struct {
	Password types.String `tfsdk:"password"`
}

type ResponseApplianceGetNetworkApplianceVpnBgpNeighborsIpv6 struct {
	Address types.String `tfsdk:"address"`
}

type ResponseApplianceGetNetworkApplianceVpnBgpNeighborsTtlSecurity struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceVpnBgpItemToBody(state NetworksApplianceVpnBgp, response *merakigosdk.ResponseApplianceGetNetworkApplianceVpnBgp) NetworksApplianceVpnBgp {
	itemState := ResponseApplianceGetNetworkApplianceVpnBgp{
		AsNumber: func() types.Int64 {
			if response.AsNumber != nil {
				return types.Int64Value(int64(*response.AsNumber))
			}
			return types.Int64{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		IbgpHoldTimer: func() types.Int64 {
			if response.IbgpHoldTimer != nil {
				return types.Int64Value(int64(*response.IbgpHoldTimer))
			}
			return types.Int64{}
		}(),
		Neighbors: func() *[]ResponseApplianceGetNetworkApplianceVpnBgpNeighbors {
			if response.Neighbors != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVpnBgpNeighbors, len(*response.Neighbors))
				for i, neighbors := range *response.Neighbors {
					result[i] = ResponseApplianceGetNetworkApplianceVpnBgpNeighbors{
						AllowTransit: func() types.Bool {
							if neighbors.AllowTransit != nil {
								return types.BoolValue(*neighbors.AllowTransit)
							}
							return types.Bool{}
						}(),
						Authentication: func() *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsAuthentication {
							if neighbors.Authentication != nil {
								return &ResponseApplianceGetNetworkApplianceVpnBgpNeighborsAuthentication{
									Password: types.StringValue(neighbors.Authentication.Password),
								}
							}
							return nil
						}(),
						EbgpHoldTimer: func() types.Int64 {
							if neighbors.EbgpHoldTimer != nil {
								return types.Int64Value(int64(*neighbors.EbgpHoldTimer))
							}
							return types.Int64{}
						}(),
						EbgpMultihop: func() types.Int64 {
							if neighbors.EbgpMultihop != nil {
								return types.Int64Value(int64(*neighbors.EbgpMultihop))
							}
							return types.Int64{}
						}(),
						IP: types.StringValue(neighbors.IP),
						IPv6: func() *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsIpv6 {
							if neighbors.IPv6 != nil {
								return &ResponseApplianceGetNetworkApplianceVpnBgpNeighborsIpv6{
									Address: types.StringValue(neighbors.IPv6.Address),
								}
							}
							return nil
						}(),
						NextHopIP: types.StringValue(neighbors.NextHopIP),
						ReceiveLimit: func() types.Int64 {
							if neighbors.ReceiveLimit != nil {
								return types.Int64Value(int64(*neighbors.ReceiveLimit))
							}
							return types.Int64{}
						}(),
						RemoteAsNumber: func() types.Int64 {
							if neighbors.RemoteAsNumber != nil {
								return types.Int64Value(int64(*neighbors.RemoteAsNumber))
							}
							return types.Int64{}
						}(),
						SourceInterface: types.StringValue(neighbors.SourceInterface),
						TtlSecurity: func() *ResponseApplianceGetNetworkApplianceVpnBgpNeighborsTtlSecurity {
							if neighbors.TtlSecurity != nil {
								return &ResponseApplianceGetNetworkApplianceVpnBgpNeighborsTtlSecurity{
									Enabled: func() types.Bool {
										if neighbors.TtlSecurity.Enabled != nil {
											return types.BoolValue(*neighbors.TtlSecurity.Enabled)
										}
										return types.Bool{}
									}(),
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
	}
	state.Item = &itemState
	return state
}
