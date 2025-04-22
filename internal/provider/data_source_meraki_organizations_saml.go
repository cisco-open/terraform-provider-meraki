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
	_ datasource.DataSource              = &OrganizationsSamlDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSamlDataSource{}
)

func NewOrganizationsSamlDataSource() datasource.DataSource {
	return &OrganizationsSamlDataSource{}
}

type OrganizationsSamlDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSamlDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSamlDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_saml"
}

func (d *OrganizationsSamlDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Toggle depicting if SAML SSO settings are enabled`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *OrganizationsSamlDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSaml OrganizationsSaml
	diags := req.Config.Get(ctx, &organizationsSaml)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSaml")
		vvOrganizationID := organizationsSaml.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSaml(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSaml",
				err.Error(),
			)
			return
		}

		organizationsSaml = ResponseOrganizationsGetOrganizationSamlItemToBody(organizationsSaml, response1)
		diags = resp.State.Set(ctx, &organizationsSaml)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSaml struct {
	OrganizationID types.String                              `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsGetOrganizationSaml `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationSaml struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSamlItemToBody(state OrganizationsSaml, response *merakigosdk.ResponseOrganizationsGetOrganizationSaml) OrganizationsSaml {
	itemState := ResponseOrganizationsGetOrganizationSaml{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
	}
	state.Item = &itemState
	return state
}
