package main

import (
	"myimagetool/models"
	_ "myimagetool/routers"

	"github.com/astaxie/beego"
)

func main() {

	models.BeegoInit()
	beego.Run()
}
