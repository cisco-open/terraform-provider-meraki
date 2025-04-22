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

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksVLANProfilesAssignmentsByDeviceDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksVLANProfilesAssignmentsByDeviceDataSource{}
)

func NewNetworksVLANProfilesAssignmentsByDeviceDataSource() datasource.DataSource {
	return &NetworksVLANProfilesAssignmentsByDeviceDataSource{}
}

type NetworksVLANProfilesAssignmentsByDeviceDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksVLANProfilesAssignmentsByDeviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksVLANProfilesAssignmentsByDeviceDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_vlan_profiles_assignments_by_device"
}

func (d *NetworksVLANProfilesAssignmentsByDeviceDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ending_before": schema.StringAttribute{
				MarkdownDescription: `endingBefore query parameter. A token used by the server to indicate the end of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"per_page": schema.Int64Attribute{
				MarkdownDescription: `perPage query parameter. The number of entries per page returned. Acceptable range is 3 1000. Default is 1000.`,
				Optional:            true,
			},
			"product_types": schema.ListAttribute{
				MarkdownDescription: `productTypes query parameter. Optional parameter to filter devices by product types.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"serials": schema.ListAttribute{
				MarkdownDescription: `serials query parameter. Optional parameter to filter devices by serials. All devices returned belong to serial numbers that are an exact match.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"stack_ids": schema.ListAttribute{
				MarkdownDescription: `stackIds query parameter. Optional parameter to filter devices by Switch Stack ids.`,
				Optional:            true,
				ElementType:         types.StringType,
			},
			"starting_after": schema.StringAttribute{
				MarkdownDescription: `startingAfter query parameter. A token used by the server to indicate the start of the page. Often this is a timestamp or an ID but it is not limited to those. This parameter should not be defined by client applications. The link for the first, last, prev, or next page in the HTTP Link header should define it.`,
				Optional:            true,
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseNetworksGetNetworkVlanProfilesAssignmentsByDevice`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"mac": schema.StringAttribute{
							MarkdownDescription: `MAC address of the device`,
							Computed:            true,
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `Name of the Device`,
							Computed:            true,
						},
						"product_type": schema.StringAttribute{
							MarkdownDescription: `The product type`,
							Computed:            true,
						},
						"serial": schema.StringAttribute{
							MarkdownDescription: `Serial of the Device`,
							Computed:            true,
						},
						"stack": schema.SingleNestedAttribute{
							MarkdownDescription: `The Switch Stack the device belongs to`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"id": schema.StringAttribute{
									MarkdownDescription: `ID of the Switch Stack`,
									Computed:            true,
								},
							},
						},
						"vlan_profile": schema.SingleNestedAttribute{
							MarkdownDescription: `The VLAN Profile`,
							Computed:            true,
							Attributes: map[string]schema.Attribute{

								"iname": schema.StringAttribute{
									MarkdownDescription: `IName of the VLAN Profile`,
									Computed:            true,
								},
								"is_default": schema.BoolAttribute{
									MarkdownDescription: `Is this VLAN profile the default for the network?`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Name of the VLAN Profile`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksVLANProfilesAssignmentsByDeviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksVLANProfilesAssignmentsByDevice NetworksVLANProfilesAssignmentsByDevice
	diags := req.Config.Get(ctx, &networksVLANProfilesAssignmentsByDevice)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkVLANProfilesAssignmentsByDevice")
		vvNetworkID := networksVLANProfilesAssignmentsByDevice.NetworkID.ValueString()
		queryParams1 := merakigosdk.GetNetworkVLANProfilesAssignmentsByDeviceQueryParams{}

		queryParams1.PerPage = int(networksVLANProfilesAssignmentsByDevice.PerPage.ValueInt64())
		queryParams1.StartingAfter = networksVLANProfilesAssignmentsByDevice.StartingAfter.ValueString()
		queryParams1.EndingBefore = networksVLANProfilesAssignmentsByDevice.EndingBefore.ValueString()
		queryParams1.Serials = elementsToStrings(ctx, networksVLANProfilesAssignmentsByDevice.Serials)
		queryParams1.ProductTypes = elementsToStrings(ctx, networksVLANProfilesAssignmentsByDevice.ProductTypes)
		queryParams1.StackIDs = elementsToStrings(ctx, networksVLANProfilesAssignmentsByDevice.StackIDs)

		// has_unknown_response: None

		response1, restyResp1, err := d.client.Networks.GetNetworkVLANProfilesAssignmentsByDevice(vvNetworkID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkVLANProfilesAssignmentsByDevice",
				err.Error(),
			)
			return
		}

		networksVLANProfilesAssignmentsByDevice = ResponseNetworksGetNetworkVLANProfilesAssignmentsByDeviceItemsToBody(networksVLANProfilesAssignmentsByDevice, response1)
		diags = resp.State.Set(ctx, &networksVLANProfilesAssignmentsByDevice)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksVLANProfilesAssignmentsByDevice struct {
	NetworkID     types.String                                                     `tfsdk:"network_id"`
	PerPage       types.Int64                                                      `tfsdk:"per_page"`
	StartingAfter types.String                                                     `tfsdk:"starting_after"`
	EndingBefore  types.String                                                     `tfsdk:"ending_before"`
	Serials       types.List                                                       `tfsdk:"serials"`
	ProductTypes  types.List                                                       `tfsdk:"product_types"`
	StackIDs      types.List                                                       `tfsdk:"stack_ids"`
	Items         *[]ResponseItemNetworksGetNetworkVlanProfilesAssignmentsByDevice `tfsdk:"items"`
}

type ResponseItemNetworksGetNetworkVlanProfilesAssignmentsByDevice struct {
	Mac         types.String                                                              `tfsdk:"mac"`
	Name        types.String                                                              `tfsdk:"name"`
	ProductType types.String                                                              `tfsdk:"product_type"`
	Serial      types.String                                                              `tfsdk:"serial"`
	Stack       *ResponseItemNetworksGetNetworkVlanProfilesAssignmentsByDeviceStack       `tfsdk:"stack"`
	VLANProfile *ResponseItemNetworksGetNetworkVlanProfilesAssignmentsByDeviceVlanProfile `tfsdk:"vlan_profile"`
}

type ResponseItemNetworksGetNetworkVlanProfilesAssignmentsByDeviceStack struct {
	ID types.String `tfsdk:"id"`
}

type ResponseItemNetworksGetNetworkVlanProfilesAssignmentsByDeviceVlanProfile struct {
	Iname     types.String `tfsdk:"iname"`
	IsDefault types.Bool   `tfsdk:"is_default"`
	Name      types.String `tfsdk:"name"`
}

// ToBody
func ResponseNetworksGetNetworkVLANProfilesAssignmentsByDeviceItemsToBody(state NetworksVLANProfilesAssignmentsByDevice, response *merakigosdk.ResponseNetworksGetNetworkVLANProfilesAssignmentsByDevice) NetworksVLANProfilesAssignmentsByDevice {
	var items []ResponseItemNetworksGetNetworkVlanProfilesAssignmentsByDevice
	for _, item := range *response {
		itemState := ResponseItemNetworksGetNetworkVlanProfilesAssignmentsByDevice{
			Mac:         types.StringValue(item.Mac),
			Name:        types.StringValue(item.Name),
			ProductType: types.StringValue(item.ProductType),
			Serial:      types.StringValue(item.Serial),
			Stack: func() *ResponseItemNetworksGetNetworkVlanProfilesAssignmentsByDeviceStack {
				if item.Stack != nil {
					return &ResponseItemNetworksGetNetworkVlanProfilesAssignmentsByDeviceStack{
						ID: types.StringValue(item.Stack.ID),
					}
				}
				return nil
			}(),
			VLANProfile: func() *ResponseItemNetworksGetNetworkVlanProfilesAssignmentsByDeviceVlanProfile {
				if item.VLANProfile != nil {
					return &ResponseItemNetworksGetNetworkVlanProfilesAssignmentsByDeviceVlanProfile{
						Iname: types.StringValue(item.VLANProfile.Iname),
						IsDefault: func() types.Bool {
							if item.VLANProfile.IsDefault != nil {
								return types.BoolValue(*item.VLANProfile.IsDefault)
							}
							return types.Bool{}
						}(),
						Name: types.StringValue(item.VLANProfile.Name),
					}
				}
				return nil
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
