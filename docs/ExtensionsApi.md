# \ExtensionsApi

All URIs are relative to *http://portainer.domain/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExtensionCreate**](ExtensionsApi.md#ExtensionCreate) | **Post** /extensions | Enable an extension
[**ExtensionDelete**](ExtensionsApi.md#ExtensionDelete) | **Delete** /extensions/{id} | Disable an extension
[**ExtensionInspect**](ExtensionsApi.md#ExtensionInspect) | **Get** /extensions/{id} | Inspect an extension
[**ExtensionList**](ExtensionsApi.md#ExtensionList) | **Get** /extensions | List extensions
[**ExtensionUpdate**](ExtensionsApi.md#ExtensionUpdate) | **Put** /extensions/{id} | Update an extension


# **ExtensionCreate**
> ExtensionCreate(ctx, body)
Enable an extension

Enable an extension. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ExtensionCreateRequest**](ExtensionCreateRequest.md)| Extension details | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtensionDelete**
> ExtensionDelete(ctx, id)
Disable an extension

Disable an extension. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Extension identifier | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtensionInspect**
> Extension ExtensionInspect(ctx, id)
Inspect an extension

Retrieve details abount an extension. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| extension identifier | 

### Return type

[**Extension**](Extension.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtensionList**
> ExtensionListResponse ExtensionList(ctx, optional)
List extensions

List all extensions registered inside Portainer. If the store parameter is set to true, will retrieve extensions details from the online repository. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ExtensionListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ExtensionListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **store** | **optional.Bool**| Retrieve online information about extensions. Possible values: true or false. | 

### Return type

[**ExtensionListResponse**](ExtensionListResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExtensionUpdate**
> ExtensionUpdate(ctx, id, body)
Update an extension

Update an extension to a specific version of the extension. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Extension identifier | 
  **body** | [**ExtensionUpdateRequest**](ExtensionUpdateRequest.md)| Extension details | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

