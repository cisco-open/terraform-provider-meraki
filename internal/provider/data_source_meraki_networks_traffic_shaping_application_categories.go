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
	_ datasource.DataSource              = &NetworksTrafficShapingApplicationCategoriesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksTrafficShapingApplicationCategoriesDataSource{}
)

func NewNetworksTrafficShapingApplicationCategoriesDataSource() datasource.DataSource {
	return &NetworksTrafficShapingApplicationCategoriesDataSource{}
}

type NetworksTrafficShapingApplicationCategoriesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksTrafficShapingApplicationCategoriesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksTrafficShapingApplicationCategoriesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_traffic_shaping_application_categories"
}

func (d *NetworksTrafficShapingApplicationCategoriesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"applications": schema.SetNestedAttribute{
									Computed: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												Computed: true,
											},
											"name": schema.StringAttribute{
												Computed: true,
											},
										},
									},
								},
								"id": schema.StringAttribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
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

func (d *NetworksTrafficShapingApplicationCategoriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksTrafficShapingApplicationCategories NetworksTrafficShapingApplicationCategories
	diags := req.Config.Get(ctx, &networksTrafficShapingApplicationCategories)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkTrafficShapingApplicationCategories")
		vvNetworkID := networksTrafficShapingApplicationCategories.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkTrafficShapingApplicationCategories(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkTrafficShapingApplicationCategories",
				err.Error(),
			)
			return
		}

		networksTrafficShapingApplicationCategories = ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesItemToBody(networksTrafficShapingApplicationCategories, response1)
		diags = resp.State.Set(ctx, &networksTrafficShapingApplicationCategories)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksTrafficShapingApplicationCategories struct {
	NetworkID types.String                                                   `tfsdk:"network_id"`
	Item      *ResponseNetworksGetNetworkTrafficShapingApplicationCategories `tfsdk:"item"`
}

type ResponseNetworksGetNetworkTrafficShapingApplicationCategories struct {
	ApplicationCategories *[]ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategories `tfsdk:"application_categories"`
}

type ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategories struct {
	Applications *[]ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategoriesApplications `tfsdk:"applications"`
	ID           types.String                                                                                      `tfsdk:"id"`
	Name         types.String                                                                                      `tfsdk:"name"`
}

type ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategoriesApplications struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesItemToBody(state NetworksTrafficShapingApplicationCategories, response *merakigosdk.ResponseNetworksGetNetworkTrafficShapingApplicationCategories) NetworksTrafficShapingApplicationCategories {
	itemState := ResponseNetworksGetNetworkTrafficShapingApplicationCategories{
		ApplicationCategories: func() *[]ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategories {
			if response.ApplicationCategories != nil {
				result := make([]ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategories, len(*response.ApplicationCategories))
				for i, applicationCategories := range *response.ApplicationCategories {
					result[i] = ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategories{
						Applications: func() *[]ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategoriesApplications {
							if applicationCategories.Applications != nil {
								result := make([]ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategoriesApplications, len(*applicationCategories.Applications))
								for i, applications := range *applicationCategories.Applications {
									result[i] = ResponseNetworksGetNetworkTrafficShapingApplicationCategoriesApplicationCategoriesApplications{
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
