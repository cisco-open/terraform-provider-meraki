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
	_ datasource.DataSource              = &NetworksWirelessAlternateManagementInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessAlternateManagementInterfaceDataSource{}
)

func NewNetworksWirelessAlternateManagementInterfaceDataSource() datasource.DataSource {
	return &NetworksWirelessAlternateManagementInterfaceDataSource{}
}

type NetworksWirelessAlternateManagementInterfaceDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessAlternateManagementInterfaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessAlternateManagementInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_alternate_management_interface"
}

func (d *NetworksWirelessAlternateManagementInterfaceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"access_points": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"alternate_management_ip": schema.StringAttribute{
									Computed: true,
								},
								"dns1": schema.StringAttribute{
									Computed: true,
								},
								"dns2": schema.StringAttribute{
									Computed: true,
								},
								"gateway": schema.StringAttribute{
									Computed: true,
								},
								"serial": schema.StringAttribute{
									Computed: true,
								},
								"subnet_mask": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"enabled": schema.BoolAttribute{
						Computed: true,
					},
					"protocols": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"vlan_id": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessAlternateManagementInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessAlternateManagementInterface NetworksWirelessAlternateManagementInterface
	diags := req.Config.Get(ctx, &networksWirelessAlternateManagementInterface)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessAlternateManagementInterface")
		vvNetworkID := networksWirelessAlternateManagementInterface.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessAlternateManagementInterface(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessAlternateManagementInterface",
				err.Error(),
			)
			return
		}

		networksWirelessAlternateManagementInterface = ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceItemToBody(networksWirelessAlternateManagementInterface, response1)
		diags = resp.State.Set(ctx, &networksWirelessAlternateManagementInterface)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessAlternateManagementInterface struct {
	NetworkID types.String                                                    `tfsdk:"network_id"`
	Item      *ResponseWirelessGetNetworkWirelessAlternateManagementInterface `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessAlternateManagementInterface struct {
	AccessPoints *[]ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPoints `tfsdk:"access_points"`
	Enabled      types.Bool                                                                    `tfsdk:"enabled"`
	Protocols    types.List                                                                    `tfsdk:"protocols"`
	VLANID       types.Int64                                                                   `tfsdk:"vlan_id"`
}

type ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPoints struct {
	AlternateManagementIP types.String `tfsdk:"alternate_management_ip"`
	DNS1                  types.String `tfsdk:"dns1"`
	DNS2                  types.String `tfsdk:"dns2"`
	Gateway               types.String `tfsdk:"gateway"`
	Serial                types.String `tfsdk:"serial"`
	SubnetMask            types.String `tfsdk:"subnet_mask"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceItemToBody(state NetworksWirelessAlternateManagementInterface, response *merakigosdk.ResponseWirelessGetNetworkWirelessAlternateManagementInterface) NetworksWirelessAlternateManagementInterface {
	itemState := ResponseWirelessGetNetworkWirelessAlternateManagementInterface{
		AccessPoints: func() *[]ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPoints {
			if response.AccessPoints != nil {
				result := make([]ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPoints, len(*response.AccessPoints))
				for i, accessPoints := range *response.AccessPoints {
					result[i] = ResponseWirelessGetNetworkWirelessAlternateManagementInterfaceAccessPoints{
						AlternateManagementIP: types.StringValue(accessPoints.AlternateManagementIP),
						DNS1:                  types.StringValue(accessPoints.DNS1),
						DNS2:                  types.StringValue(accessPoints.DNS2),
						Gateway:               types.StringValue(accessPoints.Gateway),
						Serial:                types.StringValue(accessPoints.Serial),
						SubnetMask:            types.StringValue(accessPoints.SubnetMask),
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
		Protocols: StringSliceToList(response.Protocols),
		VLANID: func() types.Int64 {
			if response.VLANID != nil {
				return types.Int64Value(int64(*response.VLANID))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
