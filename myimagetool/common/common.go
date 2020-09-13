package common

type ImageInfo struct {
	ImageID    string `orm:"column(uid);pk"`
	Time       string
	Srcpath    string
	Dstpath    string
	Procstatus int
	Method     string
}
