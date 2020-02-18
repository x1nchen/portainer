# Stack

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stack identifier | [optional] [default to null]
**Name** | **string** | Stack name | [optional] [default to null]
**Type_** | **int32** | Stack type. 1 for a Swarm stack, 2 for a Compose stack | [optional] [default to null]
**EndpointID** | **int32** | Endpoint identifier. Reference the endpoint that will be used for deployment  | [optional] [default to null]
**EntryPoint** | **string** | Path to the Stack file | [optional] [default to null]
**SwarmID** | **string** | Cluster identifier of the Swarm cluster where the stack is deployed | [optional] [default to null]
**ProjectPath** | **string** | Path on disk to the repository hosting the Stack file | [optional] [default to null]
**Env** | [**[]StackEnv**](Stack_Env.md) | A list of environment variables used during stack deployment | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


