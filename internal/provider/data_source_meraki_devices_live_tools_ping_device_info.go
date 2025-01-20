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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DevicesLiveToolsPingDeviceInfoDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesLiveToolsPingDeviceInfoDataSource{}
)

func NewDevicesLiveToolsPingDeviceInfoDataSource() datasource.DataSource {
	return &DevicesLiveToolsPingDeviceInfoDataSource{}
}

type DevicesLiveToolsPingDeviceInfoDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesLiveToolsPingDeviceInfoDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesLiveToolsPingDeviceInfoDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_live_tools_ping_device_info"
}

func (d *DevicesLiveToolsPingDeviceInfoDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: `id path parameter.`,
				Required:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"callback": schema.SingleNestedAttribute{
						MarkdownDescription: `Information for callback used to send back results`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `The ID of the callback. To check the status of the callback, use this ID in a request to /webhooks/callbacks/statuses/{id}`,
								Computed:            true,
							},
							"status": schema.StringAttribute{
								MarkdownDescription: `The status of the callback`,
								Computed:            true,
							},
							"url": schema.StringAttribute{
								MarkdownDescription: `The callback URL for the webhook target. This was either provided in the original request or comes from a configured webhook receiver`,
								Computed:            true,
							},
						},
					},
					"ping_id": schema.StringAttribute{
						MarkdownDescription: `Id to check the status of your ping request.`,
						Computed:            true,
					},
					"request": schema.SingleNestedAttribute{
						MarkdownDescription: `Ping request parameters`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"count": schema.Int64Attribute{
								MarkdownDescription: `Number of pings to send. [1..5], default 5`,
								Computed:            true,
							},
							"serial": schema.StringAttribute{
								MarkdownDescription: `Device serial number`,
								Computed:            true,
							},
						},
					},
					"results": schema.SingleNestedAttribute{
						MarkdownDescription: `Results of the ping request.`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"latencies": schema.SingleNestedAttribute{
								MarkdownDescription: `Packet latency stats`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"average": schema.Float64Attribute{
										MarkdownDescription: `Average latency`,
										Computed:            true,
									},
									"maximum": schema.Float64Attribute{
										MarkdownDescription: `Maximum latency`,
										Computed:            true,
									},
									"minimum": schema.Float64Attribute{
										MarkdownDescription: `Minimum latency`,
										Computed:            true,
									},
								},
							},
							"loss": schema.SingleNestedAttribute{
								MarkdownDescription: `Lost packets`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"percentage": schema.Float64Attribute{
										MarkdownDescription: `Percentage of packets lost`,
										Computed:            true,
									},
								},
							},
							"received": schema.Int64Attribute{
								MarkdownDescription: `Number of packets received`,
								Computed:            true,
							},
							"replies": schema.SetNestedAttribute{
								MarkdownDescription: `Received packets`,
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"latency": schema.Float64Attribute{
											MarkdownDescription: `Latency of the packet in milliseconds`,
											Computed:            true,
										},
										"sequence_id": schema.Int64Attribute{
											MarkdownDescription: `Sequence ID of the packet`,
											Computed:            true,
										},
										"size": schema.Int64Attribute{
											MarkdownDescription: `Size of the packet in bytes`,
											Computed:            true,
										},
									},
								},
							},
							"sent": schema.Int64Attribute{
								MarkdownDescription: `Number of packets sent`,
								Computed:            true,
							},
						},
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `Status of the ping request.`,
						Computed:            true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `GET this url to check the status of your ping request.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *DevicesLiveToolsPingDeviceInfoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesLiveToolsPingDeviceInfo DevicesLiveToolsPingDeviceInfo
	diags := req.Config.Get(ctx, &devicesLiveToolsPingDeviceInfo)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceLiveToolsPingDevice")
		vvSerial := devicesLiveToolsPingDeviceInfo.Serial.ValueString()
		vvID := devicesLiveToolsPingDeviceInfo.ID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Devices.GetDeviceLiveToolsPingDevice(vvSerial, vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceLiveToolsPingDevice",
				err.Error(),
			)
			return
		}

		devicesLiveToolsPingDeviceInfo = ResponseDevicesGetDeviceLiveToolsPingDeviceItemToBody(devicesLiveToolsPingDeviceInfo, response1)
		diags = resp.State.Set(ctx, &devicesLiveToolsPingDeviceInfo)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesLiveToolsPingDeviceInfo struct {
	Serial types.String                                 `tfsdk:"serial"`
	ID     types.String                                 `tfsdk:"id"`
	Item   *ResponseDevicesGetDeviceLiveToolsPingDevice `tfsdk:"item"`
}

type ResponseDevicesGetDeviceLiveToolsPingDevice struct {
	Callback *ResponseDevicesGetDeviceLiveToolsPingDeviceCallback `tfsdk:"callback"`
	PingID   types.String                                         `tfsdk:"ping_id"`
	Request  *ResponseDevicesGetDeviceLiveToolsPingDeviceRequest  `tfsdk:"request"`
	Results  *ResponseDevicesGetDeviceLiveToolsPingDeviceResults  `tfsdk:"results"`
	Status   types.String                                         `tfsdk:"status"`
	URL      types.String                                         `tfsdk:"url"`
}

type ResponseDevicesGetDeviceLiveToolsPingDeviceCallback struct {
	ID     types.String `tfsdk:"id"`
	Status types.String `tfsdk:"status"`
	URL    types.String `tfsdk:"url"`
}

type ResponseDevicesGetDeviceLiveToolsPingDeviceRequest struct {
	Count  types.Int64  `tfsdk:"count"`
	Serial types.String `tfsdk:"serial"`
}

type ResponseDevicesGetDeviceLiveToolsPingDeviceResults struct {
	Latencies *ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLatencies `tfsdk:"latencies"`
	Loss      *ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLoss      `tfsdk:"loss"`
	Received  types.Int64                                                  `tfsdk:"received"`
	Replies   *[]ResponseDevicesGetDeviceLiveToolsPingDeviceResultsReplies `tfsdk:"replies"`
	Sent      types.Int64                                                  `tfsdk:"sent"`
}

type ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLatencies struct {
	Average types.Float64 `tfsdk:"average"`
	Maximum types.Float64 `tfsdk:"maximum"`
	Minimum types.Float64 `tfsdk:"minimum"`
}

type ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLoss struct {
	Percentage types.Float64 `tfsdk:"percentage"`
}

type ResponseDevicesGetDeviceLiveToolsPingDeviceResultsReplies struct {
	Latency    types.Float64 `tfsdk:"latency"`
	SequenceID types.Int64   `tfsdk:"sequence_id"`
	Size       types.Int64   `tfsdk:"size"`
}

// ToBody
func ResponseDevicesGetDeviceLiveToolsPingDeviceItemToBody(state DevicesLiveToolsPingDeviceInfo, response *merakigosdk.ResponseDevicesGetDeviceLiveToolsPingDevice) DevicesLiveToolsPingDeviceInfo {
	itemState := ResponseDevicesGetDeviceLiveToolsPingDevice{
		Callback: func() *ResponseDevicesGetDeviceLiveToolsPingDeviceCallback {
			if response.Callback != nil {
				return &ResponseDevicesGetDeviceLiveToolsPingDeviceCallback{
					ID:     types.StringValue(response.Callback.ID),
					Status: types.StringValue(response.Callback.Status),
					URL:    types.StringValue(response.Callback.URL),
				}
			}
			return nil
		}(),
		PingID: types.StringValue(response.PingID),
		Request: func() *ResponseDevicesGetDeviceLiveToolsPingDeviceRequest {
			if response.Request != nil {
				return &ResponseDevicesGetDeviceLiveToolsPingDeviceRequest{
					Count: func() types.Int64 {
						if response.Request.Count != nil {
							return types.Int64Value(int64(*response.Request.Count))
						}
						return types.Int64{}
					}(),
					Serial: types.StringValue(response.Request.Serial),
				}
			}
			return nil
		}(),
		Results: func() *ResponseDevicesGetDeviceLiveToolsPingDeviceResults {
			if response.Results != nil {
				return &ResponseDevicesGetDeviceLiveToolsPingDeviceResults{
					Latencies: func() *ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLatencies {
						if response.Results.Latencies != nil {
							return &ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLatencies{
								Average: func() types.Float64 {
									if response.Results.Latencies.Average != nil {
										return types.Float64Value(float64(*response.Results.Latencies.Average))
									}
									return types.Float64{}
								}(),
								Maximum: func() types.Float64 {
									if response.Results.Latencies.Maximum != nil {
										return types.Float64Value(float64(*response.Results.Latencies.Maximum))
									}
									return types.Float64{}
								}(),
								Minimum: func() types.Float64 {
									if response.Results.Latencies.Minimum != nil {
										return types.Float64Value(float64(*response.Results.Latencies.Minimum))
									}
									return types.Float64{}
								}(),
							}
						}
						return nil
					}(),
					Loss: func() *ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLoss {
						if response.Results.Loss != nil {
							return &ResponseDevicesGetDeviceLiveToolsPingDeviceResultsLoss{
								Percentage: func() types.Float64 {
									if response.Results.Loss.Percentage != nil {
										return types.Float64Value(float64(*response.Results.Loss.Percentage))
									}
									return types.Float64{}
								}(),
							}
						}
						return nil
					}(),
					Received: func() types.Int64 {
						if response.Results.Received != nil {
							return types.Int64Value(int64(*response.Results.Received))
						}
						return types.Int64{}
					}(),
					Replies: func() *[]ResponseDevicesGetDeviceLiveToolsPingDeviceResultsReplies {
						if response.Results.Replies != nil {
							result := make([]ResponseDevicesGetDeviceLiveToolsPingDeviceResultsReplies, len(*response.Results.Replies))
							for i, replies := range *response.Results.Replies {
								result[i] = ResponseDevicesGetDeviceLiveToolsPingDeviceResultsReplies{
									Latency: func() types.Float64 {
										if replies.Latency != nil {
											return types.Float64Value(float64(*replies.Latency))
										}
										return types.Float64{}
									}(),
									SequenceID: func() types.Int64 {
										if replies.SequenceID != nil {
											return types.Int64Value(int64(*replies.SequenceID))
										}
										return types.Int64{}
									}(),
									Size: func() types.Int64 {
										if replies.Size != nil {
											return types.Int64Value(int64(*replies.Size))
										}
										return types.Int64{}
									}(),
								}
							}
							return &result
						}
						return nil
					}(),
					Sent: func() types.Int64 {
						if response.Results.Sent != nil {
							return types.Int64Value(int64(*response.Results.Sent))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
		Status: types.StringValue(response.Status),
		URL:    types.StringValue(response.URL),
	}
	state.Item = &itemState
	return state
}
