package controllers

import "fmt"

type ImageUploadController struct {
	MainController
}

func (c *ImageUploadController) Download() {
	filename := c.GetString("filename")
	fmt.Println("filename666", filename)
	c.Ctx.Output.Download("./upload/demo.txt")
}

func (c *ImageUploadController) Upload() {
	f, h, _ := c.GetFile("myfile")
	path := "./upload/" + h.Filename
	f.Close()
	c.SaveToFile("myfile", path)
}
