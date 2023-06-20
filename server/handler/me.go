package handler

import (
	"h23s_15/api"
	"h23s_15/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
)

// 自分の投稿に対する通知の設定
// (PUT /words/me/)
func (s Server) PutWordsMe(ctx echo.Context) error {
	data := &api.PutWordsMeJSONRequestBody{}
	err := ctx.Bind(data)
	if err != nil {
		slog.Info("Error creating PutWordsMe JSON request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	model.PutWordsMe(data.IncludeMe, data.Word)
	return nil
}

// 自分の投稿に対する通知の一括設定
// (POST /words/me/all)
func (s Server) PostWordsMeAll(ctx echo.Context) error {
	data := &api.PostWordsMeAllJSONRequestBody{}
	err := ctx.Bind(data)
	if err != nil {
		slog.Info("Error creating PostWordsMeAll JSON request body: %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	// model.PostWordsMeAll(data.IncludeMe)
	return nil
}

// func ConvertWordMeSetting(model.WordMeSetting) (api.WordMeSetting, error) {
// 	return nil, nil
// }
