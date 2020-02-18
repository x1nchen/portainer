# \AuthApi

All URIs are relative to *http://portainer.domain/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateUser**](AuthApi.md#AuthenticateUser) | **Post** /auth | Authenticate a user


# **AuthenticateUser**
> AuthenticateUserResponse AuthenticateUser(ctx, body)
Authenticate a user

Use this endpoint to authenticate against Portainer using a username and password. **Access policy**: public 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AuthenticateUserRequest**](AuthenticateUserRequest.md)| Credentials used for authentication | 

### Return type

[**AuthenticateUserResponse**](AuthenticateUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

