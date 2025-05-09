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
	_ datasource.DataSource              = &NetworksApplianceVpnSiteToSiteVpnDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceVpnSiteToSiteVpnDataSource{}
)

func NewNetworksApplianceVpnSiteToSiteVpnDataSource() datasource.DataSource {
	return &NetworksApplianceVpnSiteToSiteVpnDataSource{}
}

type NetworksApplianceVpnSiteToSiteVpnDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceVpnSiteToSiteVpnDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceVpnSiteToSiteVpnDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_vpn_site_to_site_vpn"
}

func (d *NetworksApplianceVpnSiteToSiteVpnDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"hubs": schema.SetNestedAttribute{
						MarkdownDescription: `The list of VPN hubs, in order of preference.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"hub_id": schema.StringAttribute{
									MarkdownDescription: `The network ID of the hub.`,
									Computed:            true,
								},
								"use_default_route": schema.BoolAttribute{
									MarkdownDescription: `Indicates whether default route traffic should be sent to this hub.`,
									Computed:            true,
								},
							},
						},
					},
					"mode": schema.StringAttribute{
						MarkdownDescription: `The site-to-site VPN mode.`,
						Computed:            true,
					},
					"subnet": schema.SingleNestedAttribute{
						MarkdownDescription: `Configuration of subnet features`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"nat": schema.SingleNestedAttribute{
								MarkdownDescription: `Configuration of NAT for subnets`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"is_allowed": schema.BoolAttribute{
										MarkdownDescription: `If enabled, VPN subnet translation can be used to translate any local subnets that are allowed to use the VPN into a new subnet with the same number of addresses.`,
										Computed:            true,
									},
								},
							},
						},
					},
					"subnets": schema.SetNestedAttribute{
						MarkdownDescription: `The list of subnets and their VPN presence.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"local_subnet": schema.StringAttribute{
									MarkdownDescription: `The CIDR notation subnet used within the VPN`,
									Computed:            true,
								},
								"nat": schema.SingleNestedAttribute{
									MarkdownDescription: `Configuration of NAT for the subnet`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"enabled": schema.BoolAttribute{
											MarkdownDescription: `Whether or not VPN subnet translation is enabled for the subnet`,
											Computed:            true,
										},
										"remote_subnet": schema.StringAttribute{
											MarkdownDescription: `The translated subnet to be used in the VPN`,
											Computed:            true,
										},
									},
								},
								"use_vpn": schema.BoolAttribute{
									MarkdownDescription: `Indicates the presence of the subnet in the VPN`,
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

func (d *NetworksApplianceVpnSiteToSiteVpnDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceVpnSiteToSiteVpn NetworksApplianceVpnSiteToSiteVpn
	diags := req.Config.Get(ctx, &networksApplianceVpnSiteToSiteVpn)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceVpnSiteToSiteVpn")
		vvNetworkID := networksApplianceVpnSiteToSiteVpn.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceVpnSiteToSiteVpn(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceVpnSiteToSiteVpn",
				err.Error(),
			)
			return
		}

		networksApplianceVpnSiteToSiteVpn = ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnItemToBody(networksApplianceVpnSiteToSiteVpn, response1)
		diags = resp.State.Set(ctx, &networksApplianceVpnSiteToSiteVpn)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceVpnSiteToSiteVpn struct {
	NetworkID types.String                                          `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpn `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpn struct {
	Hubs    *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubs    `tfsdk:"hubs"`
	Mode    types.String                                                   `tfsdk:"mode"`
	Subnet  *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnet    `tfsdk:"subnet"`
	Subnets *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnets `tfsdk:"subnets"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubs struct {
	HubID           types.String `tfsdk:"hub_id"`
	UseDefaultRoute types.Bool   `tfsdk:"use_default_route"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnet struct {
	Nat *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetNat `tfsdk:"nat"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetNat struct {
	IsAllowed types.Bool `tfsdk:"is_allowed"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnets struct {
	LocalSubnet types.String                                                    `tfsdk:"local_subnet"`
	Nat         *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsNat `tfsdk:"nat"`
	UseVpn      types.Bool                                                      `tfsdk:"use_vpn"`
}

type ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsNat struct {
	Enabled      types.Bool   `tfsdk:"enabled"`
	RemoteSubnet types.String `tfsdk:"remote_subnet"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnItemToBody(state NetworksApplianceVpnSiteToSiteVpn, response *merakigosdk.ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpn) NetworksApplianceVpnSiteToSiteVpn {
	itemState := ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpn{
		Hubs: func() *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubs {
			if response.Hubs != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubs, len(*response.Hubs))
				for i, hubs := range *response.Hubs {
					result[i] = ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnHubs{
						HubID: types.StringValue(hubs.HubID),
						UseDefaultRoute: func() types.Bool {
							if hubs.UseDefaultRoute != nil {
								return types.BoolValue(*hubs.UseDefaultRoute)
							}
							return types.Bool{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Mode: types.StringValue(response.Mode),
		Subnet: func() *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnet {
			if response.Subnet != nil {
				return &ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnet{
					Nat: func() *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetNat {
						if response.Subnet.Nat != nil {
							return &ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetNat{
								IsAllowed: func() types.Bool {
									if response.Subnet.Nat.IsAllowed != nil {
										return types.BoolValue(*response.Subnet.Nat.IsAllowed)
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
		Subnets: func() *[]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnets {
			if response.Subnets != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnets, len(*response.Subnets))
				for i, subnets := range *response.Subnets {
					result[i] = ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnets{
						LocalSubnet: types.StringValue(subnets.LocalSubnet),
						Nat: func() *ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsNat {
							if subnets.Nat != nil {
								return &ResponseApplianceGetNetworkApplianceVpnSiteToSiteVpnSubnetsNat{
									Enabled: func() types.Bool {
										if subnets.Nat.Enabled != nil {
											return types.BoolValue(*subnets.Nat.Enabled)
										}
										return types.Bool{}
									}(),
									RemoteSubnet: types.StringValue(subnets.Nat.RemoteSubnet),
								}
							}
							return nil
						}(),
						UseVpn: func() types.Bool {
							if subnets.UseVpn != nil {
								return types.BoolValue(*subnets.UseVpn)
							}
							return types.Bool{}
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
