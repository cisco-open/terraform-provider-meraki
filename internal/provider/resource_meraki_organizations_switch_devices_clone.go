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
	_ resource.Resource              = &OrganizationsSwitchDevicesCloneResource{}
	_ resource.ResourceWithConfigure = &OrganizationsSwitchDevicesCloneResource{}
)

func NewOrganizationsSwitchDevicesCloneResource() resource.Resource {
	return &OrganizationsSwitchDevicesCloneResource{}
}

type OrganizationsSwitchDevicesCloneResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsSwitchDevicesCloneResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsSwitchDevicesCloneResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_switch_devices_clone"
}

// resourceAction
func (r *OrganizationsSwitchDevicesCloneResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"source_serial": schema.StringAttribute{
						MarkdownDescription: `Serial number of the source switch (must be on a network not bound to a template)`,
						Computed:            true,
					},
					"target_serials": schema.SetAttribute{
						MarkdownDescription: `Array of serial numbers of one or more target switches (must be on a network not bound to a template)`,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"source_serial": schema.StringAttribute{
						MarkdownDescription: `Serial number of the source switch (must be on a network not bound to a template)`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
					"target_serials": schema.SetAttribute{
						MarkdownDescription: `Array of serial numbers of one or more target switches (must be on a network not bound to a template)`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *OrganizationsSwitchDevicesCloneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsSwitchDevicesClone

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
	response, restyResp1, err := r.client.Switch.CloneOrganizationSwitchDevices(vvOrganizationID, dataRequest)

	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CloneOrganizationSwitchDevices",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CloneOrganizationSwitchDevices",
			err.Error(),
		)
		return
	}
	//Item
	// //entro aqui 2
	data = ResponseSwitchCloneOrganizationSwitchDevicesItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsSwitchDevicesCloneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsSwitchDevicesCloneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsSwitchDevicesCloneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsSwitchDevicesClone struct {
	OrganizationID types.String                                   `tfsdk:"organization_id"`
	Item           *ResponseSwitchCloneOrganizationSwitchDevices  `tfsdk:"item"`
	Parameters     *RequestSwitchCloneOrganizationSwitchDevicesRs `tfsdk:"parameters"`
}

type ResponseSwitchCloneOrganizationSwitchDevices struct {
	SourceSerial  types.String `tfsdk:"source_serial"`
	TargetSerials types.Set    `tfsdk:"target_serials"`
}

type RequestSwitchCloneOrganizationSwitchDevicesRs struct {
	SourceSerial  types.String `tfsdk:"source_serial"`
	TargetSerials types.Set    `tfsdk:"target_serials"`
}

// FromBody
func (r *OrganizationsSwitchDevicesClone) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCloneOrganizationSwitchDevices {
	emptyString := ""
	re := *r.Parameters
	sourceSerial := new(string)
	if !re.SourceSerial.IsUnknown() && !re.SourceSerial.IsNull() {
		*sourceSerial = re.SourceSerial.ValueString()
	} else {
		sourceSerial = &emptyString
	}
	var targetSerials []string = nil
	re.TargetSerials.ElementsAs(ctx, &targetSerials, false)
	out := merakigosdk.RequestSwitchCloneOrganizationSwitchDevices{
		SourceSerial:  *sourceSerial,
		TargetSerials: targetSerials,
	}
	return &out
}

// ToBody
func ResponseSwitchCloneOrganizationSwitchDevicesItemToBody(state OrganizationsSwitchDevicesClone, response *merakigosdk.ResponseSwitchCloneOrganizationSwitchDevices) OrganizationsSwitchDevicesClone {
	itemState := ResponseSwitchCloneOrganizationSwitchDevices{
		SourceSerial:  types.StringValue(response.SourceSerial),
		TargetSerials: StringSliceToSet(response.TargetSerials),
	}
	state.Item = &itemState
	return state
}
