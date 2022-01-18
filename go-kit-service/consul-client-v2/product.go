package main

import (
	"context"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"gomicro2/Services"
	"net/url"
	"os"
)

func main() {
	tgt, _ := url.Parse("http://192.168.1.5:8080")
	//创建一个直连client，这里我们必须写两个func,一个是如何请求,一个是响应我们怎么处理
	client := httptransport.NewClient("GET", tgt, Services.GetUserInfo_Request, Services.GetUserInfo_Response)

	//通过这个拿到了定义在服务端的endpoint也就是上面这段代码return出来的函数，直接在本地就可以调用服务端的代码
	getUserInfo := client.Endpoint()

	ctx := context.Background() //创建一个上下文

	//执行
	res, err := getUserInfo(ctx, Services.UserRequest{Uid: 11}) //使用go-kit插件来直接调用服务
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//断言，获取返回值
	userinfo := res.(Services.UserResponse)
	fmt.Println(userinfo.Result)

}
