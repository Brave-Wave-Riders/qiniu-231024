package logic

import (
	"TokTik/app/user/model/users"
	"TokTik/common/vo"
	"context"
	"github.com/pkg/errors"

	"TokTik/app/user/cmd/rpc/internal/svc"
	"TokTik/app/user/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	// 查找用户
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil && err != users.ErrNotFound {
		return nil, errors.Wrap(vo.ErrDBError, "数据库查询出错")
	}
	if err == users.ErrNotFound || user == nil {
		return nil, errors.Wrap(vo.ErrUserNoExistsError, "用户不存在")
	}

	return &pb.GetUserInfoResp{
		Status: vo.OK,
		Data: &pb.User{
			Id:       user.Id,
			Username: user.Username,
			Avatar:   user.Avatar.String,
			Email:    user.Email,
		},
		Message: vo.SUCCESS,
		Error:   "",
	}, nil
}
