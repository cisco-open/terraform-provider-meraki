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
	_ datasource.DataSource              = &NetworksSmDevicesCellularUsageHistoryDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmDevicesCellularUsageHistoryDataSource{}
)

func NewNetworksSmDevicesCellularUsageHistoryDataSource() datasource.DataSource {
	return &NetworksSmDevicesCellularUsageHistoryDataSource{}
}

type NetworksSmDevicesCellularUsageHistoryDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmDevicesCellularUsageHistoryDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmDevicesCellularUsageHistoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_cellular_usage_history"
}

func (d *NetworksSmDevicesCellularUsageHistoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_id": schema.StringAttribute{
				MarkdownDescription: `deviceId path parameter. Device ID`,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetNetworkSmDeviceCellularUsageHistory`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"received": schema.Float64Attribute{
							MarkdownDescription: `The amount of cellular data received by the device.`,
							Computed:            true,
						},
						"sent": schema.Float64Attribute{
							MarkdownDescription: `The amount of cellular sent received by the device.`,
							Computed:            true,
						},
						"ts": schema.StringAttribute{
							MarkdownDescription: `When the cellular usage data was collected.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmDevicesCellularUsageHistoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmDevicesCellularUsageHistory NetworksSmDevicesCellularUsageHistory
	diags := req.Config.Get(ctx, &networksSmDevicesCellularUsageHistory)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmDeviceCellularUsageHistory")
		vvNetworkID := networksSmDevicesCellularUsageHistory.NetworkID.ValueString()
		vvDeviceID := networksSmDevicesCellularUsageHistory.DeviceID.ValueString()

		response1, restyResp1, err := d.client.Sm.GetNetworkSmDeviceCellularUsageHistory(vvNetworkID, vvDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmDeviceCellularUsageHistory",
				err.Error(),
			)
			return
		}

		networksSmDevicesCellularUsageHistory = ResponseSmGetNetworkSmDeviceCellularUsageHistoryItemsToBody(networksSmDevicesCellularUsageHistory, response1)
		diags = resp.State.Set(ctx, &networksSmDevicesCellularUsageHistory)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmDevicesCellularUsageHistory struct {
	NetworkID types.String                                            `tfsdk:"network_id"`
	DeviceID  types.String                                            `tfsdk:"device_id"`
	Items     *[]ResponseItemSmGetNetworkSmDeviceCellularUsageHistory `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmDeviceCellularUsageHistory struct {
	Received types.Float64 `tfsdk:"received"`
	Sent     types.Float64 `tfsdk:"sent"`
	Ts       types.String  `tfsdk:"ts"`
}

// ToBody
func ResponseSmGetNetworkSmDeviceCellularUsageHistoryItemsToBody(state NetworksSmDevicesCellularUsageHistory, response *merakigosdk.ResponseSmGetNetworkSmDeviceCellularUsageHistory) NetworksSmDevicesCellularUsageHistory {
	var items []ResponseItemSmGetNetworkSmDeviceCellularUsageHistory
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmDeviceCellularUsageHistory{
			Received: func() types.Float64 {
				if item.Received != nil {
					return types.Float64Value(float64(*item.Received))
				}
				return types.Float64{}
			}(),
			Sent: func() types.Float64 {
				if item.Sent != nil {
					return types.Float64Value(float64(*item.Sent))
				}
				return types.Float64{}
			}(),
			Ts: types.StringValue(item.Ts),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
