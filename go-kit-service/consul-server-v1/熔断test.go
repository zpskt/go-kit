package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"math/rand"
	"sync"
	"time"
)

type Product struct {
	ID    int
	Title string
	Price int
}

//假设这就是远程调用商品信息
func getProduct() (Product, error) {
	r := rand.Intn(10) //随机一个数
	if r < 6 {         //模拟api卡顿和超时效果
		//time.Sleep(time.Second * 3)
	}
	return Product{
		ID:    101,
		Title: "Golang从入门到精通",
		Price: 12,
	}, nil
}

//推荐商品
func RecProduct() (Product, error) {
	return Product{
		ID:    999,
		Title: "推荐商品",
		Price: 120,
	}, nil

}
func main() {
	//设置随机因子
	rand.Seed(time.Now().UnixNano())
	configA := hystrix.CommandConfig{ //hystrix.CommandConfig 修改配置文件
		Timeout:               2000, //设置延时参数
		MaxConcurrentRequests: 5,    //控制最大并发数为5，如果超过5会调用我们传入的回调函数降级
	}
	hystrix.ConfigureCommand("get_prod", configA) //设置name是get_prod的配置参数为configA
	resultChan := make(chan Product, 1)

	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		go (func() {

			wg.Add(1)
			defer wg.Done()

			// hystrix.Do同步执行。有三个参数：name名字，func1要调用的服务，func2超时后执行代码
			//hystrix.Go是异步执行。有三个参数：name名字，func1要调用的服务，func2超时后执行代码
			errs := hystrix.Go("get_prod", func() error {
				p, _ := getProduct()
				resultChan <- p
				return nil
			}, func(err error) error {
				fmt.Println(err)
				//推荐商品,如果这里的err不是nil,那么就会忘errs中写入这个err，下面的select就可以监控到
				rcp, err := RecProduct()
				resultChan <- rcp
				return err
			})
			select {
			case getProd := <-resultChan:
				fmt.Println(getProd)
			case err := <-errs: //使用hystrix.Go时返回值是chan error各个协程的错误都放到errs中
				fmt.Println(err, 1)
			}
		})()
	}
	wg.Wait()
}
