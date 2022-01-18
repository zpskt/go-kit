package util

import (
	"context"
	"net/http"
)

func MyErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	w.Header().Set("Content-type", contentType) //设置请求头
	w.WriteHeader(429)                          //写入返回码
	w.Write(body)
}
