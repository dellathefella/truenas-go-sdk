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

// UpdateVMParams struct for UpdateVMParams
type UpdateVMParams struct {
	Name *string `json:"name,omitempty"`
	CpuMode *string `json:"cpu_mode,omitempty"`
	CpuModel NullableString `json:"cpu_model,omitempty"`
	Description *string `json:"description,omitempty"`
	// Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying `cores` attributes. `vcpus` specifies total number of CPU sockets. `cores` specifies number of cores per socket. `threads` specifies number of threads per core.
	Vcpus *int32 `json:"vcpus,omitempty"`
	// Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying `cores` attributes. `vcpus` specifies total number of CPU sockets. `cores` specifies number of cores per socket. `threads` specifies number of threads per core.
	Cores *int32 `json:"cores,omitempty"`
	// Maximum of 16 guest virtual CPUs are allowed. By default, every virtual CPU is configured as a separate package. Multiple cores can be configured per CPU by specifying `cores` attributes. `vcpus` specifies total number of CPU sockets. `cores` specifies number of cores per socket. `threads` specifies number of threads per core.
	Threads *int32 `json:"threads,omitempty"`
	Memory *int64 `json:"memory,omitempty"`
	Bootloader *string `json:"bootloader,omitempty"`
	// `devices` is a list of virtualized hardware to add to the newly created Virtual Machine. Failure to attach a device destroys the VM and any resources allocated by the VM devices.
	Devices []VMDevice `json:"devices,omitempty"`
	Autostart *bool `json:"autostart,omitempty"`
	// `hide_from_msr` is a boolean which when set will hide the KVM hypervisor from standard MSR based discovery and is useful to enable when doing GPU passthrough.
	HideFromMsr *bool `json:"hide_from_msr,omitempty"`
	// `ensure_display_device` when set ( the default ) will ensure that the guest always has access to a video device. For headless installations like ubuntu server this is required for the guest to operate properly. However for cases where consumer would like to use GPU passthrough and does not want a display device added should set this to `false`.
	EnsureDisplayDevice *bool `json:"ensure_display_device,omitempty"`
	Time *string `json:"time,omitempty"`
	// `shutdown_timeout` indicates the time in seconds the system waits for the VM to cleanly shutdown. During system shutdown, if the VM hasn't exited after a hardware shutdown signal has been sent by the system within `shutdown_timeout` seconds, system initiates poweroff for the VM to stop it.
	ShutdownTimeout *int32 `json:"shutdown_timeout,omitempty"`
	// `arch_type` refers to architecture type and can be specified for the guest. By default the value is `null` and system in this case will choose a reasonable default based on host. `machine_type` refers to machine type of the guest based on the architecture type selected with `arch_type`. By default the value is `null` and system in this case will choose a reasonable default based on `arch_type` configuration.
	ArchType NullableString `json:"arch_type,omitempty"`
	// `machine_type` refers to machine type of the guest based on the architecture type selected with `arch_type`. By default the value is `null` and system in this case will choose a reasonable default based on `arch_type` configuration.
	MachineType NullableString `json:"machine_type,omitempty"`
	Uuid NullableString `json:"uuid,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateVMParams UpdateVMParams

// NewUpdateVMParams instantiates a new UpdateVMParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVMParams() *UpdateVMParams {
	this := UpdateVMParams{}
	var cpuMode string = "CUSTOM"
	this.CpuMode = &cpuMode
	var vcpus int32 = 1
	this.Vcpus = &vcpus
	var cores int32 = 1
	this.Cores = &cores
	var threads int32 = 1
	this.Threads = &threads
	var bootloader string = "UEFI"
	this.Bootloader = &bootloader
	var autostart bool = true
	this.Autostart = &autostart
	var hideFromMsr bool = false
	this.HideFromMsr = &hideFromMsr
	var ensureDisplayDevice bool = true
	this.EnsureDisplayDevice = &ensureDisplayDevice
	var time string = "LOCAL"
	this.Time = &time
	var shutdownTimeout int32 = 90
	this.ShutdownTimeout = &shutdownTimeout
	return &this
}

// NewUpdateVMParamsWithDefaults instantiates a new UpdateVMParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVMParamsWithDefaults() *UpdateVMParams {
	this := UpdateVMParams{}
	var cpuMode string = "CUSTOM"
	this.CpuMode = &cpuMode
	var vcpus int32 = 1
	this.Vcpus = &vcpus
	var cores int32 = 1
	this.Cores = &cores
	var threads int32 = 1
	this.Threads = &threads
	var bootloader string = "UEFI"
	this.Bootloader = &bootloader
	var autostart bool = true
	this.Autostart = &autostart
	var hideFromMsr bool = false
	this.HideFromMsr = &hideFromMsr
	var ensureDisplayDevice bool = true
	this.EnsureDisplayDevice = &ensureDisplayDevice
	var time string = "LOCAL"
	this.Time = &time
	var shutdownTimeout int32 = 90
	this.ShutdownTimeout = &shutdownTimeout
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateVMParams) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateVMParams) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateVMParams) SetName(v string) {
	o.Name = &v
}

// GetCpuMode returns the CpuMode field value if set, zero value otherwise.
func (o *UpdateVMParams) GetCpuMode() string {
	if o == nil || o.CpuMode == nil {
		var ret string
		return ret
	}
	return *o.CpuMode
}

// GetCpuModeOk returns a tuple with the CpuMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetCpuModeOk() (*string, bool) {
	if o == nil || o.CpuMode == nil {
		return nil, false
	}
	return o.CpuMode, true
}

// HasCpuMode returns a boolean if a field has been set.
func (o *UpdateVMParams) HasCpuMode() bool {
	if o != nil && o.CpuMode != nil {
		return true
	}

	return false
}

// SetCpuMode gets a reference to the given string and assigns it to the CpuMode field.
func (o *UpdateVMParams) SetCpuMode(v string) {
	o.CpuMode = &v
}

// GetCpuModel returns the CpuModel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateVMParams) GetCpuModel() string {
	if o == nil || o.CpuModel.Get() == nil {
		var ret string
		return ret
	}
	return *o.CpuModel.Get()
}

// GetCpuModelOk returns a tuple with the CpuModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateVMParams) GetCpuModelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuModel.Get(), o.CpuModel.IsSet()
}

// HasCpuModel returns a boolean if a field has been set.
func (o *UpdateVMParams) HasCpuModel() bool {
	if o != nil && o.CpuModel.IsSet() {
		return true
	}

	return false
}

// SetCpuModel gets a reference to the given NullableString and assigns it to the CpuModel field.
func (o *UpdateVMParams) SetCpuModel(v string) {
	o.CpuModel.Set(&v)
}
// SetCpuModelNil sets the value for CpuModel to be an explicit nil
func (o *UpdateVMParams) SetCpuModelNil() {
	o.CpuModel.Set(nil)
}

// UnsetCpuModel ensures that no value is present for CpuModel, not even an explicit nil
func (o *UpdateVMParams) UnsetCpuModel() {
	o.CpuModel.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateVMParams) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateVMParams) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateVMParams) SetDescription(v string) {
	o.Description = &v
}

// GetVcpus returns the Vcpus field value if set, zero value otherwise.
func (o *UpdateVMParams) GetVcpus() int32 {
	if o == nil || o.Vcpus == nil {
		var ret int32
		return ret
	}
	return *o.Vcpus
}

// GetVcpusOk returns a tuple with the Vcpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetVcpusOk() (*int32, bool) {
	if o == nil || o.Vcpus == nil {
		return nil, false
	}
	return o.Vcpus, true
}

// HasVcpus returns a boolean if a field has been set.
func (o *UpdateVMParams) HasVcpus() bool {
	if o != nil && o.Vcpus != nil {
		return true
	}

	return false
}

// SetVcpus gets a reference to the given int32 and assigns it to the Vcpus field.
func (o *UpdateVMParams) SetVcpus(v int32) {
	o.Vcpus = &v
}

// GetCores returns the Cores field value if set, zero value otherwise.
func (o *UpdateVMParams) GetCores() int32 {
	if o == nil || o.Cores == nil {
		var ret int32
		return ret
	}
	return *o.Cores
}

// GetCoresOk returns a tuple with the Cores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetCoresOk() (*int32, bool) {
	if o == nil || o.Cores == nil {
		return nil, false
	}
	return o.Cores, true
}

// HasCores returns a boolean if a field has been set.
func (o *UpdateVMParams) HasCores() bool {
	if o != nil && o.Cores != nil {
		return true
	}

	return false
}

// SetCores gets a reference to the given int32 and assigns it to the Cores field.
func (o *UpdateVMParams) SetCores(v int32) {
	o.Cores = &v
}

// GetThreads returns the Threads field value if set, zero value otherwise.
func (o *UpdateVMParams) GetThreads() int32 {
	if o == nil || o.Threads == nil {
		var ret int32
		return ret
	}
	return *o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetThreadsOk() (*int32, bool) {
	if o == nil || o.Threads == nil {
		return nil, false
	}
	return o.Threads, true
}

// HasThreads returns a boolean if a field has been set.
func (o *UpdateVMParams) HasThreads() bool {
	if o != nil && o.Threads != nil {
		return true
	}

	return false
}

// SetThreads gets a reference to the given int32 and assigns it to the Threads field.
func (o *UpdateVMParams) SetThreads(v int32) {
	o.Threads = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *UpdateVMParams) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *UpdateVMParams) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *UpdateVMParams) SetMemory(v int64) {
	o.Memory = &v
}

// GetBootloader returns the Bootloader field value if set, zero value otherwise.
func (o *UpdateVMParams) GetBootloader() string {
	if o == nil || o.Bootloader == nil {
		var ret string
		return ret
	}
	return *o.Bootloader
}

// GetBootloaderOk returns a tuple with the Bootloader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetBootloaderOk() (*string, bool) {
	if o == nil || o.Bootloader == nil {
		return nil, false
	}
	return o.Bootloader, true
}

// HasBootloader returns a boolean if a field has been set.
func (o *UpdateVMParams) HasBootloader() bool {
	if o != nil && o.Bootloader != nil {
		return true
	}

	return false
}

// SetBootloader gets a reference to the given string and assigns it to the Bootloader field.
func (o *UpdateVMParams) SetBootloader(v string) {
	o.Bootloader = &v
}

// GetDevices returns the Devices field value if set, zero value otherwise.
func (o *UpdateVMParams) GetDevices() []VMDevice {
	if o == nil || o.Devices == nil {
		var ret []VMDevice
		return ret
	}
	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetDevicesOk() ([]VMDevice, bool) {
	if o == nil || o.Devices == nil {
		return nil, false
	}
	return o.Devices, true
}

// HasDevices returns a boolean if a field has been set.
func (o *UpdateVMParams) HasDevices() bool {
	if o != nil && o.Devices != nil {
		return true
	}

	return false
}

// SetDevices gets a reference to the given []VMDevice and assigns it to the Devices field.
func (o *UpdateVMParams) SetDevices(v []VMDevice) {
	o.Devices = v
}

// GetAutostart returns the Autostart field value if set, zero value otherwise.
func (o *UpdateVMParams) GetAutostart() bool {
	if o == nil || o.Autostart == nil {
		var ret bool
		return ret
	}
	return *o.Autostart
}

// GetAutostartOk returns a tuple with the Autostart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetAutostartOk() (*bool, bool) {
	if o == nil || o.Autostart == nil {
		return nil, false
	}
	return o.Autostart, true
}

// HasAutostart returns a boolean if a field has been set.
func (o *UpdateVMParams) HasAutostart() bool {
	if o != nil && o.Autostart != nil {
		return true
	}

	return false
}

// SetAutostart gets a reference to the given bool and assigns it to the Autostart field.
func (o *UpdateVMParams) SetAutostart(v bool) {
	o.Autostart = &v
}

// GetHideFromMsr returns the HideFromMsr field value if set, zero value otherwise.
func (o *UpdateVMParams) GetHideFromMsr() bool {
	if o == nil || o.HideFromMsr == nil {
		var ret bool
		return ret
	}
	return *o.HideFromMsr
}

// GetHideFromMsrOk returns a tuple with the HideFromMsr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetHideFromMsrOk() (*bool, bool) {
	if o == nil || o.HideFromMsr == nil {
		return nil, false
	}
	return o.HideFromMsr, true
}

// HasHideFromMsr returns a boolean if a field has been set.
func (o *UpdateVMParams) HasHideFromMsr() bool {
	if o != nil && o.HideFromMsr != nil {
		return true
	}

	return false
}

// SetHideFromMsr gets a reference to the given bool and assigns it to the HideFromMsr field.
func (o *UpdateVMParams) SetHideFromMsr(v bool) {
	o.HideFromMsr = &v
}

// GetEnsureDisplayDevice returns the EnsureDisplayDevice field value if set, zero value otherwise.
func (o *UpdateVMParams) GetEnsureDisplayDevice() bool {
	if o == nil || o.EnsureDisplayDevice == nil {
		var ret bool
		return ret
	}
	return *o.EnsureDisplayDevice
}

// GetEnsureDisplayDeviceOk returns a tuple with the EnsureDisplayDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetEnsureDisplayDeviceOk() (*bool, bool) {
	if o == nil || o.EnsureDisplayDevice == nil {
		return nil, false
	}
	return o.EnsureDisplayDevice, true
}

// HasEnsureDisplayDevice returns a boolean if a field has been set.
func (o *UpdateVMParams) HasEnsureDisplayDevice() bool {
	if o != nil && o.EnsureDisplayDevice != nil {
		return true
	}

	return false
}

// SetEnsureDisplayDevice gets a reference to the given bool and assigns it to the EnsureDisplayDevice field.
func (o *UpdateVMParams) SetEnsureDisplayDevice(v bool) {
	o.EnsureDisplayDevice = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *UpdateVMParams) GetTime() string {
	if o == nil || o.Time == nil {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetTimeOk() (*string, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *UpdateVMParams) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *UpdateVMParams) SetTime(v string) {
	o.Time = &v
}

// GetShutdownTimeout returns the ShutdownTimeout field value if set, zero value otherwise.
func (o *UpdateVMParams) GetShutdownTimeout() int32 {
	if o == nil || o.ShutdownTimeout == nil {
		var ret int32
		return ret
	}
	return *o.ShutdownTimeout
}

// GetShutdownTimeoutOk returns a tuple with the ShutdownTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVMParams) GetShutdownTimeoutOk() (*int32, bool) {
	if o == nil || o.ShutdownTimeout == nil {
		return nil, false
	}
	return o.ShutdownTimeout, true
}

// HasShutdownTimeout returns a boolean if a field has been set.
func (o *UpdateVMParams) HasShutdownTimeout() bool {
	if o != nil && o.ShutdownTimeout != nil {
		return true
	}

	return false
}

// SetShutdownTimeout gets a reference to the given int32 and assigns it to the ShutdownTimeout field.
func (o *UpdateVMParams) SetShutdownTimeout(v int32) {
	o.ShutdownTimeout = &v
}

// GetArchType returns the ArchType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateVMParams) GetArchType() string {
	if o == nil || o.ArchType.Get() == nil {
		var ret string
		return ret
	}
	return *o.ArchType.Get()
}

// GetArchTypeOk returns a tuple with the ArchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateVMParams) GetArchTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ArchType.Get(), o.ArchType.IsSet()
}

// HasArchType returns a boolean if a field has been set.
func (o *UpdateVMParams) HasArchType() bool {
	if o != nil && o.ArchType.IsSet() {
		return true
	}

	return false
}

// SetArchType gets a reference to the given NullableString and assigns it to the ArchType field.
func (o *UpdateVMParams) SetArchType(v string) {
	o.ArchType.Set(&v)
}
// SetArchTypeNil sets the value for ArchType to be an explicit nil
func (o *UpdateVMParams) SetArchTypeNil() {
	o.ArchType.Set(nil)
}

// UnsetArchType ensures that no value is present for ArchType, not even an explicit nil
func (o *UpdateVMParams) UnsetArchType() {
	o.ArchType.Unset()
}

// GetMachineType returns the MachineType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateVMParams) GetMachineType() string {
	if o == nil || o.MachineType.Get() == nil {
		var ret string
		return ret
	}
	return *o.MachineType.Get()
}

// GetMachineTypeOk returns a tuple with the MachineType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateVMParams) GetMachineTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MachineType.Get(), o.MachineType.IsSet()
}

// HasMachineType returns a boolean if a field has been set.
func (o *UpdateVMParams) HasMachineType() bool {
	if o != nil && o.MachineType.IsSet() {
		return true
	}

	return false
}

// SetMachineType gets a reference to the given NullableString and assigns it to the MachineType field.
func (o *UpdateVMParams) SetMachineType(v string) {
	o.MachineType.Set(&v)
}
// SetMachineTypeNil sets the value for MachineType to be an explicit nil
func (o *UpdateVMParams) SetMachineTypeNil() {
	o.MachineType.Set(nil)
}

// UnsetMachineType ensures that no value is present for MachineType, not even an explicit nil
func (o *UpdateVMParams) UnsetMachineType() {
	o.MachineType.Unset()
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateVMParams) GetUuid() string {
	if o == nil || o.Uuid.Get() == nil {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateVMParams) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *UpdateVMParams) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *UpdateVMParams) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *UpdateVMParams) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *UpdateVMParams) UnsetUuid() {
	o.Uuid.Unset()
}

func (o UpdateVMParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.CpuMode != nil {
		toSerialize["cpu_mode"] = o.CpuMode
	}
	if o.CpuModel.IsSet() {
		toSerialize["cpu_model"] = o.CpuModel.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Vcpus != nil {
		toSerialize["vcpus"] = o.Vcpus
	}
	if o.Cores != nil {
		toSerialize["cores"] = o.Cores
	}
	if o.Threads != nil {
		toSerialize["threads"] = o.Threads
	}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	if o.Bootloader != nil {
		toSerialize["bootloader"] = o.Bootloader
	}
	if o.Devices != nil {
		toSerialize["devices"] = o.Devices
	}
	if o.Autostart != nil {
		toSerialize["autostart"] = o.Autostart
	}
	if o.HideFromMsr != nil {
		toSerialize["hide_from_msr"] = o.HideFromMsr
	}
	if o.EnsureDisplayDevice != nil {
		toSerialize["ensure_display_device"] = o.EnsureDisplayDevice
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.ShutdownTimeout != nil {
		toSerialize["shutdown_timeout"] = o.ShutdownTimeout
	}
	if o.ArchType.IsSet() {
		toSerialize["arch_type"] = o.ArchType.Get()
	}
	if o.MachineType.IsSet() {
		toSerialize["machine_type"] = o.MachineType.Get()
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *UpdateVMParams) UnmarshalJSON(bytes []byte) (err error) {
	varUpdateVMParams := _UpdateVMParams{}

	if err = json.Unmarshal(bytes, &varUpdateVMParams); err == nil {
		*o = UpdateVMParams(varUpdateVMParams)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "cpu_mode")
		delete(additionalProperties, "cpu_model")
		delete(additionalProperties, "description")
		delete(additionalProperties, "vcpus")
		delete(additionalProperties, "cores")
		delete(additionalProperties, "threads")
		delete(additionalProperties, "memory")
		delete(additionalProperties, "bootloader")
		delete(additionalProperties, "devices")
		delete(additionalProperties, "autostart")
		delete(additionalProperties, "hide_from_msr")
		delete(additionalProperties, "ensure_display_device")
		delete(additionalProperties, "time")
		delete(additionalProperties, "shutdown_timeout")
		delete(additionalProperties, "arch_type")
		delete(additionalProperties, "machine_type")
		delete(additionalProperties, "uuid")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableUpdateVMParams struct {
	value *UpdateVMParams
	isSet bool
}

func (v NullableUpdateVMParams) Get() *UpdateVMParams {
	return v.value
}

func (v *NullableUpdateVMParams) Set(val *UpdateVMParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVMParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVMParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVMParams(val *UpdateVMParams) *NullableUpdateVMParams {
	return &NullableUpdateVMParams{value: val, isSet: true}
}

func (v NullableUpdateVMParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVMParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


