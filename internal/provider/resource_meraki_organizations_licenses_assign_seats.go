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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsLicensesAssignSeatsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsLicensesAssignSeatsResource{}
)

func NewOrganizationsLicensesAssignSeatsResource() resource.Resource {
	return &OrganizationsLicensesAssignSeatsResource{}
}

type OrganizationsLicensesAssignSeatsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsLicensesAssignSeatsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsLicensesAssignSeatsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_licenses_assign_seats"
}

// resourceAction
func (r *OrganizationsLicensesAssignSeatsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"resulting_licenses": schema.SetNestedAttribute{
						MarkdownDescription: `Resulting licenses from the move`,
						Computed:            true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"activation_date": schema.StringAttribute{
									MarkdownDescription: `The date the license started burning`,
									Computed:            true,
								},
								"claim_date": schema.StringAttribute{
									MarkdownDescription: `The date the license was claimed into the organization`,
									Computed:            true,
								},
								"device_serial": schema.StringAttribute{
									MarkdownDescription: `Serial number of the device the license is assigned to`,
									Computed:            true,
								},
								"duration_in_days": schema.Int64Attribute{
									MarkdownDescription: `The duration of the individual license`,
									Computed:            true,
								},
								"expiration_date": schema.StringAttribute{
									MarkdownDescription: `The date the license will expire`,
									Computed:            true,
								},
								"head_license_id": schema.StringAttribute{
									MarkdownDescription: `The id of the head license this license is queued behind. If there is no head license, it returns nil.`,
									Computed:            true,
								},
								"id": schema.StringAttribute{
									MarkdownDescription: `License ID`,
									Computed:            true,
								},
								"license_key": schema.StringAttribute{
									MarkdownDescription: `License key`,
									Computed:            true,
								},
								"license_type": schema.StringAttribute{
									MarkdownDescription: `License type`,
									Computed:            true,
								},
								"network_id": schema.StringAttribute{
									MarkdownDescription: `ID of the network the license is assigned to`,
									Computed:            true,
								},
								"order_number": schema.StringAttribute{
									MarkdownDescription: `Order number`,
									Computed:            true,
								},
								"permanently_queued_licenses": schema.SetNestedAttribute{
									MarkdownDescription: `DEPRECATED List of permanently queued licenses attached to the license. Instead, use /organizations/{organizationId}/licenses?deviceSerial= to retrieved queued licenses for a given device.`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"duration_in_days": schema.Int64Attribute{
												MarkdownDescription: `The duration of the individual license`,
												Computed:            true,
											},
											"id": schema.StringAttribute{
												MarkdownDescription: `Permanently queued license ID`,
												Computed:            true,
											},
											"license_key": schema.StringAttribute{
												MarkdownDescription: `License key`,
												Computed:            true,
											},
											"license_type": schema.StringAttribute{
												MarkdownDescription: `License type`,
												Computed:            true,
											},
											"order_number": schema.StringAttribute{
												MarkdownDescription: `Order number`,
												Computed:            true,
											},
										},
									},
								},
								"seat_count": schema.Int64Attribute{
									MarkdownDescription: `The number of seats of the license. Only applicable to SM licenses.`,
									Computed:            true,
								},
								"state": schema.StringAttribute{
									MarkdownDescription: `The state of the license. All queued licenses have a status of **recentlyQueued**.
                                                Allowed values: [active,expired,expiring,recentlyQueued,unused,unusedActive]`,
									Computed: true,
								},
								"total_duration_in_days": schema.Int64Attribute{
									MarkdownDescription: `The duration of the license plus all permanently queued licenses associated with it`,
									Computed:            true,
								},
							},
						},
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"license_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the SM license to assign seats from`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"network_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the SM network to assign the seats to`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"seat_count": schema.Int64Attribute{
						MarkdownDescription: `The number of seats to assign to the SM network. Must be less than or equal to the total number of seats of the license`,
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
func (r *OrganizationsLicensesAssignSeatsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsLicensesAssignSeats

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
	response, restyResp1, err := r.client.Organizations.AssignOrganizationLicensesSeats(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing AssignOrganizationLicensesSeats",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing AssignOrganizationLicensesSeats",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseOrganizationsAssignOrganizationLicensesSeatsItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsLicensesAssignSeatsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsLicensesAssignSeatsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsLicensesAssignSeatsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsLicensesAssignSeats struct {
	OrganizationID types.String                                           `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsAssignOrganizationLicensesSeats  `tfsdk:"item"`
	Parameters     *RequestOrganizationsAssignOrganizationLicensesSeatsRs `tfsdk:"parameters"`
}

type ResponseOrganizationsAssignOrganizationLicensesSeats struct {
	ResultingLicenses *[]ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicenses `tfsdk:"resulting_licenses"`
}

type ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicenses struct {
	ActivationDate            types.String                                                                                      `tfsdk:"activation_date"`
	ClaimDate                 types.String                                                                                      `tfsdk:"claim_date"`
	DeviceSerial              types.String                                                                                      `tfsdk:"device_serial"`
	DurationInDays            types.Int64                                                                                       `tfsdk:"duration_in_days"`
	ExpirationDate            types.String                                                                                      `tfsdk:"expiration_date"`
	HeadLicenseID             types.String                                                                                      `tfsdk:"head_license_id"`
	ID                        types.String                                                                                      `tfsdk:"id"`
	LicenseKey                types.String                                                                                      `tfsdk:"license_key"`
	LicenseType               types.String                                                                                      `tfsdk:"license_type"`
	NetworkID                 types.String                                                                                      `tfsdk:"network_id"`
	OrderNumber               types.String                                                                                      `tfsdk:"order_number"`
	PermanentlyQueuedLicenses *[]ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses `tfsdk:"permanently_queued_licenses"`
	SeatCount                 types.Int64                                                                                       `tfsdk:"seat_count"`
	State                     types.String                                                                                      `tfsdk:"state"`
	TotalDurationInDays       types.Int64                                                                                       `tfsdk:"total_duration_in_days"`
}

type ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses struct {
	DurationInDays types.Int64  `tfsdk:"duration_in_days"`
	ID             types.String `tfsdk:"id"`
	LicenseKey     types.String `tfsdk:"license_key"`
	LicenseType    types.String `tfsdk:"license_type"`
	OrderNumber    types.String `tfsdk:"order_number"`
}

type RequestOrganizationsAssignOrganizationLicensesSeatsRs struct {
	LicenseID types.String `tfsdk:"license_id"`
	NetworkID types.String `tfsdk:"network_id"`
	SeatCount types.Int64  `tfsdk:"seat_count"`
}

// FromBody
func (r *OrganizationsLicensesAssignSeats) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsAssignOrganizationLicensesSeats {
	emptyString := ""
	re := *r.Parameters
	licenseID := new(string)
	if !re.LicenseID.IsUnknown() && !re.LicenseID.IsNull() {
		*licenseID = re.LicenseID.ValueString()
	} else {
		licenseID = &emptyString
	}
	networkID := new(string)
	if !re.NetworkID.IsUnknown() && !re.NetworkID.IsNull() {
		*networkID = re.NetworkID.ValueString()
	} else {
		networkID = &emptyString
	}
	seatCount := new(int64)
	if !re.SeatCount.IsUnknown() && !re.SeatCount.IsNull() {
		*seatCount = re.SeatCount.ValueInt64()
	} else {
		seatCount = nil
	}
	out := merakigosdk.RequestOrganizationsAssignOrganizationLicensesSeats{
		LicenseID: *licenseID,
		NetworkID: *networkID,
		SeatCount: int64ToIntPointer(seatCount),
	}
	return &out
}

// ToBody
func ResponseOrganizationsAssignOrganizationLicensesSeatsItemToBody(state OrganizationsLicensesAssignSeats, response *merakigosdk.ResponseOrganizationsAssignOrganizationLicensesSeats) OrganizationsLicensesAssignSeats {
	itemState := ResponseOrganizationsAssignOrganizationLicensesSeats{
		ResultingLicenses: func() *[]ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicenses {
			if response.ResultingLicenses != nil {
				result := make([]ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicenses, len(*response.ResultingLicenses))
				for i, resultingLicenses := range *response.ResultingLicenses {
					result[i] = ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicenses{
						ActivationDate: types.StringValue(resultingLicenses.ActivationDate),
						ClaimDate:      types.StringValue(resultingLicenses.ClaimDate),
						DeviceSerial:   types.StringValue(resultingLicenses.DeviceSerial),
						DurationInDays: func() types.Int64 {
							if resultingLicenses.DurationInDays != nil {
								return types.Int64Value(int64(*resultingLicenses.DurationInDays))
							}
							return types.Int64{}
						}(),
						ExpirationDate: types.StringValue(resultingLicenses.ExpirationDate),
						HeadLicenseID:  types.StringValue(resultingLicenses.HeadLicenseID),
						ID:             types.StringValue(resultingLicenses.ID),
						LicenseKey:     types.StringValue(resultingLicenses.LicenseKey),
						LicenseType:    types.StringValue(resultingLicenses.LicenseType),
						NetworkID:      types.StringValue(resultingLicenses.NetworkID),
						OrderNumber:    types.StringValue(resultingLicenses.OrderNumber),
						PermanentlyQueuedLicenses: func() *[]ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses {
							if resultingLicenses.PermanentlyQueuedLicenses != nil {
								result := make([]ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses, len(*resultingLicenses.PermanentlyQueuedLicenses))
								for i, permanentlyQueuedLicenses := range *resultingLicenses.PermanentlyQueuedLicenses {
									result[i] = ResponseOrganizationsAssignOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses{
										DurationInDays: func() types.Int64 {
											if permanentlyQueuedLicenses.DurationInDays != nil {
												return types.Int64Value(int64(*permanentlyQueuedLicenses.DurationInDays))
											}
											return types.Int64{}
										}(),
										ID:          types.StringValue(permanentlyQueuedLicenses.ID),
										LicenseKey:  types.StringValue(permanentlyQueuedLicenses.LicenseKey),
										LicenseType: types.StringValue(permanentlyQueuedLicenses.LicenseType),
										OrderNumber: types.StringValue(permanentlyQueuedLicenses.OrderNumber),
									}
								}
								return &result
							}
							return nil
						}(),
						SeatCount: func() types.Int64 {
							if resultingLicenses.SeatCount != nil {
								return types.Int64Value(int64(*resultingLicenses.SeatCount))
							}
							return types.Int64{}
						}(),
						State: types.StringValue(resultingLicenses.State),
						TotalDurationInDays: func() types.Int64 {
							if resultingLicenses.TotalDurationInDays != nil {
								return types.Int64Value(int64(*resultingLicenses.TotalDurationInDays))
							}
							return types.Int64{}
						}(),
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
