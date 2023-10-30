package fans

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

var (
	cacheQiniuTiktokFollowingsIdPrefix = "cache:qiniuTiktok:following:id:"
)

type (
	myFansModel interface {
		SelectFollowings(ctx context.Context, uid int64) ([]int64, error)
		SelectFans(ctx context.Context, uid int64) ([]int64, error)
	}
)

// SelectFollowings 查询关注的用户id
func (m *defaultFansModel) SelectFollowings(ctx context.Context, uid int64) ([]int64, error) {
	key := fmt.Sprintf("%s%v", cacheQiniuTiktokFollowingsIdPrefix, uid)
	var fid []int64
	err := m.QueryCtx(ctx, &fid, key, func(conn *gorm.DB, v interface{}) error {
		err := conn.Where("fan_id = ?", uid).Find(&fid).Error
		return err
	})
	return fid, err
}

// SelectFans 查询粉丝id
func (m *defaultFansModel) SelectFans(ctx context.Context, uid int64) ([]int64, error) {
	key := fmt.Sprintf("%s%v", cacheQiniuTiktokFansIdPrefix, uid)
	var fid []int64
	err := m.QueryCtx(ctx, &fid, key, func(conn *gorm.DB, v interface{}) error {
		err := conn.Where("user_id = ?", uid).Find(&fid).Error
		return err
	})
	return fid, err

}
