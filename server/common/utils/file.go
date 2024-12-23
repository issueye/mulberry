package utils

import (
	"fmt"
	"os"
)

// 获取程序运行目录
func GetWorkDir() string {
	pwd, _ := os.Getwd()
	return pwd
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
	return false, err
}

func IsExistsCreatePath(path, name string) string {
	p := fmt.Sprintf("%s/%s", path, name)
	exists, err := PathExists(p)
	if err != nil {
		panic(err.Error())
	}

	if !exists {
		panic("创建【config】文件夹失败")
	}

	return p
}
