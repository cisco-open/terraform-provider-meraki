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
	_ datasource.DataSource              = &NetworksSmDevicesDesktopLogsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmDevicesDesktopLogsDataSource{}
)

func NewNetworksSmDevicesDesktopLogsDataSource() datasource.DataSource {
	return &NetworksSmDevicesDesktopLogsDataSource{}
}

type NetworksSmDevicesDesktopLogsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmDevicesDesktopLogsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmDevicesDesktopLogsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_desktop_logs"
}

func (d *NetworksSmDevicesDesktopLogsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_id": schema.StringAttribute{
				MarkdownDescription: `deviceId path parameter. Device ID`,
				Required:            true,
			},
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

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetNetworkSmDeviceDesktopLogs`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"dhcp_server": schema.StringAttribute{
							MarkdownDescription: `The IP address of the DCHP Server.`,
							Computed:            true,
						},
						"dns_server": schema.StringAttribute{
							MarkdownDescription: `The DNS Server during the connection.`,
							Computed:            true,
						},
						"gateway": schema.StringAttribute{
							MarkdownDescription: `The gateway IP the device was connected to.`,
							Computed:            true,
						},
						"ip": schema.StringAttribute{
							MarkdownDescription: `The IP of the device during connection.`,
							Computed:            true,
						},
						"measured_at": schema.StringAttribute{
							MarkdownDescription: `The time the data was measured at.`,
							Computed:            true,
						},
						"network_device": schema.StringAttribute{
							MarkdownDescription: `The network device for the device used for connection.`,
							Computed:            true,
						},
						"network_driver": schema.StringAttribute{
							MarkdownDescription: `The network driver for the device.`,
							Computed:            true,
						},
						"network_mtu": schema.StringAttribute{
							MarkdownDescription: `The network max transmission unit.`,
							Computed:            true,
						},
						"public_ip": schema.StringAttribute{
							MarkdownDescription: `The public IP address of the device.`,
							Computed:            true,
						},
						"subnet": schema.StringAttribute{
							MarkdownDescription: `The subnet of the device connection.`,
							Computed:            true,
						},
						"ts": schema.StringAttribute{
							MarkdownDescription: `The time the connection was logged.`,
							Computed:            true,
						},
						"user": schema.StringAttribute{
							MarkdownDescription: `The user during connection.`,
							Computed:            true,
						},
						"wifi_auth": schema.StringAttribute{
							MarkdownDescription: `The type of authentication used by the SSID.`,
							Computed:            true,
						},
						"wifi_bssid": schema.StringAttribute{
							MarkdownDescription: `The MAC of the access point the device is connected to.`,
							Computed:            true,
						},
						"wifi_channel": schema.StringAttribute{
							MarkdownDescription: `Channel through which the connection is routing.`,
							Computed:            true,
						},
						"wifi_noise": schema.StringAttribute{
							MarkdownDescription: `The wireless signal power level received by the device.`,
							Computed:            true,
						},
						"wifi_rssi": schema.StringAttribute{
							MarkdownDescription: `The Received Signal Strength Indicator for the device.`,
							Computed:            true,
						},
						"wifi_ssid": schema.StringAttribute{
							MarkdownDescription: `The name of the network the device is connected to.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmDevicesDesktopLogsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmDevicesDesktopLogs NetworksSmDevicesDesktopLogs
	diags := req.Config.Get(ctx, &networksSmDevicesDesktopLogs)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmDeviceDesktopLogs")
		vvNetworkID := networksSmDevicesDesktopLogs.NetworkID.ValueString()
		vvDeviceID := networksSmDevicesDesktopLogs.DeviceID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSmDeviceDesktopLogsQueryParams{}

		queryParams1.PerPage = int(networksSmDevicesDesktopLogs.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksSmDevicesDesktopLogs.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksSmDevicesDesktopLogs.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetNetworkSmDeviceDesktopLogs(vvNetworkID, vvDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmDeviceDesktopLogs",
				err.Error(),
			)
			return
		}

		networksSmDevicesDesktopLogs = ResponseSmGetNetworkSmDeviceDesktopLogsItemsToBody(networksSmDevicesDesktopLogs, response1)
		diags = resp.State.Set(ctx, &networksSmDevicesDesktopLogs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmDevicesDesktopLogs struct {
	NetworkID     types.String                                   `tfsdk:"network_id"`
	DeviceID      types.String                                   `tfsdk:"device_id"`
	PerPage       types.Int64                                    `tfsdk:"per_page"`
	StartingAfter types.String                                   `tfsdk:"starting_after"`
	EndingBefore  types.String                                   `tfsdk:"ending_before"`
	Items         *[]ResponseItemSmGetNetworkSmDeviceDesktopLogs `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmDeviceDesktopLogs struct {
	DhcpServer    types.String `tfsdk:"dhcp_server"`
	DNSServer     types.String `tfsdk:"dns_server"`
	Gateway       types.String `tfsdk:"gateway"`
	IP            types.String `tfsdk:"ip"`
	MeasuredAt    types.String `tfsdk:"measured_at"`
	NetworkDevice types.String `tfsdk:"network_device"`
	NetworkDriver types.String `tfsdk:"network_driver"`
	NetworkMTU    types.String `tfsdk:"network_mtu"`
	PublicIP      types.String `tfsdk:"public_ip"`
	Subnet        types.String `tfsdk:"subnet"`
	Ts            types.String `tfsdk:"ts"`
	User          types.String `tfsdk:"user"`
	WifiAuth      types.String `tfsdk:"wifi_auth"`
	WifiBssid     types.String `tfsdk:"wifi_bssid"`
	WifiChannel   types.String `tfsdk:"wifi_channel"`
	WifiNoise     types.String `tfsdk:"wifi_noise"`
	WifiRssi      types.String `tfsdk:"wifi_rssi"`
	WifiSSID      types.String `tfsdk:"wifi_ssid"`
}

// ToBody
func ResponseSmGetNetworkSmDeviceDesktopLogsItemsToBody(state NetworksSmDevicesDesktopLogs, response *merakigosdk.ResponseSmGetNetworkSmDeviceDesktopLogs) NetworksSmDevicesDesktopLogs {
	var items []ResponseItemSmGetNetworkSmDeviceDesktopLogs
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmDeviceDesktopLogs{
			DhcpServer: func() types.String {
				if item.DhcpServer != "" {
					return types.StringValue(item.DhcpServer)
				}
				return types.String{}
			}(),
			DNSServer: func() types.String {
				if item.DNSServer != "" {
					return types.StringValue(item.DNSServer)
				}
				return types.String{}
			}(),
			Gateway: func() types.String {
				if item.Gateway != "" {
					return types.StringValue(item.Gateway)
				}
				return types.String{}
			}(),
			IP: func() types.String {
				if item.IP != "" {
					return types.StringValue(item.IP)
				}
				return types.String{}
			}(),
			MeasuredAt: func() types.String {
				if item.MeasuredAt != "" {
					return types.StringValue(item.MeasuredAt)
				}
				return types.String{}
			}(),
			NetworkDevice: func() types.String {
				if item.NetworkDevice != "" {
					return types.StringValue(item.NetworkDevice)
				}
				return types.String{}
			}(),
			NetworkDriver: func() types.String {
				if item.NetworkDriver != "" {
					return types.StringValue(item.NetworkDriver)
				}
				return types.String{}
			}(),
			NetworkMTU: func() types.String {
				if item.NetworkMTU != "" {
					return types.StringValue(item.NetworkMTU)
				}
				return types.String{}
			}(),
			PublicIP: func() types.String {
				if item.PublicIP != "" {
					return types.StringValue(item.PublicIP)
				}
				return types.String{}
			}(),
			Subnet: func() types.String {
				if item.Subnet != "" {
					return types.StringValue(item.Subnet)
				}
				return types.String{}
			}(),
			Ts: func() types.String {
				if item.Ts != "" {
					return types.StringValue(item.Ts)
				}
				return types.String{}
			}(),
			User: func() types.String {
				if item.User != "" {
					return types.StringValue(item.User)
				}
				return types.String{}
			}(),
			WifiAuth: func() types.String {
				if item.WifiAuth != "" {
					return types.StringValue(item.WifiAuth)
				}
				return types.String{}
			}(),
			WifiBssid: func() types.String {
				if item.WifiBssid != "" {
					return types.StringValue(item.WifiBssid)
				}
				return types.String{}
			}(),
			WifiChannel: func() types.String {
				if item.WifiChannel != "" {
					return types.StringValue(item.WifiChannel)
				}
				return types.String{}
			}(),
			WifiNoise: func() types.String {
				if item.WifiNoise != "" {
					return types.StringValue(item.WifiNoise)
				}
				return types.String{}
			}(),
			WifiRssi: func() types.String {
				if item.WifiRssi != "" {
					return types.StringValue(item.WifiRssi)
				}
				return types.String{}
			}(),
			WifiSSID: func() types.String {
				if item.WifiSSID != "" {
					return types.StringValue(item.WifiSSID)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
