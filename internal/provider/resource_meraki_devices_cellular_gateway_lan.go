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
	_ resource.Resource              = &DevicesCellularGatewayLanResource{}
	_ resource.ResourceWithConfigure = &DevicesCellularGatewayLanResource{}
)

func NewDevicesCellularGatewayLanResource() resource.Resource {
	return &DevicesCellularGatewayLanResource{}
}

type DevicesCellularGatewayLanResource struct {
	client *merakigosdk.Client
}

func (r *DevicesCellularGatewayLanResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesCellularGatewayLanResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_cellular_gateway_lan"
}

func (r *DevicesCellularGatewayLanResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"device_lan_ip": schema.StringAttribute{
				MarkdownDescription: `Lan IP of the MG`,
				Computed:            true,
			},
			"device_name": schema.StringAttribute{
				MarkdownDescription: `Name of the MG.`,
				Computed:            true,
			},
			"device_subnet": schema.StringAttribute{
				MarkdownDescription: `Subnet configuration of the MG.`,
				Computed:            true,
			},
			"fixed_ip_assignments": schema.SetNestedAttribute{
				MarkdownDescription: `list of all fixed IP assignments for a single MG`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"ip": schema.StringAttribute{
							MarkdownDescription: `The IP address you want to assign to a specific server or device`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"mac": schema.StringAttribute{
							MarkdownDescription: `The MAC address of the server or device that hosts the internal resource that you wish to receive the specified IP address`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `A descriptive name of the assignment`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"reserved_ip_ranges": schema.SetNestedAttribute{
				MarkdownDescription: `list of all reserved IP ranges for a single MG`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"comment": schema.StringAttribute{
							MarkdownDescription: `Comment explaining the reserved IP range`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"end": schema.StringAttribute{
							MarkdownDescription: `Ending IP included in the reserved range of IPs`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"start": schema.StringAttribute{
							MarkdownDescription: `Starting IP included in the reserved range of IPs`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
		},
	}
}

func (r *DevicesCellularGatewayLanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesCellularGatewayLanRs

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
	//Has Item and not has items

	if vvSerial != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.CellularGateway.GetDeviceCellularGatewayLan(vvSerial)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource DevicesCellularGatewayLan  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource DevicesCellularGatewayLan only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.CellularGateway.UpdateDeviceCellularGatewayLan(vvSerial, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCellularGatewayLan",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCellularGatewayLan",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.CellularGateway.GetDeviceCellularGatewayLan(vvSerial)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCellularGatewayLan",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCellularGatewayLan",
			err.Error(),
		)
		return
	}

	data = ResponseCellularGatewayGetDeviceCellularGatewayLanItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *DevicesCellularGatewayLanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesCellularGatewayLanRs

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
	responseGet, restyRespGet, err := r.client.CellularGateway.GetDeviceCellularGatewayLan(vvSerial)
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
				"Failure when executing GetDeviceCellularGatewayLan",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCellularGatewayLan",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCellularGatewayGetDeviceCellularGatewayLanItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesCellularGatewayLanResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesCellularGatewayLanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesCellularGatewayLanRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.CellularGateway.UpdateDeviceCellularGatewayLan(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCellularGatewayLan",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCellularGatewayLan",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCellularGatewayLanResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesCellularGatewayLan", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesCellularGatewayLanRs struct {
	Serial             types.String                                                              `tfsdk:"serial"`
	DeviceLanIP        types.String                                                              `tfsdk:"device_lan_ip"`
	DeviceName         types.String                                                              `tfsdk:"device_name"`
	DeviceSubnet       types.String                                                              `tfsdk:"device_subnet"`
	FixedIPAssignments *[]ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIpAssignmentsRs `tfsdk:"fixed_ip_assignments"`
	ReservedIPRanges   *[]ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIpRangesRs   `tfsdk:"reserved_ip_ranges"`
}

type ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIpAssignmentsRs struct {
	IP   types.String `tfsdk:"ip"`
	Mac  types.String `tfsdk:"mac"`
	Name types.String `tfsdk:"name"`
}

type ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIpRangesRs struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// FromBody
func (r *DevicesCellularGatewayLanRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayLan {
	var requestCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments []merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments

	if r.FixedIPAssignments != nil {
		for _, rItem1 := range *r.FixedIPAssignments {
			ip := rItem1.IP.ValueString()
			mac := rItem1.Mac.ValueString()
			name := rItem1.Name.ValueString()
			requestCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments = append(requestCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments, merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments{
				IP:   ip,
				Mac:  mac,
				Name: name,
			})
			//[debug] Is Array: True
		}
	}
	var requestCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges []merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges

	if r.ReservedIPRanges != nil {
		for _, rItem1 := range *r.ReservedIPRanges {
			comment := rItem1.Comment.ValueString()
			end := rItem1.End.ValueString()
			start := rItem1.Start.ValueString()
			requestCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges = append(requestCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges, merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges{
				Comment: comment,
				End:     end,
				Start:   start,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayLan{
		FixedIPAssignments: func() *[]merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments {
			if len(requestCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments) > 0 {
				return &requestCellularGatewayUpdateDeviceCellularGatewayLanFixedIPAssignments
			}
			return nil
		}(),
		ReservedIPRanges: func() *[]merakigosdk.RequestCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges {
			if len(requestCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges) > 0 {
				return &requestCellularGatewayUpdateDeviceCellularGatewayLanReservedIPRanges
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCellularGatewayGetDeviceCellularGatewayLanItemToBodyRs(state DevicesCellularGatewayLanRs, response *merakigosdk.ResponseCellularGatewayGetDeviceCellularGatewayLan, is_read bool) DevicesCellularGatewayLanRs {
	itemState := DevicesCellularGatewayLanRs{
		DeviceLanIP:  types.StringValue(response.DeviceLanIP),
		DeviceName:   types.StringValue(response.DeviceName),
		DeviceSubnet: types.StringValue(response.DeviceSubnet),
		FixedIPAssignments: func() *[]ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIpAssignmentsRs {
			if response.FixedIPAssignments != nil {
				result := make([]ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIpAssignmentsRs, len(*response.FixedIPAssignments))
				for i, fixedIPAssignments := range *response.FixedIPAssignments {
					result[i] = ResponseCellularGatewayGetDeviceCellularGatewayLanFixedIpAssignmentsRs{
						IP:   types.StringValue(fixedIPAssignments.IP),
						Mac:  types.StringValue(fixedIPAssignments.Mac),
						Name: types.StringValue(fixedIPAssignments.Name),
					}
				}
				return &result
			}
			return nil
		}(),
		ReservedIPRanges: func() *[]ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIpRangesRs {
			if response.ReservedIPRanges != nil {
				result := make([]ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIpRangesRs, len(*response.ReservedIPRanges))
				for i, reservedIPRanges := range *response.ReservedIPRanges {
					result[i] = ResponseCellularGatewayGetDeviceCellularGatewayLanReservedIpRangesRs{
						Comment: types.StringValue(reservedIPRanges.Comment),
						End:     types.StringValue(reservedIPRanges.End),
						Start:   types.StringValue(reservedIPRanges.Start),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesCellularGatewayLanRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesCellularGatewayLanRs)
}
