# go-kit
go-kit实践
##案例
### [server-v1](https://github.com/zpskt/go-kit/tree/main/go-kit-service/consul-server-v1) 
日志，API限流，consulapi功能
### [client-v1](https://github.com/zpskt/go-kit/tree/main/go-kit-service/consul-client-v1)
负载均衡，consulapi
### [server-v2](https://github.com/zpskt/go-kit/tree/main/go-kit-service/consul-server-v2)
jwt验证
### 秒杀系统 
在seckill文件夹下
##工具
### [kong Apigateway](https://github.com/zpskt/go-kit/tree/main/tools/apigateway)
## 介绍
go-kit 是什么
>Go kit 是一个微服务工具包集合。利用它提供的额API和规范可以创建健壮、可维护性高的微服务体系

Go-kit的三层架构
>1、Service 
这里就是我们的业务类、接口等相关信息存放

2、EndPoint
定义Request、Response格式，并可以使用装饰器(闭包)包装函数,以此来实现各个中间件嵌套

3、Transport
主要负责与HTTP、gRPC、thrift等相关逻辑  

