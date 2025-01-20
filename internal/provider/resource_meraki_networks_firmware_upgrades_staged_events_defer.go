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

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksFirmwareUpgradesStagedEventsDeferResource{}
	_ resource.ResourceWithConfigure = &NetworksFirmwareUpgradesStagedEventsDeferResource{}
)

func NewNetworksFirmwareUpgradesStagedEventsDeferResource() resource.Resource {
	return &NetworksFirmwareUpgradesStagedEventsDeferResource{}
}

type NetworksFirmwareUpgradesStagedEventsDeferResource struct {
	client *merakigosdk.Client
}

func (r *NetworksFirmwareUpgradesStagedEventsDeferResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksFirmwareUpgradesStagedEventsDeferResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_firmware_upgrades_staged_events_defer"
}

// resourceAction
func (r *NetworksFirmwareUpgradesStagedEventsDeferResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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
									MarkdownDescription: `Reason for the rollback
                                                Allowed values: [broke old features,other,performance,stability,testing,unifying networks versions]`,
									Computed: true,
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
func (r *NetworksFirmwareUpgradesStagedEventsDeferResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksFirmwareUpgradesStagedEventsDefer

	var item types.Object
	resp.Diagnostics.Append(req.Plan.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	response, restyResp1, err := r.client.Networks.DeferNetworkFirmwareUpgradesStagedEvents(vvNetworkID)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing DeferNetworkFirmwareUpgradesStagedEvents",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing DeferNetworkFirmwareUpgradesStagedEvents",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksFirmwareUpgradesStagedEventsDeferResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksFirmwareUpgradesStagedEventsDeferResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksFirmwareUpgradesStagedEventsDeferResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksFirmwareUpgradesStagedEventsDefer struct {
	NetworkID  types.String                                               `tfsdk:"network_id"`
	Item       *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEvents  `tfsdk:"item"`
	Parameters *RequestNetworksDeferNetworkFirmwareUpgradesStagedEventsRs `tfsdk:"parameters"`
}

type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEvents struct {
	Products *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProducts  `tfsdk:"products"`
	Reasons  *[]ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsReasons `tfsdk:"reasons"`
	Stages   *[]ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStages  `tfsdk:"stages"`
}

type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProducts struct {
	Switch *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitch `tfsdk:"switch"`
}

type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitch struct {
	NextUpgrade *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade `tfsdk:"next_upgrade"`
}

type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade struct {
	ToVersion *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion `tfsdk:"to_version"`
}

type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion struct {
	ID        types.String `tfsdk:"id"`
	ShortName types.String `tfsdk:"short_name"`
}

type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsReasons struct {
	Category types.String `tfsdk:"category"`
	Comment  types.String `tfsdk:"comment"`
}

type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStages struct {
	Group      *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStagesGroup      `tfsdk:"group"`
	Milestones *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStagesMilestones `tfsdk:"milestones"`
	Status     types.String                                                              `tfsdk:"status"`
}

type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStagesGroup struct {
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
}

type ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStagesMilestones struct {
	CanceledAt   types.String `tfsdk:"canceled_at"`
	CompletedAt  types.String `tfsdk:"completed_at"`
	ScheduledFor types.String `tfsdk:"scheduled_for"`
	StartedAt    types.String `tfsdk:"started_at"`
}

type RequestNetworksDeferNetworkFirmwareUpgradesStagedEventsRs interface{}

// FromBody
// ToBody
func ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsItemToBody(state NetworksFirmwareUpgradesStagedEventsDefer, response *merakigosdk.ResponseNetworksDeferNetworkFirmwareUpgradesStagedEvents) NetworksFirmwareUpgradesStagedEventsDefer {
	itemState := ResponseNetworksDeferNetworkFirmwareUpgradesStagedEvents{
		Products: func() *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProducts {
			if response.Products != nil {
				return &ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProducts{
					Switch: func() *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitch {
						if response.Products.Switch != nil {
							return &ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitch{
								NextUpgrade: func() *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade {
									if response.Products.Switch.NextUpgrade != nil {
										return &ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade{
											ToVersion: func() *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion {
												if response.Products.Switch.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion{
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
		Reasons: func() *[]ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsReasons {
			if response.Reasons != nil {
				result := make([]ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsReasons, len(*response.Reasons))
				for i, reasons := range *response.Reasons {
					result[i] = ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsReasons{
						Category: types.StringValue(reasons.Category),
						Comment:  types.StringValue(reasons.Comment),
					}
				}
				return &result
			}
			return nil
		}(),
		Stages: func() *[]ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStages {
			if response.Stages != nil {
				result := make([]ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStages, len(*response.Stages))
				for i, stages := range *response.Stages {
					result[i] = ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStages{
						Group: func() *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStagesGroup {
							if stages.Group != nil {
								return &ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStagesGroup{
									Description: types.StringValue(stages.Group.Description),
									ID:          types.StringValue(stages.Group.ID),
									Name:        types.StringValue(stages.Group.Name),
								}
							}
							return nil
						}(),
						Milestones: func() *ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStagesMilestones {
							if stages.Milestones != nil {
								return &ResponseNetworksDeferNetworkFirmwareUpgradesStagedEventsStagesMilestones{
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
