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
	_ resource.Resource              = &NetworksSmDevicesFieldsResource{}
	_ resource.ResourceWithConfigure = &NetworksSmDevicesFieldsResource{}
)

func NewNetworksSmDevicesFieldsResource() resource.Resource {
	return &NetworksSmDevicesFieldsResource{}
}

type NetworksSmDevicesFieldsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSmDevicesFieldsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSmDevicesFieldsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_fields"
}

// resourceAction
func (r *NetworksSmDevicesFieldsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"device_fields": schema.SingleNestedAttribute{
						MarkdownDescription: `The new fields of the device. Each field of this object is optional.`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"name": schema.StringAttribute{
								MarkdownDescription: `New name for the device`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
							"notes": schema.StringAttribute{
								MarkdownDescription: `New notes for the device`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
						},
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `The id of the device to be modified.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"items": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"items": schema.ListNestedAttribute{
									MarkdownDescription: `Array of ResponseSmUpdateNetworkSmDevicesFields`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"id": schema.StringAttribute{
												MarkdownDescription: `The Meraki Id of the device record.`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"name": schema.StringAttribute{
												MarkdownDescription: `The name of the device.`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"notes": schema.StringAttribute{
												MarkdownDescription: `Notes associated with the device.`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"serial": schema.StringAttribute{
												MarkdownDescription: `The device serial.`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"wifi_mac": schema.StringAttribute{
												MarkdownDescription: `The MAC of the device.`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
										},
									},
								},
							},
						},
					},
					"serial": schema.StringAttribute{
						MarkdownDescription: `The serial of the device to be modified.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"wifi_mac": schema.StringAttribute{
						MarkdownDescription: `The wifiMac of the device to be modified.`,
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
func (r *NetworksSmDevicesFieldsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSmDevicesFields

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
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp1, err := r.client.Sm.UpdateNetworkSmDevicesFields(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSmDevicesFields",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSmDevicesFields",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseSmUpdateNetworkSmDevicesFieldsItemsToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSmDevicesFieldsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesFieldsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesFieldsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSmDevicesFields struct {
	NetworkID  types.String                                  `tfsdk:"network_id"`
	Items      *[]ResponseItemSmUpdateNetworkSmDevicesFields `tfsdk:"items"`
	Parameters *RequestSmUpdateNetworkSmDevicesFieldsRs      `tfsdk:"parameters"`
}

type ResponseItemSmUpdateNetworkSmDevicesFields struct {
	ID      types.String `tfsdk:"id"`
	Name    types.String `tfsdk:"name"`
	Notes   types.String `tfsdk:"notes"`
	Serial  types.String `tfsdk:"serial"`
	WifiMac types.String `tfsdk:"wifi_mac"`
}

type RequestSmUpdateNetworkSmDevicesFieldsRs struct {
	DeviceFields *RequestSmUpdateNetworkSmDevicesFieldsDeviceFieldsRs `tfsdk:"device_fields"`
	ID           types.String                                         `tfsdk:"id"`
	Serial       types.String                                         `tfsdk:"serial"`
	WifiMac      types.String                                         `tfsdk:"wifi_mac"`
}

type RequestSmUpdateNetworkSmDevicesFieldsDeviceFieldsRs struct {
	Name  types.String `tfsdk:"name"`
	Notes types.String `tfsdk:"notes"`
}

// FromBody
func (r *NetworksSmDevicesFields) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSmUpdateNetworkSmDevicesFields {
	emptyString := ""
	re := *r.Parameters
	var requestSmUpdateNetworkSmDevicesFieldsDeviceFields *merakigosdk.RequestSmUpdateNetworkSmDevicesFieldsDeviceFields

	if re.DeviceFields != nil {
		name := re.DeviceFields.Name.ValueString()
		notes := re.DeviceFields.Notes.ValueString()
		requestSmUpdateNetworkSmDevicesFieldsDeviceFields = &merakigosdk.RequestSmUpdateNetworkSmDevicesFieldsDeviceFields{
			Name:  name,
			Notes: notes,
		}
		//[debug] Is Array: False
	}
	iD := new(string)
	if !re.ID.IsUnknown() && !re.ID.IsNull() {
		*iD = re.ID.ValueString()
	} else {
		iD = &emptyString
	}
	serial := new(string)
	if !re.Serial.IsUnknown() && !re.Serial.IsNull() {
		*serial = re.Serial.ValueString()
	} else {
		serial = &emptyString
	}
	wifiMac := new(string)
	if !re.WifiMac.IsUnknown() && !re.WifiMac.IsNull() {
		*wifiMac = re.WifiMac.ValueString()
	} else {
		wifiMac = &emptyString
	}
	out := merakigosdk.RequestSmUpdateNetworkSmDevicesFields{
		DeviceFields: requestSmUpdateNetworkSmDevicesFieldsDeviceFields,
		ID:           *iD,
		Serial:       *serial,
		WifiMac:      *wifiMac,
	}
	return &out
}

// ToBody
func ResponseSmUpdateNetworkSmDevicesFieldsItemsToBody(state NetworksSmDevicesFields, response *merakigosdk.ResponseSmUpdateNetworkSmDevicesFields) NetworksSmDevicesFields {
	var items []ResponseItemSmUpdateNetworkSmDevicesFields
	for _, item := range *response {
		itemState := ResponseItemSmUpdateNetworkSmDevicesFields{
			ID: func() types.String {
				if item.ID != "" {
					return types.StringValue(item.ID)
				}
				return types.String{}
			}(),
			Name: func() types.String {
				if item.Name != "" {
					return types.StringValue(item.Name)
				}
				return types.String{}
			}(),
			Notes: func() types.String {
				if item.Notes != "" {
					return types.StringValue(item.Notes)
				}
				return types.String{}
			}(),
			Serial: func() types.String {
				if item.Serial != "" {
					return types.StringValue(item.Serial)
				}
				return types.String{}
			}(),
			WifiMac: func() types.String {
				if item.WifiMac != "" {
					return types.StringValue(item.WifiMac)
				}
				return types.String{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
