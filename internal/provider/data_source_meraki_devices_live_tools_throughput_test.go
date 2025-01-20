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
	_ datasource.DataSource              = &DevicesLiveToolsThroughputTestDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesLiveToolsThroughputTestDataSource{}
)

func NewDevicesLiveToolsThroughputTestDataSource() datasource.DataSource {
	return &DevicesLiveToolsThroughputTestDataSource{}
}

type DevicesLiveToolsThroughputTestDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesLiveToolsThroughputTestDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesLiveToolsThroughputTestDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_live_tools_throughput_test"
}

func (d *DevicesLiveToolsThroughputTestDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"throughput_test_id": schema.StringAttribute{
				MarkdownDescription: `throughputTestId path parameter. Throughput test ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"error": schema.StringAttribute{
						MarkdownDescription: `Description of the error.`,
						Computed:            true,
					},
					"request": schema.SingleNestedAttribute{
						MarkdownDescription: `The parameters of the throughput test request`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"serial": schema.StringAttribute{
								MarkdownDescription: `Device serial number`,
								Computed:            true,
							},
						},
					},
					"result": schema.SingleNestedAttribute{
						MarkdownDescription: `Result of the throughput test request`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"speeds": schema.SingleNestedAttribute{
								MarkdownDescription: `Shows the speeds (Mbps)`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"downstream": schema.Float64Attribute{
										MarkdownDescription: `Shows the download speed from shard (Mbps)`,
										Computed:            true,
									},
								},
							},
						},
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `Status of the throughput test request`,
						Computed:            true,
					},
					"throughput_test_id": schema.StringAttribute{
						MarkdownDescription: `ID of throughput test job`,
						Computed:            true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `GET this url to check the status of your throughput test request`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *DevicesLiveToolsThroughputTestDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesLiveToolsThroughputTest DevicesLiveToolsThroughputTest
	diags := req.Config.Get(ctx, &devicesLiveToolsThroughputTest)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceLiveToolsThroughputTest")
		vvSerial := devicesLiveToolsThroughputTest.Serial.ValueString()
		vvThroughputTestID := devicesLiveToolsThroughputTest.ThroughputTestID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Devices.GetDeviceLiveToolsThroughputTest(vvSerial, vvThroughputTestID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceLiveToolsThroughputTest",
				err.Error(),
			)
			return
		}

		devicesLiveToolsThroughputTest = ResponseDevicesGetDeviceLiveToolsThroughputTestItemToBody(devicesLiveToolsThroughputTest, response1)
		diags = resp.State.Set(ctx, &devicesLiveToolsThroughputTest)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesLiveToolsThroughputTest struct {
	Serial           types.String                                     `tfsdk:"serial"`
	ThroughputTestID types.String                                     `tfsdk:"throughput_test_id"`
	Item             *ResponseDevicesGetDeviceLiveToolsThroughputTest `tfsdk:"item"`
}

type ResponseDevicesGetDeviceLiveToolsThroughputTest struct {
	Error            types.String                                            `tfsdk:"error"`
	Request          *ResponseDevicesGetDeviceLiveToolsThroughputTestRequest `tfsdk:"request"`
	Result           *ResponseDevicesGetDeviceLiveToolsThroughputTestResult  `tfsdk:"result"`
	Status           types.String                                            `tfsdk:"status"`
	ThroughputTestID types.String                                            `tfsdk:"throughput_test_id"`
	URL              types.String                                            `tfsdk:"url"`
}

type ResponseDevicesGetDeviceLiveToolsThroughputTestRequest struct {
	Serial types.String `tfsdk:"serial"`
}

type ResponseDevicesGetDeviceLiveToolsThroughputTestResult struct {
	Speeds *ResponseDevicesGetDeviceLiveToolsThroughputTestResultSpeeds `tfsdk:"speeds"`
}

type ResponseDevicesGetDeviceLiveToolsThroughputTestResultSpeeds struct {
	Downstream types.Float64 `tfsdk:"downstream"`
}

// ToBody
func ResponseDevicesGetDeviceLiveToolsThroughputTestItemToBody(state DevicesLiveToolsThroughputTest, response *merakigosdk.ResponseDevicesGetDeviceLiveToolsThroughputTest) DevicesLiveToolsThroughputTest {
	itemState := ResponseDevicesGetDeviceLiveToolsThroughputTest{
		Error: types.StringValue(response.Error),
		Request: func() *ResponseDevicesGetDeviceLiveToolsThroughputTestRequest {
			if response.Request != nil {
				return &ResponseDevicesGetDeviceLiveToolsThroughputTestRequest{
					Serial: types.StringValue(response.Request.Serial),
				}
			}
			return nil
		}(),
		Result: func() *ResponseDevicesGetDeviceLiveToolsThroughputTestResult {
			if response.Result != nil {
				return &ResponseDevicesGetDeviceLiveToolsThroughputTestResult{
					Speeds: func() *ResponseDevicesGetDeviceLiveToolsThroughputTestResultSpeeds {
						if response.Result.Speeds != nil {
							return &ResponseDevicesGetDeviceLiveToolsThroughputTestResultSpeeds{
								Downstream: func() types.Float64 {
									if response.Result.Speeds.Downstream != nil {
										return types.Float64Value(float64(*response.Result.Speeds.Downstream))
									}
									return types.Float64{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		Status:           types.StringValue(response.Status),
		ThroughputTestID: types.StringValue(response.ThroughputTestID),
		URL:              types.StringValue(response.URL),
	}
	state.Item = &itemState
	return state
}
