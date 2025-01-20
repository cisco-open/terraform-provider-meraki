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
	_ datasource.DataSource              = &NetworksFirmwareUpgradesStagedEventsDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksFirmwareUpgradesStagedEventsDataSource{}
)

func NewNetworksFirmwareUpgradesStagedEventsDataSource() datasource.DataSource {
	return &NetworksFirmwareUpgradesStagedEventsDataSource{}
}

type NetworksFirmwareUpgradesStagedEventsDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksFirmwareUpgradesStagedEventsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksFirmwareUpgradesStagedEventsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_firmware_upgrades_staged_events"
}

func (d *NetworksFirmwareUpgradesStagedEventsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"products": schema.SingleNestedAttribute{
						MarkdownDescription: `The network devices to be updated`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"switch": schema.SingleNestedAttribute{
								MarkdownDescription: `The Switch network to be updated`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"next_upgrade": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the next firmware upgrade`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"to_version": schema.SingleNestedAttribute{
												MarkdownDescription: `Details of the version the device will upgrade to`,
												Computed:            true,
												Attributes: map[string]schema.Attribute{

													"id": schema.StringAttribute{
														MarkdownDescription: `Id of the Version being upgraded to`,
														Computed:            true,
													},
													"short_name": schema.StringAttribute{
														MarkdownDescription: `Firmware version short name`,
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
					"reasons": schema.SetNestedAttribute{
						MarkdownDescription: `Reasons for the rollback`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"category": schema.StringAttribute{
									MarkdownDescription: `Reason for the rollback`,
									Computed:            true,
								},
								"comment": schema.StringAttribute{
									MarkdownDescription: `Additional comment about the rollback`,
									Computed:            true,
								},
							},
						},
					},
					"stages": schema.SetNestedAttribute{
						MarkdownDescription: `The ordered stages in the network`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"group": schema.SingleNestedAttribute{
									MarkdownDescription: `The staged upgrade group`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"description": schema.StringAttribute{
											MarkdownDescription: `Description of the Staged Upgrade Group`,
											Computed:            true,
										},
										"id": schema.StringAttribute{
											MarkdownDescription: `Id of the Staged Upgrade Group`,
											Computed:            true,
										},
										"name": schema.StringAttribute{
											MarkdownDescription: `Name of the Staged Upgrade Group`,
											Computed:            true,
										},
									},
								},
								"milestones": schema.SingleNestedAttribute{
									MarkdownDescription: `The Staged Upgrade Milestones for the stage`,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"canceled_at": schema.StringAttribute{
											MarkdownDescription: `Time that the group was canceled`,
											Computed:            true,
										},
										"completed_at": schema.StringAttribute{
											MarkdownDescription: `Finish time for the group`,
											Computed:            true,
										},
										"scheduled_for": schema.StringAttribute{
											MarkdownDescription: `Scheduled start time for the group`,
											Computed:            true,
										},
										"started_at": schema.StringAttribute{
											MarkdownDescription: `Start time for the group`,
											Computed:            true,
										},
									},
								},
								"status": schema.StringAttribute{
									MarkdownDescription: `Current upgrade status of the group`,
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

func (d *NetworksFirmwareUpgradesStagedEventsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksFirmwareUpgradesStagedEvents NetworksFirmwareUpgradesStagedEvents
	diags := req.Config.Get(ctx, &networksFirmwareUpgradesStagedEvents)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkFirmwareUpgradesStagedEvents")
		vvNetworkID := networksFirmwareUpgradesStagedEvents.NetworkID.ValueString()

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkFirmwareUpgradesStagedEvents(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFirmwareUpgradesStagedEvents",
				err.Error(),
			)
			return
		}

		networksFirmwareUpgradesStagedEvents = ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsItemToBody(networksFirmwareUpgradesStagedEvents, response1)
		diags = resp.State.Set(ctx, &networksFirmwareUpgradesStagedEvents)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksFirmwareUpgradesStagedEvents struct {
	NetworkID types.String                                            `tfsdk:"network_id"`
	Item      *ResponseNetworksGetNetworkFirmwareUpgradesStagedEvents `tfsdk:"item"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEvents struct {
	Products *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProducts  `tfsdk:"products"`
	Reasons  *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsReasons `tfsdk:"reasons"`
	Stages   *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStages  `tfsdk:"stages"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProducts struct {
	Switch *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitch `tfsdk:"switch"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitch struct {
	NextUpgrade *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade `tfsdk:"next_upgrade"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade struct {
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion struct {
	ID        types.String `tfsdk:"id"`
	ShortName types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsReasons struct {
	Category types.String `tfsdk:"category"`
	Comment  types.String `tfsdk:"comment"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStages struct {
	Group      *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesGroup      `tfsdk:"group"`
	Milestones *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesMilestones `tfsdk:"milestones"`
	Status     types.String                                                            `tfsdk:"status"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesGroup struct {
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesMilestones struct {
	CanceledAt   types.String `tfsdk:"canceled_at"`
	CompletedAt  types.String `tfsdk:"completed_at"`
	ScheduledFor types.String `tfsdk:"scheduled_for"`
	StartedAt    types.String `tfsdk:"started_at"`
}

// ToBody
func ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsItemToBody(state NetworksFirmwareUpgradesStagedEvents, response *merakigosdk.ResponseNetworksGetNetworkFirmwareUpgradesStagedEvents) NetworksFirmwareUpgradesStagedEvents {
	itemState := ResponseNetworksGetNetworkFirmwareUpgradesStagedEvents{
		Products: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProducts {
			if response.Products != nil {
				return &ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProducts{
					Switch: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitch {
						if response.Products.Switch != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitch{
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade {
									if response.Products.Switch.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade{
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion {
												if response.Products.Switch.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion{
														ID:        types.StringValue(response.Products.Switch.NextUpgrade.ToVersion.ID),
														ShortName: types.StringValue(response.Products.Switch.NextUpgrade.ToVersion.ShortName),
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
			}
			return nil
		}(),
		Reasons: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsReasons {
			if response.Reasons != nil {
				result := make([]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsReasons, len(*response.Reasons))
				for i, reasons := range *response.Reasons {
					result[i] = ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsReasons{
						Category: types.StringValue(reasons.Category),
						Comment:  types.StringValue(reasons.Comment),
					}
				}
				return &result
			}
			return nil
		}(),
		Stages: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStages {
			if response.Stages != nil {
				result := make([]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStages, len(*response.Stages))
				for i, stages := range *response.Stages {
					result[i] = ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStages{
						Group: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesGroup {
							if stages.Group != nil {
								return &ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesGroup{
									Description: types.StringValue(stages.Group.Description),
									ID:          types.StringValue(stages.Group.ID),
									Name:        types.StringValue(stages.Group.Name),
								}
							}
							return nil
						}(),
						Milestones: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesMilestones {
							if stages.Milestones != nil {
								return &ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesMilestones{
									CanceledAt:   types.StringValue(stages.Milestones.CanceledAt),
									CompletedAt:  types.StringValue(stages.Milestones.CompletedAt),
									ScheduledFor: types.StringValue(stages.Milestones.ScheduledFor),
									StartedAt:    types.StringValue(stages.Milestones.StartedAt),
								}
							}
							return nil
						}(),
						Status: types.StringValue(stages.Status),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
