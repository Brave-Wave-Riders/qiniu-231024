package logic

import (
	"TokTik/app/video/api/cmd/internal/svc"
	"TokTik/app/video/api/cmd/internal/types"
	"context"
	"net/http"

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

func (l *UploadVedioLogic) UploadVedio(req *types.UploapReq, r *http.Request) (resp *types.UploapResp, err error) {
	//

	_, m, err := r.FormFile("file")
	//验证大小
	if m.Size > 100*1024*1024 {

	}

	if err != nil {
		return nil, err
	}
	return
}
