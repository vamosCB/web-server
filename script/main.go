package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var root = `D:\迅雷下载\`

// 获取文件信息
func getFileInfo(path string, fp os.FileInfo, err error) error {
	if path == "" {
		return errors.New("get empty path")
	}
	if fp == nil {
		return errors.New("fp is nil")
	}
	if fp.IsDir() {
		fmt.Println(fp.Name() + "is Dir")
		return nil
	}
	if strings.Contains(fp.Name(), "阳光电影www.ygdy8.com.") {
		fmt.Println(fp.Name())
		fileName := fp.Name()
		newName := strings.TrimLeft(fileName, "阳光电影www.ygdy8.com.")
		fmt.Println(newName)
		if err := os.Rename(root+fileName, root+newName); err != nil {
			fmt.Printf("\n %c[1;40;32m%s%c[0m\n\n", 0x1B, err.Error(), 0x1B)
		}
	}
	return nil
}

func main() {
	if err := filepath.Walk(root, getFileInfo); err != nil {
		fmt.Println("!!!!!!!!!!")
		fmt.Println(err)
	}
}
