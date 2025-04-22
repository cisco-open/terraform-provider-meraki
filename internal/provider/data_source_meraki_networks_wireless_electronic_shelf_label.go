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
	_ datasource.DataSource              = &NetworksWirelessElectronicShelfLabelDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessElectronicShelfLabelDataSource{}
)

func NewNetworksWirelessElectronicShelfLabelDataSource() datasource.DataSource {
	return &NetworksWirelessElectronicShelfLabelDataSource{}
}

type NetworksWirelessElectronicShelfLabelDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessElectronicShelfLabelDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessElectronicShelfLabelDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_electronic_shelf_label"
}

func (d *NetworksWirelessElectronicShelfLabelDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Turn ESL features on and off for this network`,
						Computed:            true,
					},
					"hostname": schema.StringAttribute{
						MarkdownDescription: `Desired ESL hostname of the network`,
						Computed:            true,
					},
					"mode": schema.StringAttribute{
						MarkdownDescription: `Electronic shelf label mode of the network. Valid options are 'Bluetooth', 'high frequency'`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessElectronicShelfLabelDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessElectronicShelfLabel NetworksWirelessElectronicShelfLabel
	diags := req.Config.Get(ctx, &networksWirelessElectronicShelfLabel)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessElectronicShelfLabel")
		vvNetworkID := networksWirelessElectronicShelfLabel.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessElectronicShelfLabel(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessElectronicShelfLabel",
				err.Error(),
			)
			return
		}

		networksWirelessElectronicShelfLabel = ResponseWirelessGetNetworkWirelessElectronicShelfLabelItemToBody(networksWirelessElectronicShelfLabel, response1)
		diags = resp.State.Set(ctx, &networksWirelessElectronicShelfLabel)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessElectronicShelfLabel struct {
	NetworkID types.String                                            `tfsdk:"network_id"`
	Item      *ResponseWirelessGetNetworkWirelessElectronicShelfLabel `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessElectronicShelfLabel struct {
	Enabled  types.Bool   `tfsdk:"enabled"`
	Hostname types.String `tfsdk:"hostname"`
	Mode     types.String `tfsdk:"mode"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessElectronicShelfLabelItemToBody(state NetworksWirelessElectronicShelfLabel, response *merakigosdk.ResponseWirelessGetNetworkWirelessElectronicShelfLabel) NetworksWirelessElectronicShelfLabel {
	itemState := ResponseWirelessGetNetworkWirelessElectronicShelfLabel{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Hostname: types.StringValue(response.Hostname),
		Mode:     types.StringValue(response.Mode),
	}
	state.Item = &itemState
	return state
}
