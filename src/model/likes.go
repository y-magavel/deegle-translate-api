package model

import "time"

type Likes struct {
	ID        string    `json:"id"`
	UserID    int       `json:"user_id"`
	LikesType string    `json:"likes_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	DeleteFlg string    `json:"delete_flg"`
}
