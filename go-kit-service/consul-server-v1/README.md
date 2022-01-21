###介绍  
角色：service,发布服务  
用的consul  
### 项目结构

>-| &ensp; Services  
>----| &ensp; UserEndpoint.go  
> request,response结构体，对应业务实体的对应方法， 日志中间件，限流中间件  
>----| &ensp; Usertransport.go  
> 怎么去传？传什么，做了这些事情。把resquest，response改成我们和对方能看得懂的样子  
>----| &ensp; UserService.go  
> 业务类，要干什么都在这里写。1.写接口2.定义实现类3.继承接口，名字要和1中的抽象函数里面一样  
> -| &ensp; util  
> ----| &ensp; consul.go  
> ----| &ensp; MyError.go  
> 类似endpoint，定义错误结构体  
> ----| &ensp; MyErrorEncoder.go    
> 实现一个返回类型标准的返回错误函数，httptransport可以切片直接调用
###consul使用
1.手动添加服务
>curl \
--request PUT \
--data @myservice.json \
localhost:8500/v1/agent/service/register

2.手动取消注册  
>curl \
--request PUT \
localhost:8500/v1/agent/service/deregister/注册时候的ID

3.下载consul相关api

    go getgithub.com/hashicorp/consul
4.api的使用集成在代码中
### 使用方法
命令行执行main.go 可以实现在不同端口开启服务

    go run main.go --name userservice -p 8082
    go run main.go --name userservice -p 8083
#### 功能
中间件形式：日志，限流
util文件夹：自定义错误体
