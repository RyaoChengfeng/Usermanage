Usermanage
# 用户管理系统

## 用户

- default

可以修改自己的passwd和email

- admin

修改所有用户的username、passwd、email、permission(权限)，通过表单输入内容

查询用户的所有信息



## 内容

支持注册、登录、验证身份、管理用户（CURD）

邮件认证激活用户，否则不能登录，认证方式为访问URL

验证身份使用JWT

passwd采用md5加密

数据交互形式采用JSON

发送邮件采用TLS加密

用户数据存入mysql

验证URL、无效token存入redis

用户修改密码，用户改变权限导致登录态失效，实现方法为将现有token变为无效token



## API文档

现在不会写，咕咕咕