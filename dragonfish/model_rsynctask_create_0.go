/*
TrueNAS RESTful API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dragonfish

import (
	"encoding/json"
)

// RsynctaskCreate0 struct for RsynctaskCreate0
type RsynctaskCreate0 struct {
	Path          *string                   `json:"path,omitempty"`
	User          *string                   `json:"user,omitempty"`
	Remotehost    *string                   `json:"remotehost,omitempty"`
	Remoteport    *int32                    `json:"remoteport,omitempty"`
	Mode          *string                   `json:"mode,omitempty"`
	Remotemodule  *string                   `json:"remotemodule,omitempty"`
	Remotepath    *string                   `json:"remotepath,omitempty"`
	ValidateRpath *bool                     `json:"validate_rpath,omitempty"`
	Direction     *string                   `json:"direction,omitempty"`
	Desc          *string                   `json:"desc,omitempty"`
	Schedule      *CloudsyncCreate0Schedule `json:"schedule,omitempty"`
	Recursive     *bool                     `json:"recursive,omitempty"`
	Times         *bool                     `json:"times,omitempty"`
	Compress      *bool                     `json:"compress,omitempty"`
	Archive       *bool                     `json:"archive,omitempty"`
	Delete        *bool                     `json:"delete,omitempty"`
	Quiet         *bool                     `json:"quiet,omitempty"`
	Preserveperm  *bool                     `json:"preserveperm,omitempty"`
	Preserveattr  *bool                     `json:"preserveattr,omitempty"`
	Delayupdates  *bool                     `json:"delayupdates,omitempty"`
	Extra         []string                  `json:"extra,omitempty"`
	Enabled       *bool                     `json:"enabled,omitempty"`
}

// NewRsynctaskCreate0 instantiates a new RsynctaskCreate0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRsynctaskCreate0() *RsynctaskCreate0 {
	this := RsynctaskCreate0{}
	return &this
}

// NewRsynctaskCreate0WithDefaults instantiates a new RsynctaskCreate0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRsynctaskCreate0WithDefaults() *RsynctaskCreate0 {
	this := RsynctaskCreate0{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *RsynctaskCreate0) SetPath(v string) {
	o.Path = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *RsynctaskCreate0) SetUser(v string) {
	o.User = &v
}

// GetRemotehost returns the Remotehost field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetRemotehost() string {
	if o == nil || o.Remotehost == nil {
		var ret string
		return ret
	}
	return *o.Remotehost
}

// GetRemotehostOk returns a tuple with the Remotehost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetRemotehostOk() (*string, bool) {
	if o == nil || o.Remotehost == nil {
		return nil, false
	}
	return o.Remotehost, true
}

// HasRemotehost returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasRemotehost() bool {
	if o != nil && o.Remotehost != nil {
		return true
	}

	return false
}

// SetRemotehost gets a reference to the given string and assigns it to the Remotehost field.
func (o *RsynctaskCreate0) SetRemotehost(v string) {
	o.Remotehost = &v
}

// GetRemoteport returns the Remoteport field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetRemoteport() int32 {
	if o == nil || o.Remoteport == nil {
		var ret int32
		return ret
	}
	return *o.Remoteport
}

// GetRemoteportOk returns a tuple with the Remoteport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetRemoteportOk() (*int32, bool) {
	if o == nil || o.Remoteport == nil {
		return nil, false
	}
	return o.Remoteport, true
}

// HasRemoteport returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasRemoteport() bool {
	if o != nil && o.Remoteport != nil {
		return true
	}

	return false
}

// SetRemoteport gets a reference to the given int32 and assigns it to the Remoteport field.
func (o *RsynctaskCreate0) SetRemoteport(v int32) {
	o.Remoteport = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetMode() string {
	if o == nil || o.Mode == nil {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetModeOk() (*string, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *RsynctaskCreate0) SetMode(v string) {
	o.Mode = &v
}

// GetRemotemodule returns the Remotemodule field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetRemotemodule() string {
	if o == nil || o.Remotemodule == nil {
		var ret string
		return ret
	}
	return *o.Remotemodule
}

// GetRemotemoduleOk returns a tuple with the Remotemodule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetRemotemoduleOk() (*string, bool) {
	if o == nil || o.Remotemodule == nil {
		return nil, false
	}
	return o.Remotemodule, true
}

// HasRemotemodule returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasRemotemodule() bool {
	if o != nil && o.Remotemodule != nil {
		return true
	}

	return false
}

// SetRemotemodule gets a reference to the given string and assigns it to the Remotemodule field.
func (o *RsynctaskCreate0) SetRemotemodule(v string) {
	o.Remotemodule = &v
}

// GetRemotepath returns the Remotepath field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetRemotepath() string {
	if o == nil || o.Remotepath == nil {
		var ret string
		return ret
	}
	return *o.Remotepath
}

// GetRemotepathOk returns a tuple with the Remotepath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetRemotepathOk() (*string, bool) {
	if o == nil || o.Remotepath == nil {
		return nil, false
	}
	return o.Remotepath, true
}

// HasRemotepath returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasRemotepath() bool {
	if o != nil && o.Remotepath != nil {
		return true
	}

	return false
}

// SetRemotepath gets a reference to the given string and assigns it to the Remotepath field.
func (o *RsynctaskCreate0) SetRemotepath(v string) {
	o.Remotepath = &v
}

// GetValidateRpath returns the ValidateRpath field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetValidateRpath() bool {
	if o == nil || o.ValidateRpath == nil {
		var ret bool
		return ret
	}
	return *o.ValidateRpath
}

// GetValidateRpathOk returns a tuple with the ValidateRpath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetValidateRpathOk() (*bool, bool) {
	if o == nil || o.ValidateRpath == nil {
		return nil, false
	}
	return o.ValidateRpath, true
}

// HasValidateRpath returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasValidateRpath() bool {
	if o != nil && o.ValidateRpath != nil {
		return true
	}

	return false
}

// SetValidateRpath gets a reference to the given bool and assigns it to the ValidateRpath field.
func (o *RsynctaskCreate0) SetValidateRpath(v bool) {
	o.ValidateRpath = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetDirection() string {
	if o == nil || o.Direction == nil {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetDirectionOk() (*string, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *RsynctaskCreate0) SetDirection(v string) {
	o.Direction = &v
}

// GetDesc returns the Desc field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetDesc() string {
	if o == nil || o.Desc == nil {
		var ret string
		return ret
	}
	return *o.Desc
}

// GetDescOk returns a tuple with the Desc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetDescOk() (*string, bool) {
	if o == nil || o.Desc == nil {
		return nil, false
	}
	return o.Desc, true
}

// HasDesc returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasDesc() bool {
	if o != nil && o.Desc != nil {
		return true
	}

	return false
}

// SetDesc gets a reference to the given string and assigns it to the Desc field.
func (o *RsynctaskCreate0) SetDesc(v string) {
	o.Desc = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetSchedule() CloudsyncCreate0Schedule {
	if o == nil || o.Schedule == nil {
		var ret CloudsyncCreate0Schedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetScheduleOk() (*CloudsyncCreate0Schedule, bool) {
	if o == nil || o.Schedule == nil {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasSchedule() bool {
	if o != nil && o.Schedule != nil {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given CloudsyncCreate0Schedule and assigns it to the Schedule field.
func (o *RsynctaskCreate0) SetSchedule(v CloudsyncCreate0Schedule) {
	o.Schedule = &v
}

// GetRecursive returns the Recursive field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetRecursive() bool {
	if o == nil || o.Recursive == nil {
		var ret bool
		return ret
	}
	return *o.Recursive
}

// GetRecursiveOk returns a tuple with the Recursive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetRecursiveOk() (*bool, bool) {
	if o == nil || o.Recursive == nil {
		return nil, false
	}
	return o.Recursive, true
}

// HasRecursive returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasRecursive() bool {
	if o != nil && o.Recursive != nil {
		return true
	}

	return false
}

// SetRecursive gets a reference to the given bool and assigns it to the Recursive field.
func (o *RsynctaskCreate0) SetRecursive(v bool) {
	o.Recursive = &v
}

// GetTimes returns the Times field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetTimes() bool {
	if o == nil || o.Times == nil {
		var ret bool
		return ret
	}
	return *o.Times
}

// GetTimesOk returns a tuple with the Times field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetTimesOk() (*bool, bool) {
	if o == nil || o.Times == nil {
		return nil, false
	}
	return o.Times, true
}

// HasTimes returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasTimes() bool {
	if o != nil && o.Times != nil {
		return true
	}

	return false
}

// SetTimes gets a reference to the given bool and assigns it to the Times field.
func (o *RsynctaskCreate0) SetTimes(v bool) {
	o.Times = &v
}

// GetCompress returns the Compress field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetCompress() bool {
	if o == nil || o.Compress == nil {
		var ret bool
		return ret
	}
	return *o.Compress
}

// GetCompressOk returns a tuple with the Compress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetCompressOk() (*bool, bool) {
	if o == nil || o.Compress == nil {
		return nil, false
	}
	return o.Compress, true
}

// HasCompress returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasCompress() bool {
	if o != nil && o.Compress != nil {
		return true
	}

	return false
}

// SetCompress gets a reference to the given bool and assigns it to the Compress field.
func (o *RsynctaskCreate0) SetCompress(v bool) {
	o.Compress = &v
}

// GetArchive returns the Archive field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetArchive() bool {
	if o == nil || o.Archive == nil {
		var ret bool
		return ret
	}
	return *o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetArchiveOk() (*bool, bool) {
	if o == nil || o.Archive == nil {
		return nil, false
	}
	return o.Archive, true
}

// HasArchive returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasArchive() bool {
	if o != nil && o.Archive != nil {
		return true
	}

	return false
}

// SetArchive gets a reference to the given bool and assigns it to the Archive field.
func (o *RsynctaskCreate0) SetArchive(v bool) {
	o.Archive = &v
}

// GetDelete returns the Delete field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetDelete() bool {
	if o == nil || o.Delete == nil {
		var ret bool
		return ret
	}
	return *o.Delete
}

// GetDeleteOk returns a tuple with the Delete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetDeleteOk() (*bool, bool) {
	if o == nil || o.Delete == nil {
		return nil, false
	}
	return o.Delete, true
}

// HasDelete returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasDelete() bool {
	if o != nil && o.Delete != nil {
		return true
	}

	return false
}

// SetDelete gets a reference to the given bool and assigns it to the Delete field.
func (o *RsynctaskCreate0) SetDelete(v bool) {
	o.Delete = &v
}

// GetQuiet returns the Quiet field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetQuiet() bool {
	if o == nil || o.Quiet == nil {
		var ret bool
		return ret
	}
	return *o.Quiet
}

// GetQuietOk returns a tuple with the Quiet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetQuietOk() (*bool, bool) {
	if o == nil || o.Quiet == nil {
		return nil, false
	}
	return o.Quiet, true
}

// HasQuiet returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasQuiet() bool {
	if o != nil && o.Quiet != nil {
		return true
	}

	return false
}

// SetQuiet gets a reference to the given bool and assigns it to the Quiet field.
func (o *RsynctaskCreate0) SetQuiet(v bool) {
	o.Quiet = &v
}

// GetPreserveperm returns the Preserveperm field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetPreserveperm() bool {
	if o == nil || o.Preserveperm == nil {
		var ret bool
		return ret
	}
	return *o.Preserveperm
}

// GetPreservepermOk returns a tuple with the Preserveperm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetPreservepermOk() (*bool, bool) {
	if o == nil || o.Preserveperm == nil {
		return nil, false
	}
	return o.Preserveperm, true
}

// HasPreserveperm returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasPreserveperm() bool {
	if o != nil && o.Preserveperm != nil {
		return true
	}

	return false
}

// SetPreserveperm gets a reference to the given bool and assigns it to the Preserveperm field.
func (o *RsynctaskCreate0) SetPreserveperm(v bool) {
	o.Preserveperm = &v
}

// GetPreserveattr returns the Preserveattr field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetPreserveattr() bool {
	if o == nil || o.Preserveattr == nil {
		var ret bool
		return ret
	}
	return *o.Preserveattr
}

// GetPreserveattrOk returns a tuple with the Preserveattr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetPreserveattrOk() (*bool, bool) {
	if o == nil || o.Preserveattr == nil {
		return nil, false
	}
	return o.Preserveattr, true
}

// HasPreserveattr returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasPreserveattr() bool {
	if o != nil && o.Preserveattr != nil {
		return true
	}

	return false
}

// SetPreserveattr gets a reference to the given bool and assigns it to the Preserveattr field.
func (o *RsynctaskCreate0) SetPreserveattr(v bool) {
	o.Preserveattr = &v
}

// GetDelayupdates returns the Delayupdates field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetDelayupdates() bool {
	if o == nil || o.Delayupdates == nil {
		var ret bool
		return ret
	}
	return *o.Delayupdates
}

// GetDelayupdatesOk returns a tuple with the Delayupdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetDelayupdatesOk() (*bool, bool) {
	if o == nil || o.Delayupdates == nil {
		return nil, false
	}
	return o.Delayupdates, true
}

// HasDelayupdates returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasDelayupdates() bool {
	if o != nil && o.Delayupdates != nil {
		return true
	}

	return false
}

// SetDelayupdates gets a reference to the given bool and assigns it to the Delayupdates field.
func (o *RsynctaskCreate0) SetDelayupdates(v bool) {
	o.Delayupdates = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetExtra() []string {
	if o == nil || o.Extra == nil {
		var ret []string
		return ret
	}
	return o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetExtraOk() ([]string, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given []string and assigns it to the Extra field.
func (o *RsynctaskCreate0) SetExtra(v []string) {
	o.Extra = v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RsynctaskCreate0) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RsynctaskCreate0) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RsynctaskCreate0) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RsynctaskCreate0) SetEnabled(v bool) {
	o.Enabled = &v
}

func (o RsynctaskCreate0) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	if o.Remotehost != nil {
		toSerialize["remotehost"] = o.Remotehost
	}
	if o.Remoteport != nil {
		toSerialize["remoteport"] = o.Remoteport
	}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.Remotemodule != nil {
		toSerialize["remotemodule"] = o.Remotemodule
	}
	if o.Remotepath != nil {
		toSerialize["remotepath"] = o.Remotepath
	}
	if o.ValidateRpath != nil {
		toSerialize["validate_rpath"] = o.ValidateRpath
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.Desc != nil {
		toSerialize["desc"] = o.Desc
	}
	if o.Schedule != nil {
		toSerialize["schedule"] = o.Schedule
	}
	if o.Recursive != nil {
		toSerialize["recursive"] = o.Recursive
	}
	if o.Times != nil {
		toSerialize["times"] = o.Times
	}
	if o.Compress != nil {
		toSerialize["compress"] = o.Compress
	}
	if o.Archive != nil {
		toSerialize["archive"] = o.Archive
	}
	if o.Delete != nil {
		toSerialize["delete"] = o.Delete
	}
	if o.Quiet != nil {
		toSerialize["quiet"] = o.Quiet
	}
	if o.Preserveperm != nil {
		toSerialize["preserveperm"] = o.Preserveperm
	}
	if o.Preserveattr != nil {
		toSerialize["preserveattr"] = o.Preserveattr
	}
	if o.Delayupdates != nil {
		toSerialize["delayupdates"] = o.Delayupdates
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	return json.Marshal(toSerialize)
}

type NullableRsynctaskCreate0 struct {
	value *RsynctaskCreate0
	isSet bool
}

func (v NullableRsynctaskCreate0) Get() *RsynctaskCreate0 {
	return v.value
}

func (v *NullableRsynctaskCreate0) Set(val *RsynctaskCreate0) {
	v.value = val
	v.isSet = true
}

func (v NullableRsynctaskCreate0) IsSet() bool {
	return v.isSet
}

func (v *NullableRsynctaskCreate0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRsynctaskCreate0(val *RsynctaskCreate0) *NullableRsynctaskCreate0 {
	return &NullableRsynctaskCreate0{value: val, isSet: true}
}

func (v NullableRsynctaskCreate0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRsynctaskCreate0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
