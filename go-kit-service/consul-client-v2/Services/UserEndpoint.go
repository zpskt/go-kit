package Services

//客户端不需要知道怎么实现，就需要知道怎么定义的就可以了

//定义response和request的格式
type UserRequest struct {
	//Uid是你自定义的，想叫什么就叫什么
	Uid    int `json:"uid"`
	Method string
}
type UserResponse struct {
	Result string `json:"result"`
}
