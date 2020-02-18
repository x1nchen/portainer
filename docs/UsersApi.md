# \UsersApi

All URIs are relative to *http://portainer.domain/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAdminCheck**](UsersApi.md#UserAdminCheck) | **Get** /users/admin/check | Check administrator account existence
[**UserAdminInit**](UsersApi.md#UserAdminInit) | **Post** /users/admin/init | Initialize administrator account
[**UserCreate**](UsersApi.md#UserCreate) | **Post** /users | Create a new user
[**UserDelete**](UsersApi.md#UserDelete) | **Delete** /users/{id} | Remove a user
[**UserInspect**](UsersApi.md#UserInspect) | **Get** /users/{id} | Inspect a user
[**UserList**](UsersApi.md#UserList) | **Get** /users | List users
[**UserMembershipsInspect**](UsersApi.md#UserMembershipsInspect) | **Get** /users/{id}/memberships | Inspect a user memberships
[**UserPasswordCheck**](UsersApi.md#UserPasswordCheck) | **Post** /users/{id}/passwd | Check password validity for a user
[**UserUpdate**](UsersApi.md#UserUpdate) | **Put** /users/{id} | Update a user


# **UserAdminCheck**
> UserAdminCheck(ctx, )
Check administrator account existence

Check if an administrator account exists in the database. **Access policy**: public 

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserAdminInit**
> User UserAdminInit(ctx, body)
Initialize administrator account

Initialize the 'admin' user account. **Access policy**: public 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserAdminInitRequest**](UserAdminInitRequest.md)| User details | 

### Return type

[**User**](User.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserCreate**
> UserSubset UserCreate(ctx, body)
Create a new user

Create a new Portainer user. Only team leaders and administrators can create users. Only administrators can create an administrator user account. **Access policy**: restricted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UserCreateRequest**](UserCreateRequest.md)| User details | 

### Return type

[**UserSubset**](UserSubset.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserDelete**
> UserDelete(ctx, id)
Remove a user

Remove a user. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserInspect**
> User UserInspect(ctx, id)
Inspect a user

Retrieve details about a user. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 

### Return type

[**User**](User.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserList**
> UserListResponse UserList(ctx, )
List users

List Portainer users. Non-administrator users will only be able to list other non-administrator user accounts. **Access policy**: restricted 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserListResponse**](UserListResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserMembershipsInspect**
> UserMembershipsResponse UserMembershipsInspect(ctx, id)
Inspect a user memberships

Inspect a user memberships. **Access policy**: authenticated 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 

### Return type

[**UserMembershipsResponse**](UserMembershipsResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserPasswordCheck**
> UserPasswordCheckResponse UserPasswordCheck(ctx, id, body)
Check password validity for a user

Check if the submitted password is valid for the specified user. **Access policy**: authenticated 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 
  **body** | [**UserPasswordCheckRequest**](UserPasswordCheckRequest.md)| User details | 

### Return type

[**UserPasswordCheckResponse**](UserPasswordCheckResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserUpdate**
> User UserUpdate(ctx, id, body)
Update a user

Update user details. A regular user account can only update his details. **Access policy**: authenticated 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int32**| User identifier | 
  **body** | [**UserUpdateRequest**](UserUpdateRequest.md)| User details | 

### Return type

[**User**](User.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

