# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/code.proto](#proto/code.proto)
  
    - [Code](#protocol.Code)
  
  
  

- [proto/common.proto](#proto/common.proto)
    - [Problem](#protocol.Problem)
    - [SubmitRecord](#protocol.SubmitRecord)
    - [UserInfo](#protocol.UserInfo)
  
    - [Problem.Difficluty](#protocol.Problem.Difficluty)
    - [Role](#protocol.Role)
  
  
  

- [proto/login.proto](#proto/login.proto)
    - [LogOut](#protocol.LogOut)
    - [LoginReq](#protocol.LoginReq)
    - [LoginResp](#protocol.LoginResp)
  
  
  
  

- [proto/problem.proto](#proto/problem.proto)
    - [ProblemList](#protocol.ProblemList)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="proto/code.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/code.proto


 


<a name="protocol.Code"></a>

### Code
Code : 状态码

| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 |  |
| INTERNAL | 1 | 服务端内部错误 |
| INVAILD_DATA | 2 | 非法数据，post数据无法反序列化 |
| PERMISSION_DENIED | 3 | 没有token，无法身份认证, 或者是是token错误之类，就是无法认证用户身份 |


 

 

 



<a name="proto/common.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/common.proto



<a name="protocol.Problem"></a>

### Problem
Problem : 题目


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | 题目id（目前暂定string，可能改成int64） |
| introduction | [string](#string) |  | 题目简介 |
| details | [string](#string) |  | 题目内容 |
| category | [string](#string) | repeated | 题目类别 |
| pass_rate | [float](#float) |  | 通过率 |
| difficluty | [Problem.Difficluty](#protocol.Problem.Difficluty) |  | 难度 |






<a name="protocol.SubmitRecord"></a>

### SubmitRecord
SubmitRecord : 提交情况


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| problem | [Problem](#protocol.Problem) |  | 题目 |
| submit_time | [int64](#int64) |  | 提交时间戳 |
| is_pass | [bool](#bool) |  | 是否通过 |






<a name="protocol.UserInfo"></a>

### UserInfo
UserInfo : 用户基本信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| sex | [bool](#bool) |  |  |
| role | [Role](#protocol.Role) |  |  |
| academy | [string](#string) |  | 学院 |
| major | [string](#string) |  | 专业 |
| email | [string](#string) |  |  |
| last_login | [int64](#int64) |  |  |





 


<a name="protocol.Problem.Difficluty"></a>

### Problem.Difficluty


| Name | Number | Description |
| ---- | ------ | ----------- |
| EASY | 0 |  |
| MEDIUM | 1 |  |
| HARD | 2 |  |



<a name="protocol.Role"></a>

### Role
Role : 用户角色（学生/老师...）

| Name | Number | Description |
| ---- | ------ | ----------- |
| STUDENT | 0 |  |
| TEACHER | 1 |  |


 

 

 



<a name="proto/login.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/login.proto



<a name="protocol.LogOut"></a>

### LogOut
登出 (get)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [Code](#protocol.Code) |  |  |






<a name="protocol.LoginReq"></a>

### LoginReq
登陆 (post)


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
| user | [UserInfo](#protocol.UserInfo) |  | 用户信息 |
| submit_records | [SubmitRecord](#protocol.SubmitRecord) | repeated | submit记录 （ TODO: 可以考虑提到新的协议中） |





 

 

 

 



<a name="proto/problem.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/problem.proto



<a name="protocol.ProblemList"></a>

### ProblemList
题目分类列表（get）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| categories | [string](#string) | repeated | 所有类别 |
| problems | [Problem](#protocol.Problem) | repeated | 默认题目 |





 

 

 

 



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

