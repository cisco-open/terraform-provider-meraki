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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsPolicyObjectsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsPolicyObjectsDataSource{}
)

func NewOrganizationsPolicyObjectsDataSource() datasource.DataSource {
	return &OrganizationsPolicyObjectsDataSource{}
}

type OrganizationsPolicyObjectsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsPolicyObjectsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsPolicyObjectsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_policy_objects"
}

func (d *OrganizationsPolicyObjectsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 10 5000. Default is 5000.`,
				Optional:            true,
			},
			"policy_object_id": schema.StringAttribute{
				MarkdownDescription: `policyObjectId path parameter. Policy object ID`,
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
						MarkdownDescription: `Category of a policy object (one of: adaptivePolicy, network)`,
						Computed:            true,
					},
					"cidr": schema.StringAttribute{
						MarkdownDescription: `CIDR Value of a policy object`,
						Computed:            true,
					},
					"created_at": schema.StringAttribute{
						MarkdownDescription: `Time Stamp of policy object creation.`,
						Computed:            true,
					},
					"group_ids": schema.ListAttribute{
						MarkdownDescription: `The IDs of policy object groups the policy object belongs to.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `Policy object ID`,
						Computed:            true,
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of policy object (alphanumeric, space, dash, or underscore characters only).`,
						Computed:            true,
					},
					"network_ids": schema.ListAttribute{
						MarkdownDescription: `The IDs of the networks that use the policy object.`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"type": schema.StringAttribute{
						MarkdownDescription: `Type of a policy object (one of: adaptivePolicyIpv4Cidr, cidr, fqdn, ipAndMask)`,
						Computed:            true,
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

func (d *OrganizationsPolicyObjectsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsPolicyObjects OrganizationsPolicyObjects
	diags := req.Config.Get(ctx, &organizationsPolicyObjects)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!organizationsPolicyObjects.OrganizationID.IsNull(), !organizationsPolicyObjects.PerPage.IsNull(), !organizationsPolicyObjects.StartingAfter.IsNull(), !organizationsPolicyObjects.EndingBefore.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!organizationsPolicyObjects.OrganizationID.IsNull(), !organizationsPolicyObjects.PolicyObjectID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationPolicyObjects")
		vvOrganizationID := organizationsPolicyObjects.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationPolicyObjectsQueryParams{}

		queryParams1.PerPage = int(organizationsPolicyObjects.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsPolicyObjects.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsPolicyObjects.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationPolicyObjects(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationPolicyObjects",
				err.Error(),
			)
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetOrganizationPolicyObject")
		vvOrganizationID := organizationsPolicyObjects.OrganizationID.ValueString()
		vvPolicyObjectID := organizationsPolicyObjects.PolicyObjectID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Organizations.GetOrganizationPolicyObject(vvOrganizationID, vvPolicyObjectID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationPolicyObject",
				err.Error(),
			)
			return
		}

		organizationsPolicyObjects = ResponseOrganizationsGetOrganizationPolicyObjectItemToBody(organizationsPolicyObjects, response2)
		diags = resp.State.Set(ctx, &organizationsPolicyObjects)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsPolicyObjects struct {
	OrganizationID types.String                                      `tfsdk:"organization_id"`
	PerPage        types.Int64                                       `tfsdk:"per_page"`
	StartingAfter  types.String                                      `tfsdk:"starting_after"`
	EndingBefore   types.String                                      `tfsdk:"ending_before"`
	PolicyObjectID types.String                                      `tfsdk:"policy_object_id"`
	Item           *ResponseOrganizationsGetOrganizationPolicyObject `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationPolicyObject struct {
	Category   types.String `tfsdk:"category"`
	Cidr       types.String `tfsdk:"cidr"`
	CreatedAt  types.String `tfsdk:"created_at"`
	GroupIDs   types.List   `tfsdk:"group_ids"`
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	NetworkIDs types.List   `tfsdk:"network_ids"`
	Type       types.String `tfsdk:"type"`
	UpdatedAt  types.String `tfsdk:"updated_at"`
}

// ToBody
func ResponseOrganizationsGetOrganizationPolicyObjectItemToBody(state OrganizationsPolicyObjects, response *merakigosdk.ResponseOrganizationsGetOrganizationPolicyObject) OrganizationsPolicyObjects {
	itemState := ResponseOrganizationsGetOrganizationPolicyObject{
		Category:   types.StringValue(response.Category),
		Cidr:       types.StringValue(response.Cidr),
		CreatedAt:  types.StringValue(response.CreatedAt),
		GroupIDs:   StringSliceToList(response.GroupIDs),
		ID:         types.StringValue(response.ID),
		Name:       types.StringValue(response.Name),
		NetworkIDs: StringSliceToList(response.NetworkIDs),
		Type:       types.StringValue(response.Type),
		UpdatedAt:  types.StringValue(response.UpdatedAt),
	}
	state.Item = &itemState
	return state
}
