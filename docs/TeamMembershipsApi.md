# \TeamMembershipsApi

All URIs are relative to *http://portainer.domain/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TeamMembershipCreate**](TeamMembershipsApi.md#TeamMembershipCreate) | **Post** /team_memberships | Create a new team membership
[**TeamMembershipDelete**](TeamMembershipsApi.md#TeamMembershipDelete) | **Delete** /team_memberships/{id} | Remove a team membership
[**TeamMembershipList**](TeamMembershipsApi.md#TeamMembershipList) | **Get** /team_memberships | List team memberships
[**TeamMembershipUpdate**](TeamMembershipsApi.md#TeamMembershipUpdate) | **Put** /team_memberships/{id} | Update a team membership


# **TeamMembershipCreate**
> TeamMembership TeamMembershipCreate(ctx, body)
Create a new team membership

Create a new team memberships. Access is only available to administrators leaders of the associated team. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TeamMembershipCreateRequest**](TeamMembershipCreateRequest.md)| Team membership details | 

### Return type

[**TeamMembership**](TeamMembership.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamMembershipDelete**
> TeamMembershipDelete(ctx, id)
Remove a team membership

Remove a team membership. Access is only available to administrators leaders of the associated team. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| TeamMembership identifier | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamMembershipList**
> TeamMembershipListResponse TeamMembershipList(ctx, )
List team memberships

List team memberships. Access is only available to administrators and team leaders. **Access policy**: restricted 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TeamMembershipListResponse**](TeamMembershipListResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TeamMembershipUpdate**
> TeamMembership TeamMembershipUpdate(ctx, id, body)
Update a team membership

Update a team membership. Access is only available to administrators leaders of the associated team. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| Team membership identifier | 
  **body** | [**TeamMembershipUpdateRequest**](TeamMembershipUpdateRequest.md)| Team membership details | 

### Return type

[**TeamMembership**](TeamMembership.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

