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
	_ datasource.DataSource              = &NetworksApplianceWarmSpareDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceWarmSpareDataSource{}
)

func NewNetworksApplianceWarmSpareDataSource() datasource.DataSource {
	return &NetworksApplianceWarmSpareDataSource{}
}

type NetworksApplianceWarmSpareDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceWarmSpareDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceWarmSpareDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_warm_spare"
}

func (d *NetworksApplianceWarmSpareDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						MarkdownDescription: `Is the warm spare enabled`,
						Computed:            true,
					},
					"primary_serial": schema.StringAttribute{
						MarkdownDescription: `Serial number of the primary appliance`,
						Computed:            true,
					},
					"spare_serial": schema.StringAttribute{
						MarkdownDescription: `Serial number of the warm spare appliance`,
						Computed:            true,
					},
					"uplink_mode": schema.StringAttribute{
						MarkdownDescription: `Uplink mode, either virtual or public`,
						Computed:            true,
					},
					"wan1": schema.SingleNestedAttribute{
						MarkdownDescription: `WAN 1 IP and subnet`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"ip": schema.StringAttribute{
								MarkdownDescription: `IP address used for WAN 1`,
								Computed:            true,
							},
							"subnet": schema.StringAttribute{
								MarkdownDescription: `Subnet used for WAN 1`,
								Computed:            true,
							},
						},
					},
					"wan2": schema.SingleNestedAttribute{
						MarkdownDescription: `WAN 2 IP and subnet`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"ip": schema.StringAttribute{
								MarkdownDescription: `IP address used for WAN 2`,
								Computed:            true,
							},
							"subnet": schema.StringAttribute{
								MarkdownDescription: `Subnet used for WAN 2`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceWarmSpareDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceWarmSpare NetworksApplianceWarmSpare
	diags := req.Config.Get(ctx, &networksApplianceWarmSpare)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceWarmSpare")
		vvNetworkID := networksApplianceWarmSpare.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceWarmSpare(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceWarmSpare",
				err.Error(),
			)
			return
		}

		networksApplianceWarmSpare = ResponseApplianceGetNetworkApplianceWarmSpareItemToBody(networksApplianceWarmSpare, response1)
		diags = resp.State.Set(ctx, &networksApplianceWarmSpare)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceWarmSpare struct {
	NetworkID types.String                                   `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceWarmSpare `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceWarmSpare struct {
	Enabled       types.Bool                                         `tfsdk:"enabled"`
	PrimarySerial types.String                                       `tfsdk:"primary_serial"`
	SpareSerial   types.String                                       `tfsdk:"spare_serial"`
	UplinkMode    types.String                                       `tfsdk:"uplink_mode"`
	Wan1          *ResponseApplianceGetNetworkApplianceWarmSpareWan1 `tfsdk:"wan1"`
	Wan2          *ResponseApplianceGetNetworkApplianceWarmSpareWan2 `tfsdk:"wan2"`
}

type ResponseApplianceGetNetworkApplianceWarmSpareWan1 struct {
	IP     types.String `tfsdk:"ip"`
	Subnet types.String `tfsdk:"subnet"`
}

type ResponseApplianceGetNetworkApplianceWarmSpareWan2 struct {
	IP     types.String `tfsdk:"ip"`
	Subnet types.String `tfsdk:"subnet"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceWarmSpareItemToBody(state NetworksApplianceWarmSpare, response *merakigosdk.ResponseApplianceGetNetworkApplianceWarmSpare) NetworksApplianceWarmSpare {
	itemState := ResponseApplianceGetNetworkApplianceWarmSpare{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		PrimarySerial: types.StringValue(response.PrimarySerial),
		SpareSerial:   types.StringValue(response.SpareSerial),
		UplinkMode:    types.StringValue(response.UplinkMode),
		Wan1: func() *ResponseApplianceGetNetworkApplianceWarmSpareWan1 {
			if response.Wan1 != nil {
				return &ResponseApplianceGetNetworkApplianceWarmSpareWan1{
					IP:     types.StringValue(response.Wan1.IP),
					Subnet: types.StringValue(response.Wan1.Subnet),
				}
			}
			return nil
		}(),
		Wan2: func() *ResponseApplianceGetNetworkApplianceWarmSpareWan2 {
			if response.Wan2 != nil {
				return &ResponseApplianceGetNetworkApplianceWarmSpareWan2{
					IP:     types.StringValue(response.Wan2.IP),
					Subnet: types.StringValue(response.Wan2.Subnet),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
