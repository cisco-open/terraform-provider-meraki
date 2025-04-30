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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceContentFilteringResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceContentFilteringResource{}
)

func NewNetworksApplianceContentFilteringResource() resource.Resource {
	return &NetworksApplianceContentFilteringResource{}
}

type NetworksApplianceContentFilteringResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceContentFilteringResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceContentFilteringResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_content_filtering"
}

func (r *NetworksApplianceContentFilteringResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"allowed_url_patterns": schema.SetAttribute{
				MarkdownDescription: `A list of URL patterns that are allowed`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"blocked_url_categories_response": schema.SetNestedAttribute{
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Computed: true,
						},
						"name": schema.StringAttribute{
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Computed: true,
						},
					},
				},
			},
			"blocked_url_categories": schema.SetAttribute{
				MarkdownDescription: `A list of URL categories to block`,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"blocked_url_patterns": schema.SetAttribute{
				MarkdownDescription: `A list of URL patterns that are blocked`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"network_id": schema.StringAttribute{
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"url_category_list_size": schema.StringAttribute{
				MarkdownDescription: `URL category list size which is either 'topSites' or 'fullList'
                                  Allowed values: [fullList,topSites]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"fullList",
						"topSites",
					),
				},
			},
		},
	}
}

func (r *NetworksApplianceContentFilteringResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceContentFilteringRs

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
	vvNetworkID := data.NetworkID.ValueString()
	//Has Item and not has items

	if vvNetworkID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceContentFiltering(vvNetworkID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksApplianceContentFiltering  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource NetworksApplianceContentFiltering only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceContentFiltering(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceContentFiltering",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceContentFiltering",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceContentFiltering(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceContentFiltering",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceContentFiltering",
			err.Error(),
		)
		return
	}

	data = ResponseApplianceGetNetworkApplianceContentFilteringItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksApplianceContentFilteringResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceContentFilteringRs

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
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceContentFiltering(vvNetworkID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceContentFiltering",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceContentFiltering",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceContentFilteringItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceContentFilteringResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceContentFilteringResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceContentFilteringRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceContentFiltering(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceContentFiltering",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceContentFiltering",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceContentFilteringResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceContentFiltering", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceContentFilteringRs struct {
	NetworkID              types.String                                                                  `tfsdk:"network_id"`
	AllowedURLPatterns     types.Set                                                                     `tfsdk:"allowed_url_patterns"`
	BlockedURLCategories   *[]ResponseApplianceGetNetworkApplianceContentFilteringBlockedUrlCategoriesRs `tfsdk:"blocked_url_categories_response"`
	BlockedURLCategoriesRs types.Set                                                                     `tfsdk:"blocked_url_categories"`
	BlockedURLPatterns     types.Set                                                                     `tfsdk:"blocked_url_patterns"`
	URLCategoryListSize    types.String                                                                  `tfsdk:"url_category_list_size"`
}

type ResponseApplianceGetNetworkApplianceContentFilteringBlockedUrlCategoriesRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// FromBody
func (r *NetworksApplianceContentFilteringRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceContentFiltering {
	emptyString := ""
	var allowedURLPatterns []string = nil
	r.AllowedURLPatterns.ElementsAs(ctx, &allowedURLPatterns, false)
	var blockedURLCategories []string = nil
	r.BlockedURLCategoriesRs.ElementsAs(ctx, &blockedURLCategories, false)
	var blockedURLPatterns []string = nil
	r.BlockedURLPatterns.ElementsAs(ctx, &blockedURLPatterns, false)
	uRLCategoryListSize := new(string)
	if !r.URLCategoryListSize.IsUnknown() && !r.URLCategoryListSize.IsNull() {
		*uRLCategoryListSize = r.URLCategoryListSize.ValueString()
	} else {
		uRLCategoryListSize = &emptyString
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceContentFiltering{
		AllowedURLPatterns:   allowedURLPatterns,
		BlockedURLCategories: blockedURLCategories,
		BlockedURLPatterns:   blockedURLPatterns,
		URLCategoryListSize:  *uRLCategoryListSize,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceContentFilteringItemToBodyRs(state NetworksApplianceContentFilteringRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceContentFiltering, is_read bool) NetworksApplianceContentFilteringRs {
	itemState := NetworksApplianceContentFilteringRs{
		AllowedURLPatterns: StringSliceToSet(response.AllowedURLPatterns),
		BlockedURLCategories: func() *[]ResponseApplianceGetNetworkApplianceContentFilteringBlockedUrlCategoriesRs {
			if response.BlockedURLCategories != nil {
				result := make([]ResponseApplianceGetNetworkApplianceContentFilteringBlockedUrlCategoriesRs, len(*response.BlockedURLCategories))
				for i, blockedURLCategories := range *response.BlockedURLCategories {
					result[i] = ResponseApplianceGetNetworkApplianceContentFilteringBlockedUrlCategoriesRs{
						ID:   types.StringValue(blockedURLCategories.ID),
						Name: types.StringValue(blockedURLCategories.Name),
					}
				}
				return &result
			}
			return nil
		}(),
		BlockedURLPatterns: StringSliceToSet(response.BlockedURLPatterns),
		// URLCategoryListSize:    types.StringValue(response.URLCategoryListSize),
		BlockedURLCategoriesRs: state.BlockedURLCategoriesRs,
		URLCategoryListSize:    state.URLCategoryListSize,
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceContentFilteringRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceContentFilteringRs)
}
