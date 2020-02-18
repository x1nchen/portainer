# ResourceControl

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceID** | **string** | Docker resource identifier on which access control will be applied. In the case of a resource control applied to a stack, use the stack name as identifier | [optional] [default to null]
**Type_** | **string** | Type of Docker resource. Valid values are: container, volume service, secret, config or stack | [optional] [default to null]
**Public** | **bool** | Permit access to the associated resource to any user | [optional] [default to null]
**Users** | **[]int32** | List of user identifiers with access to the associated resource | [optional] [default to null]
**Teams** | **[]int32** | List of team identifiers with access to the associated resource | [optional] [default to null]
**SubResourceIDs** | **[]string** | List of Docker resources that will inherit this access control | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


