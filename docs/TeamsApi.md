# \TeamsApi

All URIs are relative to *http://portainer.domain/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamCreate**](TeamsApi.md#TeamCreate) | **Post** /teams | Create a new team
[**TeamDelete**](TeamsApi.md#TeamDelete) | **Delete** /teams/{id} | Remove a team
[**TeamInspect**](TeamsApi.md#TeamInspect) | **Get** /teams/{id} | Inspect a team
[**TeamList**](TeamsApi.md#TeamList) | **Get** /teams | List teams
[**TeamMembershipsInspect**](TeamsApi.md#TeamMembershipsInspect) | **Get** /teams/{id}/memberships | Inspect a team memberships
[**TeamUpdate**](TeamsApi.md#TeamUpdate) | **Put** /teams/{id} | Update a team


# **TeamCreate**
> Team TeamCreate(ctx, body)
Create a new team

Create a new team. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TeamCreateRequest**](TeamCreateRequest.md)| Team details | 

### Return type

[**Team**](Team.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamDelete**
> TeamDelete(ctx, id)
Remove a team

Remove a team. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Team identifier | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamInspect**
> Team TeamInspect(ctx, id)
Inspect a team

Retrieve details about a team. Access is only available for administrator and leaders of that team. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Team identifier | 

### Return type

[**Team**](Team.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamList**
> TeamListResponse TeamList(ctx, )
List teams

List teams. For non-administrator users, will only list the teams they are member of. **Access policy**: restricted 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TeamListResponse**](TeamListResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamMembershipsInspect**
> TeamMembershipsResponse TeamMembershipsInspect(ctx, id)
Inspect a team memberships

Inspect a team memberships. Access is only available for administrator and leaders of that team. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Team identifier | 

### Return type

[**TeamMembershipsResponse**](TeamMembershipsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamUpdate**
> TeamUpdate(ctx, id, body)
Update a team

Update a team. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Team identifier | 
  **body** | [**TeamUpdateRequest**](TeamUpdateRequest.md)| Team details | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

