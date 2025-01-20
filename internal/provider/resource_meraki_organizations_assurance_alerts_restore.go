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
	_ resource.Resource              = &OrganizationsAssuranceAlertsRestoreResource{}
	_ resource.ResourceWithConfigure = &OrganizationsAssuranceAlertsRestoreResource{}
)

func NewOrganizationsAssuranceAlertsRestoreResource() resource.Resource {
	return &OrganizationsAssuranceAlertsRestoreResource{}
}

type OrganizationsAssuranceAlertsRestoreResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsAssuranceAlertsRestoreResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsAssuranceAlertsRestoreResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_assurance_alerts_restore"
}

// resourceAction
func (r *OrganizationsAssuranceAlertsRestoreResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					"alert_ids": schema.SetAttribute{
						MarkdownDescription: `Array of alert IDs to restore`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *OrganizationsAssuranceAlertsRestoreResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsAssuranceAlertsRestore

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
	restyResp1, err := r.client.Organizations.RestoreOrganizationAssuranceAlerts(vvOrganizationID, dataRequest)

	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing RestoreOrganizationAssuranceAlerts",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing RestoreOrganizationAssuranceAlerts",
			err.Error(),
		)
		return
	}
	//Item
	// //entro aqui 2
	// data2 := ResponseOrganizationsRestoreOrganizationAssuranceAlerts(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsAssuranceAlertsRestoreResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsAssuranceAlertsRestoreResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsAssuranceAlertsRestoreResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsAssuranceAlertsRestore struct {
	OrganizationID types.String                                              `tfsdk:"organization_id"`
	Parameters     *RequestOrganizationsRestoreOrganizationAssuranceAlertsRs `tfsdk:"parameters"`
}

type RequestOrganizationsRestoreOrganizationAssuranceAlertsRs struct {
	AlertIDs types.Set `tfsdk:"alert_ids"`
}

// FromBody
func (r *OrganizationsAssuranceAlertsRestore) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsRestoreOrganizationAssuranceAlerts {
	re := *r.Parameters
	var alertIDs []string = nil
	re.AlertIDs.ElementsAs(ctx, &alertIDs, false)
	out := merakigosdk.RequestOrganizationsRestoreOrganizationAssuranceAlerts{
		AlertIDs: alertIDs,
	}
	return &out
}
