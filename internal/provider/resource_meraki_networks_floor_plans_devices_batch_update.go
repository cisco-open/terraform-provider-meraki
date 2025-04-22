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
	_ resource.Resource              = &NetworksFloorPlansDevicesBatchUpdateResource{}
	_ resource.ResourceWithConfigure = &NetworksFloorPlansDevicesBatchUpdateResource{}
)

func NewNetworksFloorPlansDevicesBatchUpdateResource() resource.Resource {
	return &NetworksFloorPlansDevicesBatchUpdateResource{}
}

type NetworksFloorPlansDevicesBatchUpdateResource struct {
	client *merakigosdk.Client
}

func (r *NetworksFloorPlansDevicesBatchUpdateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksFloorPlansDevicesBatchUpdateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_floor_plans_devices_batch_update"
}

// resourceAction
func (r *NetworksFloorPlansDevicesBatchUpdateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"success": schema.BoolAttribute{
						MarkdownDescription: `Status of attempt to update device floorplan assignments`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"assignments": schema.SetNestedAttribute{
						MarkdownDescription: `List of floorplan assignments to update. Up to 100 floor plan assignments can be provided in a request.`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"floor_plan": schema.SingleNestedAttribute{
									MarkdownDescription: `Floorplan to be assigned or unassigned`,
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `The ID of the floor plan to assign the device to, or null to unassign the device from its floor plan`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
											},
										},
									},
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Serial of the device to change the floor plan assignment for`,
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
func (r *NetworksFloorPlansDevicesBatchUpdateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksFloorPlansDevicesBatchUpdate

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
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Networks.BatchNetworkFloorPlansDevicesUpdate(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing BatchNetworkFloorPlansDevicesUpdate",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing BatchNetworkFloorPlansDevicesUpdate",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksBatchNetworkFloorPlansDevicesUpdateItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksFloorPlansDevicesBatchUpdateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksFloorPlansDevicesBatchUpdateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksFloorPlansDevicesBatchUpdateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksFloorPlansDevicesBatchUpdate struct {
	NetworkID  types.String                                          `tfsdk:"network_id"`
	Item       *ResponseNetworksBatchNetworkFloorPlansDevicesUpdate  `tfsdk:"item"`
	Parameters *RequestNetworksBatchNetworkFloorPlansDevicesUpdateRs `tfsdk:"parameters"`
}

type ResponseNetworksBatchNetworkFloorPlansDevicesUpdate struct {
	Success types.Bool `tfsdk:"success"`
}

type RequestNetworksBatchNetworkFloorPlansDevicesUpdateRs struct {
	Assignments *[]RequestNetworksBatchNetworkFloorPlansDevicesUpdateAssignmentsRs `tfsdk:"assignments"`
}

type RequestNetworksBatchNetworkFloorPlansDevicesUpdateAssignmentsRs struct {
	FloorPlan *RequestNetworksBatchNetworkFloorPlansDevicesUpdateAssignmentsFloorPlanRs `tfsdk:"floor_plan"`
	Serial    types.String                                                              `tfsdk:"serial"`
}

type RequestNetworksBatchNetworkFloorPlansDevicesUpdateAssignmentsFloorPlanRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *NetworksFloorPlansDevicesBatchUpdate) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksBatchNetworkFloorPlansDevicesUpdate {
	re := *r.Parameters
	var requestNetworksBatchNetworkFloorPlansDevicesUpdateAssignments []merakigosdk.RequestNetworksBatchNetworkFloorPlansDevicesUpdateAssignments

	if re.Assignments != nil {
		for _, rItem1 := range *re.Assignments {
			var requestNetworksBatchNetworkFloorPlansDevicesUpdateAssignmentsFloorPlan *merakigosdk.RequestNetworksBatchNetworkFloorPlansDevicesUpdateAssignmentsFloorPlan

			if rItem1.FloorPlan != nil {
				id := rItem1.FloorPlan.ID.ValueString()
				requestNetworksBatchNetworkFloorPlansDevicesUpdateAssignmentsFloorPlan = &merakigosdk.RequestNetworksBatchNetworkFloorPlansDevicesUpdateAssignmentsFloorPlan{
					ID: id,
				}
				//[debug] Is Array: False
			}
			serial := rItem1.Serial.ValueString()
			requestNetworksBatchNetworkFloorPlansDevicesUpdateAssignments = append(requestNetworksBatchNetworkFloorPlansDevicesUpdateAssignments, merakigosdk.RequestNetworksBatchNetworkFloorPlansDevicesUpdateAssignments{
				FloorPlan: requestNetworksBatchNetworkFloorPlansDevicesUpdateAssignmentsFloorPlan,
				Serial:    serial,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestNetworksBatchNetworkFloorPlansDevicesUpdate{
		Assignments: &requestNetworksBatchNetworkFloorPlansDevicesUpdateAssignments,
	}
	return &out
}

// ToBody
func ResponseNetworksBatchNetworkFloorPlansDevicesUpdateItemToBody(state NetworksFloorPlansDevicesBatchUpdate, response *merakigosdk.ResponseNetworksBatchNetworkFloorPlansDevicesUpdate) NetworksFloorPlansDevicesBatchUpdate {
	itemState := ResponseNetworksBatchNetworkFloorPlansDevicesUpdate{
		Success: func() types.Bool {
			if response.Success != nil {
				return types.BoolValue(*response.Success)
			}
			return types.Bool{}
		}(),
	}
	state.Item = &itemState
	return state
}
