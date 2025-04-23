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
	_ datasource.DataSource              = &OrganizationsAssuranceAlertsOverviewDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsAssuranceAlertsOverviewDataSource{}
)

func NewOrganizationsAssuranceAlertsOverviewDataSource() datasource.DataSource {
	return &OrganizationsAssuranceAlertsOverviewDataSource{}
}

type OrganizationsAssuranceAlertsOverviewDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsAssuranceAlertsOverviewDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsAssuranceAlertsOverviewDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_assurance_alerts_overview"
}

func (d *OrganizationsAssuranceAlertsOverviewDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
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
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId query parameter. Optional parameter to filter alerts overview by network ids.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
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

					"counts": schema.SingleNestedAttribute{
						MarkdownDescription: `Counts of alerts on the organization`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"by_severity": schema.SetNestedAttribute{
								MarkdownDescription: `Counts of alerts on organization by severity`,
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"count": schema.Int64Attribute{
											MarkdownDescription: `Total count of the given severity type`,
											Computed:            true,
										},
										"type": schema.StringAttribute{
											MarkdownDescription: `Severity Type`,
											Computed:            true,
										},
									},
								},
							},
							"total": schema.Int64Attribute{
								MarkdownDescription: `Total number of alerts on the organization`,
								Computed:            true,
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsAssuranceAlertsOverviewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsAssuranceAlertsOverview OrganizationsAssuranceAlertsOverview
	diags := req.Config.Get(ctx, &organizationsAssuranceAlertsOverview)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationAssuranceAlertsOverview")
		vvOrganizationID := organizationsAssuranceAlertsOverview.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationAssuranceAlertsOverviewQueryParams{}

		queryParams1.NetworkID = organizationsAssuranceAlertsOverview.NetworkID.ValueString()
		queryParams1.Severity = organizationsAssuranceAlertsOverview.Severity.ValueString()
		queryParams1.Types = elementsToStrings(ctx, organizationsAssuranceAlertsOverview.Types)
		queryParams1.TsStart = organizationsAssuranceAlertsOverview.TsStart.ValueString()
		queryParams1.TsEnd = organizationsAssuranceAlertsOverview.TsEnd.ValueString()
		queryParams1.Category = organizationsAssuranceAlertsOverview.Category.ValueString()
		queryParams1.Serials = elementsToStrings(ctx, organizationsAssuranceAlertsOverview.Serials)
		queryParams1.DeviceTypes = elementsToStrings(ctx, organizationsAssuranceAlertsOverview.DeviceTypes)
		queryParams1.DeviceTags = elementsToStrings(ctx, organizationsAssuranceAlertsOverview.DeviceTags)
		queryParams1.Active = organizationsAssuranceAlertsOverview.Active.ValueBool()
		queryParams1.Dismissed = organizationsAssuranceAlertsOverview.Dismissed.ValueBool()
		queryParams1.Resolved = organizationsAssuranceAlertsOverview.Resolved.ValueBool()
		queryParams1.SuppressAlertsForOfflineNodes = organizationsAssuranceAlertsOverview.SuppressAlertsForOfflineNodes.ValueBool()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationAssuranceAlertsOverview(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAssuranceAlertsOverview",
				err.Error(),
			)
			return
		}

		organizationsAssuranceAlertsOverview = ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewItemToBody(organizationsAssuranceAlertsOverview, response1)
		diags = resp.State.Set(ctx, &organizationsAssuranceAlertsOverview)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsAssuranceAlertsOverview struct {
	OrganizationID                types.String                                                 `tfsdk:"organization_id"`
	NetworkID                     types.String                                                 `tfsdk:"network_id"`
	Severity                      types.String                                                 `tfsdk:"severity"`
	Types                         types.List                                                   `tfsdk:"types"`
	TsStart                       types.String                                                 `tfsdk:"ts_start"`
	TsEnd                         types.String                                                 `tfsdk:"ts_end"`
	Category                      types.String                                                 `tfsdk:"category"`
	Serials                       types.List                                                   `tfsdk:"serials"`
	DeviceTypes                   types.List                                                   `tfsdk:"device_types"`
	DeviceTags                    types.List                                                   `tfsdk:"device_tags"`
	Active                        types.Bool                                                   `tfsdk:"active"`
	Dismissed                     types.Bool                                                   `tfsdk:"dismissed"`
	Resolved                      types.Bool                                                   `tfsdk:"resolved"`
	SuppressAlertsForOfflineNodes types.Bool                                                   `tfsdk:"suppress_alerts_for_offline_nodes"`
	Item                          *ResponseOrganizationsGetOrganizationAssuranceAlertsOverview `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverview struct {
	Counts *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewCounts `tfsdk:"counts"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewCounts struct {
	BySeverity *[]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewCountsBySeverity `tfsdk:"by_severity"`
	Total      types.Int64                                                                    `tfsdk:"total"`
}

type ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewCountsBySeverity struct {
	Count types.Int64  `tfsdk:"count"`
	Type  types.String `tfsdk:"type"`
}

// ToBody
func ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewItemToBody(state OrganizationsAssuranceAlertsOverview, response *merakigosdk.ResponseOrganizationsGetOrganizationAssuranceAlertsOverview) OrganizationsAssuranceAlertsOverview {
	itemState := ResponseOrganizationsGetOrganizationAssuranceAlertsOverview{
		Counts: func() *ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewCounts {
			if response.Counts != nil {
				return &ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewCounts{
					BySeverity: func() *[]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewCountsBySeverity {
						if response.Counts.BySeverity != nil {
							result := make([]ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewCountsBySeverity, len(*response.Counts.BySeverity))
							for i, bySeverity := range *response.Counts.BySeverity {
								result[i] = ResponseOrganizationsGetOrganizationAssuranceAlertsOverviewCountsBySeverity{
									Count: func() types.Int64 {
										if bySeverity.Count != nil {
											return types.Int64Value(int64(*bySeverity.Count))
										}
										return types.Int64{}
									}(),
									Type: types.StringValue(bySeverity.Type),
								}
							}
							return &result
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
