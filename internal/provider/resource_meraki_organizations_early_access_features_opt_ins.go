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
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v2/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsEarlyAccessFeaturesOptInsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsEarlyAccessFeaturesOptInsResource{}
)

func NewOrganizationsEarlyAccessFeaturesOptInsResource() resource.Resource {
	return &OrganizationsEarlyAccessFeaturesOptInsResource{}
}

type OrganizationsEarlyAccessFeaturesOptInsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsEarlyAccessFeaturesOptInsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsEarlyAccessFeaturesOptInsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_early_access_features_opt_ins"
}

func (r *OrganizationsEarlyAccessFeaturesOptInsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"limit_scope_to_networks": schema.SetAttribute{
				MarkdownDescription: `A list of network IDs to apply the opt-in to`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"opt_in_id": schema.StringAttribute{
				MarkdownDescription: `optInId path parameter. Opt in ID`,
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
			"short_name": schema.StringAttribute{
				MarkdownDescription: `Short name of the early access feature`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
			},
		},
	}
}

//path params to set ['optInId']
//path params to assign NOT EDITABLE ['shortName']

func (r *OrganizationsEarlyAccessFeaturesOptInsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsEarlyAccessFeaturesOptInsRs

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
	//Reviw This  Has Item and item
	//HAS CREATE

	vvOptInID := data.OptInID.ValueString()
	if vvOptInID != "" {
		responseVerifyItem, restyRespGet, err := r.client.Organizations.GetOrganizationEarlyAccessFeaturesOptIn(vvOrganizationID, vvOptInID)
		if err != nil || responseVerifyItem == nil {
			if restyRespGet != nil {
				if restyRespGet.StatusCode() != 404 {

					resp.Diagnostics.AddError(
						"Failure when executing GetOrganizationEarlyAccessFeaturesOptIn",
						err.Error(),
					)
					return
				}
			}
		}

		if responseVerifyItem != nil {
			data = ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInItemToBodyRs(data, responseVerifyItem, false)
			diags := resp.State.Set(ctx, &data)
			resp.Diagnostics.Append(diags...)
			return
		}
	}

	response, restyResp2, err := r.client.Organizations.CreateOrganizationEarlyAccessFeaturesOptIn(vvOrganizationID, data.toSdkApiRequestCreate(ctx))

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ",
			err.Error(),
		)
		return
	}
	//Items
	vvOptInID = response.ID
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationEarlyAccessFeaturesOptIn(vvOrganizationID, vvOptInID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationEarlyAccessFeaturesOptIns",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationEarlyAccessFeaturesOptIns",
			err.Error(),
		)
		return
	} else {
		data = ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInItemToBodyRs(data, responseGet, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	}
}

func (r *OrganizationsEarlyAccessFeaturesOptInsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsEarlyAccessFeaturesOptInsRs

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
	// Has Item2

	vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	vvOptInID := data.OptInID.ValueString()
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationEarlyAccessFeaturesOptIn(vvOrganizationID, vvOptInID)
	if err != nil || restyRespGet == nil {
		if restyRespGet != nil {
			if restyRespGet.StatusCode() == 404 {
				resp.Diagnostics.AddWarning(
					"Resource not found",
					"Deleting resource",
				)
				resp.State.RemoveResource(ctx)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationEarlyAccessFeaturesOptIn",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationEarlyAccessFeaturesOptIn",
			err.Error(),
		)
		return
	}

	data = ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsEarlyAccessFeaturesOptInsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("opt_in_id"), idParts[1])...)
}

func (r *OrganizationsEarlyAccessFeaturesOptInsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsEarlyAccessFeaturesOptInsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	vvOptInID := data.OptInID.ValueString()
	// opt_in_id
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Organizations.UpdateOrganizationEarlyAccessFeaturesOptIn(vvOrganizationID, vvOptInID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationEarlyAccessFeaturesOptIn",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationEarlyAccessFeaturesOptIn",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsEarlyAccessFeaturesOptInsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsEarlyAccessFeaturesOptInsRs
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &state, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	vvOrganizationID := state.OrganizationID.ValueString()
	vvOptInID := state.OptInID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationEarlyAccessFeaturesOptIn(vvOrganizationID, vvOptInID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationEarlyAccessFeaturesOptIn", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsEarlyAccessFeaturesOptInsRs struct {
	OrganizationID         types.String                                                                          `tfsdk:"organization_id"`
	OptInID                types.String                                                                          `tfsdk:"opt_in_id"`
	CreatedAt              types.String                                                                          `tfsdk:"created_at"`
	ID                     types.String                                                                          `tfsdk:"id"`
	LimitScopeToNetworks   *[]ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworksRs `tfsdk:"limit_scope_to_networks"`
	LimitScopeToNetworksRs types.Set                                                                             `tfsdk:"limit_scope_to_networks_rs"`
	ShortName              types.String                                                                          `tfsdk:"short_name"`
}

type ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworksRs struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// FromBody
func (r *OrganizationsEarlyAccessFeaturesOptInsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationEarlyAccessFeaturesOptIn {
	emptyString := ""
	var limitScopeToNetworks []string = nil
	r.LimitScopeToNetworksRs.ElementsAs(ctx, &limitScopeToNetworks, false)
	shortName := new(string)
	if !r.ShortName.IsUnknown() && !r.ShortName.IsNull() {
		*shortName = r.ShortName.ValueString()
	} else {
		shortName = &emptyString
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationEarlyAccessFeaturesOptIn{
		LimitScopeToNetworks: limitScopeToNetworks,
		ShortName:            *shortName,
	}
	return &out
}
func (r *OrganizationsEarlyAccessFeaturesOptInsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationEarlyAccessFeaturesOptIn {
	var limitScopeToNetworks []string = nil
	r.LimitScopeToNetworksRs.ElementsAs(ctx, &limitScopeToNetworks, false)
	out := merakigosdk.RequestOrganizationsUpdateOrganizationEarlyAccessFeaturesOptIn{
		LimitScopeToNetworks: limitScopeToNetworks,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInItemToBodyRs(state OrganizationsEarlyAccessFeaturesOptInsRs, response *merakigosdk.ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptIn, is_read bool) OrganizationsEarlyAccessFeaturesOptInsRs {
	itemState := OrganizationsEarlyAccessFeaturesOptInsRs{
		CreatedAt: types.StringValue(response.CreatedAt),
		ID:        types.StringValue(response.ID),
		LimitScopeToNetworks: func() *[]ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworksRs {
			if response.LimitScopeToNetworks != nil {
				result := make([]ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworksRs, len(*response.LimitScopeToNetworks))
				for i, limitScopeToNetworks := range *response.LimitScopeToNetworks {
					result[i] = ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworksRs{
						ID:   types.StringValue(limitScopeToNetworks.ID),
						Name: types.StringValue(limitScopeToNetworks.Name),
					}
				}
				return &result
			}
			return &[]ResponseOrganizationsGetOrganizationEarlyAccessFeaturesOptInLimitScopeToNetworksRs{}
		}(),
		ShortName: types.StringValue(response.ShortName),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsEarlyAccessFeaturesOptInsRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsEarlyAccessFeaturesOptInsRs)
}
