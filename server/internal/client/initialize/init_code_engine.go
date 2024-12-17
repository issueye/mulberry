package initialize

import (
	"mulberry/internal/global"
	"mulberry/pkg/code_engine"
	"path/filepath"
)

func InitCodeEngine() {
	logPath := filepath.Join(global.ROOT_PATH, "logs")
	global.CodeEngine = code_engine.NewCore(
		code_engine.OptionLog(logPath, global.Logger.Named("code_engine")),
	)
	scriptPath := filepath.Join(global.ROOT_PATH, "scripts")
	global.CodeEngine.SetGlobalPath(scriptPath)
}
