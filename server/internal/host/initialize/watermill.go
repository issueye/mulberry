package initialize

import (
	"fmt"
	"mulberry/internal/global"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

func InitEventBus() {
	pubSub := gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false),
	)

	global.PubSub = pubSub
}

func FreeEventBus() {
	if global.PubSub == nil {
		return
	}

	// 处理事件总线
	global.WriteLog("事件总线关闭")
	err := global.PubSub.Close()
	if err != nil {
		global.MsgChannel <- err.Error()
		global.WriteLog(fmt.Sprintf("事件总线关闭失败 %s", err.Error()))
	} else {
		global.WriteLog("事件总线关闭成功")
	}
}
