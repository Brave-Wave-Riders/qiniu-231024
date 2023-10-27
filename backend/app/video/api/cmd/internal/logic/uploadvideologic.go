package logic

import (
	"TokTik/app/video/rpc/cmd/video"
	"TokTik/common/dyerr"
	"TokTik/common/type_"
	"context"
	"encoding/json"
	"net/http"

	"TokTik/app/video/api/cmd/internal/svc"
	"TokTik/app/video/api/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadVideoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadVideoLogic {
	return &UploadVideoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadVideoLogic) UploadVideo(req *types.UploapReq, r *http.Request) (resp *types.UploapResp, err error) {
	// todo: add your logic here and delete this line
	//
	id := l.ctx.Value("id").(int)
	resp = new(types.UploapResp)
	_, m, err := r.FormFile("file")
	//验证大小
	if m.Size > 100*1024*1024 {
		resp.Msg = "视频大小不能超过100M"
		resp.Code = 1
		return
	}

	file, err := json.Marshal(type_.VideoFile{
		Files:  m,
		Userid: id,
	})
	if err != nil {

	}
	//异步检验视频内容
	err = l.svcCtx.KqPusherClient.Push(string(file))
	if err != nil {
		return nil, dyerr.NewErrMsg(err.Error())
	}

	//上传
	_, err = l.svcCtx.Video.Upload(l.ctx, &video.UploadReq{
		Title: "",
		Kind:  0,
	})
	if err != nil {
		return nil, err
	}

	return
}
