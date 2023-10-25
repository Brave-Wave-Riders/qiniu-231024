package logic

import (
	"TokTik/app/video/api/cmd/internal/svc"
	"TokTik/app/video/api/cmd/internal/types"
	"TokTik/app/video/rpc/cmd/video"
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
	resp = new(types.UploapResp)
	_, m, err := r.FormFile("file")
	//验证大小
	if m.Size > 100*1024*1024 {
		resp.Msg = "视频大小不能超过100M"
		resp.Code = 1
		return
	}
	//异步检验视频内容

	//上传

	//
	_, err = l.svcCtx.Video.Upload(l.ctx, &video.UploadReq{
		Title: "",
		Kind:  0,
	})
	if err != nil {
		return nil, err
	}

	return
}
