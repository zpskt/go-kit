package util

import (
	"context"
	"net/http"
)

func MyErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	w.Header().Set("Content-type", contentType) //设置请求头
	if myerr, ok := err.(*MyError); ok {        //通过类型断言判断当前error的类型，走相应的处理
		w.WriteHeader(myerr.Code)
		w.Write(body)
	} else {
		w.WriteHeader(500)
		w.Write(body)
	}
}
