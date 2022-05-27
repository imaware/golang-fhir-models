// Copyright 2019 - 2022 The imaware Community
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/imaware/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ActionRequiredBehavior is documented here http://hl7.org/fhir/ValueSet/action-required-behavior
type ActionRequiredBehavior int

const (
	ActionRequiredBehaviorMust ActionRequiredBehavior = iota
	ActionRequiredBehaviorCould
	ActionRequiredBehaviorMustUnlessDocumented
)

func (code ActionRequiredBehavior) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ActionRequiredBehavior) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "must":
		*code = ActionRequiredBehaviorMust
	case "could":
		*code = ActionRequiredBehaviorCould
	case "must-unless-documented":
		*code = ActionRequiredBehaviorMustUnlessDocumented
	default:
		return fmt.Errorf("unknown ActionRequiredBehavior code `%s`", s)
	}
	return nil
}
func (code ActionRequiredBehavior) String() string {
	return code.Code()
}
func (code ActionRequiredBehavior) Code() string {
	switch code {
	case ActionRequiredBehaviorMust:
		return "must"
	case ActionRequiredBehaviorCould:
		return "could"
	case ActionRequiredBehaviorMustUnlessDocumented:
		return "must-unless-documented"
	}
	return "<unknown>"
}
func (code ActionRequiredBehavior) Display() string {
	switch code {
	case ActionRequiredBehaviorMust:
		return "Must"
	case ActionRequiredBehaviorCould:
		return "Could"
	case ActionRequiredBehaviorMustUnlessDocumented:
		return "Must Unless Documented"
	}
	return "<unknown>"
}
func (code ActionRequiredBehavior) Definition() string {
	switch code {
	case ActionRequiredBehaviorMust:
		return "An action with this behavior must be included in the actions processed by the end user; the end user SHALL NOT choose not to include this action."
	case ActionRequiredBehaviorCould:
		return "An action with this behavior may be included in the set of actions processed by the end user."
	case ActionRequiredBehaviorMustUnlessDocumented:
		return "An action with this behavior must be included in the set of actions processed by the end user, unless the end user provides documentation as to why the action was not included."
	}
	return "<unknown>"
}
