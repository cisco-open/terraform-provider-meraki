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
	_ datasource.DataSource              = &NetworksWirelessSSIDsEapOverrideDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsEapOverrideDataSource{}
)

func NewNetworksWirelessSSIDsEapOverrideDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsEapOverrideDataSource{}
}

type NetworksWirelessSSIDsEapOverrideDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsEapOverrideDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsEapOverrideDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_eap_override"
}

func (d *NetworksWirelessSSIDsEapOverrideDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"eapol_key": schema.SingleNestedAttribute{
						MarkdownDescription: `EAPOL Key settings.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"retries": schema.Int64Attribute{
								MarkdownDescription: `Maximum number of EAPOL key retries.`,
								Computed:            true,
							},
							"timeout_in_ms": schema.Int64Attribute{
								MarkdownDescription: `EAPOL Key timeout in milliseconds.`,
								Computed:            true,
							},
						},
					},
					"identity": schema.SingleNestedAttribute{
						MarkdownDescription: `EAP settings for identity requests.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"retries": schema.Int64Attribute{
								MarkdownDescription: `Maximum number of EAP retries.`,
								Computed:            true,
							},
							"timeout": schema.Int64Attribute{
								MarkdownDescription: `EAP timeout in seconds.`,
								Computed:            true,
							},
						},
					},
					"max_retries": schema.Int64Attribute{
						MarkdownDescription: `Maximum number of general EAP retries.`,
						Computed:            true,
					},
					"timeout": schema.Int64Attribute{
						MarkdownDescription: `General EAP timeout in seconds.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessSSIDsEapOverrideDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsEapOverride NetworksWirelessSSIDsEapOverride
	diags := req.Config.Get(ctx, &networksWirelessSSIDsEapOverride)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDEapOverride")
		vvNetworkID := networksWirelessSSIDsEapOverride.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsEapOverride.Number.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDEapOverride(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDEapOverride",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsEapOverride = ResponseWirelessGetNetworkWirelessSSIDEapOverrideItemToBody(networksWirelessSSIDsEapOverride, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsEapOverride)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsEapOverride struct {
	NetworkID types.String                                       `tfsdk:"network_id"`
	Number    types.String                                       `tfsdk:"number"`
	Item      *ResponseWirelessGetNetworkWirelessSsidEapOverride `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSsidEapOverride struct {
	EapolKey   *ResponseWirelessGetNetworkWirelessSsidEapOverrideEapolKey `tfsdk:"eapol_key"`
	IDentity   *ResponseWirelessGetNetworkWirelessSsidEapOverrideIdentity `tfsdk:"identity"`
	MaxRetries types.Int64                                                `tfsdk:"max_retries"`
	Timeout    types.Int64                                                `tfsdk:"timeout"`
}

type ResponseWirelessGetNetworkWirelessSsidEapOverrideEapolKey struct {
	Retries     types.Int64 `tfsdk:"retries"`
	TimeoutInMs types.Int64 `tfsdk:"timeout_in_ms"`
}

type ResponseWirelessGetNetworkWirelessSsidEapOverrideIdentity struct {
	Retries types.Int64 `tfsdk:"retries"`
	Timeout types.Int64 `tfsdk:"timeout"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDEapOverrideItemToBody(state NetworksWirelessSSIDsEapOverride, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDEapOverride) NetworksWirelessSSIDsEapOverride {
	itemState := ResponseWirelessGetNetworkWirelessSsidEapOverride{
		EapolKey: func() *ResponseWirelessGetNetworkWirelessSsidEapOverrideEapolKey {
			if response.EapolKey != nil {
				return &ResponseWirelessGetNetworkWirelessSsidEapOverrideEapolKey{
					Retries: func() types.Int64 {
						if response.EapolKey.Retries != nil {
							return types.Int64Value(int64(*response.EapolKey.Retries))
						}
						return types.Int64{}
					}(),
					TimeoutInMs: func() types.Int64 {
						if response.EapolKey.TimeoutInMs != nil {
							return types.Int64Value(int64(*response.EapolKey.TimeoutInMs))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		IDentity: func() *ResponseWirelessGetNetworkWirelessSsidEapOverrideIdentity {
			if response.IDentity != nil {
				return &ResponseWirelessGetNetworkWirelessSsidEapOverrideIdentity{
					Retries: func() types.Int64 {
						if response.IDentity.Retries != nil {
							return types.Int64Value(int64(*response.IDentity.Retries))
						}
						return types.Int64{}
					}(),
					Timeout: func() types.Int64 {
						if response.IDentity.Timeout != nil {
							return types.Int64Value(int64(*response.IDentity.Timeout))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		MaxRetries: func() types.Int64 {
			if response.MaxRetries != nil {
				return types.Int64Value(int64(*response.MaxRetries))
			}
			return types.Int64{}
		}(),
		Timeout: func() types.Int64 {
			if response.Timeout != nil {
				return types.Int64Value(int64(*response.Timeout))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
