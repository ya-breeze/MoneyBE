// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Geek Budget - OpenAPI 3.0
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Contact: ilya.korolev@outlook.com
 */

package goserver


import (
	"time"
)



type TransactionNoId struct {

	Date time.Time `json:"date"`

	Description string `json:"description,omitempty"`

	Place string `json:"place,omitempty"`

	Tags []string `json:"tags,omitempty"`

	PartnerName string `json:"partnerName,omitempty"`

	PartnerAccount string `json:"partnerAccount,omitempty"`

	// Internal bank's ID to be able to match later if necessary
	PartnerInternalID string `json:"partnerInternalID,omitempty"`

	// Stores extra data about transaction. For example could hold \"variable symbol\" to distinguish payment for the same account, but with different meaning
	Extra string `json:"extra,omitempty"`

	// Stores FULL unprocessed transactions which was source of this transaction. Could be used later for detailed analysis
	UnprocessedSources string `json:"unprocessedSources,omitempty"`

	// IDs of unprocessed transaction - to match later
	ExternalIDs string `json:"externalIDs,omitempty"`

	Movements []Movement `json:"movements"`
}

// AssertTransactionNoIdRequired checks if the required fields are not zero-ed
func AssertTransactionNoIdRequired(obj TransactionNoId) error {
	elements := map[string]interface{}{
		"date": obj.Date,
		"movements": obj.Movements,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Movements {
		if err := AssertMovementRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertTransactionNoIdConstraints checks if the values respects the defined constraints
func AssertTransactionNoIdConstraints(obj TransactionNoId) error {
	for _, el := range obj.Movements {
		if err := AssertMovementConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
