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
	_ datasource.DataSource              = &OrganizationsDevicesUplinksLossAndLatencyDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesUplinksLossAndLatencyDataSource{}
)

func NewOrganizationsDevicesUplinksLossAndLatencyDataSource() datasource.DataSource {
	return &OrganizationsDevicesUplinksLossAndLatencyDataSource{}
}

type OrganizationsDevicesUplinksLossAndLatencyDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesUplinksLossAndLatencyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesUplinksLossAndLatencyDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_uplinks_loss_and_latency"
}

func (d *OrganizationsDevicesUplinksLossAndLatencyDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ip": schema.StringAttribute{
				MarkdownDescription: `ip query parameter. Optional filter for a specific destination IP. Default will return all destination IPs.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data. The maximum lookback period is 60 days from today.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 5 minutes after t0. The latest possible time that t1 can be is 2 minutes into the past.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 5 minutes. The default is 5 minutes.`,
				Optional:            true,
			},
			"uplink": schema.StringAttribute{
				MarkdownDescription: `uplink query parameter. Optional filter for a specific WAN uplink. Valid uplinks are wan1, wan2, wan3, cellular. Default will return all uplinks.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationDevicesUplinksLossAndLatency`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"ip": schema.StringAttribute{
							MarkdownDescription: `IP address of uplink`,
							Computed:            true,
						},
						"network_id": schema.StringAttribute{
							MarkdownDescription: `Network ID`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial of MX device`,
							Computed:            true,
						},
						"time_series": schema.SetNestedAttribute{
							MarkdownDescription: `Loss and latency timeseries data`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"latency_ms": schema.Float64Attribute{
										MarkdownDescription: `Latency in milliseconds`,
										Computed:            true,
									},
									"loss_percent": schema.Float64Attribute{
										MarkdownDescription: `Loss percentage`,
										Computed:            true,
									},
									"ts": schema.StringAttribute{
										MarkdownDescription: `Timestamp for this data point`,
										Computed:            true,
									},
								},
							},
						},
						"uplink": schema.StringAttribute{
							MarkdownDescription: `Uplink interface (wan1, wan2, or cellular)`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsDevicesUplinksLossAndLatencyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevicesUplinksLossAndLatency OrganizationsDevicesUplinksLossAndLatency
	diags := req.Config.Get(ctx, &organizationsDevicesUplinksLossAndLatency)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevicesUplinksLossAndLatency")
		vvOrganizationID := organizationsDevicesUplinksLossAndLatency.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesUplinksLossAndLatencyQueryParams{}

		queryParams1.T0 = organizationsDevicesUplinksLossAndLatency.T0.ValueString()
		queryParams1.T1 = organizationsDevicesUplinksLossAndLatency.T1.ValueString()
		queryParams1.Timespan = organizationsDevicesUplinksLossAndLatency.Timespan.ValueFloat64()
		queryParams1.Uplink = organizationsDevicesUplinksLossAndLatency.Uplink.ValueString()
		queryParams1.IP = organizationsDevicesUplinksLossAndLatency.IP.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevicesUplinksLossAndLatency(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevicesUplinksLossAndLatency",
				err.Error(),
			)
			return
		}

		organizationsDevicesUplinksLossAndLatency = ResponseOrganizationsGetOrganizationDevicesUplinksLossAndLatencyItemsToBody(organizationsDevicesUplinksLossAndLatency, response1)
		diags = resp.State.Set(ctx, &organizationsDevicesUplinksLossAndLatency)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevicesUplinksLossAndLatency struct {
	OrganizationID types.String                                                            `tfsdk:"organization_id"`
	T0             types.String                                                            `tfsdk:"t0"`
	T1             types.String                                                            `tfsdk:"t1"`
	Timespan       types.Float64                                                           `tfsdk:"timespan"`
	Uplink         types.String                                                            `tfsdk:"uplink"`
	IP             types.String                                                            `tfsdk:"ip"`
	Items          *[]ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatency `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatency struct {
	IP         types.String                                                                      `tfsdk:"ip"`
	NetworkID  types.String                                                                      `tfsdk:"network_id"`
	Serial     types.String                                                                      `tfsdk:"serial"`
	TimeSeries *[]ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatencyTimeSeries `tfsdk:"time_series"`
	Uplink     types.String                                                                      `tfsdk:"uplink"`
}

type ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatencyTimeSeries struct {
	LatencyMs   types.Float64 `tfsdk:"latency_ms"`
	LossPercent types.Float64 `tfsdk:"loss_percent"`
	Ts          types.String  `tfsdk:"ts"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesUplinksLossAndLatencyItemsToBody(state OrganizationsDevicesUplinksLossAndLatency, response *merakigosdk.ResponseOrganizationsGetOrganizationDevicesUplinksLossAndLatency) OrganizationsDevicesUplinksLossAndLatency {
	var items []ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatency
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatency{
			IP:        types.StringValue(item.IP),
			NetworkID: types.StringValue(item.NetworkID),
			Serial:    types.StringValue(item.Serial),
			TimeSeries: func() *[]ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatencyTimeSeries {
				if item.TimeSeries != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatencyTimeSeries, len(*item.TimeSeries))
					for i, timeSeries := range *item.TimeSeries {
						result[i] = ResponseItemOrganizationsGetOrganizationDevicesUplinksLossAndLatencyTimeSeries{
							LatencyMs: func() types.Float64 {
								if timeSeries.LatencyMs != nil {
									return types.Float64Value(float64(*timeSeries.LatencyMs))
								}
								return types.Float64{}
							}(),
							LossPercent: func() types.Float64 {
								if timeSeries.LossPercent != nil {
									return types.Float64Value(float64(*timeSeries.LossPercent))
								}
								return types.Float64{}
							}(),
							Ts: types.StringValue(timeSeries.Ts),
						}
					}
					return &result
				}
				return nil
			}(),
			Uplink: types.StringValue(item.Uplink),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
