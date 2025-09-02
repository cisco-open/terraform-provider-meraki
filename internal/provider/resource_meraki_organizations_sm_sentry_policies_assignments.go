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

	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsSmSentryPoliciesAssignmentsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsSmSentryPoliciesAssignmentsResource{}
)

func NewOrganizationsSmSentryPoliciesAssignmentsResource() resource.Resource {
	return &OrganizationsSmSentryPoliciesAssignmentsResource{}
}

type OrganizationsSmSentryPoliciesAssignmentsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsSmSentryPoliciesAssignmentsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsSmSentryPoliciesAssignmentsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_sm_sentry_policies_assignments"
}

// resourceAction
func (r *OrganizationsSmSentryPoliciesAssignmentsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `Sentry Group Policies for the Organization keyed by Network Id`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"network_id": schema.StringAttribute{
									MarkdownDescription: `The Id of the Network`,
									Computed:            true,
								},
								"policies": schema.SetNestedAttribute{
									MarkdownDescription: `Array of Sentry Group Policies for the Network`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"created_at": schema.StringAttribute{
												MarkdownDescription: `The creation time of the Sentry Policy`,
												Computed:            true,
											},
											"group_number": schema.StringAttribute{
												MarkdownDescription: `The number of the Group Policy`,
												Computed:            true,
											},
											"group_policy_id": schema.StringAttribute{
												MarkdownDescription: `The Id of the Group Policy. This is associated with the network specified by the networkId.`,
												Computed:            true,
											},
											"last_updated_at": schema.StringAttribute{
												MarkdownDescription: `The last update time of the Sentry Policy`,
												Computed:            true,
											},
											"network_id": schema.StringAttribute{
												MarkdownDescription: `The Id of the Network the Sentry Policy is associated with. In a locale, this should be the Wireless Group if present, otherwise the Wired Group.`,
												Computed:            true,
											},
											"policy_id": schema.StringAttribute{
												MarkdownDescription: `The Id of the Sentry Policy`,
												Computed:            true,
											},
											"priority": schema.StringAttribute{
												MarkdownDescription: `The priority of the Sentry Policy`,
												Computed:            true,
											},
											"scope": schema.StringAttribute{
												MarkdownDescription: `The scope of the Sentry Policy
                                                      Allowed values: [all,none,withAll,withAny,withoutAll,withoutAny]`,
												Computed: true,
											},
											"sm_network_id": schema.StringAttribute{
												MarkdownDescription: `The Id of the Systems Manager Network the Sentry Policy is assigned to`,
												Computed:            true,
											},
											"tags": schema.ListAttribute{
												MarkdownDescription: `The tags of the Sentry Policy`,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"items": schema.ListNestedAttribute{
						MarkdownDescription: `Sentry Group Policies for the Organization keyed by Network Id`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"network_id": schema.StringAttribute{
									MarkdownDescription: `The Id of the Network`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"policies": schema.SetNestedAttribute{
									MarkdownDescription: `Array of Sentry Group Policies for the Network`,
									Optional:            true,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"group_policy_id": schema.StringAttribute{
												MarkdownDescription: `The Group Policy Id`,
												Optional:            true,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"policy_id": schema.StringAttribute{
												MarkdownDescription: `The Sentry Policy Id, if updating an existing Sentry Policy`,
												Optional:            true,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"scope": schema.StringAttribute{
												MarkdownDescription: `The scope of the Sentry Policy
                                                    Allowed values: [all,none,withAll,withAny,withoutAll,withoutAny]`,
												Optional: true,
												Computed: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"sm_network_id": schema.StringAttribute{
												MarkdownDescription: `The Id of the Systems Manager Network`,
												Optional:            true,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"tags": schema.ListAttribute{
												MarkdownDescription: `The tags for the Sentry Policy`,
												Optional:            true,
												Computed:            true,
												ElementType:         types.StringType,
											},
										},
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
func (r *OrganizationsSmSentryPoliciesAssignmentsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsSmSentryPoliciesAssignments

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
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp1, err := r.client.Sm.UpdateOrganizationSmSentryPoliciesAssignments(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationSmSentryPoliciesAssignments",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationSmSentryPoliciesAssignments",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseSmUpdateOrganizationSmSentryPoliciesAssignmentsItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsSmSentryPoliciesAssignmentsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsSmSentryPoliciesAssignmentsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsSmSentryPoliciesAssignmentsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsSmSentryPoliciesAssignments struct {
	OrganizationID types.String                                              `tfsdk:"organization_id"`
	Item           *ResponseSmUpdateOrganizationSmSentryPoliciesAssignments  `tfsdk:"item"`
	Parameters     *RequestSmUpdateOrganizationSmSentryPoliciesAssignmentsRs `tfsdk:"parameters"`
}

type ResponseSmUpdateOrganizationSmSentryPoliciesAssignments struct {
	Items *[]ResponseSmUpdateOrganizationSmSentryPoliciesAssignmentsItems `tfsdk:"items"`
}

type ResponseSmUpdateOrganizationSmSentryPoliciesAssignmentsItems struct {
	NetworkID types.String                                                            `tfsdk:"network_id"`
	Policies  *[]ResponseSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies `tfsdk:"policies"`
}

type ResponseSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies struct {
	CreatedAt     types.String `tfsdk:"created_at"`
	GroupNumber   types.String `tfsdk:"group_number"`
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
	LastUpdatedAt types.String `tfsdk:"last_updated_at"`
	NetworkID     types.String `tfsdk:"network_id"`
	PolicyID      types.String `tfsdk:"policy_id"`
	Priority      types.String `tfsdk:"priority"`
	Scope         types.String `tfsdk:"scope"`
	SmNetworkID   types.String `tfsdk:"sm_network_id"`
	Tags          types.List   `tfsdk:"tags"`
}

type RequestSmUpdateOrganizationSmSentryPoliciesAssignmentsRs struct {
	Items *[]RequestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsRs `tfsdk:"items"`
}

type RequestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsRs struct {
	NetworkID types.String                                                             `tfsdk:"network_id"`
	Policies  *[]RequestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPoliciesRs `tfsdk:"policies"`
}

type RequestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPoliciesRs struct {
	GroupPolicyID types.String `tfsdk:"group_policy_id"`
	PolicyID      types.String `tfsdk:"policy_id"`
	Scope         types.String `tfsdk:"scope"`
	SmNetworkID   types.String `tfsdk:"sm_network_id"`
	Tags          types.List   `tfsdk:"tags"`
}

// FromBody
func (r *OrganizationsSmSentryPoliciesAssignments) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSmUpdateOrganizationSmSentryPoliciesAssignments {
	re := *r.Parameters
	var requestSmUpdateOrganizationSmSentryPoliciesAssignmentsItems []merakigosdk.RequestSmUpdateOrganizationSmSentryPoliciesAssignmentsItems

	if re.Items != nil {
		for _, rItem1 := range *re.Items {
			networkID := rItem1.NetworkID.ValueString()

			log.Printf("[DEBUG] #TODO []RequestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies")
			var requestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies []merakigosdk.RequestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies

			if rItem1.Policies != nil {
				for _, rItem2 := range *rItem1.Policies {
					groupPolicyID := rItem2.GroupPolicyID.ValueString()
					policyID := rItem2.PolicyID.ValueString()
					scope := rItem2.Scope.ValueString()
					smNetworkID := rItem2.SmNetworkID.ValueString()

					var tags []string = nil
					rItem2.Tags.ElementsAs(ctx, &tags, false)
					requestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies = append(requestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies, merakigosdk.RequestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies{
						GroupPolicyID: groupPolicyID,
						PolicyID:      policyID,
						Scope:         scope,
						SmNetworkID:   smNetworkID,
						Tags:          tags,
					})
					//[debug] Is Array: True
				}
			}
			requestSmUpdateOrganizationSmSentryPoliciesAssignmentsItems = append(requestSmUpdateOrganizationSmSentryPoliciesAssignmentsItems, merakigosdk.RequestSmUpdateOrganizationSmSentryPoliciesAssignmentsItems{
				NetworkID: networkID,
				Policies: func() *[]merakigosdk.RequestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies {
					if len(requestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies) > 0 {
						return &requestSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies
					}
					return nil
				}(),
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestSmUpdateOrganizationSmSentryPoliciesAssignments{
		Items: &requestSmUpdateOrganizationSmSentryPoliciesAssignmentsItems,
	}
	return &out
}

// ToBody
func ResponseSmUpdateOrganizationSmSentryPoliciesAssignmentsItemToBody(state OrganizationsSmSentryPoliciesAssignments, response *merakigosdk.ResponseSmUpdateOrganizationSmSentryPoliciesAssignments) OrganizationsSmSentryPoliciesAssignments {
	itemState := ResponseSmUpdateOrganizationSmSentryPoliciesAssignments{
		Items: func() *[]ResponseSmUpdateOrganizationSmSentryPoliciesAssignmentsItems {
			if response.Items != nil {
				result := make([]ResponseSmUpdateOrganizationSmSentryPoliciesAssignmentsItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseSmUpdateOrganizationSmSentryPoliciesAssignmentsItems{
						NetworkID: func() types.String {
							if items.NetworkID != "" {
								return types.StringValue(items.NetworkID)
							}
							return types.String{}
						}(),
						Policies: func() *[]ResponseSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies {
							if items.Policies != nil {
								result := make([]ResponseSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies, len(*items.Policies))
								for i, policies := range *items.Policies {
									result[i] = ResponseSmUpdateOrganizationSmSentryPoliciesAssignmentsItemsPolicies{
										CreatedAt: func() types.String {
											if policies.CreatedAt != "" {
												return types.StringValue(policies.CreatedAt)
											}
											return types.String{}
										}(),
										GroupNumber: func() types.String {
											if policies.GroupNumber != "" {
												return types.StringValue(policies.GroupNumber)
											}
											return types.String{}
										}(),
										GroupPolicyID: func() types.String {
											if policies.GroupPolicyID != "" {
												return types.StringValue(policies.GroupPolicyID)
											}
											return types.String{}
										}(),
										LastUpdatedAt: func() types.String {
											if policies.LastUpdatedAt != "" {
												return types.StringValue(policies.LastUpdatedAt)
											}
											return types.String{}
										}(),
										NetworkID: func() types.String {
											if policies.NetworkID != "" {
												return types.StringValue(policies.NetworkID)
											}
											return types.String{}
										}(),
										PolicyID: func() types.String {
											if policies.PolicyID != "" {
												return types.StringValue(policies.PolicyID)
											}
											return types.String{}
										}(),
										Priority: func() types.String {
											if policies.Priority != "" {
												return types.StringValue(policies.Priority)
											}
											return types.String{}
										}(),
										Scope: func() types.String {
											if policies.Scope != "" {
												return types.StringValue(policies.Scope)
											}
											return types.String{}
										}(),
										SmNetworkID: func() types.String {
											if policies.SmNetworkID != "" {
												return types.StringValue(policies.SmNetworkID)
											}
											return types.String{}
										}(),
										Tags: StringSliceToList(policies.Tags),
									}
								}
								return &result
							}
							return nil
						}(),
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
