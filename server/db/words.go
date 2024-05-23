package db

import (
	"database/sql"
	"errors"
	"fmt"
	"traQ-gazer/model"

	"log"

	"golang.org/x/exp/slog"
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
		slog.Info(err.Error())
		return false, err
	}

	// 存在する
	return true, nil
}

func GetWords() (model.WordsAllList, error) {
	words := []model.WordAllListItem{}
	err := db.Select(&words, "SELECT * FROM `words`")
	if err != nil {
		log.Printf("Error: %s\n", err)
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
		slog.Info("Error selecting words: %v", err)
		return nil, err
	}
	return wordsList, nil
}

func GetMatchedWordList(messageItem model.MessageItem) ([]model.MatchedWords, error) {
	matchedWordsList := []model.MatchedWords{}
	err := db.Select(&matchedWordsList, `
			SELECT
				group_concat(words.word SEPARATOR '\n') AS contacted_words,
				words.trap_id AS trap_id,
				users.traq_uuid AS traq_uuid
			FROM words
			JOIN users ON words.trap_id = users.trap_id
				WHERE (
				    	((word NOT LIKE '/%/') 
				    		AND (? LIKE concat('%', word, '%')))
            			OR ((word LIKE '/%/') 
            				AND (BINARY ? REGEXP trim(BOTH '/' FROM word)))
				    )
				AND (me_notification OR
					 users.traq_uuid != ?)
				AND (bot_notification OR
					 (SELECT is_bot FROM users WHERE traq_uuid = ? LIMIT 1) = FALSE)
			GROUP BY words.trap_id`,
		messageItem.Content, messageItem.Content, messageItem.TraqUuid, messageItem.TraqUuid)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to search words with message: `%s`", messageItem.Id))
		return nil, err
	}
	return matchedWordsList, nil
}
