package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksWirelessSSIDsBonjourForwardingResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessSSIDsBonjourForwardingResource{}
)

func NewNetworksWirelessSSIDsBonjourForwardingResource() resource.Resource {
	return &NetworksWirelessSSIDsBonjourForwardingResource{}
}

type NetworksWirelessSSIDsBonjourForwardingResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessSSIDsBonjourForwardingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessSSIDsBonjourForwardingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_ssids_bonjour_forwarding"
}

func (r *NetworksWirelessSSIDsBonjourForwardingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `If true, Bonjour forwarding is enabled on the SSID.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"exception": schema.SingleNestedAttribute{
				MarkdownDescription: `Bonjour forwarding exception`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `If true, Bonjour forwarding exception is enabled on this SSID. Exception is required to enable L2 isolation and Bonjour forwarding to work together.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
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
				MarkdownDescription: `Bonjour forwarding rules`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"description": schema.StringAttribute{
							MarkdownDescription: `Desctiption of the bonjour forwarding rule`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"services": schema.SetAttribute{
							MarkdownDescription: `A list of Bonjour services. At least one service must be specified. Available services are 'All Services', 'AirPlay', 'AFP', 'BitTorrent', 'FTP', 'iChat', 'iTunes', 'Printers', 'Samba', 'Scanners' and 'SSH'`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},

							ElementType: types.StringType,
						},
						"vlan_id": schema.StringAttribute{
							MarkdownDescription: `The ID of the service VLAN. Required`,
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

func (r *NetworksWirelessSSIDsBonjourForwardingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessSSIDsBonjourForwardingRs

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
	responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDBonjourForwarding(vvNetworkID, vvNumber)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDsBonjourForwarding only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksWirelessSSIDsBonjourForwarding only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDBonjourForwarding(vvNetworkID, vvNumber, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDBonjourForwarding",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDBonjourForwarding",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessSSIDBonjourForwarding(vvNetworkID, vvNumber)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessSSIDBonjourForwarding",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDBonjourForwarding",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsBonjourForwardingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessSSIDsBonjourForwardingRs

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
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessSSIDBonjourForwarding(vvNetworkID, vvNumber)
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
				"Failure when executing GetNetworkWirelessSSIDBonjourForwarding",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessSSIDBonjourForwarding",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessSSIDsBonjourForwardingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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

func (r *NetworksWirelessSSIDsBonjourForwardingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessSSIDsBonjourForwardingRs
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
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessSSIDBonjourForwarding(vvNetworkID, vvNumber, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessSSIDBonjourForwarding",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessSSIDBonjourForwarding",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessSSIDsBonjourForwardingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksWirelessSSIDsBonjourForwarding", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksWirelessSSIDsBonjourForwardingRs struct {
	NetworkID types.String                                                        `tfsdk:"network_id"`
	Number    types.String                                                        `tfsdk:"number"`
	Enabled   types.Bool                                                          `tfsdk:"enabled"`
	Exception *ResponseWirelessGetNetworkWirelessSsidBonjourForwardingExceptionRs `tfsdk:"exception"`
	Rules     *[]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRulesRs   `tfsdk:"rules"`
}

type ResponseWirelessGetNetworkWirelessSsidBonjourForwardingExceptionRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRulesRs struct {
	Description types.String `tfsdk:"description"`
	Services    types.Set    `tfsdk:"services"`
	VLANID      types.String `tfsdk:"vlan_id"`
}

// FromBody
func (r *NetworksWirelessSSIDsBonjourForwardingRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDBonjourForwarding {
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingException *merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingException
	if r.Exception != nil {
		enabled := func() *bool {
			if !r.Exception.Enabled.IsUnknown() && !r.Exception.Enabled.IsNull() {
				return r.Exception.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingException = &merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingException{
			Enabled: enabled,
		}
	}
	var requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules []merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules
	if r.Rules != nil {
		for _, rItem1 := range *r.Rules {
			description := rItem1.Description.ValueString()
			var services []string = nil
			//Hoola aqui
			rItem1.Services.ElementsAs(ctx, &services, false)
			vLANID := rItem1.VLANID.ValueString()
			requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules = append(requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules, merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules{
				Description: description,
				Services:    services,
				VLANID:      vLANID,
			})
		}
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDBonjourForwarding{
		Enabled:   enabled,
		Exception: requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingException,
		Rules: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules {
			if len(requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules) > 0 {
				return &requestWirelessUpdateNetworkWirelessSSIDBonjourForwardingRules
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessSSIDBonjourForwardingItemToBodyRs(state NetworksWirelessSSIDsBonjourForwardingRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessSSIDBonjourForwarding, is_read bool) NetworksWirelessSSIDsBonjourForwardingRs {
	itemState := NetworksWirelessSSIDsBonjourForwardingRs{
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		Exception: func() *ResponseWirelessGetNetworkWirelessSsidBonjourForwardingExceptionRs {
			if response.Exception != nil {
				return &ResponseWirelessGetNetworkWirelessSsidBonjourForwardingExceptionRs{
					Enabled: func() types.Bool {
						if response.Exception.Enabled != nil {
							return types.BoolValue(*response.Exception.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseWirelessGetNetworkWirelessSsidBonjourForwardingExceptionRs{}
		}(),
		Rules: func() *[]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRulesRs {
			if response.Rules != nil {
				result := make([]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRulesRs, len(*response.Rules))
				for i, rules := range *response.Rules {
					result[i] = ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRulesRs{
						Description: types.StringValue(rules.Description),
						Services:    StringSliceToSet(rules.Services),
						VLANID:      types.StringValue(rules.VLANID),
					}
				}
				return &result
			}
			return &[]ResponseWirelessGetNetworkWirelessSsidBonjourForwardingRulesRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessSSIDsBonjourForwardingRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessSSIDsBonjourForwardingRs)
}
