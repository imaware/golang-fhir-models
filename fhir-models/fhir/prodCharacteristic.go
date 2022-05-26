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

// THIS FILE IS GENERATED BY https://github.com/imaware/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// ProdCharacteristic is documented here http://hl7.org/fhir/StructureDefinition/ProdCharacteristic
type ProdCharacteristic struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Height            *Quantity        `bson:"height,omitempty" json:"height,omitempty"`
	Width             *Quantity        `bson:"width,omitempty" json:"width,omitempty"`
	Depth             *Quantity        `bson:"depth,omitempty" json:"depth,omitempty"`
	Weight            *Quantity        `bson:"weight,omitempty" json:"weight,omitempty"`
	NominalVolume     *Quantity        `bson:"nominalVolume,omitempty" json:"nominalVolume,omitempty"`
	ExternalDiameter  *Quantity        `bson:"externalDiameter,omitempty" json:"externalDiameter,omitempty"`
	Shape             *string          `bson:"shape,omitempty" json:"shape,omitempty"`
	Color             []string         `bson:"color,omitempty" json:"color,omitempty"`
	Imprint           []string         `bson:"imprint,omitempty" json:"imprint,omitempty"`
	Image             []Attachment     `bson:"image,omitempty" json:"image,omitempty"`
	Scoring           *CodeableConcept `bson:"scoring,omitempty" json:"scoring,omitempty"`
}
