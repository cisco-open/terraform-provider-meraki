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
	_ datasource.DataSource              = &OrganizationsDevicesControllerMigrationsDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsDevicesControllerMigrationsDataSource{}
)

func NewOrganizationsDevicesControllerMigrationsDataSource() datasource.DataSource {
	return &OrganizationsDevicesControllerMigrationsDataSource{}
}

type OrganizationsDevicesControllerMigrationsDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsDevicesControllerMigrationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsDevicesControllerMigrationsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_controller_migrations"
}

func (d *OrganizationsDevicesControllerMigrationsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Filter device migrations by network IDs`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 100.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. A list of Meraki Serials for which to retrieve migrations`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"target": schema.StringAttribute{
				MarkdownDescription: `target query parameter. Filter device migrations by target destination`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"items": schema.ListNestedAttribute{
						MarkdownDescription: `List of migrations for the specified devices`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"created_at": schema.StringAttribute{
									MarkdownDescription: `The time at which a migration was created`,
									Computed:            true,
								},
								"migrated_at": schema.StringAttribute{
									MarkdownDescription: `The time at which the device initiated migration`,
									Computed:            true,
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `The device serial`,
									Computed:            true,
								},
								"target": schema.StringAttribute{
									MarkdownDescription: `The migration target destination`,
									Computed:            true,
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
	}
}

func (d *OrganizationsDevicesControllerMigrationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsDevicesControllerMigrations OrganizationsDevicesControllerMigrations
	diags := req.Config.Get(ctx, &organizationsDevicesControllerMigrations)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationDevicesControllerMigrations")
		vvOrganizationID := organizationsDevicesControllerMigrations.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationDevicesControllerMigrationsQueryParams{}

		queryParams1.Serials = elementsToStrings(ctx, organizationsDevicesControllerMigrations.Serials)
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsDevicesControllerMigrations.NetworkIDs)
		queryParams1.Target = organizationsDevicesControllerMigrations.Target.ValueString()
		queryParams1.PerPage = int(organizationsDevicesControllerMigrations.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsDevicesControllerMigrations.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsDevicesControllerMigrations.EndingBefore.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationDevicesControllerMigrations(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevicesControllerMigrations",
				err.Error(),
			)
			return
		}

		organizationsDevicesControllerMigrations = ResponseOrganizationsGetOrganizationDevicesControllerMigrationsItemToBody(organizationsDevicesControllerMigrations, response1)
		diags = resp.State.Set(ctx, &organizationsDevicesControllerMigrations)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsDevicesControllerMigrations struct {
	OrganizationID types.String                                                     `tfsdk:"organization_id"`
	Serials        types.List                                                       `tfsdk:"serials"`
	NetworkIDs     types.List                                                       `tfsdk:"network_ids"`
	Target         types.String                                                     `tfsdk:"target"`
	PerPage        types.Int64                                                      `tfsdk:"per_page"`
	StartingAfter  types.String                                                     `tfsdk:"starting_after"`
	EndingBefore   types.String                                                     `tfsdk:"ending_before"`
	Item           *ResponseOrganizationsGetOrganizationDevicesControllerMigrations `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationDevicesControllerMigrations struct {
	Items *[]ResponseOrganizationsGetOrganizationDevicesControllerMigrationsItems `tfsdk:"items"`
	Meta  *ResponseOrganizationsGetOrganizationDevicesControllerMigrationsMeta    `tfsdk:"meta"`
}

type ResponseOrganizationsGetOrganizationDevicesControllerMigrationsItems struct {
	CreatedAt  types.String `tfsdk:"created_at"`
	MigratedAt types.String `tfsdk:"migrated_at"`
	Serial     types.String `tfsdk:"serial"`
	Target     types.String `tfsdk:"target"`
}

type ResponseOrganizationsGetOrganizationDevicesControllerMigrationsMeta struct {
	Counts *ResponseOrganizationsGetOrganizationDevicesControllerMigrationsMetaCounts `tfsdk:"counts"`
}

type ResponseOrganizationsGetOrganizationDevicesControllerMigrationsMetaCounts struct {
	Items *ResponseOrganizationsGetOrganizationDevicesControllerMigrationsMetaCountsItems `tfsdk:"items"`
}

type ResponseOrganizationsGetOrganizationDevicesControllerMigrationsMetaCountsItems struct {
	Remaining types.Int64 `tfsdk:"remaining"`
	Total     types.Int64 `tfsdk:"total"`
}

// ToBody
func ResponseOrganizationsGetOrganizationDevicesControllerMigrationsItemToBody(state OrganizationsDevicesControllerMigrations, response *merakigosdk.ResponseOrganizationsGetOrganizationDevicesControllerMigrations) OrganizationsDevicesControllerMigrations {
	itemState := ResponseOrganizationsGetOrganizationDevicesControllerMigrations{
		Items: func() *[]ResponseOrganizationsGetOrganizationDevicesControllerMigrationsItems {
			if response.Items != nil {
				result := make([]ResponseOrganizationsGetOrganizationDevicesControllerMigrationsItems, len(*response.Items))
				for i, items := range *response.Items {
					result[i] = ResponseOrganizationsGetOrganizationDevicesControllerMigrationsItems{
						CreatedAt: func() types.String {
							if items.CreatedAt != "" {
								return types.StringValue(items.CreatedAt)
							}
							return types.String{}
						}(),
						MigratedAt: func() types.String {
							if items.MigratedAt != "" {
								return types.StringValue(items.MigratedAt)
							}
							return types.String{}
						}(),
						Serial: func() types.String {
							if items.Serial != "" {
								return types.StringValue(items.Serial)
							}
							return types.String{}
						}(),
						Target: func() types.String {
							if items.Target != "" {
								return types.StringValue(items.Target)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Meta: func() *ResponseOrganizationsGetOrganizationDevicesControllerMigrationsMeta {
			if response.Meta != nil {
				return &ResponseOrganizationsGetOrganizationDevicesControllerMigrationsMeta{
					Counts: func() *ResponseOrganizationsGetOrganizationDevicesControllerMigrationsMetaCounts {
						if response.Meta.Counts != nil {
							return &ResponseOrganizationsGetOrganizationDevicesControllerMigrationsMetaCounts{
								Items: func() *ResponseOrganizationsGetOrganizationDevicesControllerMigrationsMetaCountsItems {
									if response.Meta.Counts.Items != nil {
										return &ResponseOrganizationsGetOrganizationDevicesControllerMigrationsMetaCountsItems{
											Remaining: func() types.Int64 {
												if response.Meta.Counts.Items.Remaining != nil {
													return types.Int64Value(int64(*response.Meta.Counts.Items.Remaining))
												}
												return types.Int64{}
											}(),
											Total: func() types.Int64 {
												if response.Meta.Counts.Items.Total != nil {
													return types.Int64Value(int64(*response.Meta.Counts.Items.Total))
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
	state.Item = &itemState
	return state
}
