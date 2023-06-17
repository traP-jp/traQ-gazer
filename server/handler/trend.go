package handler

import (
	"h23s_15/api"

	"github.com/labstack/echo/v4"
)

// 今日のトレンド
// (GET /trend/day/today)
func (s Server) GetTodayTrendingWords(ctx echo.Context, params api.GetTodayTrendingWordsParams) error {
	return nil
}

// ある日のトレンド
// (GET /trend/day/{day})
func (s Server) GetTrendingWordsForDay(ctx echo.Context, day string, params api.GetTrendingWordsForDayParams) error {
	return nil
}

// ある月のトレンド
// (GET /trend/month/{month})
func (s Server) GetTrendingWordsForMonth(ctx echo.Context, month string, params api.GetTrendingWordsForMonthParams) error {
	return nil
}

// ある年のトレンド
// (GET /trend/year/{year})
func (s Server) GetTrendingWordsForYear(ctx echo.Context, year string, params api.GetTrendingWordsForYearParams) error {
	return nil
}
