package util

import (
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

var ConsulClient *consulapi.Client //因为客户端都会用所以放在外面

func init() {
	config := consulapi.DefaultConfig()
	config.Address = "localhost:8500"          //这里写你的服务注册地址
	client, err := consulapi.NewClient(config) //创建客户端
	if err != nil {
		log.Fatal(err)
	}
	ConsulClient = client
}
func RegService() {

	//注册信息
	reg := consulapi.AgentServiceRegistration{}
	reg.ID = "userserviceid"    //不可重复
	reg.Name = "userservice"    //注册service的名字，可以重复
	reg.Address = "192.168.1.5" //注册service的ip
	reg.Port = 8080             //注册service的端口
	reg.Tags = []string{"primary"}

	//健康检查，心跳
	check := consulapi.AgentServiceCheck{}        //创建consul的检查器
	check.Interval = "5s"                         //设置consul心跳检查时间间隔
	check.HTTP = "http://192.168.1.5:8080/health" //设置检查使用的url

	reg.Check = &check //传入我们写的check，地址传递
	//完成基本数据

	err := ConsulClient.Agent().ServiceRegister(&reg)
	if err != nil {
		log.Fatal(err)
	}
}
func UnRegService() {
	ConsulClient.Agent().ServiceDeregister("userserviceid")
}
