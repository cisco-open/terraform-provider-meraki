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
	_ resource.Resource              = &OrganizationsDevicesControllerMigrationsCreateResource{}
	_ resource.ResourceWithConfigure = &OrganizationsDevicesControllerMigrationsCreateResource{}
)

func NewOrganizationsDevicesControllerMigrationsCreateResource() resource.Resource {
	return &OrganizationsDevicesControllerMigrationsCreateResource{}
}

type OrganizationsDevicesControllerMigrationsCreateResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsDevicesControllerMigrationsCreateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsDevicesControllerMigrationsCreateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_devices_controller_migrations_create"
}

// resourceAction
func (r *OrganizationsDevicesControllerMigrationsCreateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
					"items": schema.ListNestedAttribute{
						Computed: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"items": schema.ListNestedAttribute{
									MarkdownDescription: `Array of ResponseOrganizationsCreateOrganizationDevicesControllerMigration`,
									Computed:            true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{

											"created_at": schema.StringAttribute{
												MarkdownDescription: `The time at which a migration was created`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"migrated_at": schema.StringAttribute{
												MarkdownDescription: `The time at which the device initiated migration`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"serial": schema.StringAttribute{
												MarkdownDescription: `The device serial`,
												Computed:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
											"target": schema.StringAttribute{
												MarkdownDescription: `The migration target destination
                                                    Allowed values: [wirelessController]`,
												Computed: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplace(),
												},
											},
										},
									},
								},
							},
						},
					},
					"serials": schema.ListAttribute{
						MarkdownDescription: `A list of Meraki Serials to migrate`,
						Optional:            true,
						Computed:            true,
						ElementType:         types.StringType,
					},
					"target": schema.StringAttribute{
						MarkdownDescription: `The controller or management mode to which the devices will be migrated
                                        Allowed values: [wirelessController]`,
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *OrganizationsDevicesControllerMigrationsCreateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsDevicesControllerMigrationsCreate

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
	response, restyResp1, err := r.client.Organizations.CreateOrganizationDevicesControllerMigration(vvOrganizationID, dataRequest)
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationDevicesControllerMigration",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationDevicesControllerMigration",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseOrganizationsCreateOrganizationDevicesControllerMigrationItemsToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsDevicesControllerMigrationsCreateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsDevicesControllerMigrationsCreateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *OrganizationsDevicesControllerMigrationsCreateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type OrganizationsDevicesControllerMigrationsCreate struct {
	OrganizationID types.String                                                             `tfsdk:"organization_id"`
	Items          *[]ResponseItemOrganizationsCreateOrganizationDevicesControllerMigration `tfsdk:"items"`
	Parameters     *RequestOrganizationsCreateOrganizationDevicesControllerMigrationRs      `tfsdk:"parameters"`
}

type ResponseItemOrganizationsCreateOrganizationDevicesControllerMigration struct {
	CreatedAt  types.String `tfsdk:"created_at"`
	MigratedAt types.String `tfsdk:"migrated_at"`
	Serial     types.String `tfsdk:"serial"`
	Target     types.String `tfsdk:"target"`
}

type RequestOrganizationsCreateOrganizationDevicesControllerMigrationRs struct {
	Serials types.Set    `tfsdk:"serials"`
	Target  types.String `tfsdk:"target"`
}

// FromBody
func (r *OrganizationsDevicesControllerMigrationsCreate) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationDevicesControllerMigration {
	emptyString := ""
	re := *r.Parameters
	var serials []string = nil
	re.Serials.ElementsAs(ctx, &serials, false)
	target := new(string)
	if !re.Target.IsUnknown() && !re.Target.IsNull() {
		*target = re.Target.ValueString()
	} else {
		target = &emptyString
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationDevicesControllerMigration{
		Serials: serials,
		Target:  *target,
	}
	return &out
}

// ToBody
func ResponseOrganizationsCreateOrganizationDevicesControllerMigrationItemsToBody(state OrganizationsDevicesControllerMigrationsCreate, response *merakigosdk.ResponseOrganizationsCreateOrganizationDevicesControllerMigration) OrganizationsDevicesControllerMigrationsCreate {
	var items []ResponseItemOrganizationsCreateOrganizationDevicesControllerMigration
	for _, item := range *response {
		itemState := ResponseItemOrganizationsCreateOrganizationDevicesControllerMigration{
			CreatedAt:  types.StringValue(item.CreatedAt),
			MigratedAt: types.StringValue(item.MigratedAt),
			Serial:     types.StringValue(item.Serial),
			Target:     types.StringValue(item.Target),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}
