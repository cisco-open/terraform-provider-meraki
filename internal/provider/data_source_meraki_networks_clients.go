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
	_ datasource.DataSource              = &NetworksClientsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksClientsDataSource{}
)

func NewNetworksClientsDataSource() datasource.DataSource {
	return &NetworksClientsDataSource{}
}

type NetworksClientsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksClientsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksClientsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_clients"
}

func (d *NetworksClientsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"client_id": schema.StringAttribute{
				MarkdownDescription: `clientId path parameter. Client ID`,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"cdp": schema.SetNestedAttribute{
						MarkdownDescription: `The Cisco discover protocol settings for the client`,
						Computed:            true,
					},
					"client_vpn_connections": schema.SetNestedAttribute{
						MarkdownDescription: `VPN connections associated with the client`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"connected_at": schema.Int64Attribute{
									MarkdownDescription: `The time the client last connected to the VPN`,
									Computed:            true,
								},
								"disconnected_at": schema.Int64Attribute{
									MarkdownDescription: `The time the client last disconnectd from the VPN`,
									Computed:            true,
								},
								"remote_ip": schema.StringAttribute{
									MarkdownDescription: `The IP address of the VPN the client last connected to`,
									Computed:            true,
								},
							},
						},
					},
					"description": schema.StringAttribute{
						MarkdownDescription: `Short description of the client`,
						Computed:            true,
					},
					"first_seen": schema.Int64Attribute{
						MarkdownDescription: `Timestamp client was first seen in the network`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `The ID of the client`,
						Computed:            true,
					},
					"ip": schema.StringAttribute{
						MarkdownDescription: `The IP address of the client`,
						Computed:            true,
					},
					"ip6": schema.StringAttribute{
						MarkdownDescription: `The IPv6 address of the client`,
						Computed:            true,
					},
					"last_seen": schema.Int64Attribute{
						MarkdownDescription: `Timestamp client was last seen in the network`,
						Computed:            true,
					},
					"lldp": schema.SetNestedAttribute{
						MarkdownDescription: `The link layer discover protocol settings for the client`,
						Computed:            true,
					},
					"mac": schema.StringAttribute{
						MarkdownDescription: `The MAC address of the client`,
						Computed:            true,
					},
					"manufacturer": schema.StringAttribute{
						MarkdownDescription: `Manufacturer of the client`,
						Computed:            true,
					},
					"notes": schema.StringAttribute{
						MarkdownDescription: `The notes associated with the client`,
						Computed:            true,
					},
					"os": schema.StringAttribute{
						MarkdownDescription: `The operating system of the client`,
						Computed:            true,
					},
					"recent_device_connection": schema.StringAttribute{
						MarkdownDescription: `Client's most recent connection type`,
						Computed:            true,
					},
					"recent_device_mac": schema.StringAttribute{
						MarkdownDescription: `The MAC address of the node that the device was last connected to`,
						Computed:            true,
					},
					"sm_installed": schema.BoolAttribute{
						MarkdownDescription: `Status of SM for the client`,
						Computed:            true,
					},
					"ssid": schema.StringAttribute{
						MarkdownDescription: `The name of the SSID that the client is connected to`,
						Computed:            true,
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `The connection status of the client`,
						Computed:            true,
					},
					"switchport": schema.StringAttribute{
						MarkdownDescription: `The switch port that the client is connected to`,
						Computed:            true,
					},
					"user": schema.StringAttribute{
						MarkdownDescription: `The username of the user of the client`,
						Computed:            true,
					},
					"vlan": schema.StringAttribute{
						MarkdownDescription: `The name of the VLAN that the client is connected to`,
						Computed:            true,
					},
					"wireless_capabilities": schema.StringAttribute{
						MarkdownDescription: `Wireless capabilities of the client`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksClientsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksClients NetworksClients
	diags := req.Config.Get(ctx, &networksClients)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkClient")
		vvNetworkID := networksClients.NetworkID.ValueString()
		vvClientID := networksClients.ClientID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkClient(vvNetworkID, vvClientID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkClient",
				err.Error(),
			)
			return
		}

		networksClients = ResponseNetworksGetNetworkClientItemToBody(networksClients, response1)
		diags = resp.State.Set(ctx, &networksClients)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksClients struct {
	NetworkID types.String                      `tfsdk:"network_id"`
	ClientID  types.String                      `tfsdk:"client_id"`
	Item      *ResponseNetworksGetNetworkClient `tfsdk:"item"`
}

type ResponseNetworksGetNetworkClient struct {
	Cdp                    *[][]string                                             `tfsdk:"cdp"`
	ClientVpnConnections   *[]ResponseNetworksGetNetworkClientClientVpnConnections `tfsdk:"client_vpn_connections"`
	Description            types.String                                            `tfsdk:"description"`
	FirstSeen              types.Int64                                             `tfsdk:"first_seen"`
	ID                     types.String                                            `tfsdk:"id"`
	IP                     types.String                                            `tfsdk:"ip"`
	IP6                    types.String                                            `tfsdk:"ip6"`
	LastSeen               types.Int64                                             `tfsdk:"last_seen"`
	Lldp                   *[][]string                                             `tfsdk:"lldp"`
	Mac                    types.String                                            `tfsdk:"mac"`
	Manufacturer           types.String                                            `tfsdk:"manufacturer"`
	Notes                  types.String                                            `tfsdk:"notes"`
	Os                     types.String                                            `tfsdk:"os"`
	RecentDeviceConnection types.String                                            `tfsdk:"recent_device_connection"`
	RecentDeviceMac        types.String                                            `tfsdk:"recent_device_mac"`
	SmInstalled            types.Bool                                              `tfsdk:"sm_installed"`
	SSID                   types.String                                            `tfsdk:"ssid"`
	Status                 types.String                                            `tfsdk:"status"`
	Switchport             types.String                                            `tfsdk:"switchport"`
	User                   types.String                                            `tfsdk:"user"`
	VLAN                   types.String                                            `tfsdk:"vlan"`
	WirelessCapabilities   types.String                                            `tfsdk:"wireless_capabilities"`
}

type ResponseNetworksGetNetworkClientClientVpnConnections struct {
	ConnectedAt    types.Int64  `tfsdk:"connected_at"`
	DisconnectedAt types.Int64  `tfsdk:"disconnected_at"`
	RemoteIP       types.String `tfsdk:"remote_ip"`
}

// ToBody
func ResponseNetworksGetNetworkClientItemToBody(state NetworksClients, response *merakigosdk.ResponseNetworksGetNetworkClient) NetworksClients {
	itemState := ResponseNetworksGetNetworkClient{
		//TODO [][]
		ClientVpnConnections: func() *[]ResponseNetworksGetNetworkClientClientVpnConnections {
			if response.ClientVpnConnections != nil {
				result := make([]ResponseNetworksGetNetworkClientClientVpnConnections, len(*response.ClientVpnConnections))
				for i, clientVpnConnections := range *response.ClientVpnConnections {
					result[i] = ResponseNetworksGetNetworkClientClientVpnConnections{
						ConnectedAt: func() types.Int64 {
							if clientVpnConnections.ConnectedAt != nil {
								return types.Int64Value(int64(*clientVpnConnections.ConnectedAt))
							}
							return types.Int64{}
						}(),
						DisconnectedAt: func() types.Int64 {
							if clientVpnConnections.DisconnectedAt != nil {
								return types.Int64Value(int64(*clientVpnConnections.DisconnectedAt))
							}
							return types.Int64{}
						}(),
						RemoteIP: types.StringValue(clientVpnConnections.RemoteIP),
					}
				}
				return &result
			}
			return nil
		}(),
		Description: types.StringValue(response.Description),
		FirstSeen: func() types.Int64 {
			if response.FirstSeen != nil {
				return types.Int64Value(int64(*response.FirstSeen))
			}
			return types.Int64{}
		}(),
		ID:  types.StringValue(response.ID),
		IP:  types.StringValue(response.IP),
		IP6: types.StringValue(response.IP6),
		LastSeen: func() types.Int64 {
			if response.LastSeen != nil {
				return types.Int64Value(int64(*response.LastSeen))
			}
			return types.Int64{}
		}(),
		//TODO [][]
		Mac:                    types.StringValue(response.Mac),
		Manufacturer:           types.StringValue(response.Manufacturer),
		Notes:                  types.StringValue(response.Notes),
		Os:                     types.StringValue(response.Os),
		RecentDeviceConnection: types.StringValue(response.RecentDeviceConnection),
		RecentDeviceMac:        types.StringValue(response.RecentDeviceMac),
		SmInstalled: func() types.Bool {
			if response.SmInstalled != nil {
				return types.BoolValue(*response.SmInstalled)
			}
			return types.Bool{}
		}(),
		SSID:                 types.StringValue(response.SSID),
		Status:               types.StringValue(response.Status),
		Switchport:           types.StringValue(response.Switchport),
		User:                 types.StringValue(response.User),
		VLAN:                 types.StringValue(response.VLAN),
		WirelessCapabilities: types.StringValue(response.WirelessCapabilities),
	}
	state.Item = &itemState
	return state
}
