# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/code.proto](#proto/code.proto)
  
    - [Code](#protocol.Code)
  
  
  

- [proto/login.proto](#proto/login.proto)
    - [LogOutReq](#protocol.LogOutReq)
    - [LogOutResp](#protocol.LogOutResp)
    - [LoginReq](#protocol.LoginReq)
    - [LoginResp](#protocol.LoginResp)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="proto/code.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/code.proto


 


<a name="protocol.Code"></a>

### Code


| Name | Number | Description |
| ---- | ------ | ----------- |
| ok | 0 |  |
| internal | 1 |  |
| invaild_data | 2 |  |
| no_token | 3 |  |


 

 

 



<a name="proto/login.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/login.proto



<a name="protocol.LogOutReq"></a>

### LogOutReq
登出


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |






<a name="protocol.LogOutResp"></a>

### LogOutResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [Code](#protocol.Code) |  |  |






<a name="protocol.LoginReq"></a>

### LoginReq
登陆


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="protocol.LoginResp"></a>

### LoginResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [Code](#protocol.Code) |  |  |
| token | [string](#string) |  |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

