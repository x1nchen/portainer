# RegistryCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name that will be used to identify this registry | [default to null]
**Type_** | **int32** | Registry Type. Valid values are: 1 (Quay.io), 2 (Azure container registry) or 3 (custom registry) | [default to null]
**URL** | **string** | URL or IP address of the Docker registry | [default to null]
**Authentication** | **bool** | Is authentication against this registry enabled | [default to null]
**Username** | **string** | Username used to authenticate against this registry | [default to null]
**Password** | **string** | Password used to authenticate against this registry | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


