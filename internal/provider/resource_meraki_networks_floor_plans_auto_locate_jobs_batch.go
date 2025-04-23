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
	_ resource.Resource              = &NetworksFloorPlansAutoLocateJobsBatchResource{}
	_ resource.ResourceWithConfigure = &NetworksFloorPlansAutoLocateJobsBatchResource{}
)

func NewNetworksFloorPlansAutoLocateJobsBatchResource() resource.Resource {
	return &NetworksFloorPlansAutoLocateJobsBatchResource{}
}

type NetworksFloorPlansAutoLocateJobsBatchResource struct {
	client *merakigosdk.Client
}

func (r *NetworksFloorPlansAutoLocateJobsBatchResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksFloorPlansAutoLocateJobsBatchResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_floor_plans_auto_locate_jobs_batch"
}

// resourceAction
func (r *NetworksFloorPlansAutoLocateJobsBatchResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"jobs": schema.SetNestedAttribute{
						MarkdownDescription: `The newly created jobs`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"completed": schema.SingleNestedAttribute{
									MarkdownDescription: `Auto locate job progress information`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"percentage": schema.Int64Attribute{
											MarkdownDescription: `Approximate auto locate job completion percentage`,
											Computed:            true,
										},
									},
								},
								"errors": schema.SetNestedAttribute{
									MarkdownDescription: `List of errors that occurred during a failed run of auto locate`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"source": schema.StringAttribute{
												MarkdownDescription: `The step of the auto locate process when the error occurred. Possible values: 'gnss', 'ranging', 'positioning'`,
												Computed:            true,
											},
											"type": schema.StringAttribute{
												MarkdownDescription: `The type of error that occurred. Possible values: 'failure', 'no neighbors', 'missing anchors', 'wrong anchors', 'missing ranging data', 'calculation failure', 'scheduling failure'`,
												Computed:            true,
											},
										},
									},
								},
								"floor_plan_id": schema.StringAttribute{
									MarkdownDescription: `Floor plan ID`,
									Computed:            true,
								},
								"gnss": schema.SingleNestedAttribute{
									MarkdownDescription: `GNSS (e.g. GPS) status and progress information`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"completed": schema.SingleNestedAttribute{
											MarkdownDescription: `Progress information for the GNSS acquisition process`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"percentage": schema.Int64Attribute{
													MarkdownDescription: `Completion percentage of the GNSS acquisition process`,
													Computed:            true,
												},
											},
										},
										"status": schema.StringAttribute{
											MarkdownDescription: `GNSS status. Possible values: 'scheduled', 'in progress', 'error', 'finished', 'not applicable', 'canceled'`,
											Computed:            true,
										},
									},
								},
								"id": schema.StringAttribute{
									MarkdownDescription: `Auto locate job ID`,
									Computed:            true,
								},
								"network_id": schema.StringAttribute{
									MarkdownDescription: `Network ID`,
									Computed:            true,
								},
								"ranging": schema.SingleNestedAttribute{
									MarkdownDescription: `Ranging status and progress information`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"completed": schema.SingleNestedAttribute{
											MarkdownDescription: `Progress information for the ranging process`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"percentage": schema.Int64Attribute{
													MarkdownDescription: `Completion percentage of the ranging process`,
													Computed:            true,
												},
											},
										},
										"status": schema.StringAttribute{
											MarkdownDescription: `Ranging status. Possible values: 'scheduled', 'in progress', 'error', 'finished', 'no neighbors'`,
											Computed:            true,
										},
									},
								},
								"scheduled_at": schema.StringAttribute{
									MarkdownDescription: `Scheduled start time for auto locate job`,
									Computed:            true,
								},
								"status": schema.StringAttribute{
									MarkdownDescription: `Auto locate job status. Possible values: 'scheduled', 'in progress', 'canceling', 'error', 'finished', 'published', 'canceled'`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"jobs": schema.SetNestedAttribute{
						MarkdownDescription: `The list of auto locate jobs to be scheduled. Up to 100 jobs can be provided in a request.`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"floor_plan_id": schema.StringAttribute{
									MarkdownDescription: `The ID of the floor plan to run auto locate for`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"refresh": schema.ListAttribute{
									MarkdownDescription: `The types of location data that should be refreshed for this job. The list must either contain both 'gnss' and 'ranging' or be empty, as we currently only support refreshing both 'gnss' and 'ranging', or neither.`,
									Optional:            true,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"scheduled_at": schema.StringAttribute{
									MarkdownDescription: `Timestamp in ISO8601 format which indicates when the auto locate job should be run. If omitted, the auto locate job will start immediately.`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
							},
						},
					},
				},
			},
		},
	}
}
func (r *NetworksFloorPlansAutoLocateJobsBatchResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksFloorPlansAutoLocateJobsBatch

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
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Networks.BatchNetworkFloorPlansAutoLocateJobs(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing BatchNetworkFloorPlansAutoLocateJobs",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing BatchNetworkFloorPlansAutoLocateJobs",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksFloorPlansAutoLocateJobsBatchResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksFloorPlansAutoLocateJobsBatchResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksFloorPlansAutoLocateJobsBatchResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksFloorPlansAutoLocateJobsBatch struct {
	NetworkID  types.String                                           `tfsdk:"network_id"`
	Item       *ResponseNetworksBatchNetworkFloorPlansAutoLocateJobs  `tfsdk:"item"`
	Parameters *RequestNetworksBatchNetworkFloorPlansAutoLocateJobsRs `tfsdk:"parameters"`
}

type ResponseNetworksBatchNetworkFloorPlansAutoLocateJobs struct {
	Jobs *[]ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobs `tfsdk:"jobs"`
}

type ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobs struct {
	Completed   *ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsCompleted `tfsdk:"completed"`
	Errors      *[]ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsErrors  `tfsdk:"errors"`
	FloorPlanID types.String                                                       `tfsdk:"floor_plan_id"`
	Gnss        *ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsGnss      `tfsdk:"gnss"`
	ID          types.String                                                       `tfsdk:"id"`
	NetworkID   types.String                                                       `tfsdk:"network_id"`
	Ranging     *ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsRanging   `tfsdk:"ranging"`
	ScheduledAt types.String                                                       `tfsdk:"scheduled_at"`
	Status      types.String                                                       `tfsdk:"status"`
}

type ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsCompleted struct {
	Percentage types.Int64 `tfsdk:"percentage"`
}

type ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsErrors struct {
	Source types.String `tfsdk:"source"`
	Type   types.String `tfsdk:"type"`
}

type ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsGnss struct {
	Completed *ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsGnssCompleted `tfsdk:"completed"`
	Status    types.String                                                           `tfsdk:"status"`
}

type ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsGnssCompleted struct {
	Percentage types.Int64 `tfsdk:"percentage"`
}

type ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsRanging struct {
	Completed *ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsRangingCompleted `tfsdk:"completed"`
	Status    types.String                                                              `tfsdk:"status"`
}

type ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsRangingCompleted struct {
	Percentage types.Int64 `tfsdk:"percentage"`
}

type RequestNetworksBatchNetworkFloorPlansAutoLocateJobsRs struct {
	Jobs *[]RequestNetworksBatchNetworkFloorPlansAutoLocateJobsJobsRs `tfsdk:"jobs"`
}

type RequestNetworksBatchNetworkFloorPlansAutoLocateJobsJobsRs struct {
	FloorPlanID types.String `tfsdk:"floor_plan_id"`
	Refresh     types.Set    `tfsdk:"refresh"`
	ScheduledAt types.String `tfsdk:"scheduled_at"`
}

// FromBody
func (r *NetworksFloorPlansAutoLocateJobsBatch) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksBatchNetworkFloorPlansAutoLocateJobs {
	re := *r.Parameters
	var requestNetworksBatchNetworkFloorPlansAutoLocateJobsJobs []merakigosdk.RequestNetworksBatchNetworkFloorPlansAutoLocateJobsJobs

	if re.Jobs != nil {
		for _, rItem1 := range *re.Jobs {
			floorPlanID := rItem1.FloorPlanID.ValueString()

			var refresh []string = nil
			rItem1.Refresh.ElementsAs(ctx, &refresh, false)
			scheduledAt := rItem1.ScheduledAt.ValueString()
			requestNetworksBatchNetworkFloorPlansAutoLocateJobsJobs = append(requestNetworksBatchNetworkFloorPlansAutoLocateJobsJobs, merakigosdk.RequestNetworksBatchNetworkFloorPlansAutoLocateJobsJobs{
				FloorPlanID: floorPlanID,
				Refresh:     refresh,
				ScheduledAt: scheduledAt,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestNetworksBatchNetworkFloorPlansAutoLocateJobs{
		Jobs: &requestNetworksBatchNetworkFloorPlansAutoLocateJobsJobs,
	}
	return &out
}

// ToBody
func ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsItemToBody(state NetworksFloorPlansAutoLocateJobsBatch, response *merakigosdk.ResponseNetworksBatchNetworkFloorPlansAutoLocateJobs) NetworksFloorPlansAutoLocateJobsBatch {
	itemState := ResponseNetworksBatchNetworkFloorPlansAutoLocateJobs{
		Jobs: func() *[]ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobs {
			if response.Jobs != nil {
				result := make([]ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobs, len(*response.Jobs))
				for i, jobs := range *response.Jobs {
					result[i] = ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobs{
						Completed: func() *ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsCompleted {
							if jobs.Completed != nil {
								return &ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsCompleted{
									Percentage: func() types.Int64 {
										if jobs.Completed.Percentage != nil {
											return types.Int64Value(int64(*jobs.Completed.Percentage))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						Errors: func() *[]ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsErrors {
							if jobs.Errors != nil {
								result := make([]ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsErrors, len(*jobs.Errors))
								for i, errors := range *jobs.Errors {
									result[i] = ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsErrors{
										Source: types.StringValue(errors.Source),
										Type:   types.StringValue(errors.Type),
									}
								}
								return &result
							}
							return nil
						}(),
						FloorPlanID: types.StringValue(jobs.FloorPlanID),
						Gnss: func() *ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsGnss {
							if jobs.Gnss != nil {
								return &ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsGnss{
									Completed: func() *ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsGnssCompleted {
										if jobs.Gnss.Completed != nil {
											return &ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsGnssCompleted{
												Percentage: func() types.Int64 {
													if jobs.Gnss.Completed.Percentage != nil {
														return types.Int64Value(int64(*jobs.Gnss.Completed.Percentage))
													}
													return types.Int64{}
												}(),
											}
										}
										return nil
									}(),
									Status: types.StringValue(jobs.Gnss.Status),
								}
							}
							return nil
						}(),
						ID:        types.StringValue(jobs.ID),
						NetworkID: types.StringValue(jobs.NetworkID),
						Ranging: func() *ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsRanging {
							if jobs.Ranging != nil {
								return &ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsRanging{
									Completed: func() *ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsRangingCompleted {
										if jobs.Ranging.Completed != nil {
											return &ResponseNetworksBatchNetworkFloorPlansAutoLocateJobsJobsRangingCompleted{
												Percentage: func() types.Int64 {
													if jobs.Ranging.Completed.Percentage != nil {
														return types.Int64Value(int64(*jobs.Ranging.Completed.Percentage))
													}
													return types.Int64{}
												}(),
											}
										}
										return nil
									}(),
									Status: types.StringValue(jobs.Ranging.Status),
								}
							}
							return nil
						}(),
						ScheduledAt: types.StringValue(jobs.ScheduledAt),
						Status:      types.StringValue(jobs.Status),
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
