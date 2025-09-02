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
	_ resource.Resource              = &OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDeleteResource{}
	_ resource.ResourceWithConfigure = &OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDeleteResource{}
)

func NewOrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDeleteResource() resource.Resource {
	return &OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDeleteResource{}
}

type OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDeleteResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDeleteResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDeleteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_wireless_ssids_firewall_isolation_allowlist_entries_delete"
}

// resourceAction
func (r *OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDeleteResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"entry_id": schema.StringAttribute{
				MarkdownDescription: `entryId path parameter. Entry ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}
func (r *OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDeleteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDelete

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
	vvEntryID := data.EntryID.ValueString()
	restyResp1, err := r.client.Wireless.DeleteOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry(vvOrganizationID, vvEntryID)
	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing DeleteOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry",
			err.Error(),
		)
		return
	}
	//Item
	// //entro aqui 2
	// data2 := ResponseWirelessDeleteOrganizationWirelessSSIDsFirewallIsolationAllowlistEntry(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDeleteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDeleteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDeleteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsWirelessSSIDsFirewallIsolationAllowlistEntriesDelete struct {
	OrganizationID types.String `tfsdk:"organization_id"`
	EntryID        types.String `tfsdk:"entry_id"`
}

//FromBody
