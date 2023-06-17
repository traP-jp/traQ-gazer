package model

import (
	// "database/sql"
	// "errors"
	// "fmt"
	"database/sql"
	"errors"
	"time"

	"log"

	"golang.org/x/exp/slog"
)

type WordAllListItem struct {
	IncludeBot bool      `db:"bot_notification"`
	IncludeMe  bool      `db:"me_notification"`
	Time       time.Time `db:"register_time"`
	UserId     string    `db:"trap_id"`
	Word       string    `db:"word"`
}

func ResisterWord(word string, includeBot, includeMe bool, userId string) error {
	_, err := db.Exec("INSERT INTO words (trap_id, word, bot_notification, me_notification) VALUES (?, ?, ?, ?)", userId, word, includeBot, includeMe)
	return err
}

func DeleteWord(word string, userId string) error {
	_, err := db.Exec("DELETE FROM words WHERE trap_id = ? AND word = ?", userId, word)
	return err
}

func ExistWord(word, userId string) (bool, error) {
	var words WordAllListItem
	err := db.Get(&words, "SELECT * FROM words WHERE trap_id = ? AND word = ?", userId, word)
	// 削除するものが存在しない
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	// 予期せぬエラー
	if err != nil {
		slog.Info(err.Error())
		return false, err
	}

	// 存在する
	return true, nil
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
