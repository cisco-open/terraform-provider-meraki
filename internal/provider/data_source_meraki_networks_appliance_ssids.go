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
	_ datasource.DataSource              = &NetworksApplianceSSIDsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceSSIDsDataSource{}
)

func NewNetworksApplianceSSIDsDataSource() datasource.DataSource {
	return &NetworksApplianceSSIDsDataSource{}
}

type NetworksApplianceSSIDsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceSSIDsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceSSIDsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_ssids"
}

func (d *NetworksApplianceSSIDsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"auth_mode": schema.StringAttribute{
						MarkdownDescription: `The association control method for the SSID.`,
						Computed:            true,
					},
					"default_vlan_id": schema.Int64Attribute{
						MarkdownDescription: `The VLAN ID of the VLAN associated to this SSID.`,
						Computed:            true,
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Whether or not the SSID is enabled.`,
						Computed:            true,
					},
					"encryption_mode": schema.StringAttribute{
						MarkdownDescription: `The psk encryption mode for the SSID.`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the SSID.`,
						Computed:            true,
					},
					"number": schema.Int64Attribute{
						MarkdownDescription: `The number of the SSID.`,
						Computed:            true,
					},
					"radius_servers": schema.SetNestedAttribute{
						MarkdownDescription: `The RADIUS 802.1x servers to be used for authentication.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"host": schema.StringAttribute{
									MarkdownDescription: `The IP address of your RADIUS server.`,
									Computed:            true,
								},
								"port": schema.Int64Attribute{
									MarkdownDescription: `The UDP port your RADIUS servers listens on for Access-requests.`,
									Computed:            true,
								},
							},
						},
					},
					"visible": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating whether the MX should advertise or hide this SSID.`,
						Computed:            true,
					},
					"wpa_encryption_mode": schema.StringAttribute{
						MarkdownDescription: `WPA encryption mode for the SSID.`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseApplianceGetNetworkApplianceSsids`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"auth_mode": schema.StringAttribute{
							MarkdownDescription: `The association control method for the SSID.`,
							Computed:            true,
						},
						"default_vlan_id": schema.Int64Attribute{
							MarkdownDescription: `The VLAN ID of the VLAN associated to this SSID.`,
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: `Whether or not the SSID is enabled.`,
							Computed:            true,
						},
						"encryption_mode": schema.StringAttribute{
							MarkdownDescription: `The psk encryption mode for the SSID.`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the SSID.`,
							Computed:            true,
						},
						"number": schema.Int64Attribute{
							MarkdownDescription: `The number of the SSID.`,
							Computed:            true,
						},
						"radius_servers": schema.SetNestedAttribute{
							MarkdownDescription: `The RADIUS 802.1x servers to be used for authentication.`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"host": schema.StringAttribute{
										MarkdownDescription: `The IP address of your RADIUS server.`,
										Computed:            true,
									},
									"port": schema.Int64Attribute{
										MarkdownDescription: `The UDP port your RADIUS servers listens on for Access-requests.`,
										Computed:            true,
									},
								},
							},
						},
						"visible": schema.BoolAttribute{
							MarkdownDescription: `Boolean indicating whether the MX should advertise or hide this SSID.`,
							Computed:            true,
						},
						"wpa_encryption_mode": schema.StringAttribute{
							MarkdownDescription: `WPA encryption mode for the SSID.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceSSIDsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceSSIDs NetworksApplianceSSIDs
	diags := req.Config.Get(ctx, &networksApplianceSSIDs)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksApplianceSSIDs.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksApplianceSSIDs.NetworkID.IsNull(), !networksApplianceSSIDs.Number.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceSSIDs")
		vvNetworkID := networksApplianceSSIDs.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceSSIDs(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceSSIDs",
				err.Error(),
			)
			return
		}

		networksApplianceSSIDs = ResponseApplianceGetNetworkApplianceSSIDsItemsToBody(networksApplianceSSIDs, response1)
		diags = resp.State.Set(ctx, &networksApplianceSSIDs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceSSID")
		vvNetworkID := networksApplianceSSIDs.NetworkID.ValueString()
		vvNumber := networksApplianceSSIDs.Number.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Appliance.GetNetworkApplianceSSID(vvNetworkID, vvNumber)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceSSID",
				err.Error(),
			)
			return
		}

		networksApplianceSSIDs = ResponseApplianceGetNetworkApplianceSSIDItemToBody(networksApplianceSSIDs, response2)
		diags = resp.State.Set(ctx, &networksApplianceSSIDs)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceSSIDs struct {
	NetworkID types.String                                     `tfsdk:"network_id"`
	Number    types.String                                     `tfsdk:"number"`
	Items     *[]ResponseItemApplianceGetNetworkApplianceSsids `tfsdk:"items"`
	Item      *ResponseApplianceGetNetworkApplianceSsid        `tfsdk:"item"`
}

type ResponseItemApplianceGetNetworkApplianceSsids struct {
	AuthMode          types.String                                                  `tfsdk:"auth_mode"`
	DefaultVLANID     types.Int64                                                   `tfsdk:"default_vlan_id"`
	Enabled           types.Bool                                                    `tfsdk:"enabled"`
	EncryptionMode    types.String                                                  `tfsdk:"encryption_mode"`
	Name              types.String                                                  `tfsdk:"name"`
	Number            types.Int64                                                   `tfsdk:"number"`
	RadiusServers     *[]ResponseItemApplianceGetNetworkApplianceSsidsRadiusServers `tfsdk:"radius_servers"`
	Visible           types.Bool                                                    `tfsdk:"visible"`
	WpaEncryptionMode types.String                                                  `tfsdk:"wpa_encryption_mode"`
}

type ResponseItemApplianceGetNetworkApplianceSsidsRadiusServers struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

type ResponseApplianceGetNetworkApplianceSsid struct {
	AuthMode          types.String                                             `tfsdk:"auth_mode"`
	DefaultVLANID     types.Int64                                              `tfsdk:"default_vlan_id"`
	Enabled           types.Bool                                               `tfsdk:"enabled"`
	EncryptionMode    types.String                                             `tfsdk:"encryption_mode"`
	Name              types.String                                             `tfsdk:"name"`
	Number            types.Int64                                              `tfsdk:"number"`
	RadiusServers     *[]ResponseApplianceGetNetworkApplianceSsidRadiusServers `tfsdk:"radius_servers"`
	Visible           types.Bool                                               `tfsdk:"visible"`
	WpaEncryptionMode types.String                                             `tfsdk:"wpa_encryption_mode"`
}

type ResponseApplianceGetNetworkApplianceSsidRadiusServers struct {
	Host types.String `tfsdk:"host"`
	Port types.Int64  `tfsdk:"port"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceSSIDsItemsToBody(state NetworksApplianceSSIDs, response *merakigosdk.ResponseApplianceGetNetworkApplianceSSIDs) NetworksApplianceSSIDs {
	var items []ResponseItemApplianceGetNetworkApplianceSsids
	for _, item := range *response {
		itemState := ResponseItemApplianceGetNetworkApplianceSsids{
			AuthMode: types.StringValue(item.AuthMode),
			DefaultVLANID: func() types.Int64 {
				if item.DefaultVLANID != nil {
					return types.Int64Value(int64(*item.DefaultVLANID))
				}
				return types.Int64{}
			}(),
			Enabled: func() types.Bool {
				if item.Enabled != nil {
					return types.BoolValue(*item.Enabled)
				}
				return types.Bool{}
			}(),
			EncryptionMode: types.StringValue(item.EncryptionMode),
			Name:           types.StringValue(item.Name),
			Number: func() types.Int64 {
				if item.Number != nil {
					return types.Int64Value(int64(*item.Number))
				}
				return types.Int64{}
			}(),
			RadiusServers: func() *[]ResponseItemApplianceGetNetworkApplianceSsidsRadiusServers {
				if item.RadiusServers != nil {
					result := make([]ResponseItemApplianceGetNetworkApplianceSsidsRadiusServers, len(*item.RadiusServers))
					for i, radiusServers := range *item.RadiusServers {
						result[i] = ResponseItemApplianceGetNetworkApplianceSsidsRadiusServers{
							Host: types.StringValue(radiusServers.Host),
							Port: func() types.Int64 {
								if radiusServers.Port != nil {
									return types.Int64Value(int64(*radiusServers.Port))
								}
								return types.Int64{}
							}(),
						}
					}
					return &result
				}
				return nil
			}(),
			Visible: func() types.Bool {
				if item.Visible != nil {
					return types.BoolValue(*item.Visible)
				}
				return types.Bool{}
			}(),
			WpaEncryptionMode: types.StringValue(item.WpaEncryptionMode),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseApplianceGetNetworkApplianceSSIDItemToBody(state NetworksApplianceSSIDs, response *merakigosdk.ResponseApplianceGetNetworkApplianceSSID) NetworksApplianceSSIDs {
	itemState := ResponseApplianceGetNetworkApplianceSsid{
		AuthMode: types.StringValue(response.AuthMode),
		DefaultVLANID: func() types.Int64 {
			if response.DefaultVLANID != nil {
				return types.Int64Value(int64(*response.DefaultVLANID))
			}
			return types.Int64{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		EncryptionMode: types.StringValue(response.EncryptionMode),
		Name:           types.StringValue(response.Name),
		Number: func() types.Int64 {
			if response.Number != nil {
				return types.Int64Value(int64(*response.Number))
			}
			return types.Int64{}
		}(),
		RadiusServers: func() *[]ResponseApplianceGetNetworkApplianceSsidRadiusServers {
			if response.RadiusServers != nil {
				result := make([]ResponseApplianceGetNetworkApplianceSsidRadiusServers, len(*response.RadiusServers))
				for i, radiusServers := range *response.RadiusServers {
					result[i] = ResponseApplianceGetNetworkApplianceSsidRadiusServers{
						Host: types.StringValue(radiusServers.Host),
						Port: func() types.Int64 {
							if radiusServers.Port != nil {
								return types.Int64Value(int64(*radiusServers.Port))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Visible: func() types.Bool {
			if response.Visible != nil {
				return types.BoolValue(*response.Visible)
			}
			return types.Bool{}
		}(),
		WpaEncryptionMode: types.StringValue(response.WpaEncryptionMode),
	}
	state.Item = &itemState
	return state
}
