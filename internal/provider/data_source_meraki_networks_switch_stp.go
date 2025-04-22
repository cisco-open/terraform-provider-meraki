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
	_ datasource.DataSource              = &NetworksSwitchStpDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchStpDataSource{}
)

func NewNetworksSwitchStpDataSource() datasource.DataSource {
	return &NetworksSwitchStpDataSource{}
}

type NetworksSwitchStpDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchStpDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchStpDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stp"
}

func (d *NetworksSwitchStpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"rstp_enabled": schema.BoolAttribute{
						MarkdownDescription: `The spanning tree protocol status in network`,
						Computed:            true,
					},
					"stp_bridge_priority": schema.SetNestedAttribute{
						MarkdownDescription: `STP bridge priority for switches/stacks or switch templates. An empty array will clear the STP bridge priority settings.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"stacks": schema.ListAttribute{
									MarkdownDescription: `List of stack IDs`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"stp_priority": schema.Int64Attribute{
									MarkdownDescription: `STP priority for switch, stacks, or switch templates`,
									Computed:            true,
								},
								"switch_profiles": schema.ListAttribute{
									MarkdownDescription: `List of switch template IDs`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"switches": schema.ListAttribute{
									MarkdownDescription: `List of switch serial numbers`,
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

func (d *NetworksSwitchStpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchStp NetworksSwitchStp
	diags := req.Config.Get(ctx, &networksSwitchStp)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStp")
		vvNetworkID := networksSwitchStp.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchStp(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStp",
				err.Error(),
			)
			return
		}

		networksSwitchStp = ResponseSwitchGetNetworkSwitchStpItemToBody(networksSwitchStp, response1)
		diags = resp.State.Set(ctx, &networksSwitchStp)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchStp struct {
	NetworkID types.String                       `tfsdk:"network_id"`
	Item      *ResponseSwitchGetNetworkSwitchStp `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchStp struct {
	RstpEnabled       types.Bool                                            `tfsdk:"rstp_enabled"`
	StpBridgePriority *[]ResponseSwitchGetNetworkSwitchStpStpBridgePriority `tfsdk:"stp_bridge_priority"`
}

type ResponseSwitchGetNetworkSwitchStpStpBridgePriority struct {
	Stacks         types.List  `tfsdk:"stacks"`
	StpPriority    types.Int64 `tfsdk:"stp_priority"`
	SwitchProfiles types.List  `tfsdk:"switch_profiles"`
	Switches       types.List  `tfsdk:"switches"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchStpItemToBody(state NetworksSwitchStp, response *merakigosdk.ResponseSwitchGetNetworkSwitchStp) NetworksSwitchStp {
	itemState := ResponseSwitchGetNetworkSwitchStp{
		RstpEnabled: func() types.Bool {
			if response.RstpEnabled != nil {
				return types.BoolValue(*response.RstpEnabled)
			}
			return types.Bool{}
		}(),
		StpBridgePriority: func() *[]ResponseSwitchGetNetworkSwitchStpStpBridgePriority {
			if response.StpBridgePriority != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStpStpBridgePriority, len(*response.StpBridgePriority))
				for i, stpBridgePriority := range *response.StpBridgePriority {
					result[i] = ResponseSwitchGetNetworkSwitchStpStpBridgePriority{
						Stacks: StringSliceToList(stpBridgePriority.Stacks),
						StpPriority: func() types.Int64 {
							if stpBridgePriority.StpPriority != nil {
								return types.Int64Value(int64(*stpBridgePriority.StpPriority))
							}
							return types.Int64{}
						}(),
						SwitchProfiles: StringSliceToList(stpBridgePriority.SwitchProfiles),
						Switches:       StringSliceToList(stpBridgePriority.Switches),
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
