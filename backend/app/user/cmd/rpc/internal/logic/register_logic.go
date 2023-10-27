package logic

import (
	"TokTik/app/user/cmd/rpc/internal/svc"
	"TokTik/app/user/cmd/rpc/pb"
	"TokTik/app/user/model"
	"TokTik/common/tool"
	"TokTik/common/vo"
	"context"
	"database/sql"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	// 根据用户名查询用户是否存在
	user, err := l.svcCtx.UserModel.FindOneByEmail(l.ctx, in.Email)

	// 查询出错
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(vo.ErrDBerror, "数据库查询出错")
	}

	// 用户名已存在
	if user != nil {
		return nil, errors.Wrap(vo.ErrUserAlreadyRegisterError, "用户名已存在")
	}

	// 添加用户
	bp, _ := tool.BcryptByString(in.Password)
	user = &model.Users{
		Username: in.Username,
		//Avatar:   in.Avatar,
		Avatar: sql.NullString{
			String: in.Avatar,
			Valid:  true,
		},
		Email:    in.Email,
		Password: bp,
	}
	err = l.svcCtx.UserModel.Insert(l.ctx, nil, user)

	// 注册失败
	if err != nil {
		return nil, errors.Wrap(vo.ErrDBerror, "数据库插入出错")
	}

	return &pb.RegisterResp{
		Status:  vo.OK,
		Data:    vo.SUCCESS,
		Message: vo.SUCCESS,
		Error:   "",
	}, nil
}
