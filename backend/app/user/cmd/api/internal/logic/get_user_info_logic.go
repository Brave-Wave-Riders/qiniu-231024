package logic

import (
	"TokTik/app/user/cmd/rpc/pb"
	"TokTik/common/vo"
	"context"
	"fmt"
	"github.com/jinzhu/copier"

	"TokTik/app/user/cmd/api/internal/svc"
	"TokTik/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	res, err := l.svcCtx.UserRpcClient.GetUserInfo(l.ctx, &pb.GetUserInfoReq{Id: req.Id})

	if err != nil {
		return &types.GetUserInfoResp{
			Status:  int(vo.ErrRequestParamError.GetErrCode()),
			Message: err.Error(),
			Error:   err.Error(),
		}, nil
	}

	fmt.Println(res)
	resp = &types.GetUserInfoResp{}
	_ = copier.Copy(resp, res)
	fmt.Println(resp)
	return resp, err
}
