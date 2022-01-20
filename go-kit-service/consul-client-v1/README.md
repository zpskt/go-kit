### consul-client-v1介绍
角色：service客户端，调用服务的一方，当然他也可能是中间商，并不被浏览器直接调用。
### 项目结构  

>-| &ensp; Services  
>----| &ensp; UserEndpoint.go  
> client不关心怎么实现，只需要把request，response的结构体放在其中即可。   
>----| &ensp; Usertransport.go  
> client不关心怎么实现，只需要知道怎么对request，response处理  
>-| &ensp; util
>----| &ensp; UserUtil.go
> 使用consul-api连接consul服务中心，并获取指定的services名字，设置轮询负载均衡  
> -| &ensp; commandWay.go  
> -| &ensp; main.go
> 熔断机制，如果发生错误调用次一级的服务
### 使用方法
前提条件，需要server-v1开启。
监听的服务中心ip为：localhost:8500,服务名字为：userservice。
以上两个参数在util-UserUtil.go可以实现自定义修改。

        go run product.go
### 功能
util-UserUtil.go: consulapi , 负载均衡
客户端直连 
product.go

调用api的形式查询服务
method2.go

负载均衡器：轮询 随机
        mylb := lb.NewRoundRobin(endpointer) //使用go-kit自带的轮询

限流