# \SettingsApi

All URIs are relative to *http://portainer.domain/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublicSettingsInspect**](SettingsApi.md#PublicSettingsInspect) | **Get** /settings/public | Retrieve Portainer public settings
[**SettingsInspect**](SettingsApi.md#SettingsInspect) | **Get** /settings | Retrieve Portainer settings
[**SettingsLDAPCheck**](SettingsApi.md#SettingsLDAPCheck) | **Put** /settings/authentication/checkLDAP | Test LDAP connectivity
[**SettingsUpdate**](SettingsApi.md#SettingsUpdate) | **Put** /settings | Update Portainer settings


# **PublicSettingsInspect**
> PublicSettingsInspectResponse PublicSettingsInspect(ctx, )
Retrieve Portainer public settings

Retrieve public settings. Returns a small set of settings that are not reserved to administrators only. **Access policy**: public 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PublicSettingsInspectResponse**](PublicSettingsInspectResponse.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsInspect**
> Settings SettingsInspect(ctx, )
Retrieve Portainer settings

Retrieve Portainer settings. **Access policy**: administrator 

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Settings**](Settings.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsLDAPCheck**
> SettingsLDAPCheck(ctx, body)
Test LDAP connectivity

Test LDAP connectivity using LDAP details. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettingsLdapCheckRequest**](SettingsLdapCheckRequest.md)| LDAP settings | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SettingsUpdate**
> Settings SettingsUpdate(ctx, body)
Update Portainer settings

Update Portainer settings. **Access policy**: administrator 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SettingsUpdateRequest**](SettingsUpdateRequest.md)| New settings | 

### Return type

[**Settings**](Settings.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

