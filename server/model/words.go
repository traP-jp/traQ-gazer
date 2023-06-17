package model

import (
	"database/sql"
	"errors"
	"log"
	"h23s_15/api"
)

func GetWords() (api.WordsAllList, error) {
	var word api.WordsAllList
	err := db.Get(&word, "SELECT * FROM words")
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("No Data Found \n")
		return word, err
	} else if err != nil {
		log.Printf("Error \n")
		return word, err
	}
	return word, nil
}
