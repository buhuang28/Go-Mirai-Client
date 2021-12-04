package handler

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFileByte(fileName string) []byte {
	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		return nil
	}
	return b
}

func WriteFile(fileName string, content []byte) bool {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0777)
	defer func() {
		f.Close()
	}()
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else {
		write, e := f.Write(content)
		if e == nil && write > 0 {
			return true
		}
	}
	return false
}

func ReadDir(path string) []os.FileInfo {
	FileInfo, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("读取 img 文件夹出错")
		return nil
	}
	if FileInfo == nil || len(FileInfo) == 0 {
		return nil
	}
	return FileInfo
}
