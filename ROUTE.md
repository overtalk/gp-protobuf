# 路由

## 登陆模块

#### 登陆 
- /login 

#### 登出
- /logout   （get）




## 管理员模块

### 用户管理

#### 获得用户 
- /getUsers

#### 新增用户 
- /addUsers

#### 更新用户 
- /updateUsers

#### 删除用户 
- /delUsers

### 题目管理 （目前没有判题文件的部分）

#### 获得所有题目 （只有部分信息，包括题目标题，难度...）
- /getProblems

#### 获得题目具体信息
- /getProblemByID

#### 新增题目 
- /addProblem

#### 编辑题目
- /editProblem

### 班级管理 

#### 获得所有班级 （只有部分信息，包括班级名称，创建时间，总人数...）
- /getClasses

#### 获得班级具体信息
- /getClassByID

#### 新增班级
- /addClass

#### 编辑班级
- /editClass