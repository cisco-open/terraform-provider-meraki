package provider

// RESOURCE NORMAL
import (
	"context"
	"regexp"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &NetworksSwitchPortSchedulesResource{}
	_ resource.ResourceWithConfigure = &NetworksSwitchPortSchedulesResource{}
)

func NewNetworksSwitchPortSchedulesResource() resource.Resource {
	return &NetworksSwitchPortSchedulesResource{}
}

type NetworksSwitchPortSchedulesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksSwitchPortSchedulesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksSwitchPortSchedulesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_switch_port_schedules"
}

func (r *NetworksSwitchPortSchedulesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `The name for your port schedule. Required`,
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
			"port_schedule": schema.SingleNestedAttribute{
				MarkdownDescription: `    The schedule for switch port scheduling. Schedules are applied to days of the week.
    When it's empty, default schedule with all days of a week are configured.
    Any unspecified day in the schedule is added as a default schedule configuration of the day.
`,
				Computed: true,
				Optional: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.UseStateForUnknown(),
				},
				Attributes: map[string]schema.Attribute{

					"friday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Friday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
						},
					},
					"monday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Monday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
						},
					},
					"saturday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Saturday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
						},
					},
					"sunday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Sunday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
						},
					},
					"thursday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Thursday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
						},
					},
					"tuesday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Tuesday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
						},
					},
					"wednesday": schema.SingleNestedAttribute{
						MarkdownDescription: `The schedule object for Wednesday.`,
						Computed:            true,
						Optional:            true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.UseStateForUnknown(),
						},
						Attributes: map[string]schema.Attribute{

							"active": schema.BoolAttribute{
								MarkdownDescription: `Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.Bool{
									boolplanmodifier.UseStateForUnknown(),
								},
							},
							"from": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
							"to": schema.StringAttribute{
								MarkdownDescription: `The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.`,
								Computed:            true,
								Optional:            true,
								PlanModifiers: []planmodifier.String{
									stringplanmodifier.UseStateForUnknown(),
								},
								Validators: []validator.String{
									stringvalidator.RegexMatches(regexp.MustCompile(`^(24:00|([01]\d|2[0-3]):[0-5]\d)$`), "The time, from '00:00' to '24:00' with format xx:xx"),
								},
							},
						},
					},
				},
			},
			"port_schedule_id": schema.StringAttribute{
				MarkdownDescription: `portScheduleId path parameter. Port schedule ID`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

//path params to set ['portScheduleId']

func (r *NetworksSwitchPortSchedulesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksSwitchPortSchedulesRs

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
	responseVerifyItem, restyResp1, err := r.client.Switch.GetNetworkSwitchPortSchedules(vvNetworkID)
	//Have Create
	if err != nil {
		if restyResp1 == nil {
			resp.Diagnostics.AddError(
				"Failure when executing Get",
				err.Error(),
			)
			return
		}
		if restyResp1.StatusCode() != 404 {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchPortSchedules",
				err.Error(),
			)
			return
		}
	}

	// Create

	responseStruct := structToMap(responseVerifyItem)
	result := getDictResult(responseStruct, "Name", vvName, simpleCmp)
	var responseVerifyItem2 merakigosdk.ResponseItemSwitchGetNetworkSwitchPortSchedules
	if result != nil {

		err := mapToStruct(result.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		r.client.Switch.UpdateNetworkSwitchPortSchedule(vvNetworkID, responseVerifyItem2.ID, data.toSdkApiRequestUpdate(ctx))
		responseVerifyItem3, _, err := r.client.Switch.GetNetworkSwitchPortSchedules(vvNetworkID)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchPortSchedules in resource",
				err.Error(),
			)
			return
		}
		data.NetworkID = types.StringValue(responseVerifyItem2.NetworkID)
		responseStruct2 := structToMap(responseVerifyItem3)
		result2 := getDictResult(responseStruct2, "Name", vvName, simpleCmp)
		var responseVerifyItem4 merakigosdk.ResponseItemSwitchGetNetworkSwitchPortSchedules
		if result2 != nil {
			err := mapToStruct(result.(map[string]interface{}), &responseVerifyItem4)
			if err != nil {
				resp.Diagnostics.AddError(
					"Failure when executing mapToStruct in resource",
					err.Error(),
				)
				return
			}
			data = ResponseSwitchGetNetworkSwitchPortSchedulesItemToBodyRs(data, &responseVerifyItem4, false)
			diags := resp.State.Set(ctx, &data)
			resp.Diagnostics.Append(diags...)
			return
		}
		return
	}
	dataRequest := data.toSdkApiRequestCreate(ctx)
	_, restyResp2, err := r.client.Switch.CreateNetworkSwitchPortSchedule(vvNetworkID, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkSwitchPortSchedule",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkSwitchPortSchedule",
			err.Error(),
		)
		return
	}
	//Items
	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchPortSchedules(vvNetworkID)
	// Has only items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchPortSchedules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchPortSchedules",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet)
	result2 := getDictResult(responseStruct2, "Name", vvName, simpleCmp)
	if result2 != nil {
		err := mapToStruct(result2.(map[string]interface{}), &responseVerifyItem2)
		if err != nil {
			resp.Diagnostics.AddError(
				"Failure when executing mapToStruct in resource",
				err.Error(),
			)
			return
		}
		data = ResponseSwitchGetNetworkSwitchPortSchedulesItemToBodyRs(data, &responseVerifyItem2, false)
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

func (r *NetworksSwitchPortSchedulesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data NetworksSwitchPortSchedulesRs

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

	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvName := data.Name.ValueString()
	// name

	responseGet, restyResp1, err := r.client.Switch.GetNetworkSwitchPortSchedules(vvNetworkID)

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetNetworkSwitchPortSchedules",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetNetworkSwitchPortSchedules",
			err.Error(),
		)
		return
	}
	responseStruct2 := structToMap(responseGet)
	result2 := getDictResult(responseStruct2, "Name", vvName, simpleCmp)
	var responseVerifyItem2 merakigosdk.ResponseItemSwitchGetNetworkSwitchPortSchedules
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
		data = ResponseSwitchGetNetworkSwitchPortSchedulesItemToBodyRs(data, &responseVerifyItem2, true)
		diags := resp.State.Set(ctx, &data)
		resp.Diagnostics.Append(diags...)
		return
	} else {
		resp.Diagnostics.AddWarning(
			"Resource not found",
			"Deleting resource",
		)
		resp.State.RemoveResource(ctx)
		return
	}
}

func (r *NetworksSwitchPortSchedulesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("network_id"), req.ID)...)
}

func (r *NetworksSwitchPortSchedulesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data NetworksSwitchPortSchedulesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvNetworkID := data.NetworkID.ValueString()
	// network_id
	vvPortScheduleID := data.PortScheduleID.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	_, restyResp2, err := r.client.Switch.UpdateNetworkSwitchPortSchedule(vvNetworkID, vvPortScheduleID, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateNetworkSwitchPortSchedule",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateNetworkSwitchPortSchedule",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksSwitchPortSchedulesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state NetworksSwitchPortSchedulesRs
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
	vvPortScheduleID := state.PortScheduleID.ValueString()
	_, err := r.client.Switch.DeleteNetworkSwitchPortSchedule(vvNetworkID, vvPortScheduleID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failure when executing DeleteNetworkSwitchPortSchedule", err.Error())
		return
	}

	resp.State.RemoveResource(ctx)

}

// TF Structs Schema
type NetworksSwitchPortSchedulesRs struct {
	NetworkID      types.String `tfsdk:"network_id"`
	PortScheduleID types.String `tfsdk:"port_schedule_id"`
	//TIENE ITEMS
	ID           types.String                                                   `tfsdk:"id"`
	Name         types.String                                                   `tfsdk:"name"`
	PortSchedule *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleRs `tfsdk:"port_schedule"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleRs struct {
	Friday    *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFridayRs    `tfsdk:"friday"`
	Monday    *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMondayRs    `tfsdk:"monday"`
	Saturday  *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturdayRs  `tfsdk:"saturday"`
	Sunday    *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSundayRs    `tfsdk:"sunday"`
	Thursday  *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursdayRs  `tfsdk:"thursday"`
	Tuesday   *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesdayRs   `tfsdk:"tuesday"`
	Wednesday *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesdayRs `tfsdk:"wednesday"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFridayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMondayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturdayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSundayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursdayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesdayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

type ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesdayRs struct {
	Active types.Bool   `tfsdk:"active"`
	From   types.String `tfsdk:"from"`
	To     types.String `tfsdk:"to"`
}

// FromBody
func (r *NetworksSwitchPortSchedulesRs) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedule {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestSwitchCreateNetworkSwitchPortSchedulePortSchedule *merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortSchedule
	if r.PortSchedule != nil {
		var requestSwitchCreateNetworkSwitchPortSchedulePortScheduleFriday *merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleFriday
		if r.PortSchedule.Friday != nil {
			active := func() *bool {
				if !r.PortSchedule.Friday.Active.IsUnknown() && !r.PortSchedule.Friday.Active.IsNull() {
					return r.PortSchedule.Friday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Friday.From.ValueString()
			to := r.PortSchedule.Friday.To.ValueString()
			requestSwitchCreateNetworkSwitchPortSchedulePortScheduleFriday = &merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleFriday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestSwitchCreateNetworkSwitchPortSchedulePortScheduleMonday *merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleMonday
		if r.PortSchedule.Monday != nil {
			active := func() *bool {
				if !r.PortSchedule.Monday.Active.IsUnknown() && !r.PortSchedule.Monday.Active.IsNull() {
					return r.PortSchedule.Monday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Monday.From.ValueString()
			to := r.PortSchedule.Monday.To.ValueString()
			requestSwitchCreateNetworkSwitchPortSchedulePortScheduleMonday = &merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleMonday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestSwitchCreateNetworkSwitchPortSchedulePortScheduleSaturday *merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleSaturday
		if r.PortSchedule.Saturday != nil {
			active := func() *bool {
				if !r.PortSchedule.Saturday.Active.IsUnknown() && !r.PortSchedule.Saturday.Active.IsNull() {
					return r.PortSchedule.Saturday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Saturday.From.ValueString()
			to := r.PortSchedule.Saturday.To.ValueString()
			requestSwitchCreateNetworkSwitchPortSchedulePortScheduleSaturday = &merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleSaturday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestSwitchCreateNetworkSwitchPortSchedulePortScheduleSunday *merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleSunday
		if r.PortSchedule.Sunday != nil {
			active := func() *bool {
				if !r.PortSchedule.Sunday.Active.IsUnknown() && !r.PortSchedule.Sunday.Active.IsNull() {
					return r.PortSchedule.Sunday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Sunday.From.ValueString()
			to := r.PortSchedule.Sunday.To.ValueString()
			requestSwitchCreateNetworkSwitchPortSchedulePortScheduleSunday = &merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleSunday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestSwitchCreateNetworkSwitchPortSchedulePortScheduleThursday *merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleThursday
		if r.PortSchedule.Thursday != nil {
			active := func() *bool {
				if !r.PortSchedule.Thursday.Active.IsUnknown() && !r.PortSchedule.Thursday.Active.IsNull() {
					return r.PortSchedule.Thursday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Thursday.From.ValueString()
			to := r.PortSchedule.Thursday.To.ValueString()
			requestSwitchCreateNetworkSwitchPortSchedulePortScheduleThursday = &merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleThursday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestSwitchCreateNetworkSwitchPortSchedulePortScheduleTuesday *merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleTuesday
		if r.PortSchedule.Tuesday != nil {
			active := func() *bool {
				if !r.PortSchedule.Tuesday.Active.IsUnknown() && !r.PortSchedule.Tuesday.Active.IsNull() {
					return r.PortSchedule.Tuesday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Tuesday.From.ValueString()
			to := r.PortSchedule.Tuesday.To.ValueString()
			requestSwitchCreateNetworkSwitchPortSchedulePortScheduleTuesday = &merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleTuesday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestSwitchCreateNetworkSwitchPortSchedulePortScheduleWednesday *merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleWednesday
		if r.PortSchedule.Wednesday != nil {
			active := func() *bool {
				if !r.PortSchedule.Wednesday.Active.IsUnknown() && !r.PortSchedule.Wednesday.Active.IsNull() {
					return r.PortSchedule.Wednesday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Wednesday.From.ValueString()
			to := r.PortSchedule.Wednesday.To.ValueString()
			requestSwitchCreateNetworkSwitchPortSchedulePortScheduleWednesday = &merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortScheduleWednesday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		requestSwitchCreateNetworkSwitchPortSchedulePortSchedule = &merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedulePortSchedule{
			Friday:    requestSwitchCreateNetworkSwitchPortSchedulePortScheduleFriday,
			Monday:    requestSwitchCreateNetworkSwitchPortSchedulePortScheduleMonday,
			Saturday:  requestSwitchCreateNetworkSwitchPortSchedulePortScheduleSaturday,
			Sunday:    requestSwitchCreateNetworkSwitchPortSchedulePortScheduleSunday,
			Thursday:  requestSwitchCreateNetworkSwitchPortSchedulePortScheduleThursday,
			Tuesday:   requestSwitchCreateNetworkSwitchPortSchedulePortScheduleTuesday,
			Wednesday: requestSwitchCreateNetworkSwitchPortSchedulePortScheduleWednesday,
		}
	}
	out := merakigosdk.RequestSwitchCreateNetworkSwitchPortSchedule{
		Name:         *name,
		PortSchedule: requestSwitchCreateNetworkSwitchPortSchedulePortSchedule,
	}
	return &out
}
func (r *NetworksSwitchPortSchedulesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedule {
	emptyString := ""
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	var requestSwitchUpdateNetworkSwitchPortSchedulePortSchedule *merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortSchedule
	if r.PortSchedule != nil {
		var requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleFriday *merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleFriday
		if r.PortSchedule.Friday != nil {
			active := func() *bool {
				if !r.PortSchedule.Friday.Active.IsUnknown() && !r.PortSchedule.Friday.Active.IsNull() {
					return r.PortSchedule.Friday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Friday.From.ValueString()
			to := r.PortSchedule.Friday.To.ValueString()
			requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleFriday = &merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleFriday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleMonday *merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleMonday
		if r.PortSchedule.Monday != nil {
			active := func() *bool {
				if !r.PortSchedule.Monday.Active.IsUnknown() && !r.PortSchedule.Monday.Active.IsNull() {
					return r.PortSchedule.Monday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Monday.From.ValueString()
			to := r.PortSchedule.Monday.To.ValueString()
			requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleMonday = &merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleMonday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSaturday *merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSaturday
		if r.PortSchedule.Saturday != nil {
			active := func() *bool {
				if !r.PortSchedule.Saturday.Active.IsUnknown() && !r.PortSchedule.Saturday.Active.IsNull() {
					return r.PortSchedule.Saturday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Saturday.From.ValueString()
			to := r.PortSchedule.Saturday.To.ValueString()
			requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSaturday = &merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSaturday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSunday *merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSunday
		if r.PortSchedule.Sunday != nil {
			active := func() *bool {
				if !r.PortSchedule.Sunday.Active.IsUnknown() && !r.PortSchedule.Sunday.Active.IsNull() {
					return r.PortSchedule.Sunday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Sunday.From.ValueString()
			to := r.PortSchedule.Sunday.To.ValueString()
			requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSunday = &merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSunday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleThursday *merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleThursday
		if r.PortSchedule.Thursday != nil {
			active := func() *bool {
				if !r.PortSchedule.Thursday.Active.IsUnknown() && !r.PortSchedule.Thursday.Active.IsNull() {
					return r.PortSchedule.Thursday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Thursday.From.ValueString()
			to := r.PortSchedule.Thursday.To.ValueString()
			requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleThursday = &merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleThursday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleTuesday *merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleTuesday
		if r.PortSchedule.Tuesday != nil {
			active := func() *bool {
				if !r.PortSchedule.Tuesday.Active.IsUnknown() && !r.PortSchedule.Tuesday.Active.IsNull() {
					return r.PortSchedule.Tuesday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Tuesday.From.ValueString()
			to := r.PortSchedule.Tuesday.To.ValueString()
			requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleTuesday = &merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleTuesday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		var requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleWednesday *merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleWednesday
		if r.PortSchedule.Wednesday != nil {
			active := func() *bool {
				if !r.PortSchedule.Wednesday.Active.IsUnknown() && !r.PortSchedule.Wednesday.Active.IsNull() {
					return r.PortSchedule.Wednesday.Active.ValueBoolPointer()
				}
				return nil
			}()
			from := r.PortSchedule.Wednesday.From.ValueString()
			to := r.PortSchedule.Wednesday.To.ValueString()
			requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleWednesday = &merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortScheduleWednesday{
				Active: active,
				From:   from,
				To:     to,
			}
		}
		requestSwitchUpdateNetworkSwitchPortSchedulePortSchedule = &merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedulePortSchedule{
			Friday:    requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleFriday,
			Monday:    requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleMonday,
			Saturday:  requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSaturday,
			Sunday:    requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleSunday,
			Thursday:  requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleThursday,
			Tuesday:   requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleTuesday,
			Wednesday: requestSwitchUpdateNetworkSwitchPortSchedulePortScheduleWednesday,
		}
	}
	out := merakigosdk.RequestSwitchUpdateNetworkSwitchPortSchedule{
		Name:         *name,
		PortSchedule: requestSwitchUpdateNetworkSwitchPortSchedulePortSchedule,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseSwitchGetNetworkSwitchPortSchedulesItemToBodyRs(state NetworksSwitchPortSchedulesRs, response *merakigosdk.ResponseItemSwitchGetNetworkSwitchPortSchedules, is_read bool) NetworksSwitchPortSchedulesRs {
	itemState := NetworksSwitchPortSchedulesRs{
		ID:   types.StringValue(response.ID),
		Name: types.StringValue(response.Name),
		PortSchedule: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleRs {
			if response.PortSchedule != nil {
				return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleRs{
					Friday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFridayRs {
						if response.PortSchedule.Friday != nil {
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFridayRs{
								Active: func() types.Bool {
									if response.PortSchedule.Friday.Active != nil {
										return types.BoolValue(*response.PortSchedule.Friday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.PortSchedule.Friday.From),
								To:   types.StringValue(response.PortSchedule.Friday.To),
							}
						}
						return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleFridayRs{}
					}(),
					Monday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMondayRs {
						if response.PortSchedule.Monday != nil {
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMondayRs{
								Active: func() types.Bool {
									if response.PortSchedule.Monday.Active != nil {
										return types.BoolValue(*response.PortSchedule.Monday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.PortSchedule.Monday.From),
								To:   types.StringValue(response.PortSchedule.Monday.To),
							}
						}
						return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleMondayRs{}
					}(),
					Saturday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturdayRs {
						if response.PortSchedule.Saturday != nil {
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturdayRs{
								Active: func() types.Bool {
									if response.PortSchedule.Saturday.Active != nil {
										return types.BoolValue(*response.PortSchedule.Saturday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.PortSchedule.Saturday.From),
								To:   types.StringValue(response.PortSchedule.Saturday.To),
							}
						}
						return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSaturdayRs{}
					}(),
					Sunday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSundayRs {
						if response.PortSchedule.Sunday != nil {
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSundayRs{
								Active: func() types.Bool {
									if response.PortSchedule.Sunday.Active != nil {
										return types.BoolValue(*response.PortSchedule.Sunday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.PortSchedule.Sunday.From),
								To:   types.StringValue(response.PortSchedule.Sunday.To),
							}
						}
						return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleSundayRs{}
					}(),
					Thursday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursdayRs {
						if response.PortSchedule.Thursday != nil {
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursdayRs{
								Active: func() types.Bool {
									if response.PortSchedule.Thursday.Active != nil {
										return types.BoolValue(*response.PortSchedule.Thursday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.PortSchedule.Thursday.From),
								To:   types.StringValue(response.PortSchedule.Thursday.To),
							}
						}
						return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleThursdayRs{}
					}(),
					Tuesday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesdayRs {
						if response.PortSchedule.Tuesday != nil {
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesdayRs{
								Active: func() types.Bool {
									if response.PortSchedule.Tuesday.Active != nil {
										return types.BoolValue(*response.PortSchedule.Tuesday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.PortSchedule.Tuesday.From),
								To:   types.StringValue(response.PortSchedule.Tuesday.To),
							}
						}
						return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleTuesdayRs{}
					}(),
					Wednesday: func() *ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesdayRs {
						if response.PortSchedule.Wednesday != nil {
							return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesdayRs{
								Active: func() types.Bool {
									if response.PortSchedule.Wednesday.Active != nil {
										return types.BoolValue(*response.PortSchedule.Wednesday.Active)
									}
									return types.Bool{}
								}(),
								From: types.StringValue(response.PortSchedule.Wednesday.From),
								To:   types.StringValue(response.PortSchedule.Wednesday.To),
							}
						}
						return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleWednesdayRs{}
					}(),
				}
			}
			return &ResponseItemSwitchGetNetworkSwitchPortSchedulesPortScheduleRs{}
		}(),
		NetworkID: state.NetworkID,
	}
	state = itemState
	return state
}
