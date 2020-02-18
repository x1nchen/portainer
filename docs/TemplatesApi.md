# \TemplatesApi

All URIs are relative to *http://portainer.domain/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplateCreate**](TemplatesApi.md#TemplateCreate) | **Post** /templates | Create a new template
[**TemplateDelete**](TemplatesApi.md#TemplateDelete) | **Delete** /templates/{id} | Remove a template
[**TemplateInspect**](TemplatesApi.md#TemplateInspect) | **Get** /templates/{id} | Inspect a template
[**TemplateList**](TemplatesApi.md#TemplateList) | **Get** /templates | List available templates
[**TemplateUpdate**](TemplatesApi.md#TemplateUpdate) | **Put** /templates/{id} | Update a template


# **TemplateCreate**
> Template TemplateCreate(ctx, body)
Create a new template

Create a new template. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TemplateCreateRequest**](TemplateCreateRequest.md)| Template details | 

### Return type

[**Template**](Template.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplateDelete**
> TemplateDelete(ctx, id)
Remove a template

Remove a template. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Template identifier | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplateInspect**
> Template TemplateInspect(ctx, id)
Inspect a template

Retrieve details about a template. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Template identifier | 

### Return type

[**Template**](Template.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplateList**
> TemplateListResponse TemplateList(ctx, )
List available templates

List available templates. Administrator templates will not be listed for non-administrator users. **Access policy**: restricted 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TemplateListResponse**](TemplateListResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplateUpdate**
> TemplateUpdate(ctx, id, body)
Update a template

Update a template. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Template identifier | 
  **body** | [**TemplateUpdateRequest**](TemplateUpdateRequest.md)| Template details | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

