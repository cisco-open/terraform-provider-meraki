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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksDevicesClaimResource{}
	_ resource.ResourceWithConfigure = &NetworksDevicesClaimResource{}
)

func NewNetworksDevicesClaimResource() resource.Resource {
	return &NetworksDevicesClaimResource{}
}

type NetworksDevicesClaimResource struct {
	client *merakigosdk.Client
}

func (r *NetworksDevicesClaimResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksDevicesClaimResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_devices_claim"
}

// resourceAction
func (r *NetworksDevicesClaimResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"add_atomically": schema.BoolAttribute{
				MarkdownDescription: `addAtomically query parameter. Whether to claim devices atomically. If true, all devices will be claimed or none will be claimed. Default is true.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
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

					"errors": schema.SetNestedAttribute{
						MarkdownDescription: `Errors for devices that were not added`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"errors": schema.ListAttribute{
									MarkdownDescription: `The errors for the device`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"serial": schema.StringAttribute{
									MarkdownDescription: `The serial of the device`,
									Computed:            true,
								},
							},
						},
					},
					"serials": schema.ListAttribute{
						MarkdownDescription: `The serials of the devices`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"serials": schema.ListAttribute{
						MarkdownDescription: `A list of serials of devices to claim`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *NetworksDevicesClaimResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksDevicesClaim

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
	response, restyResp1, err := r.client.Networks.ClaimNetworkDevices(vvNetworkID, dataRequest, nil)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ClaimNetworkDevices",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ClaimNetworkDevices",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksClaimNetworkDevicesItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksDevicesClaimResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksDevicesClaimResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksDevicesClaimResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksDevicesClaim struct {
	NetworkID     types.String                          `tfsdk:"network_id"`
	Item          *ResponseNetworksClaimNetworkDevices  `tfsdk:"item"`
	Parameters    *RequestNetworksClaimNetworkDevicesRs `tfsdk:"parameters"`
	AddAtomically types.Bool                            `tfsdk:"add_atomically"`
}

type ResponseNetworksClaimNetworkDevices struct {
	Errors  *[]ResponseNetworksClaimNetworkDevicesErrors `tfsdk:"errors"`
	Serials types.List                                   `tfsdk:"serials"`
}

type ResponseNetworksClaimNetworkDevicesErrors struct {
	Errors types.List   `tfsdk:"errors"`
	Serial types.String `tfsdk:"serial"`
}

type RequestNetworksClaimNetworkDevicesRs struct {
	Serials types.List `tfsdk:"serials"`
}

// FromBody
func (r *NetworksDevicesClaim) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksClaimNetworkDevices {
	re := *r.Parameters
	var serials []string = nil
	re.Serials.ElementsAs(ctx, &serials, false)
	out := merakigosdk.RequestNetworksClaimNetworkDevices{
		Serials: serials,
	}
	return &out
}

// ToBody
func ResponseNetworksClaimNetworkDevicesItemToBody(state NetworksDevicesClaim, response *merakigosdk.ResponseNetworksClaimNetworkDevices) NetworksDevicesClaim {
	itemState := ResponseNetworksClaimNetworkDevices{
		Errors: func() *[]ResponseNetworksClaimNetworkDevicesErrors {
			if response.Errors != nil {
				result := make([]ResponseNetworksClaimNetworkDevicesErrors, len(*response.Errors))
				for i, errors := range *response.Errors {
					result[i] = ResponseNetworksClaimNetworkDevicesErrors{
						Errors: StringSliceToList(errors.Errors),
						Serial: func() types.String {
							if errors.Serial != "" {
								return types.StringValue(errors.Serial)
							}
							return types.String{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		Serials: StringSliceToList(response.Serials),
	}
	state.Item = &itemState
	return state
}
