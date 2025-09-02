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
	_ datasource.DataSource              = &OrganizationsCameraRolesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsCameraRolesDataSource{}
)

func NewOrganizationsCameraRolesDataSource() datasource.DataSource {
	return &OrganizationsCameraRolesDataSource{}
}

type OrganizationsCameraRolesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsCameraRolesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsCameraRolesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_camera_roles"
}

func (d *OrganizationsCameraRolesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"role_id": schema.StringAttribute{
				MarkdownDescription: `roleId path parameter. Role ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"applied_on_devices": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									Computed: true,
								},
								"permission_level": schema.StringAttribute{
									Computed: true,
								},
								"permission_scope": schema.StringAttribute{
									Computed: true,
								},
								"permission_scope_id": schema.StringAttribute{
									Computed: true,
								},
								"tag": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"applied_on_networks": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									Computed: true,
								},
								"permission_level": schema.StringAttribute{
									Computed: true,
								},
								"permission_scope": schema.StringAttribute{
									Computed: true,
								},
								"permission_scope_id": schema.StringAttribute{
									Computed: true,
								},
								"tag": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"applied_org_wide": schema.SetNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"permission_level": schema.StringAttribute{
									Computed: true,
								},
								"permission_scope": schema.StringAttribute{
									Computed: true,
								},
								"permission_scope_id": schema.StringAttribute{
									Computed: true,
								},
								"tag": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseCameraGetOrganizationCameraRoles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"applied_on_devices": schema.SetNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"id": schema.StringAttribute{
										Computed: true,
									},
									"permission_level": schema.StringAttribute{
										Computed: true,
									},
									"permission_scope": schema.StringAttribute{
										Computed: true,
									},
									"permission_scope_id": schema.StringAttribute{
										Computed: true,
									},
									"tag": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"applied_on_networks": schema.SetNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"id": schema.StringAttribute{
										Computed: true,
									},
									"permission_level": schema.StringAttribute{
										Computed: true,
									},
									"permission_scope": schema.StringAttribute{
										Computed: true,
									},
									"permission_scope_id": schema.StringAttribute{
										Computed: true,
									},
									"tag": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"applied_org_wide": schema.SetNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"permission_level": schema.StringAttribute{
										Computed: true,
									},
									"permission_scope": schema.StringAttribute{
										Computed: true,
									},
									"permission_scope_id": schema.StringAttribute{
										Computed: true,
									},
									"tag": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsCameraRolesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsCameraRoles OrganizationsCameraRoles
	diags := req.Config.Get(ctx, &organizationsCameraRoles)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsCameraRoles.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsCameraRoles.OrganizationID.IsNull(), !organizationsCameraRoles.RoleID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCameraRoles")
		vvOrganizationID := organizationsCameraRoles.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Camera.GetOrganizationCameraRoles(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraRoles",
				err.Error(),
			)
			return
		}

		organizationsCameraRoles = ResponseCameraGetOrganizationCameraRolesItemsToBody(organizationsCameraRoles, response1)
		diags = resp.State.Set(ctx, &organizationsCameraRoles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCameraRole")
		vvOrganizationID := organizationsCameraRoles.OrganizationID.ValueString()
		vvRoleID := organizationsCameraRoles.RoleID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Camera.GetOrganizationCameraRole(vvOrganizationID, vvRoleID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraRole",
				err.Error(),
			)
			return
		}

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraRole",
				err.Error(),
			)
			return
		}

		organizationsCameraRoles = ResponseCameraGetOrganizationCameraRoleItemToBody(organizationsCameraRoles, response2)
		diags = resp.State.Set(ctx, &organizationsCameraRoles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsCameraRoles struct {
	OrganizationID types.String                                    `tfsdk:"organization_id"`
	RoleID         types.String                                    `tfsdk:"role_id"`
	Items          *[]ResponseItemCameraGetOrganizationCameraRoles `tfsdk:"items"`
	Item           *ResponseCameraGetOrganizationCameraRole        `tfsdk:"item"`
}

type ResponseItemCameraGetOrganizationCameraRoles struct {
	AppliedOnDevices  *[]ResponseItemCameraGetOrganizationCameraRolesAppliedOnDevices  `tfsdk:"applied_on_devices"`
	AppliedOnNetworks *[]ResponseItemCameraGetOrganizationCameraRolesAppliedOnNetworks `tfsdk:"applied_on_networks"`
	AppliedOrgWide    *[]ResponseItemCameraGetOrganizationCameraRolesAppliedOrgWide    `tfsdk:"applied_org_wide"`
	Name              types.String                                                     `tfsdk:"name"`
}

type ResponseItemCameraGetOrganizationCameraRolesAppliedOnDevices struct {
	ID                types.String `tfsdk:"id"`
	PermissionLevel   types.String `tfsdk:"permission_level"`
	PermissionScope   types.String `tfsdk:"permission_scope"`
	PermissionScopeID types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
}

type ResponseItemCameraGetOrganizationCameraRolesAppliedOnNetworks struct {
	ID                types.String `tfsdk:"id"`
	PermissionLevel   types.String `tfsdk:"permission_level"`
	PermissionScope   types.String `tfsdk:"permission_scope"`
	PermissionScopeID types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
}

type ResponseItemCameraGetOrganizationCameraRolesAppliedOrgWide struct {
	PermissionLevel   types.String `tfsdk:"permission_level"`
	PermissionScope   types.String `tfsdk:"permission_scope"`
	PermissionScopeID types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
}

type ResponseCameraGetOrganizationCameraRole struct {
	AppliedOnDevices  *[]ResponseCameraGetOrganizationCameraRoleAppliedOnDevices  `tfsdk:"applied_on_devices"`
	AppliedOnNetworks *[]ResponseCameraGetOrganizationCameraRoleAppliedOnNetworks `tfsdk:"applied_on_networks"`
	AppliedOrgWide    *[]ResponseCameraGetOrganizationCameraRoleAppliedOrgWide    `tfsdk:"applied_org_wide"`
	Name              types.String                                                `tfsdk:"name"`
}

type ResponseCameraGetOrganizationCameraRoleAppliedOnDevices struct {
	ID                types.String `tfsdk:"id"`
	PermissionLevel   types.String `tfsdk:"permission_level"`
	PermissionScope   types.String `tfsdk:"permission_scope"`
	PermissionScopeID types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
}

type ResponseCameraGetOrganizationCameraRoleAppliedOnNetworks struct {
	ID                types.String `tfsdk:"id"`
	PermissionLevel   types.String `tfsdk:"permission_level"`
	PermissionScope   types.String `tfsdk:"permission_scope"`
	PermissionScopeID types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
}

type ResponseCameraGetOrganizationCameraRoleAppliedOrgWide struct {
	PermissionLevel   types.String `tfsdk:"permission_level"`
	PermissionScope   types.String `tfsdk:"permission_scope"`
	PermissionScopeID types.String `tfsdk:"permission_scope_id"`
	Tag               types.String `tfsdk:"tag"`
}

// ToBody
func ResponseCameraGetOrganizationCameraRolesItemsToBody(state OrganizationsCameraRoles, response *merakigosdk.ResponseCameraGetOrganizationCameraRoles) OrganizationsCameraRoles {
	var items []ResponseItemCameraGetOrganizationCameraRoles
	for _, item := range *response {
		itemState := ResponseItemCameraGetOrganizationCameraRoles{
			AppliedOnDevices: func() *[]ResponseItemCameraGetOrganizationCameraRolesAppliedOnDevices {
				if item.AppliedOnDevices != nil {
					result := make([]ResponseItemCameraGetOrganizationCameraRolesAppliedOnDevices, len(*item.AppliedOnDevices))
					for i, appliedOnDevices := range *item.AppliedOnDevices {
						result[i] = ResponseItemCameraGetOrganizationCameraRolesAppliedOnDevices{
							ID: func() types.String {
								if appliedOnDevices.ID != "" {
									return types.StringValue(appliedOnDevices.ID)
								}
								return types.String{}
							}(),
							PermissionLevel: func() types.String {
								if appliedOnDevices.PermissionLevel != "" {
									return types.StringValue(appliedOnDevices.PermissionLevel)
								}
								return types.String{}
							}(),
							PermissionScope: func() types.String {
								if appliedOnDevices.PermissionScope != "" {
									return types.StringValue(appliedOnDevices.PermissionScope)
								}
								return types.String{}
							}(),
							PermissionScopeID: func() types.String {
								if appliedOnDevices.PermissionScopeID != "" {
									return types.StringValue(appliedOnDevices.PermissionScopeID)
								}
								return types.String{}
							}(),
							Tag: func() types.String {
								if appliedOnDevices.Tag != "" {
									return types.StringValue(appliedOnDevices.Tag)
								}
								return types.String{}
							}(),
						}
					}
					return &result
				}
				return nil
			}(),
			AppliedOnNetworks: func() *[]ResponseItemCameraGetOrganizationCameraRolesAppliedOnNetworks {
				if item.AppliedOnNetworks != nil {
					result := make([]ResponseItemCameraGetOrganizationCameraRolesAppliedOnNetworks, len(*item.AppliedOnNetworks))
					for i, appliedOnNetworks := range *item.AppliedOnNetworks {
						result[i] = ResponseItemCameraGetOrganizationCameraRolesAppliedOnNetworks{
							ID: func() types.String {
								if appliedOnNetworks.ID != "" {
									return types.StringValue(appliedOnNetworks.ID)
								}
								return types.String{}
							}(),
							PermissionLevel: func() types.String {
								if appliedOnNetworks.PermissionLevel != "" {
									return types.StringValue(appliedOnNetworks.PermissionLevel)
								}
								return types.String{}
							}(),
							PermissionScope: func() types.String {
								if appliedOnNetworks.PermissionScope != "" {
									return types.StringValue(appliedOnNetworks.PermissionScope)
								}
								return types.String{}
							}(),
							PermissionScopeID: func() types.String {
								if appliedOnNetworks.PermissionScopeID != "" {
									return types.StringValue(appliedOnNetworks.PermissionScopeID)
								}
								return types.String{}
							}(),
							Tag: func() types.String {
								if appliedOnNetworks.Tag != "" {
									return types.StringValue(appliedOnNetworks.Tag)
								}
								return types.String{}
							}(),
						}
					}
					return &result
				}
				return nil
			}(),
			AppliedOrgWide: func() *[]ResponseItemCameraGetOrganizationCameraRolesAppliedOrgWide {
				if item.AppliedOrgWide != nil {
					result := make([]ResponseItemCameraGetOrganizationCameraRolesAppliedOrgWide, len(*item.AppliedOrgWide))
					for i, appliedOrgWide := range *item.AppliedOrgWide {
						result[i] = ResponseItemCameraGetOrganizationCameraRolesAppliedOrgWide{
							PermissionLevel: func() types.String {
								if appliedOrgWide.PermissionLevel != "" {
									return types.StringValue(appliedOrgWide.PermissionLevel)
								}
								return types.String{}
							}(),
							PermissionScope: func() types.String {
								if appliedOrgWide.PermissionScope != "" {
									return types.StringValue(appliedOrgWide.PermissionScope)
								}
								return types.String{}
							}(),
							PermissionScopeID: func() types.String {
								if appliedOrgWide.PermissionScopeID != "" {
									return types.StringValue(appliedOrgWide.PermissionScopeID)
								}
								return types.String{}
							}(),
							Tag: func() types.String {
								if appliedOrgWide.Tag != "" {
									return types.StringValue(appliedOrgWide.Tag)
								}
								return types.String{}
							}(),
						}
					}
					return &result
				}
				return nil
			}(),
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseCameraGetOrganizationCameraRoleItemToBody(state OrganizationsCameraRoles, response *merakigosdk.ResponseCameraGetOrganizationCameraRole) OrganizationsCameraRoles {
	itemState := ResponseCameraGetOrganizationCameraRole{
		AppliedOnDevices: func() *[]ResponseCameraGetOrganizationCameraRoleAppliedOnDevices {
			if response.AppliedOnDevices != nil {
				result := make([]ResponseCameraGetOrganizationCameraRoleAppliedOnDevices, len(*response.AppliedOnDevices))
				for i, appliedOnDevices := range *response.AppliedOnDevices {
					result[i] = ResponseCameraGetOrganizationCameraRoleAppliedOnDevices{
						ID: func() types.String {
							if appliedOnDevices.ID != "" {
								return types.StringValue(appliedOnDevices.ID)
							}
							return types.String{}
						}(),
						PermissionLevel: func() types.String {
							if appliedOnDevices.PermissionLevel != "" {
								return types.StringValue(appliedOnDevices.PermissionLevel)
							}
							return types.String{}
						}(),
						PermissionScope: func() types.String {
							if appliedOnDevices.PermissionScope != "" {
								return types.StringValue(appliedOnDevices.PermissionScope)
							}
							return types.String{}
						}(),
						PermissionScopeID: func() types.String {
							if appliedOnDevices.PermissionScopeID != "" {
								return types.StringValue(appliedOnDevices.PermissionScopeID)
							}
							return types.String{}
						}(),
						Tag: func() types.String {
							if appliedOnDevices.Tag != "" {
								return types.StringValue(appliedOnDevices.Tag)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		AppliedOnNetworks: func() *[]ResponseCameraGetOrganizationCameraRoleAppliedOnNetworks {
			if response.AppliedOnNetworks != nil {
				result := make([]ResponseCameraGetOrganizationCameraRoleAppliedOnNetworks, len(*response.AppliedOnNetworks))
				for i, appliedOnNetworks := range *response.AppliedOnNetworks {
					result[i] = ResponseCameraGetOrganizationCameraRoleAppliedOnNetworks{
						ID: func() types.String {
							if appliedOnNetworks.ID != "" {
								return types.StringValue(appliedOnNetworks.ID)
							}
							return types.String{}
						}(),
						PermissionLevel: func() types.String {
							if appliedOnNetworks.PermissionLevel != "" {
								return types.StringValue(appliedOnNetworks.PermissionLevel)
							}
							return types.String{}
						}(),
						PermissionScope: func() types.String {
							if appliedOnNetworks.PermissionScope != "" {
								return types.StringValue(appliedOnNetworks.PermissionScope)
							}
							return types.String{}
						}(),
						PermissionScopeID: func() types.String {
							if appliedOnNetworks.PermissionScopeID != "" {
								return types.StringValue(appliedOnNetworks.PermissionScopeID)
							}
							return types.String{}
						}(),
						Tag: func() types.String {
							if appliedOnNetworks.Tag != "" {
								return types.StringValue(appliedOnNetworks.Tag)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		AppliedOrgWide: func() *[]ResponseCameraGetOrganizationCameraRoleAppliedOrgWide {
			if response.AppliedOrgWide != nil {
				result := make([]ResponseCameraGetOrganizationCameraRoleAppliedOrgWide, len(*response.AppliedOrgWide))
				for i, appliedOrgWide := range *response.AppliedOrgWide {
					result[i] = ResponseCameraGetOrganizationCameraRoleAppliedOrgWide{
						PermissionLevel: func() types.String {
							if appliedOrgWide.PermissionLevel != "" {
								return types.StringValue(appliedOrgWide.PermissionLevel)
							}
							return types.String{}
						}(),
						PermissionScope: func() types.String {
							if appliedOrgWide.PermissionScope != "" {
								return types.StringValue(appliedOrgWide.PermissionScope)
							}
							return types.String{}
						}(),
						PermissionScopeID: func() types.String {
							if appliedOrgWide.PermissionScopeID != "" {
								return types.StringValue(appliedOrgWide.PermissionScopeID)
							}
							return types.String{}
						}(),
						Tag: func() types.String {
							if appliedOrgWide.Tag != "" {
								return types.StringValue(appliedOrgWide.Tag)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
	}
	state.Item = &itemState
	return state
}
