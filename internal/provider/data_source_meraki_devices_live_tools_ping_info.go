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
	_ datasource.DataSource              = &DevicesLiveToolsPingInfoDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesLiveToolsPingInfoDataSource{}
)

func NewDevicesLiveToolsPingInfoDataSource() datasource.DataSource {
	return &DevicesLiveToolsPingInfoDataSource{}
}

type DevicesLiveToolsPingInfoDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesLiveToolsPingInfoDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesLiveToolsPingInfoDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_live_tools_ping_info"
}

func (d *DevicesLiveToolsPingInfoDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
							"target": schema.StringAttribute{
								MarkdownDescription: `IP address or FQDN to ping`,
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

func (d *DevicesLiveToolsPingInfoDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesLiveToolsPingInfo DevicesLiveToolsPingInfo
	diags := req.Config.Get(ctx, &devicesLiveToolsPingInfo)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceLiveToolsPing")
		vvSerial := devicesLiveToolsPingInfo.Serial.ValueString()
		vvID := devicesLiveToolsPingInfo.ID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Devices.GetDeviceLiveToolsPing(vvSerial, vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceLiveToolsPing",
				err.Error(),
			)
			return
		}

		devicesLiveToolsPingInfo = ResponseDevicesGetDeviceLiveToolsPingItemToBody(devicesLiveToolsPingInfo, response1)
		diags = resp.State.Set(ctx, &devicesLiveToolsPingInfo)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesLiveToolsPingInfo struct {
	Serial types.String                           `tfsdk:"serial"`
	ID     types.String                           `tfsdk:"id"`
	Item   *ResponseDevicesGetDeviceLiveToolsPing `tfsdk:"item"`
}

type ResponseDevicesGetDeviceLiveToolsPing struct {
	PingID  types.String                                  `tfsdk:"ping_id"`
	Request *ResponseDevicesGetDeviceLiveToolsPingRequest `tfsdk:"request"`
	Results *ResponseDevicesGetDeviceLiveToolsPingResults `tfsdk:"results"`
	Status  types.String                                  `tfsdk:"status"`
	URL     types.String                                  `tfsdk:"url"`
}

type ResponseDevicesGetDeviceLiveToolsPingRequest struct {
	Count  types.Int64  `tfsdk:"count"`
	Serial types.String `tfsdk:"serial"`
	Target types.String `tfsdk:"target"`
}

type ResponseDevicesGetDeviceLiveToolsPingResults struct {
	Latencies *ResponseDevicesGetDeviceLiveToolsPingResultsLatencies `tfsdk:"latencies"`
	Loss      *ResponseDevicesGetDeviceLiveToolsPingResultsLoss      `tfsdk:"loss"`
	Received  types.Int64                                            `tfsdk:"received"`
	Replies   *[]ResponseDevicesGetDeviceLiveToolsPingResultsReplies `tfsdk:"replies"`
	Sent      types.Int64                                            `tfsdk:"sent"`
}

type ResponseDevicesGetDeviceLiveToolsPingResultsLatencies struct {
	Average types.Float64 `tfsdk:"average"`
	Maximum types.Float64 `tfsdk:"maximum"`
	Minimum types.Float64 `tfsdk:"minimum"`
}

type ResponseDevicesGetDeviceLiveToolsPingResultsLoss struct {
	Percentage types.Float64 `tfsdk:"percentage"`
}

type ResponseDevicesGetDeviceLiveToolsPingResultsReplies struct {
	Latency    types.Float64 `tfsdk:"latency"`
	SequenceID types.Int64   `tfsdk:"sequence_id"`
	Size       types.Int64   `tfsdk:"size"`
}

// ToBody
func ResponseDevicesGetDeviceLiveToolsPingItemToBody(state DevicesLiveToolsPingInfo, response *merakigosdk.ResponseDevicesGetDeviceLiveToolsPing) DevicesLiveToolsPingInfo {
	itemState := ResponseDevicesGetDeviceLiveToolsPing{
		PingID: types.StringValue(response.PingID),
		Request: func() *ResponseDevicesGetDeviceLiveToolsPingRequest {
			if response.Request != nil {
				return &ResponseDevicesGetDeviceLiveToolsPingRequest{
					Count: func() types.Int64 {
						if response.Request.Count != nil {
							return types.Int64Value(int64(*response.Request.Count))
						}
						return types.Int64{}
					}(),
					Serial: types.StringValue(response.Request.Serial),
					Target: types.StringValue(response.Request.Target),
				}
			}
			return nil
		}(),
		Results: func() *ResponseDevicesGetDeviceLiveToolsPingResults {
			if response.Results != nil {
				return &ResponseDevicesGetDeviceLiveToolsPingResults{
					Latencies: func() *ResponseDevicesGetDeviceLiveToolsPingResultsLatencies {
						if response.Results.Latencies != nil {
							return &ResponseDevicesGetDeviceLiveToolsPingResultsLatencies{
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
					Loss: func() *ResponseDevicesGetDeviceLiveToolsPingResultsLoss {
						if response.Results.Loss != nil {
							return &ResponseDevicesGetDeviceLiveToolsPingResultsLoss{
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
					Replies: func() *[]ResponseDevicesGetDeviceLiveToolsPingResultsReplies {
						if response.Results.Replies != nil {
							result := make([]ResponseDevicesGetDeviceLiveToolsPingResultsReplies, len(*response.Results.Replies))
							for i, replies := range *response.Results.Replies {
								result[i] = ResponseDevicesGetDeviceLiveToolsPingResultsReplies{
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
