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
	_ resource.Resource              = &OrganizationsInventoryReleaseResource{}
	_ resource.ResourceWithConfigure = &OrganizationsInventoryReleaseResource{}
)

func NewOrganizationsInventoryReleaseResource() resource.Resource {
	return &OrganizationsInventoryReleaseResource{}
}

type OrganizationsInventoryReleaseResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsInventoryReleaseResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsInventoryReleaseResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_inventory_release"
}

// resourceAction
func (r *OrganizationsInventoryReleaseResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					"serials": schema.SetAttribute{
						MarkdownDescription: `Serials of the devices that should be released`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
				},
			},
		},
	}
}
func (r *OrganizationsInventoryReleaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsInventoryRelease

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
	restyResp1, err := r.client.Organizations.ReleaseFromOrganizationInventory(vvOrganizationID, dataRequest)

	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ReleaseFromOrganizationInventory",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ReleaseFromOrganizationInventory",
			err.Error(),
		)
		return
	}
	//Item

	// data2 := ResponseOrganizationsReleaseFromOrganizationInventory(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsInventoryReleaseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsInventoryReleaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsInventoryReleaseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsInventoryRelease struct {
	OrganizationID types.String                                            `tfsdk:"organization_id"`
	Parameters     *RequestOrganizationsReleaseFromOrganizationInventoryRs `tfsdk:"parameters"`
}

type RequestOrganizationsReleaseFromOrganizationInventoryRs struct {
	Serials types.Set `tfsdk:"serials"`
}

// FromBody
func (r *OrganizationsInventoryRelease) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsReleaseFromOrganizationInventory {
	re := *r.Parameters
	var serials []string = nil
	re.Serials.ElementsAs(ctx, &serials, false)
	out := merakigosdk.RequestOrganizationsReleaseFromOrganizationInventory{
		Serials: serials,
	}
	return &out
}

//ToBody
