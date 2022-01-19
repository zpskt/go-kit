package Services

import (
	"context"
	"encoding/json"
	"errors"
	mymux "github.com/gorilla/mux" //第三方路由
	"gomicro/util"
	"net/http"
	"strconv"
)

//怎么去传？传什么，做了这些事情
//当有外部请求过来时，对request解码
func DecodeUserRequest(c context.Context, r *http.Request) (interface{}, error) {
	//http://localhost:xxx/?uid=101
	//判断是不是url来的参数，获取指定参数
	/* 如果用的路由路径，那么就不能用这个了
	if r.URL.Query().Get("uid") != "" {
		//用strconv.Atoi进行转化
		uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
		//这里就和endpoint里面的UserRequest里面对应了
		return UserRequest{
			Uid: uid,
		}, nil
	}*/
	vars := mymux.Vars(r)
	if uid, ok := vars["uid"]; ok { //？？？
		//用strconv.Atoi进行转化
		uid, _ := strconv.Atoi(uid)
		//这里就和endpoint里面的UserRequest里面对应了

		//请求必须携带token过来，如果找不到这里返回空字符串，
		//因为request访问的先后顺序是先DecodeUserRequest，
		//再EncodeUserResponse再到我们的EndPoint，
		//所以这里就已经给我们的request结构体存入了Token，
		//那么我们EndPoint里面的request类型断言成UserRequest结构体实例后里面就有Token了
		return UserRequest{Uid: uid, Method: r.Method, Token: r.URL.Query().Get("token")}, nil

	}
	return nil, errors.New("参数错误")
}
func EncodeUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	//这一步把uid的类型转化成application/json格式，开发者工具里面就可以看到
	w.Header().Set("Content-type", "application/json")
	//因为系统异构，所以要变成大家都认识的形式，json格式
	return json.NewEncoder(w).Encode(response)
}
func MyErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	w.Header().Set("Content-type", contentType) //设置请求头
	if myerr, ok := err.(*util.MyError); ok {
		w.WriteHeader(myerr.Code)
		w.Write(body)
	} else {
		w.WriteHeader(500)
		w.Write(body)
	}

}
