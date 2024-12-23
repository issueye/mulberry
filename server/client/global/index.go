package global

import (
	"carambola/common/code_engine"
	"carambola/common/ipc/client"
	"carambola/common/logger"

	"gorm.io/gorm"
)

var (
	MsgChannel = make(chan string, 10)
	Logger     *logger.LoggerWrapper
	DB         *gorm.DB
	ClientID   string
	DBMap      = make(map[string]*gorm.DB)
	CodeEngine *code_engine.Core
	IpcClient  *client.Client
)

var (
	APP_NAME = "杨桃桃服务"
	VERSION  = "v1.0.0.1"
)

const (
	TOPIC_CONSOLE_LOG = "TOPIC_CONSOLE_LOG"
	ROOT_PATH         = "root"
	DEFAULT_PWD       = "123456"
	DB_Key            = "data_base:info"
)
