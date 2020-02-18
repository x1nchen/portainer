# EndpointGroup

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | Endpoint group identifier | [optional] [default to null]
**Name** | **string** | Endpoint group name | [optional] [default to null]
**Description** | **string** | Endpoint group description | [optional] [default to null]
**AuthorizedUsers** | **[]int32** | List of user identifiers authorized to connect to this endpoint group. Will be inherited by endpoints that are part of the group | [optional] [default to null]
**AuthorizedTeams** | **[]int32** | List of team identifiers authorized to connect to this endpoint. Will be inherited by endpoints that are part of the group | [optional] [default to null]
**Labels** | [**[]Pair**](Pair.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


