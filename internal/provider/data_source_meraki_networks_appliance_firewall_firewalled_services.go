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
	_ datasource.DataSource              = &NetworksApplianceFirewallFirewalledServicesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceFirewallFirewalledServicesDataSource{}
)

func NewNetworksApplianceFirewallFirewalledServicesDataSource() datasource.DataSource {
	return &NetworksApplianceFirewallFirewalledServicesDataSource{}
}

type NetworksApplianceFirewallFirewalledServicesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceFirewallFirewalledServicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceFirewallFirewalledServicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_firewalled_services"
}

func (d *NetworksApplianceFirewallFirewalledServicesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"service": schema.StringAttribute{
				MarkdownDescription: `service path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"access": schema.StringAttribute{
						MarkdownDescription: `A string indicating the rule for which IPs are allowed to use the specified service`,
						Computed:            true,
					},
					"allowed_ips": schema.ListAttribute{
						MarkdownDescription: `An array of allowed IPs that can access the service`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"service": schema.StringAttribute{
						MarkdownDescription: `Appliance service name`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceFirewallFirewalledServicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceFirewallFirewalledServices NetworksApplianceFirewallFirewalledServices
	diags := req.Config.Get(ctx, &networksApplianceFirewallFirewalledServices)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceFirewallFirewalledService")
		vvNetworkID := networksApplianceFirewallFirewalledServices.NetworkID.ValueString()
		vvService := networksApplianceFirewallFirewalledServices.Service.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceFirewallFirewalledService(vvNetworkID, vvService)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallFirewalledService",
				err.Error(),
			)
			return
		}

		networksApplianceFirewallFirewalledServices = ResponseApplianceGetNetworkApplianceFirewallFirewalledServiceItemToBody(networksApplianceFirewallFirewalledServices, response1)
		diags = resp.State.Set(ctx, &networksApplianceFirewallFirewalledServices)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceFirewallFirewalledServices struct {
	NetworkID types.String                                                   `tfsdk:"network_id"`
	Service   types.String                                                   `tfsdk:"service"`
	Item      *ResponseApplianceGetNetworkApplianceFirewallFirewalledService `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceFirewallFirewalledService struct {
	Access     types.String `tfsdk:"access"`
	AllowedIPs types.List   `tfsdk:"allowed_ips"`
	Service    types.String `tfsdk:"service"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceFirewallFirewalledServiceItemToBody(state NetworksApplianceFirewallFirewalledServices, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallFirewalledService) NetworksApplianceFirewallFirewalledServices {
	itemState := ResponseApplianceGetNetworkApplianceFirewallFirewalledService{
		Access:     types.StringValue(response.Access),
		AllowedIPs: StringSliceToList(response.AllowedIPs),
		Service:    types.StringValue(response.Service),
	}
	state.Item = &itemState
	return state
}
