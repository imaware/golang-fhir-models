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

// MedicinalProductPackaged is documented here http://hl7.org/fhir/StructureDefinition/MedicinalProductPackaged
type MedicinalProductPackaged struct {
	Id                     *string                                   `bson:"id,omitempty" json:"id,omitempty"`
	Meta                   *Meta                                     `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules          *string                                   `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language               *string                                   `bson:"language,omitempty" json:"language,omitempty"`
	Text                   *Narrative                                `bson:"text,omitempty" json:"text,omitempty"`
	Extension              []Extension                               `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension      []Extension                               `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier             []Identifier                              `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Subject                []Reference                               `bson:"subject,omitempty" json:"subject,omitempty"`
	Description            *string                                   `bson:"description,omitempty" json:"description,omitempty"`
	LegalStatusOfSupply    *CodeableConcept                          `bson:"legalStatusOfSupply,omitempty" json:"legalStatusOfSupply,omitempty"`
	MarketingStatus        []MarketingStatus                         `bson:"marketingStatus,omitempty" json:"marketingStatus,omitempty"`
	MarketingAuthorization *Reference                                `bson:"marketingAuthorization,omitempty" json:"marketingAuthorization,omitempty"`
	Manufacturer           []Reference                               `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
	BatchIdentifier        []MedicinalProductPackagedBatchIdentifier `bson:"batchIdentifier,omitempty" json:"batchIdentifier,omitempty"`
	PackageItem            []MedicinalProductPackagedPackageItem     `bson:"packageItem" json:"packageItem"`
}
type MedicinalProductPackagedBatchIdentifier struct {
	Id                 *string     `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	OuterPackaging     Identifier  `bson:"outerPackaging" json:"outerPackaging"`
	ImmediatePackaging *Identifier `bson:"immediatePackaging,omitempty" json:"immediatePackaging,omitempty"`
}
type MedicinalProductPackagedPackageItem struct {
	Id                      *string                               `bson:"id,omitempty" json:"id,omitempty"`
	Extension               []Extension                           `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension       []Extension                           `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier              []Identifier                          `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                    CodeableConcept                       `bson:"type" json:"type"`
	Quantity                Quantity                              `bson:"quantity" json:"quantity"`
	Material                []CodeableConcept                     `bson:"material,omitempty" json:"material,omitempty"`
	AlternateMaterial       []CodeableConcept                     `bson:"alternateMaterial,omitempty" json:"alternateMaterial,omitempty"`
	Device                  []Reference                           `bson:"device,omitempty" json:"device,omitempty"`
	ManufacturedItem        []Reference                           `bson:"manufacturedItem,omitempty" json:"manufacturedItem,omitempty"`
	PackageItem             []MedicinalProductPackagedPackageItem `bson:"packageItem,omitempty" json:"packageItem,omitempty"`
	PhysicalCharacteristics *ProdCharacteristic                   `bson:"physicalCharacteristics,omitempty" json:"physicalCharacteristics,omitempty"`
	OtherCharacteristics    []CodeableConcept                     `bson:"otherCharacteristics,omitempty" json:"otherCharacteristics,omitempty"`
	ShelfLifeStorage        []ProductShelfLife                    `bson:"shelfLifeStorage,omitempty" json:"shelfLifeStorage,omitempty"`
	Manufacturer            []Reference                           `bson:"manufacturer,omitempty" json:"manufacturer,omitempty"`
}
type OtherMedicinalProductPackaged MedicinalProductPackaged

// MarshalJSON marshals the given MedicinalProductPackaged as JSON into a byte slice
func (r MedicinalProductPackaged) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherMedicinalProductPackaged
		ResourceType string `json:"resourceType"`
	}{
		OtherMedicinalProductPackaged: OtherMedicinalProductPackaged(r),
		ResourceType:                  "MedicinalProductPackaged",
	})
}

// UnmarshalMedicinalProductPackaged unmarshals a MedicinalProductPackaged.
func UnmarshalMedicinalProductPackaged(b []byte) (MedicinalProductPackaged, error) {
	var medicinalProductPackaged MedicinalProductPackaged
	if err := json.Unmarshal(b, &medicinalProductPackaged); err != nil {
		return medicinalProductPackaged, err
	}
	return medicinalProductPackaged, nil
}
