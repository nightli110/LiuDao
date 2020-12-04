package controllers

import (
	"fmt"
	"path"
	"strings"

	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (c *UploadController) Get() {
	beego.ReadFromRequest(&c.Controller)

	c.TplName = "index.html"
}

func (c *UploadController) Post() {

	file, information, err := c.GetFile("file")
	defer file.Close()
	if err != nil {
		c.Ctx.WriteString("File retrieval failure")
		return
	} else {
		filename := information.Filename
		picture := strings.Split(filename, ".")
		layout := strings.ToLower(picture[len(picture)-1])

		if layout != "jpg" && layout != "png" && layout != "gif" && layout != "jpeg" {
			c.Ctx.WriteString("请上传符合格式的图片（png、jpg、gif）")
			return //结束整个程序，不执行保存文件
		}

		err = c.SaveToFile("file", path.Join("static/upload", filename))
		fmt.Printf(filename)
		if err != nil {
			c.Ctx.WriteString("File upload failed！")
		} else {
			c.Ctx.WriteString("File upload succeed!")
		}
	}
	return
}
