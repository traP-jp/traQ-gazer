package model

import "log"

type TrendingWord struct {
	Number int    `db:"count"`
	Word   string `db:"word"`
}

type TrendingWords = []TrendingWord

func GetTodayTrendingWords(limit int) (TrendingWords, error) {
	words := TrendingWords{}
	err := db.Select(&words, "SELECT word, COUNT(word) AS count FROM words WHERE DATE(register_time) = CURDATE() GROUP BY word ORDER BY count DESC LIMIT ?", limit)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return nil, err
	}
	return words, nil
}

func GetTrendingWordsForDay(limit int, day string) (TrendingWords, error) {
	words := TrendingWords{}
	err := db.Select(&words, "SELECT word, COUNT(word) AS count FROM words WHERE DATE(register_time) = ? GROUP BY word ORDER BY count DESC LIMIT ?", day, limit)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return nil, err
	}
	return words, nil
}

func GetTrendingWordsForMonth(limit int, month string) (TrendingWords, error) {
	words := TrendingWords{}
	err := db.Select(&words, "SELECT word, COUNT(word) AS count FROM words WHERE DATE_FORMAT(register_time, '%Y-%m') = ? GROUP BY word ORDER BY count DESC LIMIT ?", month, limit)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return nil, err
	}
	return words, nil
}

func GetTrendingWordsForYear(limit int, year string) (TrendingWords, error) {
	words := TrendingWords{}
	err := db.Select(&words, "SELECT word, COUNT(word) AS count FROM words WHERE DATE_FORMAT(register_time, '%Y') = ? GROUP BY word ORDER BY count DESC LIMIT ?", year, limit)
	if err != nil {
		log.Printf("Error: %s\n", err)
		return nil, err
	}
	return words, nil
}
