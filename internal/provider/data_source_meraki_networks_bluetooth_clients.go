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
	_ datasource.DataSource              = &NetworksBluetoothClientsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksBluetoothClientsDataSource{}
)

func NewNetworksBluetoothClientsDataSource() datasource.DataSource {
	return &NetworksBluetoothClientsDataSource{}
}

type NetworksBluetoothClientsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksBluetoothClientsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksBluetoothClientsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_bluetooth_clients"
}

func (d *NetworksBluetoothClientsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bluetooth_client_id": schema.StringAttribute{
				MarkdownDescription: `bluetoothClientId path parameter. Bluetooth client ID`,
				Required:            true,
			},
			"connectivity_history_timespan": schema.Int64Attribute{
				MarkdownDescription: `connectivityHistoryTimespan query parameter. The timespan, in seconds, for the connectivityHistory data. By default 1 day, 86400, will be used.`,
				Optional:            true,
			},
			"include_connectivity_history": schema.BoolAttribute{
				MarkdownDescription: `includeConnectivityHistory query parameter. Include the connectivity history for this client`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"device_name": schema.StringAttribute{
						Computed: true,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"in_sight_alert": schema.BoolAttribute{
						Computed: true,
					},
					"last_seen": schema.Int64Attribute{
						Computed: true,
					},
					"mac": schema.StringAttribute{
						Computed: true,
					},
					"manufacturer": schema.StringAttribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"network_id": schema.StringAttribute{
						Computed: true,
					},
					"out_of_sight_alert": schema.BoolAttribute{
						Computed: true,
					},
					"seen_by_device_mac": schema.StringAttribute{
						Computed: true,
					},
					"tags": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func (d *NetworksBluetoothClientsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksBluetoothClients NetworksBluetoothClients
	diags := req.Config.Get(ctx, &networksBluetoothClients)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkBluetoothClient")
		vvNetworkID := networksBluetoothClients.NetworkID.ValueString()
		vvBluetoothClientID := networksBluetoothClients.BluetoothClientID.ValueString()
		queryParams1 := merakigosdk.GetNetworkBluetoothClientQueryParams{}

		queryParams1.IncludeConnectivityHistory = networksBluetoothClients.IncludeConnectivityHistory.ValueBool()
		queryParams1.ConnectivityHistoryTimespan = int(networksBluetoothClients.ConnectivityHistoryTimespan.ValueInt64())

		response1, restyResp1, err := d.client.Networks.GetNetworkBluetoothClient(vvNetworkID, vvBluetoothClientID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkBluetoothClient",
				err.Error(),
			)
			return
		}

		networksBluetoothClients = ResponseNetworksGetNetworkBluetoothClientItemToBody(networksBluetoothClients, response1)
		diags = resp.State.Set(ctx, &networksBluetoothClients)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksBluetoothClients struct {
	NetworkID                   types.String                               `tfsdk:"network_id"`
	BluetoothClientID           types.String                               `tfsdk:"bluetooth_client_id"`
	IncludeConnectivityHistory  types.Bool                                 `tfsdk:"include_connectivity_history"`
	ConnectivityHistoryTimespan types.Int64                                `tfsdk:"connectivity_history_timespan"`
	Item                        *ResponseNetworksGetNetworkBluetoothClient `tfsdk:"item"`
}

type ResponseNetworksGetNetworkBluetoothClient struct {
	DeviceName      types.String `tfsdk:"device_name"`
	ID              types.String `tfsdk:"id"`
	InSightAlert    types.Bool   `tfsdk:"in_sight_alert"`
	LastSeen        types.Int64  `tfsdk:"last_seen"`
	Mac             types.String `tfsdk:"mac"`
	Manufacturer    types.String `tfsdk:"manufacturer"`
	Name            types.String `tfsdk:"name"`
	NetworkID       types.String `tfsdk:"network_id"`
	OutOfSightAlert types.Bool   `tfsdk:"out_of_sight_alert"`
	SeenByDeviceMac types.String `tfsdk:"seen_by_device_mac"`
	Tags            types.List   `tfsdk:"tags"`
}

// ToBody
func ResponseNetworksGetNetworkBluetoothClientItemToBody(state NetworksBluetoothClients, response *merakigosdk.ResponseNetworksGetNetworkBluetoothClient) NetworksBluetoothClients {
	itemState := ResponseNetworksGetNetworkBluetoothClient{
		DeviceName: types.StringValue(response.DeviceName),
		ID:         types.StringValue(response.ID),
		InSightAlert: func() types.Bool {
			if response.InSightAlert != nil {
				return types.BoolValue(*response.InSightAlert)
			}
			return types.Bool{}
		}(),
		LastSeen: func() types.Int64 {
			if response.LastSeen != nil {
				return types.Int64Value(int64(*response.LastSeen))
			}
			return types.Int64{}
		}(),
		Mac:          types.StringValue(response.Mac),
		Manufacturer: types.StringValue(response.Manufacturer),
		Name:         types.StringValue(response.Name),
		NetworkID:    types.StringValue(response.NetworkID),
		OutOfSightAlert: func() types.Bool {
			if response.OutOfSightAlert != nil {
				return types.BoolValue(*response.OutOfSightAlert)
			}
			return types.Bool{}
		}(),
		SeenByDeviceMac: types.StringValue(response.SeenByDeviceMac),
		Tags:            StringSliceToList(response.Tags),
	}
	state.Item = &itemState
	return state
}
