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
	_ resource.Resource              = &OrganizationsNetworksCombineResource{}
	_ resource.ResourceWithConfigure = &OrganizationsNetworksCombineResource{}
)

func NewOrganizationsNetworksCombineResource() resource.Resource {
	return &OrganizationsNetworksCombineResource{}
}

type OrganizationsNetworksCombineResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsNetworksCombineResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsNetworksCombineResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_networks_combine"
}

// resourceAction
func (r *OrganizationsNetworksCombineResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"resulting_network": schema.SingleNestedAttribute{
						MarkdownDescription: `Network after the combination`,
						Computed:            true,
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
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"enrollment_string": schema.StringAttribute{
						MarkdownDescription: `A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break. All networks that are part of this combined network will have their enrollment string appended by '-network_type'. If left empty, all exisitng enrollment strings will be deleted.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the combined network`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"network_ids": schema.ListAttribute{
						MarkdownDescription: `A list of the network IDs that will be combined. If an ID of a combined network is included in this list, the other networks in the list will be grouped into that network`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *OrganizationsNetworksCombineResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsNetworksCombine

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
	vvOrganizationID := data.OrganizationID.ValueString()
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Organizations.CombineOrganizationNetworks(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CombineOrganizationNetworks",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CombineOrganizationNetworks",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseOrganizationsCombineOrganizationNetworksItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsNetworksCombineResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsNetworksCombineResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsNetworksCombineResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsNetworksCombine struct {
	OrganizationID types.String                                       `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsCombineOrganizationNetworks  `tfsdk:"item"`
	Parameters     *RequestOrganizationsCombineOrganizationNetworksRs `tfsdk:"parameters"`
}

type ResponseOrganizationsCombineOrganizationNetworks struct {
	ResultingNetwork *ResponseOrganizationsCombineOrganizationNetworksResultingNetwork `tfsdk:"resulting_network"`
}

type ResponseOrganizationsCombineOrganizationNetworksResultingNetwork struct {
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

type RequestOrganizationsCombineOrganizationNetworksRs struct {
	EnrollmentString types.String `tfsdk:"enrollment_string"`
	Name             types.String `tfsdk:"name"`
	NetworkIDs       types.List   `tfsdk:"network_ids"`
}

// FromBody
func (r *OrganizationsNetworksCombine) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCombineOrganizationNetworks {
	emptyString := ""
	re := *r.Parameters
	enrollmentString := new(string)
	if !re.EnrollmentString.IsUnknown() && !re.EnrollmentString.IsNull() {
		*enrollmentString = re.EnrollmentString.ValueString()
	} else {
		enrollmentString = &emptyString
	}
	name := new(string)
	if !re.Name.IsUnknown() && !re.Name.IsNull() {
		*name = re.Name.ValueString()
	} else {
		name = &emptyString
	}
	var networkIDs []string = nil
	re.NetworkIDs.ElementsAs(ctx, &networkIDs, false)
	out := merakigosdk.RequestOrganizationsCombineOrganizationNetworks{
		EnrollmentString: *enrollmentString,
		Name:             *name,
		NetworkIDs:       networkIDs,
	}
	return &out
}

// ToBody
func ResponseOrganizationsCombineOrganizationNetworksItemToBody(state OrganizationsNetworksCombine, response *merakigosdk.ResponseOrganizationsCombineOrganizationNetworks) OrganizationsNetworksCombine {
	itemState := ResponseOrganizationsCombineOrganizationNetworks{
		ResultingNetwork: func() *ResponseOrganizationsCombineOrganizationNetworksResultingNetwork {
			if response.ResultingNetwork != nil {
				return &ResponseOrganizationsCombineOrganizationNetworksResultingNetwork{
					EnrollmentString: types.StringValue(response.ResultingNetwork.EnrollmentString),
					ID:               types.StringValue(response.ResultingNetwork.ID),
					IsBoundToConfigTemplate: func() types.Bool {
						if response.ResultingNetwork.IsBoundToConfigTemplate != nil {
							return types.BoolValue(*response.ResultingNetwork.IsBoundToConfigTemplate)
						}
						return types.Bool{}
					}(),
					Name:           types.StringValue(response.ResultingNetwork.Name),
					Notes:          types.StringValue(response.ResultingNetwork.Notes),
					OrganizationID: types.StringValue(response.ResultingNetwork.OrganizationID),
					ProductTypes:   StringSliceToList(response.ResultingNetwork.ProductTypes),
					Tags:           StringSliceToList(response.ResultingNetwork.Tags),
					TimeZone:       types.StringValue(response.ResultingNetwork.TimeZone),
					URL:            types.StringValue(response.ResultingNetwork.URL),
				}
			}
			return nil
		}(),
	}
	state.Item = &itemState
	return state
}
