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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

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
	_ resource.Resource              = &NetworksSwitchStacksResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchStacksResource{}
)

func NewNetworksSwitchStacksResource() resource.Resource {
	return &NetworksSwitchStacksResource{}
}

type NetworksSwitchStacksResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchStacksResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchStacksResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stacks"
}

func (r *NetworksSwitchStacksResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: `ID of the Switch stack`,
				Computed:            true,
			},
			"is_monitor_only": schema.BoolAttribute{
				MarkdownDescription: `Tells if stack is Monitored Stack.`,
				Computed:            true,
			},
			"members": schema.SetNestedAttribute{
				MarkdownDescription: `Members of the Stack`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"mac": schema.StringAttribute{
							MarkdownDescription: `MAC address of the device`,
							Computed:            true,
						},
						"model": schema.StringAttribute{
							MarkdownDescription: `Model of the device`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the device`,
							Computed:            true,
						},
						"role": schema.StringAttribute{
							MarkdownDescription: `Role of the device
                                        Allowed values: [active,member,standby]`,
							Computed: true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial number of the device`,
							Computed:            true,
						},
					},
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the Switch stack`,
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
			"serials": schema.SetAttribute{
				MarkdownDescription: `Serials of the switches in the switch stack`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: `switchStackId path parameter. Switch stack ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *NetworksSwitchStacksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchStacksRs

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
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchStacks(vvNetworkID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkSwitchStacks",
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
			vvSwitchStackID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter SwitchStackID",
					"Fail Parsing SwitchStackID",
				)
				return
			}
			// r.client..( data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Switch.GetNetworkSwitchStack(vvNetworkID, vvSwitchStackID)
			if responseVerifyItem2 != nil {
				data = ResponseSwitchGetNetworkSwitchStackItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Switch.CreateNetworkSwitchStack(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkSwitchStack",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkSwitchStack",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchStacks(vvNetworkID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStacks",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchStacks",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvSwitchStackID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter SwitchStackID",
				"Error",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Switch.GetNetworkSwitchStack(vvNetworkID, vvSwitchStackID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseSwitchGetNetworkSwitchStackItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkSwitchStack",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchStack",
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

func (r *NetworksSwitchStacksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchStacksRs

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
	vvSwitchStackID := data.SwitchStackID.ValueString()
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchStack(vvNetworkID, vvSwitchStackID)
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
				"Failure when executing GetNetworkSwitchStack",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchStack",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchStackItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSwitchStacksResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("switch_stack_id"), idParts[1])...)
}

func (r *NetworksSwitchStacksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchStacksRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update
	// No update
	resp.Diagnostics.AddError(
		"Update operation not supported in NetworksSwitchStacks",
		"Update operation not supported in NetworksSwitchStacks",
	)
	return
}

func (r *NetworksSwitchStacksResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksSwitchStacksRs
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
	vvSwitchStackID := state.SwitchStackID.ValueString()
	_, err := r.client.Switch.DeleteNetworkSwitchStack(vvNetworkID, vvSwitchStackID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkSwitchStack", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksSwitchStacksRs struct {
	NetworkID     types.String                                    `tfsdk:"network_id"`
	SwitchStackID types.String                                    `tfsdk:"switch_stack_id"`
	ID            types.String                                    `tfsdk:"id"`
	IsMonitorOnly types.Bool                                      `tfsdk:"is_monitor_only"`
	Members       *[]ResponseSwitchGetNetworkSwitchStackMembersRs `tfsdk:"members"`
	Name          types.String                                    `tfsdk:"name"`
	Serials       types.Set                                       `tfsdk:"serials"`
}

type ResponseSwitchGetNetworkSwitchStackMembersRs struct {
	Mac    types.String `tfsdk:"mac"`
	Model  types.String `tfsdk:"model"`
	Name   types.String `tfsdk:"name"`
	Role   types.String `tfsdk:"role"`
	Serial types.String `tfsdk:"serial"`
}

// FromBody
func (r *NetworksSwitchStacksRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCreateNetworkSwitchStack {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var serials []string = nil
	r.Serials.ElementsAs(ctx, &serials, false)
	out := merakigosdk.RequestSwitchCreateNetworkSwitchStack{
		Name:    *name,
		Serials: serials,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchStackItemToBodyRs(state NetworksSwitchStacksRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchStack, is_read bool) NetworksSwitchStacksRs {
	itemState := NetworksSwitchStacksRs{
		ID: types.StringValue(response.ID),
		IsMonitorOnly: func() types.Bool {
			if response.IsMonitorOnly != nil {
				return types.BoolValue(*response.IsMonitorOnly)
			}
			return types.Bool{}
		}(),
		Members: func() *[]ResponseSwitchGetNetworkSwitchStackMembersRs {
			if response.Members != nil {
				result := make([]ResponseSwitchGetNetworkSwitchStackMembersRs, len(*response.Members))
				for i, members := range *response.Members {
					result[i] = ResponseSwitchGetNetworkSwitchStackMembersRs{
						Mac:    types.StringValue(members.Mac),
						Model:  types.StringValue(members.Model),
						Name:   types.StringValue(members.Name),
						Role:   types.StringValue(members.Role),
						Serial: types.StringValue(members.Serial),
					}
				}
				return &result
			}
			return nil
		}(),
		Name:    types.StringValue(response.Name),
		Serials: StringSliceToSet(response.Serials),
	}
	itemState.SwitchStackID = itemState.ID
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchStacksRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchStacksRs)
}
