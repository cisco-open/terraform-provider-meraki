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

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &OrganizationsLicensesDataSource{}
	_ datasource.DataSourceWithConfigure = &OrganizationsLicensesDataSource{}
)

func NewOrganizationsLicensesDataSource() datasource.DataSource {
	return &OrganizationsLicensesDataSource{}
}

type OrganizationsLicensesDataSource struct {
	client *merakigosdk.Client
}

func (d *OrganizationsLicensesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *OrganizationsLicensesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_licenses"
}

func (d *OrganizationsLicensesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"license_id": schema.StringAttribute{
				MarkdownDescription: `licenseId path parameter. License ID`,
				Required:            true,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
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
	}
}

func (d *OrganizationsLicensesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var organizationsLicenses OrganizationsLicenses
	diags := req.Config.Get(ctx, &organizationsLicenses)
	if resp.Diagnostics.HasError() {
		return
	}

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetOrganizationLicense")
		vvOrganizationID := organizationsLicenses.OrganizationID.ValueString()
		vvLicenseID := organizationsLicenses.LicenseID.ValueString()

		response1, restyResp1, err := d.client.Organizations.GetOrganizationLicense(vvOrganizationID, vvLicenseID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationLicense",
				err.Error(),
			)
			return
		}

		organizationsLicenses = ResponseOrganizationsGetOrganizationLicenseItemToBody(organizationsLicenses, response1)
		diags = resp.State.Set(ctx, &organizationsLicenses)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type OrganizationsLicenses struct {
	OrganizationID types.String                                 `tfsdk:"organization_id"`
	LicenseID      types.String                                 `tfsdk:"license_id"`
	Item           *ResponseOrganizationsGetOrganizationLicense `tfsdk:"item"`
}

type ResponseOrganizationsGetOrganizationLicense struct {
	ActivationDate            types.String                                                            `tfsdk:"activation_date"`
	ClaimDate                 types.String                                                            `tfsdk:"claim_date"`
	DeviceSerial              types.String                                                            `tfsdk:"device_serial"`
	DurationInDays            types.Int64                                                             `tfsdk:"duration_in_days"`
	ExpirationDate            types.String                                                            `tfsdk:"expiration_date"`
	HeadLicenseID             types.String                                                            `tfsdk:"head_license_id"`
	ID                        types.String                                                            `tfsdk:"id"`
	LicenseKey                types.String                                                            `tfsdk:"license_key"`
	LicenseType               types.String                                                            `tfsdk:"license_type"`
	NetworkID                 types.String                                                            `tfsdk:"network_id"`
	OrderNumber               types.String                                                            `tfsdk:"order_number"`
	PermanentlyQueuedLicenses *[]ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicenses `tfsdk:"permanently_queued_licenses"`
	SeatCount                 types.Int64                                                             `tfsdk:"seat_count"`
	State                     types.String                                                            `tfsdk:"state"`
	TotalDurationInDays       types.Int64                                                             `tfsdk:"total_duration_in_days"`
}

type ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicenses struct {
	DurationInDays types.Int64  `tfsdk:"duration_in_days"`
	ID             types.String `tfsdk:"id"`
	LicenseKey     types.String `tfsdk:"license_key"`
	LicenseType    types.String `tfsdk:"license_type"`
	OrderNumber    types.String `tfsdk:"order_number"`
}

// ToBody
func ResponseOrganizationsGetOrganizationLicenseItemToBody(state OrganizationsLicenses, response *merakigosdk.ResponseOrganizationsGetOrganizationLicense) OrganizationsLicenses {
	itemState := ResponseOrganizationsGetOrganizationLicense{
		ActivationDate: types.StringValue(response.ActivationDate),
		ClaimDate:      types.StringValue(response.ClaimDate),
		DeviceSerial:   types.StringValue(response.DeviceSerial),
		DurationInDays: func() types.Int64 {
			if response.DurationInDays != nil {
				return types.Int64Value(int64(*response.DurationInDays))
			}
			return types.Int64{}
		}(),
		ExpirationDate: types.StringValue(response.ExpirationDate),
		HeadLicenseID:  types.StringValue(response.HeadLicenseID),
		ID:             types.StringValue(response.ID),
		LicenseKey:     types.StringValue(response.LicenseKey),
		LicenseType:    types.StringValue(response.LicenseType),
		NetworkID:      types.StringValue(response.NetworkID),
		OrderNumber:    types.StringValue(response.OrderNumber),
		PermanentlyQueuedLicenses: func() *[]ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicenses {
			if response.PermanentlyQueuedLicenses != nil {
				result := make([]ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicenses, len(*response.PermanentlyQueuedLicenses))
				for i, permanentlyQueuedLicenses := range *response.PermanentlyQueuedLicenses {
					result[i] = ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicenses{
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
			return &[]ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicenses{}
		}(),
		SeatCount: func() types.Int64 {
			if response.SeatCount != nil {
				return types.Int64Value(int64(*response.SeatCount))
			}
			return types.Int64{}
		}(),
		State: types.StringValue(response.State),
		TotalDurationInDays: func() types.Int64 {
			if response.TotalDurationInDays != nil {
				return types.Int64Value(int64(*response.TotalDurationInDays))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
