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

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSmDevicesModifyTagsResource{}
	_ resource.ResourceWithConfigure = &NetworksSmDevicesModifyTagsResource{}
)

func NewNetworksSmDevicesModifyTagsResource() resource.Resource {
	return &NetworksSmDevicesModifyTagsResource{}
}

type NetworksSmDevicesModifyTagsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSmDevicesModifyTagsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSmDevicesModifyTagsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_sm_devices_modify_tags"
}

// resourceAction
func (r *NetworksSmDevicesModifyTagsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					"ids": schema.ListAttribute{
						MarkdownDescription: `The ids of the devices to be modified.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"items": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"items": schema.ListNestedAttribute{
									MarkdownDescription: `Array of ResponseSmModifyNetworkSmDevicesTags`,
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
											"serial": schema.StringAttribute{
												MarkdownDescription: `The device serial.`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"tags": schema.ListAttribute{
												MarkdownDescription: `An array of tags associated with the device.`,
												Computed:            true,
												ElementType:         types.StringType,
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
					"scope": schema.ListAttribute{
						MarkdownDescription: `The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be modified.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"serials": schema.ListAttribute{
						MarkdownDescription: `The serials of the devices to be modified.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"tags": schema.ListAttribute{
						MarkdownDescription: `The tags to be added, deleted, or updated.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"update_action": schema.StringAttribute{
						MarkdownDescription: `One of add, delete, or update. Only devices that have been modified will be returned.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"wifi_macs": schema.ListAttribute{
						MarkdownDescription: `The wifiMacs of the devices to be modified.`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *NetworksSmDevicesModifyTagsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSmDevicesModifyTags

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
	response, restyResp1, err := r.client.Sm.ModifyNetworkSmDevicesTags(vvNetworkID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ModifyNetworkSmDevicesTags",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ModifyNetworkSmDevicesTags",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseSmModifyNetworkSmDevicesTagsItemsToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSmDevicesModifyTagsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesModifyTagsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSmDevicesModifyTagsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSmDevicesModifyTags struct {
	NetworkID  types.String                                `tfsdk:"network_id"`
	Items      *[]ResponseItemSmModifyNetworkSmDevicesTags `tfsdk:"items"`
	Parameters *RequestSmModifyNetworkSmDevicesTagsRs      `tfsdk:"parameters"`
}

type ResponseItemSmModifyNetworkSmDevicesTags struct {
	ID      types.String `tfsdk:"id"`
	Serial  types.String `tfsdk:"serial"`
	Tags    types.List   `tfsdk:"tags"`
	WifiMac types.String `tfsdk:"wifi_mac"`
}

type RequestSmModifyNetworkSmDevicesTagsRs struct {
	IDs          types.Set    `tfsdk:"ids"`
	Scope        types.Set    `tfsdk:"scope"`
	Serials      types.Set    `tfsdk:"serials"`
	Tags         types.Set    `tfsdk:"tags"`
	UpdateAction types.String `tfsdk:"update_action"`
	WifiMacs     types.Set    `tfsdk:"wifi_macs"`
}

// FromBody
func (r *NetworksSmDevicesModifyTags) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSmModifyNetworkSmDevicesTags {
	emptyString := ""
	re := *r.Parameters
	var iDs []string = nil
	re.IDs.ElementsAs(ctx, &iDs, false)
	var scope []string = nil
	re.Scope.ElementsAs(ctx, &scope, false)
	var serials []string = nil
	re.Serials.ElementsAs(ctx, &serials, false)
	var tags []string = nil
	re.Tags.ElementsAs(ctx, &tags, false)
	updateAction := new(string)
	if !re.UpdateAction.IsUnknown() && !re.UpdateAction.IsNull() {
		*updateAction = re.UpdateAction.ValueString()
	} else {
		updateAction = &emptyString
	}
	var wifiMacs []string = nil
	re.WifiMacs.ElementsAs(ctx, &wifiMacs, false)
	out := merakigosdk.RequestSmModifyNetworkSmDevicesTags{
		IDs:          iDs,
		Scope:        scope,
		Serials:      serials,
		Tags:         tags,
		UpdateAction: *updateAction,
		WifiMacs:     wifiMacs,
	}
	return &out
}

// ToBody
func ResponseSmModifyNetworkSmDevicesTagsItemsToBody(state NetworksSmDevicesModifyTags, response *merakigosdk.ResponseSmModifyNetworkSmDevicesTags) NetworksSmDevicesModifyTags {
	var items []ResponseItemSmModifyNetworkSmDevicesTags
	for _, item := range *response {
		itemState := ResponseItemSmModifyNetworkSmDevicesTags{
			ID:      types.StringValue(item.ID),
			Serial:  types.StringValue(item.Serial),
			Tags:    StringSliceToList(item.Tags),
			WifiMac: types.StringValue(item.WifiMac),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
