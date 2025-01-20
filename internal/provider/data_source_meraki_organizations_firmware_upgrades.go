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
	_ datasource.DataSource              = &OrganizationsFirmwareUpgradesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsFirmwareUpgradesDataSource{}
)

func NewOrganizationsFirmwareUpgradesDataSource() datasource.DataSource {
	return &OrganizationsFirmwareUpgradesDataSource{}
}

type OrganizationsFirmwareUpgradesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsFirmwareUpgradesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsFirmwareUpgradesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_firmware_upgrades"
}

func (d *OrganizationsFirmwareUpgradesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
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
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. Optional parameter to filter the upgrade by product type.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"status": schema.ListAttribute{
				MarkdownDescription: `status query parameter. Optional parameter to filter the upgrade by status.`,
				Optional:            true,
				ElementType:         types.StringType,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationFirmwareUpgrades`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"completed_at": schema.StringAttribute{
							MarkdownDescription: `Timestamp when upgrade completed. Null if status pending.`,
							Computed:            true,
						},
						"from_version": schema.SingleNestedAttribute{
							MarkdownDescription: `ID of the upgrade's starting version`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"firmware": schema.StringAttribute{
									MarkdownDescription: `Firmware name`,
									Computed:            true,
								},
								"id": schema.StringAttribute{
									MarkdownDescription: `Firmware version ID`,
									Computed:            true,
								},
								"release_date": schema.StringAttribute{
									MarkdownDescription: `Release date of the firmware version`,
									Computed:            true,
								},
								"release_type": schema.StringAttribute{
									MarkdownDescription: `Release type of the firmware version`,
									Computed:            true,
								},
								"short_name": schema.StringAttribute{
									MarkdownDescription: `Firmware version short name`,
									Computed:            true,
								},
							},
						},
						"network": schema.SingleNestedAttribute{
							MarkdownDescription: `Network of the upgrade`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID of network`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The network`,
									Computed:            true,
								},
							},
						},
						"product_types": schema.StringAttribute{
							MarkdownDescription: `product upgraded [wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor]`,
							Computed:            true,
						},
						"status": schema.StringAttribute{
							MarkdownDescription: `Status of upgrade event: [Cancelled, Completed]`,
							Computed:            true,
						},
						"time": schema.StringAttribute{
							MarkdownDescription: `Scheduled start time`,
							Computed:            true,
						},
						"to_version": schema.SingleNestedAttribute{
							MarkdownDescription: `ID of the upgrade's target version`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"firmware": schema.StringAttribute{
									MarkdownDescription: `Firmware name`,
									Computed:            true,
								},
								"id": schema.StringAttribute{
									MarkdownDescription: `Firmware version ID`,
									Computed:            true,
								},
								"release_date": schema.StringAttribute{
									MarkdownDescription: `Release date of the firmware version`,
									Computed:            true,
								},
								"release_type": schema.StringAttribute{
									MarkdownDescription: `Release type of the firmware version`,
									Computed:            true,
								},
								"short_name": schema.StringAttribute{
									MarkdownDescription: `Firmware version short name`,
									Computed:            true,
								},
							},
						},
						"upgrade_batch_id": schema.StringAttribute{
							MarkdownDescription: `The upgrade batch`,
							Computed:            true,
						},
						"upgrade_id": schema.StringAttribute{
							MarkdownDescription: `The upgrade`,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsFirmwareUpgradesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsFirmwareUpgrades OrganizationsFirmwareUpgrades
	diags := req.Config.Get(ctx, &organizationsFirmwareUpgrades)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationFirmwareUpgrades")
		vvOrganizationID := organizationsFirmwareUpgrades.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationFirmwareUpgradesQueryParams{}

		queryParams1.PerPage = int(organizationsFirmwareUpgrades.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsFirmwareUpgrades.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsFirmwareUpgrades.EndingBefore.ValueString()
		queryParams1.Status = elementsToStrings(ctx, organizationsFirmwareUpgrades.Status)
		queryParams1.ProductTypes = elementsToStrings(ctx, organizationsFirmwareUpgrades.ProductTypes)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationFirmwareUpgrades(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationFirmwareUpgrades",
				err.Error(),
			)
			return
		}

		organizationsFirmwareUpgrades = ResponseOrganizationsGetOrganizationFirmwareUpgradesItemsToBody(organizationsFirmwareUpgrades, response1)
		diags = resp.State.Set(ctx, &organizationsFirmwareUpgrades)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsFirmwareUpgrades struct {
	OrganizationID types.String                                                `tfsdk:"organization_id"`
	PerPage        types.Int64                                                 `tfsdk:"per_page"`
	StartingAfter  types.String                                                `tfsdk:"starting_after"`
	EndingBefore   types.String                                                `tfsdk:"ending_before"`
	Status         types.List                                                  `tfsdk:"status"`
	ProductTypes   types.List                                                  `tfsdk:"product_types"`
	Items          *[]ResponseItemOrganizationsGetOrganizationFirmwareUpgrades `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationFirmwareUpgrades struct {
	CompletedAt    types.String                                                         `tfsdk:"completed_at"`
	FromVersion    *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesFromVersion `tfsdk:"from_version"`
	Network        *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesNetwork     `tfsdk:"network"`
	ProductTypes   types.String                                                         `tfsdk:"product_types"`
	Status         types.String                                                         `tfsdk:"status"`
	Time           types.String                                                         `tfsdk:"time"`
	ToVersion      *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesToVersion   `tfsdk:"to_version"`
	UpgradeBatchID types.String                                                         `tfsdk:"upgrade_batch_id"`
	UpgradeID      types.String                                                         `tfsdk:"upgrade_id"`
}

type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesFromVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

// ToBody
func ResponseOrganizationsGetOrganizationFirmwareUpgradesItemsToBody(state OrganizationsFirmwareUpgrades, response *merakigosdk.ResponseOrganizationsGetOrganizationFirmwareUpgrades) OrganizationsFirmwareUpgrades {
	var items []ResponseItemOrganizationsGetOrganizationFirmwareUpgrades
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationFirmwareUpgrades{
			CompletedAt: types.StringValue(item.CompletedAt),
			FromVersion: func() *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesFromVersion {
				if item.FromVersion != nil {
					return &ResponseItemOrganizationsGetOrganizationFirmwareUpgradesFromVersion{
						Firmware:    types.StringValue(item.FromVersion.Firmware),
						ID:          types.StringValue(item.FromVersion.ID),
						ReleaseDate: types.StringValue(item.FromVersion.ReleaseDate),
						ReleaseType: types.StringValue(item.FromVersion.ReleaseType),
						ShortName:   types.StringValue(item.FromVersion.ShortName),
					}
				}
				return nil
			}(),
			Network: func() *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesNetwork {
				if item.Network != nil {
					return &ResponseItemOrganizationsGetOrganizationFirmwareUpgradesNetwork{
						ID:   types.StringValue(item.Network.ID),
						Name: types.StringValue(item.Network.Name),
					}
				}
				return nil
			}(),
			ProductTypes: types.StringValue(item.ProductTypes),
			Status:       types.StringValue(item.Status),
			Time:         types.StringValue(item.Time),
			ToVersion: func() *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesToVersion {
				if item.ToVersion != nil {
					return &ResponseItemOrganizationsGetOrganizationFirmwareUpgradesToVersion{
						Firmware:    types.StringValue(item.ToVersion.Firmware),
						ID:          types.StringValue(item.ToVersion.ID),
						ReleaseDate: types.StringValue(item.ToVersion.ReleaseDate),
						ReleaseType: types.StringValue(item.ToVersion.ReleaseType),
						ShortName:   types.StringValue(item.ToVersion.ShortName),
					}
				}
				return nil
			}(),
			UpgradeBatchID: types.StringValue(item.UpgradeBatchID),
			UpgradeID:      types.StringValue(item.UpgradeID),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
