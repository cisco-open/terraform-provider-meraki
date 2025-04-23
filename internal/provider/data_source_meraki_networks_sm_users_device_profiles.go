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
	_ datasource.DataSource              = &NetworksSmUsersDeviceProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSmUsersDeviceProfilesDataSource{}
)

func NewNetworksSmUsersDeviceProfilesDataSource() datasource.DataSource {
	return &NetworksSmUsersDeviceProfilesDataSource{}
}

type NetworksSmUsersDeviceProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSmUsersDeviceProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSmUsersDeviceProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_users_device_profiles"
}

func (d *NetworksSmUsersDeviceProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"user_id": schema.StringAttribute{
				MarkdownDescription: `userId path parameter. User ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetNetworkSmUserDeviceProfiles`,
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

func (d *NetworksSmUsersDeviceProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSmUsersDeviceProfiles NetworksSmUsersDeviceProfiles
	diags := req.Config.Get(ctx, &networksSmUsersDeviceProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSmUserDeviceProfiles")
		vvNetworkID := networksSmUsersDeviceProfiles.NetworkID.ValueString()
		vvUserID := networksSmUsersDeviceProfiles.UserID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetNetworkSmUserDeviceProfiles(vvNetworkID, vvUserID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmUserDeviceProfiles",
				err.Error(),
			)
			return
		}

		networksSmUsersDeviceProfiles = ResponseSmGetNetworkSmUserDeviceProfilesItemsToBody(networksSmUsersDeviceProfiles, response1)
		diags = resp.State.Set(ctx, &networksSmUsersDeviceProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSmUsersDeviceProfiles struct {
	NetworkID types.String                                    `tfsdk:"network_id"`
	UserID    types.String                                    `tfsdk:"user_id"`
	Items     *[]ResponseItemSmGetNetworkSmUserDeviceProfiles `tfsdk:"items"`
}

type ResponseItemSmGetNetworkSmUserDeviceProfiles struct {
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
func ResponseSmGetNetworkSmUserDeviceProfilesItemsToBody(state NetworksSmUsersDeviceProfiles, response *merakigosdk.ResponseSmGetNetworkSmUserDeviceProfiles) NetworksSmUsersDeviceProfiles {
	var items []ResponseItemSmGetNetworkSmUserDeviceProfiles
	for _, item := range *response {
		itemState := ResponseItemSmGetNetworkSmUserDeviceProfiles{
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
