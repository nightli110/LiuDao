package common

type Image struct {
	ImageID    string `orm:"column(imageid);pk"`
	Time       string
	Srcpath    string
	Dstpath    string
	Procstatus int
	Method     string
}

var Updateimagedir = "static/upload"
