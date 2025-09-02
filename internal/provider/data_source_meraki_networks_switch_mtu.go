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
	_ datasource.DataSource              = &NetworksSwitchMtuDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchMtuDataSource{}
)

func NewNetworksSwitchMtuDataSource() datasource.DataSource {
	return &NetworksSwitchMtuDataSource{}
}

type NetworksSwitchMtuDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchMtuDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchMtuDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_mtu"
}

func (d *NetworksSwitchMtuDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"default_mtu_size": schema.Int64Attribute{
						MarkdownDescription: `MTU size for the entire network. Default value is 9578.`,
						Computed:            true,
					},
					"overrides": schema.SetNestedAttribute{
						MarkdownDescription: `Override MTU size for individual switches or switch templates.
      An empty array will clear overrides.`,
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"mtu_size": schema.Int64Attribute{
									MarkdownDescription: `MTU size for the switches or switch templates.`,
									Computed:            true,
								},
								"switch_profiles": schema.ListAttribute{
									MarkdownDescription: `List of switch template IDs. Applicable only for template network.`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"switches": schema.ListAttribute{
									MarkdownDescription: `List of switch serials. Applicable only for switch network.`,
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchMtuDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchMtu NetworksSwitchMtu
	diags := req.Config.Get(ctx, &networksSwitchMtu)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchMtu")
		vvNetworkID := networksSwitchMtu.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchMtu(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchMtu",
				err.Error(),
			)
			return
		}

		networksSwitchMtu = ResponseSwitchGetNetworkSwitchMtuItemToBody(networksSwitchMtu, response1)
		diags = resp.State.Set(ctx, &networksSwitchMtu)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchMtu struct {
	NetworkID types.String                       `tfsdk:"network_id"`
	Item      *ResponseSwitchGetNetworkSwitchMtu `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchMtu struct {
	DefaultMtuSize types.Int64                                   `tfsdk:"default_mtu_size"`
	Overrides      *[]ResponseSwitchGetNetworkSwitchMtuOverrides `tfsdk:"overrides"`
}

type ResponseSwitchGetNetworkSwitchMtuOverrides struct {
	MtuSize        types.Int64 `tfsdk:"mtu_size"`
	SwitchProfiles types.List  `tfsdk:"switch_profiles"`
	Switches       types.List  `tfsdk:"switches"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchMtuItemToBody(state NetworksSwitchMtu, response *merakigosdk.ResponseSwitchGetNetworkSwitchMtu) NetworksSwitchMtu {
	itemState := ResponseSwitchGetNetworkSwitchMtu{
		DefaultMtuSize: func() types.Int64 {
			if response.DefaultMtuSize != nil {
				return types.Int64Value(int64(*response.DefaultMtuSize))
			}
			return types.Int64{}
		}(),
		Overrides: func() *[]ResponseSwitchGetNetworkSwitchMtuOverrides {
			if response.Overrides != nil {
				result := make([]ResponseSwitchGetNetworkSwitchMtuOverrides, len(*response.Overrides))
				for i, overrides := range *response.Overrides {
					result[i] = ResponseSwitchGetNetworkSwitchMtuOverrides{
						MtuSize: func() types.Int64 {
							if overrides.MtuSize != nil {
								return types.Int64Value(int64(*overrides.MtuSize))
							}
							return types.Int64{}
						}(),
						SwitchProfiles: StringSliceToList(overrides.SwitchProfiles),
						Switches:       StringSliceToList(overrides.Switches),
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
