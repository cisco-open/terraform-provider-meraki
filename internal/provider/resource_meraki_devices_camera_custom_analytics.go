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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesCameraCustomAnalyticsResource{}
	_ resource.ResourceWithConfigure = &DevicesCameraCustomAnalyticsResource{}
)

func NewDevicesCameraCustomAnalyticsResource() resource.Resource {
	return &DevicesCameraCustomAnalyticsResource{}
}

type DevicesCameraCustomAnalyticsResource struct {
	client *merakigosdk.Client
}

func (r *DevicesCameraCustomAnalyticsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesCameraCustomAnalyticsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_custom_analytics"
}

func (r *DevicesCameraCustomAnalyticsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"artifact_id": schema.StringAttribute{
				MarkdownDescription: `Custom analytics artifact ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether custom analytics is enabled`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"parameters": schema.SetNestedAttribute{
				MarkdownDescription: `Parameters for the custom analytics workload`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the parameter`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"value": schema.Float64Attribute{
							MarkdownDescription: `Value of the parameter`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Float64{
								float64planmodifier.UseStateForUnknown(),
							},
							//                  Differents_types: `   parameter: schema.TypeString, item: schema.TypeFloat`,
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

func (r *DevicesCameraCustomAnalyticsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesCameraCustomAnalyticsRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Camera.GetDeviceCameraCustomAnalytics(vvSerial)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesCameraCustomAnalytics only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesCameraCustomAnalytics only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Camera.UpdateDeviceCameraCustomAnalytics(vvSerial, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCameraCustomAnalytics",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCameraCustomAnalytics",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Camera.GetDeviceCameraCustomAnalytics(vvSerial)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraCustomAnalytics",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCameraCustomAnalytics",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCameraGetDeviceCameraCustomAnalyticsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCameraCustomAnalyticsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesCameraCustomAnalyticsRs

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
	responseGet, restyRespGet, err := r.client.Camera.GetDeviceCameraCustomAnalytics(vvSerial)
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
				"Failure when executing GetDeviceCameraCustomAnalytics",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCameraCustomAnalytics",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCameraGetDeviceCameraCustomAnalyticsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesCameraCustomAnalyticsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesCameraCustomAnalyticsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesCameraCustomAnalyticsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Camera.UpdateDeviceCameraCustomAnalytics(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCameraCustomAnalytics",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCameraCustomAnalytics",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCameraCustomAnalyticsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesCameraCustomAnalytics", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesCameraCustomAnalyticsRs struct {
	Serial     types.String                                                `tfsdk:"serial"`
	ArtifactID types.String                                                `tfsdk:"artifact_id"`
	Enabled    types.Bool                                                  `tfsdk:"enabled"`
	Parameters *[]ResponseCameraGetDeviceCameraCustomAnalyticsParametersRs `tfsdk:"parameters"`
}

type ResponseCameraGetDeviceCameraCustomAnalyticsParametersRs struct {
	Name  types.String  `tfsdk:"name"`
	Value types.Float64 `tfsdk:"value"`
}

// FromBody
func (r *DevicesCameraCustomAnalyticsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCameraUpdateDeviceCameraCustomAnalytics {
	emptyString := ""
	artifactID := new(string)
	if !r.ArtifactID.IsUnknown() && !r.ArtifactID.IsNull() {
		*artifactID = r.ArtifactID.ValueString()
	} else {
		artifactID = &emptyString
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var requestCameraUpdateDeviceCameraCustomAnalyticsParameters []merakigosdk.RequestCameraUpdateDeviceCameraCustomAnalyticsParameters
	if r.Parameters != nil {
		for _, rItem1 := range *r.Parameters {
			name := rItem1.Name.ValueString()
			value := ""
			if rItem1.Value.ValueFloat64Pointer() != nil {
				value = fmt.Sprintf("%f", *rItem1.Value.ValueFloat64Pointer())
			}
			requestCameraUpdateDeviceCameraCustomAnalyticsParameters = append(requestCameraUpdateDeviceCameraCustomAnalyticsParameters, merakigosdk.RequestCameraUpdateDeviceCameraCustomAnalyticsParameters{
				Name:  name,
				Value: value,
			})
		}
	}
	out := merakigosdk.RequestCameraUpdateDeviceCameraCustomAnalytics{
		ArtifactID: *artifactID,
		Enabled:    enabled,
		Parameters: func() *[]merakigosdk.RequestCameraUpdateDeviceCameraCustomAnalyticsParameters {
			if len(requestCameraUpdateDeviceCameraCustomAnalyticsParameters) > 0 {
				return &requestCameraUpdateDeviceCameraCustomAnalyticsParameters
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCameraGetDeviceCameraCustomAnalyticsItemToBodyRs(state DevicesCameraCustomAnalyticsRs, response *merakigosdk.ResponseCameraGetDeviceCameraCustomAnalytics, is_read bool) DevicesCameraCustomAnalyticsRs {
	itemState := DevicesCameraCustomAnalyticsRs{
		ArtifactID: types.StringValue(response.ArtifactID),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Parameters: func() *[]ResponseCameraGetDeviceCameraCustomAnalyticsParametersRs {
			if response.Parameters != nil {
				result := make([]ResponseCameraGetDeviceCameraCustomAnalyticsParametersRs, len(*response.Parameters))
				for i, parameters := range *response.Parameters {
					result[i] = ResponseCameraGetDeviceCameraCustomAnalyticsParametersRs{
						Name: types.StringValue(parameters.Name),
						Value: func() types.Float64 {
							if parameters.Value != nil {
								return types.Float64Value(float64(*parameters.Value))
							}
							return types.Float64{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesCameraCustomAnalyticsRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesCameraCustomAnalyticsRs)
}
