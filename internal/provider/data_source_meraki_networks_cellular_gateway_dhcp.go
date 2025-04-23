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
	_ datasource.DataSource              = &NetworksCellularGatewayDhcpDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksCellularGatewayDhcpDataSource{}
)

func NewNetworksCellularGatewayDhcpDataSource() datasource.DataSource {
	return &NetworksCellularGatewayDhcpDataSource{}
}

type NetworksCellularGatewayDhcpDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksCellularGatewayDhcpDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksCellularGatewayDhcpDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_cellular_gateway_dhcp"
}

func (d *NetworksCellularGatewayDhcpDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"dhcp_lease_time": schema.StringAttribute{
						MarkdownDescription: `DHCP Lease time for all MG in the network.`,
						Computed:            true,
					},
					"dns_custom_nameservers": schema.ListAttribute{
						MarkdownDescription: `List of fixed IPs representing the the DNS Name servers when the mode is 'custom'.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"dns_nameservers": schema.StringAttribute{
						MarkdownDescription: `DNS name servers mode for all MG in the network.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksCellularGatewayDhcpDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksCellularGatewayDhcp NetworksCellularGatewayDhcp
	diags := req.Config.Get(ctx, &networksCellularGatewayDhcp)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkCellularGatewayDhcp")
		vvNetworkID := networksCellularGatewayDhcp.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.CellularGateway.GetNetworkCellularGatewayDhcp(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCellularGatewayDhcp",
				err.Error(),
			)
			return
		}

		networksCellularGatewayDhcp = ResponseCellularGatewayGetNetworkCellularGatewayDhcpItemToBody(networksCellularGatewayDhcp, response1)
		diags = resp.State.Set(ctx, &networksCellularGatewayDhcp)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksCellularGatewayDhcp struct {
	NetworkID types.String                                          `tfsdk:"network_id"`
	Item      *ResponseCellularGatewayGetNetworkCellularGatewayDhcp `tfsdk:"item"`
}

type ResponseCellularGatewayGetNetworkCellularGatewayDhcp struct {
	DhcpLeaseTime        types.String `tfsdk:"dhcp_lease_time"`
	DNSCustomNameservers types.List   `tfsdk:"dns_custom_nameservers"`
	DNSNameservers       types.String `tfsdk:"dns_nameservers"`
}

// ToBody
func ResponseCellularGatewayGetNetworkCellularGatewayDhcpItemToBody(state NetworksCellularGatewayDhcp, response *merakigosdk.ResponseCellularGatewayGetNetworkCellularGatewayDhcp) NetworksCellularGatewayDhcp {
	itemState := ResponseCellularGatewayGetNetworkCellularGatewayDhcp{
		DhcpLeaseTime:        types.StringValue(response.DhcpLeaseTime),
		DNSCustomNameservers: StringSliceToList(response.DNSCustomNameservers),
		DNSNameservers:       types.StringValue(response.DNSNameservers),
	}
	state.Item = &itemState
	return state
}
