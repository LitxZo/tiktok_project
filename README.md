## 目录结构说明
- api：相当于controllor，处理web请求
- cmd：启动项目的相关命令
- conf：系统配置相关
- dao：数据库处理相关
- global：全局相关
- log：日志
- middleware：中间件
- model：数据库model
- router：路由
- service：服务层，具体服务方法
- - dto：各种request和response的对应结构体
- utils：各种工具

## 下载依赖
```
go mod tidy
```

## 运行项目
```
go run main.go
```
  

### github项目地址：https://github.com/LitxZo/tiktok_project.git 
<br>
<br>


## 项目说明
### 处理流程，以注册为例
1. 客户端发来请求
2. 通过router/core.go中对应的方法进行处理，这里调用的方法是api/user.go中的UserRegister方法。
3. 进行binding验证，没有错误调用register服务UserRegisterService，此方法在service/user_service.go中。
4. 服务方法中又调用了dao/user_dao.go中的UserRegisterDao方法，将用户名和密码存入数据库，并返回用户id。dao文件夹中的方法都是数据库处理相关。
5. 生成response并返回