package logic

import (
	"context"

	"TokTik/app/video/rpc/cmd/internal/svc"
	"TokTik/app/video/rpc/cmd/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UploadLogic) Upload(in *video.UploadReq) (*video.UploadResp, error) {
	// todo: add your logic here and delete this line

	return &video.UploadResp{}, nil
}
