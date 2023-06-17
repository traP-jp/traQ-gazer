package handler

import "github.com/labstack/echo/v4"

// 似たような者を探す
// (GET /similar/{userId})
func (s Server) GetUsersWithSimilarWords(ctx echo.Context, userId string) error {
	return nil
}

// おすすめの単語を出す
// (GET /similar/{userId}/recommend)
func (s Server) GetRecommendedWordsForUser(ctx echo.Context, userId string) error {
	return nil
}
