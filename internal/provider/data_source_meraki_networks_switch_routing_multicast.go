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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksSwitchRoutingMulticastDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchRoutingMulticastDataSource{}
)

func NewNetworksSwitchRoutingMulticastDataSource() datasource.DataSource {
	return &NetworksSwitchRoutingMulticastDataSource{}
}

type NetworksSwitchRoutingMulticastDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchRoutingMulticastDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchRoutingMulticastDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_routing_multicast"
}

func (d *NetworksSwitchRoutingMulticastDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"default_settings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"flood_unknown_multicast_traffic_enabled": schema.BoolAttribute{
								Computed: true,
							},
							"igmp_snooping_enabled": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
					"overrides": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"flood_unknown_multicast_traffic_enabled": schema.BoolAttribute{
									Computed: true,
								},
								"igmp_snooping_enabled": schema.BoolAttribute{
									Computed: true,
								},
								"switches": schema.ListAttribute{
									Computed:    true,
									ElementType: types.StringType,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchRoutingMulticastDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchRoutingMulticast NetworksSwitchRoutingMulticast
	diags := req.Config.Get(ctx, &networksSwitchRoutingMulticast)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchRoutingMulticast")
		vvNetworkID := networksSwitchRoutingMulticast.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchRoutingMulticast(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchRoutingMulticast",
				err.Error(),
			)
			return
		}

		networksSwitchRoutingMulticast = ResponseSwitchGetNetworkSwitchRoutingMulticastItemToBody(networksSwitchRoutingMulticast, response1)
		diags = resp.State.Set(ctx, &networksSwitchRoutingMulticast)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchRoutingMulticast struct {
	NetworkID types.String                                    `tfsdk:"network_id"`
	Item      *ResponseSwitchGetNetworkSwitchRoutingMulticast `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchRoutingMulticast struct {
	DefaultSettings *ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings `tfsdk:"default_settings"`
	Overrides       *[]ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides     `tfsdk:"overrides"`
}

type ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings struct {
	FloodUnknownMulticastTrafficEnabled types.Bool `tfsdk:"flood_unknown_multicast_traffic_enabled"`
	IgmpSnoopingEnabled                 types.Bool `tfsdk:"igmp_snooping_enabled"`
}

type ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides struct {
	FloodUnknownMulticastTrafficEnabled types.Bool `tfsdk:"flood_unknown_multicast_traffic_enabled"`
	IgmpSnoopingEnabled                 types.Bool `tfsdk:"igmp_snooping_enabled"`
	Switches                            types.List `tfsdk:"switches"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchRoutingMulticastItemToBody(state NetworksSwitchRoutingMulticast, response *merakigosdk.ResponseSwitchGetNetworkSwitchRoutingMulticast) NetworksSwitchRoutingMulticast {
	itemState := ResponseSwitchGetNetworkSwitchRoutingMulticast{
		DefaultSettings: func() *ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings {
			if response.DefaultSettings != nil {
				return &ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings{
					FloodUnknownMulticastTrafficEnabled: func() types.Bool {
						if response.DefaultSettings.FloodUnknownMulticastTrafficEnabled != nil {
							return types.BoolValue(*response.DefaultSettings.FloodUnknownMulticastTrafficEnabled)
						}
						return types.Bool{}
					}(),
					IgmpSnoopingEnabled: func() types.Bool {
						if response.DefaultSettings.IgmpSnoopingEnabled != nil {
							return types.BoolValue(*response.DefaultSettings.IgmpSnoopingEnabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseSwitchGetNetworkSwitchRoutingMulticastDefaultSettings{}
		}(),
		Overrides: func() *[]ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides {
			if response.Overrides != nil {
				result := make([]ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides, len(*response.Overrides))
				for i, overrides := range *response.Overrides {
					result[i] = ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides{
						FloodUnknownMulticastTrafficEnabled: func() types.Bool {
							if overrides.FloodUnknownMulticastTrafficEnabled != nil {
								return types.BoolValue(*overrides.FloodUnknownMulticastTrafficEnabled)
							}
							return types.Bool{}
						}(),
						IgmpSnoopingEnabled: func() types.Bool {
							if overrides.IgmpSnoopingEnabled != nil {
								return types.BoolValue(*overrides.IgmpSnoopingEnabled)
							}
							return types.Bool{}
						}(),
						Switches: StringSliceToList(overrides.Switches),
					}
				}
				return &result
			}
			return &[]ResponseSwitchGetNetworkSwitchRoutingMulticastOverrides{}
		}(),
	}
	state.Item = &itemState
	return state
}
