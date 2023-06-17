package model

import (
	"database/sql"
	"errors"
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

func ChengeBotNotification(word string, includeBot bool, userId string) error {
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
	_, err = db.Exec("UPDATE words SET bot_notification = ? WHERE word = ? AND trap_id = ?", includeBot, word, userId)
	return err
}

func ChangeAllBotNotification(includeBot bool, userId string) error {
	_, err := db.Exec("UPDATE words SET bot_notification = ? WHERE trap_id = ?", includeBot, userId)
	return err
}
