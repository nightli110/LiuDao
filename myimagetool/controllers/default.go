package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Upload() {
	f, h, _ := c.GetFile("myfile") //获取上传的文件
	path := "./" + h.Filename
	f.Close()                    //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	c.SaveToFile("myfile", path) //存文件
}
