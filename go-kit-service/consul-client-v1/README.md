客户端直连 
product.go

调用api的形式查询服务
method2.go

负载均衡器：轮询 随机
        mylb := lb.NewRoundRobin(endpointer) //使用go-kit自带的轮询

限流