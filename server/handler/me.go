package handler

import "github.com/labstack/echo/v4"

// 自分の投稿に対する通知の設定
// (PUT /words/me/)
func (s Server) PutWordsMe(ctx echo.Context) error {
	return nil
}

// 自分の投稿に対する通知の一括設定
// (POST /words/me/all)
func (s Server) PostWordsMeAll(ctx echo.Context) error {
	return nil
}
