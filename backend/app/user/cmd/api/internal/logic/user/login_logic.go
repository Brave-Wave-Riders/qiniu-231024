package user

import (
	"TokTik/app/user/cmd/rpc/pb"
	"TokTik/common/vo"
	"context"
	"github.com/jinzhu/copier"

	"TokTik/app/user/cmd/api/internal/svc"
	"TokTik/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp, err := l.svcCtx.UserRpcClient.Login(l.ctx, &pb.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		return &types.LoginResp{
			Status:  int(vo.ErrRequestParamError.GetErrCode()),
			Message: err.Error(),
			Error:   err.Error(),
		}, nil
	}

	resp = &types.LoginResp{}
	_ = copier.Copy(resp, loginResp)
	return resp, err
}
