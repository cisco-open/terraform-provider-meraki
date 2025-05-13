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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesCameraSenseResource{}
	_ resource.ResourceWithConfigure = &DevicesCameraSenseResource{}
)

func NewDevicesCameraSenseResource() resource.Resource {
	return &DevicesCameraSenseResource{}
}

type DevicesCameraSenseResource struct {
	client *merakigosdk.Client
}

func (r *DevicesCameraSenseResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesCameraSenseResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_sense"
}

func (r *DevicesCameraSenseResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"audio_detection": schema.SingleNestedAttribute{
				MarkdownDescription: `The details of the audio detection config.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating if audio detection is enabled(true) or disabled(false) on the camera`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"detection_model_id": schema.StringAttribute{
				MarkdownDescription: `The ID of the object detection model`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"mqtt_broker_id": schema.StringAttribute{
				MarkdownDescription: `The ID of the MQTT broker to be enabled on the camera. A value of null will disable MQTT on the camera`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"mqtt_topics": schema.SetAttribute{
				Computed:    true,
				ElementType: types.StringType,
			},
			"sense_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating if sense(license) is enabled(true) or disabled(false) on the camera`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
		},
	}
}

func (r *DevicesCameraSenseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesCameraSenseRs

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
	// Has Paths
	vvSerial := data.Serial.ValueString()
	//Has Item and not has items

	if vvSerial != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Camera.GetDeviceCameraSense(vvSerial)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource DevicesCameraSense  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource DevicesCameraSense only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Camera.UpdateDeviceCameraSense(vvSerial, dataRequest)
	//Update
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCameraSense",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCameraSense",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Camera.GetDeviceCameraSense(vvSerial)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraSense",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCameraSense",
			err.Error(),
		)
		return
	}

	data = ResponseCameraGetDeviceCameraSenseItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *DevicesCameraSenseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesCameraSenseRs

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

	vvSerial := data.Serial.ValueString()
	responseGet, restyRespGet, err := r.client.Camera.GetDeviceCameraSense(vvSerial)
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
				"Failure when executing GetDeviceCameraSense",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCameraSense",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCameraGetDeviceCameraSenseItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesCameraSenseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesCameraSenseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesCameraSenseRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Camera.UpdateDeviceCameraSense(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCameraSense",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCameraSense",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCameraSenseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesCameraSense", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesCameraSenseRs struct {
	Serial           types.String                                        `tfsdk:"serial"`
	AudioDetection   *ResponseCameraGetDeviceCameraSenseAudioDetectionRs `tfsdk:"audio_detection"`
	MqttBrokerID     types.String                                        `tfsdk:"mqtt_broker_id"`
	MqttTopics       types.Set                                           `tfsdk:"mqtt_topics"`
	SenseEnabled     types.Bool                                          `tfsdk:"sense_enabled"`
	DetectionModelID types.String                                        `tfsdk:"detection_model_id"`
}

type ResponseCameraGetDeviceCameraSenseAudioDetectionRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// FromBody
func (r *DevicesCameraSenseRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCameraUpdateDeviceCameraSense {
	emptyString := ""
	var requestCameraUpdateDeviceCameraSenseAudioDetection *merakigosdk.RequestCameraUpdateDeviceCameraSenseAudioDetection

	if r.AudioDetection != nil {
		enabled := func() *bool {
			if !r.AudioDetection.Enabled.IsUnknown() && !r.AudioDetection.Enabled.IsNull() {
				return r.AudioDetection.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestCameraUpdateDeviceCameraSenseAudioDetection = &merakigosdk.RequestCameraUpdateDeviceCameraSenseAudioDetection{
			Enabled: enabled,
		}
		//[debug] Is Array: False
	}
	detectionModelID := new(string)
	if !r.DetectionModelID.IsUnknown() && !r.DetectionModelID.IsNull() {
		*detectionModelID = r.DetectionModelID.ValueString()
	} else {
		detectionModelID = &emptyString
	}
	mqttBrokerID := new(string)
	if !r.MqttBrokerID.IsUnknown() && !r.MqttBrokerID.IsNull() {
		*mqttBrokerID = r.MqttBrokerID.ValueString()
	} else {
		mqttBrokerID = &emptyString
	}
	senseEnabled := new(bool)
	if !r.SenseEnabled.IsUnknown() && !r.SenseEnabled.IsNull() {
		*senseEnabled = r.SenseEnabled.ValueBool()
	} else {
		senseEnabled = nil
	}
	out := merakigosdk.RequestCameraUpdateDeviceCameraSense{
		AudioDetection:   requestCameraUpdateDeviceCameraSenseAudioDetection,
		DetectionModelID: *detectionModelID,
		MqttBrokerID:     *mqttBrokerID,
		SenseEnabled:     senseEnabled,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCameraGetDeviceCameraSenseItemToBodyRs(state DevicesCameraSenseRs, response *merakigosdk.ResponseCameraGetDeviceCameraSense, is_read bool) DevicesCameraSenseRs {
	itemState := DevicesCameraSenseRs{
		AudioDetection: func() *ResponseCameraGetDeviceCameraSenseAudioDetectionRs {
			if response.AudioDetection != nil {
				return &ResponseCameraGetDeviceCameraSenseAudioDetectionRs{
					Enabled: func() types.Bool {
						if response.AudioDetection.Enabled != nil {
							return types.BoolValue(*response.AudioDetection.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		MqttBrokerID: types.StringValue(response.MqttBrokerID),
		MqttTopics:   StringSliceToSet(response.MqttTopics),
		SenseEnabled: func() types.Bool {
			if response.SenseEnabled != nil {
				return types.BoolValue(*response.SenseEnabled)
			}
			return types.Bool{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesCameraSenseRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesCameraSenseRs)
}
