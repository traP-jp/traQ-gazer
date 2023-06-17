package model

import (
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


func PutWords(data *api.WordBotSetting, userId string) error {
	// 該当するword trap_idの組が存在するかチェック
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
	_, err = db.Exec("UPDATE words SET bot_notification = ? WHERE word = ? AND trap_id = ?", data.IncludeBot, data.Word, userId)
	return err
}

func PostWordsBot(data *api.Bot, userId string) error {
	_, err := db.Exec("UPDATE words SET bot_notification = ? WHERE word = ? AND trap_id = ?", data.IncludeBot, userId)
	return err
}
