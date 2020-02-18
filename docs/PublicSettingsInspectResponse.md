# PublicSettingsInspectResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogoURL** | **string** | URL to a logo that will be displayed on the login page as well as on top of the sidebar. Will use default Portainer logo when value is empty string | [optional] [default to null]
**DisplayExternalContributors** | **bool** | Whether to display or not external templates contributions as sub-menus in the UI. | [optional] [default to null]
**AuthenticationMethod** | **int32** | Active authentication method for the Portainer instance. Valid values are: 1 for managed or 2 for LDAP. | [optional] [default to null]
**AllowBindMountsForRegularUsers** | **bool** | Whether non-administrator should be able to use bind mounts when creating containers | [optional] [default to null]
**AllowPrivilegedModeForRegularUsers** | **bool** | Whether non-administrator should be able to use privileged mode when creating containers | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


