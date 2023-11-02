package fan

import (
	"TokTik/app/user/cmd/rpc/pb"
	"TokTik/common/vo"
	"context"
	"github.com/jinzhu/copier"

	"TokTik/app/user/cmd/api/internal/svc"
	"TokTik/app/user/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserFollowingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserFollowingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFollowingLogic {
	return &GetUserFollowingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserFollowingLogic) GetUserFollowing(req *types.GetFollowingsRequest) (resp *types.GetFollowingsResponse, err error) {
	id := req.UserId

	fans, err := l.svcCtx.UserRpcClient.GetFollowings(l.ctx, &pb.GetFollowingsReq{Id: id})
	if err != nil {
		return &types.GetFollowingsResponse{
			Status:  int(vo.ErrServerCommonError.GetErrCode()),
			Message: err.Error(),
			Error:   err.Error(),
		}, nil
	}

	resp = &types.GetFollowingsResponse{}
	_ = copier.Copy(resp, fans)
	return resp, err
}
