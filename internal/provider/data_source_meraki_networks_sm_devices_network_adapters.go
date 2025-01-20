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
	_ datasource.DataSource              = &NetworksSmDevicesNetworkAdaptersDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmDevicesNetworkAdaptersDataSource{}
)

func NewNetworksSmDevicesNetworkAdaptersDataSource() datasource.DataSource {
	return &NetworksSmDevicesNetworkAdaptersDataSource{}
}

type NetworksSmDevicesNetworkAdaptersDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmDevicesNetworkAdaptersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmDevicesNetworkAdaptersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_network_adapters"
}

func (d *NetworksSmDevicesNetworkAdaptersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `Array of ResponseSmGetNetworkSmDeviceNetworkAdapters`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"dhcp_server": schema.StringAttribute{
							MarkdownDescription: `The IP address of the DCHP Server.`,
							Computed:            true,
						},
						"dns_server": schema.StringAttribute{
							MarkdownDescription: `The IP address of the DNS Server.`,
							Computed:            true,
						},
						"gateway": schema.StringAttribute{
							MarkdownDescription: `The IP address of the Gateway.`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `The Meraki Id of the network adapter record.`,
							Computed:            true,
						},
						"ip": schema.StringAttribute{
							MarkdownDescription: `The IP address of the network adapter.`,
							Computed:            true,
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `The MAC associated with the network adapter.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the newtwork adapter.`,
							Computed:            true,
						},
						"subnet": schema.StringAttribute{
							MarkdownDescription: `The subnet for the network adapter.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmDevicesNetworkAdaptersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmDevicesNetworkAdapters NetworksSmDevicesNetworkAdapters
	diags := req.Config.Get(ctx, &networksSmDevicesNetworkAdapters)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmDeviceNetworkAdapters")
		vvNetworkID := networksSmDevicesNetworkAdapters.NetworkID.ValueString()
		vvDeviceID := networksSmDevicesNetworkAdapters.DeviceID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetNetworkSmDeviceNetworkAdapters(vvNetworkID, vvDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmDeviceNetworkAdapters",
				err.Error(),
			)
			return
		}

		networksSmDevicesNetworkAdapters = ResponseSmGetNetworkSmDeviceNetworkAdaptersItemsToBody(networksSmDevicesNetworkAdapters, response1)
		diags = resp.State.Set(ctx, &networksSmDevicesNetworkAdapters)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmDevicesNetworkAdapters struct {
	NetworkID types.String                                       `tfsdk:"network_id"`
	DeviceID  types.String                                       `tfsdk:"device_id"`
	Items     *[]ResponseItemSmGetNetworkSmDeviceNetworkAdapters `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmDeviceNetworkAdapters struct {
	DhcpServer types.String `tfsdk:"dhcp_server"`
	DNSServer  types.String `tfsdk:"dns_server"`
	Gateway    types.String `tfsdk:"gateway"`
	ID         types.String `tfsdk:"id"`
	IP         types.String `tfsdk:"ip"`
	Mac        types.String `tfsdk:"mac"`
	Name       types.String `tfsdk:"name"`
	Subnet     types.String `tfsdk:"subnet"`
}

// ToBody
func ResponseSmGetNetworkSmDeviceNetworkAdaptersItemsToBody(state NetworksSmDevicesNetworkAdapters, response *merakigosdk.ResponseSmGetNetworkSmDeviceNetworkAdapters) NetworksSmDevicesNetworkAdapters {
	var items []ResponseItemSmGetNetworkSmDeviceNetworkAdapters
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmDeviceNetworkAdapters{
			DhcpServer: types.StringValue(item.DhcpServer),
			DNSServer:  types.StringValue(item.DNSServer),
			Gateway:    types.StringValue(item.Gateway),
			ID:         types.StringValue(item.ID),
			IP:         types.StringValue(item.IP),
			Mac:        types.StringValue(item.Mac),
			Name:       types.StringValue(item.Name),
			Subnet:     types.StringValue(item.Subnet),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
