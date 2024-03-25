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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksUnbindResource{}
	_ resource.ResourceWithConfigure = &NetworksUnbindResource{}
)

func NewNetworksUnbindResource() resource.Resource {
	return &NetworksUnbindResource{}
}

type NetworksUnbindResource struct {
	client *merakigosdk.Client
}

func (r *NetworksUnbindResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksUnbindResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_unbind"
}

// resourceAction
func (r *NetworksUnbindResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					"product_types": schema.SetAttribute{
						MarkdownDescription: `List of the product types that the network supports`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"tags": schema.SetAttribute{
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
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"retain_configs": schema.BoolAttribute{
						MarkdownDescription: `Optional boolean to retain all the current configs given by the template.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *NetworksUnbindResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksUnbind

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
	// network_id
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Networks.UnbindNetwork(vvNetworkID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UnbindNetwork",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UnbindNetwork",
			err.Error(),
		)
		return
	}
	//Item
	data2 := ResponseNetworksUnbindNetworkItemToBody(data, response)

	diags := resp.State.Set(ctx, &data2)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksUnbindResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksUnbindResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksUnbindResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksUnbind struct {
	NetworkID  types.String                    `tfsdk:"network_id"`
	Item       *ResponseNetworksUnbindNetwork  `tfsdk:"item"`
	Parameters *RequestNetworksUnbindNetworkRs `tfsdk:"parameters"`
}

type ResponseNetworksUnbindNetwork struct {
	EnrollmentString        types.String `tfsdk:"enrollment_string"`
	ID                      types.String `tfsdk:"id"`
	IsBoundToConfigTemplate types.Bool   `tfsdk:"is_bound_to_config_template"`
	Name                    types.String `tfsdk:"name"`
	Notes                   types.String `tfsdk:"notes"`
	OrganizationID          types.String `tfsdk:"organization_id"`
	ProductTypes            types.Set    `tfsdk:"product_types"`
	Tags                    types.Set    `tfsdk:"tags"`
	TimeZone                types.String `tfsdk:"time_zone"`
	URL                     types.String `tfsdk:"url"`
}

type RequestNetworksUnbindNetworkRs struct {
	RetainConfigs types.Bool `tfsdk:"retain_configs"`
}

// FromBody
func (r *NetworksUnbind) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksUnbindNetwork {
	re := *r.Parameters
	retainConfigs := new(bool)
	if !re.RetainConfigs.IsUnknown() && !re.RetainConfigs.IsNull() {
		*retainConfigs = re.RetainConfigs.ValueBool()
	} else {
		retainConfigs = nil
	}
	out := merakigosdk.RequestNetworksUnbindNetwork{
		RetainConfigs: retainConfigs,
	}
	return &out
}

// ToBody
func ResponseNetworksUnbindNetworkItemToBody(state NetworksUnbind, response *merakigosdk.ResponseNetworksUnbindNetwork) NetworksUnbind {
	itemState := ResponseNetworksUnbindNetwork{
		EnrollmentString: types.StringValue(response.EnrollmentString),
		ID:               types.StringValue(response.ID),
		IsBoundToConfigTemplate: func() types.Bool {
			if response.IsBoundToConfigTemplate != nil {
				return types.BoolValue(*response.IsBoundToConfigTemplate)
			}
			return types.Bool{}
		}(),
		Name:           types.StringValue(response.Name),
		Notes:          types.StringValue(response.Notes),
		OrganizationID: types.StringValue(response.OrganizationID),
		ProductTypes:   StringSliceToSet(response.ProductTypes),
		Tags:           StringSliceToSet(response.Tags),
		TimeZone:       types.StringValue(response.TimeZone),
		URL:            types.StringValue(response.URL),
	}
	state.Item = &itemState
	return state
}
