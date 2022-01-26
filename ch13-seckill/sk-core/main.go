package main

import (
	"ch13-seckill/sk-core/setup"
)

func main() {
	//fmt.Println("hello,world")
	setup.InitZk()
	setup.InitRedis()
	setup.RunService()

}
