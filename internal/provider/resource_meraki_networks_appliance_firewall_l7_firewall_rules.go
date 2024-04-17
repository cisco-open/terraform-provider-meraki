package provider

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
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
	_ resource.Resource              = &NetworksApplianceFirewallL7FirewallRulesResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceFirewallL7FirewallRulesResource{}
)

func NewNetworksApplianceFirewallL7FirewallRulesResource() resource.Resource {
	return &NetworksApplianceFirewallL7FirewallRulesResource{}
}

type NetworksApplianceFirewallL7FirewallRulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceFirewallL7FirewallRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceFirewallL7FirewallRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_l7_firewall_rules"
}

func (r *NetworksApplianceFirewallL7FirewallRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"rules": schema.SetNestedAttribute{
				MarkdownDescription: `An ordered array of the MX L7 firewall rules`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"policy": schema.StringAttribute{
							MarkdownDescription: `'Deny' traffic specified by this rule`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"deny",
								),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `Type of the L7 rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'`,
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
									"port",
								),
							},
						},
						"value": schema.StringAttribute{
							MarkdownDescription: `The 'value' of what you want to block. Format of 'value' varies depending on type of the rule. The application categories and application ids can be retrieved from the the 'MX L7 application categories' endpoint. The countries follow the two-letter ISO 3166-1 alpha-2 format.`,
							Computed:            true,
							Optional:            true,
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
		},
	}
}

func (r *NetworksApplianceFirewallL7FirewallRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceFirewallL7FirewallRulesRs

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
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceFirewallL7FirewallRules(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceFirewallL7FirewallRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceFirewallL7FirewallRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallL7FirewallRules(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallL7FirewallRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallL7FirewallRules",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceFirewallL7FirewallRules(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallL7FirewallRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallL7FirewallRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceFirewallL7FirewallRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceFirewallL7FirewallRulesRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceFirewallL7FirewallRules(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceFirewallL7FirewallRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallL7FirewallRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceFirewallL7FirewallRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceFirewallL7FirewallRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceFirewallL7FirewallRulesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallL7FirewallRules(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallL7FirewallRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallL7FirewallRules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceFirewallL7FirewallRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceFirewallL7FirewallRules", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceFirewallL7FirewallRulesRs struct {
	NetworkID types.String                                                          `tfsdk:"network_id"`
	Rules     *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesRs `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesRs struct {
	Policy    types.String                                                              `tfsdk:"policy"`
	Type      types.String                                                              `tfsdk:"type"`
	Value     types.String                                                              `tfsdk:"value"`
	ValueList types.Set                                                                 `tfsdk:"value_list"`
	ValueObj  *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj `tfsdk:"value_obj"`
}

type ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj struct {
	ID   types.String `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// FromBody
func (r *NetworksApplianceFirewallL7FirewallRulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRules {
	var requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRules []merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRules
	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			var valueR interface{}
			policy := rItem1.Policy.ValueString()
			typeR := rItem1.Type.ValueString()
			value := rItem1.Value.ValueString()
			var valueList []string
			rItem1.ValueList.ElementsAs(ctx, &valueList, false)
			var requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRulesValue *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRulesValue
			if rItem1.ValueObj != nil {
				name := rItem1.ValueObj.Name.ValueString()
				id := rItem1.ValueObj.ID.ValueString()
				requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRulesValue = &merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRulesValue{
					ID:   id,
					Name: name,
				}
			}
			if !rItem1.Value.IsNull() && !rItem1.Value.IsUnknown() && rItem1.Type.ValueString() != "blockedCountries" && rItem1.Type.ValueString() != "applicationCategory" {
				valueR = value
			} else {
				if !rItem1.ValueList.IsNull() && !rItem1.ValueList.IsUnknown() && rItem1.Type.ValueString() == "blockedCountries" {
					valueR = valueList
				} else {
					valueR = requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRulesValue
				}
			}
			requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRules = append(requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRules, merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRules{
				Policy: policy,
				Type:   typeR,
				Value:  valueR,
			})
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRules{
		Rules: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRules {
			if len(requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRules) > 0 {
				return &requestApplianceUpdateNetworkApplianceFirewallL7FirewallRulesRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesItemToBodyRs(state NetworksApplianceFirewallL7FirewallRulesRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallL7FirewallRules, is_read bool) NetworksApplianceFirewallL7FirewallRulesRs {
	itemState := NetworksApplianceFirewallL7FirewallRulesRs{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesRs {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesRs{
						Policy: types.StringValue(rules.Policy),
						Type:   types.StringValue(rules.Type),
						Value: func() types.String {
							if rules.Value == nil {
								return types.StringNull()
							}
							return types.StringValue(*rules.Value)
						}(),
						ValueList: func() types.Set {
							if rules.ValueList == nil {
								return types.SetNull(types.StringType)
							}
							return StringSliceToSet(*rules.ValueList)
						}(),
						ValueObj: func() *ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj {
							if rules.ValueObj == nil {
								return nil
							}
							return &ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesValueObj{
								ID:   types.StringValue(rules.ValueObj.ID),
								Name: types.StringValue(rules.ValueObj.Name),
							}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceFirewallL7FirewallRulesRulesRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceFirewallL7FirewallRulesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceFirewallL7FirewallRulesRs)
}
