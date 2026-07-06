package repo

import (
	"database/sql"
	"errors"
	"log/slog"

	"traQ-gazer/model"
)

func ResisterWord(word string, includeBot, includeMe bool, userId string) error {
	_, err := db.Exec(
		"INSERT INTO `words` (`trap_id`, `word`, `bot_notification`, `me_notification`) VALUES (?, ?, ?, ?)",
		userId,
		word,
		includeBot,
		includeMe,
	)
	return err
}

func DeleteWord(word string, userId string) error {
	_, err := db.Exec("DELETE FROM `words` WHERE `trap_id` = ? AND `word` = ?", userId, word)
	return err
}

func ExistWord(word, userId string) (bool, error) {
	var words model.WordAllListItem
	err := db.Get(&words, "SELECT * FROM `words` WHERE `trap_id` = ? AND `word` = ?", userId, word)
	// 削除するものが存在しない
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	// 予期せぬエラー
	if err != nil {
		slog.Error("failed to check existing word", "err", err)
		return false, err
	}

	// 存在する
	return true, nil
}

func GetWords() (model.WordsAllList, error) {
	words := []model.WordAllListItem{}
	err := db.Select(&words, "SELECT * FROM `words`")
	if err != nil {
		slog.Error("failed to select words", "err", err)
		return nil, err
	}
	return words, nil
}

func GetWordsWithoutTime() ([]model.WordsItem, error) {
	wordsList := []model.WordsItem{}
	err := db.Select(&wordsList, `
			SELECT
				word, bot_notification, me_notification, trap_id
			FROM
				words`)
	if err != nil {
		return nil, err
	}
	return wordsList, nil
}
