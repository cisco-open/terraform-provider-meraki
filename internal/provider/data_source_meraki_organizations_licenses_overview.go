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
	_ datasource.DataSource              = &OrganizationsLicensesOverviewDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsLicensesOverviewDataSource{}
)

func NewOrganizationsLicensesOverviewDataSource() datasource.DataSource {
	return &OrganizationsLicensesOverviewDataSource{}
}

type OrganizationsLicensesOverviewDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsLicensesOverviewDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsLicensesOverviewDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_licenses_overview"
}

func (d *OrganizationsLicensesOverviewDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"expiration_date": schema.StringAttribute{
						MarkdownDescription: `License expiration date (Co-termination licensing only)`,
						Computed:            true,
					},
					"license_count": schema.Int64Attribute{
						MarkdownDescription: `Total number of licenses (Per-device licensing only)`,
						Computed:            true,
					},
					"license_types": schema.SetNestedAttribute{
						MarkdownDescription: `Data by license type (Per-device licensing only)`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"counts": schema.SingleNestedAttribute{
									MarkdownDescription: `Aggregated count data for the license type`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"unassigned": schema.Int64Attribute{
											MarkdownDescription: `The number of unassigned licenses`,
											Computed:            true,
										},
									},
								},
								"license_type": schema.StringAttribute{
									MarkdownDescription: `License type`,
									Computed:            true,
								},
							},
						},
					},
					"licensed_device_counts": schema.StringAttribute{
						//Entro en string ds
						//TODO interface
						MarkdownDescription: `License counts (Co-termination licensing only)`,
						Computed:            true,
					},
					"states": schema.SingleNestedAttribute{
						MarkdownDescription: `Aggregated data for licenses by state (Per-device licensing only)`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"active": schema.SingleNestedAttribute{
								MarkdownDescription: `Data for active licenses`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"count": schema.Int64Attribute{
										MarkdownDescription: `The number of active licenses`,
										Computed:            true,
									},
								},
							},
							"expired": schema.SingleNestedAttribute{
								MarkdownDescription: `Data for expired licenses`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"count": schema.Int64Attribute{
										MarkdownDescription: `The number of expired licenses`,
										Computed:            true,
									},
								},
							},
							"expiring": schema.SingleNestedAttribute{
								MarkdownDescription: `Data for expiring licenses`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"count": schema.Int64Attribute{
										MarkdownDescription: `The number of expiring licenses`,
										Computed:            true,
									},
									"critical": schema.SingleNestedAttribute{
										MarkdownDescription: `Data for the critical threshold`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"expiring_count": schema.Int64Attribute{
												MarkdownDescription: `The number of licenses that will expire in this window`,
												Computed:            true,
											},
											"threshold_in_days": schema.Int64Attribute{
												MarkdownDescription: `The number of days from now denoting the critical threshold for an expiring license`,
												Computed:            true,
											},
										},
									},
									"warning": schema.SingleNestedAttribute{
										MarkdownDescription: `Data for the warning threshold`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"expiring_count": schema.Int64Attribute{
												MarkdownDescription: `The number of licenses that will expire in this window`,
												Computed:            true,
											},
											"threshold_in_days": schema.Int64Attribute{
												MarkdownDescription: `The number of days from now denoting the warning threshold for an expiring license`,
												Computed:            true,
											},
										},
									},
								},
							},
							"recently_queued": schema.SingleNestedAttribute{
								MarkdownDescription: `Data for recently queued licenses`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"count": schema.Int64Attribute{
										MarkdownDescription: `The number of recently queued licenses`,
										Computed:            true,
									},
								},
							},
							"unused": schema.SingleNestedAttribute{
								MarkdownDescription: `Data for unused licenses`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"count": schema.Int64Attribute{
										MarkdownDescription: `The number of unused licenses`,
										Computed:            true,
									},
									"soonest_activation": schema.SingleNestedAttribute{
										MarkdownDescription: `Information about the soonest forthcoming license activation`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"activation_date": schema.StringAttribute{
												MarkdownDescription: `The soonest license activation date`,
												Computed:            true,
											},
											"to_activate_count": schema.Int64Attribute{
												MarkdownDescription: `The number of licenses that will activate on this date`,
												Computed:            true,
											},
										},
									},
								},
							},
							"unused_active": schema.SingleNestedAttribute{
								MarkdownDescription: `Data for unused, active licenses`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"count": schema.Int64Attribute{
										MarkdownDescription: `The number of unused, active licenses`,
										Computed:            true,
									},
									"oldest_activation": schema.SingleNestedAttribute{
										MarkdownDescription: `Information about the oldest historical license activation`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"activation_date": schema.StringAttribute{
												MarkdownDescription: `The oldest license activation date`,
												Computed:            true,
											},
											"active_count": schema.Int64Attribute{
												MarkdownDescription: `The number of licenses that activated on this date`,
												Computed:            true,
											},
										},
									},
								},
							},
						},
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `License status (Co-termination licensing only)`,
						Computed:            true,
					},
					"systems_manager": schema.SingleNestedAttribute{
						MarkdownDescription: `Aggregated data for Systems Manager licenses (Per-device licensing only)`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"counts": schema.SingleNestedAttribute{
								MarkdownDescription: `Aggregated license count data for Systems Manager`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"active_seats": schema.Int64Attribute{
										MarkdownDescription: `The number of Systems Manager seats in use`,
										Computed:            true,
									},
									"orgwide_enrolled_devices": schema.Int64Attribute{
										MarkdownDescription: `The total number of enrolled Systems Manager devices`,
										Computed:            true,
									},
									"total_seats": schema.Int64Attribute{
										MarkdownDescription: `The total number of Systems Manager seats`,
										Computed:            true,
									},
									"unassigned_seats": schema.Int64Attribute{
										MarkdownDescription: `The number of unused Systems Manager seats`,
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

func (d *OrganizationsLicensesOverviewDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsLicensesOverview OrganizationsLicensesOverview
	diags := req.Config.Get(ctx, &organizationsLicensesOverview)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationLicensesOverview")
		vvOrganizationID := organizationsLicensesOverview.OrganizationID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Organizations.GetOrganizationLicensesOverview(vvOrganizationID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationLicensesOverview",
				err.Error(),
			)
			return
		}

		organizationsLicensesOverview = ResponseOrganizationsGetOrganizationLicensesOverviewItemToBody(organizationsLicensesOverview, response1)
		diags = resp.State.Set(ctx, &organizationsLicensesOverview)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsLicensesOverview struct {
	OrganizationID types.String                                          `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsGetOrganizationLicensesOverview `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationLicensesOverview struct {
	ExpirationDate       types.String                                                              `tfsdk:"expiration_date"`
	LicenseCount         types.Int64                                                               `tfsdk:"license_count"`
	LicenseTypes         *[]ResponseOrganizationsGetOrganizationLicensesOverviewLicenseTypes       `tfsdk:"license_types"`
	LicensedDeviceCounts *ResponseOrganizationsGetOrganizationLicensesOverviewLicensedDeviceCounts `tfsdk:"licensed_device_counts"`
	States               *ResponseOrganizationsGetOrganizationLicensesOverviewStates               `tfsdk:"states"`
	Status               types.String                                                              `tfsdk:"status"`
	SystemsManager       *ResponseOrganizationsGetOrganizationLicensesOverviewSystemsManager       `tfsdk:"systems_manager"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewLicenseTypes struct {
	Counts      *ResponseOrganizationsGetOrganizationLicensesOverviewLicenseTypesCounts `tfsdk:"counts"`
	LicenseType types.String                                                            `tfsdk:"license_type"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewLicenseTypesCounts struct {
	Unassigned types.Int64 `tfsdk:"unassigned"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewLicensedDeviceCounts interface{}

type ResponseOrganizationsGetOrganizationLicensesOverviewStates struct {
	Active         *ResponseOrganizationsGetOrganizationLicensesOverviewStatesActive         `tfsdk:"active"`
	Expired        *ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpired        `tfsdk:"expired"`
	Expiring       *ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpiring       `tfsdk:"expiring"`
	RecentlyQueued *ResponseOrganizationsGetOrganizationLicensesOverviewStatesRecentlyQueued `tfsdk:"recently_queued"`
	Unused         *ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnused         `tfsdk:"unused"`
	UnusedActive   *ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnusedActive   `tfsdk:"unused_active"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewStatesActive struct {
	Count types.Int64 `tfsdk:"count"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpired struct {
	Count types.Int64 `tfsdk:"count"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpiring struct {
	Count    types.Int64                                                                 `tfsdk:"count"`
	Critical *ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpiringCritical `tfsdk:"critical"`
	Warning  *ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpiringWarning  `tfsdk:"warning"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpiringCritical struct {
	ExpiringCount   types.Int64 `tfsdk:"expiring_count"`
	ThresholdInDays types.Int64 `tfsdk:"threshold_in_days"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpiringWarning struct {
	ExpiringCount   types.Int64 `tfsdk:"expiring_count"`
	ThresholdInDays types.Int64 `tfsdk:"threshold_in_days"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewStatesRecentlyQueued struct {
	Count types.Int64 `tfsdk:"count"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnused struct {
	Count             types.Int64                                                                        `tfsdk:"count"`
	SoonestActivation *ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnusedSoonestActivation `tfsdk:"soonest_activation"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnusedSoonestActivation struct {
	ActivationDate  types.String `tfsdk:"activation_date"`
	ToActivateCount types.Int64  `tfsdk:"to_activate_count"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnusedActive struct {
	Count            types.Int64                                                                             `tfsdk:"count"`
	OldestActivation *ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnusedActiveOldestActivation `tfsdk:"oldest_activation"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnusedActiveOldestActivation struct {
	ActivationDate types.String `tfsdk:"activation_date"`
	ActiveCount    types.Int64  `tfsdk:"active_count"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewSystemsManager struct {
	Counts *ResponseOrganizationsGetOrganizationLicensesOverviewSystemsManagerCounts `tfsdk:"counts"`
}

type ResponseOrganizationsGetOrganizationLicensesOverviewSystemsManagerCounts struct {
	ActiveSeats            types.Int64 `tfsdk:"active_seats"`
	OrgwideEnrolledDevices types.Int64 `tfsdk:"orgwide_enrolled_devices"`
	TotalSeats             types.Int64 `tfsdk:"total_seats"`
	UnassignedSeats        types.Int64 `tfsdk:"unassigned_seats"`
}

// ToBody
func ResponseOrganizationsGetOrganizationLicensesOverviewItemToBody(state OrganizationsLicensesOverview, response *merakigosdk.ResponseOrganizationsGetOrganizationLicensesOverview) OrganizationsLicensesOverview {
	itemState := ResponseOrganizationsGetOrganizationLicensesOverview{
		ExpirationDate: types.StringValue(response.ExpirationDate),
		LicenseCount: func() types.Int64 {
			if response.LicenseCount != nil {
				return types.Int64Value(int64(*response.LicenseCount))
			}
			return types.Int64{}
		}(),
		LicenseTypes: func() *[]ResponseOrganizationsGetOrganizationLicensesOverviewLicenseTypes {
			if response.LicenseTypes != nil {
				result := make([]ResponseOrganizationsGetOrganizationLicensesOverviewLicenseTypes, len(*response.LicenseTypes))
				for i, licenseTypes := range *response.LicenseTypes {
					result[i] = ResponseOrganizationsGetOrganizationLicensesOverviewLicenseTypes{
						Counts: func() *ResponseOrganizationsGetOrganizationLicensesOverviewLicenseTypesCounts {
							if licenseTypes.Counts != nil {
								return &ResponseOrganizationsGetOrganizationLicensesOverviewLicenseTypesCounts{
									Unassigned: func() types.Int64 {
										if licenseTypes.Counts.Unassigned != nil {
											return types.Int64Value(int64(*licenseTypes.Counts.Unassigned))
										}
										return types.Int64{}
									}(),
								}
							}
							return nil
						}(),
						LicenseType: types.StringValue(licenseTypes.LicenseType),
					}
				}
				return &result
			}
			return nil
		}(),
		// LicensedDeviceCounts: types.StringValue(response.LicensedDeviceCounts), //TODO POSIBLE interface
		States: func() *ResponseOrganizationsGetOrganizationLicensesOverviewStates {
			if response.States != nil {
				return &ResponseOrganizationsGetOrganizationLicensesOverviewStates{
					Active: func() *ResponseOrganizationsGetOrganizationLicensesOverviewStatesActive {
						if response.States.Active != nil {
							return &ResponseOrganizationsGetOrganizationLicensesOverviewStatesActive{
								Count: func() types.Int64 {
									if response.States.Active.Count != nil {
										return types.Int64Value(int64(*response.States.Active.Count))
									}
									return types.Int64{}
								}(),
							}
						}
						return nil
					}(),
					Expired: func() *ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpired {
						if response.States.Expired != nil {
							return &ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpired{
								Count: func() types.Int64 {
									if response.States.Expired.Count != nil {
										return types.Int64Value(int64(*response.States.Expired.Count))
									}
									return types.Int64{}
								}(),
							}
						}
						return nil
					}(),
					Expiring: func() *ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpiring {
						if response.States.Expiring != nil {
							return &ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpiring{
								Count: func() types.Int64 {
									if response.States.Expiring.Count != nil {
										return types.Int64Value(int64(*response.States.Expiring.Count))
									}
									return types.Int64{}
								}(),
								Critical: func() *ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpiringCritical {
									if response.States.Expiring.Critical != nil {
										return &ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpiringCritical{
											ExpiringCount: func() types.Int64 {
												if response.States.Expiring.Critical.ExpiringCount != nil {
													return types.Int64Value(int64(*response.States.Expiring.Critical.ExpiringCount))
												}
												return types.Int64{}
											}(),
											ThresholdInDays: func() types.Int64 {
												if response.States.Expiring.Critical.ThresholdInDays != nil {
													return types.Int64Value(int64(*response.States.Expiring.Critical.ThresholdInDays))
												}
												return types.Int64{}
											}(),
										}
									}
									return nil
								}(),
								Warning: func() *ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpiringWarning {
									if response.States.Expiring.Warning != nil {
										return &ResponseOrganizationsGetOrganizationLicensesOverviewStatesExpiringWarning{
											ExpiringCount: func() types.Int64 {
												if response.States.Expiring.Warning.ExpiringCount != nil {
													return types.Int64Value(int64(*response.States.Expiring.Warning.ExpiringCount))
												}
												return types.Int64{}
											}(),
											ThresholdInDays: func() types.Int64 {
												if response.States.Expiring.Warning.ThresholdInDays != nil {
													return types.Int64Value(int64(*response.States.Expiring.Warning.ThresholdInDays))
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
					RecentlyQueued: func() *ResponseOrganizationsGetOrganizationLicensesOverviewStatesRecentlyQueued {
						if response.States.RecentlyQueued != nil {
							return &ResponseOrganizationsGetOrganizationLicensesOverviewStatesRecentlyQueued{
								Count: func() types.Int64 {
									if response.States.RecentlyQueued.Count != nil {
										return types.Int64Value(int64(*response.States.RecentlyQueued.Count))
									}
									return types.Int64{}
								}(),
							}
						}
						return nil
					}(),
					Unused: func() *ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnused {
						if response.States.Unused != nil {
							return &ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnused{
								Count: func() types.Int64 {
									if response.States.Unused.Count != nil {
										return types.Int64Value(int64(*response.States.Unused.Count))
									}
									return types.Int64{}
								}(),
								SoonestActivation: func() *ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnusedSoonestActivation {
									if response.States.Unused.SoonestActivation != nil {
										return &ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnusedSoonestActivation{
											ActivationDate: types.StringValue(response.States.Unused.SoonestActivation.ActivationDate),
											ToActivateCount: func() types.Int64 {
												if response.States.Unused.SoonestActivation.ToActivateCount != nil {
													return types.Int64Value(int64(*response.States.Unused.SoonestActivation.ToActivateCount))
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
					UnusedActive: func() *ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnusedActive {
						if response.States.UnusedActive != nil {
							return &ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnusedActive{
								Count: func() types.Int64 {
									if response.States.UnusedActive.Count != nil {
										return types.Int64Value(int64(*response.States.UnusedActive.Count))
									}
									return types.Int64{}
								}(),
								OldestActivation: func() *ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnusedActiveOldestActivation {
									if response.States.UnusedActive.OldestActivation != nil {
										return &ResponseOrganizationsGetOrganizationLicensesOverviewStatesUnusedActiveOldestActivation{
											ActivationDate: types.StringValue(response.States.UnusedActive.OldestActivation.ActivationDate),
											ActiveCount: func() types.Int64 {
												if response.States.UnusedActive.OldestActivation.ActiveCount != nil {
													return types.Int64Value(int64(*response.States.UnusedActive.OldestActivation.ActiveCount))
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
		Status: types.StringValue(response.Status),
		SystemsManager: func() *ResponseOrganizationsGetOrganizationLicensesOverviewSystemsManager {
			if response.SystemsManager != nil {
				return &ResponseOrganizationsGetOrganizationLicensesOverviewSystemsManager{
					Counts: func() *ResponseOrganizationsGetOrganizationLicensesOverviewSystemsManagerCounts {
						if response.SystemsManager.Counts != nil {
							return &ResponseOrganizationsGetOrganizationLicensesOverviewSystemsManagerCounts{
								ActiveSeats: func() types.Int64 {
									if response.SystemsManager.Counts.ActiveSeats != nil {
										return types.Int64Value(int64(*response.SystemsManager.Counts.ActiveSeats))
									}
									return types.Int64{}
								}(),
								OrgwideEnrolledDevices: func() types.Int64 {
									if response.SystemsManager.Counts.OrgwideEnrolledDevices != nil {
										return types.Int64Value(int64(*response.SystemsManager.Counts.OrgwideEnrolledDevices))
									}
									return types.Int64{}
								}(),
								TotalSeats: func() types.Int64 {
									if response.SystemsManager.Counts.TotalSeats != nil {
										return types.Int64Value(int64(*response.SystemsManager.Counts.TotalSeats))
									}
									return types.Int64{}
								}(),
								UnassignedSeats: func() types.Int64 {
									if response.SystemsManager.Counts.UnassignedSeats != nil {
										return types.Int64Value(int64(*response.SystemsManager.Counts.UnassignedSeats))
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
