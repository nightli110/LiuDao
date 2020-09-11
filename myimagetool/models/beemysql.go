package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL) //注册驱动
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("image", "mysql", "root:root@/orm_test?charset=utf8", maxIdle, maxConn)

}
