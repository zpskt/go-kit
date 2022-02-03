package setup

import (
	conf "ch13-seckill/pkg/config"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"time"
)

//初始化zookeeper
func InitZk() {
	var hosts = []string{"localhost:2181"}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		fmt.Println(err)
		return
	}
	conf.Zk.ZkConn = conn
	//设置关键字
	conf.Zk.SecProductKey = "/product"
}
