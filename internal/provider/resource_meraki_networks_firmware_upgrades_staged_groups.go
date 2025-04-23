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

	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksFirmwareUpgradesStagedGroupsResource{}
	_ resource.ResourceWithConfigure = &NetworksFirmwareUpgradesStagedGroupsResource{}
)

func NewNetworksFirmwareUpgradesStagedGroupsResource() resource.Resource {
	return &NetworksFirmwareUpgradesStagedGroupsResource{}
}

type NetworksFirmwareUpgradesStagedGroupsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksFirmwareUpgradesStagedGroupsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksFirmwareUpgradesStagedGroupsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_firmware_upgrades_staged_groups"
}

func (r *NetworksFirmwareUpgradesStagedGroupsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"assigned_devices": schema.SingleNestedAttribute{
				MarkdownDescription: `The devices and Switch Stacks assigned to the Group`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"devices": schema.SetNestedAttribute{
						MarkdownDescription: `Data Array of Devices containing the name and serial`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the device`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `Serial of the device`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
					"switch_stacks": schema.SetNestedAttribute{
						MarkdownDescription: `Data Array of Switch Stacks containing the name and id`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID of the Switch Stack`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the Switch Stack`,
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
			"description": schema.StringAttribute{
				MarkdownDescription: `Description of the Staged Upgrade Group`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"group_id": schema.StringAttribute{
				MarkdownDescription: `Id of staged upgrade group`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"is_default": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the Staged Upgrade Group`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
		},
	}
}

//path params to set ['groupId']

func (r *NetworksFirmwareUpgradesStagedGroupsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksFirmwareUpgradesStagedGroupsRs

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
	//Has Item and has items and post

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Networks.GetNetworkFirmwareUpgradesStagedGroups(vvNetworkID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkFirmwareUpgradesStagedGroups",
					err.Error(),
				)
				return
			}
		}
	}

	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvGroupID, ok := result2["GroupID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter GroupID",
					"Fail Parsing GroupID",
				)
				return
			}
			r.client.Networks.UpdateNetworkFirmwareUpgradesStagedGroup(vvNetworkID, vvGroupID, data.toSdkApiRequestUpdate(ctx))

			responseVerifyItem2, _, _ := r.client.Networks.GetNetworkFirmwareUpgradesStagedGroup(vvNetworkID, vvGroupID)
			if responseVerifyItem2 != nil {
				data = ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Networks.CreateNetworkFirmwareUpgradesStagedGroup(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkFirmwareUpgradesStagedGroup",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkFirmwareUpgradesStagedGroup",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Networks.GetNetworkFirmwareUpgradesStagedGroups(vvNetworkID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFirmwareUpgradesStagedGroups",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkFirmwareUpgradesStagedGroups",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvGroupID, ok := result2["GroupID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter GroupID",
				"Fail Parsing GroupID",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Networks.GetNetworkFirmwareUpgradesStagedGroup(vvNetworkID, vvGroupID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkFirmwareUpgradesStagedGroup",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkFirmwareUpgradesStagedGroup",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}

}

func (r *NetworksFirmwareUpgradesStagedGroupsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksFirmwareUpgradesStagedGroupsRs

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
	vvGroupID := data.GroupID.ValueString()
	responseGet, restyRespGet, err := r.client.Networks.GetNetworkFirmwareUpgradesStagedGroup(vvNetworkID, vvGroupID)
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
				"Failure when executing GetNetworkFirmwareUpgradesStagedGroup",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkFirmwareUpgradesStagedGroup",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksFirmwareUpgradesStagedGroupsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("group_id"), idParts[1])...)
}

func (r *NetworksFirmwareUpgradesStagedGroupsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksFirmwareUpgradesStagedGroupsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvGroupID := data.GroupID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Networks.UpdateNetworkFirmwareUpgradesStagedGroup(vvNetworkID, vvGroupID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkFirmwareUpgradesStagedGroup",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkFirmwareUpgradesStagedGroup",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksFirmwareUpgradesStagedGroupsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksFirmwareUpgradesStagedGroupsRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvNetworkID := state.NetworkID.ValueString()
	vvGroupID := state.GroupID.ValueString()
	_, err := r.client.Networks.DeleteNetworkFirmwareUpgradesStagedGroup(vvNetworkID, vvGroupID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkFirmwareUpgradesStagedGroup", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksFirmwareUpgradesStagedGroupsRs struct {
	NetworkID       types.String                                                            `tfsdk:"network_id"`
	GroupID         types.String                                                            `tfsdk:"group_id"`
	AssignedDevices *ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesRs `tfsdk:"assigned_devices"`
	Description     types.String                                                            `tfsdk:"description"`
	IsDefault       types.Bool                                                              `tfsdk:"is_default"`
	Name            types.String                                                            `tfsdk:"name"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesRs struct {
	Devices      *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevicesRs      `tfsdk:"devices"`
	SwitchStacks *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacksRs `tfsdk:"switch_stacks"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevicesRs struct {
	Name   types.String `tfsdk:"name"`
	Serial types.String `tfsdk:"serial"`
}

type ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacksRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// FromBody
func (r *NetworksFirmwareUpgradesStagedGroupsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedGroup {
	emptyString := ""
	var requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevices *merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevices

	if r.AssignedDevices != nil {

		log.Printf("[DEBUG] #TODO []RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices")
		var requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices []merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices

		if r.AssignedDevices.Devices != nil {
			for _, rItem1 := range *r.AssignedDevices.Devices {
				name := rItem1.Name.ValueString()
				serial := rItem1.Serial.ValueString()
				requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices = append(requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices, merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices{
					Name:   name,
					Serial: serial,
				})
				//[debug] Is Array: True
			}
		}

		log.Printf("[DEBUG] #TODO []RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks")
		var requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks []merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks

		if r.AssignedDevices.SwitchStacks != nil {
			for _, rItem1 := range *r.AssignedDevices.SwitchStacks {
				id := rItem1.ID.ValueString()
				name := rItem1.Name.ValueString()
				requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks = append(requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks, merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks{
					ID:   id,
					Name: name,
				})
				//[debug] Is Array: True
			}
		}
		requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevices = &merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevices{
			Devices: func() *[]merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices {
				if len(requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices) > 0 {
					return &requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices
				}
				return nil
			}(),
			SwitchStacks: func() *[]merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks {
				if len(requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks) > 0 {
					return &requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks
				}
				return nil
			}(),
		}
		//[debug] Is Array: False
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = &emptyString
	}
	isDefault := new(bool)
	if !r.IsDefault.IsUnknown() && !r.IsDefault.IsNull() {
		*isDefault = r.IsDefault.ValueBool()
	} else {
		isDefault = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestNetworksCreateNetworkFirmwareUpgradesStagedGroup{
		AssignedDevices: requestNetworksCreateNetworkFirmwareUpgradesStagedGroupAssignedDevices,
		Description:     *description,
		IsDefault:       isDefault,
		Name:            *name,
	}
	return &out
}
func (r *NetworksFirmwareUpgradesStagedGroupsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroup {
	emptyString := ""
	var requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevices *merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevices

	if r.AssignedDevices != nil {

		log.Printf("[DEBUG] #TODO []RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices")
		var requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices []merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices

		if r.AssignedDevices.Devices != nil {
			for _, rItem1 := range *r.AssignedDevices.Devices {
				name := rItem1.Name.ValueString()
				serial := rItem1.Serial.ValueString()
				requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices = append(requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices, merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices{
					Name:   name,
					Serial: serial,
				})
				//[debug] Is Array: True
			}
		}

		log.Printf("[DEBUG] #TODO []RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks")
		var requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks []merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks

		if r.AssignedDevices.SwitchStacks != nil {
			for _, rItem1 := range *r.AssignedDevices.SwitchStacks {
				id := rItem1.ID.ValueString()
				name := rItem1.Name.ValueString()
				requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks = append(requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks, merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks{
					ID:   id,
					Name: name,
				})
				//[debug] Is Array: True
			}
		}
		requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevices = &merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevices{
			Devices: func() *[]merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices {
				if len(requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices) > 0 {
					return &requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevices
				}
				return nil
			}(),
			SwitchStacks: func() *[]merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks {
				if len(requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks) > 0 {
					return &requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacks
				}
				return nil
			}(),
		}
		//[debug] Is Array: False
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = &emptyString
	}
	isDefault := new(bool)
	if !r.IsDefault.IsUnknown() && !r.IsDefault.IsNull() {
		*isDefault = r.IsDefault.ValueBool()
	} else {
		isDefault = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestNetworksUpdateNetworkFirmwareUpgradesStagedGroup{
		AssignedDevices: requestNetworksUpdateNetworkFirmwareUpgradesStagedGroupAssignedDevices,
		Description:     *description,
		IsDefault:       isDefault,
		Name:            *name,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupItemToBodyRs(state NetworksFirmwareUpgradesStagedGroupsRs, response *merakigosdk.ResponseNetworksGetNetworkFirmwareUpgradesStagedGroup, is_read bool) NetworksFirmwareUpgradesStagedGroupsRs {
	itemState := NetworksFirmwareUpgradesStagedGroupsRs{
		AssignedDevices: func() *ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesRs {
			if response.AssignedDevices != nil {
				return &ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesRs{
					Devices: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevicesRs {
						if response.AssignedDevices.Devices != nil {
							result := make([]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevicesRs, len(*response.AssignedDevices.Devices))
							for i, devices := range *response.AssignedDevices.Devices {
								result[i] = ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesDevicesRs{
									Name:   types.StringValue(devices.Name),
									Serial: types.StringValue(devices.Serial),
								}
							}
							return &result
						}
						return nil
					}(),
					SwitchStacks: func() *[]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacksRs {
						if response.AssignedDevices.SwitchStacks != nil {
							result := make([]ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacksRs, len(*response.AssignedDevices.SwitchStacks))
							for i, switchStacks := range *response.AssignedDevices.SwitchStacks {
								result[i] = ResponseNetworksGetNetworkFirmwareUpgradesStagedGroupAssignedDevicesSwitchStacksRs{
									ID:   types.StringValue(switchStacks.ID),
									Name: types.StringValue(switchStacks.Name),
								}
							}
							return &result
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		Description: types.StringValue(response.Description),
		GroupID:     types.StringValue(response.GroupID),
		IsDefault: func() types.Bool {
			if response.IsDefault != nil {
				return types.BoolValue(*response.IsDefault)
			}
			return types.Bool{}
		}(),
		Name: types.StringValue(response.Name),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksFirmwareUpgradesStagedGroupsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksFirmwareUpgradesStagedGroupsRs)
}
