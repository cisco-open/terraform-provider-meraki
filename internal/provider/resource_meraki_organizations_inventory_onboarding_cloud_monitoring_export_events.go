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
	_ resource.Resource              = &OrganizationsInventoryOnboardingCloudMonitoringExportEventsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsInventoryOnboardingCloudMonitoringExportEventsResource{}
)

func NewOrganizationsInventoryOnboardingCloudMonitoringExportEventsResource() resource.Resource {
	return &OrganizationsInventoryOnboardingCloudMonitoringExportEventsResource{}
}

type OrganizationsInventoryOnboardingCloudMonitoringExportEventsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsInventoryOnboardingCloudMonitoringExportEventsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsInventoryOnboardingCloudMonitoringExportEventsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_inventory_onboarding_cloud_monitoring_export_events"
}

// resourceAction
func (r *OrganizationsInventoryOnboardingCloudMonitoringExportEventsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					"log_event": schema.StringAttribute{
						MarkdownDescription: `The type of log event this is recording, e.g. download or opening a banner`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"request": schema.StringAttribute{
						MarkdownDescription: `Used to describe if this event was the result of a redirect. E.g. a query param if an info banner is being used`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"target_os": schema.StringAttribute{
						MarkdownDescription: `The name of the onboarding distro being downloaded`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"timestamp": schema.Int64Attribute{
						MarkdownDescription: `A JavaScript UTC datetime stamp for when the even occurred`,
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
func (r *OrganizationsInventoryOnboardingCloudMonitoringExportEventsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsInventoryOnboardingCloudMonitoringExportEvents

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
	restyResp1, err := r.client.Organizations.CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(vvOrganizationID, dataRequest)
	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationInventoryOnboardingCloudMonitoringExportEvent",
			err.Error(),
		)
		return
	}
	//Item
	// //entro aqui 2
	// data2 := ResponseOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringExportEvent(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsInventoryOnboardingCloudMonitoringExportEventsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsInventoryOnboardingCloudMonitoringExportEventsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsInventoryOnboardingCloudMonitoringExportEventsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsInventoryOnboardingCloudMonitoringExportEvents struct {
	OrganizationID types.String                                                                           `tfsdk:"organization_id"`
	Parameters     *RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRs `tfsdk:"parameters"`
}

type RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringExportEventRs struct {
	LogEvent  types.String `tfsdk:"log_event"`
	Request   types.String `tfsdk:"request"`
	TargetOS  types.String `tfsdk:"target_os"`
	Timestamp types.Int64  `tfsdk:"timestamp"`
}

// FromBody
func (r *OrganizationsInventoryOnboardingCloudMonitoringExportEvents) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringExportEvent {
	emptyString := ""
	re := *r.Parameters
	logEvent := new(string)
	if !re.LogEvent.IsUnknown() && !re.LogEvent.IsNull() {
		*logEvent = re.LogEvent.ValueString()
	} else {
		logEvent = &emptyString
	}
	request := new(string)
	if !re.Request.IsUnknown() && !re.Request.IsNull() {
		*request = re.Request.ValueString()
	} else {
		request = &emptyString
	}
	targetOS := new(string)
	if !re.TargetOS.IsUnknown() && !re.TargetOS.IsNull() {
		*targetOS = re.TargetOS.ValueString()
	} else {
		targetOS = &emptyString
	}
	timestamp := new(int64)
	if !re.Timestamp.IsUnknown() && !re.Timestamp.IsNull() {
		*timestamp = re.Timestamp.ValueInt64()
	} else {
		timestamp = nil
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationInventoryOnboardingCloudMonitoringExportEvent{
		LogEvent:  *logEvent,
		Request:   *request,
		TargetOS:  *targetOS,
		Timestamp: int64ToIntPointer(timestamp),
	}
	return &out
}
