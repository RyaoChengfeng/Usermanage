# 用户管理系统

## 用户

- default

可以修改自己的passwd和email

- admin

修改所有用户的username、passwd、email、permission(权限)，通过表单输入内容

查询用户的所有信息



## 内容

- 支持注册、登录、验证身份、管理用户（CURD）


- 邮件认证激活用户，否则不能登录，认证方式为访问URL
    - 实现：token中包含username信息， 进入相应链接http://localhost:1323/activate?verify=<tokenString>,匹配链接中tokenstring和访问的客户端发来的token，匹配成功后解析出username，进行激活账号
    - 激活状态在token中，未激活报错401

- 验证身份使用JWT


- passwd采用md5加密


- 数据交互形式采用JSON


- 发送邮件采用TLS加密


- 用户数据存入mysql,用户信息都在mysql中读写


- 用户修改密码，用户改变权限导致登录态失效
    - 实现：将有效token从redis中改为无效token，用户重新登录后redis重新获得对应的有效token，登录态查询redis有效token中是否有客户端发出的token



## API文档

不会写，咕咕咕



## 还未解决的问题

- jwt验证和登录态验证的中间件没有写出来

## 有时间可以尝试做的

- 有关数据改动写入log

- email发送后过一段时间可以重新发送一个，发送后上一个邮件失效
