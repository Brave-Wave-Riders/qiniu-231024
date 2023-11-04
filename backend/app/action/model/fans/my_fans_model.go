package fans

import (
	"TokTik/app/user/model/fans"
	"context"
	"fmt"
	"gorm.io/gorm"
)

var (
	cacheQiniuTiktokUserFanPrefix = "cache:qiniuTiktok:user:fan"
)

type (
	myFansModel interface {
		IsFollowed(ctx context.Context, uid int64, fid int64) (bool, error)
		DeleteFollow(ctx context.Context, uid int64, fid int64) error
	}
)

// IsFollowed 查询关注的用户id
func (m *defaultFansModel) IsFollowed(ctx context.Context, uid int64, fid int64) (bool, error) {
	key := fmt.Sprintf("%s%v%v", cacheQiniuTiktokUserFanPrefix, uid, fid)
	var resp int64
	err := m.QueryCtx(ctx, &resp, key, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Fans{}).Where("user_id = ? and fan_id = ?", uid, fid).Count(&resp).Error
	})

	switch err {
	case nil:
		return resp > 0, nil
	case gorm.ErrRecordNotFound:
		return false, nil
	default:
		return false, err
	}
}

// DeleteFollow 查询关注关系
func (m *defaultFansModel) DeleteFollow(ctx context.Context, uid int64, fid int64) error {
	key := fmt.Sprintf("%s%v%v", cacheQiniuTiktokUserFanPrefix, uid, fid)
	var resp fans.Fans
	err := m.QueryCtx(ctx, &resp, key, func(conn *gorm.DB, v interface{}) error {
		return conn.Model(&Fans{}).Where("user_id = ? and fan_id = ?", uid, fid).Delete(&resp).Error
	})

	return err
}
