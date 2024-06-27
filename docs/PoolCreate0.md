# PoolCreate0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDuplicateSerials** | Pointer to **bool** |  | [optional] 
**Checksum** | Pointer to **NullableString** |  | [optional] 
**Deduplication** | Pointer to **NullableString** |  | [optional] 
**Encryption** | Pointer to **bool** |  | [optional] 
**EncryptionOptions** | Pointer to [**PoolCreate0EncryptionOptions**](PoolCreate0EncryptionOptions.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Topology** | Pointer to [**PoolCreate0Topology**](PoolCreate0Topology.md) |  | [optional] 

## Methods

### NewPoolCreate0

`func NewPoolCreate0() *PoolCreate0`

NewPoolCreate0 instantiates a new PoolCreate0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolCreate0WithDefaults

`func NewPoolCreate0WithDefaults() *PoolCreate0`

NewPoolCreate0WithDefaults instantiates a new PoolCreate0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDuplicateSerials

`func (o *PoolCreate0) GetAllowDuplicateSerials() bool`

GetAllowDuplicateSerials returns the AllowDuplicateSerials field if non-nil, zero value otherwise.

### GetAllowDuplicateSerialsOk

`func (o *PoolCreate0) GetAllowDuplicateSerialsOk() (*bool, bool)`

GetAllowDuplicateSerialsOk returns a tuple with the AllowDuplicateSerials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicateSerials

`func (o *PoolCreate0) SetAllowDuplicateSerials(v bool)`

SetAllowDuplicateSerials sets AllowDuplicateSerials field to given value.

### HasAllowDuplicateSerials

`func (o *PoolCreate0) HasAllowDuplicateSerials() bool`

HasAllowDuplicateSerials returns a boolean if a field has been set.

### GetChecksum

`func (o *PoolCreate0) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *PoolCreate0) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *PoolCreate0) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *PoolCreate0) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### SetChecksumNil

`func (o *PoolCreate0) SetChecksumNil(b bool)`

 SetChecksumNil sets the value for Checksum to be an explicit nil

### UnsetChecksum
`func (o *PoolCreate0) UnsetChecksum()`

UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
### GetDeduplication

`func (o *PoolCreate0) GetDeduplication() string`

GetDeduplication returns the Deduplication field if non-nil, zero value otherwise.

### GetDeduplicationOk

`func (o *PoolCreate0) GetDeduplicationOk() (*string, bool)`

GetDeduplicationOk returns a tuple with the Deduplication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeduplication

`func (o *PoolCreate0) SetDeduplication(v string)`

SetDeduplication sets Deduplication field to given value.

### HasDeduplication

`func (o *PoolCreate0) HasDeduplication() bool`

HasDeduplication returns a boolean if a field has been set.

### SetDeduplicationNil

`func (o *PoolCreate0) SetDeduplicationNil(b bool)`

 SetDeduplicationNil sets the value for Deduplication to be an explicit nil

### UnsetDeduplication
`func (o *PoolCreate0) UnsetDeduplication()`

UnsetDeduplication ensures that no value is present for Deduplication, not even an explicit nil
### GetEncryption

`func (o *PoolCreate0) GetEncryption() bool`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *PoolCreate0) GetEncryptionOk() (*bool, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *PoolCreate0) SetEncryption(v bool)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *PoolCreate0) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetEncryptionOptions

`func (o *PoolCreate0) GetEncryptionOptions() PoolCreate0EncryptionOptions`

GetEncryptionOptions returns the EncryptionOptions field if non-nil, zero value otherwise.

### GetEncryptionOptionsOk

`func (o *PoolCreate0) GetEncryptionOptionsOk() (*PoolCreate0EncryptionOptions, bool)`

GetEncryptionOptionsOk returns a tuple with the EncryptionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionOptions

`func (o *PoolCreate0) SetEncryptionOptions(v PoolCreate0EncryptionOptions)`

SetEncryptionOptions sets EncryptionOptions field to given value.

### HasEncryptionOptions

`func (o *PoolCreate0) HasEncryptionOptions() bool`

HasEncryptionOptions returns a boolean if a field has been set.

### GetName

`func (o *PoolCreate0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoolCreate0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoolCreate0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PoolCreate0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTopology

`func (o *PoolCreate0) GetTopology() PoolCreate0Topology`

GetTopology returns the Topology field if non-nil, zero value otherwise.

### GetTopologyOk

`func (o *PoolCreate0) GetTopologyOk() (*PoolCreate0Topology, bool)`

GetTopologyOk returns a tuple with the Topology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopology

`func (o *PoolCreate0) SetTopology(v PoolCreate0Topology)`

SetTopology sets Topology field to given value.

### HasTopology

`func (o *PoolCreate0) HasTopology() bool`

HasTopology returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


