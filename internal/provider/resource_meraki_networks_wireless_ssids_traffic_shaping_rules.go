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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsTrafficShapingRulesResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsTrafficShapingRulesResource{}
)

func NewNetworksWirelessSSIDsTrafficShapingRulesResource() resource.Resource {
	return &NetworksWirelessSSIDsTrafficShapingRulesResource{}
}

type NetworksWirelessSSIDsTrafficShapingRulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsTrafficShapingRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsTrafficShapingRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_traffic_shaping_rules"
}

func (r *NetworksWirelessSSIDsTrafficShapingRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"default_rules_enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether default traffic shaping rules are enabled (true) or disabled (false). There are 4 default rules, which can be seen on your network's traffic shaping page. Note that default rules count against the rule limit of 8.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"rules": schema.SetNestedAttribute{
				MarkdownDescription: `    An array of traffic shaping rules. Rules are applied in the order that
    they are specified in. An empty list (or null) means no rules. Note that
    you are allowed a maximum of 8 rules.
`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"definitions": schema.SetNestedAttribute{
							MarkdownDescription: `    A list of objects describing the definitions of your traffic shaping rule. At least one definition is required.
`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"type": schema.StringAttribute{
										MarkdownDescription: `The type of definition. Can be one of 'application', 'applicationCategory', 'host', 'port', 'ipRange' or 'localNet'.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
									"value": schema.StringAttribute{
										MarkdownDescription: `    If "type" is 'host', 'port', 'ipRange' or 'localNet', then "value" must be a string, matching either
    a hostname (e.g. "somesite.com"), a port (e.g. 8080), or an IP range ("192.1.0.0",
    "192.1.0.0/16", or "10.1.0.0/16:80"). 'localNet' also supports CIDR notation, excluding
    custom ports.
     If "type" is 'application' or 'applicationCategory', then "value" must be an object
    with the structure { "id": "meraki:layer7/..." }, where "id" is the application category or
    application ID (for a list of IDs for your network, use the trafficShaping/applicationCategories
    endpoint).
`,
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
								},
							},
						},
						"dscp_tag_value": schema.Int64Attribute{
							MarkdownDescription: `    The DSCP tag applied by your rule. null means 'Do not change DSCP tag'.
    For a list of possible tag values, use the trafficShaping/dscpTaggingOptions endpoint.
`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"pcp_tag_value": schema.Int64Attribute{
							MarkdownDescription: `    The PCP tag applied by your rule. Can be 0 (lowest priority) through 7 (highest priority).
    null means 'Do not set PCP tag'.
`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"per_client_bandwidth_limits": schema.SingleNestedAttribute{
							MarkdownDescription: `    An object describing the bandwidth settings for your rule.
`,
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"bandwidth_limits": schema.SingleNestedAttribute{
									MarkdownDescription: `The bandwidth limits object, specifying the upload ('limitUp') and download ('limitDown') speed in Kbps. These are only enforced if 'settings' is set to 'custom'.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"limit_down": schema.Int64Attribute{
											MarkdownDescription: `The maximum download limit (integer, in Kbps).`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
										"limit_up": schema.Int64Attribute{
											MarkdownDescription: `The maximum upload limit (integer, in Kbps).`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Int64{
												int64planmodifier.UseStateForUnknown(),
											},
										},
									},
								},
								"settings": schema.StringAttribute{
									MarkdownDescription: `How bandwidth limits are applied by your rule. Can be one of 'network default', 'ignore' or 'custom'.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
				},
			},
			"traffic_shaping_enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether traffic shaping rules are applied to clients on your SSID.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *NetworksWirelessSSIDsTrafficShapingRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsTrafficShapingRulesRs

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
	vvNumber := data.Number.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDTrafficShapingRules(vvNetworkID, vvNumber)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDsTrafficShapingRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDsTrafficShapingRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDTrafficShapingRules(vvNetworkID, vvNumber, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDTrafficShapingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDTrafficShapingRules",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDTrafficShapingRules(vvNetworkID, vvNumber)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDTrafficShapingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDTrafficShapingRules",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsTrafficShapingRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsTrafficShapingRulesRs

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

	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvNumber := data.Number.ValueString()
	// number
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDTrafficShapingRules(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkWirelessSSIDTrafficShapingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDTrafficShapingRules",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessSSIDsTrafficShapingRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("number"), idParts[1])...)
}

func (r *NetworksWirelessSSIDsTrafficShapingRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessSSIDsTrafficShapingRulesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvNumber := data.Number.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDTrafficShapingRules(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDTrafficShapingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDTrafficShapingRules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsTrafficShapingRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsTrafficShapingRulesRs struct {
	NetworkID             types.String                                                        `tfsdk:"network_id"`
	Number                types.String                                                        `tfsdk:"number"`
	DefaultRulesEnabled   types.Bool                                                          `tfsdk:"default_rules_enabled"`
	Rules                 *[]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesRs `tfsdk:"rules"`
	TrafficShapingEnabled types.Bool                                                          `tfsdk:"traffic_shaping_enabled"`
}

type ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesRs struct {
	Definitions              *[]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesDefinitionsRs            `tfsdk:"definitions"`
	DscpTagValue             types.Int64                                                                               `tfsdk:"dscp_tag_value"`
	PcpTagValue              types.Int64                                                                               `tfsdk:"pcp_tag_value"`
	PerClientBandwidthLimits *ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsRs `tfsdk:"per_client_bandwidth_limits"`
}

type ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesDefinitionsRs struct {
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsRs struct {
	BandwidthLimits *ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimitsRs `tfsdk:"bandwidth_limits"`
	Settings        types.String                                                                                             `tfsdk:"settings"`
}

type ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimitsRs struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

// FromBody
func (r *NetworksWirelessSSIDsTrafficShapingRulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRules {
	defaultRulesEnabled := new(bool)
	if !r.DefaultRulesEnabled.IsUnknown() && !r.DefaultRulesEnabled.IsNull() {
		*defaultRulesEnabled = r.DefaultRulesEnabled.ValueBool()
	} else {
		defaultRulesEnabled = nil
	}
	var requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRules []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRules
	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			var requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions
			if rItem1.Definitions != nil {
				for _, rItem2 := range *rItem1.Definitions { //Definitions// name: definitions
					typeR := rItem2.Type.ValueString()
					value := rItem2.Value.ValueString()
					requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions = append(requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions{
						Type:  typeR,
						Value: value,
					})
				}
			}
			dscpTagValue := func() *int64 {
				if !rItem1.DscpTagValue.IsUnknown() && !rItem1.DscpTagValue.IsNull() {
					return rItem1.DscpTagValue.ValueInt64Pointer()
				}
				return nil
			}()
			pcpTagValue := func() *int64 {
				if !rItem1.PcpTagValue.IsUnknown() && !rItem1.PcpTagValue.IsNull() {
					return rItem1.PcpTagValue.ValueInt64Pointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits
			if rItem1.PerClientBandwidthLimits != nil {
				var requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits
				if rItem1.PerClientBandwidthLimits.BandwidthLimits != nil {
					limitDown := func() *int64 {
						if !rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitDown.IsUnknown() && !rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitDown.IsNull() {
							return rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitDown.ValueInt64Pointer()
						}
						return nil
					}()
					limitUp := func() *int64 {
						if !rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitUp.IsUnknown() && !rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitUp.IsNull() {
							return rItem1.PerClientBandwidthLimits.BandwidthLimits.LimitUp.ValueInt64Pointer()
						}
						return nil
					}()
					requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits{
						LimitDown: int64ToIntPointer(limitDown),
						LimitUp:   int64ToIntPointer(limitUp),
					}
				}
				settings := rItem1.PerClientBandwidthLimits.Settings.ValueString()
				requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits{
					BandwidthLimits: requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits,
					Settings:        settings,
				}
			}
			requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRules = append(requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRules, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRules{
				Definitions: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions {
					if len(requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions) > 0 {
						return &requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesDefinitions
					}
					return nil
				}(),
				DscpTagValue:             int64ToIntPointer(dscpTagValue),
				PcpTagValue:              int64ToIntPointer(pcpTagValue),
				PerClientBandwidthLimits: requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRulesPerClientBandwidthLimits,
			})
		}
	}
	trafficShapingEnabled := new(bool)
	if !r.TrafficShapingEnabled.IsUnknown() && !r.TrafficShapingEnabled.IsNull() {
		*trafficShapingEnabled = r.TrafficShapingEnabled.ValueBool()
	} else {
		trafficShapingEnabled = nil
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRules{
		DefaultRulesEnabled: defaultRulesEnabled,
		Rules: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRules {
			if len(requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRules) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDTrafficShapingRulesRules
			}
			return nil
		}(),
		TrafficShapingEnabled: trafficShapingEnabled,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRulesItemToBodyRs(state NetworksWirelessSSIDsTrafficShapingRulesRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDTrafficShapingRules, is_read bool) NetworksWirelessSSIDsTrafficShapingRulesRs {
	itemState := NetworksWirelessSSIDsTrafficShapingRulesRs{
		DefaultRulesEnabled: func() types.Bool {
			if response.DefaultRulesEnabled != nil {
				return types.BoolValue(*response.DefaultRulesEnabled)
			}
			return types.Bool{}
		}(),
		Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesRs {
			if response.Rules != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesRs{
						Definitions: func() *[]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesDefinitionsRs {
							if rules.Definitions != nil {
								result := make([]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesDefinitionsRs, len(*rules.Definitions))
								for i, definitions := range *rules.Definitions {
									result[i] = ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesDefinitionsRs{
										Type:  types.StringValue(definitions.Type),
										Value: types.StringValue(definitions.Value),
									}
								}
								return &result
							}
							return &[]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesDefinitionsRs{}
						}(),
						DscpTagValue: func() types.Int64 {
							if rules.DscpTagValue != nil {
								return types.Int64Value(int64(*rules.DscpTagValue))
							}
							return types.Int64{}
						}(),
						PcpTagValue: func() types.Int64 {
							if rules.PcpTagValue != nil {
								return types.Int64Value(int64(*rules.PcpTagValue))
							}
							return types.Int64{}
						}(),
						PerClientBandwidthLimits: func() *ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsRs {
							if rules.PerClientBandwidthLimits != nil {
								return &ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsRs{
									BandwidthLimits: func() *ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimitsRs {
										if rules.PerClientBandwidthLimits.BandwidthLimits != nil {
											return &ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimitsRs{
												LimitDown: func() types.Int64 {
													if rules.PerClientBandwidthLimits.BandwidthLimits.LimitDown != nil {
														return types.Int64Value(int64(*rules.PerClientBandwidthLimits.BandwidthLimits.LimitDown))
													}
													return types.Int64{}
												}(),
												LimitUp: func() types.Int64 {
													if rules.PerClientBandwidthLimits.BandwidthLimits.LimitUp != nil {
														return types.Int64Value(int64(*rules.PerClientBandwidthLimits.BandwidthLimits.LimitUp))
													}
													return types.Int64{}
												}(),
											}
										}
										return &ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimitsRs{}
									}(),
									Settings: types.StringValue(rules.PerClientBandwidthLimits.Settings),
								}
							}
							return &ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesPerClientBandwidthLimitsRs{}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidTrafficShapingRulesRulesRs{}
		}(),
		TrafficShapingEnabled: func() types.Bool {
			if response.TrafficShapingEnabled != nil {
				return types.BoolValue(*response.TrafficShapingEnabled)
			}
			return types.Bool{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsTrafficShapingRulesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsTrafficShapingRulesRs)
}
