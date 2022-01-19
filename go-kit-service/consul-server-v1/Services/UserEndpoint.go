package Services

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"golang.org/x/time/rate"
	"gomicro/util"
	"strconv"
)

//定义response和request的格式
type UserRequest struct {
	//Uid是你自定义的，想叫什么就叫什么
	Uid    int `json:"uid"`
	Method string
	Token  string
}
type UserResponse struct {
	Result string `json:"result"`
}

//token验证中间件
func CheckTokenMiddleware() endpoint.Middleware { //Middleware type Middleware func(Endpoint) Endpoint
	return func(next endpoint.Endpoint) endpoint.Endpoint { //Endpoint type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			r := request.(UserRequest) //通过类型断言获取请求结构体
			uc := UserClaim{}
			//下面的r.Token是在代码DecodeUserRequest那里封装进去的
			getToken, err := jwt.ParseWithClaims(r.Token, &uc, func(token *jwt.Token) (i interface{}, e error) {
				return []byte(secKey), err
			})
			fmt.Println(err, 123)
			if getToken != nil && getToken.Valid { //验证通过
				newCtx := context.WithValue(ctx, "LoginUser", getToken.Claims.(*UserClaim).Uname)
				return next(newCtx, request)
			} else {
				return nil, util.NewMyError(403, "error token")
			}

			//logger.Log("method", r.Method, "event", "get user", "userid", r.Uid)

		}
	}
}

//日志中间件，增加限流功能
func UserServiceLogMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {

			r := request.(UserRequest) //获取到了r，就可以用我们的服务了
			//要成双成对 k-v
			logger.Log("method", r.Method, "event", "get user", "userid", r.Uid)
			return next(ctx, request)
		}
	}
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
		fmt.Println("当前登录用户为", ctx.Value("LoginUser"))
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
