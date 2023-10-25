package logic

import (
	"context"

	"TokTik/app/video/cmd/internal/svc"
	"TokTik/app/video/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVedioLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadVedioLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVedioLogic {
	return &UploadVedioLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadVedioLogic) UploadVedio(req *types.UploapReq) (resp *types.UploapResp, err error) {
	// todo: add your logic here and delete this line

	return
}
