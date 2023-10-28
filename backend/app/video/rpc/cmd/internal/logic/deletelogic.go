package logic

import (
	"TokTik/app/video/rpc/model"
	"TokTik/common/dyerr"
	"context"

	"TokTik/app/video/rpc/cmd/internal/svc"
	"TokTik/app/video/rpc/cmd/video"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteLogic {
	return &DeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteLogic) Delete(in *video.DeleteReq) (*video.DeleteResp, error) {
	// todo: add your logic here and delete this line
	//判断是否存在
	videos := new(model.Video)
	ok, err2 := l.svcCtx.Model.Where("id=?", in.Id).Get(videos)
	if err2 != nil {
		return nil, err2
	}
	if !ok {
		return nil, dyerr.NewErrMsg("非法数据！")
	}
	_, err := l.svcCtx.Model.Where("id=?", in.Id).Delete(videos)
	if err != nil {
		return nil, dyerr.NewErrMsg("删除失败！")
	}
	//删除redis中的值,更新缓存
	del, err := l.svcCtx.Redis.Del(string(in.Id))
	if err != nil {
		return nil, err
	}
	if del == 0 {
		return nil, dyerr.NewErrMsg("缓存删除失败！")
	}
	if err != nil {
		return nil, err
	}
	return &video.DeleteResp{
		Msg:  "删除成功！",
		Code: 200,
	}, nil
}
