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
	_ datasource.DataSource              = &NetworksApplianceContentFilteringCategoriesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceContentFilteringCategoriesDataSource{}
)

func NewNetworksApplianceContentFilteringCategoriesDataSource() datasource.DataSource {
	return &NetworksApplianceContentFilteringCategoriesDataSource{}
}

type NetworksApplianceContentFilteringCategoriesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceContentFilteringCategoriesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceContentFilteringCategoriesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_content_filtering_categories"
}

func (d *NetworksApplianceContentFilteringCategoriesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"categories": schema.SetNestedAttribute{
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
				},
			},
		},
	}
}

func (d *NetworksApplianceContentFilteringCategoriesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceContentFilteringCategories NetworksApplianceContentFilteringCategories
	diags := req.Config.Get(ctx, &networksApplianceContentFilteringCategories)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceContentFilteringCategories")
		vvNetworkID := networksApplianceContentFilteringCategories.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceContentFilteringCategories(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceContentFilteringCategories",
				err.Error(),
			)
			return
		}

		networksApplianceContentFilteringCategories = ResponseApplianceGetNetworkApplianceContentFilteringCategoriesItemToBody(networksApplianceContentFilteringCategories, response1)
		diags = resp.State.Set(ctx, &networksApplianceContentFilteringCategories)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceContentFilteringCategories struct {
	NetworkID types.String                                                    `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceContentFilteringCategories `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceContentFilteringCategories struct {
	Categories *[]ResponseApplianceGetNetworkApplianceContentFilteringCategoriesCategories `tfsdk:"categories"`
}

type ResponseApplianceGetNetworkApplianceContentFilteringCategoriesCategories struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceContentFilteringCategoriesItemToBody(state NetworksApplianceContentFilteringCategories, response *merakigosdk.ResponseApplianceGetNetworkApplianceContentFilteringCategories) NetworksApplianceContentFilteringCategories {
	itemState := ResponseApplianceGetNetworkApplianceContentFilteringCategories{
		Categories: func() *[]ResponseApplianceGetNetworkApplianceContentFilteringCategoriesCategories {
			if response.Categories != nil {
				result := make([]ResponseApplianceGetNetworkApplianceContentFilteringCategoriesCategories, len(*response.Categories))
				for i, categories := range *response.Categories {
					result[i] = ResponseApplianceGetNetworkApplianceContentFilteringCategoriesCategories{
						ID:   types.StringValue(categories.ID),
						Name: types.StringValue(categories.Name),
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
