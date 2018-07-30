package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type BaseModel struct {
	queryBuilder orm.QueryBuilder
	db           orm.Ormer
}

func (this *BaseModel) init() {
	User := beego.AppConfig.String("mysqluser")
	Password := beego.AppConfig.String("mysqlpass")
	Host := beego.AppConfig.String("mysqlhost")
	Db := beego.AppConfig.String("mysqldb")
	ConnectString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", User, Password, Host, Db)
	orm.RegisterDataBase("default", "mysql", ConnectString, 30)
	this.db = orm.NewOrm()
}
