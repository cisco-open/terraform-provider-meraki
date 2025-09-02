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
	_ datasource.DataSource              = &OrganizationsCameraPermissionsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsCameraPermissionsDataSource{}
)

func NewOrganizationsCameraPermissionsDataSource() datasource.DataSource {
	return &OrganizationsCameraPermissionsDataSource{}
}

type OrganizationsCameraPermissionsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsCameraPermissionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsCameraPermissionsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_camera_permissions"
}

func (d *OrganizationsCameraPermissionsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"permission_scope_id": schema.StringAttribute{
				MarkdownDescription: `permissionScopeId path parameter. Permission scope ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `Permission scope id`,
						Computed:            true,
					},
					"level": schema.StringAttribute{
						MarkdownDescription: `Permission scope level`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of permission scope`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *OrganizationsCameraPermissionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsCameraPermissions OrganizationsCameraPermissions
	diags := req.Config.Get(ctx, &organizationsCameraPermissions)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCameraPermission")
		vvOrganizationID := organizationsCameraPermissions.OrganizationID.ValueString()
		vvPermissionScopeID := organizationsCameraPermissions.PermissionScopeID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Camera.GetOrganizationCameraPermission(vvOrganizationID, vvPermissionScopeID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraPermission",
				err.Error(),
			)
			return
		}

		organizationsCameraPermissions = ResponseCameraGetOrganizationCameraPermissionItemToBody(organizationsCameraPermissions, response1)
		diags = resp.State.Set(ctx, &organizationsCameraPermissions)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsCameraPermissions struct {
	OrganizationID    types.String                                   `tfsdk:"organization_id"`
	PermissionScopeID types.String                                   `tfsdk:"permission_scope_id"`
	Item              *ResponseCameraGetOrganizationCameraPermission `tfsdk:"item"`
}

type ResponseCameraGetOrganizationCameraPermission struct {
	ID    types.String `tfsdk:"id"`
	Level types.String `tfsdk:"level"`
	Name  types.String `tfsdk:"name"`
}

// ToBody
func ResponseCameraGetOrganizationCameraPermissionItemToBody(state OrganizationsCameraPermissions, response *merakigosdk.ResponseCameraGetOrganizationCameraPermission) OrganizationsCameraPermissions {
	itemState := ResponseCameraGetOrganizationCameraPermission{
		ID: func() types.String {
			if response.ID != "" {
				return types.StringValue(response.ID)
			}
			return types.String{}
		}(),
		Level: func() types.String {
			if response.Level != "" {
				return types.StringValue(response.Level)
			}
			return types.String{}
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
