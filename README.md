# 用户管理系统

## 角色
* 添加分组 /role/group/create [put]
* 删除分组 /role/group/delete [delete]
* 添加角色 /role/create [put]
* 修改角色 /role/modify [post]
* 删除角色 /role/delete [delete]
* 组内添加角色 /role/group/add [post]
* 组内删除角色 /role/group/remove [post]
* 关联权限 /role/permission/add [post]
* 取消权限关联 /role/permission/remove [post]

## 权限
* 添加分组 /permission/group/create [put]
* 删除分组 /permission/group/delete [delete]
* 添加权限 /permission/create [put]
* 修改权限 /permission/modify [post]
* 删除权限 /permission/delete [delete]
* 组内添加权限 /permission/group/add [post]
* 组内删除权限 /permission/group/remove [post]
* 关联角色 /permission/role/add [post]
* 取消角色关联 /permission/role/remove [post]

## 用户
* 注册 /register [post]
* 登陆 /login [post]
* 删除 /user/delete [delete]
* 查询 /user [get]
* 修改信息 /user/modify [post]
* 修改密码 /user/passwd [post]
* 关联角色 /user/role/add [post]
* 取消角色关联 /user/role/remove [post]

## 

* 用户基本信息
* 用户登陆验证：手机、邮箱、账号、QQ、微信、微博
* 位置信息
* 日志信息
* 扩展信息
