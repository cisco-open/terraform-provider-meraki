package provider

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceTrafficShapingRulesResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceTrafficShapingRulesResource{}
)

func NewNetworksApplianceTrafficShapingRulesResource() resource.Resource {
	return &NetworksApplianceTrafficShapingRulesResource{}
}

type NetworksApplianceTrafficShapingRulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceTrafficShapingRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceTrafficShapingRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_traffic_shaping_rules"
}

func (r *NetworksApplianceTrafficShapingRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
										Validators: []validator.String{
											stringvalidator.OneOf(
												"application",
												"applicationCategory",
												"host",
												"ipRange",
												"localNet",
												"port",
											),
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
									"value_list": schema.SetAttribute{
										MarkdownDescription: `The 'value_list' of what you want to block. Send a list in request`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},
										ElementType: types.StringType,
										Default:     setdefault.StaticValue(types.SetNull(basetypes.StringType{})),
									},
									"value_obj": schema.SingleNestedAttribute{
										MarkdownDescription: `The 'value_obj' of what you want to block. Send a dict in request`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{
											"id": schema.StringAttribute{
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"name": schema.StringAttribute{
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
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
						"priority": schema.StringAttribute{
							MarkdownDescription: `    A string, indicating the priority level for packets bound to your rule.
    Can be 'low', 'normal' or 'high'.
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
		},
	}
}

func (r *NetworksApplianceTrafficShapingRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceTrafficShapingRulesRs

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
	//Item
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceTrafficShapingRules(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceTrafficShapingRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceTrafficShapingRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceTrafficShapingRules(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceTrafficShapingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceTrafficShapingRules",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceTrafficShapingRules(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceTrafficShapingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceTrafficShapingRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceTrafficShapingRulesItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceTrafficShapingRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceTrafficShapingRulesRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceTrafficShapingRules(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceTrafficShapingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceTrafficShapingRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceTrafficShapingRulesItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceTrafficShapingRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceTrafficShapingRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceTrafficShapingRulesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceTrafficShapingRules(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceTrafficShapingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceTrafficShapingRules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceTrafficShapingRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceTrafficShapingRules", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceTrafficShapingRulesRs struct {
	NetworkID           types.String                                                      `tfsdk:"network_id"`
	DefaultRulesEnabled types.Bool                                                        `tfsdk:"default_rules_enabled"`
	Rules               *[]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesRs `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesRs struct {
	Definitions              *[]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitionsRs            `tfsdk:"definitions"`
	DscpTagValue             types.Int64                                                                             `tfsdk:"dscp_tag_value"`
	PerClientBandwidthLimits *ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsRs `tfsdk:"per_client_bandwidth_limits"`
	Priority                 types.String                                                                            `tfsdk:"priority"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitionsRs struct {
	Type      types.String                                                              `tfsdk:"type"`
	Value     types.String                                                              `tfsdk:"value"`
	ValueList types.Set                                                                 `tfsdk:"value_list"`
	ValueObj  *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj `tfsdk:"value_obj"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsRs struct {
	BandwidthLimits *ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimitsRs `tfsdk:"bandwidth_limits"`
	Settings        types.String                                                                                           `tfsdk:"settings"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimitsRs struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

// FromBody
func (r *NetworksApplianceTrafficShapingRulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingRules {
	defaultRulesEnabled := new(bool)
	if !r.DefaultRulesEnabled.IsUnknown() && !r.DefaultRulesEnabled.IsNull() {
		*defaultRulesEnabled = r.DefaultRulesEnabled.ValueBool()
	} else {
		defaultRulesEnabled = nil
	}
	var requestApplianceUpdateNetworkApplianceTrafficShapingRulesRules []merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRules
	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			var requestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesDefinitions []merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesDefinitions
			if rItem1.Definitions != nil {
				for _, rItem2 := range *rItem1.Definitions { //Definitions// name: definitions
					var valueR interface{}
					typeR := rItem2.Type.ValueString()
					value := rItem2.Value.ValueString()
					var valueList []string
					rItem2.ValueList.ElementsAs(ctx, &valueList, false)
					var requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRulesValue *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRulesValue
					if rItem2.ValueObj != nil {
						name := rItem2.ValueObj.Name.ValueString()
						id := rItem2.ValueObj.ID.ValueString()
						requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRulesValue = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRulesValue{
							ID:   id,
							Name: name,
						}
					}
					if !rItem2.Value.IsNull() && !rItem2.Value.IsUnknown() && rItem2.Type.ValueString() != "blockedCountries" && rItem2.Type.ValueString() != "applicationCategory" {
						valueR = value
					} else {
						if !rItem2.ValueList.IsNull() && !rItem2.ValueList.IsUnknown() && rItem2.Type.ValueString() == "blockedCountries" {
							valueR = valueList
						} else {
							valueR = requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRulesValue
						}
					}
					requestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesDefinitions = append(requestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesDefinitions, merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesDefinitions{
						Type:  typeR,
						Value: valueR,
					})
				}
			}
			dscpTagValue := func() *int64 {
				if !rItem1.DscpTagValue.IsUnknown() && !rItem1.DscpTagValue.IsNull() {
					return rItem1.DscpTagValue.ValueInt64Pointer()
				}
				return nil
			}()
			var requestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits
			if rItem1.PerClientBandwidthLimits != nil {
				var requestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits
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
					requestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits{
						LimitDown: int64ToIntPointer(limitDown),
						LimitUp:   int64ToIntPointer(limitUp),
					}
				}
				settings := rItem1.PerClientBandwidthLimits.Settings.ValueString()
				requestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits{
					BandwidthLimits: requestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimits,
					Settings:        settings,
				}
			}
			priority := rItem1.Priority.ValueString()
			requestApplianceUpdateNetworkApplianceTrafficShapingRulesRules = append(requestApplianceUpdateNetworkApplianceTrafficShapingRulesRules, merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRules{
				Definitions: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesDefinitions {
					if len(requestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesDefinitions) > 0 {
						return &requestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesDefinitions
					}
					return nil
				}(),
				DscpTagValue:             int64ToIntPointer(dscpTagValue),
				PerClientBandwidthLimits: requestApplianceUpdateNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimits,
				Priority:                 priority,
			})
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingRules{
		DefaultRulesEnabled: defaultRulesEnabled,
		Rules: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingRulesRules {
			if len(requestApplianceUpdateNetworkApplianceTrafficShapingRulesRules) > 0 {
				return &requestApplianceUpdateNetworkApplianceTrafficShapingRulesRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceTrafficShapingRulesItemToBodyRs(state NetworksApplianceTrafficShapingRulesRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceTrafficShapingRules, is_read bool) NetworksApplianceTrafficShapingRulesRs {
	itemState := NetworksApplianceTrafficShapingRulesRs{
		DefaultRulesEnabled: func() types.Bool {
			if response.DefaultRulesEnabled != nil {
				return types.BoolValue(*response.DefaultRulesEnabled)
			}
			return types.Bool{}
		}(),
		Rules: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesRs {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesRs{
						Definitions: func() *[]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitionsRs {
							if rules.Definitions != nil {
								result := make([]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitionsRs, len(*rules.Definitions))
								for i, definitions := range *rules.Definitions {
									result[i] = ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitionsRs{
										Type: types.StringValue(definitions.Type),
										Value: func() types.String {
											if definitions.Value == nil {
												return types.StringNull()
											}
											return types.StringValue(*definitions.Value)
										}(),
										ValueList: func() types.Set {
											if definitions.ValueList == nil {
												return types.SetNull(types.StringType)
											}
											return StringSliceToSet(*definitions.ValueList)
										}(),
										ValueObj: func() *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj {
											if definitions.ValueObj == nil {
												return nil
											}
											return &ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj{
												ID:   types.StringValue(definitions.ValueObj.ID),
												Name: types.StringValue(definitions.ValueObj.Name),
											}
										}(),
									}
								}
								return &result
							}
							return &[]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesDefinitionsRs{}
						}(),
						DscpTagValue: func() types.Int64 {
							if rules.DscpTagValue != nil {
								return types.Int64Value(int64(*rules.DscpTagValue))
							}
							return types.Int64{}
						}(),
						PerClientBandwidthLimits: func() *ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsRs {
							if rules.PerClientBandwidthLimits != nil {
								return &ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsRs{
									BandwidthLimits: func() *ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimitsRs {
										if rules.PerClientBandwidthLimits.BandwidthLimits != nil {
											return &ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimitsRs{
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
										return &ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsBandwidthLimitsRs{}
									}(),
									Settings: types.StringValue(rules.PerClientBandwidthLimits.Settings),
								}
							}
							return &ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesPerClientBandwidthLimitsRs{}
						}(),
						Priority: types.StringValue(rules.Priority),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceTrafficShapingRulesRulesRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceTrafficShapingRulesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceTrafficShapingRulesRs)
}
