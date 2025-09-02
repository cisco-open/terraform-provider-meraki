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
	"strconv"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsApplianceDNSLocalRecordsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsApplianceDNSLocalRecordsResource{}
)

func NewOrganizationsApplianceDNSLocalRecordsResource() resource.Resource {
	return &OrganizationsApplianceDNSLocalRecordsResource{}
}

type OrganizationsApplianceDNSLocalRecordsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsApplianceDNSLocalRecordsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsApplianceDNSLocalRecordsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_appliance_dns_local_records"
}

func (r *OrganizationsApplianceDNSLocalRecordsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				MarkdownDescription: `IP for the DNS record`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: `Hostname for the DNS record`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"profile": schema.SingleNestedAttribute{
				MarkdownDescription: `The profile the DNS record is associated with`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"id": schema.StringAttribute{
						MarkdownDescription: `Profile ID`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"record_id": schema.StringAttribute{
				MarkdownDescription: `Record ID`,
				Computed:            true,
			},
		},
	}
}

func (r *OrganizationsApplianceDNSLocalRecordsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsApplianceDNSLocalRecordsRs

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
	//Only Items

	vvHostname := data.Hostname.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Appliance.GetOrganizationApplianceDNSLocalRecords(vvOrganizationID, nil)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationApplianceDNSLocalRecords",
					restyResp1.String(),
				)
				return
			}
		}
	}

	var responseVerifyItem2 merakigosdk.ResponseItemApplianceGetOrganizationApplianceDNSLocalRecords
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Hostname", vvHostname, simpleCmp)
		if result != nil {
			err := mapToStruct(result.(map[string]interface{}), &responseVerifyItem2)
			if err != nil {
				resp.Diagnostics.AddError(
					"Failure when executing mapToStruct in resource",
					err.Error(),
				)
				return
			}
			data = ResponseApplianceGetOrganizationApplianceDNSLocalRecordsItemToBodyRs(data, &responseVerifyItem2, false)
			// Path params update assigned
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return

		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	_, restyResp2, err := r.client.Appliance.CreateOrganizationApplianceDNSLocalRecord(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationApplianceDNSLocalRecord",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationApplianceDNSLocalRecord",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Appliance.GetOrganizationApplianceDNSLocalRecords(vvOrganizationID, nil)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceDNSLocalRecords",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationApplianceDNSLocalRecords",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result2 := getDictResult(responseStruct, "Hostname", vvHostname, simpleCmp)
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		data = ResponseApplianceGetOrganizationApplianceDNSLocalRecordsItemToBodyRs(data, &responseVerifyItem2, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationApplianceDNSLocalRecords Result",
			"Not Found",
		)
		return
	}

}

func (r *OrganizationsApplianceDNSLocalRecordsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsApplianceDNSLocalRecordsRs

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
	// Not has Item

	vvOrganizationID := data.OrganizationID.ValueString()
	vvHostname := data.Hostname.ValueString()

	responseGet, restyResp1, err := r.client.Appliance.GetOrganizationApplianceDNSLocalRecords(vvOrganizationID, nil)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationApplianceDNSLocalRecords",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationApplianceDNSLocalRecords",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet)
	result2 := getDictResult(responseStruct2, "Hostname", vvHostname, simpleCmp)
	var responseVerifyItem2 merakigosdk.ResponseItemApplianceGetOrganizationApplianceDNSLocalRecords
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		//entro aqui
		data = ResponseApplianceGetOrganizationApplianceDNSLocalRecordsItemToBodyRs(data, &responseVerifyItem2, true)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationApplianceDNSLocalRecords Result",
			err.Error(),
		)
		return
	}
}

func (r *OrganizationsApplianceDNSLocalRecordsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("hostname"), idParts[1])...)
}

func (r *OrganizationsApplianceDNSLocalRecordsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsApplianceDNSLocalRecordsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update
	// No update
	resp.Diagnostics.AddError(
		"Update operation not supported in OrganizationsApplianceDNSLocalRecords",
		"Update operation not supported in OrganizationsApplianceDNSLocalRecords",
	)
	return
}

func (r *OrganizationsApplianceDNSLocalRecordsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting OrganizationsApplianceDNSLocalRecords", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsApplianceDNSLocalRecordsRs struct {
	OrganizationID types.String `tfsdk:"organization_id"`
	//TIENE ITEMS
	Address  types.String                                                           `tfsdk:"address"`
	Hostname types.String                                                           `tfsdk:"hostname"`
	Profile  *ResponseItemApplianceGetOrganizationApplianceDnsLocalRecordsProfileRs `tfsdk:"profile"`
	RecordID types.String                                                           `tfsdk:"record_id"`
}

type ResponseItemApplianceGetOrganizationApplianceDnsLocalRecordsProfileRs struct {
	ID types.String `tfsdk:"id"`
}

// FromBody
func (r *OrganizationsApplianceDNSLocalRecordsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestApplianceCreateOrganizationApplianceDNSLocalRecord {
	emptyString := ""
	address := new(string)
	if !r.Address.IsUnknown() && !r.Address.IsNull() {
		*address = r.Address.ValueString()
	} else {
		address = &emptyString
	}
	Hostname := new(string)
	if !r.Hostname.IsUnknown() && !r.Hostname.IsNull() {
		*Hostname = r.Hostname.ValueString()
	} else {
		Hostname = &emptyString
	}
	var requestApplianceCreateOrganizationApplianceDNSLocalRecordProfile *merakigosdk.RequestApplianceCreateOrganizationApplianceDNSLocalRecordProfile

	if r.Profile != nil {
		id := r.Profile.ID.ValueString()
		requestApplianceCreateOrganizationApplianceDNSLocalRecordProfile = &merakigosdk.RequestApplianceCreateOrganizationApplianceDNSLocalRecordProfile{
			ID: id,
		}
		//[debug] Is Array: False
	}
	out := merakigosdk.RequestApplianceCreateOrganizationApplianceDNSLocalRecord{
		Address:  *address,
		Hostname: *Hostname,
		Profile:  requestApplianceCreateOrganizationApplianceDNSLocalRecordProfile,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetOrganizationApplianceDNSLocalRecordsItemToBodyRs(state OrganizationsApplianceDNSLocalRecordsRs, response *merakigosdk.ResponseItemApplianceGetOrganizationApplianceDNSLocalRecords, is_read bool) OrganizationsApplianceDNSLocalRecordsRs {
	itemState := OrganizationsApplianceDNSLocalRecordsRs{
		Address: func() types.String {
			if response.Address != "" {
				return types.StringValue(response.Address)
			}
			return types.String{}
		}(),
		Hostname: func() types.String {
			if response.Hostname != "" {
				return types.StringValue(response.Hostname)
			}
			return types.String{}
		}(),
		Profile: func() *ResponseItemApplianceGetOrganizationApplianceDnsLocalRecordsProfileRs {
			if response.Profile != nil {
				return &ResponseItemApplianceGetOrganizationApplianceDnsLocalRecordsProfileRs{
					ID: func() types.String {
						if response.Profile.ID != "" {
							return types.StringValue(response.Profile.ID)
						}
						return types.String{}
					}(),
				}
			}
			return nil
		}(),
		RecordID: func() types.String {
			if response.RecordID != "" {
				return types.StringValue(response.RecordID)
			}
			return types.String{}
		}(),
	}
	state = itemState
	return state
}
