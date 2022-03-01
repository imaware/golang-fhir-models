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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/mharper-imaware/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Basic is documented here http://hl7.org/fhir/StructureDefinition/Basic
type Basic struct {
	Id                *string         `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta           `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string         `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string         `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative      `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension     `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension     `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier    `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Code              CodeableConcept `bson:"code" json:"code"`
	Subject           *Reference      `bson:"subject,omitempty" json:"subject,omitempty"`
	Created           *string         `bson:"created,omitempty" json:"created,omitempty"`
	Author            *Reference      `bson:"author,omitempty" json:"author,omitempty"`
}
type OtherBasic Basic

// MarshalJSON marshals the given Basic as JSON into a byte slice
func (r Basic) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBasic
		ResourceType string `json:"resourceType"`
	}{
		OtherBasic:   OtherBasic(r),
		ResourceType: "Basic",
	})
}

// UnmarshalBasic unmarshals a Basic.
func UnmarshalBasic(b []byte) (Basic, error) {
	var basic Basic
	if err := json.Unmarshal(b, &basic); err != nil {
		return basic, err
	}
	return basic, nil
}
