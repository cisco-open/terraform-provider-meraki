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
	_ resource.Resource              = &OrganizationsInventoryOnboardingCloudMonitoringImportsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsInventoryOnboardingCloudMonitoringImportsResource{}
)

func NewOrganizationsInventoryOnboardingCloudMonitoringImportsResource() resource.Resource {
	return &OrganizationsInventoryOnboardingCloudMonitoringImportsResource{}
}

type OrganizationsInventoryOnboardingCloudMonitoringImportsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsInventoryOnboardingCloudMonitoringImportsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsInventoryOnboardingCloudMonitoringImportsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_inventory_onboarding_cloud_monitoring_imports"
}

// resourceAction
func (r *OrganizationsInventoryOnboardingCloudMonitoringImportsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"devices": schema.SetNestedAttribute{
						MarkdownDescription: `A set of device imports to commit`,
						Optional:            true,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"device_id": schema.StringAttribute{
									MarkdownDescription: `Import ID from the Import operation`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"network_id": schema.StringAttribute{
									MarkdownDescription: `Network Id`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
								"udi": schema.StringAttribute{
									MarkdownDescription: `Device UDI certificate`,
									Optional:            true,
									Computed:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplace(),
									},
								},
							},
						},
					},
					"items": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"items": schema.ListNestedAttribute{
									MarkdownDescription: `Array of ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"import_id": schema.StringAttribute{
												MarkdownDescription: `Unique id associated with the import of the device`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"message": schema.StringAttribute{
												MarkdownDescription: `Response method`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"status": schema.StringAttribute{
												MarkdownDescription: `Cloud monitor import status`,
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
				},
			},
		},
	}
}
func (r *OrganizationsInventoryOnboardingCloudMonitoringImportsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsInventoryOnboardingCloudMonitoringImportsRs

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
	response, restyResp1, err := r.client.Organizations.CreateOrganizationInventoryOnboardingCloudMonitoringImport(vvOrganizationID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationInventoryOnboardingCloudMonitoringImport",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationInventoryOnboardingCloudMonitoringImport",
			err.Error(),
		)
		return
	}
	//Item
	// //entro aqui 2
	// data2 := ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsInventoryOnboardingCloudMonitoringImportsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsInventoryOnboardingCloudMonitoringImportsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsInventoryOnboardingCloudMonitoringImportsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsInventoryOnboardingCloudMonitoringImportsRs struct {
	OrganizationID types.String                                                                           `tfsdk:"organization_id"`
	Items          *[]ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport `tfsdk:"items"`
	Parameters     *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportRs      `tfsdk:"items"`
}

type ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport struct {
	ImportID types.String `tfsdk:"import_id"`
	Message  types.String `tfsdk:"message"`
	Status   types.String `tfsdk:"status"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportRs struct {
	Devices *[]RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportDevicesRs `tfsdk:"devices"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportDevicesRs struct {
	DeviceID  types.String `tfsdk:"device_id"`
	NetworkID types.String `tfsdk:"network_id"`
	Udi       types.String `tfsdk:"udi"`
}

// FromBody
func (r *OrganizationsInventoryOnboardingCloudMonitoringImportsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport {
	re := *r.Parameters
	var requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportDevices []merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportDevices
	if re.Devices != nil {
		for _, rItem1 := range *re.Devices {
			deviceID := rItem1.DeviceID.ValueString()
			networkID := rItem1.NetworkID.ValueString()
			udi := rItem1.Udi.ValueString()
			requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportDevices = append(requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportDevices, merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportDevices{
				DeviceID:  deviceID,
				NetworkID: networkID,
				Udi:       udi,
			})
		}
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport{
		Devices: func() *[]merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportDevices {
			if len(requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportDevices) > 0 {
				return &requestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportDevices
			}
			return nil
		}(),
	}
	return &out
}

// ToBody
func ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImportItemsToBody(state OrganizationsInventoryOnboardingCloudMonitoringImportsRs, response *merakigosdk.ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport) OrganizationsInventoryOnboardingCloudMonitoringImportsRs {
	var items []ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport
	for _, item := range *response {
		itemState := ResponseItemOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringImport{
			ImportID: types.StringValue(item.ImportID),
			Message:  types.StringValue(item.Message),
			Status:   types.StringValue(item.Status),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
