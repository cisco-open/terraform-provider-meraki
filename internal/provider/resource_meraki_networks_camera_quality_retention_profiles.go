// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0
package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"strconv"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v5/sdk"

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
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"cloud_archive_enabled": schema.BoolAttribute{
				MarkdownDescription: `Create redundant video backup using Cloud Archive. Can be either true or false. Defaults to false.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"max_retention_days": schema.Int64Attribute{
				MarkdownDescription: `The maximum number of days for which the data will be stored, or 'null' to keep data until storage space runs out. If the former, it can be in the range of one to ninety days.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"motion_based_retention_enabled": schema.BoolAttribute{
				MarkdownDescription: `Deletes footage older than 3 days in which no motion was detected. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"motion_detector_version": schema.Int64Attribute{
				MarkdownDescription: `The version of the motion detector that will be used by the camera. Only applies to Gen 2 cameras. Defaults to v2.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the new profile. Must be unique. This parameter is required.`,
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
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"restricted_bandwidth_mode_enabled": schema.BoolAttribute{
				MarkdownDescription: `Disable features that require additional bandwidth such as Motion Recap. Can be either true or false. Defaults to false. This setting does not apply to MV2 cameras.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"schedule_id": schema.StringAttribute{
				MarkdownDescription: `Schedule for which this camera will record video, or 'null' to always record.`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"smart_retention": schema.SingleNestedAttribute{
				MarkdownDescription: `Smart Retention records footage in two qualities and intelligently retains higher quality when motion, people or vehicles are detected.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Boolean indicating if Smart Retention is enabled(true) or disabled(false).`,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"video_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Video quality and resolution settings for all the camera models.`,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"m_v12_m_v22_m_v72": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV12/MV22/MV72 camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1280x720' or '1920x1080'.
                                              Allowed values: [1280x720,1920x1080]`,
								Optional: true,
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
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1280x720' or '1920x1080'.
                                              Allowed values: [1280x720,1920x1080]`,
								Optional: true,
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
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
					"m_v13_m": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV13M camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1280x720'.
                                              Allowed values: [1280x720]`,
								Optional: true,
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
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1280x720', '1920x1080' or '2688x1512'.
                                              Allowed values: [1280x720,1920x1080,2688x1512]`,
								Optional: true,
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
					"m_v23": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV23 camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
					"m_v23_m": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV23M camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
					"m_v23_x": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV23X camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
					"m_v32": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV32 camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1080x1080' or '2112x2112'.
                                              Allowed values: [1080x1080,2112x2112]`,
								Optional: true,
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
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
                                              Allowed values: [1080x1080,2112x2112,2880x2880]`,
								Optional: true,
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
					"m_v33_m": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV33M camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
                                              Allowed values: [1080x1080,2112x2112,2880x2880]`,
								Optional: true,
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
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1280x720', '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1280x720,1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
					"m_v53_x": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV53X camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
					"m_v63": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV63 camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
					"m_v63_m": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV63M camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
					"m_v63_x": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV63X camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
					"m_v73": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV73 camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
					"m_v73_m": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV73M camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
					"m_v73_x": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV73X camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1920x1080', '2688x1512' or '3840x2160'.
                                              Allowed values: [1920x1080,2688x1512,3840x2160]`,
								Optional: true,
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
					"m_v84_x": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV84X camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard' or 'Enhanced'.
                                              Allowed values: [Enhanced,Standard]`,
								Optional: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"Enhanced",
										"Standard",
									),
								},
							},
							"resolution": schema.StringAttribute{
								MarkdownDescription: `Resolution of the camera. Can be one of '1440x1080' or '2560x1920'.
                                              Allowed values: [1440x1080,2560x1920]`,
								Optional: true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"1440x1080",
										"2560x1920",
									),
								},
							},
						},
					},
					"m_v93": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV93 camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
                                              Allowed values: [1080x1080,2112x2112,2880x2880]`,
								Optional: true,
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
					"m_v93_m": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV93M camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
                                              Allowed values: [1080x1080,2112x2112,2880x2880]`,
								Optional: true,
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
					"m_v93_x": schema.SingleNestedAttribute{
						MarkdownDescription: `Quality and resolution for MV93X camera models.`,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"quality": schema.StringAttribute{
								MarkdownDescription: `Quality of the camera. Can be one of 'Standard', 'Enhanced' or 'High'.
                                              Allowed values: [Enhanced,High,Standard]`,
								Optional: true,
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
								MarkdownDescription: `Resolution of the camera. Can be one of '1080x1080', '2112x2112' or '2880x2880'.
                                              Allowed values: [1080x1080,2112x2112,2880x2880]`,
								Optional: true,
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
	// Has Paths
	vvNetworkID := data.NetworkID.ValueString()
	//Has Item and has items and post

	vvName := data.Name.ValueString()

	responseVerifyItem, restyResp1, err := r.client.Camera.GetNetworkCameraQualityRetentionProfiles(vvNetworkID)
	//Has Post
	if err != nil {
		if restyResp1 != nil {
			if restyResp1.StatusCode() != 404 {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkCameraQualityRetentionProfiles",
					restyResp1.String(),
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
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkCameraQualityRetentionProfile",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkCameraQualityRetentionProfile",
			err.Error(),
		)
		return
	}

	responseGet, restyResp1, err := r.client.Camera.GetNetworkCameraQualityRetentionProfiles(vvNetworkID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkCameraQualityRetentionProfiles",
				restyResp1.String(),
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
				"Fail Parsing QualityRetentionProfileID",
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
					restyRespGet.String(),
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

	diags := req.State.Get(ctx, &data)
	if resp.Diagnostics.Append(diags...); resp.Diagnostics.HasError() {
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
				restyRespGet.String(),
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
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
func (r *NetworksCameraQualityRetentionProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: networkId,qualityRetentionProfileId. Got: %q", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("quality_retention_profile_id"), idParts[1])...)
}

func (r *NetworksCameraQualityRetentionProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan NetworksCameraQualityRetentionProfilesRs
	merge(ctx, req, resp, &plan)

	if resp.Diagnostics.HasError() {
		return
	}
	//Path Params
	vvNetworkID := plan.NetworkID.ValueString()
	vvQualityRetentionProfileID := plan.QualityRetentionProfileID.ValueString()
	dataRequest := plan.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Camera.UpdateNetworkCameraQualityRetentionProfile(vvNetworkID, vvQualityRetentionProfileID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkCameraQualityRetentionProfile",
				"Status: "+strconv.Itoa(restyResp2.StatusCode())+"\n"+restyResp2.String(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkCameraQualityRetentionProfile",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
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
	NetworkID                      types.String                                                           `tfsdk:"network_id"`
	QualityRetentionProfileID      types.String                                                           `tfsdk:"quality_retention_profile_id"`
	AudioRecordingEnabled          types.Bool                                                             `tfsdk:"audio_recording_enabled"`
	CloudArchiveEnabled            types.Bool                                                             `tfsdk:"cloud_archive_enabled"`
	ID                             types.String                                                           `tfsdk:"id"`
	MaxRetentionDays               types.Int64                                                            `tfsdk:"max_retention_days"`
	MotionBasedRetentionEnabled    types.Bool                                                             `tfsdk:"motion_based_retention_enabled"`
	MotionDetectorVersion          types.Int64                                                            `tfsdk:"motion_detector_version"`
	Name                           types.String                                                           `tfsdk:"name"`
	RestrictedBandwidthModeEnabled types.Bool                                                             `tfsdk:"restricted_bandwidth_mode_enabled"`
	ScheduleID                     types.String                                                           `tfsdk:"schedule_id"`
	SmartRetention                 *ResponseCameraGetNetworkCameraQualityRetentionProfileSmartRetentionRs `tfsdk:"smart_retention"`
	VideoSettings                  *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsRs  `tfsdk:"video_settings"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfileSmartRetentionRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsRs struct {
	MV12MV22MV72 *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72Rs `tfsdk:"m_v12_m_v22_m_v72"`
	MV12WE       *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WERs       `tfsdk:"m_v12_we"`
	MV21MV71     *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71Rs     `tfsdk:"m_v21_m_v71"`
	MV32         *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32Rs         `tfsdk:"m_v32"`
	MV13         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13Rs       `tfsdk:"mv13"`
	MV13M        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13MRs      `tfsdk:"mv13_m"`
	MV22XMV72X   *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72XRs `tfsdk:"mv22_x/mv72_x"`
	MV23         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23Rs       `tfsdk:"mv23"`
	MV23M        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23MRs      `tfsdk:"mv23_m"`
	MV23X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23XRs      `tfsdk:"mv23_x"`
	MV33         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33Rs       `tfsdk:"mv33"`
	MV33M        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33MRs      `tfsdk:"mv33_m"`
	MV52         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52Rs       `tfsdk:"mv52"`
	MV53X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV53XRs      `tfsdk:"mv53_x"`
	MV63         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63Rs       `tfsdk:"mv63"`
	MV63M        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63MRs      `tfsdk:"mv63_m"`
	MV63X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63XRs      `tfsdk:"mv63_x"`
	MV73         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73Rs       `tfsdk:"mv73"`
	MV73M        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73MRs      `tfsdk:"mv73_m"`
	MV73X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73XRs      `tfsdk:"mv73_x"`
	MV84X        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV84XRs      `tfsdk:"mv84_x"`
	MV93         *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93Rs       `tfsdk:"mv93"`
	MV93M        *RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93MRs      `tfsdk:"mv93_m"`
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

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13MRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72XRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23MRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23XRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33MRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV53XRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63MRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63XRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73MRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73XRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV84XRs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93Rs struct {
	Quality    types.String `tfsdk:"quality"`
	Resolution types.String `tfsdk:"resolution"`
}

type RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93MRs struct {
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
	var requestCameraCreateNetworkCameraQualityRetentionProfileSmartRetention *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileSmartRetention

	if r.SmartRetention != nil {
		enabled := func() *bool {
			if !r.SmartRetention.Enabled.IsUnknown() && !r.SmartRetention.Enabled.IsNull() {
				return r.SmartRetention.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestCameraCreateNetworkCameraQualityRetentionProfileSmartRetention = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileSmartRetention{
			Enabled: enabled,
		}
		//[debug] Is Array: False
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
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE

		if r.VideoSettings.MV12WE != nil {
			quality := r.VideoSettings.MV12WE.Quality.ValueString()
			resolution := r.VideoSettings.MV12WE.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13

		if r.VideoSettings.MV13 != nil {
			quality := r.VideoSettings.MV13.Quality.ValueString()
			resolution := r.VideoSettings.MV13.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13M *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13M

		if r.VideoSettings.MV13M != nil {
			quality := r.VideoSettings.MV13M.Quality.ValueString()
			resolution := r.VideoSettings.MV13M.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13M = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13M{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71

		if r.VideoSettings.MV21MV71 != nil {
			quality := r.VideoSettings.MV21MV71.Quality.ValueString()
			resolution := r.VideoSettings.MV21MV71.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X

		if r.VideoSettings.MV22XMV72X != nil {
			quality := r.VideoSettings.MV22XMV72X.Quality.ValueString()
			resolution := r.VideoSettings.MV22XMV72X.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23

		if r.VideoSettings.MV23 != nil {
			quality := r.VideoSettings.MV23.Quality.ValueString()
			resolution := r.VideoSettings.MV23.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23M *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23M

		if r.VideoSettings.MV23M != nil {
			quality := r.VideoSettings.MV23M.Quality.ValueString()
			resolution := r.VideoSettings.MV23M.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23M = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23M{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23X *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23X

		if r.VideoSettings.MV23X != nil {
			quality := r.VideoSettings.MV23X.Quality.ValueString()
			resolution := r.VideoSettings.MV23X.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23X = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32

		if r.VideoSettings.MV32 != nil {
			quality := r.VideoSettings.MV32.Quality.ValueString()
			resolution := r.VideoSettings.MV32.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33

		if r.VideoSettings.MV33 != nil {
			quality := r.VideoSettings.MV33.Quality.ValueString()
			resolution := r.VideoSettings.MV33.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33M *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33M

		if r.VideoSettings.MV33M != nil {
			quality := r.VideoSettings.MV33M.Quality.ValueString()
			resolution := r.VideoSettings.MV33M.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33M = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33M{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52

		if r.VideoSettings.MV52 != nil {
			quality := r.VideoSettings.MV52.Quality.ValueString()
			resolution := r.VideoSettings.MV52.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV53X *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV53X

		if r.VideoSettings.MV53X != nil {
			quality := r.VideoSettings.MV53X.Quality.ValueString()
			resolution := r.VideoSettings.MV53X.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV53X = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV53X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63

		if r.VideoSettings.MV63 != nil {
			quality := r.VideoSettings.MV63.Quality.ValueString()
			resolution := r.VideoSettings.MV63.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63M *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63M

		if r.VideoSettings.MV63M != nil {
			quality := r.VideoSettings.MV63M.Quality.ValueString()
			resolution := r.VideoSettings.MV63M.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63M = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63M{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X

		if r.VideoSettings.MV63X != nil {
			quality := r.VideoSettings.MV63X.Quality.ValueString()
			resolution := r.VideoSettings.MV63X.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73

		if r.VideoSettings.MV73 != nil {
			quality := r.VideoSettings.MV73.Quality.ValueString()
			resolution := r.VideoSettings.MV73.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73M *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73M

		if r.VideoSettings.MV73M != nil {
			quality := r.VideoSettings.MV73M.Quality.ValueString()
			resolution := r.VideoSettings.MV73M.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73M = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73M{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73X *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73X

		if r.VideoSettings.MV73X != nil {
			quality := r.VideoSettings.MV73X.Quality.ValueString()
			resolution := r.VideoSettings.MV73X.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73X = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV84X *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV84X

		if r.VideoSettings.MV84X != nil {
			quality := r.VideoSettings.MV84X.Quality.ValueString()
			resolution := r.VideoSettings.MV84X.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV84X = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV84X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93 *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93

		if r.VideoSettings.MV93 != nil {
			quality := r.VideoSettings.MV93.Quality.ValueString()
			resolution := r.VideoSettings.MV93.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93 = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93M *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93M

		if r.VideoSettings.MV93M != nil {
			quality := r.VideoSettings.MV93M.Quality.ValueString()
			resolution := r.VideoSettings.MV93M.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93M = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93M{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X *merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X

		if r.VideoSettings.MV93X != nil {
			quality := r.VideoSettings.MV93X.Quality.ValueString()
			resolution := r.VideoSettings.MV93X.Resolution.ValueString()
			requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettings = &merakigosdk.RequestCameraCreateNetworkCameraQualityRetentionProfileVideoSettings{
			MV12MV22MV72: requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72,
			MV12WE:       requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE,
			MV13:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13,
			MV13M:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV13M,
			MV21MV71:     requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71,
			MV22XMV72X:   requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X,
			MV23:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23,
			MV23M:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23M,
			MV23X:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV23X,
			MV32:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV32,
			MV33:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33,
			MV33M:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV33M,
			MV52:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV52,
			MV53X:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV53X,
			MV63:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63,
			MV63M:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63M,
			MV63X:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV63X,
			MV73:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73,
			MV73M:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73M,
			MV73X:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV73X,
			MV84X:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV84X,
			MV93:         requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93,
			MV93M:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93M,
			MV93X:        requestCameraCreateNetworkCameraQualityRetentionProfileVideoSettingsMV93X,
		}
		//[debug] Is Array: False
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
		SmartRetention:                 requestCameraCreateNetworkCameraQualityRetentionProfileSmartRetention,
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
	var requestCameraUpdateNetworkCameraQualityRetentionProfileSmartRetention *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileSmartRetention

	if r.SmartRetention != nil {
		enabled := func() *bool {
			if !r.SmartRetention.Enabled.IsUnknown() && !r.SmartRetention.Enabled.IsNull() {
				return r.SmartRetention.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestCameraUpdateNetworkCameraQualityRetentionProfileSmartRetention = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileSmartRetention{
			Enabled: enabled,
		}
		//[debug] Is Array: False
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
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE

		if r.VideoSettings.MV12WE != nil {
			quality := r.VideoSettings.MV12WE.Quality.ValueString()
			resolution := r.VideoSettings.MV12WE.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13

		if r.VideoSettings.MV13 != nil {
			quality := r.VideoSettings.MV13.Quality.ValueString()
			resolution := r.VideoSettings.MV13.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13M *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13M

		if r.VideoSettings.MV13M != nil {
			quality := r.VideoSettings.MV13M.Quality.ValueString()
			resolution := r.VideoSettings.MV13M.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13M = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13M{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71

		if r.VideoSettings.MV21MV71 != nil {
			quality := r.VideoSettings.MV21MV71.Quality.ValueString()
			resolution := r.VideoSettings.MV21MV71.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X

		if r.VideoSettings.MV22XMV72X != nil {
			quality := r.VideoSettings.MV22XMV72X.Quality.ValueString()
			resolution := r.VideoSettings.MV22XMV72X.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23

		if r.VideoSettings.MV23 != nil {
			quality := r.VideoSettings.MV23.Quality.ValueString()
			resolution := r.VideoSettings.MV23.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23M *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23M

		if r.VideoSettings.MV23M != nil {
			quality := r.VideoSettings.MV23M.Quality.ValueString()
			resolution := r.VideoSettings.MV23M.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23M = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23M{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23X *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23X

		if r.VideoSettings.MV23X != nil {
			quality := r.VideoSettings.MV23X.Quality.ValueString()
			resolution := r.VideoSettings.MV23X.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23X = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32

		if r.VideoSettings.MV32 != nil {
			quality := r.VideoSettings.MV32.Quality.ValueString()
			resolution := r.VideoSettings.MV32.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33

		if r.VideoSettings.MV33 != nil {
			quality := r.VideoSettings.MV33.Quality.ValueString()
			resolution := r.VideoSettings.MV33.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33M *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33M

		if r.VideoSettings.MV33M != nil {
			quality := r.VideoSettings.MV33M.Quality.ValueString()
			resolution := r.VideoSettings.MV33M.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33M = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33M{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52

		if r.VideoSettings.MV52 != nil {
			quality := r.VideoSettings.MV52.Quality.ValueString()
			resolution := r.VideoSettings.MV52.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV53X *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV53X

		if r.VideoSettings.MV53X != nil {
			quality := r.VideoSettings.MV53X.Quality.ValueString()
			resolution := r.VideoSettings.MV53X.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV53X = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV53X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63

		if r.VideoSettings.MV63 != nil {
			quality := r.VideoSettings.MV63.Quality.ValueString()
			resolution := r.VideoSettings.MV63.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63M *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63M

		if r.VideoSettings.MV63M != nil {
			quality := r.VideoSettings.MV63M.Quality.ValueString()
			resolution := r.VideoSettings.MV63M.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63M = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63M{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X

		if r.VideoSettings.MV63X != nil {
			quality := r.VideoSettings.MV63X.Quality.ValueString()
			resolution := r.VideoSettings.MV63X.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73

		if r.VideoSettings.MV73 != nil {
			quality := r.VideoSettings.MV73.Quality.ValueString()
			resolution := r.VideoSettings.MV73.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73M *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73M

		if r.VideoSettings.MV73M != nil {
			quality := r.VideoSettings.MV73M.Quality.ValueString()
			resolution := r.VideoSettings.MV73M.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73M = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73M{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73X *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73X

		if r.VideoSettings.MV73X != nil {
			quality := r.VideoSettings.MV73X.Quality.ValueString()
			resolution := r.VideoSettings.MV73X.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73X = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV84X *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV84X

		if r.VideoSettings.MV84X != nil {
			quality := r.VideoSettings.MV84X.Quality.ValueString()
			resolution := r.VideoSettings.MV84X.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV84X = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV84X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93 *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93

		if r.VideoSettings.MV93 != nil {
			quality := r.VideoSettings.MV93.Quality.ValueString()
			resolution := r.VideoSettings.MV93.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93 = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93M *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93M

		if r.VideoSettings.MV93M != nil {
			quality := r.VideoSettings.MV93M.Quality.ValueString()
			resolution := r.VideoSettings.MV93M.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93M = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93M{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		var requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93X *merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93X

		if r.VideoSettings.MV93X != nil {
			quality := r.VideoSettings.MV93X.Quality.ValueString()
			resolution := r.VideoSettings.MV93X.Resolution.ValueString()
			requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93X = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93X{
				Quality:    quality,
				Resolution: resolution,
			}
			//[debug] Is Array: False
		}
		requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettings = &merakigosdk.RequestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettings{
			MV12MV22MV72: requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72,
			MV12WE:       requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV12WE,
			MV13:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13,
			MV13M:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV13M,
			MV21MV71:     requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71,
			MV22XMV72X:   requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV22XMV72X,
			MV23:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23,
			MV23M:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23M,
			MV23X:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV23X,
			MV32:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV32,
			MV33:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33,
			MV33M:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV33M,
			MV52:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV52,
			MV53X:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV53X,
			MV63:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63,
			MV63M:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63M,
			MV63X:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV63X,
			MV73:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73,
			MV73M:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73M,
			MV73X:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV73X,
			MV84X:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV84X,
			MV93:         requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93,
			MV93M:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93M,
			MV93X:        requestCameraUpdateNetworkCameraQualityRetentionProfileVideoSettingsMV93X,
		}
		//[debug] Is Array: False
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
		SmartRetention:                 requestCameraUpdateNetworkCameraQualityRetentionProfileSmartRetention,
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
		ID: func() types.String {
			if response.ID != "" {
				return types.StringValue(response.ID)
			}
			return types.String{}
		}(),
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
		Name: func() types.String {
			if response.Name != "" {
				return types.StringValue(response.Name)
			}
			return types.String{}
		}(),
		NetworkID: func() types.String {
			if response.NetworkID != "" {
				return types.StringValue(response.NetworkID)
			}
			return types.String{}
		}(),
		RestrictedBandwidthModeEnabled: func() types.Bool {
			if response.RestrictedBandwidthModeEnabled != nil {
				return types.BoolValue(*response.RestrictedBandwidthModeEnabled)
			}
			return types.Bool{}
		}(),
		ScheduleID: func() types.String {
			if response.ScheduleID != "" {
				return types.StringValue(response.ScheduleID)
			}
			return types.String{}
		}(),
		SmartRetention: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileSmartRetentionRs {
			if response.SmartRetention != nil {
				return &ResponseCameraGetNetworkCameraQualityRetentionProfileSmartRetentionRs{
					Enabled: func() types.Bool {
						if response.SmartRetention.Enabled != nil {
							return types.BoolValue(*response.SmartRetention.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return nil
		}(),
		VideoSettings: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsRs {
			if response.VideoSettings != nil {
				return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsRs{
					MV12MV22MV72: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72Rs {
						if response.VideoSettings.MV12MV22MV72 != nil {
							return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12MV22MV72Rs{
								Quality: func() types.String {
									if response.VideoSettings.MV12MV22MV72.Quality != "" {
										return types.StringValue(response.VideoSettings.MV12MV22MV72.Quality)
									}
									return types.String{}
								}(),
								Resolution: func() types.String {
									if response.VideoSettings.MV12MV22MV72.Resolution != "" {
										return types.StringValue(response.VideoSettings.MV12MV22MV72.Resolution)
									}
									return types.String{}
								}(),
							}
						}
						return nil
					}(),
					MV12WE: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WERs {
						if response.VideoSettings.MV12WE != nil {
							return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV12WERs{
								Quality: func() types.String {
									if response.VideoSettings.MV12WE.Quality != "" {
										return types.StringValue(response.VideoSettings.MV12WE.Quality)
									}
									return types.String{}
								}(),
								Resolution: func() types.String {
									if response.VideoSettings.MV12WE.Resolution != "" {
										return types.StringValue(response.VideoSettings.MV12WE.Resolution)
									}
									return types.String{}
								}(),
							}
						}
						return nil
					}(),
					MV21MV71: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71Rs {
						if response.VideoSettings.MV21MV71 != nil {
							return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV21MV71Rs{
								Quality: func() types.String {
									if response.VideoSettings.MV21MV71.Quality != "" {
										return types.StringValue(response.VideoSettings.MV21MV71.Quality)
									}
									return types.String{}
								}(),
								Resolution: func() types.String {
									if response.VideoSettings.MV21MV71.Resolution != "" {
										return types.StringValue(response.VideoSettings.MV21MV71.Resolution)
									}
									return types.String{}
								}(),
							}
						}
						return nil
					}(),
					MV32: func() *ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32Rs {
						if response.VideoSettings.MV32 != nil {
							return &ResponseCameraGetNetworkCameraQualityRetentionProfileVideoSettingsMV32Rs{
								Quality: func() types.String {
									if response.VideoSettings.MV32.Quality != "" {
										return types.StringValue(response.VideoSettings.MV32.Quality)
									}
									return types.String{}
								}(),
								Resolution: func() types.String {
									if response.VideoSettings.MV32.Resolution != "" {
										return types.StringValue(response.VideoSettings.MV32.Resolution)
									}
									return types.String{}
								}(),
							}
						}
						return nil
					}(),
				}
			}
			return nil
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksCameraQualityRetentionProfilesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksCameraQualityRetentionProfilesRs)
}
