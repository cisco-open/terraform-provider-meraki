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
	_ datasource.DataSource              = &NetworksClientsSplashAuthorizationStatusDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksClientsSplashAuthorizationStatusDataSource{}
)

func NewNetworksClientsSplashAuthorizationStatusDataSource() datasource.DataSource {
	return &NetworksClientsSplashAuthorizationStatusDataSource{}
}

type NetworksClientsSplashAuthorizationStatusDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksClientsSplashAuthorizationStatusDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksClientsSplashAuthorizationStatusDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_clients_splash_authorization_status"
}

func (d *NetworksClientsSplashAuthorizationStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"client_id": schema.StringAttribute{
				MarkdownDescription: `clientId path parameter. Client ID`,
				Required:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"ssids": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"status_0": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"authorized_at": schema.StringAttribute{
										Computed: true,
									},
									"expires_at": schema.StringAttribute{
										Computed: true,
									},
									"is_authorized": schema.BoolAttribute{
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksClientsSplashAuthorizationStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksClientsSplashAuthorizationStatus NetworksClientsSplashAuthorizationStatus
	diags := req.Config.Get(ctx, &networksClientsSplashAuthorizationStatus)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkClientSplashAuthorizationStatus")
		vvNetworkID := networksClientsSplashAuthorizationStatus.NetworkID.ValueString()
		vvClientID := networksClientsSplashAuthorizationStatus.ClientID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkClientSplashAuthorizationStatus(vvNetworkID, vvClientID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkClientSplashAuthorizationStatus",
				err.Error(),
			)
			return
		}

		networksClientsSplashAuthorizationStatus = ResponseNetworksGetNetworkClientSplashAuthorizationStatusItemToBody(networksClientsSplashAuthorizationStatus, response1)
		diags = resp.State.Set(ctx, &networksClientsSplashAuthorizationStatus)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksClientsSplashAuthorizationStatus struct {
	NetworkID types.String                                               `tfsdk:"network_id"`
	ClientID  types.String                                               `tfsdk:"client_id"`
	Item      *ResponseNetworksGetNetworkClientSplashAuthorizationStatus `tfsdk:"item"`
}

type ResponseNetworksGetNetworkClientSplashAuthorizationStatus struct {
	SSIDs *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids `tfsdk:"ssids"`
}

type ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids struct {
	Status0 *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids0 `tfsdk:"status_0"`
}

type ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids0 struct {
	AuthorizedAt types.String `tfsdk:"authorized_at"`
	ExpiresAt    types.String `tfsdk:"expires_at"`
	IsAuthorized types.Bool   `tfsdk:"is_authorized"`
}

// ToBody
func ResponseNetworksGetNetworkClientSplashAuthorizationStatusItemToBody(state NetworksClientsSplashAuthorizationStatus, response *merakigosdk.ResponseNetworksGetNetworkClientSplashAuthorizationStatus) NetworksClientsSplashAuthorizationStatus {
	itemState := ResponseNetworksGetNetworkClientSplashAuthorizationStatus{
		SSIDs: func() *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids {
			if response.SSIDs != nil {
				return &ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids{
					Status0: func() *ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids0 {
						if response.SSIDs.Status0 != nil {
							return &ResponseNetworksGetNetworkClientSplashAuthorizationStatusSsids0{
								AuthorizedAt: types.StringValue(response.SSIDs.Status0.AuthorizedAt),
								ExpiresAt:    types.StringValue(response.SSIDs.Status0.ExpiresAt),
								IsAuthorized: func() types.Bool {
									if response.SSIDs.Status0.IsAuthorized != nil {
										return types.BoolValue(*response.SSIDs.Status0.IsAuthorized)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
