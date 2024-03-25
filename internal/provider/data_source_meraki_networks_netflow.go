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
	_ datasource.DataSource              = &NetworksNetflowDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksNetflowDataSource{}
)

func NewNetworksNetflowDataSource() datasource.DataSource {
	return &NetworksNetflowDataSource{}
}

type NetworksNetflowDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksNetflowDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksNetflowDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_netflow"
}

func (d *NetworksNetflowDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"collector_ip": schema.StringAttribute{
						Computed: true,
					},
					"collector_port": schema.Int64Attribute{
						Computed: true,
					},
					"eta_dst_port": schema.Int64Attribute{
						Computed: true,
					},
					"eta_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"reporting_enabled": schema.BoolAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *NetworksNetflowDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksNetflow NetworksNetflow
	diags := req.Config.Get(ctx, &networksNetflow)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkNetflow")
		vvNetworkID := networksNetflow.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkNetflow(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkNetflow",
				err.Error(),
			)
			return
		}

		networksNetflow = ResponseNetworksGetNetworkNetflowItemToBody(networksNetflow, response1)
		diags = resp.State.Set(ctx, &networksNetflow)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksNetflow struct {
	NetworkID types.String                       `tfsdk:"network_id"`
	Item      *ResponseNetworksGetNetworkNetflow `tfsdk:"item"`
}

type ResponseNetworksGetNetworkNetflow struct {
	CollectorIP      types.String `tfsdk:"collector_ip"`
	CollectorPort    types.Int64  `tfsdk:"collector_port"`
	EtaDstPort       types.Int64  `tfsdk:"eta_dst_port"`
	EtaEnabled       types.Bool   `tfsdk:"eta_enabled"`
	ReportingEnabled types.Bool   `tfsdk:"reporting_enabled"`
}

// ToBody
func ResponseNetworksGetNetworkNetflowItemToBody(state NetworksNetflow, response *merakigosdk.ResponseNetworksGetNetworkNetflow) NetworksNetflow {
	itemState := ResponseNetworksGetNetworkNetflow{
		CollectorIP: types.StringValue(response.CollectorIP),
		CollectorPort: func() types.Int64 {
			if response.CollectorPort != nil {
				return types.Int64Value(int64(*response.CollectorPort))
			}
			return types.Int64{}
		}(),
		EtaDstPort: func() types.Int64 {
			if response.EtaDstPort != nil {
				return types.Int64Value(int64(*response.EtaDstPort))
			}
			return types.Int64{}
		}(),
		EtaEnabled: func() types.Bool {
			if response.EtaEnabled != nil {
				return types.BoolValue(*response.EtaEnabled)
			}
			return types.Bool{}
		}(),
		ReportingEnabled: func() types.Bool {
			if response.ReportingEnabled != nil {
				return types.BoolValue(*response.ReportingEnabled)
			}
			return types.Bool{}
		}(),
	}
	state.Item = &itemState
	return state
}
