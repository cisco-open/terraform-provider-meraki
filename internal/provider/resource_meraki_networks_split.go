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
	_ resource.Resource              = &NetworksSplitResource{}
	_ resource.ResourceWithConfigure = &NetworksSplitResource{}
)

func NewNetworksSplitResource() resource.Resource {
	return &NetworksSplitResource{}
}

type NetworksSplitResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSplitResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSplitResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_split"
}

// resourceAction
func (r *NetworksSplitResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"resulting_networks": schema.SetNestedAttribute{
						MarkdownDescription: `Networks after the split`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"enrollment_string": schema.StringAttribute{
									MarkdownDescription: `Enrollment string for the network`,
									Computed:            true,
								},
								"id": schema.StringAttribute{
									MarkdownDescription: `Network ID`,
									Computed:            true,
								},
								"is_bound_to_config_template": schema.BoolAttribute{
									MarkdownDescription: `If the network is bound to a config template`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `Network name`,
									Computed:            true,
								},
								"notes": schema.StringAttribute{
									MarkdownDescription: `Notes for the network`,
									Computed:            true,
								},
								"organization_id": schema.StringAttribute{
									MarkdownDescription: `Organization ID`,
									Computed:            true,
								},
								"product_types": schema.ListAttribute{
									MarkdownDescription: `List of the product types that the network supports`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"tags": schema.ListAttribute{
									MarkdownDescription: `Network tags`,
									Computed:            true,
									ElementType:         types.StringType,
								},
								"time_zone": schema.StringAttribute{
									MarkdownDescription: `Timezone of the network`,
									Computed:            true,
								},
								"url": schema.StringAttribute{
									MarkdownDescription: `URL to the network Dashboard UI`,
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
func (r *NetworksSplitResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSplit

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
	response, restyResp1, err := r.client.Networks.SplitNetwork(vvNetworkID)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing SplitNetwork",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing SplitNetwork",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseNetworksSplitNetworkItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSplitResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSplitResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksSplitResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSplit struct {
	NetworkID  types.String                   `tfsdk:"network_id"`
	Item       *ResponseNetworksSplitNetwork  `tfsdk:"item"`
	Parameters *RequestNetworksSplitNetworkRs `tfsdk:"parameters"`
}

type ResponseNetworksSplitNetwork struct {
	ResultingNetworks *[]ResponseNetworksSplitNetworkResultingNetworks `tfsdk:"resulting_networks"`
}

type ResponseNetworksSplitNetworkResultingNetworks struct {
	EnrollmentString        types.String `tfsdk:"enrollment_string"`
	ID                      types.String `tfsdk:"id"`
	IsBoundToConfigTemplate types.Bool   `tfsdk:"is_bound_to_config_template"`
	Name                    types.String `tfsdk:"name"`
	Notes                   types.String `tfsdk:"notes"`
	OrganizationID          types.String `tfsdk:"organization_id"`
	ProductTypes            types.List   `tfsdk:"product_types"`
	Tags                    types.List   `tfsdk:"tags"`
	TimeZone                types.String `tfsdk:"time_zone"`
	URL                     types.String `tfsdk:"url"`
}

type RequestNetworksSplitNetworkRs interface{}

// FromBody
// ToBody
func ResponseNetworksSplitNetworkItemToBody(state NetworksSplit, response *merakigosdk.ResponseNetworksSplitNetwork) NetworksSplit {
	itemState := ResponseNetworksSplitNetwork{
		ResultingNetworks: func() *[]ResponseNetworksSplitNetworkResultingNetworks {
			if response.ResultingNetworks != nil {
				result := make([]ResponseNetworksSplitNetworkResultingNetworks, len(*response.ResultingNetworks))
				for i, resultingNetworks := range *response.ResultingNetworks {
					result[i] = ResponseNetworksSplitNetworkResultingNetworks{
						EnrollmentString: types.StringValue(resultingNetworks.EnrollmentString),
						ID:               types.StringValue(resultingNetworks.ID),
						IsBoundToConfigTemplate: func() types.Bool {
							if resultingNetworks.IsBoundToConfigTemplate != nil {
								return types.BoolValue(*resultingNetworks.IsBoundToConfigTemplate)
							}
							return types.Bool{}
						}(),
						Name:           types.StringValue(resultingNetworks.Name),
						Notes:          types.StringValue(resultingNetworks.Notes),
						OrganizationID: types.StringValue(resultingNetworks.OrganizationID),
						ProductTypes:   StringSliceToList(resultingNetworks.ProductTypes),
						Tags:           StringSliceToList(resultingNetworks.Tags),
						TimeZone:       types.StringValue(resultingNetworks.TimeZone),
						URL:            types.StringValue(resultingNetworks.URL),
					}
				}
				return &result
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
