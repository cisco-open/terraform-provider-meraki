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
	_ datasource.DataSource              = &OrganizationsAdaptivePolicyGroupsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAdaptivePolicyGroupsDataSource{}
)

func NewOrganizationsAdaptivePolicyGroupsDataSource() datasource.DataSource {
	return &OrganizationsAdaptivePolicyGroupsDataSource{}
}

type OrganizationsAdaptivePolicyGroupsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAdaptivePolicyGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAdaptivePolicyGroupsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_adaptive_policy_groups"
}

func (d *OrganizationsAdaptivePolicyGroupsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: `id path parameter.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"created_at": schema.StringAttribute{
						MarkdownDescription: `Created at timestamp for the adaptive policy group`,
						Computed:            true,
					},
					"description": schema.StringAttribute{
						MarkdownDescription: `The description for the adaptive policy group`,
						Computed:            true,
					},
					"group_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the adaptive policy group`,
						Computed:            true,
					},
					"is_default_group": schema.BoolAttribute{
						MarkdownDescription: `Whether the adaptive policy group is the default group`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the adaptive policy group`,
						Computed:            true,
					},
					"policy_objects": schema.SetNestedAttribute{
						MarkdownDescription: `The policy objects for the adaptive policy group`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `The ID of the policy object`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the policy object`,
									Computed:            true,
								},
							},
						},
					},
					"required_ip_mappings": schema.ListAttribute{
						MarkdownDescription: `List of required IP mappings for the adaptive policy group`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"sgt": schema.Int64Attribute{
						MarkdownDescription: `The security group tag for the adaptive policy group`,
						Computed:            true,
					},
					"updated_at": schema.StringAttribute{
						MarkdownDescription: `Updated at timestamp for the adaptive policy group`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationAdaptivePolicyGroups`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"created_at": schema.StringAttribute{
							MarkdownDescription: `Created at timestamp for the adaptive policy group`,
							Computed:            true,
						},
						"description": schema.StringAttribute{
							MarkdownDescription: `The description for the adaptive policy group`,
							Computed:            true,
						},
						"group_id": schema.StringAttribute{
							MarkdownDescription: `The ID of the adaptive policy group`,
							Computed:            true,
						},
						"is_default_group": schema.BoolAttribute{
							MarkdownDescription: `Whether the adaptive policy group is the default group`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `The name of the adaptive policy group`,
							Computed:            true,
						},
						"policy_objects": schema.SetNestedAttribute{
							MarkdownDescription: `The policy objects for the adaptive policy group`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"id": schema.StringAttribute{
										MarkdownDescription: `The ID of the policy object`,
										Computed:            true,
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `The name of the policy object`,
										Computed:            true,
									},
								},
							},
						},
						"required_ip_mappings": schema.ListAttribute{
							MarkdownDescription: `List of required IP mappings for the adaptive policy group`,
							Computed:            true,
							ElementType:         types.StringType,
						},
						"sgt": schema.Int64Attribute{
							MarkdownDescription: `The security group tag for the adaptive policy group`,
							Computed:            true,
						},
						"updated_at": schema.StringAttribute{
							MarkdownDescription: `Updated at timestamp for the adaptive policy group`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsAdaptivePolicyGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAdaptivePolicyGroups OrganizationsAdaptivePolicyGroups
	diags := req.Config.Get(ctx, &organizationsAdaptivePolicyGroups)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsAdaptivePolicyGroups.OrganizationID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsAdaptivePolicyGroups.OrganizationID.IsNull(), !organizationsAdaptivePolicyGroups.ID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAdaptivePolicyGroups")
		vvOrganizationID := organizationsAdaptivePolicyGroups.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAdaptivePolicyGroups(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyGroups",
				err.Error(),
			)
			return
		}

		organizationsAdaptivePolicyGroups = ResponseOrganizationsGetOrganizationAdaptivePolicyGroupsItemsToBody(organizationsAdaptivePolicyGroups, response1)
		diags = resp.State.Set(ctx, &organizationsAdaptivePolicyGroups)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAdaptivePolicyGroup")
		vvOrganizationID := organizationsAdaptivePolicyGroups.OrganizationID.ValueString()
		vvID := organizationsAdaptivePolicyGroups.ID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Organizations.GetOrganizationAdaptivePolicyGroup(vvOrganizationID, vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyGroup",
				err.Error(),
			)
			return
		}

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyGroup",
				err.Error(),
			)
			return
		}

		organizationsAdaptivePolicyGroups = ResponseOrganizationsGetOrganizationAdaptivePolicyGroupItemToBody(organizationsAdaptivePolicyGroups, response2)
		diags = resp.State.Set(ctx, &organizationsAdaptivePolicyGroups)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAdaptivePolicyGroups struct {
	OrganizationID types.String                                                    `tfsdk:"organization_id"`
	ID             types.String                                                    `tfsdk:"id"`
	Items          *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroups `tfsdk:"items"`
	Item           *ResponseOrganizationsGetOrganizationAdaptivePolicyGroup        `tfsdk:"item"`
}

type ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroups struct {
	CreatedAt          types.String                                                                 `tfsdk:"created_at"`
	Description        types.String                                                                 `tfsdk:"description"`
	GroupID            types.String                                                                 `tfsdk:"group_id"`
	IsDefaultGroup     types.Bool                                                                   `tfsdk:"is_default_group"`
	Name               types.String                                                                 `tfsdk:"name"`
	PolicyObjects      *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects `tfsdk:"policy_objects"`
	RequiredIPMappings types.List                                                                   `tfsdk:"required_ip_mappings"`
	Sgt                types.Int64                                                                  `tfsdk:"sgt"`
	UpdatedAt          types.String                                                                 `tfsdk:"updated_at"`
}

type ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyGroup struct {
	CreatedAt          types.String                                                            `tfsdk:"created_at"`
	Description        types.String                                                            `tfsdk:"description"`
	GroupID            types.String                                                            `tfsdk:"group_id"`
	IsDefaultGroup     types.Bool                                                              `tfsdk:"is_default_group"`
	Name               types.String                                                            `tfsdk:"name"`
	PolicyObjects      *[]ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects `tfsdk:"policy_objects"`
	RequiredIPMappings types.List                                                              `tfsdk:"required_ip_mappings"`
	Sgt                types.Int64                                                             `tfsdk:"sgt"`
	UpdatedAt          types.String                                                            `tfsdk:"updated_at"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAdaptivePolicyGroupsItemsToBody(state OrganizationsAdaptivePolicyGroups, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicyGroups) OrganizationsAdaptivePolicyGroups {
	var items []ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroups
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroups{
			CreatedAt: func() types.String {
				if item.CreatedAt != "" {
					return types.StringValue(item.CreatedAt)
				}
				return types.String{}
			}(),
			Description: func() types.String {
				if item.Description != "" {
					return types.StringValue(item.Description)
				}
				return types.String{}
			}(),
			GroupID: func() types.String {
				if item.GroupID != "" {
					return types.StringValue(item.GroupID)
				}
				return types.String{}
			}(),
			IsDefaultGroup: func() types.Bool {
				if item.IsDefaultGroup != nil {
					return types.BoolValue(*item.IsDefaultGroup)
				}
				return types.Bool{}
			}(),
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
			PolicyObjects: func() *[]ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects {
				if item.PolicyObjects != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects, len(*item.PolicyObjects))
					for i, policyObjects := range *item.PolicyObjects {
						result[i] = ResponseItemOrganizationsGetOrganizationAdaptivePolicyGroupsPolicyObjects{
							ID: func() types.String {
								if policyObjects.ID != "" {
									return types.StringValue(policyObjects.ID)
								}
								return types.String{}
							}(),
							Name: func() types.String {
								if policyObjects.Name != "" {
									return types.StringValue(policyObjects.Name)
								}
								return types.String{}
							}(),
						}
					}
					return &result
				}
				return nil
			}(),
			RequiredIPMappings: StringSliceToList(item.RequiredIPMappings),
			Sgt: func() types.Int64 {
				if item.Sgt != nil {
					return types.Int64Value(int64(*item.Sgt))
				}
				return types.Int64{}
			}(),
			UpdatedAt: func() types.String {
				if item.UpdatedAt != "" {
					return types.StringValue(item.UpdatedAt)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseOrganizationsGetOrganizationAdaptivePolicyGroupItemToBody(state OrganizationsAdaptivePolicyGroups, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicyGroup) OrganizationsAdaptivePolicyGroups {
	itemState := ResponseOrganizationsGetOrganizationAdaptivePolicyGroup{
		CreatedAt: func() types.String {
			if response.CreatedAt != "" {
				return types.StringValue(response.CreatedAt)
			}
			return types.String{}
		}(),
		Description: func() types.String {
			if response.Description != "" {
				return types.StringValue(response.Description)
			}
			return types.String{}
		}(),
		GroupID: func() types.String {
			if response.GroupID != "" {
				return types.StringValue(response.GroupID)
			}
			return types.String{}
		}(),
		IsDefaultGroup: func() types.Bool {
			if response.IsDefaultGroup != nil {
				return types.BoolValue(*response.IsDefaultGroup)
			}
			return types.Bool{}
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
		PolicyObjects: func() *[]ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects {
			if response.PolicyObjects != nil {
				result := make([]ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects, len(*response.PolicyObjects))
				for i, policyObjects := range *response.PolicyObjects {
					result[i] = ResponseOrganizationsGetOrganizationAdaptivePolicyGroupPolicyObjects{
						ID: func() types.String {
							if policyObjects.ID != "" {
								return types.StringValue(policyObjects.ID)
							}
							return types.String{}
						}(),
						Name: func() types.String {
							if policyObjects.Name != "" {
								return types.StringValue(policyObjects.Name)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		RequiredIPMappings: StringSliceToList(response.RequiredIPMappings),
		Sgt: func() types.Int64 {
			if response.Sgt != nil {
				return types.Int64Value(int64(*response.Sgt))
			}
			return types.Int64{}
		}(),
		UpdatedAt: func() types.String {
			if response.UpdatedAt != "" {
				return types.StringValue(response.UpdatedAt)
			}
			return types.String{}
		}(),
	}
	state.Item = &itemState
	return state
}
