package service

import (
	"errors"
	"strings"
)

// Service constants
const (
	StrMaxSize = 1024
)

// Service errors
var (
	ErrMaxSize = errors.New("maximum size of 1024 bytes exceeded")

	ErrStrValue = errors.New("maximum size of 1024 bytes exceeded")
)

//定义传入参数和返回参数的数据结构
type StringRequest struct {
	A string
	B string
}

//定义服务对象
type Service interface {
	// 字符串拼接函数
	Concat(req StringRequest, ret *string) error

	// 字符串差异函数
	Diff(req StringRequest, ret *string) error
}

//ArithmeticService implement Service interface
//定义结构体
type StringService struct {
}

//实现Service接口
func (s StringService) Concat(req StringRequest, ret *string) error {
	// test for length overflow
	if len(req.A)+len(req.B) > StrMaxSize {
		*ret = ""
		return ErrMaxSize
	}
	*ret = req.A + req.B
	return nil
}

func (s StringService) Diff(req StringRequest, ret *string) error {
	if len(req.A) < 1 || len(req.B) < 1 {
		*ret = ""
		return nil
	}
	res := ""
	if len(req.A) >= len(req.B) {
		for _, char := range req.B {
			if strings.Contains(req.A, string(char)) {
				res = res + string(char)
			}
		}
	} else {
		for _, char := range req.A {
			if strings.Contains(req.B, string(char)) {
				res = res + string(char)
			}
		}
	}
	*ret = res
	return nil
}

// ServiceMiddleware define service middleware
type ServiceMiddleware func(Service) Service
