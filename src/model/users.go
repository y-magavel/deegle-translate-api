package model

import "time"

type User struct {
	UserID    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
	DeleteFlg string    `json:"delete_flg"`
}
