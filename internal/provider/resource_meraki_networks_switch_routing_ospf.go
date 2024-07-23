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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchRoutingOspfResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchRoutingOspfResource{}
)

func NewNetworksSwitchRoutingOspfResource() resource.Resource {
	return &NetworksSwitchRoutingOspfResource{}
}

type NetworksSwitchRoutingOspfResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchRoutingOspfResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchRoutingOspfResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_routing_ospf"
}

func (r *NetworksSwitchRoutingOspfResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"areas": schema.SetNestedAttribute{
				MarkdownDescription: `OSPF areas`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"area_id": schema.Int64Attribute{
							MarkdownDescription: `OSPF area ID`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Int64{
								int64planmodifier.UseStateForUnknown(),
							},
						},
						"area_name": schema.StringAttribute{
							MarkdownDescription: `Name of the OSPF area`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
						},
						"area_type": schema.StringAttribute{
							MarkdownDescription: `Area types in OSPF. Must be one of: ["normal", "stub", "nssa"]`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"normal",
									"nssa",
									"stub",
								),
							},
						},
					},
				},
			},
			"dead_timer_in_seconds": schema.Int64Attribute{
				MarkdownDescription: `Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean value to enable or disable OSPF routing. OSPF routing is disabled by default.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"hello_timer_in_seconds": schema.Int64Attribute{
				MarkdownDescription: `Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"md5_authentication_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean value to enable or disable MD5 authentication. MD5 authentication is disabled by default.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"md5_authentication_key": schema.SingleNestedAttribute{
				MarkdownDescription: `MD5 authentication credentials. This param is only relevant if md5AuthenticationEnabled is true`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"id": schema.Int64Attribute{
						MarkdownDescription: `MD5 authentication key index. Key index must be between 1 to 255`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"passphrase": schema.StringAttribute{
						MarkdownDescription: `MD5 authentication passphrase`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
			},
			"v3": schema.SingleNestedAttribute{
				MarkdownDescription: `OSPF v3 configuration`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"areas": schema.SetNestedAttribute{
						MarkdownDescription: `OSPF v3 areas`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"area_id": schema.Int64Attribute{
									MarkdownDescription: `OSPF area ID`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Int64{
										int64planmodifier.UseStateForUnknown(),
									},
								},
								"area_name": schema.StringAttribute{
									MarkdownDescription: `Name of the OSPF area`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
								},
								"area_type": schema.StringAttribute{
									MarkdownDescription: `Area types in OSPF. Must be one of: ["normal", "stub", "nssa"]`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.UseStateForUnknown(),
									},
									Validators: []validator.String{
										stringvalidator.OneOf(
											"normal",
											"nssa",
											"stub",
										),
									},
								},
							},
						},
					},
					"dead_timer_in_seconds": schema.Int64Attribute{
						MarkdownDescription: `Time interval to determine when the peer will be declared inactive/dead. Value must be between 1 and 65535`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean value to enable or disable V3 OSPF routing. OSPF V3 routing is disabled by default.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"hello_timer_in_seconds": schema.Int64Attribute{
						MarkdownDescription: `Time interval in seconds at which hello packet will be sent to OSPF neighbors to maintain connectivity. Value must be between 1 and 255. Default is 10 seconds.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
		},
	}
}

func (r *NetworksSwitchRoutingOspfResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchRoutingOspfRs

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
	responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchRoutingOspf(vvNetworkID)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchRoutingOspf only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource NetworksSwitchRoutingOspf only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Switch.UpdateNetworkSwitchRoutingOspf(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchRoutingOspf",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchRoutingOspf",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchRoutingOspf(vvNetworkID)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchRoutingOspf",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchRoutingOspf",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchRoutingOspfItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchRoutingOspfResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchRoutingOspfRs

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
	responseGet, restyRespGet, err := r.client.Switch.GetNetworkSwitchRoutingOspf(vvNetworkID)
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
				"Failure when executing GetNetworkSwitchRoutingOspf",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchRoutingOspf",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseSwitchGetNetworkSwitchRoutingOspfItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksSwitchRoutingOspfResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchRoutingOspfResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchRoutingOspfRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Switch.UpdateNetworkSwitchRoutingOspf(vvNetworkID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchRoutingOspf",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchRoutingOspf",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchRoutingOspfResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting NetworksSwitchRoutingOspf", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksSwitchRoutingOspfRs struct {
	NetworkID                types.String                                                     `tfsdk:"network_id"`
	Areas                    *[]ResponseSwitchGetNetworkSwitchRoutingOspfAreasRs              `tfsdk:"areas"`
	DeadTimerInSeconds       types.Int64                                                      `tfsdk:"dead_timer_in_seconds"`
	Enabled                  types.Bool                                                       `tfsdk:"enabled"`
	HelloTimerInSeconds      types.Int64                                                      `tfsdk:"hello_timer_in_seconds"`
	Md5AuthenticationEnabled types.Bool                                                       `tfsdk:"md5_authentication_enabled"`
	Md5AuthenticationKey     *ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKeyRs `tfsdk:"md5_authentication_key"`
	V3                       *ResponseSwitchGetNetworkSwitchRoutingOspfV3Rs                   `tfsdk:"v3"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspfAreasRs struct {
	AreaID   types.Int64  `tfsdk:"area_id"`
	AreaName types.String `tfsdk:"area_name"`
	AreaType types.String `tfsdk:"area_type"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKeyRs struct {
	ID         types.Int64  `tfsdk:"id"`
	Passphrase types.String `tfsdk:"passphrase"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspfV3Rs struct {
	Areas               *[]ResponseSwitchGetNetworkSwitchRoutingOspfV3AreasRs `tfsdk:"areas"`
	DeadTimerInSeconds  types.Int64                                           `tfsdk:"dead_timer_in_seconds"`
	Enabled             types.Bool                                            `tfsdk:"enabled"`
	HelloTimerInSeconds types.Int64                                           `tfsdk:"hello_timer_in_seconds"`
}

type ResponseSwitchGetNetworkSwitchRoutingOspfV3AreasRs struct {
	AreaID   types.Int64  `tfsdk:"area_id"`
	AreaName types.String `tfsdk:"area_name"`
	AreaType types.String `tfsdk:"area_type"`
}

// FromBody
func (r *NetworksSwitchRoutingOspfRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspf {
	var requestSwitchUpdateNetworkSwitchRoutingOspfAreas []merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfAreas
	if r.Areas != nil {
		for _, rItem1 := range *r.Areas {
			areaID := rItem1.AreaID.ValueInt64()
			areaName := rItem1.AreaName.ValueString()
			areaType := rItem1.AreaType.ValueString()
			requestSwitchUpdateNetworkSwitchRoutingOspfAreas = append(requestSwitchUpdateNetworkSwitchRoutingOspfAreas, merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfAreas{
				AreaID:   int64ToString(&areaID),
				AreaName: areaName,
				AreaType: areaType,
			})
		}
	}
	deadTimerInSeconds := new(int64)
	if !r.DeadTimerInSeconds.IsUnknown() && !r.DeadTimerInSeconds.IsNull() {
		*deadTimerInSeconds = r.DeadTimerInSeconds.ValueInt64()
	} else {
		deadTimerInSeconds = nil
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	helloTimerInSeconds := new(int64)
	if !r.HelloTimerInSeconds.IsUnknown() && !r.HelloTimerInSeconds.IsNull() {
		*helloTimerInSeconds = r.HelloTimerInSeconds.ValueInt64()
	} else {
		helloTimerInSeconds = nil
	}
	md5AuthenticationEnabled := new(bool)
	if !r.Md5AuthenticationEnabled.IsUnknown() && !r.Md5AuthenticationEnabled.IsNull() {
		*md5AuthenticationEnabled = r.Md5AuthenticationEnabled.ValueBool()
	} else {
		md5AuthenticationEnabled = nil
	}
	var requestSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey *merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey
	if r.Md5AuthenticationKey != nil {
		iD := func() *int64 {
			if !r.Md5AuthenticationKey.ID.IsUnknown() && !r.Md5AuthenticationKey.ID.IsNull() {
				return r.Md5AuthenticationKey.ID.ValueInt64Pointer()
			}
			return nil
		}()
		passphrase := r.Md5AuthenticationKey.Passphrase.ValueString()
		requestSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey = &merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey{
			ID:         int64ToIntPointer(iD),
			Passphrase: passphrase,
		}
	}
	var requestSwitchUpdateNetworkSwitchRoutingOspfV3 *merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfV3
	if r.V3 != nil {
		var requestSwitchUpdateNetworkSwitchRoutingOspfV3Areas []merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfV3Areas
		if r.V3.Areas != nil {
			for _, rItem1 := range *r.V3.Areas { //V3.Areas// name: areas
				areaID := rItem1.AreaID.ValueInt64()
				areaName := rItem1.AreaName.ValueString()
				areaType := rItem1.AreaType.ValueString()
				requestSwitchUpdateNetworkSwitchRoutingOspfV3Areas = append(requestSwitchUpdateNetworkSwitchRoutingOspfV3Areas, merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfV3Areas{
					AreaID:   int64ToString(&areaID),
					AreaName: areaName,
					AreaType: areaType,
				})
			}
		}
		deadTimerInSeconds := func() *int64 {
			if !r.DeadTimerInSeconds.IsUnknown() && !r.DeadTimerInSeconds.IsNull() {
				return r.DeadTimerInSeconds.ValueInt64Pointer()
			}
			return nil
		}()
		enabled := func() *bool {
			if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
				return r.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		helloTimerInSeconds := func() *int64 {
			if !r.HelloTimerInSeconds.IsUnknown() && !r.HelloTimerInSeconds.IsNull() {
				return r.HelloTimerInSeconds.ValueInt64Pointer()
			}
			return nil
		}()
		requestSwitchUpdateNetworkSwitchRoutingOspfV3 = &merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfV3{
			Areas: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfV3Areas {
				if len(requestSwitchUpdateNetworkSwitchRoutingOspfV3Areas) > 0 {
					return &requestSwitchUpdateNetworkSwitchRoutingOspfV3Areas
				}
				return nil
			}(),
			DeadTimerInSeconds:  int64ToIntPointer(deadTimerInSeconds),
			Enabled:             enabled,
			HelloTimerInSeconds: int64ToIntPointer(helloTimerInSeconds),
		}
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspf{
		Areas: func() *[]merakigosdk.RequestSwitchUpdateNetworkSwitchRoutingOspfAreas {
			if len(requestSwitchUpdateNetworkSwitchRoutingOspfAreas) > 0 {
				return &requestSwitchUpdateNetworkSwitchRoutingOspfAreas
			}
			return nil
		}(),
		DeadTimerInSeconds:       int64ToIntPointer(deadTimerInSeconds),
		Enabled:                  enabled,
		HelloTimerInSeconds:      int64ToIntPointer(helloTimerInSeconds),
		Md5AuthenticationEnabled: md5AuthenticationEnabled,
		Md5AuthenticationKey:     requestSwitchUpdateNetworkSwitchRoutingOspfMd5AuthenticationKey,
		V3:                       requestSwitchUpdateNetworkSwitchRoutingOspfV3,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchRoutingOspfItemToBodyRs(state NetworksSwitchRoutingOspfRs, response *merakigosdk.ResponseSwitchGetNetworkSwitchRoutingOspf, is_read bool) NetworksSwitchRoutingOspfRs {
	itemState := NetworksSwitchRoutingOspfRs{
		Areas: func() *[]ResponseSwitchGetNetworkSwitchRoutingOspfAreasRs {
			if response.Areas != nil {
				result := make([]ResponseSwitchGetNetworkSwitchRoutingOspfAreasRs, len(*response.Areas))
				for i, areas := range *response.Areas {
					result[i] = ResponseSwitchGetNetworkSwitchRoutingOspfAreasRs{
						AreaID:   types.Int64Value(int64(*areas.AreaID)),
						AreaName: types.StringValue(areas.AreaName),
						AreaType: types.StringValue(areas.AreaType),
					}
				}
				return &result
			}
			return nil
		}(),
		DeadTimerInSeconds: func() types.Int64 {
			if response.DeadTimerInSeconds != nil {
				return types.Int64Value(int64(*response.DeadTimerInSeconds))
			}
			return types.Int64{}
		}(),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		HelloTimerInSeconds: func() types.Int64 {
			if response.HelloTimerInSeconds != nil {
				return types.Int64Value(int64(*response.HelloTimerInSeconds))
			}
			return types.Int64{}
		}(),
		Md5AuthenticationEnabled: func() types.Bool {
			if response.Md5AuthenticationEnabled != nil {
				return types.BoolValue(*response.Md5AuthenticationEnabled)
			}
			return types.Bool{}
		}(),
		Md5AuthenticationKey: func() *ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKeyRs {
			if response.Md5AuthenticationKey != nil {
				return &ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKeyRs{
					ID: func() types.Int64 {
						if response.Md5AuthenticationKey.ID != nil {
							return types.Int64Value(int64(*response.Md5AuthenticationKey.ID))
						}
						return types.Int64{}
					}(),
					Passphrase: types.StringValue(response.Md5AuthenticationKey.Passphrase),
				}
			}
			return &ResponseSwitchGetNetworkSwitchRoutingOspfMd5AuthenticationKeyRs{}
		}(),
		V3: func() *ResponseSwitchGetNetworkSwitchRoutingOspfV3Rs {
			if response.V3 != nil {
				return &ResponseSwitchGetNetworkSwitchRoutingOspfV3Rs{
					Areas: func() *[]ResponseSwitchGetNetworkSwitchRoutingOspfV3AreasRs {
						if response.V3.Areas != nil {
							result := make([]ResponseSwitchGetNetworkSwitchRoutingOspfV3AreasRs, len(*response.V3.Areas))
							for i, areas := range *response.V3.Areas {
								result[i] = ResponseSwitchGetNetworkSwitchRoutingOspfV3AreasRs{
									AreaID:   types.Int64Value(int64(*areas.AreaID)),
									AreaName: types.StringValue(areas.AreaName),
									AreaType: types.StringValue(areas.AreaType),
								}
							}
							return &result
						}
						return nil
					}(),
					DeadTimerInSeconds: func() types.Int64 {
						if response.V3.DeadTimerInSeconds != nil {
							return types.Int64Value(int64(*response.V3.DeadTimerInSeconds))
						}
						return types.Int64{}
					}(),
					Enabled: func() types.Bool {
						if response.V3.Enabled != nil {
							return types.BoolValue(*response.V3.Enabled)
						}
						return types.Bool{}
					}(),
					HelloTimerInSeconds: func() types.Int64 {
						if response.V3.HelloTimerInSeconds != nil {
							return types.Int64Value(int64(*response.V3.HelloTimerInSeconds))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseSwitchGetNetworkSwitchRoutingOspfV3Rs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksSwitchRoutingOspfRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksSwitchRoutingOspfRs)
}
