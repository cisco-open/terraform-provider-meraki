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
	_ resource.Resource              = &OrganizationsClaimResource{}
	_ resource.ResourceWithConfigure = &OrganizationsClaimResource{}
)

func NewOrganizationsClaimResource() resource.Resource {
	return &OrganizationsClaimResource{}
}

type OrganizationsClaimResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsClaimResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsClaimResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_claim"
}

// resourceAction
func (r *OrganizationsClaimResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"licenses": schema.SetNestedAttribute{
						MarkdownDescription: `The licenses claimed`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"key": schema.StringAttribute{
									MarkdownDescription: `The key of the license`,
									Computed:            true,
								},
								"mode": schema.StringAttribute{
									MarkdownDescription: `The mode of the license`,
									Computed:            true,
								},
							},
						},
					},
					"orders": schema.ListAttribute{
						MarkdownDescription: `The numbers of the orders claimed`,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"serials": schema.ListAttribute{
						MarkdownDescription: `The serials of the devices claimed`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"licenses": schema.SetNestedAttribute{
						MarkdownDescription: `The licenses that should be claimed`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"key": schema.StringAttribute{
									MarkdownDescription: `The key of the license`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"mode": schema.StringAttribute{
									MarkdownDescription: `Either 'renew' or 'addDevices'. 'addDevices' will increase the license limit, while 'renew' will extend the amount of time until expiration. Defaults to 'addDevices'. All licenses must be claimed with the same mode, and at most one renewal can be claimed at a time. This parameter is legacy and does not apply to organizations with per-device licensing enabled.
                                              Allowed values: [addDevices,renew]`,
									Optional: true,
									Computed: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
							},
						},
					},
					"orders": schema.ListAttribute{
						MarkdownDescription: `The numbers of the orders that should be claimed`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"serials": schema.ListAttribute{
						MarkdownDescription: `The serials of the devices that should be claimed`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *OrganizationsClaimResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsClaim

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
	response, restyResp1, err := r.client.Organizations.ClaimIntoOrganization(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ClaimIntoOrganization",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ClaimIntoOrganization",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseOrganizationsClaimIntoOrganizationItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsClaimResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsClaimResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsClaimResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsClaim struct {
	OrganizationID types.String                                 `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsClaimIntoOrganization  `tfsdk:"item"`
	Parameters     *RequestOrganizationsClaimIntoOrganizationRs `tfsdk:"parameters"`
}

type ResponseOrganizationsClaimIntoOrganization struct {
	Licenses *[]ResponseOrganizationsClaimIntoOrganizationLicenses `tfsdk:"licenses"`
	Orders   types.List                                            `tfsdk:"orders"`
	Serials  types.List                                            `tfsdk:"serials"`
}

type ResponseOrganizationsClaimIntoOrganizationLicenses struct {
	Key  types.String `tfsdk:"key"`
	Mode types.String `tfsdk:"mode"`
}

type RequestOrganizationsClaimIntoOrganizationRs struct {
	Licenses *[]RequestOrganizationsClaimIntoOrganizationLicensesRs `tfsdk:"licenses"`
	Orders   types.Set                                              `tfsdk:"orders"`
	Serials  types.Set                                              `tfsdk:"serials"`
}

type RequestOrganizationsClaimIntoOrganizationLicensesRs struct {
	Key  types.String `tfsdk:"key"`
	Mode types.String `tfsdk:"mode"`
}

// FromBody
func (r *OrganizationsClaim) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsClaimIntoOrganization {
	re := *r.Parameters
	var requestOrganizationsClaimIntoOrganizationLicenses []merakigosdk.RequestOrganizationsClaimIntoOrganizationLicenses

	if re.Licenses != nil {
		for _, rItem1 := range *re.Licenses {
			key := rItem1.Key.ValueString()
			mode := rItem1.Mode.ValueString()
			requestOrganizationsClaimIntoOrganizationLicenses = append(requestOrganizationsClaimIntoOrganizationLicenses, merakigosdk.RequestOrganizationsClaimIntoOrganizationLicenses{
				Key:  key,
				Mode: mode,
			})
			//[debug] Is Array: True
		}
	}
	var orders []string = nil
	re.Orders.ElementsAs(ctx, &orders, false)
	var serials []string = nil
	re.Serials.ElementsAs(ctx, &serials, false)
	out := merakigosdk.RequestOrganizationsClaimIntoOrganization{
		Licenses: func() *[]merakigosdk.RequestOrganizationsClaimIntoOrganizationLicenses {
			if len(requestOrganizationsClaimIntoOrganizationLicenses) > 0 {
				return &requestOrganizationsClaimIntoOrganizationLicenses
			}
			return nil
		}(),
		Orders:  orders,
		Serials: serials,
	}
	return &out
}

// ToBody
func ResponseOrganizationsClaimIntoOrganizationItemToBody(state OrganizationsClaim, response *merakigosdk.ResponseOrganizationsClaimIntoOrganization) OrganizationsClaim {
	itemState := ResponseOrganizationsClaimIntoOrganization{
		Licenses: func() *[]ResponseOrganizationsClaimIntoOrganizationLicenses {
			if response.Licenses != nil {
				result := make([]ResponseOrganizationsClaimIntoOrganizationLicenses, len(*response.Licenses))
				for i, licenses := range *response.Licenses {
					result[i] = ResponseOrganizationsClaimIntoOrganizationLicenses{
						Key:  types.StringValue(licenses.Key),
						Mode: types.StringValue(licenses.Mode),
					}
				}
				return &result
			}
			return nil
		}(),
		Orders:  StringSliceToList(response.Orders),
		Serials: StringSliceToList(response.Serials),
	}
	state.Item = &itemState
	return state
}
