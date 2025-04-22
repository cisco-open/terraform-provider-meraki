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
	_ datasource.DataSource              = &NetworksAppliancePortsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksAppliancePortsDataSource{}
)

func NewNetworksAppliancePortsDataSource() datasource.DataSource {
	return &NetworksAppliancePortsDataSource{}
}

type NetworksAppliancePortsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksAppliancePortsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksAppliancePortsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_ports"
}

func (d *NetworksAppliancePortsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"port_id": schema.StringAttribute{
				MarkdownDescription: `portId path parameter. Port ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"access_policy": schema.StringAttribute{
						MarkdownDescription: `The name of the policy. Only applicable to Access ports.`,
						Computed:            true,
					},
					"allowed_vlans": schema.StringAttribute{
						MarkdownDescription: `Comma-delimited list of the VLAN ID's allowed on the port, or 'all' to permit all VLAN's on the port.`,
						Computed:            true,
					},
					"drop_untagged_traffic": schema.BoolAttribute{
						MarkdownDescription: `Whether the trunk port can drop all untagged traffic.`,
						Computed:            true,
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `The status of the port`,
						Computed:            true,
					},
					"number": schema.Int64Attribute{
						MarkdownDescription: `Number of the port`,
						Computed:            true,
					},
					"type": schema.StringAttribute{
						MarkdownDescription: `The type of the port: 'access' or 'trunk'.`,
						Computed:            true,
					},
					"vlan": schema.Int64Attribute{
						MarkdownDescription: `Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseApplianceGetNetworkAppliancePorts`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"access_policy": schema.StringAttribute{
							MarkdownDescription: `The name of the policy. Only applicable to Access ports.`,
							Computed:            true,
						},
						"allowed_vlans": schema.StringAttribute{
							MarkdownDescription: `Comma-delimited list of the VLAN ID's allowed on the port, or 'all' to permit all VLAN's on the port.`,
							Computed:            true,
						},
						"drop_untagged_traffic": schema.BoolAttribute{
							MarkdownDescription: `Whether the trunk port can drop all untagged traffic.`,
							Computed:            true,
						},
						"enabled": schema.BoolAttribute{
							MarkdownDescription: `The status of the port`,
							Computed:            true,
						},
						"number": schema.Int64Attribute{
							MarkdownDescription: `Number of the port`,
							Computed:            true,
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The type of the port: 'access' or 'trunk'.`,
							Computed:            true,
						},
						"vlan": schema.Int64Attribute{
							MarkdownDescription: `Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *NetworksAppliancePortsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksAppliancePorts NetworksAppliancePorts
	diags := req.Config.Get(ctx, &networksAppliancePorts)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksAppliancePorts.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksAppliancePorts.NetworkID.IsNull(), !networksAppliancePorts.PortID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkAppliancePorts")
		vvNetworkID := networksAppliancePorts.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkAppliancePorts(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkAppliancePorts",
				err.Error(),
			)
			return
		}

		networksAppliancePorts = ResponseApplianceGetNetworkAppliancePortsItemsToBody(networksAppliancePorts, response1)
		diags = resp.State.Set(ctx, &networksAppliancePorts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkAppliancePort")
		vvNetworkID := networksAppliancePorts.NetworkID.ValueString()
		vvPortID := networksAppliancePorts.PortID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Appliance.GetNetworkAppliancePort(vvNetworkID, vvPortID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkAppliancePort",
				err.Error(),
			)
			return
		}

		networksAppliancePorts = ResponseApplianceGetNetworkAppliancePortItemToBody(networksAppliancePorts, response2)
		diags = resp.State.Set(ctx, &networksAppliancePorts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksAppliancePorts struct {
	NetworkID types.String                                     `tfsdk:"network_id"`
	PortID    types.String                                     `tfsdk:"port_id"`
	Items     *[]ResponseItemApplianceGetNetworkAppliancePorts `tfsdk:"items"`
	Item      *ResponseApplianceGetNetworkAppliancePort        `tfsdk:"item"`
}

type ResponseItemApplianceGetNetworkAppliancePorts struct {
	AccessPolicy        types.String `tfsdk:"access_policy"`
	AllowedVLANs        types.String `tfsdk:"allowed_vlans"`
	DropUntaggedTraffic types.Bool   `tfsdk:"drop_untagged_traffic"`
	Enabled             types.Bool   `tfsdk:"enabled"`
	Number              types.Int64  `tfsdk:"number"`
	Type                types.String `tfsdk:"type"`
	VLAN                types.Int64  `tfsdk:"vlan"`
}

type ResponseApplianceGetNetworkAppliancePort struct {
	AccessPolicy        types.String `tfsdk:"access_policy"`
	AllowedVLANs        types.String `tfsdk:"allowed_vlans"`
	DropUntaggedTraffic types.Bool   `tfsdk:"drop_untagged_traffic"`
	Enabled             types.Bool   `tfsdk:"enabled"`
	Number              types.Int64  `tfsdk:"number"`
	Type                types.String `tfsdk:"type"`
	VLAN                types.Int64  `tfsdk:"vlan"`
}

// ToBody
func ResponseApplianceGetNetworkAppliancePortsItemsToBody(state NetworksAppliancePorts, response *merakigosdk.ResponseApplianceGetNetworkAppliancePorts) NetworksAppliancePorts {
	var items []ResponseItemApplianceGetNetworkAppliancePorts
	for _, item := range *response {
		itemState := ResponseItemApplianceGetNetworkAppliancePorts{
			AccessPolicy: types.StringValue(item.AccessPolicy),
			AllowedVLANs: types.StringValue(item.AllowedVLANs),
			DropUntaggedTraffic: func() types.Bool {
				if item.DropUntaggedTraffic != nil {
					return types.BoolValue(*item.DropUntaggedTraffic)
				}
				return types.Bool{}
			}(),
			Enabled: func() types.Bool {
				if item.Enabled != nil {
					return types.BoolValue(*item.Enabled)
				}
				return types.Bool{}
			}(),
			Number: func() types.Int64 {
				if item.Number != nil {
					return types.Int64Value(int64(*item.Number))
				}
				return types.Int64{}
			}(),
			Type: types.StringValue(item.Type),
			VLAN: func() types.Int64 {
				if item.VLAN != nil {
					return types.Int64Value(int64(*item.VLAN))
				}
				return types.Int64{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseApplianceGetNetworkAppliancePortItemToBody(state NetworksAppliancePorts, response *merakigosdk.ResponseApplianceGetNetworkAppliancePort) NetworksAppliancePorts {
	itemState := ResponseApplianceGetNetworkAppliancePort{
		AccessPolicy: types.StringValue(response.AccessPolicy),
		AllowedVLANs: types.StringValue(response.AllowedVLANs),
		DropUntaggedTraffic: func() types.Bool {
			if response.DropUntaggedTraffic != nil {
				return types.BoolValue(*response.DropUntaggedTraffic)
			}
			return types.Bool{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Number: func() types.Int64 {
			if response.Number != nil {
				return types.Int64Value(int64(*response.Number))
			}
			return types.Int64{}
		}(),
		Type: types.StringValue(response.Type),
		VLAN: func() types.Int64 {
			if response.VLAN != nil {
				return types.Int64Value(int64(*response.VLAN))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
