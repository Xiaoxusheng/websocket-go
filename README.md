##  **Gin+websocket 实时通信**



### 1. swagger文档地址

```
http://localhost:8080/swagger/index.html#/
```



### 2. api 接口

| 接口路径               | 请求方式 |                          Param参数                           | 请求头是否携带token |        Body类型         |                           Body参数                           | 返回数据类型     | 接口说明       |
| :--------------------- | :------: | :----------------------------------------------------------: | :-----------------: | :---------------------: | :----------------------------------------------------------: | ---------------- | -------------- |
| /group/exit            |   GET    |                      **account** (账号)                      |          √          |           无            |                              无                              | application/json | 退出群聊       |
| /group/group           |   POST   |                              无                              |          √          |   multipart/form-data   |                             info                             | application/json | 创建群聊       |
| /group/grouperlist     |   GET    |                              无                              |          √          |           无            |                              无                              | application/json | 群列表         |
| /group/grouplist       |   GET    |                      **room_id** (群号)                      |          √          |           无            |                              无                              | application/json | 获取群成员列表 |
| /group/join            |   GET    |                      **room_id** (群号)                      |          √          |           无            |                              无                              | application/json | 加入群聊列表   |
| /user/delete           |   GET    |                      **account** (账号)                      |          √          |           无            |                              无                              | application/json | 删除好友列表   |
| /user/file             |   POST   |                              无                              |          √          | **multipart/form-data** |                     表单的name **file**                      | application/json | 上传文件       |
| /user/friend_list      |   GET    |                              无                              |          √          |           无            |                              无                              | application/json | 好友列表       |
| /user/get_message      |   GET    | **room_id** (房间号)  **pageSize(**分页查询最大数，**pageSize**默认值为1 ) |          √          |           无            |                              无                              | application/json | 聊天记录       |
| /user/join             |   GET    |                      **account** (账号)                      |          √          |           无            |                              无                              | application/json | 添加好友       |
| /user/login            |   POST   |                              无                              |          √          | **multipart/form-data** | **username** (用户名) **password(**密码)   **code**(验证码)  | application/json | 登录           |
| /user/online           |   GET    |                      **account** (账号)                      |          √          |           无            |                              无                              | application/json | 好友在线状态   |
| /user/recallchatrecord |   GET    |                   **message_id** (消息id)                    |          √          |           无            |                              无                              | application/json | 消息撤回       |
| /user/register         |   POST   |                              无                              |          ×          | **multipart/form-data** | **username**（用户名）   **password** (密码)       **email**(邮箱) | application/json | 注册           |
| /user/send_code        |   GRET   |                    **username** (用户名)                     |          ×          |           无            |                              无                              | application/json | 获取验证码     |
| /user/setheadpicture   |   POST   |                              无                              |          √          | **multipart/form-data** |                 表单的name为**HeadPicture**                  | application/json | 上传头像       |
| /user/userinfo         |   GET    |                              无                              |          √          |           无            |                              无                              | application/json | 个人信息       |
| /websocket        |   GET    |                       **token**(token)                       |          ×          |           无            |                              无                              | application/json | websocket连接  |



### 3 打包为二进制文件 ，可直接在服务器运行（Gin）

#### 1 打包

```
go build -o Gin
```

#### 2 后台运行

```
nohup ./Gin &
```

