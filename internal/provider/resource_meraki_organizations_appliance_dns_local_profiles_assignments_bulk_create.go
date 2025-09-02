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
	_ resource.Resource              = &OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreateResource{}
	_ resource.ResourceWithConfigure = &OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreateResource{}
)

func NewOrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreateResource() resource.Resource {
	return &OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreateResource{}
}

type OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreateResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_dns_local_profiles_assignments_bulk_create"
}

// resourceAction
func (r *OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
						MarkdownDescription: `List of local DNS profile assignment`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"assignment_id": schema.StringAttribute{
									MarkdownDescription: `ID of the assignment`,
									Computed:            true,
								},
								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `The network attached to the profile`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the network`,
											Computed:            true,
										},
									},
								},
								"profile": schema.SingleNestedAttribute{
									MarkdownDescription: `The profile the network is attached to`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the profile`,
											Computed:            true,
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
						MarkdownDescription: `List containing the network ID and Profile ID`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"network": schema.SingleNestedAttribute{
									MarkdownDescription: `The network attached to the profile`,
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the network`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
											},
										},
									},
								},
								"profile": schema.SingleNestedAttribute{
									MarkdownDescription: `The profile the network is attached to`,
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the profile`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
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
func (r *OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreate

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Appliance.BulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing BulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing BulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreate struct {
	OrganizationID types.String                                                                  `tfsdk:"organization_id"`
	Item           *ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate  `tfsdk:"item"`
	Parameters     *RequestApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateRs `tfsdk:"parameters"`
}

type ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate struct {
	Items *[]ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItems `tfsdk:"items"`
}

type ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItems struct {
	AssignmentID types.String                                                                             `tfsdk:"assignment_id"`
	Network      *ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsNetwork `tfsdk:"network"`
	Profile      *ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsProfile `tfsdk:"profile"`
}

type ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsProfile struct {
	ID types.String `tfsdk:"id"`
}

type RequestApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateRs struct {
	Items *[]RequestApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsRs `tfsdk:"items"`
}

type RequestApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsRs struct {
	Network *RequestApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsNetworkRs `tfsdk:"network"`
	Profile *RequestApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsProfileRs `tfsdk:"profile"`
}

type RequestApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsNetworkRs struct {
	ID types.String `tfsdk:"id"`
}

type RequestApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsProfileRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreate) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate {
	re := *r.Parameters
	var requestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItems []merakigosdk.RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItems

	if re.Items != nil {
		for _, rItem1 := range *re.Items {
			var requestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsNetwork *merakigosdk.RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsNetwork

			if rItem1.Network != nil {
				id := rItem1.Network.ID.ValueString()
				requestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsNetwork = &merakigosdk.RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsNetwork{
					ID: id,
				}
				//[debug] Is Array: False
			}
			var requestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsProfile *merakigosdk.RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsProfile

			if rItem1.Profile != nil {
				id := rItem1.Profile.ID.ValueString()
				requestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsProfile = &merakigosdk.RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsProfile{
					ID: id,
				}
				//[debug] Is Array: False
			}
			requestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItems = append(requestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItems, merakigosdk.RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItems{
				Network: requestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsNetwork,
				Profile: requestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemsProfile,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate{
		Items: &requestApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItems,
	}
	return &out
}

// ToBody
func ResponseApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreateItemToBody(state OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreate, response *merakigosdk.ResponseApplianceBulkOrganizationApplianceDNSLocalProfilesAssignmentsCreate) OrganizationsApplianceDNSLocalProfilesAssignmentsBulkCreate {
	itemState := ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreate{
		Items: func() *[]ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItems {
			if response.Items != nil {
				result := make([]ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItems{
						AssignmentID: func() types.String {
							if items.AssignmentID != "" {
								return types.StringValue(items.AssignmentID)
							}
							return types.String{}
						}(),
						Network: func() *ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsNetwork {
							if items.Network != nil {
								return &ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsNetwork{
									ID: func() types.String {
										if items.Network.ID != "" {
											return types.StringValue(items.Network.ID)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
						Profile: func() *ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsProfile {
							if items.Profile != nil {
								return &ResponseApplianceBulkOrganizationApplianceDnsLocalProfilesAssignmentsCreateItemsProfile{
									ID: func() types.String {
										if items.Profile.ID != "" {
											return types.StringValue(items.Profile.ID)
										}
										return types.String{}
									}(),
								}
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
