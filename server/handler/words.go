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

	// Wordの取得
	data := &api.PostWordsJSONRequestBody{}
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

	if exist {
		// 同じword user_idが存在する
		return echo.NewHTTPError(http.StatusBadRequest, "Already Resistered")
	}

	err = model.ResisterWord(data.Word, data.IncludeBot, data.IncludeMe, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, "Successful registration")
}

// wordの削除
// (DELETE /words)
func (s Server) DeleteWords(ctx echo.Context) error {

	data := &api.DeleteWordsJSONRequestBody{}
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

	err = model.DeleteWord(data.Word, userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, "Successful deletion")
}

// 全データの取得
// (GET /words)
func (s Server) GetWords(ctx echo.Context) error {
	return nil
}
