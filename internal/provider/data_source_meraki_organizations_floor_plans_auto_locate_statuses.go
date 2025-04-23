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
	_ datasource.DataSource              = &OrganizationsFloorPlansAutoLocateStatusesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsFloorPlansAutoLocateStatusesDataSource{}
)

func NewOrganizationsFloorPlansAutoLocateStatusesDataSource() datasource.DataSource {
	return &OrganizationsFloorPlansAutoLocateStatusesDataSource{}
}

type OrganizationsFloorPlansAutoLocateStatusesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsFloorPlansAutoLocateStatusesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsFloorPlansAutoLocateStatusesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_floor_plans_auto_locate_statuses"
}

func (d *OrganizationsFloorPlansAutoLocateStatusesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"floor_plan_ids": schema.ListAttribute{
				MarkdownDescription: `floorPlanIds query parameter. Optional parameter to filter floorplans by one or more floorplan IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter floorplans by one or more network IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 10000. Default is 1000.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationFloorPlansAutoLocateStatuses`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"items": schema.ListNestedAttribute{
							MarkdownDescription: `Items in the paginated dataset`,
							Computed:            true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"counts": schema.SingleNestedAttribute{
										MarkdownDescription: `Counts for this floor plan`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"devices": schema.SingleNestedAttribute{
												MarkdownDescription: `Device counts for this floor plan`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"total": schema.Int64Attribute{
														MarkdownDescription: `The total number of devices that will participate if an auto locate job is started`,
														Computed:            true,
													},
												},
											},
										},
									},
									"floor_plan_id": schema.StringAttribute{
										MarkdownDescription: `Floor plan ID`,
										Computed:            true,
									},
									"jobs": schema.SetNestedAttribute{
										MarkdownDescription: `The most recent job for this floor plan`,
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
									"name": schema.StringAttribute{
										MarkdownDescription: `Floor plan name`,
										Computed:            true,
									},
									"network": schema.SingleNestedAttribute{
										MarkdownDescription: `Network info`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												MarkdownDescription: `ID for the network containing the floorplan`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
						"meta": schema.SingleNestedAttribute{
							MarkdownDescription: `Metadata relevant to the paginated dataset`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"counts": schema.SingleNestedAttribute{
									MarkdownDescription: `Counts relating to the paginated dataset`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"items": schema.SingleNestedAttribute{
											MarkdownDescription: `Counts relating to the paginated items`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"remaining": schema.Int64Attribute{
													MarkdownDescription: `The number of items in the dataset that are available on subsequent pages`,
													Computed:            true,
												},
												"total": schema.Int64Attribute{
													MarkdownDescription: `The total number of items in the dataset`,
													Computed:            true,
												},
											},
										},
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

func (d *OrganizationsFloorPlansAutoLocateStatusesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsFloorPlansAutoLocateStatuses OrganizationsFloorPlansAutoLocateStatuses
	diags := req.Config.Get(ctx, &organizationsFloorPlansAutoLocateStatuses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationFloorPlansAutoLocateStatuses")
		vvOrganizationID := organizationsFloorPlansAutoLocateStatuses.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationFloorPlansAutoLocateStatusesQueryParams{}

		queryParams1.PerPage = int(organizationsFloorPlansAutoLocateStatuses.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsFloorPlansAutoLocateStatuses.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsFloorPlansAutoLocateStatuses.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsFloorPlansAutoLocateStatuses.NetworkIDs)
		queryParams1.FloorPlanIDs = elementsToStrings(ctx, organizationsFloorPlansAutoLocateStatuses.FloorPlanIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationFloorPlansAutoLocateStatuses(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationFloorPlansAutoLocateStatuses",
				err.Error(),
			)
			return
		}

		organizationsFloorPlansAutoLocateStatuses = ResponseOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsToBody(organizationsFloorPlansAutoLocateStatuses, response1)
		diags = resp.State.Set(ctx, &organizationsFloorPlansAutoLocateStatuses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsFloorPlansAutoLocateStatuses struct {
	OrganizationID types.String                                                            `tfsdk:"organization_id"`
	PerPage        types.Int64                                                             `tfsdk:"per_page"`
	StartingAfter  types.String                                                            `tfsdk:"starting_after"`
	EndingBefore   types.String                                                            `tfsdk:"ending_before"`
	NetworkIDs     types.List                                                              `tfsdk:"network_ids"`
	FloorPlanIDs   types.List                                                              `tfsdk:"floor_plan_ids"`
	Items          *[]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatuses `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatuses struct {
	Items *[]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItems `tfsdk:"items"`
	Meta  *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesMeta    `tfsdk:"meta"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItems struct {
	Counts      *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsCounts  `tfsdk:"counts"`
	FloorPlanID types.String                                                                      `tfsdk:"floor_plan_id"`
	Jobs        *[]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobs  `tfsdk:"jobs"`
	Name        types.String                                                                      `tfsdk:"name"`
	Network     *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsNetwork `tfsdk:"network"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsCounts struct {
	Devices *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsCountsDevices `tfsdk:"devices"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsCountsDevices struct {
	Total types.Int64 `tfsdk:"total"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobs struct {
	Completed   *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsCompleted `tfsdk:"completed"`
	Errors      *[]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsErrors  `tfsdk:"errors"`
	Gnss        *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsGnss      `tfsdk:"gnss"`
	ID          types.String                                                                            `tfsdk:"id"`
	Ranging     *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsRanging   `tfsdk:"ranging"`
	ScheduledAt types.String                                                                            `tfsdk:"scheduled_at"`
	Status      types.String                                                                            `tfsdk:"status"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsCompleted struct {
	Percentage types.Int64 `tfsdk:"percentage"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsErrors struct {
	Source types.String `tfsdk:"source"`
	Type   types.String `tfsdk:"type"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsGnss struct {
	Completed *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsGnssCompleted `tfsdk:"completed"`
	Status    types.String                                                                                `tfsdk:"status"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsGnssCompleted struct {
	Percentage types.Int64 `tfsdk:"percentage"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsRanging struct {
	Completed *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsRangingCompleted `tfsdk:"completed"`
	Status    types.String                                                                                   `tfsdk:"status"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsRangingCompleted struct {
	Percentage types.Int64 `tfsdk:"percentage"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsNetwork struct {
	ID types.String `tfsdk:"id"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesMeta struct {
	Counts *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesMetaCounts `tfsdk:"counts"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesMetaCounts struct {
	Items *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesMetaCountsItems `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsToBody(state OrganizationsFloorPlansAutoLocateStatuses, response *merakigosdk.ResponseOrganizationsGetOrganizationFloorPlansAutoLocateStatuses) OrganizationsFloorPlansAutoLocateStatuses {
	var items []ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatuses
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatuses{
			Items: func() *[]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItems {
				if item.Items != nil {
					result := make([]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItems, len(*item.Items))
					for i, items := range *item.Items {
						result[i] = ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItems{
							Counts: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsCounts {
								if items.Counts != nil {
									return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsCounts{
										Devices: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsCountsDevices {
											if items.Counts.Devices != nil {
												return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsCountsDevices{
													Total: func() types.Int64 {
														if items.Counts.Devices.Total != nil {
															return types.Int64Value(int64(*items.Counts.Devices.Total))
														}
														return types.Int64{}
													}(),
												}
											}
											return nil
										}(),
									}
								}
								return nil
							}(),
							FloorPlanID: types.StringValue(items.FloorPlanID),
							Jobs: func() *[]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobs {
								if items.Jobs != nil {
									result := make([]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobs, len(*items.Jobs))
									for i, jobs := range *items.Jobs {
										result[i] = ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobs{
											Completed: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsCompleted {
												if jobs.Completed != nil {
													return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsCompleted{
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
											Errors: func() *[]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsErrors {
												if jobs.Errors != nil {
													result := make([]ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsErrors, len(*jobs.Errors))
													for i, errors := range *jobs.Errors {
														result[i] = ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsErrors{
															Source: types.StringValue(errors.Source),
															Type:   types.StringValue(errors.Type),
														}
													}
													return &result
												}
												return nil
											}(),
											Gnss: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsGnss {
												if jobs.Gnss != nil {
													return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsGnss{
														Completed: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsGnssCompleted {
															if jobs.Gnss.Completed != nil {
																return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsGnssCompleted{
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
											ID: types.StringValue(jobs.ID),
											Ranging: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsRanging {
												if jobs.Ranging != nil {
													return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsRanging{
														Completed: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsRangingCompleted {
															if jobs.Ranging.Completed != nil {
																return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsJobsRangingCompleted{
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
							Name: types.StringValue(items.Name),
							Network: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsNetwork {
								if items.Network != nil {
									return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesItemsNetwork{
										ID: types.StringValue(items.Network.ID),
									}
								}
								return nil
							}(),
						}
					}
					return &result
				}
				return nil
			}(),
			Meta: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesMeta {
				if item.Meta != nil {
					return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesMeta{
						Counts: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesMetaCounts {
							if item.Meta.Counts != nil {
								return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesMetaCounts{
									Items: func() *ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesMetaCountsItems {
										if item.Meta.Counts.Items != nil {
											return &ResponseItemOrganizationsGetOrganizationFloorPlansAutoLocateStatusesMetaCountsItems{
												Remaining: func() types.Int64 {
													if item.Meta.Counts.Items.Remaining != nil {
														return types.Int64Value(int64(*item.Meta.Counts.Items.Remaining))
													}
													return types.Int64{}
												}(),
												Total: func() types.Int64 {
													if item.Meta.Counts.Items.Total != nil {
														return types.Int64Value(int64(*item.Meta.Counts.Items.Total))
													}
													return types.Int64{}
												}(),
											}
										}
										return nil
									}(),
								}
							}
							return nil
						}(),
					}
				}
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
