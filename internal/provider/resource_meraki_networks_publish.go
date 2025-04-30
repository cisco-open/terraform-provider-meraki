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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksPublishResource{}
	_ resource.ResourceWithConfigure = &NetworksPublishResource{}
)

func NewNetworksPublishResource() resource.Resource {
	return &NetworksPublishResource{}
}

type NetworksPublishResource struct {
	client *merakigosdk.Client
}

func (r *NetworksPublishResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksPublishResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_publish"
}

// resourceAction
func (r *NetworksPublishResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
						MarkdownDescription: `Status of attempt to publish auto locate job`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"devices": schema.SetNestedAttribute{
						MarkdownDescription: `The list of devices to publish positions for`,
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
											MarkdownDescription: `Whether or not this device's location should be saved as a user-defined anchor`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.Bool{
												boolplanmodifier.RequiresReplace(),
											},
										},
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
								"serial": schema.StringAttribute{
									MarkdownDescription: `Serial for device to publish position for`,
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
func (r *NetworksPublishResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksPublish

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
	response, restyResp1, err := r.client.Networks.PublishNetworkFloorPlansAutoLocateJob(vvNetworkID, vvJobID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing PublishNetworkFloorPlansAutoLocateJob",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing PublishNetworkFloorPlansAutoLocateJob",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksPublishNetworkFloorPlansAutoLocateJobItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksPublishResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksPublishResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksPublishResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksPublish struct {
	NetworkID  types.String                                            `tfsdk:"network_id"`
	JobID      types.String                                            `tfsdk:"job_id"`
	Item       *ResponseNetworksPublishNetworkFloorPlansAutoLocateJob  `tfsdk:"item"`
	Parameters *RequestNetworksPublishNetworkFloorPlansAutoLocateJobRs `tfsdk:"parameters"`
}

type ResponseNetworksPublishNetworkFloorPlansAutoLocateJob struct {
	Success types.Bool `tfsdk:"success"`
}

type RequestNetworksPublishNetworkFloorPlansAutoLocateJobRs struct {
	Devices *[]RequestNetworksPublishNetworkFloorPlansAutoLocateJobDevicesRs `tfsdk:"devices"`
}

type RequestNetworksPublishNetworkFloorPlansAutoLocateJobDevicesRs struct {
	AutoLocate *RequestNetworksPublishNetworkFloorPlansAutoLocateJobDevicesAutoLocateRs `tfsdk:"auto_locate"`
	Lat        types.Float64                                                            `tfsdk:"lat"`
	Lng        types.Float64                                                            `tfsdk:"lng"`
	Serial     types.String                                                             `tfsdk:"serial"`
}

type RequestNetworksPublishNetworkFloorPlansAutoLocateJobDevicesAutoLocateRs struct {
	IsAnchor types.Bool `tfsdk:"is_anchor"`
}

// FromBody
func (r *NetworksPublish) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksPublishNetworkFloorPlansAutoLocateJob {
	re := *r.Parameters
	var requestNetworksPublishNetworkFloorPlansAutoLocateJobDevices []merakigosdk.RequestNetworksPublishNetworkFloorPlansAutoLocateJobDevices

	if re.Devices != nil {
		for _, rItem1 := range *re.Devices {
			var requestNetworksPublishNetworkFloorPlansAutoLocateJobDevicesAutoLocate *merakigosdk.RequestNetworksPublishNetworkFloorPlansAutoLocateJobDevicesAutoLocate

			if rItem1.AutoLocate != nil {
				isAnchor := func() *bool {
					if !rItem1.AutoLocate.IsAnchor.IsUnknown() && !rItem1.AutoLocate.IsAnchor.IsNull() {
						return rItem1.AutoLocate.IsAnchor.ValueBoolPointer()
					}
					return nil
				}()
				requestNetworksPublishNetworkFloorPlansAutoLocateJobDevicesAutoLocate = &merakigosdk.RequestNetworksPublishNetworkFloorPlansAutoLocateJobDevicesAutoLocate{
					IsAnchor: isAnchor,
				}
				//[debug] Is Array: False
			}
			lat := func() *float64 {
				if !rItem1.Lat.IsUnknown() && !rItem1.Lat.IsNull() {
					return rItem1.Lat.ValueFloat64Pointer()
				}
				return nil
			}()
			lng := func() *float64 {
				if !rItem1.Lng.IsUnknown() && !rItem1.Lng.IsNull() {
					return rItem1.Lng.ValueFloat64Pointer()
				}
				return nil
			}()
			serial := rItem1.Serial.ValueString()
			requestNetworksPublishNetworkFloorPlansAutoLocateJobDevices = append(requestNetworksPublishNetworkFloorPlansAutoLocateJobDevices, merakigosdk.RequestNetworksPublishNetworkFloorPlansAutoLocateJobDevices{
				AutoLocate: requestNetworksPublishNetworkFloorPlansAutoLocateJobDevicesAutoLocate,
				Lat:        lat,
				Lng:        lng,
				Serial:     serial,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestNetworksPublishNetworkFloorPlansAutoLocateJob{
		Devices: &requestNetworksPublishNetworkFloorPlansAutoLocateJobDevices,
	}
	return &out
}

// ToBody
func ResponseNetworksPublishNetworkFloorPlansAutoLocateJobItemToBody(state NetworksPublish, response *merakigosdk.ResponseNetworksPublishNetworkFloorPlansAutoLocateJob) NetworksPublish {
	itemState := ResponseNetworksPublishNetworkFloorPlansAutoLocateJob{
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
