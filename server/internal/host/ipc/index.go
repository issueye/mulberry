package ipc

import (
	"context"
	"mulberry/internal/global"
	"mulberry/pkg/ipc/grpc/pb"
)

var commonHelper map[string]*CommonHelper

func GetCommonHelper(clientID string) *CommonHelper {
	if commonHelper == nil {
		commonHelper = make(map[string]*CommonHelper)
	}

	if commonHelper[clientID] == nil {
		commonHelper[clientID] = &CommonHelper{}
	}

	return commonHelper[clientID]
}

type CommonHelper struct {
	ClientID      string
	CommandStream pb.HostHelper_CommandServer
}

type HostHelper struct{}

// 获取任务列表
func (HostHelper) GetTask(ctx context.Context, req *pb.TaskListRequest) (*pb.TaskResponse, error) {
	// req.ClientID
	return &pb.TaskResponse{}, nil
}

// 管理端给客户端发送命令 启动/停止/运行 任务
func (HostHelper) Command(client *pb.ClientRequest, req pb.HostHelper_CommandServer) error {
	helper := GetCommonHelper(client.ClientID)
	helper.ClientID = client.ClientID
	helper.CommandStream = req
	select {}
}

// 客户端在之后完任务之后，就会向管理端发送任务完成消息
func (HostHelper) TaskComplete(context.Context, *pb.TaskInfo) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

type RDBHelper struct{}

func (RDBHelper) ListLen(ctx context.Context, in *pb.RDBKey) (*pb.RDBLen, error) {
	i, err := global.RedkaDB.List().Len(in.Key)
	return &pb.RDBLen{Key: in.Key, Len: int32(i)}, err
}
func (RDBHelper) ListGet(ctx context.Context, in *pb.RDBKeyGet) (*pb.RDBKV, error) {
	value, err := global.RedkaDB.List().Get(in.Key, int(in.Index))
	return &pb.RDBKV{
		Key:   in.Key,
		Value: value.String(),
	}, err
}
func (RDBHelper) ListPopFront(ctx context.Context, in *pb.RDBKey) (*pb.RDBKV, error) {
	value, err := global.RedkaDB.List().PopFront(in.Key)
	return &pb.RDBKV{
		Key:   in.Key,
		Value: value.String(),
	}, err
}
func (RDBHelper) ListPopBack(ctx context.Context, in *pb.RDBKey) (*pb.RDBKV, error) {
	value, err := global.RedkaDB.List().PopBack(in.Key)
	return &pb.RDBKV{
		Key:   in.Key,
		Value: value.String(),
	}, err
}
func (RDBHelper) ListPushFront(ctx context.Context, in *pb.RDBKV) (*pb.Empty, error) {
	_, err := global.RedkaDB.List().PushFront(in.Key, in.Value)
	return &pb.Empty{}, err
}
func (RDBHelper) ListPushBack(ctx context.Context, in *pb.RDBKV) (*pb.Empty, error) {
	_, err := global.RedkaDB.List().PushBack(in.Key, in.Value)
	return &pb.Empty{}, err
}
func (RDBHelper) ListRemove(ctx context.Context, in *pb.RDBKey) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}
func (RDBHelper) ListRange(ctx context.Context, in *pb.RDBStartEnd) (*pb.RDBKVs, error) {
	list, err := global.RedkaDB.List().Range(in.Key, int(in.Start), int(in.End))
	if err != nil {
		return &pb.RDBKVs{}, err
	}

	var values []string

	for _, v := range list {
		values = append(values, v.String())
	}

	return &pb.RDBKVs{Key: in.Key, Values: values}, nil
}
func (RDBHelper) ListTrim(ctx context.Context, in *pb.RDBStartEnd) (*pb.Empty, error) {
	_, err := global.RedkaDB.List().Trim(in.Key, int(in.Start), int(in.End))
	return &pb.Empty{}, err
}
func (RDBHelper) StrSet(ctx context.Context, in *pb.RDBKV) (*pb.Empty, error) {
	return &pb.Empty{}, global.RedkaDB.Str().Set(in.Key, in.Value)
}
func (RDBHelper) StrGet(ctx context.Context, in *pb.RDBKey) (*pb.RDBKV, error) {
	value, err := global.RedkaDB.Str().Get(in.Key)
	if err != nil {
		return &pb.RDBKV{}, err
	}

	return &pb.RDBKV{
		Key:   in.Key,
		Value: value.String(),
	}, nil
}
func (RDBHelper) StrDelete(ctx context.Context, in *pb.RDBKey) (*pb.Empty, error) {
	return &pb.Empty{}, global.RedkaDB.Str().SetExpires(in.Key, "", 100)
}
