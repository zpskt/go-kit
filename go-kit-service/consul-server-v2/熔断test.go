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
		time.Sleep(time.Second * 10)
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
		Timeout:                3000,                  //设置延时参数
		MaxConcurrentRequests:  5,                     //控制最大并发数为5，并且在一个统计窗口内处理的请求数量达到阈值会调用我们传入的降级回调函数
		RequestVolumeThreshold: 5,                     //判断熔断的最少请求数，默认是5；只有在一个统计窗口内处理的请求数量达到这个阈值，才会进行熔断与否的判断
		ErrorPercentThreshold:  5,                     //判断熔断的阈值，默认值5，表示在一个统计窗口内有50%的请求处理失败，比如有20个请求有10个以上失败了会触发熔断器短路直接熔断服务
		SleepWindow:            int(time.Second * 10), //熔断器短路多久以后开始尝试是否恢复，这里设置的是10
	}
	hystrix.ConfigureCommand("get_prod", configA) //设置name是get_prod的配置参数为configA
	//返回值有三个，第一个是熔断器指针,第二个是bool表示是否能够取到，第三个是error
	c, _, _ := hystrix.GetCircuit("get_prod")

	resultChan := make(chan Product, 1)

	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()

			// hystrix.Do同步执行。有三个参数：name名字，func1要调用的服务，func2超时后执行代码
			//hystrix.Go是异步执行。有三个参数：name名字，func1要调用的服务，func2超时后执行代码
			errs := hystrix.Do("get_prod", func() error {
				p, _ := getProduct() //这里会随机延迟0-4秒
				resultChan <- p
				return nil
			}, func(e error) error { //这里返回的error在回调中可以获取到，也就是下面的er变量
				//fmt.Println(e)
				//推荐商品,如果这里的err不是nil,那么就会忘errs中写入这个err，下面的select就可以监控到
				rcp, err := RecProduct() //这里返回的error在回调中可以获取到，也就是下面的e变量
				resultChan <- rcp
				return err
			})
			//select {
			//case getProd := <-resultChan:
			//	fmt.Println(getProd)
			//case err := <-errs: //使用hystrix.Go时返回值是chan error各个协程的错误都放到errs中
			//	fmt.Println(err, 1)
			//}
			if errs != nil { //这里errs是error接口，但是使用hystrix.Go异步执行时返回值是chan error各个协程的错误都放到errs中
				fmt.Println(errs)
			} else {
				select {
				case prod := <-resultChan:
					//fmt.Println("hello")
					fmt.Println(prod)
				}
			}
			fmt.Println(c.IsOpen())       //判断熔断器是否打开，一旦打开所有的请求都会走fallback
			fmt.Println(c.AllowRequest()) //判断是否允许请求服务，一旦打开
		}()
	}
	wg.Wait()
}
