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
	_ resource.Resource              = &OrganizationsAlertsProfilesResource{}
	_ resource.ResourceWithConfigure = &OrganizationsAlertsProfilesResource{}
)

func NewOrganizationsAlertsProfilesResource() resource.Resource {
	return &OrganizationsAlertsProfilesResource{}
}

type OrganizationsAlertsProfilesResource struct {
	client *merakigosdk.Client
}

func (r *OrganizationsAlertsProfilesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *OrganizationsAlertsProfilesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organizations_alerts_profiles"
}

func (r *OrganizationsAlertsProfilesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"alert_condition": schema.SingleNestedAttribute{
				MarkdownDescription: `The conditions that determine if the alert triggers`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"bit_rate_bps": schema.Int64Attribute{
						MarkdownDescription: `The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"duration": schema.Int64Attribute{
						MarkdownDescription: `The total duration in seconds that the threshold should be crossed before alerting`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"interface": schema.StringAttribute{
						MarkdownDescription: `The uplink observed for the alert`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf(
								"cellular",
								"wan1",
								"wan2",
								"wan3",
							),
						},
					},
					"jitter_ms": schema.Int64Attribute{
						MarkdownDescription: `The threshold the metric must cross to be valid for alerting. Used only for VoIP Jitter alerts.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"latency_ms": schema.Int64Attribute{
						MarkdownDescription: `The threshold the metric must cross to be valid for alerting. Used only for WAN Latency alerts.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
					"loss_ratio": schema.Float64Attribute{
						MarkdownDescription: `The threshold the metric must cross to be valid for alerting. Used only for Packet Loss alerts.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"mos": schema.Float64Attribute{
						MarkdownDescription: `The threshold the metric must drop below to be valid for alerting. Used only for VoIP MOS alerts.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Float64{
							float64planmodifier.UseStateForUnknown(),
						},
					},
					"window": schema.Int64Attribute{
						MarkdownDescription: `The look back period in seconds for sensing the alert`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.UseStateForUnknown(),
						},
					},
				},
			},
			"alert_config_id": schema.StringAttribute{
				MarkdownDescription: `alertConfigId path parameter. Alert config ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: `User supplied description of the alert`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"enabled": schema.BoolAttribute{
				MarkdownDescription: `Is the alert config enabled`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"network_tags": schema.SetAttribute{
				MarkdownDescription: `Networks with these tags will be monitored for the alert`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
			"organization_id": schema.StringAttribute{
				MarkdownDescription: `organizationId path parameter. Organization ID`,
				Required:            true,
			},
			"recipients": schema.SingleNestedAttribute{
				MarkdownDescription: `List of recipients that will recieve the alert.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"emails": schema.SetAttribute{
						MarkdownDescription: `A list of emails that will receive information about the alert`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
					"http_server_ids": schema.SetAttribute{
						MarkdownDescription: `A list base64 encoded urls of webhook endpoints that will receive information about the alert`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Set{
							setplanmodifier.UseStateForUnknown(),
						},

						ElementType: types.StringType,
					},
				},
			},
			"type": schema.StringAttribute{
				MarkdownDescription: `The alert type`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(
						"appOutage",
						"voipJitter",
						"voipMos",
						"voipPacketLoss",
						"wanLatency",
						"wanPacketLoss",
						"wanStatus",
						"wanUtilization",
					),
				},
			},
		},
	}
}

//path params to set ['alertConfigId']

func (r *OrganizationsAlertsProfilesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data OrganizationsAlertsProfilesRs

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
	vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	vvAlertConfigID := data.AlertConfigID.ValueString()
	//Reviw This  Has Item Not item

	responseVerifyItem, restyResp1, err := r.client.Organizations.GetOrganizationAlertsProfiles(vvOrganizationID)

	//Only Item
	var responseVerifyItem2 merakigosdk.ResponseItemOrganizationsGetOrganizationAlertsProfiles
	responseStruct2 := structToMap(responseVerifyItem)
	result2 := getDictResult(responseStruct2, "ID", vvAlertConfigID, simpleCmp)
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		//entro aqui
		data = ResponseOrganizationsGetOrganizationAlertsProfilesItemToBodyRs(data, &responseVerifyItem2, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	}
	// HAS CRREATE
	response, restyResp2, err := r.client.Organizations.CreateOrganizationAlertsProfile(vvOrganizationID, data.toSdkApiRequestCreate(ctx))

	if err != nil || restyResp2 == nil || response == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing ",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing ",
			err.Error(),
		)
		return
	}

	vvAlertConfigID = response.ID
	//Items
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationAlertsProfiles(vvOrganizationID)
	// Has only items

	responseStruct2 = structToMap(responseGet)
	result2 = getDictResult(responseStruct2, "ID", vvAlertConfigID, simpleCmp)
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		//entro aqui
		data = ResponseOrganizationsGetOrganizationAlertsProfilesItemToBodyRs(data, &responseVerifyItem2, false)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchLinkAggregations Result",
			"Not Found",
		)
		return
	}
}

func (r *OrganizationsAlertsProfilesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data OrganizationsAlertsProfilesRs

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
	// Not has Item

	vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	vvAlertConfigID := data.AlertConfigID.ValueString()
	// alert_config_id
	responseGet, restyResp1, err := r.client.Organizations.GetOrganizationAlertsProfiles(vvOrganizationID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationAlertsProfiles",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationAlertsProfiles",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet)
	result2 := getDictResult(responseStruct2, "ID", vvAlertConfigID, simpleCmp)
	var responseVerifyItem2 merakigosdk.ResponseItemOrganizationsGetOrganizationAlertsProfiles
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		//entro aqui
		data = ResponseOrganizationsGetOrganizationAlertsProfilesItemToBodyRs(data, &responseVerifyItem2, true)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchLinkAggregations Result",
			"Not Found",
		)
		return
	}
}

func (r *OrganizationsAlertsProfilesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("organization_id"), req.ID)...)
}

func (r *OrganizationsAlertsProfilesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data OrganizationsAlertsProfilesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	vvAlertConfigID := data.AlertConfigID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	_, restyResp2, err := r.client.Organizations.UpdateOrganizationAlertsProfile(vvOrganizationID, vvAlertConfigID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateOrganizationAlertsProfile",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateOrganizationAlertsProfile",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *OrganizationsAlertsProfilesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state OrganizationsAlertsProfilesRs
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

	vvOrganizationID := state.OrganizationID.ValueString()
	vvAlertConfigID := state.AlertConfigID.ValueString()
	_, err := r.client.Organizations.DeleteOrganizationAlertsProfile(vvOrganizationID, vvAlertConfigID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteOrganizationAlertsProfile", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type OrganizationsAlertsProfilesRs struct {
	OrganizationID types.String `tfsdk:"organization_id"`
	AlertConfigID  types.String `tfsdk:"alert_config_id"`
	//TIENE ITEMS
	AlertCondition *ResponseItemOrganizationsGetOrganizationAlertsProfilesAlertConditionRs `tfsdk:"alert_condition"`
	Description    types.String                                                            `tfsdk:"description"`
	Enabled        types.Bool                                                              `tfsdk:"enabled"`
	ID             types.String                                                            `tfsdk:"id"`
	NetworkTags    types.Set                                                               `tfsdk:"network_tags"`
	Recipients     *ResponseItemOrganizationsGetOrganizationAlertsProfilesRecipientsRs     `tfsdk:"recipients"`
	Type           types.String                                                            `tfsdk:"type"`
}

type ResponseItemOrganizationsGetOrganizationAlertsProfilesAlertConditionRs struct {
	BitRateBps types.Int64   `tfsdk:"bit_rate_bps"`
	Duration   types.Int64   `tfsdk:"duration"`
	Interface  types.String  `tfsdk:"interface"`
	Window     types.Int64   `tfsdk:"window"`
	JitterMs   types.Int64   `tfsdk:"jitter_ms"`
	LatencyMs  types.Int64   `tfsdk:"latency_ms"`
	LossRatio  types.Float64 `tfsdk:"loss_ratio"`
	Mos        types.Float64 `tfsdk:"mos"`
}

type ResponseItemOrganizationsGetOrganizationAlertsProfilesRecipientsRs struct {
	Emails        types.Set `tfsdk:"emails"`
	HTTPServerIDs types.Set `tfsdk:"http_server_ids"`
}

// FromBody
func (r *OrganizationsAlertsProfilesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestOrganizationsCreateOrganizationAlertsProfile {
	emptyString := ""
	var requestOrganizationsCreateOrganizationAlertsProfileAlertCondition *merakigosdk.RequestOrganizationsCreateOrganizationAlertsProfileAlertCondition
	if r.AlertCondition != nil {
		bitRateBps := func() *int64 {
			if !r.AlertCondition.BitRateBps.IsUnknown() && !r.AlertCondition.BitRateBps.IsNull() {
				return r.AlertCondition.BitRateBps.ValueInt64Pointer()
			}
			return nil
		}()
		duration := func() *int64 {
			if !r.AlertCondition.Duration.IsUnknown() && !r.AlertCondition.Duration.IsNull() {
				return r.AlertCondition.Duration.ValueInt64Pointer()
			}
			return nil
		}()
		interfaceR := r.AlertCondition.Interface.ValueString()
		jitterMs := func() *int64 {
			if !r.AlertCondition.JitterMs.IsUnknown() && !r.AlertCondition.JitterMs.IsNull() {
				return r.AlertCondition.JitterMs.ValueInt64Pointer()
			}
			return nil
		}()
		latencyMs := func() *int64 {
			if !r.AlertCondition.LatencyMs.IsUnknown() && !r.AlertCondition.LatencyMs.IsNull() {
				return r.AlertCondition.LatencyMs.ValueInt64Pointer()
			}
			return nil
		}()
		lossRatio := func() *float64 {
			if !r.AlertCondition.LossRatio.IsUnknown() && !r.AlertCondition.LossRatio.IsNull() {
				return r.AlertCondition.LossRatio.ValueFloat64Pointer()
			}
			return nil
		}()
		mos := func() *float64 {
			if !r.AlertCondition.Mos.IsUnknown() && !r.AlertCondition.Mos.IsNull() {
				return r.AlertCondition.Mos.ValueFloat64Pointer()
			}
			return nil
		}()
		window := func() *int64 {
			if !r.AlertCondition.Window.IsUnknown() && !r.AlertCondition.Window.IsNull() {
				return r.AlertCondition.Window.ValueInt64Pointer()
			}
			return nil
		}()
		requestOrganizationsCreateOrganizationAlertsProfileAlertCondition = &merakigosdk.RequestOrganizationsCreateOrganizationAlertsProfileAlertCondition{
			BitRateBps: int64ToIntPointer(bitRateBps),
			Duration:   int64ToIntPointer(duration),
			Interface:  interfaceR,
			JitterMs:   int64ToIntPointer(jitterMs),
			LatencyMs:  int64ToIntPointer(latencyMs),
			LossRatio:  lossRatio,
			Mos:        mos,
			Window:     int64ToIntPointer(window),
		}
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = &emptyString
	}
	var networkTags []string = nil
	r.NetworkTags.ElementsAs(ctx, &networkTags, false)
	var requestOrganizationsCreateOrganizationAlertsProfileRecipients *merakigosdk.RequestOrganizationsCreateOrganizationAlertsProfileRecipients
	if r.Recipients != nil {
		var emails []string = nil

		r.Recipients.Emails.ElementsAs(ctx, &emails, false)
		var httpServerIDs []string = nil

		r.Recipients.HTTPServerIDs.ElementsAs(ctx, &httpServerIDs, false)
		requestOrganizationsCreateOrganizationAlertsProfileRecipients = &merakigosdk.RequestOrganizationsCreateOrganizationAlertsProfileRecipients{
			Emails:        emails,
			HTTPServerIDs: httpServerIDs,
		}
	}
	typeR := new(string)
	if !r.Type.IsUnknown() && !r.Type.IsNull() {
		*typeR = r.Type.ValueString()
	} else {
		typeR = &emptyString
	}
	out := merakigosdk.RequestOrganizationsCreateOrganizationAlertsProfile{
		AlertCondition: requestOrganizationsCreateOrganizationAlertsProfileAlertCondition,
		Description:    *description,
		NetworkTags:    networkTags,
		Recipients:     requestOrganizationsCreateOrganizationAlertsProfileRecipients,
		Type:           *typeR,
	}
	return &out
}
func (r *OrganizationsAlertsProfilesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestOrganizationsUpdateOrganizationAlertsProfile {
	emptyString := ""
	var requestOrganizationsUpdateOrganizationAlertsProfileAlertCondition *merakigosdk.RequestOrganizationsUpdateOrganizationAlertsProfileAlertCondition
	if r.AlertCondition != nil {
		bitRateBps := func() *int64 {
			if !r.AlertCondition.BitRateBps.IsUnknown() && !r.AlertCondition.BitRateBps.IsNull() {
				return r.AlertCondition.BitRateBps.ValueInt64Pointer()
			}
			return nil
		}()
		duration := func() *int64 {
			if !r.AlertCondition.Duration.IsUnknown() && !r.AlertCondition.Duration.IsNull() {
				return r.AlertCondition.Duration.ValueInt64Pointer()
			}
			return nil
		}()
		interfaceR := r.AlertCondition.Interface.ValueString()
		jitterMs := func() *int64 {
			if !r.AlertCondition.JitterMs.IsUnknown() && !r.AlertCondition.JitterMs.IsNull() {
				return r.AlertCondition.JitterMs.ValueInt64Pointer()
			}
			return nil
		}()
		latencyMs := func() *int64 {
			if !r.AlertCondition.LatencyMs.IsUnknown() && !r.AlertCondition.LatencyMs.IsNull() {
				return r.AlertCondition.LatencyMs.ValueInt64Pointer()
			}
			return nil
		}()
		lossRatio := func() *float64 {
			if !r.AlertCondition.LossRatio.IsUnknown() && !r.AlertCondition.LossRatio.IsNull() {
				return r.AlertCondition.LossRatio.ValueFloat64Pointer()
			}
			return nil
		}()
		mos := func() *float64 {
			if !r.AlertCondition.Mos.IsUnknown() && !r.AlertCondition.Mos.IsNull() {
				return r.AlertCondition.Mos.ValueFloat64Pointer()
			}
			return nil
		}()
		window := func() *int64 {
			if !r.AlertCondition.Window.IsUnknown() && !r.AlertCondition.Window.IsNull() {
				return r.AlertCondition.Window.ValueInt64Pointer()
			}
			return nil
		}()
		requestOrganizationsUpdateOrganizationAlertsProfileAlertCondition = &merakigosdk.RequestOrganizationsUpdateOrganizationAlertsProfileAlertCondition{
			BitRateBps: int64ToIntPointer(bitRateBps),
			Duration:   int64ToIntPointer(duration),
			Interface:  interfaceR,
			JitterMs:   int64ToIntPointer(jitterMs),
			LatencyMs:  int64ToIntPointer(latencyMs),
			LossRatio:  lossRatio,
			Mos:        mos,
			Window:     int64ToIntPointer(window),
		}
	}
	description := new(string)
	if !r.Description.IsUnknown() && !r.Description.IsNull() {
		*description = r.Description.ValueString()
	} else {
		description = &emptyString
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	var networkTags []string = nil
	r.NetworkTags.ElementsAs(ctx, &networkTags, false)
	var requestOrganizationsUpdateOrganizationAlertsProfileRecipients *merakigosdk.RequestOrganizationsUpdateOrganizationAlertsProfileRecipients
	if r.Recipients != nil {
		var emails []string = nil

		r.Recipients.Emails.ElementsAs(ctx, &emails, false)
		var httpServerIDs []string = nil

		r.Recipients.HTTPServerIDs.ElementsAs(ctx, &httpServerIDs, false)
		requestOrganizationsUpdateOrganizationAlertsProfileRecipients = &merakigosdk.RequestOrganizationsUpdateOrganizationAlertsProfileRecipients{
			Emails:        emails,
			HTTPServerIDs: httpServerIDs,
		}
	}
	typeR := new(string)
	if !r.Type.IsUnknown() && !r.Type.IsNull() {
		*typeR = r.Type.ValueString()
	} else {
		typeR = &emptyString
	}
	out := merakigosdk.RequestOrganizationsUpdateOrganizationAlertsProfile{
		AlertCondition: requestOrganizationsUpdateOrganizationAlertsProfileAlertCondition,
		Description:    *description,
		Enabled:        enabled,
		NetworkTags:    networkTags,
		Recipients:     requestOrganizationsUpdateOrganizationAlertsProfileRecipients,
		Type:           *typeR,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseOrganizationsGetOrganizationAlertsProfilesItemToBodyRs(state OrganizationsAlertsProfilesRs, response *merakigosdk.ResponseItemOrganizationsGetOrganizationAlertsProfiles, is_read bool) OrganizationsAlertsProfilesRs {
	itemState := OrganizationsAlertsProfilesRs{
		AlertCondition: func() *ResponseItemOrganizationsGetOrganizationAlertsProfilesAlertConditionRs {
			if response.AlertCondition != nil {
				return &ResponseItemOrganizationsGetOrganizationAlertsProfilesAlertConditionRs{
					BitRateBps: func() types.Int64 {
						if response.AlertCondition.BitRateBps != nil {
							return types.Int64Value(int64(*response.AlertCondition.BitRateBps))
						}
						return types.Int64{}
					}(),
					Duration: func() types.Int64 {
						if response.AlertCondition.Duration != nil {
							return types.Int64Value(int64(*response.AlertCondition.Duration))
						}
						return types.Int64{}
					}(),
					Interface: types.StringValue(response.AlertCondition.Interface),
					Window: func() types.Int64 {
						if response.AlertCondition.Window != nil {
							return types.Int64Value(int64(*response.AlertCondition.Window))
						}
						return types.Int64{}
					}(),
				}
			}
			return &ResponseItemOrganizationsGetOrganizationAlertsProfilesAlertConditionRs{}
		}(),
		Description: types.StringValue(response.Description),
		Enabled: func() types.Bool {
			if response.Enabled != nil {
				return types.BoolValue(*response.Enabled)
			}
			return types.Bool{}
		}(),
		ID:          types.StringValue(response.ID),
		NetworkTags: StringSliceToSet(response.NetworkTags),
		Recipients: func() *ResponseItemOrganizationsGetOrganizationAlertsProfilesRecipientsRs {
			if response.Recipients != nil {
				return &ResponseItemOrganizationsGetOrganizationAlertsProfilesRecipientsRs{
					Emails:        StringSliceToSet(response.Recipients.Emails),
					HTTPServerIDs: StringSliceToSet(response.Recipients.HTTPServerIDs),
				}
			}
			return &ResponseItemOrganizationsGetOrganizationAlertsProfilesRecipientsRs{}
		}(),
		Type: types.StringValue(response.Type),
	}
	state = itemState
	return state
}
