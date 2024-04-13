/*
Car info

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Car struct for Car
type Car struct {
	RegNum string `json:"regNum"`
	Mark string `json:"mark"`
	Model string `json:"model"`
	Year  *int32 `json:"year,omitempty"`
	Owner People `json:"owner"`
}

// NewCar instantiates a new Car object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCar(regNum string, mark string, model string, owner People) *Car {
	this := Car{}
	this.RegNum = regNum
	this.Mark = mark
	this.Model = model
	this.Owner = owner
	return &this
}

// NewCarWithDefaults instantiates a new Car object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCarWithDefaults() *Car {
	this := Car{}
	return &this
}

// GetRegNum returns the RegNum field value
func (o *Car) GetRegNum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegNum
}

// GetRegNumOk returns a tuple with the RegNum field value
// and a boolean to check if the value has been set.
func (o *Car) GetRegNumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegNum, true
}

// SetRegNum sets field value
func (o *Car) SetRegNum(v string) {
	o.RegNum = v
}

// GetMark returns the Mark field value
func (o *Car) GetMark() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mark
}

// GetMarkOk returns a tuple with the Mark field value
// and a boolean to check if the value has been set.
func (o *Car) GetMarkOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mark, true
}

// SetMark sets field value
func (o *Car) SetMark(v string) {
	o.Mark = v
}

// GetModel returns the Model field value
func (o *Car) GetModel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Model
}

// GetModelOk returns a tuple with the Model field value
// and a boolean to check if the value has been set.
func (o *Car) GetModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Model, true
}

// SetModel sets field value
func (o *Car) SetModel(v string) {
	o.Model = v
}

// GetYear returns the Year field value if set, zero value otherwise.
func (o *Car) GetYear() int32 {
	if o == nil || o.Year == nil {
		var ret int32
		return ret
	}
	return *o.Year
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Car) GetYearOk() (*int32, bool) {
	if o == nil || o.Year == nil {
		return nil, false
	}
	return o.Year, true
}

// HasYear returns a boolean if a field has been set.
func (o *Car) HasYear() bool {
	if o != nil && o.Year != nil {
		return true
	}

	return false
}

// SetYear gets a reference to the given int32 and assigns it to the Year field.
func (o *Car) SetYear(v int32) {
	o.Year = &v
}

// GetOwner returns the Owner field value
func (o *Car) GetOwner() People {
	if o == nil {
		var ret People
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *Car) GetOwnerOk() (*People, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *Car) SetOwner(v People) {
	o.Owner = v
}

func (o Car) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["regNum"] = o.RegNum
	}
	if true {
		toSerialize["mark"] = o.Mark
	}
	if true {
		toSerialize["model"] = o.Model
	}
	if o.Year != nil {
		toSerialize["year"] = o.Year
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	return json.Marshal(toSerialize)
}

type NullableCar struct {
	value *Car
	isSet bool
}

func (v NullableCar) Get() *Car {
	return v.value
}

func (v *NullableCar) Set(val *Car) {
	v.value = val
	v.isSet = true
}

func (v NullableCar) IsSet() bool {
	return v.isSet
}

func (v *NullableCar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCar(val *Car) *NullableCar {
	return &NullableCar{value: val, isSet: true}
}

func (v NullableCar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


