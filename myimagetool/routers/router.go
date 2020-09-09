package routers

import (
	"myimagetool/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/image/upload", &controllers.ImageUploadController{}, "post:Upload")
}
