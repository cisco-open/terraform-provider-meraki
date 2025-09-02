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
	_ datasource.DataSource              = &OrganizationsCameraDetectionsHistoryByBoundaryByIntervalDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsCameraDetectionsHistoryByBoundaryByIntervalDataSource{}
)

func NewOrganizationsCameraDetectionsHistoryByBoundaryByIntervalDataSource() datasource.DataSource {
	return &OrganizationsCameraDetectionsHistoryByBoundaryByIntervalDataSource{}
}

type OrganizationsCameraDetectionsHistoryByBoundaryByIntervalDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsCameraDetectionsHistoryByBoundaryByIntervalDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsCameraDetectionsHistoryByBoundaryByIntervalDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_camera_detections_history_by_boundary_by_interval"
}

func (d *OrganizationsCameraDetectionsHistoryByBoundaryByIntervalDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"boundary_ids": schema.ListAttribute{
				MarkdownDescription: `boundaryIds query parameter. A list of boundary ids. The returned cameras will be filtered to only include these ids.`,
				Required:            true,
				ElementType:         types.StringType,
			},
			"boundary_types": schema.ListAttribute{
				MarkdownDescription: `boundaryTypes query parameter. The detection types. Defaults to 'person'.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"duration": schema.Int64Attribute{
				MarkdownDescription: `duration query parameter. The minimum time, in seconds, that the person or car remains in the area to be counted. Defaults to boundary configuration or 60.`,
				Optional:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 1 1000. Defaults to 1000.`,
				Optional:            true,
			},
			"ranges": schema.ListAttribute{
				MarkdownDescription: `ranges query parameter. A list of time ranges with intervals`,
				Required:            true,
				ElementType:         types.StringType,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseCameraGetOrganizationCameraDetectionsHistoryByBoundaryByInterval`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"boundary_id": schema.StringAttribute{
							MarkdownDescription: `The boundary id`,
							Computed:            true,
						},
						"results": schema.SingleNestedAttribute{
							MarkdownDescription: `The analytics data`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"end_time": schema.StringAttribute{
									MarkdownDescription: `The period end time`,
									Computed:            true,
								},
								"in": schema.Int64Attribute{
									MarkdownDescription: `The number of detections entered`,
									Computed:            true,
								},
								"object_type": schema.StringAttribute{
									MarkdownDescription: `The detection type`,
									Computed:            true,
								},
								"out": schema.Int64Attribute{
									MarkdownDescription: `The number of detections exited`,
									Computed:            true,
								},
								"start_time": schema.StringAttribute{
									MarkdownDescription: `The period start time`,
									Computed:            true,
								},
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The boundary type`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsCameraDetectionsHistoryByBoundaryByIntervalDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsCameraDetectionsHistoryByBoundaryByInterval OrganizationsCameraDetectionsHistoryByBoundaryByInterval
	diags := req.Config.Get(ctx, &organizationsCameraDetectionsHistoryByBoundaryByInterval)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationCameraDetectionsHistoryByBoundaryByInterval")
		vvOrganizationID := organizationsCameraDetectionsHistoryByBoundaryByInterval.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationCameraDetectionsHistoryByBoundaryByIntervalQueryParams{}

		queryParams1.BoundaryIDs = elementsToStrings(ctx, organizationsCameraDetectionsHistoryByBoundaryByInterval.BoundaryIDs)

		queryParams1.Ranges = elementsToStrings(ctx, organizationsCameraDetectionsHistoryByBoundaryByInterval.Ranges)

		queryParams1.Duration = int(organizationsCameraDetectionsHistoryByBoundaryByInterval.Duration.ValueInt64())
		queryParams1.PerPage = int(organizationsCameraDetectionsHistoryByBoundaryByInterval.PerPage.ValueInt64())
		queryParams1.BoundaryTypes = elementsToStrings(ctx, organizationsCameraDetectionsHistoryByBoundaryByInterval.BoundaryTypes)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Camera.GetOrganizationCameraDetectionsHistoryByBoundaryByInterval(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationCameraDetectionsHistoryByBoundaryByInterval",
				err.Error(),
			)
			return
		}

		organizationsCameraDetectionsHistoryByBoundaryByInterval = ResponseCameraGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalItemsToBody(organizationsCameraDetectionsHistoryByBoundaryByInterval, response1)
		diags = resp.State.Set(ctx, &organizationsCameraDetectionsHistoryByBoundaryByInterval)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsCameraDetectionsHistoryByBoundaryByInterval struct {
	OrganizationID types.String                                                                    `tfsdk:"organization_id"`
	BoundaryIDs    types.List                                                                      `tfsdk:"boundary_ids"`
	Ranges         types.List                                                                      `tfsdk:"ranges"`
	Duration       types.Int64                                                                     `tfsdk:"duration"`
	PerPage        types.Int64                                                                     `tfsdk:"per_page"`
	BoundaryTypes  types.List                                                                      `tfsdk:"boundary_types"`
	Items          *[]ResponseItemCameraGetOrganizationCameraDetectionsHistoryByBoundaryByInterval `tfsdk:"items"`
}

type ResponseItemCameraGetOrganizationCameraDetectionsHistoryByBoundaryByInterval struct {
	BoundaryID types.String                                                                         `tfsdk:"boundary_id"`
	Results    *ResponseItemCameraGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalResults `tfsdk:"results"`
	Type       types.String                                                                         `tfsdk:"type"`
}

type ResponseItemCameraGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalResults struct {
	EndTime    types.String `tfsdk:"end_time"`
	In         types.Int64  `tfsdk:"in"`
	ObjectType types.String `tfsdk:"object_type"`
	Out        types.Int64  `tfsdk:"out"`
	StartTime  types.String `tfsdk:"start_time"`
}

// ToBody
func ResponseCameraGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalItemsToBody(state OrganizationsCameraDetectionsHistoryByBoundaryByInterval, response *merakigosdk.ResponseCameraGetOrganizationCameraDetectionsHistoryByBoundaryByInterval) OrganizationsCameraDetectionsHistoryByBoundaryByInterval {
	var items []ResponseItemCameraGetOrganizationCameraDetectionsHistoryByBoundaryByInterval
	for _, item := range *response {
		itemState := ResponseItemCameraGetOrganizationCameraDetectionsHistoryByBoundaryByInterval{
			BoundaryID: func() types.String {
				if item.BoundaryID != "" {
					return types.StringValue(item.BoundaryID)
				}
				return types.String{}
			}(),
			Results: func() *ResponseItemCameraGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalResults {
				if item.Results != nil {
					return &ResponseItemCameraGetOrganizationCameraDetectionsHistoryByBoundaryByIntervalResults{
						EndTime: func() types.String {
							if item.Results.EndTime != "" {
								return types.StringValue(item.Results.EndTime)
							}
							return types.String{}
						}(),
						In: func() types.Int64 {
							if item.Results.In != nil {
								return types.Int64Value(int64(*item.Results.In))
							}
							return types.Int64{}
						}(),
						ObjectType: func() types.String {
							if item.Results.ObjectType != "" {
								return types.StringValue(item.Results.ObjectType)
							}
							return types.String{}
						}(),
						Out: func() types.Int64 {
							if item.Results.Out != nil {
								return types.Int64Value(int64(*item.Results.Out))
							}
							return types.Int64{}
						}(),
						StartTime: func() types.String {
							if item.Results.StartTime != "" {
								return types.StringValue(item.Results.StartTime)
							}
							return types.String{}
						}(),
					}
				}
				return nil
			}(),
			Type: func() types.String {
				if item.Type != "" {
					return types.StringValue(item.Type)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
