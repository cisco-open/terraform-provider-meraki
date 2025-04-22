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
	_ datasource.DataSource              = &OrganizationsConfigTemplatesSwitchProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsConfigTemplatesSwitchProfilesDataSource{}
)

func NewOrganizationsConfigTemplatesSwitchProfilesDataSource() datasource.DataSource {
	return &OrganizationsConfigTemplatesSwitchProfilesDataSource{}
}

type OrganizationsConfigTemplatesSwitchProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsConfigTemplatesSwitchProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsConfigTemplatesSwitchProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_config_templates_switch_profiles"
}

func (d *OrganizationsConfigTemplatesSwitchProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"config_template_id": schema.StringAttribute{
				MarkdownDescription: `configTemplateId path parameter. Config template ID`,
				Required:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSwitchGetOrganizationConfigTemplateSwitchProfiles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"model": schema.StringAttribute{
							MarkdownDescription: `Switch model`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Switch template name`,
							Computed:            true,
						},
						"switch_profile_id": schema.StringAttribute{
							MarkdownDescription: `Switch template id`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsConfigTemplatesSwitchProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsConfigTemplatesSwitchProfiles OrganizationsConfigTemplatesSwitchProfiles
	diags := req.Config.Get(ctx, &organizationsConfigTemplatesSwitchProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationConfigTemplateSwitchProfiles")
		vvOrganizationID := organizationsConfigTemplatesSwitchProfiles.OrganizationID.ValueString()
		vvConfigTemplateID := organizationsConfigTemplatesSwitchProfiles.ConfigTemplateID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetOrganizationConfigTemplateSwitchProfiles(vvOrganizationID, vvConfigTemplateID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationConfigTemplateSwitchProfiles",
				err.Error(),
			)
			return
		}

		organizationsConfigTemplatesSwitchProfiles = ResponseSwitchGetOrganizationConfigTemplateSwitchProfilesItemsToBody(organizationsConfigTemplatesSwitchProfiles, response1)
		diags = resp.State.Set(ctx, &organizationsConfigTemplatesSwitchProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsConfigTemplatesSwitchProfiles struct {
	OrganizationID   types.String                                                     `tfsdk:"organization_id"`
	ConfigTemplateID types.String                                                     `tfsdk:"config_template_id"`
	Items            *[]ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfiles `tfsdk:"items"`
}

type ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfiles struct {
	Model           types.String `tfsdk:"model"`
	Name            types.String `tfsdk:"name"`
	SwitchProfileID types.String `tfsdk:"switch_profile_id"`
}

// ToBody
func ResponseSwitchGetOrganizationConfigTemplateSwitchProfilesItemsToBody(state OrganizationsConfigTemplatesSwitchProfiles, response *merakigosdk.ResponseSwitchGetOrganizationConfigTemplateSwitchProfiles) OrganizationsConfigTemplatesSwitchProfiles {
	var items []ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfiles
	for _, item := range *response {
		itemState := ResponseItemSwitchGetOrganizationConfigTemplateSwitchProfiles{
			Model:           types.StringValue(item.Model),
			Name:            types.StringValue(item.Name),
			SwitchProfileID: types.StringValue(item.SwitchProfileID),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
