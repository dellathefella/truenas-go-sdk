/*
TrueNAS RESTful API

Go SDK for interacting with TrueNAS APIs (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package angelfish

import (
	"encoding/json"
)

// UpdateUserParams struct for UpdateUserParams
type UpdateUserParams struct {
	Uid                  *int32                 `json:"uid,omitempty"`
	Username             *string                `json:"username,omitempty"`
	Group                *int32                 `json:"group,omitempty"`
	Home                 *string                `json:"home,omitempty"`
	HomeMode             *string                `json:"home_mode,omitempty"`
	Shell                *string                `json:"shell,omitempty"`
	FullName             *string                `json:"full_name,omitempty"`
	Email                NullableString         `json:"email,omitempty"`
	Password             *string                `json:"password,omitempty"`
	PasswordDisabled     *bool                  `json:"password_disabled,omitempty"`
	Locked               *bool                  `json:"locked,omitempty"`
	MicrosoftAccount     *bool                  `json:"microsoft_account,omitempty"`
	Smb                  *bool                  `json:"smb,omitempty"`
	Sudo                 *bool                  `json:"sudo,omitempty"`
	SudoNopasswd         *bool                  `json:"sudo_nopasswd,omitempty"`
	SudoCommands         []string               `json:"sudo_commands,omitempty"`
	Sshpubkey            NullableString         `json:"sshpubkey,omitempty"`
	Groups               []int32                `json:"groups,omitempty"`
	Attributes           map[string]interface{} `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateUserParams UpdateUserParams

// NewUpdateUserParams instantiates a new UpdateUserParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserParams() *UpdateUserParams {
	this := UpdateUserParams{}
	return &this
}

// NewUpdateUserParamsWithDefaults instantiates a new UpdateUserParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserParamsWithDefaults() *UpdateUserParams {
	this := UpdateUserParams{}
	return &this
}

// GetUid returns the Uid field value if set, zero value otherwise.
func (o *UpdateUserParams) GetUid() int32 {
	if o == nil || isNil(o.Uid) {
		var ret int32
		return ret
	}
	return *o.Uid
}

// GetUidOk returns a tuple with the Uid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetUidOk() (*int32, bool) {
	if o == nil || isNil(o.Uid) {
		return nil, false
	}
	return o.Uid, true
}

// HasUid returns a boolean if a field has been set.
func (o *UpdateUserParams) HasUid() bool {
	if o != nil && !isNil(o.Uid) {
		return true
	}

	return false
}

// SetUid gets a reference to the given int32 and assigns it to the Uid field.
func (o *UpdateUserParams) SetUid(v int32) {
	o.Uid = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateUserParams) GetUsername() string {
	if o == nil || isNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetUsernameOk() (*string, bool) {
	if o == nil || isNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UpdateUserParams) HasUsername() bool {
	if o != nil && !isNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateUserParams) SetUsername(v string) {
	o.Username = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *UpdateUserParams) GetGroup() int32 {
	if o == nil || isNil(o.Group) {
		var ret int32
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetGroupOk() (*int32, bool) {
	if o == nil || isNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *UpdateUserParams) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given int32 and assigns it to the Group field.
func (o *UpdateUserParams) SetGroup(v int32) {
	o.Group = &v
}

// GetHome returns the Home field value if set, zero value otherwise.
func (o *UpdateUserParams) GetHome() string {
	if o == nil || isNil(o.Home) {
		var ret string
		return ret
	}
	return *o.Home
}

// GetHomeOk returns a tuple with the Home field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetHomeOk() (*string, bool) {
	if o == nil || isNil(o.Home) {
		return nil, false
	}
	return o.Home, true
}

// HasHome returns a boolean if a field has been set.
func (o *UpdateUserParams) HasHome() bool {
	if o != nil && !isNil(o.Home) {
		return true
	}

	return false
}

// SetHome gets a reference to the given string and assigns it to the Home field.
func (o *UpdateUserParams) SetHome(v string) {
	o.Home = &v
}

// GetHomeMode returns the HomeMode field value if set, zero value otherwise.
func (o *UpdateUserParams) GetHomeMode() string {
	if o == nil || isNil(o.HomeMode) {
		var ret string
		return ret
	}
	return *o.HomeMode
}

// GetHomeModeOk returns a tuple with the HomeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetHomeModeOk() (*string, bool) {
	if o == nil || isNil(o.HomeMode) {
		return nil, false
	}
	return o.HomeMode, true
}

// HasHomeMode returns a boolean if a field has been set.
func (o *UpdateUserParams) HasHomeMode() bool {
	if o != nil && !isNil(o.HomeMode) {
		return true
	}

	return false
}

// SetHomeMode gets a reference to the given string and assigns it to the HomeMode field.
func (o *UpdateUserParams) SetHomeMode(v string) {
	o.HomeMode = &v
}

// GetShell returns the Shell field value if set, zero value otherwise.
func (o *UpdateUserParams) GetShell() string {
	if o == nil || isNil(o.Shell) {
		var ret string
		return ret
	}
	return *o.Shell
}

// GetShellOk returns a tuple with the Shell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetShellOk() (*string, bool) {
	if o == nil || isNil(o.Shell) {
		return nil, false
	}
	return o.Shell, true
}

// HasShell returns a boolean if a field has been set.
func (o *UpdateUserParams) HasShell() bool {
	if o != nil && !isNil(o.Shell) {
		return true
	}

	return false
}

// SetShell gets a reference to the given string and assigns it to the Shell field.
func (o *UpdateUserParams) SetShell(v string) {
	o.Shell = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *UpdateUserParams) GetFullName() string {
	if o == nil || isNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetFullNameOk() (*string, bool) {
	if o == nil || isNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *UpdateUserParams) HasFullName() bool {
	if o != nil && !isNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *UpdateUserParams) SetFullName(v string) {
	o.FullName = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserParams) GetEmail() string {
	if o == nil || isNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserParams) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateUserParams) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UpdateUserParams) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *UpdateUserParams) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UpdateUserParams) UnsetEmail() {
	o.Email.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateUserParams) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateUserParams) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateUserParams) SetPassword(v string) {
	o.Password = &v
}

// GetPasswordDisabled returns the PasswordDisabled field value if set, zero value otherwise.
func (o *UpdateUserParams) GetPasswordDisabled() bool {
	if o == nil || isNil(o.PasswordDisabled) {
		var ret bool
		return ret
	}
	return *o.PasswordDisabled
}

// GetPasswordDisabledOk returns a tuple with the PasswordDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetPasswordDisabledOk() (*bool, bool) {
	if o == nil || isNil(o.PasswordDisabled) {
		return nil, false
	}
	return o.PasswordDisabled, true
}

// HasPasswordDisabled returns a boolean if a field has been set.
func (o *UpdateUserParams) HasPasswordDisabled() bool {
	if o != nil && !isNil(o.PasswordDisabled) {
		return true
	}

	return false
}

// SetPasswordDisabled gets a reference to the given bool and assigns it to the PasswordDisabled field.
func (o *UpdateUserParams) SetPasswordDisabled(v bool) {
	o.PasswordDisabled = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *UpdateUserParams) GetLocked() bool {
	if o == nil || isNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetLockedOk() (*bool, bool) {
	if o == nil || isNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *UpdateUserParams) HasLocked() bool {
	if o != nil && !isNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *UpdateUserParams) SetLocked(v bool) {
	o.Locked = &v
}

// GetMicrosoftAccount returns the MicrosoftAccount field value if set, zero value otherwise.
func (o *UpdateUserParams) GetMicrosoftAccount() bool {
	if o == nil || isNil(o.MicrosoftAccount) {
		var ret bool
		return ret
	}
	return *o.MicrosoftAccount
}

// GetMicrosoftAccountOk returns a tuple with the MicrosoftAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetMicrosoftAccountOk() (*bool, bool) {
	if o == nil || isNil(o.MicrosoftAccount) {
		return nil, false
	}
	return o.MicrosoftAccount, true
}

// HasMicrosoftAccount returns a boolean if a field has been set.
func (o *UpdateUserParams) HasMicrosoftAccount() bool {
	if o != nil && !isNil(o.MicrosoftAccount) {
		return true
	}

	return false
}

// SetMicrosoftAccount gets a reference to the given bool and assigns it to the MicrosoftAccount field.
func (o *UpdateUserParams) SetMicrosoftAccount(v bool) {
	o.MicrosoftAccount = &v
}

// GetSmb returns the Smb field value if set, zero value otherwise.
func (o *UpdateUserParams) GetSmb() bool {
	if o == nil || isNil(o.Smb) {
		var ret bool
		return ret
	}
	return *o.Smb
}

// GetSmbOk returns a tuple with the Smb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetSmbOk() (*bool, bool) {
	if o == nil || isNil(o.Smb) {
		return nil, false
	}
	return o.Smb, true
}

// HasSmb returns a boolean if a field has been set.
func (o *UpdateUserParams) HasSmb() bool {
	if o != nil && !isNil(o.Smb) {
		return true
	}

	return false
}

// SetSmb gets a reference to the given bool and assigns it to the Smb field.
func (o *UpdateUserParams) SetSmb(v bool) {
	o.Smb = &v
}

// GetSudo returns the Sudo field value if set, zero value otherwise.
func (o *UpdateUserParams) GetSudo() bool {
	if o == nil || isNil(o.Sudo) {
		var ret bool
		return ret
	}
	return *o.Sudo
}

// GetSudoOk returns a tuple with the Sudo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetSudoOk() (*bool, bool) {
	if o == nil || isNil(o.Sudo) {
		return nil, false
	}
	return o.Sudo, true
}

// HasSudo returns a boolean if a field has been set.
func (o *UpdateUserParams) HasSudo() bool {
	if o != nil && !isNil(o.Sudo) {
		return true
	}

	return false
}

// SetSudo gets a reference to the given bool and assigns it to the Sudo field.
func (o *UpdateUserParams) SetSudo(v bool) {
	o.Sudo = &v
}

// GetSudoNopasswd returns the SudoNopasswd field value if set, zero value otherwise.
func (o *UpdateUserParams) GetSudoNopasswd() bool {
	if o == nil || isNil(o.SudoNopasswd) {
		var ret bool
		return ret
	}
	return *o.SudoNopasswd
}

// GetSudoNopasswdOk returns a tuple with the SudoNopasswd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetSudoNopasswdOk() (*bool, bool) {
	if o == nil || isNil(o.SudoNopasswd) {
		return nil, false
	}
	return o.SudoNopasswd, true
}

// HasSudoNopasswd returns a boolean if a field has been set.
func (o *UpdateUserParams) HasSudoNopasswd() bool {
	if o != nil && !isNil(o.SudoNopasswd) {
		return true
	}

	return false
}

// SetSudoNopasswd gets a reference to the given bool and assigns it to the SudoNopasswd field.
func (o *UpdateUserParams) SetSudoNopasswd(v bool) {
	o.SudoNopasswd = &v
}

// GetSudoCommands returns the SudoCommands field value if set, zero value otherwise.
func (o *UpdateUserParams) GetSudoCommands() []string {
	if o == nil || isNil(o.SudoCommands) {
		var ret []string
		return ret
	}
	return o.SudoCommands
}

// GetSudoCommandsOk returns a tuple with the SudoCommands field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetSudoCommandsOk() ([]string, bool) {
	if o == nil || isNil(o.SudoCommands) {
		return nil, false
	}
	return o.SudoCommands, true
}

// HasSudoCommands returns a boolean if a field has been set.
func (o *UpdateUserParams) HasSudoCommands() bool {
	if o != nil && !isNil(o.SudoCommands) {
		return true
	}

	return false
}

// SetSudoCommands gets a reference to the given []string and assigns it to the SudoCommands field.
func (o *UpdateUserParams) SetSudoCommands(v []string) {
	o.SudoCommands = v
}

// GetSshpubkey returns the Sshpubkey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateUserParams) GetSshpubkey() string {
	if o == nil || isNil(o.Sshpubkey.Get()) {
		var ret string
		return ret
	}
	return *o.Sshpubkey.Get()
}

// GetSshpubkeyOk returns a tuple with the Sshpubkey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateUserParams) GetSshpubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Sshpubkey.Get(), o.Sshpubkey.IsSet()
}

// HasSshpubkey returns a boolean if a field has been set.
func (o *UpdateUserParams) HasSshpubkey() bool {
	if o != nil && o.Sshpubkey.IsSet() {
		return true
	}

	return false
}

// SetSshpubkey gets a reference to the given NullableString and assigns it to the Sshpubkey field.
func (o *UpdateUserParams) SetSshpubkey(v string) {
	o.Sshpubkey.Set(&v)
}

// SetSshpubkeyNil sets the value for Sshpubkey to be an explicit nil
func (o *UpdateUserParams) SetSshpubkeyNil() {
	o.Sshpubkey.Set(nil)
}

// UnsetSshpubkey ensures that no value is present for Sshpubkey, not even an explicit nil
func (o *UpdateUserParams) UnsetSshpubkey() {
	o.Sshpubkey.Unset()
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *UpdateUserParams) GetGroups() []int32 {
	if o == nil || isNil(o.Groups) {
		var ret []int32
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetGroupsOk() ([]int32, bool) {
	if o == nil || isNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *UpdateUserParams) HasGroups() bool {
	if o != nil && !isNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []int32 and assigns it to the Groups field.
func (o *UpdateUserParams) SetGroups(v []int32) {
	o.Groups = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpdateUserParams) GetAttributes() map[string]interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserParams) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdateUserParams) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *UpdateUserParams) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o UpdateUserParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Uid) {
		toSerialize["uid"] = o.Uid
	}
	if !isNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !isNil(o.Home) {
		toSerialize["home"] = o.Home
	}
	if !isNil(o.HomeMode) {
		toSerialize["home_mode"] = o.HomeMode
	}
	if !isNil(o.Shell) {
		toSerialize["shell"] = o.Shell
	}
	if !isNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !isNil(o.PasswordDisabled) {
		toSerialize["password_disabled"] = o.PasswordDisabled
	}
	if !isNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !isNil(o.MicrosoftAccount) {
		toSerialize["microsoft_account"] = o.MicrosoftAccount
	}
	if !isNil(o.Smb) {
		toSerialize["smb"] = o.Smb
	}
	if !isNil(o.Sudo) {
		toSerialize["sudo"] = o.Sudo
	}
	if !isNil(o.SudoNopasswd) {
		toSerialize["sudo_nopasswd"] = o.SudoNopasswd
	}
	if !isNil(o.SudoCommands) {
		toSerialize["sudo_commands"] = o.SudoCommands
	}
	if o.Sshpubkey.IsSet() {
		toSerialize["sshpubkey"] = o.Sshpubkey.Get()
	}
	if !isNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateUserParams) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateUserParams := _UpdateUserParams{}

	if err = json.Unmarshal(bytes, &varUpdateUserParams); err == nil {
		*o = UpdateUserParams(varUpdateUserParams)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "uid")
		delete(additionalProperties, "username")
		delete(additionalProperties, "group")
		delete(additionalProperties, "home")
		delete(additionalProperties, "home_mode")
		delete(additionalProperties, "shell")
		delete(additionalProperties, "full_name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "password")
		delete(additionalProperties, "password_disabled")
		delete(additionalProperties, "locked")
		delete(additionalProperties, "microsoft_account")
		delete(additionalProperties, "smb")
		delete(additionalProperties, "sudo")
		delete(additionalProperties, "sudo_nopasswd")
		delete(additionalProperties, "sudo_commands")
		delete(additionalProperties, "sshpubkey")
		delete(additionalProperties, "groups")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateUserParams struct {
	value *UpdateUserParams
	isSet bool
}

func (v NullableUpdateUserParams) Get() *UpdateUserParams {
	return v.value
}

func (v *NullableUpdateUserParams) Set(val *UpdateUserParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserParams(val *UpdateUserParams) *NullableUpdateUserParams {
	return &NullableUpdateUserParams{value: val, isSet: true}
}

func (v NullableUpdateUserParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}