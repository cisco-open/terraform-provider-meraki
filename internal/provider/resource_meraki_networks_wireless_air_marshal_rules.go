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

	merakigosdk "github.com/meraki/dashboard-api-go/v4/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessAirMarshalRulesResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessAirMarshalRulesResource{}
)

func NewNetworksWirelessAirMarshalRulesResource() resource.Resource {
	return &NetworksWirelessAirMarshalRulesResource{}
}

type NetworksWirelessAirMarshalRulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessAirMarshalRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessAirMarshalRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_air_marshal_rules"
}

// resourceAction
func (r *NetworksWirelessAirMarshalRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"created_at": schema.StringAttribute{
						MarkdownDescription: `Created at timestamp`,
						Computed:            true,
					},
					"match": schema.SingleNestedAttribute{
						MarkdownDescription: `Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default)`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"string": schema.StringAttribute{
								MarkdownDescription: `Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default)`,
								Computed:            true,
							},
							"type": schema.StringAttribute{
								MarkdownDescription: `Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default)`,
								Computed:            true,
							},
						},
					},
					"network": schema.SingleNestedAttribute{
						MarkdownDescription: `Network details`,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"id": schema.StringAttribute{
								MarkdownDescription: `Network ID`,
								Computed:            true,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Network name`,
								Computed:            true,
							},
						},
					},
					"rule_id": schema.StringAttribute{
						MarkdownDescription: `Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default)`,
						Computed:            true,
					},
					"type": schema.StringAttribute{
						MarkdownDescription: `Indicates whether or not clients are allowed to       connect to rogue SSIDs by default. (blocked by default)`,
						Computed:            true,
					},
					"updated_at": schema.StringAttribute{
						MarkdownDescription: `Updated at timestamp`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"match": schema.SingleNestedAttribute{
						MarkdownDescription: `Object describing the rule specification.`,
						Optional:            true,
						Computed:            true,
						Attributes: map[string]schema.Attribute{

							"string": schema.StringAttribute{
								MarkdownDescription: `The string used to match.`,
								Optional:            true,
								Computed:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
							"type": schema.StringAttribute{
								MarkdownDescription: `The type of match.
                                              Allowed values: [bssid,contains,exact,wildcard]`,
								Optional: true,
								Computed: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.RequiresReplace(),
								},
							},
						},
					},
					"type": schema.StringAttribute{
						MarkdownDescription: `Indicates if this rule will allow, block, or alert.
                                        Allowed values: [alert,allow,block]`,
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
func (r *NetworksWirelessAirMarshalRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessAirMarshalRules

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Wireless.CreateNetworkWirelessAirMarshalRule(vvNetworkID, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkWirelessAirMarshalRule",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkWirelessAirMarshalRule",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseWirelessCreateNetworkWirelessAirMarshalRuleItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessAirMarshalRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksWirelessAirMarshalRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksWirelessAirMarshalRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessAirMarshalRules struct {
	NetworkID  types.String                                          `tfsdk:"network_id"`
	Item       *ResponseWirelessCreateNetworkWirelessAirMarshalRule  `tfsdk:"item"`
	Parameters *RequestWirelessCreateNetworkWirelessAirMarshalRuleRs `tfsdk:"parameters"`
}

type ResponseWirelessCreateNetworkWirelessAirMarshalRule struct {
	CreatedAt types.String                                                `tfsdk:"created_at"`
	Match     *ResponseWirelessCreateNetworkWirelessAirMarshalRuleMatch   `tfsdk:"match"`
	Network   *ResponseWirelessCreateNetworkWirelessAirMarshalRuleNetwork `tfsdk:"network"`
	RuleID    types.String                                                `tfsdk:"rule_id"`
	Type      types.String                                                `tfsdk:"type"`
	UpdatedAt types.String                                                `tfsdk:"updated_at"`
}

type ResponseWirelessCreateNetworkWirelessAirMarshalRuleMatch struct {
	String types.String `tfsdk:"string"`
	Type   types.String `tfsdk:"type"`
}

type ResponseWirelessCreateNetworkWirelessAirMarshalRuleNetwork struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

type RequestWirelessCreateNetworkWirelessAirMarshalRuleRs struct {
	Match *RequestWirelessCreateNetworkWirelessAirMarshalRuleMatchRs `tfsdk:"match"`
	Type  types.String                                               `tfsdk:"type"`
}

type RequestWirelessCreateNetworkWirelessAirMarshalRuleMatchRs struct {
	String types.String `tfsdk:"string"`
	Type   types.String `tfsdk:"type"`
}

// FromBody
func (r *NetworksWirelessAirMarshalRules) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestWirelessCreateNetworkWirelessAirMarshalRule {
	emptyString := ""
	re := *r.Parameters
	var requestWirelessCreateNetworkWirelessAirMarshalRuleMatch *merakigosdk.RequestWirelessCreateNetworkWirelessAirMarshalRuleMatch
	if re.Match != nil {
		string := re.Match.String.ValueString()
		typeR := re.Match.Type.ValueString()
		requestWirelessCreateNetworkWirelessAirMarshalRuleMatch = &merakigosdk.RequestWirelessCreateNetworkWirelessAirMarshalRuleMatch{
			String: string,
			Type:   typeR,
		}
	}
	typeR := new(string)
	if !re.Type.IsUnknown() && !re.Type.IsNull() {
		*typeR = re.Type.ValueString()
	} else {
		typeR = &emptyString
	}
	out := merakigosdk.RequestWirelessCreateNetworkWirelessAirMarshalRule{
		Match: requestWirelessCreateNetworkWirelessAirMarshalRuleMatch,
		Type:  *typeR,
	}
	return &out
}

// ToBody
func ResponseWirelessCreateNetworkWirelessAirMarshalRuleItemToBody(state NetworksWirelessAirMarshalRules, response *merakigosdk.ResponseWirelessCreateNetworkWirelessAirMarshalRule) NetworksWirelessAirMarshalRules {
	itemState := ResponseWirelessCreateNetworkWirelessAirMarshalRule{
		CreatedAt: types.StringValue(response.CreatedAt),
		Match: func() *ResponseWirelessCreateNetworkWirelessAirMarshalRuleMatch {
			if response.Match != nil {
				return &ResponseWirelessCreateNetworkWirelessAirMarshalRuleMatch{
					String: types.StringValue(response.Match.String),
					Type:   types.StringValue(response.Match.Type),
				}
			}
			return nil
		}(),
		Network: func() *ResponseWirelessCreateNetworkWirelessAirMarshalRuleNetwork {
			if response.Network != nil {
				return &ResponseWirelessCreateNetworkWirelessAirMarshalRuleNetwork{
					ID:   types.StringValue(response.Network.ID),
					Name: types.StringValue(response.Network.Name),
				}
			}
			return nil
		}(),
		RuleID:    types.StringValue(response.RuleID),
		Type:      types.StringValue(response.Type),
		UpdatedAt: types.StringValue(response.UpdatedAt),
	}
	state.Item = &itemState
	return state
}
