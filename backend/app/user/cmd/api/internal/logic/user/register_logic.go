package user

import (
	"TokTik/app/user/cmd/api/internal/svc"
	"TokTik/app/user/cmd/api/internal/types"
	"TokTik/app/user/cmd/rpc/pb"
	"TokTik/common/vo"
	"context"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	// 向 rpc 发送请求
	registerResp, err := l.svcCtx.UserRpcClient.Register(l.ctx, &pb.RegisterReq{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Avatar:   req.Avatar,
	})

	// 出现错误
	if err != nil {
		return &types.RegisterResp{
			Status:  int(vo.ErrRequestParamError.GetErrCode()),
			Data:    vo.ErrRequestParamError.GetErrMsg(),
			Message: err.Error(),
			Error:   err.Error(),
		}, nil
	}

	resp = &types.RegisterResp{}
	_ = copier.Copy(resp, registerResp)
	return resp, err
}
