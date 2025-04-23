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
	_ datasource.DataSource              = &OrganizationsConfigTemplatesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsConfigTemplatesDataSource{}
)

func NewOrganizationsConfigTemplatesDataSource() datasource.DataSource {
	return &OrganizationsConfigTemplatesDataSource{}
}

type OrganizationsConfigTemplatesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsConfigTemplatesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsConfigTemplatesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_config_templates"
}

func (d *OrganizationsConfigTemplatesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"config_template_id": schema.StringAttribute{
				MarkdownDescription: `configTemplateId path parameter. Config template ID`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `The ID of the network or config template to copy configuration from`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the configuration template`,
						Computed:            true,
					},
					"product_types": schema.ListAttribute{
						MarkdownDescription: `The product types of the configuration template`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"time_zone": schema.StringAttribute{
						MarkdownDescription: `The timezone of the configuration template. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article</a>. Not applicable if copying from existing network or template`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationConfigTemplates`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `The ID of the network or config template to copy configuration from`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the configuration template`,
							Computed:            true,
						},
						"product_types": schema.ListAttribute{
							MarkdownDescription: `The product types of the configuration template`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"time_zone": schema.StringAttribute{
							MarkdownDescription: `The timezone of the configuration template. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article</a>. Not applicable if copying from existing network or template`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsConfigTemplatesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsConfigTemplates OrganizationsConfigTemplates
	diags := req.Config.Get(ctx, &organizationsConfigTemplates)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsConfigTemplates.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsConfigTemplates.OrganizationID.IsNull(), !organizationsConfigTemplates.ConfigTemplateID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationConfigTemplates")
		vvOrganizationID := organizationsConfigTemplates.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationConfigTemplates(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationConfigTemplates",
				err.Error(),
			)
			return
		}

		organizationsConfigTemplates = ResponseOrganizationsGetOrganizationConfigTemplatesItemsToBody(organizationsConfigTemplates, response1)
		diags = resp.State.Set(ctx, &organizationsConfigTemplates)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationConfigTemplate")
		vvOrganizationID := organizationsConfigTemplates.OrganizationID.ValueString()
		vvConfigTemplateID := organizationsConfigTemplates.ConfigTemplateID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Organizations.GetOrganizationConfigTemplate(vvOrganizationID, vvConfigTemplateID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationConfigTemplate",
				err.Error(),
			)
			return
		}

		organizationsConfigTemplates = ResponseOrganizationsGetOrganizationConfigTemplateItemToBody(organizationsConfigTemplates, response2)
		diags = resp.State.Set(ctx, &organizationsConfigTemplates)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsConfigTemplates struct {
	OrganizationID   types.String                                               `tfsdk:"organization_id"`
	ConfigTemplateID types.String                                               `tfsdk:"config_template_id"`
	Items            *[]ResponseItemOrganizationsGetOrganizationConfigTemplates `tfsdk:"items"`
	Item             *ResponseOrganizationsGetOrganizationConfigTemplate        `tfsdk:"item"`
}

type ResponseItemOrganizationsGetOrganizationConfigTemplates struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	ProductTypes types.List   `tfsdk:"product_types"`
	TimeZone     types.String `tfsdk:"time_zone"`
}

type ResponseOrganizationsGetOrganizationConfigTemplate struct {
	ID           types.String `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	ProductTypes types.List   `tfsdk:"product_types"`
	TimeZone     types.String `tfsdk:"time_zone"`
}

// ToBody
func ResponseOrganizationsGetOrganizationConfigTemplatesItemsToBody(state OrganizationsConfigTemplates, response *merakigosdk.ResponseOrganizationsGetOrganizationConfigTemplates) OrganizationsConfigTemplates {
	var items []ResponseItemOrganizationsGetOrganizationConfigTemplates
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationConfigTemplates{
			ID:           types.StringValue(item.ID),
			Name:         types.StringValue(item.Name),
			ProductTypes: StringSliceToList(item.ProductTypes),
			TimeZone:     types.StringValue(item.TimeZone),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseOrganizationsGetOrganizationConfigTemplateItemToBody(state OrganizationsConfigTemplates, response *merakigosdk.ResponseOrganizationsGetOrganizationConfigTemplate) OrganizationsConfigTemplates {
	itemState := ResponseOrganizationsGetOrganizationConfigTemplate{
		ID:           types.StringValue(response.ID),
		Name:         types.StringValue(response.Name),
		ProductTypes: StringSliceToList(response.ProductTypes),
		TimeZone:     types.StringValue(response.TimeZone),
	}
	state.Item = &itemState
	return state
}
