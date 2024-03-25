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

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesSwitchPortsCycleResource{}
	_ resource.ResourceWithConfigure = &DevicesSwitchPortsCycleResource{}
)

func NewDevicesSwitchPortsCycleResource() resource.Resource {
	return &DevicesSwitchPortsCycleResource{}
}

type DevicesSwitchPortsCycleResource struct {
	client *merakigosdk.Client
}

func (r *DevicesSwitchPortsCycleResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesSwitchPortsCycleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_switch_ports_cycle"
}

// resourceAction
func (r *DevicesSwitchPortsCycleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{

			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"ports": schema.SetAttribute{
						MarkdownDescription: `List of switch ports`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"ports": schema.SetAttribute{
						MarkdownDescription: `List of switch ports`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *DevicesSwitchPortsCycleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesSwitchPortsCycle

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
	vvSerial := data.Serial.ValueString()
	// serial
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Switch.CycleDeviceSwitchPorts(vvSerial, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CycleDeviceSwitchPorts",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CycleDeviceSwitchPorts",
			err.Error(),
		)
		return
	}
	//Item
	data2 := ResponseSwitchCycleDeviceSwitchPortsItemToBody(data, response)

	diags := resp.State.Set(ctx, &data2)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesSwitchPortsCycleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesSwitchPortsCycleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesSwitchPortsCycleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesSwitchPortsCycle struct {
	Serial     types.String                           `tfsdk:"serial"`
	Item       *ResponseSwitchCycleDeviceSwitchPorts  `tfsdk:"item"`
	Parameters *RequestSwitchCycleDeviceSwitchPortsRs `tfsdk:"parameters"`
}

type ResponseSwitchCycleDeviceSwitchPorts struct {
	Ports types.Set `tfsdk:"ports"`
}

type RequestSwitchCycleDeviceSwitchPortsRs struct {
	Ports types.Set `tfsdk:"ports"`
}

// FromBody
func (r *DevicesSwitchPortsCycle) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCycleDeviceSwitchPorts {
	re := *r.Parameters
	var ports []string = nil
	re.Ports.ElementsAs(ctx, &ports, false)
	out := merakigosdk.RequestSwitchCycleDeviceSwitchPorts{
		Ports: ports,
	}
	return &out
}

// ToBody
func ResponseSwitchCycleDeviceSwitchPortsItemToBody(state DevicesSwitchPortsCycle, response *merakigosdk.ResponseSwitchCycleDeviceSwitchPorts) DevicesSwitchPortsCycle {
	itemState := ResponseSwitchCycleDeviceSwitchPorts{
		Ports: StringSliceToSet(response.Ports),
	}
	state.Item = &itemState
	return state
}
