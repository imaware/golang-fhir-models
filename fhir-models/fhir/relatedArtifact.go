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

// RelatedArtifact is documented here http://hl7.org/fhir/StructureDefinition/RelatedArtifact
type RelatedArtifact struct {
	Id        *string             `bson:"id,omitempty" json:"id,omitempty"`
	Extension []Extension         `bson:"extension,omitempty" json:"extension,omitempty"`
	Type      RelatedArtifactType `bson:"type" json:"type"`
	Label     *string             `bson:"label,omitempty" json:"label,omitempty"`
	Display   *string             `bson:"display,omitempty" json:"display,omitempty"`
	Citation  *string             `bson:"citation,omitempty" json:"citation,omitempty"`
	Url       *string             `bson:"url,omitempty" json:"url,omitempty"`
	Document  *Attachment         `bson:"document,omitempty" json:"document,omitempty"`
	Resource  *string             `bson:"resource,omitempty" json:"resource,omitempty"`
}
