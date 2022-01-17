package Services

import "errors"

//业务类，我们要干什么
//1.写接口
type IUserService interface {
	GetName(userid int) string
	DelUser(userid int) error
}

//2.定义实现类
type UserService struct {
}

//3.继承接口，写个名字一摸一样的方法
func (this UserService) GetName(userid int) string {
	//这里可以用数据库
	if userid == 101 {
		return "zp"
	}
	return "guest"

}
func (this UserService) DelUser(userid int) error {
	//这里可以用数据库
	if userid == 101 {
		return errors.New("无权限")
	}
	return nil

}
