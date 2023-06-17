package model

import (
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/exp/slog"
)

type WordsMe struct {
	IncludeMe bool   `db:"me_notification"`
	Word      string `db:"word"`
}

func PutWordsMe(includeMe bool, word string, userId string) error {
	wordMes := []WordsMe{}
	err := db.Select(&wordMes, "SELECT me_notification, word FROM words WHERE word = ? AND trap_id = ?", word, userId)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("No Data Found")
		return err
	} else if err != nil {
		slog.Info("Error: %s", err)
		return err
	}
	if len(wordMes) != 1 {
		return fmt.Errorf("Already exist too many data same trap_id and word: %d", len(wordMes))
	}
	if wordMes[0].IncludeMe == includeMe {
		return fmt.Errorf("Already %t", includeMe)
	}
	_, err = db.Exec("UPDATE words SET me_notification = ? WHERE word = ? AND trap_id = ?", includeMe, word, userId)
	if err != nil {
		return err
	}
	return nil
}

func PostWordsMeAll(includeMe bool, userId string) error {
	wordMes := []WordsMe{}
	err := db.Select(&wordMes, "SELECT me_notification, word FROM words WHERE trap_id = ?", userId)
	if errors.Is(err, sql.ErrNoRows) {
		slog.Info("No Data Found")
		return err
	} else if err != nil {
		slog.Info("Error: %s", err)
		return err
	}
	_, err = db.Exec("UPDATE words SET me_notification = ? WHERE trap_id = ?", includeMe, userId)
	if err != nil {
		return err
	}
	return nil
}
