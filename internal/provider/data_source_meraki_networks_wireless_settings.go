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
	_ datasource.DataSource              = &NetworksWirelessSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSettingsDataSource{}
)

func NewNetworksWirelessSettingsDataSource() datasource.DataSource {
	return &NetworksWirelessSettingsDataSource{}
}

type NetworksWirelessSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_settings"
}

func (d *NetworksWirelessSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"ipv6_bridge_enabled": schema.BoolAttribute{
						MarkdownDescription: `Toggle for enabling or disabling IPv6 bridging in a network (Note: if enabled, SSIDs must also be configured to use bridge mode)`,
						Computed:            true,
					},
					"led_lights_on": schema.BoolAttribute{
						MarkdownDescription: `Toggle for enabling or disabling LED lights on all APs in the network (making them run dark)`,
						Computed:            true,
					},
					"location_analytics_enabled": schema.BoolAttribute{
						MarkdownDescription: `Toggle for enabling or disabling location analytics for your network`,
						Computed:            true,
					},
					"meshing_enabled": schema.BoolAttribute{
						MarkdownDescription: `Toggle for enabling or disabling meshing in a network`,
						Computed:            true,
					},
					"named_vlans": schema.SingleNestedAttribute{
						MarkdownDescription: `Named VLAN settings for wireless networks.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"pool_dhcp_monitoring": schema.SingleNestedAttribute{
								MarkdownDescription: `Named VLAN Pool DHCP Monitoring settings.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"duration": schema.Int64Attribute{
										MarkdownDescription: `The duration in minutes that devices will refrain from using dirty VLANs before adding them back to the pool.`,
										Computed:            true,
									},
									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Whether or not devices using named VLAN pools should remove dirty VLANs from the pool, thereby preventing clients from being assigned to VLANs where they would be unable to obtain an IP address via DHCP`,
										Computed:            true,
									},
								},
							},
						},
					},
					"regulatory_domain": schema.SingleNestedAttribute{
						MarkdownDescription: `Regulatory domain information for this network.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"country_code": schema.StringAttribute{
								MarkdownDescription: `The country code of the regulatory domain.`,
								Computed:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `The name of the regulatory domain for this network.`,
								Computed:            true,
							},
							"permits6e": schema.BoolAttribute{
								MarkdownDescription: `Whether or not the regulatory domain for this network permits Wifi 6E.`,
								Computed:            true,
							},
						},
					},
					"upgradestrategy": schema.StringAttribute{
						MarkdownDescription: `The default strategy that network devices will use to perform an upgrade. Requires firmware version MR 26.8 or higher.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSettings NetworksWirelessSettings
	diags := req.Config.Get(ctx, &networksWirelessSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSettings")
		vvNetworkID := networksWirelessSettings.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSettings(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSettings",
				err.Error(),
			)
			return
		}

		networksWirelessSettings = ResponseWirelessGetNetworkWirelessSettingsItemToBody(networksWirelessSettings, response1)
		diags = resp.State.Set(ctx, &networksWirelessSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSettings struct {
	NetworkID types.String                                `tfsdk:"network_id"`
	Item      *ResponseWirelessGetNetworkWirelessSettings `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSettings struct {
	IPv6BridgeEnabled        types.Bool                                                  `tfsdk:"ipv6_bridge_enabled"`
	LedLightsOn              types.Bool                                                  `tfsdk:"led_lights_on"`
	LocationAnalyticsEnabled types.Bool                                                  `tfsdk:"location_analytics_enabled"`
	MeshingEnabled           types.Bool                                                  `tfsdk:"meshing_enabled"`
	NamedVLANs               *ResponseWirelessGetNetworkWirelessSettingsNamedVlans       `tfsdk:"named_vlans"`
	RegulatoryDomain         *ResponseWirelessGetNetworkWirelessSettingsRegulatoryDomain `tfsdk:"regulatory_domain"`
	Upgradestrategy          types.String                                                `tfsdk:"upgrade_strategy"`
}

type ResponseWirelessGetNetworkWirelessSettingsNamedVlans struct {
	PoolDhcpMonitoring *ResponseWirelessGetNetworkWirelessSettingsNamedVlansPoolDhcpMonitoring `tfsdk:"pool_dhcp_monitoring"`
}

type ResponseWirelessGetNetworkWirelessSettingsNamedVlansPoolDhcpMonitoring struct {
	Duration types.Int64 `tfsdk:"duration"`
	Enabled  types.Bool  `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessSettingsRegulatoryDomain struct {
	CountryCode types.String `tfsdk:"country_code"`
	Name        types.String `tfsdk:"name"`
	Permits6E   types.Bool   `tfsdk:"permits6e"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSettingsItemToBody(state NetworksWirelessSettings, response *merakigosdk.ResponseWirelessGetNetworkWirelessSettings) NetworksWirelessSettings {
	itemState := ResponseWirelessGetNetworkWirelessSettings{
		IPv6BridgeEnabled: func() types.Bool {
			if response.IPv6BridgeEnabled != nil {
				return types.BoolValue(*response.IPv6BridgeEnabled)
			}
			return types.Bool{}
		}(),
		LedLightsOn: func() types.Bool {
			if response.LedLightsOn != nil {
				return types.BoolValue(*response.LedLightsOn)
			}
			return types.Bool{}
		}(),
		LocationAnalyticsEnabled: func() types.Bool {
			if response.LocationAnalyticsEnabled != nil {
				return types.BoolValue(*response.LocationAnalyticsEnabled)
			}
			return types.Bool{}
		}(),
		MeshingEnabled: func() types.Bool {
			if response.MeshingEnabled != nil {
				return types.BoolValue(*response.MeshingEnabled)
			}
			return types.Bool{}
		}(),
		NamedVLANs: func() *ResponseWirelessGetNetworkWirelessSettingsNamedVlans {
			if response.NamedVLANs != nil {
				return &ResponseWirelessGetNetworkWirelessSettingsNamedVlans{
					PoolDhcpMonitoring: func() *ResponseWirelessGetNetworkWirelessSettingsNamedVlansPoolDhcpMonitoring {
						if response.NamedVLANs.PoolDhcpMonitoring != nil {
							return &ResponseWirelessGetNetworkWirelessSettingsNamedVlansPoolDhcpMonitoring{
								Duration: func() types.Int64 {
									if response.NamedVLANs.PoolDhcpMonitoring.Duration != nil {
										return types.Int64Value(int64(*response.NamedVLANs.PoolDhcpMonitoring.Duration))
									}
									return types.Int64{}
								}(),
								Enabled: func() types.Bool {
									if response.NamedVLANs.PoolDhcpMonitoring.Enabled != nil {
										return types.BoolValue(*response.NamedVLANs.PoolDhcpMonitoring.Enabled)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		RegulatoryDomain: func() *ResponseWirelessGetNetworkWirelessSettingsRegulatoryDomain {
			if response.RegulatoryDomain != nil {
				return &ResponseWirelessGetNetworkWirelessSettingsRegulatoryDomain{
					CountryCode: types.StringValue(response.RegulatoryDomain.CountryCode),
					Name:        types.StringValue(response.RegulatoryDomain.Name),
					Permits6E: func() types.Bool {
						if response.RegulatoryDomain.Permits6E != nil {
							return types.BoolValue(*response.RegulatoryDomain.Permits6E)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		Upgradestrategy: types.StringValue(response.Upgradestrategy),
	}
	state.Item = &itemState
	return state
}
