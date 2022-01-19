package main

import (
	"flag"
	"fmt"
)

func main() {
	//自己写个命令参数
	name := flag.String("name", "", "服务名")
	flag.Parse()
	fmt.Println(*name)
}
