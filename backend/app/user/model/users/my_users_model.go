package users

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

var (
	cacheQiniuTiktokUserIdsPrefix = "cache:qiniuTiktok:user:ids:"
)

type (
	myUsersModel interface {
		SelectUsersByIds(ctx context.Context, uids []int64) ([]Users, error)
	}
)

// SelectUsersByIds 根据id列表查询用户信息
func (m *defaultUsersModel) SelectUsersByIds(ctx context.Context, uids []int64) ([]Users, error) {
	key := fmt.Sprintf("%s%v", cacheQiniuTiktokUserIdsPrefix, uids)
	var users []Users
	err := m.QueryCtx(ctx, &users, key, func(conn *gorm.DB, v interface{}) error {
		err := conn.Find(&users, uids).Error
		return err
	})
	return users, err
}
