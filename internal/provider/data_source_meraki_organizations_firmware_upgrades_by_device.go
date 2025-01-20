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
	_ datasource.DataSource              = &OrganizationsFirmwareUpgradesByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsFirmwareUpgradesByDeviceDataSource{}
)

func NewOrganizationsFirmwareUpgradesByDeviceDataSource() datasource.DataSource {
	return &OrganizationsFirmwareUpgradesByDeviceDataSource{}
}

type OrganizationsFirmwareUpgradesByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsFirmwareUpgradesByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsFirmwareUpgradesByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_firmware_upgrades_by_device"
}

func (d *OrganizationsFirmwareUpgradesByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"current_upgrades_only": schema.BoolAttribute{
				MarkdownDescription: `currentUpgradesOnly query parameter. Optional parameter to filter to only current or pending upgrade statuses`,
				Optional:            true,
			},
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"firmware_upgrade_batch_ids": schema.ListAttribute{
				MarkdownDescription: `firmwareUpgradeBatchIds query parameter. Optional parameter to filter by firmware upgrade batch ids.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"macs": schema.ListAttribute{
				MarkdownDescription: `macs query parameter. Optional parameter to filter by one or more MAC addresses belonging to devices. All devices returned belong to MAC addresses that are an exact match.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"network_ids": schema.ListAttribute{
				MarkdownDescription: `networkIds query parameter. Optional parameter to filter by network`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 50.`,
				Optional:            true,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter by serial number.  All returned devices will have a serial number that is an exact match.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"upgradestatuses": schema.ListAttribute{
				MarkdownDescription: `upgradeStatuses query parameter. Optional parameter to filter by firmware upgrade statuses.`,
				Optional:            true,
				ElementType:         types.StringType,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseOrganizationsGetOrganizationFirmwareUpgradesByDevice`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"device_status": schema.StringAttribute{
							MarkdownDescription: `Status of the device upgrade`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name assigned to the device`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial of the device`,
							Computed:            true,
						},
						"upgrade": schema.SingleNestedAttribute{
							MarkdownDescription: `The devices upgrade details and status`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"from_version": schema.SingleNestedAttribute{
									MarkdownDescription: `The initial version of the device`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the initial firmware version`,
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
								"id": schema.StringAttribute{
									MarkdownDescription: `ID of the upgrade`,
									Computed:            true,
								},
								"staged": schema.SingleNestedAttribute{
									MarkdownDescription: `Staged upgrade`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"group": schema.SingleNestedAttribute{
											MarkdownDescription: `The staged upgrade group`,
											Computed:            true,
											Attributes: map[string]schema.Attribute{

												"id": schema.StringAttribute{
													MarkdownDescription: `Id of the staged upgrade group`,
													Computed:            true,
												},
											},
										},
									},
								},
								"status": schema.StringAttribute{
									MarkdownDescription: `Status of the upgrade`,
									Computed:            true,
								},
								"time": schema.StringAttribute{
									MarkdownDescription: `Start time of the upgrade`,
									Computed:            true,
								},
								"to_version": schema.SingleNestedAttribute{
									MarkdownDescription: `Version the device is upgrading to`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the initial firmware version`,
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
									MarkdownDescription: `ID of the upgrade batch`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *OrganizationsFirmwareUpgradesByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsFirmwareUpgradesByDevice OrganizationsFirmwareUpgradesByDevice
	diags := req.Config.Get(ctx, &organizationsFirmwareUpgradesByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationFirmwareUpgradesByDevice")
		vvOrganizationID := organizationsFirmwareUpgradesByDevice.OrganizationID.ValueString()
		queryParams1 := merakigosdk.GetOrganizationFirmwareUpgradesByDeviceQueryParams{}

		queryParams1.PerPage = int(organizationsFirmwareUpgradesByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = organizationsFirmwareUpgradesByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = organizationsFirmwareUpgradesByDevice.EndingBefore.ValueString()
		queryParams1.NetworkIDs = elementsToStrings(ctx, organizationsFirmwareUpgradesByDevice.NetworkIDs)
		queryParams1.Serials = elementsToStrings(ctx, organizationsFirmwareUpgradesByDevice.Serials)
		queryParams1.Macs = elementsToStrings(ctx, organizationsFirmwareUpgradesByDevice.Macs)
		queryParams1.FirmwareUpgradeBatchIDs = elementsToStrings(ctx, organizationsFirmwareUpgradesByDevice.FirmwareUpgradeBatchIDs)
		queryParams1.Upgradestatuses = elementsToStrings(ctx, organizationsFirmwareUpgradesByDevice.Upgradestatuses)
		queryParams1.CurrentUpgradesOnly = organizationsFirmwareUpgradesByDevice.CurrentUpgradesOnly.ValueBool()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationFirmwareUpgradesByDevice(vvOrganizationID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationFirmwareUpgradesByDevice",
				err.Error(),
			)
			return
		}

		organizationsFirmwareUpgradesByDevice = ResponseOrganizationsGetOrganizationFirmwareUpgradesByDeviceItemsToBody(organizationsFirmwareUpgradesByDevice, response1)
		diags = resp.State.Set(ctx, &organizationsFirmwareUpgradesByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsFirmwareUpgradesByDevice struct {
	OrganizationID          types.String                                                        `tfsdk:"organization_id"`
	PerPage                 types.Int64                                                         `tfsdk:"per_page"`
	StartingAfter           types.String                                                        `tfsdk:"starting_after"`
	EndingBefore            types.String                                                        `tfsdk:"ending_before"`
	NetworkIDs              types.List                                                          `tfsdk:"network_ids"`
	Serials                 types.List                                                          `tfsdk:"serials"`
	Macs                    types.List                                                          `tfsdk:"macs"`
	FirmwareUpgradeBatchIDs types.List                                                          `tfsdk:"firmware_upgrade_batch_ids"`
	Upgradestatuses         types.List                                                          `tfsdk:"upgradestatuses"`
	CurrentUpgradesOnly     types.Bool                                                          `tfsdk:"current_upgrades_only"`
	Items                   *[]ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDevice `tfsdk:"items"`
}

type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDevice struct {
	DeviceStatus types.String                                                             `tfsdk:"device_status"`
	Name         types.String                                                             `tfsdk:"name"`
	Serial       types.String                                                             `tfsdk:"serial"`
	Upgrade      *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgrade `tfsdk:"upgrade"`
}

type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgrade struct {
	FromVersion    *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradeFromVersion `tfsdk:"from_version"`
	ID             types.String                                                                        `tfsdk:"id"`
	Staged         *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradestaged      `tfsdk:"staged"`
	Status         types.String                                                                        `tfsdk:"status"`
	Time           types.String                                                                        `tfsdk:"time"`
	ToVersion      *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradeToVersion   `tfsdk:"to_version"`
	UpgradeBatchID types.String                                                                        `tfsdk:"upgrade_batch_id"`
}

type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradeFromVersion struct {
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradestaged struct {
	Group *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradestagedGroup `tfsdk:"group"`
}

type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradestagedGroup struct {
	ID types.String `tfsdk:"id"`
}

type ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradeToVersion struct {
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

// ToBody
func ResponseOrganizationsGetOrganizationFirmwareUpgradesByDeviceItemsToBody(state OrganizationsFirmwareUpgradesByDevice, response *merakigosdk.ResponseOrganizationsGetOrganizationFirmwareUpgradesByDevice) OrganizationsFirmwareUpgradesByDevice {
	var items []ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDevice
	for _, item := range *response {
		itemState := ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDevice{
			DeviceStatus: types.StringValue(item.DeviceStatus),
			Name:         types.StringValue(item.Name),
			Serial:       types.StringValue(item.Serial),
			Upgrade: func() *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgrade {
				if item.Upgrade != nil {
					return &ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgrade{
						FromVersion: func() *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradeFromVersion {
							if item.Upgrade.FromVersion != nil {
								return &ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradeFromVersion{
									ID:          types.StringValue(item.Upgrade.FromVersion.ID),
									ReleaseDate: types.StringValue(item.Upgrade.FromVersion.ReleaseDate),
									ReleaseType: types.StringValue(item.Upgrade.FromVersion.ReleaseType),
									ShortName:   types.StringValue(item.Upgrade.FromVersion.ShortName),
								}
							}
							return nil
						}(),
						ID: types.StringValue(item.Upgrade.ID),
						Staged: func() *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradestaged {
							if item.Upgrade.Staged != nil {
								return &ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradestaged{
									Group: func() *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradestagedGroup {
										if item.Upgrade.Staged.Group != nil {
											return &ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradestagedGroup{
												ID: types.StringValue(item.Upgrade.Staged.Group.ID),
											}
										}
										return nil
									}(),
								}
							}
							return nil
						}(),
						Status: types.StringValue(item.Upgrade.Status),
						Time:   types.StringValue(item.Upgrade.Time),
						ToVersion: func() *ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradeToVersion {
							if item.Upgrade.ToVersion != nil {
								return &ResponseItemOrganizationsGetOrganizationFirmwareUpgradesByDeviceUpgradeToVersion{
									ID:          types.StringValue(item.Upgrade.ToVersion.ID),
									ReleaseDate: types.StringValue(item.Upgrade.ToVersion.ReleaseDate),
									ReleaseType: types.StringValue(item.Upgrade.ToVersion.ReleaseType),
									ShortName:   types.StringValue(item.Upgrade.ToVersion.ShortName),
								}
							}
							return nil
						}(),
						UpgradeBatchID: types.StringValue(item.Upgrade.UpgradeBatchID),
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
