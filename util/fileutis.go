package util

import (
	"errors"
	"os"
	"strconv"
)

var fileCnt int = 0

//get the icon path
func GetIconPath(buf []byte) string {
	if len(buf) == 0 {
		return ""
	}
	path := ("photos/" + strconv.Itoa(fileCnt))
	file, err := os.Create(path)
	if err != nil {
		return ""
	}
	fileCnt++
	file.Write(buf)
	file.Close()
	return path
}

func GetImageBytes(path string) ([]byte, error) {
	if len(path) <= 0 {
		return nil, errors.New("invalid path of icon")
	}
	filePath := "E:/hz/workspace/go/src/cst_im/model/"
	file, err := os.Open(filePath + path)
	if err != nil {
		return nil, err
	}
	fileBuf := make([]byte, 0)
	buf := make([]byte, 1024*4)
	for {
		n, _ := file.Read(buf)
		if n <= 0 {
			break
		}
		fileBuf = append(fileBuf, buf[:n]...)
	}
	file.Close()
	return fileBuf, nil
}
