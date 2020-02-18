# Endpoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Endpoint identifier | [optional] [default to null]
**Name** | **string** | Endpoint name | [optional] [default to null]
**Type_** | **int32** | Endpoint environment type. 1 for a Docker environment, 2 for an agent on Docker environment or 3 for an Azure environment. | [optional] [default to null]
**URL** | **string** | URL or IP address of the Docker host associated to this endpoint | [optional] [default to null]
**PublicURL** | **string** | URL or IP address where exposed containers will be reachable | [optional] [default to null]
**GroupID** | **int32** | Endpoint group identifier | [optional] [default to null]
**AuthorizedUsers** | **[]int32** | List of user identifiers authorized to connect to this endpoint | [optional] [default to null]
**AuthorizedTeams** | **[]int32** | List of team identifiers authorized to connect to this endpoint | [optional] [default to null]
**TLSConfig** | [***TlsConfiguration**](TLSConfiguration.md) |  | [optional] [default to null]
**AzureCredentials** | [***AzureCredentials**](AzureCredentials.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


