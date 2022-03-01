// Copyright 2019 - 2021 The mharper-imaware Community
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

// THIS FILE IS GENERATED BY https://github.com/mharper-imaware/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ParticipantRequired is documented here http://hl7.org/fhir/ValueSet/participantrequired
type ParticipantRequired int

const (
	ParticipantRequiredRequired ParticipantRequired = iota
	ParticipantRequiredOptional
	ParticipantRequiredInformationOnly
)

func (code ParticipantRequired) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ParticipantRequired) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "required":
		*code = ParticipantRequiredRequired
	case "optional":
		*code = ParticipantRequiredOptional
	case "information-only":
		*code = ParticipantRequiredInformationOnly
	default:
		return fmt.Errorf("unknown ParticipantRequired code `%s`", s)
	}
	return nil
}
func (code ParticipantRequired) String() string {
	return code.Code()
}
func (code ParticipantRequired) Code() string {
	switch code {
	case ParticipantRequiredRequired:
		return "required"
	case ParticipantRequiredOptional:
		return "optional"
	case ParticipantRequiredInformationOnly:
		return "information-only"
	}
	return "<unknown>"
}
func (code ParticipantRequired) Display() string {
	switch code {
	case ParticipantRequiredRequired:
		return "Required"
	case ParticipantRequiredOptional:
		return "Optional"
	case ParticipantRequiredInformationOnly:
		return "Information Only"
	}
	return "<unknown>"
}
func (code ParticipantRequired) Definition() string {
	switch code {
	case ParticipantRequiredRequired:
		return "The participant is required to attend the appointment."
	case ParticipantRequiredOptional:
		return "The participant may optionally attend the appointment."
	case ParticipantRequiredInformationOnly:
		return "The participant is excluded from the appointment, and might not be informed of the appointment taking place. (Appointment is about them, not for them - such as 2 doctors discussing results about a patient's test)."
	}
	return "<unknown>"
}
