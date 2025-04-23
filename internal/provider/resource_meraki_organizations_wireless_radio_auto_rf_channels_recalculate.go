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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsWirelessRadioAutoRfChannelsRecalculateResource{}
	_ resource.ResourceWithConfigure = &OrganizationsWirelessRadioAutoRfChannelsRecalculateResource{}
)

func NewOrganizationsWirelessRadioAutoRfChannelsRecalculateResource() resource.Resource {
	return &OrganizationsWirelessRadioAutoRfChannelsRecalculateResource{}
}

type OrganizationsWirelessRadioAutoRfChannelsRecalculateResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsWirelessRadioAutoRfChannelsRecalculateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsWirelessRadioAutoRfChannelsRecalculateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_radio_auto_rf_channels_recalculate"
}

// resourceAction
func (r *OrganizationsWirelessRadioAutoRfChannelsRecalculateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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

					"estimated_completed_at": schema.StringAttribute{
						MarkdownDescription: `Estimated time of completion.`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"network_ids": schema.ListAttribute{
						MarkdownDescription: `A list of network ids (limit: 15).`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *OrganizationsWirelessRadioAutoRfChannelsRecalculateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsWirelessRadioAutoRfChannelsRecalculate

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
	response, restyResp1, err := r.client.Wireless.RecalculateOrganizationWirelessRadioAutoRfChannels(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing RecalculateOrganizationWirelessRadioAutoRfChannels",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing RecalculateOrganizationWirelessRadioAutoRfChannels",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseWirelessRecalculateOrganizationWirelessRadioAutoRfChannelsItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsWirelessRadioAutoRfChannelsRecalculateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsWirelessRadioAutoRfChannelsRecalculateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsWirelessRadioAutoRfChannelsRecalculateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsWirelessRadioAutoRfChannelsRecalculate struct {
	OrganizationID types.String                                                         `tfsdk:"organization_id"`
	Item           *ResponseWirelessRecalculateOrganizationWirelessRadioAutoRfChannels  `tfsdk:"item"`
	Parameters     *RequestWirelessRecalculateOrganizationWirelessRadioAutoRfChannelsRs `tfsdk:"parameters"`
}

type ResponseWirelessRecalculateOrganizationWirelessRadioAutoRfChannels struct {
	EstimatedCompletedAt types.String `tfsdk:"estimated_completed_at"`
}

type RequestWirelessRecalculateOrganizationWirelessRadioAutoRfChannelsRs struct {
	NetworkIDs types.Set `tfsdk:"network_ids"`
}

// FromBody
func (r *OrganizationsWirelessRadioAutoRfChannelsRecalculate) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestWirelessRecalculateOrganizationWirelessRadioAutoRfChannels {
	re := *r.Parameters
	var networkIDs []string = nil
	re.NetworkIDs.ElementsAs(ctx, &networkIDs, false)
	out := merakigosdk.RequestWirelessRecalculateOrganizationWirelessRadioAutoRfChannels{
		NetworkIDs: networkIDs,
	}
	return &out
}

// ToBody
func ResponseWirelessRecalculateOrganizationWirelessRadioAutoRfChannelsItemToBody(state OrganizationsWirelessRadioAutoRfChannelsRecalculate, response *merakigosdk.ResponseWirelessRecalculateOrganizationWirelessRadioAutoRfChannels) OrganizationsWirelessRadioAutoRfChannelsRecalculate {
	itemState := ResponseWirelessRecalculateOrganizationWirelessRadioAutoRfChannels{
		EstimatedCompletedAt: types.StringValue(response.EstimatedCompletedAt),
	}
	state.Item = &itemState
	return state
}
