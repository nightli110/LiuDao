package common

import (
	"os"
)

type ImageInfo struct {
	ImageID    string
	Time       string
	Srcpath    string
	Dstpath    string
	Procstatus int
	Method     string
}

func IsExists(filepath string) bool {
	if _, err := os.Stat(filepath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
