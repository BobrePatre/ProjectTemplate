/*
Example API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ApiexampleExampleRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiexampleExampleRequest{}

// ApiexampleExampleRequest struct for ApiexampleExampleRequest
type ApiexampleExampleRequest struct {
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Body *string `json:"body,omitempty"`
}

// NewApiexampleExampleRequest instantiates a new ApiexampleExampleRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiexampleExampleRequest() *ApiexampleExampleRequest {
	this := ApiexampleExampleRequest{}
	return &this
}

// NewApiexampleExampleRequestWithDefaults instantiates a new ApiexampleExampleRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiexampleExampleRequestWithDefaults() *ApiexampleExampleRequest {
	this := ApiexampleExampleRequest{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiexampleExampleRequest) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiexampleExampleRequest) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiexampleExampleRequest) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiexampleExampleRequest) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiexampleExampleRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiexampleExampleRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiexampleExampleRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiexampleExampleRequest) SetDescription(v string) {
	o.Description = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ApiexampleExampleRequest) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiexampleExampleRequest) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ApiexampleExampleRequest) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *ApiexampleExampleRequest) SetBody(v string) {
	o.Body = &v
}

func (o ApiexampleExampleRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiexampleExampleRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Body) {
		toSerialize["body"] = o.Body
	}
	return toSerialize, nil
}

type NullableApiexampleExampleRequest struct {
	value *ApiexampleExampleRequest
	isSet bool
}

func (v NullableApiexampleExampleRequest) Get() *ApiexampleExampleRequest {
	return v.value
}

func (v *NullableApiexampleExampleRequest) Set(val *ApiexampleExampleRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiexampleExampleRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiexampleExampleRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiexampleExampleRequest(val *ApiexampleExampleRequest) *NullableApiexampleExampleRequest {
	return &NullableApiexampleExampleRequest{value: val, isSet: true}
}

func (v NullableApiexampleExampleRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiexampleExampleRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


