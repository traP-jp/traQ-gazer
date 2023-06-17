package handler

import (
	"h23s_15/api"
	"h23s_15/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// bot投稿に対する通知の設定
// (PUT /words)
func (s Server) PutWords(ctx echo.Context) error {
	data := &api.WordBotSetting{}
	err := ctx.Bind(data)
	
	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// traPIdの取得
	userId, err := getUserIdFromSession(ctx)
	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	exist, err := model.ExistWord(data.Word, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if !exist {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	err = model.ChengeBotNotification(data.Word, data.IncludeBot, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	
	return ctx.JSON(http.StatusOK, "Successful Change")
}

// bot投稿に対する通知の一括設定
// (POST /words/bot)
func (s Server) PostWordsBot(ctx echo.Context) error {
	data := &api.Bot{}
	err := ctx.Bind(data)
	
	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	// traPIdの取得
	userId, err := getUserIdFromSession(ctx)
	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	err = model.ChangeAllBotNotification(data.IncludeBot, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, "Successful Change")
}
