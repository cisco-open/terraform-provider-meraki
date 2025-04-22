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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DevicesLiveToolsCableDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesLiveToolsCableDataSource{}
)

func NewDevicesLiveToolsCableDataSource() datasource.DataSource {
	return &DevicesLiveToolsCableDataSource{}
}

type DevicesLiveToolsCableDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesLiveToolsCableDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesLiveToolsCableDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_live_tools_cable"
}

func (d *DevicesLiveToolsCableDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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

					"cable_test_id": schema.StringAttribute{
						MarkdownDescription: `Id of the cable test request. Used to check the status of the request.`,
						Computed:            true,
					},
					"error": schema.StringAttribute{
						MarkdownDescription: `An error message for a failed execution`,
						Computed:            true,
					},
					"request": schema.SingleNestedAttribute{
						MarkdownDescription: `Cable test request parameters`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"ports": schema.ListAttribute{
								MarkdownDescription: `A list of ports for which to perform the cable test.`,
								Computed:            true,
								ElementType:         types.StringType,
							},
							"serial": schema.StringAttribute{
								MarkdownDescription: `Device serial number`,
								Computed:            true,
							},
						},
					},
					"results": schema.SetNestedAttribute{
						MarkdownDescription: `Results of the cable test request, one for each requested port.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"error": schema.StringAttribute{
									MarkdownDescription: `If an error occurred during the cable test, the error message will be populated here.`,
									Computed:            true,
								},
								"pairs": schema.SetNestedAttribute{
									MarkdownDescription: `Results for each twisted pair within the cable.`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"index": schema.Int64Attribute{
												MarkdownDescription: `The index of the twisted pair tested.`,
												Computed:            true,
											},
											"length_meters": schema.Int64Attribute{
												MarkdownDescription: `The detected length of the twisted pair.`,
												Computed:            true,
											},
											"status": schema.StringAttribute{
												MarkdownDescription: `The test result of the twisted pair tested.`,
												Computed:            true,
											},
										},
									},
								},
								"port": schema.StringAttribute{
									MarkdownDescription: `The port for which the test was performed.`,
									Computed:            true,
								},
								"speed_mbps": schema.Int64Attribute{
									MarkdownDescription: `Speed in Mbps.  A speed of 0 indicates the port is down or the port speed is automatic.`,
									Computed:            true,
								},
								"status": schema.StringAttribute{
									MarkdownDescription: `The current status of the port. If the cable test is still being performed on the port, "in-progress" is used. If an error occurred during the cable test, "error" is used and the error property will be populated.`,
									Computed:            true,
								},
							},
						},
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `Status of the cable test request.`,
						Computed:            true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `GET this url to check the status of your cable test request.`,
						Computed:            true,
					},
				},
			},
		},
	}
}

func (d *DevicesLiveToolsCableDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesLiveToolsCable DevicesLiveToolsCable
	diags := req.Config.Get(ctx, &devicesLiveToolsCable)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceLiveToolsCableTest")
		vvSerial := devicesLiveToolsCable.Serial.ValueString()
		vvID := devicesLiveToolsCable.ID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Devices.GetDeviceLiveToolsCableTest(vvSerial, vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceLiveToolsCableTest",
				err.Error(),
			)
			return
		}

		devicesLiveToolsCable = ResponseDevicesGetDeviceLiveToolsCableTestItemToBody(devicesLiveToolsCable, response1)
		diags = resp.State.Set(ctx, &devicesLiveToolsCable)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesLiveToolsCable struct {
	Serial types.String                                `tfsdk:"serial"`
	ID     types.String                                `tfsdk:"id"`
	Item   *ResponseDevicesGetDeviceLiveToolsCableTest `tfsdk:"item"`
}

type ResponseDevicesGetDeviceLiveToolsCableTest struct {
	CableTestID types.String                                         `tfsdk:"cable_test_id"`
	Error       types.String                                         `tfsdk:"error"`
	Request     *ResponseDevicesGetDeviceLiveToolsCableTestRequest   `tfsdk:"request"`
	Results     *[]ResponseDevicesGetDeviceLiveToolsCableTestResults `tfsdk:"results"`
	Status      types.String                                         `tfsdk:"status"`
	URL         types.String                                         `tfsdk:"url"`
}

type ResponseDevicesGetDeviceLiveToolsCableTestRequest struct {
	Ports  types.List   `tfsdk:"ports"`
	Serial types.String `tfsdk:"serial"`
}

type ResponseDevicesGetDeviceLiveToolsCableTestResults struct {
	Error     types.String                                              `tfsdk:"error"`
	Pairs     *[]ResponseDevicesGetDeviceLiveToolsCableTestResultsPairs `tfsdk:"pairs"`
	Port      types.String                                              `tfsdk:"port"`
	SpeedMbps types.Int64                                               `tfsdk:"speed_mbps"`
	Status    types.String                                              `tfsdk:"status"`
}

type ResponseDevicesGetDeviceLiveToolsCableTestResultsPairs struct {
	Index        types.Int64  `tfsdk:"index"`
	LengthMeters types.Int64  `tfsdk:"length_meters"`
	Status       types.String `tfsdk:"status"`
}

// ToBody
func ResponseDevicesGetDeviceLiveToolsCableTestItemToBody(state DevicesLiveToolsCable, response *merakigosdk.ResponseDevicesGetDeviceLiveToolsCableTest) DevicesLiveToolsCable {
	itemState := ResponseDevicesGetDeviceLiveToolsCableTest{
		CableTestID: types.StringValue(response.CableTestID),
		Error:       types.StringValue(response.Error),
		Request: func() *ResponseDevicesGetDeviceLiveToolsCableTestRequest {
			if response.Request != nil {
				return &ResponseDevicesGetDeviceLiveToolsCableTestRequest{
					Ports:  StringSliceToList(response.Request.Ports),
					Serial: types.StringValue(response.Request.Serial),
				}
			}
			return nil
		}(),
		Results: func() *[]ResponseDevicesGetDeviceLiveToolsCableTestResults {
			if response.Results != nil {
				result := make([]ResponseDevicesGetDeviceLiveToolsCableTestResults, len(*response.Results))
				for i, results := range *response.Results {
					result[i] = ResponseDevicesGetDeviceLiveToolsCableTestResults{
						Error: types.StringValue(results.Error),
						Pairs: func() *[]ResponseDevicesGetDeviceLiveToolsCableTestResultsPairs {
							if results.Pairs != nil {
								result := make([]ResponseDevicesGetDeviceLiveToolsCableTestResultsPairs, len(*results.Pairs))
								for i, pairs := range *results.Pairs {
									result[i] = ResponseDevicesGetDeviceLiveToolsCableTestResultsPairs{
										Index: func() types.Int64 {
											if pairs.Index != nil {
												return types.Int64Value(int64(*pairs.Index))
											}
											return types.Int64{}
										}(),
										LengthMeters: func() types.Int64 {
											if pairs.LengthMeters != nil {
												return types.Int64Value(int64(*pairs.LengthMeters))
											}
											return types.Int64{}
										}(),
										Status: types.StringValue(pairs.Status),
									}
								}
								return &result
							}
							return nil
						}(),
						Port: types.StringValue(results.Port),
						SpeedMbps: func() types.Int64 {
							if results.SpeedMbps != nil {
								return types.Int64Value(int64(*results.SpeedMbps))
							}
							return types.Int64{}
						}(),
						Status: types.StringValue(results.Status),
					}
				}
				return &result
			}
			return nil
		}(),
		Status: types.StringValue(response.Status),
		URL:    types.StringValue(response.URL),
	}
	state.Item = &itemState
	return state
}
