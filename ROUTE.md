# 路由

## 登陆模块


| 功能描述 | 路由 | HTTP方法 |  
| ----- | ---- | ----- |
| 登陆 | /login  | POST |  
| 登出 | /logout | GET |  
 



## 管理员模块

### 用户管理

| 功能描述 | 路由 | HTTP方法 |  
| ----- | ---- | ----- | 
| 获得用户 | /getUsers  | POST |  
| 新增用户 | /addUsers | POST |  
| 更新用户 | /updateUsers | POST |  
| 删除用户 | /delUsers | POST |  

### 题目管理 （目前没有判题文件的部分）

| 功能描述 | 路由 | HTTP方法 |  
| ----- | ---- | ----- | 
| 获得所有题目 （只有部分信息，包括题目标题，难度...） | /getProblems  | POST |  
| 获得题目具体信息 | /getProblemByID | POST |  
| 新增题目 | /addProblem | POST |  
| 编辑题目 | /editProblem | POST |  


### 班级管理 

| 功能描述 | 路由 | HTTP方法 |  
| ----- | ---- | ----- |
| 获得所有班级 （只有部分信息，包括班级名称，创建时间，总人数...） | /getClasses  | POST |  
| 获得班级具体信息 | /getClassByID | POST |  
| 新增班级 | /addClass | POST |  
| 编辑班级 | /editClass | POST |  

