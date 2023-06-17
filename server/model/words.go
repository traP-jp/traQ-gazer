package model

import (
	"h23s_15/api"
)


func PostWords(data *api.PostWordsJSONBody, userId string) error {
	_, err := db.Exec("INSERT INTO words (trap_id, word, bot_notification, me_notification) VALUES (?, ?, ?, ?)", userId, data.Word, data.IncludeBot, data.IncludeMe) 
	return err
}

func DeleteWords(data *api.DeleteWordsJSONRequestBody, userId string) error {
	_, err := db.Exec("DELETE FROM words WHERE trap_id = ? AND word = ?", userId, data.Word) 
	return err
}