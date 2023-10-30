// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userrpc

import (
	"context"

	"TokTik/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Data              = pb.Data
	GenerateTokenReq  = pb.GenerateTokenReq
	GenerateTokenResp = pb.GenerateTokenResp
	GetFansReq        = pb.GetFansReq
	GetFansResp       = pb.GetFansResp
	GetFollowingsReq  = pb.GetFollowingsReq
	GetFollowingsResp = pb.GetFollowingsResp
	GetUserInfoReq    = pb.GetUserInfoReq
	GetUserInfoResp   = pb.GetUserInfoResp
	LoginReq          = pb.LoginReq
	LoginResp         = pb.LoginResp
	RegisterReq       = pb.RegisterReq
	RegisterResp      = pb.RegisterResp
	User              = pb.User

	Userrpc interface {
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
		GetFollowings(ctx context.Context, in *GetFollowingsReq, opts ...grpc.CallOption) (*GetFollowingsResp, error)
		GetFans(ctx context.Context, in *GetFansReq, opts ...grpc.CallOption) (*GetFansResp, error)
	}

	defaultUserrpc struct {
		cli zrpc.Client
	}
)

func NewUserrpc(cli zrpc.Client) Userrpc {
	return &defaultUserrpc{
		cli: cli,
	}
}

func (m *defaultUserrpc) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUserrpcClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserrpc) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUserrpcClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserrpc) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewUserrpcClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultUserrpc) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewUserrpcClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUserrpc) GetFollowings(ctx context.Context, in *GetFollowingsReq, opts ...grpc.CallOption) (*GetFollowingsResp, error) {
	client := pb.NewUserrpcClient(m.cli.Conn())
	return client.GetFollowings(ctx, in, opts...)
}

func (m *defaultUserrpc) GetFans(ctx context.Context, in *GetFansReq, opts ...grpc.CallOption) (*GetFansResp, error) {
	client := pb.NewUserrpcClient(m.cli.Conn())
	return client.GetFans(ctx, in, opts...)
}
