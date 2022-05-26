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

// ExposureState is documented here http://hl7.org/fhir/ValueSet/exposure-state
type ExposureState int

const (
	ExposureStateExposure ExposureState = iota
	ExposureStateExposureAlternative
)

func (code ExposureState) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ExposureState) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "exposure":
		*code = ExposureStateExposure
	case "exposure-alternative":
		*code = ExposureStateExposureAlternative
	default:
		return fmt.Errorf("unknown ExposureState code `%s`", s)
	}
	return nil
}
func (code ExposureState) String() string {
	return code.Code()
}
func (code ExposureState) Code() string {
	switch code {
	case ExposureStateExposure:
		return "exposure"
	case ExposureStateExposureAlternative:
		return "exposure-alternative"
	}
	return "<unknown>"
}
func (code ExposureState) Display() string {
	switch code {
	case ExposureStateExposure:
		return "Exposure"
	case ExposureStateExposureAlternative:
		return "Exposure Alternative"
	}
	return "<unknown>"
}
func (code ExposureState) Definition() string {
	switch code {
	case ExposureStateExposure:
		return "used when the results by exposure is describing the results for the primary exposure of interest."
	case ExposureStateExposureAlternative:
		return "used when the results by exposure is describing the results for the alternative exposure state, control state or comparator state."
	}
	return "<unknown>"
}
