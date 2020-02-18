# LdapSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnonymousMode** | **bool** | Enable this option if the server is configured for Anonymous access. When enabled, ReaderDN and Password will not be used. | [optional] [default to null]
**ReaderDN** | **string** | Account that will be used to search for users | [optional] [default to null]
**Password** | **string** | Password of the account that will be used to search users | [optional] [default to null]
**URL** | **string** | URL or IP address of the LDAP server | [optional] [default to null]
**TLSConfig** | [***TlsConfiguration**](TLSConfiguration.md) |  | [optional] [default to null]
**StartTLS** | **bool** | Whether LDAP connection should use StartTLS | [optional] [default to null]
**SearchSettings** | [**[]LdapSearchSettings**](LDAPSearchSettings.md) |  | [optional] [default to null]
**GroupSearchSettings** | [**[]LdapGroupSearchSettings**](LDAPGroupSearchSettings.md) |  | [optional] [default to null]
**AutoCreateUsers** | **bool** | Automatically provision users and assign them to matching LDAP group names | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


