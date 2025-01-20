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
	_ datasource.DataSource              = &OrganizationsSplashThemesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSplashThemesDataSource{}
)

func NewOrganizationsSplashThemesDataSource() datasource.DataSource {
	return &OrganizationsSplashThemesDataSource{}
}

type OrganizationsSplashThemesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSplashThemesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSplashThemesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_splash_themes"
}

func (d *OrganizationsSplashThemesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationSplashThemes`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `theme id`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `theme name`,
							Computed:            true,
						},
						"theme_assets": schema.SetNestedAttribute{
							MarkdownDescription: `list of theme assets`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"id": schema.StringAttribute{
										MarkdownDescription: `asset id`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `asset name`,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsSplashThemesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSplashThemes OrganizationsSplashThemes
	diags := req.Config.Get(ctx, &organizationsSplashThemes)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSplashThemes")
		vvOrganizationID := organizationsSplashThemes.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSplashThemes(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSplashThemes",
				err.Error(),
			)
			return
		}

		organizationsSplashThemes = ResponseOrganizationsGetOrganizationSplashThemesItemsToBody(organizationsSplashThemes, response1)
		diags = resp.State.Set(ctx, &organizationsSplashThemes)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSplashThemes struct {
	OrganizationID types.String                                            `tfsdk:"organization_id"`
	Items          *[]ResponseItemOrganizationsGetOrganizationSplashThemes `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationSplashThemes struct {
	ID          types.String                                                       `tfsdk:"id"`
	Name        types.String                                                       `tfsdk:"name"`
	ThemeAssets *[]ResponseItemOrganizationsGetOrganizationSplashThemesThemeAssets `tfsdk:"theme_assets"`
}

type ResponseItemOrganizationsGetOrganizationSplashThemesThemeAssets struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSplashThemesItemsToBody(state OrganizationsSplashThemes, response *merakigosdk.ResponseOrganizationsGetOrganizationSplashThemes) OrganizationsSplashThemes {
	var items []ResponseItemOrganizationsGetOrganizationSplashThemes
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationSplashThemes{
			ID:   types.StringValue(item.ID),
			Name: types.StringValue(item.Name),
			ThemeAssets: func() *[]ResponseItemOrganizationsGetOrganizationSplashThemesThemeAssets {
				if item.ThemeAssets != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationSplashThemesThemeAssets, len(*item.ThemeAssets))
					for i, themeAssets := range *item.ThemeAssets {
						result[i] = ResponseItemOrganizationsGetOrganizationSplashThemesThemeAssets{
							ID:   types.StringValue(themeAssets.ID),
							Name: types.StringValue(themeAssets.Name),
						}
					}
					return &result
				}
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
