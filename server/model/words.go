package model

import (
	"log"
	"time"
)

type WordAllListItem struct {
	IncludeBot bool      `db:"bot_notification"`
	IncludeMe  bool      `db:"me_notification"`
	Time       time.Time `db:"register_time"`
	UserId     string    `db:"trap_id"`
	Word       string    `db:"word"`
}

type WordsAllList []WordAllListItem

func GetWords() (WordsAllList, error) {
	words := []WordAllListItem{}
	err := db.Select(&words, "SELECT * FROM words")
	if err != nil {
		log.Printf("Error: %s\n", err)
		return nil, err
	}
	return words, nil
}
