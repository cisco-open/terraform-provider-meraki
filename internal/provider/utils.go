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

func elementsToStrings(ctx context.Context, elements types.List) []string {
	var strings []string
	elements.ElementsAs(ctx, &strings, false)

	return strings
}

// getDictResult busca y devuelve un diccionario específico dentro de una estructura de datos flexible.
func getDictResult(result interface{}, key string, value interface{}, cmpFn cmpFunc) interface{} {
	// Mutex local para sincronizar el acceso al mapa
	var mu sync.Mutex

	// Mapa para almacenar el resultado
	resultMap := make(map[string]interface{})

	// Función para verificar si un valor es un mapa[string]interface{}
	isMapStringInterface := func(v interface{}) (map[string]interface{}, bool) {
		if m, ok := v.(map[string]interface{}); ok {
			return m, true
		}
		return nil, false
	}

	// Función interna para procesar un elemento
	processItem := func(item interface{}) {
		if itemMap, ok := isMapStringInterface(item); ok {
			// Verificar si la clave y el valor coinciden con los proporcionados.
			log.Print(" Verificar si la clave y el valor coinciden con los proporcionados.")
			if val, ok := itemMap[key]; ok && cmpFn(val, value) {
				// Bloquear el acceso al mapa antes de realizar la asignación
				log.Print(" Bloquear el acceso al mapa antes de realizar la asignación")
				mu.Lock()
				defer mu.Unlock()
				resultMap = itemMap
			}
		}
	}

	// Verificar si el resultado es un slice.
	log.Print("Verificar si el resultado es un slice.")
	if reflect.TypeOf(result).Kind() == reflect.Slice {
		resultSlice := reflect.ValueOf(result)
		log.Print("Verificar si el resultado es un slice. 2")
		if resultSlice.Len() == 0 {
			return nil
		}
		log.Print("Si el slice tiene un solo elemento.")
		// Si el slice tiene un solo elemento.
		if resultSlice.Len() == 1 {
			if itemMap, ok := isMapStringInterface(resultSlice.Index(0).Interface()); ok {

				log.Print(" Verificar si la clave y el valor coinciden con los proporcionados.")
				// Verificar si la clave y el valor coinciden con los proporcionados.
				if val, ok := itemMap[key]; ok && !cmpFn(val, value) {
					return nil
				}
				return itemMap
			}
			return nil
		}

		// Iterar sobre los elementos del slice.
		log.Print("Iterar sobre los elementos del slice.")

		for i := 0; i < resultSlice.Len(); i++ {
			processItem(resultSlice.Index(i).Interface())
		}
		if len(reflect.ValueOf(resultMap).MapKeys()) == 0 {
			return nil
		}
		return resultMap
	}

	// Si el resultado no es un slice.
	if itemMap, ok := isMapStringInterface(result); ok {
		// Verificar si la clave y el valor coinciden con los proporcionados.
		if val, ok := itemMap[key]; ok && !cmpFn(val, value) {
			return nil
		}
		return itemMap
	}

	// Si el resultado no es ni un slice ni un mapa.
	return nil
}
func structToMap(data interface{}) []map[string]interface{} {
	// Mutex local para sincronizar el acceso al resultado
	var mu sync.Mutex

	// Verificar si data es un puntero, si es así obtener el valor
	if reflect.ValueOf(data).Kind() == reflect.Ptr {
		data = reflect.ValueOf(data).Elem().Interface()
	}

	// Verificar si se pasa una slice
	val := reflect.ValueOf(data)
	if val.Kind() != reflect.Slice {
		log.Printf("[DEBUG] El valor proporcionado no es una slice.")
		fmt.Println("El valor proporcionado no es una slice.")
		// return nil
	}

	// Obtener el tipo de la estructura dentro de la slice
	elementType := val.Type().Elem()
	if elementType.Kind() == reflect.Ptr {
		elementType = elementType.Elem()
	}
	if elementType.Kind() != reflect.Struct {
		log.Printf("[DEBUG] El tipo dentro de la slice no es una estructura.")
		fmt.Println("El tipo dentro de la slice no es una estructura.")
		// return nil
	}

	// Convertir cada elemento de la slice a un mapa
	result := make([]map[string]interface{}, val.Len())
	for i := 0; i < val.Len(); i++ {
		element := val.Index(i)

		// Función interna para convertir una estructura a un mapa
		structToMapFunc := func(elem reflect.Value) map[string]interface{} {
			elemMap := make(map[string]interface{})
			for j := 0; j < elem.NumField(); j++ {
				fieldName := elementType.Field(j).Name
				fieldValue := elem.Field(j).Interface()
				elemMap[fieldName] = fieldValue
			}
			return elemMap
		}

		// Bloquear el acceso al resultado antes de asignar
		mu.Lock()
		result[i] = structToMapFunc(element)
		mu.Unlock()
	}

	return result
}

// mapToStruct asigna valores desde un mapa a una estructura (target)
func mapToStruct(data map[string]interface{}, target interface{}) error {
	// Mutex local para sincronizar el acceso a la estructura de destino
	var mu sync.Mutex

	// Verificar si target es un puntero a una estructura
	targetType := reflect.TypeOf(target)
	if targetType.Kind() != reflect.Ptr || targetType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("El objetivo no es un puntero a una estructura")
	}
	targetValue := reflect.ValueOf(target).Elem()

	// Iterar sobre los campos de la estructura de destino
	for i := 0; i < targetType.Elem().NumField(); i++ {
		field := targetType.Elem().Field(i)
		fieldName := field.Name

		// Obtener el valor del mapa para el nombre del campo
		value, ok := data[fieldName]
		if !ok {
			continue // Si no existe en el mapa, continuar con el siguiente campo
		}

		// Función interna para asignar el valor al campo de la estructura
		assignValue := func(field reflect.Value, value interface{}) error {
			// Bloquear el acceso al campo antes de la asignación
			mu.Lock()
			defer mu.Unlock()

			fieldValue := reflect.ValueOf(value)

			// Verificar si el tipo de valor es convertible al tipo del campo
			if fieldValue.Type().ConvertibleTo(field.Type()) {
				field.Set(fieldValue.Convert(field.Type()))
			} else {
				return fmt.Errorf("el tipo de valor para el campo %s no es convertible al tipo del campo", fieldName)
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
		fmt.Println("Esperaba un puntero no nulo.")
		return nil
	}

	val = dereferencePtr(val)

	if val.Kind() != reflect.Slice {
		fmt.Println("Esperaba un slice.")
		return nil
	}

	for i := 0; i < val.Len(); i++ {
		elem := val.Index(i)

		if elem.Kind() != reflect.Struct {
			fmt.Println("Esperaba una estructura en el slice.")
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

	// Desreferenciar punteros si es necesario
	valueA = dereferencePtr(valueA)
	valueB = dereferencePtr(valueB)

	if !valueA.IsValid() {
		valueA = reflect.ValueOf(a)
	}
	if !valueB.IsValid() {
		valueB = reflect.ValueOf(b)
	}
	// Verificar si ambos valores son slices
	if valueA.Kind() == reflect.Slice && valueB.Kind() == reflect.Slice {
		lenA := valueA.Len()
		lenB := valueB.Len()

		// Verificar si los slices tienen la misma longitud
		if lenA != lenB {
			// Si no tienen la misma longitud, devolvemos el slice más grande
			if lenB > lenA {
				return b
			}
			return a
		}

		// Mezclar los elementos de los slices
		resultSlice := reflect.MakeSlice(valueA.Type(), lenA, lenA)
		for i := 0; i < lenA; i++ {
			resultSlice.Index(i).Set(reflect.ValueOf(mergeInterfaces(valueA.Index(i).Interface(), valueB.Index(i).Interface(), false)))
		}
		return resultSlice.Interface()
	}

	if valueA.Kind() != reflect.Struct || valueB.Kind() != reflect.Struct {
		// Si no son slices ni structs, simplemente devolvemos el segundo valor
		return b
	}

	// Mezclar los campos de los structs
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
			if fieldA.Kind() == reflect.Slice && fieldB.Kind() == reflect.Slice && valueA.Field(i).Kind() == reflect.Slice && valueB.Field(i).Kind() == reflect.Slice {
				log.Printf("IF Slice:")
				// Si ambos campos son slices, mezclarlos recursivamente
				if field := replaceUnknownFields(valueA.Field(i)); field != nil {
					log.Printf("IF Slice Replace:")
					resultStruct.Field(i).Set(reflect.ValueOf(field))
				} else {
					log.Printf("Else Slice Replace:")
					resultStruct.Field(i).Set(valueB.Field(i))
				}
			} else if reflect.DeepEqual(fieldA.Interface(), reflect.Zero(fieldA.Type()).Interface()) && !reflect.DeepEqual(fieldB.Interface(), reflect.Zero(fieldB.Type()).Interface()) {
				// Si el campo de a es nulo y el campo de b no lo es, utilizamos el campo de b
				log.Printf("IF DeepEqual:")
				resultStruct.Field(i).Set(valueB.Field(i))
			} else {
				if fmt.Sprint(fieldA.Interface()) == "<unknown>" || fmt.Sprint(fieldA.Interface()) == "<nil>" || fmt.Sprint(fieldA.Interface()) == "<null>" {
					log.Printf("IF unkown:")
					if fmt.Sprint(fieldB.Interface()) != "<unknown>" {
						resultStruct.Field(i).Set(valueB.Field(i))
					}
				} else {
					if valueA.Field(i).Type() != reflect.TypeOf(types.String{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Bool{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Int64{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Float64{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Set{}) {
						// log.Printf("IF Struct:")
						// log.Printf("[DEBUG] Result Struct: %v", valueA.Field(i).Interface())
						// log.Printf("[DEBUG] Kind: %v", valueA.Field(i).Kind())
						// log.Printf("[DEBUG] Kind: %v", valueA.Field(i).Type())
						// log.Printf("[DEBUG] NameA: %v", valueA.Type().Field(i).Name)
						// log.Printf("[DEBUG] NameB: %v", valueB.Type().Field(i).Name)
						mergedValues := changeStructUnknowns(fieldA.Interface(), fieldB.Interface())
						// log.Printf(" Sali 2")
						// log.Printf(" ssss %v", PrintKeyValue(mergedValues))
						fieldValueBPtr := reflect.New(fieldA.Type())
						fieldValueBPtr.Elem().Set(reflect.ValueOf(mergedValues))
						resultStruct.Field(i).Set(fieldValueBPtr)
					} else {
						// log.Printf("ELSE Struct:")
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
			// resultStruct.Field(i).Set(valueB.Field(i))
			// log.Printf("Antes")

			// log.Printf("Despues ZERO")
			for _, path := range pathParams {
				// log.Printf("Despues FOR path  %v", path)
				// log.Printf("Despues FOR valueB.Type().Field(i).Name  %v", valueB.Type().Field(i).Name)
				if valueB.Type().Field(i).Name == path && fieldB.IsZero() && !fieldA.IsZero() && fmt.Sprint(fieldA.Interface()) != "<unknown>" {
					resultStruct.Field(i).Set(fieldA)
				} else {
					// resultStruct.Field(i).Set(fieldB)
					if valueB.Type().Field(i).Name == path && valueB.Type().Field(i).Name != "ID" && fieldB.IsZero() {
						// log.Printf("Despues FOR fieldB.Type().Name() %v", valueB.Type().Field(i).Name)
						// log.Printf("Despues FOR fieldB.Type().Value() %v", fieldB.Interface())
						// log.Printf("Despues FOR fieldB.Type().Value() 22%v", fieldB.IsZero())
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

// func verifyIsSimpleStruct(typeVar string){
// 	for
// }

func changeStructUnknowns(a interface{}, b interface{}) interface{} {
	valueA := reflect.ValueOf(a)
	justA := valueA
	valueB := reflect.ValueOf(b)
	// if valueA.IsZero() {
	// 	valueA = reflect.ValueOf(b)
	// }
	// log.Printf("Entre: ")
	if valueA.Kind() == reflect.Ptr {
		log.Printf("valueA.Kind() == reflect.Ptr")
		valueA = valueA.Elem()
	}
	if valueB.Kind() == reflect.Ptr {
		log.Printf("valueB.Kind() == reflect.Ptr")
		valueB = valueB.Elem()
	}

	if justA.IsZero() {
		log.Printf("Field A = B")
		valueA = valueB
	}

	resultStruct := reflect.New(valueA.Type()).Elem()
	if valueA.Kind() == reflect.Slice && valueB.Kind() == reflect.Slice {
		lenA := valueA.Len()
		lenB := valueB.Len()

		// Verificar si los slices tienen la misma longitud
		if lenA != lenB {
			// Si no tienen la misma longitud, devolvemos el slice más grande
			if lenB > lenA {
				return b
			}
			return a
		}

		// Mezclar los elementos de los slices
		resultSlice := reflect.MakeSlice(valueA.Type(), lenA, lenA)
		for i := 0; i < lenA; i++ {
			resultSlice.Index(i).Set(reflect.ValueOf(changeStructUnknowns(valueA.Index(i).Interface(), valueB.Index(i).Interface())))
		}
		return resultSlice.Interface()
	}
	for i := 0; i < valueA.NumField(); i++ {
		// log.Printf("Entre 2: ")
		fieldValueA := valueA.Field(i)
		// log.Printf("Entre 2: fieldValueA %s", fieldValueA.Interface())
		// log.Printf("Entre 2: fieldValueA %s", fieldValueA.Kind())
		// log.Printf("Entre 2: fieldValueB %v", valueB.IsValid())
		// log.Printf("Entre 2: fieldValueB %v", valueB.Field(i).Kind())

		var fieldValueB reflect.Value

		if valueB.IsValid() {
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

		// Obtener el nombre del campo
		fieldName := valueA.Type().Field(i).Name
		log.Printf("Entre 2: fieldName %s", fieldName)
		if fmt.Sprint(fieldValueA.Interface()) == "<unknown>" {
			if fmt.Sprint(fieldValueB.Interface()) != "<unknown>" {
				log.Printf("fieldValueB %t", fmt.Sprint(fieldValueB.Interface()) == "<unknown>")
				log.Printf("Assigned %v to field %v\n", fieldValueB.Interface(), fieldName)
				resultStruct.Field(i).Set(fieldValueB)
			} else {
				log.Printf("Not Assigned")
			}
		} else {
			if valueA.Field(i).Type() != reflect.TypeOf(types.String{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Bool{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Int64{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Float64{}) && valueA.Field(i).Type() != reflect.TypeOf(types.Set{}) {
				log.Printf("Entre 3 %v to field %v\n", fieldValueB.Interface(), fieldName)
				if !fieldValueA.IsZero() || !fieldValueB.IsZero() {
					nestedResult := changeStructUnknowns(fieldValueA.Interface(), fieldValueB.Interface())

					log.Printf("A pointer: %t", fieldValueA.Kind() != reflect.Ptr)
					log.Printf("B pointer: %t", fieldValueB.Kind() != reflect.Ptr)
					if fieldValueB.Kind() != reflect.Ptr {
						log.Printf("Entre 3 %v to field %v\n", fieldValueB.Interface(), fieldName)
						fieldValueBPtr := reflect.New(fieldValueB.Type())
						fieldValueBPtr.Elem().Set(reflect.ValueOf(nestedResult))
						resultStruct.Field(i).Set(fieldValueBPtr)
					} else {
						if fieldValueA.Kind() != reflect.Ptr {
							log.Printf("Entre 3 %v to field %v\n", fieldValueB.Interface(), fieldName)
							fieldValueAPtr := reflect.New(fieldValueA.Type())
							fieldValueAPtr.Elem().Set(reflect.ValueOf(nestedResult))
							resultStruct.Field(i).Set(fieldValueAPtr)
						}
					}
				} else {
					log.Printf("Both null")
				}
			} else {
				log.Printf("2. Assigned %v to field %v\n", fieldValueA.Interface(), fieldName)

				resultStruct.Field(i).Set(fieldValueA)
			}
		}
	}
	log.Printf("Sali: ")
	// if resultStruct.Kind() != reflect.Ptr {
	// 	log.Print("Diferente de puntero: ", resultStruct.Kind())
	// 	resultStruct = reflect.ValueOf(resultStruct)
	// 	log.Print("Ahora puntero: ", resultStruct.Kind())
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
		if reflect.TypeOf(fieldB).Kind() == reflect.Slice {
			// Obtener la longitud del slice usando reflexión
			length := reflect.ValueOf(fieldB).Elem().Len()
			if length > 0 {
				resultStruct.Field(i).Set(valueA.Field(i))
			}
		} else {
			resultStruct.Field(i).Set(valueB.Field(i))
		}

		log.Printf("Antes")

		for _, path := range pathParams {
			// if valueB.Type().Field(i).Name == "OrganizationID" {
			// 	log.Printf("fieldB.IsZero() %v", fieldB.IsZero())
			// 	log.Printf("valueB.Type().Field(i).Name %v", valueB.Type().Field(i).Name)
			// 	log.Printf("valueB.Type().Field(i).Name %v", !fieldA.IsZero())
			// }

			if valueB.Type().Field(i).Name == path && fieldB.IsZero() && !fieldA.IsZero() && fmt.Sprint(fieldA.Interface()) != "<unknown>" {
				resultStruct.Field(i).Set(fieldA)
			} else {
				if valueB.Type().Field(i).Name == path && valueB.Type().Field(i).Name != "ID" && fieldB.IsZero() {
					// log.Printf("Despues FOR path  %v", path)
					// log.Printf("Despues FOR fieldB.Type().Name() %v", valueB.Type().Field(i).Name)
					// log.Printf("Despues FOR fieldB.Type().Value() %v", fieldB.Interface())
					// log.Printf("Despues FOR fieldB.Type().Value() 22%v", fieldB.IsZero())
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

	// Si v es un puntero, desreferenciarlo
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Asegurarse de que v es una estructura
	if val.Kind() != reflect.Struct {
		return "PrintKeyValue: v no es una estructura"
	}
	var buffer bytes.Buffer

	// Iterar sobre los campos de la estructura
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		// Imprimir el nombre del campo y su valor
		buffer.WriteString(fmt.Sprintf("%s value is %v, ", fieldName, field.Interface()))
	}

	// Eliminar la última coma y espacio
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
