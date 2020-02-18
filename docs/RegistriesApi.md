# \RegistriesApi

All URIs are relative to *http://portainer.domain/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RegistryCreate**](RegistriesApi.md#RegistryCreate) | **Post** /registries | Create a new registry
[**RegistryDelete**](RegistriesApi.md#RegistryDelete) | **Delete** /registries/{id} | Remove a registry
[**RegistryInspect**](RegistriesApi.md#RegistryInspect) | **Get** /registries/{id} | Inspect a registry
[**RegistryList**](RegistriesApi.md#RegistryList) | **Get** /registries | List registries
[**RegistryUpdate**](RegistriesApi.md#RegistryUpdate) | **Put** /registries/{id} | Update a registry


# **RegistryCreate**
> Registry RegistryCreate(ctx, body)
Create a new registry

Create a new registry. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RegistryCreateRequest**](RegistryCreateRequest.md)| Registry details | 

### Return type

[**Registry**](Registry.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistryDelete**
> RegistryDelete(ctx, id)
Remove a registry

Remove a registry. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Registry identifier | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistryInspect**
> Registry RegistryInspect(ctx, id)
Inspect a registry

Retrieve details about a registry. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Registry identifier | 

### Return type

[**Registry**](Registry.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistryList**
> RegistryListResponse RegistryList(ctx, )
List registries

List all registries based on the current user authorizations. Will return all registries if using an administrator account otherwise it will only return authorized registries. **Access policy**: restricted 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RegistryListResponse**](RegistryListResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegistryUpdate**
> Registry RegistryUpdate(ctx, id, body)
Update a registry

Update a registry. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Registry identifier | 
  **body** | [**RegistryUpdateRequest**](RegistryUpdateRequest.md)| Registry details | 

### Return type

[**Registry**](Registry.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

