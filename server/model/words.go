package model

import (
	// "database/sql"
	// "errors"
	// "fmt"
	"database/sql"
	"errors"
	"h23s_15/api"
	"time"

	"golang.org/x/exp/slog"
)

type WordAllListItem struct {
	IncludeBot bool      `db:"bot_notification"`
	IncludeMe  bool      `db:"me_notification"`
	Time       time.Time `db:"register_time"`
	UserId     string    `db:"trap_id"`
	Word       string    `db:"word"`
}

func PostWords(data *api.PostWordsJSONBody, userId string) error {
	var words WordAllListItem
	err := db.Get(&words, "SELECT * FROM words WHERE trap_id = ? AND word = ?", userId, data.Word)

	// 同じword trap_idの組が存在する
	if err == nil {
		slog.Info("Already Resistered")
		return errors.New("Already Resistered")
	}

	// 予期せぬエラー
	if !errors.Is(err, sql.ErrNoRows) {
		slog.Info("!!!!!!")
		return err
	}

	_, err = db.Exec("INSERT INTO words (trap_id, word, bot_notification, me_notification) VALUES (?, ?, ?, ?)", userId, data.Word, data.IncludeBot, data.IncludeMe)
	return err
}

func DeleteWords(data *api.DeleteWordsJSONRequestBody, userId string) error {
	var words WordAllListItem
	err := db.Get(&words, "SELECT * FROM words WHERE trap_id = ? AND word = ?", userId, data.Word)

	if err != nil {
		// 予期せぬエラー
		if !errors.Is(err, sql.ErrNoRows) {
			slog.Info("!!!!!!!!!!")
			return err
		}
		// 同じword trap_idの組が存在しない
		slog.Info("Not Found")
		return errors.New("Not Found")
	}

	_, err = db.Exec("DELETE FROM words WHERE trap_id = ? AND word = ?", userId, data.Word)
	return err
}
