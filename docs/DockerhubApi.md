# \DockerhubApi

All URIs are relative to *http://portainer.domain/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DockerHubInspect**](DockerhubApi.md#DockerHubInspect) | **Get** /dockerhub | Retrieve DockerHub information
[**DockerHubUpdate**](DockerhubApi.md#DockerHubUpdate) | **Put** /dockerhub | Update DockerHub information


# **DockerHubInspect**
> DockerHubSubset DockerHubInspect(ctx, )
Retrieve DockerHub information

Use this endpoint to retrieve the information used to connect to the DockerHub **Access policy**: authenticated 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DockerHubSubset**](DockerHubSubset.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DockerHubUpdate**
> DockerHub DockerHubUpdate(ctx, body)
Update DockerHub information

Use this endpoint to update the information used to connect to the DockerHub **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DockerHubUpdateRequest**](DockerHubUpdateRequest.md)| DockerHub information | 

### Return type

[**DockerHub**](DockerHub.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

