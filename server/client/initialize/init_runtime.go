package initialize

import (
	"carambola/client/global"
	"fmt"
	"os"
)

func InitRuntime() {
	// 检查本地是否存在runtime文件夹
	rtPath := global.IpcClient.RootPath
	// fmt.Println("rtPath", rtPath)
	isExistsCreatePath(rtPath, "data")
	isExistsCreatePath(rtPath, "config")
	isExistsCreatePath(rtPath, "scripts")
	isExistsCreatePath(rtPath, "logs")
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return false, err
		} else {
			return true, nil
		}
	}
	return false, err
}

func isExistsCreatePath(path, name string) string {
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
