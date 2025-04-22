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
	_ resource.Resource              = &OrganizationsIntegrationsXdrNetworksDisableResource{}
	_ resource.ResourceWithConfigure = &OrganizationsIntegrationsXdrNetworksDisableResource{}
)

func NewOrganizationsIntegrationsXdrNetworksDisableResource() resource.Resource {
	return &OrganizationsIntegrationsXdrNetworksDisableResource{}
}

type OrganizationsIntegrationsXdrNetworksDisableResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsIntegrationsXdrNetworksDisableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsIntegrationsXdrNetworksDisableResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_integrations_xdr_networks_disable"
}

// resourceAction
func (r *OrganizationsIntegrationsXdrNetworksDisableResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"networks": schema.SetNestedAttribute{
						MarkdownDescription: `List of networks that have XDR disabled`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Represents whether XDR is enabled for the network`,
									Computed:            true,
								},
								"is_eligible": schema.BoolAttribute{
									MarkdownDescription: `Represents whether the network is eligible for XDR`,
									Computed:            true,
								},
								"name": schema.StringAttribute{
									MarkdownDescription: `The name of the network`,
									Computed:            true,
								},
								"network_id": schema.StringAttribute{
									MarkdownDescription: `Network ID`,
									Computed:            true,
								},
								"product_types": schema.ListAttribute{
									MarkdownDescription: `List of products that have XDR disabled`,
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"networks": schema.SetNestedAttribute{
						MarkdownDescription: `List containing the network ID and the product type to disable XDR on`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"network_id": schema.StringAttribute{
									MarkdownDescription: `Network ID`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"product_types": schema.ListAttribute{
									MarkdownDescription: `List of products for which to disable XDR`,
									Optional:            true,
									Computed:            true,
									ElementType:         types.StringType,
								},
							},
						},
					},
				},
			},
		},
	}
}
func (r *OrganizationsIntegrationsXdrNetworksDisableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsIntegrationsXdrNetworksDisable

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
	response, restyResp1, err := r.client.Organizations.DisableOrganizationIntegrationsXdrNetworks(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing DisableOrganizationIntegrationsXdrNetworks",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing DisableOrganizationIntegrationsXdrNetworks",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseOrganizationsDisableOrganizationIntegrationsXdrNetworksItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsIntegrationsXdrNetworksDisableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsIntegrationsXdrNetworksDisableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsIntegrationsXdrNetworksDisableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsIntegrationsXdrNetworksDisable struct {
	OrganizationID types.String                                                      `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsDisableOrganizationIntegrationsXdrNetworks  `tfsdk:"item"`
	Parameters     *RequestOrganizationsDisableOrganizationIntegrationsXdrNetworksRs `tfsdk:"parameters"`
}

type ResponseOrganizationsDisableOrganizationIntegrationsXdrNetworks struct {
	Networks *[]ResponseOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworks `tfsdk:"networks"`
}

type ResponseOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworks struct {
	Enabled      types.Bool   `tfsdk:"enabled"`
	IsEligible   types.Bool   `tfsdk:"is_eligible"`
	Name         types.String `tfsdk:"name"`
	NetworkID    types.String `tfsdk:"network_id"`
	ProductTypes types.List   `tfsdk:"product_types"`
}

type RequestOrganizationsDisableOrganizationIntegrationsXdrNetworksRs struct {
	Networks *[]RequestOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworksRs `tfsdk:"networks"`
}

type RequestOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworksRs struct {
	NetworkID    types.String `tfsdk:"network_id"`
	ProductTypes types.Set    `tfsdk:"product_types"`
}

// FromBody
func (r *OrganizationsIntegrationsXdrNetworksDisable) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsDisableOrganizationIntegrationsXdrNetworks {
	re := *r.Parameters
	var requestOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworks []merakigosdk.RequestOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworks

	if re.Networks != nil {
		for _, rItem1 := range *re.Networks {
			networkID := rItem1.NetworkID.ValueString()

			var productTypes []string = nil
			rItem1.ProductTypes.ElementsAs(ctx, &productTypes, false)
			requestOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworks = append(requestOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworks, merakigosdk.RequestOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworks{
				NetworkID:    networkID,
				ProductTypes: productTypes,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestOrganizationsDisableOrganizationIntegrationsXdrNetworks{
		Networks: &requestOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworks,
	}
	return &out
}

// ToBody
func ResponseOrganizationsDisableOrganizationIntegrationsXdrNetworksItemToBody(state OrganizationsIntegrationsXdrNetworksDisable, response *merakigosdk.ResponseOrganizationsDisableOrganizationIntegrationsXdrNetworks) OrganizationsIntegrationsXdrNetworksDisable {
	itemState := ResponseOrganizationsDisableOrganizationIntegrationsXdrNetworks{
		Networks: func() *[]ResponseOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworks {
			if response.Networks != nil {
				result := make([]ResponseOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworks, len(*response.Networks))
				for i, networks := range *response.Networks {
					result[i] = ResponseOrganizationsDisableOrganizationIntegrationsXdrNetworksNetworks{
						Enabled: func() types.Bool {
							if networks.Enabled != nil {
								return types.BoolValue(*networks.Enabled)
							}
							return types.Bool{}
						}(),
						IsEligible: func() types.Bool {
							if networks.IsEligible != nil {
								return types.BoolValue(*networks.IsEligible)
							}
							return types.Bool{}
						}(),
						Name:         types.StringValue(networks.Name),
						NetworkID:    types.StringValue(networks.NetworkID),
						ProductTypes: StringSliceToList(networks.ProductTypes),
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
