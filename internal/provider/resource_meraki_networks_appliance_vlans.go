package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceVLANsResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceVLANsResource{}
)

func NewNetworksApplianceVLANsResource() resource.Resource {
	return &NetworksApplianceVLANsResource{}
}

type NetworksApplianceVLANsResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceVLANsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceVLANsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_vlans"
}

func (r *NetworksApplianceVLANsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"appliance_ip": schema.StringAttribute{
				MarkdownDescription: `The local IP of the appliance on the VLAN`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cidr": schema.StringAttribute{
				MarkdownDescription: `CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"dhcp_boot_filename": schema.StringAttribute{
				MarkdownDescription: `DHCP boot option for boot filename`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"dhcp_boot_next_server": schema.StringAttribute{
				MarkdownDescription: `DHCP boot option to direct boot clients to the server to load the boot file from`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"dhcp_boot_options_enabled": schema.BoolAttribute{
				MarkdownDescription: `Use DHCP boot options specified in other properties`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"dhcp_handling": schema.StringAttribute{
				MarkdownDescription: `The appliance's handling of DHCP requests on this VLAN. One of: 'Run a DHCP server', 'Relay DHCP to another server' or 'Do not respond to DHCP requests'`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"Do not respond to DHCP requests",
						"Relay DHCP to another server",
						"Run a DHCP server",
					),
				},
			},
			"dhcp_lease_time": schema.StringAttribute{
				MarkdownDescription: `The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"1 day",
						"1 hour",
						"1 week",
						"12 hours",
						"30 minutes",
						"4 hours",
					),
				},
			},
			"dhcp_options": schema.SetNestedAttribute{
				MarkdownDescription: `The list of DHCP options that will be included in DHCP responses. Each object in the list should have "code", "type", and "value" properties.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"code": schema.StringAttribute{
							MarkdownDescription: `The code for the DHCP option. This should be an integer between 2 and 254.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"type": schema.StringAttribute{
							MarkdownDescription: `The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"hex",
									"integer",
									"ip",
									"text",
								),
							},
						},
						"value": schema.StringAttribute{
							MarkdownDescription: `The value for the DHCP option`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"dhcp_relay_server_ips": schema.SetAttribute{
				MarkdownDescription: `The IPs of the DHCP servers that DHCP requests should be relayed to`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"dns_nameservers": schema.StringAttribute{
				MarkdownDescription: `The DNS nameservers used for DHCP responses, either "upstream_dns", "google_dns", "opendns", or a newline seperated string of IP addresses or domain names`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			// "fixed_ip_assignments": schema.StringAttribute{
			// 	//Todo interface
			// 	MarkdownDescription: `The DHCP fixed IP assignments on the VLAN. This should be an object that contains mappings from MAC addresses to objects that themselves each contain "ip" and "name" string fields. See the sample request/response for more details.`,
			// 	Computed:            true,
			// 	Optional:            true,
			// },
			"group_policy_id": schema.StringAttribute{
				MarkdownDescription: `The id of the desired group policy to apply to the VLAN`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `The VLAN ID of the VLAN`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
					SuppressDiffString(),
				},
			},
			"interface_id": schema.StringAttribute{
				MarkdownDescription: `The interface ID of the VLAN`,
				Computed:            true,
			},
			"ipv6": schema.SingleNestedAttribute{
				MarkdownDescription: `IPv6 configuration on the VLAN`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Enable IPv6 on VLAN`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"prefix_assignments": schema.SetNestedAttribute{
						MarkdownDescription: `Prefix assignments on the VLAN`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"autonomous": schema.BoolAttribute{
									MarkdownDescription: `Auto assign a /64 prefix from the origin to the VLAN`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Bool{
										boolplanmodifier.UseStateForUnknown(),
									},
								},
								"origin": schema.SingleNestedAttribute{
									MarkdownDescription: `The origin of the prefix`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Object{
										objectplanmodifier.UseStateForUnknown(),
									},
									Attributes: map[string]schema.Attribute{

										"interfaces": schema.SetAttribute{
											MarkdownDescription: `Interfaces associated with the prefix`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.Set{
												setplanmodifier.UseStateForUnknown(),
											},

											ElementType: types.StringType,
										},
										"type": schema.StringAttribute{
											MarkdownDescription: `Type of the origin`,
											Computed:            true,
											Optional:            true,
											PlanModifiers: []planmodifier.String{
												stringplanmodifier.UseStateForUnknown(),
											},
											Validators: []validator.String{
												stringvalidator.OneOf(
													"independent",
													"internet",
												),
											},
										},
									},
								},
								"static_appliance_ip6": schema.StringAttribute{
									MarkdownDescription: `Manual configuration of the IPv6 Appliance IP`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"static_prefix": schema.StringAttribute{
									MarkdownDescription: `Manual configuration of a /64 prefix on the VLAN`,
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
			"mandatory_dhcp": schema.SingleNestedAttribute{
				MarkdownDescription: `Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Enable Mandatory DHCP on VLAN.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"mask": schema.Int64Attribute{
				MarkdownDescription: `Mask used for the subnet of all bound to the template networks. Applicable only for template network.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the VLAN`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"reserved_ip_ranges": schema.SetNestedAttribute{
				MarkdownDescription: `The DHCP reserved IP ranges on the VLAN`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"comment": schema.StringAttribute{
							MarkdownDescription: `A text comment for the reserved range`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"end": schema.StringAttribute{
							MarkdownDescription: `The last IP in the reserved range`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"start": schema.StringAttribute{
							MarkdownDescription: `The first IP in the reserved range`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
					},
				},
			},
			"subnet": schema.StringAttribute{
				MarkdownDescription: `The subnet of the VLAN`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"template_vlan_type": schema.StringAttribute{
				MarkdownDescription: `Type of subnetting of the VLAN. Applicable only for template network.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"same",
						"unique",
					),
				},
			},
			"vlan_id": schema.StringAttribute{
				MarkdownDescription: `vlanId path parameter. Vlan ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"vpn_nat_subnet": schema.StringAttribute{
				MarkdownDescription: `The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['vlanId']
//path params to assign NOT EDITABLE ['id']

func (r *NetworksApplianceVLANsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceVLANsRs

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
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceVLANs(vvNetworkID)
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceVLANs",
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
			vvVLANID := strconv.Itoa(*result2["ID"].(*int))
			r.client.Appliance.UpdateNetworkApplianceVLAN(vvNetworkID, vvVLANID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Appliance.GetNetworkApplianceVLAN(vvNetworkID, vvVLANID)
			if responseVerifyItem2 != nil {
				data = ResponseApplianceGetNetworkApplianceVLANItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Appliance.CreateNetworkApplianceVLAN(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkApplianceVLAN",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkApplianceVLAN",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceVLANs(vvNetworkID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceVLANs",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceVLANs",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvVLANID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter VLANID",
				err.Error(),
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Appliance.GetNetworkApplianceVLAN(vvNetworkID, vvVLANID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseApplianceGetNetworkApplianceVLANItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkApplianceVLAN",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceVLAN",
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

func (r *NetworksApplianceVLANsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceVLANsRs

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
	vvVLANID := data.ID.ValueString()
	// vlan_id
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceVLAN(vvNetworkID, vvVLANID)
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
				"Failure when executing GetNetworkApplianceVLAN",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceVLAN",
			err.Error(),
		)
		return
	}

	data = ResponseApplianceGetNetworkApplianceVLANItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceVLANsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("vlan_id"), idParts[1])...)
}

func (r *NetworksApplianceVLANsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceVLANsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvVLANID := data.VLANID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Appliance.UpdateNetworkApplianceVLAN(vvNetworkID, vvVLANID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceVLAN",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceVLAN",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceVLANsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksApplianceVLANsRs
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

	vvNetworkID := state.NetworkID.ValueString()
	vvVLANID := state.VLANID.ValueString()
	_, err := r.client.Appliance.DeleteNetworkApplianceVLAN(vvNetworkID, vvVLANID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkApplianceVLAN", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksApplianceVLANsRs struct {
	NetworkID              types.String                                             `tfsdk:"network_id"`
	VLANID                 types.String                                             `tfsdk:"vlan_id"`
	ApplianceIP            types.String                                             `tfsdk:"appliance_ip"`
	Cidr                   types.String                                             `tfsdk:"cidr"`
	DhcpBootFilename       types.String                                             `tfsdk:"dhcp_boot_filename"`
	DhcpBootNextServer     types.String                                             `tfsdk:"dhcp_boot_next_server"`
	DhcpBootOptionsEnabled types.Bool                                               `tfsdk:"dhcp_boot_options_enabled"`
	DhcpHandling           types.String                                             `tfsdk:"dhcp_handling"`
	DhcpLeaseTime          types.String                                             `tfsdk:"dhcp_lease_time"`
	DhcpOptions            *[]ResponseApplianceGetNetworkApplianceVlanDhcpOptionsRs `tfsdk:"dhcp_options"`
	DhcpRelayServerIPs     types.Set                                                `tfsdk:"dhcp_relay_server_ips"`
	DNSNameservers         types.String                                             `tfsdk:"dns_nameservers"`
	// FixedIPAssignments     *ResponseApplianceGetNetworkApplianceVlanFixedIpAssignmentsRs `tfsdk:"fixed_ip_assignments"`
	GroupPolicyID    types.String                                                  `tfsdk:"group_policy_id"`
	ID               types.String                                                  `tfsdk:"id"`
	InterfaceID      types.String                                                  `tfsdk:"interface_id"`
	IPv6             *ResponseApplianceGetNetworkApplianceVlanIpv6Rs               `tfsdk:"ipv6"`
	MandatoryDhcp    *ResponseApplianceGetNetworkApplianceVlanMandatoryDhcpRs      `tfsdk:"mandatory_dhcp"`
	Mask             types.Int64                                                   `tfsdk:"mask"`
	Name             types.String                                                  `tfsdk:"name"`
	ReservedIPRanges *[]ResponseApplianceGetNetworkApplianceVlanReservedIpRangesRs `tfsdk:"reserved_ip_ranges"`
	Subnet           types.String                                                  `tfsdk:"subnet"`
	TemplateVLANType types.String                                                  `tfsdk:"template_vlan_type"`
	VpnNatSubnet     types.String                                                  `tfsdk:"vpn_nat_subnet"`
}

type ResponseApplianceGetNetworkApplianceVlanDhcpOptionsRs struct {
	Code  types.String `tfsdk:"code"`
	Type  types.String `tfsdk:"type"`
	Value types.String `tfsdk:"value"`
}

type ResponseApplianceGetNetworkApplianceVlanFixedIpAssignmentsRs interface{}

type ResponseApplianceGetNetworkApplianceVlanIpv6Rs struct {
	Enabled           types.Bool                                                         `tfsdk:"enabled"`
	PrefixAssignments *[]ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsRs `tfsdk:"prefix_assignments"`
}

type ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsRs struct {
	Autonomous         types.Bool                                                             `tfsdk:"autonomous"`
	Origin             *ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsOriginRs `tfsdk:"origin"`
	StaticApplianceIP6 types.String                                                           `tfsdk:"static_appliance_ip6"`
	StaticPrefix       types.String                                                           `tfsdk:"static_prefix"`
}

type ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsOriginRs struct {
	Interfaces types.Set    `tfsdk:"interfaces"`
	Type       types.String `tfsdk:"type"`
}

type ResponseApplianceGetNetworkApplianceVlanMandatoryDhcpRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseApplianceGetNetworkApplianceVlanReservedIpRangesRs struct {
	Comment types.String `tfsdk:"comment"`
	End     types.String `tfsdk:"end"`
	Start   types.String `tfsdk:"start"`
}

// FromBody
func (r *NetworksApplianceVLANsRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestApplianceCreateNetworkApplianceVLAN {
	log.Printf("ResquestCreate: %v", r.IPv6)
	emptyString := ""
	applianceIP := new(string)
	if !r.ApplianceIP.IsUnknown() && !r.ApplianceIP.IsNull() {
		*applianceIP = r.ApplianceIP.ValueString()
	} else {
		applianceIP = &emptyString
	}
	cidr := new(string)
	if !r.Cidr.IsUnknown() && !r.Cidr.IsNull() {
		*cidr = r.Cidr.ValueString()
	} else {
		cidr = &emptyString
	}
	groupPolicyID := new(string)
	if !r.GroupPolicyID.IsUnknown() && !r.GroupPolicyID.IsNull() {
		*groupPolicyID = r.GroupPolicyID.ValueString()
	} else {
		groupPolicyID = &emptyString
	}
	iD := new(string)
	if !r.ID.IsUnknown() && !r.ID.IsNull() {
		iD = r.ID.ValueStringPointer()
	} else {
		iD = &emptyString
	}
	var requestApplianceCreateNetworkApplianceVLANIPv6 *merakigosdk.RequestApplianceCreateNetworkApplianceVLANIPv6
	if r.IPv6 != nil {
		enabled := func() *bool {
			if !r.IPv6.Enabled.IsUnknown() && !r.IPv6.Enabled.IsNull() {
				return r.IPv6.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		var requestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignments []merakigosdk.RequestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignments
		if r.IPv6.PrefixAssignments != nil {
			for _, rItem1 := range *r.IPv6.PrefixAssignments {
				autonomous := func() *bool {
					if !rItem1.Autonomous.IsUnknown() && !rItem1.Autonomous.IsNull() {
						return rItem1.Autonomous.ValueBoolPointer()
					}
					return nil
				}()
				var requestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin *merakigosdk.RequestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin
				if rItem1.Origin != nil {
					var interfaces []string
					rItem1.Origin.Interfaces.ElementsAs(ctx, &interfaces, false)
					typeR := rItem1.Origin.Type.ValueString()
					requestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin = &merakigosdk.RequestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin{
						Interfaces: interfaces,
						Type:       typeR,
					}
				}
				staticApplianceIP6 := rItem1.StaticApplianceIP6.ValueString()
				staticPrefix := rItem1.StaticPrefix.ValueString()
				requestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignments = append(requestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignments, merakigosdk.RequestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignments{
					Autonomous:         autonomous,
					Origin:             requestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin,
					StaticApplianceIP6: staticApplianceIP6,
					StaticPrefix:       staticPrefix,
				})
			}
		}
		requestApplianceCreateNetworkApplianceVLANIPv6 = &merakigosdk.RequestApplianceCreateNetworkApplianceVLANIPv6{
			Enabled: enabled,
			PrefixAssignments: func() *[]merakigosdk.RequestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignments {
				if len(requestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignments) > 0 {
					return &requestApplianceCreateNetworkApplianceVLANIPv6PrefixAssignments
				}
				return nil
			}(),
		}
	}
	var requestApplianceCreateNetworkApplianceVLANMandatoryDhcp *merakigosdk.RequestApplianceCreateNetworkApplianceVLANMandatoryDhcp
	if r.MandatoryDhcp != nil {
		enabled := func() *bool {
			if !r.MandatoryDhcp.Enabled.IsUnknown() && !r.MandatoryDhcp.Enabled.IsNull() {
				return r.MandatoryDhcp.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestApplianceCreateNetworkApplianceVLANMandatoryDhcp = &merakigosdk.RequestApplianceCreateNetworkApplianceVLANMandatoryDhcp{
			Enabled: enabled,
		}
	}
	mask := new(int64)
	if !r.Mask.IsUnknown() && !r.Mask.IsNull() {
		*mask = r.Mask.ValueInt64()
	} else {
		mask = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	subnet := new(string)
	if !r.Subnet.IsUnknown() && !r.Subnet.IsNull() {
		*subnet = r.Subnet.ValueString()
	} else {
		subnet = &emptyString
	}
	templateVLANType := new(string)
	if !r.TemplateVLANType.IsUnknown() && !r.TemplateVLANType.IsNull() {
		*templateVLANType = r.TemplateVLANType.ValueString()
	} else {
		templateVLANType = &emptyString
	}
	out := merakigosdk.RequestApplianceCreateNetworkApplianceVLAN{
		ApplianceIP:      *applianceIP,
		Cidr:             *cidr,
		GroupPolicyID:    *groupPolicyID,
		ID:               *iD,
		IPv6:             requestApplianceCreateNetworkApplianceVLANIPv6,
		MandatoryDhcp:    requestApplianceCreateNetworkApplianceVLANMandatoryDhcp,
		Mask:             int64ToIntPointer(mask),
		Name:             *name,
		Subnet:           *subnet,
		TemplateVLANType: *templateVLANType,
	}
	return &out
}
func (r *NetworksApplianceVLANsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceVLAN {
	if r.IPv6 != nil {
		log.Printf("ResquestUpdate: %v", *r.IPv6.PrefixAssignments)
	} else {
		log.Printf("ResquestUpdate: nil")
	}

	applianceIP := ""
	if !r.ApplianceIP.IsUnknown() && !r.ApplianceIP.IsNull() {
		applianceIP = r.ApplianceIP.ValueString()
	}

	cidr := ""
	if !r.Cidr.IsUnknown() && !r.Cidr.IsNull() {
		cidr = r.Cidr.ValueString()
	}

	dhcpBootFilename := ""
	if !r.DhcpBootFilename.IsUnknown() && !r.DhcpBootFilename.IsNull() {
		dhcpBootFilename = r.DhcpBootFilename.ValueString()
	}

	dhcpBootNextServer := ""
	if !r.DhcpBootNextServer.IsUnknown() && !r.DhcpBootNextServer.IsNull() {
		dhcpBootNextServer = r.DhcpBootNextServer.ValueString()
	}

	var dhcpBootOptionsEnabled *bool
	if !r.DhcpBootOptionsEnabled.IsUnknown() && !r.DhcpBootOptionsEnabled.IsNull() {
		enabled := r.DhcpBootOptionsEnabled.ValueBool()
		dhcpBootOptionsEnabled = &enabled
	}

	dhcpHandling := ""
	if !r.DhcpHandling.IsUnknown() && !r.DhcpHandling.IsNull() {
		dhcpHandling = r.DhcpHandling.ValueString()
	}

	dhcpLeaseTime := ""
	if !r.DhcpLeaseTime.IsUnknown() && !r.DhcpLeaseTime.IsNull() {
		dhcpLeaseTime = r.DhcpLeaseTime.ValueString()
	}

	var requestApplianceUpdateNetworkApplianceVLANDhcpOptions []merakigosdk.RequestApplianceUpdateNetworkApplianceVLANDhcpOptions
	if r.DhcpOptions != nil {
		for _, rItem1 := range *r.DhcpOptions {
			code := rItem1.Code.ValueString()
			typeR := rItem1.Type.ValueString()
			value := rItem1.Value.ValueString()
			requestApplianceUpdateNetworkApplianceVLANDhcpOptions = append(requestApplianceUpdateNetworkApplianceVLANDhcpOptions, merakigosdk.RequestApplianceUpdateNetworkApplianceVLANDhcpOptions{
				Code:  code,
				Type:  typeR,
				Value: value,
			})
		}
	}

	var dhcpRelayServerIPs []string
	r.DhcpRelayServerIPs.ElementsAs(ctx, &dhcpRelayServerIPs, false)

	dNSNameservers := ""
	if !r.DNSNameservers.IsUnknown() && !r.DNSNameservers.IsNull() {
		dNSNameservers = r.DNSNameservers.ValueString()
	}

	// requestApplianceUpdateNetworkApplianceVLANFixedIPAssignments := r.FixedIPAssignments.ValueString()
	// var intf interface{} = requestApplianceUpdateNetworkApplianceVLANFixedIPAssignments
	// requestApplianceUpdateNetworkApplianceVLANFixedIPAssignments2, ok := intf.(merakigosdk.RequestApplianceUpdateNetworkApplianceVLANFixedIPAssignments)
	// if !ok {
	// 	requestApplianceUpdateNetworkApplianceVLANFixedIPAssignments2 = nil
	// }
	groupPolicyID := ""
	if !r.GroupPolicyID.IsUnknown() && !r.GroupPolicyID.IsNull() {
		groupPolicyID = r.GroupPolicyID.ValueString()
	}

	var requestApplianceUpdateNetworkApplianceVLANIPv6 *merakigosdk.RequestApplianceUpdateNetworkApplianceVLANIPv6
	if r.IPv6 != nil {
		enabled := func() *bool {
			if !r.IPv6.Enabled.IsUnknown() && !r.IPv6.Enabled.IsNull() {
				return r.IPv6.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		var requestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignments []merakigosdk.RequestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignments
		if r.IPv6.PrefixAssignments != nil && len(*r.IPv6.PrefixAssignments) > 0 {
			for _, rItem1 := range *r.IPv6.PrefixAssignments {
				autonomous := func() *bool {
					if !rItem1.Autonomous.IsUnknown() && !rItem1.Autonomous.IsNull() {
						return rItem1.Autonomous.ValueBoolPointer()
					}
					return nil
				}()
				var requestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin *merakigosdk.RequestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin
				if rItem1.Origin != nil {
					var interfaces []string
					rItem1.Origin.Interfaces.ElementsAs(ctx, &interfaces, false)
					typeR := rItem1.Origin.Type.ValueString()
					requestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin = &merakigosdk.RequestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin{
						Interfaces: interfaces,
						Type:       typeR,
					}
				}
				staticApplianceIP6 := rItem1.StaticApplianceIP6.ValueString()
				staticPrefix := rItem1.StaticPrefix.ValueString()
				requestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignments = append(requestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignments, merakigosdk.RequestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignments{
					Autonomous:         autonomous,
					Origin:             requestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignmentsOrigin,
					StaticApplianceIP6: staticApplianceIP6,
					StaticPrefix:       staticPrefix,
				})
			}
		}
		requestApplianceUpdateNetworkApplianceVLANIPv6 = &merakigosdk.RequestApplianceUpdateNetworkApplianceVLANIPv6{
			Enabled: enabled,
			PrefixAssignments: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignments {
				if len(requestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignments) > 0 {
					return &requestApplianceUpdateNetworkApplianceVLANIPv6PrefixAssignments
				}
				return nil
			}(),
		}
	}

	var requestApplianceUpdateNetworkApplianceVLANMandatoryDhcp *merakigosdk.RequestApplianceUpdateNetworkApplianceVLANMandatoryDhcp
	if r.MandatoryDhcp != nil {
		enabled := func() *bool {
			if !r.MandatoryDhcp.Enabled.IsUnknown() && !r.MandatoryDhcp.Enabled.IsNull() {
				return r.MandatoryDhcp.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestApplianceUpdateNetworkApplianceVLANMandatoryDhcp = &merakigosdk.RequestApplianceUpdateNetworkApplianceVLANMandatoryDhcp{
			Enabled: enabled,
		}
	}

	mask := new(int64)
	if !r.Mask.IsUnknown() && !r.Mask.IsNull() {
		*mask = r.Mask.ValueInt64()
	}

	name := ""
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		name = r.Name.ValueString()
	}

	var requestApplianceUpdateNetworkApplianceVLANReservedIPRanges []merakigosdk.RequestApplianceUpdateNetworkApplianceVLANReservedIPRanges
	if r.ReservedIPRanges != nil {
		for _, rItem1 := range *r.ReservedIPRanges {
			comment := rItem1.Comment.ValueString()
			end := rItem1.End.ValueString()
			start := rItem1.Start.ValueString()
			requestApplianceUpdateNetworkApplianceVLANReservedIPRanges = append(requestApplianceUpdateNetworkApplianceVLANReservedIPRanges, merakigosdk.RequestApplianceUpdateNetworkApplianceVLANReservedIPRanges{
				Comment: comment,
				End:     end,
				Start:   start,
			})
		}
	}

	subnet := ""
	if !r.Subnet.IsUnknown() && !r.Subnet.IsNull() {
		subnet = r.Subnet.ValueString()
	}

	templateVLANType := ""
	if !r.TemplateVLANType.IsUnknown() && !r.TemplateVLANType.IsNull() {
		templateVLANType = r.TemplateVLANType.ValueString()
	}

	vpnNatSubnet := ""
	if !r.VpnNatSubnet.IsUnknown() && !r.VpnNatSubnet.IsNull() {
		vpnNatSubnet = r.VpnNatSubnet.ValueString()
	}

	out := merakigosdk.RequestApplianceUpdateNetworkApplianceVLAN{
		ApplianceIP:            applianceIP,
		Cidr:                   cidr,
		DhcpBootFilename:       dhcpBootFilename,
		DhcpBootNextServer:     dhcpBootNextServer,
		DhcpBootOptionsEnabled: dhcpBootOptionsEnabled,
		DhcpHandling:           dhcpHandling,
		DhcpLeaseTime:          dhcpLeaseTime,
		DhcpOptions: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceVLANDhcpOptions {
			if len(requestApplianceUpdateNetworkApplianceVLANDhcpOptions) > 0 {
				return &requestApplianceUpdateNetworkApplianceVLANDhcpOptions
			}
			return nil
		}(),
		DhcpRelayServerIPs: dhcpRelayServerIPs,
		DNSNameservers:     dNSNameservers,
		// FixedIPAssignments:     func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceVLANFixedIPAssignments2 {
		GroupPolicyID: groupPolicyID,
		IPv6:          requestApplianceUpdateNetworkApplianceVLANIPv6,
		MandatoryDhcp: requestApplianceUpdateNetworkApplianceVLANMandatoryDhcp,
		Mask:          int64ToIntPointer(mask),
		Name:          name,
		ReservedIPRanges: func() *[]merakigosdk.RequestApplianceUpdateNetworkApplianceVLANReservedIPRanges {
			if len(requestApplianceUpdateNetworkApplianceVLANReservedIPRanges) > 0 {
				return &requestApplianceUpdateNetworkApplianceVLANReservedIPRanges
			}
			return nil
		}(),
		Subnet:           subnet,
		TemplateVLANType: templateVLANType,
		VpnNatSubnet:     vpnNatSubnet,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceVLANItemToBodyRs(state NetworksApplianceVLANsRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceVLAN, is_read bool) NetworksApplianceVLANsRs {
	itemState := NetworksApplianceVLANsRs{
		ApplianceIP:        types.StringValue(response.ApplianceIP),
		Cidr:               state.Cidr,
		DhcpBootFilename:   types.StringValue(response.DhcpBootFilename),
		DhcpBootNextServer: types.StringValue(response.DhcpBootNextServer),
		DhcpBootOptionsEnabled: func() types.Bool {
			if response.DhcpBootOptionsEnabled != nil {
				return types.BoolValue(*response.DhcpBootOptionsEnabled)
			}
			return types.BoolNull()
		}(),
		DhcpHandling:  types.StringValue(response.DhcpHandling),
		DhcpLeaseTime: types.StringValue(response.DhcpLeaseTime),
		DhcpOptions: func() *[]ResponseApplianceGetNetworkApplianceVlanDhcpOptionsRs {
			if response.DhcpOptions != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVlanDhcpOptionsRs, len(*response.DhcpOptions))
				for i, dhcpOptions := range *response.DhcpOptions {
					result[i] = ResponseApplianceGetNetworkApplianceVlanDhcpOptionsRs{
						Code:  types.StringValue(dhcpOptions.Code),
						Type:  types.StringValue(dhcpOptions.Type),
						Value: types.StringValue(dhcpOptions.Value),
					}
				}
				return &result
			}
			return nil
		}(),
		DhcpRelayServerIPs: StringSliceToSet(response.DhcpRelayServerIPs),
		DNSNameservers:     types.StringValue(response.DNSNameservers),
		GroupPolicyID:      types.StringValue(response.GroupPolicyID),
		ID:                 types.StringValue(strconv.Itoa(*response.ID)),
		InterfaceID:        types.StringValue(response.InterfaceID),
		IPv6: func() *ResponseApplianceGetNetworkApplianceVlanIpv6Rs {
			if response.IPv6 != nil {
				return &ResponseApplianceGetNetworkApplianceVlanIpv6Rs{
					Enabled: func() types.Bool {
						if response.IPv6.Enabled != nil {
							return types.BoolValue(*response.IPv6.Enabled)
						}
						return types.BoolNull()
					}(),
					PrefixAssignments: func() *[]ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsRs {
						if response.IPv6.PrefixAssignments != nil {
							result := make([]ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsRs, len(*response.IPv6.PrefixAssignments))
							for i, prefixAssignments := range *response.IPv6.PrefixAssignments {
								result[i] = ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsRs{
									Autonomous: func() types.Bool {
										if prefixAssignments.Autonomous != nil {
											return types.BoolValue(*prefixAssignments.Autonomous)
										}
										return types.BoolNull()
									}(),
									Origin: func() *ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsOriginRs {
										if prefixAssignments.Origin != nil {
											return &ResponseApplianceGetNetworkApplianceVlanIpv6PrefixAssignmentsOriginRs{
												Interfaces: StringSliceToSet(prefixAssignments.Origin.Interfaces),
												Type:       types.StringValue(prefixAssignments.Origin.Type),
											}
										}
										return nil
									}(),
									StaticApplianceIP6: types.StringValue(prefixAssignments.StaticApplianceIP6),
									StaticPrefix:       types.StringValue(prefixAssignments.StaticPrefix),
								}
							}
							return &result
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
		MandatoryDhcp: func() *ResponseApplianceGetNetworkApplianceVlanMandatoryDhcpRs {
			if response.MandatoryDhcp != nil {
				return &ResponseApplianceGetNetworkApplianceVlanMandatoryDhcpRs{
					Enabled: func() types.Bool {
						if response.MandatoryDhcp.Enabled != nil {
							return types.BoolValue(*response.MandatoryDhcp.Enabled)
						}
						return types.BoolNull()
					}(),
				}
			}
			return nil
		}(),
		Mask: func() types.Int64 {
			if response.Mask != nil {
				return types.Int64Value(int64(*response.Mask))
			}
			if !state.Mask.IsNull() {
				return types.Int64Value(state.Mask.ValueInt64())
			}
			return types.Int64Null()
		}(),
		Name: types.StringValue(response.Name),
		ReservedIPRanges: func() *[]ResponseApplianceGetNetworkApplianceVlanReservedIpRangesRs {
			if response.ReservedIPRanges != nil {
				result := make([]ResponseApplianceGetNetworkApplianceVlanReservedIpRangesRs, len(*response.ReservedIPRanges))
				for i, reservedIPRanges := range *response.ReservedIPRanges {
					result[i] = ResponseApplianceGetNetworkApplianceVlanReservedIpRangesRs{
						Comment: types.StringValue(reservedIPRanges.Comment),
						End:     types.StringValue(reservedIPRanges.End),
						Start:   types.StringValue(reservedIPRanges.Start),
					}
				}
				return &result
			}
			return nil
		}(),
		Subnet:           types.StringValue(response.Subnet),
		TemplateVLANType: types.StringValue(response.TemplateVLANType),
		VpnNatSubnet:     types.StringValue(response.VpnNatSubnet),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceVLANsRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceVLANsRs)
}
