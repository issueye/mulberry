package pages

import "carambola/host/global"

func WriteLog(log string) {
	if global.Logger != nil {
		global.Logger.Info(log)
	}

	global.MsgChannel <- log
}
