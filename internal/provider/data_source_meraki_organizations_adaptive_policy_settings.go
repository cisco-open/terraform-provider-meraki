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
	_ datasource.DataSource              = &OrganizationsAdaptivePolicySettingsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAdaptivePolicySettingsDataSource{}
)

func NewOrganizationsAdaptivePolicySettingsDataSource() datasource.DataSource {
	return &OrganizationsAdaptivePolicySettingsDataSource{}
}

type OrganizationsAdaptivePolicySettingsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAdaptivePolicySettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAdaptivePolicySettingsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_adaptive_policy_settings"
}

func (d *OrganizationsAdaptivePolicySettingsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"enabled_networks": schema.ListAttribute{
						MarkdownDescription: `List of network IDs with adaptive policy enabled`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}

func (d *OrganizationsAdaptivePolicySettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAdaptivePolicySettings OrganizationsAdaptivePolicySettings
	diags := req.Config.Get(ctx, &organizationsAdaptivePolicySettings)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAdaptivePolicySettings")
		vvOrganizationID := organizationsAdaptivePolicySettings.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAdaptivePolicySettings(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicySettings",
				err.Error(),
			)
			return
		}

		organizationsAdaptivePolicySettings = ResponseOrganizationsGetOrganizationAdaptivePolicySettingsItemToBody(organizationsAdaptivePolicySettings, response1)
		diags = resp.State.Set(ctx, &organizationsAdaptivePolicySettings)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAdaptivePolicySettings struct {
	OrganizationID types.String                                                `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsGetOrganizationAdaptivePolicySettings `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicySettings struct {
	EnabledNetworks types.List `tfsdk:"enabled_networks"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAdaptivePolicySettingsItemToBody(state OrganizationsAdaptivePolicySettings, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicySettings) OrganizationsAdaptivePolicySettings {
	itemState := ResponseOrganizationsGetOrganizationAdaptivePolicySettings{
		EnabledNetworks: StringSliceToList(response.EnabledNetworks),
	}
	state.Item = &itemState
	return state
}
