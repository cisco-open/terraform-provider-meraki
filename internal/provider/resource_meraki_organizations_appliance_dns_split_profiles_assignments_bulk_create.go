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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreateResource{}
	_ resource.ResourceWithConfigure = &ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreateResource{}
)

func NewApplianceApplianceDNSSplitProfilesAssignmentsBulkCreateResource() resource.Resource {
	return &ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreateResource{}
}

type ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreateResource struct {
	client *merakigosdk.Client
}

func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_Appliance_appliance_dns_split_profiles_assignments_bulk_create"
}

// resourceAction
func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreate

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
	response, restyResp1, err := r.client.Appliance.CreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreate struct {
	OrganizationID types.String                                                                        `tfsdk:"organization_id"`
	Item           *ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate  `tfsdk:"item"`
	Parameters     *RequestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateRs `tfsdk:"parameters"`
}

type ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate struct {
	Items *[]ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItems `tfsdk:"items"`
}

type ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItems struct {
	AssignmentID types.String                                                                                   `tfsdk:"assignment_id"`
	Network      *ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsNetwork `tfsdk:"network"`
	Profile      *ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsProfile `tfsdk:"profile"`
}

type ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsProfile struct {
	ID types.String `tfsdk:"id"`
}

type RequestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateRs struct {
	Items *[]RequestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsRs `tfsdk:"items"`
}

type RequestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsRs struct {
	Network *RequestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsNetworkRs `tfsdk:"network"`
	Profile *RequestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsProfileRs `tfsdk:"profile"`
}

type RequestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsNetworkRs struct {
	ID types.String `tfsdk:"id"`
}

type RequestApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsProfileRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreate) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate {
	re := *r.Parameters
	var requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItems []merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItems

	if re.Items != nil {
		for _, rItem1 := range *re.Items {
			var requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsNetwork *merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsNetwork

			if rItem1.Network != nil {
				id := rItem1.Network.ID.ValueString()
				requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsNetwork = &merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsNetwork{
					ID: id,
				}
				//[debug] Is Array: False
			}
			var requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsProfile *merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsProfile

			if rItem1.Profile != nil {
				id := rItem1.Profile.ID.ValueString()
				requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsProfile = &merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsProfile{
					ID: id,
				}
				//[debug] Is Array: False
			}
			requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItems = append(requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItems, merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItems{
				Network: requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsNetwork,
				Profile: requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemsProfile,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate{
		Items: &requestApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItems,
	}
	return &out
}

// ToBody
func ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreateItemToBody(state ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreate, response *merakigosdk.ResponseApplianceCreateOrganizationApplianceDNSSplitProfilesAssignmentsBulkCreate) ApplianceApplianceDNSSplitProfilesAssignmentsBulkCreate {
	itemState := ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreate{
		Items: func() *[]ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItems {
			if response.Items != nil {
				result := make([]ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItems{
						AssignmentID: types.StringValue(items.AssignmentID),
						Network: func() *ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsNetwork {
							if items.Network != nil {
								return &ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsNetwork{
									ID: types.StringValue(items.Network.ID),
								}
							}
							return nil
						}(),
						Profile: func() *ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsProfile {
							if items.Profile != nil {
								return &ResponseApplianceCreateOrganizationApplianceDnsSplitProfilesAssignmentsBulkCreateItemsProfile{
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
