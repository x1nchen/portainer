# \EndpointsApi

All URIs are relative to *http://portainer.domain/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EndpointCreate**](EndpointsApi.md#EndpointCreate) | **Post** /endpoints | Create a new endpoint
[**EndpointDelete**](EndpointsApi.md#EndpointDelete) | **Delete** /endpoints/{id} | Remove an endpoint
[**EndpointInspect**](EndpointsApi.md#EndpointInspect) | **Get** /endpoints/{id} | Inspect an endpoint
[**EndpointJob**](EndpointsApi.md#EndpointJob) | **Post** /endpoints/{id}/job | Execute a job on the endpoint host
[**EndpointList**](EndpointsApi.md#EndpointList) | **Get** /endpoints | List endpoints
[**EndpointUpdate**](EndpointsApi.md#EndpointUpdate) | **Put** /endpoints/{id} | Update an endpoint


# **EndpointCreate**
> Endpoint EndpointCreate(ctx, name, endpointType, optional)
Create a new endpoint

Create a new endpoint that will be used to manage a Docker environment. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name that will be used to identify this endpoint (example: my-endpoint) | 
  **endpointType** | **int32**| Environment type. Value must be one of: 1 (Docker environment), 2 (Agent environment), 3 (Azure environment) or 4 (Edge agent environment) | 
 **optional** | ***EndpointCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **uRL** | **optional.String**| URL or IP address of a Docker host (example: docker.mydomain.tld:2375). Defaults to local if not specified (Linux: /var/run/docker.sock, Windows: //./pipe/docker_engine) | 
 **publicURL** | **optional.String**| URL or IP address where exposed containers will be reachable. Defaults to URL if not specified (example: docker.mydomain.tld:2375) | 
 **groupID** | **optional.String**| Endpoint group identifier. If not specified will default to 1 (unassigned). | 
 **tLS** | **optional.String**| Require TLS to connect against this endpoint (example: true) | 
 **tLSSkipVerify** | **optional.String**| Skip server verification when using TLS (example: false) | 
 **tLSSkipClientVerify** | **optional.String**| Skip client verification when using TLS (example: false) | 
 **tLSCACertFile** | **optional.Interface of *os.File**| TLS CA certificate file | 
 **tLSCertFile** | **optional.Interface of *os.File**| TLS client certificate file | 
 **tLSKeyFile** | **optional.Interface of *os.File**| TLS client key file | 
 **azureApplicationID** | **optional.String**| Azure application ID. Required if endpoint type is set to 3 | 
 **azureTenantID** | **optional.String**| Azure tenant ID. Required if endpoint type is set to 3 | 
 **azureAuthenticationKey** | **optional.String**| Azure authentication key. Required if endpoint type is set to 3 | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointDelete**
> EndpointDelete(ctx, id)
Remove an endpoint

Remove an endpoint. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Endpoint identifier | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointInspect**
> Endpoint EndpointInspect(ctx, id)
Inspect an endpoint

Retrieve details abount an endpoint. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Endpoint identifier | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointJob**
> Endpoint EndpointJob(ctx, id, method, nodeName, body, optional)
Execute a job on the endpoint host

Execute a job (script) on the underlying host of the endpoint. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Endpoint identifier | 
  **method** | **string**| Job execution method. Possible values: file or string. | 
  **nodeName** | **string**| Optional. Hostname of a node when targeting a Portainer agent cluster. | 
  **body** | [**EndpointJobRequest**](EndpointJobRequest.md)| Job details. Required when method equals string. | 
 **optional** | ***EndpointJobOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointJobOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **image** | **optional.String**| Container image which will be used to execute the job. Required when method equals file. | 
 **file** | **optional.Interface of *os.File**| Job script file. Required when method equals file. | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: multipart/form-data
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointList**
> EndpointListResponse EndpointList(ctx, )
List endpoints

List all endpoints based on the current user authorizations. Will return all endpoints if using an administrator account otherwise it will only return authorized endpoints. **Access policy**: restricted 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EndpointListResponse**](EndpointListResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EndpointUpdate**
> EndpointUpdate(ctx, id, body)
Update an endpoint

Update an endpoint. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Endpoint identifier | 
  **body** | [**EndpointUpdateRequest**](EndpointUpdateRequest.md)| Endpoint details | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

