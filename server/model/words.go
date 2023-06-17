package model

import (
	// "database/sql"
	// "errors"
	// "fmt"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/labstack/echo"
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
	var words WordAllListItem
	err := db.Get(&words, "SELECT * FROM words WHERE trap_id = ? AND word = ?", userId, word)

	// 同じword trap_idの組が存在する
	if err == nil {
		slog.Info("Already Resistered")
		return echo.NewHTTPError(http.StatusBadRequest, "Already Resistered")
	}

	// 予期せぬエラー
	if !errors.Is(err, sql.ErrNoRows) {
		slog.Info(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	_, err = db.Exec("INSERT INTO words (trap_id, word, bot_notification, me_notification) VALUES (?, ?, ?, ?)", userId, word, includeBot, includeMe)
	return err
}

func DeleteWord(word string, userId string) error {
	var words WordAllListItem
	err := db.Get(&words, "SELECT * FROM words WHERE trap_id = ? AND word = ?", userId, word)

	// 削除するものが存在しない
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("Not Found")
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	// 予期せぬエラー
	if err != nil {
		slog.Info(err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	_, err = db.Exec("DELETE FROM words WHERE trap_id = ? AND word = ?", userId, word)
	return err
}
