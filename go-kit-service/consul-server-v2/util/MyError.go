package util

//自定义错误结构体
type MyError struct {
	Code    int
	Message string
}

func NewMyError(code int, msg string) error {
	return &MyError{Code: code, Message: msg}
}

func (this *MyError) Error() string {
	return this.Message
}
