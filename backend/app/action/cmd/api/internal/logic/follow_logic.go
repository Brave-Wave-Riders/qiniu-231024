package logic

import (
	"TokTik/app/action/cmd/rpc/pb"
	"TokTik/common/ctxdata"
	"TokTik/common/vo"
	"context"
	"github.com/jinzhu/copier"

	"TokTik/app/action/cmd/api/internal/svc"
	"TokTik/app/action/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowLogic) Follow(req *types.FollowRequest) (resp *types.FollowResponse, err error) {
	uid := ctxdata.GetUidFromCtx(l.ctx)
	res, err := l.svcCtx.ActionRpcClient.Follow(l.ctx, &pb.FollowReq{
		UserId: uid,
		FansId: req.Id,
	})

	if err != nil {
		return &types.FollowResponse{
			Status:  int(vo.ErrServerCommonError.GetErrCode()),
			Data:    err.Error(),
			Message: err.Error(),
			Error:   err.Error(),
		}, nil
	}

	resp = &types.FollowResponse{}
	_ = copier.Copy(resp, res)
	return resp, err
}
