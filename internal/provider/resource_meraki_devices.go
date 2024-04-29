package provider

// RESOURCE NORMAL
import (
	"context"
	"fmt"
	"net/url"
	"strings"

	merakigosdk "github.com/meraki/dashboard-api-go/v3/sdk"

	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var (
	_ resource.Resource              = &DevicesResource{}
	_ resource.ResourceWithConfigure = &DevicesResource{}
)

func NewDevicesResource() resource.Resource {
	return &DevicesResource{}
}

type DevicesResource struct {
	client *merakigosdk.Client
}

func (r *DevicesResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client := req.ProviderData.(MerakiProviderData).Client
	r.client = client
}

// Metadata returns the data source type name.
func (r *DevicesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices"
}

func (r *DevicesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"address": schema.StringAttribute{
				MarkdownDescription: `Physical address of the device`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"details": schema.SetNestedAttribute{
				MarkdownDescription: `Additional device information`,
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{

						"name": schema.StringAttribute{
							MarkdownDescription: `Additional property name`,
							Computed:            true,
						},
						"value": schema.StringAttribute{
							MarkdownDescription: `Additional property value`,
							Computed:            true,
						},
					},
				},
			},
			"firmware": schema.StringAttribute{
				MarkdownDescription: `Firmware version of the device`,
				Computed:            true,
			},
			"floor_plan_id": schema.StringAttribute{
				MarkdownDescription: `The floor plan to associate to this device. null disassociates the device from the floorplan.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"imei": schema.StringAttribute{
				MarkdownDescription: `IMEI of the device, if applicable`,
				Computed:            true,
			},
			"lan_ip": schema.StringAttribute{
				MarkdownDescription: `LAN IP address of the device`,
				Computed:            true,
			},
			"lat": schema.Float64Attribute{
				MarkdownDescription: `Latitude of the device`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
			},
			"lng": schema.Float64Attribute{
				MarkdownDescription: `Longitude of the device`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
			},
			"mac": schema.StringAttribute{
				MarkdownDescription: `MAC address of the device`,
				Computed:            true,
			},
			"model": schema.StringAttribute{
				MarkdownDescription: `Model of the device`,
				Computed:            true,
			},
			"move_map_marker": schema.BoolAttribute{
				MarkdownDescription: `Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: `Name of the device`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"network_id": schema.StringAttribute{
				MarkdownDescription: `ID of the network the device belongs to`,
				Computed:            true,
			},
			"notes": schema.StringAttribute{
				MarkdownDescription: `Notes for the device, limited to 255 characters`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			// "organization_id": schema.StringAttribute{
			// 	MarkdownDescription: `organizationId path parameter. Organization ID`,
			// 	Required:            true,
			// },
			"product_type": schema.StringAttribute{
				MarkdownDescription: `Product type of the device`,
				Computed:            true,
			},
			"serial": schema.StringAttribute{
				MarkdownDescription: `Serial number of the device`,
				Required:            true,
			},
			"switch_profile_id": schema.StringAttribute{
				MarkdownDescription: `The ID of a switch template to bind to the device (for available switch templates, see the 'Switch Templates' endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch template, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"tags": schema.SetAttribute{
				MarkdownDescription: `List of tags assigned to the device`,
				Computed:            true,
				Optional:            true,
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},

				ElementType: types.StringType,
			},
		},
	}
}

//path params to set ['serial']

func (r *DevicesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var data DevicesRs

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
	// vvOrganizationID := data.OrganizationID.ValueString()
	// organization_id
	vvSerial := data.Serial.ValueString()
	//Item
	responseVerifyItem, restyResp1, err := r.client.Devices.GetDevice(vvSerial)
	if err != nil || restyResp1 == nil || responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource Devices only have update context, not create.",
			err.Error(),
		)
		return
	}
	//Only Item
	if responseVerifyItem == nil {
		resp.Diagnostics.AddError(
			"Resource Devices only have update context, not create.",
			err.Error(),
		)
		return
	}
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Devices.UpdateDevice(vvSerial, dataRequest)

	if err != nil || restyResp2 == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDevice",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDevice",
			err.Error(),
		)
		return
	}
	//Item
	responseGet, restyResp1, err := r.client.Devices.GetDevice(vvSerial)
	// Has only items

	if err != nil || responseGet == nil {
		if restyResp1 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing GetOrganizationDevices",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetOrganizationDevices",
			err.Error(),
		)
		return
	}
	data = ResponseDevicesGetDeviceItemToBodyRs(data, responseGet, false)

	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data DevicesRs

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
	responseGet, restyRespGet, err := r.client.Devices.GetDevice(vvSerial)
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
				"Failure when executing GetDevice",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing GetDevice",
			err.Error(),
		)
		return
	}
	//entro aqui 2
	data = ResponseDevicesGetDeviceItemToBodyRs(data, responseGet, true)
	diags := resp.State.Set(ctx, &data)
	//update path params assigned
	resp.Diagnostics.Append(diags...)
}
func (r *DevicesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("serial"), req.ID)...)
}

func (r *DevicesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data DevicesRs
	merge(ctx, req, resp, &data)

	if resp.Diagnostics.HasError() {
		return
	}
	//Has Paths
	//Update

	//Path Params
	vvSerial := data.Serial.ValueString()
	dataRequest := data.toSdkApiRequestUpdate(ctx)
	restyResp2, err := r.client.Devices.UpdateDevice(vvSerial, dataRequest)
	if err != nil || restyResp2 == nil {
		if restyResp2 != nil {
			resp.Diagnostics.AddError(
				"Failure when executing UpdateDevice",
				err.Error(),
			)
			return
		}
		resp.Diagnostics.AddError(
			"Failure when executing UpdateDevice",
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(req.Plan.Set(ctx, &data)...)
	diags := resp.State.Set(ctx, &data)
	resp.Diagnostics.Append(diags...)
}

func (r *DevicesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	//missing delete
	resp.Diagnostics.AddWarning("Error deleting Devices", "This resource has no delete method in the meraki lab, the resource was deleted only in terraform.")
	resp.State.RemoveResource(ctx)
}

// TF Structs Schema
type DevicesRs struct {
	Serial          types.String                         `tfsdk:"serial"`
	Address         types.String                         `tfsdk:"address"`
	Details         *[]ResponseDevicesGetDeviceDetailsRs `tfsdk:"details"`
	Firmware        types.String                         `tfsdk:"firmware"`
	Imei            types.String                         `tfsdk:"imei"`
	LanIP           types.String                         `tfsdk:"lan_ip"`
	Lat             types.Float64                        `tfsdk:"lat"`
	Lng             types.Float64                        `tfsdk:"lng"`
	Mac             types.String                         `tfsdk:"mac"`
	Model           types.String                         `tfsdk:"model"`
	Name            types.String                         `tfsdk:"name"`
	NetworkID       types.String                         `tfsdk:"network_id"`
	Notes           types.String                         `tfsdk:"notes"`
	ProductType     types.String                         `tfsdk:"product_type"`
	Tags            types.Set                            `tfsdk:"tags"`
	FloorPlanID     types.String                         `tfsdk:"floor_plan_id"`
	MoveMapMarker   types.Bool                           `tfsdk:"move_map_marker"`
	SwitchProfileID types.String                         `tfsdk:"switch_profile_id"`
}

type ResponseDevicesGetDeviceDetailsRs struct {
	Name  types.String `tfsdk:"name"`
	Value types.String `tfsdk:"value"`
}

// FromBody
func (r *DevicesRs) toSdkApiRequestUpdate(ctx context.Context) *merakigosdk.RequestDevicesUpdateDevice {
	emptyString := ""
	address := new(string)
	if !r.Address.IsUnknown() && !r.Address.IsNull() {
		*address = r.Address.ValueString()
	} else {
		address = &emptyString
	}
	floorPlanID := new(string)
	if !r.FloorPlanID.IsUnknown() && !r.FloorPlanID.IsNull() {
		*floorPlanID = r.FloorPlanID.ValueString()
	} else {
		floorPlanID = &emptyString
	}
	lat := new(float64)
	if !r.Lat.IsUnknown() && !r.Lat.IsNull() {
		*lat = r.Lat.ValueFloat64()
	} else {
		lat = nil
	}
	lng := new(float64)
	if !r.Lng.IsUnknown() && !r.Lng.IsNull() {
		*lng = r.Lng.ValueFloat64()
	} else {
		lng = nil
	}
	moveMapMarker := new(bool)
	if !r.MoveMapMarker.IsUnknown() && !r.MoveMapMarker.IsNull() {
		*moveMapMarker = r.MoveMapMarker.ValueBool()
	} else {
		moveMapMarker = nil
	}
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = &emptyString
	}
	notes := new(string)
	if !r.Notes.IsUnknown() && !r.Notes.IsNull() {
		*notes = r.Notes.ValueString()
	} else {
		notes = &emptyString
	}
	switchProfileID := new(string)
	if !r.SwitchProfileID.IsUnknown() && !r.SwitchProfileID.IsNull() {
		*switchProfileID = r.SwitchProfileID.ValueString()
	} else {
		switchProfileID = &emptyString
	}
	var tags []string = nil
	r.Tags.ElementsAs(ctx, &tags, false)
	out := merakigosdk.RequestDevicesUpdateDevice{
		Address:         *address,
		FloorPlanID:     *floorPlanID,
		Lat:             lat,
		Lng:             lng,
		MoveMapMarker:   moveMapMarker,
		Name:            *name,
		Notes:           *notes,
		SwitchProfileID: *switchProfileID,
		Tags:            tags,
	}
	return &out
}

// From gosdk to TF Structs Schema
func ResponseDevicesGetDeviceItemToBodyRs(state DevicesRs, response *merakigosdk.ResponseDevicesGetDevice, is_read bool) DevicesRs {
	itemState := DevicesRs{
		Address: types.StringValue(response.Address),
		Details: func() *[]ResponseDevicesGetDeviceDetailsRs {
			if response.Details != nil {
				result := make([]ResponseDevicesGetDeviceDetailsRs, len(*response.Details))
				for i, details := range *response.Details {
					result[i] = ResponseDevicesGetDeviceDetailsRs{
						Name:  types.StringValue(details.Name),
						Value: types.StringValue(details.Value),
					}
				}
				return &result
			}
			return &[]ResponseDevicesGetDeviceDetailsRs{}
		}(),
		Firmware: types.StringValue(response.Firmware),
		Imei:     types.StringValue(response.Imei),
		LanIP:    types.StringValue(response.LanIP),
		Lat: func() types.Float64 {
			if response.Lat != nil {
				return types.Float64Value(float64(*response.Lat))
			}
			return types.Float64{}
		}(),
		Lng: func() types.Float64 {
			if response.Lng != nil {
				return types.Float64Value(float64(*response.Lng))
			}
			return types.Float64{}
		}(),
		Mac:         types.StringValue(response.Mac),
		Model:       types.StringValue(response.Model),
		Name:        types.StringValue(response.Name),
		NetworkID:   types.StringValue(response.NetworkID),
		Notes:       types.StringValue(response.Notes),
		ProductType: types.StringValue(response.ProductType),
		Serial:      types.StringValue(response.Serial),
		Tags:        StringSliceToSet(response.Tags),
	}
	if is_read {
		return mergeInterfacesOnlyPath(state, itemState).(DevicesRs)
	}
	return mergeInterfaces(state, itemState, true).(DevicesRs)
}

func getAllItemsDevices(client merakigosdk.Client, organizationId string) (merakigosdk.ResponseOrganizationsGetOrganizationDevices, *resty.Response, error) {
	var all_response merakigosdk.ResponseOrganizationsGetOrganizationDevices
	response, r2, er := client.Organizations.GetOrganizationDevices(organizationId, &merakigosdk.GetOrganizationDevicesQueryParams{
		PerPage: 1000,
	})
	count := 0
	all_response = append(all_response, *response...)
	for len(*response) >= 1000 {
		count += 1
		fmt.Println(count)
		links := strings.Split(r2.Header().Get("Link"), ",")
		var link string
		if count > 1 {
			link = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Split(links[2], ";")[0], ">", ""), "<", ""), client.RestyClient().BaseURL, "")
		} else {
			link = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Split(links[1], ";")[0], ">", ""), "<", ""), client.RestyClient().BaseURL, "")
		}
		myUrl, _ := url.Parse(link)
		params, _ := url.ParseQuery(myUrl.RawQuery)
		if params["endingBefore"] != nil {
			response, r2, er = client.Organizations.GetOrganizationDevices(organizationId, &merakigosdk.GetOrganizationDevicesQueryParams{
				PerPage:      1000,
				EndingBefore: params["endingBefore"][0],
			})
			all_response = append(all_response, *response...)
		} else {
			break
		}
	}

	return all_response, r2, er
}
