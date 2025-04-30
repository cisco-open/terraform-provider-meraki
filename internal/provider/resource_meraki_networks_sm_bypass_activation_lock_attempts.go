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
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSmBypassActivationLockAttemptsResource{}
	_ resource.ResourceWithConfigure = &NetworksSmBypassActivationLockAttemptsResource{}
)

func NewNetworksSmBypassActivationLockAttemptsResource() resource.Resource {
	return &NetworksSmBypassActivationLockAttemptsResource{}
}

type NetworksSmBypassActivationLockAttemptsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSmBypassActivationLockAttemptsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSmBypassActivationLockAttemptsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_bypass_activation_lock_attempts"
}

func (r *NetworksSmBypassActivationLockAttemptsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"attempt_id": schema.StringAttribute{
				MarkdownDescription: `attemptId path parameter. Attempt ID`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"data": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"status_2090938209": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"errors": schema.SetAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
							"success": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
					"status_38290139892": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"success": schema.BoolAttribute{
								Computed: true,
							},
						},
					},
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"ids": schema.SetAttribute{
				MarkdownDescription: `The ids of the devices to attempt activation lock bypass.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"status": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (r *NetworksSmBypassActivationLockAttemptsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSmBypassActivationLockAttemptsRs

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
	vvAttemptID := data.AttemptID.ValueString()
	//Has Item and not has items

	if vvNetworkID != "" && vvAttemptID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Sm.GetNetworkSmBypassActivationLockAttempt(vvNetworkID, vvAttemptID)
		//Has Post
		if err != nil {
			if restyResp1 != nil {
				if restyResp1.StatusCode() != 404 {
					resp.Diagnostics.AddError(
						"Failure when executing GetNetworkSmBypassActivationLockAttempt",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data = ResponseSmGetNetworkSmBypassActivationLockAttemptItemToBodyRs(data, responseVerifyItem, false)
			//Path params in update assigned
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp2, err := r.client.Sm.CreateNetworkSmBypassActivationLockAttempt(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkSmBypassActivationLockAttempt",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkSmBypassActivationLockAttempt",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Sm.GetNetworkSmBypassActivationLockAttempt(vvNetworkID, vvAttemptID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSmBypassActivationLockAttempt",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSmBypassActivationLockAttempt",
			err.Error(),
		)
		return
	}

	data = ResponseSmGetNetworkSmBypassActivationLockAttemptItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksSmBypassActivationLockAttemptsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSmBypassActivationLockAttemptsRs

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
	vvAttemptID := data.AttemptID.ValueString()
	responseGet, restyRespGet, err := r.client.Sm.GetNetworkSmBypassActivationLockAttempt(vvNetworkID, vvAttemptID)
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
				"Failure when executing GetNetworkSmBypassActivationLockAttempt",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSmBypassActivationLockAttempt",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSmGetNetworkSmBypassActivationLockAttemptItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSmBypassActivationLockAttemptsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("attempt_id"), idParts[1])...)
}

func (r *NetworksSmBypassActivationLockAttemptsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSmBypassActivationLockAttemptsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update
	// No update
	resp.Diagnostics.AddError(
		"Update operation not supported in NetworksSmBypassActivationLockAttempts",
		"Update operation not supported in NetworksSmBypassActivationLockAttempts",
	)
	return
}

func (r *NetworksSmBypassActivationLockAttemptsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSmBypassActivationLockAttempts", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSmBypassActivationLockAttemptsRs struct {
	NetworkID types.String                                             `tfsdk:"network_id"`
	AttemptID types.String                                             `tfsdk:"attempt_id"`
	Data      *ResponseSmGetNetworkSmBypassActivationLockAttemptDataRs `tfsdk:"data"`
	ID        types.String                                             `tfsdk:"id"`
	Status    types.String                                             `tfsdk:"status"`
	IDs       types.Set                                                `tfsdk:"ids"`
}

type ResponseSmGetNetworkSmBypassActivationLockAttemptDataRs struct {
	Status2090938209  *ResponseSmGetNetworkSmBypassActivationLockAttemptData2090938209Rs  `tfsdk:"status_2090938209"`
	Status38290139892 *ResponseSmGetNetworkSmBypassActivationLockAttemptData38290139892Rs `tfsdk:"status_38290139892"`
}

type ResponseSmGetNetworkSmBypassActivationLockAttemptData2090938209Rs struct {
	Errors  types.Set  `tfsdk:"errors"`
	Success types.Bool `tfsdk:"success"`
}

type ResponseSmGetNetworkSmBypassActivationLockAttemptData38290139892Rs struct {
	Success types.Bool `tfsdk:"success"`
}

// FromBody
func (r *NetworksSmBypassActivationLockAttemptsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSmCreateNetworkSmBypassActivationLockAttempt {
	var iDs []string = nil
	r.IDs.ElementsAs(ctx, &iDs, false)
	out := merakigosdk.RequestSmCreateNetworkSmBypassActivationLockAttempt{
		IDs: iDs,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSmGetNetworkSmBypassActivationLockAttemptItemToBodyRs(state NetworksSmBypassActivationLockAttemptsRs, response *merakigosdk.ResponseSmGetNetworkSmBypassActivationLockAttempt, is_read bool) NetworksSmBypassActivationLockAttemptsRs {
	itemState := NetworksSmBypassActivationLockAttemptsRs{
		Data: func() *ResponseSmGetNetworkSmBypassActivationLockAttemptDataRs {
			if response.Data != nil {
				return &ResponseSmGetNetworkSmBypassActivationLockAttemptDataRs{
					Status2090938209: func() *ResponseSmGetNetworkSmBypassActivationLockAttemptData2090938209Rs {
						if response.Data.Status2090938209 != nil {
							return &ResponseSmGetNetworkSmBypassActivationLockAttemptData2090938209Rs{
								Errors: StringSliceToSet(response.Data.Status2090938209.Errors),
								Success: func() types.Bool {
									if response.Data.Status2090938209.Success != nil {
										return types.BoolValue(*response.Data.Status2090938209.Success)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
					Status38290139892: func() *ResponseSmGetNetworkSmBypassActivationLockAttemptData38290139892Rs {
						if response.Data.Status38290139892 != nil {
							return &ResponseSmGetNetworkSmBypassActivationLockAttemptData38290139892Rs{
								Success: func() types.Bool {
									if response.Data.Status38290139892.Success != nil {
										return types.BoolValue(*response.Data.Status38290139892.Success)
									}
									return types.Bool{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		ID:     types.StringValue(response.ID),
		Status: types.StringValue(response.Status),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSmBypassActivationLockAttemptsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSmBypassActivationLockAttemptsRs)
}
