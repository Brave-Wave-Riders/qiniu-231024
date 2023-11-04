package logic

import (
	"TokTik/app/action/model/fans"
	"TokTik/common/vo"
	"context"
	"gorm.io/gorm"

	"TokTik/app/action/cmd/rpc/internal/svc"
	"TokTik/app/action/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FollowLogic) Follow(in *pb.FollowReq) (*pb.FollowResp, error) {
	fan := fans.Fans{
		UserId: in.UserId,
		FanId:  in.FansId,
	}

	err := l.svcCtx.FansModel.Transaction(l.ctx, func(db *gorm.DB) error {
		// 查询是否关注过
		followed, err := l.svcCtx.FansModel.IsFollowed(l.ctx, fan.UserId, fan.FanId)
		if err != nil {
			return err
		}

		if followed {
			// 已关注, 取消关注
			return l.svcCtx.FansModel.DeleteFollow(l.ctx, fan.UserId, fan.FanId)
		} else {
			// 未关注, 关注
			return l.svcCtx.FansModel.Insert(l.ctx, nil, &fan)
		}
	})

	if err != nil {
		return nil, err
	}

	return &pb.FollowResp{
		Status:  vo.OK,
		Data:    "操作成功",
		Message: "操作成功",
	}, nil
}
