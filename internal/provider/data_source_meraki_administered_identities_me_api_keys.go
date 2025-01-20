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
	_ datasource.DataSource              = &AdministeredIDentitiesMeAPIKeysDataSource{}
	_ datasource.DataSourceWithConfigure = &AdministeredIDentitiesMeAPIKeysDataSource{}
)

func NewAdministeredIDentitiesMeAPIKeysDataSource() datasource.DataSource {
	return &AdministeredIDentitiesMeAPIKeysDataSource{}
}

type AdministeredIDentitiesMeAPIKeysDataSource struct {
	client *merakigosdk.Client
}

func (d *AdministeredIDentitiesMeAPIKeysDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *AdministeredIDentitiesMeAPIKeysDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_administered_identities_me_api_keys"
}

func (d *AdministeredIDentitiesMeAPIKeysDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseAdministeredGetAdministeredIdentitiesMeApiKeys`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"created_at": schema.StringAttribute{
							MarkdownDescription: `Date that the API key was created`,
							Computed:            true,
						},
						"suffix": schema.StringAttribute{
							MarkdownDescription: `Last 4 characters of the API key`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *AdministeredIDentitiesMeAPIKeysDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var administeredIDentitiesMeAPIKeys AdministeredIDentitiesMeAPIKeys
	diags := req.Config.Get(ctx, &administeredIDentitiesMeAPIKeys)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAdministeredIDentitiesMeAPIKeys")

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Administered.GetAdministeredIDentitiesMeAPIKeys()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetAdministeredIDentitiesMeAPIKeys",
				err.Error(),
			)
			return
		}

		administeredIDentitiesMeAPIKeys = ResponseAdministeredGetAdministeredIDentitiesMeAPIKeysItemsToBody(administeredIDentitiesMeAPIKeys, response1)
		diags = resp.State.Set(ctx, &administeredIDentitiesMeAPIKeys)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type AdministeredIDentitiesMeAPIKeys struct {
	Items *[]ResponseItemAdministeredGetAdministeredIdentitiesMeApiKeys `tfsdk:"items"`
}

type ResponseItemAdministeredGetAdministeredIdentitiesMeApiKeys struct {
	CreatedAt types.String `tfsdk:"created_at"`
	Suffix    types.String `tfsdk:"suffix"`
}

// ToBody
func ResponseAdministeredGetAdministeredIDentitiesMeAPIKeysItemsToBody(state AdministeredIDentitiesMeAPIKeys, response *merakigosdk.ResponseAdministeredGetAdministeredIDentitiesMeAPIKeys) AdministeredIDentitiesMeAPIKeys {
	var items []ResponseItemAdministeredGetAdministeredIdentitiesMeApiKeys
	for _, item := range *response {
		itemState := ResponseItemAdministeredGetAdministeredIdentitiesMeApiKeys{
			CreatedAt: types.StringValue(item.CreatedAt),
			Suffix:    types.StringValue(item.Suffix),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
