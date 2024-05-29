package oapi

import (
	"net/http"
	"traQ-gazer/repo"

	"github.com/labstack/echo/v4"
)

// 自分の投稿に対する通知の設定
// (PUT /words/me/)
func (s Server) PutWordsMe(ctx echo.Context) error {
	data := &WordMeSetting{}
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

	exist, err := repo.ExistWord(data.Word, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if !exist {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	err = repo.ChengeMeNotification(data.Word, data.IncludeMe, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, "Successful Change")
}

// 自分の投稿に対する通知の一括設定
// (POST /words/me/all)
func (s Server) PostWordsMeAll(ctx echo.Context) error {
	return nil
}
