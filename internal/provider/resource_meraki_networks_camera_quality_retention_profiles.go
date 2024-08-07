package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksCameraQualityRetentionProfilesResource{}
	_ resource.ResourceWithConfigure = &NetworksCameraQualityRetentionProfilesResource{}
)

func NewNetworksCameraQualityRetentionProfilesResource() resource.Resource {
	return &NetworksCameraQualityRetentionProfilesResource{}
}

type NetworksCameraQualityRetentionProfilesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksCameraQualityRetentionProfilesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksCameraQualityRetentionProfilesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_camera_quality_retention_profiles"
}

func (r *NetworksCameraQualityRetentionProfilesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"audio_recording_enabled": schema.BoolAttribute{
				MarkdownDescription: `Whether or not to record audio. Can be either true or false. Defaults to false.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"cloud_archive_enabled": schema.BoolAttribute{
				MarkdownDescription: `Create redundant video backup using Cloud Archive. Can be either true or false. Defaults to false.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"max_retention_days": schema.Int64Attribute{
				MarkdownDescription: `The maximum number of days for which the data will be stored, or 'null' to keep data until storage space runs out. If the former, it can be one of [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 14, 30, 60, 90] days.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"motion_based_retention_enabled": schema.BoolAttribute{
				MarkdownDescription: `Deletes footage older than 3 days in which no motion was detected. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"motion_detector_version": schema.Int64Attribute{
				MarkdownDescription: `The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the new profile. Must be unique. This parameter is required.`,
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
			"quality_retention_profile_id": schema.StringAttribute{
				MarkdownDescription: `qualityRetentionProfileId path parameter. Quality retention profile ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"restricted_bandwidth_mode_enabled": schema.BoolAttribute{
				MarkdownDescription: `Disable features that require additional bandwidth such as Motion Recap. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"schedule_id": schema.StringAttribute{
				MarkdownDescription: `Schedule for which this camera will record video, or 'null' to always record.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"video_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Video quality and resolution settings for all the camera models.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"m_v12_m_v22_m_v72": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV12/MV22/MV72 camera models.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"High",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1280x720' or '1920x1080'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1280x720",
										"1920x1080",
									),
								},
							},
						},
					},
					"m_v12_we": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV12WE camera models.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"High",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1280x720' or '1920x1080'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1280x720",
										"1920x1080",
									),
								},
							},
						},
					},
					"m_v13": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV13 camera models.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"High",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1920x1080",
										"2688x1512",
										"3840x2160",
									),
								},
							},
						},
					},
					"m_v21_m_v71": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV21/MV71 camera models.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"High",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1280x720'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1280x720",
									),
								},
							},
						},
					},
					"m_v22_xmv72_x": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV22X/MV72X camera models.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"High",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1280x720', '1920x1080' or '2688x1512'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1280x720",
										"1920x1080",
										"2688x1512",
									),
								},
							},
						},
					},
					"m_v32": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV32 camera models.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"High",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1080x1080' or '2112x2112'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1080x1080",
										"2112x2112",
									),
								},
							},
						},
					},
					"m_v33": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV33 camera models.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"High",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1080x1080",
										"2112x2112",
										"2880x2880",
									),
								},
							},
						},
					},
					"m_v52": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV52 camera models.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"High",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1280x720', '1920x1080', '2688x1512' or '3840x2160'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1280x720",
										"1920x1080",
										"2688x1512",
										"3840x2160",
									),
								},
							},
						},
					},
					"m_v63": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV63 camera models.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"High",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080' or '2688x1512'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1920x1080",
										"2688x1512",
									),
								},
							},
						},
					},
					"m_v63_x": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV63X camera models.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"High",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1920x1080",
										"2688x1512",
										"3840x2160",
									),
								},
							},
						},
					},
					"m_v93": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV93 camera models.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"High",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1080x1080' or '2112x2112'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1080x1080",
										"2112x2112",
									),
								},
							},
						},
					},
					"m_v93_x": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV93X camera models.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"High",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1080x1080",
										"2112x2112",
										"2880x2880",
									),
								},
							},
						},
					},
				},
			},
		},
	}
}

//path params to set ['qualityRetentionProfileId']

func (r *NetworksCameraQualityRetentionProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksCameraQualityRetentionProfilesRs

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
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Camera.GetNetworkCameraQualityRetentionProfiles(vvNetworkID)
	//Have Create
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkCameraQualityRetentionProfiles",
					err.Error(),
				)
				return
			}
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvQualityRetentionProfileID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter QualityRetentionProfileID",
					"Fail Parsing QualityRetentionProfileID",
				)
				return
			}
			r.client.Camera.UpdateNetworkCameraQualityRetentionProfile(vvNetworkID, vvQualityRetentionProfileID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Camera.GetNetworkCameraQualityRetentionProfile(vvNetworkID, vvQualityRetentionProfileID)
			if responseVerifyItem2 != nil {
				data = ResponseCameraGetNetworkCameraQualityRetentionProfileItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp2, err := r.client.Camera.CreateNetworkCameraQualityRetentionProfile(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkCameraQualityRetentionProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkCameraQualityRetentionProfile",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Camera.GetNetworkCameraQualityRetentionProfiles(vvNetworkID)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCameraQualityRetentionProfiles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkCameraQualityRetentionProfiles",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvQualityRetentionProfileID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter QualityRetentionProfileID",
				err.Error(),
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Camera.GetNetworkCameraQualityRetentionProfile(vvNetworkID, vvQualityRetentionProfileID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseCameraGetNetworkCameraQualityRetentionProfileItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkCameraQualityRetentionProfile",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCameraQualityRetentionProfile",
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

func (r *NetworksCameraQualityRetentionProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksCameraQualityRetentionProfilesRs

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
	vvQualityRetentionProfileID := data.QualityRetentionProfileID.ValueString()
	responseGet, restyRespGet, err := r.client.Camera.GetNetworkCameraQualityRetentionProfile(vvNetworkID, vvQualityRetentionProfileID)
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
				"Failure when executing GetNetworkCameraQualityRetentionProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkCameraQualityRetentionProfile",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseCameraGetNetworkCameraQualityRetentionProfileItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksCameraQualityRetentionProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("qualityRetention_profile_id"), idParts[1])...)
}

func (r *NetworksCameraQualityRetentionProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksCameraQualityRetentionProfilesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	vvQualityRetentionProfileID := data.QualityRetentionProfileID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Camera.UpdateNetworkCameraQualityRetentionProfile(vvNetworkID, vvQualityRetentionProfileID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkCameraQualityRetentionProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkCameraQualityRetentionProfile",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksCameraQualityRetentionProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksCameraQualityRetentionProfilesRs
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
	vvQualityRetentionProfileID := state.QualityRetentionProfileID.ValueString()
	_, err := r.client.Camera.DeleteNetworkCameraQualityRetentionProfile(vvNetworkID, vvQualityRetentionProfileID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkCameraQualityRetentionProfile", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksCameraQualityRetentionProfilesRs struct {
	NetworkID                      types.String                                                          `tfsdk:"network_id"`
	QualityRetentionProfileID      types.String                                                          `tfsdk:"quality_retention_profile_id"`
	AudioRecordingEnabled          types.Bool                                                            `tfsdk:"audio_recording_enabled"`
	CloudArchiveEnabled            types.Bool                                                            `tfsdk:"cloud_archive_enabled"`
	ID                             types.String                                                          `tfsdk:"id"`
	MaxRetentionDays               types.Int64                                                           `tfsdk:"max_retention_days"`
	MotionBasedRetentionEnabled    types.Bool                                                            `tfsdk:"motion_based_retention_enabled"`
	MotionDetectorVersion          types.Int64                                                           `tfsdk:"motion_detector_version"`
	Name                           types.String                                                          `tfsdk:"name"`
	RestrictedBandwidthModeEnabled types.Bool                                                            `tfsdk:"restricted_bandwidth_mode_enabled"`
	ScheduleID                     types.String                                                          `tfsdk:"schedule_id"`
	VideoSettings                  *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsRs `tfsdk:"video_settings"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsRs struct {
	MV12MV22MV72 *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72Rs `tfsdk:"mv12/mv22/mv72"`
	MV12WE       *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WERs       `tfsdk:"mv12_we"`
	MV21MV71     *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71Rs     `tfsdk:"mv21/mv71"`
	MV32         *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32Rs         `tfsdk:"mv32"`
	MV13         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13Rs       `tfsdk:"mv13"`
	MV22XMV72X   *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72XRs `tfsdk:"mv22_x/mv72_x"`
	MV33         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33Rs       `tfsdk:"mv33"`
	MV52         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52Rs       `tfsdk:"mv52"`
	MV63         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63Rs       `tfsdk:"mv63"`
	MV63X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63XRs      `tfsdk:"mv63_x"`
	MV93         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93Rs       `tfsdk:"mv93"`
	MV93X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93XRs      `tfsdk:"mv93_x"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WERs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72XRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63XRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93XRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

// FromBody
func (r *NetworksCameraQualityRetentionProfilesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfile {
	emptyString := ""
	audioRecordingEnabled := new(bool)
	if !r.AudioRecordingEnabled.IsUnknown() && !r.AudioRecordingEnabled.IsNull() {
		*audioRecordingEnabled = r.AudioRecordingEnabled.ValueBool()
	} else {
		audioRecordingEnabled = nil
	}
	cloudArchiveEnabled := new(bool)
	if !r.CloudArchiveEnabled.IsUnknown() && !r.CloudArchiveEnabled.IsNull() {
		*cloudArchiveEnabled = r.CloudArchiveEnabled.ValueBool()
	} else {
		cloudArchiveEnabled = nil
	}
	maxRetentionDays := new(int64)
	if !r.MaxRetentionDays.IsUnknown() && !r.MaxRetentionDays.IsNull() {
		*maxRetentionDays = r.MaxRetentionDays.ValueInt64()
	} else {
		maxRetentionDays = nil
	}
	motionBasedRetentionEnabled := new(bool)
	if !r.MotionBasedRetentionEnabled.IsUnknown() && !r.MotionBasedRetentionEnabled.IsNull() {
		*motionBasedRetentionEnabled = r.MotionBasedRetentionEnabled.ValueBool()
	} else {
		motionBasedRetentionEnabled = nil
	}
	motionDetectorVersion := new(int64)
	if !r.MotionDetectorVersion.IsUnknown() && !r.MotionDetectorVersion.IsNull() {
		*motionDetectorVersion = r.MotionDetectorVersion.ValueInt64()
	} else {
		motionDetectorVersion = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	restrictedBandwidthModeEnabled := new(bool)
	if !r.RestrictedBandwidthModeEnabled.IsUnknown() && !r.RestrictedBandwidthModeEnabled.IsNull() {
		*restrictedBandwidthModeEnabled = r.RestrictedBandwidthModeEnabled.ValueBool()
	} else {
		restrictedBandwidthModeEnabled = nil
	}
	scheduleID := new(string)
	if !r.ScheduleID.IsUnknown() && !r.ScheduleID.IsNull() {
		*scheduleID = r.ScheduleID.ValueString()
	} else {
		scheduleID = &emptyString
	}
	var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettings *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettings
	if r.VideoSettings != nil {
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72
		if r.VideoSettings.MV12MV22MV72 != nil {
			quality := r.VideoSettings.MV12MV22MV72.Quality.ValueString()
			resolution := r.VideoSettings.MV12MV22MV72.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE
		if r.VideoSettings.MV12WE != nil {
			quality := r.VideoSettings.MV12WE.Quality.ValueString()
			resolution := r.VideoSettings.MV12WE.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13
		if r.VideoSettings.MV13 != nil {
			quality := r.VideoSettings.MV13.Quality.ValueString()
			resolution := r.VideoSettings.MV13.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71
		if r.VideoSettings.MV21MV71 != nil {
			quality := r.VideoSettings.MV21MV71.Quality.ValueString()
			resolution := r.VideoSettings.MV21MV71.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X
		if r.VideoSettings.MV22XMV72X != nil {
			quality := r.VideoSettings.MV22XMV72X.Quality.ValueString()
			resolution := r.VideoSettings.MV22XMV72X.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32
		if r.VideoSettings.MV32 != nil {
			quality := r.VideoSettings.MV32.Quality.ValueString()
			resolution := r.VideoSettings.MV32.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33
		if r.VideoSettings.MV33 != nil {
			quality := r.VideoSettings.MV33.Quality.ValueString()
			resolution := r.VideoSettings.MV33.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52
		if r.VideoSettings.MV52 != nil {
			quality := r.VideoSettings.MV52.Quality.ValueString()
			resolution := r.VideoSettings.MV52.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63
		if r.VideoSettings.MV63 != nil {
			quality := r.VideoSettings.MV63.Quality.ValueString()
			resolution := r.VideoSettings.MV63.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X
		if r.VideoSettings.MV63X != nil {
			quality := r.VideoSettings.MV63X.Quality.ValueString()
			resolution := r.VideoSettings.MV63X.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93
		if r.VideoSettings.MV93 != nil {
			quality := r.VideoSettings.MV93.Quality.ValueString()
			resolution := r.VideoSettings.MV93.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X
		if r.VideoSettings.MV93X != nil {
			quality := r.VideoSettings.MV93X.Quality.ValueString()
			resolution := r.VideoSettings.MV93X.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettings = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettings{
			MV12MV22MV72: requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72,
			MV12WE:       requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE,
			MV13:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13,
			MV21MV71:     requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71,
			MV22XMV72X:   requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X,
			MV32:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32,
			MV33:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33,
			MV52:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52,
			MV63:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63,
			MV63X:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X,
			MV93:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93,
			MV93X:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X,
		}
	}
	out := merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfile{
		AudioRecordingEnabled:          audioRecordingEnabled,
		CloudArchiveEnabled:            cloudArchiveEnabled,
		MaxRetentionDays:               int64ToIntPointer(maxRetentionDays),
		MotionBasedRetentionEnabled:    motionBasedRetentionEnabled,
		MotionDetectorVersion:          int64ToIntPointer(motionDetectorVersion),
		Name:                           *name,
		RestrictedBandwidthModeEnabled: restrictedBandwidthModeEnabled,
		ScheduleID:                     *scheduleID,
		VideoSettings:                  requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettings,
	}
	return &out
}
func (r *NetworksCameraQualityRetentionProfilesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfile {
	emptyString := ""
	audioRecordingEnabled := new(bool)
	if !r.AudioRecordingEnabled.IsUnknown() && !r.AudioRecordingEnabled.IsNull() {
		*audioRecordingEnabled = r.AudioRecordingEnabled.ValueBool()
	} else {
		audioRecordingEnabled = nil
	}
	cloudArchiveEnabled := new(bool)
	if !r.CloudArchiveEnabled.IsUnknown() && !r.CloudArchiveEnabled.IsNull() {
		*cloudArchiveEnabled = r.CloudArchiveEnabled.ValueBool()
	} else {
		cloudArchiveEnabled = nil
	}
	maxRetentionDays := new(int64)
	if !r.MaxRetentionDays.IsUnknown() && !r.MaxRetentionDays.IsNull() {
		*maxRetentionDays = r.MaxRetentionDays.ValueInt64()
	} else {
		maxRetentionDays = nil
	}
	motionBasedRetentionEnabled := new(bool)
	if !r.MotionBasedRetentionEnabled.IsUnknown() && !r.MotionBasedRetentionEnabled.IsNull() {
		*motionBasedRetentionEnabled = r.MotionBasedRetentionEnabled.ValueBool()
	} else {
		motionBasedRetentionEnabled = nil
	}
	motionDetectorVersion := new(int64)
	if !r.MotionDetectorVersion.IsUnknown() && !r.MotionDetectorVersion.IsNull() {
		*motionDetectorVersion = r.MotionDetectorVersion.ValueInt64()
	} else {
		motionDetectorVersion = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	restrictedBandwidthModeEnabled := new(bool)
	if !r.RestrictedBandwidthModeEnabled.IsUnknown() && !r.RestrictedBandwidthModeEnabled.IsNull() {
		*restrictedBandwidthModeEnabled = r.RestrictedBandwidthModeEnabled.ValueBool()
	} else {
		restrictedBandwidthModeEnabled = nil
	}
	scheduleID := new(string)
	if !r.ScheduleID.IsUnknown() && !r.ScheduleID.IsNull() {
		*scheduleID = r.ScheduleID.ValueString()
	} else {
		scheduleID = &emptyString
	}
	var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettings *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettings
	if r.VideoSettings != nil {
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72
		if r.VideoSettings.MV12MV22MV72 != nil {
			quality := r.VideoSettings.MV12MV22MV72.Quality.ValueString()
			resolution := r.VideoSettings.MV12MV22MV72.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE
		if r.VideoSettings.MV12WE != nil {
			quality := r.VideoSettings.MV12WE.Quality.ValueString()
			resolution := r.VideoSettings.MV12WE.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13
		if r.VideoSettings.MV13 != nil {
			quality := r.VideoSettings.MV13.Quality.ValueString()
			resolution := r.VideoSettings.MV13.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71
		if r.VideoSettings.MV21MV71 != nil {
			quality := r.VideoSettings.MV21MV71.Quality.ValueString()
			resolution := r.VideoSettings.MV21MV71.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X
		if r.VideoSettings.MV22XMV72X != nil {
			quality := r.VideoSettings.MV22XMV72X.Quality.ValueString()
			resolution := r.VideoSettings.MV22XMV72X.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32
		if r.VideoSettings.MV32 != nil {
			quality := r.VideoSettings.MV32.Quality.ValueString()
			resolution := r.VideoSettings.MV32.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33
		if r.VideoSettings.MV33 != nil {
			quality := r.VideoSettings.MV33.Quality.ValueString()
			resolution := r.VideoSettings.MV33.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52
		if r.VideoSettings.MV52 != nil {
			quality := r.VideoSettings.MV52.Quality.ValueString()
			resolution := r.VideoSettings.MV52.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63
		if r.VideoSettings.MV63 != nil {
			quality := r.VideoSettings.MV63.Quality.ValueString()
			resolution := r.VideoSettings.MV63.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X
		if r.VideoSettings.MV63X != nil {
			quality := r.VideoSettings.MV63X.Quality.ValueString()
			resolution := r.VideoSettings.MV63X.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93
		if r.VideoSettings.MV93 != nil {
			quality := r.VideoSettings.MV93.Quality.ValueString()
			resolution := r.VideoSettings.MV93.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93X *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93X
		if r.VideoSettings.MV93X != nil {
			quality := r.VideoSettings.MV93X.Quality.ValueString()
			resolution := r.VideoSettings.MV93X.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93X = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93X{
				Quality:    quality,
				Resolution: resolution,
			}
		}
		requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettings = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettings{
			MV12MV22MV72: requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72,
			MV12WE:       requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE,
			MV13:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13,
			MV21MV71:     requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71,
			MV22XMV72X:   requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X,
			MV32:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32,
			MV33:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33,
			MV52:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52,
			MV63:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63,
			MV63X:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X,
			MV93:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93,
			MV93X:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93X,
		}
	}
	out := merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfile{
		AudioRecordingEnabled:          audioRecordingEnabled,
		CloudArchiveEnabled:            cloudArchiveEnabled,
		MaxRetentionDays:               int64ToIntPointer(maxRetentionDays),
		MotionBasedRetentionEnabled:    motionBasedRetentionEnabled,
		MotionDetectorVersion:          int64ToIntPointer(motionDetectorVersion),
		Name:                           *name,
		RestrictedBandwidthModeEnabled: restrictedBandwidthModeEnabled,
		ScheduleID:                     *scheduleID,
		VideoSettings:                  requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettings,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseCameraGetNetworkCameraQualityRetentionProfileItemToBodyRs(state NetworksCameraQualityRetentionProfilesRs, response *merakigosdk.ResponseCameraGetNetworkCameraQualityRetentionProfile, is_read bool) NetworksCameraQualityRetentionProfilesRs {
	itemState := NetworksCameraQualityRetentionProfilesRs{
		AudioRecordingEnabled: func() types.Bool {
			if response.AudioRecordingEnabled != nil {
				return types.BoolValue(*response.AudioRecordingEnabled)
			}
			return types.Bool{}
		}(),
		CloudArchiveEnabled: func() types.Bool {
			if response.CloudArchiveEnabled != nil {
				return types.BoolValue(*response.CloudArchiveEnabled)
			}
			return types.Bool{}
		}(),
		ID: types.StringValue(response.ID),
		MaxRetentionDays: func() types.Int64 {
			if response.MaxRetentionDays != nil {
				return types.Int64Value(int64(*response.MaxRetentionDays))
			}
			return types.Int64{}
		}(),
		MotionBasedRetentionEnabled: func() types.Bool {
			if response.MotionBasedRetentionEnabled != nil {
				return types.BoolValue(*response.MotionBasedRetentionEnabled)
			}
			return types.Bool{}
		}(),
		MotionDetectorVersion: func() types.Int64 {
			if response.MotionDetectorVersion != nil {
				return types.Int64Value(int64(*response.MotionDetectorVersion))
			}
			return types.Int64{}
		}(),
		Name:      types.StringValue(response.Name),
		NetworkID: types.StringValue(response.NetworkID),
		RestrictedBandwidthModeEnabled: func() types.Bool {
			if response.RestrictedBandwidthModeEnabled != nil {
				return types.BoolValue(*response.RestrictedBandwidthModeEnabled)
			}
			return types.Bool{}
		}(),
		ScheduleID: types.StringValue(response.ScheduleID),
		VideoSettings: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsRs {
			if response.VideoSettings != nil {
				return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsRs{
					MV12MV22MV72: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72Rs {
						if response.VideoSettings.MV12MV22MV72 != nil {
							return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72Rs{
								Quality:    types.StringValue(response.VideoSettings.MV12MV22MV72.Quality),
								Resolution: types.StringValue(response.VideoSettings.MV12MV22MV72.Resolution),
							}
						}
						return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72Rs{}
					}(),
					MV12WE: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WERs {
						if response.VideoSettings.MV12WE != nil {
							return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WERs{
								Quality:    types.StringValue(response.VideoSettings.MV12WE.Quality),
								Resolution: types.StringValue(response.VideoSettings.MV12WE.Resolution),
							}
						}
						return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WERs{}
					}(),
					MV21MV71: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71Rs {
						if response.VideoSettings.MV21MV71 != nil {
							return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71Rs{
								Quality:    types.StringValue(response.VideoSettings.MV21MV71.Quality),
								Resolution: types.StringValue(response.VideoSettings.MV21MV71.Resolution),
							}
						}
						return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71Rs{}
					}(),
					MV32: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32Rs {
						if response.VideoSettings.MV32 != nil {
							return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32Rs{
								Quality:    types.StringValue(response.VideoSettings.MV32.Quality),
								Resolution: types.StringValue(response.VideoSettings.MV32.Resolution),
							}
						}
						return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32Rs{}
					}(),
				}
			}
			return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksCameraQualityRetentionProfilesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksCameraQualityRetentionProfilesRs)
}
