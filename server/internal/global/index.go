package global

import (
	"io/fs"
	"mulberry/pkg/code_engine"
	"mulberry/pkg/ipc/client"
	"mulberry/pkg/logger"
	"net/http"

	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/gin-gonic/gin"
	"github.com/nalgeon/redka"
	"gorm.io/gorm"
)

var (
	MsgChannel = make(chan string, 10)
	PubSub     *gochannel.GoChannel
	Logger     *logger.LoggerWrapper
	HttpEngine *gin.Engine
	HttpServer *http.Server
	DB         *gorm.DB
	ClientID   string
	RedkaDB    *redka.DB
	StaticWEB  fs.FS
	DBMap      = make(map[string]*gorm.DB)
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

// client
var (
	CodeEngine *code_engine.Core
	IpcClient  *client.Client
)

var (
	APP_NAME = "桑葚网关"
	VERSION  = "v1.0.0.2"
)
