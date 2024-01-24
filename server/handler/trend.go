package handler

import (
	"traQ-gazer/api"
	"traQ-gazer/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

const LIMIT_DEFAULT = 10

// 今日のトレンド
// (GET /trend/day/today)
func (s Server) GetTodayTrendingWords(ctx echo.Context, params api.GetTodayTrendingWordsParams) error {
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// limit(初期値LIMIT_DEFAULT)を更新
	limit := LIMIT_DEFAULT
	if params.Limit != nil {
		limit = *params.Limit
	}

	trends, err := model.GetTrendToday(limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, ConvertSliceTrendingWord(trends))
}

// ある日のトレンド
// (GET /trend/day/{day})
func (s Server) GetTrendingWordsForDay(ctx echo.Context, day string, params api.GetTrendingWordsForDayParams) error {
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// limit(初期値LIMIT_DEFAULT)を更新
	limit := LIMIT_DEFAULT
	if params.Limit != nil {
		limit = *params.Limit
	}
	trends, err := model.GetTrendOneday(day, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, ConvertSliceTrendingWord(trends))
}

// ある月のトレンド
// (GET /trend/month/{month})
func (s Server) GetTrendingWordsForMonth(ctx echo.Context, month string, params api.GetTrendingWordsForMonthParams) error {
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// limit(初期値LIMIT_DEFAULT)を更新
	limit := LIMIT_DEFAULT
	if params.Limit != nil {
		limit = *params.Limit
	}

	trends, err := model.GetTrendOneMonth(month, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, ConvertSliceTrendingWord(trends))
}

// ある年のトレンド
// (GET /trend/year/{year})
func (s Server) GetTrendingWordsForYear(ctx echo.Context, year string, params api.GetTrendingWordsForYearParams) error {
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// limit(初期値LIMIT_DEFAULT)を更新
	limit := LIMIT_DEFAULT
	if params.Limit != nil {
		limit = *params.Limit
	}

	trends, err := model.GetTrendOneYear(year, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, ConvertSliceTrendingWord(trends))
}

// model.TrendingWordからapi.WordsAllListへの型の変換
func ConvertSliceTrendingWord(modelList model.TrendingWords) api.TrendingWords {
	apiList := make([]api.TrendingWord, len(modelList))
	for i, WordType := range modelList {
		apiList[i] = api.TrendingWord{
			Number: WordType.Number,
			Word:   WordType.Word,
		}
	}
	return apiList
}
