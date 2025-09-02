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
	_ datasource.DataSource              = &OrganizationsSmAdminsRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSmAdminsRolesDataSource{}
)

func NewOrganizationsSmAdminsRolesDataSource() datasource.DataSource {
	return &OrganizationsSmAdminsRolesDataSource{}
}

type OrganizationsSmAdminsRolesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSmAdminsRolesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSmAdminsRolesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_sm_admins_roles"
}

func (d *OrganizationsSmAdminsRolesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 50.`,
				Optional:            true,
			},
			"role_id": schema.StringAttribute{
				MarkdownDescription: `roleId path parameter. Role ID`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the limited access role`,
						Computed:            true,
					},
					"role_id": schema.StringAttribute{
						MarkdownDescription: `The Id of the limited access role`,
						Computed:            true,
					},
					"scope": schema.StringAttribute{
						MarkdownDescription: `The scope of the limited access role`,
						Computed:            true,
					},
					"tags": schema.ListAttribute{
						MarkdownDescription: `The tags of the limited access role`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}

func (d *OrganizationsSmAdminsRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSmAdminsRoles OrganizationsSmAdminsRoles
	diags := req.Config.Get(ctx, &organizationsSmAdminsRoles)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsSmAdminsRoles.OrganizationID.IsNull(), !organizationsSmAdminsRoles.PerPage.IsNull(), !organizationsSmAdminsRoles.StartingAfter.IsNull(), !organizationsSmAdminsRoles.EndingBefore.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsSmAdminsRoles.OrganizationID.IsNull(), !organizationsSmAdminsRoles.RoleID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSmAdminsRoles")
		vvOrganizationID := organizationsSmAdminsRoles.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSmAdminsRolesQueryParams{}

		queryParams1.PerPage = int(organizationsSmAdminsRoles.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsSmAdminsRoles.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsSmAdminsRoles.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sm.GetOrganizationSmAdminsRoles(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmAdminsRoles",
				err.Error(),
			)
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSmAdminsRole")
		vvOrganizationID := organizationsSmAdminsRoles.OrganizationID.ValueString()
		vvRoleID := organizationsSmAdminsRoles.RoleID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Sm.GetOrganizationSmAdminsRole(vvOrganizationID, vvRoleID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmAdminsRole",
				err.Error(),
			)
			return
		}

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSmAdminsRole",
				err.Error(),
			)
			return
		}

		organizationsSmAdminsRoles = ResponseSmGetOrganizationSmAdminsRoleItemToBody(organizationsSmAdminsRoles, response2)
		diags = resp.State.Set(ctx, &organizationsSmAdminsRoles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSmAdminsRoles struct {
	OrganizationID types.String                           `tfsdk:"organization_id"`
	PerPage        types.Int64                            `tfsdk:"per_page"`
	StartingAfter  types.String                           `tfsdk:"starting_after"`
	EndingBefore   types.String                           `tfsdk:"ending_before"`
	RoleID         types.String                           `tfsdk:"role_id"`
	Item           *ResponseSmGetOrganizationSmAdminsRole `tfsdk:"item"`
}

type ResponseSmGetOrganizationSmAdminsRole struct {
	Name   types.String `tfsdk:"name"`
	RoleID types.String `tfsdk:"role_id"`
	Scope  types.String `tfsdk:"scope"`
	Tags   types.List   `tfsdk:"tags"`
}

// ToBody
func ResponseSmGetOrganizationSmAdminsRoleItemToBody(state OrganizationsSmAdminsRoles, response *merakigosdk.ResponseSmGetOrganizationSmAdminsRole) OrganizationsSmAdminsRoles {
	itemState := ResponseSmGetOrganizationSmAdminsRole{
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
		RoleID: func() types.String {
			if response.RoleID != "" {
				return types.StringValue(response.RoleID)
			}
			return types.String{}
		}(),
		Scope: func() types.String {
			if response.Scope != "" {
				return types.StringValue(response.Scope)
			}
			return types.String{}
		}(),
		Tags: StringSliceToList(response.Tags),
	}
	state.Item = &itemState
	return state
}
