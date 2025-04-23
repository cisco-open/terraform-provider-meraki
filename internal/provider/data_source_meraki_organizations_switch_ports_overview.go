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
	_ datasource.DataSource              = &OrganizationsSwitchPortsOverviewDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsSwitchPortsOverviewDataSource{}
)

func NewOrganizationsSwitchPortsOverviewDataSource() datasource.DataSource {
	return &OrganizationsSwitchPortsOverviewDataSource{}
}

type OrganizationsSwitchPortsOverviewDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsSwitchPortsOverviewDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsSwitchPortsOverviewDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_switch_ports_overview"
}

func (d *OrganizationsSwitchPortsOverviewDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"t0": schema.StringAttribute{
				MarkdownDescription: `t0 query parameter. The beginning of the timespan for the data.`,
				Optional:            true,
			},
			"t1": schema.StringAttribute{
				MarkdownDescription: `t1 query parameter. The end of the timespan for the data. t1 can be a maximum of 186 days after t0.`,
				Optional:            true,
			},
			"timespan": schema.Float64Attribute{
				MarkdownDescription: `timespan query parameter. The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be greater than or equal to 12 hours and be less than or equal to 186 days. The default is 1 day.`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"counts": schema.SingleNestedAttribute{
						MarkdownDescription: `The count data of all ports`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"by_status": schema.SingleNestedAttribute{
								MarkdownDescription: `The count data, indexed by active or inactive status`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"active": schema.SingleNestedAttribute{
										MarkdownDescription: `The count data for active ports`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"by_media_and_link_speed": schema.SingleNestedAttribute{
												MarkdownDescription: `The active count data, indexed by media type (RJ45 or SFP)`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"rj45": schema.SingleNestedAttribute{
														MarkdownDescription: `The count data for RJ45 ports, indexed by speed in Mb`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"status_10": schema.Int64Attribute{
																MarkdownDescription: `The number of active 10 Mbps RJ45 ports`,
																Computed:            true,
															},
															"status_100": schema.Int64Attribute{
																MarkdownDescription: `The number of active 100 Mbps RJ45 ports`,
																Computed:            true,
															},
															"status_1000": schema.Int64Attribute{
																MarkdownDescription: `The number of active 1 Gbps RJ45 ports`,
																Computed:            true,
															},
															"status_10000": schema.Int64Attribute{
																MarkdownDescription: `The number of active 10 Gbps RJ45 ports`,
																Computed:            true,
															},
															"status_2500": schema.Int64Attribute{
																MarkdownDescription: `The number of active 2 Gbps RJ45 ports`,
																Computed:            true,
															},
															"status_5000": schema.Int64Attribute{
																MarkdownDescription: `The number of active 5 Gbps RJ45 ports`,
																Computed:            true,
															},
															"total": schema.Int64Attribute{
																MarkdownDescription: `The total number of active RJ45 ports`,
																Computed:            true,
															},
														},
													},
													"sfp": schema.SingleNestedAttribute{
														MarkdownDescription: `The count data for SFP ports, indexed by speed in Mb`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"status_100": schema.Int64Attribute{
																MarkdownDescription: `The number of active 100 Mbps SFP ports`,
																Computed:            true,
															},
															"status_1000": schema.Int64Attribute{
																MarkdownDescription: `The number of active 1 Gbps SFP ports`,
																Computed:            true,
															},
															"status_10000": schema.Int64Attribute{
																MarkdownDescription: `The number of active 10 Gbps SFP ports`,
																Computed:            true,
															},
															"status_100000": schema.Int64Attribute{
																MarkdownDescription: `The number of active 100 Gbps SFP ports`,
																Computed:            true,
															},
															"status_20000": schema.Int64Attribute{
																MarkdownDescription: `The number of active 20 Gbps SFP ports`,
																Computed:            true,
															},
															"status_25000": schema.Int64Attribute{
																MarkdownDescription: `The number of active 25 Gbps SFP ports`,
																Computed:            true,
															},
															"status_40000": schema.Int64Attribute{
																MarkdownDescription: `The number of active 40 Gbps SFP ports`,
																Computed:            true,
															},
															"status_50000": schema.Int64Attribute{
																MarkdownDescription: `The number of active 50 Gbps SFP ports`,
																Computed:            true,
															},
															"total": schema.Int64Attribute{
																MarkdownDescription: `The total number of active SFP ports`,
																Computed:            true,
															},
														},
													},
												},
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `The total number of active ports`,
												Computed:            true,
											},
										},
									},
									"inactive": schema.SingleNestedAttribute{
										MarkdownDescription: `The count data for inactive ports`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"by_media": schema.SingleNestedAttribute{
												MarkdownDescription: `The inactive count data, indexed by media type (RJ45 or SFP)`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"rj45": schema.SingleNestedAttribute{
														MarkdownDescription: `The count data for inactive RJ45 ports`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"total": schema.Int64Attribute{
																MarkdownDescription: `The total number of inactive RJ45 ports`,
																Computed:            true,
															},
														},
													},
													"sfp": schema.SingleNestedAttribute{
														MarkdownDescription: `The count data for inactive SFP ports`,
														Computed:            true,
														Attributes: map[string]schema.Attribute{

															"total": schema.Int64Attribute{
																MarkdownDescription: `The total number of inactive SFP ports`,
																Computed:            true,
															},
														},
													},
												},
											},
											"total": schema.Int64Attribute{
												MarkdownDescription: `The total number of inactive ports`,
												Computed:            true,
											},
										},
									},
								},
							},
							"total": schema.Int64Attribute{
								MarkdownDescription: `The total number of ports`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsSwitchPortsOverviewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsSwitchPortsOverview OrganizationsSwitchPortsOverview
	diags := req.Config.Get(ctx, &organizationsSwitchPortsOverview)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationSwitchPortsOverview")
		vvOrganizationID := organizationsSwitchPortsOverview.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationSwitchPortsOverviewQueryParams{}

		queryParams1.T0 = organizationsSwitchPortsOverview.T0.ValueString()
		queryParams1.T1 = organizationsSwitchPortsOverview.T1.ValueString()
		queryParams1.Timespan = organizationsSwitchPortsOverview.Timespan.ValueFloat64()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Switch.GetOrganizationSwitchPortsOverview(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationSwitchPortsOverview",
				err.Error(),
			)
			return
		}

		organizationsSwitchPortsOverview = ResponseSwitchGetOrganizationSwitchPortsOverviewItemToBody(organizationsSwitchPortsOverview, response1)
		diags = resp.State.Set(ctx, &organizationsSwitchPortsOverview)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsSwitchPortsOverview struct {
	OrganizationID types.String                                      `tfsdk:"organization_id"`
	T0             types.String                                      `tfsdk:"t0"`
	T1             types.String                                      `tfsdk:"t1"`
	Timespan       types.Float64                                     `tfsdk:"timespan"`
	Item           *ResponseSwitchGetOrganizationSwitchPortsOverview `tfsdk:"item"`
}

type ResponseSwitchGetOrganizationSwitchPortsOverview struct {
	Counts *ResponseSwitchGetOrganizationSwitchPortsOverviewCounts `tfsdk:"counts"`
}

type ResponseSwitchGetOrganizationSwitchPortsOverviewCounts struct {
	ByStatus *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatus `tfsdk:"by_status"`
	Total    types.Int64                                                     `tfsdk:"total"`
}

type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatus struct {
	Active   *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActive   `tfsdk:"active"`
	Inactive *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactive `tfsdk:"inactive"`
}

type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActive struct {
	ByMediaAndLinkSpeed *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeed `tfsdk:"by_media_and_link_speed"`
	Total               types.Int64                                                                              `tfsdk:"total"`
}

type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeed struct {
	Rj45 *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeedRj45 `tfsdk:"rj45"`
	Sfp  *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeedSfp  `tfsdk:"sfp"`
}

type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeedRj45 struct {
	Status10    types.Int64 `tfsdk:"status_10"`
	Status100   types.Int64 `tfsdk:"status_100"`
	Status1000  types.Int64 `tfsdk:"status_1000"`
	Status10000 types.Int64 `tfsdk:"status_10000"`
	Status2500  types.Int64 `tfsdk:"status_2500"`
	Status5000  types.Int64 `tfsdk:"status_5000"`
	Total       types.Int64 `tfsdk:"total"`
}

type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeedSfp struct {
	Status100    types.Int64 `tfsdk:"status_100"`
	Status1000   types.Int64 `tfsdk:"status_1000"`
	Status10000  types.Int64 `tfsdk:"status_10000"`
	Status100000 types.Int64 `tfsdk:"status_100000"`
	Status20000  types.Int64 `tfsdk:"status_20000"`
	Status25000  types.Int64 `tfsdk:"status_25000"`
	Status40000  types.Int64 `tfsdk:"status_40000"`
	Status50000  types.Int64 `tfsdk:"status_50000"`
	Total        types.Int64 `tfsdk:"total"`
}

type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactive struct {
	ByMedia *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMedia `tfsdk:"by_media"`
	Total   types.Int64                                                                    `tfsdk:"total"`
}

type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMedia struct {
	Rj45 *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMediaRj45 `tfsdk:"rj45"`
	Sfp  *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMediaSfp  `tfsdk:"sfp"`
}

type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMediaRj45 struct {
	Total types.Int64 `tfsdk:"total"`
}

type ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMediaSfp struct {
	Total types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseSwitchGetOrganizationSwitchPortsOverviewItemToBody(state OrganizationsSwitchPortsOverview, response *merakigosdk.ResponseSwitchGetOrganizationSwitchPortsOverview) OrganizationsSwitchPortsOverview {
	itemState := ResponseSwitchGetOrganizationSwitchPortsOverview{
		Counts: func() *ResponseSwitchGetOrganizationSwitchPortsOverviewCounts {
			if response.Counts != nil {
				return &ResponseSwitchGetOrganizationSwitchPortsOverviewCounts{
					ByStatus: func() *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatus {
						if response.Counts.ByStatus != nil {
							return &ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatus{
								Active: func() *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActive {
									if response.Counts.ByStatus.Active != nil {
										return &ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActive{
											ByMediaAndLinkSpeed: func() *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeed {
												if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed != nil {
													return &ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeed{
														Rj45: func() *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeedRj45 {
															if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45 != nil {
																return &ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeedRj45{
																	Status10: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Status10 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Status10))
																		}
																		return types.Int64{}
																	}(),
																	Status100: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Status100 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Status100))
																		}
																		return types.Int64{}
																	}(),
																	Status1000: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Status1000 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Status1000))
																		}
																		return types.Int64{}
																	}(),
																	Status10000: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Status10000 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Status10000))
																		}
																		return types.Int64{}
																	}(),
																	Status2500: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Status2500 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Status2500))
																		}
																		return types.Int64{}
																	}(),
																	Status5000: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Status5000 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Status5000))
																		}
																		return types.Int64{}
																	}(),
																	Total: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Total != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Rj45.Total))
																		}
																		return types.Int64{}
																	}(),
																}
															}
															return nil
														}(),
														Sfp: func() *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeedSfp {
															if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp != nil {
																return &ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusActiveByMediaAndLinkSpeedSfp{
																	Status100: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status100 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status100))
																		}
																		return types.Int64{}
																	}(),
																	Status1000: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status1000 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status1000))
																		}
																		return types.Int64{}
																	}(),
																	Status10000: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status10000 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status10000))
																		}
																		return types.Int64{}
																	}(),
																	Status100000: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status100000 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status100000))
																		}
																		return types.Int64{}
																	}(),
																	Status20000: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status20000 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status20000))
																		}
																		return types.Int64{}
																	}(),
																	Status25000: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status25000 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status25000))
																		}
																		return types.Int64{}
																	}(),
																	Status40000: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status40000 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status40000))
																		}
																		return types.Int64{}
																	}(),
																	Status50000: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status50000 != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Status50000))
																		}
																		return types.Int64{}
																	}(),
																	Total: func() types.Int64 {
																		if response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Total != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Active.ByMediaAndLinkSpeed.Sfp.Total))
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
											Total: func() types.Int64 {
												if response.Counts.ByStatus.Active.Total != nil {
													return types.Int64Value(int64(*response.Counts.ByStatus.Active.Total))
												}
												return types.Int64{}
											}(),
										}
									}
									return nil
								}(),
								Inactive: func() *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactive {
									if response.Counts.ByStatus.Inactive != nil {
										return &ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactive{
											ByMedia: func() *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMedia {
												if response.Counts.ByStatus.Inactive.ByMedia != nil {
													return &ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMedia{
														Rj45: func() *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMediaRj45 {
															if response.Counts.ByStatus.Inactive.ByMedia.Rj45 != nil {
																return &ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMediaRj45{
																	Total: func() types.Int64 {
																		if response.Counts.ByStatus.Inactive.ByMedia.Rj45.Total != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Inactive.ByMedia.Rj45.Total))
																		}
																		return types.Int64{}
																	}(),
																}
															}
															return nil
														}(),
														Sfp: func() *ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMediaSfp {
															if response.Counts.ByStatus.Inactive.ByMedia.Sfp != nil {
																return &ResponseSwitchGetOrganizationSwitchPortsOverviewCountsByStatusInactiveByMediaSfp{
																	Total: func() types.Int64 {
																		if response.Counts.ByStatus.Inactive.ByMedia.Sfp.Total != nil {
																			return types.Int64Value(int64(*response.Counts.ByStatus.Inactive.ByMedia.Sfp.Total))
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
											Total: func() types.Int64 {
												if response.Counts.ByStatus.Inactive.Total != nil {
													return types.Int64Value(int64(*response.Counts.ByStatus.Inactive.Total))
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
					Total: func() types.Int64 {
						if response.Counts.Total != nil {
							return types.Int64Value(int64(*response.Counts.Total))
						}
						return types.Int64{}
					}(),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
