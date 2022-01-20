###介绍
此版本在consul-server-v1演化而来，因为单机继承jwt，所以和v1版本的client不匹配，所以单独变成此版本。
熔断限流等功能在v1都可以看到并测试，所以我们在这里关注与jwt。
###文档说明
AccessEndpoint.go文件里放的主要是jwt文档里面相关的有：  
1. token设置的60s过期
2. 定义request和response的结构体和endpoint方法
AccessTransport.go文件是对新的request和response的结构体进行解析。  


在UserEndpoint。go文件中新增一个验证中间件AccessEndpoint，所有的中间件都在main方法里面包裹调用。

    	endp := Services.RateLimit(limit)(Services.UserServiceLogMiddleware(logger)(Services.CheckTokenMiddleware()(Services.GenUserEnpoint(user)))) //调用限流代码生成的中间件
如果把AccessEndpoint中间件注销，那么server就可以变成v1版本的样子了
###使用方法

首先生成证书  

    go run genras.go  
会在./pem/下生成公私钥文件。
需要本机运行consul服务中心8500端口
运行主文件

    go run main.go -name userservice -p 8082
此时访问 http://localhost:8082/user/101 会提示error token。

在main.go中新加入了一个路由/access-token POST方法访问提供username和userpass，会返回token。
![](/Users/zp/Downloads/iShot2022-01-20 13.45.03.png)
获取到token后，加入到Query里面，就可以访问
![](/Users/zp/Downloads/iShot2022-01-20 13.48.19.png)