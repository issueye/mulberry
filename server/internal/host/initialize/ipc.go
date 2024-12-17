package initialize

import (
	"context"
	"mulberry/internal/host/ipc"
	"mulberry/pkg/ipc/grpc/pb"
	"mulberry/pkg/ipc/server"

	grpc "google.golang.org/grpc"
)

func InitIpc(ctx context.Context) {
	RunIpcServer(ctx)
}

func RunIpcServer(ctx context.Context) {
	server.NewServer(true, &ipc.HostHelper{}, nil, func(s *grpc.Server) {
		pb.RegisterRDBHelperServer(s, &ipc.RDBHelper{})
	})
}
