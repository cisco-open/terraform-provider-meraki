package provider

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksApplianceTrafficShapingResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceTrafficShapingResource{}
)

func NewNetworksApplianceTrafficShapingResource() resource.Resource {
	return &NetworksApplianceTrafficShapingResource{}
}

type NetworksApplianceTrafficShapingResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceTrafficShapingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceTrafficShapingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_traffic_shaping"
}

func (r *NetworksApplianceTrafficShapingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"global_bandwidth_limits": schema.SingleNestedAttribute{
				MarkdownDescription: `Global per-client bandwidth limit`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"limit_down": schema.Int64Attribute{
						MarkdownDescription: `The download bandwidth limit in Kbps. (0 represents no limit.)`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"limit_up": schema.Int64Attribute{
						MarkdownDescription: `The upload bandwidth limit in Kbps. (0 represents no limit.)`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
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

func (r *NetworksApplianceTrafficShapingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceTrafficShapingRs

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
	responseVerifyItem, restyResp1, err := r.client.Appliance.GetNetworkApplianceTrafficShaping(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceTrafficShaping only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksApplianceTrafficShaping only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceTrafficShaping(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceTrafficShaping",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceTrafficShaping",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Appliance.GetNetworkApplianceTrafficShaping(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkApplianceTrafficShaping",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceTrafficShaping",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceTrafficShapingItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceTrafficShapingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksApplianceTrafficShapingRs

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
	responseGet, restyRespGet, err := r.client.Appliance.GetNetworkApplianceTrafficShaping(vvNetworkID)
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
				"Failure when executing GetNetworkApplianceTrafficShaping",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkApplianceTrafficShaping",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseApplianceGetNetworkApplianceTrafficShapingItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksApplianceTrafficShapingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksApplianceTrafficShapingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksApplianceTrafficShapingRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Appliance.UpdateNetworkApplianceTrafficShaping(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkApplianceTrafficShaping",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkApplianceTrafficShaping",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceTrafficShapingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksApplianceTrafficShaping", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceTrafficShapingRs struct {
	NetworkID             types.String                                                               `tfsdk:"network_id"`
	GlobalBandwidthLimits *ResponseApplianceGetNetworkApplianceTrafficShapingGlobalBandwidthLimitsRs `tfsdk:"global_bandwidth_limits"`
}

type ResponseApplianceGetNetworkApplianceTrafficShapingGlobalBandwidthLimitsRs struct {
	LimitDown types.Int64 `tfsdk:"limit_down"`
	LimitUp   types.Int64 `tfsdk:"limit_up"`
}

// FromBody
func (r *NetworksApplianceTrafficShapingRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShaping {
	var requestApplianceUpdateNetworkApplianceTrafficShapingGlobalBandwidthLimits *merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingGlobalBandwidthLimits
	if r.GlobalBandwidthLimits != nil {
		limitDown := func() *int64 {
			if !r.GlobalBandwidthLimits.LimitDown.IsUnknown() && !r.GlobalBandwidthLimits.LimitDown.IsNull() {
				return r.GlobalBandwidthLimits.LimitDown.ValueInt64Pointer()
			}
			return nil
		}()
		limitUp := func() *int64 {
			if !r.GlobalBandwidthLimits.LimitUp.IsUnknown() && !r.GlobalBandwidthLimits.LimitUp.IsNull() {
				return r.GlobalBandwidthLimits.LimitUp.ValueInt64Pointer()
			}
			return nil
		}()
		requestApplianceUpdateNetworkApplianceTrafficShapingGlobalBandwidthLimits = &merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShapingGlobalBandwidthLimits{
			LimitDown: int64ToIntPointer(limitDown),
			LimitUp:   int64ToIntPointer(limitUp),
		}
	}
	out := merakigosdk.RequestApplianceUpdateNetworkApplianceTrafficShaping{
		GlobalBandwidthLimits: requestApplianceUpdateNetworkApplianceTrafficShapingGlobalBandwidthLimits,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseApplianceGetNetworkApplianceTrafficShapingItemToBodyRs(state NetworksApplianceTrafficShapingRs, response *merakigosdk.ResponseApplianceGetNetworkApplianceTrafficShaping, is_read bool) NetworksApplianceTrafficShapingRs {
	itemState := NetworksApplianceTrafficShapingRs{
		GlobalBandwidthLimits: func() *ResponseApplianceGetNetworkApplianceTrafficShapingGlobalBandwidthLimitsRs {
			if response.GlobalBandwidthLimits != nil {
				return &ResponseApplianceGetNetworkApplianceTrafficShapingGlobalBandwidthLimitsRs{
					LimitDown: func() types.Int64 {
						if response.GlobalBandwidthLimits.LimitDown != nil {
							return types.Int64Value(int64(*response.GlobalBandwidthLimits.LimitDown))
						}
						return types.Int64{}
					}(),
					LimitUp: func() types.Int64 {
						if response.GlobalBandwidthLimits.LimitUp != nil {
							return types.Int64Value(int64(*response.GlobalBandwidthLimits.LimitUp))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseApplianceGetNetworkApplianceTrafficShapingGlobalBandwidthLimitsRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksApplianceTrafficShapingRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksApplianceTrafficShapingRs)
}
