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
	_ datasource.DataSource              = &NetworksSmDevicesDeviceProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmDevicesDeviceProfilesDataSource{}
)

func NewNetworksSmDevicesDeviceProfilesDataSource() datasource.DataSource {
	return &NetworksSmDevicesDeviceProfilesDataSource{}
}

type NetworksSmDevicesDeviceProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmDevicesDeviceProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmDevicesDeviceProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_device_profiles"
}

func (d *NetworksSmDevicesDeviceProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `Array of ResponseSmGetNetworkSmDeviceDeviceProfiles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"device_id": schema.StringAttribute{
							MarkdownDescription: `The Meraki managed device Id.`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `The numerical Meraki Id of the profile.`,
							Computed:            true,
						},
						"is_encrypted": schema.BoolAttribute{
							MarkdownDescription: `A boolean indicating if the profile is encrypted.`,
							Computed:            true,
						},
						"is_managed": schema.BoolAttribute{
							MarkdownDescription: `Whether or not the profile is managed by Meraki.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the profile.`,
							Computed:            true,
						},
						"profile_data": schema.StringAttribute{
							MarkdownDescription: `A string containing a JSON object with the profile data.`,
							Computed:            true,
						},
						"profile_identifier": schema.StringAttribute{
							MarkdownDescription: `The identifier of the profile.`,
							Computed:            true,
						},
						"version": schema.StringAttribute{
							MarkdownDescription: `The verison of the profile.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSmDevicesDeviceProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmDevicesDeviceProfiles NetworksSmDevicesDeviceProfiles
	diags := req.Config.Get(ctx, &networksSmDevicesDeviceProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmDeviceDeviceProfiles")
		vvNetworkID := networksSmDevicesDeviceProfiles.NetworkID.ValueString()
		vvDeviceID := networksSmDevicesDeviceProfiles.DeviceID.ValueString()

		response1, restyResp1, err := d.client.Sm.GetNetworkSmDeviceDeviceProfiles(vvNetworkID, vvDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmDeviceDeviceProfiles",
				err.Error(),
			)
			return
		}

		networksSmDevicesDeviceProfiles = ResponseSmGetNetworkSmDeviceDeviceProfilesItemsToBody(networksSmDevicesDeviceProfiles, response1)
		diags = resp.State.Set(ctx, &networksSmDevicesDeviceProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmDevicesDeviceProfiles struct {
	NetworkID types.String                                      `tfsdk:"network_id"`
	DeviceID  types.String                                      `tfsdk:"device_id"`
	Items     *[]ResponseItemSmGetNetworkSmDeviceDeviceProfiles `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmDeviceDeviceProfiles struct {
	DeviceID          types.String `tfsdk:"device_id"`
	ID                types.String `tfsdk:"id"`
	IsEncrypted       types.Bool   `tfsdk:"is_encrypted"`
	IsManaged         types.Bool   `tfsdk:"is_managed"`
	Name              types.String `tfsdk:"name"`
	ProfileData       types.String `tfsdk:"profile_data"`
	ProfileIDentifier types.String `tfsdk:"profile_identifier"`
	Version           types.String `tfsdk:"version"`
}

// ToBody
func ResponseSmGetNetworkSmDeviceDeviceProfilesItemsToBody(state NetworksSmDevicesDeviceProfiles, response *merakigosdk.ResponseSmGetNetworkSmDeviceDeviceProfiles) NetworksSmDevicesDeviceProfiles {
	var items []ResponseItemSmGetNetworkSmDeviceDeviceProfiles
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmDeviceDeviceProfiles{
			DeviceID: types.StringValue(item.DeviceID),
			ID:       types.StringValue(item.ID),
			IsEncrypted: func() types.Bool {
				if item.IsEncrypted != nil {
					return types.BoolValue(*item.IsEncrypted)
				}
				return types.Bool{}
			}(),
			IsManaged: func() types.Bool {
				if item.IsManaged != nil {
					return types.BoolValue(*item.IsManaged)
				}
				return types.Bool{}
			}(),
			Name:              types.StringValue(item.Name),
			ProfileData:       types.StringValue(item.ProfileData),
			ProfileIDentifier: types.StringValue(item.ProfileIDentifier),
			Version:           types.StringValue(item.Version),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
