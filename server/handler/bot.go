package handler

import "github.com/labstack/echo/v4"

// bot投稿に対する通知の設定
// (PUT /words)
func (s Server) PutWords(ctx echo.Context) error {
	return nil
}

// bot投稿に対する通知の一括設定
// (POST /words/bot)
func (s Server) PostWordsBot(ctx echo.Context) error {
	return nil
}
