package handler

import "github.com/labstack/echo/v4"

// アクセスしているuserのwordたち
// (GET /list/user/me)
func (s Server) GetListUserMe(ctx echo.Context) error {
	return nil
}

// あるuserのwordたち
// (GET /list/user/{userId})
func (s Server) GetListUserUserId(ctx echo.Context, userId string) error {
	return nil
}

// あるuserのwordたちを登録しているuserたち
// (GET /list/user/{userId}/users)
func (s Server) GetListUserUserIdUsers(ctx echo.Context, userId string) error {
	return nil
}

// ある単語を見ているuserたち
// (GET /list/word/{word})
func (s Server) GetListWordWord(ctx echo.Context, word string) error {
	return nil
}

// あるwordのuserたちが登録しているwordたち
// (GET /list/word/{word}/words)
func (s Server) GetListWordWordWords(ctx echo.Context, word string) error {
	return nil
}
