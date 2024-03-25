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
	_ datasource.DataSource              = &NetworksSwitchStormControlDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchStormControlDataSource{}
)

func NewNetworksSwitchStormControlDataSource() datasource.DataSource {
	return &NetworksSwitchStormControlDataSource{}
}

type NetworksSwitchStormControlDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchStormControlDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchStormControlDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_storm_control"
}

func (d *NetworksSwitchStormControlDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"broadcast_threshold": schema.Int64Attribute{
						MarkdownDescription: `Broadcast threshold.`,
						Computed:            true,
					},
					"multicast_threshold": schema.Int64Attribute{
						MarkdownDescription: `Multicast threshold.`,
						Computed:            true,
					},
					"unknown_unicast_threshold": schema.Int64Attribute{
						MarkdownDescription: `Unknown Unicast threshold.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchStormControlDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchStormControl NetworksSwitchStormControl
	diags := req.Config.Get(ctx, &networksSwitchStormControl)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchStormControl")
		vvNetworkID := networksSwitchStormControl.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchStormControl(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStormControl",
				err.Error(),
			)
			return
		}

		networksSwitchStormControl = ResponseSwitchGetNetworkSwitchStormControlItemToBody(networksSwitchStormControl, response1)
		diags = resp.State.Set(ctx, &networksSwitchStormControl)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchStormControl struct {
	NetworkID types.String                                `tfsdk:"network_id"`
	Item      *ResponseSwitchGetNetworkSwitchStormControl `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchStormControl struct {
	BroadcastThreshold      types.Int64 `tfsdk:"broadcast_threshold"`
	MulticastThreshold      types.Int64 `tfsdk:"multicast_threshold"`
	UnknownUnicastThreshold types.Int64 `tfsdk:"unknown_unicast_threshold"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchStormControlItemToBody(state NetworksSwitchStormControl, response *merakigosdk.ResponseSwitchGetNetworkSwitchStormControl) NetworksSwitchStormControl {
	itemState := ResponseSwitchGetNetworkSwitchStormControl{
		BroadcastThreshold: func() types.Int64 {
			if response.BroadcastThreshold != nil {
				return types.Int64Value(int64(*response.BroadcastThreshold))
			}
			return types.Int64{}
		}(),
		MulticastThreshold: func() types.Int64 {
			if response.MulticastThreshold != nil {
				return types.Int64Value(int64(*response.MulticastThreshold))
			}
			return types.Int64{}
		}(),
		UnknownUnicastThreshold: func() types.Int64 {
			if response.UnknownUnicastThreshold != nil {
				return types.Int64Value(int64(*response.UnknownUnicastThreshold))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
