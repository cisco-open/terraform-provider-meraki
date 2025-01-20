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
	_ datasource.DataSource              = &OrganizationsAssuranceAlertsOverviewHistoricalDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAssuranceAlertsOverviewHistoricalDataSource{}
)

func NewOrganizationsAssuranceAlertsOverviewHistoricalDataSource() datasource.DataSource {
	return &OrganizationsAssuranceAlertsOverviewHistoricalDataSource{}
}

type OrganizationsAssuranceAlertsOverviewHistoricalDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAssuranceAlertsOverviewHistoricalDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAssuranceAlertsOverviewHistoricalDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_assurance_alerts_overview_historical"
}

func (d *OrganizationsAssuranceAlertsOverviewHistoricalDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"category": schema.StringAttribute{
				MarkdownDescription: `category query parameter. Optional parameter to filter by category.`,
				Optional:            true,
			},
			"device_types": schema.ListAttribute{
				MarkdownDescription: `deviceTypes query parameter. Optional parameter to filter by device types`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId query parameter. Optional parameter to filter alerts overview by network ids.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"segment_duration": schema.Int64Attribute{
				MarkdownDescription: `segmentDuration query parameter. Amount of time in seconds for each segment in the returned dataset`,
				Required:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter by primary device serial`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"severity": schema.StringAttribute{
				MarkdownDescription: `severity query parameter. Optional parameter to filter alerts overview by severity type.`,
				Optional:            true,
			},
			"ts_end": schema.StringAttribute{
				MarkdownDescription: `tsEnd query parameter. Optional parameter to filter by end timestamp defaults to the current time`,
				Optional:            true,
			},
			"ts_start": schema.StringAttribute{
				MarkdownDescription: `tsStart query parameter. Parameter to define starting timestamp of historical totals`,
				Required:            true,
			},
			"types": schema.ListAttribute{
				MarkdownDescription: `types query parameter. Optional parameter to filter by alert type.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `Historical Severity Counts`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"by_alert_type": schema.SetNestedAttribute{
									MarkdownDescription: `Totals by Type`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"critical": schema.Int64Attribute{
												MarkdownDescription: `Critical Severity Count`,
												Computed:            true,
											},
											"informational": schema.Int64Attribute{
												MarkdownDescription: `Informational Severity Count`,
												Computed:            true,
											},
											"type": schema.StringAttribute{
												MarkdownDescription: `Alert Type`,
												Computed:            true,
											},
											"warning": schema.Int64Attribute{
												MarkdownDescription: `Warning Severity Count`,
												Computed:            true,
											},
										},
									},
								},
								"segment_start": schema.StringAttribute{
									MarkdownDescription: `Starting datetime of the segment in iso8601 format`,
									Computed:            true,
								},
								"totals": schema.SingleNestedAttribute{
									MarkdownDescription: `Totals by Severity`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"critical": schema.Int64Attribute{
											MarkdownDescription: `Critical Severity Count`,
											Computed:            true,
										},
										"informational": schema.Int64Attribute{
											MarkdownDescription: `Informational Severity Count`,
											Computed:            true,
										},
										"warning": schema.Int64Attribute{
											MarkdownDescription: `Warning Severity Count`,
											Computed:            true,
										},
									},
								},
							},
						},
					},
					"meta": schema.SingleNestedAttribute{
						MarkdownDescription: `Metadata about the response`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"counts": schema.SingleNestedAttribute{
								MarkdownDescription: `Counts`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"items": schema.Int64Attribute{
										MarkdownDescription: `Total Segments`,
										Computed:            true,
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

func (d *OrganizationsAssuranceAlertsOverviewHistoricalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAssuranceAlertsOverviewHistorical OrganizationsAssuranceAlertsOverviewHistorical
	diags := req.Config.Get(ctx, &organizationsAssuranceAlertsOverviewHistorical)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAssuranceAlertsOverviewHistorical")
		vvOrganizationID := organizationsAssuranceAlertsOverviewHistorical.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationAssuranceAlertsOverviewHistoricalQueryParams{}

		queryParams1.SegmentDuration = int(organizationsAssuranceAlertsOverviewHistorical.SegmentDuration.ValueInt64())

		queryParams1.NetworkID = organizationsAssuranceAlertsOverviewHistorical.NetworkID.ValueString()
		queryParams1.Severity = organizationsAssuranceAlertsOverviewHistorical.Severity.ValueString()
		queryParams1.Types = elementsToStrings(ctx, organizationsAssuranceAlertsOverviewHistorical.Types)
		queryParams1.TsStart = organizationsAssuranceAlertsOverviewHistorical.TsStart.ValueString()

		queryParams1.TsEnd = organizationsAssuranceAlertsOverviewHistorical.TsEnd.ValueString()
		queryParams1.Category = organizationsAssuranceAlertsOverviewHistorical.Category.ValueString()
		queryParams1.Serials = elementsToStrings(ctx, organizationsAssuranceAlertsOverviewHistorical.Serials)
		queryParams1.DeviceTypes = elementsToStrings(ctx, organizationsAssuranceAlertsOverviewHistorical.DeviceTypes)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAssuranceAlertsOverviewHistorical(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAssuranceAlertsOverviewHistorical",
				err.Error(),
			)
			return
		}

		organizationsAssuranceAlertsOverviewHistorical = ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItemToBody(organizationsAssuranceAlertsOverviewHistorical, response1)
		diags = resp.State.Set(ctx, &organizationsAssuranceAlertsOverviewHistorical)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAssuranceAlertsOverviewHistorical struct {
	OrganizationID  types.String                                                           `tfsdk:"organization_id"`
	SegmentDuration types.Int64                                                            `tfsdk:"segment_duration"`
	NetworkID       types.String                                                           `tfsdk:"network_id"`
	Severity        types.String                                                           `tfsdk:"severity"`
	Types           types.List                                                             `tfsdk:"types"`
	TsStart         types.String                                                           `tfsdk:"ts_start"`
	TsEnd           types.String                                                           `tfsdk:"ts_end"`
	Category        types.String                                                           `tfsdk:"category"`
	Serials         types.List                                                             `tfsdk:"serials"`
	DeviceTypes     types.List                                                             `tfsdk:"device_types"`
	Item            *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistorical `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistorical struct {
	Items *[]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItems `tfsdk:"items"`
	Meta  *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalMeta    `tfsdk:"meta"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItems struct {
	ByAlertType  *[]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItemsByAlertType `tfsdk:"by_alert_type"`
	SegmentStart types.String                                                                             `tfsdk:"segment_start"`
	Totals       *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItemsTotals        `tfsdk:"totals"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItemsByAlertType struct {
	Critical      types.Int64  `tfsdk:"critical"`
	Informational types.Int64  `tfsdk:"informational"`
	Type          types.String `tfsdk:"type"`
	Warning       types.Int64  `tfsdk:"warning"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItemsTotals struct {
	Critical      types.Int64 `tfsdk:"critical"`
	Informational types.Int64 `tfsdk:"informational"`
	Warning       types.Int64 `tfsdk:"warning"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalMeta struct {
	Counts *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalMetaCounts `tfsdk:"counts"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalMetaCounts struct {
	Items types.Int64 `tfsdk:"items"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItemToBody(state OrganizationsAssuranceAlertsOverviewHistorical, response *merakigosdk.ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistorical) OrganizationsAssuranceAlertsOverviewHistorical {
	itemState := ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistorical{
		Items: func() *[]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItems {
			if response.Items != nil {
				result := make([]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItems{
						ByAlertType: func() *[]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItemsByAlertType {
							if items.ByAlertType != nil {
								result := make([]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItemsByAlertType, len(*items.ByAlertType))
								for i, byAlertType := range *items.ByAlertType {
									result[i] = ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItemsByAlertType{
										Critical: func() types.Int64 {
											if byAlertType.Critical != nil {
												return types.Int64Value(int64(*byAlertType.Critical))
											}
											return types.Int64{}
										}(),
										Informational: func() types.Int64 {
											if byAlertType.Informational != nil {
												return types.Int64Value(int64(*byAlertType.Informational))
											}
											return types.Int64{}
										}(),
										Type: types.StringValue(byAlertType.Type),
										Warning: func() types.Int64 {
											if byAlertType.Warning != nil {
												return types.Int64Value(int64(*byAlertType.Warning))
											}
											return types.Int64{}
										}(),
									}
								}
								return &result
							}
							return nil
						}(),
						SegmentStart: types.StringValue(items.SegmentStart),
						Totals: func() *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItemsTotals {
							if items.Totals != nil {
								return &ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalItemsTotals{
									Critical: func() types.Int64 {
										if items.Totals.Critical != nil {
											return types.Int64Value(int64(*items.Totals.Critical))
										}
										return types.Int64{}
									}(),
									Informational: func() types.Int64 {
										if items.Totals.Informational != nil {
											return types.Int64Value(int64(*items.Totals.Informational))
										}
										return types.Int64{}
									}(),
									Warning: func() types.Int64 {
										if items.Totals.Warning != nil {
											return types.Int64Value(int64(*items.Totals.Warning))
										}
										return types.Int64{}
									}(),
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
		Meta: func() *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalMeta {
			if response.Meta != nil {
				return &ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalMeta{
					Counts: func() *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewHistoricalMetaCounts{
								Items: func() types.Int64 {
									if response.Meta.Counts.Items != nil {
										return types.Int64Value(int64(*response.Meta.Counts.Items))
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
	state.Item = &itemState
	return state
}
