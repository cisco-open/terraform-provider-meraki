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
	_ datasource.DataSource              = &NetworksApplianceFirewallL7FirewallRulesApplicationCategoriesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceFirewallL7FirewallRulesApplicationCategoriesDataSource{}
)

func NewNetworksApplianceFirewallL7FirewallRulesApplicationCategoriesDataSource() datasource.DataSource {
	return &NetworksApplianceFirewallL7FirewallRulesApplicationCategoriesDataSource{}
}

type NetworksApplianceFirewallL7FirewallRulesApplicationCategoriesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceFirewallL7FirewallRulesApplicationCategoriesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceFirewallL7FirewallRulesApplicationCategoriesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_l7_firewall_rules_application_categories"
}

func (d *NetworksApplianceFirewallL7FirewallRulesApplicationCategoriesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"application_categories": schema.SetNestedAttribute{
						MarkdownDescription: ` The L7 firewall application categories and their associated applications for an MX network`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"applications": schema.SetNestedAttribute{
									MarkdownDescription: `Details of the associated applications`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												MarkdownDescription: `The id of the application`,
												Computed:            true,
											},
											"name": schema.StringAttribute{
												MarkdownDescription: `The name of the application`,
												Computed:            true,
											},
										},
									},
								},
								"id": schema.StringAttribute{
									MarkdownDescription: `The id of the category`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the category`,
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

func (d *NetworksApplianceFirewallL7FirewallRulesApplicationCategoriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceFirewallL7FirewallRulesApplicationCategories NetworksApplianceFirewallL7FirewallRulesApplicationCategories
	diags := req.Config.Get(ctx, &networksApplianceFirewallL7FirewallRulesApplicationCategories)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories")
		vvNetworkID := networksApplianceFirewallL7FirewallRulesApplicationCategories.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallL7FirewallRulesApplicationCategories",
				err.Error(),
			)
			return
		}

		networksApplianceFirewallL7FirewallRulesApplicationCategories = ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesItemToBody(networksApplianceFirewallL7FirewallRulesApplicationCategories, response1)
		diags = resp.State.Set(ctx, &networksApplianceFirewallL7FirewallRulesApplicationCategories)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceFirewallL7FirewallRulesApplicationCategories struct {
	NetworkID types.String                                                                      `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories struct {
	ApplicationCategories *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategories `tfsdk:"application_categories"`
}

type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategories struct {
	Applications *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategoriesApplications `tfsdk:"applications"`
	ID           types.String                                                                                                         `tfsdk:"id"`
	Name         types.String                                                                                                         `tfsdk:"name"`
}

type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategoriesApplications struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesItemToBody(state NetworksApplianceFirewallL7FirewallRulesApplicationCategories, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories) NetworksApplianceFirewallL7FirewallRulesApplicationCategories {
	itemState := ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategories{
		ApplicationCategories: func() *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategories {
			if response.ApplicationCategories != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategories, len(*response.ApplicationCategories))
				for i, applicationCategories := range *response.ApplicationCategories {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategories{
						Applications: func() *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategoriesApplications {
							if applicationCategories.Applications != nil {
								result := make([]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategoriesApplications, len(*applicationCategories.Applications))
								for i, applications := range *applicationCategories.Applications {
									result[i] = ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesApplicationCategoriesApplicationCategoriesApplications{
										ID:   types.StringValue(applications.ID),
										Name: types.StringValue(applications.Name),
									}
								}
								return &result
							}
							return nil
						}(),
						ID:   types.StringValue(applicationCategories.ID),
						Name: types.StringValue(applicationCategories.Name),
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
