package model

import "time"

type Video struct {
	ID            int        `json:"id"`
	Title         string     `json:"title"`
	AuthorID      int        `json:"author_id"`
	PlayURL       string     `json:"play_url"`
	CoverURL      string     `json:"cover_url"`
	FavoriteCount int        `json:"favorite_count"`
	CommentCount  int        `json:"comment_count"`
	CreateTime    time.Time  `json:"create_time"`
	UpdateTime    time.Time  `json:"update_time"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Kind          int        `json:"kind"`
}
