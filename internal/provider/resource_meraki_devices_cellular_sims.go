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
	_ resource.Resource              = &DevicesCellularSimsResource{}
	_ resource.ResourceWithConfigure = &DevicesCellularSimsResource{}
)

func NewDevicesCellularSimsResource() resource.Resource {
	return &DevicesCellularSimsResource{}
}

type DevicesCellularSimsResource struct {
	client *merakigosdk.Client
}

func (r *DevicesCellularSimsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesCellularSimsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_cellular_sims"
}

func (r *DevicesCellularSimsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
			"sim_failover": schema.SingleNestedAttribute{
				MarkdownDescription: `SIM Failover settings.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Failover to secondary SIM (optional)`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"timeout": schema.Int64Attribute{
						MarkdownDescription: `Failover timeout in seconds (optional)`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"sims": schema.SetNestedAttribute{
				MarkdownDescription: `List of SIMs. If a SIM was previously configured and not specified in this request, it will remain unchanged.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"apns": schema.SetNestedAttribute{
							MarkdownDescription: `APN configurations. If empty, the default APN will be used.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Set{
								setplanmodifier.UseStateForUnknown(),
							},
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{

									"allowed_ip_types": schema.SetAttribute{
										MarkdownDescription: `IP versions to support (permitted values include 'ipv4', 'ipv6').`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
									"authentication": schema.SingleNestedAttribute{
										MarkdownDescription: `APN authentication configurations.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.UseStateForUnknown(),
										},
										Attributes: map[string]schema.Attribute{

											"password": schema.StringAttribute{
												MarkdownDescription: `APN password, if type is set (if APN password is not supplied, the password is left unchanged).`,
												Sensitive:           true,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
											"type": schema.StringAttribute{
												MarkdownDescription: `APN auth type.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
												Validators: []validator.String{
													stringvalidator.OneOf(
														"chap",
														"none",
														"pap",
													),
												},
											},
											"username": schema.StringAttribute{
												MarkdownDescription: `APN username, if type is set.`,
												Computed:            true,
												Optional:            true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.UseStateForUnknown(),
												},
											},
										},
									},
									"name": schema.StringAttribute{
										MarkdownDescription: `APN name.`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.UseStateForUnknown(),
										},
									},
								},
							},
						},
						"is_primary": schema.BoolAttribute{
							MarkdownDescription: `If true, this SIM is used for boot. Must be true on single-sim devices.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.UseStateForUnknown(),
							},
						},
						"slot": schema.StringAttribute{
							MarkdownDescription: `SIM slot being configured. Must be 'sim1' on single-sim devices.`,
							Computed:            true,
							Optional:            true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.UseStateForUnknown(),
							},
							Validators: []validator.String{
								stringvalidator.OneOf(
									"sim1",
									"sim2",
								),
							},
						},
					},
				},
			},
		},
	}
}

func (r *DevicesCellularSimsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesCellularSimsRs

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
	vvSerial := data.Serial.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Devices.GetDeviceCellularSims(vvSerial)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesCellularSims only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesCellularSims only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Devices.UpdateDeviceCellularSims(vvSerial, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCellularSims",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCellularSims",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Devices.GetDeviceCellularSims(vvSerial)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCellularSims",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCellularSims",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseDevicesGetDeviceCellularSimsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCellularSimsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesCellularSimsRs

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

	vvSerial := data.Serial.ValueString()
	responseGet, restyRespGet, err := r.client.Devices.GetDeviceCellularSims(vvSerial)
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
				"Failure when executing GetDeviceCellularSims",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCellularSims",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseDevicesGetDeviceCellularSimsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesCellularSimsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesCellularSimsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesCellularSimsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Devices.UpdateDeviceCellularSims(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCellularSims",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCellularSims",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCellularSimsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesCellularSims", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesCellularSimsRs struct {
	Serial      types.String                                         `tfsdk:"serial"`
	Sims        *[]ResponseDevicesGetDeviceCellularSimsSimsRs        `tfsdk:"sims"`
	SimFailover *RequestDevicesUpdateDeviceCellularSimsSimFailoverRs `tfsdk:"sim_failover"`
}

type ResponseDevicesGetDeviceCellularSimsSimsRs struct {
	Apns      *[]ResponseDevicesGetDeviceCellularSimsSimsApnsRs `tfsdk:"apns"`
	IsPrimary types.Bool                                        `tfsdk:"is_primary"`
	Slot      types.String                                      `tfsdk:"slot"`
}

type ResponseDevicesGetDeviceCellularSimsSimsApnsRs struct {
	AllowedIPTypes types.Set                                                     `tfsdk:"allowed_ip_types"`
	Authentication *ResponseDevicesGetDeviceCellularSimsSimsApnsAuthenticationRs `tfsdk:"authentication"`
	Name           types.String                                                  `tfsdk:"name"`
}

type ResponseDevicesGetDeviceCellularSimsSimsApnsAuthenticationRs struct {
	Type     types.String `tfsdk:"type"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

type RequestDevicesUpdateDeviceCellularSimsSimFailoverRs struct {
	Enabled types.Bool  `tfsdk:"enabled"`
	Timeout types.Int64 `tfsdk:"timeout"`
}

// FromBody
func (r *DevicesCellularSimsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestDevicesUpdateDeviceCellularSims {
	var requestDevicesUpdateDeviceCellularSimsSimFailover *merakigosdk.RequestDevicesUpdateDeviceCellularSimsSimFailover
	if r.SimFailover != nil {
		enabled := func() *bool {
			if !r.SimFailover.Enabled.IsUnknown() && !r.SimFailover.Enabled.IsNull() {
				return r.SimFailover.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		timeout := func() *int64 {
			if !r.SimFailover.Timeout.IsUnknown() && !r.SimFailover.Timeout.IsNull() {
				return r.SimFailover.Timeout.ValueInt64Pointer()
			}
			return nil
		}()
		requestDevicesUpdateDeviceCellularSimsSimFailover = &merakigosdk.RequestDevicesUpdateDeviceCellularSimsSimFailover{
			Enabled: enabled,
			Timeout: int64ToIntPointer(timeout),
		}
	}
	var requestDevicesUpdateDeviceCellularSimsSims []merakigosdk.RequestDevicesUpdateDeviceCellularSimsSims
	if r.Sims != nil {
		for _, rItem1 := range *r.Sims {
			var requestDevicesUpdateDeviceCellularSimsSimsApns []merakigosdk.RequestDevicesUpdateDeviceCellularSimsSimsApns
			if rItem1.Apns != nil {
				for _, rItem2 := range *rItem1.Apns { //Apns// name: apns
					var allowedIPTypes []string = nil
					//Hoola aqui
					rItem2.AllowedIPTypes.ElementsAs(ctx, &allowedIPTypes, false)
					var requestDevicesUpdateDeviceCellularSimsSimsApnsAuthentication *merakigosdk.RequestDevicesUpdateDeviceCellularSimsSimsApnsAuthentication
					if rItem2.Authentication != nil {
						password := rItem2.Authentication.Password.ValueString()
						typeR := rItem2.Authentication.Type.ValueString()
						username := rItem2.Authentication.Username.ValueString()
						requestDevicesUpdateDeviceCellularSimsSimsApnsAuthentication = &merakigosdk.RequestDevicesUpdateDeviceCellularSimsSimsApnsAuthentication{
							Password: password,
							Type:     typeR,
							Username: username,
						}
					}
					name := rItem2.Name.ValueString()
					requestDevicesUpdateDeviceCellularSimsSimsApns = append(requestDevicesUpdateDeviceCellularSimsSimsApns, merakigosdk.RequestDevicesUpdateDeviceCellularSimsSimsApns{
						AllowedIPTypes: allowedIPTypes,
						Authentication: requestDevicesUpdateDeviceCellularSimsSimsApnsAuthentication,
						Name:           name,
					})
				}
			}
			isPrimary := func() *bool {
				if !rItem1.IsPrimary.IsUnknown() && !rItem1.IsPrimary.IsNull() {
					return rItem1.IsPrimary.ValueBoolPointer()
				}
				return nil
			}()
			slot := rItem1.Slot.ValueString()
			requestDevicesUpdateDeviceCellularSimsSims = append(requestDevicesUpdateDeviceCellularSimsSims, merakigosdk.RequestDevicesUpdateDeviceCellularSimsSims{
				Apns: func() *[]merakigosdk.RequestDevicesUpdateDeviceCellularSimsSimsApns {
					if len(requestDevicesUpdateDeviceCellularSimsSimsApns) > 0 {
						return &requestDevicesUpdateDeviceCellularSimsSimsApns
					}
					return nil
				}(),
				IsPrimary: isPrimary,
				Slot:      slot,
			})
		}
	}
	out := merakigosdk.RequestDevicesUpdateDeviceCellularSims{
		SimFailover: requestDevicesUpdateDeviceCellularSimsSimFailover,
		Sims: func() *[]merakigosdk.RequestDevicesUpdateDeviceCellularSimsSims {
			if len(requestDevicesUpdateDeviceCellularSimsSims) > 0 {
				return &requestDevicesUpdateDeviceCellularSimsSims
			}
			return nil
		}(),
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseDevicesGetDeviceCellularSimsItemToBodyRs(state DevicesCellularSimsRs, response *merakigosdk.ResponseDevicesGetDeviceCellularSims, is_read bool) DevicesCellularSimsRs {
	itemState := DevicesCellularSimsRs{
		Sims: func() *[]ResponseDevicesGetDeviceCellularSimsSimsRs {
			if response.Sims != nil {
				result := make([]ResponseDevicesGetDeviceCellularSimsSimsRs, len(*response.Sims))
				for i, sims := range *response.Sims {
					result[i] = ResponseDevicesGetDeviceCellularSimsSimsRs{
						Apns: func() *[]ResponseDevicesGetDeviceCellularSimsSimsApnsRs {
							if sims.Apns != nil {
								result := make([]ResponseDevicesGetDeviceCellularSimsSimsApnsRs, len(*sims.Apns))
								for i, apns := range *sims.Apns {
									result[i] = ResponseDevicesGetDeviceCellularSimsSimsApnsRs{
										AllowedIPTypes: StringSliceToSet(apns.AllowedIPTypes),
										Authentication: func() *ResponseDevicesGetDeviceCellularSimsSimsApnsAuthenticationRs {
											if apns.Authentication != nil {
												return &ResponseDevicesGetDeviceCellularSimsSimsApnsAuthenticationRs{
													Type:     types.StringValue(apns.Authentication.Type),
													Username: types.StringValue(apns.Authentication.Username),
												}
											}
											return &ResponseDevicesGetDeviceCellularSimsSimsApnsAuthenticationRs{}
										}(),
										Name: types.StringValue(apns.Name),
									}
								}
								return &result
							}
							return &[]ResponseDevicesGetDeviceCellularSimsSimsApnsRs{}
						}(),
						IsPrimary: func() types.Bool {
							if sims.IsPrimary != nil {
								return types.BoolValue(*sims.IsPrimary)
							}
							return types.Bool{}
						}(),
						Slot: types.StringValue(sims.Slot),
					}
				}
				return &result
			}
			return &[]ResponseDevicesGetDeviceCellularSimsSimsRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesCellularSimsRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesCellularSimsRs)
}
