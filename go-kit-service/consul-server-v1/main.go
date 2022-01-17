package main

import (
	"flag"
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	mymux "github.com/gorilla/mux" //第三方路由
	. "gomicro/Services"
	"gomicro/util"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func main() {

	name := flag.String("name", "", "服务名")
	port := flag.Int("p", 0, "服务端口")
	flag.Parse()
	if *name == "" {
		log.Fatal("请指定服务名")
	}
	if *port == 0 {
		log.Fatal("请指定端口")
	}
	util.SetServiceNameAndPort(*name, *port) //设置服务名和端口

	fmt.Println(*name)
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

	errChan := make(chan error)
	go func() {
		util.RegService() //调用注册服务程序
		err := http.ListenAndServe(":"+strconv.Itoa(*port), r)
		if err != nil {
			log.Println(err)
			errChan <- err
		}
	}()
	go func() { //监听关闭信号，ctrl+c 等
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sigChan)
	}()
	getErr := <-errChan //只要报错 或者service关闭阻塞在这里的会进行下去
	util.UnRegService()
	log.Println(getErr)

}
