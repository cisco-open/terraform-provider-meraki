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

// RESOURCE NORMAL
import (
	"context"
	"strconv"

	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsResource{}
)

func NewOrganizationsResource() resource.Resource {
	return &OrganizationsResource{}
}

type OrganizationsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations"
}

func (r *OrganizationsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"api": schema.SingleNestedAttribute{
				MarkdownDescription: `API related settings`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Enable API access`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
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
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"details": schema.SetNestedAttribute{
						MarkdownDescription: `Details related to organization management, possibly empty. Details may be named 'MSP ID', 'customer number', 'IP restriction mode for API', or 'IP restriction mode for dashboard', if the organization admin has configured any.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"name": schema.StringAttribute{
									MarkdownDescription: `Name of management data`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"value": schema.StringAttribute{
									MarkdownDescription: `Value of management data`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Organization name`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"url": schema.StringAttribute{
				MarkdownDescription: `Organization URL`,
				Computed:            true,
			},
		},
	}
}

//path params to set ['organizationId']

func (r *OrganizationsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsRs

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
	//Has Item and has items and post

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizations(&merakigosdk.GetOrganizationsQueryParams{
		PerPage: -1,
	})
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizations",
					restyResp1.String(),
				)
				return
			}
		}
	}

	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvOrganizationID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter OrganizationID",
					"Fail Parsing OrganizationID",
				)
				return
			}
			r.client.Organizations.UpdateOrganization(vvOrganizationID, data.toSdkApiRequestUpdate(ctx))

			responseVerifyItem2, _, _ := r.client.Organizations.GetOrganization(vvOrganizationID)
			if responseVerifyItem2 != nil {
				data = ResponseOrganizationsGetOrganizationItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Organizations.CreateOrganization(dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganization",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganization",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Organizations.GetOrganizations(&merakigosdk.GetOrganizationsQueryParams{
		PerPage: -1})

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizations",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizations",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvOrganizationID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter OrganizationID",
				"Fail Parsing OrganizationID",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Organizations.GetOrganization(vvOrganizationID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseOrganizationsGetOrganizationItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganization",
					restyRespGet.String(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganization",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}

}

func (r *OrganizationsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsRs

	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
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
	// Has Item2

	vvOrganizationID := data.OrganizationID.ValueString()
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganization(vvOrganizationID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganization",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganization",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), req.ID)...)
}

func (r *OrganizationsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganization(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganization",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganization",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvOrganizationID := state.OrganizationID.ValueString()
	_, err := r.client.Organizations.DeleteOrganization(vvOrganizationID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganization", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsRs struct {
	OrganizationID types.String                                      `tfsdk:"organization_id"`
	API            *ResponseOrganizationsGetOrganizationApiRs        `tfsdk:"api"`
	Cloud          *ResponseOrganizationsGetOrganizationCloudRs      `tfsdk:"cloud"`
	ID             types.String                                      `tfsdk:"id"`
	Licensing      *ResponseOrganizationsGetOrganizationLicensingRs  `tfsdk:"licensing"`
	Management     *ResponseOrganizationsGetOrganizationManagementRs `tfsdk:"management"`
	Name           types.String                                      `tfsdk:"name"`
	URL            types.String                                      `tfsdk:"url"`
}

type ResponseOrganizationsGetOrganizationApiRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseOrganizationsGetOrganizationCloudRs struct {
	Region *ResponseOrganizationsGetOrganizationCloudRegionRs `tfsdk:"region"`
}

type ResponseOrganizationsGetOrganizationCloudRegionRs struct {
	Host *ResponseOrganizationsGetOrganizationCloudRegionHostRs `tfsdk:"host"`
	Name types.String                                           `tfsdk:"name"`
}

type ResponseOrganizationsGetOrganizationCloudRegionHostRs struct {
	Name types.String `tfsdk:"name"`
}

type ResponseOrganizationsGetOrganizationLicensingRs struct {
	Model types.String `tfsdk:"model"`
}

type ResponseOrganizationsGetOrganizationManagementRs struct {
	Details *[]ResponseOrganizationsGetOrganizationManagementDetailsRs `tfsdk:"details"`
}

type ResponseOrganizationsGetOrganizationManagementDetailsRs struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// FromBody
func (r *OrganizationsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganization {
	emptyString := ""
	var requestOrganizationsCreateOrganizationManagement *merakigosdk.RequestOrganizationsCreateOrganizationManagement

	if r.Management != nil {

		log.Printf("[DEBUG] #TODO []RequestOrganizationsCreateOrganizationManagementDetails")
		var requestOrganizationsCreateOrganizationManagementDetails []merakigosdk.RequestOrganizationsCreateOrganizationManagementDetails

		if r.Management.Details != nil {
			for _, rItem1 := range *r.Management.Details {
				name := rItem1.Name.ValueString()
				value := rItem1.Value.ValueString()
				requestOrganizationsCreateOrganizationManagementDetails = append(requestOrganizationsCreateOrganizationManagementDetails, merakigosdk.RequestOrganizationsCreateOrganizationManagementDetails{
					Name:  name,
					Value: value,
				})
				//[debug] Is Array: True
			}
		}
		requestOrganizationsCreateOrganizationManagement = &merakigosdk.RequestOrganizationsCreateOrganizationManagement{
			Details: func() *[]merakigosdk.RequestOrganizationsCreateOrganizationManagementDetails {
				if len(requestOrganizationsCreateOrganizationManagementDetails) > 0 {
					return &requestOrganizationsCreateOrganizationManagementDetails
				}
				return nil
			}(),
		}
		//[debug] Is Array: False
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestOrganizationsCreateOrganization{
		Management: requestOrganizationsCreateOrganizationManagement,
		Name:       *name,
	}
	return &out
}
func (r *OrganizationsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganization {
	emptyString := ""
	var requestOrganizationsUpdateOrganizationAPI *merakigosdk.RequestOrganizationsUpdateOrganizationAPI

	if r.API != nil {
		enabled := func() *bool {
			if !r.API.Enabled.IsUnknown() && !r.API.Enabled.IsNull() {
				return r.API.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestOrganizationsUpdateOrganizationAPI = &merakigosdk.RequestOrganizationsUpdateOrganizationAPI{
			Enabled: enabled,
		}
		//[debug] Is Array: False
	}
	var requestOrganizationsUpdateOrganizationManagement *merakigosdk.RequestOrganizationsUpdateOrganizationManagement

	if r.Management != nil {

		log.Printf("[DEBUG] #TODO []RequestOrganizationsUpdateOrganizationManagementDetails")
		var requestOrganizationsUpdateOrganizationManagementDetails []merakigosdk.RequestOrganizationsUpdateOrganizationManagementDetails

		if r.Management.Details != nil {
			for _, rItem1 := range *r.Management.Details {
				name := rItem1.Name.ValueString()
				value := rItem1.Value.ValueString()
				requestOrganizationsUpdateOrganizationManagementDetails = append(requestOrganizationsUpdateOrganizationManagementDetails, merakigosdk.RequestOrganizationsUpdateOrganizationManagementDetails{
					Name:  name,
					Value: value,
				})
				//[debug] Is Array: True
			}
		}
		requestOrganizationsUpdateOrganizationManagement = &merakigosdk.RequestOrganizationsUpdateOrganizationManagement{
			Details: func() *[]merakigosdk.RequestOrganizationsUpdateOrganizationManagementDetails {
				if len(requestOrganizationsUpdateOrganizationManagementDetails) > 0 {
					return &requestOrganizationsUpdateOrganizationManagementDetails
				}
				return nil
			}(),
		}
		//[debug] Is Array: False
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganization{
		API:        requestOrganizationsUpdateOrganizationAPI,
		Management: requestOrganizationsUpdateOrganizationManagement,
		Name:       *name,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationItemToBodyRs(state OrganizationsRs, response *merakigosdk.ResponseOrganizationsGetOrganization, is_read bool) OrganizationsRs {
	itemState := OrganizationsRs{
		API: func() *ResponseOrganizationsGetOrganizationApiRs {
			if response.API != nil {
				return &ResponseOrganizationsGetOrganizationApiRs{
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
		Cloud: func() *ResponseOrganizationsGetOrganizationCloudRs {
			if response.Cloud != nil {
				return &ResponseOrganizationsGetOrganizationCloudRs{
					Region: func() *ResponseOrganizationsGetOrganizationCloudRegionRs {
						if response.Cloud.Region != nil {
							return &ResponseOrganizationsGetOrganizationCloudRegionRs{
								Host: func() *ResponseOrganizationsGetOrganizationCloudRegionHostRs {
									if response.Cloud.Region.Host != nil {
										return &ResponseOrganizationsGetOrganizationCloudRegionHostRs{
											Name: func() types.String {
												if response.Cloud.Region.Host.Name != "" {
													return types.StringValue(response.Cloud.Region.Host.Name)
												}
												return types.String{}
											}(),
										}
									}
									return nil
								}(),
								Name: func() types.String {
									if response.Cloud.Region.Name != "" {
										return types.StringValue(response.Cloud.Region.Name)
									}
									return types.String{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		ID: func() types.String {
			if response.ID != "" {
				return types.StringValue(response.ID)
			}
			return types.String{}
		}(),
		Licensing: func() *ResponseOrganizationsGetOrganizationLicensingRs {
			if response.Licensing != nil {
				return &ResponseOrganizationsGetOrganizationLicensingRs{
					Model: func() types.String {
						if response.Licensing.Model != "" {
							return types.StringValue(response.Licensing.Model)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		Management: func() *ResponseOrganizationsGetOrganizationManagementRs {
			if response.Management != nil {
				return &ResponseOrganizationsGetOrganizationManagementRs{
					Details: func() *[]ResponseOrganizationsGetOrganizationManagementDetailsRs {
						if response.Management.Details != nil {
							result := make([]ResponseOrganizationsGetOrganizationManagementDetailsRs, len(*response.Management.Details))
							for i, details := range *response.Management.Details {
								result[i] = ResponseOrganizationsGetOrganizationManagementDetailsRs{
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
				}
			}
			return nil
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
		URL: func() types.String {
			if response.URL != "" {
				return types.StringValue(response.URL)
			}
			return types.String{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsRs)
}
