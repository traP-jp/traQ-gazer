package model

import "h23s_15/api"

func CheckMessageFromWords(message string) (api.UsersList, error) {
	words := WordsAllList{}
	if err := db.Select(&words, "SELECT * FROM words"); err != nil {
		return nil, err
	}
	return nil, nil
}
