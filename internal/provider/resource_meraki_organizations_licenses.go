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
	"fmt"
	"strings"

	merakigosdk "dashboard-api-go/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsLicensesResource{}
	_ resource.ResourceWithConfigure = &OrganizationsLicensesResource{}
)

func NewOrganizationsLicensesResource() resource.Resource {
	return &OrganizationsLicensesResource{}
}

type OrganizationsLicensesResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsLicensesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsLicensesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_licenses"
}

func (r *OrganizationsLicensesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
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
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
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
			"license_id": schema.StringAttribute{
				MarkdownDescription: `licenseId path parameter. License ID`,
				Required:            true,
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
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
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
	}
}

func (r *OrganizationsLicensesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsLicensesRs

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
	// Has Paths
	vvOrganizationID := data.OrganizationID.ValueString()
	vvLicenseID := data.LicenseID.ValueString()
	//Has Item and not has items

	if vvOrganizationID != "" && vvLicenseID != "" {
		//dentro
		responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationLicense(vvOrganizationID, vvLicenseID)
		// No Post
		if err != nil || restyResp1 == nil || responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource OrganizationsLicenses  only have update context, not create.",
				err.Error(),
			)
			return
		}

		if responseVerifyItem == nil {
			resp.Diagnostics.AddError(
				"Resource OrganizationsLicenses only have update context, not create.",
				err.Error(),
			)
			return
		}
	}

	// UPDATE NO CREATE
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationLicense(vvOrganizationID, vvLicenseID, dataRequest)
	//Update
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationLicense",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationLicense",
			err.Error(),
		)
		return
	}

	//Assign Path Params required

	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationLicense(vvOrganizationID, vvLicenseID)
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationLicense",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationLicense",
			err.Error(),
		)
		return
	}

	data = ResponseOrganizationsGetOrganizationLicenseItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)

}

func (r *OrganizationsLicensesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsLicensesRs

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
	vvLicenseID := data.LicenseID.ValueString()
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationLicense(vvOrganizationID, vvLicenseID)
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
				"Failure when executing GetOrganizationLicense",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationLicense",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationLicenseItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsLicensesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("license_id"), idParts[1])...)
}

func (r *OrganizationsLicensesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsLicensesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvLicenseID := data.LicenseID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationLicense(vvOrganizationID, vvLicenseID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationLicense",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationLicense",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsLicensesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting OrganizationsLicenses", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsLicensesRs struct {
	OrganizationID            types.String                                                              `tfsdk:"organization_id"`
	LicenseID                 types.String                                                              `tfsdk:"license_id"`
	ActivationDate            types.String                                                              `tfsdk:"activation_date"`
	ClaimDate                 types.String                                                              `tfsdk:"claim_date"`
	DeviceSerial              types.String                                                              `tfsdk:"device_serial"`
	DurationInDays            types.Int64                                                               `tfsdk:"duration_in_days"`
	ExpirationDate            types.String                                                              `tfsdk:"expiration_date"`
	HeadLicenseID             types.String                                                              `tfsdk:"head_license_id"`
	ID                        types.String                                                              `tfsdk:"id"`
	LicenseKey                types.String                                                              `tfsdk:"license_key"`
	LicenseType               types.String                                                              `tfsdk:"license_type"`
	NetworkID                 types.String                                                              `tfsdk:"network_id"`
	OrderNumber               types.String                                                              `tfsdk:"order_number"`
	PermanentlyQueuedLicenses *[]ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicensesRs `tfsdk:"permanently_queued_licenses"`
	SeatCount                 types.Int64                                                               `tfsdk:"seat_count"`
	State                     types.String                                                              `tfsdk:"state"`
	TotalDurationInDays       types.Int64                                                               `tfsdk:"total_duration_in_days"`
}

type ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicensesRs struct {
	DurationInDays types.Int64  `tfsdk:"duration_in_days"`
	ID             types.String `tfsdk:"id"`
	LicenseKey     types.String `tfsdk:"license_key"`
	LicenseType    types.String `tfsdk:"license_type"`
	OrderNumber    types.String `tfsdk:"order_number"`
}

// FromBody
func (r *OrganizationsLicensesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationLicense {
	emptyString := ""
	deviceSerial := new(string)
	if !r.DeviceSerial.IsUnknown() && !r.DeviceSerial.IsNull() {
		*deviceSerial = r.DeviceSerial.ValueString()
	} else {
		deviceSerial = &emptyString
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationLicense{
		DeviceSerial: *deviceSerial,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationLicenseItemToBodyRs(state OrganizationsLicensesRs, response *merakigosdk.ResponseOrganizationsGetOrganizationLicense, is_read bool) OrganizationsLicensesRs {
	itemState := OrganizationsLicensesRs{
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
		PermanentlyQueuedLicenses: func() *[]ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicensesRs {
			if response.PermanentlyQueuedLicenses != nil {
				result := make([]ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicensesRs, len(*response.PermanentlyQueuedLicenses))
				for i, permanentlyQueuedLicenses := range *response.PermanentlyQueuedLicenses {
					result[i] = ResponseOrganizationsGetOrganizationLicensePermanentlyQueuedLicensesRs{
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
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsLicensesRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsLicensesRs)
}
