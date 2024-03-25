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
	_ datasource.DataSource              = &NetworksApplianceSettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceSettingsDataSource{}
)

func NewNetworksApplianceSettingsDataSource() datasource.DataSource {
	return &NetworksApplianceSettingsDataSource{}
}

type NetworksApplianceSettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceSettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceSettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_settings"
}

func (d *NetworksApplianceSettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"client_tracking_method": schema.StringAttribute{
						MarkdownDescription: `Client tracking method of a network`,
						Computed:            true,
					},
					"deployment_mode": schema.StringAttribute{
						MarkdownDescription: `Deployment mode of a network`,
						Computed:            true,
					},
					"dynamic_dns": schema.SingleNestedAttribute{
						MarkdownDescription: `Dynamic DNS settings for a network`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Dynamic DNS enabled`,
								Computed:            true,
							},
							"prefix": schema.StringAttribute{
								MarkdownDescription: `Dynamic DNS url prefix. DDNS must be enabled to update`,
								Computed:            true,
							},
							"url": schema.StringAttribute{
								MarkdownDescription: `Dynamic DNS url. DDNS must be enabled to update`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceSettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceSettings NetworksApplianceSettings
	diags := req.Config.Get(ctx, &networksApplianceSettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceSettings")
		vvNetworkID := networksApplianceSettings.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceSettings(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceSettings",
				err.Error(),
			)
			return
		}

		networksApplianceSettings = ResponseApplianceGetNetworkApplianceSettingsItemToBody(networksApplianceSettings, response1)
		diags = resp.State.Set(ctx, &networksApplianceSettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceSettings struct {
	NetworkID types.String                                  `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceSettings `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceSettings struct {
	ClientTrackingMethod types.String                                            `tfsdk:"client_tracking_method"`
	DeploymentMode       types.String                                            `tfsdk:"deployment_mode"`
	DynamicDNS           *ResponseApplianceGetNetworkApplianceSettingsDynamicDns `tfsdk:"dynamic_dns"`
}

type ResponseApplianceGetNetworkApplianceSettingsDynamicDns struct {
	Enabled types.Bool   `tfsdk:"enabled"`
	Prefix  types.String `tfsdk:"prefix"`
	URL     types.String `tfsdk:"url"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceSettingsItemToBody(state NetworksApplianceSettings, response *merakigosdk.ResponseApplianceGetNetworkApplianceSettings) NetworksApplianceSettings {
	itemState := ResponseApplianceGetNetworkApplianceSettings{
		ClientTrackingMethod: types.StringValue(response.ClientTrackingMethod),
		DeploymentMode:       types.StringValue(response.DeploymentMode),
		DynamicDNS: func() *ResponseApplianceGetNetworkApplianceSettingsDynamicDns {
			if response.DynamicDNS != nil {
				return &ResponseApplianceGetNetworkApplianceSettingsDynamicDns{
					Enabled: func() types.Bool {
						if response.DynamicDNS.Enabled != nil {
							return types.BoolValue(*response.DynamicDNS.Enabled)
						}
						return types.Bool{}
					}(),
					Prefix: types.StringValue(response.DynamicDNS.Prefix),
					URL:    types.StringValue(response.DynamicDNS.URL),
				}
			}
			return &ResponseApplianceGetNetworkApplianceSettingsDynamicDns{}
		}(),
	}
	state.Item = &itemState
	return state
}
