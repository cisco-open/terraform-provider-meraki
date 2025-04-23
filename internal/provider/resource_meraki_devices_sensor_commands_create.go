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
	_ resource.Resource              = &DevicesSensorCommandsCreateResource{}
	_ resource.ResourceWithConfigure = &DevicesSensorCommandsCreateResource{}
)

func NewDevicesSensorCommandsCreateResource() resource.Resource {
	return &DevicesSensorCommandsCreateResource{}
}

type DevicesSensorCommandsCreateResource struct {
	client *merakigosdk.Client
}

func (r *DevicesSensorCommandsCreateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesSensorCommandsCreateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_sensor_commands_create"
}

// resourceAction
func (r *DevicesSensorCommandsCreateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"command_id": schema.StringAttribute{
						MarkdownDescription: `ID to check the status of the command request`,
						Computed:            true,
					},
					"completed_at": schema.StringAttribute{
						MarkdownDescription: `Time when the command was completed`,
						Computed:            true,
					},
					"created_at": schema.StringAttribute{
						MarkdownDescription: `Time when the command was triggered`,
						Computed:            true,
					},
					"created_by": schema.SingleNestedAttribute{
						MarkdownDescription: `Information about the admin who triggered the command`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"admin_id": schema.StringAttribute{
								MarkdownDescription: `ID of the admin`,
								Computed:            true,
							},
							"email": schema.StringAttribute{
								MarkdownDescription: `Email of the admin`,
								Computed:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of the admin`,
								Computed:            true,
							},
						},
					},
					"errors": schema.ListAttribute{
						MarkdownDescription: `Array of errors if failed`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"operation": schema.StringAttribute{
						MarkdownDescription: `Operation run on the sensor
                                          Allowed values: [cycleDownstreamPower,disableDownstreamPower,enableDownstreamPower,refreshData]`,
						Computed: true,
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `Status of the command request
                                          Allowed values: [completed,failed,in_progress,pending]`,
						Computed: true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"operation": schema.StringAttribute{
						MarkdownDescription: `Operation to run on the sensor. 'enableDownstreamPower', 'disableDownstreamPower', and 'cycleDownstreamPower' turn power on/off to the device that is connected downstream of an MT40 power monitor. 'refreshData' causes an MT15 or MT40 device to upload its latest readings so that they are immediately available in the Dashboard API.
                                        Allowed values: [cycleDownstreamPower,disableDownstreamPower,enableDownstreamPower,refreshData]`,
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *DevicesSensorCommandsCreateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesSensorCommandsCreate

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Sensor.CreateDeviceSensorCommand(vvSerial, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateDeviceSensorCommand",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateDeviceSensorCommand",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseSensorCreateDeviceSensorCommandItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesSensorCommandsCreateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesSensorCommandsCreateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesSensorCommandsCreateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesSensorCommandsCreate struct {
	Serial     types.String                              `tfsdk:"serial"`
	Item       *ResponseSensorCreateDeviceSensorCommand  `tfsdk:"item"`
	Parameters *RequestSensorCreateDeviceSensorCommandRs `tfsdk:"parameters"`
}

type ResponseSensorCreateDeviceSensorCommand struct {
	CommandID   types.String                                      `tfsdk:"command_id"`
	CompletedAt types.String                                      `tfsdk:"completed_at"`
	CreatedAt   types.String                                      `tfsdk:"created_at"`
	CreatedBy   *ResponseSensorCreateDeviceSensorCommandCreatedBy `tfsdk:"created_by"`
	Errors      types.List                                        `tfsdk:"errors"`
	Operation   types.String                                      `tfsdk:"operation"`
	Status      types.String                                      `tfsdk:"status"`
}

type ResponseSensorCreateDeviceSensorCommandCreatedBy struct {
	AdminID types.String `tfsdk:"admin_id"`
	Email   types.String `tfsdk:"email"`
	Name    types.String `tfsdk:"name"`
}

type RequestSensorCreateDeviceSensorCommandRs struct {
	Operation types.String `tfsdk:"operation"`
}

// FromBody
func (r *DevicesSensorCommandsCreate) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSensorCreateDeviceSensorCommand {
	emptyString := ""
	re := *r.Parameters
	operation := new(string)
	if !re.Operation.IsUnknown() && !re.Operation.IsNull() {
		*operation = re.Operation.ValueString()
	} else {
		operation = &emptyString
	}
	out := merakigosdk.RequestSensorCreateDeviceSensorCommand{
		Operation: *operation,
	}
	return &out
}

// ToBody
func ResponseSensorCreateDeviceSensorCommandItemToBody(state DevicesSensorCommandsCreate, response *merakigosdk.ResponseSensorCreateDeviceSensorCommand) DevicesSensorCommandsCreate {
	itemState := ResponseSensorCreateDeviceSensorCommand{
		CommandID:   types.StringValue(response.CommandID),
		CompletedAt: types.StringValue(response.CompletedAt),
		CreatedAt:   types.StringValue(response.CreatedAt),
		CreatedBy: func() *ResponseSensorCreateDeviceSensorCommandCreatedBy {
			if response.CreatedBy != nil {
				return &ResponseSensorCreateDeviceSensorCommandCreatedBy{
					AdminID: types.StringValue(response.CreatedBy.AdminID),
					Email:   types.StringValue(response.CreatedBy.Email),
					Name:    types.StringValue(response.CreatedBy.Name),
				}
			}
			return nil
		}(),
		Errors:    StringSliceToList(response.Errors),
		Operation: types.StringValue(response.Operation),
		Status:    types.StringValue(response.Status),
	}
	state.Item = &itemState
	return state
}
