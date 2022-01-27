package main

import (
	d "ch13-seckill/pkg/discover"
	"log"
	"os"
)

var ConsulService d.DiscoveryClient

func main() {
	//bootstrap.HttpConfig.Host = "localhost"
	//bootstrap.HttpConfig.Port = "9030"

	Logger := log.New(os.Stderr, "", log.LstdFlags)
	//instanceId := bootstrap.DiscoverConfig.InstanceId
	if !ConsulService.Register("asdf", "127.0.0.1", "/health",
		"9030", "sk-admin",
		0,
		map[string]string{
			"rpcPort": "1111",
		}, nil, Logger) {
		//Logger.Printf("register service %s failed.", bootstrap.DiscoverConfig.ServiceName)
		Logger.Printf("register service %s failed.", "abcd")
		// 注册失败，服务启动失败
		panic(0)
	}
}
