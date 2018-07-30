package models

import (
	"github.com/astaxie/beego/orm"
)

type UsersModel struct {
	BaseModel
}
type UserDao struct {
	Uid        int
	Email      string
	Password   string
	Created_at string
	Updated_at string
}

var usersModel *UsersModel

func GetUsersModel() *UsersModel {
	if usersModel == nil {
		usersModel = &UsersModel{}
	}
	usersModel.init()
	return usersModel
}

//根据用户名获取用户
func (this *UsersModel) GetUserByEmail(email string) (*UserDao, error) {
	var result *UserDao
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("users").Where("email=?").Limit(1).String()
	err := this.db.Raw(sql, email).QueryRow(&result)

	return result, err
}

//根据id获取用户
func (this *UsersModel) GetUserById(id int) (*UserDao, error) {
	var result *UserDao
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("*").From("users").Where("`uid`= ?").Limit(1).String()
	err := this.db.Raw(sql, id).QueryRow(&result)

	return result, err
}
