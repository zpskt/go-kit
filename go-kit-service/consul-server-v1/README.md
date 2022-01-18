手动添加服务
>curl \
--request PUT \
--data @myservice.json \
localhost:8500/v1/agent/service/register

手动取消注册  
>curl \
--request PUT \
localhost:8500/v1/agent/service/deregister/注册时候的ID

下载consul相关api

    go getgithub.com/hashicorp/consul

命令行执行main.go 可以实现在不同端口开启服务

    go run main.go --name userservice -p 8081
    go run main.go --name userservice -p 8080

API限流
自定义错误体

熔断
go get github.com/afex/hystrix-go
>场景：假设
>
熔断补救：服务降级
异步执行，和服务降级