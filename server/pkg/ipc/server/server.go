package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"

	"mulberry/internal/global"
	pb "mulberry/pkg/ipc/grpc/pb" // 替换为你的 proto 生成代码路径
	"mulberry/pkg/ipc/vars"

	"github.com/Microsoft/go-winio"
	"google.golang.org/grpc"
)

type PluginInfo struct {
	Version           string `json:"version"`             // 版本号
	AppName           string `json:"appName"`             // 程序名称
	BuildTime         string `json:"buildTime"`           // 构建时间
	GoVersion         string `json:"goVersion"`           // go 版本
	ClientID          string `json:"client_id"`           // 客户端ID
	State             int32  `json:"state"`               // 状态
	LastHeartbeatTime int64  `json:"last_heartbeat_time"` // 最后心跳时间
}

type PluginState struct {
	ClientID          string `json:"client_id"`           // 客户端ID
	LastHeartbeatTime int64  `json:"last_heartbeat_time"` // 最后心跳时间
	State             int32  `json:"state"`               // 状态
}

type CommonHelpServer struct {
	callback        func(info *PluginInfo) error
	updateHeartbeat func(cookieKey string, lastHeartbeatTime int64)
	states          map[string]*PluginState
}

func (s *CommonHelpServer) CheckPlugin(cookieKey string) (*PluginState, error) {
	state, ok := s.states[cookieKey]
	if !ok {
		return nil, fmt.Errorf("cookie key 不存在")
	}

	return state, nil
}

// 测试网络
func (s *CommonHelpServer) Ping(context.Context, *pb.Empty) (*pb.PublicResponse, error) {
	return &pb.PublicResponse{Message: "pong", Timestamp: time.Now().Unix()}, nil
}

// 获取插件信息
func (s *CommonHelpServer) Register(ctx context.Context, data *pb.InfoRequest) (*pb.PublicResponse, error) {
	pi := &PluginInfo{
		Version:   data.Version,
		AppName:   data.AppName,
		BuildTime: data.BuildTime,
		State:     1,
		ClientID:  data.ClientID,
	}

	fmt.Printf("程序名称：%s，客户端ID：%s，版本号：%s\n", pi.AppName, data.ClientID, pi.Version)
	if pi.ClientID == "" {
		return nil, fmt.Errorf("客户端ID不能为空")
	}

	err := s.callback(pi)
	if err != nil {
		return nil, fmt.Errorf("插件注册失败: %s", err.Error())
	}

	s.states[data.ClientID] = &PluginState{
		ClientID:          data.ClientID,
		LastHeartbeatTime: time.Now().Unix(),
		State:             1,
	}

	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, global.ROOT_PATH)
	return &pb.PublicResponse{Message: "ok", Timestamp: time.Now().Unix(), RootPath: path}, nil
}

// 心跳检测
func (s *CommonHelpServer) Heartbeat(stream pb.CommonHelper_HeartbeatServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break // 客户端已关闭流
		}

		if err != nil {
			return stream.SendAndClose(&pb.PublicResponse{Message: err.Error(), Timestamp: time.Now().Unix()})
		}

		if req.ClientID == "" {
			return stream.SendAndClose(&pb.PublicResponse{Message: "客户端ID不能为空", Timestamp: time.Now().Unix()})
		}

		state, err := s.CheckPlugin(req.ClientID)
		if err != nil {
			return stream.SendAndClose(&pb.PublicResponse{Message: err.Error(), Timestamp: time.Now().Unix()})
		}

		fmt.Printf("收到心跳包，客户端：%s  内存使用：%.2f MB CPU使用: %.2f\n", req.ClientID, req.MemoryUsage/1024.0/1024.0, req.CpuUsage/100.0)
		state.LastHeartbeatTime = time.Now().Unix()
		state.State = 1

		s.updateHeartbeat(req.ClientID, req.Timestamp)
	}

	return nil
}

func NewCommonHelpServer(callback func(*PluginInfo) error, updateHeartbeat func(cookieKey string, lastHeartbeatTime int64)) *CommonHelpServer {
	return &CommonHelpServer{
		callback:        callback,
		updateHeartbeat: updateHeartbeat,
		states:          make(map[string]*PluginState),
	}
}

type Server struct {
	Plugins          map[string]*PluginInfo
	hostHelper       pb.HostHelperServer
	commonHelpServer *CommonHelpServer
	checkPlugin      func(PluginInfo) error
}

func NewServer(isASync bool, hostHelper pb.HostHelperServer, checkPlugin func(PluginInfo) error, reg func(s *grpc.Server)) (server *Server, err error) {
	var (
		lis net.Listener
	)

	lis, err = winio.ListenPipe(vars.GetPipeName(), nil)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer()

	callBack := func(pi *PluginInfo) error {
		_, ok := server.Plugins[pi.ClientID]
		if !ok {
			if checkPlugin != nil {
				err := checkPlugin(*pi)
				if err != nil {
					return err
				}
			}

			server.Plugins[pi.ClientID] = pi
			return nil
		} else {
			return fmt.Errorf("cookie key 已存在")
		}
	}

	updateHeartbeat := func(cookieKey string, lastHeartbeatTime int64) {
		plugin, ok := server.Plugins[cookieKey]
		if !ok {
			return
		}

		plugin.LastHeartbeatTime = lastHeartbeatTime
	}

	commonHelpServer := NewCommonHelpServer(callBack, updateHeartbeat)
	server = &Server{
		Plugins:          make(map[string]*PluginInfo),
		hostHelper:       hostHelper,
		commonHelpServer: commonHelpServer,
		checkPlugin:      checkPlugin,
	}

	pb.RegisterCommonHelperServer(s, commonHelpServer)
	pb.RegisterHostHelperServer(s, hostHelper)
	reg(s)
	fmt.Println("服务端启动...")

	if isASync {
		go func() {
			err = s.Serve(lis)
			if err != nil {
				fmt.Println("grpc server start failed: ", err.Error())
			}
		}()
	} else {
		err = s.Serve(lis)
		if err != nil {
			return nil, err
		}
	}

	return
}
