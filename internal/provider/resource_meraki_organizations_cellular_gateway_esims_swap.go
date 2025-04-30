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
	_ resource.Resource              = &OrganizationsCellularGatewayEsimsSwapResource{}
	_ resource.ResourceWithConfigure = &OrganizationsCellularGatewayEsimsSwapResource{}
)

func NewOrganizationsCellularGatewayEsimsSwapResource() resource.Resource {
	return &OrganizationsCellularGatewayEsimsSwapResource{}
}

type OrganizationsCellularGatewayEsimsSwapResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsCellularGatewayEsimsSwapResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsCellularGatewayEsimsSwapResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_cellular_gateway_esims_swap"
}

// resourceAction
func (r *OrganizationsCellularGatewayEsimsSwapResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"eid": schema.StringAttribute{
						MarkdownDescription: `eSIM EID`,
						Computed:            true,
					},
					"iccid": schema.StringAttribute{
						MarkdownDescription: `eSIM ICCID`,
						Computed:            true,
					},
					"status": schema.StringAttribute{
						MarkdownDescription: `Swap status
                                          Allowed values: [Completed,Failed,In progress]`,
						Computed: true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"swaps": schema.SetNestedAttribute{
						MarkdownDescription: `Each object represents a swap for one eSIM`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"eid": schema.StringAttribute{
									MarkdownDescription: `eSIM EID`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"target": schema.SingleNestedAttribute{
									MarkdownDescription: `Target Profile attributes`,
									Optional:            true,
									Computed:            true,
									Attributes: map[string]schema.Attribute{

										"account_id": schema.StringAttribute{
											MarkdownDescription: `ID of the target account; can be the account currently tied to the eSIM`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
											},
										},
										"communication_plan": schema.StringAttribute{
											MarkdownDescription: `Name of the target communication plan`,
											Optional:            true,
											Computed:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.RequiresReplace(),
											},
										},
										"rate_plan": schema.StringAttribute{
											MarkdownDescription: `Name of the target rate plan`,
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
func (r *OrganizationsCellularGatewayEsimsSwapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsCellularGatewayEsimsSwap

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
	vvOrganizationID := data.OrganizationID.ValueString()
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.CellularGateway.CreateOrganizationCellularGatewayEsimsSwap(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationCellularGatewayEsimsSwap",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationCellularGatewayEsimsSwap",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsSwapItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsCellularGatewayEsimsSwapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsCellularGatewayEsimsSwapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsCellularGatewayEsimsSwapResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsCellularGatewayEsimsSwap struct {
	OrganizationID types.String                                                        `tfsdk:"organization_id"`
	Item           *ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsSwap  `tfsdk:"item"`
	Parameters     *RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapRs `tfsdk:"parameters"`
}

type ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsSwap struct {
	Eid    types.String `tfsdk:"eid"`
	Iccid  types.String `tfsdk:"iccid"`
	Status types.String `tfsdk:"status"`
}

type RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapRs struct {
	Swaps *[]RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwapsRs `tfsdk:"swaps"`
}

type RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwapsRs struct {
	Eid    types.String                                                                   `tfsdk:"eid"`
	Target *RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwapsTargetRs `tfsdk:"target"`
}

type RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwapsTargetRs struct {
	AccountID         types.String `tfsdk:"account_id"`
	CommunicationPlan types.String `tfsdk:"communication_plan"`
	RatePlan          types.String `tfsdk:"rate_plan"`
}

// FromBody
func (r *OrganizationsCellularGatewayEsimsSwap) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwap {
	re := *r.Parameters
	var requestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwaps []merakigosdk.RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwaps

	if re.Swaps != nil {
		for _, rItem1 := range *re.Swaps {
			eid := rItem1.Eid.ValueString()
			var requestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwapsTarget *merakigosdk.RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwapsTarget

			if rItem1.Target != nil {
				accountID := rItem1.Target.AccountID.ValueString()
				communicationPlan := rItem1.Target.CommunicationPlan.ValueString()
				ratePlan := rItem1.Target.RatePlan.ValueString()
				requestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwapsTarget = &merakigosdk.RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwapsTarget{
					AccountID:         accountID,
					CommunicationPlan: communicationPlan,
					RatePlan:          ratePlan,
				}
				//[debug] Is Array: False
			}
			requestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwaps = append(requestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwaps, merakigosdk.RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwaps{
				Eid:    eid,
				Target: requestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwapsTarget,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwap{
		Swaps: func() *[]merakigosdk.RequestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwaps {
			if len(requestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwaps) > 0 {
				return &requestCellularGatewayCreateOrganizationCellularGatewayEsimsSwapSwaps
			}
			return nil
		}(),
	}
	return &out
}

// ToBody
func ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsSwapItemToBody(state OrganizationsCellularGatewayEsimsSwap, response *merakigosdk.ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsSwap) OrganizationsCellularGatewayEsimsSwap {
	itemState := ResponseCellularGatewayCreateOrganizationCellularGatewayEsimsSwap{
		Eid:    types.StringValue(response.Eid),
		Iccid:  types.StringValue(response.Iccid),
		Status: types.StringValue(response.Status),
	}
	state.Item = &itemState
	return state
}
