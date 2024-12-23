package register

import (
	"carambola/client/global"
	"context"
	"encoding/base64"
	"os"
	"os/exec"
)

// FileToString 将pdf文件转换为base64字符串
func FileToBase64(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(data)
}

func RunApp(path string, args ...string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	out, err := exec.CommandContext(ctx, path, args...).Output()
	if err != nil {
		global.Logger.Sugar().Errorf("执行程序[%s]失败：%s", err.Error())
		return
	}

	global.Logger.Sugar().Infof("执行程序[%s]成功：%s", path, string(out))
}
