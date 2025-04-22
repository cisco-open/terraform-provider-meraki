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
	_ datasource.DataSource              = &NetworksCameraWirelessProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksCameraWirelessProfilesDataSource{}
)

func NewNetworksCameraWirelessProfilesDataSource() datasource.DataSource {
	return &NetworksCameraWirelessProfilesDataSource{}
}

type NetworksCameraWirelessProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksCameraWirelessProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksCameraWirelessProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_camera_wireless_profiles"
}

func (d *NetworksCameraWirelessProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"wireless_profile_id": schema.StringAttribute{
				MarkdownDescription: `wirelessProfileId path parameter. Wireless profile ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"applied_device_count": schema.Int64Attribute{
						MarkdownDescription: `The count of the applied devices.`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `The ID of the camera wireless profile.`,
						Computed:            true,
					},
					"identity": schema.SingleNestedAttribute{
						MarkdownDescription: `The identity of the wireless profile. Required for creating wireless profiles in 8021x-radius auth mode.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"password": schema.StringAttribute{
								MarkdownDescription: `The password of the identity.`,
								Sensitive:           true,
								Computed:            true,
							},
							"username": schema.StringAttribute{
								MarkdownDescription: `The username of the identity.`,
								Computed:            true,
							},
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the camera wireless profile.`,
						Computed:            true,
					},
					"ssid": schema.SingleNestedAttribute{
						MarkdownDescription: `The details of the SSID config.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"auth_mode": schema.StringAttribute{
								MarkdownDescription: `The auth mode of the SSID. It can be set to ('psk', '8021x-radius').`,
								Computed:            true,
							},
							"encryption_mode": schema.StringAttribute{
								MarkdownDescription: `The encryption mode of the SSID.`,
								Computed:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `The name of the SSID.`,
								Computed:            true,
							},
							"psk": schema.StringAttribute{
								MarkdownDescription: `The pre-shared key of the SSID, if mode is PSK`,
								Computed:            true,
							},
						},
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseCameraGetNetworkCameraWirelessProfiles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"applied_device_count": schema.Int64Attribute{
							MarkdownDescription: `The count of the applied devices.`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `The ID of the camera wireless profile.`,
							Computed:            true,
						},
						"identity": schema.SingleNestedAttribute{
							MarkdownDescription: `The identity of the wireless profile. Required for creating wireless profiles in 8021x-radius auth mode.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"password": schema.StringAttribute{
									MarkdownDescription: `The password of the identity.`,
									Sensitive:           true,
									Computed:            true,
								},
								"username": schema.StringAttribute{
									MarkdownDescription: `The username of the identity.`,
									Computed:            true,
								},
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the camera wireless profile.`,
							Computed:            true,
						},
						"ssid": schema.SingleNestedAttribute{
							MarkdownDescription: `The details of the SSID config.`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"auth_mode": schema.StringAttribute{
									MarkdownDescription: `The auth mode of the SSID. It can be set to ('psk', '8021x-radius').`,
									Computed:            true,
								},
								"encryption_mode": schema.StringAttribute{
									MarkdownDescription: `The encryption mode of the SSID.`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the SSID.`,
									Computed:            true,
								},
								"psk": schema.StringAttribute{
									MarkdownDescription: `The pre-shared key of the SSID, if mode is PSK`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksCameraWirelessProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksCameraWirelessProfiles NetworksCameraWirelessProfiles
	diags := req.Config.Get(ctx, &networksCameraWirelessProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksCameraWirelessProfiles.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksCameraWirelessProfiles.NetworkID.IsNull(), !networksCameraWirelessProfiles.WirelessProfileID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkCameraWirelessProfiles")
		vvNetworkID := networksCameraWirelessProfiles.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Camera.GetNetworkCameraWirelessProfiles(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCameraWirelessProfiles",
				err.Error(),
			)
			return
		}

		networksCameraWirelessProfiles = ResponseCameraGetNetworkCameraWirelessProfilesItemsToBody(networksCameraWirelessProfiles, response1)
		diags = resp.State.Set(ctx, &networksCameraWirelessProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkCameraWirelessProfile")
		vvNetworkID := networksCameraWirelessProfiles.NetworkID.ValueString()
		vvWirelessProfileID := networksCameraWirelessProfiles.WirelessProfileID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Camera.GetNetworkCameraWirelessProfile(vvNetworkID, vvWirelessProfileID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCameraWirelessProfile",
				err.Error(),
			)
			return
		}

		networksCameraWirelessProfiles = ResponseCameraGetNetworkCameraWirelessProfileItemToBody(networksCameraWirelessProfiles, response2)
		diags = resp.State.Set(ctx, &networksCameraWirelessProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksCameraWirelessProfiles struct {
	NetworkID         types.String                                          `tfsdk:"network_id"`
	WirelessProfileID types.String                                          `tfsdk:"wireless_profile_id"`
	Items             *[]ResponseItemCameraGetNetworkCameraWirelessProfiles `tfsdk:"items"`
	Item              *ResponseCameraGetNetworkCameraWirelessProfile        `tfsdk:"item"`
}

type ResponseItemCameraGetNetworkCameraWirelessProfiles struct {
	AppliedDeviceCount types.Int64                                                 `tfsdk:"applied_device_count"`
	ID                 types.String                                                `tfsdk:"id"`
	IDentity           *ResponseItemCameraGetNetworkCameraWirelessProfilesIdentity `tfsdk:"identity"`
	Name               types.String                                                `tfsdk:"name"`
	SSID               *ResponseItemCameraGetNetworkCameraWirelessProfilesSsid     `tfsdk:"ssid"`
}

type ResponseItemCameraGetNetworkCameraWirelessProfilesIdentity struct {
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}

type ResponseItemCameraGetNetworkCameraWirelessProfilesSsid struct {
	AuthMode       types.String `tfsdk:"auth_mode"`
	EncryptionMode types.String `tfsdk:"encryption_mode"`
	Name           types.String `tfsdk:"name"`
	Psk            types.String `tfsdk:"psk"`
}

type ResponseCameraGetNetworkCameraWirelessProfile struct {
	AppliedDeviceCount types.Int64                                            `tfsdk:"applied_device_count"`
	ID                 types.String                                           `tfsdk:"id"`
	IDentity           *ResponseCameraGetNetworkCameraWirelessProfileIdentity `tfsdk:"identity"`
	Name               types.String                                           `tfsdk:"name"`
	SSID               *ResponseCameraGetNetworkCameraWirelessProfileSsid     `tfsdk:"ssid"`
}

type ResponseCameraGetNetworkCameraWirelessProfileIdentity struct {
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}

type ResponseCameraGetNetworkCameraWirelessProfileSsid struct {
	AuthMode       types.String `tfsdk:"auth_mode"`
	EncryptionMode types.String `tfsdk:"encryption_mode"`
	Name           types.String `tfsdk:"name"`
	Psk            types.String `tfsdk:"psk"`
}

// ToBody
func ResponseCameraGetNetworkCameraWirelessProfilesItemsToBody(state NetworksCameraWirelessProfiles, response *merakigosdk.ResponseCameraGetNetworkCameraWirelessProfiles) NetworksCameraWirelessProfiles {
	var items []ResponseItemCameraGetNetworkCameraWirelessProfiles
	for _, item := range *response {
		itemState := ResponseItemCameraGetNetworkCameraWirelessProfiles{
			AppliedDeviceCount: func() types.Int64 {
				if item.AppliedDeviceCount != nil {
					return types.Int64Value(int64(*item.AppliedDeviceCount))
				}
				return types.Int64{}
			}(),
			ID: types.StringValue(item.ID),
			IDentity: func() *ResponseItemCameraGetNetworkCameraWirelessProfilesIdentity {
				if item.IDentity != nil {
					return &ResponseItemCameraGetNetworkCameraWirelessProfilesIdentity{
						Password: types.StringValue(item.IDentity.Password),
						Username: types.StringValue(item.IDentity.Username),
					}
				}
				return nil
			}(),
			Name: types.StringValue(item.Name),
			SSID: func() *ResponseItemCameraGetNetworkCameraWirelessProfilesSsid {
				if item.SSID != nil {
					return &ResponseItemCameraGetNetworkCameraWirelessProfilesSsid{
						AuthMode:       types.StringValue(item.SSID.AuthMode),
						EncryptionMode: types.StringValue(item.SSID.EncryptionMode),
						Name:           types.StringValue(item.SSID.Name),
						Psk:            types.StringValue(item.SSID.Psk),
					}
				}
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseCameraGetNetworkCameraWirelessProfileItemToBody(state NetworksCameraWirelessProfiles, response *merakigosdk.ResponseCameraGetNetworkCameraWirelessProfile) NetworksCameraWirelessProfiles {
	itemState := ResponseCameraGetNetworkCameraWirelessProfile{
		AppliedDeviceCount: func() types.Int64 {
			if response.AppliedDeviceCount != nil {
				return types.Int64Value(int64(*response.AppliedDeviceCount))
			}
			return types.Int64{}
		}(),
		ID: types.StringValue(response.ID),
		IDentity: func() *ResponseCameraGetNetworkCameraWirelessProfileIdentity {
			if response.IDentity != nil {
				return &ResponseCameraGetNetworkCameraWirelessProfileIdentity{
					Password: types.StringValue(response.IDentity.Password),
					Username: types.StringValue(response.IDentity.Username),
				}
			}
			return nil
		}(),
		Name: types.StringValue(response.Name),
		SSID: func() *ResponseCameraGetNetworkCameraWirelessProfileSsid {
			if response.SSID != nil {
				return &ResponseCameraGetNetworkCameraWirelessProfileSsid{
					AuthMode:       types.StringValue(response.SSID.AuthMode),
					EncryptionMode: types.StringValue(response.SSID.EncryptionMode),
					Name:           types.StringValue(response.SSID.Name),
					Psk:            types.StringValue(response.SSID.Psk),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
