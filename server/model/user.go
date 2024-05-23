package model

import "time"

type UserListItem struct {
	IncludeBot bool      `db:"bot_notification"`
	IncludeMe  bool      `db:"me_notification"`
	Time       time.Time `db:"register_time"`
	UserId     string    `db:"trap_id"`
}

type UsersList = []UserListItem

type UsersOfWordListItem struct {
	UserIds []UserListItem `db:"user_ids" json:"user_ids"`
	Word    string         `db:"word" json:"word"`
}

type UsersOfWordsList = []UsersOfWordListItem

type WordListItem struct {
	IncludeBot bool      `db:"bot_notification"`
	IncludeMe  bool      `db:"me_notification"`
	Time       time.Time `db:"register_time"`
	Word       string    `db:"word"`
}

type WordsList []WordListItem
