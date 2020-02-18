# \EndpointGroupsApi

All URIs are relative to *http://portainer.domain/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndpointGroupAddEndpoint**](EndpointGroupsApi.md#EndpointGroupAddEndpoint) | **Put** /endpoint_groups/{id}/endpoints/{endpointId} | Add an endpoint to an endpoint group
[**EndpointGroupCreate**](EndpointGroupsApi.md#EndpointGroupCreate) | **Post** /endpoint_groups | Create a new endpoint
[**EndpointGroupDelete**](EndpointGroupsApi.md#EndpointGroupDelete) | **Delete** /endpoint_groups/{id} | Remove an endpoint group
[**EndpointGroupDeleteEndpoint**](EndpointGroupsApi.md#EndpointGroupDeleteEndpoint) | **Delete** /endpoint_groups/{id}/endpoints/{endpointId} | Remove an endpoint group
[**EndpointGroupInspect**](EndpointGroupsApi.md#EndpointGroupInspect) | **Get** /endpoint_groups/{id} | Inspect an endpoint group
[**EndpointGroupList**](EndpointGroupsApi.md#EndpointGroupList) | **Get** /endpoint_groups | List endpoint groups
[**EndpointGroupUpdate**](EndpointGroupsApi.md#EndpointGroupUpdate) | **Put** /endpoint_groups/{id} | Update an endpoint group


# **EndpointGroupAddEndpoint**
> EndpointGroupAddEndpoint(ctx, id, endpointId)
Add an endpoint to an endpoint group

Add an endpoint to an endpoint group **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| EndpointGroup identifier | 
  **endpointId** | **int32**| Endpoint identifier | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGroupCreate**
> EndpointGroup EndpointGroupCreate(ctx, body)
Create a new endpoint

Create a new endpoint group. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**EndpointGroupCreateRequest**](EndpointGroupCreateRequest.md)| Registry details | 

### Return type

[**EndpointGroup**](EndpointGroup.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGroupDelete**
> EndpointGroupDelete(ctx, id)
Remove an endpoint group

Remove an endpoint group. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| EndpointGroup identifier | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGroupDeleteEndpoint**
> EndpointGroupDeleteEndpoint(ctx, id, endpointId)
Remove an endpoint group

Remove an endpoint group. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| EndpointGroup identifier | 
  **endpointId** | **int32**| Endpoint identifier | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGroupInspect**
> EndpointGroup EndpointGroupInspect(ctx, id)
Inspect an endpoint group

Retrieve details abount an endpoint group. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Endpoint group identifier | 

### Return type

[**EndpointGroup**](EndpointGroup.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGroupList**
> EndpointGroupListResponse EndpointGroupList(ctx, )
List endpoint groups

List all endpoint groups based on the current user authorizations. Will return all endpoint groups if using an administrator account otherwise it will only return authorized endpoint groups. **Access policy**: restricted 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EndpointGroupListResponse**](EndpointGroupListResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointGroupUpdate**
> EndpointGroup EndpointGroupUpdate(ctx, id, body)
Update an endpoint group

Update an endpoint group. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| EndpointGroup identifier | 
  **body** | [**EndpointGroupUpdateRequest**](EndpointGroupUpdateRequest.md)| EndpointGroup details | 

### Return type

[**EndpointGroup**](EndpointGroup.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

