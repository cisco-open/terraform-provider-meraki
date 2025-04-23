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
	_ datasource.DataSource              = &OrganizationsAssuranceAlertsOverviewByTypeDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAssuranceAlertsOverviewByTypeDataSource{}
)

func NewOrganizationsAssuranceAlertsOverviewByTypeDataSource() datasource.DataSource {
	return &OrganizationsAssuranceAlertsOverviewByTypeDataSource{}
}

type OrganizationsAssuranceAlertsOverviewByTypeDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAssuranceAlertsOverviewByTypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAssuranceAlertsOverviewByTypeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_assurance_alerts_overview_by_type"
}

func (d *OrganizationsAssuranceAlertsOverviewByTypeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
				MarkdownDescription: `networkId query parameter. Optional parameter to filter alerts overview by network ids.`,
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
			"sort_by": schema.StringAttribute{
				MarkdownDescription: `sortBy query parameter. Optional parameter to set column to sort by.`,
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
						MarkdownDescription: `Organization Alert counts by type`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"count": schema.Int64Attribute{
									MarkdownDescription: `Total count of the given alert type`,
									Computed:            true,
								},
								"type": schema.StringAttribute{
									MarkdownDescription: `Alert Type`,
									Computed:            true,
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

func (d *OrganizationsAssuranceAlertsOverviewByTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAssuranceAlertsOverviewByType OrganizationsAssuranceAlertsOverviewByType
	diags := req.Config.Get(ctx, &organizationsAssuranceAlertsOverviewByType)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAssuranceAlertsOverviewByType")
		vvOrganizationID := organizationsAssuranceAlertsOverviewByType.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationAssuranceAlertsOverviewByTypeQueryParams{}

		queryParams1.PerPage = int(organizationsAssuranceAlertsOverviewByType.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsAssuranceAlertsOverviewByType.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsAssuranceAlertsOverviewByType.EndingBefore.ValueString()
		queryParams1.SortOrder = organizationsAssuranceAlertsOverviewByType.SortOrder.ValueString()
		queryParams1.NetworkID = organizationsAssuranceAlertsOverviewByType.NetworkID.ValueString()
		queryParams1.Severity = organizationsAssuranceAlertsOverviewByType.Severity.ValueString()
		queryParams1.Types = elementsToStrings(ctx, organizationsAssuranceAlertsOverviewByType.Types)
		queryParams1.TsStart = organizationsAssuranceAlertsOverviewByType.TsStart.ValueString()
		queryParams1.TsEnd = organizationsAssuranceAlertsOverviewByType.TsEnd.ValueString()
		queryParams1.Category = organizationsAssuranceAlertsOverviewByType.Category.ValueString()
		queryParams1.SortBy = organizationsAssuranceAlertsOverviewByType.SortBy.ValueString()
		queryParams1.Serials = elementsToStrings(ctx, organizationsAssuranceAlertsOverviewByType.Serials)
		queryParams1.DeviceTypes = elementsToStrings(ctx, organizationsAssuranceAlertsOverviewByType.DeviceTypes)
		queryParams1.DeviceTags = elementsToStrings(ctx, organizationsAssuranceAlertsOverviewByType.DeviceTags)
		queryParams1.Active = organizationsAssuranceAlertsOverviewByType.Active.ValueBool()
		queryParams1.Dismissed = organizationsAssuranceAlertsOverviewByType.Dismissed.ValueBool()
		queryParams1.Resolved = organizationsAssuranceAlertsOverviewByType.Resolved.ValueBool()
		queryParams1.SuppressAlertsForOfflineNodes = organizationsAssuranceAlertsOverviewByType.SuppressAlertsForOfflineNodes.ValueBool()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAssuranceAlertsOverviewByType(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAssuranceAlertsOverviewByType",
				err.Error(),
			)
			return
		}

		organizationsAssuranceAlertsOverviewByType = ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeItemToBody(organizationsAssuranceAlertsOverviewByType, response1)
		diags = resp.State.Set(ctx, &organizationsAssuranceAlertsOverviewByType)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAssuranceAlertsOverviewByType struct {
	OrganizationID                types.String                                                       `tfsdk:"organization_id"`
	PerPage                       types.Int64                                                        `tfsdk:"per_page"`
	StartingAfter                 types.String                                                       `tfsdk:"starting_after"`
	EndingBefore                  types.String                                                       `tfsdk:"ending_before"`
	SortOrder                     types.String                                                       `tfsdk:"sort_order"`
	NetworkID                     types.String                                                       `tfsdk:"network_id"`
	Severity                      types.String                                                       `tfsdk:"severity"`
	Types                         types.List                                                         `tfsdk:"types"`
	TsStart                       types.String                                                       `tfsdk:"ts_start"`
	TsEnd                         types.String                                                       `tfsdk:"ts_end"`
	Category                      types.String                                                       `tfsdk:"category"`
	SortBy                        types.String                                                       `tfsdk:"sort_by"`
	Serials                       types.List                                                         `tfsdk:"serials"`
	DeviceTypes                   types.List                                                         `tfsdk:"device_types"`
	DeviceTags                    types.List                                                         `tfsdk:"device_tags"`
	Active                        types.Bool                                                         `tfsdk:"active"`
	Dismissed                     types.Bool                                                         `tfsdk:"dismissed"`
	Resolved                      types.Bool                                                         `tfsdk:"resolved"`
	SuppressAlertsForOfflineNodes types.Bool                                                         `tfsdk:"suppress_alerts_for_offline_nodes"`
	Item                          *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByType `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByType struct {
	Items *[]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeItems `tfsdk:"items"`
	Meta  *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeMeta    `tfsdk:"meta"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeItems struct {
	Count types.Int64  `tfsdk:"count"`
	Type  types.String `tfsdk:"type"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeMeta struct {
	Counts *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeMetaCounts `tfsdk:"counts"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeMetaCounts struct {
	Items types.Int64 `tfsdk:"items"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeItemToBody(state OrganizationsAssuranceAlertsOverviewByType, response *merakigosdk.ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByType) OrganizationsAssuranceAlertsOverviewByType {
	itemState := ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByType{
		Items: func() *[]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeItems {
			if response.Items != nil {
				result := make([]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeItems{
						Count: func() types.Int64 {
							if items.Count != nil {
								return types.Int64Value(int64(*items.Count))
							}
							return types.Int64{}
						}(),
						Type: types.StringValue(items.Type),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeMeta {
			if response.Meta != nil {
				return &ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeMeta{
					Counts: func() *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewByTypeMetaCounts{
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
