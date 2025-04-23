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
	"context"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
)

// trueReflectValue returns the reflect.Value for `in` after derefencing all
// the pointers and unwrapping all the interfaces. It's the concrete value
// beneath it all.
func trueReflectValue(val reflect.Value) reflect.Value {
	kind := val.Type().Kind()
	for kind == reflect.Interface || kind == reflect.Ptr {
		innerVal := val.Elem()
		if !innerVal.IsValid() {
			break
		}
		val = innerVal
		kind = val.Type().Kind()
	}
	return val
}

// commaSeparatedString returns an English joining of the strings in `in`,
// using "and" and commas as appropriate.
func commaSeparatedString(in []string) string {
	switch len(in) {
	case 0:
		return ""
	case 1:
		return in[0]
	case 2:
		return strings.Join(in, " and ")
	default:
		in[len(in)-1] = "and " + in[len(in)-1]
		return strings.Join(in, ", ")
	}
}

// getStructTags returns a map of Terraform field names to their position in
// the tags of the struct `in`. `in` must be a struct.
func getStructTags(_ context.Context, in reflect.Value, path path.Path) (map[string]int, error) {
	tags := map[string]int{}
	typ := trueReflectValue(in).Type()
	if typ.Kind() != reflect.Struct {
		return nil, fmt.Errorf("%s: can't get struct tags of %s, is not a struct", path, in.Type())
	}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.PkgPath != "" {
			// skip unexported fields
			continue
		}
		tag := field.Tag.Get(`tfsdk`)
		if tag == "-" {
			// skip explicitly excluded fields
			continue
		}
		if tag == "" {
			return nil, fmt.Errorf(`%s: need a struct tag for "tfsdk" on %s`, path, field.Name)
		}
		path := path.AtName(tag)
		if !isValidFieldName(tag) {
			return nil, fmt.Errorf("%s: invalid field name, must only use lowercase letters, underscores, and numbers, and must start with a letter", path)
		}
		if other, ok := tags[tag]; ok {
			return nil, fmt.Errorf("%s: can't use field name for both %s and %s", path, typ.Field(other).Name, field.Name)
		}
		tags[tag] = i
	}
	return tags, nil
}

// isValidFieldName returns true if `name` can be used as a field name in a
// Terraform resource or data source.
func isValidFieldName(name string) bool {
	re := regexp.MustCompile("^[a-z][a-z0-9_]*$")
	return re.MatchString(name)
}

// canBeNil returns true if `target`'s type can hold a nil value
func canBeNil(target reflect.Value) bool {
	switch target.Kind() {
	case reflect.Ptr, reflect.Slice, reflect.Map, reflect.Interface:
		// these types can all hold nils
		return true
	default:
		// nothing else can be set to nil
		return false
	}
}
