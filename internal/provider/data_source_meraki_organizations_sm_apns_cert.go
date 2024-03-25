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
	_ datasource.DataSource              = &OrganizationsSmApnsCertDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSmApnsCertDataSource{}
)

func NewOrganizationsSmApnsCertDataSource() datasource.DataSource {
	return &OrganizationsSmApnsCertDataSource{}
}

type OrganizationsSmApnsCertDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSmApnsCertDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSmApnsCertDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_sm_apns_cert"
}

func (d *OrganizationsSmApnsCertDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"certificate": schema.StringAttribute{
						MarkdownDescription: `Organization APNS Certificate used by devices to communication with Apple`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *OrganizationsSmApnsCertDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSmApnsCert OrganizationsSmApnsCert
	diags := req.Config.Get(ctx, &organizationsSmApnsCert)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSmApnsCert")
		vvOrganizationID := organizationsSmApnsCert.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Sm.GetOrganizationSmApnsCert(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmApnsCert",
				err.Error(),
			)
			return
		}

		organizationsSmApnsCert = ResponseSmGetOrganizationSmApnsCertItemToBody(organizationsSmApnsCert, response1)
		diags = resp.State.Set(ctx, &organizationsSmApnsCert)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSmApnsCert struct {
	OrganizationID types.String                         `tfsdk:"organization_id"`
	Item           *ResponseSmGetOrganizationSmApnsCert `tfsdk:"item"`
}

type ResponseSmGetOrganizationSmApnsCert struct {
	Certificate types.String `tfsdk:"certificate"`
}

// ToBody
func ResponseSmGetOrganizationSmApnsCertItemToBody(state OrganizationsSmApnsCert, response *merakigosdk.ResponseSmGetOrganizationSmApnsCert) OrganizationsSmApnsCert {
	itemState := ResponseSmGetOrganizationSmApnsCert{
		Certificate: types.StringValue(response.Certificate),
	}
	state.Item = &itemState
	return state
}
