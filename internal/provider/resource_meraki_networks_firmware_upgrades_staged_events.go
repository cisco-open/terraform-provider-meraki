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

// RESOURCE NORMAL
import (
	"context"
	"strconv"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksFirmwareUpgradesStagedEventsResource{}
	_ resource.ResourceWithConfigure = &NetworksFirmwareUpgradesStagedEventsResource{}
)

func NewNetworksFirmwareUpgradesStagedEventsResource() resource.Resource {
	return &NetworksFirmwareUpgradesStagedEventsResource{}
}

type NetworksFirmwareUpgradesStagedEventsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksFirmwareUpgradesStagedEventsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksFirmwareUpgradesStagedEventsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_firmware_upgrades_staged_events"
}

func (r *NetworksFirmwareUpgradesStagedEventsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"products": schema.SingleNestedAttribute{
				MarkdownDescription: `The network devices to be updated`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"switch": schema.SingleNestedAttribute{
						MarkdownDescription: `The Switch network to be updated`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"next_upgrade": schema.SingleNestedAttribute{
								MarkdownDescription: `Details of the next firmware upgrade`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"to_version": schema.SingleNestedAttribute{
										MarkdownDescription: `Details of the version the device will upgrade to`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												MarkdownDescription: `Id of the Version being upgraded to`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
													SuppressDiffString(),
												},
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
					"switch_catalyst": schema.SingleNestedAttribute{
						MarkdownDescription: `Version information for the switch network being upgraded`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"next_upgrade": schema.SingleNestedAttribute{
								MarkdownDescription: `The next upgrade version for the switch network`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"to_version": schema.SingleNestedAttribute{
										MarkdownDescription: `The version to be updated to for switch Catalyst devices`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												MarkdownDescription: `The version ID`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
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
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"group": schema.SingleNestedAttribute{
							MarkdownDescription: `The staged upgrade group`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"description": schema.StringAttribute{
									MarkdownDescription: `Description of the Staged Upgrade Group`,
									Computed:            true,
								},
								"id": schema.StringAttribute{
									MarkdownDescription: `Id of the Staged Upgrade Group`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
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
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
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
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
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
	}
}

//path params to assign NOT EDITABLE ['products']

func (r *NetworksFirmwareUpgradesStagedEventsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksFirmwareUpgradesStagedEventsRs

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
	// Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	//Has Item and not has items

	if vvNetworkID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkFirmwareUpgradesStagedEvents(vvNetworkID)
		//Has Post
		if err != nil {
			if restyResp1 != nil {
				if restyResp1.StatusCode() != 404 {
					resp.Diagnostics.AddError(
						"Failure when executing GetNetworkFirmwareUpgradesStagedEvents",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data = ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsItemToBodyRs(data, responseVerifyItem, false)
			//Path params in update assigned
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Networks.CreateNetworkFirmwareUpgradesStagedEvent(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkFirmwareUpgradesStagedEvent",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkFirmwareUpgradesStagedEvent",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Networks.GetNetworkFirmwareUpgradesStagedEvents(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFirmwareUpgradesStagedEvents",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkFirmwareUpgradesStagedEvents",
			err.Error(),
		)
		return
	}

	data = ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksFirmwareUpgradesStagedEventsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksFirmwareUpgradesStagedEventsRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	// Has Item2

	vvNetworkID := data.NetworkID.ValueString()
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkFirmwareUpgradesStagedEvents(vvNetworkID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFirmwareUpgradesStagedEvents",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkFirmwareUpgradesStagedEvents",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksFirmwareUpgradesStagedEventsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksFirmwareUpgradesStagedEventsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksFirmwareUpgradesStagedEventsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkFirmwareUpgradesStagedEvents(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkFirmwareUpgradesStagedEvents",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkFirmwareUpgradesStagedEvents",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksFirmwareUpgradesStagedEventsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksFirmwareUpgradesStagedEvents", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksFirmwareUpgradesStagedEventsRs struct {
	NetworkID types.String                                                       `tfsdk:"network_id"`
	Products  *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsRs  `tfsdk:"products"`
	Reasons   *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsReasonsRs `tfsdk:"reasons"`
	Stages    *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesRs  `tfsdk:"stages"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsRs struct {
	Switch         *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchRs `tfsdk:"switch"`
	SwitchCatalyst *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchRs `tfsdk:"switch_catalyst"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchRs struct {
	NextUpgrade *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeRs `tfsdk:"next_upgrade"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeRs struct {
	ToVersion *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersionRs `tfsdk:"to_version"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersionRs struct {
	ID        types.String `tfsdk:"id"`
	ShortName types.String `tfsdk:"short_name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsReasonsRs struct {
	Category types.String `tfsdk:"category"`
	Comment  types.String `tfsdk:"comment"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesRs struct {
	Group      *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesGroupRs      `tfsdk:"group"`
	Milestones *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesMilestonesRs `tfsdk:"milestones"`
	Status     types.String                                                              `tfsdk:"status"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesGroupRs struct {
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesMilestonesRs struct {
	CanceledAt   types.String `tfsdk:"canceled_at"`
	CompletedAt  types.String `tfsdk:"completed_at"`
	ScheduledFor types.String `tfsdk:"scheduled_for"`
	StartedAt    types.String `tfsdk:"started_at"`
}

// FromBody
func (r *NetworksFirmwareUpgradesStagedEventsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEvent {
	var requestNetworksCreateNetworkFirmwareUpgradesStagedEventProducts *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProducts

	if r.Products != nil {
		var requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitch *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitch

		if r.Products.Switch != nil {
			var requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgrade *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgrade

			if r.Products.Switch.NextUpgrade != nil {
				var requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgradeToVersion *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgradeToVersion

				if r.Products.Switch.NextUpgrade.ToVersion != nil {
					id := r.Products.Switch.NextUpgrade.ToVersion.ID.ValueString()
					requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgradeToVersion = &merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgradeToVersion{
						ID: id,
					}
					//[debug] Is Array: False
				}
				requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgrade = &merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgrade{
					ToVersion: requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgradeToVersion,
				}
				//[debug] Is Array: False
			}
			requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitch = &merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitch{
				NextUpgrade: requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchNextUpgrade,
			}
			//[debug] Is Array: False
		}
		var requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalyst *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalyst

		if r.Products.SwitchCatalyst != nil {
			var requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalystNextUpgrade *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalystNextUpgrade

			if r.Products.SwitchCatalyst.NextUpgrade != nil {
				var requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalystNextUpgradeToVersion *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalystNextUpgradeToVersion

				if r.Products.SwitchCatalyst.NextUpgrade.ToVersion != nil {
					id := r.Products.SwitchCatalyst.NextUpgrade.ToVersion.ID.ValueString()
					requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalystNextUpgradeToVersion = &merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalystNextUpgradeToVersion{
						ID: id,
					}
					//[debug] Is Array: False
				}
				requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalystNextUpgrade = &merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalystNextUpgrade{
					ToVersion: requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalystNextUpgradeToVersion,
				}
				//[debug] Is Array: False
			}
			requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalyst = &merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalyst{
				NextUpgrade: requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalystNextUpgrade,
			}
			//[debug] Is Array: False
		}
		requestNetworksCreateNetworkFirmwareUpgradesStagedEventProducts = &merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventProducts{
			Switch:         requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitch,
			SwitchCatalyst: requestNetworksCreateNetworkFirmwareUpgradesStagedEventProductsSwitchCatalyst,
		}
		//[debug] Is Array: False
	}
	var requestNetworksCreateNetworkFirmwareUpgradesStagedEventStages []merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStages

	if r.Stages != nil {
		for _, rItem1 := range *r.Stages {
			var requestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesGroup *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesGroup

			if rItem1.Group != nil {
				id := rItem1.Group.ID.ValueString()
				requestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesGroup = &merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesGroup{
					ID: id,
				}
				//[debug] Is Array: False
			}
			var requestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesMilestones *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesMilestones

			if rItem1.Milestones != nil {
				scheduledFor := rItem1.Milestones.ScheduledFor.ValueString()
				requestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesMilestones = &merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesMilestones{
					ScheduledFor: scheduledFor,
				}
				//[debug] Is Array: False
			}
			requestNetworksCreateNetworkFirmwareUpgradesStagedEventStages = append(requestNetworksCreateNetworkFirmwareUpgradesStagedEventStages, merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStages{
				Group:      requestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesGroup,
				Milestones: requestNetworksCreateNetworkFirmwareUpgradesStagedEventStagesMilestones,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEvent{
		Products: requestNetworksCreateNetworkFirmwareUpgradesStagedEventProducts,
		Stages: func() *[]merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedEventStages {
			if len(requestNetworksCreateNetworkFirmwareUpgradesStagedEventStages) > 0 {
				return &requestNetworksCreateNetworkFirmwareUpgradesStagedEventStages
			}
			return nil
		}(),
	}
	return &out
}
func (r *NetworksFirmwareUpgradesStagedEventsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedEvents {
	var requestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStages []merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStages

	if r.Stages != nil {
		for _, rItem1 := range *r.Stages {
			var requestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesGroup *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesGroup

			if rItem1.Group != nil {
				id := rItem1.Group.ID.ValueString()
				requestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesGroup = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesGroup{
					ID: id,
				}
				//[debug] Is Array: False
			}
			var requestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesMilestones *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesMilestones

			if rItem1.Milestones != nil {
				scheduledFor := rItem1.Milestones.ScheduledFor.ValueString()
				requestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesMilestones = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesMilestones{
					ScheduledFor: scheduledFor,
				}
				//[debug] Is Array: False
			}
			requestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStages = append(requestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStages, merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStages{
				Group:      requestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesGroup,
				Milestones: requestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStagesMilestones,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedEvents{
		Stages: func() *[]merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStages {
			if len(requestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStages) > 0 {
				return &requestNetworksUpdateNetworkFirmwareUpgradesStagedEventsStages
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsItemToBodyRs(state NetworksFirmwareUpgradesStagedEventsRs, response *merakigosdk.ResponseNetworksGetNetworkFirmwareUpgradesStagedEvents, is_read bool) NetworksFirmwareUpgradesStagedEventsRs {
	itemState := NetworksFirmwareUpgradesStagedEventsRs{
		Products: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsRs {
			if response.Products != nil {
				return &ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsRs{
					Switch: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchRs {
						if response.Products.Switch != nil {
							return &ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchRs{
								NextUpgrade: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeRs {
									if response.Products.Switch.NextUpgrade != nil {
										return &ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeRs{
											ToVersion: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersionRs {
												if response.Products.Switch.NextUpgrade.ToVersion != nil {
													return &ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsProductsSwitchNextUpgradeToVersionRs{
														ID: func() types.String {
															if response.Products.Switch.NextUpgrade.ToVersion.ID != "" {
																return types.StringValue(response.Products.Switch.NextUpgrade.ToVersion.ID)
															}
															return types.String{}
														}(),
														ShortName: func() types.String {
															if response.Products.Switch.NextUpgrade.ToVersion.ShortName != "" {
																return types.StringValue(response.Products.Switch.NextUpgrade.ToVersion.ShortName)
															}
															return types.String{}
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
			}
			return nil
		}(),
		Reasons: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsReasonsRs {
			if response.Reasons != nil {
				result := make([]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsReasonsRs, len(*response.Reasons))
				for i, reasons := range *response.Reasons {
					result[i] = ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsReasonsRs{
						Category: func() types.String {
							if reasons.Category != "" {
								return types.StringValue(reasons.Category)
							}
							return types.String{}
						}(),
						Comment: func() types.String {
							if reasons.Comment != "" {
								return types.StringValue(reasons.Comment)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Stages: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesRs {
			if response.Stages != nil {
				result := make([]ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesRs, len(*response.Stages))
				for i, stages := range *response.Stages {
					result[i] = ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesRs{
						Group: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesGroupRs {
							if stages.Group != nil {
								return &ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesGroupRs{
									Description: func() types.String {
										if stages.Group.Description != "" {
											return types.StringValue(stages.Group.Description)
										}
										return types.String{}
									}(),
									ID: func() types.String {
										if stages.Group.ID != "" {
											return types.StringValue(stages.Group.ID)
										}
										return types.String{}
									}(),
									Name: func() types.String {
										if stages.Group.Name != "" {
											return types.StringValue(stages.Group.Name)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
						Milestones: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesMilestonesRs {
							if stages.Milestones != nil {
								return &ResponseNetworksGetNetworkFirmwareUpgradesStagedEventsStagesMilestonesRs{
									CanceledAt: func() types.String {
										if stages.Milestones.CanceledAt != "" {
											return types.StringValue(stages.Milestones.CanceledAt)
										}
										return types.String{}
									}(),
									CompletedAt: func() types.String {
										if stages.Milestones.CompletedAt != "" {
											return types.StringValue(stages.Milestones.CompletedAt)
										}
										return types.String{}
									}(),
									ScheduledFor: func() types.String {
										if stages.Milestones.ScheduledFor != "" {
											return types.StringValue(stages.Milestones.ScheduledFor)
										}
										return types.String{}
									}(),
									StartedAt: func() types.String {
										if stages.Milestones.StartedAt != "" {
											return types.StringValue(stages.Milestones.StartedAt)
										}
										return types.String{}
									}(),
								}
							}
							return nil
						}(),
						Status: func() types.String {
							if stages.Status != "" {
								return types.StringValue(stages.Status)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksFirmwareUpgradesStagedEventsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksFirmwareUpgradesStagedEventsRs)
}
