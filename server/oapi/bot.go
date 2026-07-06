package oapi

import (
	"net/http"
	"traQ-gazer/repo"

	"github.com/labstack/echo/v5"
)

// bot投稿に対する通知の設定
// (PUT /words)
func (s Server) PutWords(ctx *echo.Context) error {
	data := &WordBotSetting{}
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

	err = repo.ChengeBotNotification(data.Word, data.IncludeBot, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Successful Change")
}

// bot投稿に対する通知の一括設定
// (POST /words/bot)
func (s Server) PostWordsBot(ctx *echo.Context) error {
	data := &Bot{}
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

	err = repo.ChangeAllBotNotification(data.IncludeBot, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Successful Change")
}
