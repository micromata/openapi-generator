/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"encoding/json"
)

// Foo struct for Foo
type Foo struct {
	Bar *string `json:"bar,omitempty"`
}

// NewFoo instantiates a new Foo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFoo() *Foo {
	this := Foo{}
	var bar string = "bar"
	this.Bar = &bar
	return &this
}

// NewFooWithDefaults instantiates a new Foo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFooWithDefaults() *Foo {
	this := Foo{}
	var bar string = "bar"
	this.Bar = &bar
	return &this
}

// GetBar returns the Bar field value if set, zero value otherwise.
func (o *Foo) GetBar() string {
	if o == nil || o.Bar == nil {
		var ret string
		return ret
	}
	return *o.Bar
}

// GetBarOk returns a tuple with the Bar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Foo) GetBarOk() (*string, bool) {
	if o == nil || o.Bar == nil {
		return nil, false
	}
	return o.Bar, true
}

// HasBar returns a boolean if a field has been set.
func (o *Foo) HasBar() bool {
	if o != nil && o.Bar != nil {
		return true
	}

	return false
}

// SetBar gets a reference to the given string and assigns it to the Bar field.
func (o *Foo) SetBar(v string) {
	o.Bar = &v
}

func (o Foo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bar != nil {
		toSerialize["bar"] = o.Bar
	}
	return json.Marshal(toSerialize)
}

type NullableFoo struct {
	value *Foo
	isSet bool
}

func (v NullableFoo) Get() *Foo {
	return v.value
}

func (v *NullableFoo) Set(val *Foo) {
	v.value = val
	v.isSet = true
}

func (v NullableFoo) IsSet() bool {
	return v.isSet
}

func (v *NullableFoo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFoo(val *Foo) *NullableFoo {
	return &NullableFoo{value: val, isSet: true}
}

func (v NullableFoo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFoo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

