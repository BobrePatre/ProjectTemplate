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

// checks if the ApiexampleExampleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiexampleExampleResponse{}

// ApiexampleExampleResponse struct for ApiexampleExampleResponse
type ApiexampleExampleResponse struct {
	Title *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Body *string `json:"body,omitempty"`
}

// NewApiexampleExampleResponse instantiates a new ApiexampleExampleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiexampleExampleResponse() *ApiexampleExampleResponse {
	this := ApiexampleExampleResponse{}
	return &this
}

// NewApiexampleExampleResponseWithDefaults instantiates a new ApiexampleExampleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiexampleExampleResponseWithDefaults() *ApiexampleExampleResponse {
	this := ApiexampleExampleResponse{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiexampleExampleResponse) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiexampleExampleResponse) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiexampleExampleResponse) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ApiexampleExampleResponse) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApiexampleExampleResponse) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiexampleExampleResponse) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApiexampleExampleResponse) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApiexampleExampleResponse) SetDescription(v string) {
	o.Description = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ApiexampleExampleResponse) GetBody() string {
	if o == nil || IsNil(o.Body) {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiexampleExampleResponse) GetBodyOk() (*string, bool) {
	if o == nil || IsNil(o.Body) {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ApiexampleExampleResponse) HasBody() bool {
	if o != nil && !IsNil(o.Body) {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *ApiexampleExampleResponse) SetBody(v string) {
	o.Body = &v
}

func (o ApiexampleExampleResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiexampleExampleResponse) ToMap() (map[string]interface{}, error) {
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

type NullableApiexampleExampleResponse struct {
	value *ApiexampleExampleResponse
	isSet bool
}

func (v NullableApiexampleExampleResponse) Get() *ApiexampleExampleResponse {
	return v.value
}

func (v *NullableApiexampleExampleResponse) Set(val *ApiexampleExampleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiexampleExampleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiexampleExampleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiexampleExampleResponse(val *ApiexampleExampleResponse) *NullableApiexampleExampleResponse {
	return &NullableApiexampleExampleResponse{value: val, isSet: true}
}

func (v NullableApiexampleExampleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiexampleExampleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


