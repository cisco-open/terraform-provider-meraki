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
	_ datasource.DataSource              = &NetworksSNMPDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSNMPDataSource{}
)

func NewNetworksSNMPDataSource() datasource.DataSource {
	return &NetworksSNMPDataSource{}
}

type NetworksSNMPDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSNMPDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSNMPDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_snmp"
}

func (d *NetworksSNMPDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"access": schema.StringAttribute{
						Computed: true,
					},
					"users": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"passphrase": schema.StringAttribute{
									Computed: true,
								},
								"username": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSNMPDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSNMP NetworksSNMP
	diags := req.Config.Get(ctx, &networksSNMP)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSNMP")
		vvNetworkID := networksSNMP.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Networks.GetNetworkSNMP(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSNMP",
				err.Error(),
			)
			return
		}

		networksSNMP = ResponseNetworksGetNetworkSNMPItemToBody(networksSNMP, response1)
		diags = resp.State.Set(ctx, &networksSNMP)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSNMP struct {
	NetworkID types.String                    `tfsdk:"network_id"`
	Item      *ResponseNetworksGetNetworkSnmp `tfsdk:"item"`
}

type ResponseNetworksGetNetworkSnmp struct {
	Access types.String                           `tfsdk:"access"`
	Users  *[]ResponseNetworksGetNetworkSnmpUsers `tfsdk:"users"`
}

type ResponseNetworksGetNetworkSnmpUsers struct {
	Passphrase types.String `tfsdk:"passphrase"`
	Username   types.String `tfsdk:"username"`
}

// ToBody
func ResponseNetworksGetNetworkSNMPItemToBody(state NetworksSNMP, response *merakigosdk.ResponseNetworksGetNetworkSNMP) NetworksSNMP {
	itemState := ResponseNetworksGetNetworkSnmp{
		Access: types.StringValue(response.Access),
		Users: func() *[]ResponseNetworksGetNetworkSnmpUsers {
			if response.Users != nil {
				result := make([]ResponseNetworksGetNetworkSnmpUsers, len(*response.Users))
				for i, users := range *response.Users {
					result[i] = ResponseNetworksGetNetworkSnmpUsers{
						Passphrase: types.StringValue(users.Passphrase),
						Username:   types.StringValue(users.Username),
					}
				}
				return &result
			}
			return &[]ResponseNetworksGetNetworkSnmpUsers{}
		}(),
	}
	state.Item = &itemState
	return state
}
