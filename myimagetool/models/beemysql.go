package main

import (
	"fmt"
	"myimagetool/common"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL) //注册驱动
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("image", "mysql", "root:123456{Ljj}@tcp(localhost:3306)?charset=utf8", maxIdle, maxConn)

	orm.RegisterModel(new(common.ImageInfo))
}

func insert(info common.ImageInfo) {
	o := orm.NewOrm()

	id, err := o.Insert(&info)

	fmt.Printf("NUM: %d, ERR: %v\n", id, err)
}

type my struct {
	my string
}

func main() {
	myinfo := new(common.ImageInfo)

	myinfo.ImageID = "t"
	myinfo.Time = "r"
	myinfo.Srcpath = "1"
	myinfo.Dstpath = "2"
	myinfo.Procstatus = 1
	insert(*myinfo)
}
