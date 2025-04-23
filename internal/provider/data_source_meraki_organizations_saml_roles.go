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
	_ datasource.DataSource              = &OrganizationsSamlRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSamlRolesDataSource{}
)

func NewOrganizationsSamlRolesDataSource() datasource.DataSource {
	return &OrganizationsSamlRolesDataSource{}
}

type OrganizationsSamlRolesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSamlRolesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSamlRolesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_saml_roles"
}

func (d *OrganizationsSamlRolesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"saml_role_id": schema.StringAttribute{
				MarkdownDescription: `samlRoleId path parameter. Saml role ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"camera": schema.SetNestedAttribute{
						MarkdownDescription: `The list of camera access privileges for SAML administrator`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"access": schema.StringAttribute{
									MarkdownDescription: `Camera access ability`,
									Computed:            true,
								},
								"org_wide": schema.BoolAttribute{
									MarkdownDescription: `Whether or not SAML administrator has org-wide access`,
									Computed:            true,
								},
							},
						},
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `ID associated with the SAML role`,
						Computed:            true,
					},
					"networks": schema.SetNestedAttribute{
						MarkdownDescription: `The list of networks that the SAML administrator has privileges on`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"access": schema.StringAttribute{
									MarkdownDescription: `The privilege of the SAML administrator on the network`,
									Computed:            true,
								},
								"id": schema.StringAttribute{
									MarkdownDescription: `The network ID`,
									Computed:            true,
								},
							},
						},
					},
					"org_access": schema.StringAttribute{
						MarkdownDescription: `The privilege of the SAML administrator on the organization`,
						Computed:            true,
					},
					"role": schema.StringAttribute{
						MarkdownDescription: `The role of the SAML administrator`,
						Computed:            true,
					},
					"tags": schema.SetNestedAttribute{
						MarkdownDescription: `The list of tags that the SAML administrator has privleges on`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"access": schema.StringAttribute{
									MarkdownDescription: `The privilege of the SAML administrator on the tag`,
									Computed:            true,
								},
								"tag": schema.StringAttribute{
									MarkdownDescription: `The name of the tag`,
									Computed:            true,
								},
							},
						},
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationSamlRoles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"camera": schema.SetNestedAttribute{
							MarkdownDescription: `The list of camera access privileges for SAML administrator`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"access": schema.StringAttribute{
										MarkdownDescription: `Camera access ability`,
										Computed:            true,
									},
									"org_wide": schema.BoolAttribute{
										MarkdownDescription: `Whether or not SAML administrator has org-wide access`,
										Computed:            true,
									},
								},
							},
						},
						"id": schema.StringAttribute{
							MarkdownDescription: `ID associated with the SAML role`,
							Computed:            true,
						},
						"networks": schema.SetNestedAttribute{
							MarkdownDescription: `The list of networks that the SAML administrator has privileges on`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"access": schema.StringAttribute{
										MarkdownDescription: `The privilege of the SAML administrator on the network`,
										Computed:            true,
									},
									"id": schema.StringAttribute{
										MarkdownDescription: `The network ID`,
										Computed:            true,
									},
								},
							},
						},
						"org_access": schema.StringAttribute{
							MarkdownDescription: `The privilege of the SAML administrator on the organization`,
							Computed:            true,
						},
						"role": schema.StringAttribute{
							MarkdownDescription: `The role of the SAML administrator`,
							Computed:            true,
						},
						"tags": schema.SetNestedAttribute{
							MarkdownDescription: `The list of tags that the SAML administrator has privleges on`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"access": schema.StringAttribute{
										MarkdownDescription: `The privilege of the SAML administrator on the tag`,
										Computed:            true,
									},
									"tag": schema.StringAttribute{
										MarkdownDescription: `The name of the tag`,
										Computed:            true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsSamlRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSamlRoles OrganizationsSamlRoles
	diags := req.Config.Get(ctx, &organizationsSamlRoles)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsSamlRoles.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsSamlRoles.OrganizationID.IsNull(), !organizationsSamlRoles.SamlRoleID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSamlRoles")
		vvOrganizationID := organizationsSamlRoles.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationSamlRoles(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSamlRoles",
				err.Error(),
			)
			return
		}

		organizationsSamlRoles = ResponseOrganizationsGetOrganizationSamlRolesItemsToBody(organizationsSamlRoles, response1)
		diags = resp.State.Set(ctx, &organizationsSamlRoles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSamlRole")
		vvOrganizationID := organizationsSamlRoles.OrganizationID.ValueString()
		vvSamlRoleID := organizationsSamlRoles.SamlRoleID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Organizations.GetOrganizationSamlRole(vvOrganizationID, vvSamlRoleID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSamlRole",
				err.Error(),
			)
			return
		}

		organizationsSamlRoles = ResponseOrganizationsGetOrganizationSamlRoleItemToBody(organizationsSamlRoles, response2)
		diags = resp.State.Set(ctx, &organizationsSamlRoles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSamlRoles struct {
	OrganizationID types.String                                         `tfsdk:"organization_id"`
	SamlRoleID     types.String                                         `tfsdk:"saml_role_id"`
	Items          *[]ResponseItemOrganizationsGetOrganizationSamlRoles `tfsdk:"items"`
	Item           *ResponseOrganizationsGetOrganizationSamlRole        `tfsdk:"item"`
}

type ResponseItemOrganizationsGetOrganizationSamlRoles struct {
	Camera    *[]ResponseItemOrganizationsGetOrganizationSamlRolesCamera   `tfsdk:"camera"`
	ID        types.String                                                 `tfsdk:"id"`
	Networks  *[]ResponseItemOrganizationsGetOrganizationSamlRolesNetworks `tfsdk:"networks"`
	OrgAccess types.String                                                 `tfsdk:"org_access"`
	Role      types.String                                                 `tfsdk:"role"`
	Tags      *[]ResponseItemOrganizationsGetOrganizationSamlRolesTags     `tfsdk:"tags"`
}

type ResponseItemOrganizationsGetOrganizationSamlRolesCamera struct {
	Access  types.String `tfsdk:"access"`
	OrgWide types.Bool   `tfsdk:"org_wide"`
}

type ResponseItemOrganizationsGetOrganizationSamlRolesNetworks struct {
	Access types.String `tfsdk:"access"`
	ID     types.String `tfsdk:"id"`
}

type ResponseItemOrganizationsGetOrganizationSamlRolesTags struct {
	Access types.String `tfsdk:"access"`
	Tag    types.String `tfsdk:"tag"`
}

type ResponseOrganizationsGetOrganizationSamlRole struct {
	Camera    *[]ResponseOrganizationsGetOrganizationSamlRoleCamera   `tfsdk:"camera"`
	ID        types.String                                            `tfsdk:"id"`
	Networks  *[]ResponseOrganizationsGetOrganizationSamlRoleNetworks `tfsdk:"networks"`
	OrgAccess types.String                                            `tfsdk:"org_access"`
	Role      types.String                                            `tfsdk:"role"`
	Tags      *[]ResponseOrganizationsGetOrganizationSamlRoleTags     `tfsdk:"tags"`
}

type ResponseOrganizationsGetOrganizationSamlRoleCamera struct {
	Access  types.String `tfsdk:"access"`
	OrgWide types.Bool   `tfsdk:"org_wide"`
}

type ResponseOrganizationsGetOrganizationSamlRoleNetworks struct {
	Access types.String `tfsdk:"access"`
	ID     types.String `tfsdk:"id"`
}

type ResponseOrganizationsGetOrganizationSamlRoleTags struct {
	Access types.String `tfsdk:"access"`
	Tag    types.String `tfsdk:"tag"`
}

// ToBody
func ResponseOrganizationsGetOrganizationSamlRolesItemsToBody(state OrganizationsSamlRoles, response *merakigosdk.ResponseOrganizationsGetOrganizationSamlRoles) OrganizationsSamlRoles {
	var items []ResponseItemOrganizationsGetOrganizationSamlRoles
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationSamlRoles{
			Camera: func() *[]ResponseItemOrganizationsGetOrganizationSamlRolesCamera {
				if item.Camera != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationSamlRolesCamera, len(*item.Camera))
					for i, camera := range *item.Camera {
						result[i] = ResponseItemOrganizationsGetOrganizationSamlRolesCamera{
							Access: types.StringValue(camera.Access),
							OrgWide: func() types.Bool {
								if camera.OrgWide != nil {
									return types.BoolValue(*camera.OrgWide)
								}
								return types.Bool{}
							}(),
						}
					}
					return &result
				}
				return nil
			}(),
			ID: types.StringValue(item.ID),
			Networks: func() *[]ResponseItemOrganizationsGetOrganizationSamlRolesNetworks {
				if item.Networks != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationSamlRolesNetworks, len(*item.Networks))
					for i, networks := range *item.Networks {
						result[i] = ResponseItemOrganizationsGetOrganizationSamlRolesNetworks{
							Access: types.StringValue(networks.Access),
							ID:     types.StringValue(networks.ID),
						}
					}
					return &result
				}
				return nil
			}(),
			OrgAccess: types.StringValue(item.OrgAccess),
			Role:      types.StringValue(item.Role),
			Tags: func() *[]ResponseItemOrganizationsGetOrganizationSamlRolesTags {
				if item.Tags != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationSamlRolesTags, len(*item.Tags))
					for i, tags := range *item.Tags {
						result[i] = ResponseItemOrganizationsGetOrganizationSamlRolesTags{
							Access: types.StringValue(tags.Access),
							Tag:    types.StringValue(tags.Tag),
						}
					}
					return &result
				}
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseOrganizationsGetOrganizationSamlRoleItemToBody(state OrganizationsSamlRoles, response *merakigosdk.ResponseOrganizationsGetOrganizationSamlRole) OrganizationsSamlRoles {
	itemState := ResponseOrganizationsGetOrganizationSamlRole{
		Camera: func() *[]ResponseOrganizationsGetOrganizationSamlRoleCamera {
			if response.Camera != nil {
				result := make([]ResponseOrganizationsGetOrganizationSamlRoleCamera, len(*response.Camera))
				for i, camera := range *response.Camera {
					result[i] = ResponseOrganizationsGetOrganizationSamlRoleCamera{
						Access: types.StringValue(camera.Access),
						OrgWide: func() types.Bool {
							if camera.OrgWide != nil {
								return types.BoolValue(*camera.OrgWide)
							}
							return types.Bool{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		ID: types.StringValue(response.ID),
		Networks: func() *[]ResponseOrganizationsGetOrganizationSamlRoleNetworks {
			if response.Networks != nil {
				result := make([]ResponseOrganizationsGetOrganizationSamlRoleNetworks, len(*response.Networks))
				for i, networks := range *response.Networks {
					result[i] = ResponseOrganizationsGetOrganizationSamlRoleNetworks{
						Access: types.StringValue(networks.Access),
						ID:     types.StringValue(networks.ID),
					}
				}
				return &result
			}
			return nil
		}(),
		OrgAccess: types.StringValue(response.OrgAccess),
		Role:      types.StringValue(response.Role),
		Tags: func() *[]ResponseOrganizationsGetOrganizationSamlRoleTags {
			if response.Tags != nil {
				result := make([]ResponseOrganizationsGetOrganizationSamlRoleTags, len(*response.Tags))
				for i, tags := range *response.Tags {
					result[i] = ResponseOrganizationsGetOrganizationSamlRoleTags{
						Access: types.StringValue(tags.Access),
						Tag:    types.StringValue(tags.Tag),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
