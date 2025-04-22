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
	_ resource.Resource              = &OrganizationsCloneResource{}
	_ resource.ResourceWithConfigure = &OrganizationsCloneResource{}
)

func NewOrganizationsCloneResource() resource.Resource {
	return &OrganizationsCloneResource{}
}

type OrganizationsCloneResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsCloneResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsCloneResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_clone"
}

// resourceAction
func (r *OrganizationsCloneResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"api": schema.SingleNestedAttribute{
						MarkdownDescription: `API related settings`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"enabled": schema.BoolAttribute{
								MarkdownDescription: `Enable API access`,
								Computed:            true,
							},
						},
					},
					"cloud": schema.SingleNestedAttribute{
						MarkdownDescription: `Data for this organization`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"region": schema.SingleNestedAttribute{
								MarkdownDescription: `Region info`,
								Computed:            true,
								Attributes: map[string]schema.Attribute{

									"host": schema.SingleNestedAttribute{
										MarkdownDescription: `Where organization data is hosted`,
										Computed:            true,
										Attributes: map[string]schema.Attribute{

											"name": schema.StringAttribute{
												MarkdownDescription: `Name of location`,
												Computed:            true,
											},
										},
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `Name of region`,
										Computed:            true,
									},
								},
							},
						},
					},
					"id": schema.StringAttribute{
						MarkdownDescription: `Organization ID`,
						Computed:            true,
					},
					"licensing": schema.SingleNestedAttribute{
						MarkdownDescription: `Licensing related settings`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"model": schema.StringAttribute{
								MarkdownDescription: `Organization licensing model. Can be 'co-term', 'per-device', or 'subscription'.
                                                Allowed values: [co-term,per-device,subscription]`,
								Computed: true,
							},
						},
					},
					"management": schema.SingleNestedAttribute{
						MarkdownDescription: `Information about the organization's management system`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"details": schema.SetNestedAttribute{
								MarkdownDescription: `Details related to organization management, possibly empty. Details may be named 'MSP ID', 'customer number', 'IP restriction mode for API', or 'IP restriction mode for dashboard', if the organization admin has configured any.`,
								Computed:            true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{

										"name": schema.StringAttribute{
											MarkdownDescription: `Name of management data`,
											Computed:            true,
										},
										"value": schema.StringAttribute{
											MarkdownDescription: `Value of management data`,
											Computed:            true,
										},
									},
								},
							},
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Organization name`,
						Computed:            true,
					},
					"url": schema.StringAttribute{
						MarkdownDescription: `Organization URL`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						MarkdownDescription: `The name of the new organization`,
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
func (r *OrganizationsCloneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsClone

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
	response, restyResp1, err := r.client.Organizations.CloneOrganization(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CloneOrganization",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CloneOrganization",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseOrganizationsCloneOrganizationItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsCloneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsCloneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsCloneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsClone struct {
	OrganizationID types.String                             `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsCloneOrganization  `tfsdk:"item"`
	Parameters     *RequestOrganizationsCloneOrganizationRs `tfsdk:"parameters"`
}

type ResponseOrganizationsCloneOrganization struct {
	API        *ResponseOrganizationsCloneOrganizationApi        `tfsdk:"api"`
	Cloud      *ResponseOrganizationsCloneOrganizationCloud      `tfsdk:"cloud"`
	ID         types.String                                      `tfsdk:"id"`
	Licensing  *ResponseOrganizationsCloneOrganizationLicensing  `tfsdk:"licensing"`
	Management *ResponseOrganizationsCloneOrganizationManagement `tfsdk:"management"`
	Name       types.String                                      `tfsdk:"name"`
	URL        types.String                                      `tfsdk:"url"`
}

type ResponseOrganizationsCloneOrganizationApi struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseOrganizationsCloneOrganizationCloud struct {
	Region *ResponseOrganizationsCloneOrganizationCloudRegion `tfsdk:"region"`
}

type ResponseOrganizationsCloneOrganizationCloudRegion struct {
	Host *ResponseOrganizationsCloneOrganizationCloudRegionHost `tfsdk:"host"`
	Name types.String                                           `tfsdk:"name"`
}

type ResponseOrganizationsCloneOrganizationCloudRegionHost struct {
	Name types.String `tfsdk:"name"`
}

type ResponseOrganizationsCloneOrganizationLicensing struct {
	Model types.String `tfsdk:"model"`
}

type ResponseOrganizationsCloneOrganizationManagement struct {
	Details *[]ResponseOrganizationsCloneOrganizationManagementDetails `tfsdk:"details"`
}

type ResponseOrganizationsCloneOrganizationManagementDetails struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

type RequestOrganizationsCloneOrganizationRs struct {
	Name types.String `tfsdk:"name"`
}

// FromBody
func (r *OrganizationsClone) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCloneOrganization {
	emptyString := ""
	re := *r.Parameters
	name := new(string)
	if !re.Name.IsUnknown() && !re.Name.IsNull() {
		*name = re.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestOrganizationsCloneOrganization{
		Name: *name,
	}
	return &out
}

// ToBody
func ResponseOrganizationsCloneOrganizationItemToBody(state OrganizationsClone, response *merakigosdk.ResponseOrganizationsCloneOrganization) OrganizationsClone {
	itemState := ResponseOrganizationsCloneOrganization{
		API: func() *ResponseOrganizationsCloneOrganizationApi {
			if response.API != nil {
				return &ResponseOrganizationsCloneOrganizationApi{
					Enabled: func() types.Bool {
						if response.API.Enabled != nil {
							return types.BoolValue(*response.API.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		Cloud: func() *ResponseOrganizationsCloneOrganizationCloud {
			if response.Cloud != nil {
				return &ResponseOrganizationsCloneOrganizationCloud{
					Region: func() *ResponseOrganizationsCloneOrganizationCloudRegion {
						if response.Cloud.Region != nil {
							return &ResponseOrganizationsCloneOrganizationCloudRegion{
								Host: func() *ResponseOrganizationsCloneOrganizationCloudRegionHost {
									if response.Cloud.Region.Host != nil {
										return &ResponseOrganizationsCloneOrganizationCloudRegionHost{
											Name: types.StringValue(response.Cloud.Region.Host.Name),
										}
									}
									return nil
								}(),
								Name: types.StringValue(response.Cloud.Region.Name),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		ID: types.StringValue(response.ID),
		Licensing: func() *ResponseOrganizationsCloneOrganizationLicensing {
			if response.Licensing != nil {
				return &ResponseOrganizationsCloneOrganizationLicensing{
					Model: types.StringValue(response.Licensing.Model),
				}
			}
			return nil
		}(),
		Management: func() *ResponseOrganizationsCloneOrganizationManagement {
			if response.Management != nil {
				return &ResponseOrganizationsCloneOrganizationManagement{
					Details: func() *[]ResponseOrganizationsCloneOrganizationManagementDetails {
						if response.Management.Details != nil {
							result := make([]ResponseOrganizationsCloneOrganizationManagementDetails, len(*response.Management.Details))
							for i, details := range *response.Management.Details {
								result[i] = ResponseOrganizationsCloneOrganizationManagementDetails{
									Name:  types.StringValue(details.Name),
									Value: types.StringValue(details.Value),
								}
							}
							return &result
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		Name: types.StringValue(response.Name),
		URL:  types.StringValue(response.URL),
	}
	state.Item = &itemState
	return state
}
