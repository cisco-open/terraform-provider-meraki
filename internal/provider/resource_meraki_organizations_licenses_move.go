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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsLicensesMoveResource{}
	_ resource.ResourceWithConfigure = &OrganizationsLicensesMoveResource{}
)

func NewOrganizationsLicensesMoveResource() resource.Resource {
	return &OrganizationsLicensesMoveResource{}
}

type OrganizationsLicensesMoveResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsLicensesMoveResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsLicensesMoveResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_licenses_move"
}

// resourceAction
func (r *OrganizationsLicensesMoveResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"dest_organization_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the organization to move the licenses to`,
						Computed:            true,
					},
					"license_ids": schema.SetAttribute{
						MarkdownDescription: `A list of IDs of licenses to move to the new organization`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"dest_organization_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the organization to move the licenses to`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"license_ids": schema.SetAttribute{
						MarkdownDescription: `A list of IDs of licenses to move to the new organization`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *OrganizationsLicensesMoveResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsLicensesMove

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
	// organization_id
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Organizations.MoveOrganizationLicenses(vvOrganizationID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing MoveOrganizationLicenses",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing MoveOrganizationLicenses",
			err.Error(),
		)
		return
	}
	//Item
	data2 := ResponseOrganizationsMoveOrganizationLicensesItemToBody(data, response)

	diags := resp.State.Set(ctx, &data2)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsLicensesMoveResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsLicensesMoveResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsLicensesMoveResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsLicensesMove struct {
	OrganizationID types.String                                    `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsMoveOrganizationLicenses  `tfsdk:"item"`
	Parameters     *RequestOrganizationsMoveOrganizationLicensesRs `tfsdk:"parameters"`
}

type ResponseOrganizationsMoveOrganizationLicenses struct {
	DestOrganizationID types.String `tfsdk:"dest_organization_id"`
	LicenseIDs         types.Set    `tfsdk:"license_ids"`
}

type RequestOrganizationsMoveOrganizationLicensesRs struct {
	DestOrganizationID types.String `tfsdk:"dest_organization_id"`
	LicenseIDs         types.Set    `tfsdk:"license_ids"`
}

// FromBody
func (r *OrganizationsLicensesMove) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsMoveOrganizationLicenses {
	emptyString := ""
	re := *r.Parameters
	destOrganizationID := new(string)
	if !re.DestOrganizationID.IsUnknown() && !re.DestOrganizationID.IsNull() {
		*destOrganizationID = re.DestOrganizationID.ValueString()
	} else {
		destOrganizationID = &emptyString
	}
	var licenseIDs []string = nil
	re.LicenseIDs.ElementsAs(ctx, &licenseIDs, false)
	out := merakigosdk.RequestOrganizationsMoveOrganizationLicenses{
		DestOrganizationID: *destOrganizationID,
		LicenseIDs:         licenseIDs,
	}
	return &out
}

// ToBody
func ResponseOrganizationsMoveOrganizationLicensesItemToBody(state OrganizationsLicensesMove, response *merakigosdk.ResponseOrganizationsMoveOrganizationLicenses) OrganizationsLicensesMove {
	itemState := ResponseOrganizationsMoveOrganizationLicenses{
		DestOrganizationID: types.StringValue(response.DestOrganizationID),
		LicenseIDs:         StringSliceToSet(response.LicenseIDs),
	}
	state.Item = &itemState
	return state
}
