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
	"strconv"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsPolicyObjectsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsPolicyObjectsResource{}
)

func NewOrganizationsPolicyObjectsResource() resource.Resource {
	return &OrganizationsPolicyObjectsResource{}
}

type OrganizationsPolicyObjectsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsPolicyObjectsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsPolicyObjectsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_policy_objects"
}

func (r *OrganizationsPolicyObjectsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"category": schema.StringAttribute{
				MarkdownDescription: `Category of a policy object (one of: adaptivePolicy, network)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
			},
			"cidr": schema.StringAttribute{
				MarkdownDescription: `CIDR Value of a policy object`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: `Time Stamp of policy object creation.`,
				Computed:            true,
			},
			"fqdn": schema.StringAttribute{
				MarkdownDescription: `Fully qualified domain name of policy object (e.g. "example.com")`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"group_ids": schema.SetAttribute{
				MarkdownDescription: `The IDs of policy object groups the policy object belongs to.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
				Default:     setdefault.StaticValue(types.SetNull(types.StringType)),
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `Policy object ID`,
				Computed:            true,
			},
			"ip": schema.StringAttribute{
				MarkdownDescription: `IP Address of a policy object (e.g. "1.2.3.4")`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"mask": schema.StringAttribute{
				MarkdownDescription: `Mask of a policy object (e.g. "255.255.0.0")`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of policy object (alphanumeric, space, dash, or underscore characters only).`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_ids": schema.SetAttribute{
				MarkdownDescription: `The IDs of the networks that use the policy object.`,
				Computed:            true,
				ElementType:         types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"policy_object_id": schema.StringAttribute{
				MarkdownDescription: `policyObjectId path parameter. Policy object ID`,
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: `Type of a policy object (one of: adaptivePolicyIpv4Cidr, cidr, fqdn, ipAndMask)`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: `Time Stamp of policy object updation.`,
				Computed:            true,
			},
		},
	}
}

//path params to assign NOT EDITABLE ['category', 'type']

func (r *OrganizationsPolicyObjectsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsPolicyObjectsRs

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
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationPolicyObjects(vvOrganizationID, &merakigosdk.GetOrganizationPolicyObjectsQueryParams{
		PerPage: -1,
	})
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationPolicyObjects",
				err.Error(),
			)
			return
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvPolicyObjectID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter PolicyObjectID",
					"Error",
				)
				return
			}
			r.client.Organizations.UpdateOrganizationPolicyObject(vvOrganizationID, vvPolicyObjectID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Organizations.GetOrganizationPolicyObject(vvOrganizationID, vvPolicyObjectID)
			if responseVerifyItem2 != nil {
				data = ResponseOrganizationsGetOrganizationPolicyObjectItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	_, restyResp2, err := r.client.Organizations.CreateOrganizationPolicyObject(vvOrganizationID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationPolicyObject",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationPolicyObject",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationPolicyObjects(vvOrganizationID, &merakigosdk.GetOrganizationPolicyObjectsQueryParams{
		PerPage: -1,
	})
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationPolicyObjects",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationPolicyObjects",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvPolicyObjectID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter PolicyObjectID",
				"Error",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Organizations.GetOrganizationPolicyObject(vvOrganizationID, vvPolicyObjectID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseOrganizationsGetOrganizationPolicyObjectItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationPolicyObject",
					restyRespGet.String(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationPolicyObject",
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

func (r *OrganizationsPolicyObjectsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsPolicyObjectsRs

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
	vvPolicyObjectID := data.PolicyObjectID.ValueString()
	// policy_object_id
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationPolicyObject(vvOrganizationID, vvPolicyObjectID)
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
				"Failure when executing GetOrganizationPolicyObjects",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationPolicyObjects",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationPolicyObjectItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsPolicyObjectsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("policy_object_id"), idParts[1])...)
}

func (r *OrganizationsPolicyObjectsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsPolicyObjectsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvPolicyObjectID := data.PolicyObjectID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationPolicyObject(vvOrganizationID, vvPolicyObjectID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationPolicyObject",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationPolicyObject",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsPolicyObjectsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsPolicyObjectsRs
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
	vvPolicyObjectID := state.PolicyObjectID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationPolicyObject(vvOrganizationID, vvPolicyObjectID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationPolicyObject", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsPolicyObjectsRs struct {
	OrganizationID types.String `tfsdk:"organization_id"`
	PolicyObjectID types.String `tfsdk:"policy_object_id"`
	Category       types.String `tfsdk:"category"`
	Cidr           types.String `tfsdk:"cidr"`
	CreatedAt      types.String `tfsdk:"created_at"`
	GroupIDs       types.Set    `tfsdk:"group_ids"`
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	NetworkIDs     types.Set    `tfsdk:"network_ids"`
	Type           types.String `tfsdk:"type"`
	UpdatedAt      types.String `tfsdk:"updated_at"`
	Fqdn           types.String `tfsdk:"fqdn"`
	IP             types.String `tfsdk:"ip"`
	Mask           types.String `tfsdk:"mask"`
}

// FromBody
func (r *OrganizationsPolicyObjectsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationPolicyObject {
	emptyString := ""
	category := new(string)
	if !r.Category.IsUnknown() && !r.Category.IsNull() {
		*category = r.Category.ValueString()
	} else {
		category = &emptyString
	}
	cidr := new(string)
	if !r.Cidr.IsUnknown() && !r.Cidr.IsNull() {
		*cidr = r.Cidr.ValueString()
	} else {
		cidr = &emptyString
	}
	fqdn := new(string)
	if !r.Fqdn.IsUnknown() && !r.Fqdn.IsNull() {
		*fqdn = r.Fqdn.ValueString()
	} else {
		fqdn = &emptyString
	}
	var groupIDs []string = nil
	r.GroupIDs.ElementsAs(ctx, &groupIDs, false)
	iP := new(string)
	if !r.IP.IsUnknown() && !r.IP.IsNull() {
		*iP = r.IP.ValueString()
	} else {
		iP = &emptyString
	}
	mask := new(string)
	if !r.Mask.IsUnknown() && !r.Mask.IsNull() {
		*mask = r.Mask.ValueString()
	} else {
		mask = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	typeR := new(string)
	if !r.Type.IsUnknown() && !r.Type.IsNull() {
		*typeR = r.Type.ValueString()
	} else {
		typeR = &emptyString
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationPolicyObject{
		Category: *category,
		Cidr:     *cidr,
		Fqdn:     *fqdn,
		GroupIDs: groupIDs,
		IP:       *iP,
		Mask:     *mask,
		Name:     *name,
		Type:     *typeR,
	}
	return &out
}
func (r *OrganizationsPolicyObjectsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationPolicyObject {
	emptyString := ""
	cidr := new(string)
	if !r.Cidr.IsUnknown() && !r.Cidr.IsNull() {
		*cidr = r.Cidr.ValueString()
	} else {
		cidr = &emptyString
	}
	fqdn := new(string)
	if !r.Fqdn.IsUnknown() && !r.Fqdn.IsNull() {
		*fqdn = r.Fqdn.ValueString()
	} else {
		fqdn = &emptyString
	}
	var groupIDs []string = nil
	r.GroupIDs.ElementsAs(ctx, &groupIDs, false)
	iP := new(string)
	if !r.IP.IsUnknown() && !r.IP.IsNull() {
		*iP = r.IP.ValueString()
	} else {
		iP = &emptyString
	}
	mask := new(string)
	if !r.Mask.IsUnknown() && !r.Mask.IsNull() {
		*mask = r.Mask.ValueString()
	} else {
		mask = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationPolicyObject{
		Cidr:     *cidr,
		Fqdn:     *fqdn,
		GroupIDs: groupIDs,
		IP:       *iP,
		Mask:     *mask,
		Name:     *name,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationPolicyObjectItemToBodyRs(state OrganizationsPolicyObjectsRs, response *merakigosdk.ResponseOrganizationsGetOrganizationPolicyObject, is_read bool) OrganizationsPolicyObjectsRs {
	itemState := OrganizationsPolicyObjectsRs{
		Category: func() types.String {
			if response.Category != "" {
				return types.StringValue(response.Category)
			}
			return types.String{}
		}(),
		Cidr: func() types.String {
			if response.Cidr != "" {
				return types.StringValue(response.Cidr)
			}
			return types.String{}
		}(),
		CreatedAt: func() types.String {
			if response.CreatedAt != "" {
				return types.StringValue(response.CreatedAt)
			}
			return types.String{}
		}(),
		GroupIDs: StringSliceToSet(response.GroupIDs),
		ID: func() types.String {
			if response.ID != "" {
				return types.StringValue(response.ID)
			}
			return types.String{}
		}(),
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
		NetworkIDs: StringSliceToSet(response.NetworkIDs),
		Type: func() types.String {
			if response.Type != "" {
				return types.StringValue(response.Type)
			}
			return types.String{}
		}(),
		UpdatedAt: func() types.String {
			if response.UpdatedAt != "" {
				return types.StringValue(response.UpdatedAt)
			}
			return types.String{}
		}(),
		PolicyObjectID: func() types.String {
			if response.ID != "" {
				return types.StringValue(response.ID)
			}
			return types.String{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsPolicyObjectsRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsPolicyObjectsRs)
}
