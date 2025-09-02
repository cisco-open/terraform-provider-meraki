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
	_ datasource.DataSource              = &NetworksSmDevicesWLANListsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmDevicesWLANListsDataSource{}
)

func NewNetworksSmDevicesWLANListsDataSource() datasource.DataSource {
	return &NetworksSmDevicesWLANListsDataSource{}
}

type NetworksSmDevicesWLANListsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmDevicesWLANListsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmDevicesWLANListsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_wlan_lists"
}

func (d *NetworksSmDevicesWLANListsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `Array of ResponseSmGetNetworkSmDeviceWlanLists`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"created_at": schema.StringAttribute{
							MarkdownDescription: `When the Meraki record for the wlanList was created.`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `The Meraki managed Id of the wlanList record.`,
							Computed:            true,
						},
						"xml": schema.StringAttribute{
							MarkdownDescription: `An XML string containing the WLAN List for the device.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmDevicesWLANListsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmDevicesWLANLists NetworksSmDevicesWLANLists
	diags := req.Config.Get(ctx, &networksSmDevicesWLANLists)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmDeviceWLANLists")
		vvNetworkID := networksSmDevicesWLANLists.NetworkID.ValueString()
		vvDeviceID := networksSmDevicesWLANLists.DeviceID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetNetworkSmDeviceWLANLists(vvNetworkID, vvDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmDeviceWLANLists",
				err.Error(),
			)
			return
		}

		networksSmDevicesWLANLists = ResponseSmGetNetworkSmDeviceWLANListsItemsToBody(networksSmDevicesWLANLists, response1)
		diags = resp.State.Set(ctx, &networksSmDevicesWLANLists)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmDevicesWLANLists struct {
	NetworkID types.String                                 `tfsdk:"network_id"`
	DeviceID  types.String                                 `tfsdk:"device_id"`
	Items     *[]ResponseItemSmGetNetworkSmDeviceWlanLists `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmDeviceWlanLists struct {
	CreatedAt types.String `tfsdk:"created_at"`
	ID        types.String `tfsdk:"id"`
	Xml       types.String `tfsdk:"xml"`
}

// ToBody
func ResponseSmGetNetworkSmDeviceWLANListsItemsToBody(state NetworksSmDevicesWLANLists, response *merakigosdk.ResponseSmGetNetworkSmDeviceWLANLists) NetworksSmDevicesWLANLists {
	var items []ResponseItemSmGetNetworkSmDeviceWlanLists
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmDeviceWlanLists{
			CreatedAt: func() types.String {
				if item.CreatedAt != "" {
					return types.StringValue(item.CreatedAt)
				}
				return types.String{}
			}(),
			ID: func() types.String {
				if item.ID != "" {
					return types.StringValue(item.ID)
				}
				return types.String{}
			}(),
			Xml: func() types.String {
				if item.Xml != "" {
					return types.StringValue(item.Xml)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
