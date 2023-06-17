package handler

import (
	"h23s_15/model"

	"github.com/labstack/echo/v4"
)

// 似たような者を探す
// (GET /similar/{userId})
func (s Server) GetUsersWithSimilarWords(ctx echo.Context, userId string) error {
	model.GetUsersWithSimilarWords(userId, 10)
	return nil
}

// おすすめの単語を出す
// (GET /similar/{userId}/recommend)
func (s Server) GetRecommendedWordsForUser(ctx echo.Context, userId string) error {
	model.GetRecommendedWordsForUser(userId, 10)
	return nil
}
