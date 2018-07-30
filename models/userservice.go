package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

//用户service 用户对象
type UserService struct {
	usersModel *UsersModel
	userObj    *User
}
type User struct {
	Uid         int
	Email       string
	CreatedAt   string
	UpdatedAt   string
	LoginTime   int64
	CookieToken string
	LoginToken  string
}

var userserviceInstance *UserService

//单例
func UserServiceInstance() *UserService {
	if userserviceInstance == nil {
		um := GetUsersModel()
		userserviceInstance = &UserService{usersModel: um}
	}
	return userserviceInstance
}

//检查登录状态
func (this *UserService) CheckLoginStatus() bool {
	return true
}

//首次登录 验证密码
func (this *UserService) CheckPassword(email string, pwd string) (*User, error) {
	if len(email) == 0 || len(pwd) == 0 {
		return nil, errors.New("email or password can't be empty")
	}
	inputPwd := this.GeneratePassword(pwd)
	user, err := usersModel.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not exist")
	}
	if user.Password != inputPwd {
		return nil, errors.New("wrong password ")
	} else {
		return this.InitUserById(user.Uid), nil
	}
}

//根据uid初始化用户
func (this *UserService) InitUserById(uid int) *User {
	userdao, err := this.usersModel.GetUserById(uid)
	if err != nil {
		beego.Error(err)
		panic(err)
	}
	if userdao == nil {
		return nil
	}
	loginTime := time.Now().Unix()
	token := this.GenerateLoginToken(userdao.Uid, loginTime, userdao.Password)
	ctoken := this.GenerateCookieToken(userdao.Uid)

	return &User{
		Uid:         userdao.Uid,
		Email:       userdao.Email,
		CreatedAt:   userdao.Created_at,
		UpdatedAt:   userdao.Updated_at,
		LoginToken:  token,
		CookieToken: ctoken,
		LoginTime:   loginTime,
	}
}

//生成/获取 Logintoken
func (this *UserService) GenerateLoginToken(uid int, time int64, pwd string) string {
	str := []byte(strconv.Itoa(uid) + strconv.FormatInt(time, 10) + pwd)
	return fmt.Sprintf("%x", md5.Sum(str))
}

//生成cookietoken
func (this *UserService) GenerateCookieToken(uid int) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(uid)+"3arrows")))
}

//生成加密后的密码
func (this *UserService) GeneratePassword(pwd string) string {
	//md5(md5($pwd."3aarroowwss"));
	str := pwd + "3aarroowwss"
	str1 := fmt.Sprintf("%x", md5.Sum([]byte(str)))
	str2 := fmt.Sprintf("%x", md5.Sum([]byte(str1)))
	return str2
}
