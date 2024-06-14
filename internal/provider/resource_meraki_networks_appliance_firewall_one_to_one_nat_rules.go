package provider

// RESOURCE NORMAL
import (
	"context"

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
	_ resource.Resource              = &NetworksApplianceFirewallOneToOneNatRulesResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceFirewallOneToOneNatRulesResource{}
)

func NewNetworksApplianceFirewallOneToOneNatRulesResource() resource.Resource {
	return &NetworksApplianceFirewallOneToOneNatRulesResource{}
}

type NetworksApplianceFirewallOneToOneNatRulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceFirewallOneToOneNatRulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceFirewallOneToOneNatRulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_firewall_one_to_one_nat_rules"
}

func (r *NetworksApplianceFirewallOneToOneNatRulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"rules": schema.SetNestedAttribute{
				MarkdownDescription: `An array of 1:1 nat rules`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"allowed_inbound": schema.SetNestedAttribute{
							MarkdownDescription: `The ports this mapping will provide access on, and the remote IPs that will be allowed access to the resource`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"allowed_ips": schema.SetAttribute{
										MarkdownDescription: `An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges, or 'any'`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
									"destination_ports": schema.SetAttribute{
										MarkdownDescription: `An array of ports or port ranges that will be forwarded to the host on the LAN`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
									"protocol": schema.StringAttribute{
										MarkdownDescription: `Either of the following: 'tcp', 'udp', 'icmp-ping' or 'any'`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
										Validators: []validator.String{
											stringvalidator.OneOf(
												"any",
												"icmp-ping",
												"tcp",
												"udp",
											),
										},
									},
								},
							},
						},
						"lan_ip": schema.StringAttribute{
							MarkdownDescription: `The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN`,
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
						"public_ip": schema.StringAttribute{
							MarkdownDescription: `The IP address that will be used to access the internal resource from the WAN`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"uplink": schema.StringAttribute{
							MarkdownDescription: `The physical WAN interface on which the traffic will arrive ('internet1' or, if available, 'internet2')`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"internet1",
									"internet2",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (r *NetworksApplianceFirewallOneToOneNatRulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceFirewallOneToOneNatRulesRs

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
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceFirewallOneToOneNatRules(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceFirewallOneToOneNatRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceFirewallOneToOneNatRules only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallOneToOneNatRules(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallOneToOneNatRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallOneToOneNatRules",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceFirewallOneToOneNatRules(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceFirewallOneToOneNatRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallOneToOneNatRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceFirewallOneToOneNatRulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceFirewallOneToOneNatRulesRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceFirewallOneToOneNatRules(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceFirewallOneToOneNatRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceFirewallOneToOneNatRules",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceFirewallOneToOneNatRulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceFirewallOneToOneNatRulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceFirewallOneToOneNatRulesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceFirewallOneToOneNatRules(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceFirewallOneToOneNatRules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceFirewallOneToOneNatRules",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceFirewallOneToOneNatRulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceFirewallOneToOneNatRules", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceFirewallOneToOneNatRulesRs struct {
	NetworkID types.String                                                           `tfsdk:"network_id"`
	Rules     *[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesRs `tfsdk:"rules"`
}

type ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesRs struct {
	AllowedInbound *[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInboundRs `tfsdk:"allowed_inbound"`
	LanIP          types.String                                                                         `tfsdk:"lan_ip"`
	Name           types.String                                                                         `tfsdk:"name"`
	PublicIP       types.String                                                                         `tfsdk:"public_ip"`
	Uplink         types.String                                                                         `tfsdk:"uplink"`
}

type ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInboundRs struct {
	AllowedIPs       types.Set    `tfsdk:"allowed_ips"`
	DestinationPorts types.Set    `tfsdk:"destination_ports"`
	Protocol         types.String `tfsdk:"protocol"`
}

// FromBody
func (r *NetworksApplianceFirewallOneToOneNatRulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRules {
	var requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRules []merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRules
	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			var requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound []merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound
			if rItem1.AllowedInbound != nil {
				for _, rItem2 := range *rItem1.AllowedInbound { //AllowedInbound// name: allowedInbound
					var allowedIPs []string = nil
					//Hoola aqui
					rItem2.AllowedIPs.ElementsAs(ctx, &allowedIPs, false)
					var destinationPorts []string = nil
					//Hoola aqui
					rItem2.DestinationPorts.ElementsAs(ctx, &destinationPorts, false)
					protocol := rItem2.Protocol.ValueString()
					requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound = append(requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound, merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound{
						AllowedIPs:       allowedIPs,
						DestinationPorts: destinationPorts,
						Protocol:         protocol,
					})
				}
			}
			lanIP := rItem1.LanIP.ValueString()
			name := rItem1.Name.ValueString()
			publicIP := rItem1.PublicIP.ValueString()
			uplink := rItem1.Uplink.ValueString()
			requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRules = append(requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRules, merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRules{
				AllowedInbound: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound {
					if len(requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound) > 0 {
						return &requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInbound
					}
					return nil
				}(),
				LanIP:    lanIP,
				Name:     name,
				PublicIP: publicIP,
				Uplink:   uplink,
			})
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRules{
		Rules: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRules {
			return &requestApplianceUpdateNetworkApplianceFirewallOneToOneNatRulesRules
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesItemToBodyRs(state NetworksApplianceFirewallOneToOneNatRulesRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRules, is_read bool) NetworksApplianceFirewallOneToOneNatRulesRs {
	itemState := NetworksApplianceFirewallOneToOneNatRulesRs{
		Rules: func() *[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesRs {
			if response.Rules != nil {
				result := make([]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesRs{
						AllowedInbound: func() *[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInboundRs {
							if rules.AllowedInbound != nil {
								result := make([]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInboundRs, len(*rules.AllowedInbound))
								for i, allowedInbound := range *rules.AllowedInbound {
									result[i] = ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInboundRs{
										AllowedIPs:       StringSliceToSet(allowedInbound.AllowedIPs),
										DestinationPorts: StringSliceToSet(allowedInbound.DestinationPorts),
										Protocol:         types.StringValue(allowedInbound.Protocol),
									}
								}
								return &result
							}
							return &[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesAllowedInboundRs{}
						}(),
						LanIP:    types.StringValue(rules.LanIP),
						Name:     types.StringValue(rules.Name),
						PublicIP: types.StringValue(rules.PublicIP),
						Uplink:   types.StringValue(rules.Uplink),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceFirewallOneToOneNatRulesRulesRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceFirewallOneToOneNatRulesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceFirewallOneToOneNatRulesRs)
}
