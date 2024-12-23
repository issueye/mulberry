package initialize

import (
	"carambola/common/ipc/grpc/pb"
	"carambola/common/ipc/server"
	"carambola/common/ipc/vars"
	"carambola/host/ipc"
	"context"

	grpc "google.golang.org/grpc"
)

func InitIpc(ctx context.Context) {
	vars.PIPE_NAME = vars.PIPE_NAME + "_carambola"
	RunIpcServer(ctx)
}

func RunIpcServer(ctx context.Context) {
	server.NewServer(true, &ipc.HostHelper{}, nil, func(s *grpc.Server) {
		pb.RegisterRDBHelperServer(s, &ipc.RDBHelper{})
	})
}
