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
	_ datasource.DataSource              = &NetworksPiiPiiKeysDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksPiiPiiKeysDataSource{}
)

func NewNetworksPiiPiiKeysDataSource() datasource.DataSource {
	return &NetworksPiiPiiKeysDataSource{}
}

type NetworksPiiPiiKeysDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksPiiPiiKeysDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksPiiPiiKeysDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_pii_pii_keys"
}

func (d *NetworksPiiPiiKeysDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

					"n_1234": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"bluetooth_macs": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"emails": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"imeis": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"macs": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"serials": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"usernames": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksPiiPiiKeysDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksPiiPiiKeys NetworksPiiPiiKeys
	diags := req.Config.Get(ctx, &networksPiiPiiKeys)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkPiiPiiKeys")
		vvNetworkID := networksPiiPiiKeys.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkPiiPiiKeysQueryParams{}

		queryParams1.Username = networksPiiPiiKeys.Username.ValueString()
		queryParams1.Email = networksPiiPiiKeys.Email.ValueString()
		queryParams1.Mac = networksPiiPiiKeys.Mac.ValueString()
		queryParams1.Serial = networksPiiPiiKeys.Serial.ValueString()
		queryParams1.Imei = networksPiiPiiKeys.Imei.ValueString()
		queryParams1.BluetoothMac = networksPiiPiiKeys.BluetoothMac.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkPiiPiiKeys(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkPiiPiiKeys",
				err.Error(),
			)
			return
		}

		networksPiiPiiKeys = ResponseNetworksGetNetworkPiiPiiKeysItemToBody(networksPiiPiiKeys, response1)
		diags = resp.State.Set(ctx, &networksPiiPiiKeys)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksPiiPiiKeys struct {
	NetworkID    types.String                          `tfsdk:"network_id"`
	Username     types.String                          `tfsdk:"username"`
	Email        types.String                          `tfsdk:"email"`
	Mac          types.String                          `tfsdk:"mac"`
	Serial       types.String                          `tfsdk:"serial"`
	Imei         types.String                          `tfsdk:"imei"`
	BluetoothMac types.String                          `tfsdk:"bluetooth_mac"`
	Item         *ResponseNetworksGetNetworkPiiPiiKeys `tfsdk:"item"`
}

type ResponseNetworksGetNetworkPiiPiiKeys struct {
	N1234 *ResponseNetworksGetNetworkPiiPiiKeysN1234 `tfsdk:"n_1234"`
}

type ResponseNetworksGetNetworkPiiPiiKeysN1234 struct {
	BluetoothMacs types.List `tfsdk:"bluetooth_macs"`
	Emails        types.List `tfsdk:"emails"`
	Imeis         types.List `tfsdk:"imeis"`
	Macs          types.List `tfsdk:"macs"`
	Serials       types.List `tfsdk:"serials"`
	Usernames     types.List `tfsdk:"usernames"`
}

// ToBody
func ResponseNetworksGetNetworkPiiPiiKeysItemToBody(state NetworksPiiPiiKeys, response *merakigosdk.ResponseNetworksGetNetworkPiiPiiKeys) NetworksPiiPiiKeys {
	itemState := ResponseNetworksGetNetworkPiiPiiKeys{
		N1234: func() *ResponseNetworksGetNetworkPiiPiiKeysN1234 {
			if response.N1234 != nil {
				return &ResponseNetworksGetNetworkPiiPiiKeysN1234{
					BluetoothMacs: StringSliceToList(response.N1234.BluetoothMacs),
					Emails:        StringSliceToList(response.N1234.Emails),
					Imeis:         StringSliceToList(response.N1234.Imeis),
					Macs:          StringSliceToList(response.N1234.Macs),
					Serials:       StringSliceToList(response.N1234.Serials),
					Usernames:     StringSliceToList(response.N1234.Usernames),
				}
			}
			return &ResponseNetworksGetNetworkPiiPiiKeysN1234{}
		}(),
	}
	state.Item = &itemState
	return state
}
