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
	_ datasource.DataSource              = &NetworksSmDevicesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmDevicesDataSource{}
)

func NewNetworksSmDevicesDataSource() datasource.DataSource {
	return &NetworksSmDevicesDataSource{}
}

type NetworksSmDevicesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmDevicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmDevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices"
}

func (d *NetworksSmDevicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"fields": schema.ListAttribute{
				MarkdownDescription: `fields query parameter. Additional fields that will be displayed for each device.
    The default fields are: id, name, tags, ssid, wifiMac, osName, systemModel, uuid, and serialNumber. The additional fields are: ip,
    systemType, availableDeviceCapacity, kioskAppName, biosVersion, lastConnected, missingAppsCount, userSuppliedAddress, location, lastUser,
    ownerEmail, ownerUsername, osBuild, publicIp, phoneNumber, diskInfoJson, deviceCapacity, isManaged, hadMdm, isSupervised, meid, imei, iccid,
    simCarrierNetwork, cellularDataUsed, isHotspotEnabled, createdAt, batteryEstCharge, quarantined, avName, avRunning, asName, fwName,
    isRooted, loginRequired, screenLockEnabled, screenLockDelay, autoLoginDisabled, autoTags, hasMdm, hasDesktopAgent, diskEncryptionEnabled,
    hardwareEncryptionCaps, passCodeLock, usesHardwareKeystore, androidSecurityPatchVersion, cellular, and url.`,
				Optional:    true,
				ElementType: types.StringType,
			},
			"ids": schema.ListAttribute{
				MarkdownDescription: `ids query parameter. Filter devices by id(s).`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"scope": schema.ListAttribute{
				MarkdownDescription: `scope query parameter. Specify a scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Filter devices by serial(s).`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"system_types": schema.ListAttribute{
				MarkdownDescription: `systemTypes query parameter. Filter devices by system type(s).`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"uuids": schema.ListAttribute{
				MarkdownDescription: `uuids query parameter. Filter devices by uuid(s).`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"wifi_macs": schema.ListAttribute{
				MarkdownDescription: `wifiMacs query parameter. Filter devices by wifi mac(s).`,
				Optional:            true,
				ElementType:         types.StringType,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetNetworkSmDevices`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `The Meraki Id of the device record.`,
							Computed:            true,
						},
						"ip": schema.StringAttribute{
							MarkdownDescription: `The IP address of the device.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the device.`,
							Computed:            true,
						},
						"notes": schema.StringAttribute{
							MarkdownDescription: `Notes associated with the device.`,
							Computed:            true,
						},
						"os_name": schema.StringAttribute{
							MarkdownDescription: `The name of the device OS.`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `The device serial.`,
							Computed:            true,
						},
						"serial_number": schema.StringAttribute{
							MarkdownDescription: `The device serial number.`,
							Computed:            true,
						},
						"ssid": schema.StringAttribute{
							MarkdownDescription: `The name of the SSID the device was last connected to.`,
							Computed:            true,
						},
						"system_model": schema.StringAttribute{
							MarkdownDescription: `The device model.`,
							Computed:            true,
						},
						"tags": schema.ListAttribute{
							MarkdownDescription: `An array of tags associated with the device.`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"uuid": schema.StringAttribute{
							MarkdownDescription: `The UUID of the device.`,
							Computed:            true,
						},
						"wifi_mac": schema.StringAttribute{
							MarkdownDescription: `The MAC of the device.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmDevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmDevices NetworksSmDevices
	diags := req.Config.Get(ctx, &networksSmDevices)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmDevices")
		vvNetworkID := networksSmDevices.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkSmDevicesQueryParams{}

		queryParams1.Fields = elementsToStrings(ctx, networksSmDevices.Fields)
		queryParams1.WifiMacs = elementsToStrings(ctx, networksSmDevices.WifiMacs)
		queryParams1.Serials = elementsToStrings(ctx, networksSmDevices.Serials)
		queryParams1.IDs = elementsToStrings(ctx, networksSmDevices.IDs)
		queryParams1.UUIDs = elementsToStrings(ctx, networksSmDevices.UUIDs)
		queryParams1.SystemTypes = elementsToStrings(ctx, networksSmDevices.SystemTypes)
		queryParams1.Scope = elementsToStrings(ctx, networksSmDevices.Scope)
		queryParams1.PerPage = int(networksSmDevices.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksSmDevices.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksSmDevices.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetNetworkSmDevices(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmDevices",
				err.Error(),
			)
			return
		}

		networksSmDevices = ResponseSmGetNetworkSmDevicesItemsToBody(networksSmDevices, response1)
		diags = resp.State.Set(ctx, &networksSmDevices)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmDevices struct {
	NetworkID     types.String                         `tfsdk:"network_id"`
	Fields        types.List                           `tfsdk:"fields"`
	WifiMacs      types.List                           `tfsdk:"wifi_macs"`
	Serials       types.List                           `tfsdk:"serials"`
	IDs           types.List                           `tfsdk:"ids"`
	UUIDs         types.List                           `tfsdk:"uuids"`
	SystemTypes   types.List                           `tfsdk:"system_types"`
	Scope         types.List                           `tfsdk:"scope"`
	PerPage       types.Int64                          `tfsdk:"per_page"`
	StartingAfter types.String                         `tfsdk:"starting_after"`
	EndingBefore  types.String                         `tfsdk:"ending_before"`
	Items         *[]ResponseItemSmGetNetworkSmDevices `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmDevices struct {
	ID           types.String `tfsdk:"id"`
	IP           types.String `tfsdk:"ip"`
	Name         types.String `tfsdk:"name"`
	Notes        types.String `tfsdk:"notes"`
	OsName       types.String `tfsdk:"os_name"`
	Serial       types.String `tfsdk:"serial"`
	SerialNumber types.String `tfsdk:"serial_number"`
	SSID         types.String `tfsdk:"ssid"`
	SystemModel  types.String `tfsdk:"system_model"`
	Tags         types.List   `tfsdk:"tags"`
	UUID         types.String `tfsdk:"uuid"`
	WifiMac      types.String `tfsdk:"wifi_mac"`
}

// ToBody
func ResponseSmGetNetworkSmDevicesItemsToBody(state NetworksSmDevices, response *merakigosdk.ResponseSmGetNetworkSmDevices) NetworksSmDevices {
	var items []ResponseItemSmGetNetworkSmDevices
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmDevices{
			ID:           types.StringValue(item.ID),
			IP:           types.StringValue(item.IP),
			Name:         types.StringValue(item.Name),
			Notes:        types.StringValue(item.Notes),
			OsName:       types.StringValue(item.OsName),
			Serial:       types.StringValue(item.Serial),
			SerialNumber: types.StringValue(item.SerialNumber),
			SSID:         types.StringValue(item.SSID),
			SystemModel:  types.StringValue(item.SystemModel),
			Tags:         StringSliceToList(item.Tags),
			UUID:         types.StringValue(item.UUID),
			WifiMac:      types.StringValue(item.WifiMac),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
