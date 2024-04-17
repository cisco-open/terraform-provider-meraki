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
	_ resource.Resource              = &NetworksApplianceTrafficShapingCustomPerformanceClassesResource{}
	_ resource.ResourceWithConfigure = &NetworksApplianceTrafficShapingCustomPerformanceClassesResource{}
)

func NewNetworksApplianceTrafficShapingCustomPerformanceClassesResource() resource.Resource {
	return &NetworksApplianceTrafficShapingCustomPerformanceClassesResource{}
}

type NetworksApplianceTrafficShapingCustomPerformanceClassesResource struct {
	client *merakigosdk.Client
}

func (r *NetworksApplianceTrafficShapingCustomPerformanceClassesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *NetworksApplianceTrafficShapingCustomPerformanceClassesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_networks_appliance_traffic_shaping_custom_performance_classes"
}

// resourceAction
func (r *NetworksApplianceTrafficShapingCustomPerformanceClassesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{

			"network_id": schema.StringAttribute{
				MarkdownDescription: `networkId path parameter. Network ID`,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"parameters": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"max_jitter": schema.Int64Attribute{
						MarkdownDescription: `Maximum jitter in milliseconds`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"max_latency": schema.Int64Attribute{
						MarkdownDescription: `Maximum latency in milliseconds`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"max_loss_percentage": schema.Int64Attribute{
						MarkdownDescription: `Maximum percentage of packet loss`,
						Optional:            true,
						Computed:            true,
						PlanModifiers: []planmodifier.Int64{
							int64planmodifier.RequiresReplace(),
						},
					},
					"name": schema.StringAttribute{
						MarkdownDescription: `Name of the custom performance class`,
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
func (r *NetworksApplianceTrafficShapingCustomPerformanceClassesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data NetworksApplianceTrafficShapingCustomPerformanceClasses

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
	dataRequest := data.toSdkApiRequestCreate(ctx)
	restyResp1, err := r.client.Appliance.CreateNetworkApplianceTrafficShapingCustomPerformanceClass(vvNetworkID, dataRequest)

	if err != nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing CreateNetworkApplianceTrafficShapingCustomPerformanceClass",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing CreateNetworkApplianceTrafficShapingCustomPerformanceClass",
			err.Error(),
		)
		return
	}
	//Item

	// data = ResponseApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass(data, response)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *NetworksApplianceTrafficShapingCustomPerformanceClassesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// resp.Diagnostics.AddWarning("Error deleting Resource", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksApplianceTrafficShapingCustomPerformanceClassesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// resp.Diagnostics.AddWarning("Error Update Resource", "This resource has no update method in the meraki lab, the resource was deleted only in terraform.")
}

func (r *NetworksApplianceTrafficShapingCustomPerformanceClassesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type NetworksApplianceTrafficShapingCustomPerformanceClasses struct {
	NetworkID  types.String                                                                  `tfsdk:"network_id"`
	Parameters *RequestApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClassRs `tfsdk:"parameters"`
}

type RequestApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClassRs struct {
	MaxJitter         types.Int64  `tfsdk:"max_jitter"`
	MaxLatency        types.Int64  `tfsdk:"max_latency"`
	MaxLossPercentage types.Int64  `tfsdk:"max_loss_percentage"`
	Name              types.String `tfsdk:"name"`
}

// FromBody
func (r *NetworksApplianceTrafficShapingCustomPerformanceClasses) toSdkApiRequestCreate(ctx context.Context) *merakigosdk.RequestApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass {
	emptyString := ""
	re := *r.Parameters
	maxJitter := new(int64)
	if !re.MaxJitter.IsUnknown() && !re.MaxJitter.IsNull() {
		*maxJitter = re.MaxJitter.ValueInt64()
	} else {
		maxJitter = nil
	}
	maxLatency := new(int64)
	if !re.MaxLatency.IsUnknown() && !re.MaxLatency.IsNull() {
		*maxLatency = re.MaxLatency.ValueInt64()
	} else {
		maxLatency = nil
	}
	maxLossPercentage := new(int64)
	if !re.MaxLossPercentage.IsUnknown() && !re.MaxLossPercentage.IsNull() {
		*maxLossPercentage = re.MaxLossPercentage.ValueInt64()
	} else {
		maxLossPercentage = nil
	}
	name := new(string)
	if !re.Name.IsUnknown() && !re.Name.IsNull() {
		*name = re.Name.ValueString()
	} else {
		name = &emptyString
	}
	out := merakigosdk.RequestApplianceCreateNetworkApplianceTrafficShapingCustomPerformanceClass{
		MaxJitter:         int64ToIntPointer(maxJitter),
		MaxLatency:        int64ToIntPointer(maxLatency),
		MaxLossPercentage: int64ToIntPointer(maxLossPercentage),
		Name:              *name,
	}
	return &out
}

//ToBody
