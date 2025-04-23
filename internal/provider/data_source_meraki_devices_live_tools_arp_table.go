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
	_ datasource.DataSource              = &DevicesLiveToolsArpTableDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesLiveToolsArpTableDataSource{}
)

func NewDevicesLiveToolsArpTableDataSource() datasource.DataSource {
	return &DevicesLiveToolsArpTableDataSource{}
}

type DevicesLiveToolsArpTableDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesLiveToolsArpTableDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesLiveToolsArpTableDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_live_tools_arp_table"
}

func (d *DevicesLiveToolsArpTableDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"arp_table_id": schema.StringAttribute{
				MarkdownDescription: `arpTableId path parameter. Arp table ID`,
				Required:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"arp_table_id": schema.StringAttribute{
						MarkdownDescription: `Id of the ARP table request. Used to check the status of the request.`,
						Computed:            true,
					},
					"entries": schema.SetNestedAttribute{
						MarkdownDescription: `The ARP table entries`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"ip": schema.StringAttribute{
									MarkdownDescription: `The IP address of the ARP table entry`,
									Computed:            true,
								},
								"last_updated_at": schema.StringAttribute{
									MarkdownDescription: `Time of the last update of the ARP table entry`,
									Computed:            true,
								},
								"mac": schema.StringAttribute{
									MarkdownDescription: `The MAC address of the ARP table entry`,
									Computed:            true,
								},
								"vlan_id": schema.Int64Attribute{
									MarkdownDescription: `The VLAN ID of the ARP table entry`,
									Computed:            true,
								},
							},
						},
					},
					"error": schema.StringAttribute{
						MarkdownDescription: `An error message for a failed execution`,
						Computed:            true,
					},
					"request": schema.SingleNestedAttribute{
						MarkdownDescription: `ARP table request parameters`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"serial": schema.StringAttribute{
								MarkdownDescription: `Device serial number`,
								Computed:            true,
							},
						},
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `Status of the ARP table request.`,
						Computed:            true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `GET this url to check the status of your ARP table request.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *DevicesLiveToolsArpTableDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesLiveToolsArpTable DevicesLiveToolsArpTable
	diags := req.Config.Get(ctx, &devicesLiveToolsArpTable)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceLiveToolsArpTable")
		vvSerial := devicesLiveToolsArpTable.Serial.ValueString()
		vvArpTableID := devicesLiveToolsArpTable.ArpTableID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Devices.GetDeviceLiveToolsArpTable(vvSerial, vvArpTableID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceLiveToolsArpTable",
				err.Error(),
			)
			return
		}

		devicesLiveToolsArpTable = ResponseDevicesGetDeviceLiveToolsArpTableItemToBody(devicesLiveToolsArpTable, response1)
		diags = resp.State.Set(ctx, &devicesLiveToolsArpTable)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesLiveToolsArpTable struct {
	Serial     types.String                               `tfsdk:"serial"`
	ArpTableID types.String                               `tfsdk:"arp_table_id"`
	Item       *ResponseDevicesGetDeviceLiveToolsArpTable `tfsdk:"item"`
}

type ResponseDevicesGetDeviceLiveToolsArpTable struct {
	ArpTableID types.String                                        `tfsdk:"arp_table_id"`
	Entries    *[]ResponseDevicesGetDeviceLiveToolsArpTableEntries `tfsdk:"entries"`
	Error      types.String                                        `tfsdk:"error"`
	Request    *ResponseDevicesGetDeviceLiveToolsArpTableRequest   `tfsdk:"request"`
	Status     types.String                                        `tfsdk:"status"`
	URL        types.String                                        `tfsdk:"url"`
}

type ResponseDevicesGetDeviceLiveToolsArpTableEntries struct {
	IP            types.String `tfsdk:"ip"`
	LastUpdatedAt types.String `tfsdk:"last_updated_at"`
	Mac           types.String `tfsdk:"mac"`
	VLANID        types.Int64  `tfsdk:"vlan_id"`
}

type ResponseDevicesGetDeviceLiveToolsArpTableRequest struct {
	Serial types.String `tfsdk:"serial"`
}

// ToBody
func ResponseDevicesGetDeviceLiveToolsArpTableItemToBody(state DevicesLiveToolsArpTable, response *merakigosdk.ResponseDevicesGetDeviceLiveToolsArpTable) DevicesLiveToolsArpTable {
	itemState := ResponseDevicesGetDeviceLiveToolsArpTable{
		ArpTableID: types.StringValue(response.ArpTableID),
		Entries: func() *[]ResponseDevicesGetDeviceLiveToolsArpTableEntries {
			if response.Entries != nil {
				result := make([]ResponseDevicesGetDeviceLiveToolsArpTableEntries, len(*response.Entries))
				for i, entries := range *response.Entries {
					result[i] = ResponseDevicesGetDeviceLiveToolsArpTableEntries{
						IP:            types.StringValue(entries.IP),
						LastUpdatedAt: types.StringValue(entries.LastUpdatedAt),
						Mac:           types.StringValue(entries.Mac),
						VLANID: func() types.Int64 {
							if entries.VLANID != nil {
								return types.Int64Value(int64(*entries.VLANID))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Error: types.StringValue(response.Error),
		Request: func() *ResponseDevicesGetDeviceLiveToolsArpTableRequest {
			if response.Request != nil {
				return &ResponseDevicesGetDeviceLiveToolsArpTableRequest{
					Serial: types.StringValue(response.Request.Serial),
				}
			}
			return nil
		}(),
		Status: types.StringValue(response.Status),
		URL:    types.StringValue(response.URL),
	}
	state.Item = &itemState
	return state
}
