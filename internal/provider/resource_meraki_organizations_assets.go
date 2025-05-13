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

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsAssetsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsAssetsResource{}
)

func NewOrganizationsAssetsResource() resource.Resource {
	return &OrganizationsAssetsResource{}
}

type OrganizationsAssetsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsAssetsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsAssetsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_assets"
}

// resourceAction
func (r *OrganizationsAssetsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"theme_identifier": schema.StringAttribute{
				MarkdownDescription: `themeIdentifier path parameter. Theme identifier`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"file_data": schema.StringAttribute{
						MarkdownDescription: `Splash theme asset file date base64 encoded`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `Splash theme asset id`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Splash theme asset name`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"content": schema.StringAttribute{
						MarkdownDescription: `a file containing the asset content`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `File name. Will overwrite files with same name.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *OrganizationsAssetsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsAssets

	var item types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	vvOrganizationID := data.OrganizationID.ValueString()
	vvThemeIDentifier := data.ThemeIDentifier.ValueString()
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Organizations.CreateOrganizationSplashThemeAsset(vvOrganizationID, vvThemeIDentifier, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationSplashThemeAsset",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationSplashThemeAsset",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseOrganizationsCreateOrganizationSplashThemeAssetItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsAssetsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsAssetsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsAssetsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsAssets struct {
	OrganizationID  types.String                                              `tfsdk:"organization_id"`
	ThemeIDentifier types.String                                              `tfsdk:"theme_identifier"`
	Item            *ResponseOrganizationsCreateOrganizationSplashThemeAsset  `tfsdk:"item"`
	Parameters      *RequestOrganizationsCreateOrganizationSplashThemeAssetRs `tfsdk:"parameters"`
}

type ResponseOrganizationsCreateOrganizationSplashThemeAsset struct {
	FileData types.String `tfsdk:"file_data"`
	ID       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
}

type RequestOrganizationsCreateOrganizationSplashThemeAssetRs struct {
	Content types.String `tfsdk:"content"`
	Name    types.String `tfsdk:"name"`
}

// FromBody
func (r *OrganizationsAssets) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationSplashThemeAsset {
	emptyString := ""
	re := *r.Parameters
	content := new(string)
	if !re.Content.IsUnknown() && !re.Content.IsNull() {
		*content = re.Content.ValueString()
	} else {
		content = &emptyString
	}
	name := new(string)
	if !re.Name.IsUnknown() && !re.Name.IsNull() {
		*name = re.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationSplashThemeAsset{
		Content: *content,
		Name:    *name,
	}
	return &out
}

// ToBody
func ResponseOrganizationsCreateOrganizationSplashThemeAssetItemToBody(state OrganizationsAssets, response *merakigosdk.ResponseOrganizationsCreateOrganizationSplashThemeAsset) OrganizationsAssets {
	itemState := ResponseOrganizationsCreateOrganizationSplashThemeAsset{
		FileData: types.StringValue(response.FileData),
		ID:       types.StringValue(response.ID),
		Name:     types.StringValue(response.Name),
	}
	state.Item = &itemState
	return state
}
