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
	_ datasource.DataSource              = &NetworksSyslogServersDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksSyslogServersDataSource{}
)

func NewNetworksSyslogServersDataSource() datasource.DataSource {
	return &NetworksSyslogServersDataSource{}
}

type NetworksSyslogServersDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksSyslogServersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksSyslogServersDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_syslog_servers"
}

func (d *NetworksSyslogServersDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"servers": schema.SetNestedAttribute{
						MarkdownDescription: `List of the syslog servers for this network`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"host": schema.StringAttribute{
									MarkdownDescription: `The IP address or FQDN of the syslog server`,
									Computed:            true,
								},
								"port": schema.StringAttribute{
									MarkdownDescription: `The port of the syslog server`,
									Computed:            true,
								},
								"roles": schema.ListAttribute{
									MarkdownDescription: `A list of roles for the syslog server. Options (case-insensitive): 'Wireless event log', 'Appliance event log', 'Switch event log', 'Air Marshal events', 'Flows', 'URLs', 'IDS alerts', 'Security events'`,
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksSyslogServersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksSyslogServers NetworksSyslogServers
	diags := req.Config.Get(ctx, &networksSyslogServers)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkSyslogServers")
		vvNetworkID := networksSyslogServers.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkSyslogServers(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSyslogServers",
				err.Error(),
			)
			return
		}

		networksSyslogServers = ResponseNetworksGetNetworkSyslogServersItemToBody(networksSyslogServers, response1)
		diags = resp.State.Set(ctx, &networksSyslogServers)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksSyslogServers struct {
	NetworkID types.String                             `tfsdk:"network_id"`
	Item      *ResponseNetworksGetNetworkSyslogServers `tfsdk:"item"`
}

type ResponseNetworksGetNetworkSyslogServers struct {
	Servers *[]ResponseNetworksGetNetworkSyslogServersServers `tfsdk:"servers"`
}

type ResponseNetworksGetNetworkSyslogServersServers struct {
	Host  types.String `tfsdk:"host"`
	Port  types.String `tfsdk:"port"`
	Roles types.List   `tfsdk:"roles"`
}

// ToBody
func ResponseNetworksGetNetworkSyslogServersItemToBody(state NetworksSyslogServers, response *merakigosdk.ResponseNetworksGetNetworkSyslogServers) NetworksSyslogServers {
	itemState := ResponseNetworksGetNetworkSyslogServers{
		Servers: func() *[]ResponseNetworksGetNetworkSyslogServersServers {
			if response.Servers != nil {
				result := make([]ResponseNetworksGetNetworkSyslogServersServers, len(*response.Servers))
				for i, servers := range *response.Servers {
					result[i] = ResponseNetworksGetNetworkSyslogServersServers{
						Host:  types.StringValue(servers.Host),
						Port:  types.StringValue(servers.Port),
						Roles: StringSliceToList(servers.Roles),
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
