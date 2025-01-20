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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesSensorRelationshipsResource{}
	_ resource.ResourceWithConfigure = &DevicesSensorRelationshipsResource{}
)

func NewDevicesSensorRelationshipsResource() resource.Resource {
	return &DevicesSensorRelationshipsResource{}
}

type DevicesSensorRelationshipsResource struct {
	client *merakigosdk.Client
}

func (r *DevicesSensorRelationshipsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesSensorRelationshipsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_sensor_relationships"
}

func (r *DevicesSensorRelationshipsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"livestream": schema.SingleNestedAttribute{
				MarkdownDescription: `A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"related_devices": schema.SetNestedAttribute{
						MarkdownDescription: `An array of the related devices for the role`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"product_type": schema.StringAttribute{
									MarkdownDescription: `The product type of the related device
                                              Allowed values: [camera,sensor]`,
									Computed: true,
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `The serial of the related device`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
				},
			},
			"livestream_request": schema.SetNestedAttribute{
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: `A role defined between an MT sensor and an MV camera that adds the camera's r.Livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.`,
				Computed:            true,
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"related_devices": schema.SetNestedAttribute{
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							MarkdownDescription: `An array of the related devices for the role`,
							Computed:            true,
							Optional:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"product_type": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: `The product type of the related device`,
										Computed:            true,
									},
									"serial": schema.StringAttribute{
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										MarkdownDescription: `The serial of the related device`,
										Computed:            true,
										Optional:            true,
									},
								},
							},
						},
					},
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
		},
	}
}

func (r *DevicesSensorRelationshipsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesSensorRelationshipsRs

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
	vvSerial := data.Serial.ValueString()
	// serial
	//Reviw This  Has Item Not response
	//Esta bien

	//Items
	responseVerifyItem, restyResp1, err := r.client.Sensor.GetDeviceSensorRelationships(vvSerial)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesSensorRelationships only have update context, not create.",
			err.Error(),
		)
		return
	}
	//TODO HAS ONLY ITEMS UPDATE

	response, restyResp2, err := r.client.Sensor.UpdateDeviceSensorRelationships(vvSerial, data.toSdkApiRequestUpdate(ctx))

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceSensorRelationships",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceSensorRelationships",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Sensor.GetDeviceSensorRelationships(vvSerial)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSensorRelationships",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceSensorRelationships",
			err.Error(),
		)
		return
	}
	data = ResponseSensorGetDeviceSensorRelationshipsItemsToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesSensorRelationshipsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesSensorRelationshipsRs

	var response types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &response)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(response.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	// Has Item2

	vvSerial := data.Serial.ValueString()
	// serial
	responseGet, restyResp1, err := r.client.Sensor.GetDeviceSensorRelationships(vvSerial)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSensorRelationships",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceSensorRelationships",
			err.Error(),
		)
		return
	}
	//entro aqui
	data = ResponseSensorGetDeviceSensorRelationshipsItemsToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesSensorRelationshipsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesSensorRelationshipsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesSensorRelationshipsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Sensor.UpdateDeviceSensorRelationships(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceSensorRelationships",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceSensorRelationships",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesSensorRelationshipsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesSensorRelationships", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesSensorRelationshipsRs struct {
	Serial types.String `tfsdk:"serial"`
	//TIENE ITEMS
	Livestream   *ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRs `tfsdk:"livestream"`
	LivestreamRs *ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRs `tfsdk:"livestream_request"`
}

type ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRs struct {
	RelatedDevices *[]ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevicesRs `tfsdk:"related_devices"`
}

type ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevicesRs struct {
	ProductType types.String `tfsdk:"product_type"`
	Serial      types.String `tfsdk:"serial"`
}

// FromBody
func (r *DevicesSensorRelationshipsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSensorUpdateDeviceSensorRelationships {
	var requestSensorUpdateDeviceSensorRelationshipsLivestream *merakigosdk.RequestSensorUpdateDeviceSensorRelationshipsLivestream
	if r.LivestreamRs != nil {
		var requestSensorUpdateDeviceSensorRelationshipsLivestreamRelatedDevices []merakigosdk.RequestSensorUpdateDeviceSensorRelationshipsLivestreamRelatedDevices

		for _, rItem1 := range *r.LivestreamRs.RelatedDevices {
			serial := rItem1.Serial.ValueString()
			requestSensorUpdateDeviceSensorRelationshipsLivestreamRelatedDevices = append(requestSensorUpdateDeviceSensorRelationshipsLivestreamRelatedDevices, merakigosdk.RequestSensorUpdateDeviceSensorRelationshipsLivestreamRelatedDevices{
				Serial: serial,
			})
		}

		requestSensorUpdateDeviceSensorRelationshipsLivestream = &merakigosdk.RequestSensorUpdateDeviceSensorRelationshipsLivestream{
			RelatedDevices: func() *[]merakigosdk.RequestSensorUpdateDeviceSensorRelationshipsLivestreamRelatedDevices {
				if len(requestSensorUpdateDeviceSensorRelationshipsLivestreamRelatedDevices) > 0 {
					return &requestSensorUpdateDeviceSensorRelationshipsLivestreamRelatedDevices
				}
				return nil
			}(),
		}
	}
	out := merakigosdk.RequestSensorUpdateDeviceSensorRelationships{
		Livestream: requestSensorUpdateDeviceSensorRelationshipsLivestream,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSensorGetDeviceSensorRelationshipsItemsToBodyRs(state DevicesSensorRelationshipsRs, response *merakigosdk.ResponseSensorGetDeviceSensorRelationships, is_read bool) DevicesSensorRelationshipsRs {

	var newItem ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRs

	newItem = func() ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRs {
		if response.Livestream != nil {
			return ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRs{
				RelatedDevices: func() *[]ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevicesRs {
					if response.Livestream.RelatedDevices != nil {
						result := make([]ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevicesRs, len(*response.Livestream.RelatedDevices))
						for i, relatedDevices := range *response.Livestream.RelatedDevices {
							result[i] = ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRelatedDevicesRs{
								ProductType: types.StringValue(relatedDevices.ProductType),
								Serial:      types.StringValue(relatedDevices.Serial),
							}
						}
						return &result
					}
					return nil
				}(),
			}
		}
		return ResponseItemSensorGetDeviceSensorRelationshipsLivestreamRs{}
	}()
	state.Livestream = &newItem
	return state
}
