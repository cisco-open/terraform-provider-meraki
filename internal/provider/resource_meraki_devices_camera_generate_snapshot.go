package provider

// RESOURCE ACTION

import (
	"context"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesCameraGenerateSnapshotResource{}
	_ resource.ResourceWithConfigure = &DevicesCameraGenerateSnapshotResource{}
)

func NewDevicesCameraGenerateSnapshotResource() resource.Resource {
	return &DevicesCameraGenerateSnapshotResource{}
}

type DevicesCameraGenerateSnapshotResource struct {
	client *merakigosdk.Client
}

func (r *DevicesCameraGenerateSnapshotResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesCameraGenerateSnapshotResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices_camera_generate_snapshot"
}

// resourceAction
func (r *DevicesCameraGenerateSnapshotResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"serial": schema.StringAttribute{
				MarkdownDescription: `serial path parameter.`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"fullframe": schema.BoolAttribute{
						MarkdownDescription: `[optional] If set to "true" the snapshot will be taken at full sensor resolution. This will error if used with timestamp.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.RequiresReplace(),
						},
					},
					"timestamp": schema.StringAttribute{
						MarkdownDescription: `[optional] The snapshot will be taken from this time on the camera. The timestamp is expected to be in ISO 8601 format. If no timestamp is specified, we will assume current time.`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplace(),
						},
					},
				},
			},
		},
	}
}
func (r *DevicesCameraGenerateSnapshotResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesCameraGenerateSnapshot

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
	restyResp1, err := r.client.Camera.GenerateDeviceCameraSnapshot(vvSerial, dataRequest)

	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GenerateDeviceCameraSnapshot",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GenerateDeviceCameraSnapshot",
			err.Error(),
		)
		return
	}
	//Item
	// //entro aqui 2
	// data = ResponseCameraGenerateDeviceCameraSnapshot(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesCameraGenerateSnapshotResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesCameraGenerateSnapshotResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *DevicesCameraGenerateSnapshotResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesCameraGenerateSnapshot struct {
	Serial     types.String                                 `tfsdk:"serial"`
	Parameters *RequestCameraGenerateDeviceCameraSnapshotRs `tfsdk:"parameters"`
}

type RequestCameraGenerateDeviceCameraSnapshotRs struct {
	Fullframe types.Bool   `tfsdk:"fullframe"`
	Timestamp types.String `tfsdk:"timestamp"`
}

// FromBody
func (r *DevicesCameraGenerateSnapshot) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestCameraGenerateDeviceCameraSnapshot {
	emptyString := ""
	re := *r.Parameters
	fullframe := new(bool)
	if !re.Fullframe.IsUnknown() && !re.Fullframe.IsNull() {
		*fullframe = re.Fullframe.ValueBool()
	} else {
		fullframe = nil
	}
	timestamp := new(string)
	if !re.Timestamp.IsUnknown() && !re.Timestamp.IsNull() {
		*timestamp = re.Timestamp.ValueString()
	} else {
		timestamp = &emptyString
	}
	out := merakigosdk.RequestCameraGenerateDeviceCameraSnapshot{
		Fullframe: fullframe,
		Timestamp: *timestamp,
	}
	return &out
}
