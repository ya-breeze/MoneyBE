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
)

// checks if the CheckMatcher200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckMatcher200Response{}

// CheckMatcher200Response struct for CheckMatcher200Response
type CheckMatcher200Response struct {
	Result *bool `json:"result,omitempty"`
}

// NewCheckMatcher200Response instantiates a new CheckMatcher200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckMatcher200Response() *CheckMatcher200Response {
	this := CheckMatcher200Response{}
	return &this
}

// NewCheckMatcher200ResponseWithDefaults instantiates a new CheckMatcher200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckMatcher200ResponseWithDefaults() *CheckMatcher200Response {
	this := CheckMatcher200Response{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *CheckMatcher200Response) GetResult() bool {
	if o == nil || IsNil(o.Result) {
		var ret bool
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckMatcher200Response) GetResultOk() (*bool, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *CheckMatcher200Response) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given bool and assigns it to the Result field.
func (o *CheckMatcher200Response) SetResult(v bool) {
	o.Result = &v
}

func (o CheckMatcher200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckMatcher200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableCheckMatcher200Response struct {
	value *CheckMatcher200Response
	isSet bool
}

func (v NullableCheckMatcher200Response) Get() *CheckMatcher200Response {
	return v.value
}

func (v *NullableCheckMatcher200Response) Set(val *CheckMatcher200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckMatcher200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckMatcher200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckMatcher200Response(val *CheckMatcher200Response) *NullableCheckMatcher200Response {
	return &NullableCheckMatcher200Response{value: val, isSet: true}
}

func (v NullableCheckMatcher200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckMatcher200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


