package model

import (
	"database/sql"
	"errors"
	"h23s_15/api"
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

func GetWords() (api.WordsAllList, error) {
	words := []WordAllListItem{}
	err := db.Select(&words, "SELECT * FROM words")
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("No Data Found \n")
		return nil, err
	} else if err != nil {
		log.Printf("Error: %s\n", err)
		return nil, err
	}

	change := ConvertSliceToA1(words)

	return change, nil
}

func ConvertSliceToA1(WordsListSlice WordsAllList) api.WordsAllList {
	WordsAllListSlice := make([]api.WordAllListItem, len(WordsListSlice))
	for i, WordType := range WordsListSlice {
		WordsAllListSlice[i] = api.WordAllListItem{
			IncludeBot: WordType.IncludeBot,
			IncludeMe:  WordType.IncludeMe,
			Time:       WordType.Time,
			UserId:     WordType.UserId,
			Word:       WordType.Word,
		}
	}
	return WordsAllListSlice
}
