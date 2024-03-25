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
	_ datasource.DataSource              = &NetworksApplianceContentFilteringDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksApplianceContentFilteringDataSource{}
)

func NewNetworksApplianceContentFilteringDataSource() datasource.DataSource {
	return &NetworksApplianceContentFilteringDataSource{}
}

type NetworksApplianceContentFilteringDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksApplianceContentFilteringDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksApplianceContentFilteringDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_content_filtering"
}

func (d *NetworksApplianceContentFilteringDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"allowed_url_patterns": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"blocked_url_categories": schema.SetNestedAttribute{
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
					"blocked_url_patterns": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"url_category_list_size": schema.StringAttribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *NetworksApplianceContentFilteringDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksApplianceContentFiltering NetworksApplianceContentFiltering
	diags := req.Config.Get(ctx, &networksApplianceContentFiltering)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkApplianceContentFiltering")
		vvNetworkID := networksApplianceContentFiltering.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetNetworkApplianceContentFiltering(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceContentFiltering",
				err.Error(),
			)
			return
		}

		networksApplianceContentFiltering = ResponseApplianceGetNetworkApplianceContentFilteringItemToBody(networksApplianceContentFiltering, response1)
		diags = resp.State.Set(ctx, &networksApplianceContentFiltering)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksApplianceContentFiltering struct {
	NetworkID types.String                                          `tfsdk:"network_id"`
	Item      *ResponseApplianceGetNetworkApplianceContentFiltering `tfsdk:"item"`
}

type ResponseApplianceGetNetworkApplianceContentFiltering struct {
	AllowedURLPatterns   types.List                                                                  `tfsdk:"allowed_url_patterns"`
	BlockedURLCategories *[]ResponseApplianceGetNetworkApplianceContentFilteringBlockedUrlCategories `tfsdk:"blocked_url_categories"`
	BlockedURLPatterns   types.List                                                                  `tfsdk:"blocked_url_patterns"`
	URLCategoryListSize  types.String                                                                `tfsdk:"url_category_list_size"`
}

type ResponseApplianceGetNetworkApplianceContentFilteringBlockedUrlCategories struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseApplianceGetNetworkApplianceContentFilteringItemToBody(state NetworksApplianceContentFiltering, response *merakigosdk.ResponseApplianceGetNetworkApplianceContentFiltering) NetworksApplianceContentFiltering {
	itemState := ResponseApplianceGetNetworkApplianceContentFiltering{
		AllowedURLPatterns: StringSliceToList(response.AllowedURLPatterns),
		BlockedURLCategories: func() *[]ResponseApplianceGetNetworkApplianceContentFilteringBlockedUrlCategories {
			if response.BlockedURLCategories != nil {
				result := make([]ResponseApplianceGetNetworkApplianceContentFilteringBlockedUrlCategories, len(*response.BlockedURLCategories))
				for i, blockedURLCategories := range *response.BlockedURLCategories {
					result[i] = ResponseApplianceGetNetworkApplianceContentFilteringBlockedUrlCategories{
						ID:   types.StringValue(blockedURLCategories.ID),
						Name: types.StringValue(blockedURLCategories.Name),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceContentFilteringBlockedUrlCategories{}
		}(),
		BlockedURLPatterns:  StringSliceToList(response.BlockedURLPatterns),
		URLCategoryListSize: types.StringValue(response.URLCategoryListSize),
	}
	state.Item = &itemState
	return state
}
