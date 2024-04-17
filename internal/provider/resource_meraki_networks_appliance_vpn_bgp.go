package provider

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

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
	_ resource.Resource              = &NetworksApplianceVpnBgpResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceVpnBgpResource{}
)

func NewNetworksApplianceVpnBgpResource() resource.Resource {
	return &NetworksApplianceVpnBgpResource{}
}

type NetworksApplianceVpnBgpResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceVpnBgpResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceVpnBgpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_vpn_bgp"
}

func (r *NetworksApplianceVpnBgpResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"as_number": schema.Int64Attribute{
				MarkdownDescription: `An Autonomous System Number (ASN) is required if you are to run BGP and peer with another BGP Speaker outside of the Auto VPN domain. This ASN will be applied to the entire Auto VPN domain. The entire 4-byte ASN range is supported. So, the ASN must be an integer between 1 and 4294967295. When absent, this field is not updated. If no value exists then it defaults to 64512.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean value to enable or disable the BGP configuration. When BGP is enabled, the asNumber (ASN) will be autopopulated with the preconfigured ASN at other Hubs or a default value if there is no ASN configured.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"ibgp_hold_timer": schema.Int64Attribute{
				MarkdownDescription: `The iBGP holdtimer in seconds. The iBGP holdtimer must be an integer between 12 and 240. When absent, this field is not updated. If no value exists then it defaults to 240.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"neighbors": schema.SetNestedAttribute{
				MarkdownDescription: `List of BGP neighbors. This list replaces the existing set of neighbors. When absent, this field is not updated.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"allow_transit": schema.BoolAttribute{
							MarkdownDescription: `When this feature is on, the Meraki device will advertise routes learned from other Autonomous Systems, thereby allowing traffic between Autonomous Systems to transit this AS. When absent, it defaults to false.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"authentication": schema.SingleNestedAttribute{
							MarkdownDescription: `Authentication settings between BGP peers.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"password": schema.StringAttribute{
									MarkdownDescription: `Password to configure MD5 authentication between BGP peers.`,
									Sensitive:           true,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"ebgp_hold_timer": schema.Int64Attribute{
							MarkdownDescription: `The eBGP hold timer in seconds for each neighbor. The eBGP hold timer must be an integer between 12 and 240.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"ebgp_multihop": schema.Int64Attribute{
							MarkdownDescription: `Configure this if the neighbor is not adjacent. The eBGP multi-hop must be an integer between 1 and 255.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"ip": schema.StringAttribute{
							MarkdownDescription: `The IPv4 address of the neighbor`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"ipv6": schema.SingleNestedAttribute{
							MarkdownDescription: `Information regarding IPv6 address of the neighbor, Required if *ip* is not present.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"address": schema.StringAttribute{
									MarkdownDescription: `The IPv6 address of the neighbor.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
						"next_hop_ip": schema.StringAttribute{
							MarkdownDescription: `The IPv4 address of the remote BGP peer that will establish a TCP session with the local MX.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"receive_limit": schema.Int64Attribute{
							MarkdownDescription: `The receive limit is the maximum number of routes that can be received from any BGP peer. The receive limit must be an integer between 0 and 4294967295. When absent, it defaults to 0.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"remote_as_number": schema.Int64Attribute{
							MarkdownDescription: `Remote ASN of the neighbor. The remote ASN must be an integer between 1 and 4294967295.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"source_interface": schema.StringAttribute{
							MarkdownDescription: `The output interface for peering with the remote BGP peer. Valid values are: 'wan1', 'wan2' or 'vlan{VLAN ID}'(e.g. 'vlan123').`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"ttl_security": schema.SingleNestedAttribute{
							MarkdownDescription: `Settings for BGP TTL security to protect BGP peering sessions from forged IP attacks.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.UseStateForUnknown(),
							},
							Attributes: map[string]schema.Attribute{

								"enabled": schema.BoolAttribute{
									MarkdownDescription: `Boolean value to enable or disable BGP TTL security.`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
								},
							},
						},
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
		},
	}
}

func (r *NetworksApplianceVpnBgpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceVpnBgpRs

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
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceVpnBgp(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceVpnBgp only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceVpnBgp only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceVpnBgp(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceVpnBgp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceVpnBgp",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceVpnBgp(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceVpnBgp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceVpnBgp",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceVpnBgpItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceVpnBgpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceVpnBgpRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceVpnBgp(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceVpnBgp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceVpnBgp",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceVpnBgpItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceVpnBgpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceVpnBgpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceVpnBgpRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceVpnBgp(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceVpnBgp",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceVpnBgp",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceVpnBgpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceVpnBgp", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceVpnBgpRs struct {
	NetworkID     types.String                                             `tfsdk:"network_id"`
	AsNumber      types.Int64                                              `tfsdk:"as_number"`
	Enabled       types.Bool                                               `tfsdk:"enabled"`
	IbgpHoldTimer types.Int64                                              `tfsdk:"ibgp_hold_timer"`
	Neighbors     *[]ResponseApplianceGetNetworkApplianceVpnBgpNeighborsRs `tfsdk:"neighbors"`
}

type ResponseApplianceGetNetworkApplianceVpnBgpNeighborsRs struct {
	AllowTransit    types.Bool                                                             `tfsdk:"allow_transit"`
	EbgpHoldTimer   types.Int64                                                            `tfsdk:"ebgp_hold_timer"`
	EbgpMultihop    types.Int64                                                            `tfsdk:"ebgp_multihop"`
	IP              types.String                                                           `tfsdk:"ip"`
	ReceiveLimit    types.Int64                                                            `tfsdk:"receive_limit"`
	RemoteAsNumber  types.Int64                                                            `tfsdk:"remote_as_number"`
	Authentication  *RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthenticationRs `tfsdk:"authentication"`
	IPv6            *RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsIpv6Rs           `tfsdk:"ipv6"`
	NextHopIP       types.String                                                           `tfsdk:"next_hop_ip"`
	SourceInterface types.String                                                           `tfsdk:"source_interface"`
	TtlSecurity     *RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurityRs    `tfsdk:"ttl_security"`
}

type RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthenticationRs struct {
	Password types.String `tfsdk:"password"`
}

type RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsIpv6Rs struct {
	Address types.String `tfsdk:"address"`
}

type RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurityRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

// FromBody
func (r *NetworksApplianceVpnBgpRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgp {
	asNumber := new(int64)
	if !r.AsNumber.IsUnknown() && !r.AsNumber.IsNull() {
		*asNumber = r.AsNumber.ValueInt64()
	} else {
		asNumber = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	ibgpHoldTimer := new(int64)
	if !r.IbgpHoldTimer.IsUnknown() && !r.IbgpHoldTimer.IsNull() {
		*ibgpHoldTimer = r.IbgpHoldTimer.ValueInt64()
	} else {
		ibgpHoldTimer = nil
	}
	var requestApplianceUpdateNetworkApplianceVpnBgpNeighbors []merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighbors
	if r.Neighbors != nil {
		for _, rItem1 := range *r.Neighbors {
			allowTransit := func() *bool {
				if !rItem1.AllowTransit.IsUnknown() && !rItem1.AllowTransit.IsNull() {
					return rItem1.AllowTransit.ValueBoolPointer()
				}
				return nil
			}()
			var requestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication *merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication
			if rItem1.Authentication != nil {
				password := rItem1.Authentication.Password.ValueString()
				requestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication = &merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication{
					Password: password,
				}
			}
			ebgpHoldTimer := func() *int64 {
				if !rItem1.EbgpHoldTimer.IsUnknown() && !rItem1.EbgpHoldTimer.IsNull() {
					return rItem1.EbgpHoldTimer.ValueInt64Pointer()
				}
				return nil
			}()
			ebgpMultihop := func() *int64 {
				if !rItem1.EbgpMultihop.IsUnknown() && !rItem1.EbgpMultihop.IsNull() {
					return rItem1.EbgpMultihop.ValueInt64Pointer()
				}
				return nil
			}()
			iP := rItem1.IP.ValueString()
			var requestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6 *merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6
			if rItem1.IPv6 != nil {
				address := rItem1.IPv6.Address.ValueString()
				requestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6 = &merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6{
					Address: address,
				}
			}
			nextHopIP := rItem1.NextHopIP.ValueString()
			receiveLimit := func() *int64 {
				if !rItem1.ReceiveLimit.IsUnknown() && !rItem1.ReceiveLimit.IsNull() {
					return rItem1.ReceiveLimit.ValueInt64Pointer()
				}
				return nil
			}()
			remoteAsNumber := func() *int64 {
				if !rItem1.RemoteAsNumber.IsUnknown() && !rItem1.RemoteAsNumber.IsNull() {
					return rItem1.RemoteAsNumber.ValueInt64Pointer()
				}
				return nil
			}()
			sourceInterface := rItem1.SourceInterface.ValueString()
			var requestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity *merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity
			if rItem1.TtlSecurity != nil {
				enabled := func() *bool {
					if !rItem1.TtlSecurity.Enabled.IsUnknown() && !rItem1.TtlSecurity.Enabled.IsNull() {
						return rItem1.TtlSecurity.Enabled.ValueBoolPointer()
					}
					return nil
				}()
				requestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity = &merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity{
					Enabled: enabled,
				}
			}
			requestApplianceUpdateNetworkApplianceVpnBgpNeighbors = append(requestApplianceUpdateNetworkApplianceVpnBgpNeighbors, merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighbors{
				AllowTransit:    allowTransit,
				Authentication:  requestApplianceUpdateNetworkApplianceVpnBgpNeighborsAuthentication,
				EbgpHoldTimer:   int64ToIntPointer(ebgpHoldTimer),
				EbgpMultihop:    int64ToIntPointer(ebgpMultihop),
				IP:              iP,
				IPv6:            requestApplianceUpdateNetworkApplianceVpnBgpNeighborsIPv6,
				NextHopIP:       nextHopIP,
				ReceiveLimit:    int64ToIntPointer(receiveLimit),
				RemoteAsNumber:  int64ToIntPointer(remoteAsNumber),
				SourceInterface: sourceInterface,
				TtlSecurity:     requestApplianceUpdateNetworkApplianceVpnBgpNeighborsTtlSecurity,
			})
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgp{
		AsNumber:      int64ToIntPointer(asNumber),
		Enabled:       enabled,
		IbgpHoldTimer: int64ToIntPointer(ibgpHoldTimer),
		Neighbors: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceVpnBgpNeighbors {
			if len(requestApplianceUpdateNetworkApplianceVpnBgpNeighbors) > 0 {
				return &requestApplianceUpdateNetworkApplianceVpnBgpNeighbors
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceVpnBgpItemToBodyRs(state NetworksApplianceVpnBgpRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceVpnBgp, is_read bool) NetworksApplianceVpnBgpRs {
	itemState := NetworksApplianceVpnBgpRs{
		AsNumber: func() types.Int64 {
			if response.AsNumber != nil {
				return types.Int64Value(int64(*response.AsNumber))
			}
			return types.Int64{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		IbgpHoldTimer: func() types.Int64 {
			if response.IbgpHoldTimer != nil {
				return types.Int64Value(int64(*response.IbgpHoldTimer))
			}
			return types.Int64{}
		}(),
		Neighbors: func() *[]ResponseApplianceGetNetworkApplianceVpnBgpNeighborsRs {
			if response.Neighbors != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVpnBgpNeighborsRs, len(*response.Neighbors))
				for i, neighbors := range *response.Neighbors {
					result[i] = ResponseApplianceGetNetworkApplianceVpnBgpNeighborsRs{
						AllowTransit: func() types.Bool {
							if neighbors.AllowTransit != nil {
								return types.BoolValue(*neighbors.AllowTransit)
							}
							return types.Bool{}
						}(),
						EbgpHoldTimer: func() types.Int64 {
							if neighbors.EbgpHoldTimer != nil {
								return types.Int64Value(int64(*neighbors.EbgpHoldTimer))
							}
							return types.Int64{}
						}(),
						EbgpMultihop: func() types.Int64 {
							if neighbors.EbgpMultihop != nil {
								return types.Int64Value(int64(*neighbors.EbgpMultihop))
							}
							return types.Int64{}
						}(),
						IP: types.StringValue(neighbors.IP),
						ReceiveLimit: func() types.Int64 {
							if neighbors.ReceiveLimit != nil {
								return types.Int64Value(int64(*neighbors.ReceiveLimit))
							}
							return types.Int64{}
						}(),
						RemoteAsNumber: func() types.Int64 {
							if neighbors.RemoteAsNumber != nil {
								return types.Int64Value(int64(*neighbors.RemoteAsNumber))
							}
							return types.Int64{}
						}(),
					}
				}
				return &result
			}
			return &[]ResponseApplianceGetNetworkApplianceVpnBgpNeighborsRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceVpnBgpRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceVpnBgpRs)
}
