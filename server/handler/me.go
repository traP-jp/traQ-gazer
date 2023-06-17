package handler

import (
	"h23s_15/api"
	"h23s_15/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// 自分の投稿に対する通知の設定
// (PUT /words/me/)
func (s Server) PutWordsMe(ctx echo.Context) error {
	data := &api.PutWordsMeJSONRequestBody{}
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	userId, err := getUserIdFromSession(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = model.PutWordsMe(data.IncludeMe, data.Word, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, nil)
}

// 自分の投稿に対する通知の一括設定
// (POST /words/me/all)
func (s Server) PostWordsMeAll(ctx echo.Context) error {
	data := &api.PostWordsMeAllJSONRequestBody{}
	if err := ctx.Bind(data); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	userId, err := getUserIdFromSession(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	err = model.PostWordsMeAll(data.IncludeMe, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	return ctx.JSON(http.StatusOK, nil)
}
