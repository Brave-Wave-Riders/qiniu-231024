package fans

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)

var _ FansModel = (*customFansModel)(nil)

type (
	// FansModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFansModel.
	FansModel interface {
		fansModel
		customFansLogicModel
	}

	customFansModel struct {
		*defaultFansModel
	}

	customFansLogicModel interface {
		myFansModel
	}
)

// NewFansModel returns a model for the database table.
func NewFansModel(conn *gorm.DB, c cache.CacheConf) FansModel {
	return &customFansModel{
		defaultFansModel: newFansModel(conn, c),
	}
}

func (m *defaultFansModel) customCacheKeys(data *Fans) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}
