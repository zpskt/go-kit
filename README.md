# go-kit
go-kit实践
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


## 安装
        go get -u github.com/gorilla/mux
        github.com/go-kit/kit/transport/http
## 组件架构
consul-client-v1和consul-server-v1可以连起来用
consul-server-v2新加入了jwt验证，需要本机用接口测试看结果