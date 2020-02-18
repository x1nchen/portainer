# EndpointUpdateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name that will be used to identify this endpoint | [optional] [default to null]
**URL** | **string** | URL or IP address of a Docker host | [optional] [default to null]
**PublicURL** | **string** | URL or IP address where exposed containers will be reachable. Defaults to URL if not specified | [optional] [default to null]
**GroupID** | **int32** | Group identifier | [optional] [default to null]
**TLS** | **bool** | Require TLS to connect against this endpoint | [optional] [default to null]
**TLSSkipVerify** | **bool** | Skip server verification when using TLS | [optional] [default to null]
**TLSSkipClientVerify** | **bool** | Skip client verification when using TLS | [optional] [default to null]
**ApplicationID** | **string** | Azure application ID | [optional] [default to null]
**TenantID** | **string** | Azure tenant ID | [optional] [default to null]
**AuthenticationKey** | **string** | Azure authentication key | [optional] [default to null]
**UserAccessPolicies** | [***UserAccessPolicies**](UserAccessPolicies.md) |  | [optional] [default to null]
**TeamAccessPolicies** | [***TeamAccessPolicies**](TeamAccessPolicies.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


