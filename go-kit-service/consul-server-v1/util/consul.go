package util

import (
	"fmt"
	"github.com/google/uuid"
	consulapi "github.com/hashicorp/consul/api"
	"log"
)

var ConsulClient *consulapi.Client //因为客户端都会用所以放在外面
var ServiceID string
var ServiceName string
var ServicePort int

func init() {
	config := consulapi.DefaultConfig()
	config.Address = "localhost:8500"
	client, err := consulapi.NewClient(config) //创建客户端
	if err != nil {
		log.Fatal(err)
	}
	ConsulClient = client
	//因为最终这段代码是在不同的机器上跑的，是分布式的，有好几台机器提供相同的server，
	//所以这里存到consul中的id必须是唯一的，否则只有一台服务器可以注册进去，这里使用uuid保证唯一性。
	ServiceID = "userservice" + uuid.New().String()
}
func SetServiceNameAndPort(name string, port int) {
	ServiceName = name
	ServicePort = port
}
func RegService() {

	//注册信息
	reg := consulapi.AgentServiceRegistration{}
	reg.ID = ServiceID        //不可重复
	reg.Name = ServiceName    //注册service的名字，可以重复
	reg.Address = "localhost" //注册service的ip
	reg.Port = ServicePort    //注册service的端口
	reg.Tags = []string{"primary"}

	//健康检查，心跳
	check := consulapi.AgentServiceCheck{}                                    //创建consul的检查器
	check.Interval = "5s"                                                     //设置consul心跳检查时间间隔
	check.HTTP = fmt.Sprintf("http://%s:%d/health", reg.Address, ServicePort) //设置检查心跳API

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
