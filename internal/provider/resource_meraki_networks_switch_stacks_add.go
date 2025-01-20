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
	_ resource.Resource              = &NetworksSwitchStacksAddResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchStacksAddResource{}
)

func NewNetworksSwitchStacksAddResource() resource.Resource {
	return &NetworksSwitchStacksAddResource{}
}

type NetworksSwitchStacksAddResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchStacksAddResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchStacksAddResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_stacks_add"
}

// resourceAction
func (r *NetworksSwitchStacksAddResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					"serials": schema.SetAttribute{
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
						MarkdownDescription: `The serial of the switch to be added`,
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
func (r *NetworksSwitchStacksAddResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchStacksAdd

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
	response, restyResp1, err := r.client.Switch.AddNetworkSwitchStack(vvNetworkID, vvSwitchStackID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing AddNetworkSwitchStack",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing AddNetworkSwitchStack",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseSwitchAddNetworkSwitchStackItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchStacksAddResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSwitchStacksAddResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSwitchStacksAddResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchStacksAdd struct {
	NetworkID     types.String                          `tfsdk:"network_id"`
	SwitchStackID types.String                          `tfsdk:"switch_stack_id"`
	Item          *ResponseSwitchAddNetworkSwitchStack  `tfsdk:"item"`
	Parameters    *RequestSwitchAddNetworkSwitchStackRs `tfsdk:"parameters"`
}

type ResponseSwitchAddNetworkSwitchStack struct {
	ID            types.String                                  `tfsdk:"id"`
	IsMonitorOnly types.Bool                                    `tfsdk:"is_monitor_only"`
	Members       *[]ResponseSwitchAddNetworkSwitchStackMembers `tfsdk:"members"`
	Name          types.String                                  `tfsdk:"name"`
	Serials       types.Set                                     `tfsdk:"serials"`
}

type ResponseSwitchAddNetworkSwitchStackMembers struct {
	Mac    types.String `tfsdk:"mac"`
	Model  types.String `tfsdk:"model"`
	Name   types.String `tfsdk:"name"`
	Role   types.String `tfsdk:"role"`
	Serial types.String `tfsdk:"serial"`
}

type RequestSwitchAddNetworkSwitchStackRs struct {
	Serial types.String `tfsdk:"serial"`
}

// FromBody
func (r *NetworksSwitchStacksAdd) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchAddNetworkSwitchStack {
	emptyString := ""
	re := *r.Parameters
	serial := new(string)
	if !re.Serial.IsUnknown() && !re.Serial.IsNull() {
		*serial = re.Serial.ValueString()
	} else {
		serial = &emptyString
	}
	out := merakigosdk.RequestSwitchAddNetworkSwitchStack{
		Serial: *serial,
	}
	return &out
}

// ToBody
func ResponseSwitchAddNetworkSwitchStackItemToBody(state NetworksSwitchStacksAdd, response *merakigosdk.ResponseSwitchAddNetworkSwitchStack) NetworksSwitchStacksAdd {
	itemState := ResponseSwitchAddNetworkSwitchStack{
		ID: types.StringValue(response.ID),
		IsMonitorOnly: func() types.Bool {
			if response.IsMonitorOnly != nil {
				return types.BoolValue(*response.IsMonitorOnly)
			}
			return types.Bool{}
		}(),
		Members: func() *[]ResponseSwitchAddNetworkSwitchStackMembers {
			if response.Members != nil {
				result := make([]ResponseSwitchAddNetworkSwitchStackMembers, len(*response.Members))
				for i, members := range *response.Members {
					result[i] = ResponseSwitchAddNetworkSwitchStackMembers{
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
	state.Item = &itemState
	return state
}
