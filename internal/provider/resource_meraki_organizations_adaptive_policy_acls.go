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

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &OrganizationsAdaptivePolicyACLsResource{}
	_ resource.ResourceWithConfigure = &OrganizationsAdaptivePolicyACLsResource{}
)

func NewOrganizationsAdaptivePolicyACLsResource() resource.Resource {
	return &OrganizationsAdaptivePolicyACLsResource{}
}

type OrganizationsAdaptivePolicyACLsResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsAdaptivePolicyACLsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsAdaptivePolicyACLsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_adaptive_policy_acls"
}

func (r *OrganizationsAdaptivePolicyACLsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"acl_id": schema.StringAttribute{
				MarkdownDescription: `ID of the adaptive policy ACL`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: `When the adaptive policy ACL was created`,
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: `Description of the adaptive policy ACL`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"ip_version": schema.StringAttribute{
				MarkdownDescription: `IP version of adpative policy ACL
                                  Allowed values: [any,ipv4,ipv6]`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"any",
						"ipv4",
						"ipv6",
					),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the adaptive policy ACL`,
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
			"rules": schema.SetNestedAttribute{
				MarkdownDescription: `An ordered array of the adaptive policy ACL rules`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"dst_port": schema.StringAttribute{
							MarkdownDescription: `Destination port`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"log": schema.BoolAttribute{
							MarkdownDescription: `If enabled, when this rule is hit an entry will be logged to the event log
`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"policy": schema.StringAttribute{
							MarkdownDescription: `'allow' or 'deny' traffic specified by this rule
                                        Allowed values: [allow,deny]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"allow",
									"deny",
								),
							},
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: `The type of protocol
                                        Allowed values: [any,icmp,tcp,udp]`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"any",
									"icmp",
									"tcp",
									"udp",
								),
							},
						},
						"src_port": schema.StringAttribute{
							MarkdownDescription: `Source port`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"tcp_established": schema.BoolAttribute{
							MarkdownDescription: `If enabled, means TCP connection with this node must be established.
`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: `When the adaptive policy ACL was last updated`,
				Computed:            true,
			},
		},
	}
}

//path params to set ['aclId']

func (r *OrganizationsAdaptivePolicyACLsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsAdaptivePolicyACLsRs

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

	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationAdaptivePolicyACLs(vvOrganizationID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationAdaptivePolicyACLs",
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
			vvACLID, ok := result2["ACLID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter ACLID",
					"Fail Parsing ACLID",
				)
				return
			}
			r.client.Organizations.UpdateOrganizationAdaptivePolicyACL(vvOrganizationID, vvACLID, data.toSdkApiRequestUpdate(ctx))

			responseVerifyItem2, _, _ := r.client.Organizations.GetOrganizationAdaptivePolicyACL(vvOrganizationID, vvACLID)
			if responseVerifyItem2 != nil {
				data = ResponseOrganizationsGetOrganizationAdaptivePolicyACLItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}

	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Organizations.CreateOrganizationAdaptivePolicyACL(vvOrganizationID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateOrganizationAdaptivePolicyACL",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateOrganizationAdaptivePolicyACL",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationAdaptivePolicyACLs(vvOrganizationID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyACLs",
				restyResp1.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationAdaptivePolicyACLs",
			err.Error(),
		)
		return
	}

	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvACLID, ok := result2["ACLID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter ACLID",
				"Fail Parsing ACLID",
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Organizations.GetOrganizationAdaptivePolicyACL(vvOrganizationID, vvACLID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseOrganizationsGetOrganizationAdaptivePolicyACLItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetOrganizationAdaptivePolicyACL",
					restyRespGet.String(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAdaptivePolicyACL",
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

func (r *OrganizationsAdaptivePolicyACLsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsAdaptivePolicyACLsRs

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
	vvACLID := data.ACLID.ValueString()
	responseGet, restyRespGet, err := r.client.Organizations.GetOrganizationAdaptivePolicyACL(vvOrganizationID, vvACLID)
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
				"Failure when executing GetOrganizationAdaptivePolicyACL",
				restyRespGet.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationAdaptivePolicyACL",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseOrganizationsGetOrganizationAdaptivePolicyACLItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *OrganizationsAdaptivePolicyACLsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("acl_id"), idParts[1])...)
}

func (r *OrganizationsAdaptivePolicyACLsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsAdaptivePolicyACLsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	vvACLID := data.ACLID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Organizations.UpdateOrganizationAdaptivePolicyACL(vvOrganizationID, vvACLID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationAdaptivePolicyACL",
				restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationAdaptivePolicyACL",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsAdaptivePolicyACLsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsAdaptivePolicyACLsRs
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
	vvACLID := state.ACLID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationAdaptivePolicyACL(vvOrganizationID, vvACLID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationAdaptivePolicyACL", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsAdaptivePolicyACLsRs struct {
	OrganizationID types.String                                                    `tfsdk:"organization_id"`
	ACLID          types.String                                                    `tfsdk:"acl_id"`
	CreatedAt      types.String                                                    `tfsdk:"created_at"`
	Description    types.String                                                    `tfsdk:"description"`
	IPVersion      types.String                                                    `tfsdk:"ip_version"`
	Name           types.String                                                    `tfsdk:"name"`
	Rules          *[]ResponseOrganizationsGetOrganizationAdaptivePolicyAclRulesRs `tfsdk:"rules"`
	UpdatedAt      types.String                                                    `tfsdk:"updated_at"`
}

type ResponseOrganizationsGetOrganizationAdaptivePolicyAclRulesRs struct {
	DstPort        types.String `tfsdk:"dst_port"`
	Log            types.Bool   `tfsdk:"log"`
	Policy         types.String `tfsdk:"policy"`
	Protocol       types.String `tfsdk:"protocol"`
	SrcPort        types.String `tfsdk:"src_port"`
	TCPEstablished types.Bool   `tfsdk:"tcp_established"`
}

// FromBody
func (r *OrganizationsAdaptivePolicyACLsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyACL {
	emptyString := ""
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = &emptyString
	}
	iPVersion := new(string)
	if !r.IPVersion.IsUnknown() && !r.IPVersion.IsNull() {
		*iPVersion = r.IPVersion.ValueString()
	} else {
		iPVersion = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestOrganizationsCreateOrganizationAdaptivePolicyACLRules []merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyACLRules

	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			dstPort := rItem1.DstPort.ValueString()
			log := func() *bool {
				if !rItem1.Log.IsUnknown() && !rItem1.Log.IsNull() {
					return rItem1.Log.ValueBoolPointer()
				}
				return nil
			}()
			policy := rItem1.Policy.ValueString()
			protocol := rItem1.Protocol.ValueString()
			srcPort := rItem1.SrcPort.ValueString()
			tcpEstablished := func() *bool {
				if !rItem1.TCPEstablished.IsUnknown() && !rItem1.TCPEstablished.IsNull() {
					return rItem1.TCPEstablished.ValueBoolPointer()
				}
				return nil
			}()
			requestOrganizationsCreateOrganizationAdaptivePolicyACLRules = append(requestOrganizationsCreateOrganizationAdaptivePolicyACLRules, merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyACLRules{
				DstPort:        dstPort,
				Log:            log,
				Policy:         policy,
				Protocol:       protocol,
				SrcPort:        srcPort,
				TCPEstablished: tcpEstablished,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyACL{
		Description: *description,
		IPVersion:   *iPVersion,
		Name:        *name,
		Rules: func() *[]merakigosdk.RequestOrganizationsCreateOrganizationAdaptivePolicyACLRules {
			if len(requestOrganizationsCreateOrganizationAdaptivePolicyACLRules) > 0 {
				return &requestOrganizationsCreateOrganizationAdaptivePolicyACLRules
			}
			return nil
		}(),
	}
	return &out
}
func (r *OrganizationsAdaptivePolicyACLsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyACL {
	emptyString := ""
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = &emptyString
	}
	iPVersion := new(string)
	if !r.IPVersion.IsUnknown() && !r.IPVersion.IsNull() {
		*iPVersion = r.IPVersion.ValueString()
	} else {
		iPVersion = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestOrganizationsUpdateOrganizationAdaptivePolicyACLRules []merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyACLRules

	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			dstPort := rItem1.DstPort.ValueString()
			log := func() *bool {
				if !rItem1.Log.IsUnknown() && !rItem1.Log.IsNull() {
					return rItem1.Log.ValueBoolPointer()
				}
				return nil
			}()
			policy := rItem1.Policy.ValueString()
			protocol := rItem1.Protocol.ValueString()
			srcPort := rItem1.SrcPort.ValueString()
			tcpEstablished := func() *bool {
				if !rItem1.TCPEstablished.IsUnknown() && !rItem1.TCPEstablished.IsNull() {
					return rItem1.TCPEstablished.ValueBoolPointer()
				}
				return nil
			}()
			requestOrganizationsUpdateOrganizationAdaptivePolicyACLRules = append(requestOrganizationsUpdateOrganizationAdaptivePolicyACLRules, merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyACLRules{
				DstPort:        dstPort,
				Log:            log,
				Policy:         policy,
				Protocol:       protocol,
				SrcPort:        srcPort,
				TCPEstablished: tcpEstablished,
			})
			//[debug] Is Array: True
		}
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyACL{
		Description: *description,
		IPVersion:   *iPVersion,
		Name:        *name,
		Rules: func() *[]merakigosdk.RequestOrganizationsUpdateOrganizationAdaptivePolicyACLRules {
			if len(requestOrganizationsUpdateOrganizationAdaptivePolicyACLRules) > 0 {
				return &requestOrganizationsUpdateOrganizationAdaptivePolicyACLRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationAdaptivePolicyACLItemToBodyRs(state OrganizationsAdaptivePolicyACLsRs, response *merakigosdk.ResponseOrganizationsGetOrganizationAdaptivePolicyACL, is_read bool) OrganizationsAdaptivePolicyACLsRs {
	itemState := OrganizationsAdaptivePolicyACLsRs{
		ACLID:       types.StringValue(response.ACLID),
		CreatedAt:   types.StringValue(response.CreatedAt),
		Description: types.StringValue(response.Description),
		IPVersion:   types.StringValue(response.IPVersion),
		Name:        types.StringValue(response.Name),
		Rules: func() *[]ResponseOrganizationsGetOrganizationAdaptivePolicyAclRulesRs {
			if response.Rules != nil {
				result := make([]ResponseOrganizationsGetOrganizationAdaptivePolicyAclRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseOrganizationsGetOrganizationAdaptivePolicyAclRulesRs{
						DstPort: types.StringValue(rules.DstPort),
						Log: func() types.Bool {
							if rules.Log != nil {
								return types.BoolValue(*rules.Log)
							}
							return types.Bool{}
						}(),
						Policy:   types.StringValue(rules.Policy),
						Protocol: types.StringValue(rules.Protocol),
						SrcPort:  types.StringValue(rules.SrcPort),
						TCPEstablished: func() types.Bool {
							if rules.TCPEstablished != nil {
								return types.BoolValue(*rules.TCPEstablished)
							}
							return types.Bool{}
						}(),
					}
				}
				return &result
			}
			return nil
		}(),
		UpdatedAt: types.StringValue(response.UpdatedAt),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(OrganizationsAdaptivePolicyACLsRs)
	}
	return mergeInterfaces(state, itemState, true).(OrganizationsAdaptivePolicyACLsRs)
}
