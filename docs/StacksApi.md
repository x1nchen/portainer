# \StacksApi

All URIs are relative to *http://portainer.domain/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StackCreate**](StacksApi.md#StackCreate) | **Post** /stacks | Deploy a new stack
[**StackDelete**](StacksApi.md#StackDelete) | **Delete** /stacks/{id} | Remove a stack
[**StackFileInspect**](StacksApi.md#StackFileInspect) | **Get** /stacks/{id}/file | Retrieve the content of the Stack file for the specified stack
[**StackInspect**](StacksApi.md#StackInspect) | **Get** /stacks/{id} | Inspect a stack
[**StackList**](StacksApi.md#StackList) | **Get** /stacks | List stacks
[**StackMigrate**](StacksApi.md#StackMigrate) | **Post** /stacks/{id}/migrate | Migrate a stack to another endpoint
[**StackUpdate**](StacksApi.md#StackUpdate) | **Put** /stacks/{id} | Update a stack


# **StackCreate**
> Stack StackCreate(ctx, type_, method, endpointId, optional)
Deploy a new stack

Deploy a new stack into a Docker environment specified via the endpoint identifier. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **int32**| Stack deployment type. Possible values: 1 (Swarm stack) or 2 (Compose stack). | 
  **method** | **string**| Stack deployment method. Possible values: file, string or repository. | 
  **endpointId** | **int32**| Identifier of the endpoint that will be used to deploy the stack. | 
 **optional** | ***StackCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StackCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of StackCreateRequest**](StackCreateRequest.md)| Stack details. Required when method equals string or repository. | 
 **name** | **optional.String**| Name of the stack. Required when method equals file. | 
 **endpointID** | **optional.String**| Endpoint identifier used to deploy the stack. Required when method equals file. | 
 **swarmID** | **optional.String**| Swarm cluster identifier. Required when method equals file and type equals 1. | 
 **file** | **optional.Interface of *os.File**| Stack file. Required when method equals file. | 
 **env** | **optional.String**| Environment variables passed during deployment, represented as a JSON array [{&#39;name&#39;: &#39;name&#39;, &#39;value&#39;: &#39;value&#39;}]. Optional, used when method equals file and type equals 1. | 

### Return type

[**Stack**](Stack.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackDelete**
> StackDelete(ctx, id, optional)
Remove a stack

Remove a stack. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 
 **optional** | ***StackDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StackDeleteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **external** | **optional.Bool**| Set to true to delete an external stack. Only external Swarm stacks are supported. | 
 **endpointId** | **optional.String**| Endpoint identifier used to remove an external stack (required when external is set to true) | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackFileInspect**
> StackFileInspectResponse StackFileInspect(ctx, id)
Retrieve the content of the Stack file for the specified stack

Get Stack file content. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 

### Return type

[**StackFileInspectResponse**](StackFileInspectResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackInspect**
> Stack StackInspect(ctx, id)
Inspect a stack

Retrieve details about a stack. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 

### Return type

[**Stack**](Stack.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackList**
> StackListResponse StackList(ctx, optional)
List stacks

List all stacks based on the current user authorizations. Will return all stacks if using an administrator account otherwise it will only return the list of stacks the user have access to. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***StackListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StackListOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| Filters to process on the stack list. Encoded as JSON (a map[string]string). For example, {\&quot;SwarmID\&quot;: \&quot;jpofkc0i9uo9wtx1zesuk649w\&quot;} will only return stacks that are part of the specified Swarm cluster. Available filters: EndpointID, SwarmID.  | 

### Return type

[**StackListResponse**](StackListResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackMigrate**
> Stack StackMigrate(ctx, id, optional)
Migrate a stack to another endpoint

Migrate a stack from an endpoint to another endpoint. It will re-create the stack inside the target endpoint before removing the original stack. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 
 **optional** | ***StackMigrateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StackMigrateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **endpointId** | **optional.Int32**| Stacks created before version 1.18.0 might not have an associated endpoint identifier. Use this optional parameter to set the endpoint identifier used by the stack. | 
 **body** | [**optional.Interface of StackMigrateRequest**](StackMigrateRequest.md)| Stack migration details. | 

### Return type

[**Stack**](Stack.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StackUpdate**
> Stack StackUpdate(ctx, id, body, optional)
Update a stack

Update a stack. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Stack identifier | 
  **body** | [**StackUpdateRequest**](StackUpdateRequest.md)| Stack details | 
 **optional** | ***StackUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a StackUpdateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **endpointId** | **optional.Int32**| Stacks created before version 1.18.0 might not have an associated endpoint identifier. Use this optional parameter to set the endpoint identifier used by the stack. | 

### Return type

[**Stack**](Stack.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

