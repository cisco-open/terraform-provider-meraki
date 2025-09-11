package provider

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"sync"

	tfsdkr "github.com/cisco-open/terraform-provider-meraki/internal/provider/reflects"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

var simpleTypes = []string{
	"basetypes.StringValue",
}

var pathParams = []string{
	"PortID",
	"InterfaceID",
	"StaticRouteID",
	"Serial",
	"StaticDelegatedPrefixID",
	"Number",
	"VLANID",
	"QualityRetentionProfileID",
	"WirelessProfileID",
	"GroupID",
	"FloorPlanID",
	"GroupPolicyID",
	"MerakiAuthUserID",
	"ID",
	"MqttBrokerID",
	"TargetGroupID",
	"AccessPolicyNumber",
	"TrustedServerID",
	"LinkAggregationID",
	"PortScheduleID",
	"QosRuleID",
	"RendezvousPointID",
	"HTTPServerID",
	"PayloadTemplateID",
	"RfProfileID",
	"IdentityPskID",
	"NetworkID",
	"ActionBatchID",
	"ACLID",
	"AdminID",
	"AlertConfigID",
	"BrandingPolicyID",
	"ConfigTemplateID",
	"OptInID",
	"MonitoredMediaServerID",
	"PolicyObjectGroupID",
	"PolicyObjectID",
	"SamlRoleID",
	"OrganizationID",
	"ProfileID",
}

type cmpFunc func(interface{}, interface{}) bool

// Simple comparison function that compares if two values are equal.
func simpleCmp(a, b interface{}) bool {
	return a == b
}

// func int64ToIntPointer(i int64) *int64 {
// 	return &i
// }

func pickMethodAux(method []bool) float64 {
	lenM := len(method)
	countM := 0
	for _, em := range method {
		if em {
			countM += 1
		}
	}
	var percentM float64 = float64(countM) / float64(lenM)
	return percentM
}

func pickMethod(methods [][]bool) int {
	methodN := 0
	maxPercentM := 0.0
	for i, method := range methods {
		percentM := pickMethodAux(method)
		if maxPercentM < percentM {
			methodN = i
			maxPercentM = percentM
		}
	}
	// Add 1 to match number method and not index
	return methodN + 1
}

func StringSliceToList(items []string) types.List {
	var eles []attr.Value
	log.Printf("[DEBUG] StringSliceToList: %v", items)
	if len(items) == 0 || items == nil {
		log.Printf("[DEBUG] StringSliceToList: returning null for empty list")
		return types.ListNull(types.StringType)
	}
	for _, item := range items {
		eles = append(eles, types.StringValue(item))
	}

	return types.ListValueMust(types.StringType, eles)
}

// isSetOrListType checks if a reflect.Type is a Set or List type
func isSetOrListType(fieldType reflect.Type) bool {
	// Check for the primitive types
	if fieldType == reflect.TypeOf(types.Set{}) || fieldType == reflect.TypeOf(types.List{}) {
		return true
	}

	// Check for pointer types
	if fieldType.Kind() == reflect.Ptr {
		elemType := fieldType.Elem()
		if elemType == reflect.TypeOf(types.Set{}) || elemType == reflect.TypeOf(types.List{}) {
			return true
		}
	}

	return false
}

// createEmptySetOrList creates an empty Set or List of the appropriate type
func createEmptySetOrList(fieldType reflect.Type) interface{} {
	if fieldType == reflect.TypeOf(types.Set{}) {
		return types.SetNull(types.StringType)
	}
	if fieldType == reflect.TypeOf(types.List{}) {
		return types.ListNull(types.StringType)
	}

	// For pointer types
	if fieldType.Kind() == reflect.Ptr {
		elemType := fieldType.Elem()
		if elemType == reflect.TypeOf(types.Set{}) {
			return types.SetNull(types.StringType)
		}
		if elemType == reflect.TypeOf(types.List{}) {
			return types.ListNull(types.StringType)
		}
	}

	return reflect.Zero(fieldType).Interface()
}

func StringSliceToSet(items []string) basetypes.SetValue {
	var eles []attr.Value
	if len(items) == 0 {
		return types.SetNull(types.StringType)
	}
	for _, item := range items {
		eles = append(eles, types.StringValue(item))
	}

	return types.SetValueMust(types.StringType, eles)
}

func StringSliceToSetInt(items *[]int) basetypes.SetValue {
	var eles []attr.Value
	if items == nil {
		return types.SetNull(types.Int64Type)
	}

	if len(*items) == 0 {
		return types.SetNull(types.Int64Type)
	}
	for _, item := range *items {
		eles = append(eles, basetypes.NewInt64Value(int64(item)))
	}

	return types.SetValueMust(types.Int64Type, eles)
}

func StringSliceToListInt(items *[]int) basetypes.ListValue {
	var eles []attr.Value
	if items == nil {
		return types.ListNull(types.Int64Type)
	}

	if len(*items) == 0 {
		return types.ListNull(types.Int64Type)
	}
	for _, item := range *items {
		eles = append(eles, basetypes.NewInt64Value(int64(item)))
	}

	return types.ListValueMust(types.Int64Type, eles)
}

func elementsToStrings(ctx context.Context, elements types.List) []string {
	var strings []string
	elements.ElementsAs(ctx, &strings, false)

	return strings
}

// getDictResult looks up and returns a specific dictionary within a flexible data structure.
func getDictResult(result interface{}, key string, value interface{}, cmpFn cmpFunc) interface{} {
	// Local mutex to synchronize map access
	var mu sync.Mutex

	// Map to store the result
	resultMap := make(map[string]interface{})

	// Function to check if a value is a map[string]interface{}
	isMapStringInterface := func(v interface{}) (map[string]interface{}, bool) {
		if m, ok := v.(map[string]interface{}); ok {
			return m, true
		}
		return nil, false
	}

	// Internal function to process an element
	processItem := func(item interface{}) {
		if itemMap, ok := isMapStringInterface(item); ok {
			// Check if the key and value match those provided.
			log.Print(" Check if the key and value match those provided.")
			if val, ok := itemMap[key]; ok && cmpFn(val, value) {
				// Block map access before mapping
				log.Print(" Block map access before mapping")
				mu.Lock()
				defer mu.Unlock()
				resultMap = itemMap
			}
		}
	}

	// Check if the result is a slice.
	log.Print("Check if the result is a slice.")
	if reflect.TypeOf(result).Kind() == reflect.Slice {
		resultSlice := reflect.ValueOf(result)
		log.Print("Check if the result is a slice. 2")
		if resultSlice.Len() == 0 {
			return nil
		}
		log.Print("If the slice has only one element.")
		// If the slice has only one element.
		if resultSlice.Len() == 1 {
			if itemMap, ok := isMapStringInterface(resultSlice.Index(0).Interface()); ok {

				log.Print(" Check if the key and value match those provided.")
				// Check if the key and value match those provided.
				if val, ok := itemMap[key]; ok && !cmpFn(val, value) {
					return nil
				}
				return itemMap
			}
			return nil
		}

		// Iterate over the elements of the slice.
		log.Print("Iterate over the elements of the slice.")

		for i := 0; i < resultSlice.Len(); i++ {
			processItem(resultSlice.Index(i).Interface())
		}
		if len(reflect.ValueOf(resultMap).MapKeys()) == 0 {
			return nil
		}
		return resultMap
	}

	// If the result is not a slice.
	if itemMap, ok := isMapStringInterface(result); ok {
		// Check if the key and value match those provided.
		if val, ok := itemMap[key]; ok && !cmpFn(val, value) {
			return nil
		}
		return itemMap
	}

	// If the result is neither a slice nor a map.
	return nil
}
func structToMap(data interface{}) []map[string]interface{} {
	// Local mutex to synchronize access to the result
	var mu sync.Mutex

	// Check if data is a pointer, if so get the value
	if reflect.ValueOf(data).Kind() == reflect.Ptr {
		data = reflect.ValueOf(data).Elem().Interface()
	}

	// Check if a slice is passed
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		log.Printf("[DEBUG] the value is not a slice.")
		fmt.Println("the value is not a slice.")
		// return nil
	}

	// Get the type of the structure inside the slice
	elementType := val.Type().Elem()
	if elementType.Kind() == reflect.Ptr {
		elementType = elementType.Elem()
	}
	if elementType.Kind() != reflect.Struct {
		log.Printf("[DEBUG] the type inside the slice is not a structure.")
		fmt.Println("the type inside the slice is not a structure.")
		// return nil
	}

	// Convert each element of the slice to a map
	result := make([]map[string]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		element := val.Index(i)

		// Internal function to convert a structure to a map
		structToMapFunc := func(elem reflect.Value) map[string]interface{} {
			elemMap := make(map[string]interface{})
			for j := 0; j < elem.NumField(); j++ {
				fieldName := elementType.Field(j).Name
				fieldValue := elem.Field(j).Interface()
				elemMap[fieldName] = fieldValue
			}
			return elemMap
		}

		// Block access to the result before assigning
		mu.Lock()
		result[i] = structToMapFunc(element)
		mu.Unlock()
	}

	return result
}

// mapToStruct assigns values ​​from a map to a structure (target)
func mapToStruct(data map[string]interface{}, target interface{}) error {
	// Local mutex to synchronize access to the target structure
	var mu sync.Mutex

	// Check if target is a pointer to a structure
	targetType := reflect.TypeOf(target)
	if targetType.Kind() != reflect.Ptr || targetType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("the target is not a pointer to a structure")
	}
	targetValue := reflect.ValueOf(target).Elem()

	// Iterate over the fields of the target structure
	for i := 0; i < targetType.Elem().NumField(); i++ {
		field := targetType.Elem().Field(i)
		fieldName := field.Name

		// Get map value for field name
		value, ok := data[fieldName]
		if !ok {
			continue // If it does not exist on the map, continue with the next field
		}

		// Internal function to assign the value to the structure field
		assignValue := func(field reflect.Value, value interface{}) error {
			// Block field access before assignment
			mu.Lock()
			defer mu.Unlock()

			fieldValue := reflect.ValueOf(value)

			// Check if the value type is convertible to the field type
			if fieldValue.Type().ConvertibleTo(field.Type()) {
				field.Set(fieldValue.Convert(field.Type()))
			} else {
				return fmt.Errorf("the value type for field %s is not convertible to the field type", fieldName)
			}
			return nil
		}

		// Llamar a la función interna para asignar el valor al campo
		if err := assignValue(targetValue.FieldByName(fieldName), value); err != nil {
			return err
		}
	}

	return nil
}

func int64ToString(i *int64) string {
	// log.Printf("INT: %v", i)
	if i == nil {
		return ""
	}
	a := strconv.Itoa(int(*i))
	return a
}

func int64ToIntPointer(i *int64) *int {
	// log.Printf("INT: %v", i)
	if i == nil {
		return nil
	}
	a := int(*i)
	return &a
}
func int64ToIntPointer3(i *int64) *int {
	if i == nil {
		return nil
	}
	a := int(*i)
	return &a
}

// func int64ToIntPointer(i int64) *int {
// 	a := int(i)
// 	return &a
// }

func int64ToFloat(i *int64) *float64 {
	if i == nil {
		return nil
	}
	a := float64(*i)
	return &a
}

const (
	// ExplicitSuppress strategy suppresses "(known after changes)" messages unless we're in the initial creation
	ExplicitSuppress = iota
)

// SuppressDiffString returns a plan modifier that propagates a state value into the planned value, when it is Known, and the Plan Value is Unknown
func SuppressDiffString() planmodifier.String {
	return suppressDiffString{}
}

// suppressDiffString implements the plan modifier.
type suppressDiffString struct{}

// Description returns a human-readable description of the plan modifier.
func (m suppressDiffString) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m suppressDiffString) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyString implements the plan modification logic.
func (m suppressDiffString) PlanModifyString(ctx context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	// Do nothing if there is a known planned value.
	// a, _ := resp.PlanValue.ToStringValue(ctx)
	log.Printf("[DEBUG] resp: %v", resp)
	log.Printf("[DEBUG] state: %v", req.Path)
	if !req.StateValue.IsNull() {
		if resp.PlanValue != req.StateValue {
			resp.Diagnostics.AddWarning("Attribute not editable changed "+req.Path.String(), "An attribute that is not editable has been changed, you may lose consistency because of it.")
		}
		return
	}
	if req.PlanValue.IsUnknown() {
		return
	}
	if resp.PlanValue.IsUnknown() {
		return
	}
	// Do nothing if there is an unknown configuration value
	if req.ConfigValue.IsUnknown() {
		return
	}
	if IsAllStateUnknown(ctx, req.State) {
		return
	}
	resp.PlanValue = req.StateValue
}

func IsAllStateUnknown(ctx context.Context, state tfsdk.State) bool {
	attrs := state.Schema.GetAttributes()
	anyFound := false
	for k, _ := range attrs {
		attrValue := new(attr.Value)
		state.GetAttribute(ctx, path.Root(k), attrValue)
		if attrValue != nil && !(*attrValue).IsUnknown() && !(*attrValue).IsNull() {
			anyFound = true
			break
		}
	}
	return !anyFound
}

func SuppressDiffBool() planmodifier.Bool {
	return suppressDiffBool{}
}

type suppressDiffBool struct{}

func (m suppressDiffBool) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

func (m suppressDiffBool) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

func (m suppressDiffBool) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	log.Printf("[DEBUG] resp: %v", resp)
	log.Printf("[DEBUG] state: %v", req.Path)
	if !req.StateValue.IsNull() {
		if resp.PlanValue != req.StateValue {
			resp.Diagnostics.AddWarning("Attribute not editable changed "+req.Path.String(), "An attribute that is not editable has been changed, you may lose consistency because of it.")
		}
		return
	}
	if req.PlanValue.IsUnknown() {
		return
	}
	if resp.PlanValue.IsUnknown() {
		return
	}
	if req.ConfigValue.IsUnknown() {
		return
	}
	if IsAllStateUnknown(ctx, req.State) {
		return
	}
	resp.PlanValue = req.StateValue
}

// SetNil returns a plan modifier that propagates a state value into the planned value, when it is Known, and the Plan Value is Unknown
func SetNil() planmodifier.Set {
	return setNil{}
}

// setNil implements the plan modifier.
type setNil struct{}

// Description returns a human-readable description of the plan modifier.
func (m setNil) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m setNil) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyString implements the plan modification logic.
func (m setNil) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	// req.StateValue = ctx.
	// resp.PlanValue = req.StateValue
}

func SuppressDiffObj() planmodifier.Object {
	return suppressDiffObj{}
}

// suppressDiffObj implements the plan modifier.
type suppressDiffObj struct{}

// Description returns a human-readable description of the plan modifier.
func (m suppressDiffObj) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m suppressDiffObj) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifyObject implements the plan modification logic.
func (m suppressDiffObj) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	// Do nothing if there is a known planned value.
	// log.Printf("[DEBUG] resp: %v", resp.PlanValue)
	// log.Printf("[DEBUG] state: %v", req.StateValue)
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value
	if req.ConfigValue.IsUnknown() {
		return
	}

	if IsAllStateUnknown(ctx, req.State) {
		return
	}

	resp.PlanValue = req.StateValue
}

func SuppressDiffSet() planmodifier.Set {
	return suppressDiffSet{}
}

// suppressDiffSet implements the plan modifier.
type suppressDiffSet struct{}

// Description returns a human-readable description of the plan modifier.
func (m suppressDiffSet) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// MarkdownDescription returns a markdown description of the plan modifier.
func (m suppressDiffSet) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

// PlanModifySet implements the plan modification logic.
func (m suppressDiffSet) PlanModifySet(ctx context.Context, req planmodifier.SetRequest, resp *planmodifier.SetResponse) {
	// log.Printf("[DEBUG] resp: %v", resp)
	// log.Printf("[DEBUG] state: %v", req.Path)
	if !req.StateValue.IsNull() {
		if !reflect.DeepEqual(resp.PlanValue, req.StateValue) {
			resp.Diagnostics.AddWarning("Attribute not editable changed "+req.Path.String(), "An attribute that is not editable has been changed, you may lose consistency because of it.")
			return
		}
		return
	}
	// Do nothing if there is a known planned value.

	if req.StateValue.IsNull() && !resp.PlanValue.IsNull() {
		// log.Printf("[DEBUG] resp: %v", resp.PlanValue)
		// log.Printf("[DEBUG] state: %v", req.StateValue)
		req.StateValue = resp.PlanValue
		return
	}
	if !req.PlanValue.IsUnknown() {
		return
	}

	// Do nothing if there is an unknown configuration value
	if req.ConfigValue.IsUnknown() {
		return
	}

	if IsAllStateUnknown(ctx, req.State) {
		return
	}

	resp.PlanValue = req.StateValue
}

func merge(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse, target interface{}) {
	var plan types.Object
	var state types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(state.As(ctx, target, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)
	if resp.Diagnostics.HasError() {
		return
	}

	// we need a tftypes.Value for this Object to be able to use it with
	// our reflection code
	obj := types.ObjectType{AttrTypes: plan.AttributeTypes(ctx)}
	val, err := plan.ToTerraformValue(ctx)
	if err != nil {
		resp.Diagnostics.Append(diag.NewErrorDiagnostic("Object Conversion Error", "An unexpected error was encountered trying to convert object. This is always an error in the provider. Please report the following to the provider developer:\n\n"+err.Error()))
		return
	}
	resp.Diagnostics.Append(tfsdkr.Into(ctx, obj, val, target, tfsdkr.Options{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	}, path.Empty())...)
}

func dereferencePtr(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr {
		log.Printf("Dereference")
		return v.Elem()
	}
	log.Printf("Dereference NO")
	return v
}

func replaceUnknownFields(data interface{}) interface{} {
	val := reflect.ValueOf(data)
	// log.Printf("Name: %v", val.Type().Name())
	// log.Printf("val: %v", val.Interface())
	// log.Printf("val: %v", val.Interface())
	// log.Printf("val: %v", val.IsValid())
	// log.Printf("val: %v", val.IsValid())
	// log.Printf("val: %v", val.Kind())
	if val.Kind() != reflect.Ptr || val.IsNil() {
		fmt.Println("expecting a non-null pointer.")
		return nil
	}

	val = dereferencePtr(val)

	if val.Kind() != reflect.Slice {
		fmt.Println("expecting a slice.")
		return nil
	}

	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)

		if elem.Kind() != reflect.Struct {
			fmt.Println("expecting some structure in the slice.")
			return nil
		}

		for j := 0; j < elem.NumField(); j++ {
			field := elem.Field(j)
			if field.Kind() == reflect.String && field.Interface().(string) == "<unknown>" {
				field.Set(reflect.Zero(field.Type()))
			}
		}
	}

	return val
}

// isNullOrZero checks if a value is null, zero, or represents an unknown/empty state
func isNullOrZero(value reflect.Value) bool {
	if !value.IsValid() {
		return true
	}

	// Check for unknown values
	if value.CanInterface() {
		valStr := fmt.Sprint(value.Interface())
		if valStr == "<unknown>" || valStr == "<nil>" || valStr == "<null>" {
			return true
		}
	}

	// Check if it's zero value
	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

// prioritizeNonNullValues returns the non-null value, prioritizing b over a when both have values
func prioritizeNonNullValues(fieldA, fieldB reflect.Value) reflect.Value {
	isANull := isNullOrZero(fieldA)
	isBNull := isNullOrZero(fieldB)

	// If A is null and B is not null, use B
	if isANull && !isBNull {
		return fieldB
	}

	// If B is null and A is not null, use A
	if isBNull && !isANull {
		return fieldA
	}

	// If both are null, return A (preserve original null state)
	if isANull && isBNull {
		return fieldA
	}

	// If both have values, prioritize B
	return fieldB
}

// prioritizeNonNullValuesForMerge returns the non-null value, prioritizing a over b when both have values
// and prioritizing unknown values from a to get values from b if available, otherwise null
func prioritizeNonNullValuesForMerge(fieldA, fieldB reflect.Value) reflect.Value {
	isANull := isNullOrZero(fieldA)
	isBNull := isNullOrZero(fieldB)

	// Check if fieldA is unknown
	isAUnknown := false
	if fieldA.CanInterface() {
		valStr := fmt.Sprint(fieldA.Interface())
		isAUnknown = valStr == "<unknown>"
	}

	// Check if fieldA is empty string and treat it as null
	isAEmptyString := false
	if fieldA.Kind() == reflect.String {
		if fieldA.CanInterface() {
			valStr := fieldA.Interface().(string)
			isAEmptyString = valStr == ""
		}
	}

	// If A is unknown, prioritize B if available, otherwise return null
	if isAUnknown {
		if !isBNull {
			return fieldB
		}
		// Return null for unknown values when B is not available
		return reflect.Zero(fieldA.Type())
	}

	// If A is empty string, treat it as null and use B if available, otherwise return null
	if isAEmptyString {
		if !isBNull {
			return fieldB
		}
		// Return null for empty string values when B is not available
		return reflect.Zero(fieldA.Type())
	}

	// If A is null and B is not null, use B
	// if isANull && !isBNull {
	// 	return fieldB
	// }

	// If B is null and A is not null, use A
	if isBNull && !isANull {
		return fieldA
	}

	// If both are null, return null (except for List/Set which return empty)
	if isANull && isBNull {
		// If it's a List or Set, return an empty slice/set
		if fieldA.Type() == reflect.TypeOf(types.List{}) || fieldA.Type() == reflect.TypeOf(types.Set{}) {
			return reflect.ValueOf(createEmptySetOrList(fieldA.Type()))
		}
		// For all other types, return zero value (which represents null in Terraform)
		return reflect.Zero(fieldA.Type())
	}

	// If both have values, prioritize A (original behavior for mergeInterfaces)
	return fieldA
}

// prioritizeNonNullValuesForMergeWithPriority allows specifying whether to prioritize A or B when both have values
func prioritizeNonNullValuesForMergeWithPriority(fieldA, fieldB reflect.Value, prioritizeA bool) reflect.Value {
	isANull := isNullOrZero(fieldA)
	isBNull := isNullOrZero(fieldB)
	log.Printf("About to check if fieldA is nil")
	log.Printf("fieldA is nil: %v", isANull)
	log.Printf("fieldB is nil: %v", isBNull)

	// Check if fieldA is unknown
	isAUnknown := false
	if fieldA.CanInterface() {
		valStr := fmt.Sprint(fieldA.Interface())
		isAUnknown = valStr == "<unknown>"
	}

	// Check if fieldA is empty string and treat it as null
	isAEmptyString := false
	if fieldA.Kind() == reflect.String {
		if fieldA.CanInterface() {
			valStr := fieldA.Interface().(string)
			isAEmptyString = valStr == ""
		}
	}

	// If A is unknown, prioritize B if available, otherwise return null
	if isAUnknown {
		if !isBNull {
			return fieldB
		}
		// Return null for unknown values when B is not available
		return reflect.Zero(fieldA.Type())
	}

	// If A is empty string, treat it as null and use B if available, otherwise return null
	if isAEmptyString {
		// if !isBNull {
		// 	return fieldB
		// }
		// Return null for empty string values when B is not available
		return reflect.Zero(fieldA.Type())
	}
	if isANull && !isBNull && (reflect.TypeOf(types.List{}) != fieldA.Type() && reflect.TypeOf(types.Set{}) != fieldA.Type()) {
		return reflect.Zero(fieldB.Type())
	} else if isANull && !isBNull && (reflect.TypeOf(types.List{}) == fieldA.Type() || reflect.TypeOf(types.Set{}) == fieldA.Type()) {
		return reflect.ValueOf(createEmptySetOrList(fieldA.Type()))
	}

	// If A is null and B is not null, use B
	// if isANull && !isBNull {
	// 	return fieldB
	// }

	// If B is null and A is not null, use A
	if isBNull && !isANull {
		return fieldA
	}

	// If both are null, return null (except for List/Set which return empty)
	if isANull && isBNull {
		// If it's a List or Set, return an empty slice/set
		if fieldA.Type() == reflect.TypeOf(types.List{}) || fieldA.Type() == reflect.TypeOf(types.Set{}) {
			return reflect.ValueOf(createEmptySetOrList(fieldA.Type()))
		}
		// For all other types, return zero value (which represents null in Terraform)
		return reflect.Zero(fieldA.Type())
	}

	// If both have values, use the specified priority
	if prioritizeA {
		return fieldA
	}
	log.Printf("About to check if fieldA is an Struct want to know the type and kind: %v", fieldA.Type())
	log.Printf("Field A kind: %v", fieldA.Kind())
	// Verify if is an Struct
	if fieldA.Kind() == reflect.Slice {
		log.Printf("Field A is an Slice")
		log.Printf("About to iterate through slice elements")
		for i := 0; i < fieldA.Len(); i++ {
			sliceElemA := fieldA.Index(i)
			sliceElemB := fieldB.Index(i)
			log.Printf("Processing slice element %d", i)
			if sliceElemA.Kind() == reflect.Struct {
				foreachStructFieldBoth(sliceElemA, sliceElemB)
			}
		}
		log.Printf("Finished iterating through slice elements")
	} else {
		if fieldA.Kind() == reflect.Struct {
			//FieldName print
			log.Printf("Field A is an Struct")
			log.Printf("About to call foreachStructField with both fieldA and fieldB")
			foreachStructFieldBoth(fieldA, fieldB)
			log.Printf("Finished foreachStructField call")
		} else {
			// si fieldA es null, hacer que fieldB sea null}
			log.Printf("Field A is not an Struct")
			log.Printf("About to check if fieldA is nil")
			log.Printf("fieldA is nil: %v", fieldA.IsNil())
			if fieldA.IsNil() {
				fieldB.Set(reflect.Zero(fieldB.Type()))
			}

		}
	}

	return fieldB
}

func foreachStructField(field reflect.Value) {
	log.Printf("foreachStructField called with field type: %v", field.Type())
	log.Printf("Field kind: %v, NumField: %d", field.Kind(), field.NumField())

	for i := 0; i < field.NumField(); i++ {
		fieldName := field.Type().Field(i).Name
		fieldValue := field.Field(i)
		log.Printf("Campo: %s", fieldName)
		log.Printf("Field value kind: %v", fieldValue.Kind())

		// Handle direct struct
		if fieldValue.Kind() == reflect.Struct {
			foreachStructField(fieldValue)
		}

		// Handle pointer to struct
		if fieldValue.Kind() == reflect.Ptr && !fieldValue.IsNil() {
			elem := fieldValue.Elem()
			if elem.Kind() == reflect.Struct {
				log.Printf("  -> Diving into pointer to struct: %s", fieldName)
				foreachStructField(elem)
			}
		}
	}
}

func foreachStructFieldBoth(fieldA, fieldB reflect.Value) {
	log.Printf("foreachStructFieldBoth called with fieldA type: %v", fieldA.Type())
	log.Printf("fieldA kind: %v, NumField: %d", fieldA.Kind(), fieldA.NumField())
	log.Printf("fieldB type: %v", fieldB.Type())
	log.Printf("fieldB kind: %v, NumField: %d", fieldB.Kind(), fieldB.NumField())

	for i := 0; i < fieldA.NumField(); i++ {
		fieldTypeA := fieldA.Type().Field(i)
		fieldTypeB := fieldB.Type().Field(i)

		// Skip unexported fields
		if !fieldTypeA.IsExported() || !fieldTypeB.IsExported() {
			log.Printf("Skipping unexported field: %s", fieldTypeA.Name)
			continue
		}

		fieldNameA := fieldTypeA.Name
		fieldValueA := fieldA.Field(i)
		fieldNameB := fieldTypeB.Name
		fieldValueB := fieldB.Field(i)

		log.Printf("Comparing field %s (A) and field %s (B)", fieldNameA, fieldNameB)
		log.Printf("Field A value kind: %v", fieldValueA.Kind())
		log.Printf("Field B value kind: %v", fieldValueB.Kind())

		// Handle direct structs
		if fieldValueA.Kind() == reflect.Struct && fieldValueB.Kind() == reflect.Struct {
			log.Printf("Both fields are structs, diving into recursion")
			foreachStructFieldBoth(fieldValueA, fieldValueB)
		}

		// Handle pointer to structs
		if fieldValueA.Kind() == reflect.Ptr && !fieldValueA.IsNil() && fieldValueB.Kind() == reflect.Ptr && !fieldValueB.IsNil() {
			elemA := fieldValueA.Elem()
			elemB := fieldValueB.Elem()
			if elemA.Kind() == reflect.Struct && elemB.Kind() == reflect.Struct {
				log.Printf("Both fields are pointers to structs, diving into recursion")
				foreachStructFieldBoth(elemA, elemB)
			}
		}

		// Handle direct structs and pointers to structs
		if fieldValueA.Kind() == reflect.Struct && fieldValueB.Kind() == reflect.Ptr && !fieldValueB.IsNil() {
			elemB := fieldValueB.Elem()
			if elemB.Kind() == reflect.Struct {
				log.Printf("Field A is struct, Field B is pointer to struct, diving into recursion")
				foreachStructFieldBoth(fieldValueA, elemB)
			}
		}

		// Handle pointers to structs and direct structs
		if fieldValueA.Kind() == reflect.Ptr && !fieldValueA.IsNil() && fieldValueB.Kind() == reflect.Struct {
			elemA := fieldValueA.Elem()
			if elemA.Kind() == reflect.Struct {
				log.Printf("Field A is pointer to struct, Field B is struct, diving into recursion")
				foreachStructFieldBoth(elemA, fieldValueB)
			}
		}

		// Handle simple types - safely check if we can access the values
		if fieldValueA.Kind() == reflect.String && fieldValueB.Kind() == reflect.String {
			// Only try to access values if they are exported fields
			if fieldValueA.CanInterface() && fieldValueB.CanInterface() {
				valueA := fieldValueA.Interface().(string)
				valueB := fieldValueB.Interface().(string)
				log.Printf("String values - A: %s, B: %s", valueA, valueB)

				if valueA == "<unknown>" {
					fieldValueA.Set(reflect.Zero(fieldValueA.Type()))
				}
				if valueB == "<unknown>" {
					fieldValueB.Set(reflect.Zero(fieldValueB.Type()))
				}
			}
		}

		// Handle null/nil values - if A is null, make B null too
		// Only call IsNil() on types that support it
		if (fieldValueA.Kind() == reflect.Ptr || fieldValueA.Kind() == reflect.Interface ||
			fieldValueA.Kind() == reflect.Slice || fieldValueA.Kind() == reflect.Map ||
			fieldValueA.Kind() == reflect.Chan) && fieldValueA.IsNil() &&
			!fieldValueB.IsNil() {
			log.Printf("Field A is null, setting field B to null as well")
			fieldValueB.Set(reflect.Zero(fieldValueB.Type()))
		}

		// Handle zero values for basic types - if A is zero, make B zero too
		// Only check zero values for types that can be compared
		if fieldValueA.CanInterface() && fieldValueB.CanInterface() {
			// Check if A is a zero value for its type
			zeroValue := reflect.Zero(fieldValueA.Type())
			if reflect.DeepEqual(fieldValueA.Interface(), zeroValue.Interface()) {
				log.Printf("Field A is zero value, setting field B to zero as well")
				fieldValueB.Set(zeroValue)
			}
		}
	}
}

// safeSetField safely sets a field value, handling pointer types correctly
func safeSetField(field reflect.Value, value interface{}) bool {
	if !field.CanSet() {
		return false
	}

	valueReflect := reflect.ValueOf(value)

	// If the field expects a pointer and we have a value
	if field.Type().Kind() == reflect.Ptr {
		if valueReflect.Kind() != reflect.Ptr {
			// Create a pointer to the value
			ptr := reflect.New(field.Type().Elem())
			ptr.Elem().Set(valueReflect)
			field.Set(ptr)
		} else {
			// Both are pointers, set directly
			field.Set(valueReflect)
		}
	} else {
		// Field expects a value, dereference if needed
		if valueReflect.Kind() == reflect.Ptr {
			field.Set(valueReflect.Elem())
		} else {
			field.Set(valueReflect)
		}
	}

	return true
}

func isUnknown(value reflect.Value) bool {
	if value.CanInterface() {
		valStr := fmt.Sprint(value.Interface())
		return valStr == "<unknown>"
	}
	return false
}

func mergeInterfaces(a, b interface{}, isFirstTime bool) interface{} {

	if a == nil {
		return a
	}
	if isUnknown(reflect.ValueOf(a)) {
		return b
	}
	if b == nil {
		return a
	}

	log.Printf("State: %v", PrintKeyValue(a))
	log.Printf("Resp: %v", PrintKeyValue(b))
	valueA := reflect.ValueOf(a)
	valueB := reflect.ValueOf(b)

	// Dereference pointers if necessary
	valueA = dereferencePtr(valueA)
	valueB = dereferencePtr(valueB)

	if !valueA.IsValid() {
		valueA = reflect.ValueOf(a)
	}
	if !valueB.IsValid() {
		valueB = reflect.ValueOf(b)
	}
	// Check if both values ​​are slices
	if valueA.Kind() == reflect.Slice && valueB.Kind() == reflect.Slice {
		lenA := valueA.Len()
		lenB := valueB.Len()

		// Check if slices have the same length
		if lenA != lenB {
			// If they are not the same length, we return the largest slice
			if lenB > lenA {
				return a
			}
			return a
		}

		// Mix the elements of the slices
		resultSlice := reflect.MakeSlice(valueA.Type(), lenA, lenA)
		for i := 0; i < lenA; i++ {
			resultSlice.Index(i).Set(reflect.ValueOf(mergeInterfaces(valueA.Index(i).Interface(), valueB.Index(i).Interface(), false)))
		}
		return resultSlice.Interface()
	}

	if valueA.Kind() != reflect.Struct || valueB.Kind() != reflect.Struct {
		// Use the new prioritization logic for non-struct values
		prioritizedValue := prioritizeNonNullValuesForMerge(valueA, valueB)
		return prioritizedValue.Interface()
	}

	// Mixing fields of structs
	numFields := valueA.NumField()
	resultStruct := reflect.New(valueA.Type()).Elem()

	for i := 0; i < numFields; i++ {
		fieldA := valueA.Field(i)
		fieldB := valueB.Field(i)
		fieldA = dereferencePtr(fieldA)
		fieldB = dereferencePtr(fieldB)
		if !fieldA.IsValid() {
			fieldA = valueA.Field(i)
		}
		if !fieldB.IsValid() {
			fieldB = valueB.Field(i)
		}

		// log.Printf("Name: %v", valueB.Type().Field(i).Name)
		// log.Printf("ValueA: %v", valueA.Field(i).Interface())
		// log.Printf("ValueB: %v", valueB.Field(i).Interface())
		// log.Printf("ValueA: %v", valueA.Field(i).IsValid())
		// log.Printf("ValueB: %v", valueB.Field(i).IsValid())
		// log.Printf("ValueB: %v", valueB.Field(i).Kind())

		// log.Printf("fieldA: %v", fieldA.IsValid())
		// log.Printf("fieldB: %v", fieldB.IsValid())
		// log.Printf("fieldA: %v", fieldA.Interface())
		// log.Printf("fieldB: %v", fieldB.Interface())
		// log.Printf("Kind: %v", fieldA.Kind())

		// Check if both fields are valid before proceeding
		if fieldA.IsValid() && fieldB.IsValid() {
			// Handle slices recursively
			if fieldA.Kind() == reflect.Slice && fieldB.Kind() == reflect.Slice && valueA.Field(i).Kind() == reflect.Slice && valueB.Field(i).Kind() == reflect.Slice {
				log.Printf("IF Slice:")
				// If both fields are slices, merge them recursively
				// In mergeInterfaces, always prioritize A (state) over B (respRead)
				if field := replaceUnknownFields(valueA.Field(i)); field != nil {
					log.Printf("IF Slice Replace:")
					if resultStruct.Field(i).CanSet() {
						resultStruct.Field(i).Set(reflect.ValueOf(field))
					}
				} else {
					log.Printf("Else Slice Replace:")
					// Even when replaceUnknownFields returns nil, prioritize A over B
					if resultStruct.Field(i).CanSet() {
						resultStruct.Field(i).Set(valueA.Field(i))
					}
				}
			} else if valueA.Field(i).Type() != reflect.TypeOf(types.String{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Bool{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Int64{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Float64{}) && !isSetOrListType(valueA.Field(i).Type()) {
				// Handle complex struct types recursively
				mergedValues := changeStructUnknowns(fieldA.Interface(), fieldB.Interface())
				safeSetField(resultStruct.Field(i), mergedValues)
			} else {
				// Use the new prioritization logic for simple types
				prioritizedValue := prioritizeNonNullValuesForMerge(fieldA, fieldB)
				safeSetField(resultStruct.Field(i), prioritizedValue.Interface())
			}
		} else {
			log.Printf("One or both fields are invalid.")
		}
	}

	if isFirstTime {
		for i := 0; i < numFields; i++ {
			fieldA := valueA.Field(i)
			fieldB := valueB.Field(i)
			// resultStruct.Field(i).Set(valueB.Field(i))
			// log.Printf("Before")

			// log.Printf("After ZERO")
			for _, path := range pathParams {
				// log.Printf("After FOR path  %v", path)
				// log.Printf("After FOR valueB.Type().Field(i).Name  %v", valueB.Type().Field(i).Name)
				if valueB.Type().Field(i).Name == path && fieldB.IsZero() && !fieldA.IsZero() && fmt.Sprint(fieldA.Interface()) != "<unknown>" {
					if resultStruct.Field(i).CanSet() {
						resultStruct.Field(i).Set(fieldA)
					}
				} else {
					// resultStruct.Field(i).Set(fieldB)
					if valueB.Type().Field(i).Name == path && valueB.Type().Field(i).Name != "ID" && fieldB.IsZero() {
						// log.Printf("After FOR fieldB.Type().Name() %v", valueB.Type().Field(i).Name)
						// log.Printf("After FOR fieldB.Type().Value() %v", fieldB.Interface())
						// log.Printf("After FOR fieldB.Type().Value() 22%v", fieldB.IsZero())
						if valueB.FieldByName("ID").IsValid() {
							if !valueB.FieldByName("ID").IsZero() {
								if resultStruct.Field(i).CanSet() {
									resultStruct.Field(i).Set(valueB.FieldByName("ID"))
								}
							}
						}

					}
				}
			}
		}
	}
	log.Printf("[DEBUG] result: %v", PrintKeyValue(resultStruct.Interface()))

	return resultStruct.Interface()
}

// func verifyIsSimpleStruct(typeVar string){
// 	for
// }

func changeStructUnknowns(a interface{}, b interface{}) interface{} {
	valueA := reflect.ValueOf(a)
	valueB := reflect.ValueOf(b)

	// Handle nil/zero values early
	if !valueA.IsValid() || valueA.IsZero() {
		if isUnknown(valueA) {
			return b
		}
		if valueB.IsValid() && !valueB.IsZero() {
			return a
		}
		return a
	}

	if !valueB.IsValid() || valueB.IsZero() {
		return a
	}

	if valueA.Kind() == reflect.Ptr {
		log.Printf("valueA.Kind() == reflect.Ptr")
		valueA = valueA.Elem()
	}
	if valueB.Kind() == reflect.Ptr {
		log.Printf("valueB.Kind() == reflect.Ptr")
		valueB = valueB.Elem()
	}

	// Check if values are still valid after dereferencing
	if !valueA.IsValid() || !valueB.IsValid() {
		if valueA.IsValid() {
			return a
		}
		if valueB.IsValid() {
			return b
		}
		return a
	}

	resultStruct := reflect.New(valueA.Type()).Elem()
	if valueA.Kind() == reflect.Slice && valueB.Kind() == reflect.Slice {
		lenA := valueA.Len()
		lenB := valueB.Len()

		// Check if slices have the same length
		if lenA != lenB {
			// If they are not the same length, we return the largest slice
			if lenB > lenA {
				return b
			}
			return a
		}

		// Mix the elements of the slices
		resultSlice := reflect.MakeSlice(valueA.Type(), lenA, lenA)
		for i := 0; i < lenA; i++ {
			resultSlice.Index(i).Set(reflect.ValueOf(changeStructUnknowns(valueA.Index(i).Interface(), valueB.Index(i).Interface())))
		}
		return resultSlice.Interface()
	}
	for i := 0; i < valueA.NumField(); i++ {
		// log.Printf("In 2: ")
		fieldValueA := valueA.Field(i)
		// log.Printf("In 2: fieldValueA %s", fieldValueA.Interface())
		// log.Printf("In 2: fieldValueA %s", fieldValueA.Kind())
		// log.Printf("In 2: fieldValueB %v", valueB.IsValid())
		// log.Printf("In 2: fieldValueB %v", valueB.Field(i).Kind())

		var fieldValueB reflect.Value

		if valueB.IsValid() && i < valueB.NumField() {
			fieldValueB = valueB.Field(i)
		} else {
			fieldValueB = valueA.Field(i)
		}

		fieldValueA = dereferencePtr(fieldValueA)
		fieldValueB = dereferencePtr(fieldValueB)
		if !fieldValueA.IsValid() {
			fieldValueA = valueA.Field(i)
		}
		if !fieldValueB.IsValid() {

			fieldValueB = valueB.Field(i)
		}

		// Get the field name - add safety check
		if !valueA.IsValid() || i >= valueA.Type().NumField() {
			continue
		}
		fieldName := valueA.Type().Field(i).Name
		log.Printf("In 2: fieldName %s", fieldName)

		// Check if we can safely call Interface() on the field
		canInterfaceA := fieldValueA.IsValid() && fieldValueA.CanInterface()
		canInterfaceB := fieldValueB.IsValid() && fieldValueB.CanInterface()

		// Use the new prioritization logic for simple types
		if fieldValueA.IsValid() && (fieldValueA.Type() == reflect.TypeOf(types.String{}) || fieldValueA.Type() == reflect.TypeOf(types.Bool{}) || fieldValueA.Type() == reflect.TypeOf(types.Int64{}) || fieldValueA.Type() == reflect.TypeOf(types.Float64{}) || isSetOrListType(fieldValueA.Type())) {
			prioritizedValue := prioritizeNonNullValuesForMerge(fieldValueA, fieldValueB)
			safeSetField(resultStruct.Field(i), prioritizedValue.Interface())
		} else {
			// Handle complex struct types recursively
			if fieldValueA.IsValid() && fieldValueA.Type() != reflect.TypeOf(types.String{}) && fieldValueA.Type() != reflect.TypeOf(types.Bool{}) && fieldValueA.Type() != reflect.TypeOf(types.Int64{}) && fieldValueA.Type() != reflect.TypeOf(types.Float64{}) && !isSetOrListType(fieldValueA.Type()) {
				if canInterfaceA && canInterfaceB {
					nestedResult := changeStructUnknowns(fieldValueA.Interface(), fieldValueB.Interface())
					safeSetField(resultStruct.Field(i), nestedResult)
				} else {
					// If we can't interface, use prioritization logic
					prioritizedValue := prioritizeNonNullValuesForMerge(fieldValueA, fieldValueB)
					safeSetField(resultStruct.Field(i), prioritizedValue.Interface())
				}
			} else {
				// For simple types, use prioritization logic
				prioritizedValue := prioritizeNonNullValuesForMerge(fieldValueA, fieldValueB)
				safeSetField(resultStruct.Field(i), prioritizedValue.Interface())
			}
		}
	}
	// if resultStruct.Kind() != reflect.Ptr {
	// 	log.Print("Not a pointer: ", resultStruct.Kind())
	// 	resultStruct = reflect.ValueOf(resultStruct)
	// 	log.Print("Pointer: ", resultStruct.Kind())
	// }
	return resultStruct.Interface()
}

func mergeInterfacesOnlyPath(a, b interface{}) interface{} {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	log.Printf("StateRead: %v", PrintKeyValue(a))
	log.Printf("RespRead: %v", PrintKeyValue(b))
	valueA := reflect.ValueOf(a)
	valueB := reflect.ValueOf(b)
	numFields := valueA.NumField()
	resultStruct := reflect.New(valueA.Type()).Elem()

	for i := 0; i < numFields; i++ {
		fieldA := valueA.Field(i)
		fieldB := valueB.Field(i)
		fieldA = dereferencePtr(fieldA)
		fieldB = dereferencePtr(fieldB)
		if !fieldA.IsValid() {
			fieldA = valueA.Field(i)
		}
		if !fieldB.IsValid() {
			fieldB = valueB.Field(i)
		}
		log.Printf("fieldname %v", valueA.Type().Field(i).Name)
		log.Printf("fieldA %v", fieldA)
		log.Printf("fieldB %v", fieldB)
		log.Printf("fieldB %v", fieldB)

		// Check if both fields are valid before proceeding
		if fieldA.IsValid() && fieldB.IsValid() {
			// Handle slices
			if reflect.TypeOf(fieldB).Kind() == reflect.Slice {
				// Get the length of the slice using reflect
				fieldBValue := reflect.ValueOf(fieldB)
				if fieldBValue.Kind() == reflect.Ptr && !fieldBValue.IsNil() {
					fieldBValue = fieldBValue.Elem()
				}
				if fieldBValue.Kind() == reflect.Slice {
					length := fieldBValue.Len()
					if length > 0 {
						if resultStruct.Field(i).CanSet() {
							resultStruct.Field(i).Set(valueA.Field(i))
						}
					}
				}
			} else {
				// Use the merge prioritization logic for non-slice types
				// In mergeInterfacesOnlyPath, prioritize B (respRead) over A (state)
				prioritizedValue := prioritizeNonNullValuesForMergeWithPriority(fieldA, fieldB, false)
				if resultStruct.Field(i).CanSet() {
					// Check if the prioritized value is null/zero and the field type is a pointer
					if isNullOrZero(prioritizedValue) && resultStruct.Field(i).Type().Kind() == reflect.Ptr {
						// Set to nil pointer for null values
						resultStruct.Field(i).Set(reflect.Zero(resultStruct.Field(i).Type()))
					} else {
						// Handle type compatibility for slices and structs
						if prioritizedValue.Kind() == reflect.Slice && resultStruct.Field(i).Type().Kind() == reflect.Ptr {
							// Field expects pointer to slice, but we have a slice
							if prioritizedValue.IsNil() {
								resultStruct.Field(i).Set(prioritizedValue)
							} else {
								// Create a pointer to the slice
								slicePtr := reflect.New(prioritizedValue.Type())
								slicePtr.Elem().Set(prioritizedValue)
								resultStruct.Field(i).Set(slicePtr)
							}
						} else if prioritizedValue.Kind() == reflect.Struct && resultStruct.Field(i).Type().Kind() == reflect.Ptr {
							// Field expects pointer to struct, but we have a struct
							// Create a pointer to the struct
							structPtr := reflect.New(prioritizedValue.Type())
							structPtr.Elem().Set(prioritizedValue)
							resultStruct.Field(i).Set(structPtr)
						} else {
							resultStruct.Field(i).Set(prioritizedValue)
						}
					}
				}
			}
		} else if fieldA.IsValid() && !fieldB.IsValid() {
			// If fieldA is valid but fieldB is not valid (null), keep fieldA
			log.Printf("IF Keep A when B is invalid/null:")
			if resultStruct.Field(i).CanSet() {
				// Handle type compatibility for slices and structs
				if fieldA.Kind() == reflect.Slice && resultStruct.Field(i).Type().Kind() == reflect.Ptr {
					// Field expects pointer to slice, but we have a slice
					if fieldA.IsNil() {
						resultStruct.Field(i).Set(fieldA)
					} else {
						// Create a pointer to the slice
						slicePtr := reflect.New(fieldA.Type())
						slicePtr.Elem().Set(fieldA)
						resultStruct.Field(i).Set(slicePtr)
					}
				} else if fieldA.Kind() == reflect.Struct && resultStruct.Field(i).Type().Kind() == reflect.Ptr {
					// Field expects pointer to struct, but we have a struct
					// Create a pointer to the struct
					structPtr := reflect.New(fieldA.Type())
					structPtr.Elem().Set(fieldA)
					resultStruct.Field(i).Set(structPtr)
				} else {
					resultStruct.Field(i).Set(valueA.Field(i))
				}
			}
		} else if !fieldA.IsValid() && fieldB.IsValid() {
			// If fieldA is not valid (null) but fieldB is valid, use fieldB
			log.Printf("IF Use B when A is invalid/null:")
			if resultStruct.Field(i).CanSet() {
				// Handle type compatibility for slices and structs
				if fieldB.Kind() == reflect.Slice && resultStruct.Field(i).Type().Kind() == reflect.Ptr {
					// Field expects pointer to slice, but we have a slice
					if fieldB.IsNil() {
						resultStruct.Field(i).Set(fieldB)
					} else {
						// Create a pointer to the slice
						slicePtr := reflect.New(fieldB.Type())
						slicePtr.Elem().Set(fieldB)
						resultStruct.Field(i).Set(slicePtr)
					}
				} else if fieldB.Kind() == reflect.Struct && resultStruct.Field(i).Type().Kind() == reflect.Ptr {
					// Field expects pointer to struct, but we have a struct
					// Create a pointer to the struct
					structPtr := reflect.New(fieldB.Type())
					structPtr.Elem().Set(fieldB)
					resultStruct.Field(i).Set(structPtr)
				} else {
					resultStruct.Field(i).Set(fieldB)
				}
			}
		} else {
			// Both fields are invalid/null - set to nil pointer if field type is pointer
			log.Printf("Both fields are invalid/null - setting to nil")
			if resultStruct.Field(i).CanSet() && resultStruct.Field(i).Type().Kind() == reflect.Ptr {
				resultStruct.Field(i).Set(reflect.Zero(resultStruct.Field(i).Type()))
			}
		}

		log.Printf("Before")

		for _, path := range pathParams {
			// if valueB.Type().Field(i).Name == "OrganizationID" {
			// 	log.Printf("fieldB.IsZero() %v", fieldB.IsZero())
			// 	log.Printf("valueB.Type().Field(i).Name %v", valueB.Type().Field(i).Name)
			// 	log.Printf("valueB.Type().Field(i).Name %v", !fieldA.IsZero())
			// }

			if valueB.Type().Field(i).Name == path && fieldB.IsZero() && !fieldA.IsZero() && fmt.Sprint(fieldA.Interface()) != "<unknown>" {
				if resultStruct.Field(i).CanSet() {
					// Handle type compatibility for slices and structs
					fieldValue := valueA.Field(i)
					if fieldValue.Kind() == reflect.Slice && resultStruct.Field(i).Type().Kind() == reflect.Ptr {
						// Field expects pointer to slice, but we have a slice
						if fieldValue.IsNil() {
							resultStruct.Field(i).Set(fieldA)
						} else {
							// Create a pointer to the slice
							slicePtr := reflect.New(fieldValue.Type())
							slicePtr.Elem().Set(fieldValue)
							resultStruct.Field(i).Set(slicePtr)
						}
					} else if fieldValue.Kind() == reflect.Struct && resultStruct.Field(i).Type().Kind() == reflect.Ptr {
						// Field expects pointer to struct, but we have a struct
						// Create a pointer to the struct
						structPtr := reflect.New(fieldValue.Type())
						structPtr.Elem().Set(fieldValue)
						resultStruct.Field(i).Set(structPtr)
					} else {
						resultStruct.Field(i).Set(fieldValue)
					}
				}
			} else {
				if valueB.Type().Field(i).Name == path && valueB.Type().Field(i).Name != "ID" && fieldB.IsZero() {
					// Only use ID from B if A is also zero/null, otherwise preserve A's null state
					if fieldA.IsZero() || !fieldA.IsValid() {
						if valueB.FieldByName("ID").IsValid() {
							if !valueB.FieldByName("ID").IsZero() {
								if resultStruct.Field(i).CanSet() {
									resultStruct.Field(i).Set(valueB.FieldByName("ID"))
								}
							}
						}
					}
				}
			}

		}

	}
	log.Printf("Result Read: %v", PrintKeyValue(resultStruct.Interface()))
	return resultStruct.Interface()
}

func PrintKeyValue(v interface{}) string {
	val := reflect.ValueOf(v)

	// If v is a pointer, dereference it
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Make sure v is a structure
	if val.Kind() != reflect.Struct {
		return "PrintKeyValue: v is not a struct"
	}
	var buffer bytes.Buffer

	// Iterate over the fields of the structure
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		// Print the field name and its value
		buffer.WriteString(fmt.Sprintf("%s value is %v, ", fieldName, field.Interface()))
	}

	// Remove the last comma and space
	if buffer.Len() > 0 {
		buffer.Truncate(buffer.Len() - 2)
	}

	return buffer.String()
}

var _ validator.String = notEditable{}

// notEditable validates if the provided value is of type string and can be parsed as JSON.
type notEditable struct {
}

func (validator notEditable) ValidateString(ctx context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	// Only validate the attribute configuration value if it is known.
	log.Printf("REQ: %v", req.ConfigValue)

	// log.Printf("RESP: %v", resp.)
}

func (validator notEditable) Description(ctx context.Context) string {
	return "value must have exactly one child attribute defined"
}

func (validator notEditable) MarkdownDescription(ctx context.Context) string {
	return validator.Description(ctx)
}

// ExactlyOneChild returns an AttributeValidator which ensures that any configured
// attribute object has only one child attribute.
// Null (unconfigured) and unknown values are skipped.
func ExactlyOneChild() validator.String {
	return notEditable{}
}

func SuppressDiffList() planmodifier.List {
	return suppressDiffList{}
}

type suppressDiffList struct{}

func (m suppressDiffList) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

func (m suppressDiffList) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

func (m suppressDiffList) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
	log.Printf("[DEBUG] resp: %v", resp)
	log.Printf("[DEBUG] state: %v", req.Path)
	if !req.StateValue.IsNull() {
		if !reflect.DeepEqual(resp.PlanValue, req.StateValue) {
			resp.Diagnostics.AddWarning("Attribute not editable changed "+req.Path.String(), "An attribute that is not editable has been changed, you may lose consistency because of it.")
			return
		}
		return
	}
	if req.PlanValue.IsUnknown() {
		return
	}
	if resp.PlanValue.IsUnknown() {
		return
	}
	if req.ConfigValue.IsUnknown() {
		return
	}
	if IsAllStateUnknown(ctx, req.State) {
		return
	}
	resp.PlanValue = req.StateValue
}

func SuppressDiffBool2() planmodifier.Bool {
	return suppressDiffBool2{}
}

type suppressDiffBool2 struct{}

func (m suppressDiffBool2) Description(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

func (m suppressDiffBool2) MarkdownDescription(_ context.Context) string {
	return "Once set, the value of this attribute in state will not change."
}

func (m suppressDiffBool2) PlanModifyBool(ctx context.Context, req planmodifier.BoolRequest, resp *planmodifier.BoolResponse) {
	// log.Printf("[DEBUG] resp: %v", resp.PlanValue)
	// log.Printf("[DEBUG] resp2: %v", req.StateValue)
	// log.Printf("[DEBUG] path: %v", req.Path)
	if req.PlanValue.IsUnknown() {
		req.StateValue = resp.PlanValue
		return
	}
	if resp.PlanValue.IsUnknown() {
		resp.PlanValue = req.StateValue
		return
	}
	if req.ConfigValue.IsUnknown() {
		req.ConfigValue = resp.PlanValue
		return
	}
	if IsAllStateUnknown(ctx, req.State) {
		return
	}
	resp.PlanValue = req.StateValue
}
