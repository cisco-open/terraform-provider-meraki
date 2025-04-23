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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsBrandingPoliciesPrioritiesResource{}
	_ resource.ResourceWithConfigure = &OrganizationsBrandingPoliciesPrioritiesResource{}
)

func NewOrganizationsBrandingPoliciesPrioritiesResource() resource.Resource {
	return &OrganizationsBrandingPoliciesPrioritiesResource{}
}

type OrganizationsBrandingPoliciesPrioritiesResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsBrandingPoliciesPrioritiesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsBrandingPoliciesPrioritiesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_branding_policies_priorities"
}

func (r *OrganizationsBrandingPoliciesPrioritiesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"branding_policy_ids": schema.SetAttribute{
				MarkdownDescription: `      An ordered list of branding policy IDs that determines the priority order of how to apply the policies
`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
		},
	}
}

func (r *OrganizationsBrandingPoliciesPrioritiesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsBrandingPoliciesPrioritiesRs

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
	//Has Item and not has items

	if vvOrganizationID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationBrandingPoliciesPriorities(vvOrganizationID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource OrganizationsBrandingPoliciesPriorities  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource OrganizationsBrandingPoliciesPriorities only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationBrandingPoliciesPriorities(vvOrganizationID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationBrandingPoliciesPriorities",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationBrandingPoliciesPriorities",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationBrandingPoliciesPriorities(vvOrganizationID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationBrandingPoliciesPriorities",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationBrandingPoliciesPriorities",
			err.Error(),
		)
		return
	}

	data = ResponseOrganizationsGetOrganizationBrandingPoliciesPrioritiesItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *OrganizationsBrandingPoliciesPrioritiesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsBrandingPoliciesPrioritiesRs

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

	vvOrganizationID := data.OrganizationID.ValueString()
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationBrandingPoliciesPriorities(vvOrganizationID)
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
				"Failure when executing GetOrganizationBrandingPoliciesPriorities",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationBrandingPoliciesPriorities",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationBrandingPoliciesPrioritiesItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsBrandingPoliciesPrioritiesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), req.ID)...)
}

func (r *OrganizationsBrandingPoliciesPrioritiesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsBrandingPoliciesPrioritiesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationBrandingPoliciesPriorities(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationBrandingPoliciesPriorities",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationBrandingPoliciesPriorities",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsBrandingPoliciesPrioritiesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting OrganizationsBrandingPoliciesPriorities", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsBrandingPoliciesPrioritiesRs struct {
	OrganizationID    types.String `tfsdk:"organization_id"`
	BrandingPolicyIDs types.Set    `tfsdk:"branding_policy_ids"`
}

// FromBody
func (r *OrganizationsBrandingPoliciesPrioritiesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationBrandingPoliciesPriorities {
	var brandingPolicyIDs []string = nil
	r.BrandingPolicyIDs.ElementsAs(ctx, &brandingPolicyIDs, false)
	out := merakigosdk.RequestOrganizationsUpdateOrganizationBrandingPoliciesPriorities{
		BrandingPolicyIDs: brandingPolicyIDs,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationBrandingPoliciesPrioritiesItemToBodyRs(state OrganizationsBrandingPoliciesPrioritiesRs, response *merakigosdk.ResponseOrganizationsGetOrganizationBrandingPoliciesPriorities, is_read bool) OrganizationsBrandingPoliciesPrioritiesRs {
	itemState := OrganizationsBrandingPoliciesPrioritiesRs{
		BrandingPolicyIDs: StringSliceToSet(response.BrandingPolicyIDs),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsBrandingPoliciesPrioritiesRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsBrandingPoliciesPrioritiesRs)
}
