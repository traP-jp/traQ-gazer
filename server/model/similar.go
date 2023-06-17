package model

type SimilarUser struct {
	UserId string `db:"traq_id"`
}

type SimilarUsers = []SimilarUser

func GetUsersWithSimilarWords(userId string, limit int) (SimilarUsers, error) {
	similarUsers := SimilarUsers{}
	err := db.Select(&similarUsers, "SELECT w.trap_id, COUNT(*) AS word_count FROM words w WHERE w.word IN (SELECT word FROM words WHERE trap_id = ? ) AND w.trap_id != ? GROUP BY w.trap_id ORDER BY word_count DESC LIMIT ?", userId, limit)
	if err != nil {
		return nil, err
	}
	return similarUsers, nil
}

type RecommendedWord struct {
	Number int    `db:"number"`
	Word   string `db:"word"`
}

type RecommendedWords = []RecommendedWord

func GetRecommendedWordsForUser(userId string, limit int) (RecommendedWords, error) {
	return nil, nil
}
