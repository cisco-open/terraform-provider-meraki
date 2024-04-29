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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
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
	_ resource.Resource              = &NetworksWirelessRfProfilesResource{}
	_ resource.ResourceWithConfigure = &NetworksWirelessRfProfilesResource{}
)

func NewNetworksWirelessRfProfilesResource() resource.Resource {
	return &NetworksWirelessRfProfilesResource{}
}

type NetworksWirelessRfProfilesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksWirelessRfProfilesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksWirelessRfProfilesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_wireless_rf_profiles"
}

func (r *NetworksWirelessRfProfilesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ap_band_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Settings that will be enabled if selectionType is set to 'ap'.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"band_operation_mode": schema.StringAttribute{
						MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'. Defaults to dual.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"2.4ghz",
								"5ghz",
								"6ghz",
								"dual",
								"multi",
							),
						},
					},
					"band_steering_enabled": schema.BoolAttribute{
						MarkdownDescription: `Steers client to most open band. Can be either true or false. Defaults to true.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"bands": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings related to all bands`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"enabled": schema.SetAttribute{
								MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Set{
									setplanmodifier.UseStateForUnknown(),
								},

								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"band_selection_type": schema.StringAttribute{
				MarkdownDescription: `Band selection can be set to either 'ssid' or 'ap'. This param is required on creation.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"ap",
						"ssid",
					),
				},
			},
			"client_balancing_enabled": schema.BoolAttribute{
				MarkdownDescription: `Steers client to best available access point. Can be either true or false. Defaults to true.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"five_ghz_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Settings related to 5Ghz band`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"channel_width": schema.StringAttribute{
						MarkdownDescription: `Sets channel width (MHz) for 5Ghz band. Can be one of 'auto', '20', '40' or '80'. Defaults to auto.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"max_power": schema.Int64Attribute{
						MarkdownDescription: `Sets max power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 30.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"min_bitrate": schema.Int64Attribute{
						MarkdownDescription: `Sets min bitrate (Mbps) of 5Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"min_power": schema.Int64Attribute{
						MarkdownDescription: `Sets min power (dBm) of 5Ghz band. Can be integer between 2 and 30. Defaults to 8.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"rxsop": schema.Int64Attribute{
						MarkdownDescription: `The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"valid_auto_channels": schema.SetAttribute{
						MarkdownDescription: `Sets valid auto channels for 5Ghz band. Can be one of '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161' or '165'.Defaults to [36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165].`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.Int64Type,
					},
				},
			},
			"flex_radios": schema.SingleNestedAttribute{
				MarkdownDescription: `Flex radio settings.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"by_model": schema.SetNestedAttribute{
						MarkdownDescription: `Flex radios by model.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{

								"bands": schema.SetAttribute{
									MarkdownDescription: `Band to use for each flex radio. For example, ['6'] will set the AP's first flex radio to 6 GHz`,
									Computed:            true,
									Optional:            true,
									PlanModifiers: []planmodifier.Set{
										setplanmodifier.UseStateForUnknown(),
									},

									ElementType: types.StringType,
								},
								"model": schema.StringAttribute{
									MarkdownDescription: `Model of the AP`,
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
			},
			"id": schema.StringAttribute{
				MarkdownDescription: `The name of the new profile. Must be unique.`,
				Computed:            true,
			},
			"min_bitrate_type": schema.StringAttribute{
				MarkdownDescription: `Minimum bitrate can be set to either 'band' or 'ssid'. Defaults to band.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"band",
						"ssid",
					),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name of the new profile. Must be unique. This param is required on creation.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `The network ID of the RF Profile`,
				Required:            true,
			},
			"per_ssid_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Per-SSID radio settings by number.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"status_0": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 0`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_1": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 1`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_10": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 10`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_11": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 11`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_12": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 12`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_13": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 13`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_14": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 14`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_2": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 2`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_3": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 3`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_4": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 4`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_5": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 5`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_6": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 6`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_7": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 7`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_8": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 8`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
					"status_9": schema.SingleNestedAttribute{
						MarkdownDescription: `Settings for SSID 9`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"band_operation_mode": schema.StringAttribute{
								MarkdownDescription: `Choice between 'dual', '2.4ghz', '5ghz', '6ghz' or 'multi'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.OneOf(
										"2.4ghz",
										"5ghz",
										"6ghz",
										"dual",
										"multi",
									),
								},
							},
							"band_steering_enabled": schema.BoolAttribute{
								MarkdownDescription: `Steers client to most open band between 2.4 GHz and 5 GHz. Can be either true or false.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"bands": schema.SingleNestedAttribute{
								MarkdownDescription: `Settings related to all bands`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.UseStateForUnknown(),
								},
								Attributes: map[string]schema.Attribute{

									"enabled": schema.SetAttribute{
										MarkdownDescription: `List of enabled bands. Can include ["2.4", "5", "6", "disabled"`,
										Computed:            true,
										Optional:            true,
										PlanModifiers: []planmodifier.Set{
											setplanmodifier.UseStateForUnknown(),
										},

										ElementType: types.StringType,
									},
								},
							},
							"min_bitrate": schema.Int64Attribute{
								MarkdownDescription: `Sets min bitrate (Mbps) of this SSID. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Int64{
									int64planmodifier.UseStateForUnknown(),
								},
								//                        Differents_types: `   parameter: schema.TypeFloat, item: schema.TypeInt`,
							},
							"name": schema.StringAttribute{
								MarkdownDescription: `Name of SSID`,
								Computed:            true,
							},
						},
					},
				},
			},
			"rf_profile_id": schema.StringAttribute{
				MarkdownDescription: `rfProfileId path parameter. Rf profile ID`,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"six_ghz_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Settings related to 6Ghz band. Only applicable to networks with 6Ghz capable APs`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"channel_width": schema.StringAttribute{
						MarkdownDescription: `Sets channel width (MHz) for 6Ghz band. Can be one of '0', '20', '40', '80' or '160'. Defaults to auto.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
					},
					"max_power": schema.Int64Attribute{
						MarkdownDescription: `Sets max power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 30.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"min_bitrate": schema.Int64Attribute{
						MarkdownDescription: `Sets min bitrate (Mbps) of 6Ghz band. Can be one of '6', '9', '12', '18', '24', '36', '48' or '54'. Defaults to 12.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"min_power": schema.Int64Attribute{
						MarkdownDescription: `Sets min power (dBm) of 6Ghz band. Can be integer between 2 and 30. Defaults to 8.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"rxsop": schema.Int64Attribute{
						MarkdownDescription: `The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"valid_auto_channels": schema.SetAttribute{
						MarkdownDescription: `Sets valid auto channels for 6Ghz band. Can be one of '1', '5', '9', '13', '17', '21', '25', '29', '33', '37', '41', '45', '49', '53', '57', '61', '65', '69', '73', '77', '81', '85', '89', '93', '97', '101', '105', '109', '113', '117', '121', '125', '129', '133', '137', '141', '145', '149', '153', '157', '161', '165', '169', '173', '177', '181', '185', '189', '193', '197', '201', '205', '209', '213', '217', '221', '225', '229' or '233'. Defaults to auto.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.Int64Type,
					},
				},
			},
			"transmission": schema.SingleNestedAttribute{
				MarkdownDescription: `Settings related to radio transmission.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"enabled": schema.BoolAttribute{
						MarkdownDescription: `Toggle for radio transmission. When false, radios will not transmit at all.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"two_four_ghz_settings": schema.SingleNestedAttribute{
				MarkdownDescription: `Settings related to 2.4Ghz band`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"ax_enabled": schema.BoolAttribute{
						MarkdownDescription: `Determines whether ax radio on 2.4Ghz band is on or off. Can be either true or false. If false, we highly recommend disabling band steering. Defaults to true.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Bool{
							boolplanmodifier.UseStateForUnknown(),
						},
					},
					"max_power": schema.Int64Attribute{
						MarkdownDescription: `Sets max power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 30.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"min_bitrate": schema.Float64Attribute{
						MarkdownDescription: `Sets min bitrate (Mbps) of 2.4Ghz band. Can be one of '1', '2', '5.5', '6', '9', '11', '12', '18', '24', '36', '48' or '54'. Defaults to 11.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"min_power": schema.Int64Attribute{
						MarkdownDescription: `Sets min power (dBm) of 2.4Ghz band. Can be integer between 2 and 30. Defaults to 5.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"rxsop": schema.Int64Attribute{
						MarkdownDescription: `The RX-SOP level controls the sensitivity of the radio. It is strongly recommended to use RX-SOP only after consulting a wireless expert. RX-SOP can be configured in the range of -65 to -95 (dBm). A value of null will reset this to the default.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"valid_auto_channels": schema.SetAttribute{
						MarkdownDescription: `Sets valid auto channels for 2.4Ghz band. Can be one of '1', '6' or '11'. Defaults to [1, 6, 11].`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.Int64Type,
					},
				},
			},
		},
	}
}

func (r *NetworksWirelessRfProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksWirelessRfProfilesRs

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
	// network_id
	vvName := data.Name.ValueString()
	//Items
	responseVerifyItem, restyResp1, err := r.client.Wireless.GetNetworkWirelessRfProfiles(vvNetworkID, nil)
	//Have Create
	if err != nil || restyResp1 == nil {
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessRfProfiles",
				err.Error(),
			)
			return
		}
	}
	if responseVerifyItem != nil {
		responseStruct := structToMap(responseVerifyItem)
		result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
		if result != nil {
			result2 := result.(map[string]interface{})
			vvRfProfileID, ok := result2["ID"].(string)
			if !ok {
				resp.Diagnostics.AddError(
					"Failure when parsing path parameter RfProfileID",
					err.Error(),
				)
				return
			}
			r.client.Wireless.UpdateNetworkWirelessRfProfile(vvNetworkID, vvRfProfileID, data.toSdkApiRequestUpdate(ctx))
			responseVerifyItem2, _, _ := r.client.Wireless.GetNetworkWirelessRfProfile(vvNetworkID, vvRfProfileID)
			if responseVerifyItem2 != nil {
				data = ResponseWirelessGetNetworkWirelessRfProfileItemToBodyRs(data, responseVerifyItem2, false)
				// Path params update assigned
				resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
				return
			}
		}
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	response, restyResp2, err := r.client.Wireless.CreateNetworkWirelessRfProfile(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkWirelessRfProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkWirelessRfProfile",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Wireless.GetNetworkWirelessRfProfiles(vvNetworkID, nil)
	// Has item and has items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessRfProfiles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessRfProfiles",
			err.Error(),
		)
		return
	}
	responseStruct := structToMap(responseGet)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	if result != nil {
		result2 := result.(map[string]interface{})
		vvRfProfileID, ok := result2["ID"].(string)
		if !ok {
			resp.Diagnostics.AddError(
				"Failure when parsing path parameter RfProfileID",
				err.Error(),
			)
			return
		}
		responseVerifyItem2, restyRespGet, err := r.client.Wireless.GetNetworkWirelessRfProfile(vvNetworkID, vvRfProfileID)
		if responseVerifyItem2 != nil && err == nil {
			data = ResponseWirelessGetNetworkWirelessRfProfileItemToBodyRs(data, responseVerifyItem2, false)
			resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
			return
		} else {
			if restyRespGet != nil {
				resp.Diagnostics.AddError(
					"Failure when executing GetNetworkWirelessRfProfile",
					err.Error(),
				)
				return
			}
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkWirelessRfProfile",
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

func (r *NetworksWirelessRfProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksWirelessRfProfilesRs

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
	// network_id
	vvRfProfileID := data.RfProfileID.ValueString()
	// rf_profile_id
	responseGet, restyRespGet, err := r.client.Wireless.GetNetworkWirelessRfProfile(vvNetworkID, vvRfProfileID)
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
				"Failure when executing GetNetworkWirelessRfProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkWirelessRfProfile",
			err.Error(),
		)
		return
	}

	data = ResponseWirelessGetNetworkWirelessRfProfileItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *NetworksWirelessRfProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	idParts := strings.Split(req.ID, ",")

	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			fmt.Sprintf("Expected import identifier with format: attr_one,attr_two. Got: %q", req.ID),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), idParts[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("rf_profile_id"), idParts[1])...)
}

func (r *NetworksWirelessRfProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksWirelessRfProfilesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvRfProfileID := data.RfProfileID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	response, restyResp2, err := r.client.Wireless.UpdateNetworkWirelessRfProfile(vvNetworkID, vvRfProfileID, dataRequest)
	if err != nil || restyResp2 == nil || response == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkWirelessRfProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkWirelessRfProfile",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksWirelessRfProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksWirelessRfProfilesRs
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
	vvRfProfileID := state.RfProfileID.ValueString()
	_, err := r.client.Wireless.DeleteNetworkWirelessRfProfile(vvNetworkID, vvRfProfileID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkWirelessRfProfile", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksWirelessRfProfilesRs struct {
	NetworkID              types.String                                                     `tfsdk:"network_id"`
	RfProfileID            types.String                                                     `tfsdk:"rf_profile_id"`
	ApBandSettings         *ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsRs     `tfsdk:"ap_band_settings"`
	BandSelectionType      types.String                                                     `tfsdk:"band_selection_type"`
	ClientBalancingEnabled types.Bool                                                       `tfsdk:"client_balancing_enabled"`
	FiveGhzSettings        *ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettingsRs    `tfsdk:"five_ghz_settings"`
	ID                     types.String                                                     `tfsdk:"id"`
	MinBitrateType         types.String                                                     `tfsdk:"min_bitrate_type"`
	Name                   types.String                                                     `tfsdk:"name"`
	PerSSIDSettings        *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettingsRs    `tfsdk:"per_ssid_settings"`
	SixGhzSettings         *ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettingsRs     `tfsdk:"six_ghz_settings"`
	Transmission           *ResponseWirelessGetNetworkWirelessRfProfileTransmissionRs       `tfsdk:"transmission"`
	TwoFourGhzSettings     *ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettingsRs `tfsdk:"two_four_ghz_settings"`
	FlexRadios             *RequestWirelessUpdateNetworkWirelessRfProfileFlexRadiosRs       `tfsdk:"flex_radios"`
}

type ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsRs struct {
	BandOperationMode   types.String                                                      `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                        `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsBandsRs `tfsdk:"bands"`
}

type ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsBandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettingsRs struct {
	ChannelWidth      types.String `tfsdk:"channel_width"`
	MaxPower          types.Int64  `tfsdk:"max_power"`
	MinBitrate        types.Int64  `tfsdk:"min_bitrate"`
	MinPower          types.Int64  `tfsdk:"min_power"`
	Rxsop             types.Int64  `tfsdk:"rxsop"`
	ValidAutoChannels types.Set    `tfsdk:"valid_auto_channels"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettingsRs struct {
	Status0  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0Rs  `tfsdk:"status_0"`
	Status1  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1Rs  `tfsdk:"status_1"`
	Status10 *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10Rs `tfsdk:"status_10"`
	Status11 *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11Rs `tfsdk:"status_11"`
	Status12 *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12Rs `tfsdk:"status_12"`
	Status13 *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13Rs `tfsdk:"status_13"`
	Status14 *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14Rs `tfsdk:"status_14"`
	Status2  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2Rs  `tfsdk:"status_2"`
	Status3  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3Rs  `tfsdk:"status_3"`
	Status4  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4Rs  `tfsdk:"status_4"`
	Status5  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5Rs  `tfsdk:"status_5"`
	Status6  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6Rs  `tfsdk:"status_6"`
	Status7  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7Rs  `tfsdk:"status_7"`
	Status8  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8Rs  `tfsdk:"status_8"`
	Status9  *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9Rs  `tfsdk:"status_9"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0Rs struct {
	BandOperationMode   types.String                                                        `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                          `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                         `tfsdk:"min_bitrate"`
	Name                types.String                                                        `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1Rs struct {
	BandOperationMode   types.String                                                        `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                          `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                         `tfsdk:"min_bitrate"`
	Name                types.String                                                        `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10Rs struct {
	BandOperationMode   types.String                                                         `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                           `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                          `tfsdk:"min_bitrate"`
	Name                types.String                                                         `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11Rs struct {
	BandOperationMode   types.String                                                         `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                           `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                          `tfsdk:"min_bitrate"`
	Name                types.String                                                         `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12Rs struct {
	BandOperationMode   types.String                                                         `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                           `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                          `tfsdk:"min_bitrate"`
	Name                types.String                                                         `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13Rs struct {
	BandOperationMode   types.String                                                         `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                           `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                          `tfsdk:"min_bitrate"`
	Name                types.String                                                         `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14Rs struct {
	BandOperationMode   types.String                                                         `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                           `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                          `tfsdk:"min_bitrate"`
	Name                types.String                                                         `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2Rs struct {
	BandOperationMode   types.String                                                        `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                          `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                         `tfsdk:"min_bitrate"`
	Name                types.String                                                        `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3Rs struct {
	BandOperationMode   types.String                                                        `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                          `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                         `tfsdk:"min_bitrate"`
	Name                types.String                                                        `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4Rs struct {
	BandOperationMode   types.String                                                        `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                          `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                         `tfsdk:"min_bitrate"`
	Name                types.String                                                        `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5Rs struct {
	BandOperationMode   types.String                                                        `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                          `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                         `tfsdk:"min_bitrate"`
	Name                types.String                                                        `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6Rs struct {
	BandOperationMode   types.String                                                        `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                          `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                         `tfsdk:"min_bitrate"`
	Name                types.String                                                        `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7Rs struct {
	BandOperationMode   types.String                                                        `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                          `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                         `tfsdk:"min_bitrate"`
	Name                types.String                                                        `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8Rs struct {
	BandOperationMode   types.String                                                        `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                          `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                         `tfsdk:"min_bitrate"`
	Name                types.String                                                        `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9Rs struct {
	BandOperationMode   types.String                                                        `tfsdk:"band_operation_mode"`
	BandSteeringEnabled types.Bool                                                          `tfsdk:"band_steering_enabled"`
	Bands               *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9BandsRs `tfsdk:"bands"`
	MinBitrate          types.Int64                                                         `tfsdk:"min_bitrate"`
	Name                types.String                                                        `tfsdk:"name"`
}

type ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9BandsRs struct {
	Enabled types.Set `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettingsRs struct {
	ChannelWidth      types.String `tfsdk:"channel_width"`
	MaxPower          types.Int64  `tfsdk:"max_power"`
	MinBitrate        types.Int64  `tfsdk:"min_bitrate"`
	MinPower          types.Int64  `tfsdk:"min_power"`
	Rxsop             types.Int64  `tfsdk:"rxsop"`
	ValidAutoChannels types.Set    `tfsdk:"valid_auto_channels"`
}

type ResponseWirelessGetNetworkWirelessRfProfileTransmissionRs struct {
	Enabled types.Bool `tfsdk:"enabled"`
}

type ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettingsRs struct {
	AxEnabled         types.Bool    `tfsdk:"ax_enabled"`
	MaxPower          types.Int64   `tfsdk:"max_power"`
	MinBitrate        types.Float64 `tfsdk:"min_bitrate"`
	MinPower          types.Int64   `tfsdk:"min_power"`
	Rxsop             types.Int64   `tfsdk:"rxsop"`
	ValidAutoChannels types.Set     `tfsdk:"valid_auto_channels"`
}

type RequestWirelessUpdateNetworkWirelessRfProfileFlexRadiosRs struct {
	ByModel *[]RequestWirelessUpdateNetworkWirelessRfProfileFlexRadiosByModelRs `tfsdk:"by_model"`
}

type RequestWirelessUpdateNetworkWirelessRfProfileFlexRadiosByModelRs struct {
	Bands types.Set    `tfsdk:"bands"`
	Model types.String `tfsdk:"model"`
}

// FromBody
func (r *NetworksWirelessRfProfilesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfile {
	emptyString := ""
	var requestWirelessCreateNetworkWirelessRfProfileApBandSettings *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileApBandSettings
	if r.ApBandSettings != nil {
		bandOperationMode := r.ApBandSettings.BandOperationMode.ValueString()
		bandSteeringEnabled := func() *bool {
			if !r.ApBandSettings.BandSteeringEnabled.IsUnknown() && !r.ApBandSettings.BandSteeringEnabled.IsNull() {
				return r.ApBandSettings.BandSteeringEnabled.ValueBoolPointer()
			}
			return nil
		}()
		var requestWirelessCreateNetworkWirelessRfProfileApBandSettingsBands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileApBandSettingsBands
		if r.ApBandSettings.Bands != nil {
			var enabled []string = nil
			//Hoola aqui
			r.ApBandSettings.Bands.Enabled.ElementsAs(ctx, &enabled, false)
			requestWirelessCreateNetworkWirelessRfProfileApBandSettingsBands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileApBandSettingsBands{
				Enabled: enabled,
			}
		}
		requestWirelessCreateNetworkWirelessRfProfileApBandSettings = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileApBandSettings{
			BandOperationMode:   bandOperationMode,
			BandSteeringEnabled: bandSteeringEnabled,
			Bands:               requestWirelessCreateNetworkWirelessRfProfileApBandSettingsBands,
		}
	}
	bandSelectionType := new(string)
	if !r.BandSelectionType.IsUnknown() && !r.BandSelectionType.IsNull() {
		*bandSelectionType = r.BandSelectionType.ValueString()
	} else {
		bandSelectionType = &emptyString
	}
	clientBalancingEnabled := new(bool)
	if !r.ClientBalancingEnabled.IsUnknown() && !r.ClientBalancingEnabled.IsNull() {
		*clientBalancingEnabled = r.ClientBalancingEnabled.ValueBool()
	} else {
		clientBalancingEnabled = nil
	}
	var requestWirelessCreateNetworkWirelessRfProfileFiveGhzSettings *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileFiveGhzSettings
	if r.FiveGhzSettings != nil {
		channelWidth := r.FiveGhzSettings.ChannelWidth.ValueString()
		maxPower := func() *int64 {
			if !r.FiveGhzSettings.MaxPower.IsUnknown() && !r.FiveGhzSettings.MaxPower.IsNull() {
				return r.FiveGhzSettings.MaxPower.ValueInt64Pointer()
			}
			return nil
		}()
		minBitrate := func() *int64 {
			if !r.FiveGhzSettings.MinBitrate.IsUnknown() && !r.FiveGhzSettings.MinBitrate.IsNull() {
				return r.FiveGhzSettings.MinBitrate.ValueInt64Pointer()
			}
			return nil
		}()
		minPower := func() *int64 {
			if !r.FiveGhzSettings.MinPower.IsUnknown() && !r.FiveGhzSettings.MinPower.IsNull() {
				return r.FiveGhzSettings.MinPower.ValueInt64Pointer()
			}
			return nil
		}()
		rxsop := func() *int64 {
			if !r.FiveGhzSettings.Rxsop.IsUnknown() && !r.FiveGhzSettings.Rxsop.IsNull() {
				return r.FiveGhzSettings.Rxsop.ValueInt64Pointer()
			}
			return nil
		}()
		var validAutoChannels *[]int = nil
		//Hoola aqui
		r.FiveGhzSettings.ValidAutoChannels.ElementsAs(ctx, &validAutoChannels, false)
		requestWirelessCreateNetworkWirelessRfProfileFiveGhzSettings = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileFiveGhzSettings{
			ChannelWidth:      channelWidth,
			MaxPower:          int64ToIntPointer(maxPower),
			MinBitrate:        int64ToIntPointer(minBitrate),
			MinPower:          int64ToIntPointer(minPower),
			Rxsop:             int64ToIntPointer(rxsop),
			ValidAutoChannels: validAutoChannels,
		}
	}
	var requestWirelessCreateNetworkWirelessRfProfileFlexRadios *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileFlexRadios
	if r.FlexRadios != nil {
		var requestWirelessCreateNetworkWirelessRfProfileFlexRadiosByModel []merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileFlexRadiosByModel
		if r.FlexRadios.ByModel != nil {
			for _, rItem1 := range *r.FlexRadios.ByModel { //FlexRadios.ByModel// name: byModel
				var bands []string = nil
				//Hoola aqui
				rItem1.Bands.ElementsAs(ctx, &bands, false)
				model := rItem1.Model.ValueString()
				requestWirelessCreateNetworkWirelessRfProfileFlexRadiosByModel = append(requestWirelessCreateNetworkWirelessRfProfileFlexRadiosByModel, merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileFlexRadiosByModel{
					Bands: bands,
					Model: model,
				})
			}
		}
		requestWirelessCreateNetworkWirelessRfProfileFlexRadios = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileFlexRadios{
			ByModel: func() *[]merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileFlexRadiosByModel {
				if len(requestWirelessCreateNetworkWirelessRfProfileFlexRadiosByModel) > 0 {
					return &requestWirelessCreateNetworkWirelessRfProfileFlexRadiosByModel
				}
				return nil
			}(),
		}
	}
	minBitrateType := new(string)
	if !r.MinBitrateType.IsUnknown() && !r.MinBitrateType.IsNull() {
		*minBitrateType = r.MinBitrateType.ValueString()
	} else {
		minBitrateType = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings
	if r.PerSSIDSettings != nil {
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0
		if r.PerSSIDSettings.Status0 != nil {
			bandOperationMode := r.PerSSIDSettings.Status0.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status0.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status0.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status0.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0Bands
			if r.PerSSIDSettings.Status0.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status0.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status0.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1
		if r.PerSSIDSettings.Status1 != nil {
			bandOperationMode := r.PerSSIDSettings.Status1.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status1.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status1.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status1.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1Bands
			if r.PerSSIDSettings.Status1.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status1.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status1.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10
		if r.PerSSIDSettings.Status10 != nil {
			bandOperationMode := r.PerSSIDSettings.Status10.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status10.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status10.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status10.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10Bands
			if r.PerSSIDSettings.Status10.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status10.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status10.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11
		if r.PerSSIDSettings.Status11 != nil {
			bandOperationMode := r.PerSSIDSettings.Status11.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status11.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status11.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status11.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11Bands
			if r.PerSSIDSettings.Status11.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status11.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status11.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12
		if r.PerSSIDSettings.Status12 != nil {
			bandOperationMode := r.PerSSIDSettings.Status12.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status12.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status12.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status12.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12Bands
			if r.PerSSIDSettings.Status12.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status12.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status12.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13
		if r.PerSSIDSettings.Status13 != nil {
			bandOperationMode := r.PerSSIDSettings.Status13.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status13.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status13.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status13.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13Bands
			if r.PerSSIDSettings.Status13.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status13.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status13.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14
		if r.PerSSIDSettings.Status14 != nil {
			bandOperationMode := r.PerSSIDSettings.Status14.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status14.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status14.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status14.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14Bands
			if r.PerSSIDSettings.Status14.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status14.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status14.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2
		if r.PerSSIDSettings.Status2 != nil {
			bandOperationMode := r.PerSSIDSettings.Status2.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status2.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status2.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status2.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2Bands
			if r.PerSSIDSettings.Status2.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status2.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status2.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3
		if r.PerSSIDSettings.Status3 != nil {
			bandOperationMode := r.PerSSIDSettings.Status3.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status3.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status3.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status3.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3Bands
			if r.PerSSIDSettings.Status3.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status3.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status3.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4
		if r.PerSSIDSettings.Status4 != nil {
			bandOperationMode := r.PerSSIDSettings.Status4.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status4.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status4.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status4.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4Bands
			if r.PerSSIDSettings.Status4.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status4.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status4.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5
		if r.PerSSIDSettings.Status5 != nil {
			bandOperationMode := r.PerSSIDSettings.Status5.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status5.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status5.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status5.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5Bands
			if r.PerSSIDSettings.Status5.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status5.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status5.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6
		if r.PerSSIDSettings.Status6 != nil {
			bandOperationMode := r.PerSSIDSettings.Status6.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status6.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status6.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status6.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6Bands
			if r.PerSSIDSettings.Status6.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status6.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status6.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7
		if r.PerSSIDSettings.Status7 != nil {
			bandOperationMode := r.PerSSIDSettings.Status7.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status7.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status7.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status7.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7Bands
			if r.PerSSIDSettings.Status7.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status7.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status7.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8
		if r.PerSSIDSettings.Status8 != nil {
			bandOperationMode := r.PerSSIDSettings.Status8.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status8.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status8.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status8.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8Bands
			if r.PerSSIDSettings.Status8.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status8.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status8.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9 *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9
		if r.PerSSIDSettings.Status9 != nil {
			bandOperationMode := r.PerSSIDSettings.Status9.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status9.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status9.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status9.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9Bands *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9Bands
			if r.PerSSIDSettings.Status9.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status9.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9Bands = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status9.MinBitrate.ValueInt64Pointer())
			requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9 = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9Bands,
				MinBitrate:          minBitrate,
			}
		}
		requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings{
			Status0:  requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings0,
			Status1:  requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings1,
			Status10: requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings10,
			Status11: requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings11,
			Status12: requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings12,
			Status13: requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings13,
			Status14: requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings14,
			Status2:  requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings2,
			Status3:  requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings3,
			Status4:  requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings4,
			Status5:  requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings5,
			Status6:  requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings6,
			Status7:  requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings7,
			Status8:  requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings8,
			Status9:  requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings9,
		}
	}
	var requestWirelessCreateNetworkWirelessRfProfileSixGhzSettings *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileSixGhzSettings
	if r.SixGhzSettings != nil {
		channelWidth := r.SixGhzSettings.ChannelWidth.ValueString()
		maxPower := func() *int64 {
			if !r.SixGhzSettings.MaxPower.IsUnknown() && !r.SixGhzSettings.MaxPower.IsNull() {
				return r.SixGhzSettings.MaxPower.ValueInt64Pointer()
			}
			return nil
		}()
		minBitrate := func() *int64 {
			if !r.SixGhzSettings.MinBitrate.IsUnknown() && !r.SixGhzSettings.MinBitrate.IsNull() {
				return r.SixGhzSettings.MinBitrate.ValueInt64Pointer()
			}
			return nil
		}()
		minPower := func() *int64 {
			if !r.SixGhzSettings.MinPower.IsUnknown() && !r.SixGhzSettings.MinPower.IsNull() {
				return r.SixGhzSettings.MinPower.ValueInt64Pointer()
			}
			return nil
		}()
		rxsop := func() *int64 {
			if !r.SixGhzSettings.Rxsop.IsUnknown() && !r.SixGhzSettings.Rxsop.IsNull() {
				return r.SixGhzSettings.Rxsop.ValueInt64Pointer()
			}
			return nil
		}()
		var validAutoChannels *[]int = nil
		//Hoola aqui
		r.SixGhzSettings.ValidAutoChannels.ElementsAs(ctx, &validAutoChannels, false)
		requestWirelessCreateNetworkWirelessRfProfileSixGhzSettings = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileSixGhzSettings{
			ChannelWidth:      channelWidth,
			MaxPower:          int64ToIntPointer(maxPower),
			MinBitrate:        int64ToIntPointer(minBitrate),
			MinPower:          int64ToIntPointer(minPower),
			Rxsop:             int64ToIntPointer(rxsop),
			ValidAutoChannels: validAutoChannels,
		}
	}
	var requestWirelessCreateNetworkWirelessRfProfileTransmission *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileTransmission
	if r.Transmission != nil {
		enabled := func() *bool {
			if !r.Transmission.Enabled.IsUnknown() && !r.Transmission.Enabled.IsNull() {
				return r.Transmission.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestWirelessCreateNetworkWirelessRfProfileTransmission = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileTransmission{
			Enabled: enabled,
		}
	}
	var requestWirelessCreateNetworkWirelessRfProfileTwoFourGhzSettings *merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileTwoFourGhzSettings
	if r.TwoFourGhzSettings != nil {
		axEnabled := func() *bool {
			if !r.TwoFourGhzSettings.AxEnabled.IsUnknown() && !r.TwoFourGhzSettings.AxEnabled.IsNull() {
				return r.TwoFourGhzSettings.AxEnabled.ValueBoolPointer()
			}
			return nil
		}()
		maxPower := func() *int64 {
			if !r.TwoFourGhzSettings.MaxPower.IsUnknown() && !r.TwoFourGhzSettings.MaxPower.IsNull() {
				return r.TwoFourGhzSettings.MaxPower.ValueInt64Pointer()
			}
			return nil
		}()
		minBitrate := r.TwoFourGhzSettings.MinBitrate.ValueFloat64Pointer()
		minPower := func() *int64 {
			if !r.TwoFourGhzSettings.MinPower.IsUnknown() && !r.TwoFourGhzSettings.MinPower.IsNull() {
				return r.TwoFourGhzSettings.MinPower.ValueInt64Pointer()
			}
			return nil
		}()
		rxsop := func() *int64 {
			if !r.TwoFourGhzSettings.Rxsop.IsUnknown() && !r.TwoFourGhzSettings.Rxsop.IsNull() {
				return r.TwoFourGhzSettings.Rxsop.ValueInt64Pointer()
			}
			return nil
		}()
		var validAutoChannels *[]int = nil
		//Hoola aqui
		r.TwoFourGhzSettings.ValidAutoChannels.ElementsAs(ctx, &validAutoChannels, false)
		requestWirelessCreateNetworkWirelessRfProfileTwoFourGhzSettings = &merakigosdk.RequestWirelessCreateNetworkWirelessRfProfileTwoFourGhzSettings{
			AxEnabled:         axEnabled,
			MaxPower:          int64ToIntPointer(maxPower),
			MinBitrate:        minBitrate,
			MinPower:          int64ToIntPointer(minPower),
			Rxsop:             int64ToIntPointer(rxsop),
			ValidAutoChannels: validAutoChannels,
		}
	}
	out := merakigosdk.RequestWirelessCreateNetworkWirelessRfProfile{
		ApBandSettings:         requestWirelessCreateNetworkWirelessRfProfileApBandSettings,
		BandSelectionType:      *bandSelectionType,
		ClientBalancingEnabled: clientBalancingEnabled,
		FiveGhzSettings:        requestWirelessCreateNetworkWirelessRfProfileFiveGhzSettings,
		FlexRadios:             requestWirelessCreateNetworkWirelessRfProfileFlexRadios,
		MinBitrateType:         *minBitrateType,
		Name:                   *name,
		PerSSIDSettings:        requestWirelessCreateNetworkWirelessRfProfilePerSSIDSettings,
		SixGhzSettings:         requestWirelessCreateNetworkWirelessRfProfileSixGhzSettings,
		Transmission:           requestWirelessCreateNetworkWirelessRfProfileTransmission,
		TwoFourGhzSettings:     requestWirelessCreateNetworkWirelessRfProfileTwoFourGhzSettings,
	}
	return &out
}
func (r *NetworksWirelessRfProfilesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfile {
	emptyString := ""
	var requestWirelessUpdateNetworkWirelessRfProfileApBandSettings *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileApBandSettings
	if r.ApBandSettings != nil {
		bandOperationMode := r.ApBandSettings.BandOperationMode.ValueString()
		bandSteeringEnabled := func() *bool {
			if !r.ApBandSettings.BandSteeringEnabled.IsUnknown() && !r.ApBandSettings.BandSteeringEnabled.IsNull() {
				return r.ApBandSettings.BandSteeringEnabled.ValueBoolPointer()
			}
			return nil
		}()
		var requestWirelessUpdateNetworkWirelessRfProfileApBandSettingsBands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileApBandSettingsBands
		if r.ApBandSettings.Bands != nil {
			var enabled []string = nil
			//Hoola aqui
			r.ApBandSettings.Bands.Enabled.ElementsAs(ctx, &enabled, false)
			requestWirelessUpdateNetworkWirelessRfProfileApBandSettingsBands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileApBandSettingsBands{
				Enabled: enabled,
			}
		}
		requestWirelessUpdateNetworkWirelessRfProfileApBandSettings = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileApBandSettings{
			BandOperationMode:   bandOperationMode,
			BandSteeringEnabled: bandSteeringEnabled,
			Bands:               requestWirelessUpdateNetworkWirelessRfProfileApBandSettingsBands,
		}
	}
	bandSelectionType := new(string)
	if !r.BandSelectionType.IsUnknown() && !r.BandSelectionType.IsNull() {
		*bandSelectionType = r.BandSelectionType.ValueString()
	} else {
		bandSelectionType = &emptyString
	}
	clientBalancingEnabled := new(bool)
	if !r.ClientBalancingEnabled.IsUnknown() && !r.ClientBalancingEnabled.IsNull() {
		*clientBalancingEnabled = r.ClientBalancingEnabled.ValueBool()
	} else {
		clientBalancingEnabled = nil
	}
	var requestWirelessUpdateNetworkWirelessRfProfileFiveGhzSettings *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileFiveGhzSettings
	if r.FiveGhzSettings != nil {
		channelWidth := r.FiveGhzSettings.ChannelWidth.ValueString()
		maxPower := func() *int64 {
			if !r.FiveGhzSettings.MaxPower.IsUnknown() && !r.FiveGhzSettings.MaxPower.IsNull() {
				return r.FiveGhzSettings.MaxPower.ValueInt64Pointer()
			}
			return nil
		}()
		minBitrate := func() *int64 {
			if !r.FiveGhzSettings.MinBitrate.IsUnknown() && !r.FiveGhzSettings.MinBitrate.IsNull() {
				return r.FiveGhzSettings.MinBitrate.ValueInt64Pointer()
			}
			return nil
		}()
		minPower := func() *int64 {
			if !r.FiveGhzSettings.MinPower.IsUnknown() && !r.FiveGhzSettings.MinPower.IsNull() {
				return r.FiveGhzSettings.MinPower.ValueInt64Pointer()
			}
			return nil
		}()
		rxsop := func() *int64 {
			if !r.FiveGhzSettings.Rxsop.IsUnknown() && !r.FiveGhzSettings.Rxsop.IsNull() {
				return r.FiveGhzSettings.Rxsop.ValueInt64Pointer()
			}
			return nil
		}()
		var validAutoChannels *[]int = nil
		//Hoola aqui
		r.FiveGhzSettings.ValidAutoChannels.ElementsAs(ctx, &validAutoChannels, false)
		requestWirelessUpdateNetworkWirelessRfProfileFiveGhzSettings = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileFiveGhzSettings{
			ChannelWidth:      channelWidth,
			MaxPower:          int64ToIntPointer(maxPower),
			MinBitrate:        int64ToIntPointer(minBitrate),
			MinPower:          int64ToIntPointer(minPower),
			Rxsop:             int64ToIntPointer(rxsop),
			ValidAutoChannels: validAutoChannels,
		}
	}
	var requestWirelessUpdateNetworkWirelessRfProfileFlexRadios *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileFlexRadios
	if r.FlexRadios != nil {
		var requestWirelessUpdateNetworkWirelessRfProfileFlexRadiosByModel []merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileFlexRadiosByModel
		if r.FlexRadios.ByModel != nil {
			for _, rItem1 := range *r.FlexRadios.ByModel { //FlexRadios.ByModel// name: byModel
				var bands []string = nil
				//Hoola aqui
				rItem1.Bands.ElementsAs(ctx, &bands, false)
				model := rItem1.Model.ValueString()
				requestWirelessUpdateNetworkWirelessRfProfileFlexRadiosByModel = append(requestWirelessUpdateNetworkWirelessRfProfileFlexRadiosByModel, merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileFlexRadiosByModel{
					Bands: bands,
					Model: model,
				})
			}
		}
		requestWirelessUpdateNetworkWirelessRfProfileFlexRadios = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileFlexRadios{
			ByModel: func() *[]merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileFlexRadiosByModel {
				if len(requestWirelessUpdateNetworkWirelessRfProfileFlexRadiosByModel) > 0 {
					return &requestWirelessUpdateNetworkWirelessRfProfileFlexRadiosByModel
				}
				return nil
			}(),
		}
	}
	minBitrateType := new(string)
	if !r.MinBitrateType.IsUnknown() && !r.MinBitrateType.IsNull() {
		*minBitrateType = r.MinBitrateType.ValueString()
	} else {
		minBitrateType = &emptyString
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings
	if r.PerSSIDSettings != nil {
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0
		if r.PerSSIDSettings.Status0 != nil {
			bandOperationMode := r.PerSSIDSettings.Status0.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status0.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status0.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status0.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0Bands
			if r.PerSSIDSettings.Status0.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status0.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status0.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1
		if r.PerSSIDSettings.Status1 != nil {
			bandOperationMode := r.PerSSIDSettings.Status1.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status1.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status1.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status1.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1Bands
			if r.PerSSIDSettings.Status1.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status1.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status1.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10
		if r.PerSSIDSettings.Status10 != nil {
			bandOperationMode := r.PerSSIDSettings.Status10.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status10.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status10.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status10.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10Bands
			if r.PerSSIDSettings.Status10.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status10.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status10.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11
		if r.PerSSIDSettings.Status11 != nil {
			bandOperationMode := r.PerSSIDSettings.Status11.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status11.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status11.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status11.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11Bands
			if r.PerSSIDSettings.Status11.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status11.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status11.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12
		if r.PerSSIDSettings.Status12 != nil {
			bandOperationMode := r.PerSSIDSettings.Status12.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status12.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status12.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status12.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12Bands
			if r.PerSSIDSettings.Status12.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status12.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status12.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13
		if r.PerSSIDSettings.Status13 != nil {
			bandOperationMode := r.PerSSIDSettings.Status13.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status13.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status13.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status13.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13Bands
			if r.PerSSIDSettings.Status13.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status13.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status13.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14
		if r.PerSSIDSettings.Status14 != nil {
			bandOperationMode := r.PerSSIDSettings.Status14.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status14.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status14.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status14.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14Bands
			if r.PerSSIDSettings.Status14.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status14.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status14.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2
		if r.PerSSIDSettings.Status2 != nil {
			bandOperationMode := r.PerSSIDSettings.Status2.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status2.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status2.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status2.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2Bands
			if r.PerSSIDSettings.Status2.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status2.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status2.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3
		if r.PerSSIDSettings.Status3 != nil {
			bandOperationMode := r.PerSSIDSettings.Status3.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status3.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status3.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status3.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3Bands
			if r.PerSSIDSettings.Status3.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status3.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status3.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4
		if r.PerSSIDSettings.Status4 != nil {
			bandOperationMode := r.PerSSIDSettings.Status4.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status4.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status4.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status4.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4Bands
			if r.PerSSIDSettings.Status4.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status4.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status4.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5
		if r.PerSSIDSettings.Status5 != nil {
			bandOperationMode := r.PerSSIDSettings.Status5.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status5.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status5.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status5.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5Bands
			if r.PerSSIDSettings.Status5.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status5.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status5.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6
		if r.PerSSIDSettings.Status6 != nil {
			bandOperationMode := r.PerSSIDSettings.Status6.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status6.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status6.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status6.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6Bands
			if r.PerSSIDSettings.Status6.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status6.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status6.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7
		if r.PerSSIDSettings.Status7 != nil {
			bandOperationMode := r.PerSSIDSettings.Status7.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status7.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status7.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status7.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7Bands
			if r.PerSSIDSettings.Status7.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status7.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status7.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8
		if r.PerSSIDSettings.Status8 != nil {
			bandOperationMode := r.PerSSIDSettings.Status8.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status8.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status8.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status8.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8Bands
			if r.PerSSIDSettings.Status8.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status8.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status8.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8Bands,
				MinBitrate:          minBitrate,
			}
		}
		var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9 *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9
		if r.PerSSIDSettings.Status9 != nil {
			bandOperationMode := r.PerSSIDSettings.Status9.BandOperationMode.ValueString()
			bandSteeringEnabled := func() *bool {
				if !r.PerSSIDSettings.Status9.BandSteeringEnabled.IsUnknown() && !r.PerSSIDSettings.Status9.BandSteeringEnabled.IsNull() {
					return r.PerSSIDSettings.Status9.BandSteeringEnabled.ValueBoolPointer()
				}
				return nil
			}()
			var requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9Bands *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9Bands
			if r.PerSSIDSettings.Status9.Bands != nil {
				var enabled []string = nil
				//Hoola aqui
				r.PerSSIDSettings.Status9.Bands.Enabled.ElementsAs(ctx, &enabled, false)
				requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9Bands = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9Bands{
					Enabled: enabled,
				}
			}
			minBitrate := int64ToFloat(r.PerSSIDSettings.Status9.MinBitrate.ValueInt64Pointer())
			requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9 = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9{
				BandOperationMode:   bandOperationMode,
				BandSteeringEnabled: bandSteeringEnabled,
				Bands:               requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9Bands,
				MinBitrate:          minBitrate,
			}
		}
		requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings{
			Status0:  requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings0,
			Status1:  requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings1,
			Status10: requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings10,
			Status11: requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings11,
			Status12: requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings12,
			Status13: requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings13,
			Status14: requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings14,
			Status2:  requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings2,
			Status3:  requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings3,
			Status4:  requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings4,
			Status5:  requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings5,
			Status6:  requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings6,
			Status7:  requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings7,
			Status8:  requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings8,
			Status9:  requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings9,
		}
	}
	var requestWirelessUpdateNetworkWirelessRfProfileSixGhzSettings *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileSixGhzSettings
	if r.SixGhzSettings != nil {
		channelWidth := r.SixGhzSettings.ChannelWidth.ValueString()
		maxPower := func() *int64 {
			if !r.SixGhzSettings.MaxPower.IsUnknown() && !r.SixGhzSettings.MaxPower.IsNull() {
				return r.SixGhzSettings.MaxPower.ValueInt64Pointer()
			}
			return nil
		}()
		minBitrate := func() *int64 {
			if !r.SixGhzSettings.MinBitrate.IsUnknown() && !r.SixGhzSettings.MinBitrate.IsNull() {
				return r.SixGhzSettings.MinBitrate.ValueInt64Pointer()
			}
			return nil
		}()
		minPower := func() *int64 {
			if !r.SixGhzSettings.MinPower.IsUnknown() && !r.SixGhzSettings.MinPower.IsNull() {
				return r.SixGhzSettings.MinPower.ValueInt64Pointer()
			}
			return nil
		}()
		rxsop := func() *int64 {
			if !r.SixGhzSettings.Rxsop.IsUnknown() && !r.SixGhzSettings.Rxsop.IsNull() {
				return r.SixGhzSettings.Rxsop.ValueInt64Pointer()
			}
			return nil
		}()
		var validAutoChannels *[]int = nil
		//Hoola aqui
		r.SixGhzSettings.ValidAutoChannels.ElementsAs(ctx, &validAutoChannels, false)
		requestWirelessUpdateNetworkWirelessRfProfileSixGhzSettings = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileSixGhzSettings{
			ChannelWidth:      channelWidth,
			MaxPower:          int64ToIntPointer(maxPower),
			MinBitrate:        int64ToIntPointer(minBitrate),
			MinPower:          int64ToIntPointer(minPower),
			Rxsop:             int64ToIntPointer(rxsop),
			ValidAutoChannels: validAutoChannels,
		}
	}
	var requestWirelessUpdateNetworkWirelessRfProfileTransmission *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileTransmission
	if r.Transmission != nil {
		enabled := func() *bool {
			if !r.Transmission.Enabled.IsUnknown() && !r.Transmission.Enabled.IsNull() {
				return r.Transmission.Enabled.ValueBoolPointer()
			}
			return nil
		}()
		requestWirelessUpdateNetworkWirelessRfProfileTransmission = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileTransmission{
			Enabled: enabled,
		}
	}
	var requestWirelessUpdateNetworkWirelessRfProfileTwoFourGhzSettings *merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileTwoFourGhzSettings
	if r.TwoFourGhzSettings != nil {
		axEnabled := func() *bool {
			if !r.TwoFourGhzSettings.AxEnabled.IsUnknown() && !r.TwoFourGhzSettings.AxEnabled.IsNull() {
				return r.TwoFourGhzSettings.AxEnabled.ValueBoolPointer()
			}
			return nil
		}()
		maxPower := func() *int64 {
			if !r.TwoFourGhzSettings.MaxPower.IsUnknown() && !r.TwoFourGhzSettings.MaxPower.IsNull() {
				return r.TwoFourGhzSettings.MaxPower.ValueInt64Pointer()
			}
			return nil
		}()
		minBitrate := r.TwoFourGhzSettings.MinBitrate.ValueFloat64Pointer()
		minPower := func() *int64 {
			if !r.TwoFourGhzSettings.MinPower.IsUnknown() && !r.TwoFourGhzSettings.MinPower.IsNull() {
				return r.TwoFourGhzSettings.MinPower.ValueInt64Pointer()
			}
			return nil
		}()
		rxsop := func() *int64 {
			if !r.TwoFourGhzSettings.Rxsop.IsUnknown() && !r.TwoFourGhzSettings.Rxsop.IsNull() {
				return r.TwoFourGhzSettings.Rxsop.ValueInt64Pointer()
			}
			return nil
		}()
		var validAutoChannels *[]int = nil
		//Hoola aqui
		r.TwoFourGhzSettings.ValidAutoChannels.ElementsAs(ctx, &validAutoChannels, false)
		requestWirelessUpdateNetworkWirelessRfProfileTwoFourGhzSettings = &merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfileTwoFourGhzSettings{
			AxEnabled:         axEnabled,
			MaxPower:          int64ToIntPointer(maxPower),
			MinBitrate:        minBitrate,
			MinPower:          int64ToIntPointer(minPower),
			Rxsop:             int64ToIntPointer(rxsop),
			ValidAutoChannels: validAutoChannels,
		}
	}
	out := merakigosdk.RequestWirelessUpdateNetworkWirelessRfProfile{
		ApBandSettings:         requestWirelessUpdateNetworkWirelessRfProfileApBandSettings,
		BandSelectionType:      *bandSelectionType,
		ClientBalancingEnabled: clientBalancingEnabled,
		FiveGhzSettings:        requestWirelessUpdateNetworkWirelessRfProfileFiveGhzSettings,
		FlexRadios:             requestWirelessUpdateNetworkWirelessRfProfileFlexRadios,
		MinBitrateType:         *minBitrateType,
		Name:                   *name,
		PerSSIDSettings:        requestWirelessUpdateNetworkWirelessRfProfilePerSSIDSettings,
		SixGhzSettings:         requestWirelessUpdateNetworkWirelessRfProfileSixGhzSettings,
		Transmission:           requestWirelessUpdateNetworkWirelessRfProfileTransmission,
		TwoFourGhzSettings:     requestWirelessUpdateNetworkWirelessRfProfileTwoFourGhzSettings,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseWirelessGetNetworkWirelessRfProfileItemToBodyRs(state NetworksWirelessRfProfilesRs, response *merakigosdk.ResponseWirelessGetNetworkWirelessRfProfile, is_read bool) NetworksWirelessRfProfilesRs {
	itemState := NetworksWirelessRfProfilesRs{
		ApBandSettings: func() *ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsRs {
			if response.ApBandSettings != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsRs{
					BandOperationMode: types.StringValue(response.ApBandSettings.BandOperationMode),
					BandSteeringEnabled: func() types.Bool {
						if response.ApBandSettings.BandSteeringEnabled != nil {
							return types.BoolValue(*response.ApBandSettings.BandSteeringEnabled)
						}
						return types.Bool{}
					}(),
					Bands: func() *ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsBandsRs {
						if response.ApBandSettings.Bands != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsBandsRs{
								Enabled: StringSliceToSet(response.ApBandSettings.Bands.Enabled),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsBandsRs{}
					}(),
				}
			}
			return &ResponseWirelessGetNetworkWirelessRfProfileApBandSettingsRs{}
		}(),
		BandSelectionType: types.StringValue(response.BandSelectionType),
		ClientBalancingEnabled: func() types.Bool {
			if response.ClientBalancingEnabled != nil {
				return types.BoolValue(*response.ClientBalancingEnabled)
			}
			return types.Bool{}
		}(),
		FiveGhzSettings: func() *ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettingsRs {
			if response.FiveGhzSettings != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettingsRs{
					ChannelWidth: types.StringValue(response.FiveGhzSettings.ChannelWidth),
					MaxPower: func() types.Int64 {
						if response.FiveGhzSettings.MaxPower != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.MaxPower))
						}
						return types.Int64{}
					}(),
					MinBitrate: func() types.Int64 {
						if response.FiveGhzSettings.MinBitrate != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.MinBitrate))
						}
						return types.Int64{}
					}(),
					MinPower: func() types.Int64 {
						if response.FiveGhzSettings.MinPower != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.MinPower))
						}
						return types.Int64{}
					}(),
					Rxsop: func() types.Int64 {
						if response.FiveGhzSettings.Rxsop != nil {
							return types.Int64Value(int64(*response.FiveGhzSettings.Rxsop))
						}
						return types.Int64{}
					}(),
					ValidAutoChannels: StringSliceToSetInt(response.FiveGhzSettings.ValidAutoChannels),
				}
			}
			return &ResponseWirelessGetNetworkWirelessRfProfileFiveGhzSettingsRs{}
		}(),
		ID:             types.StringValue(response.ID),
		MinBitrateType: types.StringValue(response.MinBitrateType),
		Name:           types.StringValue(response.Name),
		NetworkID:      types.StringValue(response.NetworkID),
		PerSSIDSettings: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettingsRs {
			if response.PerSSIDSettings != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettingsRs{
					Status0: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0Rs {
						if response.PerSSIDSettings.Status0 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status0.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status0.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status0.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0BandsRs {
									if response.PerSSIDSettings.Status0.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status0.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status0.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status0.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status0.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings0Rs{}
					}(),
					Status1: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1Rs {
						if response.PerSSIDSettings.Status1 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status1.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status1.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status1.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1BandsRs {
									if response.PerSSIDSettings.Status1.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status1.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status1.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status1.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status1.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings1Rs{}
					}(),
					Status10: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10Rs {
						if response.PerSSIDSettings.Status10 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status10.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status10.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status10.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10BandsRs {
									if response.PerSSIDSettings.Status10.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status10.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status10.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status10.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status10.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings10Rs{}
					}(),
					Status11: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11Rs {
						if response.PerSSIDSettings.Status11 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status11.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status11.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status11.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11BandsRs {
									if response.PerSSIDSettings.Status11.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status11.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status11.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status11.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status11.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings11Rs{}
					}(),
					Status12: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12Rs {
						if response.PerSSIDSettings.Status12 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status12.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status12.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status12.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12BandsRs {
									if response.PerSSIDSettings.Status12.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status12.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status12.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status12.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status12.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings12Rs{}
					}(),
					Status13: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13Rs {
						if response.PerSSIDSettings.Status13 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status13.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status13.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status13.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13BandsRs {
									if response.PerSSIDSettings.Status13.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status13.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status13.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status13.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status13.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings13Rs{}
					}(),
					Status14: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14Rs {
						if response.PerSSIDSettings.Status14 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status14.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status14.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status14.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14BandsRs {
									if response.PerSSIDSettings.Status14.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status14.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status14.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status14.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status14.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings14Rs{}
					}(),
					Status2: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2Rs {
						if response.PerSSIDSettings.Status2 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status2.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status2.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status2.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2BandsRs {
									if response.PerSSIDSettings.Status2.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status2.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status2.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status2.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status2.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings2Rs{}
					}(),
					Status3: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3Rs {
						if response.PerSSIDSettings.Status3 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status3.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status3.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status3.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3BandsRs {
									if response.PerSSIDSettings.Status3.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status3.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status3.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status3.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status3.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings3Rs{}
					}(),
					Status4: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4Rs {
						if response.PerSSIDSettings.Status4 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status4.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status4.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status4.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4BandsRs {
									if response.PerSSIDSettings.Status4.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status4.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status4.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status4.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status4.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings4Rs{}
					}(),
					Status5: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5Rs {
						if response.PerSSIDSettings.Status5 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status5.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status5.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status5.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5BandsRs {
									if response.PerSSIDSettings.Status5.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status5.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status5.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status5.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status5.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings5Rs{}
					}(),
					Status6: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6Rs {
						if response.PerSSIDSettings.Status6 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status6.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status6.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status6.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6BandsRs {
									if response.PerSSIDSettings.Status6.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status6.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status6.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status6.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status6.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings6Rs{}
					}(),
					Status7: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7Rs {
						if response.PerSSIDSettings.Status7 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status7.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status7.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status7.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7BandsRs {
									if response.PerSSIDSettings.Status7.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status7.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status7.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status7.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status7.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings7Rs{}
					}(),
					Status8: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8Rs {
						if response.PerSSIDSettings.Status8 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status8.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status8.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status8.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8BandsRs {
									if response.PerSSIDSettings.Status8.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status8.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status8.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status8.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status8.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings8Rs{}
					}(),
					Status9: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9Rs {
						if response.PerSSIDSettings.Status9 != nil {
							return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9Rs{
								BandOperationMode: types.StringValue(response.PerSSIDSettings.Status9.BandOperationMode),
								BandSteeringEnabled: func() types.Bool {
									if response.PerSSIDSettings.Status9.BandSteeringEnabled != nil {
										return types.BoolValue(*response.PerSSIDSettings.Status9.BandSteeringEnabled)
									}
									return types.Bool{}
								}(),
								Bands: func() *ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9BandsRs {
									if response.PerSSIDSettings.Status9.Bands != nil {
										return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9BandsRs{
											Enabled: StringSliceToSet(response.PerSSIDSettings.Status9.Bands.Enabled),
										}
									}
									return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9BandsRs{}
								}(),
								MinBitrate: func() types.Int64 {
									if response.PerSSIDSettings.Status9.MinBitrate != nil {
										return types.Int64Value(int64(*response.PerSSIDSettings.Status9.MinBitrate))
									}
									return types.Int64{}
								}(),
								Name: types.StringValue(response.PerSSIDSettings.Status9.Name),
							}
						}
						return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettings9Rs{}
					}(),
				}
			}
			return &ResponseWirelessGetNetworkWirelessRfProfilePerSsidSettingsRs{}
		}(),
		SixGhzSettings: func() *ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettingsRs {
			if response.SixGhzSettings != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettingsRs{
					ChannelWidth: types.StringValue(response.SixGhzSettings.ChannelWidth),
					MaxPower: func() types.Int64 {
						if response.SixGhzSettings.MaxPower != nil {
							return types.Int64Value(int64(*response.SixGhzSettings.MaxPower))
						}
						return types.Int64{}
					}(),
					MinBitrate: func() types.Int64 {
						if response.SixGhzSettings.MinBitrate != nil {
							return types.Int64Value(int64(*response.SixGhzSettings.MinBitrate))
						}
						return types.Int64{}
					}(),
					MinPower: func() types.Int64 {
						if response.SixGhzSettings.MinPower != nil {
							return types.Int64Value(int64(*response.SixGhzSettings.MinPower))
						}
						return types.Int64{}
					}(),
					Rxsop: func() types.Int64 {
						if response.SixGhzSettings.Rxsop != nil {
							return types.Int64Value(int64(*response.SixGhzSettings.Rxsop))
						}
						return types.Int64{}
					}(),
					ValidAutoChannels: StringSliceToSetInt(response.SixGhzSettings.ValidAutoChannels),
				}
			}
			return &ResponseWirelessGetNetworkWirelessRfProfileSixGhzSettingsRs{}
		}(),
		Transmission: func() *ResponseWirelessGetNetworkWirelessRfProfileTransmissionRs {
			if response.Transmission != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfileTransmissionRs{
					Enabled: func() types.Bool {
						if response.Transmission.Enabled != nil {
							return types.BoolValue(*response.Transmission.Enabled)
						}
						return types.Bool{}
					}(),
				}
			}
			return &ResponseWirelessGetNetworkWirelessRfProfileTransmissionRs{}
		}(),
		TwoFourGhzSettings: func() *ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettingsRs {
			if response.TwoFourGhzSettings != nil {
				return &ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettingsRs{
					AxEnabled: func() types.Bool {
						if response.TwoFourGhzSettings.AxEnabled != nil {
							return types.BoolValue(*response.TwoFourGhzSettings.AxEnabled)
						}
						return types.Bool{}
					}(),
					MaxPower: func() types.Int64 {
						if response.TwoFourGhzSettings.MaxPower != nil {
							return types.Int64Value(int64(*response.TwoFourGhzSettings.MaxPower))
						}
						return types.Int64{}
					}(),
					MinBitrate: func() types.Float64 {
						if response.TwoFourGhzSettings.MinBitrate != nil {
							return types.Float64Value(float64(*response.TwoFourGhzSettings.MinBitrate))
						}
						return types.Float64{}
					}(),
					MinPower: func() types.Int64 {
						if response.TwoFourGhzSettings.MinPower != nil {
							return types.Int64Value(int64(*response.TwoFourGhzSettings.MinPower))
						}
						return types.Int64{}
					}(),
					Rxsop: func() types.Int64 {
						if response.TwoFourGhzSettings.Rxsop != nil {
							return types.Int64Value(int64(*response.TwoFourGhzSettings.Rxsop))
						}
						return types.Int64{}
					}(),
					ValidAutoChannels: StringSliceToSetInt(response.SixGhzSettings.ValidAutoChannels),
				}
			}
			return &ResponseWirelessGetNetworkWirelessRfProfileTwoFourGhzSettingsRs{}
		}(),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(NetworksWirelessRfProfilesRs)
	}
	return mergeInterfaces(state, itemState, true).(NetworksWirelessRfProfilesRs)
}
