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
	_ datasource.DataSource              = &OrganizationsAssuranceAlertsOverviewByNetworkDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAssuranceAlertsOverviewByNetworkDataSource{}
)

func NewOrganizationsAssuranceAlertsOverviewByNetworkDataSource() datasource.DataSource {
	return &OrganizationsAssuranceAlertsOverviewByNetworkDataSource{}
}

type OrganizationsAssuranceAlertsOverviewByNetworkDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAssuranceAlertsOverviewByNetworkDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAssuranceAlertsOverviewByNetworkDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_assurance_alerts_overview_by_network"
}

func (d *OrganizationsAssuranceAlertsOverviewByNetworkDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"active": schema.BoolAttribute{
				MarkdownDescription: `active query parameter. Optional parameter to filter by active alerts defaults to true`,
				Optional:            true,
			},
			"category": schema.StringAttribute{
				MarkdownDescription: `category query parameter. Optional parameter to filter by category.`,
				Optional:            true,
			},
			"device_tags": schema.ListAttribute{
				MarkdownDescription: `deviceTags query parameter. Optional parameter to filter by device tags`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"device_types": schema.ListAttribute{
				MarkdownDescription: `deviceTypes query parameter. Optional parameter to filter by device types`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"dismissed": schema.BoolAttribute{
				MarkdownDescription: `dismissed query parameter. Optional parameter to filter by dismissed alerts defaults to false`,
				Optional:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId query parameter. Optional parameter to filter alerts overview by network id.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"resolved": schema.BoolAttribute{
				MarkdownDescription: `resolved query parameter. Optional parameter to filter by resolved alerts defaults to false`,
				Optional:            true,
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
			"sort_order": schema.StringAttribute{
				MarkdownDescription: `sortOrder query parameter. Sorted order of entries. Order options are 'ascending' and 'descending'. Default is 'ascending'.`,
				Optional:            true,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"suppress_alerts_for_offline_nodes": schema.BoolAttribute{
				MarkdownDescription: `suppressAlertsForOfflineNodes query parameter. When set to true the api will only return connectivity alerts for a given device if that device is in an offline state. This only applies to devices. This is ignored when resolved is true. Example:  If a Switch has a VLan Mismatch and is Unreachable. only the Unreachable alert will be returned. Defaults to false.`,
				Optional:            true,
			},
			"ts_end": schema.StringAttribute{
				MarkdownDescription: `tsEnd query parameter. Optional parameter to filter by end timestamp`,
				Optional:            true,
			},
			"ts_start": schema.StringAttribute{
				MarkdownDescription: `tsStart query parameter. Optional parameter to filter by starting timestamp`,
				Optional:            true,
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
						MarkdownDescription: `Alert Counts by Network`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"alert_count": schema.Int64Attribute{
									MarkdownDescription: `Total Alerts`,
									Computed:            true,
								},
								"network_id": schema.StringAttribute{
									MarkdownDescription: `id`,
									Computed:            true,
								},
								"network_name": schema.StringAttribute{
									MarkdownDescription: `Name`,
									Computed:            true,
								},
								"severity_counts": schema.SetNestedAttribute{
									MarkdownDescription: `Alerts By Severity`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"count": schema.Int64Attribute{
												MarkdownDescription: `Count`,
												Computed:            true,
											},
											"type": schema.StringAttribute{
												MarkdownDescription: `Type`,
												Computed:            true,
											},
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
										MarkdownDescription: `Total Alerts`,
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

func (d *OrganizationsAssuranceAlertsOverviewByNetworkDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAssuranceAlertsOverviewByNetwork OrganizationsAssuranceAlertsOverviewByNetwork
	diags := req.Config.Get(ctx, &organizationsAssuranceAlertsOverviewByNetwork)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAssuranceAlertsOverviewByNetwork")
		vvOrganizationID := organizationsAssuranceAlertsOverviewByNetwork.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationAssuranceAlertsOverviewByNetworkQueryParams{}

		queryParams1.PerPage = int(organizationsAssuranceAlertsOverviewByNetwork.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsAssuranceAlertsOverviewByNetwork.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsAssuranceAlertsOverviewByNetwork.EndingBefore.ValueString()
		queryParams1.SortOrder = organizationsAssuranceAlertsOverviewByNetwork.SortOrder.ValueString()
		queryParams1.NetworkID = organizationsAssuranceAlertsOverviewByNetwork.NetworkID.ValueString()
		queryParams1.Severity = organizationsAssuranceAlertsOverviewByNetwork.Severity.ValueString()
		queryParams1.Types = elementsToStrings(ctx, organizationsAssuranceAlertsOverviewByNetwork.Types)
		queryParams1.TsStart = organizationsAssuranceAlertsOverviewByNetwork.TsStart.ValueString()
		queryParams1.TsEnd = organizationsAssuranceAlertsOverviewByNetwork.TsEnd.ValueString()
		queryParams1.Category = organizationsAssuranceAlertsOverviewByNetwork.Category.ValueString()
		queryParams1.Serials = elementsToStrings(ctx, organizationsAssuranceAlertsOverviewByNetwork.Serials)
		queryParams1.DeviceTypes = elementsToStrings(ctx, organizationsAssuranceAlertsOverviewByNetwork.DeviceTypes)
		queryParams1.DeviceTags = elementsToStrings(ctx, organizationsAssuranceAlertsOverviewByNetwork.DeviceTags)
		queryParams1.Active = organizationsAssuranceAlertsOverviewByNetwork.Active.ValueBool()
		queryParams1.Dismissed = organizationsAssuranceAlertsOverviewByNetwork.Dismissed.ValueBool()
		queryParams1.Resolved = organizationsAssuranceAlertsOverviewByNetwork.Resolved.ValueBool()
		queryParams1.SuppressAlertsForOfflineNodes = organizationsAssuranceAlertsOverviewByNetwork.SuppressAlertsForOfflineNodes.ValueBool()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAssuranceAlertsOverviewByNetwork(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAssuranceAlertsOverviewByNetwork",
				err.Error(),
			)
			return
		}

		organizationsAssuranceAlertsOverviewByNetwork = ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkItemToBody(organizationsAssuranceAlertsOverviewByNetwork, response1)
		diags = resp.State.Set(ctx, &organizationsAssuranceAlertsOverviewByNetwork)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAssuranceAlertsOverviewByNetwork struct {
	OrganizationID                types.String                                                          `tfsdk:"organization_id"`
	PerPage                       types.Int64                                                           `tfsdk:"per_page"`
	StartingAfter                 types.String                                                          `tfsdk:"starting_after"`
	EndingBefore                  types.String                                                          `tfsdk:"ending_before"`
	SortOrder                     types.String                                                          `tfsdk:"sort_order"`
	NetworkID                     types.String                                                          `tfsdk:"network_id"`
	Severity                      types.String                                                          `tfsdk:"severity"`
	Types                         types.List                                                            `tfsdk:"types"`
	TsStart                       types.String                                                          `tfsdk:"ts_start"`
	TsEnd                         types.String                                                          `tfsdk:"ts_end"`
	Category                      types.String                                                          `tfsdk:"category"`
	Serials                       types.List                                                            `tfsdk:"serials"`
	DeviceTypes                   types.List                                                            `tfsdk:"device_types"`
	DeviceTags                    types.List                                                            `tfsdk:"device_tags"`
	Active                        types.Bool                                                            `tfsdk:"active"`
	Dismissed                     types.Bool                                                            `tfsdk:"dismissed"`
	Resolved                      types.Bool                                                            `tfsdk:"resolved"`
	SuppressAlertsForOfflineNodes types.Bool                                                            `tfsdk:"suppress_alerts_for_offline_nodes"`
	Item                          *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetwork `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetwork struct {
	Items *[]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkItems `tfsdk:"items"`
	Meta  *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkMeta    `tfsdk:"meta"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkItems struct {
	AlertCount     types.Int64                                                                                `tfsdk:"alert_count"`
	NetworkID      types.String                                                                               `tfsdk:"network_id"`
	NetworkName    types.String                                                                               `tfsdk:"network_name"`
	SeverityCounts *[]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkItemsSeverityCounts `tfsdk:"severity_counts"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkItemsSeverityCounts struct {
	Count types.Int64  `tfsdk:"count"`
	Type  types.String `tfsdk:"type"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkMeta struct {
	Counts *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkMetaCounts `tfsdk:"counts"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkMetaCounts struct {
	Items types.Int64 `tfsdk:"items"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkItemToBody(state OrganizationsAssuranceAlertsOverviewByNetwork, response *merakigosdk.ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetwork) OrganizationsAssuranceAlertsOverviewByNetwork {
	itemState := ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetwork{
		Items: func() *[]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkItems {
			if response.Items != nil {
				result := make([]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkItems{
						AlertCount: func() types.Int64 {
							if items.AlertCount != nil {
								return types.Int64Value(int64(*items.AlertCount))
							}
							return types.Int64{}
						}(),
						NetworkID:   types.StringValue(items.NetworkID),
						NetworkName: types.StringValue(items.NetworkName),
						SeverityCounts: func() *[]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkItemsSeverityCounts {
							if items.SeverityCounts != nil {
								result := make([]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkItemsSeverityCounts, len(*items.SeverityCounts))
								for i, severityCounts := range *items.SeverityCounts {
									result[i] = ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkItemsSeverityCounts{
										Count: func() types.Int64 {
											if severityCounts.Count != nil {
												return types.Int64Value(int64(*severityCounts.Count))
											}
											return types.Int64{}
										}(),
										Type: types.StringValue(severityCounts.Type),
									}
								}
								return &result
							}
							return nil
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkMeta {
			if response.Meta != nil {
				return &ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkMeta{
					Counts: func() *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByNetworkMetaCounts{
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
