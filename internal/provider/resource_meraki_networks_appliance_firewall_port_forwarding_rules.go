package provider

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

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
	_ resource.Resource              = &NetworksApplianceFirewallPortForwardingRulesResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceFirewallPortForwardingRulesResource{}
)

func NewNetworksApplianceFirewallPortForwardingRulesResource() resource.Resource {
	return &NetworksApplianceFirewallPortForwardingRulesResource{}
}

type NetworksApplianceFirewallPortForwardingRulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceFirewallPortForwardingRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceFirewallPortForwardingRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_port_forwarding_rules"
}

func (r *NetworksApplianceFirewallPortForwardingRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"rules": schema.SetNestedAttribute{
				MarkdownDescription: `An array of port forwarding params`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"allowed_ips": schema.SetAttribute{
							MarkdownDescription: `An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges (or any)`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"lan_ip": schema.StringAttribute{
							MarkdownDescription: `The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"local_port": schema.StringAttribute{
							MarkdownDescription: `A port or port ranges that will receive the forwarded traffic from the WAN`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"name": schema.StringAttribute{
							MarkdownDescription: `A descriptive name for the rule`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: `TCP or UDP`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"public_port": schema.StringAttribute{
							MarkdownDescription: `A port or port ranges that will be forwarded to the host on the LAN`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"uplink": schema.StringAttribute{
							MarkdownDescription: `The physical WAN interface on which the traffic will arrive ('internet1' or, if available, 'internet2' or 'both')`,
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

func (r *NetworksApplianceFirewallPortForwardingRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceFirewallPortForwardingRulesRs

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
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceFirewallPortForwardingRules(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceFirewallPortForwardingRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceFirewallPortForwardingRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallPortForwardingRules(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallPortForwardingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallPortForwardingRules",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceFirewallPortForwardingRules(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallPortForwardingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallPortForwardingRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceFirewallPortForwardingRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceFirewallPortForwardingRulesRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceFirewallPortForwardingRules(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceFirewallPortForwardingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallPortForwardingRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceFirewallPortForwardingRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceFirewallPortForwardingRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceFirewallPortForwardingRulesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallPortForwardingRules(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallPortForwardingRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallPortForwardingRules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceFirewallPortForwardingRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceFirewallPortForwardingRules", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceFirewallPortForwardingRulesRs struct {
	NetworkID types.String                                                              `tfsdk:"network_id"`
	Rules     *[]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRulesRs `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRulesRs struct {
	AllowedIPs types.Set    `tfsdk:"allowed_ips"`
	LanIP      types.String `tfsdk:"lan_ip"`
	LocalPort  types.String `tfsdk:"local_port"`
	Name       types.String `tfsdk:"name"`
	Protocol   types.String `tfsdk:"protocol"`
	PublicPort types.String `tfsdk:"public_port"`
	Uplink     types.String `tfsdk:"uplink"`
}

// FromBody
func (r *NetworksApplianceFirewallPortForwardingRulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRules {
	var requestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules []merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules
	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			var allowedIPs []string = nil
			//Hoola aqui
			rItem1.AllowedIPs.ElementsAs(ctx, &allowedIPs, false)
			lanIP := rItem1.LanIP.ValueString()
			localPort := rItem1.LocalPort.ValueString()
			name := rItem1.Name.ValueString()
			protocol := rItem1.Protocol.ValueString()
			publicPort := rItem1.PublicPort.ValueString()
			uplink := rItem1.Uplink.ValueString()
			requestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules = append(requestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules, merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules{
				AllowedIPs: allowedIPs,
				LanIP:      lanIP,
				LocalPort:  localPort,
				Name:       name,
				Protocol:   protocol,
				PublicPort: publicPort,
				Uplink:     uplink,
			})
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRules{
		Rules: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules {

			return &requestApplianceUpdateNetworkApplianceFirewallPortForwardingRulesRules

		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesItemToBodyRs(state NetworksApplianceFirewallPortForwardingRulesRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallPortForwardingRules, is_read bool) NetworksApplianceFirewallPortForwardingRulesRs {
	itemState := NetworksApplianceFirewallPortForwardingRulesRs{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRulesRs {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRulesRs{
						AllowedIPs: StringSliceToSet(rules.AllowedIPs),
						LanIP:      types.StringValue(rules.LanIP),
						LocalPort:  types.StringValue(rules.LocalPort),
						Name:       types.StringValue(rules.Name),
						Protocol:   types.StringValue(rules.Protocol),
						PublicPort: types.StringValue(rules.PublicPort),
						Uplink:     types.StringValue(rules.Uplink),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceFirewallPortForwardingRulesRulesRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceFirewallPortForwardingRulesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceFirewallPortForwardingRulesRs)
}
