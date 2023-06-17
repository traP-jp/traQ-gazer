package handler

import (
	"h23s_15/api"
	"h23s_15/model"

	"github.com/labstack/echo/v4"
)

// 今日のトレンド
// (GET /trend/day/today)
func (s Server) GetTodayTrendingWords(ctx echo.Context, params api.GetTodayTrendingWordsParams) error {
	trendingWordsModel, err := model.GetTodayTrendingWords(*params.Limit)
	if err != nil {
		return echo.NewHTTPError(500, err)
	}
	trendingWordsApi := ConvertTrendingWords(trendingWordsModel)
	return ctx.JSON(200, trendingWordsApi)
}

// ある日のトレンド
// (GET /trend/day/{day})
func (s Server) GetTrendingWordsForDay(ctx echo.Context, day string, params api.GetTrendingWordsForDayParams) error {
	trendingWordsModel, err := model.GetTrendingWordsForDay(*params.Limit, day)
	if err != nil {
		return echo.NewHTTPError(500, err)
	}
	trendingWordsApi := ConvertTrendingWords(trendingWordsModel)
	return ctx.JSON(200, trendingWordsApi)
}

// ある月のトレンド
// (GET /trend/month/{month})
func (s Server) GetTrendingWordsForMonth(ctx echo.Context, month string, params api.GetTrendingWordsForMonthParams) error {
	trendingWordsModel, err := model.GetTrendingWordsForMonth(*params.Limit, month)
	if err != nil {
		return echo.NewHTTPError(500, err)
	}
	trendingWordsApi := ConvertTrendingWords(trendingWordsModel)
	return ctx.JSON(200, trendingWordsApi)
}

// ある年のトレンド
// (GET /trend/year/{year})
func (s Server) GetTrendingWordsForYear(ctx echo.Context, year string, params api.GetTrendingWordsForYearParams) error {
	trendingWordsModel, err := model.GetTrendingWordsForYear(*params.Limit, year)
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
