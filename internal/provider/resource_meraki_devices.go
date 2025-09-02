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
	"strconv"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesResource{}
	_ resource.ResourceWithConfigure = &DevicesResource{}
)

func NewDevicesResource() resource.Resource {
	return &DevicesResource{}
}

type DevicesResource struct {
	client *merakigosdk.Client
}

func (r *DevicesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices"
}

func (r *DevicesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				MarkdownDescription: `Physical address of the device`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"firmware": schema.StringAttribute{
				MarkdownDescription: `Firmware version of the device`,
				Optional:            true,
			},
			"floor_plan_id": schema.StringAttribute{
				MarkdownDescription: `The floor plan to associate to this device. null disassociates the device from the floorplan.`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"lan_ip": schema.StringAttribute{
				MarkdownDescription: `LAN IP address of the device`,
				Optional:            true,
			},
			"lat": schema.Float64Attribute{
				MarkdownDescription: `Latitude of the device`,
				Optional:            true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
			},
			"lng": schema.Float64Attribute{
				MarkdownDescription: `Longitude of the device`,
				Optional:            true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `MAC address of the device`,
				Optional:            true,
			},
			"model": schema.StringAttribute{
				MarkdownDescription: `Model of the device`,
				Optional:            true,
			},
			"move_map_marker": schema.BoolAttribute{
				MarkdownDescription: `Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the device`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `ID of the network the device belongs to`,
				Optional:            true,
			},
			"notes": schema.StringAttribute{
				MarkdownDescription: `Notes for the device, limited to 255 characters`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			// "organization_id": schema.StringAttribute{
			// 	MarkdownDescription: `organizationId path parameter. Organization ID`,
			// 	Required:            true,
			// },
			"serial": schema.StringAttribute{
				MarkdownDescription: `Serial number of the device`,
				Required:            true,
			},
			"switch_profile_id": schema.StringAttribute{
				MarkdownDescription: `The ID of a switch template to bind to the device (for available switch templates, see the 'Switch Templates' endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch template, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"tags": schema.ListAttribute{
				MarkdownDescription: `List of tags assigned to the device`,
				Optional:            true,
				Computed:            true,
				Default:             listdefault.StaticValue(types.ListNull(types.StringType)),
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
		},
	}
}

//path params to set ['serial']

func (r *DevicesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesRs

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
	vvSerial := data.Serial.ValueString()
	//Has Item and has items and not post

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Devices.UpdateDevice(vvSerial, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDevice",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDevice",
			err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)

}

func (r *DevicesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesRs

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
		return
	}

	//Has Paths
	// Has Item2

	vvSerial := data.Serial.ValueString()
	responseGet, restyRespGet, err := r.client.Devices.GetDevice(vvSerial)
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
				"Failure when executing GetDevice",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDevice",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseDevicesGetDeviceItemToBodyRs(data, responseGet, true)
	fmt.Printf("SEE HEEREEEEEE: %+v\n", data)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *DevicesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DevicesRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvSerial := plan.Serial.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Devices.UpdateDevice(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDevice",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDevice",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *DevicesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Devices", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesRs struct {
	Serial          types.String  `tfsdk:"serial"`
	Address         types.String  `tfsdk:"address"`
	Firmware        types.String  `tfsdk:"firmware"`
	FloorPlanID     types.String  `tfsdk:"floor_plan_id"`
	LanIP           types.String  `tfsdk:"lan_ip"`
	Lat             types.Float64 `tfsdk:"lat"`
	Lng             types.Float64 `tfsdk:"lng"`
	Mac             types.String  `tfsdk:"mac"`
	Model           types.String  `tfsdk:"model"`
	Name            types.String  `tfsdk:"name"`
	NetworkID       types.String  `tfsdk:"network_id"`
	Notes           types.String  `tfsdk:"notes"`
	Tags            types.List    `tfsdk:"tags"`
	MoveMapMarker   types.Bool    `tfsdk:"move_map_marker"`
	SwitchProfileID types.String  `tfsdk:"switch_profile_id"`
}

type ResponseDevicesGetDeviceBeaconIdParamsRs struct {
	Major types.Int64  `tfsdk:"major"`
	Minor types.Int64  `tfsdk:"minor"`
	UUID  types.String `tfsdk:"uuid"`
}

// FromBody
func (r *DevicesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestDevicesUpdateDevice {
	emptyString := ""
	address := new(string)
	if !r.Address.IsUnknown() && !r.Address.IsNull() {
		*address = r.Address.ValueString()
	} else {
		address = &emptyString
	}
	floorPlanID := new(string)
	if !r.FloorPlanID.IsUnknown() && !r.FloorPlanID.IsNull() {
		*floorPlanID = r.FloorPlanID.ValueString()
	} else {
		floorPlanID = &emptyString
	}
	lat := new(float64)
	if !r.Lat.IsUnknown() && !r.Lat.IsNull() {
		*lat = r.Lat.ValueFloat64()
	} else {
		lat = nil
	}
	lng := new(float64)
	if !r.Lng.IsUnknown() && !r.Lng.IsNull() {
		*lng = r.Lng.ValueFloat64()
	} else {
		lng = nil
	}
	moveMapMarker := new(bool)
	if !r.MoveMapMarker.IsUnknown() && !r.MoveMapMarker.IsNull() {
		*moveMapMarker = r.MoveMapMarker.ValueBool()
	} else {
		moveMapMarker = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	notes := new(string)
	if !r.Notes.IsUnknown() && !r.Notes.IsNull() {
		*notes = r.Notes.ValueString()
	} else {
		notes = &emptyString
	}
	switchProfileID := new(string)
	if !r.SwitchProfileID.IsUnknown() && !r.SwitchProfileID.IsNull() {
		*switchProfileID = r.SwitchProfileID.ValueString()
	} else {
		switchProfileID = &emptyString
	}
	var tags []string = nil
	r.Tags.ElementsAs(ctx, &tags, false)
	out := merakigosdk.RequestDevicesUpdateDevice{
		Address:         *address,
		FloorPlanID:     *floorPlanID,
		Lat:             lat,
		Lng:             lng,
		MoveMapMarker:   moveMapMarker,
		Name:            *name,
		Notes:           *notes,
		SwitchProfileID: *switchProfileID,
		Tags:            tags,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseDevicesGetDeviceItemToBodyRs(state DevicesRs, response *merakigosdk.ResponseDevicesGetDevice, is_read bool) DevicesRs {
	fmt.Printf("Aqui llego: %+v\n", response)
	itemState := DevicesRs{
		Address: func() types.String {
			if response.Address != "" {
				return types.StringValue(response.Address)
			}
			return types.StringNull()
		}(),
		Firmware: func() types.String {
			if response.Firmware != "" {
				return types.StringValue(response.Firmware)
			}
			return types.StringNull()
		}(),
		FloorPlanID: func() types.String {
			if response.FloorPlanID != "" {
				return types.StringValue(response.FloorPlanID)
			}
			return types.StringNull()
		}(),
		LanIP: func() types.String {
			if response.LanIP != "" {
				return types.StringValue(response.LanIP)
			}
			return types.StringNull()
		}(),
		Lat: func() types.Float64 {
			if response.Lat != nil {
				return types.Float64Value(float64(*response.Lat))
			}
			return types.Float64{}
		}(),
		Lng: func() types.Float64 {
			if response.Lng != nil {
				return types.Float64Value(float64(*response.Lng))
			}
			return types.Float64{}
		}(),
		Mac: func() types.String {
			if response.Mac != "" {
				return types.StringValue(response.Mac)
			}
			return types.StringNull()
		}(),
		Model: func() types.String {
			if response.Model != "" {
				return types.StringValue(response.Model)
			}
			return types.StringNull()
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.StringNull()
		}(),
		NetworkID: func() types.String {
			if response.NetworkID != "" {
				return types.StringValue(response.NetworkID)
			}
			return types.StringNull()
		}(),
		Notes: func() types.String {
			if response.Notes != "" {
				return types.StringValue(response.Notes)
			}
			return types.StringNull()
		}(),
		Serial: func() types.String {
			if response.Serial != "" {
				return types.StringValue(response.Serial)
			}
			return types.StringNull()
		}(),
		Tags: func() types.List {
			if len(response.Tags) > 0 {
				return StringSliceToList(response.Tags)
			}
			return state.Tags
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesRs)
}
