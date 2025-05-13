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

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsSplashThemesResource{}
	_ resource.ResourceWithConfigure = &OrganizationsSplashThemesResource{}
)

func NewOrganizationsSplashThemesResource() resource.Resource {
	return &OrganizationsSplashThemesResource{}
}

type OrganizationsSplashThemesResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsSplashThemesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsSplashThemesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_splash_themes"
}

func (r *OrganizationsSplashThemesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"base_theme": schema.StringAttribute{
				MarkdownDescription: `base theme id `,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `theme id`,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `theme name`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"theme_assets": schema.SetNestedAttribute{
				MarkdownDescription: `list of theme assets`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `asset id`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `asset name`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (r *OrganizationsSplashThemesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsSplashThemesRs

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
	// Has Paths
	vvOrganizationID := data.OrganizationID.ValueString()
	//Only Items

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationSplashThemes(vvOrganizationID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationSplashThemes",
					restyResp1.String(),
				)
				return
			}
		}
	}

	var responseVerifyItem2 merakigosdk.ResponseItemOrganizationsGetOrganizationSplashThemes
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			err := mapToStruct(result.(map[string]interface{}), &responseVerifyItem2)
			if err != nil {
				resp.Diagnostics.AddError(
					"Failure when executing mapToStruct in resource",
					err.Error(),
				)
				return
			}
			data = ResponseOrganizationsGetOrganizationSplashThemesItemToBodyRs(data, &responseVerifyItem2, false)
			// Path params update assigned
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return

		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Organizations.CreateOrganizationSplashTheme(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationSplashTheme",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationSplashTheme",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationSplashThemes(vvOrganizationID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSplashThemes",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationSplashThemes",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result2 := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		data = ResponseOrganizationsGetOrganizationSplashThemesItemToBodyRs(data, &responseVerifyItem2, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationSplashThemes Result",
			"Not Found",
		)
		return
	}

}

func (r *OrganizationsSplashThemesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsSplashThemesRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	// Not has Item

	vvOrganizationID := data.OrganizationID.ValueString()
	vvName := data.Name.ValueString()

	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationSplashThemes(vvOrganizationID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSplashThemes",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationSplashThemes",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet)
	result2 := getDictResult(responseStruct2, "Name", vvName, simpleCmp)
	var responseVerifyItem2 merakigosdk.ResponseItemOrganizationsGetOrganizationSplashThemes
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		//entro aqui
		data = ResponseOrganizationsGetOrganizationSplashThemesItemToBodyRs(data, &responseVerifyItem2, true)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationSplashThemes Result",
			err.Error(),
		)
		return
	}
}
func (r *OrganizationsSplashThemesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), req.ID)...)
}

func (r *OrganizationsSplashThemesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsSplashThemesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update
	// No update
	resp.Diagnostics.AddError(
		"Update operation not supported in OrganizationsSplashThemes",
		"Update operation not supported in OrganizationsSplashThemes",
	)
	return
}

func (r *OrganizationsSplashThemesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsSplashThemesRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvOrganizationID := state.OrganizationID.ValueString()
	vvID := state.ID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationSplashTheme(vvOrganizationID, vvID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationSplashTheme", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsSplashThemesRs struct {
	OrganizationID types.String `tfsdk:"organization_id"`
	ID             types.String `tfsdk:"id"`
	//TIENE ITEMS
	Name        types.String                                                         `tfsdk:"name"`
	ThemeAssets *[]ResponseItemOrganizationsGetOrganizationSplashThemesThemeAssetsRs `tfsdk:"theme_assets"`
	BaseTheme   types.String                                                         `tfsdk:"base_theme"`
}

type ResponseItemOrganizationsGetOrganizationSplashThemesThemeAssetsRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// FromBody
func (r *OrganizationsSplashThemesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationSplashTheme {
	emptyString := ""
	baseTheme := new(string)
	if !r.BaseTheme.IsUnknown() && !r.BaseTheme.IsNull() {
		*baseTheme = r.BaseTheme.ValueString()
	} else {
		baseTheme = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationSplashTheme{
		BaseTheme: *baseTheme,
		Name:      *name,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationSplashThemesItemToBodyRs(state OrganizationsSplashThemesRs, response *merakigosdk.ResponseItemOrganizationsGetOrganizationSplashThemes, is_read bool) OrganizationsSplashThemesRs {
	itemState := OrganizationsSplashThemesRs{
		ID:   types.StringValue(response.ID),
		Name: types.StringValue(response.Name),
		ThemeAssets: func() *[]ResponseItemOrganizationsGetOrganizationSplashThemesThemeAssetsRs {
			if response.ThemeAssets != nil {
				result := make([]ResponseItemOrganizationsGetOrganizationSplashThemesThemeAssetsRs, len(*response.ThemeAssets))
				for i, themeAssets := range *response.ThemeAssets {
					result[i] = ResponseItemOrganizationsGetOrganizationSplashThemesThemeAssetsRs{
						ID:   types.StringValue(themeAssets.ID),
						Name: types.StringValue(themeAssets.Name),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state = itemState
	return state
}
