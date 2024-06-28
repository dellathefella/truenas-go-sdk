/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package bluefin

import (
	"encoding/json"
)

// CreateShareSMBParams struct for CreateShareSMBParams
type CreateShareSMBParams struct {
	Purpose *string `json:"purpose,omitempty"`
	Path string `json:"path"`
	PathSuffix *string `json:"path_suffix,omitempty"`
	Home *bool `json:"home,omitempty"`
	Name *string `json:"name,omitempty"`
	Comment *string `json:"comment,omitempty"`
	Ro *bool `json:"ro,omitempty"`
	Browsable *bool `json:"browsable,omitempty"`
	Timemachine *bool `json:"timemachine,omitempty"`
	Recyclebin *bool `json:"recyclebin,omitempty"`
	Guestok *bool `json:"guestok,omitempty"`
	Abe *bool `json:"abe,omitempty"`
	Hostsallow []string `json:"hostsallow,omitempty"`
	Hostsdeny []string `json:"hostsdeny,omitempty"`
	AaplNameMangling *bool `json:"aapl_name_mangling,omitempty"`
	Acl *bool `json:"acl,omitempty"`
	Durablehandle *bool `json:"durablehandle,omitempty"`
	Shadowcopy *bool `json:"shadowcopy,omitempty"`
	Streams *bool `json:"streams,omitempty"`
	Fsrvp *bool `json:"fsrvp,omitempty"`
	Auxsmbconf *string `json:"auxsmbconf,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateShareSMBParams CreateShareSMBParams

// NewCreateShareSMBParams instantiates a new CreateShareSMBParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShareSMBParams(path string) *CreateShareSMBParams {
	this := CreateShareSMBParams{}
	this.Path = path
	return &this
}

// NewCreateShareSMBParamsWithDefaults instantiates a new CreateShareSMBParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShareSMBParamsWithDefaults() *CreateShareSMBParams {
	this := CreateShareSMBParams{}
	return &this
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetPurpose() string {
	if o == nil || o.Purpose == nil {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetPurposeOk() (*string, bool) {
	if o == nil || o.Purpose == nil {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasPurpose() bool {
	if o != nil && o.Purpose != nil {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *CreateShareSMBParams) SetPurpose(v string) {
	o.Purpose = &v
}

// GetPath returns the Path field value
func (o *CreateShareSMBParams) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *CreateShareSMBParams) SetPath(v string) {
	o.Path = v
}

// GetPathSuffix returns the PathSuffix field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetPathSuffix() string {
	if o == nil || o.PathSuffix == nil {
		var ret string
		return ret
	}
	return *o.PathSuffix
}

// GetPathSuffixOk returns a tuple with the PathSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetPathSuffixOk() (*string, bool) {
	if o == nil || o.PathSuffix == nil {
		return nil, false
	}
	return o.PathSuffix, true
}

// HasPathSuffix returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasPathSuffix() bool {
	if o != nil && o.PathSuffix != nil {
		return true
	}

	return false
}

// SetPathSuffix gets a reference to the given string and assigns it to the PathSuffix field.
func (o *CreateShareSMBParams) SetPathSuffix(v string) {
	o.PathSuffix = &v
}

// GetHome returns the Home field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetHome() bool {
	if o == nil || o.Home == nil {
		var ret bool
		return ret
	}
	return *o.Home
}

// GetHomeOk returns a tuple with the Home field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetHomeOk() (*bool, bool) {
	if o == nil || o.Home == nil {
		return nil, false
	}
	return o.Home, true
}

// HasHome returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasHome() bool {
	if o != nil && o.Home != nil {
		return true
	}

	return false
}

// SetHome gets a reference to the given bool and assigns it to the Home field.
func (o *CreateShareSMBParams) SetHome(v bool) {
	o.Home = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateShareSMBParams) SetName(v string) {
	o.Name = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetComment() string {
	if o == nil || o.Comment == nil {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetCommentOk() (*string, bool) {
	if o == nil || o.Comment == nil {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasComment() bool {
	if o != nil && o.Comment != nil {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *CreateShareSMBParams) SetComment(v string) {
	o.Comment = &v
}

// GetRo returns the Ro field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetRo() bool {
	if o == nil || o.Ro == nil {
		var ret bool
		return ret
	}
	return *o.Ro
}

// GetRoOk returns a tuple with the Ro field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetRoOk() (*bool, bool) {
	if o == nil || o.Ro == nil {
		return nil, false
	}
	return o.Ro, true
}

// HasRo returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasRo() bool {
	if o != nil && o.Ro != nil {
		return true
	}

	return false
}

// SetRo gets a reference to the given bool and assigns it to the Ro field.
func (o *CreateShareSMBParams) SetRo(v bool) {
	o.Ro = &v
}

// GetBrowsable returns the Browsable field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetBrowsable() bool {
	if o == nil || o.Browsable == nil {
		var ret bool
		return ret
	}
	return *o.Browsable
}

// GetBrowsableOk returns a tuple with the Browsable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetBrowsableOk() (*bool, bool) {
	if o == nil || o.Browsable == nil {
		return nil, false
	}
	return o.Browsable, true
}

// HasBrowsable returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasBrowsable() bool {
	if o != nil && o.Browsable != nil {
		return true
	}

	return false
}

// SetBrowsable gets a reference to the given bool and assigns it to the Browsable field.
func (o *CreateShareSMBParams) SetBrowsable(v bool) {
	o.Browsable = &v
}

// GetTimemachine returns the Timemachine field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetTimemachine() bool {
	if o == nil || o.Timemachine == nil {
		var ret bool
		return ret
	}
	return *o.Timemachine
}

// GetTimemachineOk returns a tuple with the Timemachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetTimemachineOk() (*bool, bool) {
	if o == nil || o.Timemachine == nil {
		return nil, false
	}
	return o.Timemachine, true
}

// HasTimemachine returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasTimemachine() bool {
	if o != nil && o.Timemachine != nil {
		return true
	}

	return false
}

// SetTimemachine gets a reference to the given bool and assigns it to the Timemachine field.
func (o *CreateShareSMBParams) SetTimemachine(v bool) {
	o.Timemachine = &v
}

// GetRecyclebin returns the Recyclebin field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetRecyclebin() bool {
	if o == nil || o.Recyclebin == nil {
		var ret bool
		return ret
	}
	return *o.Recyclebin
}

// GetRecyclebinOk returns a tuple with the Recyclebin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetRecyclebinOk() (*bool, bool) {
	if o == nil || o.Recyclebin == nil {
		return nil, false
	}
	return o.Recyclebin, true
}

// HasRecyclebin returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasRecyclebin() bool {
	if o != nil && o.Recyclebin != nil {
		return true
	}

	return false
}

// SetRecyclebin gets a reference to the given bool and assigns it to the Recyclebin field.
func (o *CreateShareSMBParams) SetRecyclebin(v bool) {
	o.Recyclebin = &v
}

// GetGuestok returns the Guestok field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetGuestok() bool {
	if o == nil || o.Guestok == nil {
		var ret bool
		return ret
	}
	return *o.Guestok
}

// GetGuestokOk returns a tuple with the Guestok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetGuestokOk() (*bool, bool) {
	if o == nil || o.Guestok == nil {
		return nil, false
	}
	return o.Guestok, true
}

// HasGuestok returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasGuestok() bool {
	if o != nil && o.Guestok != nil {
		return true
	}

	return false
}

// SetGuestok gets a reference to the given bool and assigns it to the Guestok field.
func (o *CreateShareSMBParams) SetGuestok(v bool) {
	o.Guestok = &v
}

// GetAbe returns the Abe field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetAbe() bool {
	if o == nil || o.Abe == nil {
		var ret bool
		return ret
	}
	return *o.Abe
}

// GetAbeOk returns a tuple with the Abe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetAbeOk() (*bool, bool) {
	if o == nil || o.Abe == nil {
		return nil, false
	}
	return o.Abe, true
}

// HasAbe returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasAbe() bool {
	if o != nil && o.Abe != nil {
		return true
	}

	return false
}

// SetAbe gets a reference to the given bool and assigns it to the Abe field.
func (o *CreateShareSMBParams) SetAbe(v bool) {
	o.Abe = &v
}

// GetHostsallow returns the Hostsallow field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetHostsallow() []string {
	if o == nil || o.Hostsallow == nil {
		var ret []string
		return ret
	}
	return o.Hostsallow
}

// GetHostsallowOk returns a tuple with the Hostsallow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetHostsallowOk() ([]string, bool) {
	if o == nil || o.Hostsallow == nil {
		return nil, false
	}
	return o.Hostsallow, true
}

// HasHostsallow returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasHostsallow() bool {
	if o != nil && o.Hostsallow != nil {
		return true
	}

	return false
}

// SetHostsallow gets a reference to the given []string and assigns it to the Hostsallow field.
func (o *CreateShareSMBParams) SetHostsallow(v []string) {
	o.Hostsallow = v
}

// GetHostsdeny returns the Hostsdeny field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetHostsdeny() []string {
	if o == nil || o.Hostsdeny == nil {
		var ret []string
		return ret
	}
	return o.Hostsdeny
}

// GetHostsdenyOk returns a tuple with the Hostsdeny field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetHostsdenyOk() ([]string, bool) {
	if o == nil || o.Hostsdeny == nil {
		return nil, false
	}
	return o.Hostsdeny, true
}

// HasHostsdeny returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasHostsdeny() bool {
	if o != nil && o.Hostsdeny != nil {
		return true
	}

	return false
}

// SetHostsdeny gets a reference to the given []string and assigns it to the Hostsdeny field.
func (o *CreateShareSMBParams) SetHostsdeny(v []string) {
	o.Hostsdeny = v
}

// GetAaplNameMangling returns the AaplNameMangling field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetAaplNameMangling() bool {
	if o == nil || o.AaplNameMangling == nil {
		var ret bool
		return ret
	}
	return *o.AaplNameMangling
}

// GetAaplNameManglingOk returns a tuple with the AaplNameMangling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetAaplNameManglingOk() (*bool, bool) {
	if o == nil || o.AaplNameMangling == nil {
		return nil, false
	}
	return o.AaplNameMangling, true
}

// HasAaplNameMangling returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasAaplNameMangling() bool {
	if o != nil && o.AaplNameMangling != nil {
		return true
	}

	return false
}

// SetAaplNameMangling gets a reference to the given bool and assigns it to the AaplNameMangling field.
func (o *CreateShareSMBParams) SetAaplNameMangling(v bool) {
	o.AaplNameMangling = &v
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetAcl() bool {
	if o == nil || o.Acl == nil {
		var ret bool
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetAclOk() (*bool, bool) {
	if o == nil || o.Acl == nil {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasAcl() bool {
	if o != nil && o.Acl != nil {
		return true
	}

	return false
}

// SetAcl gets a reference to the given bool and assigns it to the Acl field.
func (o *CreateShareSMBParams) SetAcl(v bool) {
	o.Acl = &v
}

// GetDurablehandle returns the Durablehandle field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetDurablehandle() bool {
	if o == nil || o.Durablehandle == nil {
		var ret bool
		return ret
	}
	return *o.Durablehandle
}

// GetDurablehandleOk returns a tuple with the Durablehandle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetDurablehandleOk() (*bool, bool) {
	if o == nil || o.Durablehandle == nil {
		return nil, false
	}
	return o.Durablehandle, true
}

// HasDurablehandle returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasDurablehandle() bool {
	if o != nil && o.Durablehandle != nil {
		return true
	}

	return false
}

// SetDurablehandle gets a reference to the given bool and assigns it to the Durablehandle field.
func (o *CreateShareSMBParams) SetDurablehandle(v bool) {
	o.Durablehandle = &v
}

// GetShadowcopy returns the Shadowcopy field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetShadowcopy() bool {
	if o == nil || o.Shadowcopy == nil {
		var ret bool
		return ret
	}
	return *o.Shadowcopy
}

// GetShadowcopyOk returns a tuple with the Shadowcopy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetShadowcopyOk() (*bool, bool) {
	if o == nil || o.Shadowcopy == nil {
		return nil, false
	}
	return o.Shadowcopy, true
}

// HasShadowcopy returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasShadowcopy() bool {
	if o != nil && o.Shadowcopy != nil {
		return true
	}

	return false
}

// SetShadowcopy gets a reference to the given bool and assigns it to the Shadowcopy field.
func (o *CreateShareSMBParams) SetShadowcopy(v bool) {
	o.Shadowcopy = &v
}

// GetStreams returns the Streams field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetStreams() bool {
	if o == nil || o.Streams == nil {
		var ret bool
		return ret
	}
	return *o.Streams
}

// GetStreamsOk returns a tuple with the Streams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetStreamsOk() (*bool, bool) {
	if o == nil || o.Streams == nil {
		return nil, false
	}
	return o.Streams, true
}

// HasStreams returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasStreams() bool {
	if o != nil && o.Streams != nil {
		return true
	}

	return false
}

// SetStreams gets a reference to the given bool and assigns it to the Streams field.
func (o *CreateShareSMBParams) SetStreams(v bool) {
	o.Streams = &v
}

// GetFsrvp returns the Fsrvp field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetFsrvp() bool {
	if o == nil || o.Fsrvp == nil {
		var ret bool
		return ret
	}
	return *o.Fsrvp
}

// GetFsrvpOk returns a tuple with the Fsrvp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetFsrvpOk() (*bool, bool) {
	if o == nil || o.Fsrvp == nil {
		return nil, false
	}
	return o.Fsrvp, true
}

// HasFsrvp returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasFsrvp() bool {
	if o != nil && o.Fsrvp != nil {
		return true
	}

	return false
}

// SetFsrvp gets a reference to the given bool and assigns it to the Fsrvp field.
func (o *CreateShareSMBParams) SetFsrvp(v bool) {
	o.Fsrvp = &v
}

// GetAuxsmbconf returns the Auxsmbconf field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetAuxsmbconf() string {
	if o == nil || o.Auxsmbconf == nil {
		var ret string
		return ret
	}
	return *o.Auxsmbconf
}

// GetAuxsmbconfOk returns a tuple with the Auxsmbconf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetAuxsmbconfOk() (*string, bool) {
	if o == nil || o.Auxsmbconf == nil {
		return nil, false
	}
	return o.Auxsmbconf, true
}

// HasAuxsmbconf returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasAuxsmbconf() bool {
	if o != nil && o.Auxsmbconf != nil {
		return true
	}

	return false
}

// SetAuxsmbconf gets a reference to the given string and assigns it to the Auxsmbconf field.
func (o *CreateShareSMBParams) SetAuxsmbconf(v string) {
	o.Auxsmbconf = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *CreateShareSMBParams) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShareSMBParams) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *CreateShareSMBParams) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *CreateShareSMBParams) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o CreateShareSMBParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Purpose != nil {
		toSerialize["purpose"] = o.Purpose
	}
	if true {
		toSerialize["path"] = o.Path
	}
	if o.PathSuffix != nil {
		toSerialize["path_suffix"] = o.PathSuffix
	}
	if o.Home != nil {
		toSerialize["home"] = o.Home
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Comment != nil {
		toSerialize["comment"] = o.Comment
	}
	if o.Ro != nil {
		toSerialize["ro"] = o.Ro
	}
	if o.Browsable != nil {
		toSerialize["browsable"] = o.Browsable
	}
	if o.Timemachine != nil {
		toSerialize["timemachine"] = o.Timemachine
	}
	if o.Recyclebin != nil {
		toSerialize["recyclebin"] = o.Recyclebin
	}
	if o.Guestok != nil {
		toSerialize["guestok"] = o.Guestok
	}
	if o.Abe != nil {
		toSerialize["abe"] = o.Abe
	}
	if o.Hostsallow != nil {
		toSerialize["hostsallow"] = o.Hostsallow
	}
	if o.Hostsdeny != nil {
		toSerialize["hostsdeny"] = o.Hostsdeny
	}
	if o.AaplNameMangling != nil {
		toSerialize["aapl_name_mangling"] = o.AaplNameMangling
	}
	if o.Acl != nil {
		toSerialize["acl"] = o.Acl
	}
	if o.Durablehandle != nil {
		toSerialize["durablehandle"] = o.Durablehandle
	}
	if o.Shadowcopy != nil {
		toSerialize["shadowcopy"] = o.Shadowcopy
	}
	if o.Streams != nil {
		toSerialize["streams"] = o.Streams
	}
	if o.Fsrvp != nil {
		toSerialize["fsrvp"] = o.Fsrvp
	}
	if o.Auxsmbconf != nil {
		toSerialize["auxsmbconf"] = o.Auxsmbconf
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateShareSMBParams) UnmarshalJSON(bytes []byte) (err error) {
	varCreateShareSMBParams := _CreateShareSMBParams{}

	if err = json.Unmarshal(bytes, &varCreateShareSMBParams); err == nil {
		*o = CreateShareSMBParams(varCreateShareSMBParams)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "purpose")
		delete(additionalProperties, "path")
		delete(additionalProperties, "path_suffix")
		delete(additionalProperties, "home")
		delete(additionalProperties, "name")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "ro")
		delete(additionalProperties, "browsable")
		delete(additionalProperties, "timemachine")
		delete(additionalProperties, "recyclebin")
		delete(additionalProperties, "guestok")
		delete(additionalProperties, "abe")
		delete(additionalProperties, "hostsallow")
		delete(additionalProperties, "hostsdeny")
		delete(additionalProperties, "aapl_name_mangling")
		delete(additionalProperties, "acl")
		delete(additionalProperties, "durablehandle")
		delete(additionalProperties, "shadowcopy")
		delete(additionalProperties, "streams")
		delete(additionalProperties, "fsrvp")
		delete(additionalProperties, "auxsmbconf")
		delete(additionalProperties, "enabled")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateShareSMBParams struct {
	value *CreateShareSMBParams
	isSet bool
}

func (v NullableCreateShareSMBParams) Get() *CreateShareSMBParams {
	return v.value
}

func (v *NullableCreateShareSMBParams) Set(val *CreateShareSMBParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShareSMBParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShareSMBParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShareSMBParams(val *CreateShareSMBParams) *NullableCreateShareSMBParams {
	return &NullableCreateShareSMBParams{value: val, isSet: true}
}

func (v NullableCreateShareSMBParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShareSMBParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


