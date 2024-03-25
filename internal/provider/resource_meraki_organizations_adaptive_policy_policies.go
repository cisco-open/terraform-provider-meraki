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
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsAdaptivePolicyPoliciesResource{}
	_ resource.ResourceWithConfigure = &OrganizationsAdaptivePolicyPoliciesResource{}
)

func NewOrganizationsAdaptivePolicyPoliciesResource() resource.Resource {
	return &OrganizationsAdaptivePolicyPoliciesResource{}
}

type OrganizationsAdaptivePolicyPoliciesResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsAdaptivePolicyPoliciesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsAdaptivePolicyPoliciesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_adaptive_policy_policies"
}

func (r *OrganizationsAdaptivePolicyPoliciesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"acls": schema.SetNestedAttribute{
				MarkdownDescription: `An ordered array of adaptive policy ACLs (each requires one unique attribute) that apply to this policy (default: [])`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"id": schema.StringAttribute{
							MarkdownDescription: `The ID of the adaptive policy ACL`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the adaptive policy ACL`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"adaptive_policy_id": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"destination_group": schema.SingleNestedAttribute{
				MarkdownDescription: `The destination adaptive policy group (requires one unique attribute)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `The ID of the destination adaptive policy group`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the destination adaptive policy group`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"sgt": schema.Int64Attribute{
						MarkdownDescription: `The SGT of the destination adaptive policy group`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `id path parameter.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_entry_rule": schema.StringAttribute{
				MarkdownDescription: `The rule to apply if there is no matching ACL (default: "default")`,
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
			"source_group": schema.SingleNestedAttribute{
				MarkdownDescription: `The source adaptive policy group (requires one unique attribute)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `The ID of the source adaptive policy group`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the source adaptive policy group`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"sgt": schema.Int64Attribute{
						MarkdownDescription: `The SGT of the source adaptive policy group`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"updated_at": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

//path params to set ['id']

func (r *OrganizationsAdaptivePolicyPoliciesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsAdaptivePolicyPoliciesRs

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
	// organization_id
	//Reviw This  Has Item and item
	//HAS CREATE

	vvID := data.ID.ValueString()
	if vvID != "" {
		responseVerifyItem, restyRespGet, err := r.client.Organizations.GetOrganizationAdaptivePolicyPolicy(vvOrganizationID, vvID)
		if err != nil || responseVerifyItem == nil {
			if restyRespGet != nil {
				if restyRespGet.StatusCode() != 404 {

					resp.Diagnostics.AddError(
						"Failure when executing GetOrganizationAdaptivePolicyPolicy",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data = ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyItemToBodyRs(data, responseVerifyItem, false)
			diags := resp.State.Set(ctx, &data)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	response, restyResp1, err := r.client.Organizations.CreateOrganizationAdaptivePolicyPolicy(vvOrganizationID, data.toSdkApiRequestCreate(ctx))

	if err != nil || restyResp1 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ",
			err.Error(),
		)
		return
	}
	//Items
	vvID = response.AdaptivePolicyID
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationAdaptivePolicyPolicy(vvOrganizationID, vvID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyPolicies",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationAdaptivePolicyPolicies",
			err.Error(),
		)
		return
	} else {
		data = ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyItemToBodyRs(data, responseGet, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	}
}

func (r *OrganizationsAdaptivePolicyPoliciesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsAdaptivePolicyPoliciesRs

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
	// organization_id
	vvID := data.ID.ValueString()
	// id
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationAdaptivePolicyPolicy(vvOrganizationID, vvID)
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
				"Failure when executing GetOrganizationAdaptivePolicyPolicy",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationAdaptivePolicyPolicy",
			err.Error(),
		)
		return
	}

	data = ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsAdaptivePolicyPoliciesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), idParts[1])...)
}

func (r *OrganizationsAdaptivePolicyPoliciesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsAdaptivePolicyPoliciesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	vvID := data.ID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Organizations.UpdateOrganizationAdaptivePolicyPolicy(vvOrganizationID, vvID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationAdaptivePolicyPolicy",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationAdaptivePolicyPolicy",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsAdaptivePolicyPoliciesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsAdaptivePolicyPoliciesRs
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
	_, err := r.client.Organizations.DeleteOrganizationAdaptivePolicyPolicy(vvOrganizationID, vvID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationAdaptivePolicyPolicy", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsAdaptivePolicyPoliciesRs struct {
	OrganizationID   types.String                                                                `tfsdk:"organization_id"`
	ID               types.String                                                                `tfsdk:"id"`
	ACLs             *[]ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyAclsRs           `tfsdk:"acls"`
	AdaptivePolicyID types.String                                                                `tfsdk:"adaptive_policy_id"`
	CreatedAt        types.String                                                                `tfsdk:"created_at"`
	DestinationGroup *ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyDestinationGroupRs `tfsdk:"destination_group"`
	LastEntryRule    types.String                                                                `tfsdk:"last_entry_rule"`
	SourceGroup      *ResponseOrganizationsGetOrganizationAdaptivePolicyPolicySourceGroupRs      `tfsdk:"source_group"`
	UpdatedAt        types.String                                                                `tfsdk:"updated_at"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyAclsRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyDestinationGroupRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Sgt  types.Int64  `tfsdk:"sgt"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyPolicySourceGroupRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Sgt  types.Int64  `tfsdk:"sgt"`
}

// FromBody
func (r *OrganizationsAdaptivePolicyPoliciesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyPolicy {
	emptyString := ""
	var requestOrganizationsCreateOrganizationAdaptivePolicyPolicyACLs []merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyPolicyACLs
	if r.ACLs != nil {
		for _, rItem1 := range *r.ACLs {
			iD := rItem1.ID.ValueString()
			name := rItem1.Name.ValueString()
			requestOrganizationsCreateOrganizationAdaptivePolicyPolicyACLs = append(requestOrganizationsCreateOrganizationAdaptivePolicyPolicyACLs, merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyPolicyACLs{
				ID:   iD,
				Name: name,
			})
		}
	}
	var requestOrganizationsCreateOrganizationAdaptivePolicyPolicyDestinationGroup *merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyPolicyDestinationGroup
	if r.DestinationGroup != nil {
		iD := r.DestinationGroup.ID.ValueString()
		name := r.DestinationGroup.Name.ValueString()
		sgt := func() *int64 {
			if !r.DestinationGroup.Sgt.IsUnknown() && !r.DestinationGroup.Sgt.IsNull() {
				return r.DestinationGroup.Sgt.ValueInt64Pointer()
			}
			return nil
		}()
		requestOrganizationsCreateOrganizationAdaptivePolicyPolicyDestinationGroup = &merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyPolicyDestinationGroup{
			ID:   iD,
			Name: name,
			Sgt:  int64ToIntPointer(sgt),
		}
	}
	lastEntryRule := new(string)
	if !r.LastEntryRule.IsUnknown() && !r.LastEntryRule.IsNull() {
		*lastEntryRule = r.LastEntryRule.ValueString()
	} else {
		lastEntryRule = &emptyString
	}
	var requestOrganizationsCreateOrganizationAdaptivePolicyPolicySourceGroup *merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyPolicySourceGroup
	if r.SourceGroup != nil {
		iD := r.SourceGroup.ID.ValueString()
		name := r.SourceGroup.Name.ValueString()
		sgt := func() *int64 {
			if !r.SourceGroup.Sgt.IsUnknown() && !r.SourceGroup.Sgt.IsNull() {
				return r.SourceGroup.Sgt.ValueInt64Pointer()
			}
			return nil
		}()
		requestOrganizationsCreateOrganizationAdaptivePolicyPolicySourceGroup = &merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyPolicySourceGroup{
			ID:   iD,
			Name: name,
			Sgt:  int64ToIntPointer(sgt),
		}
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyPolicy{
		ACLs: func() *[]merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyPolicyACLs {
			if len(requestOrganizationsCreateOrganizationAdaptivePolicyPolicyACLs) > 0 {
				return &requestOrganizationsCreateOrganizationAdaptivePolicyPolicyACLs
			}
			return nil
		}(),
		DestinationGroup: requestOrganizationsCreateOrganizationAdaptivePolicyPolicyDestinationGroup,
		LastEntryRule:    *lastEntryRule,
		SourceGroup:      requestOrganizationsCreateOrganizationAdaptivePolicyPolicySourceGroup,
	}
	return &out
}
func (r *OrganizationsAdaptivePolicyPoliciesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicy {
	emptyString := ""
	var requestOrganizationsUpdateOrganizationAdaptivePolicyPolicyACLs []merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicyACLs
	if r.ACLs != nil {
		for _, rItem1 := range *r.ACLs {
			iD := rItem1.ID.ValueString()
			name := rItem1.Name.ValueString()
			requestOrganizationsUpdateOrganizationAdaptivePolicyPolicyACLs = append(requestOrganizationsUpdateOrganizationAdaptivePolicyPolicyACLs, merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicyACLs{
				ID:   iD,
				Name: name,
			})
		}
	}
	var requestOrganizationsUpdateOrganizationAdaptivePolicyPolicyDestinationGroup *merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicyDestinationGroup
	if r.DestinationGroup != nil {
		iD := r.DestinationGroup.ID.ValueString()
		name := r.DestinationGroup.Name.ValueString()
		sgt := func() *int64 {
			if !r.DestinationGroup.Sgt.IsUnknown() && !r.DestinationGroup.Sgt.IsNull() {
				return r.DestinationGroup.Sgt.ValueInt64Pointer()
			}
			return nil
		}()
		requestOrganizationsUpdateOrganizationAdaptivePolicyPolicyDestinationGroup = &merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicyDestinationGroup{
			ID:   iD,
			Name: name,
			Sgt:  int64ToIntPointer(sgt),
		}
	}
	lastEntryRule := new(string)
	if !r.LastEntryRule.IsUnknown() && !r.LastEntryRule.IsNull() {
		*lastEntryRule = r.LastEntryRule.ValueString()
	} else {
		lastEntryRule = &emptyString
	}
	var requestOrganizationsUpdateOrganizationAdaptivePolicyPolicySourceGroup *merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicySourceGroup
	if r.SourceGroup != nil {
		iD := r.SourceGroup.ID.ValueString()
		name := r.SourceGroup.Name.ValueString()
		sgt := func() *int64 {
			if !r.SourceGroup.Sgt.IsUnknown() && !r.SourceGroup.Sgt.IsNull() {
				return r.SourceGroup.Sgt.ValueInt64Pointer()
			}
			return nil
		}()
		requestOrganizationsUpdateOrganizationAdaptivePolicyPolicySourceGroup = &merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicySourceGroup{
			ID:   iD,
			Name: name,
			Sgt:  int64ToIntPointer(sgt),
		}
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicy{
		ACLs: func() *[]merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyPolicyACLs {
			if len(requestOrganizationsUpdateOrganizationAdaptivePolicyPolicyACLs) > 0 {
				return &requestOrganizationsUpdateOrganizationAdaptivePolicyPolicyACLs
			}
			return nil
		}(),
		DestinationGroup: requestOrganizationsUpdateOrganizationAdaptivePolicyPolicyDestinationGroup,
		LastEntryRule:    *lastEntryRule,
		SourceGroup:      requestOrganizationsUpdateOrganizationAdaptivePolicyPolicySourceGroup,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyItemToBodyRs(state OrganizationsAdaptivePolicyPoliciesRs, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicyPolicy, is_read bool) OrganizationsAdaptivePolicyPoliciesRs {
	itemState := OrganizationsAdaptivePolicyPoliciesRs{
		ACLs: func() *[]ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyAclsRs {
			if response.ACLs != nil {
				result := make([]ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyAclsRs, len(*response.ACLs))
				for i, aCLs := range *response.ACLs {
					result[i] = ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyAclsRs{
						ID:   types.StringValue(aCLs.ID),
						Name: types.StringValue(aCLs.Name),
					}
				}
				return &result
			}
			return &[]ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyAclsRs{}
		}(),
		AdaptivePolicyID: types.StringValue(response.AdaptivePolicyID),
		CreatedAt:        types.StringValue(response.CreatedAt),
		DestinationGroup: func() *ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyDestinationGroupRs {
			if response.DestinationGroup != nil {
				return &ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyDestinationGroupRs{
					ID:   types.StringValue(response.DestinationGroup.ID),
					Name: types.StringValue(response.DestinationGroup.Name),
					Sgt: func() types.Int64 {
						if response.DestinationGroup.Sgt != nil {
							return types.Int64Value(int64(*response.DestinationGroup.Sgt))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseOrganizationsGetOrganizationAdaptivePolicyPolicyDestinationGroupRs{}
		}(),
		LastEntryRule: types.StringValue(response.LastEntryRule),
		SourceGroup: func() *ResponseOrganizationsGetOrganizationAdaptivePolicyPolicySourceGroupRs {
			if response.SourceGroup != nil {
				return &ResponseOrganizationsGetOrganizationAdaptivePolicyPolicySourceGroupRs{
					ID:   types.StringValue(response.SourceGroup.ID),
					Name: types.StringValue(response.SourceGroup.Name),
					Sgt: func() types.Int64 {
						if response.SourceGroup.Sgt != nil {
							return types.Int64Value(int64(*response.SourceGroup.Sgt))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseOrganizationsGetOrganizationAdaptivePolicyPolicySourceGroupRs{}
		}(),
		UpdatedAt: types.StringValue(response.UpdatedAt),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsAdaptivePolicyPoliciesRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsAdaptivePolicyPoliciesRs)
}
