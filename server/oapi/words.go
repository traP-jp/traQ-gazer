package oapi

import (
	"net/http"
	"traQ-gazer/model"
	"traQ-gazer/repo"
	"traQ-gazer/wordpattern"

	"github.com/labstack/echo/v5"
)

// wordの登録
// (POST /words)
func (s Server) PostWords(ctx *echo.Context) error {

	// Wordの取得
	data := &PostWordsJSONRequestBody{}
	err := ctx.Bind(data)
	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := wordpattern.ValidateRegisteredWord(data.Word); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// traPIdの取得
	userId, err := getUserIdFromSession(ctx)
	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	exist, err := repo.ExistWord(data.Word, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if exist {
		// 同じword user_idが存在する
		return echo.NewHTTPError(http.StatusBadRequest, "Already Resistered")
	}

	err = repo.ResisterWord(data.Word, data.IncludeBot, data.IncludeMe, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Successful registration")
}

// wordの削除
// (DELETE /words)
func (s Server) DeleteWords(ctx *echo.Context) error {

	data := &DeleteWordsJSONRequestBody{}
	err := ctx.Bind(data)

	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// traPIdの取得
	userId, err := getUserIdFromSession(ctx)
	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	exist, err := repo.ExistWord(data.Word, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if !exist {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	err = repo.DeleteWord(data.Word, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Successful deletion")
}

// 全データの取得
// (GET /words)
func (s Server) GetWords(ctx *echo.Context) error {
	return echo.NewHTTPError(http.StatusGone, "this endpoint is not available")
}

// model.WordsAllListからoapi.WordsAllListへの型の変換
func convertSliceToA1(WordsListSlice model.WordsAllList) WordsAllList {
	WordsAllListSlice := make([]WordAllListItem, len(WordsListSlice))
	for i, WordType := range WordsListSlice {
		WordsAllListSlice[i] = WordAllListItem{
			IncludeBot: WordType.IncludeBot,
			IncludeMe:  WordType.IncludeMe,
			Time:       WordType.Time,
			UserId:     WordType.UserId,
			Word:       WordType.Word,
		}
	}
	return WordsAllListSlice
}
