package logic

import (
	"TokTik/app/video/rpc/cmd/video"
	"context"

	"TokTik/app/video/api/cmd/internal/svc"
	"TokTik/app/video/api/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteVideoLogic {
	return &DeleteVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteVideoLogic) DeleteVideo(req *types.DeleteReq) (resp *types.DeleteResp, err error) {
	// todo: add your logic here and delete this line
	deleteResp, err := l.svcCtx.Video.Delete(l.ctx, &video.DeleteReq{Id: int32(req.Id)})
	if err != nil {
		return nil, err
	}
	logx.Info(deleteResp)
	return &types.DeleteResp{
		Code: int(deleteResp.Code),
		Msg:  deleteResp.Msg,
	}, nil
}
