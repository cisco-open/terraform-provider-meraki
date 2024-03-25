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
	_ resource.Resource              = &OrganizationsLicensesRenewSeatsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsLicensesRenewSeatsResource{}
)

func NewOrganizationsLicensesRenewSeatsResource() resource.Resource {
	return &OrganizationsLicensesRenewSeatsResource{}
}

type OrganizationsLicensesRenewSeatsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsLicensesRenewSeatsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsLicensesRenewSeatsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_licenses_renew_seats"
}

// resourceAction
func (r *OrganizationsLicensesRenewSeatsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
									MarkdownDescription: `The state of the license. All queued licenses have a status of *recentlyQueued*.`,
									Computed:            true,
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
					"license_id_to_renew": schema.StringAttribute{
						MarkdownDescription: `The ID of the SM license to renew. This license must already be assigned to an SM network`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"unused_license_id": schema.StringAttribute{
						MarkdownDescription: `The SM license to use to renew the seats on 'licenseIdToRenew'. This license must have at least as many seats available as there are seats on 'licenseIdToRenew'`,
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
func (r *OrganizationsLicensesRenewSeatsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsLicensesRenewSeats

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
	response, restyResp1, err := r.client.Organizations.RenewOrganizationLicensesSeats(vvOrganizationID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing RenewOrganizationLicensesSeats",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing RenewOrganizationLicensesSeats",
			err.Error(),
		)
		return
	}
	//Item
	data2 := ResponseOrganizationsRenewOrganizationLicensesSeatsItemToBody(data, response)

	diags := resp.State.Set(ctx, &data2)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsLicensesRenewSeatsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsLicensesRenewSeatsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsLicensesRenewSeatsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsLicensesRenewSeats struct {
	OrganizationID types.String                                          `tfsdk:"organization_id"`
	Item           *ResponseOrganizationsRenewOrganizationLicensesSeats  `tfsdk:"item"`
	Parameters     *RequestOrganizationsRenewOrganizationLicensesSeatsRs `tfsdk:"parameters"`
}

type ResponseOrganizationsRenewOrganizationLicensesSeats struct {
	ResultingLicenses *[]ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicenses `tfsdk:"resulting_licenses"`
}

type ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicenses struct {
	ActivationDate            types.String                                                                                     `tfsdk:"activation_date"`
	ClaimDate                 types.String                                                                                     `tfsdk:"claim_date"`
	DeviceSerial              types.String                                                                                     `tfsdk:"device_serial"`
	DurationInDays            types.Int64                                                                                      `tfsdk:"duration_in_days"`
	ExpirationDate            types.String                                                                                     `tfsdk:"expiration_date"`
	HeadLicenseID             types.String                                                                                     `tfsdk:"head_license_id"`
	ID                        types.String                                                                                     `tfsdk:"id"`
	LicenseKey                types.String                                                                                     `tfsdk:"license_key"`
	LicenseType               types.String                                                                                     `tfsdk:"license_type"`
	NetworkID                 types.String                                                                                     `tfsdk:"network_id"`
	OrderNumber               types.String                                                                                     `tfsdk:"order_number"`
	PermanentlyQueuedLicenses *[]ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses `tfsdk:"permanently_queued_licenses"`
	SeatCount                 types.Int64                                                                                      `tfsdk:"seat_count"`
	State                     types.String                                                                                     `tfsdk:"state"`
	TotalDurationInDays       types.Int64                                                                                      `tfsdk:"total_duration_in_days"`
}

type ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses struct {
	DurationInDays types.Int64  `tfsdk:"duration_in_days"`
	ID             types.String `tfsdk:"id"`
	LicenseKey     types.String `tfsdk:"license_key"`
	LicenseType    types.String `tfsdk:"license_type"`
	OrderNumber    types.String `tfsdk:"order_number"`
}

type RequestOrganizationsRenewOrganizationLicensesSeatsRs struct {
	LicenseIDToRenew types.String `tfsdk:"license_id_to_renew"`
	UnusedLicenseID  types.String `tfsdk:"unused_license_id"`
}

// FromBody
func (r *OrganizationsLicensesRenewSeats) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsRenewOrganizationLicensesSeats {
	emptyString := ""
	re := *r.Parameters
	licenseIDToRenew := new(string)
	if !re.LicenseIDToRenew.IsUnknown() && !re.LicenseIDToRenew.IsNull() {
		*licenseIDToRenew = re.LicenseIDToRenew.ValueString()
	} else {
		licenseIDToRenew = &emptyString
	}
	unusedLicenseID := new(string)
	if !re.UnusedLicenseID.IsUnknown() && !re.UnusedLicenseID.IsNull() {
		*unusedLicenseID = re.UnusedLicenseID.ValueString()
	} else {
		unusedLicenseID = &emptyString
	}
	out := merakigosdk.RequestOrganizationsRenewOrganizationLicensesSeats{
		LicenseIDToRenew: *licenseIDToRenew,
		UnusedLicenseID:  *unusedLicenseID,
	}
	return &out
}

// ToBody
func ResponseOrganizationsRenewOrganizationLicensesSeatsItemToBody(state OrganizationsLicensesRenewSeats, response *merakigosdk.ResponseOrganizationsRenewOrganizationLicensesSeats) OrganizationsLicensesRenewSeats {
	itemState := ResponseOrganizationsRenewOrganizationLicensesSeats{
		ResultingLicenses: func() *[]ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicenses {
			if response.ResultingLicenses != nil {
				result := make([]ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicenses, len(*response.ResultingLicenses))
				for i, resultingLicenses := range *response.ResultingLicenses {
					result[i] = ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicenses{
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
						PermanentlyQueuedLicenses: func() *[]ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses {
							if resultingLicenses.PermanentlyQueuedLicenses != nil {
								result := make([]ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses, len(*resultingLicenses.PermanentlyQueuedLicenses))
								for i, permanentlyQueuedLicenses := range *resultingLicenses.PermanentlyQueuedLicenses {
									result[i] = ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses{
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
							return &[]ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicensesPermanentlyQueuedLicenses{}
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
			return &[]ResponseOrganizationsRenewOrganizationLicensesSeatsResultingLicenses{}
		}(),
	}
	state.Item = &itemState
	return state
}
