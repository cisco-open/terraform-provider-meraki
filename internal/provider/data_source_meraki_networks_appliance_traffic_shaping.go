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
	_ datasource.DataSource              = &NetworksApplianceTrafficShapingDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceTrafficShapingDataSource{}
)

func NewNetworksApplianceTrafficShapingDataSource() datasource.DataSource {
	return &NetworksApplianceTrafficShapingDataSource{}
}

type NetworksApplianceTrafficShapingDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceTrafficShapingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceTrafficShapingDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_traffic_shaping"
}

func (d *NetworksApplianceTrafficShapingDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"global_bandwidth_limits": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"limit_down": schema.Int64Attribute{
								Computed: true,
							},
							"limit_up": schema.Int64Attribute{
								Computed: true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceTrafficShapingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceTrafficShaping NetworksApplianceTrafficShaping
	diags := req.Config.Get(ctx, &networksApplianceTrafficShaping)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceTrafficShaping")
		vvNetworkID := networksApplianceTrafficShaping.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceTrafficShaping(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceTrafficShaping",
				err.Error(),
			)
			return
		}

		networksApplianceTrafficShaping = ResponseApplianceGetNetworkApplianceTrafficShapingItemToBody(networksApplianceTrafficShaping, response1)
		diags = resp.State.Set(ctx, &networksApplianceTrafficShaping)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceTrafficShaping struct {
	NetworkID types.String                                        `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceTrafficShaping `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceTrafficShaping struct {
	GlobalBandwidthLimits *ResponseApplianceGetNetworkApplianceTrafficShapingGlobalBandwidthLimits `tfsdk:"global_bandwidth_limits"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingGlobalBandwidthLimits struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceTrafficShapingItemToBody(state NetworksApplianceTrafficShaping, response *merakigosdk.ResponseApplianceGetNetworkApplianceTrafficShaping) NetworksApplianceTrafficShaping {
	itemState := ResponseApplianceGetNetworkApplianceTrafficShaping{
		GlobalBandwidthLimits: func() *ResponseApplianceGetNetworkApplianceTrafficShapingGlobalBandwidthLimits {
			if response.GlobalBandwidthLimits != nil {
				return &ResponseApplianceGetNetworkApplianceTrafficShapingGlobalBandwidthLimits{
					LimitDown: func() types.Int64 {
						if response.GlobalBandwidthLimits.LimitDown != nil {
							return types.Int64Value(int64(*response.GlobalBandwidthLimits.LimitDown))
						}
						return types.Int64{}
					}(),
					LimitUp: func() types.Int64 {
						if response.GlobalBandwidthLimits.LimitUp != nil {
							return types.Int64Value(int64(*response.GlobalBandwidthLimits.LimitUp))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
