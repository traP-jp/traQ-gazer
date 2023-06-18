package model

import "time"

type WordListItem struct {
	IncludeBot bool      `db:"bot_notification"`
	IncludeMe  bool      `db:"me_notification"`
	Time       time.Time `db:"register_time"`
	Word       string    `db:"word"`
}

type WordsList []WordListItem

func GetListUserUserId(userId string) (WordsList, error) {
	wordsList := WordsList{}
	err := db.Select(&wordsList, "SELECT bot_notification, me_notification, register_time, word FROM words WHERE trap_id = ?", userId)
	if err != nil {
		return nil, err
	}
	return wordsList, nil
}

type UsersOfWordListItem struct {
	UserIds []UserListItem `db:"user_ids" json:"user_ids"`
	Word    string         `db:"word" json:"word"`
}

type UsersOfWordsList = []UsersOfWordListItem

func GetListUserUserIdUsers(userId string) (UsersOfWordsList, error) {
	return nil, nil
}

type UserListItem struct {
	IncludeBot bool      `db:"bot_notification"`
	IncludeMe  bool      `db:"me_notification"`
	Time       time.Time `db:"register_time"`
	UserId     string    `db:"trap_id"`
}

type UsersList = []UserListItem

func GetListWordWord(word string) (UsersList, error) {
	usersList := UsersList{}
	err := db.Select(&usersList, "SELECT bot_notification, me_notification, register_time, trap_id FROM words WHERE word = ?", word)
	if err != nil {
		return nil, err
	}
	return usersList, nil
}

func GetListWordWordWords(word string) (WordsList, error) {
	return nil, nil
}
