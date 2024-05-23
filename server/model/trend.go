package model

type TrendingWord struct {
	Number int    `db:"number"`
	Word   string `db:"word"`
}

type TrendingWords = []TrendingWord
