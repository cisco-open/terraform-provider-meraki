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
							ID:                types.StringValue(appliedOnDevices.ID),
							PermissionLevel:   types.StringValue(appliedOnDevices.PermissionLevel),
							PermissionScope:   types.StringValue(appliedOnDevices.PermissionScope),
							PermissionScopeID: types.StringValue(appliedOnDevices.PermissionScopeID),
							Tag:               types.StringValue(appliedOnDevices.Tag),
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
							ID:                types.StringValue(appliedOnNetworks.ID),
							PermissionLevel:   types.StringValue(appliedOnNetworks.PermissionLevel),
							PermissionScope:   types.StringValue(appliedOnNetworks.PermissionScope),
							PermissionScopeID: types.StringValue(appliedOnNetworks.PermissionScopeID),
							Tag:               types.StringValue(appliedOnNetworks.Tag),
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
							PermissionLevel:   types.StringValue(appliedOrgWide.PermissionLevel),
							PermissionScope:   types.StringValue(appliedOrgWide.PermissionScope),
							PermissionScopeID: types.StringValue(appliedOrgWide.PermissionScopeID),
							Tag:               types.StringValue(appliedOrgWide.Tag),
						}
					}
					return &result
				}
				return nil
			}(),
			Name: types.StringValue(item.Name),
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
						ID:                types.StringValue(appliedOnDevices.ID),
						PermissionLevel:   types.StringValue(appliedOnDevices.PermissionLevel),
						PermissionScope:   types.StringValue(appliedOnDevices.PermissionScope),
						PermissionScopeID: types.StringValue(appliedOnDevices.PermissionScopeID),
						Tag:               types.StringValue(appliedOnDevices.Tag),
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
						ID:                types.StringValue(appliedOnNetworks.ID),
						PermissionLevel:   types.StringValue(appliedOnNetworks.PermissionLevel),
						PermissionScope:   types.StringValue(appliedOnNetworks.PermissionScope),
						PermissionScopeID: types.StringValue(appliedOnNetworks.PermissionScopeID),
						Tag:               types.StringValue(appliedOnNetworks.Tag),
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
						PermissionLevel:   types.StringValue(appliedOrgWide.PermissionLevel),
						PermissionScope:   types.StringValue(appliedOrgWide.PermissionScope),
						PermissionScopeID: types.StringValue(appliedOrgWide.PermissionScopeID),
						Tag:               types.StringValue(appliedOrgWide.Tag),
					}
				}
				return &result
			}
			return nil
		}(),
		Name: types.StringValue(response.Name),
	}
	state.Item = &itemState
	return state
}
