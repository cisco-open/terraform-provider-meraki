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
	_ resource.Resource              = &OrganizationsDevicesDetailsBulkUpdateResource{}
	_ resource.ResourceWithConfigure = &OrganizationsDevicesDetailsBulkUpdateResource{}
)

func NewOrganizationsDevicesDetailsBulkUpdateResource() resource.Resource {
	return &OrganizationsDevicesDetailsBulkUpdateResource{}
}

type OrganizationsDevicesDetailsBulkUpdateResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsDevicesDetailsBulkUpdateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsDevicesDetailsBulkUpdateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_details_bulk_update"
}

// resourceAction
func (r *OrganizationsDevicesDetailsBulkUpdateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"serials": schema.ListAttribute{
						MarkdownDescription: `A list of serials of devices updated`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"details": schema.SetNestedAttribute{
						MarkdownDescription: `An array of details`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"name": schema.StringAttribute{
									MarkdownDescription: `Name of device detail`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"value": schema.StringAttribute{
									MarkdownDescription: `Value of device detail`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
							},
						},
					},
					"serials": schema.ListAttribute{
						MarkdownDescription: `A list of serials of devices to update`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *OrganizationsDevicesDetailsBulkUpdateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsDevicesDetailsBulkUpdate

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
	response, restyResp1, err := r.client.Organizations.BulkUpdateOrganizationDevicesDetails(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing BulkUpdateOrganizationDevicesDetails",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing BulkUpdateOrganizationDevicesDetails",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseOrganizationsBulkUpdateOrganizationDevicesDetailsItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsDevicesDetailsBulkUpdateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsDevicesDetailsBulkUpdateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsDevicesDetailsBulkUpdateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsDevicesDetailsBulkUpdate struct {
	OrganizationID types.String                                                `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsBulkUpdateOrganizationDevicesDetails  `tfsdk:"item"`
	Parameters     *RequestOrganizationsBulkUpdateOrganizationDevicesDetailsRs `tfsdk:"parameters"`
}

type ResponseOrganizationsBulkUpdateOrganizationDevicesDetails struct {
	Serials types.List `tfsdk:"serials"`
}

type RequestOrganizationsBulkUpdateOrganizationDevicesDetailsRs struct {
	Details *[]RequestOrganizationsBulkUpdateOrganizationDevicesDetailsDetailsRs `tfsdk:"details"`
	Serials types.List                                                           `tfsdk:"serials"`
}

type RequestOrganizationsBulkUpdateOrganizationDevicesDetailsDetailsRs struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// FromBody
func (r *OrganizationsDevicesDetailsBulkUpdate) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsBulkUpdateOrganizationDevicesDetails {
	re := *r.Parameters
	var requestOrganizationsBulkUpdateOrganizationDevicesDetailsDetails []merakigosdk.RequestOrganizationsBulkUpdateOrganizationDevicesDetailsDetails

	if re.Details != nil {
		for _, rItem1 := range *re.Details {
			name := rItem1.Name.ValueString()
			value := rItem1.Value.ValueString()
			requestOrganizationsBulkUpdateOrganizationDevicesDetailsDetails = append(requestOrganizationsBulkUpdateOrganizationDevicesDetailsDetails, merakigosdk.RequestOrganizationsBulkUpdateOrganizationDevicesDetailsDetails{
				Name:  name,
				Value: value,
			})
			//[debug] Is Array: True
		}
	}
	var serials []string = nil
	re.Serials.ElementsAs(ctx, &serials, false)
	out := merakigosdk.RequestOrganizationsBulkUpdateOrganizationDevicesDetails{
		Details: &requestOrganizationsBulkUpdateOrganizationDevicesDetailsDetails,
		Serials: serials,
	}
	return &out
}

// ToBody
func ResponseOrganizationsBulkUpdateOrganizationDevicesDetailsItemToBody(state OrganizationsDevicesDetailsBulkUpdate, response *merakigosdk.ResponseOrganizationsBulkUpdateOrganizationDevicesDetails) OrganizationsDevicesDetailsBulkUpdate {
	itemState := ResponseOrganizationsBulkUpdateOrganizationDevicesDetails{
		Serials: StringSliceToList(response.Serials),
	}
	state.Item = &itemState
	return state
}
