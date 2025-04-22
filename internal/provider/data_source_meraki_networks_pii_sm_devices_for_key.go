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
	_ datasource.DataSource              = &NetworksPiiSmDevicesForKeyDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksPiiSmDevicesForKeyDataSource{}
)

func NewNetworksPiiSmDevicesForKeyDataSource() datasource.DataSource {
	return &NetworksPiiSmDevicesForKeyDataSource{}
}

type NetworksPiiSmDevicesForKeyDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksPiiSmDevicesForKeyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksPiiSmDevicesForKeyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_pii_sm_devices_for_key"
}

func (d *NetworksPiiSmDevicesForKeyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bluetooth_mac": schema.StringAttribute{
				MarkdownDescription: `bluetoothMac query parameter. The MAC of a Bluetooth client`,
				Optional:            true,
			},
			"email": schema.StringAttribute{
				MarkdownDescription: `email query parameter. The email of a network user account or a Systems Manager device`,
				Optional:            true,
			},
			"imei": schema.StringAttribute{
				MarkdownDescription: `imei query parameter. The IMEI of a Systems Manager device`,
				Optional:            true,
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `mac query parameter. The MAC of a network client device or a Systems Manager device`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial query parameter. The serial of a Systems Manager device`,
				Optional:            true,
			},
			"username": schema.StringAttribute{
				MarkdownDescription: `username query parameter. The username of a Systems Manager user`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"n_1234": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *NetworksPiiSmDevicesForKeyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksPiiSmDevicesForKey NetworksPiiSmDevicesForKey
	diags := req.Config.Get(ctx, &networksPiiSmDevicesForKey)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkPiiSmDevicesForKey")
		vvNetworkID := networksPiiSmDevicesForKey.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkPiiSmDevicesForKeyQueryParams{}

		queryParams1.Username = networksPiiSmDevicesForKey.Username.ValueString()
		queryParams1.Email = networksPiiSmDevicesForKey.Email.ValueString()
		queryParams1.Mac = networksPiiSmDevicesForKey.Mac.ValueString()
		queryParams1.Serial = networksPiiSmDevicesForKey.Serial.ValueString()
		queryParams1.Imei = networksPiiSmDevicesForKey.Imei.ValueString()
		queryParams1.BluetoothMac = networksPiiSmDevicesForKey.BluetoothMac.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkPiiSmDevicesForKey(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkPiiSmDevicesForKey",
				err.Error(),
			)
			return
		}

		networksPiiSmDevicesForKey = ResponseNetworksGetNetworkPiiSmDevicesForKeyItemToBody(networksPiiSmDevicesForKey, response1)
		diags = resp.State.Set(ctx, &networksPiiSmDevicesForKey)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksPiiSmDevicesForKey struct {
	NetworkID    types.String                                  `tfsdk:"network_id"`
	Username     types.String                                  `tfsdk:"username"`
	Email        types.String                                  `tfsdk:"email"`
	Mac          types.String                                  `tfsdk:"mac"`
	Serial       types.String                                  `tfsdk:"serial"`
	Imei         types.String                                  `tfsdk:"imei"`
	BluetoothMac types.String                                  `tfsdk:"bluetooth_mac"`
	Item         *ResponseNetworksGetNetworkPiiSmDevicesForKey `tfsdk:"item"`
}

type ResponseNetworksGetNetworkPiiSmDevicesForKey struct {
	N1234 types.List `tfsdk:"n_1234"`
}

// ToBody
func ResponseNetworksGetNetworkPiiSmDevicesForKeyItemToBody(state NetworksPiiSmDevicesForKey, response *merakigosdk.ResponseNetworksGetNetworkPiiSmDevicesForKey) NetworksPiiSmDevicesForKey {
	itemState := ResponseNetworksGetNetworkPiiSmDevicesForKey{
		N1234: StringSliceToList(response.N1234),
	}
	state.Item = &itemState
	return state
}
