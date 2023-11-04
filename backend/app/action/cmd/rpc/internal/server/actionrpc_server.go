// Code generated by goctl. DO NOT EDIT.
// Source: action.proto

package server

import (
	"context"

	"TokTik/app/action/cmd/rpc/internal/logic"
	"TokTik/app/action/cmd/rpc/internal/svc"
	"TokTik/app/action/cmd/rpc/pb"
)

type ActionrpcServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedActionrpcServer
}

func NewActionrpcServer(svcCtx *svc.ServiceContext) *ActionrpcServer {
	return &ActionrpcServer{
		svcCtx: svcCtx,
	}
}

func (s *ActionrpcServer) Follow(ctx context.Context, in *pb.FollowReq) (*pb.FollowResp, error) {
	l := logic.NewFollowLogic(ctx, s.svcCtx)
	return l.Follow(in)
}
