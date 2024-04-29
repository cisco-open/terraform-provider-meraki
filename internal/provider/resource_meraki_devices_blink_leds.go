package provider

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesBlinkLedsResource{}
	_ resource.ResourceWithConfigure = &DevicesBlinkLedsResource{}
)

func NewDevicesBlinkLedsResource() resource.Resource {
	return &DevicesBlinkLedsResource{}
}

type DevicesBlinkLedsResource struct {
	client *merakigosdk.Client
}

func (r *DevicesBlinkLedsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesBlinkLedsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_blink_leds"
}

// resourceAction
func (r *DevicesBlinkLedsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"duration": schema.Int64Attribute{
						MarkdownDescription: `The duration in seconds. Will be between 5 and 120. Default is 20 seconds`,
						Computed:            true,
					},
					"duty": schema.Int64Attribute{
						MarkdownDescription: `The duty cycle as the percent active. Will be between 10 and 90. Default is 50`,
						Computed:            true,
					},
					"period": schema.Int64Attribute{
						MarkdownDescription: `The period in milliseconds. Will be between 100 and 1000. Default is 160 milliseconds`,
						Computed:            true,
					},
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"duration": schema.Int64Attribute{
						MarkdownDescription: `The duration in seconds. Must be between 5 and 120. Default is 20 seconds`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"duty": schema.Int64Attribute{
						MarkdownDescription: `The duty cycle as the percent active. Must be between 10 and 90. Default is 50.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"period": schema.Int64Attribute{
						MarkdownDescription: `The period in milliseconds. Must be between 100 and 1000. Default is 160 milliseconds`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *DevicesBlinkLedsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesBlinkLeds

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp1, err := r.client.Devices.BlinkDeviceLeds(vvSerial, dataRequest)

	if err != nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing BlinkDeviceLeds",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing BlinkDeviceLeds",
			err.Error(),
		)
		return
	}
	//Item
	data = ResponseDevicesBlinkDeviceLedsItemToBody(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesBlinkLedsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesBlinkLedsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesBlinkLedsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesBlinkLeds struct {
	Serial     types.String                     `tfsdk:"serial"`
	Item       *ResponseDevicesBlinkDeviceLeds  `tfsdk:"item"`
	Parameters *RequestDevicesBlinkDeviceLedsRs `tfsdk:"parameters"`
}

type ResponseDevicesBlinkDeviceLeds struct {
	Duration types.Int64 `tfsdk:"duration"`
	Duty     types.Int64 `tfsdk:"duty"`
	Period   types.Int64 `tfsdk:"period"`
}

type RequestDevicesBlinkDeviceLedsRs struct {
	Duration types.Int64 `tfsdk:"duration"`
	Duty     types.Int64 `tfsdk:"duty"`
	Period   types.Int64 `tfsdk:"period"`
}

// FromBody
func (r *DevicesBlinkLeds) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestDevicesBlinkDeviceLeds {
	re := *r.Parameters
	duration := new(int64)
	if !re.Duration.IsUnknown() && !re.Duration.IsNull() {
		*duration = re.Duration.ValueInt64()
	} else {
		duration = nil
	}
	duty := new(int64)
	if !re.Duty.IsUnknown() && !re.Duty.IsNull() {
		*duty = re.Duty.ValueInt64()
	} else {
		duty = nil
	}
	period := new(int64)
	if !re.Period.IsUnknown() && !re.Period.IsNull() {
		*period = re.Period.ValueInt64()
	} else {
		period = nil
	}
	out := merakigosdk.RequestDevicesBlinkDeviceLeds{
		Duration: int64ToIntPointer(duration),
		Duty:     int64ToIntPointer(duty),
		Period:   int64ToIntPointer(period),
	}
	return &out
}

// ToBody
func ResponseDevicesBlinkDeviceLedsItemToBody(state DevicesBlinkLeds, response *merakigosdk.ResponseDevicesBlinkDeviceLeds) DevicesBlinkLeds {
	itemState := ResponseDevicesBlinkDeviceLeds{
		Duration: func() types.Int64 {
			if response.Duration != nil {
				return types.Int64Value(int64(*response.Duration))
			}
			return types.Int64{}
		}(),
		Duty: func() types.Int64 {
			if response.Duty != nil {
				return types.Int64Value(int64(*response.Duty))
			}
			return types.Int64{}
		}(),
		Period: func() types.Int64 {
			if response.Period != nil {
				return types.Int64Value(int64(*response.Period))
			}
			return types.Int64{}
		}(),
	}
	state.Item = &itemState
	return state
}
