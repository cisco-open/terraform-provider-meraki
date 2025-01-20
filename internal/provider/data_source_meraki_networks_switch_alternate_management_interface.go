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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksSwitchAlternateManagementInterfaceDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSwitchAlternateManagementInterfaceDataSource{}
)

func NewNetworksSwitchAlternateManagementInterfaceDataSource() datasource.DataSource {
	return &NetworksSwitchAlternateManagementInterfaceDataSource{}
}

type NetworksSwitchAlternateManagementInterfaceDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSwitchAlternateManagementInterfaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSwitchAlternateManagementInterfaceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_alternate_management_interface"
}

func (d *NetworksSwitchAlternateManagementInterfaceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean value to enable or disable AMI configuration. If enabled, VLAN and protocols must be set`,
						Computed:            true,
					},
					"protocols": schema.ListAttribute{
						MarkdownDescription: `Can be one or more of the following values: 'radius', 'snmp' or 'syslog'`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"switches": schema.SetNestedAttribute{
						MarkdownDescription: `Array of switch serial number and IP assignment. If parameter is present, it cannot have empty body. Note: switches parameter is not applicable for template networks, in other words, do not put 'switches' in the body when updating template networks. Also, an empty 'switches' array will remove all previous assignments`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"alternate_management_ip": schema.StringAttribute{
									MarkdownDescription: `Switch alternative management IP. To remove a prior IP setting, provide an empty string`,
									Computed:            true,
								},
								"gateway": schema.StringAttribute{
									MarkdownDescription: `Switch gateway must be in IP format. Only and must be specified for Polaris switches`,
									Computed:            true,
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Switch serial number`,
									Computed:            true,
								},
								"subnet_mask": schema.StringAttribute{
									MarkdownDescription: `Switch subnet mask must be in IP format. Only and must be specified for Polaris switches`,
									Computed:            true,
								},
							},
						},
					},
					"vlan_id": schema.Int64Attribute{
						MarkdownDescription: `Alternate management VLAN, must be between 1 and 4094`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksSwitchAlternateManagementInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSwitchAlternateManagementInterface NetworksSwitchAlternateManagementInterface
	diags := req.Config.Get(ctx, &networksSwitchAlternateManagementInterface)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSwitchAlternateManagementInterface")
		vvNetworkID := networksSwitchAlternateManagementInterface.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetNetworkSwitchAlternateManagementInterface(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchAlternateManagementInterface",
				err.Error(),
			)
			return
		}

		networksSwitchAlternateManagementInterface = ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceItemToBody(networksSwitchAlternateManagementInterface, response1)
		diags = resp.State.Set(ctx, &networksSwitchAlternateManagementInterface)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSwitchAlternateManagementInterface struct {
	NetworkID types.String                                                `tfsdk:"network_id"`
	Item      *ResponseSwitchGetNetworkSwitchAlternateManagementInterface `tfsdk:"item"`
}

type ResponseSwitchGetNetworkSwitchAlternateManagementInterface struct {
	Enabled   types.Bool                                                            `tfsdk:"enabled"`
	Protocols types.List                                                            `tfsdk:"protocols"`
	Switches  *[]ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitches `tfsdk:"switches"`
	VLANID    types.Int64                                                           `tfsdk:"vlan_id"`
}

type ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitches struct {
	AlternateManagementIP types.String `tfsdk:"alternate_management_ip"`
	Gateway               types.String `tfsdk:"gateway"`
	Serial                types.String `tfsdk:"serial"`
	SubnetMask            types.String `tfsdk:"subnet_mask"`
}

// ToBody
func ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceItemToBody(state NetworksSwitchAlternateManagementInterface, response *merakigosdk.ResponseSwitchGetNetworkSwitchAlternateManagementInterface) NetworksSwitchAlternateManagementInterface {
	itemState := ResponseSwitchGetNetworkSwitchAlternateManagementInterface{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Protocols: StringSliceToList(response.Protocols),
		Switches: func() *[]ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitches {
			if response.Switches != nil {
				result := make([]ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitches, len(*response.Switches))
				for i, switches := range *response.Switches {
					result[i] = ResponseSwitchGetNetworkSwitchAlternateManagementInterfaceSwitches{
						AlternateManagementIP: types.StringValue(switches.AlternateManagementIP),
						Gateway:               types.StringValue(switches.Gateway),
						Serial:                types.StringValue(switches.Serial),
						SubnetMask:            types.StringValue(switches.SubnetMask),
					}
				}
				return &result
			}
			return nil
		}(),
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
