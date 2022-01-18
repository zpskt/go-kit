package Services

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
	"gomicro/util"
	"strconv"
)

//定义response和request的格式
type UserRequest struct {
	//Uid是你自定义的，想叫什么就叫什么
	Uid    int `json:"uid"`
	Method string
}
type UserResponse struct {
	Result string `json:"result"`
}

//endpoint中间件，增加限流功能
func RateLimit(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() {
				return nil, util.NewMyError(429, "toot many request") //使用我们自定的错误结构体
			}
			return next(ctx, request)
		}
	}
}

//endpoint其实就是个func
func GenUserEnpoint(userService IUserService) endpoint.Endpoint {
	//这个func是endpoint规定的返回格式
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		//通过类型断言获取请求结构体
		r := request.(UserRequest) //获取到了r，就可以用我们的服务了
		result := "noting"
		if r.Method == "GET" {
			result = userService.GetName(r.Uid) + strconv.Itoa(util.ServicePort)
		} else if r.Method == "DELETE" { //如果是删除
			err := userService.DelUser(r.Uid)
			if err != nil { //代表有错，无法删除
				result = err.Error()
			} else {
				result = fmt.Sprintf("userid为%d的用户删除成功", r.Uid)
			}
		}

		return UserResponse{Result: result}, nil
	}
}
