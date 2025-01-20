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
	_ datasource.DataSource              = &OrganizationsAdminsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAdminsDataSource{}
)

func NewOrganizationsAdminsDataSource() datasource.DataSource {
	return &OrganizationsAdminsDataSource{}
}

type OrganizationsAdminsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAdminsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAdminsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_admins"
}

func (d *OrganizationsAdminsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter the result set by the included set of network IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationAdmins`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"account_status": schema.StringAttribute{
							MarkdownDescription: `Status of the admin's account`,
							Computed:            true,
						},
						"authentication_method": schema.StringAttribute{
							MarkdownDescription: `Admin's authentication method`,
							Computed:            true,
						},
						"email": schema.StringAttribute{
							MarkdownDescription: `Admin's email address`,
							Computed:            true,
						},
						"has_api_key": schema.BoolAttribute{
							MarkdownDescription: `Indicates whether the admin has an API key`,
							Computed:            true,
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `Admin's ID`,
							Computed:            true,
						},
						"last_active": schema.StringAttribute{
							MarkdownDescription: `Time when the admin was last active`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Admin's username`,
							Computed:            true,
						},
						"networks": schema.SetNestedAttribute{
							MarkdownDescription: `Admin network access information`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"access": schema.StringAttribute{
										MarkdownDescription: `Admin's level of access to the network`,
										Computed:            true,
									},
									"id": schema.StringAttribute{
										MarkdownDescription: `Network ID`,
										Computed:            true,
									},
								},
							},
						},
						"org_access": schema.StringAttribute{
							MarkdownDescription: `Admin's level of access to the organization`,
							Computed:            true,
						},
						"tags": schema.SetNestedAttribute{
							MarkdownDescription: `Admin tag information`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"access": schema.StringAttribute{
										MarkdownDescription: `Access level for the tag`,
										Computed:            true,
									},
									"tag": schema.StringAttribute{
										MarkdownDescription: `Tag value`,
										Computed:            true,
									},
								},
							},
						},
						"two_factor_auth_enabled": schema.BoolAttribute{
							MarkdownDescription: `Indicates whether two-factor authentication is enabled`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsAdminsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAdmins OrganizationsAdmins
	diags := req.Config.Get(ctx, &organizationsAdmins)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAdmins")
		vvOrganizationID := organizationsAdmins.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationAdminsQueryParams{}

		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsAdmins.NetworkIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAdmins(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdmins",
				err.Error(),
			)
			return
		}

		organizationsAdmins = ResponseOrganizationsGetOrganizationAdminsItemsToBody(organizationsAdmins, response1)
		diags = resp.State.Set(ctx, &organizationsAdmins)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAdmins struct {
	OrganizationID types.String                                      `tfsdk:"organization_id"`
	NetworkIDs     types.List                                        `tfsdk:"network_ids"`
	Items          *[]ResponseItemOrganizationsGetOrganizationAdmins `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationAdmins struct {
	AccountStatus        types.String                                              `tfsdk:"account_status"`
	AuthenticationMethod types.String                                              `tfsdk:"authentication_method"`
	Email                types.String                                              `tfsdk:"email"`
	HasAPIKey            types.Bool                                                `tfsdk:"has_api_key"`
	ID                   types.String                                              `tfsdk:"id"`
	LastActive           types.String                                              `tfsdk:"last_active"`
	Name                 types.String                                              `tfsdk:"name"`
	Networks             *[]ResponseItemOrganizationsGetOrganizationAdminsNetworks `tfsdk:"networks"`
	OrgAccess            types.String                                              `tfsdk:"org_access"`
	Tags                 *[]ResponseItemOrganizationsGetOrganizationAdminsTags     `tfsdk:"tags"`
	TwoFactorAuthEnabled types.Bool                                                `tfsdk:"two_factor_auth_enabled"`
}

type ResponseItemOrganizationsGetOrganizationAdminsNetworks struct {
	Access types.String `tfsdk:"access"`
	ID     types.String `tfsdk:"id"`
}

type ResponseItemOrganizationsGetOrganizationAdminsTags struct {
	Access types.String `tfsdk:"access"`
	Tag    types.String `tfsdk:"tag"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAdminsItemsToBody(state OrganizationsAdmins, response *merakigosdk.ResponseOrganizationsGetOrganizationAdmins) OrganizationsAdmins {
	var items []ResponseItemOrganizationsGetOrganizationAdmins
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationAdmins{
			AccountStatus:        types.StringValue(item.AccountStatus),
			AuthenticationMethod: types.StringValue(item.AuthenticationMethod),
			Email:                types.StringValue(item.Email),
			HasAPIKey: func() types.Bool {
				if item.HasAPIKey != nil {
					return types.BoolValue(*item.HasAPIKey)
				}
				return types.Bool{}
			}(),
			ID:         types.StringValue(item.ID),
			LastActive: types.StringValue(item.LastActive),
			Name:       types.StringValue(item.Name),
			Networks: func() *[]ResponseItemOrganizationsGetOrganizationAdminsNetworks {
				if item.Networks != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationAdminsNetworks, len(*item.Networks))
					for i, networks := range *item.Networks {
						result[i] = ResponseItemOrganizationsGetOrganizationAdminsNetworks{
							Access: types.StringValue(networks.Access),
							ID:     types.StringValue(networks.ID),
						}
					}
					return &result
				}
				return nil
			}(),
			OrgAccess: types.StringValue(item.OrgAccess),
			Tags: func() *[]ResponseItemOrganizationsGetOrganizationAdminsTags {
				if item.Tags != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationAdminsTags, len(*item.Tags))
					for i, tags := range *item.Tags {
						result[i] = ResponseItemOrganizationsGetOrganizationAdminsTags{
							Access: types.StringValue(tags.Access),
							Tag:    types.StringValue(tags.Tag),
						}
					}
					return &result
				}
				return nil
			}(),
			TwoFactorAuthEnabled: func() types.Bool {
				if item.TwoFactorAuthEnabled != nil {
					return types.BoolValue(*item.TwoFactorAuthEnabled)
				}
				return types.Bool{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
