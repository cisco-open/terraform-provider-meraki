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

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DevicesSensorCommandsDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesSensorCommandsDataSource{}
)

func NewDevicesSensorCommandsDataSource() datasource.DataSource {
	return &DevicesSensorCommandsDataSource{}
}

type DevicesSensorCommandsDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesSensorCommandsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesSensorCommandsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_sensor_commands"
}

func (d *DevicesSensorCommandsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"command_id": schema.StringAttribute{
				MarkdownDescription: `commandId path parameter. Command ID`,
				Optional:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"operations": schema.ListAttribute{
				MarkdownDescription: `operations query parameter. Optional parameter to filter commands by operation. Allowed values are disableDownstreamPower, enableDownstreamPower, cycleDownstreamPower, and refreshData.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 10.`,
				Optional:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Optional:            true,
			},
			"sort_order": schema.StringAttribute{
				MarkdownDescription: `sortOrder query parameter. Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'descending'.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 30 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 30 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 30 days. The default is 30 days.`,
				Optional:            true,
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
						MarkdownDescription: `Operation run on the sensor`,
						Computed:            true,
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `Status of the command request`,
						Computed:            true,
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseSensorGetDeviceSensorCommands`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
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
							MarkdownDescription: `Operation run on the sensor`,
							Computed:            true,
						},
						"status": schema.StringAttribute{
							MarkdownDescription: `Status of the command request`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *DevicesSensorCommandsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesSensorCommands DevicesSensorCommands
	diags := req.Config.Get(ctx, &devicesSensorCommands)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!devicesSensorCommands.Serial.IsNull(), !devicesSensorCommands.Operations.IsNull(), !devicesSensorCommands.PerPage.IsNull(), !devicesSensorCommands.StartingAfter.IsNull(), !devicesSensorCommands.EndingBefore.IsNull(), !devicesSensorCommands.SortOrder.IsNull(), !devicesSensorCommands.T0.IsNull(), !devicesSensorCommands.T1.IsNull(), !devicesSensorCommands.Timespan.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!devicesSensorCommands.Serial.IsNull(), !devicesSensorCommands.CommandID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceSensorCommands")
		vvSerial := devicesSensorCommands.Serial.ValueString()
		queryParams1 := merakigosdk.GetDeviceSensorCommandsQueryParams{}

		queryParams1.Operations = elementsToStrings(ctx, devicesSensorCommands.Operations)
		queryParams1.PerPage = int(devicesSensorCommands.PerPage.ValueInt64())
		queryParams1.StartingAfter = devicesSensorCommands.StartingAfter.ValueString()
		queryParams1.EndingBefore = devicesSensorCommands.EndingBefore.ValueString()
		queryParams1.SortOrder = devicesSensorCommands.SortOrder.ValueString()
		queryParams1.T0 = devicesSensorCommands.T0.ValueString()
		queryParams1.T1 = devicesSensorCommands.T1.ValueString()
		queryParams1.Timespan = devicesSensorCommands.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Sensor.GetDeviceSensorCommands(vvSerial, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSensorCommands",
				err.Error(),
			)
			return
		}

		devicesSensorCommands = ResponseSensorGetDeviceSensorCommandsItemsToBody(devicesSensorCommands, response1)
		diags = resp.State.Set(ctx, &devicesSensorCommands)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDeviceSensorCommand")
		vvSerial := devicesSensorCommands.Serial.ValueString()
		vvCommandID := devicesSensorCommands.CommandID.ValueString()

		// has_unknown_response: None

		response2, restyResp2, err := d.client.Sensor.GetDeviceSensorCommand(vvSerial, vvCommandID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceSensorCommand",
				err.Error(),
			)
			return
		}

		devicesSensorCommands = ResponseSensorGetDeviceSensorCommandItemToBody(devicesSensorCommands, response2)
		diags = resp.State.Set(ctx, &devicesSensorCommands)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesSensorCommands struct {
	Serial        types.String                                 `tfsdk:"serial"`
	Operations    types.List                                   `tfsdk:"operations"`
	PerPage       types.Int64                                  `tfsdk:"per_page"`
	StartingAfter types.String                                 `tfsdk:"starting_after"`
	EndingBefore  types.String                                 `tfsdk:"ending_before"`
	SortOrder     types.String                                 `tfsdk:"sort_order"`
	T0            types.String                                 `tfsdk:"t0"`
	T1            types.String                                 `tfsdk:"t1"`
	Timespan      types.Float64                                `tfsdk:"timespan"`
	CommandID     types.String                                 `tfsdk:"command_id"`
	Items         *[]ResponseItemSensorGetDeviceSensorCommands `tfsdk:"items"`
	Item          *ResponseSensorGetDeviceSensorCommand        `tfsdk:"item"`
}

type ResponseItemSensorGetDeviceSensorCommands struct {
	CommandID   types.String                                        `tfsdk:"command_id"`
	CompletedAt types.String                                        `tfsdk:"completed_at"`
	CreatedAt   types.String                                        `tfsdk:"created_at"`
	CreatedBy   *ResponseItemSensorGetDeviceSensorCommandsCreatedBy `tfsdk:"created_by"`
	Errors      types.List                                          `tfsdk:"errors"`
	Operation   types.String                                        `tfsdk:"operation"`
	Status      types.String                                        `tfsdk:"status"`
}

type ResponseItemSensorGetDeviceSensorCommandsCreatedBy struct {
	AdminID types.String `tfsdk:"admin_id"`
	Email   types.String `tfsdk:"email"`
	Name    types.String `tfsdk:"name"`
}

type ResponseSensorGetDeviceSensorCommand struct {
	CommandID   types.String                                   `tfsdk:"command_id"`
	CompletedAt types.String                                   `tfsdk:"completed_at"`
	CreatedAt   types.String                                   `tfsdk:"created_at"`
	CreatedBy   *ResponseSensorGetDeviceSensorCommandCreatedBy `tfsdk:"created_by"`
	Errors      types.List                                     `tfsdk:"errors"`
	Operation   types.String                                   `tfsdk:"operation"`
	Status      types.String                                   `tfsdk:"status"`
}

type ResponseSensorGetDeviceSensorCommandCreatedBy struct {
	AdminID types.String `tfsdk:"admin_id"`
	Email   types.String `tfsdk:"email"`
	Name    types.String `tfsdk:"name"`
}

// ToBody
func ResponseSensorGetDeviceSensorCommandsItemsToBody(state DevicesSensorCommands, response *merakigosdk.ResponseSensorGetDeviceSensorCommands) DevicesSensorCommands {
	var items []ResponseItemSensorGetDeviceSensorCommands
	for _, item := range *response {
		itemState := ResponseItemSensorGetDeviceSensorCommands{
			CommandID:   types.StringValue(item.CommandID),
			CompletedAt: types.StringValue(item.CompletedAt),
			CreatedAt:   types.StringValue(item.CreatedAt),
			CreatedBy: func() *ResponseItemSensorGetDeviceSensorCommandsCreatedBy {
				if item.CreatedBy != nil {
					return &ResponseItemSensorGetDeviceSensorCommandsCreatedBy{
						AdminID: types.StringValue(item.CreatedBy.AdminID),
						Email:   types.StringValue(item.CreatedBy.Email),
						Name:    types.StringValue(item.CreatedBy.Name),
					}
				}
				return nil
			}(),
			Errors:    StringSliceToList(item.Errors),
			Operation: types.StringValue(item.Operation),
			Status:    types.StringValue(item.Status),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseSensorGetDeviceSensorCommandItemToBody(state DevicesSensorCommands, response *merakigosdk.ResponseSensorGetDeviceSensorCommand) DevicesSensorCommands {
	itemState := ResponseSensorGetDeviceSensorCommand{
		CommandID:   types.StringValue(response.CommandID),
		CompletedAt: types.StringValue(response.CompletedAt),
		CreatedAt:   types.StringValue(response.CreatedAt),
		CreatedBy: func() *ResponseSensorGetDeviceSensorCommandCreatedBy {
			if response.CreatedBy != nil {
				return &ResponseSensorGetDeviceSensorCommandCreatedBy{
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
