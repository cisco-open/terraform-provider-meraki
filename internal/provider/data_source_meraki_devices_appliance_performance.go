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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &DevicesAppliancePerformanceDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesAppliancePerformanceDataSource{}
)

func NewDevicesAppliancePerformanceDataSource() datasource.DataSource {
	return &DevicesAppliancePerformanceDataSource{}
}

type DevicesAppliancePerformanceDataSource struct {
	client *merakigosdk.Client
}

func (d *DevicesAppliancePerformanceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *DevicesAppliancePerformanceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_appliance_performance"
}

func (d *DevicesAppliancePerformanceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"perf_score": schema.Int64Attribute{
						Computed: true,
					},
				},
			},
		},
	}
}

func (d *DevicesAppliancePerformanceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var devicesAppliancePerformance DevicesAppliancePerformance
	diags := req.Config.Get(ctx, &devicesAppliancePerformance)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceAppliancePerformance")
		vvSerial := devicesAppliancePerformance.Serial.ValueString()

		response1, restyResp1, err := d.client.Appliance.GetDeviceAppliancePerformance(vvSerial)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceAppliancePerformance",
				err.Error(),
			)
			return
		}

		devicesAppliancePerformance = ResponseApplianceGetDeviceAppliancePerformanceItemToBody(devicesAppliancePerformance, response1)
		diags = resp.State.Set(ctx, &devicesAppliancePerformance)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type DevicesAppliancePerformance struct {
	Serial types.String                                    `tfsdk:"serial"`
	Item   *ResponseApplianceGetDeviceAppliancePerformance `tfsdk:"item"`
}

type ResponseApplianceGetDeviceAppliancePerformance struct {
	PerfScore types.Int64 `tfsdk:"perf_score"`
}

// ToBody
func ResponseApplianceGetDeviceAppliancePerformanceItemToBody(state DevicesAppliancePerformance, response *merakigosdk.ResponseApplianceGetDeviceAppliancePerformance) DevicesAppliancePerformance {
	itemState := ResponseApplianceGetDeviceAppliancePerformance{
		PerfScore: func() types.Int64 {
			if response.PerfScore != nil {
				return types.Int64Value(int64(*response.PerfScore))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
