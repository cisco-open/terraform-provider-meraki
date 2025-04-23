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
	_ datasource.DataSource              = &NetworksWirelessSSIDsIDentityPsksDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsIDentityPsksDataSource{}
)

func NewNetworksWirelessSSIDsIDentityPsksDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsIDentityPsksDataSource{}
}

type NetworksWirelessSSIDsIDentityPsksDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsIDentityPsksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsIDentityPsksDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_identity_psks"
}

func (d *NetworksWirelessSSIDsIDentityPsksDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"identity_psk_id": schema.StringAttribute{
				MarkdownDescription: `identityPskId path parameter. Identity psk ID`,
				Optional:            true,
			},
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

					"email": schema.StringAttribute{
						MarkdownDescription: `The email associated with the System's Manager User`,
						Computed:            true,
					},
					"expires_at": schema.StringAttribute{
						MarkdownDescription: `Timestamp for when the Identity PSK expires, or 'null' to never expire`,
						Computed:            true,
					},
					"group_policy_id": schema.StringAttribute{
						MarkdownDescription: `The group policy to be applied to clients`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `The unique identifier of the Identity PSK`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the Identity PSK`,
						Computed:            true,
					},
					"passphrase": schema.StringAttribute{
						MarkdownDescription: `The passphrase for client authentication`,
						Computed:            true,
					},
					"wifi_personal_network_id": schema.StringAttribute{
						MarkdownDescription: `The WiFi Personal Network unique identifier`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseWirelessGetNetworkWirelessSsidIdentityPsks`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"email": schema.StringAttribute{
							MarkdownDescription: `The email associated with the System's Manager User`,
							Computed:            true,
						},
						"expires_at": schema.StringAttribute{
							MarkdownDescription: `Timestamp for when the Identity PSK expires, or 'null' to never expire`,
							Computed:            true,
						},
						"group_policy_id": schema.StringAttribute{
							MarkdownDescription: `The group policy to be applied to clients`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `The unique identifier of the Identity PSK`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the Identity PSK`,
							Computed:            true,
						},
						"passphrase": schema.StringAttribute{
							MarkdownDescription: `The passphrase for client authentication`,
							Computed:            true,
						},
						"wifi_personal_network_id": schema.StringAttribute{
							MarkdownDescription: `The WiFi Personal Network unique identifier`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessSSIDsIDentityPsksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsIDentityPsks NetworksWirelessSSIDsIDentityPsks
	diags := req.Config.Get(ctx, &networksWirelessSSIDsIDentityPsks)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksWirelessSSIDsIDentityPsks.NetworkID.IsNull(), !networksWirelessSSIDsIDentityPsks.Number.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksWirelessSSIDsIDentityPsks.NetworkID.IsNull(), !networksWirelessSSIDsIDentityPsks.Number.IsNull(), !networksWirelessSSIDsIDentityPsks.IDentityPskID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDIDentityPsks")
		vvNetworkID := networksWirelessSSIDsIDentityPsks.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsIDentityPsks.Number.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDIDentityPsks(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDIDentityPsks",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsIDentityPsks = ResponseWirelessGetNetworkWirelessSSIDIDentityPsksItemsToBody(networksWirelessSSIDsIDentityPsks, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsIDentityPsks)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDIDentityPsk")
		vvNetworkID := networksWirelessSSIDsIDentityPsks.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsIDentityPsks.Number.ValueString()
		vvIDentityPskID := networksWirelessSSIDsIDentityPsks.IDentityPskID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Wireless.GetNetworkWirelessSSIDIDentityPsk(vvNetworkID, vvNumber, vvIDentityPskID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDIDentityPsk",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsIDentityPsks = ResponseWirelessGetNetworkWirelessSSIDIDentityPskItemToBody(networksWirelessSSIDsIDentityPsks, response2)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsIDentityPsks)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsIDentityPsks struct {
	NetworkID     types.String                                              `tfsdk:"network_id"`
	Number        types.String                                              `tfsdk:"number"`
	IDentityPskID types.String                                              `tfsdk:"identity_psk_id"`
	Items         *[]ResponseItemWirelessGetNetworkWirelessSsidIdentityPsks `tfsdk:"items"`
	Item          *ResponseWirelessGetNetworkWirelessSsidIdentityPsk        `tfsdk:"item"`
}

type ResponseItemWirelessGetNetworkWirelessSsidIdentityPsks struct {
	Email                 types.String `tfsdk:"email"`
	ExpiresAt             types.String `tfsdk:"expires_at"`
	GroupPolicyID         types.String `tfsdk:"group_policy_id"`
	ID                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	Passphrase            types.String `tfsdk:"passphrase"`
	WifiPersonalNetworkID types.String `tfsdk:"wifi_personal_network_id"`
}

type ResponseWirelessGetNetworkWirelessSsidIdentityPsk struct {
	Email                 types.String `tfsdk:"email"`
	ExpiresAt             types.String `tfsdk:"expires_at"`
	GroupPolicyID         types.String `tfsdk:"group_policy_id"`
	ID                    types.String `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	Passphrase            types.String `tfsdk:"passphrase"`
	WifiPersonalNetworkID types.String `tfsdk:"wifi_personal_network_id"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDIDentityPsksItemsToBody(state NetworksWirelessSSIDsIDentityPsks, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDIDentityPsks) NetworksWirelessSSIDsIDentityPsks {
	var items []ResponseItemWirelessGetNetworkWirelessSsidIdentityPsks
	for _, item := range *response {
		itemState := ResponseItemWirelessGetNetworkWirelessSsidIdentityPsks{
			Email:                 types.StringValue(item.Email),
			ExpiresAt:             types.StringValue(item.ExpiresAt),
			GroupPolicyID:         types.StringValue(item.GroupPolicyID),
			ID:                    types.StringValue(item.ID),
			Name:                  types.StringValue(item.Name),
			Passphrase:            types.StringValue(item.Passphrase),
			WifiPersonalNetworkID: types.StringValue(item.WifiPersonalNetworkID),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseWirelessGetNetworkWirelessSSIDIDentityPskItemToBody(state NetworksWirelessSSIDsIDentityPsks, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDIDentityPsk) NetworksWirelessSSIDsIDentityPsks {
	itemState := ResponseWirelessGetNetworkWirelessSsidIdentityPsk{
		Email:                 types.StringValue(response.Email),
		ExpiresAt:             types.StringValue(response.ExpiresAt),
		GroupPolicyID:         types.StringValue(response.GroupPolicyID),
		ID:                    types.StringValue(response.ID),
		Name:                  types.StringValue(response.Name),
		Passphrase:            types.StringValue(response.Passphrase),
		WifiPersonalNetworkID: types.StringValue(response.WifiPersonalNetworkID),
	}
	state.Item = &itemState
	return state
}
