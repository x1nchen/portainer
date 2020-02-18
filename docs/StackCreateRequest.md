# StackCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the stack | [default to null]
**SwarmID** | **string** | Swarm cluster identifier. Required when creating a Swarm stack (type 1). | [optional] [default to null]
**StackFileContent** | **string** | Content of the Stack file. Required when using the &#39;string&#39; deployment method. | [optional] [default to null]
**RepositoryURL** | **string** | URL of a Git repository hosting the Stack file. Required when using the &#39;repository&#39; deployment method. | [optional] [default to null]
**RepositoryReferenceName** | **string** | Reference name of a Git repository hosting the Stack file. Used in &#39;repository&#39; deployment method. | [optional] [default to null]
**ComposeFilePathInRepository** | **string** | Path to the Stack file inside the Git repository. Will default to &#39;docker-compose.yml&#39; if not specified. | [optional] [default to null]
**RepositoryAuthentication** | **bool** | Use basic authentication to clone the Git repository. | [optional] [default to null]
**RepositoryUsername** | **string** | Username used in basic authentication. Required when RepositoryAuthentication is true. | [optional] [default to null]
**RepositoryPassword** | **string** | Password used in basic authentication. Required when RepositoryAuthentication is true. | [optional] [default to null]
**Env** | [**[]StackEnv**](Stack_Env.md) | A list of environment variables used during stack deployment | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


