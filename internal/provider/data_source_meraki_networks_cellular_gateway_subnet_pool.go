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
	_ datasource.DataSource              = &NetworksCellularGatewaySubnetPoolDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksCellularGatewaySubnetPoolDataSource{}
)

func NewNetworksCellularGatewaySubnetPoolDataSource() datasource.DataSource {
	return &NetworksCellularGatewaySubnetPoolDataSource{}
}

type NetworksCellularGatewaySubnetPoolDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksCellularGatewaySubnetPoolDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksCellularGatewaySubnetPoolDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_cellular_gateway_subnet_pool"
}

func (d *NetworksCellularGatewaySubnetPoolDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"cidr": schema.StringAttribute{
						MarkdownDescription: `CIDR of the pool of subnets. Each MG in this network will automatically pick a subnet from this pool.`,
						Computed:            true,
					},
					"deployment_mode": schema.StringAttribute{
						MarkdownDescription: `Deployment mode for the cellular gateways in the network. (Passthrough/Routed)`,
						Computed:            true,
					},
					"mask": schema.Int64Attribute{
						MarkdownDescription: `Mask used for the subnet of all MGs in  this network.`,
						Computed:            true,
					},
					"subnets": schema.SetNestedAttribute{
						MarkdownDescription: `List of subnets of all MGs in this network.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"appliance_ip": schema.StringAttribute{
									MarkdownDescription: `Appliance IP of the MG device.`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the MG.`,
									Computed:            true,
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Serial Number of the MG.`,
									Computed:            true,
								},
								"subnet": schema.StringAttribute{
									MarkdownDescription: `Subnet of MG device.`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksCellularGatewaySubnetPoolDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksCellularGatewaySubnetPool NetworksCellularGatewaySubnetPool
	diags := req.Config.Get(ctx, &networksCellularGatewaySubnetPool)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkCellularGatewaySubnetPool")
		vvNetworkID := networksCellularGatewaySubnetPool.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.CellularGateway.GetNetworkCellularGatewaySubnetPool(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCellularGatewaySubnetPool",
				err.Error(),
			)
			return
		}

		networksCellularGatewaySubnetPool = ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolItemToBody(networksCellularGatewaySubnetPool, response1)
		diags = resp.State.Set(ctx, &networksCellularGatewaySubnetPool)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksCellularGatewaySubnetPool struct {
	NetworkID types.String                                                `tfsdk:"network_id"`
	Item      *ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool `tfsdk:"item"`
}

type ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool struct {
	Cidr           types.String                                                         `tfsdk:"cidr"`
	DeploymentMode types.String                                                         `tfsdk:"deployment_mode"`
	Mask           types.Int64                                                          `tfsdk:"mask"`
	Subnets        *[]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets `tfsdk:"subnets"`
}

type ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets struct {
	ApplianceIP types.String `tfsdk:"appliance_ip"`
	Name        types.String `tfsdk:"name"`
	Serial      types.String `tfsdk:"serial"`
	Subnet      types.String `tfsdk:"subnet"`
}

// ToBody
func ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolItemToBody(state NetworksCellularGatewaySubnetPool, response *merakigosdk.ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool) NetworksCellularGatewaySubnetPool {
	itemState := ResponseCellularGatewayGetNetworkCellularGatewaySubnetPool{
		Cidr:           types.StringValue(response.Cidr),
		DeploymentMode: types.StringValue(response.DeploymentMode),
		Mask: func() types.Int64 {
			if response.Mask != nil {
				return types.Int64Value(int64(*response.Mask))
			}
			return types.Int64{}
		}(),
		Subnets: func() *[]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets {
			if response.Subnets != nil {
				result := make([]ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets, len(*response.Subnets))
				for i, subnets := range *response.Subnets {
					result[i] = ResponseCellularGatewayGetNetworkCellularGatewaySubnetPoolSubnets{
						ApplianceIP: types.StringValue(subnets.ApplianceIP),
						Name:        types.StringValue(subnets.Name),
						Serial:      types.StringValue(subnets.Serial),
						Subnet:      types.StringValue(subnets.Subnet),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
