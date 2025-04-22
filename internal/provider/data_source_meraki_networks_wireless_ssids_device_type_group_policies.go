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
	_ datasource.DataSource              = &NetworksWirelessSSIDsDeviceTypeGroupPoliciesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsDeviceTypeGroupPoliciesDataSource{}
)

func NewNetworksWirelessSSIDsDeviceTypeGroupPoliciesDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsDeviceTypeGroupPoliciesDataSource{}
}

type NetworksWirelessSSIDsDeviceTypeGroupPoliciesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsDeviceTypeGroupPoliciesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsDeviceTypeGroupPoliciesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_device_type_group_policies"
}

func (d *NetworksWirelessSSIDsDeviceTypeGroupPoliciesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

					"device_type_policies": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"device_policy": schema.StringAttribute{
									Computed: true,
								},
								"device_type": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"enabled": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessSSIDsDeviceTypeGroupPoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsDeviceTypeGroupPolicies NetworksWirelessSSIDsDeviceTypeGroupPolicies
	diags := req.Config.Get(ctx, &networksWirelessSSIDsDeviceTypeGroupPolicies)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDDeviceTypeGroupPolicies")
		vvNetworkID := networksWirelessSSIDsDeviceTypeGroupPolicies.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsDeviceTypeGroupPolicies.Number.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDDeviceTypeGroupPolicies(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDDeviceTypeGroupPolicies",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsDeviceTypeGroupPolicies = ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPoliciesItemToBody(networksWirelessSSIDsDeviceTypeGroupPolicies, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsDeviceTypeGroupPolicies)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsDeviceTypeGroupPolicies struct {
	NetworkID types.String                                                   `tfsdk:"network_id"`
	Number    types.String                                                   `tfsdk:"number"`
	Item      *ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPolicies `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPolicies struct {
	DeviceTypePolicies *[]ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPoliciesDeviceTypePolicies `tfsdk:"device_type_policies"`
	Enabled            types.Bool                                                                         `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPoliciesDeviceTypePolicies struct {
	DevicePolicy types.String `tfsdk:"device_policy"`
	DeviceType   types.String `tfsdk:"device_type"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPoliciesItemToBody(state NetworksWirelessSSIDsDeviceTypeGroupPolicies, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDDeviceTypeGroupPolicies) NetworksWirelessSSIDsDeviceTypeGroupPolicies {
	itemState := ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPolicies{
		DeviceTypePolicies: func() *[]ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPoliciesDeviceTypePolicies {
			if response.DeviceTypePolicies != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPoliciesDeviceTypePolicies, len(*response.DeviceTypePolicies))
				for i, deviceTypePolicies := range *response.DeviceTypePolicies {
					result[i] = ResponseWirelessGetNetworkWirelessSsidDeviceTypeGroupPoliciesDeviceTypePolicies{
						DevicePolicy: types.StringValue(deviceTypePolicies.DevicePolicy),
						DeviceType:   types.StringValue(deviceTypePolicies.DeviceType),
					}
				}
				return &result
			}
			return nil
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
	}
	state.Item = &itemState
	return state
}
