package main

import (
	"consul-client-v1/util"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"log"
)

func main() {

	configA := hystrix.CommandConfig{ //hystrix.CommandConfig 修改配置文件
		Timeout:                2000, //设置延时参数
		MaxConcurrentRequests:  5,    //控制最大并发数为5，并且在一个统计窗口内处理的请求数量达到阈值会调用我们传入的降级回调函数
		RequestVolumeThreshold: 3,    //判断熔断的最少请求数，默认是5；只有在一个统计窗口内处理的请求数量达到这个阈值，才会进行熔断与否的判断
		ErrorPercentThreshold:  5,    //判断熔断的阈值，默认值5，表示在一个统计窗口内有50%的请求处理失败，比如有20个请求有10个以上失败了会触发熔断器短路直接熔断服务
		//SleepWindow:            int(time.Second * 10), //熔断器短路多久以后开始尝试是否恢复，这里设置的是10
	}
	hystrix.ConfigureCommand("getuser", configA) //设置name是get_prod的配置参数为configA
	//返回值有三个，第一个是熔断器指针,第二个是bool表示是否能够取到，第三个是error
	err := hystrix.Do("getuser", func() error {
		res, err := util.GetUser()
		fmt.Println(res)
		return err
	}, func(e error) error {
		fmt.Println("降级用户")
		return e
	})

	if err != nil {
		log.Fatal(err)
		//fmt.Println("have err")
	}
}
