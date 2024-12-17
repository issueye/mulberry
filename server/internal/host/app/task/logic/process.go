package logic

import (
	"fmt"
	"mulberry/internal/config"
	"mulberry/internal/global"
	"mulberry/internal/host/app/task/service"
	"mulberry/internal/utils/helper_cmd"
	"time"
)

var (
	ClientProcessMap = make(map[string]*ClientProcess)
)

type ClientProcess struct {
	ClientAuthId string
	Process      *helper_cmd.RunResult
}

func RunProcess(id string) error {
	srv := service.NewClient(global.DB, false)
	info, err := srv.GetByField("client_auth_id", id)
	if err != nil {
		return err
	}

	clientPath := config.GetParam(config.SERVER, "client-path", "").String()
	param := fmt.Sprintf("-n=%s", info.ClientAuthId)
	rtn, err := helper_cmd.Run(0)(true, clientPath, param)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-rtn.Ctx.Done():
				return
			case msg := <-rtn.Msg:
				global.Logger.Sugar().Debugf("客户端[%s-%s] 输出: %s", info.Name, info.ClientAuthId, msg)
			}
		}
	}()

	ClientProcessMap[id] = &ClientProcess{
		ClientAuthId: id,
		Process:      rtn,
	}

	now := time.Now()
	info.StartAt = &now
	info.Pid = rtn.Pid

	global.Logger.Sugar().Debugf("客户端[%s-%s] 启动成功，PID: %d", info.Name, info.ClientAuthId, rtn.Pid)

	err = srv.Update(info.ID, info)
	if err != nil {
		return err
	}

	return nil
}

func CloseProcess(id string) error {
	process, ok := ClientProcessMap[id]
	if !ok {
		return nil
	}

	defer process.Process.Cancel()

	process.Process.Process.Kill()
	process.Process.Process.Release()

	delete(ClientProcessMap, id)
	return nil
}
