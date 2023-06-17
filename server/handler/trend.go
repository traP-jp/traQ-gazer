package handler

import (
	"h23s_15/api"
	"h23s_15/model"

	"github.com/labstack/echo/v4"
)

// 今日のトレンド
// (GET /trend/day/today)
func (s Server) GetTodayTrendingWords(ctx echo.Context, params api.GetTodayTrendingWordsParams) error {
	limitNumber := 10
	if params.Limit != nil {
		limitNumber = *params.Limit
	}
	trendingWordsModel, err := model.GetTodayTrendingWords(limitNumber)
	if err != nil {
		return echo.NewHTTPError(500, err)
	}
	trendingWordsApi := ConvertTrendingWords(trendingWordsModel)
	return ctx.JSON(200, trendingWordsApi)
}

// ある日のトレンド
// (GET /trend/day/{day})
func (s Server) GetTrendingWordsForDay(ctx echo.Context, day string, params api.GetTrendingWordsForDayParams) error {
	limitNumber := 10
	if params.Limit != nil {
		limitNumber = *params.Limit
	}
	trendingWordsModel, err := model.GetTrendingWordsForDay(limitNumber, day)
	if err != nil {
		return echo.NewHTTPError(500, err)
	}
	trendingWordsApi := ConvertTrendingWords(trendingWordsModel)
	return ctx.JSON(200, trendingWordsApi)
}

// ある月のトレンド
// (GET /trend/month/{month})
func (s Server) GetTrendingWordsForMonth(ctx echo.Context, month string, params api.GetTrendingWordsForMonthParams) error {
	limitNumber := 10
	if params.Limit != nil {
		limitNumber = *params.Limit
	}
	trendingWordsModel, err := model.GetTrendingWordsForMonth(limitNumber, month)
	if err != nil {
		return echo.NewHTTPError(500, err)
	}
	trendingWordsApi := ConvertTrendingWords(trendingWordsModel)
	return ctx.JSON(200, trendingWordsApi)
}

// ある年のトレンド
// (GET /trend/year/{year})
func (s Server) GetTrendingWordsForYear(ctx echo.Context, year string, params api.GetTrendingWordsForYearParams) error {
	limitNumber := 10
	if params.Limit != nil {
		limitNumber = *params.Limit
	}
	trendingWordsModel, err := model.GetTrendingWordsForYear(limitNumber, year)
	if err != nil {
		return echo.NewHTTPError(500, err)
	}
	trendingWordsApi := ConvertTrendingWords(trendingWordsModel)
	return ctx.JSON(200, trendingWordsApi)
}

// model.TrendingWordsからapi.TrendingWordsへの型の変換
func ConvertTrendingWords(models model.TrendingWords) api.TrendingWords {
	TrendingWordsSlice := make(api.TrendingWords, len(models))
	for i, m := range models {
		TrendingWordsSlice[i] = api.TrendingWord{
			Number: m.Number,
			Word:   m.Word,
		}
	}
	return TrendingWordsSlice
}
