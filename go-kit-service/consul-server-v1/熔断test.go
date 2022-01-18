package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"math/rand"
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
		time.Sleep(time.Second * 3)
	}
	return Product{
		ID:    101,
		Title: "Golang从入门到精通",
		Price: 12,
	}, nil
}

func main() {
	//设置随机因子
	rand.Seed(time.Now().UnixNano())
	configA := hystrix.CommandConfig{ //hystrix.CommandConfig 修改配置文件
		Timeout: 4000, //设置延时参数
	}
	hystrix.ConfigureCommand("get_prod", configA) //设置name是get_prod的配置参数为configA
	for {
		err := hystrix.Do("get_prod", func() error {
			p, _ := getProduct()
			fmt.Println(p)
			return nil
		}, nil)
		if err != nil {
			fmt.Println(err)
		}

	}
}
