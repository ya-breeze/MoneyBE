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




type CurrencyAggregation struct {

	CurrencyID string `json:"currencyID"`

	Accounts []AccountAggregation `json:"accounts"`
}

// AssertCurrencyAggregationRequired checks if the required fields are not zero-ed
func AssertCurrencyAggregationRequired(obj CurrencyAggregation) error {
	elements := map[string]interface{}{
		"currencyID": obj.CurrencyID,
		"accounts": obj.Accounts,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Accounts {
		if err := AssertAccountAggregationRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCurrencyAggregationConstraints checks if the values respects the defined constraints
func AssertCurrencyAggregationConstraints(obj CurrencyAggregation) error {
	for _, el := range obj.Accounts {
		if err := AssertAccountAggregationConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
