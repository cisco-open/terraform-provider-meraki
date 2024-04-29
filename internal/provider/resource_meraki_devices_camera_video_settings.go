package provider

// RESOURCE NORMAL
import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesCameraVideoSettingsResource{}
	_ resource.ResourceWithConfigure = &DevicesCameraVideoSettingsResource{}
)

func NewDevicesCameraVideoSettingsResource() resource.Resource {
	return &DevicesCameraVideoSettingsResource{}
}

type DevicesCameraVideoSettingsResource struct {
	client *merakigosdk.Client
}

func (r *DevicesCameraVideoSettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesCameraVideoSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_video_settings"
}

func (r *DevicesCameraVideoSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"external_rtsp_enabled": schema.BoolAttribute{
				MarkdownDescription: `Boolean indicating if external rtsp stream is exposed`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"rtsp_url": schema.StringAttribute{
				MarkdownDescription: `External rstp url. Will only be returned if external rtsp stream is exposed`,
				Computed:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
			},
		},
	}
}

func (r *DevicesCameraVideoSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesCameraVideoSettingsRs

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
	responseVerifyItem, restyResp1, err := r.client.Camera.GetDeviceCameraVideoSettings(vvSerial)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesCameraVideoSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource DevicesCameraVideoSettings only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Camera.UpdateDeviceCameraVideoSettings(vvSerial, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCameraVideoSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCameraVideoSettings",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Camera.GetDeviceCameraVideoSettings(vvSerial)
	// Has item and not has items
	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetDeviceCameraVideoSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCameraVideoSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCameraGetDeviceCameraVideoSettingsItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCameraVideoSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesCameraVideoSettingsRs

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
	responseGet, restyRespGet, err := r.client.Camera.GetDeviceCameraVideoSettings(vvSerial)
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
				"Failure when executing GetDeviceCameraVideoSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDeviceCameraVideoSettings",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCameraGetDeviceCameraVideoSettingsItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesCameraVideoSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesCameraVideoSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesCameraVideoSettingsRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Camera.UpdateDeviceCameraVideoSettings(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDeviceCameraVideoSettings",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDeviceCameraVideoSettings",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCameraVideoSettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting DevicesCameraVideoSettings", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesCameraVideoSettingsRs struct {
	Serial              types.String `tfsdk:"serial"`
	ExternalRtspEnabled types.Bool   `tfsdk:"external_rtsp_enabled"`
	RtspURL             types.String `tfsdk:"rtsp_url"`
}

// FromBody
func (r *DevicesCameraVideoSettingsRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCameraUpdateDeviceCameraVideoSettings {
	externalRtspEnabled := new(bool)
	if !r.ExternalRtspEnabled.IsUnknown() && !r.ExternalRtspEnabled.IsNull() {
		*externalRtspEnabled = r.ExternalRtspEnabled.ValueBool()
	} else {
		externalRtspEnabled = nil
	}
	out := merakigosdk.RequestCameraUpdateDeviceCameraVideoSettings{
		ExternalRtspEnabled: externalRtspEnabled,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCameraGetDeviceCameraVideoSettingsItemToBodyRs(state DevicesCameraVideoSettingsRs, response *merakigosdk.ResponseCameraGetDeviceCameraVideoSettings, is_read bool) DevicesCameraVideoSettingsRs {
	itemState := DevicesCameraVideoSettingsRs{
		ExternalRtspEnabled: func() types.Bool {
			if response.ExternalRtspEnabled != nil {
				return types.BoolValue(*response.ExternalRtspEnabled)
			}
			return types.Bool{}
		}(),
		RtspURL: types.StringValue(response.RtspURL),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesCameraVideoSettingsRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesCameraVideoSettingsRs)
}
