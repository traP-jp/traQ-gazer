package handler

import (
	"h23s_15/api"
	"h23s_15/model"

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
		return ctx.JSON(400, err)
	}

	// traPIdの取得
	userId, err := getUserIdFromSession(ctx)
	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return ctx.JSON(400, err)
	}

	err = model.PostWords(data, userId)
	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return ctx.JSON(400, err)
	}
	
	return ctx.JSON(200, "Successful registration")
}

// wordの削除
// (DELETE /words)
func (s Server) DeleteWords(ctx echo.Context) error {

	data := &api.DeleteWordsJSONRequestBody{}
	err := ctx.Bind(data)

	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return ctx.JSON(400, err)
	}

	// traPIdの取得
	userId, err := getUserIdFromSession(ctx)
	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return ctx.JSON(400, err)
	}

	err = model.DeleteWords(data, userId)
	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return ctx.JSON(400, err)
	}

	return ctx.JSON(200, "Successful deletion")
}

// 全データの取得
// (GET /words)
func (s Server) GetWords(ctx echo.Context) error {
	return nil
}
