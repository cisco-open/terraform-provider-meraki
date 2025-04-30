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

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsConfigTemplatesResource{}
	_ resource.ResourceWithConfigure = &OrganizationsConfigTemplatesResource{}
)

func NewOrganizationsConfigTemplatesResource() resource.Resource {
	return &OrganizationsConfigTemplatesResource{}
}

type OrganizationsConfigTemplatesResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsConfigTemplatesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsConfigTemplatesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_config_templates"
}

func (r *OrganizationsConfigTemplatesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"config_template_id": schema.StringAttribute{
				MarkdownDescription: `configTemplateId path parameter. Config template ID`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			// "copy_from_network_id": schema.StringAttribute{
			// 	MarkdownDescription: `The ID of the network or config template to copy configuration from`,
			// 	Computed:            true,
			// 	Optional:            true,
			// 	PlanModifiers: []planmodifier.String{
			// 		stringplanmodifier.UseStateForUnknown(),
			// 		SuppressDiffString(),
			// 	},
			// },
			"id": schema.StringAttribute{
				MarkdownDescription: `The ID of the network or config template to copy configuration from`,
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the configuration template`,
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
			"product_types": schema.SetAttribute{
				MarkdownDescription: `The product types of the configuration template`,
				Computed:            true,
				ElementType:         types.StringType,
			},
			"time_zone": schema.StringAttribute{
				MarkdownDescription: `The timezone of the configuration template. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article</a>. Not applicable if copying from existing network or template`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['configTemplateId']
//path params to assign NOT EDITABLE ['copyFromNetworkId']

func (r *OrganizationsConfigTemplatesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsConfigTemplatesRs

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
	// Has Paths
	vvOrganizationID := data.OrganizationID.ValueString()
	//Has Item and has items and post

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationConfigTemplates(vvOrganizationID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationConfigTemplates",
					restyResp1.String(),
				)
				return
			}
		}
	}

	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvConfigTemplateID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter ConfigTemplateID",
					"Fail Parsing ConfigTemplateID",
				)
				return
			}
			r.client.Organizations.UpdateOrganizationConfigTemplate(vvOrganizationID, vvConfigTemplateID, data.toSdkApiRequestUpdate(ctx))

			responseVerifyItem2, _, _ := r.client.Organizations.GetOrganizationConfigTemplate(vvOrganizationID, vvConfigTemplateID)
			if responseVerifyItem2 != nil {
				data = ResponseOrganizationsGetOrganizationConfigTemplateItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Organizations.CreateOrganizationConfigTemplate(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationConfigTemplate",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationConfigTemplate",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationConfigTemplates(vvOrganizationID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationConfigTemplates",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationConfigTemplates",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvConfigTemplateID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter ConfigTemplateID",
				"Fail Parsing ConfigTemplateID",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Organizations.GetOrganizationConfigTemplate(vvOrganizationID, vvConfigTemplateID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseOrganizationsGetOrganizationConfigTemplateItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationConfigTemplate",
					restyRespGet.String(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationConfigTemplate",
				err.Error(),
			)
			return
		}
	} else {
		resp.Diagnostics.AddError(
			"Error in result.",
			"Error in result.",
		)
		return
	}

}

func (r *OrganizationsConfigTemplatesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsConfigTemplatesRs

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
	vvConfigTemplateID := data.ConfigTemplateID.ValueString()
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationConfigTemplate(vvOrganizationID, vvConfigTemplateID)
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
				"Failure when executing GetOrganizationConfigTemplate",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationConfigTemplate",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationConfigTemplateItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsConfigTemplatesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("config_template_id"), idParts[1])...)
}

func (r *OrganizationsConfigTemplatesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsConfigTemplatesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvConfigTemplateID := data.ConfigTemplateID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationConfigTemplate(vvOrganizationID, vvConfigTemplateID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationConfigTemplate",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationConfigTemplate",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsConfigTemplatesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsConfigTemplatesRs
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
	vvConfigTemplateID := state.ConfigTemplateID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationConfigTemplate(vvOrganizationID, vvConfigTemplateID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationConfigTemplate", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsConfigTemplatesRs struct {
	OrganizationID   types.String `tfsdk:"organization_id"`
	ConfigTemplateID types.String `tfsdk:"config_template_id"`
	ID               types.String `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	ProductTypes     types.Set    `tfsdk:"product_types"`
	TimeZone         types.String `tfsdk:"time_zone"`
}

// FromBody
func (r *OrganizationsConfigTemplatesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationConfigTemplate {
	emptyString := ""
	// copyFromNetworkID := new(string)
	// if !r.CopyFromNetworkID.IsUnknown() && !r.CopyFromNetworkID.IsNull() {
	// 	*copyFromNetworkID = r.CopyFromNetworkID.ValueString()
	// } else {
	// 	copyFromNetworkID = &emptyString
	// }
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	timeZone := new(string)
	if !r.TimeZone.IsUnknown() && !r.TimeZone.IsNull() {
		*timeZone = r.TimeZone.ValueString()
	} else {
		timeZone = &emptyString
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationConfigTemplate{
		// CopyFromNetworkID: *copyFromNetworkID,
		Name:     *name,
		TimeZone: *timeZone,
	}
	return &out
}
func (r *OrganizationsConfigTemplatesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationConfigTemplate {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	timeZone := new(string)
	if !r.TimeZone.IsUnknown() && !r.TimeZone.IsNull() {
		*timeZone = r.TimeZone.ValueString()
	} else {
		timeZone = &emptyString
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationConfigTemplate{
		Name:     *name,
		TimeZone: *timeZone,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationConfigTemplateItemToBodyRs(state OrganizationsConfigTemplatesRs, response *merakigosdk.ResponseOrganizationsGetOrganizationConfigTemplate, is_read bool) OrganizationsConfigTemplatesRs {
	itemState := OrganizationsConfigTemplatesRs{
		ID:           types.StringValue(response.ID),
		Name:         types.StringValue(response.Name),
		ProductTypes: StringSliceToSet(response.ProductTypes),
		TimeZone:     types.StringValue(response.TimeZone),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsConfigTemplatesRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsConfigTemplatesRs)
}
