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
	"net/url"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesResource{}
	_ resource.ResourceWithConfigure = &DevicesResource{}
	// _ resource.ResourceWithImportState = &DevicesResource{}
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
				MarkdownDescription: `The address of a device`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"beacon_id_params": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"major": schema.Int64Attribute{
						Computed: true,
					},
					"minor": schema.Int64Attribute{
						Computed: true,
					},
					"uuid": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"firmware": schema.StringAttribute{
				Computed: true,
			},
			"floor_plan_id": schema.StringAttribute{
				MarkdownDescription: `The floor plan to associate to this device. null disassociates the device from the floorplan.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"lan_ip": schema.StringAttribute{
				Computed: true,
			},
			"lat": schema.Float64Attribute{
				MarkdownDescription: `The latitude of a device`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
			},
			"lng": schema.Float64Attribute{
				MarkdownDescription: `The longitude of a device`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `Mac.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"model": schema.StringAttribute{
				Computed: true,
			},
			"move_map_marker": schema.BoolAttribute{
				MarkdownDescription: `Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of a device`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				Computed: true,
			},
			"notes": schema.StringAttribute{
				MarkdownDescription: `The notes for the device. String. Limited to 255 characters.`,
				Computed:            true,
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
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"switch_profile_id": schema.StringAttribute{
				MarkdownDescription: `The ID of a switch profile to bind to the device (for available switch profiles, see the 'Switch Profiles' endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch profile, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"tags": schema.SetAttribute{
				MarkdownDescription: `The list of tags of a device`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
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
	//Has Paths
	// vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	vvSerial := data.Serial.ValueString()
	// serial
	//Item
	responseVerifyItem, restyResp1, err := r.client.Devices.GetDevice(vvSerial)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource Devices only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource Devices only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Devices.UpdateDevice(vvSerial, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDevice",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDevice",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Devices.GetDevice(vvSerial)
	// Has only items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevices",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationDevices",
			err.Error(),
		)
		return
	}
	data = ResponseDevicesGetDeviceItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesRs

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

	vvSerial := data.Serial.ValueString()
	// serial
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
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDevice",
			err.Error(),
		)
		return
	}

	data = ResponseDevicesGetDeviceItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	// serial
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Devices.UpdateDevice(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDevice",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDevice",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesRs struct {
	Serial types.String `tfsdk:"serial"`
	// OrganizationID  types.String                              `tfsdk:"organization_id"`
	Address         types.String                              `tfsdk:"address"`
	BeaconIDParams  *ResponseDevicesGetDeviceBeaconIdParamsRs `tfsdk:"beacon_id_params"`
	Firmware        types.String                              `tfsdk:"firmware"`
	FloorPlanID     types.String                              `tfsdk:"floor_plan_id"`
	LanIP           types.String                              `tfsdk:"lan_ip"`
	Lat             types.Float64                             `tfsdk:"lat"`
	Lng             types.Float64                             `tfsdk:"lng"`
	Mac             types.String                              `tfsdk:"mac"`
	Model           types.String                              `tfsdk:"model"`
	Name            types.String                              `tfsdk:"name"`
	NetworkID       types.String                              `tfsdk:"network_id"`
	Notes           types.String                              `tfsdk:"notes"`
	Tags            types.Set                                 `tfsdk:"tags"`
	MoveMapMarker   types.Bool                                `tfsdk:"move_map_marker"`
	SwitchProfileID types.String                              `tfsdk:"switch_profile_id"`
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
	mac := new(string)
	if !r.Mac.IsUnknown() && !r.Mac.IsNull() {
		*mac = r.Mac.ValueString()
	} else {
		mac = &emptyString
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
		Mac:             *mac,
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
	itemState := DevicesRs{
		Address: types.StringValue(response.Address),
		BeaconIDParams: func() *ResponseDevicesGetDeviceBeaconIdParamsRs {
			if response.BeaconIDParams != nil {
				return &ResponseDevicesGetDeviceBeaconIdParamsRs{
					Major: func() types.Int64 {
						if response.BeaconIDParams.Major != nil {
							return types.Int64Value(int64(*response.BeaconIDParams.Major))
						}
						return types.Int64{}
					}(),
					Minor: func() types.Int64 {
						if response.BeaconIDParams.Minor != nil {
							return types.Int64Value(int64(*response.BeaconIDParams.Minor))
						}
						return types.Int64{}
					}(),
					UUID: types.StringValue(response.BeaconIDParams.UUID),
				}
			}
			return &ResponseDevicesGetDeviceBeaconIdParamsRs{}
		}(),
		Firmware:    types.StringValue(response.Firmware),
		FloorPlanID: types.StringValue(response.FloorPlanID),
		LanIP:       types.StringValue(response.LanIP),
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
		Mac:       types.StringValue(response.Mac),
		Model:     types.StringValue(response.Model),
		Name:      types.StringValue(response.Name),
		NetworkID: types.StringValue(response.NetworkID),
		Notes:     types.StringValue(response.Notes),
		Serial:    types.StringValue(response.Serial),
		Tags:      StringSliceToSet(response.Tags),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesRs)
}

func getAllItemsDevices(client merakigosdk.Client, organizationId string) (merakigosdk.ResponseOrganizationsGetOrganizationDevices, *resty.Response, error) {
	var all_response merakigosdk.ResponseOrganizationsGetOrganizationDevices
	response, r2, er := client.Organizations.GetOrganizationDevices(organizationId, &merakigosdk.GetOrganizationDevicesQueryParams{
		PerPage: 1000,
	})
	count := 0
	all_response = append(all_response, *response...)
	for len(*response) >= 1000 {
		count += 1
		fmt.Println(count)
		links := strings.Split(r2.Header().Get("Link"), ",")
		var link string
		if count > 1 {
			link = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Split(links[2], ";")[0], ">", ""), "<", ""), client.RestyClient().BaseURL, "")
		} else {
			link = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Split(links[1], ";")[0], ">", ""), "<", ""), client.RestyClient().BaseURL, "")
		}
		myUrl, _ := url.Parse(link)
		params, _ := url.ParseQuery(myUrl.RawQuery)
		if params["endingBefore"] != nil {
			response, r2, er = client.Organizations.GetOrganizationDevices(organizationId, &merakigosdk.GetOrganizationDevicesQueryParams{
				PerPage:      1000,
				EndingBefore: params["endingBefore"][0],
			})
			all_response = append(all_response, *response...)
		} else {
			break
		}
	}

	return all_response, r2, er
}
