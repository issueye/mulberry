package home

import (
	"carambola/host/common/config"
	"carambola/host/global"
	"carambola/host/initialize"
	"carambola/host/pages"
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/ying32/govcl/vcl"
)

// 初始化数据
func (f *TFrmHome) InitData() {

}

func (f *TFrmHome) EventMonitor(ctx context.Context) {
	go func(c context.Context) {
		for {
			select {
			case msg := <-global.MsgChannel:
				f.addLog(msg)

			case <-c.Done():
				return
			}
		}
	}(ctx)
}

func (f *TFrmHome) addLog(msg string) {
	vcl.ThreadSync(func() {
		f.Mmo_run_log.Lines().Add(fmt.Sprintf("%s %s", time.Now().Format("2006-01-02 15:04:05"), msg))
	})
}

func (f *TFrmHome) StartServer() {
	f.serverRunCtx, f.serverRunCancel = context.WithCancel(context.Background())
	initialize.RunServer(f.ctx)
	f.OnBtn_clear_logClick(nil)
	pages.WriteLog("启动服务")
}

func (f *TFrmHome) StopServer() {
	defer f.serverRunCancel()

	initialize.StopServer()
	pages.WriteLog("停止服务")
	// 如果服务已经停止，就强制 GC
	runtime.GC()
}

func (f *TFrmHome) ShowRunInfo() {
	f.Lbl_name.SetCaption("名称：" + global.APP_NAME)
	port := config.GetParam(config.SERVER, "http-port", 6677).Int()
	f.Lbl_port.SetCaption(fmt.Sprintf("端口：%d", port))
	f.StartTime = time.Now()

	item_pid := f.StatusBar.Panels().Items(1)
	item_pid.SetText(fmt.Sprintf("PID：%d", os.Getpid()))
}