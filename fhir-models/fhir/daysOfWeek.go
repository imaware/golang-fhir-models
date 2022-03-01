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

// DaysOfWeek is documented here http://hl7.org/fhir/ValueSet/days-of-week
type DaysOfWeek int

const (
	DaysOfWeekMon DaysOfWeek = iota
	DaysOfWeekTue
	DaysOfWeekWed
	DaysOfWeekThu
	DaysOfWeekFri
	DaysOfWeekSat
	DaysOfWeekSun
)

func (code DaysOfWeek) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *DaysOfWeek) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "mon":
		*code = DaysOfWeekMon
	case "tue":
		*code = DaysOfWeekTue
	case "wed":
		*code = DaysOfWeekWed
	case "thu":
		*code = DaysOfWeekThu
	case "fri":
		*code = DaysOfWeekFri
	case "sat":
		*code = DaysOfWeekSat
	case "sun":
		*code = DaysOfWeekSun
	default:
		return fmt.Errorf("unknown DaysOfWeek code `%s`", s)
	}
	return nil
}
func (code DaysOfWeek) String() string {
	return code.Code()
}
func (code DaysOfWeek) Code() string {
	switch code {
	case DaysOfWeekMon:
		return "mon"
	case DaysOfWeekTue:
		return "tue"
	case DaysOfWeekWed:
		return "wed"
	case DaysOfWeekThu:
		return "thu"
	case DaysOfWeekFri:
		return "fri"
	case DaysOfWeekSat:
		return "sat"
	case DaysOfWeekSun:
		return "sun"
	}
	return "<unknown>"
}
func (code DaysOfWeek) Display() string {
	switch code {
	case DaysOfWeekMon:
		return "Monday"
	case DaysOfWeekTue:
		return "Tuesday"
	case DaysOfWeekWed:
		return "Wednesday"
	case DaysOfWeekThu:
		return "Thursday"
	case DaysOfWeekFri:
		return "Friday"
	case DaysOfWeekSat:
		return "Saturday"
	case DaysOfWeekSun:
		return "Sunday"
	}
	return "<unknown>"
}
func (code DaysOfWeek) Definition() string {
	switch code {
	case DaysOfWeekMon:
		return "Monday."
	case DaysOfWeekTue:
		return "Tuesday."
	case DaysOfWeekWed:
		return "Wednesday."
	case DaysOfWeekThu:
		return "Thursday."
	case DaysOfWeekFri:
		return "Friday."
	case DaysOfWeekSat:
		return "Saturday."
	case DaysOfWeekSun:
		return "Sunday."
	}
	return "<unknown>"
}
