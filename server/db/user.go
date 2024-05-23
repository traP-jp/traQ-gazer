package db

import (
	"traQ-gazer/model"
)

func GetUserList() ([]model.UsersItem, error) {
	usersItem := []model.UsersItem{}
	err := db.Select(&usersItem,
		"SELECT `traq_uuid`, `trap_id`, `is_bot` FROM `users`",
	)
	if err != nil {
		return nil, err
	}
	return usersItem, nil
}
