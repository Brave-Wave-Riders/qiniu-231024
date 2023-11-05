package mysql

type Video struct {
	BaseModel
	Uid       uint32 `json:"uid" gorm:"type:int;not null;index:idx_user"`
	Name      string `json:"name" gorm:"type:varchar(700);not null;unique"`
	VideoType int    `json:"type" gorm:"type:int;not null; index:idx_type"`
	ImgUrl    string `json:"img_url" gorm:"type:varchar(1024);"`
}
