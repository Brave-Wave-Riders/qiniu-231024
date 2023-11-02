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

type GetFansLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFansLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFansLogic {
	return &GetFansLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFansLogic) GetFans(req *types.GetFansRequest) (resp *types.GetFansResponse, err error) {
	id := req.UserId

	fans, err := l.svcCtx.UserRpcClient.GetFans(l.ctx, &pb.GetFansReq{Id: id})
	if err != nil {
		return &types.GetFansResponse{
			Status:  int(vo.ErrServerCommonError.GetErrCode()),
			Message: err.Error(),
			Error:   err.Error(),
		}, nil
	}

	resp = &types.GetFansResponse{}
	_ = copier.Copy(resp, fans)
	return resp, err
}
