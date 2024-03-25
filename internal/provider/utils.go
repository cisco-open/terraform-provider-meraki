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

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"reflect"
	tfsdkr "terraform-provider-meraki/internal/provider/reflects"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

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
}

type cmpFunc func(interface{}, interface{}) bool

// Simple comparison function that compares if two values are equal.
func simpleCmp(a, b interface{}) bool {
	return a == b
}

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
	if len(items) == 0 {
		return types.ListNull(types.StringType)
	}
	for _, item := range items {
		eles = append(eles, types.StringValue(item))
	}

	return types.ListValueMust(types.StringType, eles)
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

func elementsToStrings(ctx context.Context, elements types.List) []string {
	var strings []string
	elements.ElementsAs(ctx, &strings, false)

	return strings
}

// getDictResult function translates the Python function to Go.
func getDictResult(result interface{}, key string, value interface{}, cmpFn cmpFunc) interface{} {
	// Check if the result is a list.
	if reflect.TypeOf(result).Kind() == reflect.Slice {
		resultSlice := reflect.ValueOf(result)

		// If the list has a single element.
		if resultSlice.Len() == 1 {
			// Check if that element is a dictionary.
			if reflect.TypeOf(resultSlice.Index(0).Interface()).Kind() == reflect.Map {
				resultMap := resultSlice.Index(0).Interface().(map[string]interface{})
				// Check if the key and value match the provided ones.
				if val, ok := resultMap[key]; ok && !cmpFn(val, value) {
					return nil
				}
				return resultMap
			}
			return nil
		}

		// Iterate over the elements of the list.
		for i := 0; i < resultSlice.Len(); i++ {
			item := resultSlice.Index(i)
			// Check if the item is a dictionary.
			if reflect.TypeOf(item.Interface()).Kind() == reflect.Map {
				itemMap := item.Interface().(map[string]interface{})
				// Check if the key and value match the provided ones.
				if val, ok := itemMap[key]; !ok || cmpFn(val, value) {
					return itemMap
				}
			}
		}
		return nil
	}

	// If the result is not a list.
	if reflect.TypeOf(result).Kind() != reflect.Map {
		return nil
	}

	// Check if the result is a dictionary.
	resultMap := result.(map[string]interface{})
	// Check if the key and value match the provided ones.
	if val, ok := resultMap[key]; ok && !cmpFn(val, value) {
		return nil
	}

	return resultMap
}

func structToMap(data interface{}) []map[string]interface{} {
	// Check if data is a pointer, if so get the value
	if reflect.ValueOf(data).Kind() == reflect.Ptr {
		data = reflect.ValueOf(data).Elem().Interface()
	}

	// Check if a slice is passed
	val := reflect.ValueOf(data)

	// Get the type of the structure within the slice
	elementType := val.Type().Elem()
	if elementType.Kind() == reflect.Ptr {
		elementType = elementType.Elem()
	}

	// Convert each element of the slice to a map
	result := make([]map[string]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		element := val.Index(i)
		elementMap := make(map[string]interface{})
		for j := 0; j < element.NumField(); j++ {
			fieldName := elementType.Field(j).Name
			fieldValue := element.Field(j).Interface()
			elementMap[fieldName] = fieldValue
		}
		result[i] = elementMap
	}
	return result
}

func mapToStruct(data map[string]interface{}, target interface{}) error {
	targetType := reflect.TypeOf(target)
	if targetType.Kind() != reflect.Ptr || targetType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("the target is not a pointer to an struct")
	}
	targetValue := reflect.ValueOf(target).Elem()
	for i := 0; i < targetType.Elem().NumField(); i++ {
		field := targetType.Elem().Field(i)
		fieldName := field.Name
		value, ok := data[fieldName]
		if !ok {
			continue
		}
		fieldValue := reflect.ValueOf(value)
		if fieldValue.Type().ConvertibleTo(field.Type) {
			targetValue.FieldByName(fieldName).Set(fieldValue.Convert(field.Type))
		} else {
			return fmt.Errorf("the valur type %s is not on struct", fieldName)
		}
	}

	return nil
}
func int64ToIntPointer(i *int64) *int {
	if i == nil {
		return nil
	}
	a := int(*i)
	return &a
}

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
	resp.PlanValue = req.StateValue
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
	if !req.StateValue.IsNull() {
		if !reflect.DeepEqual(resp.PlanValue, req.StateValue) {
			resp.Diagnostics.AddWarning("Attribute not editable changed "+req.Path.String(), "An attribute that is not editable has been changed, you may lose consistency because of it.")
			return
		}
		return
	}
	// Do nothing if there is a known planned value.

	if req.StateValue.IsNull() && !resp.PlanValue.IsNull() {
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
		return v.Elem()
	}
	return v
}

func replaceUnknownFields(data interface{}) interface{} {
	val := reflect.ValueOf(data)

	if val.Kind() != reflect.Ptr || val.IsNil() {
		return nil
	}

	val = val.Elem()

	if val.Kind() != reflect.Slice {
		return nil
	}

	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)

		if elem.Kind() != reflect.Struct {
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

func mergeInterfaces(a, b interface{}, isFirstTime bool) interface{} {
	if a == nil {
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
	// Check if both values are slices
	if valueA.Kind() == reflect.Slice && valueB.Kind() == reflect.Slice {
		lenA := valueA.Len()
		lenB := valueB.Len()

		// Check if the slices have the same length
		if lenA != lenB {
			// If they do not have the same length, we return the largest slice
			if lenB > lenA {
				return b
			}
			return a
		}

		// Merge the elements of the slices
		resultSlice := reflect.MakeSlice(valueA.Type(), lenA, lenA)
		for i := 0; i < lenA; i++ {
			resultSlice.Index(i).Set(reflect.ValueOf(mergeInterfaces(valueA.Index(i).Interface(), valueB.Index(i).Interface(), false)))
		}
		return resultSlice.Interface()
	}

	if valueA.Kind() != reflect.Struct || valueB.Kind() != reflect.Struct {
		// If they are not slices or structs, we simply return the second value
		return b
	}

	// Merge the fields of the structs
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

		// Check if both fields are valid before proceeding
		if fieldA.IsValid() && fieldB.IsValid() {
			if fieldA.Kind() == reflect.Slice && fieldB.Kind() == reflect.Slice {
				// If both fields are slices, merge them recursively
				if field := replaceUnknownFields(valueA.Field(i)); field != nil {
					resultStruct.Field(i).Set(reflect.ValueOf(field))
				} else {
					resultStruct.Field(i).Set(valueB.Field(i))
				}
			} else if reflect.DeepEqual(fieldA.Interface(), reflect.Zero(fieldA.Type()).Interface()) && !reflect.DeepEqual(fieldB.Interface(), reflect.Zero(fieldB.Type()).Interface()) {
				// If the field of a is null and the field of b is not, we use the field of b
				resultStruct.Field(i).Set(valueB.Field(i))
			} else {
				if fmt.Sprint(fieldA.Interface()) == "<unknown>" || fmt.Sprint(fieldA.Interface()) == "<nil>" || fmt.Sprint(fieldA.Interface()) == "<null>" {
					resultStruct.Field(i).Set(valueB.Field(i))
				} else {
					if valueA.Field(i).Type() != reflect.TypeOf(types.String{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Bool{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Int64{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Float64{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Set{}) {
						mergedValues := changeStructUnknowns(fieldA.Interface(), fieldB.Interface())
						fieldValueBPtr := reflect.New(fieldB.Type())
						fieldValueBPtr.Elem().Set(reflect.ValueOf(mergedValues))
						resultStruct.Field(i).Set(fieldValueBPtr)
					} else {
						resultStruct.Field(i).Set(valueA.Field(i))
					}
				}
			}
		} else {
			log.Printf("One or both fields are invalid.")
		}
	}

	if isFirstTime {
		for i := 0; i < numFields; i++ {
			fieldA := valueA.Field(i)
			fieldB := valueB.Field(i)
			for _, path := range pathParams {
				if valueB.Type().Field(i).Name == path && fieldB.IsZero() && !fieldA.IsZero() && fmt.Sprint(fieldA.Interface()) != "<unknown>" {
					resultStruct.Field(i).Set(fieldA)
				} else {
					if valueB.Type().Field(i).Name == path && valueB.Type().Field(i).Name != "ID" && fieldB.IsZero() {
						if valueB.FieldByName("ID").IsValid() {
							if !valueB.FieldByName("ID").IsZero() {
								resultStruct.Field(i).Set(valueB.FieldByName("ID"))
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

func changeStructUnknowns(a interface{}, b interface{}) interface{} {
	valueA := reflect.ValueOf(a)
	valueB := reflect.ValueOf(b)
	if valueA.Kind() == reflect.Ptr {
		valueA = valueA.Elem()
	}
	if valueB.Kind() == reflect.Ptr {
		valueB = valueB.Elem()
	}

	if valueA.Kind() == reflect.Slice && valueB.Kind() == reflect.Slice {
		lenA := valueA.Len()
		lenB := valueB.Len()

		// Check if the slices have the same length
		if lenA != lenB {
			// If they do not have the same length, we return the largest slice
			if lenB > lenA {
				return b
			}
			return a
		}

		// Merge the elements of the slices
		resultSlice := reflect.MakeSlice(valueA.Type(), lenA, lenA)
		for i := 0; i < lenA; i++ {
			resultSlice.Index(i).Set(reflect.ValueOf(changeStructUnknowns(valueA.Index(i).Interface(), valueB.Index(i).Interface())))
		}
		return resultSlice.Interface()
	}

	resultStruct := reflect.New(valueA.Type()).Elem()

	for i := 0; i < valueA.NumField(); i++ {
		fieldValueA := valueA.Field(i)
		fieldValueB := valueB.Field(i)
		fieldValueA = dereferencePtr(fieldValueA)
		fieldValueB = dereferencePtr(fieldValueB)
		if !fieldValueA.IsValid() {
			fieldValueA = valueA.Field(i)
		}
		if !fieldValueB.IsValid() {
			fieldValueB = valueB.Field(i)
		}

		// Get the field name
		// fieldName := valueA.Type().Field(i).Name

		if fmt.Sprint(fieldValueA.Interface()) == "<unknown>" {
			resultStruct.Field(i).Set(fieldValueB)
		} else {
			if valueA.Field(i).Type() != reflect.TypeOf(types.String{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Bool{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Int64{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Float64{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Set{}) {
				nestedResult := changeStructUnknowns(fieldValueA.Interface(), fieldValueB.Interface())
				fieldValueBPtr := reflect.New(fieldValueB.Type())
				fieldValueBPtr.Elem().Set(reflect.ValueOf(nestedResult))
				resultStruct.Field(i).Set(fieldValueBPtr)
			} else {
				// log.Printf("Assigned %v to field %v\n", fieldValueA.Interface(), fieldName)
				resultStruct.Field(i).Set(fieldValueA)
			}
		}
	}
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

		if reflect.TypeOf(fieldB).Kind() == reflect.Slice {
			// Get the length of the slice using reflection
			length := reflect.ValueOf(fieldB).Elem().Len()
			if length > 0 {
				resultStruct.Field(i).Set(valueA.Field(i))
			}
		} else {
			resultStruct.Field(i).Set(valueB.Field(i))
		}

		for _, path := range pathParams {
			if valueB.Type().Field(i).Name == path && fieldB.IsZero() && !fieldA.IsZero() && fmt.Sprint(fieldA.Interface()) != "<unknown>" {
				resultStruct.Field(i).Set(fieldA)
			} else {
				if valueB.Type().Field(i).Name == path && valueB.Type().Field(i).Name != "ID" && fieldB.IsZero() {
					if valueB.FieldByName("ID").IsValid() {
						if !valueB.FieldByName("ID").IsZero() {
							resultStruct.Field(i).Set(valueB.FieldByName("ID"))
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

	// Ensure that v is a structure
	if val.Kind() != reflect.Struct {
		return "PrintKeyValue: v is not a structure"
	}
	var buffer bytes.Buffer

	// Iterate over the fields of the structure
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		// Print the name of the field and its value
		buffer.WriteString(fmt.Sprintf("%s value is %v, ", fieldName, field.Interface()))
	}

	// Remove the last comma and space
	if buffer.Len() > 0 {
		buffer.Truncate(buffer.Len() - 2)
	}

	return buffer.String()
}
