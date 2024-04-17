package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsFirewallL7FirewallRulesResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsFirewallL7FirewallRulesResource{}
)

func NewNetworksWirelessSSIDsFirewallL7FirewallRulesResource() resource.Resource {
	return &NetworksWirelessSSIDsFirewallL7FirewallRulesResource{}
}

type NetworksWirelessSSIDsFirewallL7FirewallRulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_firewall_l7_firewall_rules"
}

func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"number": schema.StringAttribute{
				MarkdownDescription: `number path parameter.`,
				Required:            true,
			},
			"rules": schema.SetNestedAttribute{
				MarkdownDescription: `An ordered array of the firewall rules for this SSID (not including the local LAN access rule or the default rule).`,
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
							MarkdownDescription: `Type of the L7 firewall rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'`,
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
							MarkdownDescription: `The value of what needs to get blocked. Format of the value varies depending on type of the firewall rule selected.`,
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
	}
}

func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsFirewallL7FirewallRulesRs

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
	vvNumber := data.Number.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDFirewallL7FirewallRules(vvNetworkID, vvNumber)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDsFirewallL7FirewallRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDsFirewallL7FirewallRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDFirewallL7FirewallRules(vvNetworkID, vvNumber, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDFirewallL7FirewallRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDFirewallL7FirewallRules",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDFirewallL7FirewallRules(vvNetworkID, vvNumber)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDFirewallL7FirewallRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDFirewallL7FirewallRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsFirewallL7FirewallRulesRs

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
	vvNumber := data.Number.ValueString()
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDFirewallL7FirewallRules(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkWirelessSSIDFirewallL7FirewallRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDFirewallL7FirewallRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessSSIDsFirewallL7FirewallRulesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvNumber := data.Number.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDFirewallL7FirewallRules(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDFirewallL7FirewallRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDFirewallL7FirewallRules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessSSIDsFirewallL7FirewallRules", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsFirewallL7FirewallRulesRs struct {
	NetworkID types.String                                                            `tfsdk:"network_id"`
	Number    types.String                                                            `tfsdk:"number"`
	Rules     *[]ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesRs `tfsdk:"rules"`
}

type ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesRs struct {
	Policy types.String `tfsdk:"policy"`
	Type   types.String `tfsdk:"type"`
	Value  types.String `tfsdk:"value"`
}

// FromBody
func (r *NetworksWirelessSSIDsFirewallL7FirewallRulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules {
	var requestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules
	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			policy := rItem1.Policy.ValueString()
			typeR := rItem1.Type.ValueString()
			value := rItem1.Value.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules = append(requestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules{
				Policy: policy,
				Type:   typeR,
				Value:  value,
			})
		}
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRules{
		Rules: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules {
			if len(requestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDFirewallL7FirewallRulesRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRulesItemToBodyRs(state NetworksWirelessSSIDsFirewallL7FirewallRulesRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDFirewallL7FirewallRules, is_read bool) NetworksWirelessSSIDsFirewallL7FirewallRulesRs {
	itemState := NetworksWirelessSSIDsFirewallL7FirewallRulesRs{
		Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesRs {
			if response.Rules != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesRs{
						Policy: types.StringValue(rules.Policy),
						Type:   types.StringValue(rules.Type),
						Value:  types.StringValue(rules.Value),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidFirewallL7FirewallRulesRulesRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsFirewallL7FirewallRulesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsFirewallL7FirewallRulesRs)
}
