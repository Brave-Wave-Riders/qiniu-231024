package logic

import (
	"TokTik/common/vo"
	"context"
	"github.com/pkg/errors"

	"TokTik/app/user/cmd/rpc/internal/svc"
	"TokTik/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowingsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingsLogic {
	return &GetFollowingsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetFollowingsLogic) GetFollowings(in *pb.GetFollowingsReq) (*pb.GetFollowingsResp, error) {
	// 查询关注列表
	followingIds, err := l.svcCtx.FansModel.SelectFollowings(l.ctx, in.Id)
	if err != nil {
		return nil, errors.Wrap(vo.ErrDBError, "数据库查询出错")
	}

	// 查询详细信息
	users, err := l.svcCtx.UserModel.SelectUsersByIds(l.ctx, followingIds)
	if err != nil {
		return nil, errors.Wrap(vo.ErrDBError, "数据库查询出错")
	}

	userResp := make([]*pb.User, len(users))
	for i := 0; i < len(users); i++ {
		u := users[i]
		userResp[i] = &pb.User{
			Id:       u.Id,
			Username: u.Username,
			Avatar:   u.Avatar.String,
			Email:    u.Email,
		}
	}
	return &pb.GetFollowingsResp{
		Status:  vo.OK,
		Data:    userResp,
		Message: "查询成功",
	}, nil
}
