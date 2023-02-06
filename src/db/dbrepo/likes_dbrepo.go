package dbrepo

import (
	"fmt"

	"github.com/tmsic/deegle-translate-api/model"
	"github.com/tmsic/deegle-translate-api/utils"
)

func AllLikes() int {
	db := utils.OpenDB()
	defer db.Close()
	var count int
	if err := db.Model(model.Likes{}).Count(&count).Error; err != nil {
		fmt.Println(err)
	}
	return count
}
