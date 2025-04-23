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
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &AdministeredIDentitiesMeAPIKeysGenerateResource{}
	_ resource.ResourceWithConfigure = &AdministeredIDentitiesMeAPIKeysGenerateResource{}
)

func NewAdministeredIDentitiesMeAPIKeysGenerateResource() resource.Resource {
	return &AdministeredIDentitiesMeAPIKeysGenerateResource{}
}

type AdministeredIDentitiesMeAPIKeysGenerateResource struct {
	client *merakigosdk.Client
}

func (r *AdministeredIDentitiesMeAPIKeysGenerateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *AdministeredIDentitiesMeAPIKeysGenerateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_administered_identities_me_api_keys_generate"
}

// resourceAction
func (r *AdministeredIDentitiesMeAPIKeysGenerateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"key": schema.StringAttribute{
						MarkdownDescription: `API key in plaintext. This value will not be accessible outside of key generation`,
						Computed:            true,
					},
				},
			},
		},
	}
}
func (r *AdministeredIDentitiesMeAPIKeysGenerateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data AdministeredIDentitiesMeAPIKeysGenerate

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
	response, restyResp1, err := r.client.Administered.GenerateAdministeredIDentitiesMeAPIKeys()
	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GenerateAdministeredIDentitiesMeAPIKeys",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GenerateAdministeredIDentitiesMeAPIKeys",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseAdministeredGenerateAdministeredIDentitiesMeAPIKeysItemToBody(data, response)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *AdministeredIDentitiesMeAPIKeysGenerateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *AdministeredIDentitiesMeAPIKeysGenerateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *AdministeredIDentitiesMeAPIKeysGenerateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type AdministeredIDentitiesMeAPIKeysGenerate struct {
	Item *ResponseAdministeredGenerateAdministeredIdentitiesMeApiKeys `tfsdk:"item"`
}

type ResponseAdministeredGenerateAdministeredIdentitiesMeApiKeys struct {
	Key types.String `tfsdk:"key"`
}

type RequestAdministeredGenerateAdministeredIdentitiesMeApiKeysRs interface{}

// FromBody
// ToBody
func ResponseAdministeredGenerateAdministeredIDentitiesMeAPIKeysItemToBody(state AdministeredIDentitiesMeAPIKeysGenerate, response *merakigosdk.ResponseAdministeredGenerateAdministeredIDentitiesMeAPIKeys) AdministeredIDentitiesMeAPIKeysGenerate {
	itemState := ResponseAdministeredGenerateAdministeredIdentitiesMeApiKeys{
		Key: types.StringValue(response.Key),
	}
	state.Item = &itemState
	return state
}
