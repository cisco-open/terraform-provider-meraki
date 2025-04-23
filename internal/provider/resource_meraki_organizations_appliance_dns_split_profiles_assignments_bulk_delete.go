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
	_ resource.Resource              = &ApplianceApplianceDNSSplitProfilesAssignmentsBulkDeleteResource{}
	_ resource.ResourceWithConfigure = &ApplianceApplianceDNSSplitProfilesAssignmentsBulkDeleteResource{}
)

func NewApplianceApplianceDNSSplitProfilesAssignmentsBulkDeleteResource() resource.Resource {
	return &ApplianceApplianceDNSSplitProfilesAssignmentsBulkDeleteResource{}
}

type ApplianceApplianceDNSSplitProfilesAssignmentsBulkDeleteResource struct {
	client *merakigosdk.Client
}

func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkDeleteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkDeleteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_Appliance_appliance_dns_split_profiles_assignments_bulk_delete"
}

// resourceAction
func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkDeleteResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
						MarkdownDescription: `List of split DNS profile assignment`,
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
						MarkdownDescription: `List containing the assignment ID`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"assignment_id": schema.StringAttribute{
									MarkdownDescription: `ID of the assignment`,
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
	}
}
func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkDeleteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data ApplianceApplianceDNSSplitProfilesAssignmentsBulkDelete

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
	response, restyResp1, err := r.client.Appliance.CreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkDeleteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkDeleteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkDeleteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type ApplianceApplianceDNSSplitProfilesAssignmentsBulkDelete struct {
	OrganizationID types.String                                                                        `tfsdk:"organization_id"`
	Item           *ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete  `tfsdk:"item"`
	Parameters     *RequestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRs `tfsdk:"parameters"`
}

type ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete struct {
	Items *[]ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItems `tfsdk:"items"`
}

type ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItems struct {
	AssignmentID types.String                                                                                   `tfsdk:"assignment_id"`
	Network      *ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItemsNetwork `tfsdk:"network"`
	Profile      *ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItemsProfile `tfsdk:"profile"`
}

type ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItemsProfile struct {
	ID types.String `tfsdk:"id"`
}

type RequestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteRs struct {
	Items *[]RequestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItemsRs `tfsdk:"items"`
}

type RequestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItemsRs struct {
	AssignmentID types.String `tfsdk:"assignment_id"`
}

// FromBody
func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkDelete) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete {
	re := *r.Parameters
	var requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItems []merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItems

	if re.Items != nil {
		for _, rItem1 := range *re.Items {
			assignmentID := rItem1.AssignmentID.ValueString()
			requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItems = append(requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItems, merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItems{
				AssignmentID: assignmentID,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete{
		Items: &requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItems,
	}
	return &out
}

// ToBody
func ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDeleteItemToBody(state ApplianceApplianceDNSSplitProfilesAssignmentsBulkDelete, response *merakigosdk.ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkDelete) ApplianceApplianceDNSSplitProfilesAssignmentsBulkDelete {
	itemState := ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDelete{
		Items: func() *[]ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItems {
			if response.Items != nil {
				result := make([]ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItems{
						AssignmentID: types.StringValue(items.AssignmentID),
						Network: func() *ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItemsNetwork {
							if items.Network != nil {
								return &ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItemsNetwork{
									ID: types.StringValue(items.Network.ID),
								}
							}
							return nil
						}(),
						Profile: func() *ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItemsProfile {
							if items.Profile != nil {
								return &ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkDeleteItemsProfile{
									ID: types.StringValue(items.Profile.ID),
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
