package main

import (
	"ch13-seckill/pkg/bootstrap"
	conf "ch13-seckill/pkg/config"
	"ch13-seckill/pkg/mysql"
	"ch13-seckill/sk-admin/setup"
	"fmt"
)

//秒杀哦管理系统，创建删除秒杀活动，商品配置
func main() {
	//初始化mysql
	mysql.InitMysql(conf.MysqlConfig.Host, conf.MysqlConfig.Port, conf.MysqlConfig.User, conf.MysqlConfig.Pwd, conf.MysqlConfig.Db) // conf.MysqlConfig.Db
	fmt.Println("mysql初始化成功")

	//setup.InitEtcd()

	setup.InitZk()
	fmt.Println("zk初始化成功")

	setup.InitServer(bootstrap.HttpConfig.Host, bootstrap.HttpConfig.Port)
	fmt.Println("服务初始化成功")

}
