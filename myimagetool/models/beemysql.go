package models

import (
	"myimagetool/common"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL) //注册驱动
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("image", "mysql", "root:123456{Ljj}@/orm_test?charset=utf8", maxIdle, maxConn)

	orm.RegisterModel(new(common.ImageInfo))
}

func insert(info common.ImageInfo) {

}
