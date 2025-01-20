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
	_ datasource.DataSource              = &NetworksWirelessSSIDsVpnDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsVpnDataSource{}
)

func NewNetworksWirelessSSIDsVpnDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsVpnDataSource{}
}

type NetworksWirelessSSIDsVpnDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsVpnDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsVpnDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_vpn"
}

func (d *NetworksWirelessSSIDsVpnDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"concentrator": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"name": schema.StringAttribute{
								Computed: true,
							},
							"network_id": schema.StringAttribute{
								Computed: true,
							},
							"vlan_id": schema.Int64Attribute{
								Computed: true,
							},
						},
					},
					"failover": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"heartbeat_interval": schema.Int64Attribute{
								Computed: true,
							},
							"idle_timeout": schema.Int64Attribute{
								Computed: true,
							},
							"request_ip": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"split_tunnel": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								Computed: true,
							},
							"rules": schema.SetNestedAttribute{
								Computed: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"comment": schema.StringAttribute{
											Computed: true,
										},
										"dest_cidr": schema.StringAttribute{
											Computed: true,
										},
										"dest_port": schema.StringAttribute{
											Computed: true,
										},
										"policy": schema.StringAttribute{
											Computed: true,
										},
										"protocol": schema.StringAttribute{
											Computed: true,
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

func (d *NetworksWirelessSSIDsVpnDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsVpn NetworksWirelessSSIDsVpn
	diags := req.Config.Get(ctx, &networksWirelessSSIDsVpn)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDVpn")
		vvNetworkID := networksWirelessSSIDsVpn.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsVpn.Number.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDVpn(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDVpn",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsVpn = ResponseWirelessGetNetworkWirelessSSIDVpnItemToBody(networksWirelessSSIDsVpn, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsVpn)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsVpn struct {
	NetworkID types.String                               `tfsdk:"network_id"`
	Number    types.String                               `tfsdk:"number"`
	Item      *ResponseWirelessGetNetworkWirelessSsidVpn `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSsidVpn struct {
	Concentrator *ResponseWirelessGetNetworkWirelessSsidVpnConcentrator `tfsdk:"concentrator"`
	Failover     *ResponseWirelessGetNetworkWirelessSsidVpnFailover     `tfsdk:"failover"`
	SplitTunnel  *ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnel  `tfsdk:"split_tunnel"`
}

type ResponseWirelessGetNetworkWirelessSsidVpnConcentrator struct {
	Name      types.String `tfsdk:"name"`
	NetworkID types.String `tfsdk:"network_id"`
	VLANID    types.Int64  `tfsdk:"vlan_id"`
}

type ResponseWirelessGetNetworkWirelessSsidVpnFailover struct {
	HeartbeatInterval types.Int64  `tfsdk:"heartbeat_interval"`
	IDleTimeout       types.Int64  `tfsdk:"idle_timeout"`
	RequestIP         types.String `tfsdk:"request_ip"`
}

type ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnel struct {
	Enabled types.Bool                                                   `tfsdk:"enabled"`
	Rules   *[]ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRules `tfsdk:"rules"`
}

type ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRules struct {
	Comment  types.String `tfsdk:"comment"`
	DestCidr types.String `tfsdk:"dest_cidr"`
	DestPort types.String `tfsdk:"dest_port"`
	Policy   types.String `tfsdk:"policy"`
	Protocol types.String `tfsdk:"protocol"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDVpnItemToBody(state NetworksWirelessSSIDsVpn, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDVpn) NetworksWirelessSSIDsVpn {
	itemState := ResponseWirelessGetNetworkWirelessSsidVpn{
		Concentrator: func() *ResponseWirelessGetNetworkWirelessSsidVpnConcentrator {
			if response.Concentrator != nil {
				return &ResponseWirelessGetNetworkWirelessSsidVpnConcentrator{
					Name:      types.StringValue(response.Concentrator.Name),
					NetworkID: types.StringValue(response.Concentrator.NetworkID),
					VLANID: func() types.Int64 {
						if response.Concentrator.VLANID != nil {
							return types.Int64Value(int64(*response.Concentrator.VLANID))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		Failover: func() *ResponseWirelessGetNetworkWirelessSsidVpnFailover {
			if response.Failover != nil {
				return &ResponseWirelessGetNetworkWirelessSsidVpnFailover{
					HeartbeatInterval: func() types.Int64 {
						if response.Failover.HeartbeatInterval != nil {
							return types.Int64Value(int64(*response.Failover.HeartbeatInterval))
						}
						return types.Int64{}
					}(),
					IDleTimeout: func() types.Int64 {
						if response.Failover.IDleTimeout != nil {
							return types.Int64Value(int64(*response.Failover.IDleTimeout))
						}
						return types.Int64{}
					}(),
					RequestIP: types.StringValue(response.Failover.RequestIP),
				}
			}
			return nil
		}(),
		SplitTunnel: func() *ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnel {
			if response.SplitTunnel != nil {
				return &ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnel{
					Enabled: func() types.Bool {
						if response.SplitTunnel.Enabled != nil {
							return types.BoolValue(*response.SplitTunnel.Enabled)
						}
						return types.Bool{}
					}(),
					Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRules {
						if response.SplitTunnel.Rules != nil {
							result := make([]ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRules, len(*response.SplitTunnel.Rules))
							for i, rules := range *response.SplitTunnel.Rules {
								result[i] = ResponseWirelessGetNetworkWirelessSsidVpnSplitTunnelRules{
									Comment:  types.StringValue(rules.Comment),
									DestCidr: types.StringValue(rules.DestCidr),
									DestPort: types.StringValue(rules.DestPort),
									Policy:   types.StringValue(rules.Policy),
									Protocol: types.StringValue(rules.Protocol),
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
	}
	state.Item = &itemState
	return state
}
