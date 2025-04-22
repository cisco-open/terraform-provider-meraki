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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsApplianceSecurityIntrusionResource{}
	_ resource.ResourceWithConfigure = &OrganizationsApplianceSecurityIntrusionResource{}
)

func NewOrganizationsApplianceSecurityIntrusionResource() resource.Resource {
	return &OrganizationsApplianceSecurityIntrusionResource{}
}

type OrganizationsApplianceSecurityIntrusionResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsApplianceSecurityIntrusionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsApplianceSecurityIntrusionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_security_intrusion"
}

func (r *OrganizationsApplianceSecurityIntrusionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"allowed_rules": schema.SetNestedAttribute{
				MarkdownDescription: `Sets a list of specific SNORT signatures to allow`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"message": schema.StringAttribute{
							MarkdownDescription: `Message is optional and is ignored on a PUT call. It is allowed in order for PUT to be compatible with GET`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"rule_id": schema.StringAttribute{
							MarkdownDescription: `A rule identifier of the format meraki:intrusion/snort/GID/<gid>/SID/<sid>. gid and sid can be obtained from either https://www.snort.org/rule-docs or as ruleIds from the security events in /organization/[orgId]/securityEvents`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
		},
	}
}

func (r *OrganizationsApplianceSecurityIntrusionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsApplianceSecurityIntrusionRs

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
		responseVerifyItem, restyResp1, err := r.client.Appliance.GetOrganizationApplianceSecurityIntrusion(vvOrganizationID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource OrganizationsApplianceSecurityIntrusion  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource OrganizationsApplianceSecurityIntrusion only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateOrganizationApplianceSecurityIntrusion(vvOrganizationID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationApplianceSecurityIntrusion",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationApplianceSecurityIntrusion",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Appliance.GetOrganizationApplianceSecurityIntrusion(vvOrganizationID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceSecurityIntrusion",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationApplianceSecurityIntrusion",
			err.Error(),
		)
		return
	}

	data = ResponseApplianceGetOrganizationApplianceSecurityIntrusionItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *OrganizationsApplianceSecurityIntrusionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsApplianceSecurityIntrusionRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetOrganizationApplianceSecurityIntrusion(vvOrganizationID)
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
				"Failure when executing GetOrganizationApplianceSecurityIntrusion",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationApplianceSecurityIntrusion",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetOrganizationApplianceSecurityIntrusionItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsApplianceSecurityIntrusionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), req.ID)...)
}

func (r *OrganizationsApplianceSecurityIntrusionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsApplianceSecurityIntrusionRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateOrganizationApplianceSecurityIntrusion(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationApplianceSecurityIntrusion",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationApplianceSecurityIntrusion",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsApplianceSecurityIntrusionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting OrganizationsApplianceSecurityIntrusion", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsApplianceSecurityIntrusionRs struct {
	OrganizationID types.String                                                                `tfsdk:"organization_id"`
	AllowedRules   *[]ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRulesRs `tfsdk:"allowed_rules"`
}

type ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRulesRs struct {
	Message types.String `tfsdk:"message"`
	RuleID  types.String `tfsdk:"rule_id"`
}

// FromBody
func (r *OrganizationsApplianceSecurityIntrusionRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateOrganizationApplianceSecurityIntrusion {
	var requestApplianceUpdateOrganizationApplianceSecurityIntrusionAllowedRules []merakigosdk.RequestApplianceUpdateOrganizationApplianceSecurityIntrusionAllowedRules

	if r.AllowedRules != nil {
		for _, rItem1 := range *r.AllowedRules {
			message := rItem1.Message.ValueString()
			ruleID := rItem1.RuleID.ValueString()
			requestApplianceUpdateOrganizationApplianceSecurityIntrusionAllowedRules = append(requestApplianceUpdateOrganizationApplianceSecurityIntrusionAllowedRules, merakigosdk.RequestApplianceUpdateOrganizationApplianceSecurityIntrusionAllowedRules{
				Message: message,
				RuleID:  ruleID,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceUpdateOrganizationApplianceSecurityIntrusion{
		AllowedRules: func() *[]merakigosdk.RequestApplianceUpdateOrganizationApplianceSecurityIntrusionAllowedRules {
			if len(requestApplianceUpdateOrganizationApplianceSecurityIntrusionAllowedRules) > 0 {
				return &requestApplianceUpdateOrganizationApplianceSecurityIntrusionAllowedRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetOrganizationApplianceSecurityIntrusionItemToBodyRs(state OrganizationsApplianceSecurityIntrusionRs, response *merakigosdk.ResponseApplianceGetOrganizationApplianceSecurityIntrusion, is_read bool) OrganizationsApplianceSecurityIntrusionRs {
	itemState := OrganizationsApplianceSecurityIntrusionRs{
		AllowedRules: func() *[]ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRulesRs {
			if response.AllowedRules != nil {
				result := make([]ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRulesRs, len(*response.AllowedRules))
				for i, allowedRules := range *response.AllowedRules {
					result[i] = ResponseApplianceGetOrganizationApplianceSecurityIntrusionAllowedRulesRs{
						Message: types.StringValue(allowedRules.Message),
						RuleID:  types.StringValue(allowedRules.RuleID),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsApplianceSecurityIntrusionRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsApplianceSecurityIntrusionRs)
}
