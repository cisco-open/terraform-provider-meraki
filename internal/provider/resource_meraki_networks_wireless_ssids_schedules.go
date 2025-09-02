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
	"strconv"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsSchedulesResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsSchedulesResource{}
)

func NewNetworksWirelessSSIDsSchedulesResource() resource.Resource {
	return &NetworksWirelessSSIDsSchedulesResource{}
}

type NetworksWirelessSSIDsSchedulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsSchedulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsSchedulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_schedules"
}

func (r *NetworksWirelessSSIDsSchedulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `If true, the SSID outage schedule is enabled.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"ranges": schema.ListNestedAttribute{
				MarkdownDescription: `List of outage ranges. Has a start date and time, and end date and time. If this parameter is passed in along with rangesInSeconds parameter, this will take precedence.`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"end_day": schema.StringAttribute{
							MarkdownDescription: `Day of when the outage ends. Can be either full day name, or three letter abbreviation`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"end_time": schema.StringAttribute{
							MarkdownDescription: `24 hour time when the outage ends.`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"start_day": schema.StringAttribute{
							MarkdownDescription: `Day of when the outage starts. Can be either full day name, or three letter abbreviation.`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"start_time": schema.StringAttribute{
							MarkdownDescription: `24 hour time when the outage starts.`,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"ranges_in_seconds": schema.ListNestedAttribute{
				MarkdownDescription: `List of outage ranges in seconds since Sunday at Midnight. Has a start and end. If this parameter is passed in along with the ranges parameter, ranges will take precedence.`,
				Optional:            true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"end": schema.Int64Attribute{
							MarkdownDescription: `Seconds since Sunday at midnight when that outage range ends.`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"start": schema.Int64Attribute{
							MarkdownDescription: `Seconds since Sunday at midnight when the outage range starts.`,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksWirelessSSIDsSchedulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsSchedulesRs

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
	// Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	vvNumber := data.Number.ValueString()
	//Has Item and not has items

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDSchedules(vvNetworkID, vvNumber, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDSchedules",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDSchedules",
			err.Error(),
		)
		return
	}

	// Assign data
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *NetworksWirelessSSIDsSchedulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsSchedulesRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	vvNumber := data.Number.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDSchedules(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkWirelessSSIDSchedules",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDSchedules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDSchedulesItemToBodyRs(data, responseGet, true)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksWirelessSSIDsSchedulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: networkId,number. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
}

func (r *NetworksWirelessSSIDsSchedulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksWirelessSSIDsSchedulesRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	vvNumber := plan.Number.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDSchedules(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDSchedules",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDSchedules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *NetworksWirelessSSIDsSchedulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessSSIDsSchedules", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsSchedulesRs struct {
	NetworkID       types.String                                                        `tfsdk:"network_id"`
	Number          types.String                                                        `tfsdk:"number"`
	Enabled         types.Bool                                                          `tfsdk:"enabled"`
	Ranges          *[]ResponseWirelessGetNetworkWirelessSsidSchedulesRangesRs          `tfsdk:"ranges"`
	RangesInSeconds *[]ResponseWirelessGetNetworkWirelessSsidSchedulesRangesInSecondsRs `tfsdk:"ranges_in_seconds"`
}

type ResponseWirelessGetNetworkWirelessSsidSchedulesRangesRs struct {
	EndDay    types.String `tfsdk:"end_day"`
	EndTime   types.String `tfsdk:"end_time"`
	StartDay  types.String `tfsdk:"start_day"`
	StartTime types.String `tfsdk:"start_time"`
}

type ResponseWirelessGetNetworkWirelessSsidSchedulesRangesInSecondsRs struct {
	End   types.Int64 `tfsdk:"end"`
	Start types.Int64 `tfsdk:"start"`
}

// FromBody
func (r *NetworksWirelessSSIDsSchedulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSchedules {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var requestWirelessUpdateNetworkWirelessSSIDSchedulesRanges []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSchedulesRanges

	if r.Ranges != nil {
		for _, rItem1 := range *r.Ranges {
			endDay := rItem1.EndDay.ValueString()
			endTime := rItem1.EndTime.ValueString()
			startDay := rItem1.StartDay.ValueString()
			startTime := rItem1.StartTime.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDSchedulesRanges = append(requestWirelessUpdateNetworkWirelessSSIDSchedulesRanges, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSchedulesRanges{
				EndDay:    endDay,
				EndTime:   endTime,
				StartDay:  startDay,
				StartTime: startTime,
			})
			//[debug] Is Array: True
		}
	}
	var requestWirelessUpdateNetworkWirelessSSIDSchedulesRangesInSeconds []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSchedulesRangesInSeconds

	if r.RangesInSeconds != nil {
		for _, rItem1 := range *r.RangesInSeconds {
			end := func() *int64 {
				if !rItem1.End.IsUnknown() && !rItem1.End.IsNull() {
					return rItem1.End.ValueInt64Pointer()
				}
				return nil
			}()
			start := func() *int64 {
				if !rItem1.Start.IsUnknown() && !rItem1.Start.IsNull() {
					return rItem1.Start.ValueInt64Pointer()
				}
				return nil
			}()
			requestWirelessUpdateNetworkWirelessSSIDSchedulesRangesInSeconds = append(requestWirelessUpdateNetworkWirelessSSIDSchedulesRangesInSeconds, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSchedulesRangesInSeconds{
				End:   int64ToIntPointer(end),
				Start: int64ToIntPointer(start),
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSchedules{
		Enabled: enabled,
		Ranges: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSchedulesRanges {
			if len(requestWirelessUpdateNetworkWirelessSSIDSchedulesRanges) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDSchedulesRanges
			}
			return nil
		}(),
		RangesInSeconds: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDSchedulesRangesInSeconds {
			if len(requestWirelessUpdateNetworkWirelessSSIDSchedulesRangesInSeconds) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDSchedulesRangesInSeconds
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDSchedulesItemToBodyRs(state NetworksWirelessSSIDsSchedulesRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDSchedules, is_read bool) NetworksWirelessSSIDsSchedulesRs {
	itemState := NetworksWirelessSSIDsSchedulesRs{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Ranges: func() *[]ResponseWirelessGetNetworkWirelessSsidSchedulesRangesRs {
			if response.Ranges != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidSchedulesRangesRs, len(*response.Ranges))
				for i, ranges := range *response.Ranges {
					result[i] = ResponseWirelessGetNetworkWirelessSsidSchedulesRangesRs{
						EndDay: func() types.String {
							if ranges.EndDay != "" {
								return types.StringValue(ranges.EndDay)
							}
							return types.String{}
						}(),
						EndTime: func() types.String {
							if ranges.EndTime != "" {
								return types.StringValue(ranges.EndTime)
							}
							return types.String{}
						}(),
						StartDay: func() types.String {
							if ranges.StartDay != "" {
								return types.StringValue(ranges.StartDay)
							}
							return types.String{}
						}(),
						StartTime: func() types.String {
							if ranges.StartTime != "" {
								return types.StringValue(ranges.StartTime)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		RangesInSeconds: func() *[]ResponseWirelessGetNetworkWirelessSsidSchedulesRangesInSecondsRs {
			if response.RangesInSeconds != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidSchedulesRangesInSecondsRs, len(*response.RangesInSeconds))
				for i, rangesInSeconds := range *response.RangesInSeconds {
					result[i] = ResponseWirelessGetNetworkWirelessSsidSchedulesRangesInSecondsRs{
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
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsSchedulesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsSchedulesRs)
}
