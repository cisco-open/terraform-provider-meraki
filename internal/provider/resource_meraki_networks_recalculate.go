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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksRecalculateResource{}
	_ resource.ResourceWithConfigure = &NetworksRecalculateResource{}
)

func NewNetworksRecalculateResource() resource.Resource {
	return &NetworksRecalculateResource{}
}

type NetworksRecalculateResource struct {
	client *merakigosdk.Client
}

func (r *NetworksRecalculateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksRecalculateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_recalculate"
}

// resourceAction
func (r *NetworksRecalculateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"job_id": schema.StringAttribute{
				MarkdownDescription: `jobId path parameter. Job ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
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
						MarkdownDescription: `Status of attempt to trigger auto locate recalculation`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"devices": schema.SetNestedAttribute{
						MarkdownDescription: `The list of devices to update anchor positions for`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"auto_locate": schema.SingleNestedAttribute{
									MarkdownDescription: `The auto locate position for this device`,
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"is_anchor": schema.BoolAttribute{
											MarkdownDescription: `Whether or not this location should be saved as a user-defined anchor`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.Bool{
												boolplanmodifier.RequiresReplace(),
											},
										},
										"lat": schema.Float64Attribute{
											MarkdownDescription: `Latitude`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.RequiresReplace(),
											},
										},
										"lng": schema.Float64Attribute{
											MarkdownDescription: `Longitude`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.Float64{
												float64planmodifier.RequiresReplace(),
											},
										},
									},
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Serial for device to update`,
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
func (r *NetworksRecalculateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksRecalculate

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
	vvJobID := data.JobID.ValueString()
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Networks.RecalculateNetworkFloorPlansAutoLocateJob(vvNetworkID, vvJobID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing RecalculateNetworkFloorPlansAutoLocateJob",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing RecalculateNetworkFloorPlansAutoLocateJob",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksRecalculateNetworkFloorPlansAutoLocateJobItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksRecalculateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksRecalculateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksRecalculateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksRecalculate struct {
	NetworkID  types.String                                                `tfsdk:"network_id"`
	JobID      types.String                                                `tfsdk:"job_id"`
	Item       *ResponseNetworksRecalculateNetworkFloorPlansAutoLocateJob  `tfsdk:"item"`
	Parameters *RequestNetworksRecalculateNetworkFloorPlansAutoLocateJobRs `tfsdk:"parameters"`
}

type ResponseNetworksRecalculateNetworkFloorPlansAutoLocateJob struct {
	Success types.Bool `tfsdk:"success"`
}

type RequestNetworksRecalculateNetworkFloorPlansAutoLocateJobRs struct {
	Devices *[]RequestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevicesRs `tfsdk:"devices"`
}

type RequestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevicesRs struct {
	AutoLocate *RequestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevicesAutoLocateRs `tfsdk:"auto_locate"`
	Serial     types.String                                                                 `tfsdk:"serial"`
}

type RequestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevicesAutoLocateRs struct {
	IsAnchor types.Bool    `tfsdk:"is_anchor"`
	Lat      types.Float64 `tfsdk:"lat"`
	Lng      types.Float64 `tfsdk:"lng"`
}

// FromBody
func (r *NetworksRecalculate) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksRecalculateNetworkFloorPlansAutoLocateJob {
	re := *r.Parameters
	var requestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevices []merakigosdk.RequestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevices

	if re.Devices != nil {
		for _, rItem1 := range *re.Devices {
			var requestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevicesAutoLocate *merakigosdk.RequestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevicesAutoLocate

			if rItem1.AutoLocate != nil {
				isAnchor := func() *bool {
					if !rItem1.AutoLocate.IsAnchor.IsUnknown() && !rItem1.AutoLocate.IsAnchor.IsNull() {
						return rItem1.AutoLocate.IsAnchor.ValueBoolPointer()
					}
					return nil
				}()
				lat := func() *float64 {
					if !rItem1.AutoLocate.Lat.IsUnknown() && !rItem1.AutoLocate.Lat.IsNull() {
						return rItem1.AutoLocate.Lat.ValueFloat64Pointer()
					}
					return nil
				}()
				lng := func() *float64 {
					if !rItem1.AutoLocate.Lng.IsUnknown() && !rItem1.AutoLocate.Lng.IsNull() {
						return rItem1.AutoLocate.Lng.ValueFloat64Pointer()
					}
					return nil
				}()
				requestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevicesAutoLocate = &merakigosdk.RequestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevicesAutoLocate{
					IsAnchor: isAnchor,
					Lat:      lat,
					Lng:      lng,
				}
				//[debug] Is Array: False
			}
			serial := rItem1.Serial.ValueString()
			requestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevices = append(requestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevices, merakigosdk.RequestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevices{
				AutoLocate: requestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevicesAutoLocate,
				Serial:     serial,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestNetworksRecalculateNetworkFloorPlansAutoLocateJob{
		Devices: &requestNetworksRecalculateNetworkFloorPlansAutoLocateJobDevices,
	}
	return &out
}

// ToBody
func ResponseNetworksRecalculateNetworkFloorPlansAutoLocateJobItemToBody(state NetworksRecalculate, response *merakigosdk.ResponseNetworksRecalculateNetworkFloorPlansAutoLocateJob) NetworksRecalculate {
	itemState := ResponseNetworksRecalculateNetworkFloorPlansAutoLocateJob{
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
