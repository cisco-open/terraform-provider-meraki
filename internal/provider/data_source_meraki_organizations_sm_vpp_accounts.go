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
	_ datasource.DataSource              = &OrganizationsSmVppAccountsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSmVppAccountsDataSource{}
)

func NewOrganizationsSmVppAccountsDataSource() datasource.DataSource {
	return &OrganizationsSmVppAccountsDataSource{}
}

type OrganizationsSmVppAccountsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSmVppAccountsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSmVppAccountsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_sm_vpp_accounts"
}

func (d *OrganizationsSmVppAccountsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"vpp_account_id": schema.StringAttribute{
				MarkdownDescription: `vppAccountId path parameter. Vpp account ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `The id of the VPP Account`,
						Computed:            true,
					},
					"vpp_service_token": schema.StringAttribute{
						MarkdownDescription: `The VPP Account's Service Token`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSmGetOrganizationSmVppAccounts`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `The id of the VPP Account`,
							Computed:            true,
						},
						"vpp_service_token": schema.StringAttribute{
							MarkdownDescription: `The VPP Account's Service Token`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsSmVppAccountsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSmVppAccounts OrganizationsSmVppAccounts
	diags := req.Config.Get(ctx, &organizationsSmVppAccounts)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsSmVppAccounts.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsSmVppAccounts.OrganizationID.IsNull(), !organizationsSmVppAccounts.VppAccountID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSmVppAccounts")
		vvOrganizationID := organizationsSmVppAccounts.OrganizationID.ValueString()

		response1, restyResp1, err := d.client.Sm.GetOrganizationSmVppAccounts(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmVppAccounts",
				err.Error(),
			)
			return
		}

		organizationsSmVppAccounts = ResponseSmGetOrganizationSmVppAccountsItemsToBody(organizationsSmVppAccounts, response1)
		diags = resp.State.Set(ctx, &organizationsSmVppAccounts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSmVppAccount")
		vvOrganizationID := organizationsSmVppAccounts.OrganizationID.ValueString()
		vvVppAccountID := organizationsSmVppAccounts.VppAccountID.ValueString()

		response2, restyResp2, err := d.client.Sm.GetOrganizationSmVppAccount(vvOrganizationID, vvVppAccountID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmVppAccount",
				err.Error(),
			)
			return
		}

		organizationsSmVppAccounts = ResponseSmGetOrganizationSmVppAccountItemToBody(organizationsSmVppAccounts, response2)
		diags = resp.State.Set(ctx, &organizationsSmVppAccounts)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSmVppAccounts struct {
	OrganizationID types.String                                  `tfsdk:"organization_id"`
	VppAccountID   types.String                                  `tfsdk:"vpp_account_id"`
	Items          *[]ResponseItemSmGetOrganizationSmVppAccounts `tfsdk:"items"`
	Item           *ResponseSmGetOrganizationSmVppAccount        `tfsdk:"item"`
}

type ResponseItemSmGetOrganizationSmVppAccounts struct {
	ID              types.String `tfsdk:"id"`
	VppServiceToken types.String `tfsdk:"vpp_service_token"`
}

type ResponseSmGetOrganizationSmVppAccount struct {
	ID              types.String `tfsdk:"id"`
	VppServiceToken types.String `tfsdk:"vpp_service_token"`
}

// ToBody
func ResponseSmGetOrganizationSmVppAccountsItemsToBody(state OrganizationsSmVppAccounts, response *merakigosdk.ResponseSmGetOrganizationSmVppAccounts) OrganizationsSmVppAccounts {
	var items []ResponseItemSmGetOrganizationSmVppAccounts
	for _, item := range *response {
		itemState := ResponseItemSmGetOrganizationSmVppAccounts{
			ID:              types.StringValue(item.ID),
			VppServiceToken: types.StringValue(item.VppServiceToken),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSmGetOrganizationSmVppAccountItemToBody(state OrganizationsSmVppAccounts, response *merakigosdk.ResponseSmGetOrganizationSmVppAccount) OrganizationsSmVppAccounts {
	itemState := ResponseSmGetOrganizationSmVppAccount{
		ID:              types.StringValue(response.ID),
		VppServiceToken: types.StringValue(response.VppServiceToken),
	}
	state.Item = &itemState
	return state
}
