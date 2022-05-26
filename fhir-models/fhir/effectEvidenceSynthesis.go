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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/imaware/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// EffectEvidenceSynthesis is documented here http://hl7.org/fhir/StructureDefinition/EffectEvidenceSynthesis
type EffectEvidenceSynthesis struct {
	Id                  *string                                    `bson:"id,omitempty" json:"id,omitempty"`
	Meta                *Meta                                      `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules       *string                                    `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language            *string                                    `bson:"language,omitempty" json:"language,omitempty"`
	Text                *Narrative                                 `bson:"text,omitempty" json:"text,omitempty"`
	Extension           []Extension                                `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension                                `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Url                 *string                                    `bson:"url,omitempty" json:"url,omitempty"`
	Identifier          []Identifier                               `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Version             *string                                    `bson:"version,omitempty" json:"version,omitempty"`
	Name                *string                                    `bson:"name,omitempty" json:"name,omitempty"`
	Title               *string                                    `bson:"title,omitempty" json:"title,omitempty"`
	Status              PublicationStatus                          `bson:"status" json:"status"`
	Date                *string                                    `bson:"date,omitempty" json:"date,omitempty"`
	Publisher           *string                                    `bson:"publisher,omitempty" json:"publisher,omitempty"`
	Contact             []ContactDetail                            `bson:"contact,omitempty" json:"contact,omitempty"`
	Description         *string                                    `bson:"description,omitempty" json:"description,omitempty"`
	Note                []Annotation                               `bson:"note,omitempty" json:"note,omitempty"`
	UseContext          []UsageContext                             `bson:"useContext,omitempty" json:"useContext,omitempty"`
	Jurisdiction        []CodeableConcept                          `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Copyright           *string                                    `bson:"copyright,omitempty" json:"copyright,omitempty"`
	ApprovalDate        *string                                    `bson:"approvalDate,omitempty" json:"approvalDate,omitempty"`
	LastReviewDate      *string                                    `bson:"lastReviewDate,omitempty" json:"lastReviewDate,omitempty"`
	EffectivePeriod     *Period                                    `bson:"effectivePeriod,omitempty" json:"effectivePeriod,omitempty"`
	Topic               []CodeableConcept                          `bson:"topic,omitempty" json:"topic,omitempty"`
	Author              []ContactDetail                            `bson:"author,omitempty" json:"author,omitempty"`
	Editor              []ContactDetail                            `bson:"editor,omitempty" json:"editor,omitempty"`
	Reviewer            []ContactDetail                            `bson:"reviewer,omitempty" json:"reviewer,omitempty"`
	Endorser            []ContactDetail                            `bson:"endorser,omitempty" json:"endorser,omitempty"`
	RelatedArtifact     []RelatedArtifact                          `bson:"relatedArtifact,omitempty" json:"relatedArtifact,omitempty"`
	SynthesisType       *CodeableConcept                           `bson:"synthesisType,omitempty" json:"synthesisType,omitempty"`
	StudyType           *CodeableConcept                           `bson:"studyType,omitempty" json:"studyType,omitempty"`
	Population          Reference                                  `bson:"population" json:"population"`
	Exposure            Reference                                  `bson:"exposure" json:"exposure"`
	ExposureAlternative Reference                                  `bson:"exposureAlternative" json:"exposureAlternative"`
	Outcome             Reference                                  `bson:"outcome" json:"outcome"`
	SampleSize          *EffectEvidenceSynthesisSampleSize         `bson:"sampleSize,omitempty" json:"sampleSize,omitempty"`
	ResultsByExposure   []EffectEvidenceSynthesisResultsByExposure `bson:"resultsByExposure,omitempty" json:"resultsByExposure,omitempty"`
	EffectEstimate      []EffectEvidenceSynthesisEffectEstimate    `bson:"effectEstimate,omitempty" json:"effectEstimate,omitempty"`
	Certainty           []EffectEvidenceSynthesisCertainty         `bson:"certainty,omitempty" json:"certainty,omitempty"`
}
type EffectEvidenceSynthesisSampleSize struct {
	Id                   *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension            []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description          *string     `bson:"description,omitempty" json:"description,omitempty"`
	NumberOfStudies      *int        `bson:"numberOfStudies,omitempty" json:"numberOfStudies,omitempty"`
	NumberOfParticipants *int        `bson:"numberOfParticipants,omitempty" json:"numberOfParticipants,omitempty"`
}
type EffectEvidenceSynthesisResultsByExposure struct {
	Id                    *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description           *string          `bson:"description,omitempty" json:"description,omitempty"`
	ExposureState         *ExposureState   `bson:"exposureState,omitempty" json:"exposureState,omitempty"`
	VariantState          *CodeableConcept `bson:"variantState,omitempty" json:"variantState,omitempty"`
	RiskEvidenceSynthesis Reference        `bson:"riskEvidenceSynthesis" json:"riskEvidenceSynthesis"`
}
type EffectEvidenceSynthesisEffectEstimate struct {
	Id                *string                                                  `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                              `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                              `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Description       *string                                                  `bson:"description,omitempty" json:"description,omitempty"`
	Type              *CodeableConcept                                         `bson:"type,omitempty" json:"type,omitempty"`
	VariantState      *CodeableConcept                                         `bson:"variantState,omitempty" json:"variantState,omitempty"`
	Value             *string                                                  `bson:"value,omitempty" json:"value,omitempty"`
	UnitOfMeasure     *CodeableConcept                                         `bson:"unitOfMeasure,omitempty" json:"unitOfMeasure,omitempty"`
	PrecisionEstimate []EffectEvidenceSynthesisEffectEstimatePrecisionEstimate `bson:"precisionEstimate,omitempty" json:"precisionEstimate,omitempty"`
}
type EffectEvidenceSynthesisEffectEstimatePrecisionEstimate struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Level             *string          `bson:"level,omitempty" json:"level,omitempty"`
	From              *string          `bson:"from,omitempty" json:"from,omitempty"`
	To                *string          `bson:"to,omitempty" json:"to,omitempty"`
}
type EffectEvidenceSynthesisCertainty struct {
	Id                    *string                                                 `bson:"id,omitempty" json:"id,omitempty"`
	Extension             []Extension                                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Rating                []CodeableConcept                                       `bson:"rating,omitempty" json:"rating,omitempty"`
	Note                  []Annotation                                            `bson:"note,omitempty" json:"note,omitempty"`
	CertaintySubcomponent []EffectEvidenceSynthesisCertaintyCertaintySubcomponent `bson:"certaintySubcomponent,omitempty" json:"certaintySubcomponent,omitempty"`
}
type EffectEvidenceSynthesisCertaintyCertaintySubcomponent struct {
	Id                *string           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept  `bson:"type,omitempty" json:"type,omitempty"`
	Rating            []CodeableConcept `bson:"rating,omitempty" json:"rating,omitempty"`
	Note              []Annotation      `bson:"note,omitempty" json:"note,omitempty"`
}
type OtherEffectEvidenceSynthesis EffectEvidenceSynthesis

// MarshalJSON marshals the given EffectEvidenceSynthesis as JSON into a byte slice
func (r EffectEvidenceSynthesis) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherEffectEvidenceSynthesis
		ResourceType string `json:"resourceType"`
	}{
		OtherEffectEvidenceSynthesis: OtherEffectEvidenceSynthesis(r),
		ResourceType:                 "EffectEvidenceSynthesis",
	})
}

// UnmarshalEffectEvidenceSynthesis unmarshals a EffectEvidenceSynthesis.
func UnmarshalEffectEvidenceSynthesis(b []byte) (EffectEvidenceSynthesis, error) {
	var effectEvidenceSynthesis EffectEvidenceSynthesis
	if err := json.Unmarshal(b, &effectEvidenceSynthesis); err != nil {
		return effectEvidenceSynthesis, err
	}
	return effectEvidenceSynthesis, nil
}
