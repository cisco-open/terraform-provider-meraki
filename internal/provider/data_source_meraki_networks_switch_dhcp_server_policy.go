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
	_ datasource.DataSource              = &NetworksSwitchDhcpServerPolicyDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchDhcpServerPolicyDataSource{}
)

func NewNetworksSwitchDhcpServerPolicyDataSource() datasource.DataSource {
	return &NetworksSwitchDhcpServerPolicyDataSource{}
}

type NetworksSwitchDhcpServerPolicyDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchDhcpServerPolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchDhcpServerPolicyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_dhcp_server_policy"
}

func (d *NetworksSwitchDhcpServerPolicyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"alerts": schema.SingleNestedAttribute{
						MarkdownDescription: `Email alert settings for DHCP servers`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"email": schema.SingleNestedAttribute{
								MarkdownDescription: `Alert settings for DHCP servers`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `When enabled, send an email if a new DHCP server is seen. Default value is false.`,
										Computed:            true,
									},
								},
							},
						},
					},
					"allowed_servers": schema.ListAttribute{
						MarkdownDescription: `List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set
      to block.An empty array will clear the entries.`,
						Computed:    true,
						ElementType: types.StringType,
					},
					"arp_inspection": schema.SingleNestedAttribute{
						MarkdownDescription: `Dynamic ARP Inspection settings`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enable or disable Dynamic ARP Inspection on the network. Default value is false.`,
								Computed:            true,
							},
							"unsupported_models": schema.ListAttribute{
								MarkdownDescription: `List of switch models that does not support dynamic ARP inspection`,
								Computed:            true,
								ElementType:         types.StringType,
							},
						},
					},
					"blocked_servers": schema.ListAttribute{
						MarkdownDescription: `List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set
      to allow.An empty array will clear the entries.`,
						Computed:    true,
						ElementType: types.StringType,
					},
					"default_policy": schema.StringAttribute{
						MarkdownDescription: `'allow' or 'block' new DHCP servers. Default value is 'allow'.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchDhcpServerPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchDhcpServerPolicy NetworksSwitchDhcpServerPolicy
	diags := req.Config.Get(ctx, &networksSwitchDhcpServerPolicy)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchDhcpServerPolicy")
		vvNetworkID := networksSwitchDhcpServerPolicy.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchDhcpServerPolicy(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchDhcpServerPolicy",
				err.Error(),
			)
			return
		}

		networksSwitchDhcpServerPolicy = ResponseSwitchGetNetworkSwitchDhcpServerPolicyItemToBody(networksSwitchDhcpServerPolicy, response1)
		diags = resp.State.Set(ctx, &networksSwitchDhcpServerPolicy)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchDhcpServerPolicy struct {
	NetworkID types.String                                    `tfsdk:"network_id"`
	Item      *ResponseSwitchGetNetworkSwitchDhcpServerPolicy `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchDhcpServerPolicy struct {
	Alerts         *ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlerts        `tfsdk:"alerts"`
	AllowedServers types.List                                                   `tfsdk:"allowed_servers"`
	ArpInspection  *ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspection `tfsdk:"arp_inspection"`
	BlockedServers types.List                                                   `tfsdk:"blocked_servers"`
	DefaultPolicy  types.String                                                 `tfsdk:"default_policy"`
}

type ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlerts struct {
	Email *ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsEmail `tfsdk:"email"`
}

type ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsEmail struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspection struct {
	Enabled           types.Bool `tfsdk:"enabled"`
	UnsupportedModels types.List `tfsdk:"unsupported_models"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchDhcpServerPolicyItemToBody(state NetworksSwitchDhcpServerPolicy, response *merakigosdk.ResponseSwitchGetNetworkSwitchDhcpServerPolicy) NetworksSwitchDhcpServerPolicy {
	itemState := ResponseSwitchGetNetworkSwitchDhcpServerPolicy{
		Alerts: func() *ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlerts {
			if response.Alerts != nil {
				return &ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlerts{
					Email: func() *ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsEmail {
						if response.Alerts.Email != nil {
							return &ResponseSwitchGetNetworkSwitchDhcpServerPolicyAlertsEmail{
								Enabled: func() types.Bool {
									if response.Alerts.Email.Enabled != nil {
										return types.BoolValue(*response.Alerts.Email.Enabled)
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
		AllowedServers: StringSliceToList(response.AllowedServers),
		ArpInspection: func() *ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspection {
			if response.ArpInspection != nil {
				return &ResponseSwitchGetNetworkSwitchDhcpServerPolicyArpInspection{
					Enabled: func() types.Bool {
						if response.ArpInspection.Enabled != nil {
							return types.BoolValue(*response.ArpInspection.Enabled)
						}
						return types.Bool{}
					}(),
					UnsupportedModels: StringSliceToList(response.ArpInspection.UnsupportedModels),
				}
			}
			return nil
		}(),
		BlockedServers: StringSliceToList(response.BlockedServers),
		DefaultPolicy:  types.StringValue(response.DefaultPolicy),
	}
	state.Item = &itemState
	return state
}
