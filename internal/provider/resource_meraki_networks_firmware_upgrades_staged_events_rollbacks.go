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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksFirmwareUpgradesStagedEventsRollbacksResource{}
	_ resource.ResourceWithConfigure = &NetworksFirmwareUpgradesStagedEventsRollbacksResource{}
)

func NewNetworksFirmwareUpgradesStagedEventsRollbacksResource() resource.Resource {
	return &NetworksFirmwareUpgradesStagedEventsRollbacksResource{}
}

type NetworksFirmwareUpgradesStagedEventsRollbacksResource struct {
	client *merakigosdk.Client
}

func (r *NetworksFirmwareUpgradesStagedEventsRollbacksResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksFirmwareUpgradesStagedEventsRollbacksResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_firmware_upgrades_staged_events_rollbacks"
}

// resourceAction
func (r *NetworksFirmwareUpgradesStagedEventsRollbacksResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"reasons": schema.SetNestedAttribute{
						MarkdownDescription: `The reason for rolling back the staged upgrade`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"category": schema.StringAttribute{
									MarkdownDescription: `Reason for the rollback
                                              Allowed values: [broke old features,other,performance,stability,testing,unifying networks versions]`,
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"comment": schema.StringAttribute{
									MarkdownDescription: `Additional comment about the rollback`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
							},
						},
					},
					"stages": schema.SetNestedAttribute{
						MarkdownDescription: `All completed or in-progress stages in the network with their new start times. All pending stages will be canceled`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"group": schema.SingleNestedAttribute{
									MarkdownDescription: `The Staged Upgrade Group containing the name and ID`,
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"id": schema.StringAttribute{
											MarkdownDescription: `ID of the Staged Upgrade Group`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
											},
										},
									},
								},
								"milestones": schema.SingleNestedAttribute{
									MarkdownDescription: `The Staged Upgrade Milestones for the specific stage`,
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"scheduled_for": schema.StringAttribute{
											MarkdownDescription: `The start time of the staged upgrade stage. (In ISO-8601 format, in the time zone of the network.)`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
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
func (r *NetworksFirmwareUpgradesStagedEventsRollbacksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksFirmwareUpgradesStagedEventsRollbacks

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Networks.RollbacksNetworkFirmwareUpgradesStagedEvents(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing RollbacksNetworkFirmwareUpgradesStagedEvents",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing RollbacksNetworkFirmwareUpgradesStagedEvents",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksFirmwareUpgradesStagedEventsRollbacksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksFirmwareUpgradesStagedEventsRollbacksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksFirmwareUpgradesStagedEventsRollbacksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksFirmwareUpgradesStagedEventsRollbacks struct {
	NetworkID  types.String                                                   `tfsdk:"network_id"`
	Item       *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEvents  `tfsdk:"item"`
	Parameters *RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsRs `tfsdk:"parameters"`
}

type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEvents struct {
	Products *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProducts  `tfsdk:"products"`
	Reasons  *[]ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons `tfsdk:"reasons"`
	Stages   *[]ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages  `tfsdk:"stages"`
}

type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProducts struct {
	Switch *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitch `tfsdk:"switch"`
}

type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitch struct {
	NextUpgrade *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade `tfsdk:"next_upgrade"`
}

type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade struct {
	ToVersion *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion `tfsdk:"to_version"`
}

type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion struct {
	ID        types.String `tfsdk:"id"`
	ShortName types.String `tfsdk:"short_name"`
}

type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons struct {
	Category types.String `tfsdk:"category"`
	Comment  types.String `tfsdk:"comment"`
}

type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages struct {
	Group      *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup      `tfsdk:"group"`
	Milestones *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones `tfsdk:"milestones"`
	Status     types.String                                                                  `tfsdk:"status"`
}

type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup struct {
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
}

type ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones struct {
	CanceledAt   types.String `tfsdk:"canceled_at"`
	CompletedAt  types.String `tfsdk:"completed_at"`
	ScheduledFor types.String `tfsdk:"scheduled_for"`
	StartedAt    types.String `tfsdk:"started_at"`
}

type RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsRs struct {
	Reasons *[]RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasonsRs `tfsdk:"reasons"`
	Stages  *[]RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesRs  `tfsdk:"stages"`
}

type RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasonsRs struct {
	Category types.String `tfsdk:"category"`
	Comment  types.String `tfsdk:"comment"`
}

type RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesRs struct {
	Group      *RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroupRs      `tfsdk:"group"`
	Milestones *RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestonesRs `tfsdk:"milestones"`
}

type RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroupRs struct {
	ID types.String `tfsdk:"id"`
}

type RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestonesRs struct {
	ScheduledFor types.String `tfsdk:"scheduled_for"`
}

// FromBody
func (r *NetworksFirmwareUpgradesStagedEventsRollbacks) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEvents {
	re := *r.Parameters
	var requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons []merakigosdk.RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons

	if re.Reasons != nil {
		for _, rItem1 := range *re.Reasons {
			category := rItem1.Category.ValueString()
			comment := rItem1.Comment.ValueString()
			requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons = append(requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons, merakigosdk.RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons{
				Category: category,
				Comment:  comment,
			})
			//[debug] Is Array: True
		}
	}
	var requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages []merakigosdk.RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages

	if re.Stages != nil {
		for _, rItem1 := range *re.Stages {
			var requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup *merakigosdk.RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup

			if rItem1.Group != nil {
				id := rItem1.Group.ID.ValueString()
				requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup = &merakigosdk.RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup{
					ID: id,
				}
				//[debug] Is Array: False
			}
			var requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones *merakigosdk.RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones

			if rItem1.Milestones != nil {
				scheduledFor := rItem1.Milestones.ScheduledFor.ValueString()
				requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones = &merakigosdk.RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones{
					ScheduledFor: scheduledFor,
				}
				//[debug] Is Array: False
			}
			requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages = append(requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages, merakigosdk.RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages{
				Group:      requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup,
				Milestones: requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestNetworksRollbacksNetworkFirmwareUpgradesStagedEvents{
		Reasons: &requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons,
		Stages:  &requestNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages,
	}
	return &out
}

// ToBody
func ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsItemToBody(state NetworksFirmwareUpgradesStagedEventsRollbacks, response *merakigosdk.ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEvents) NetworksFirmwareUpgradesStagedEventsRollbacks {
	itemState := ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEvents{
		Products: func() *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProducts {
			if response.Products != nil {
				return &ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProducts{
					Switch: func() *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitch {
						if response.Products.Switch != nil {
							return &ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitch{
								NextUpgrade: func() *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade {
									if response.Products.Switch.NextUpgrade != nil {
										return &ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgrade{
											ToVersion: func() *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion {
												if response.Products.Switch.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersion{
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
		Reasons: func() *[]ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons {
			if response.Reasons != nil {
				result := make([]ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons, len(*response.Reasons))
				for i, reasons := range *response.Reasons {
					result[i] = ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsReasons{
						Category: types.StringValue(reasons.Category),
						Comment:  types.StringValue(reasons.Comment),
					}
				}
				return &result
			}
			return nil
		}(),
		Stages: func() *[]ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages {
			if response.Stages != nil {
				result := make([]ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages, len(*response.Stages))
				for i, stages := range *response.Stages {
					result[i] = ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStages{
						Group: func() *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup {
							if stages.Group != nil {
								return &ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesGroup{
									Description: types.StringValue(stages.Group.Description),
									ID:          types.StringValue(stages.Group.ID),
									Name:        types.StringValue(stages.Group.Name),
								}
							}
							return nil
						}(),
						Milestones: func() *ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones {
							if stages.Milestones != nil {
								return &ResponseNetworksRollbacksNetworkFirmwareUpgradesStagedEventsStagesMilestones{
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
