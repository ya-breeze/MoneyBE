/*
Geek Budget - OpenAPI 3.0

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: ilya.korolev@outlook.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package goclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MatcherNoID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MatcherNoID{}

// MatcherNoID struct for MatcherNoID
type MatcherNoID struct {
	Name string `json:"name"`
	OutputDescription *string `json:"outputDescription,omitempty"`
	Amount *float64 `json:"amount,omitempty"`
	CurrencyRegExp *string `json:"currencyRegExp,omitempty"`
	PartnerNameRegExp *string `json:"partnerNameRegExp,omitempty"`
	PartnerAccountNumber *string `json:"partnerAccountNumber,omitempty"`
	DescriptionRegExp *string `json:"descriptionRegExp,omitempty"`
	ExtraRegExp *string `json:"extraRegExp,omitempty"`
	OutputMovements []Movement `json:"outputMovements,omitempty"`
}

type _MatcherNoID MatcherNoID

// NewMatcherNoID instantiates a new MatcherNoID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMatcherNoID(name string) *MatcherNoID {
	this := MatcherNoID{}
	this.Name = name
	return &this
}

// NewMatcherNoIDWithDefaults instantiates a new MatcherNoID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatcherNoIDWithDefaults() *MatcherNoID {
	this := MatcherNoID{}
	return &this
}

// GetName returns the Name field value
func (o *MatcherNoID) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MatcherNoID) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MatcherNoID) SetName(v string) {
	o.Name = v
}

// GetOutputDescription returns the OutputDescription field value if set, zero value otherwise.
func (o *MatcherNoID) GetOutputDescription() string {
	if o == nil || IsNil(o.OutputDescription) {
		var ret string
		return ret
	}
	return *o.OutputDescription
}

// GetOutputDescriptionOk returns a tuple with the OutputDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatcherNoID) GetOutputDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.OutputDescription) {
		return nil, false
	}
	return o.OutputDescription, true
}

// HasOutputDescription returns a boolean if a field has been set.
func (o *MatcherNoID) HasOutputDescription() bool {
	if o != nil && !IsNil(o.OutputDescription) {
		return true
	}

	return false
}

// SetOutputDescription gets a reference to the given string and assigns it to the OutputDescription field.
func (o *MatcherNoID) SetOutputDescription(v string) {
	o.OutputDescription = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *MatcherNoID) GetAmount() float64 {
	if o == nil || IsNil(o.Amount) {
		var ret float64
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatcherNoID) GetAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *MatcherNoID) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float64 and assigns it to the Amount field.
func (o *MatcherNoID) SetAmount(v float64) {
	o.Amount = &v
}

// GetCurrencyRegExp returns the CurrencyRegExp field value if set, zero value otherwise.
func (o *MatcherNoID) GetCurrencyRegExp() string {
	if o == nil || IsNil(o.CurrencyRegExp) {
		var ret string
		return ret
	}
	return *o.CurrencyRegExp
}

// GetCurrencyRegExpOk returns a tuple with the CurrencyRegExp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatcherNoID) GetCurrencyRegExpOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyRegExp) {
		return nil, false
	}
	return o.CurrencyRegExp, true
}

// HasCurrencyRegExp returns a boolean if a field has been set.
func (o *MatcherNoID) HasCurrencyRegExp() bool {
	if o != nil && !IsNil(o.CurrencyRegExp) {
		return true
	}

	return false
}

// SetCurrencyRegExp gets a reference to the given string and assigns it to the CurrencyRegExp field.
func (o *MatcherNoID) SetCurrencyRegExp(v string) {
	o.CurrencyRegExp = &v
}

// GetPartnerNameRegExp returns the PartnerNameRegExp field value if set, zero value otherwise.
func (o *MatcherNoID) GetPartnerNameRegExp() string {
	if o == nil || IsNil(o.PartnerNameRegExp) {
		var ret string
		return ret
	}
	return *o.PartnerNameRegExp
}

// GetPartnerNameRegExpOk returns a tuple with the PartnerNameRegExp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatcherNoID) GetPartnerNameRegExpOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerNameRegExp) {
		return nil, false
	}
	return o.PartnerNameRegExp, true
}

// HasPartnerNameRegExp returns a boolean if a field has been set.
func (o *MatcherNoID) HasPartnerNameRegExp() bool {
	if o != nil && !IsNil(o.PartnerNameRegExp) {
		return true
	}

	return false
}

// SetPartnerNameRegExp gets a reference to the given string and assigns it to the PartnerNameRegExp field.
func (o *MatcherNoID) SetPartnerNameRegExp(v string) {
	o.PartnerNameRegExp = &v
}

// GetPartnerAccountNumber returns the PartnerAccountNumber field value if set, zero value otherwise.
func (o *MatcherNoID) GetPartnerAccountNumber() string {
	if o == nil || IsNil(o.PartnerAccountNumber) {
		var ret string
		return ret
	}
	return *o.PartnerAccountNumber
}

// GetPartnerAccountNumberOk returns a tuple with the PartnerAccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatcherNoID) GetPartnerAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerAccountNumber) {
		return nil, false
	}
	return o.PartnerAccountNumber, true
}

// HasPartnerAccountNumber returns a boolean if a field has been set.
func (o *MatcherNoID) HasPartnerAccountNumber() bool {
	if o != nil && !IsNil(o.PartnerAccountNumber) {
		return true
	}

	return false
}

// SetPartnerAccountNumber gets a reference to the given string and assigns it to the PartnerAccountNumber field.
func (o *MatcherNoID) SetPartnerAccountNumber(v string) {
	o.PartnerAccountNumber = &v
}

// GetDescriptionRegExp returns the DescriptionRegExp field value if set, zero value otherwise.
func (o *MatcherNoID) GetDescriptionRegExp() string {
	if o == nil || IsNil(o.DescriptionRegExp) {
		var ret string
		return ret
	}
	return *o.DescriptionRegExp
}

// GetDescriptionRegExpOk returns a tuple with the DescriptionRegExp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatcherNoID) GetDescriptionRegExpOk() (*string, bool) {
	if o == nil || IsNil(o.DescriptionRegExp) {
		return nil, false
	}
	return o.DescriptionRegExp, true
}

// HasDescriptionRegExp returns a boolean if a field has been set.
func (o *MatcherNoID) HasDescriptionRegExp() bool {
	if o != nil && !IsNil(o.DescriptionRegExp) {
		return true
	}

	return false
}

// SetDescriptionRegExp gets a reference to the given string and assigns it to the DescriptionRegExp field.
func (o *MatcherNoID) SetDescriptionRegExp(v string) {
	o.DescriptionRegExp = &v
}

// GetExtraRegExp returns the ExtraRegExp field value if set, zero value otherwise.
func (o *MatcherNoID) GetExtraRegExp() string {
	if o == nil || IsNil(o.ExtraRegExp) {
		var ret string
		return ret
	}
	return *o.ExtraRegExp
}

// GetExtraRegExpOk returns a tuple with the ExtraRegExp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatcherNoID) GetExtraRegExpOk() (*string, bool) {
	if o == nil || IsNil(o.ExtraRegExp) {
		return nil, false
	}
	return o.ExtraRegExp, true
}

// HasExtraRegExp returns a boolean if a field has been set.
func (o *MatcherNoID) HasExtraRegExp() bool {
	if o != nil && !IsNil(o.ExtraRegExp) {
		return true
	}

	return false
}

// SetExtraRegExp gets a reference to the given string and assigns it to the ExtraRegExp field.
func (o *MatcherNoID) SetExtraRegExp(v string) {
	o.ExtraRegExp = &v
}

// GetOutputMovements returns the OutputMovements field value if set, zero value otherwise.
func (o *MatcherNoID) GetOutputMovements() []Movement {
	if o == nil || IsNil(o.OutputMovements) {
		var ret []Movement
		return ret
	}
	return o.OutputMovements
}

// GetOutputMovementsOk returns a tuple with the OutputMovements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatcherNoID) GetOutputMovementsOk() ([]Movement, bool) {
	if o == nil || IsNil(o.OutputMovements) {
		return nil, false
	}
	return o.OutputMovements, true
}

// HasOutputMovements returns a boolean if a field has been set.
func (o *MatcherNoID) HasOutputMovements() bool {
	if o != nil && !IsNil(o.OutputMovements) {
		return true
	}

	return false
}

// SetOutputMovements gets a reference to the given []Movement and assigns it to the OutputMovements field.
func (o *MatcherNoID) SetOutputMovements(v []Movement) {
	o.OutputMovements = v
}

func (o MatcherNoID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatcherNoID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.OutputDescription) {
		toSerialize["outputDescription"] = o.OutputDescription
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.CurrencyRegExp) {
		toSerialize["currencyRegExp"] = o.CurrencyRegExp
	}
	if !IsNil(o.PartnerNameRegExp) {
		toSerialize["partnerNameRegExp"] = o.PartnerNameRegExp
	}
	if !IsNil(o.PartnerAccountNumber) {
		toSerialize["partnerAccountNumber"] = o.PartnerAccountNumber
	}
	if !IsNil(o.DescriptionRegExp) {
		toSerialize["descriptionRegExp"] = o.DescriptionRegExp
	}
	if !IsNil(o.ExtraRegExp) {
		toSerialize["extraRegExp"] = o.ExtraRegExp
	}
	if !IsNil(o.OutputMovements) {
		toSerialize["outputMovements"] = o.OutputMovements
	}
	return toSerialize, nil
}

func (o *MatcherNoID) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varMatcherNoID := _MatcherNoID{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMatcherNoID)

	if err != nil {
		return err
	}

	*o = MatcherNoID(varMatcherNoID)

	return err
}

type NullableMatcherNoID struct {
	value *MatcherNoID
	isSet bool
}

func (v NullableMatcherNoID) Get() *MatcherNoID {
	return v.value
}

func (v *NullableMatcherNoID) Set(val *MatcherNoID) {
	v.value = val
	v.isSet = true
}

func (v NullableMatcherNoID) IsSet() bool {
	return v.isSet
}

func (v *NullableMatcherNoID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatcherNoID(val *MatcherNoID) *NullableMatcherNoID {
	return &NullableMatcherNoID{value: val, isSet: true}
}

func (v NullableMatcherNoID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatcherNoID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


