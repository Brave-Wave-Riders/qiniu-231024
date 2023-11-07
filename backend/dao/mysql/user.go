package mysql

type User struct {
	BaseModel
	Mobile   string `gorm:"index:idx_mobile;unique;type:varchar(11);not null" json:"mobile"`
	Password string `gorm:"type:varchar(100);not null" json:"password"`
	NickName string `gorm:"type:varchar(20)" json:"nick_name"`
	Gender   string `gorm:"column:gender;default:male;type:varchar(6) comment 'female表示女, male表示男'" json:"gender"`
	Url      string `gorm:"type:varchar(100);"`
	Role     int    `gorm:"column:role;default:1;type:int comment '1表示普通用户, 2表示管理员'"`
}
