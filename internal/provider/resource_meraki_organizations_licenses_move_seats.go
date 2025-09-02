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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsLicensesMoveSeatsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsLicensesMoveSeatsResource{}
)

func NewOrganizationsLicensesMoveSeatsResource() resource.Resource {
	return &OrganizationsLicensesMoveSeatsResource{}
}

type OrganizationsLicensesMoveSeatsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsLicensesMoveSeatsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsLicensesMoveSeatsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_licenses_move_seats"
}

// resourceAction
func (r *OrganizationsLicensesMoveSeatsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
						MarkdownDescription: `The ID of the organization to move the SM seats to`,
						Computed:            true,
					},
					"license_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the SM license to move the seats from`,
						Computed:            true,
					},
					"seat_count": schema.Int64Attribute{
						MarkdownDescription: `The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"dest_organization_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the organization to move the SM seats to`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"license_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the SM license to move the seats from`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"seat_count": schema.Int64Attribute{
						MarkdownDescription: `The number of seats to move to the new organization. Must be less than or equal to the total number of seats of the license`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *OrganizationsLicensesMoveSeatsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsLicensesMoveSeats

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
	response, restyResp1, err := r.client.Organizations.MoveOrganizationLicensesSeats(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing MoveOrganizationLicensesSeats",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing MoveOrganizationLicensesSeats",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseOrganizationsMoveOrganizationLicensesSeatsItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsLicensesMoveSeatsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsLicensesMoveSeatsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsLicensesMoveSeatsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsLicensesMoveSeats struct {
	OrganizationID types.String                                         `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsMoveOrganizationLicensesSeats  `tfsdk:"item"`
	Parameters     *RequestOrganizationsMoveOrganizationLicensesSeatsRs `tfsdk:"parameters"`
}

type ResponseOrganizationsMoveOrganizationLicensesSeats struct {
	DestOrganizationID types.String `tfsdk:"dest_organization_id"`
	LicenseID          types.String `tfsdk:"license_id"`
	SeatCount          types.Int64  `tfsdk:"seat_count"`
}

type RequestOrganizationsMoveOrganizationLicensesSeatsRs struct {
	DestOrganizationID types.String `tfsdk:"dest_organization_id"`
	LicenseID          types.String `tfsdk:"license_id"`
	SeatCount          types.Int64  `tfsdk:"seat_count"`
}

// FromBody
func (r *OrganizationsLicensesMoveSeats) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsMoveOrganizationLicensesSeats {
	emptyString := ""
	re := *r.Parameters
	destOrganizationID := new(string)
	if !re.DestOrganizationID.IsUnknown() && !re.DestOrganizationID.IsNull() {
		*destOrganizationID = re.DestOrganizationID.ValueString()
	} else {
		destOrganizationID = &emptyString
	}
	licenseID := new(string)
	if !re.LicenseID.IsUnknown() && !re.LicenseID.IsNull() {
		*licenseID = re.LicenseID.ValueString()
	} else {
		licenseID = &emptyString
	}
	seatCount := new(int64)
	if !re.SeatCount.IsUnknown() && !re.SeatCount.IsNull() {
		*seatCount = re.SeatCount.ValueInt64()
	} else {
		seatCount = nil
	}
	out := merakigosdk.RequestOrganizationsMoveOrganizationLicensesSeats{
		DestOrganizationID: *destOrganizationID,
		LicenseID:          *licenseID,
		SeatCount:          int64ToIntPointer(seatCount),
	}
	return &out
}

// ToBody
func ResponseOrganizationsMoveOrganizationLicensesSeatsItemToBody(state OrganizationsLicensesMoveSeats, response *merakigosdk.ResponseOrganizationsMoveOrganizationLicensesSeats) OrganizationsLicensesMoveSeats {
	itemState := ResponseOrganizationsMoveOrganizationLicensesSeats{
		DestOrganizationID: func() types.String {
			if response.DestOrganizationID != "" {
				return types.StringValue(response.DestOrganizationID)
			}
			return types.String{}
		}(),
		LicenseID: func() types.String {
			if response.LicenseID != "" {
				return types.StringValue(response.LicenseID)
			}
			return types.String{}
		}(),
		SeatCount: func() types.Int64 {
			if response.SeatCount != nil {
				return types.Int64Value(int64(*response.SeatCount))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
