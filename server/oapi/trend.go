package oapi

import (
	"net/http"
	"traQ-gazer/model"
	"traQ-gazer/repo"

	"github.com/labstack/echo/v4"
)

const LIMIT_DEFAULT = 10

// 今日のトレンド
// (GET /trend/day/today)
func (s Server) GetTodayTrendingWords(ctx echo.Context, params GetTodayTrendingWordsParams) error {
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// limit(初期値LIMIT_DEFAULT)を更新
	limit := LIMIT_DEFAULT
	if params.Limit != nil {
		limit = *params.Limit
	}

	trends, err := repo.GetTrendToday(limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, ConvertSliceTrendingWord(trends))
}

// ある日のトレンド
// (GET /trend/day/{day})
func (s Server) GetTrendingWordsForDay(ctx echo.Context, day string, params GetTrendingWordsForDayParams) error {
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// limit(初期値LIMIT_DEFAULT)を更新
	limit := LIMIT_DEFAULT
	if params.Limit != nil {
		limit = *params.Limit
	}
	trends, err := repo.GetTrendOneday(day, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, ConvertSliceTrendingWord(trends))
}

// ある月のトレンド
// (GET /trend/month/{month})
func (s Server) GetTrendingWordsForMonth(ctx echo.Context, month string, params GetTrendingWordsForMonthParams) error {
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// limit(初期値LIMIT_DEFAULT)を更新
	limit := LIMIT_DEFAULT
	if params.Limit != nil {
		limit = *params.Limit
	}

	trends, err := repo.GetTrendOneMonth(month, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, ConvertSliceTrendingWord(trends))
}

// ある年のトレンド
// (GET /trend/year/{year})
func (s Server) GetTrendingWordsForYear(ctx echo.Context, year string, params GetTrendingWordsForYearParams) error {
	err := ctx.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// limit(初期値LIMIT_DEFAULT)を更新
	limit := LIMIT_DEFAULT
	if params.Limit != nil {
		limit = *params.Limit
	}

	trends, err := repo.GetTrendOneYear(year, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, ConvertSliceTrendingWord(trends))
}

// model.TrendingWordからoapi.WordsAllListへの型の変換
func ConvertSliceTrendingWord(modelList model.TrendingWords) TrendingWords {
	oapiTrendingWords := make([]TrendingWord, len(modelList))
	for i, WordType := range modelList {
		oapiTrendingWords[i] = TrendingWord{
			Number: WordType.Number,
			Word:   WordType.Word,
		}
	}
	return oapiTrendingWords
}
