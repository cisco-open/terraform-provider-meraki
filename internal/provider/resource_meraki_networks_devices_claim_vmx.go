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
	_ resource.Resource              = &NetworksDevicesClaimVmxResource{}
	_ resource.ResourceWithConfigure = &NetworksDevicesClaimVmxResource{}
)

func NewNetworksDevicesClaimVmxResource() resource.Resource {
	return &NetworksDevicesClaimVmxResource{}
}

type NetworksDevicesClaimVmxResource struct {
	client *merakigosdk.Client
}

func (r *NetworksDevicesClaimVmxResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksDevicesClaimVmxResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_devices_claim_vmx"
}

// resourceAction
func (r *NetworksDevicesClaimVmxResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"address": schema.StringAttribute{
						MarkdownDescription: `Physical address of the device`,
						Computed:            true,
					},
					"details": schema.SetNestedAttribute{
						MarkdownDescription: `Additional device information`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"name": schema.StringAttribute{
									MarkdownDescription: `Additional property name`,
									Computed:            true,
								},
								"value": schema.StringAttribute{
									MarkdownDescription: `Additional property value`,
									Computed:            true,
								},
							},
						},
					},
					"firmware": schema.StringAttribute{
						MarkdownDescription: `Firmware version of the device`,
						Computed:            true,
					},
					"imei": schema.StringAttribute{
						MarkdownDescription: `IMEI of the device, if applicable`,
						Computed:            true,
					},
					"lan_ip": schema.StringAttribute{
						MarkdownDescription: `LAN IP address of the device`,
						Computed:            true,
					},
					"lat": schema.Float64Attribute{
						MarkdownDescription: `Latitude of the device`,
						Computed:            true,
					},
					"lng": schema.Float64Attribute{
						MarkdownDescription: `Longitude of the device`,
						Computed:            true,
					},
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
					"network_id": schema.StringAttribute{
						MarkdownDescription: `ID of the network the device belongs to`,
						Computed:            true,
					},
					"notes": schema.StringAttribute{
						MarkdownDescription: `Notes for the device, limited to 255 characters`,
						Computed:            true,
					},
					"product_type": schema.StringAttribute{
						MarkdownDescription: `Product type of the device`,
						Computed:            true,
					},
					"serial": schema.StringAttribute{
						MarkdownDescription: `Serial number of the device`,
						Computed:            true,
					},
					"tags": schema.ListAttribute{
						MarkdownDescription: `List of tags assigned to the device`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"size": schema.StringAttribute{
						MarkdownDescription: `The size of the vMX you claim. It can be one of: small, medium, large, xlarge, 100
                                        Allowed values: [100,large,medium,small,xlarge]`,
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *NetworksDevicesClaimVmxResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksDevicesClaimVmx

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Networks.VmxNetworkDevicesClaim(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing VmxNetworkDevicesClaim",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing VmxNetworkDevicesClaim",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksVmxNetworkDevicesClaimItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksDevicesClaimVmxResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksDevicesClaimVmxResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksDevicesClaimVmxResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksDevicesClaimVmx struct {
	NetworkID  types.String                             `tfsdk:"network_id"`
	Item       *ResponseNetworksVmxNetworkDevicesClaim  `tfsdk:"item"`
	Parameters *RequestNetworksVmxNetworkDevicesClaimRs `tfsdk:"parameters"`
}

type ResponseNetworksVmxNetworkDevicesClaim struct {
	Address     types.String                                     `tfsdk:"address"`
	Details     *[]ResponseNetworksVmxNetworkDevicesClaimDetails `tfsdk:"details"`
	Firmware    types.String                                     `tfsdk:"firmware"`
	Imei        types.String                                     `tfsdk:"imei"`
	LanIP       types.String                                     `tfsdk:"lan_ip"`
	Lat         types.Float64                                    `tfsdk:"lat"`
	Lng         types.Float64                                    `tfsdk:"lng"`
	Mac         types.String                                     `tfsdk:"mac"`
	Model       types.String                                     `tfsdk:"model"`
	Name        types.String                                     `tfsdk:"name"`
	NetworkID   types.String                                     `tfsdk:"network_id"`
	Notes       types.String                                     `tfsdk:"notes"`
	ProductType types.String                                     `tfsdk:"product_type"`
	Serial      types.String                                     `tfsdk:"serial"`
	Tags        types.List                                       `tfsdk:"tags"`
}

type ResponseNetworksVmxNetworkDevicesClaimDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type RequestNetworksVmxNetworkDevicesClaimRs struct {
	Size types.String `tfsdk:"size"`
}

// FromBody
func (r *NetworksDevicesClaimVmx) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksVmxNetworkDevicesClaim {
	emptyString := ""
	re := *r.Parameters
	size := new(string)
	if !re.Size.IsUnknown() && !re.Size.IsNull() {
		*size = re.Size.ValueString()
	} else {
		size = &emptyString
	}
	out := merakigosdk.RequestNetworksVmxNetworkDevicesClaim{
		Size: *size,
	}
	return &out
}

// ToBody
func ResponseNetworksVmxNetworkDevicesClaimItemToBody(state NetworksDevicesClaimVmx, response *merakigosdk.ResponseNetworksVmxNetworkDevicesClaim) NetworksDevicesClaimVmx {
	itemState := ResponseNetworksVmxNetworkDevicesClaim{
		Address: func() types.String {
			if response.Address != "" {
				return types.StringValue(response.Address)
			}
			return types.String{}
		}(),
		Details: func() *[]ResponseNetworksVmxNetworkDevicesClaimDetails {
			if response.Details != nil {
				result := make([]ResponseNetworksVmxNetworkDevicesClaimDetails, len(*response.Details))
				for i, details := range *response.Details {
					result[i] = ResponseNetworksVmxNetworkDevicesClaimDetails{
						Name: func() types.String {
							if details.Name != "" {
								return types.StringValue(details.Name)
							}
							return types.String{}
						}(),
						Value: func() types.String {
							if details.Value != "" {
								return types.StringValue(details.Value)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Firmware: func() types.String {
			if response.Firmware != "" {
				return types.StringValue(response.Firmware)
			}
			return types.String{}
		}(),
		Imei: func() types.String {
			if response.Imei != "" {
				return types.StringValue(response.Imei)
			}
			return types.String{}
		}(),
		LanIP: func() types.String {
			if response.LanIP != "" {
				return types.StringValue(response.LanIP)
			}
			return types.String{}
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
			return types.String{}
		}(),
		Model: func() types.String {
			if response.Model != "" {
				return types.StringValue(response.Model)
			}
			return types.String{}
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
		NetworkID: func() types.String {
			if response.NetworkID != "" {
				return types.StringValue(response.NetworkID)
			}
			return types.String{}
		}(),
		Notes: func() types.String {
			if response.Notes != "" {
				return types.StringValue(response.Notes)
			}
			return types.String{}
		}(),
		ProductType: func() types.String {
			if response.ProductType != "" {
				return types.StringValue(response.ProductType)
			}
			return types.String{}
		}(),
		Serial: func() types.String {
			if response.Serial != "" {
				return types.StringValue(response.Serial)
			}
			return types.String{}
		}(),
		Tags: StringSliceToList(response.Tags),
	}
	state.Item = &itemState
	return state
}
