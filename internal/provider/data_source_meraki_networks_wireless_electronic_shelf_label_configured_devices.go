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
	_ datasource.DataSource              = &NetworksWirelessElectronicShelfLabelConfiguredDevicesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessElectronicShelfLabelConfiguredDevicesDataSource{}
)

func NewNetworksWirelessElectronicShelfLabelConfiguredDevicesDataSource() datasource.DataSource {
	return &NetworksWirelessElectronicShelfLabelConfiguredDevicesDataSource{}
}

type NetworksWirelessElectronicShelfLabelConfiguredDevicesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessElectronicShelfLabelConfiguredDevicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessElectronicShelfLabelConfiguredDevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_electronic_shelf_label_configured_devices"
}

func (d *NetworksWirelessElectronicShelfLabelConfiguredDevicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetNetworkWirelessElectronicShelfLabelConfiguredDevices`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"enabled": schema.BoolAttribute{
							MarkdownDescription: `Turn ESL features on and off for this network`,
							Computed:            true,
						},
						"hostname": schema.StringAttribute{
							MarkdownDescription: `Desired ESL hostname of the network`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessElectronicShelfLabelConfiguredDevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessElectronicShelfLabelConfiguredDevices NetworksWirelessElectronicShelfLabelConfiguredDevices
	diags := req.Config.Get(ctx, &networksWirelessElectronicShelfLabelConfiguredDevices)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessElectronicShelfLabelConfiguredDevices")
		vvNetworkID := networksWirelessElectronicShelfLabelConfiguredDevices.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessElectronicShelfLabelConfiguredDevices(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessElectronicShelfLabelConfiguredDevices",
				err.Error(),
			)
			return
		}

		networksWirelessElectronicShelfLabelConfiguredDevices = ResponseWirelessGetNetworkWirelessElectronicShelfLabelConfiguredDevicesItemsToBody(networksWirelessElectronicShelfLabelConfiguredDevices, response1)
		diags = resp.State.Set(ctx, &networksWirelessElectronicShelfLabelConfiguredDevices)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessElectronicShelfLabelConfiguredDevices struct {
	NetworkID types.String                                                                   `tfsdk:"network_id"`
	Items     *[]ResponseItemWirelessGetNetworkWirelessElectronicShelfLabelConfiguredDevices `tfsdk:"items"`
}

type ResponseItemWirelessGetNetworkWirelessElectronicShelfLabelConfiguredDevices struct {
	Enabled  types.Bool   `tfsdk:"enabled"`
	Hostname types.String `tfsdk:"hostname"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessElectronicShelfLabelConfiguredDevicesItemsToBody(state NetworksWirelessElectronicShelfLabelConfiguredDevices, response *merakigosdk.ResponseWirelessGetNetworkWirelessElectronicShelfLabelConfiguredDevices) NetworksWirelessElectronicShelfLabelConfiguredDevices {
	var items []ResponseItemWirelessGetNetworkWirelessElectronicShelfLabelConfiguredDevices
	for _, item := range *response {
		itemState := ResponseItemWirelessGetNetworkWirelessElectronicShelfLabelConfiguredDevices{
			Enabled: func() types.Bool {
				if item.Enabled != nil {
					return types.BoolValue(*item.Enabled)
				}
				return types.Bool{}
			}(),
			Hostname: types.StringValue(item.Hostname),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
