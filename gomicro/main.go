package main

import (
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux" //第三方路由
	. "gomicro/Services"
	"net/http"
)

func main() {
	//1.第一层service
	user := UserService{}
	//通过GenUserEnpoint调用服务
	endp := GenUserEnpoint(user)

	serverHanlder := httptransport.NewServer(endp, DecodeUserRequest, EncodeUserResponse)

	//路由模块
	r := mymux.NewRouter()
	//r.Handle(`/user/{uid:\d+}`, serverHanlder)
	{
		r.Methods("GET", "DELETE").Path(`/user/{uid:\d+}`).Handler(serverHanlder)
		//手动写一个health路由，写死
		r.Methods("GET").Path(`/health`).HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			//设置json格式
			writer.Header().Set("Content-type", "application/json")
			writer.Write([]byte(`{"status":"ok"}`))
		})
	}

	http.ListenAndServe(":8080", r)

}
