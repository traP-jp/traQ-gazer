package db

import (
	"traQ-gazer/model"
)

func GetListUserUserId(userId string) (model.WordsList, error) {
	wordsList := model.WordsList{}
	err := db.Select(&wordsList, "SELECT bot_notification, me_notification, register_time, word FROM words WHERE trap_id = ?", userId)
	if err != nil {
		return nil, err
	}
	return wordsList, nil
}

func GetListUserUserIdUsers(userId string) (model.UsersOfWordsList, error) {
	return nil, nil
}

func GetListWordWord(word string) (model.UsersList, error) {
	usersList := model.UsersList{}
	err := db.Select(&usersList, "SELECT bot_notification, me_notification, register_time, trap_id FROM words WHERE word = ?", word)
	if err != nil {
		return nil, err
	}
	return usersList, nil
}

func GetListWordWordWords(word string) (model.WordsList, error) {
	return nil, nil
}
