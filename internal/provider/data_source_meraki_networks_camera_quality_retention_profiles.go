package provider

// DATA SOURCE NORMAL
import (
	"context"
	"log"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &NetworksCameraQualityRetentionProfilesDataSource{}
	_ datasource.DataSourceWithConfigure = &NetworksCameraQualityRetentionProfilesDataSource{}
)

func NewNetworksCameraQualityRetentionProfilesDataSource() datasource.DataSource {
	return &NetworksCameraQualityRetentionProfilesDataSource{}
}

type NetworksCameraQualityRetentionProfilesDataSource struct {
	client *merakigosdk.Client
}

func (d *NetworksCameraQualityRetentionProfilesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	d.client = client
}

// Metadata returns the data source type name.
func (d *NetworksCameraQualityRetentionProfilesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_camera_quality_retention_profiles"
}

func (d *NetworksCameraQualityRetentionProfilesDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Optional:            true,
			},
			"quality_retention_profile_id": schema.StringAttribute{
				MarkdownDescription: `qualityRetentionProfileId path parameter. Quality retention profile ID`,
				Optional:            true,
			},
			"item": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{

					"audio_recording_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"cloud_archive_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"id": schema.StringAttribute{
						Computed: true,
					},
					"max_retention_days": schema.Int64Attribute{
						Computed: true,
					},
					"motion_based_retention_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"motion_detector_version": schema.Int64Attribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"network_id": schema.StringAttribute{
						Computed: true,
					},
					"restricted_bandwidth_mode_enabled": schema.BoolAttribute{
						Computed: true,
					},
					"schedule_id": schema.StringAttribute{
						Computed: true,
					},
					"video_settings": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{

							"m_v12_m_v22_m_v72": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"quality": schema.StringAttribute{
										Computed: true,
									},
									"resolution": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"m_v12_we": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"quality": schema.StringAttribute{
										Computed: true,
									},
									"resolution": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"m_v21_m_v71": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"quality": schema.StringAttribute{
										Computed: true,
									},
									"resolution": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"m_v32": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{

									"quality": schema.StringAttribute{
										Computed: true,
									},
									"resolution": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
					},
				},
			},

			"items": schema.ListNestedAttribute{
				MarkdownDescription: `Array of ResponseCameraGetNetworkCameraQualityRetentionProfiles`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"audio_recording_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"cloud_archive_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"id": schema.StringAttribute{
							Computed: true,
						},
						"max_retention_days": schema.Int64Attribute{
							Computed: true,
						},
						"motion_based_retention_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"motion_detector_version": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"network_id": schema.StringAttribute{
							Computed: true,
						},
						"restricted_bandwidth_mode_enabled": schema.BoolAttribute{
							Computed: true,
						},
						"schedule_id": schema.StringAttribute{
							Computed: true,
						},
						"video_settings": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{

								"m_v12_m_v22_m_v72": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"quality": schema.StringAttribute{
											Computed: true,
										},
										"resolution": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"m_v12_we": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"quality": schema.StringAttribute{
											Computed: true,
										},
										"resolution": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"m_v21_m_v71": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"quality": schema.StringAttribute{
											Computed: true,
										},
										"resolution": schema.StringAttribute{
											Computed: true,
										},
									},
								},
								"m_v32": schema.SingleNestedAttribute{
									Computed: true,
									Attributes: map[string]schema.Attribute{

										"quality": schema.StringAttribute{
											Computed: true,
										},
										"resolution": schema.StringAttribute{
											Computed: true,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *NetworksCameraQualityRetentionProfilesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var networksCameraQualityRetentionProfiles NetworksCameraQualityRetentionProfiles
	diags := req.Config.Get(ctx, &networksCameraQualityRetentionProfiles)
	if resp.Diagnostics.HasError() {
		return
	}

	method1 := []bool{!networksCameraQualityRetentionProfiles.NetworkID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{!networksCameraQualityRetentionProfiles.NetworkID.IsNull(), !networksCameraQualityRetentionProfiles.QualityRetentionProfileID.IsNull()}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetNetworkCameraQualityRetentionProfiles")
		vvNetworkID := networksCameraQualityRetentionProfiles.NetworkID.ValueString()

		response1, restyResp1, err := d.client.Camera.GetNetworkCameraQualityRetentionProfiles(vvNetworkID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCameraQualityRetentionProfiles",
				err.Error(),
			)
			return
		}

		networksCameraQualityRetentionProfiles = ResponseCameraGetNetworkCameraQualityRetentionProfilesItemsToBody(networksCameraQualityRetentionProfiles, response1)
		diags = resp.State.Set(ctx, &networksCameraQualityRetentionProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetNetworkCameraQualityRetentionProfile")
		vvNetworkID := networksCameraQualityRetentionProfiles.NetworkID.ValueString()
		vvQualityRetentionProfileID := networksCameraQualityRetentionProfiles.QualityRetentionProfileID.ValueString()

		response2, restyResp2, err := d.client.Camera.GetNetworkCameraQualityRetentionProfile(vvNetworkID, vvQualityRetentionProfileID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCameraQualityRetentionProfile",
				err.Error(),
			)
			return
		}

		networksCameraQualityRetentionProfiles = ResponseCameraGetNetworkCameraQualityRetentionProfileItemToBody(networksCameraQualityRetentionProfiles, response2)
		diags = resp.State.Set(ctx, &networksCameraQualityRetentionProfiles)
		resp.Diagnostics.Append(diags...)
		if resp.Diagnostics.HasError() {
			return
		}

	}
}

// structs
type NetworksCameraQualityRetentionProfiles struct {
	NetworkID                 types.String                                                  `tfsdk:"network_id"`
	QualityRetentionProfileID types.String                                                  `tfsdk:"quality_retention_profile_id"`
	Items                     *[]ResponseItemCameraGetNetworkCameraQualityRetentionProfiles `tfsdk:"items"`
	Item                      *ResponseCameraGetNetworkCameraQualityRetentionProfile        `tfsdk:"item"`
}

type ResponseItemCameraGetNetworkCameraQualityRetentionProfiles struct {
	AudioRecordingEnabled          types.Bool                                                               `tfsdk:"audio_recording_enabled"`
	CloudArchiveEnabled            types.Bool                                                               `tfsdk:"cloud_archive_enabled"`
	ID                             types.String                                                             `tfsdk:"id"`
	MaxRetentionDays               types.Int64                                                              `tfsdk:"max_retention_days"`
	MotionBasedRetentionEnabled    types.Bool                                                               `tfsdk:"motion_based_retention_enabled"`
	MotionDetectorVersion          types.Int64                                                              `tfsdk:"motion_detector_version"`
	Name                           types.String                                                             `tfsdk:"name"`
	NetworkID                      types.String                                                             `tfsdk:"network_id"`
	RestrictedBandwidthModeEnabled types.Bool                                                               `tfsdk:"restricted_bandwidth_mode_enabled"`
	ScheduleID                     types.String                                                             `tfsdk:"schedule_id"`
	VideoSettings                  *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettings `tfsdk:"video_settings"`
}

type ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettings struct {
	MV12MV22MV72 *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 `tfsdk:"mv12/mv22/mv72"`
	MV12WE       *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12WE       `tfsdk:"mv12_we"`
	MV21MV71     *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV21MV71     `tfsdk:"mv21/mv71"`
	MV32         *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV32         `tfsdk:"mv32"`
}

type ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12WE struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV21MV71 struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV32 struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfile struct {
	AudioRecordingEnabled          types.Bool                                                          `tfsdk:"audio_recording_enabled"`
	CloudArchiveEnabled            types.Bool                                                          `tfsdk:"cloud_archive_enabled"`
	ID                             types.String                                                        `tfsdk:"id"`
	MaxRetentionDays               types.Int64                                                         `tfsdk:"max_retention_days"`
	MotionBasedRetentionEnabled    types.Bool                                                          `tfsdk:"motion_based_retention_enabled"`
	MotionDetectorVersion          types.Int64                                                         `tfsdk:"motion_detector_version"`
	Name                           types.String                                                        `tfsdk:"name"`
	NetworkID                      types.String                                                        `tfsdk:"network_id"`
	RestrictedBandwidthModeEnabled types.Bool                                                          `tfsdk:"restricted_bandwidth_mode_enabled"`
	ScheduleID                     types.String                                                        `tfsdk:"schedule_id"`
	VideoSettings                  *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettings `tfsdk:"video_settings"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettings struct {
	MV12MV22MV72 *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 `tfsdk:"mv12/mv22/mv72"`
	MV12WE       *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WE       `tfsdk:"mv12_we"`
	MV21MV71     *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71     `tfsdk:"mv21/mv71"`
	MV32         *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32         `tfsdk:"mv32"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WE struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32 struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

// ToBody
func ResponseCameraGetNetworkCameraQualityRetentionProfilesItemsToBody(state NetworksCameraQualityRetentionProfiles, response *merakigosdk.ResponseCameraGetNetworkCameraQualityRetentionProfiles) NetworksCameraQualityRetentionProfiles {
	var items []ResponseItemCameraGetNetworkCameraQualityRetentionProfiles
	for _, item := range *response {
		itemState := ResponseItemCameraGetNetworkCameraQualityRetentionProfiles{
			AudioRecordingEnabled: func() types.Bool {
				if item.AudioRecordingEnabled != nil {
					return types.BoolValue(*item.AudioRecordingEnabled)
				}
				return types.Bool{}
			}(),
			CloudArchiveEnabled: func() types.Bool {
				if item.CloudArchiveEnabled != nil {
					return types.BoolValue(*item.CloudArchiveEnabled)
				}
				return types.Bool{}
			}(),
			ID: types.StringValue(item.ID),
			MaxRetentionDays: func() types.Int64 {
				if item.MaxRetentionDays != nil {
					return types.Int64Value(int64(*item.MaxRetentionDays))
				}
				return types.Int64{}
			}(),
			MotionBasedRetentionEnabled: func() types.Bool {
				if item.MotionBasedRetentionEnabled != nil {
					return types.BoolValue(*item.MotionBasedRetentionEnabled)
				}
				return types.Bool{}
			}(),
			MotionDetectorVersion: func() types.Int64 {
				if item.MotionDetectorVersion != nil {
					return types.Int64Value(int64(*item.MotionDetectorVersion))
				}
				return types.Int64{}
			}(),
			Name:      types.StringValue(item.Name),
			NetworkID: types.StringValue(item.NetworkID),
			RestrictedBandwidthModeEnabled: func() types.Bool {
				if item.RestrictedBandwidthModeEnabled != nil {
					return types.BoolValue(*item.RestrictedBandwidthModeEnabled)
				}
				return types.Bool{}
			}(),
			ScheduleID: types.StringValue(item.ScheduleID),
			VideoSettings: func() *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettings {
				if item.VideoSettings != nil {
					return &ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettings{
						MV12MV22MV72: func() *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72 {
							if item.VideoSettings.MV12MV22MV72 != nil {
								return &ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72{
									Quality:    types.StringValue(item.VideoSettings.MV12MV22MV72.Quality),
									Resolution: types.StringValue(item.VideoSettings.MV12MV22MV72.Resolution),
								}
							}
							return &ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12MV22MV72{}
						}(),
						MV12WE: func() *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12WE {
							if item.VideoSettings.MV12WE != nil {
								return &ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12WE{
									Quality:    types.StringValue(item.VideoSettings.MV12WE.Quality),
									Resolution: types.StringValue(item.VideoSettings.MV12WE.Resolution),
								}
							}
							return &ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV12WE{}
						}(),
						MV21MV71: func() *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV21MV71 {
							if item.VideoSettings.MV21MV71 != nil {
								return &ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV21MV71{
									Quality:    types.StringValue(item.VideoSettings.MV21MV71.Quality),
									Resolution: types.StringValue(item.VideoSettings.MV21MV71.Resolution),
								}
							}
							return &ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV21MV71{}
						}(),
						MV32: func() *ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV32 {
							if item.VideoSettings.MV32 != nil {
								return &ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV32{
									Quality:    types.StringValue(item.VideoSettings.MV32.Quality),
									Resolution: types.StringValue(item.VideoSettings.MV32.Resolution),
								}
							}
							return &ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettingsMV32{}
						}(),
					}
				}
				return &ResponseItemCameraGetNetworkCameraQualityRetentionProfilesVideoSettings{}
			}(),
		}
		items = append(items, itemState)
	}
	state.Items = &items
	return state
}

func ResponseCameraGetNetworkCameraQualityRetentionProfileItemToBody(state NetworksCameraQualityRetentionProfiles, response *merakigosdk.ResponseCameraGetNetworkCameraQualityRetentionProfile) NetworksCameraQualityRetentionProfiles {
	itemState := ResponseCameraGetNetworkCameraQualityRetentionProfile{
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
		VideoSettings: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettings {
			if response.VideoSettings != nil {
				return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettings{
					MV12MV22MV72: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72 {
						if response.VideoSettings.MV12MV22MV72 != nil {
							return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72{
								Quality:    types.StringValue(response.VideoSettings.MV12MV22MV72.Quality),
								Resolution: types.StringValue(response.VideoSettings.MV12MV22MV72.Resolution),
							}
						}
						return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72{}
					}(),
					MV12WE: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WE {
						if response.VideoSettings.MV12WE != nil {
							return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WE{
								Quality:    types.StringValue(response.VideoSettings.MV12WE.Quality),
								Resolution: types.StringValue(response.VideoSettings.MV12WE.Resolution),
							}
						}
						return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WE{}
					}(),
					MV21MV71: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 {
						if response.VideoSettings.MV21MV71 != nil {
							return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71{
								Quality:    types.StringValue(response.VideoSettings.MV21MV71.Quality),
								Resolution: types.StringValue(response.VideoSettings.MV21MV71.Resolution),
							}
						}
						return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71{}
					}(),
					MV32: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32 {
						if response.VideoSettings.MV32 != nil {
							return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32{
								Quality:    types.StringValue(response.VideoSettings.MV32.Quality),
								Resolution: types.StringValue(response.VideoSettings.MV32.Resolution),
							}
						}
						return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32{}
					}(),
				}
			}
			return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettings{}
		}(),
	}
	state.Item = &itemState
	return state
}
