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
	_ datasource.DataSource              = &NetworksSwitchDhcpV4ServersSeenDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchDhcpV4ServersSeenDataSource{}
)

func NewNetworksSwitchDhcpV4ServersSeenDataSource() datasource.DataSource {
	return &NetworksSwitchDhcpV4ServersSeenDataSource{}
}

type NetworksSwitchDhcpV4ServersSeenDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchDhcpV4ServersSeenDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchDhcpV4ServersSeenDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_dhcp_v4_servers_seen"
}

func (d *NetworksSwitchDhcpV4ServersSeenDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 31 days from today.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameter t0. The value must be in seconds and be less than or equal to 31 days. The default is 1 day.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetNetworkSwitchDhcpV4ServersSeen`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"client_id": schema.StringAttribute{
							MarkdownDescription: `Client id of the server if available.`,
							Computed:            true,
						},
						"device": schema.SingleNestedAttribute{
							MarkdownDescription: `Attributes of the server when it's a device.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"interface": schema.SingleNestedAttribute{
									MarkdownDescription: `Interface attributes of the server. Only for configured servers.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"name": schema.StringAttribute{
											MarkdownDescription: `Interface name.`,
											Computed:            true,
										},
										"url": schema.StringAttribute{
											MarkdownDescription: `Url link to interface.`,
											Computed:            true,
										},
									},
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Device name.`,
									Computed:            true,
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Device serial.`,
									Computed:            true,
								},
								"url": schema.StringAttribute{
									MarkdownDescription: `Url link to device.`,
									Computed:            true,
								},
							},
						},
						"ipv4": schema.SingleNestedAttribute{
							MarkdownDescription: `IPv4 attributes of the server.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"address": schema.StringAttribute{
									MarkdownDescription: `IPv4 address of the server.`,
									Computed:            true,
								},
								"gateway": schema.StringAttribute{
									MarkdownDescription: `IPv4 gateway address of the server.`,
									Computed:            true,
								},
								"subnet": schema.StringAttribute{
									MarkdownDescription: `Subnet of the server.`,
									Computed:            true,
								},
							},
						},
						"is_allowed": schema.BoolAttribute{
							MarkdownDescription: `Whether the server is allowed or blocked. Always true for configured servers.`,
							Computed:            true,
						},
						"is_configured": schema.BoolAttribute{
							MarkdownDescription: `Whether the server is configured.`,
							Computed:            true,
						},
						"last_ack": schema.SingleNestedAttribute{
							MarkdownDescription: `Attributes of the server's last ack.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"ipv4": schema.SingleNestedAttribute{
									MarkdownDescription: `IPv4 attributes of the last ack.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"address": schema.StringAttribute{
											MarkdownDescription: `IPv4 address of the last ack.`,
											Computed:            true,
										},
									},
								},
								"ts": schema.StringAttribute{
									MarkdownDescription: `Last time the server was acked.`,
									Computed:            true,
								},
							},
						},
						"last_packet": schema.SingleNestedAttribute{
							MarkdownDescription: `Last packet the server received.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"destination": schema.SingleNestedAttribute{
									MarkdownDescription: `Destination of the packet.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"ipv4": schema.SingleNestedAttribute{
											MarkdownDescription: `Destination ipv4 attributes of the packet.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"address": schema.StringAttribute{
													MarkdownDescription: `Destination ipv4 address of the packet.`,
													Computed:            true,
												},
											},
										},
										"mac": schema.StringAttribute{
											MarkdownDescription: `Destination mac address of the packet.`,
											Computed:            true,
										},
										"port": schema.Int64Attribute{
											MarkdownDescription: `Destination port of the packet.`,
											Computed:            true,
										},
									},
								},
								"ethernet": schema.SingleNestedAttribute{
									MarkdownDescription: `Additional ethernet attributes of the packet.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"type": schema.StringAttribute{
											MarkdownDescription: `Ethernet type of the packet.`,
											Computed:            true,
										},
									},
								},
								"fields": schema.SingleNestedAttribute{
									MarkdownDescription: `DHCP-specific fields of the packet.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"chaddr": schema.StringAttribute{
											MarkdownDescription: `Client hardware address of the packet.`,
											Computed:            true,
										},
										"ciaddr": schema.StringAttribute{
											MarkdownDescription: `Client IP address of the packet.`,
											Computed:            true,
										},
										"flags": schema.StringAttribute{
											MarkdownDescription: `Packet flags.`,
											Computed:            true,
										},
										"giaddr": schema.StringAttribute{
											MarkdownDescription: `Gateway IP address of the packet.`,
											Computed:            true,
										},
										"hlen": schema.Int64Attribute{
											MarkdownDescription: `Hardware length of the packet.`,
											Computed:            true,
										},
										"hops": schema.Int64Attribute{
											MarkdownDescription: `Number of hops the packet took.`,
											Computed:            true,
										},
										"htype": schema.Int64Attribute{
											MarkdownDescription: `Hardware type code of the packet.`,
											Computed:            true,
										},
										"magic_cookie": schema.StringAttribute{
											MarkdownDescription: `Magic cookie of the packet.`,
											Computed:            true,
										},
										"op": schema.Int64Attribute{
											MarkdownDescription: `Operation code of the packet.`,
											Computed:            true,
										},
										"options": schema.SetNestedAttribute{
											MarkdownDescription: `Additional DHCP options of the packet.`,
											Computed:            true,
											NestedObject: schema.NestedAttributeObject{
												Attributes: map[string]schema.Attribute{

													"name": schema.StringAttribute{
														MarkdownDescription: `Option name.`,
														Computed:            true,
													},
													"value": schema.StringAttribute{
														MarkdownDescription: `Option value.`,
														Computed:            true,
													},
												},
											},
										},
										"secs": schema.Int64Attribute{
											MarkdownDescription: `Number of seconds since receiving the packet.`,
											Computed:            true,
										},
										"siaddr": schema.StringAttribute{
											MarkdownDescription: `Server IP address of the packet.`,
											Computed:            true,
										},
										"sname": schema.StringAttribute{
											MarkdownDescription: `Server identifier address of the packet.`,
											Computed:            true,
										},
										"xid": schema.StringAttribute{
											MarkdownDescription: `Transaction id of the packet.`,
											Computed:            true,
										},
										"yiaddr": schema.StringAttribute{
											MarkdownDescription: `Assigned IP address of the packet.`,
											Computed:            true,
										},
									},
								},
								"ip": schema.SingleNestedAttribute{
									MarkdownDescription: `Additional IP attributes of the packet.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"dscp": schema.SingleNestedAttribute{
											MarkdownDescription: `DSCP attributes of the packet.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"ecn": schema.Int64Attribute{
													MarkdownDescription: `ECN number of the packet.`,
													Computed:            true,
												},
												"tag": schema.Int64Attribute{
													MarkdownDescription: `DSCP tag number of the packet.`,
													Computed:            true,
												},
											},
										},
										"header_length": schema.Int64Attribute{
											MarkdownDescription: `IP header length of the packet.`,
											Computed:            true,
										},
										"id": schema.StringAttribute{
											MarkdownDescription: `IP ID of the packet.`,
											Computed:            true,
										},
										"length": schema.Int64Attribute{
											MarkdownDescription: `IP length of the packet.`,
											Computed:            true,
										},
										"protocol": schema.Int64Attribute{
											MarkdownDescription: `IP protocol number of the packet.`,
											Computed:            true,
										},
										"ttl": schema.Int64Attribute{
											MarkdownDescription: `Time to live of the packet.`,
											Computed:            true,
										},
										"version": schema.Int64Attribute{
											MarkdownDescription: `IP version of the packet.`,
											Computed:            true,
										},
									},
								},
								"source": schema.SingleNestedAttribute{
									MarkdownDescription: `Source of the packet.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"ipv4": schema.SingleNestedAttribute{
											MarkdownDescription: `Source ipv4 attributes of the packet.`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"address": schema.StringAttribute{
													MarkdownDescription: `Source ipv4 address of the packet.`,
													Computed:            true,
												},
											},
										},
										"mac": schema.StringAttribute{
											MarkdownDescription: `Source mac address of the packet.`,
											Computed:            true,
										},
										"port": schema.Int64Attribute{
											MarkdownDescription: `Source port of the packet.`,
											Computed:            true,
										},
									},
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `Packet type.`,
									Computed:            true,
								},
								"udp": schema.SingleNestedAttribute{
									MarkdownDescription: `UDP attributes of the packet.`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"checksum": schema.StringAttribute{
											MarkdownDescription: `UDP checksum of the packet.`,
											Computed:            true,
										},
										"length": schema.Int64Attribute{
											MarkdownDescription: `UDP length of the packet.`,
											Computed:            true,
										},
									},
								},
							},
						},
						"last_seen_at": schema.StringAttribute{
							MarkdownDescription: `Last time the server was seen.`,
							Computed:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `Mac address of the server.`,
							Computed:            true,
						},
						"seen_by": schema.SetNestedAttribute{
							MarkdownDescription: `Devices that saw the server.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"name": schema.StringAttribute{
										MarkdownDescription: `Device name.`,
										Computed:            true,
									},
									"serial": schema.StringAttribute{
										MarkdownDescription: `Device serial.`,
										Computed:            true,
									},
									"url": schema.StringAttribute{
										MarkdownDescription: `Url link to device.`,
										Computed:            true,
									},
								},
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `server type. Can be a 'device', 'stack', or 'discovered' (i.e client).`,
							Computed:            true,
						},
						"vlan": schema.Int64Attribute{
							MarkdownDescription: `Vlan id of the server.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchDhcpV4ServersSeenDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchDhcpV4ServersSeen NetworksSwitchDhcpV4ServersSeen
	diags := req.Config.Get(ctx, &networksSwitchDhcpV4ServersSeen)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchDhcpV4ServersSeen")
		vvNetworkID := networksSwitchDhcpV4ServersSeen.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSwitchDhcpV4ServersSeenQueryParams{}

		queryParams1.T0 = networksSwitchDhcpV4ServersSeen.T0.ValueString()
		queryParams1.Timespan = networksSwitchDhcpV4ServersSeen.Timespan.ValueFloat64()
		queryParams1.PerPage = int(networksSwitchDhcpV4ServersSeen.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksSwitchDhcpV4ServersSeen.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksSwitchDhcpV4ServersSeen.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchDhcpV4ServersSeen(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchDhcpV4ServersSeen",
				err.Error(),
			)
			return
		}

		networksSwitchDhcpV4ServersSeen = ResponseSwitchGetNetworkSwitchDhcpV4ServersSeenItemsToBody(networksSwitchDhcpV4ServersSeen, response1)
		diags = resp.State.Set(ctx, &networksSwitchDhcpV4ServersSeen)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchDhcpV4ServersSeen struct {
	NetworkID     types.String                                           `tfsdk:"network_id"`
	T0            types.String                                           `tfsdk:"t0"`
	Timespan      types.Float64                                          `tfsdk:"timespan"`
	PerPage       types.Int64                                            `tfsdk:"per_page"`
	StartingAfter types.String                                           `tfsdk:"starting_after"`
	EndingBefore  types.String                                           `tfsdk:"ending_before"`
	Items         *[]ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeen `tfsdk:"items"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeen struct {
	ClientID     types.String                                                   `tfsdk:"client_id"`
	Device       *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenDevice     `tfsdk:"device"`
	IPv4         *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenIpv4       `tfsdk:"ipv4"`
	IsAllowed    types.Bool                                                     `tfsdk:"is_allowed"`
	IsConfigured types.Bool                                                     `tfsdk:"is_configured"`
	LastAck      *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastAck    `tfsdk:"last_ack"`
	LastPacket   *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacket `tfsdk:"last_packet"`
	LastSeenAt   types.String                                                   `tfsdk:"last_seen_at"`
	Mac          types.String                                                   `tfsdk:"mac"`
	SeenBy       *[]ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenSeenBy   `tfsdk:"seen_by"`
	Type         types.String                                                   `tfsdk:"type"`
	VLAN         types.Int64                                                    `tfsdk:"vlan"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenDevice struct {
	Interface *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenDeviceInterface `tfsdk:"interface"`
	Name      types.String                                                        `tfsdk:"name"`
	Serial    types.String                                                        `tfsdk:"serial"`
	URL       types.String                                                        `tfsdk:"url"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenDeviceInterface struct {
	Name types.String `tfsdk:"name"`
	URL  types.String `tfsdk:"url"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenIpv4 struct {
	Address types.String `tfsdk:"address"`
	Gateway types.String `tfsdk:"gateway"`
	Subnet  types.String `tfsdk:"subnet"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastAck struct {
	IPv4 *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastAckIpv4 `tfsdk:"ipv4"`
	Ts   types.String                                                    `tfsdk:"ts"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastAckIpv4 struct {
	Address types.String `tfsdk:"address"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacket struct {
	Destination *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketDestination `tfsdk:"destination"`
	Ethernet    *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketEthernet    `tfsdk:"ethernet"`
	Fields      *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFields      `tfsdk:"fields"`
	IP          *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketIp          `tfsdk:"ip"`
	Source      *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketSource      `tfsdk:"source"`
	Type        types.String                                                              `tfsdk:"type"`
	UDP         *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketUdp         `tfsdk:"udp"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketDestination struct {
	IPv4 *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 `tfsdk:"ipv4"`
	Mac  types.String                                                                  `tfsdk:"mac"`
	Port types.Int64                                                                   `tfsdk:"port"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 struct {
	Address types.String `tfsdk:"address"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketEthernet struct {
	Type types.String `tfsdk:"type"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFields struct {
	Chaddr      types.String                                                                  `tfsdk:"chaddr"`
	Ciaddr      types.String                                                                  `tfsdk:"ciaddr"`
	Flags       types.String                                                                  `tfsdk:"flags"`
	Giaddr      types.String                                                                  `tfsdk:"giaddr"`
	Hlen        types.Int64                                                                   `tfsdk:"hlen"`
	Hops        types.Int64                                                                   `tfsdk:"hops"`
	Htype       types.Int64                                                                   `tfsdk:"htype"`
	MagicCookie types.String                                                                  `tfsdk:"magic_cookie"`
	Op          types.Int64                                                                   `tfsdk:"op"`
	Options     *[]ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFieldsOptions `tfsdk:"options"`
	Secs        types.Int64                                                                   `tfsdk:"secs"`
	Siaddr      types.String                                                                  `tfsdk:"siaddr"`
	Sname       types.String                                                                  `tfsdk:"sname"`
	Xid         types.String                                                                  `tfsdk:"xid"`
	Yiaddr      types.String                                                                  `tfsdk:"yiaddr"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFieldsOptions struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketIp struct {
	Dscp         *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketIpDscp `tfsdk:"dscp"`
	HeaderLength types.Int64                                                          `tfsdk:"header_length"`
	ID           types.String                                                         `tfsdk:"id"`
	Length       types.Int64                                                          `tfsdk:"length"`
	Protocol     types.Int64                                                          `tfsdk:"protocol"`
	Ttl          types.Int64                                                          `tfsdk:"ttl"`
	Version      types.Int64                                                          `tfsdk:"version"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketIpDscp struct {
	Ecn types.Int64 `tfsdk:"ecn"`
	Tag types.Int64 `tfsdk:"tag"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketSource struct {
	IPv4 *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketSourceIpv4 `tfsdk:"ipv4"`
	Mac  types.String                                                             `tfsdk:"mac"`
	Port types.Int64                                                              `tfsdk:"port"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketSourceIpv4 struct {
	Address types.String `tfsdk:"address"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketUdp struct {
	Checksum types.String `tfsdk:"checksum"`
	Length   types.Int64  `tfsdk:"length"`
}

type ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenSeenBy struct {
	Name   types.String `tfsdk:"name"`
	Serial types.String `tfsdk:"serial"`
	URL    types.String `tfsdk:"url"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchDhcpV4ServersSeenItemsToBody(state NetworksSwitchDhcpV4ServersSeen, response *merakigosdk.ResponseSwitchGetNetworkSwitchDhcpV4ServersSeen) NetworksSwitchDhcpV4ServersSeen {
	var items []ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeen
	for _, item := range *response {
		itemState := ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeen{
			ClientID: types.StringValue(item.ClientID),
			Device: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenDevice {
				if item.Device != nil {
					return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenDevice{
						Interface: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenDeviceInterface {
							if item.Device.Interface != nil {
								return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenDeviceInterface{
									Name: types.StringValue(item.Device.Interface.Name),
									URL:  types.StringValue(item.Device.Interface.URL),
								}
							}
							return nil
						}(),
						Name:   types.StringValue(item.Device.Name),
						Serial: types.StringValue(item.Device.Serial),
						URL:    types.StringValue(item.Device.URL),
					}
				}
				return nil
			}(),
			IPv4: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenIpv4 {
				if item.IPv4 != nil {
					return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenIpv4{
						Address: types.StringValue(item.IPv4.Address),
						Gateway: types.StringValue(item.IPv4.Gateway),
						Subnet:  types.StringValue(item.IPv4.Subnet),
					}
				}
				return nil
			}(),
			IsAllowed: func() types.Bool {
				if item.IsAllowed != nil {
					return types.BoolValue(*item.IsAllowed)
				}
				return types.Bool{}
			}(),
			IsConfigured: func() types.Bool {
				if item.IsConfigured != nil {
					return types.BoolValue(*item.IsConfigured)
				}
				return types.Bool{}
			}(),
			LastAck: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastAck {
				if item.LastAck != nil {
					return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastAck{
						IPv4: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastAckIpv4 {
							if item.LastAck.IPv4 != nil {
								return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastAckIpv4{
									Address: types.StringValue(item.LastAck.IPv4.Address),
								}
							}
							return nil
						}(),
						Ts: types.StringValue(item.LastAck.Ts),
					}
				}
				return nil
			}(),
			LastPacket: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacket {
				if item.LastPacket != nil {
					return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacket{
						Destination: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketDestination {
							if item.LastPacket.Destination != nil {
								return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketDestination{
									IPv4: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketDestinationIpv4 {
										if item.LastPacket.Destination.IPv4 != nil {
											return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketDestinationIpv4{
												Address: types.StringValue(item.LastPacket.Destination.IPv4.Address),
											}
										}
										return nil
									}(),
									Mac: types.StringValue(item.LastPacket.Destination.Mac),
									Port: func() types.Int64 {
										if item.LastPacket.Destination.Port != nil {
											return types.Int64Value(int64(*item.LastPacket.Destination.Port))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						Ethernet: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketEthernet {
							if item.LastPacket.Ethernet != nil {
								return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketEthernet{
									Type: types.StringValue(item.LastPacket.Ethernet.Type),
								}
							}
							return nil
						}(),
						Fields: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFields {
							if item.LastPacket.Fields != nil {
								return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFields{
									Chaddr: types.StringValue(item.LastPacket.Fields.Chaddr),
									Ciaddr: types.StringValue(item.LastPacket.Fields.Ciaddr),
									Flags:  types.StringValue(item.LastPacket.Fields.Flags),
									Giaddr: types.StringValue(item.LastPacket.Fields.Giaddr),
									Hlen: func() types.Int64 {
										if item.LastPacket.Fields.Hlen != nil {
											return types.Int64Value(int64(*item.LastPacket.Fields.Hlen))
										}
										return types.Int64{}
									}(),
									Hops: func() types.Int64 {
										if item.LastPacket.Fields.Hops != nil {
											return types.Int64Value(int64(*item.LastPacket.Fields.Hops))
										}
										return types.Int64{}
									}(),
									Htype: func() types.Int64 {
										if item.LastPacket.Fields.Htype != nil {
											return types.Int64Value(int64(*item.LastPacket.Fields.Htype))
										}
										return types.Int64{}
									}(),
									MagicCookie: types.StringValue(item.LastPacket.Fields.MagicCookie),
									Op: func() types.Int64 {
										if item.LastPacket.Fields.Op != nil {
											return types.Int64Value(int64(*item.LastPacket.Fields.Op))
										}
										return types.Int64{}
									}(),
									Options: func() *[]ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFieldsOptions {
										if item.LastPacket.Fields.Options != nil {
											result := make([]ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFieldsOptions, len(*item.LastPacket.Fields.Options))
											for i, options := range *item.LastPacket.Fields.Options {
												result[i] = ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketFieldsOptions{
													Name:  types.StringValue(options.Name),
													Value: types.StringValue(options.Value),
												}
											}
											return &result
										}
										return nil
									}(),
									Secs: func() types.Int64 {
										if item.LastPacket.Fields.Secs != nil {
											return types.Int64Value(int64(*item.LastPacket.Fields.Secs))
										}
										return types.Int64{}
									}(),
									Siaddr: types.StringValue(item.LastPacket.Fields.Siaddr),
									Sname:  types.StringValue(item.LastPacket.Fields.Sname),
									Xid:    types.StringValue(item.LastPacket.Fields.Xid),
									Yiaddr: types.StringValue(item.LastPacket.Fields.Yiaddr),
								}
							}
							return nil
						}(),
						IP: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketIp {
							if item.LastPacket.IP != nil {
								return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketIp{
									Dscp: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketIpDscp {
										if item.LastPacket.IP.Dscp != nil {
											return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketIpDscp{
												Ecn: func() types.Int64 {
													if item.LastPacket.IP.Dscp.Ecn != nil {
														return types.Int64Value(int64(*item.LastPacket.IP.Dscp.Ecn))
													}
													return types.Int64{}
												}(),
												Tag: func() types.Int64 {
													if item.LastPacket.IP.Dscp.Tag != nil {
														return types.Int64Value(int64(*item.LastPacket.IP.Dscp.Tag))
													}
													return types.Int64{}
												}(),
											}
										}
										return nil
									}(),
									HeaderLength: func() types.Int64 {
										if item.LastPacket.IP.HeaderLength != nil {
											return types.Int64Value(int64(*item.LastPacket.IP.HeaderLength))
										}
										return types.Int64{}
									}(),
									ID: types.StringValue(item.LastPacket.IP.ID),
									Length: func() types.Int64 {
										if item.LastPacket.IP.Length != nil {
											return types.Int64Value(int64(*item.LastPacket.IP.Length))
										}
										return types.Int64{}
									}(),
									Protocol: func() types.Int64 {
										if item.LastPacket.IP.Protocol != nil {
											return types.Int64Value(int64(*item.LastPacket.IP.Protocol))
										}
										return types.Int64{}
									}(),
									Ttl: func() types.Int64 {
										if item.LastPacket.IP.Ttl != nil {
											return types.Int64Value(int64(*item.LastPacket.IP.Ttl))
										}
										return types.Int64{}
									}(),
									Version: func() types.Int64 {
										if item.LastPacket.IP.Version != nil {
											return types.Int64Value(int64(*item.LastPacket.IP.Version))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						Source: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketSource {
							if item.LastPacket.Source != nil {
								return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketSource{
									IPv4: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketSourceIpv4 {
										if item.LastPacket.Source.IPv4 != nil {
											return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketSourceIpv4{
												Address: types.StringValue(item.LastPacket.Source.IPv4.Address),
											}
										}
										return nil
									}(),
									Mac: types.StringValue(item.LastPacket.Source.Mac),
									Port: func() types.Int64 {
										if item.LastPacket.Source.Port != nil {
											return types.Int64Value(int64(*item.LastPacket.Source.Port))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						Type: types.StringValue(item.LastPacket.Type),
						UDP: func() *ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketUdp {
							if item.LastPacket.UDP != nil {
								return &ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenLastPacketUdp{
									Checksum: types.StringValue(item.LastPacket.UDP.Checksum),
									Length: func() types.Int64 {
										if item.LastPacket.UDP.Length != nil {
											return types.Int64Value(int64(*item.LastPacket.UDP.Length))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
			LastSeenAt: types.StringValue(item.LastSeenAt),
			Mac:        types.StringValue(item.Mac),
			SeenBy: func() *[]ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenSeenBy {
				if item.SeenBy != nil {
					result := make([]ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenSeenBy, len(*item.SeenBy))
					for i, seenBy := range *item.SeenBy {
						result[i] = ResponseItemSwitchGetNetworkSwitchDhcpV4ServersSeenSeenBy{
							Name:   types.StringValue(seenBy.Name),
							Serial: types.StringValue(seenBy.Serial),
							URL:    types.StringValue(seenBy.URL),
						}
					}
					return &result
				}
				return nil
			}(),
			Type: types.StringValue(item.Type),
			VLAN: func() types.Int64 {
				if item.VLAN != nil {
					return types.Int64Value(int64(*item.VLAN))
				}
				return types.Int64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
