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

// Options provides configuration settings for how the reflection behavior
// works, letting callers tweak different behaviors based on their needs.
type Options struct {
	// UnhandledNullAsEmpty controls whether null values should be
	// translated into empty values without provider interaction, or if
	// they must be explicitly handled.
	UnhandledNullAsEmpty bool

	// UnhandledUnknownAsEmpty controls whether null values should be
	// translated into empty values without provider interaction, or if
	// they must be explicitly handled.
	UnhandledUnknownAsEmpty bool

	// AllowRoundingNumbers silently rounds numbers that don't fit
	// perfectly in the types they're being stored in, rather than
	// returning errors. Numbers will always be rounded towards 0.
	AllowRoundingNumbers bool
}
