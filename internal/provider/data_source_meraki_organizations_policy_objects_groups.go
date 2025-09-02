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
	_ datasource.DataSource              = &OrganizationsPolicyObjectsGroupsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsPolicyObjectsGroupsDataSource{}
)

func NewOrganizationsPolicyObjectsGroupsDataSource() datasource.DataSource {
	return &OrganizationsPolicyObjectsGroupsDataSource{}
}

type OrganizationsPolicyObjectsGroupsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsPolicyObjectsGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsPolicyObjectsGroupsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_policy_objects_groups"
}

func (d *OrganizationsPolicyObjectsGroupsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Optional:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 10 1000. Default is 1000.`,
				Optional:            true,
			},
			"policy_object_group_id": schema.StringAttribute{
				MarkdownDescription: `policyObjectGroupId path parameter. Policy object group ID`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"category": schema.StringAttribute{
						MarkdownDescription: `Type of object groups. (NetworkObjectGroup, GeoLocationGroup, PortObjectGroup, ApplicationGroup)`,
						Computed:            true,
					},
					"created_at": schema.StringAttribute{
						MarkdownDescription: `Time Stamp of policy object creation.`,
						Computed:            true,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `Policy object ID`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the Policy object group.`,
						Computed:            true,
					},
					"network_ids": schema.ListAttribute{
						MarkdownDescription: `Network ID's associated with the policy objects.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"object_ids": schema.SetAttribute{
						MarkdownDescription: `Policy objects associated with Network Object Group or Port Object Group`,
						Computed:            true,
						ElementType:         types.StringType, //TODO FINAL ELSE param_schema.Elem.Type para revisar

					},
					"updated_at": schema.StringAttribute{
						MarkdownDescription: `Time Stamp of policy object updation.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *OrganizationsPolicyObjectsGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsPolicyObjectsGroups OrganizationsPolicyObjectsGroups
	diags := req.Config.Get(ctx, &organizationsPolicyObjectsGroups)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsPolicyObjectsGroups.OrganizationID.IsNull(), !organizationsPolicyObjectsGroups.PerPage.IsNull(), !organizationsPolicyObjectsGroups.StartingAfter.IsNull(), !organizationsPolicyObjectsGroups.EndingBefore.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsPolicyObjectsGroups.OrganizationID.IsNull(), !organizationsPolicyObjectsGroups.PolicyObjectGroupID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationPolicyObjectsGroups")
		vvOrganizationID := organizationsPolicyObjectsGroups.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationPolicyObjectsGroupsQueryParams{}

		queryParams1.PerPage = int(organizationsPolicyObjectsGroups.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsPolicyObjectsGroups.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsPolicyObjectsGroups.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationPolicyObjectsGroups(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationPolicyObjectsGroups",
				err.Error(),
			)
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationPolicyObjectsGroup")
		vvOrganizationID := organizationsPolicyObjectsGroups.OrganizationID.ValueString()
		vvPolicyObjectGroupID := organizationsPolicyObjectsGroups.PolicyObjectGroupID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Organizations.GetOrganizationPolicyObjectsGroup(vvOrganizationID, vvPolicyObjectGroupID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationPolicyObjectsGroup",
				err.Error(),
			)
			return
		}

		organizationsPolicyObjectsGroups = ResponseOrganizationsGetOrganizationPolicyObjectsGroupItemToBody(organizationsPolicyObjectsGroups, response2)
		diags = resp.State.Set(ctx, &organizationsPolicyObjectsGroups)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsPolicyObjectsGroups struct {
	OrganizationID      types.String                                            `tfsdk:"organization_id"`
	PerPage             types.Int64                                             `tfsdk:"per_page"`
	StartingAfter       types.String                                            `tfsdk:"starting_after"`
	EndingBefore        types.String                                            `tfsdk:"ending_before"`
	PolicyObjectGroupID types.String                                            `tfsdk:"policy_object_group_id"`
	Item                *ResponseOrganizationsGetOrganizationPolicyObjectsGroup `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationPolicyObjectsGroup struct {
	Category   types.String `tfsdk:"category"`
	CreatedAt  types.String `tfsdk:"created_at"`
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	NetworkIDs types.List   `tfsdk:"network_ids"`
	ObjectIDs  types.List   `tfsdk:"object_ids"`
	UpdatedAt  types.String `tfsdk:"updated_at"`
}

// ToBody
func ResponseOrganizationsGetOrganizationPolicyObjectsGroupItemToBody(state OrganizationsPolicyObjectsGroups, response *merakigosdk.ResponseOrganizationsGetOrganizationPolicyObjectsGroup) OrganizationsPolicyObjectsGroups {
	itemState := ResponseOrganizationsGetOrganizationPolicyObjectsGroup{
		Category: func() types.String {
			if response.Category != "" {
				return types.StringValue(response.Category)
			}
			return types.String{}
		}(),
		CreatedAt: func() types.String {
			if response.CreatedAt != "" {
				return types.StringValue(response.CreatedAt)
			}
			return types.String{}
		}(),
		ID: func() types.String {
			if response.ID != "" {
				return types.StringValue(response.ID)
			}
			return types.String{}
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
		NetworkIDs: StringSliceToList(response.NetworkIDs),
		ObjectIDs:  StringSliceToList(*response.ObjectIDs),
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
