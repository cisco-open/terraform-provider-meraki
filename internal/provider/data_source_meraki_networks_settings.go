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
	_ datasource.DataSource              = &NetworksSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSettingsDataSource{}
)

func NewNetworksSettingsDataSource() datasource.DataSource {
	return &NetworksSettingsDataSource{}
}

type NetworksSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_settings"
}

func (d *NetworksSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"fips": schema.SingleNestedAttribute{
						MarkdownDescription: `A hash of FIPS options applied to the Network`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enables / disables FIPS on the network.`,
								Computed:            true,
							},
						},
					},
					"local_status_page": schema.SingleNestedAttribute{
						MarkdownDescription: `A hash of Local Status page(s)' authentication options applied to the Network.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"authentication": schema.SingleNestedAttribute{
								MarkdownDescription: `A hash of Local Status page(s)' authentication options applied to the Network.`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"enabled": schema.BoolAttribute{
										MarkdownDescription: `Enables / disables the authentication on Local Status page(s).`,
										Computed:            true,
									},
									"username": schema.StringAttribute{
										MarkdownDescription: `The username used for Local Status Page(s).`,
										Computed:            true,
									},
								},
							},
						},
					},
					"local_status_page_enabled": schema.BoolAttribute{
						MarkdownDescription: `Enables / disables the local device status pages (<a target='_blank' href='http://my.meraki.com/'>my.meraki.com, </a><a target='_blank' href='http://ap.meraki.com/'>ap.meraki.com, </a><a target='_blank' href='http://switch.meraki.com/'>switch.meraki.com, </a><a target='_blank' href='http://wired.meraki.com/'>wired.meraki.com</a>). Optional (defaults to false)`,
						Computed:            true,
					},
					"named_vlans": schema.SingleNestedAttribute{
						MarkdownDescription: `A hash of Named VLANs options applied to the Network.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enables / disables Named VLANs on the Network.`,
								Computed:            true,
							},
						},
					},
					"remote_status_page_enabled": schema.BoolAttribute{
						MarkdownDescription: `Enables / disables access to the device status page (<a target='_blank'>http://[device's LAN IP])</a>. Optional. Can only be set if localStatusPageEnabled is set to true`,
						Computed:            true,
					},
					"secure_port": schema.SingleNestedAttribute{
						MarkdownDescription: `A hash of SecureConnect options applied to the Network.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enables / disables SecureConnect on the network. Optional.`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSettings NetworksSettings
	diags := req.Config.Get(ctx, &networksSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSettings")
		vvNetworkID := networksSettings.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkSettings(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSettings",
				err.Error(),
			)
			return
		}

		networksSettings = ResponseNetworksGetNetworkSettingsItemToBody(networksSettings, response1)
		diags = resp.State.Set(ctx, &networksSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSettings struct {
	NetworkID types.String                        `tfsdk:"network_id"`
	Item      *ResponseNetworksGetNetworkSettings `tfsdk:"item"`
}

type ResponseNetworksGetNetworkSettings struct {
	Fips                    *ResponseNetworksGetNetworkSettingsFips            `tfsdk:"fips"`
	LocalStatusPage         *ResponseNetworksGetNetworkSettingsLocalStatusPage `tfsdk:"local_status_page"`
	LocalStatusPageEnabled  types.Bool                                         `tfsdk:"local_status_page_enabled"`
	NamedVLANs              *ResponseNetworksGetNetworkSettingsNamedVlans      `tfsdk:"named_vlans"`
	RemoteStatusPageEnabled types.Bool                                         `tfsdk:"remote_status_page_enabled"`
	SecurePort              *ResponseNetworksGetNetworkSettingsSecurePort      `tfsdk:"secure_port"`
}

type ResponseNetworksGetNetworkSettingsFips struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseNetworksGetNetworkSettingsLocalStatusPage struct {
	Authentication *ResponseNetworksGetNetworkSettingsLocalStatusPageAuthentication `tfsdk:"authentication"`
}

type ResponseNetworksGetNetworkSettingsLocalStatusPageAuthentication struct {
	Enabled  types.Bool   `tfsdk:"enabled"`
	Username types.String `tfsdk:"username"`
}

type ResponseNetworksGetNetworkSettingsNamedVlans struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseNetworksGetNetworkSettingsSecurePort struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// ToBody
func ResponseNetworksGetNetworkSettingsItemToBody(state NetworksSettings, response *merakigosdk.ResponseNetworksGetNetworkSettings) NetworksSettings {
	itemState := ResponseNetworksGetNetworkSettings{
		Fips: func() *ResponseNetworksGetNetworkSettingsFips {
			if response.Fips != nil {
				return &ResponseNetworksGetNetworkSettingsFips{
					Enabled: func() types.Bool {
						if response.Fips.Enabled != nil {
							return types.BoolValue(*response.Fips.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		LocalStatusPage: func() *ResponseNetworksGetNetworkSettingsLocalStatusPage {
			if response.LocalStatusPage != nil {
				return &ResponseNetworksGetNetworkSettingsLocalStatusPage{
					Authentication: func() *ResponseNetworksGetNetworkSettingsLocalStatusPageAuthentication {
						if response.LocalStatusPage.Authentication != nil {
							return &ResponseNetworksGetNetworkSettingsLocalStatusPageAuthentication{
								Enabled: func() types.Bool {
									if response.LocalStatusPage.Authentication.Enabled != nil {
										return types.BoolValue(*response.LocalStatusPage.Authentication.Enabled)
									}
									return types.Bool{}
								}(),
								Username: types.StringValue(response.LocalStatusPage.Authentication.Username),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		LocalStatusPageEnabled: func() types.Bool {
			if response.LocalStatusPageEnabled != nil {
				return types.BoolValue(*response.LocalStatusPageEnabled)
			}
			return types.Bool{}
		}(),
		NamedVLANs: func() *ResponseNetworksGetNetworkSettingsNamedVlans {
			if response.NamedVLANs != nil {
				return &ResponseNetworksGetNetworkSettingsNamedVlans{
					Enabled: func() types.Bool {
						if response.NamedVLANs.Enabled != nil {
							return types.BoolValue(*response.NamedVLANs.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		RemoteStatusPageEnabled: func() types.Bool {
			if response.RemoteStatusPageEnabled != nil {
				return types.BoolValue(*response.RemoteStatusPageEnabled)
			}
			return types.Bool{}
		}(),
		SecurePort: func() *ResponseNetworksGetNetworkSettingsSecurePort {
			if response.SecurePort != nil {
				return &ResponseNetworksGetNetworkSettingsSecurePort{
					Enabled: func() types.Bool {
						if response.SecurePort.Enabled != nil {
							return types.BoolValue(*response.SecurePort.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
