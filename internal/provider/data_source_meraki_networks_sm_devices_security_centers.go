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
	_ datasource.DataSource              = &NetworksSmDevicesSecurityCentersDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmDevicesSecurityCentersDataSource{}
)

func NewNetworksSmDevicesSecurityCentersDataSource() datasource.DataSource {
	return &NetworksSmDevicesSecurityCentersDataSource{}
}

type NetworksSmDevicesSecurityCentersDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmDevicesSecurityCentersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmDevicesSecurityCentersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_security_centers"
}

func (d *NetworksSmDevicesSecurityCentersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_id": schema.StringAttribute{
				MarkdownDescription: `deviceId path parameter. Device ID`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetNetworkSmDeviceSoftwares`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"app_id": schema.StringAttribute{
							MarkdownDescription: `The Meraki managed application Id for this record on a particular device.`,
							Computed:            true,
						},
						"bundle_size": schema.Int64Attribute{
							MarkdownDescription: `The size of the software bundle.`,
							Computed:            true,
						},
						"created_at": schema.StringAttribute{
							MarkdownDescription: `When the Meraki record for the software was created.`,
							Computed:            true,
						},
						"device_id": schema.StringAttribute{
							MarkdownDescription: `The Meraki managed device Id.`,
							Computed:            true,
						},
						"dynamic_size": schema.Int64Attribute{
							MarkdownDescription: `The size of the data stored in the application.`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `The Meraki software Id.`,
							Computed:            true,
						},
						"identifier": schema.StringAttribute{
							MarkdownDescription: `Software bundle identifier.`,
							Computed:            true,
						},
						"installed_at": schema.StringAttribute{
							MarkdownDescription: `When the Software was installed on the device.`,
							Computed:            true,
						},
						"ios_redemption_code": schema.BoolAttribute{
							MarkdownDescription: `A boolean indicating whether or not an iOS redemption code was used.`,
							Computed:            true,
						},
						"is_managed": schema.BoolAttribute{
							MarkdownDescription: `A boolean indicating whether or not the software is managed by Meraki.`,
							Computed:            true,
						},
						"itunes_id": schema.StringAttribute{
							MarkdownDescription: `The itunes numerical identifier.`,
							Computed:            true,
						},
						"license_key": schema.StringAttribute{
							MarkdownDescription: `The license key associated with this software installation.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the software.`,
							Computed:            true,
						},
						"path": schema.StringAttribute{
							MarkdownDescription: `The path on the device where the software record is located.`,
							Computed:            true,
						},
						"redemption_code": schema.Int64Attribute{
							MarkdownDescription: `The redemption code used for this software.`,
							Computed:            true,
						},
						"short_version": schema.StringAttribute{
							MarkdownDescription: `Short version notation for the software.`,
							Computed:            true,
						},
						"status": schema.StringAttribute{
							MarkdownDescription: `The management status of the software.`,
							Computed:            true,
						},
						"to_install": schema.BoolAttribute{
							MarkdownDescription: `A boolean indicating this software record should be installed on the associated device.`,
							Computed:            true,
						},
						"to_uninstall": schema.BoolAttribute{
							MarkdownDescription: `A boolean indicating this software record should be uninstalled on the associated device.`,
							Computed:            true,
						},
						"uninstalled_at": schema.StringAttribute{
							MarkdownDescription: `When the record was uninstalled from the device.`,
							Computed:            true,
						},
						"updated_at": schema.StringAttribute{
							MarkdownDescription: `When the record was last updated by Meraki.`,
							Computed:            true,
						},
						"vendor": schema.StringAttribute{
							MarkdownDescription: `The vendor of the software.`,
							Computed:            true,
						},
						"version": schema.StringAttribute{
							MarkdownDescription: `Full version notation for the software.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmDevicesSecurityCentersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmDevicesSecurityCenters NetworksSmDevicesSecurityCenters
	diags := req.Config.Get(ctx, &networksSmDevicesSecurityCenters)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksSmDevicesSecurityCenters.NetworkID.IsNull(), !networksSmDevicesSecurityCenters.DeviceID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksSmDevicesSecurityCenters.NetworkID.IsNull(), !networksSmDevicesSecurityCenters.DeviceID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmDeviceSecurityCenters")
		vvNetworkID := networksSmDevicesSecurityCenters.NetworkID.ValueString()
		vvDeviceID := networksSmDevicesSecurityCenters.DeviceID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetNetworkSmDeviceSecurityCenters(vvNetworkID, vvDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmDeviceSecurityCenters",
				err.Error(),
			)
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmDeviceSoftwares")
		vvNetworkID := networksSmDevicesSecurityCenters.NetworkID.ValueString()
		vvDeviceID := networksSmDevicesSecurityCenters.DeviceID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Sm.GetNetworkSmDeviceSoftwares(vvNetworkID, vvDeviceID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmDeviceSoftwares",
				err.Error(),
			)
			return
		}

		networksSmDevicesSecurityCenters = ResponseSmGetNetworkSmDeviceSoftwaresItemsToBody(networksSmDevicesSecurityCenters, response2)
		diags = resp.State.Set(ctx, &networksSmDevicesSecurityCenters)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmDevicesSecurityCenters struct {
	NetworkID types.String                                 `tfsdk:"network_id"`
	DeviceID  types.String                                 `tfsdk:"device_id"`
	Items     *[]ResponseItemSmGetNetworkSmDeviceSoftwares `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmDeviceSoftwares struct {
	AppID             types.String `tfsdk:"app_id"`
	BundleSize        types.Int64  `tfsdk:"bundle_size"`
	CreatedAt         types.String `tfsdk:"created_at"`
	DeviceID          types.String `tfsdk:"device_id"`
	DynamicSize       types.Int64  `tfsdk:"dynamic_size"`
	ID                types.String `tfsdk:"id"`
	IDentifier        types.String `tfsdk:"identifier"`
	InstalledAt       types.String `tfsdk:"installed_at"`
	IosRedemptionCode types.Bool   `tfsdk:"ios_redemption_code"`
	IsManaged         types.Bool   `tfsdk:"is_managed"`
	ItunesID          types.String `tfsdk:"itunes_id"`
	LicenseKey        types.String `tfsdk:"license_key"`
	Name              types.String `tfsdk:"name"`
	Path              types.String `tfsdk:"path"`
	RedemptionCode    types.Int64  `tfsdk:"redemption_code"`
	ShortVersion      types.String `tfsdk:"short_version"`
	Status            types.String `tfsdk:"status"`
	ToInstall         types.Bool   `tfsdk:"to_install"`
	ToUninstall       types.Bool   `tfsdk:"to_uninstall"`
	UninstalledAt     types.String `tfsdk:"uninstalled_at"`
	UpdatedAt         types.String `tfsdk:"updated_at"`
	Vendor            types.String `tfsdk:"vendor"`
	Version           types.String `tfsdk:"version"`
}

// ToBody
func ResponseSmGetNetworkSmDeviceSoftwaresItemsToBody(state NetworksSmDevicesSecurityCenters, response *merakigosdk.ResponseSmGetNetworkSmDeviceSoftwares) NetworksSmDevicesSecurityCenters {
	var items []ResponseItemSmGetNetworkSmDeviceSoftwares
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmDeviceSoftwares{
			AppID: types.StringValue(item.AppID),
			BundleSize: func() types.Int64 {
				if item.BundleSize != nil {
					return types.Int64Value(int64(*item.BundleSize))
				}
				return types.Int64{}
			}(),
			CreatedAt: types.StringValue(item.CreatedAt),
			DeviceID:  types.StringValue(item.DeviceID),
			DynamicSize: func() types.Int64 {
				if item.DynamicSize != nil {
					return types.Int64Value(int64(*item.DynamicSize))
				}
				return types.Int64{}
			}(),
			ID:          types.StringValue(item.ID),
			IDentifier:  types.StringValue(item.IDentifier),
			InstalledAt: types.StringValue(item.InstalledAt),
			IosRedemptionCode: func() types.Bool {
				if item.IosRedemptionCode != nil {
					return types.BoolValue(*item.IosRedemptionCode)
				}
				return types.Bool{}
			}(),
			IsManaged: func() types.Bool {
				if item.IsManaged != nil {
					return types.BoolValue(*item.IsManaged)
				}
				return types.Bool{}
			}(),
			ItunesID:   types.StringValue(item.ItunesID),
			LicenseKey: types.StringValue(item.LicenseKey),
			Name:       types.StringValue(item.Name),
			Path:       types.StringValue(item.Path),
			RedemptionCode: func() types.Int64 {
				if item.RedemptionCode != nil {
					return types.Int64Value(int64(*item.RedemptionCode))
				}
				return types.Int64{}
			}(),
			ShortVersion: types.StringValue(item.ShortVersion),
			Status:       types.StringValue(item.Status),
			ToInstall: func() types.Bool {
				if item.ToInstall != nil {
					return types.BoolValue(*item.ToInstall)
				}
				return types.Bool{}
			}(),
			ToUninstall: func() types.Bool {
				if item.ToUninstall != nil {
					return types.BoolValue(*item.ToUninstall)
				}
				return types.Bool{}
			}(),
			UninstalledAt: types.StringValue(item.UninstalledAt),
			UpdatedAt:     types.StringValue(item.UpdatedAt),
			Vendor:        types.StringValue(item.Vendor),
			Version:       types.StringValue(item.Version),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
