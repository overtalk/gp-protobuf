# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/common.proto](#proto/common.proto)
    - [Problem](#protocol.Problem)
    - [ProblemExample](#protocol.ProblemExample)
    - [ProblemJudgeLimit](#protocol.ProblemJudgeLimit)
    - [SubmitRecord](#protocol.SubmitRecord)
    - [UserInfo](#protocol.UserInfo)
  
    - [ProblemDifficluty](#protocol.ProblemDifficluty)
    - [Role](#protocol.Role)
  
  
  

- [proto/login.proto](#proto/login.proto)
    - [LoginReq](#protocol.LoginReq)
    - [LoginResp](#protocol.LoginResp)
  
  
  
  

- [proto/problem_manage.proto](#proto/problem_manage.proto)
    - [AddProblemReq](#protocol.AddProblemReq)
    - [AddProblemResp](#protocol.AddProblemResp)
    - [EditProblemReq](#protocol.EditProblemReq)
    - [EditProblemResp](#protocol.EditProblemResp)
    - [GetProblemByIDReq](#protocol.GetProblemByIDReq)
    - [GetProblemByIDResp](#protocol.GetProblemByIDResp)
    - [GetProblemsReq](#protocol.GetProblemsReq)
    - [GetProblemsResp](#protocol.GetProblemsResp)
  
  
  
  

- [proto/status.proto](#proto/status.proto)
    - [Status](#protocol.Status)
  
    - [Code](#protocol.Code)
  
  
  

- [proto/user_manage.proto](#proto/user_manage.proto)
    - [AddUsersReq](#protocol.AddUsersReq)
    - [AddUsersResp](#protocol.AddUsersResp)
    - [DelUsersReq](#protocol.DelUsersReq)
    - [DelUsersResp](#protocol.DelUsersResp)
    - [GetUsersReq](#protocol.GetUsersReq)
    - [GetUsersResp](#protocol.GetUsersResp)
    - [UpdateUsersReq](#protocol.UpdateUsersReq)
    - [UpdateUsersResp](#protocol.UpdateUsersResp)
  
  
  
  

- [Scalar Value Types](#scalar-value-types)



<a name="proto/common.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/common.proto



<a name="protocol.Problem"></a>

### Problem
Problem : 题目


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 题目id |
| title | [string](#string) |  | 题目标题 |
| description | [string](#string) |  | 题目描述 |
| in | [string](#string) |  | 输入 |
| out | [string](#string) |  | 输出 |
| hint | [string](#string) |  | 题目提示 |
| in_out_examples | [ProblemExample](#protocol.ProblemExample) | repeated | 输入输出样例 |
| judge_limit | [ProblemJudgeLimit](#protocol.ProblemJudgeLimit) |  | 判题限制 |
| tags | [string](#string) | repeated | 题目标签 |
| difficluty | [ProblemDifficluty](#protocol.ProblemDifficluty) |  | 难度 |
| submit_time | [int64](#int64) |  | 提交次数 |
| accept_time | [int64](#int64) |  | 通过次数 |
| judge_in_file | [bytes](#bytes) |  | 判题输入输出文件，只有在新建题目时需要用到 |
| judge_out_file | [bytes](#bytes) |  |  |






<a name="protocol.ProblemExample"></a>

### ProblemExample
ProblemExample : 题目输入输出样例


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| input | [string](#string) |  |  |
| output | [string](#string) |  |  |






<a name="protocol.ProblemJudgeLimit"></a>

### ProblemJudgeLimit
ProblemJudgeLimit : 判题的限制


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| time | [string](#string) |  |  |
| mem | [string](#string) |  |  |






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
| id | [int64](#int64) |  |  |
| role | [Role](#protocol.Role) |  |  |
| name | [string](#string) |  |  |
| sex | [bool](#bool) |  |  |
| phone | [string](#string) |  |  |
| email | [string](#string) |  |  |
| school | [string](#string) |  |  |
| last_login | [int64](#int64) |  |  |
| create | [int64](#int64) |  |  |
| account | [string](#string) |  | 这两个字段只有在用户管理中的新增用户才会用到, 客户端向服务端发送数据是填充 |
| password | [string](#string) |  |  |





 


<a name="protocol.ProblemDifficluty"></a>

### ProblemDifficluty
ProblemDifficluty : 题目难度

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
| MANAGER | 2 |  |


 

 

 



<a name="proto/login.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/login.proto



<a name="protocol.LoginReq"></a>

### LoginReq
登陆 (post)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| account | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="protocol.LoginResp"></a>

### LoginResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| token | [string](#string) |  |  |
| user | [UserInfo](#protocol.UserInfo) |  | 用户信息 |
| submit_records | [SubmitRecord](#protocol.SubmitRecord) | repeated | submit记录 （ TODO: 可以考虑提到新的协议中） |





 

 

 

 



<a name="proto/problem_manage.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/problem_manage.proto



<a name="protocol.AddProblemReq"></a>

### AddProblemReq
新增题目


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| problem | [Problem](#protocol.Problem) |  |  |






<a name="protocol.AddProblemResp"></a>

### AddProblemResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.EditProblemReq"></a>

### EditProblemReq
编辑题目


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| problem | [Problem](#protocol.Problem) |  | 需要id |






<a name="protocol.EditProblemResp"></a>

### EditProblemResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| is_success | [bool](#bool) |  |  |






<a name="protocol.GetProblemByIDReq"></a>

### GetProblemByIDReq
根据ID获得题目具体信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="protocol.GetProblemByIDResp"></a>

### GetProblemByIDResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| problem | [Problem](#protocol.Problem) |  |  |






<a name="protocol.GetProblemsReq"></a>

### GetProblemsReq
获取全部题目信息（只下发 id &amp; title &amp; difficulty &amp; pass_rate）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag | [string](#string) |  |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.GetProblemsResp"></a>

### GetProblemsResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| problems | [Problem](#protocol.Problem) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |





 

 

 

 



<a name="proto/status.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/status.proto



<a name="protocol.Status"></a>

### Status



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [Code](#protocol.Code) |  |  |
| message | [string](#string) |  |  |





 


<a name="protocol.Code"></a>

### Code


| Name | Number | Description |
| ---- | ------ | ----------- |
| OK | 0 | ok |
| INTERNAL | 1 | 服务端内部错误 |
| DATA_LOSE | 2 | 数据序列化错误 |
| NO_TOKEN | 3 | 没有token |
| UNAUTHORIZATED | 4 | token错误 |


 

 

 



<a name="proto/user_manage.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/user_manage.proto



<a name="protocol.AddUsersReq"></a>

### AddUsersReq
批量增加用户


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [UserInfo](#protocol.UserInfo) | repeated |  |






<a name="protocol.AddUsersResp"></a>

### AddUsersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| succeed | [UserInfo](#protocol.UserInfo) | repeated |  |
| fail | [UserInfo](#protocol.UserInfo) | repeated |  |






<a name="protocol.DelUsersReq"></a>

### DelUsersReq
批量删除用户


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users_id | [int64](#int64) | repeated |  |






<a name="protocol.DelUsersResp"></a>

### DelUsersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| succeed | [int64](#int64) | repeated |  |
| fail | [int64](#int64) | repeated |  |






<a name="protocol.GetUsersReq"></a>

### GetUsersReq
批量获取用户信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| role | [Role](#protocol.Role) |  |  |
| get_all | [bool](#bool) |  |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.GetUsersResp"></a>

### GetUsersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| users | [UserInfo](#protocol.UserInfo) | repeated |  |
| page_index | [int64](#int64) |  |  |
| page_num | [int64](#int64) |  |  |






<a name="protocol.UpdateUsersReq"></a>

### UpdateUsersReq
批量修改用户


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [UserInfo](#protocol.UserInfo) | repeated | 只需要填充id以及改变的field |






<a name="protocol.UpdateUsersResp"></a>

### UpdateUsersResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [Status](#protocol.Status) |  |  |
| succeed | [UserInfo](#protocol.UserInfo) | repeated |  |
| fail | [UserInfo](#protocol.UserInfo) | repeated |  |





 

 

 

 



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

