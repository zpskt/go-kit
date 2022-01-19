package main

import (
	"consul-client-v1/util"
	"fmt"
	"log"
)

func main() {
	res, err := util.GetUser()
	if err != nil {
		log.Fatal()
	}
	fmt.Println(res)
}
