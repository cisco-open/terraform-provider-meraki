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
	_ datasource.DataSource              = &NetworksWirelessSSIDsSchedulesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksWirelessSSIDsSchedulesDataSource{}
)

func NewNetworksWirelessSSIDsSchedulesDataSource() datasource.DataSource {
	return &NetworksWirelessSSIDsSchedulesDataSource{}
}

type NetworksWirelessSSIDsSchedulesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksWirelessSSIDsSchedulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksWirelessSSIDsSchedulesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_schedules"
}

func (d *NetworksWirelessSSIDsSchedulesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `If true, the SSID outage schedule is enabled.`,
						Computed:            true,
					},
					"ranges": schema.SetNestedAttribute{
						MarkdownDescription: `List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"end_day": schema.StringAttribute{
									MarkdownDescription: `Day of when the outage ends. Can be either full day name, or three letter abbreviation`,
									Computed:            true,
								},
								"end_time": schema.StringAttribute{
									MarkdownDescription: `24 hour time when the outage ends.`,
									Computed:            true,
								},
								"start_day": schema.StringAttribute{
									MarkdownDescription: `Day of when the outage starts. Can be either full day name, or three letter abbreviation.`,
									Computed:            true,
								},
								"start_time": schema.StringAttribute{
									MarkdownDescription: `24 hour time when the outage starts.`,
									Computed:            true,
								},
							},
						},
					},
					"ranges_in_seconds": schema.SetNestedAttribute{
						MarkdownDescription: `List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence.`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"end": schema.Int64Attribute{
									MarkdownDescription: `Seconds since Sunday at midnight when that outage range ends.`,
									Computed:            true,
								},
								"start": schema.Int64Attribute{
									MarkdownDescription: `Seconds since Sunday at midnight when the outage range starts.`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksWirelessSSIDsSchedulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksWirelessSSIDsSchedules NetworksWirelessSSIDsSchedules
	diags := req.Config.Get(ctx, &networksWirelessSSIDsSchedules)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkWirelessSSIDSchedules")
		vvNetworkID := networksWirelessSSIDsSchedules.NetworkID.ValueString()
		vvNumber := networksWirelessSSIDsSchedules.Number.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Wireless.GetNetworkWirelessSSIDSchedules(vvNetworkID, vvNumber)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDSchedules",
				err.Error(),
			)
			return
		}

		networksWirelessSSIDsSchedules = ResponseWirelessGetNetworkWirelessSSIDSchedulesItemToBody(networksWirelessSSIDsSchedules, response1)
		diags = resp.State.Set(ctx, &networksWirelessSSIDsSchedules)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksWirelessSSIDsSchedules struct {
	NetworkID types.String                                     `tfsdk:"network_id"`
	Number    types.String                                     `tfsdk:"number"`
	Item      *ResponseWirelessGetNetworkWirelessSsidSchedules `tfsdk:"item"`
}

type ResponseWirelessGetNetworkWirelessSsidSchedules struct {
	Enabled         types.Bool                                                        `tfsdk:"enabled"`
	Ranges          *[]ResponseWirelessGetNetworkWirelessSsidSchedulesRanges          `tfsdk:"ranges"`
	RangesInSeconds *[]ResponseWirelessGetNetworkWirelessSsidSchedulesRangesInSeconds `tfsdk:"ranges_in_seconds"`
}

type ResponseWirelessGetNetworkWirelessSsidSchedulesRanges struct {
	EndDay    types.String `tfsdk:"end_day"`
	EndTime   types.String `tfsdk:"end_time"`
	StartDay  types.String `tfsdk:"start_day"`
	StartTime types.String `tfsdk:"start_time"`
}

type ResponseWirelessGetNetworkWirelessSsidSchedulesRangesInSeconds struct {
	End   types.Int64 `tfsdk:"end"`
	Start types.Int64 `tfsdk:"start"`
}

// ToBody
func ResponseWirelessGetNetworkWirelessSSIDSchedulesItemToBody(state NetworksWirelessSSIDsSchedules, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDSchedules) NetworksWirelessSSIDsSchedules {
	itemState := ResponseWirelessGetNetworkWirelessSsidSchedules{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Ranges: func() *[]ResponseWirelessGetNetworkWirelessSsidSchedulesRanges {
			if response.Ranges != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidSchedulesRanges, len(*response.Ranges))
				for i, ranges := range *response.Ranges {
					result[i] = ResponseWirelessGetNetworkWirelessSsidSchedulesRanges{
						EndDay:    types.StringValue(ranges.EndDay),
						EndTime:   types.StringValue(ranges.EndTime),
						StartDay:  types.StringValue(ranges.StartDay),
						StartTime: types.StringValue(ranges.StartTime),
					}
				}
				return &result
			}
			return nil
		}(),
		RangesInSeconds: func() *[]ResponseWirelessGetNetworkWirelessSsidSchedulesRangesInSeconds {
			if response.RangesInSeconds != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidSchedulesRangesInSeconds, len(*response.RangesInSeconds))
				for i, rangesInSeconds := range *response.RangesInSeconds {
					result[i] = ResponseWirelessGetNetworkWirelessSsidSchedulesRangesInSeconds{
						End: func() types.Int64 {
							if rangesInSeconds.End != nil {
								return types.Int64Value(int64(*rangesInSeconds.End))
							}
							return types.Int64{}
						}(),
						Start: func() types.Int64 {
							if rangesInSeconds.Start != nil {
								return types.Int64Value(int64(*rangesInSeconds.Start))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
