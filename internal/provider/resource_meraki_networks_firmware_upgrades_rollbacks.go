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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksFirmwareUpgradesRollbacksResource{}
	_ resource.ResourceWithConfigure = &NetworksFirmwareUpgradesRollbacksResource{}
)

func NewNetworksFirmwareUpgradesRollbacksResource() resource.Resource {
	return &NetworksFirmwareUpgradesRollbacksResource{}
}

type NetworksFirmwareUpgradesRollbacksResource struct {
	client *merakigosdk.Client
}

func (r *NetworksFirmwareUpgradesRollbacksResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksFirmwareUpgradesRollbacksResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_firmware_upgrades_rollbacks"
}

// resourceAction
func (r *NetworksFirmwareUpgradesRollbacksResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"product": schema.StringAttribute{
						MarkdownDescription: `Product type to rollback (if the network is a combined network)
                                          Allowed values: [appliance,camera,cellularGateway,secureConnect,switch,switchCatalyst,wireless,wirelessController]`,
						Computed: true,
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
					"status": schema.StringAttribute{
						MarkdownDescription: `Status of the rollback
                                          Allowed values: [canceled,completed,in_progress,pending]`,
						Computed: true,
					},
					"time": schema.StringAttribute{
						MarkdownDescription: `Scheduled time for the rollback`,
						Computed:            true,
					},
					"to_version": schema.SingleNestedAttribute{
						MarkdownDescription: `Version to downgrade to (if the network has firmware flexibility)`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"firmware": schema.StringAttribute{
								MarkdownDescription: `Name of the firmware version`,
								Computed:            true,
							},
							"id": schema.StringAttribute{
								MarkdownDescription: `Firmware version identifier`,
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
						MarkdownDescription: `Batch ID of the firmware rollback`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"product": schema.StringAttribute{
						MarkdownDescription: `Product type to rollback (if the network is a combined network)
                                        Allowed values: [appliance,camera,cellularGateway,secureConnect,switch,switchCatalyst,wireless,wirelessController]`,
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"reasons": schema.SetNestedAttribute{
						MarkdownDescription: `Reasons for the rollback`,
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
					"time": schema.StringAttribute{
						MarkdownDescription: `Scheduled time for the rollback`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"to_version": schema.SingleNestedAttribute{
						MarkdownDescription: `Version to downgrade to (if the network has firmware flexibility)`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `The version ID`,
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
	}
}
func (r *NetworksFirmwareUpgradesRollbacksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksFirmwareUpgradesRollbacks

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
	response, restyResp1, err := r.client.Networks.CreateNetworkFirmwareUpgradesRollback(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkFirmwareUpgradesRollback",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkFirmwareUpgradesRollback",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksCreateNetworkFirmwareUpgradesRollbackItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksFirmwareUpgradesRollbacksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksFirmwareUpgradesRollbacksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksFirmwareUpgradesRollbacksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksFirmwareUpgradesRollbacks struct {
	NetworkID  types.String                                            `tfsdk:"network_id"`
	Item       *ResponseNetworksCreateNetworkFirmwareUpgradesRollback  `tfsdk:"item"`
	Parameters *RequestNetworksCreateNetworkFirmwareUpgradesRollbackRs `tfsdk:"parameters"`
}

type ResponseNetworksCreateNetworkFirmwareUpgradesRollback struct {
	Product        types.String                                                    `tfsdk:"product"`
	Reasons        *[]ResponseNetworksCreateNetworkFirmwareUpgradesRollbackReasons `tfsdk:"reasons"`
	Status         types.String                                                    `tfsdk:"status"`
	Time           types.String                                                    `tfsdk:"time"`
	ToVersion      *ResponseNetworksCreateNetworkFirmwareUpgradesRollbackToVersion `tfsdk:"to_version"`
	UpgradeBatchID types.String                                                    `tfsdk:"upgrade_batch_id"`
}

type ResponseNetworksCreateNetworkFirmwareUpgradesRollbackReasons struct {
	Category types.String `tfsdk:"category"`
	Comment  types.String `tfsdk:"comment"`
}

type ResponseNetworksCreateNetworkFirmwareUpgradesRollbackToVersion struct {
	Firmware    types.String `tfsdk:"firmware"`
	ID          types.String `tfsdk:"id"`
	ReleaseDate types.String `tfsdk:"release_date"`
	ReleaseType types.String `tfsdk:"release_type"`
	ShortName   types.String `tfsdk:"short_name"`
}

type RequestNetworksCreateNetworkFirmwareUpgradesRollbackRs struct {
	Product   types.String                                                     `tfsdk:"product"`
	Reasons   *[]RequestNetworksCreateNetworkFirmwareUpgradesRollbackReasonsRs `tfsdk:"reasons"`
	Time      types.String                                                     `tfsdk:"time"`
	ToVersion *RequestNetworksCreateNetworkFirmwareUpgradesRollbackToVersionRs `tfsdk:"to_version"`
}

type RequestNetworksCreateNetworkFirmwareUpgradesRollbackReasonsRs struct {
	Category types.String `tfsdk:"category"`
	Comment  types.String `tfsdk:"comment"`
}

type RequestNetworksCreateNetworkFirmwareUpgradesRollbackToVersionRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *NetworksFirmwareUpgradesRollbacks) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesRollback {
	emptyString := ""
	re := *r.Parameters
	product := new(string)
	if !re.Product.IsUnknown() && !re.Product.IsNull() {
		*product = re.Product.ValueString()
	} else {
		product = &emptyString
	}
	var requestNetworksCreateNetworkFirmwareUpgradesRollbackReasons []merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesRollbackReasons

	if re.Reasons != nil {
		for _, rItem1 := range *re.Reasons {
			category := rItem1.Category.ValueString()
			comment := rItem1.Comment.ValueString()
			requestNetworksCreateNetworkFirmwareUpgradesRollbackReasons = append(requestNetworksCreateNetworkFirmwareUpgradesRollbackReasons, merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesRollbackReasons{
				Category: category,
				Comment:  comment,
			})
			//[debug] Is Array: True
		}
	}
	time := new(string)
	if !re.Time.IsUnknown() && !re.Time.IsNull() {
		*time = re.Time.ValueString()
	} else {
		time = &emptyString
	}
	var requestNetworksCreateNetworkFirmwareUpgradesRollbackToVersion *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesRollbackToVersion

	if re.ToVersion != nil {
		id := re.ToVersion.ID.ValueString()
		requestNetworksCreateNetworkFirmwareUpgradesRollbackToVersion = &merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesRollbackToVersion{
			ID: id,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesRollback{
		Product:   *product,
		Reasons:   &requestNetworksCreateNetworkFirmwareUpgradesRollbackReasons,
		Time:      *time,
		ToVersion: requestNetworksCreateNetworkFirmwareUpgradesRollbackToVersion,
	}
	return &out
}

// ToBody
func ResponseNetworksCreateNetworkFirmwareUpgradesRollbackItemToBody(state NetworksFirmwareUpgradesRollbacks, response *merakigosdk.ResponseNetworksCreateNetworkFirmwareUpgradesRollback) NetworksFirmwareUpgradesRollbacks {
	itemState := ResponseNetworksCreateNetworkFirmwareUpgradesRollback{
		Product: types.StringValue(response.Product),
		Reasons: func() *[]ResponseNetworksCreateNetworkFirmwareUpgradesRollbackReasons {
			if response.Reasons != nil {
				result := make([]ResponseNetworksCreateNetworkFirmwareUpgradesRollbackReasons, len(*response.Reasons))
				for i, reasons := range *response.Reasons {
					result[i] = ResponseNetworksCreateNetworkFirmwareUpgradesRollbackReasons{
						Category: types.StringValue(reasons.Category),
						Comment:  types.StringValue(reasons.Comment),
					}
				}
				return &result
			}
			return nil
		}(),
		Status: types.StringValue(response.Status),
		Time:   types.StringValue(response.Time),
		ToVersion: func() *ResponseNetworksCreateNetworkFirmwareUpgradesRollbackToVersion {
			if response.ToVersion != nil {
				return &ResponseNetworksCreateNetworkFirmwareUpgradesRollbackToVersion{
					Firmware:    types.StringValue(response.ToVersion.Firmware),
					ID:          types.StringValue(response.ToVersion.ID),
					ReleaseDate: types.StringValue(response.ToVersion.ReleaseDate),
					ReleaseType: types.StringValue(response.ToVersion.ReleaseType),
					ShortName:   types.StringValue(response.ToVersion.ShortName),
				}
			}
			return nil
		}(),
		UpgradeBatchID: types.StringValue(response.UpgradeBatchID),
	}
	state.Item = &itemState
	return state
}
