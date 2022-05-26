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

// EndpointStatus is documented here http://hl7.org/fhir/ValueSet/endpoint-status
type EndpointStatus int

const (
	EndpointStatusActive EndpointStatus = iota
	EndpointStatusSuspended
	EndpointStatusError
	EndpointStatusOff
	EndpointStatusEnteredInError
	EndpointStatusTest
)

func (code EndpointStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *EndpointStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = EndpointStatusActive
	case "suspended":
		*code = EndpointStatusSuspended
	case "error":
		*code = EndpointStatusError
	case "off":
		*code = EndpointStatusOff
	case "entered-in-error":
		*code = EndpointStatusEnteredInError
	case "test":
		*code = EndpointStatusTest
	default:
		return fmt.Errorf("unknown EndpointStatus code `%s`", s)
	}
	return nil
}
func (code EndpointStatus) String() string {
	return code.Code()
}
func (code EndpointStatus) Code() string {
	switch code {
	case EndpointStatusActive:
		return "active"
	case EndpointStatusSuspended:
		return "suspended"
	case EndpointStatusError:
		return "error"
	case EndpointStatusOff:
		return "off"
	case EndpointStatusEnteredInError:
		return "entered-in-error"
	case EndpointStatusTest:
		return "test"
	}
	return "<unknown>"
}
func (code EndpointStatus) Display() string {
	switch code {
	case EndpointStatusActive:
		return "Active"
	case EndpointStatusSuspended:
		return "Suspended"
	case EndpointStatusError:
		return "Error"
	case EndpointStatusOff:
		return "Off"
	case EndpointStatusEnteredInError:
		return "Entered in error"
	case EndpointStatusTest:
		return "Test"
	}
	return "<unknown>"
}
func (code EndpointStatus) Definition() string {
	switch code {
	case EndpointStatusActive:
		return "This endpoint is expected to be active and can be used."
	case EndpointStatusSuspended:
		return "This endpoint is temporarily unavailable."
	case EndpointStatusError:
		return "This endpoint has exceeded connectivity thresholds and is considered in an error state and should no longer be attempted to connect to until corrective action is taken."
	case EndpointStatusOff:
		return "This endpoint is no longer to be used."
	case EndpointStatusEnteredInError:
		return "This instance should not have been part of this patient's medical record."
	case EndpointStatusTest:
		return "This endpoint is not intended for production usage."
	}
	return "<unknown>"
}
