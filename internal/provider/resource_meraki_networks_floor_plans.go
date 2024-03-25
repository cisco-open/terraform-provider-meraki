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

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksFloorPlansResource{}
	_ resource.ResourceWithConfigure = &NetworksFloorPlansResource{}
)

func NewNetworksFloorPlansResource() resource.Resource {
	return &NetworksFloorPlansResource{}
}

type NetworksFloorPlansResource struct {
	client *merakigosdk.Client
}

func (r *NetworksFloorPlansResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksFloorPlansResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_floor_plans"
}

func (r *NetworksFloorPlansResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"bottom_left_corner": schema.SingleNestedAttribute{
				MarkdownDescription: `The longitude and latitude of the bottom left corner of your floor plan.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"lat": schema.Float64Attribute{
						MarkdownDescription: `Latitude`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"lng": schema.Float64Attribute{
						MarkdownDescription: `Longitude`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"bottom_right_corner": schema.SingleNestedAttribute{
				MarkdownDescription: `The longitude and latitude of the bottom right corner of your floor plan.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"lat": schema.Float64Attribute{
						MarkdownDescription: `Latitude`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"lng": schema.Float64Attribute{
						MarkdownDescription: `Longitude`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"center": schema.SingleNestedAttribute{
				MarkdownDescription: `The longitude and latitude of the center of your floor plan. The 'center' or two adjacent corners (e.g. 'topLeftCorner' and 'bottomLeftCorner') must be specified. If 'center' is specified, the floor plan is placed over that point with no rotation. If two adjacent corners are specified, the floor plan is rotated to line up with the two specified points. The aspect ratio of the floor plan's image is preserved regardless of which corners/center are specified. (This means if that more than two corners are specified, only two corners may be used to preserve the floor plan's aspect ratio.). No two points can have the same latitude, longitude pair.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"lat": schema.Float64Attribute{
						MarkdownDescription: `Latitude`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"lng": schema.Float64Attribute{
						MarkdownDescription: `Longitude`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"devices": schema.SetNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"address": schema.StringAttribute{
							Computed: true,
						},
						"beacon_id_params": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"major": schema.Int64Attribute{
									Computed: true,
								},
								"minor": schema.Int64Attribute{
									Computed: true,
								},
								"uuid": schema.StringAttribute{
									Computed: true,
								},
							},
						},
						"firmware": schema.StringAttribute{
							Computed: true,
						},
						"floor_plan_id": schema.StringAttribute{
							Computed: true,
						},
						"lan_ip": schema.StringAttribute{
							Computed: true,
						},
						"lat": schema.Float64Attribute{
							Computed: true,
						},
						"lng": schema.Float64Attribute{
							Computed: true,
						},
						"mac": schema.StringAttribute{
							Computed: true,
						},
						"model": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"network_id": schema.StringAttribute{
							Computed: true,
						},
						"notes": schema.StringAttribute{
							Computed: true,
						},
						"serial": schema.StringAttribute{
							Computed: true,
						},
						"tags": schema.SetAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
			"floor_plan_id": schema.StringAttribute{
				MarkdownDescription: `floorPlanId path parameter. Floor plan ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"height": schema.Float64Attribute{
				Computed: true,
			},
			"image_contents": schema.StringAttribute{
				MarkdownDescription: `The file contents (a base 64 encoded string) of your image. Supported formats are PNG, GIF, and JPG. Note that all images are saved as PNG files, regardless of the format they are uploaded in.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"image_extension": schema.StringAttribute{
				Computed: true,
			},
			"image_md5": schema.StringAttribute{
				Computed: true,
			},
			"image_url": schema.StringAttribute{
				Computed: true,
			},
			"image_url_expires_at": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of your floor plan.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"top_left_corner": schema.SingleNestedAttribute{
				MarkdownDescription: `The longitude and latitude of the top left corner of your floor plan.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"lat": schema.Float64Attribute{
						MarkdownDescription: `Latitude`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"lng": schema.Float64Attribute{
						MarkdownDescription: `Longitude`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"top_right_corner": schema.SingleNestedAttribute{
				MarkdownDescription: `The longitude and latitude of the top right corner of your floor plan.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"lat": schema.Float64Attribute{
						MarkdownDescription: `Latitude`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"lng": schema.Float64Attribute{
						MarkdownDescription: `Longitude`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"width": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
}

//path params to set ['floorPlanId']

func (r *NetworksFloorPlansResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksFloorPlansRs

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
	// network_id
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkFloorPlans(vvNetworkID)
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFloorPlans",
				err.Error(),
			)
			return
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvFloorPlanID, ok := result2["FloorPlanID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter FloorPlanID",
					"Error",
				)
				return
			}
			r.client.Networks.UpdateNetworkFloorPlan(vvNetworkID, vvFloorPlanID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Networks.GetNetworkFloorPlan(vvNetworkID, vvFloorPlanID)
			if responseVerifyItem2 != nil {
				data = ResponseNetworksGetNetworkFloorPlanItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp2, err := r.client.Networks.CreateNetworkFloorPlan(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkFloorPlan",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkFloorPlan",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Networks.GetNetworkFloorPlans(vvNetworkID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFloorPlans",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkFloorPlans",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvFloorPlanID, ok := result2["FloorPlanID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter FloorPlanID",
				"Error",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Networks.GetNetworkFloorPlan(vvNetworkID, vvFloorPlanID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseNetworksGetNetworkFloorPlanItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkFloorPlan",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFloorPlan",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}
}

func (r *NetworksFloorPlansResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksFloorPlansRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvFloorPlanID := data.FloorPlanID.ValueString()
	// floor_plan_id
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkFloorPlan(vvNetworkID, vvFloorPlanID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFloorPlan",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkFloorPlan",
			err.Error(),
		)
		return
	}

	data = ResponseNetworksGetNetworkFloorPlanItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksFloorPlansResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("floor_plan_id"), idParts[1])...)
}

func (r *NetworksFloorPlansResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksFloorPlansRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvFloorPlanID := data.FloorPlanID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Networks.UpdateNetworkFloorPlan(vvNetworkID, vvFloorPlanID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkFloorPlan",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkFloorPlan",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksFloorPlansResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksFloorPlansRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvNetworkID := state.NetworkID.ValueString()
	vvFloorPlanID := state.FloorPlanID.ValueString()
	_, err := r.client.Networks.DeleteNetworkFloorPlan(vvNetworkID, vvFloorPlanID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkFloorPlan", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksFloorPlansRs struct {
	NetworkID         types.String                                            `tfsdk:"network_id"`
	FloorPlanID       types.String                                            `tfsdk:"floor_plan_id"`
	BottomLeftCorner  *ResponseNetworksGetNetworkFloorPlanBottomLeftCornerRs  `tfsdk:"bottom_left_corner"`
	BottomRightCorner *ResponseNetworksGetNetworkFloorPlanBottomRightCornerRs `tfsdk:"bottom_right_corner"`
	Center            *ResponseNetworksGetNetworkFloorPlanCenterRs            `tfsdk:"center"`
	Devices           *[]ResponseNetworksGetNetworkFloorPlanDevicesRs         `tfsdk:"devices"`
	Height            types.Float64                                           `tfsdk:"height"`
	ImageExtension    types.String                                            `tfsdk:"image_extension"`
	ImageMd5          types.String                                            `tfsdk:"image_md5"`
	ImageURL          types.String                                            `tfsdk:"image_url"`
	ImageURLExpiresAt types.String                                            `tfsdk:"image_url_expires_at"`
	Name              types.String                                            `tfsdk:"name"`
	TopLeftCorner     *ResponseNetworksGetNetworkFloorPlanTopLeftCornerRs     `tfsdk:"top_left_corner"`
	TopRightCorner    *ResponseNetworksGetNetworkFloorPlanTopRightCornerRs    `tfsdk:"top_right_corner"`
	Width             types.Int64                                             `tfsdk:"width"`
	ImageContents     types.String                                            `tfsdk:"image_contents"`
}

type ResponseNetworksGetNetworkFloorPlanBottomLeftCornerRs struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlanBottomRightCornerRs struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlanCenterRs struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlanDevicesRs struct {
	Address        types.String                                                `tfsdk:"address"`
	BeaconIDParams *ResponseNetworksGetNetworkFloorPlanDevicesBeaconIdParamsRs `tfsdk:"beacon_id_params"`
	Firmware       types.String                                                `tfsdk:"firmware"`
	FloorPlanID    types.String                                                `tfsdk:"floor_plan_id"`
	LanIP          types.String                                                `tfsdk:"lan_ip"`
	Lat            types.Float64                                               `tfsdk:"lat"`
	Lng            types.Float64                                               `tfsdk:"lng"`
	Mac            types.String                                                `tfsdk:"mac"`
	Model          types.String                                                `tfsdk:"model"`
	Name           types.String                                                `tfsdk:"name"`
	NetworkID      types.String                                                `tfsdk:"network_id"`
	Notes          types.String                                                `tfsdk:"notes"`
	Serial         types.String                                                `tfsdk:"serial"`
	Tags           types.Set                                                   `tfsdk:"tags"`
}

type ResponseNetworksGetNetworkFloorPlanDevicesBeaconIdParamsRs struct {
	Major types.Int64  `tfsdk:"major"`
	Minor types.Int64  `tfsdk:"minor"`
	UUID  types.String `tfsdk:"uuid"`
}

type ResponseNetworksGetNetworkFloorPlanTopLeftCornerRs struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

type ResponseNetworksGetNetworkFloorPlanTopRightCornerRs struct {
	Lat types.Float64 `tfsdk:"lat"`
	Lng types.Float64 `tfsdk:"lng"`
}

// FromBody
func (r *NetworksFloorPlansRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksCreateNetworkFloorPlan {
	emptyString := ""
	var requestNetworksCreateNetworkFloorPlanBottomLeftCorner *merakigosdk.RequestNetworksCreateNetworkFloorPlanBottomLeftCorner
	if r.BottomLeftCorner != nil {
		lat := r.BottomLeftCorner.Lat.ValueFloat64Pointer()
		lng := r.BottomLeftCorner.Lng.ValueFloat64Pointer()
		requestNetworksCreateNetworkFloorPlanBottomLeftCorner = &merakigosdk.RequestNetworksCreateNetworkFloorPlanBottomLeftCorner{
			Lat: lat,
			Lng: lng,
		}
	}
	var requestNetworksCreateNetworkFloorPlanBottomRightCorner *merakigosdk.RequestNetworksCreateNetworkFloorPlanBottomRightCorner
	if r.BottomRightCorner != nil {
		lat := r.BottomRightCorner.Lat.ValueFloat64Pointer()
		lng := r.BottomRightCorner.Lng.ValueFloat64Pointer()
		requestNetworksCreateNetworkFloorPlanBottomRightCorner = &merakigosdk.RequestNetworksCreateNetworkFloorPlanBottomRightCorner{
			Lat: lat,
			Lng: lng,
		}
	}
	var requestNetworksCreateNetworkFloorPlanCenter *merakigosdk.RequestNetworksCreateNetworkFloorPlanCenter
	if r.Center != nil {
		lat := r.Center.Lat.ValueFloat64Pointer()
		lng := r.Center.Lng.ValueFloat64Pointer()
		requestNetworksCreateNetworkFloorPlanCenter = &merakigosdk.RequestNetworksCreateNetworkFloorPlanCenter{
			Lat: lat,
			Lng: lng,
		}
	}
	imageContents := new(string)
	if !r.ImageContents.IsUnknown() && !r.ImageContents.IsNull() {
		*imageContents = r.ImageContents.ValueString()
	} else {
		imageContents = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestNetworksCreateNetworkFloorPlanTopLeftCorner *merakigosdk.RequestNetworksCreateNetworkFloorPlanTopLeftCorner
	if r.TopLeftCorner != nil {
		lat := r.TopLeftCorner.Lat.ValueFloat64Pointer()
		lng := r.TopLeftCorner.Lng.ValueFloat64Pointer()
		requestNetworksCreateNetworkFloorPlanTopLeftCorner = &merakigosdk.RequestNetworksCreateNetworkFloorPlanTopLeftCorner{
			Lat: lat,
			Lng: lng,
		}
	}
	var requestNetworksCreateNetworkFloorPlanTopRightCorner *merakigosdk.RequestNetworksCreateNetworkFloorPlanTopRightCorner
	if r.TopRightCorner != nil {
		lat := r.TopRightCorner.Lat.ValueFloat64Pointer()
		lng := r.TopRightCorner.Lng.ValueFloat64Pointer()
		requestNetworksCreateNetworkFloorPlanTopRightCorner = &merakigosdk.RequestNetworksCreateNetworkFloorPlanTopRightCorner{
			Lat: lat,
			Lng: lng,
		}
	}
	out := merakigosdk.RequestNetworksCreateNetworkFloorPlan{
		BottomLeftCorner:  requestNetworksCreateNetworkFloorPlanBottomLeftCorner,
		BottomRightCorner: requestNetworksCreateNetworkFloorPlanBottomRightCorner,
		Center:            requestNetworksCreateNetworkFloorPlanCenter,
		ImageContents:     *imageContents,
		Name:              *name,
		TopLeftCorner:     requestNetworksCreateNetworkFloorPlanTopLeftCorner,
		TopRightCorner:    requestNetworksCreateNetworkFloorPlanTopRightCorner,
	}
	return &out
}
func (r *NetworksFloorPlansRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkFloorPlan {
	emptyString := ""
	var requestNetworksUpdateNetworkFloorPlanBottomLeftCorner *merakigosdk.RequestNetworksUpdateNetworkFloorPlanBottomLeftCorner
	if r.BottomLeftCorner != nil {
		lat := r.BottomLeftCorner.Lat.ValueFloat64Pointer()
		lng := r.BottomLeftCorner.Lng.ValueFloat64Pointer()
		requestNetworksUpdateNetworkFloorPlanBottomLeftCorner = &merakigosdk.RequestNetworksUpdateNetworkFloorPlanBottomLeftCorner{
			Lat: lat,
			Lng: lng,
		}
	}
	var requestNetworksUpdateNetworkFloorPlanBottomRightCorner *merakigosdk.RequestNetworksUpdateNetworkFloorPlanBottomRightCorner
	if r.BottomRightCorner != nil {
		lat := r.BottomRightCorner.Lat.ValueFloat64Pointer()
		lng := r.BottomRightCorner.Lng.ValueFloat64Pointer()
		requestNetworksUpdateNetworkFloorPlanBottomRightCorner = &merakigosdk.RequestNetworksUpdateNetworkFloorPlanBottomRightCorner{
			Lat: lat,
			Lng: lng,
		}
	}
	var requestNetworksUpdateNetworkFloorPlanCenter *merakigosdk.RequestNetworksUpdateNetworkFloorPlanCenter
	if r.Center != nil {
		lat := r.Center.Lat.ValueFloat64Pointer()
		lng := r.Center.Lng.ValueFloat64Pointer()
		requestNetworksUpdateNetworkFloorPlanCenter = &merakigosdk.RequestNetworksUpdateNetworkFloorPlanCenter{
			Lat: lat,
			Lng: lng,
		}
	}
	imageContents := new(string)
	if !r.ImageContents.IsUnknown() && !r.ImageContents.IsNull() {
		*imageContents = r.ImageContents.ValueString()
	} else {
		imageContents = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestNetworksUpdateNetworkFloorPlanTopLeftCorner *merakigosdk.RequestNetworksUpdateNetworkFloorPlanTopLeftCorner
	if r.TopLeftCorner != nil {
		lat := r.TopLeftCorner.Lat.ValueFloat64Pointer()
		lng := r.TopLeftCorner.Lng.ValueFloat64Pointer()
		requestNetworksUpdateNetworkFloorPlanTopLeftCorner = &merakigosdk.RequestNetworksUpdateNetworkFloorPlanTopLeftCorner{
			Lat: lat,
			Lng: lng,
		}
	}
	var requestNetworksUpdateNetworkFloorPlanTopRightCorner *merakigosdk.RequestNetworksUpdateNetworkFloorPlanTopRightCorner
	if r.TopRightCorner != nil {
		lat := r.TopRightCorner.Lat.ValueFloat64Pointer()
		lng := r.TopRightCorner.Lng.ValueFloat64Pointer()
		requestNetworksUpdateNetworkFloorPlanTopRightCorner = &merakigosdk.RequestNetworksUpdateNetworkFloorPlanTopRightCorner{
			Lat: lat,
			Lng: lng,
		}
	}
	out := merakigosdk.RequestNetworksUpdateNetworkFloorPlan{
		BottomLeftCorner:  requestNetworksUpdateNetworkFloorPlanBottomLeftCorner,
		BottomRightCorner: requestNetworksUpdateNetworkFloorPlanBottomRightCorner,
		Center:            requestNetworksUpdateNetworkFloorPlanCenter,
		ImageContents:     *imageContents,
		Name:              *name,
		TopLeftCorner:     requestNetworksUpdateNetworkFloorPlanTopLeftCorner,
		TopRightCorner:    requestNetworksUpdateNetworkFloorPlanTopRightCorner,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkFloorPlanItemToBodyRs(state NetworksFloorPlansRs, response *merakigosdk.ResponseNetworksGetNetworkFloorPlan, is_read bool) NetworksFloorPlansRs {
	itemState := NetworksFloorPlansRs{
		BottomLeftCorner: func() *ResponseNetworksGetNetworkFloorPlanBottomLeftCornerRs {
			if response.BottomLeftCorner != nil {
				return &ResponseNetworksGetNetworkFloorPlanBottomLeftCornerRs{
					Lat: func() types.Float64 {
						if response.BottomLeftCorner.Lat != nil {
							return types.Float64Value(float64(*response.BottomLeftCorner.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.BottomLeftCorner.Lng != nil {
							return types.Float64Value(float64(*response.BottomLeftCorner.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanBottomLeftCornerRs{}
		}(),
		BottomRightCorner: func() *ResponseNetworksGetNetworkFloorPlanBottomRightCornerRs {
			if response.BottomRightCorner != nil {
				return &ResponseNetworksGetNetworkFloorPlanBottomRightCornerRs{
					Lat: func() types.Float64 {
						if response.BottomRightCorner.Lat != nil {
							return types.Float64Value(float64(*response.BottomRightCorner.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.BottomRightCorner.Lng != nil {
							return types.Float64Value(float64(*response.BottomRightCorner.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanBottomRightCornerRs{}
		}(),
		Center: func() *ResponseNetworksGetNetworkFloorPlanCenterRs {
			if response.Center != nil {
				return &ResponseNetworksGetNetworkFloorPlanCenterRs{
					Lat: func() types.Float64 {
						if response.Center.Lat != nil {
							return types.Float64Value(float64(*response.Center.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.Center.Lng != nil {
							return types.Float64Value(float64(*response.Center.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanCenterRs{}
		}(),
		Devices: func() *[]ResponseNetworksGetNetworkFloorPlanDevicesRs {
			if response.Devices != nil {
				result := make([]ResponseNetworksGetNetworkFloorPlanDevicesRs, len(*response.Devices))
				for i, devices := range *response.Devices {
					result[i] = ResponseNetworksGetNetworkFloorPlanDevicesRs{
						Address: types.StringValue(devices.Address),
						BeaconIDParams: func() *ResponseNetworksGetNetworkFloorPlanDevicesBeaconIdParamsRs {
							if devices.BeaconIDParams != nil {
								return &ResponseNetworksGetNetworkFloorPlanDevicesBeaconIdParamsRs{
									Major: func() types.Int64 {
										if devices.BeaconIDParams.Major != nil {
											return types.Int64Value(int64(*devices.BeaconIDParams.Major))
										}
										return types.Int64{}
									}(),
									Minor: func() types.Int64 {
										if devices.BeaconIDParams.Minor != nil {
											return types.Int64Value(int64(*devices.BeaconIDParams.Minor))
										}
										return types.Int64{}
									}(),
									UUID: types.StringValue(devices.BeaconIDParams.UUID),
								}
							}
							return &ResponseNetworksGetNetworkFloorPlanDevicesBeaconIdParamsRs{}
						}(),
						Firmware:    types.StringValue(devices.Firmware),
						FloorPlanID: types.StringValue(devices.FloorPlanID),
						LanIP:       types.StringValue(devices.LanIP),
						Lat: func() types.Float64 {
							if devices.Lat != nil {
								return types.Float64Value(float64(*devices.Lat))
							}
							return types.Float64{}
						}(),
						Lng: func() types.Float64 {
							if devices.Lng != nil {
								return types.Float64Value(float64(*devices.Lng))
							}
							return types.Float64{}
						}(),
						Mac:       types.StringValue(devices.Mac),
						Model:     types.StringValue(devices.Model),
						Name:      types.StringValue(devices.Name),
						NetworkID: types.StringValue(devices.NetworkID),
						Notes:     types.StringValue(devices.Notes),
						Serial:    types.StringValue(devices.Serial),
						Tags:      StringSliceToSet(devices.Tags),
					}
				}
				return &result
			}
			return &[]ResponseNetworksGetNetworkFloorPlanDevicesRs{}
		}(),
		FloorPlanID: types.StringValue(response.FloorPlanID),
		Height: func() types.Float64 {
			if response.Height != nil {
				return types.Float64Value(float64(*response.Height))
			}
			return types.Float64{}
		}(),
		ImageExtension:    types.StringValue(response.ImageExtension),
		ImageMd5:          types.StringValue(response.ImageMd5),
		ImageURL:          types.StringValue(response.ImageURL),
		ImageURLExpiresAt: types.StringValue(response.ImageURLExpiresAt),
		Name:              types.StringValue(response.Name),
		TopLeftCorner: func() *ResponseNetworksGetNetworkFloorPlanTopLeftCornerRs {
			if response.TopLeftCorner != nil {
				return &ResponseNetworksGetNetworkFloorPlanTopLeftCornerRs{
					Lat: func() types.Float64 {
						if response.TopLeftCorner.Lat != nil {
							return types.Float64Value(float64(*response.TopLeftCorner.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.TopLeftCorner.Lng != nil {
							return types.Float64Value(float64(*response.TopLeftCorner.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanTopLeftCornerRs{}
		}(),
		TopRightCorner: func() *ResponseNetworksGetNetworkFloorPlanTopRightCornerRs {
			if response.TopRightCorner != nil {
				return &ResponseNetworksGetNetworkFloorPlanTopRightCornerRs{
					Lat: func() types.Float64 {
						if response.TopRightCorner.Lat != nil {
							return types.Float64Value(float64(*response.TopRightCorner.Lat))
						}
						return types.Float64{}
					}(),
					Lng: func() types.Float64 {
						if response.TopRightCorner.Lng != nil {
							return types.Float64Value(float64(*response.TopRightCorner.Lng))
						}
						return types.Float64{}
					}(),
				}
			}
			return &ResponseNetworksGetNetworkFloorPlanTopRightCornerRs{}
		}(),
		Width: func() types.Int64 {
			if response.Width != nil {
				return types.Int64Value(int64(*response.Width))
			}
			return types.Int64{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksFloorPlansRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksFloorPlansRs)
}
