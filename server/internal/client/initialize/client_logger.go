package initialize

import (
	"mulberry/internal/global"
	"mulberry/pkg/logger"
	"path/filepath"
)

func InitLogger() {
	cfg := new(logger.Config)
	cfg.Level = -1
	cfg.Path = filepath.Join(global.IpcClient.RootPath, "logs")
	cfg.Name = "client_log"
	cfg.MaxSize = 100 // MB
	cfg.MaxBackups = 50
	cfg.MaxAge = 100 // days
	cfg.Compress = true

	var err error
	global.Logger, err = logger.NewLoggerWrapper(*cfg)
	if err != nil {
		panic(err)
	}

	global.Logger.Sugar().Info("初始化日志成功")
}
