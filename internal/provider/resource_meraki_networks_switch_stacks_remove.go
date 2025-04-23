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
	_ resource.Resource              = &NetworksSwitchStacksRemoveResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchStacksRemoveResource{}
)

func NewNetworksSwitchStacksRemoveResource() resource.Resource {
	return &NetworksSwitchStacksRemoveResource{}
}

type NetworksSwitchStacksRemoveResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchStacksRemoveResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchStacksRemoveResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stacks_remove"
}

// resourceAction
func (r *NetworksSwitchStacksRemoveResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"switch_stack_id": schema.StringAttribute{
				MarkdownDescription: `switchStackId path parameter. Switch stack ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
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
					},
					"serials": schema.ListAttribute{
						MarkdownDescription: `Serials of the switches in the switch stack`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"serial": schema.StringAttribute{
						MarkdownDescription: `The serial of the switch to be removed`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *NetworksSwitchStacksRemoveResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchStacksRemove

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
	vvSwitchStackID := data.SwitchStackID.ValueString()
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Switch.RemoveNetworkSwitchStack(vvNetworkID, vvSwitchStackID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing RemoveNetworkSwitchStack",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing RemoveNetworkSwitchStack",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseSwitchRemoveNetworkSwitchStackItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchStacksRemoveResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSwitchStacksRemoveResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSwitchStacksRemoveResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchStacksRemove struct {
	NetworkID     types.String                             `tfsdk:"network_id"`
	SwitchStackID types.String                             `tfsdk:"switch_stack_id"`
	Item          *ResponseSwitchRemoveNetworkSwitchStack  `tfsdk:"item"`
	Parameters    *RequestSwitchRemoveNetworkSwitchStackRs `tfsdk:"parameters"`
}

type ResponseSwitchRemoveNetworkSwitchStack struct {
	ID            types.String                                     `tfsdk:"id"`
	IsMonitorOnly types.Bool                                       `tfsdk:"is_monitor_only"`
	Members       *[]ResponseSwitchRemoveNetworkSwitchStackMembers `tfsdk:"members"`
	Name          types.String                                     `tfsdk:"name"`
	Serials       types.List                                       `tfsdk:"serials"`
}

type ResponseSwitchRemoveNetworkSwitchStackMembers struct {
	Mac    types.String `tfsdk:"mac"`
	Model  types.String `tfsdk:"model"`
	Name   types.String `tfsdk:"name"`
	Role   types.String `tfsdk:"role"`
	Serial types.String `tfsdk:"serial"`
}

type RequestSwitchRemoveNetworkSwitchStackRs struct {
	Serial types.String `tfsdk:"serial"`
}

// FromBody
func (r *NetworksSwitchStacksRemove) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchRemoveNetworkSwitchStack {
	emptyString := ""
	re := *r.Parameters
	serial := new(string)
	if !re.Serial.IsUnknown() && !re.Serial.IsNull() {
		*serial = re.Serial.ValueString()
	} else {
		serial = &emptyString
	}
	out := merakigosdk.RequestSwitchRemoveNetworkSwitchStack{
		Serial: *serial,
	}
	return &out
}

// ToBody
func ResponseSwitchRemoveNetworkSwitchStackItemToBody(state NetworksSwitchStacksRemove, response *merakigosdk.ResponseSwitchRemoveNetworkSwitchStack) NetworksSwitchStacksRemove {
	itemState := ResponseSwitchRemoveNetworkSwitchStack{
		ID: types.StringValue(response.ID),
		IsMonitorOnly: func() types.Bool {
			if response.IsMonitorOnly != nil {
				return types.BoolValue(*response.IsMonitorOnly)
			}
			return types.Bool{}
		}(),
		Members: func() *[]ResponseSwitchRemoveNetworkSwitchStackMembers {
			if response.Members != nil {
				result := make([]ResponseSwitchRemoveNetworkSwitchStackMembers, len(*response.Members))
				for i, members := range *response.Members {
					result[i] = ResponseSwitchRemoveNetworkSwitchStackMembers{
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
		Serials: StringSliceToList(response.Serials),
	}
	state.Item = &itemState
	return state
}
