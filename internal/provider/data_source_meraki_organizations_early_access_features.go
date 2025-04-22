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
	_ datasource.DataSource              = &OrganizationsEarlyAccessFeaturesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsEarlyAccessFeaturesDataSource{}
)

func NewOrganizationsEarlyAccessFeaturesDataSource() datasource.DataSource {
	return &OrganizationsEarlyAccessFeaturesDataSource{}
}

type OrganizationsEarlyAccessFeaturesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsEarlyAccessFeaturesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsEarlyAccessFeaturesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_early_access_features"
}

func (d *OrganizationsEarlyAccessFeaturesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationEarlyAccessFeatures`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"descriptions": schema.SingleNestedAttribute{
							MarkdownDescription: `Descriptions of the early access feature`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"long": schema.StringAttribute{
									MarkdownDescription: `Long description`,
									Computed:            true,
								},
								"short": schema.StringAttribute{
									MarkdownDescription: `Short description`,
									Computed:            true,
								},
							},
						},
						"documentation_link": schema.StringAttribute{
							MarkdownDescription: `Link to the documentation of this early access feature`,
							Computed:            true,
						},
						"is_org_scoped_only": schema.BoolAttribute{
							MarkdownDescription: `If this early access feature can only be opted in for the entire organization`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the early access feature`,
							Computed:            true,
						},
						"short_name": schema.StringAttribute{
							MarkdownDescription: `Short name of the early access feature`,
							Computed:            true,
						},
						"support_link": schema.StringAttribute{
							MarkdownDescription: `Link to get support for this early access feature`,
							Computed:            true,
						},
						"topic": schema.StringAttribute{
							MarkdownDescription: `Topic of the early access feature`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsEarlyAccessFeaturesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsEarlyAccessFeatures OrganizationsEarlyAccessFeatures
	diags := req.Config.Get(ctx, &organizationsEarlyAccessFeatures)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationEarlyAccessFeatures")
		vvOrganizationID := organizationsEarlyAccessFeatures.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationEarlyAccessFeatures(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationEarlyAccessFeatures",
				err.Error(),
			)
			return
		}

		organizationsEarlyAccessFeatures = ResponseOrganizationsGetOrganizationEarlyAccessFeaturesItemsToBody(organizationsEarlyAccessFeatures, response1)
		diags = resp.State.Set(ctx, &organizationsEarlyAccessFeatures)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsEarlyAccessFeatures struct {
	OrganizationID types.String                                                   `tfsdk:"organization_id"`
	Items          *[]ResponseItemOrganizationsGetOrganizationEarlyAccessFeatures `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationEarlyAccessFeatures struct {
	Descriptions      *ResponseItemOrganizationsGetOrganizationEarlyAccessFeaturesDescriptions `tfsdk:"descriptions"`
	DocumentationLink types.String                                                             `tfsdk:"documentation_link"`
	IsOrgScopedOnly   types.Bool                                                               `tfsdk:"is_org_scoped_only"`
	Name              types.String                                                             `tfsdk:"name"`
	ShortName         types.String                                                             `tfsdk:"short_name"`
	SupportLink       types.String                                                             `tfsdk:"support_link"`
	Topic             types.String                                                             `tfsdk:"topic"`
}

type ResponseItemOrganizationsGetOrganizationEarlyAccessFeaturesDescriptions struct {
	Long  types.String `tfsdk:"long"`
	Short types.String `tfsdk:"short"`
}

// ToBody
func ResponseOrganizationsGetOrganizationEarlyAccessFeaturesItemsToBody(state OrganizationsEarlyAccessFeatures, response *merakigosdk.ResponseOrganizationsGetOrganizationEarlyAccessFeatures) OrganizationsEarlyAccessFeatures {
	var items []ResponseItemOrganizationsGetOrganizationEarlyAccessFeatures
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationEarlyAccessFeatures{
			Descriptions: func() *ResponseItemOrganizationsGetOrganizationEarlyAccessFeaturesDescriptions {
				if item.Descriptions != nil {
					return &ResponseItemOrganizationsGetOrganizationEarlyAccessFeaturesDescriptions{
						Long:  types.StringValue(item.Descriptions.Long),
						Short: types.StringValue(item.Descriptions.Short),
					}
				}
				return nil
			}(),
			DocumentationLink: types.StringValue(item.DocumentationLink),
			IsOrgScopedOnly: func() types.Bool {
				if item.IsOrgScopedOnly != nil {
					return types.BoolValue(*item.IsOrgScopedOnly)
				}
				return types.Bool{}
			}(),
			Name:        types.StringValue(item.Name),
			ShortName:   types.StringValue(item.ShortName),
			SupportLink: types.StringValue(item.SupportLink),
			Topic:       types.StringValue(item.Topic),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
