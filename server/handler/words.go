package handler

import (
	"h23s_15/api"
	"h23s_15/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// wordの登録
// (POST /words)
func (s Server) PostWords(ctx echo.Context) error {
	return nil
}

// wordの削除
// (DELETE /words)
func (s Server) DeleteWords(ctx echo.Context) error {
	return nil
}

// 全データの取得
// (GET /words)
func (s Server) GetWords(ctx echo.Context) error {
	wordlist, err := model.GetWords()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	change := ConvertSliceToA1(wordlist)
	return ctx.JSON(http.StatusOK, change)
}

// model.WordsAllListからapi.WordsAllListへの型の変換
func ConvertSliceToA1(WordsListSlice model.WordsAllList) api.WordsAllList {
	WordsAllListSlice := make([]api.WordAllListItem, len(WordsListSlice))
	for i, WordType := range WordsListSlice {
		WordsAllListSlice[i] = api.WordAllListItem{
			IncludeBot: WordType.IncludeBot,
			IncludeMe:  WordType.IncludeMe,
			Time:       WordType.Time,
			UserId:     WordType.UserId,
			Word:       WordType.Word,
		}
	}
	return WordsAllListSlice
}
