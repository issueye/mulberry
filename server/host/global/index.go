package global

import (
	"carambola/common/logger"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nalgeon/redka"
	"gorm.io/gorm"
)

var (
	MsgChannel = make(chan string, 10)
	Logger     *logger.LoggerWrapper
	HttpEngine *gin.Engine
	HttpServer *http.Server
	DB         *gorm.DB
	RedkaDB    *redka.DB
	StaticWEB  fs.FS
)

const (
	TOPIC_CONSOLE_LOG = "TOPIC_CONSOLE_LOG"
	ROOT_PATH         = "root"
	DEFAULT_PWD       = "123456"
	DB_Key            = "data_base:info"
)

func WriteLog(msg string) {
	MsgChannel <- msg
}

var (
	APP_NAME = "杨桃桃服务"
	VERSION  = "v1.0.0.1"
)
