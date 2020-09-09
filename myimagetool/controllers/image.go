package controllers

import (
	"path"
	"strings"

	"github.com/astaxie/beego"
)

type UploadController struct {
	beego.Controller
}

func (this *UploadController) Get() {

	this.TplName = "upload.html"
}

func (this *UploadController) Post() {
	file, information, err := this.GetFile("file") //返回文件，文件信息头，错误信息
	if err != nil {
		this.Ctx.WriteString("File retrieval failure")
		return
	} else {
		filename := information.Filename
		picture := strings.Split(filename, ".")            //读取到字符串，并以.符号分隔开
		layout := strings.ToLower(picture[len(picture)-1]) //把字母字符转换成小写，非字母字符不做出处理,返回此字符串转换为小写形式的副本。

		if layout != "jpg" && layout != "png" && layout != "gif" {
			this.Ctx.WriteString("请上传符合格式的图片（png、jpg、gif）")
			return //结束整个程序，不执行保存文件
		}

		err = this.SaveToFile("file", path.Join("static/upload", filename))
		if err != nil {
			this.Ctx.WriteString("File upload failed！")
		} else {
			this.Ctx.WriteString("File upload succeed!")
		}
	}

	defer file.Close() //关闭上传的文件，否则出现零食文件不清除的情况
	this.TplName = "upload.html"
}
