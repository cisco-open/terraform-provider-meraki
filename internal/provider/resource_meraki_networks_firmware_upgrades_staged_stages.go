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

	merakigosdk "dashboard-api-go/sdk"

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
	_ resource.Resource              = &NetworksFirmwareUpgradesStagedStagesResource{}
	_ resource.ResourceWithConfigure = &NetworksFirmwareUpgradesStagedStagesResource{}
)

func NewNetworksFirmwareUpgradesStagedStagesResource() resource.Resource {
	return &NetworksFirmwareUpgradesStagedStagesResource{}
}

type NetworksFirmwareUpgradesStagedStagesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksFirmwareUpgradesStagedStagesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksFirmwareUpgradesStagedStagesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_firmware_upgrades_staged_stages"
}

func (r *NetworksFirmwareUpgradesStagedStagesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"json": schema.SetNestedAttribute{
				MarkdownDescription: `Array of Staged Upgrade Groups`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"group": schema.SingleNestedAttribute{
							MarkdownDescription: `The Staged Upgrade Group`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID of the Staged Upgrade Group`,
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
			"group": schema.SingleNestedAttribute{
				MarkdownDescription: `The Staged Upgrade Group`,
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
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
		},
	}
}

func (r *NetworksFirmwareUpgradesStagedStagesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksFirmwareUpgradesStagedStagesRs

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
	//Reviw This  Has Item Not item
	//Esta bien

	//dentro
	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkFirmwareUpgradesStagedStages(vvNetworkID)
	// No Post
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksFirmwareUpgradesStagedStages  only have update context, not create.",
			err.Error(),
		)
		return
	}

	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksFirmwareUpgradesStagedStages only have update context, not create.",
			err.Error(),
		)
		return
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	_, restyResp2, err := r.client.Networks.UpdateNetworkFirmwareUpgradesStagedStages(vvNetworkID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkFirmwareUpgradesStagedStages",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkFirmwareUpgradesStagedStages",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Networks.GetNetworkFirmwareUpgradesStagedStages(vvNetworkID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFirmwareUpgradesStagedStages",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkFirmwareUpgradesStagedStages",
			err.Error(),
		)
		return
	}

	data = ResponseNetworksGetNetworkFirmwareUpgradesStagedStagesItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *NetworksFirmwareUpgradesStagedStagesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksFirmwareUpgradesStagedStagesRs

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
	// Not has Item

	vvNetworkID := data.NetworkID.ValueString()
	// Read Review

	responseGet, restyResp1, err := r.client.Networks.GetNetworkFirmwareUpgradesStagedStages(vvNetworkID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFirmwareUpgradesStagedStages",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkFirmwareUpgradesStagedStages",
			err.Error(),
		)
		return
	}
	//entro aqui
	data = ResponseNetworksGetNetworkFirmwareUpgradesStagedStagesItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksFirmwareUpgradesStagedStagesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksFirmwareUpgradesStagedStagesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksFirmwareUpgradesStagedStagesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	_, restyResp2, err := r.client.Networks.UpdateNetworkFirmwareUpgradesStagedStages(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkFirmwareUpgradesStagedStages",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkFirmwareUpgradesStagedStages",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksFirmwareUpgradesStagedStagesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksFirmwareUpgradesStagedStages", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksFirmwareUpgradesStagedStagesRs struct {
	NetworkID types.String `tfsdk:"network_id"`
	//TIENE ITEMS
	Group *ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStagesGroupRs `tfsdk:"group"`
	JSON  *[]RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJsonRs  `tfsdk:"_json"`
}

type ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStagesGroupRs struct {
	Description types.String `tfsdk:"description"`
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
}

type RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJsonRs struct {
	Group *RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJsonGroupRs `tfsdk:"group"`
}

type RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJsonGroupRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *NetworksFirmwareUpgradesStagedStagesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedStages {
	var requestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSON []merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSON

	if r.JSON != nil {
		for _, rItem1 := range *r.JSON {
			var requestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSONGroup *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSONGroup

			if rItem1.Group != nil {
				id := rItem1.Group.ID.ValueString()
				requestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSONGroup = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSONGroup{
					ID: id,
				}
				//[debug] Is Array: False
			}
			requestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSON = append(requestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSON, merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSON{
				Group: requestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSONGroup,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedStages{
		JSON: func() *[]merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSON {
			if len(requestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSON) > 0 {
				return &requestNetworksUpdateNetworkFirmwareUpgradesStagedStagesJSON
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkFirmwareUpgradesStagedStagesItemToBodyRs(state NetworksFirmwareUpgradesStagedStagesRs, response *merakigosdk.ResponseNetworksGetNetworkFirmwareUpgradesStagedStages, is_read bool) NetworksFirmwareUpgradesStagedStagesRs {
	itemState := NetworksFirmwareUpgradesStagedStagesRs{

		// Group: func() *ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStagesGroupRs {
		// 	if response.Group != nil {
		// 		return &ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStagesGroupRs{
		// 			Description: types.StringValue(response.Group.Description),
		// 			ID:          types.StringValue(response.Group.ID),
		// 			Name:        types.StringValue(response.Group.Name),
		// 		}
		// 	}
		// 	return &ResponseItemNetworksGetNetworkFirmwareUpgradesStagedStagesGroupRs{}
		// }(),
	}
	state = itemState
	return state
}
