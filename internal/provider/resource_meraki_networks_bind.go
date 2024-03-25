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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksBindResource{}
	_ resource.ResourceWithConfigure = &NetworksBindResource{}
)

func NewNetworksBindResource() resource.Resource {
	return &NetworksBindResource{}
}

type NetworksBindResource struct {
	client *merakigosdk.Client
}

func (r *NetworksBindResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksBindResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_bind"
}

// resourceAction
func (r *NetworksBindResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{

			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"auto_bind": schema.BoolAttribute{
						MarkdownDescription: `Optional boolean indicating whether the network's switches should automatically bind to profiles of the same model. Defaults to false if left unspecified. This option only affects switch networks and switch templates. Auto-bind is not valid unless the switch template has at least one profile and has at most one profile per switch model.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.RequiresReplace(),
						},
					},
					"config_template_id": schema.StringAttribute{
						MarkdownDescription: `The ID of the template to which the network should be bound.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *NetworksBindResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksBind

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
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp1, err := r.client.Networks.BindNetwork(vvNetworkID, dataRequest)

	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing BindNetwork",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing BindNetwork",
			err.Error(),
		)
		return
	}
	//Item

	// data2 := ResponseNetworksBindNetwork(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksBindResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksBindResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksBindResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksBind struct {
	NetworkID  types.String                  `tfsdk:"network_id"`
	Parameters *RequestNetworksBindNetworkRs `tfsdk:"parameters"`
}

type RequestNetworksBindNetworkRs struct {
	AutoBind         types.Bool   `tfsdk:"auto_bind"`
	ConfigTemplateID types.String `tfsdk:"config_template_id"`
}

// FromBody
func (r *NetworksBind) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestNetworksBindNetwork {
	emptyString := ""
	re := *r.Parameters
	autoBind := new(bool)
	if !re.AutoBind.IsUnknown() && !re.AutoBind.IsNull() {
		*autoBind = re.AutoBind.ValueBool()
	} else {
		autoBind = nil
	}
	configTemplateID := new(string)
	if !re.ConfigTemplateID.IsUnknown() && !re.ConfigTemplateID.IsNull() {
		*configTemplateID = re.ConfigTemplateID.ValueString()
	} else {
		configTemplateID = &emptyString
	}
	out := merakigosdk.RequestNetworksBindNetwork{
		AutoBind:         autoBind,
		ConfigTemplateID: *configTemplateID,
	}
	return &out
}

//ToBody
