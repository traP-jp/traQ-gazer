// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// sessionからuserIdを取る
type Bot struct {
	IncludeBot bool `json:"includeBot"`
}

// sessionからuserIdを取る
type Me struct {
	IncludeMe bool `json:"includeMe"`
}

// RecommendedWord defines model for RecommendedWord.
type RecommendedWord struct {
	Number int    `json:"number"`
	Word   string `json:"word"`
}

// RecommendedWords defines model for RecommendedWords.
type RecommendedWords = []RecommendedWord

// SimilarUser defines model for SimilarUser.
type SimilarUser struct {
	UserId string `json:"userId"`
}

// SimilarUsers defines model for SimilarUsers.
type SimilarUsers = []SimilarUser

// TrendingWord defines model for TrendingWord.
type TrendingWord struct {
	Number int    `json:"number"`
	Word   string `json:"word"`
}

// TrendingWords defines model for TrendingWords.
type TrendingWords = []TrendingWord

// UserListItem defines model for UserListItem.
type UserListItem struct {
	IncludeBot bool      `json:"includeBot"`
	IncludeMe  bool      `json:"includeMe"`
	Time       time.Time `json:"time"`
	UserId     string    `json:"userId"`
}

// UsersList defines model for UsersList.
type UsersList = []UserListItem

// UsersOfWordListItem defines model for UsersOfWordListItem.
type UsersOfWordListItem struct {
	UserIds []UserListItem `json:"userIds"`
	Word    string         `json:"word"`
}

// UsersOfWordsList defines model for UsersOfWordsList.
type UsersOfWordsList = []UsersOfWordListItem

// WordAllListItem defines model for WordAllListItem.
type WordAllListItem struct {
	IncludeBot bool      `json:"includeBot"`
	IncludeMe  bool      `json:"includeMe"`
	Time       time.Time `json:"time"`
	UserId     string    `json:"userId"`
	Word       string    `json:"word"`
}

// sessionからuserIdを取る
type WordBotSetting struct {
	IncludeBot bool   `json:"includeBot"`
	Word       string `json:"word"`
}

// sessionからuserIdを取る
type WordDelete struct {
	Word string `json:"word"`
}

// WordListItem defines model for WordListItem.
type WordListItem struct {
	IncludeBot bool      `json:"includeBot"`
	IncludeMe  bool      `json:"includeMe"`
	Time       time.Time `json:"time"`
	Word       string    `json:"word"`
}

// sessionからuserIdを取る
type WordMeSetting struct {
	IncludeMe bool   `json:"includeMe"`
	Word      string `json:"word"`
}

// sessionからuserIdを取る
type WordRequest struct {
	IncludeBot bool   `json:"includeBot"`
	IncludeMe  bool   `json:"includeMe"`
	Word       string `json:"word"`
}

// WordsAllList defines model for WordsAllList.
type WordsAllList = []WordAllListItem

// WordsList defines model for WordsList.
type WordsList = []WordListItem

// WordsOfUserListItem defines model for WordsOfUserListItem.
type WordsOfUserListItem struct {
	UserId string         `json:"userId"`
	Words  []WordListItem `json:"words"`
}

// WordsOfUsersList defines model for WordsOfUsersList.
type WordsOfUsersList = []WordsOfUserListItem

// GetTodayTrendingWordsParams defines parameters for GetTodayTrendingWords.
type GetTodayTrendingWordsParams struct {
	// 返すwordの数
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetTrendingWordsForDayParams defines parameters for GetTrendingWordsForDay.
type GetTrendingWordsForDayParams struct {
	// 返すwordの数
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetTrendingWordsForMonthParams defines parameters for GetTrendingWordsForMonth.
type GetTrendingWordsForMonthParams struct {
	// 返すwordの数
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// GetTrendingWordsForYearParams defines parameters for GetTrendingWordsForYear.
type GetTrendingWordsForYearParams struct {
	// 返すwordの数
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}

// DeleteWordsJSONBody defines parameters for DeleteWords.
type DeleteWordsJSONBody = WordDelete

// PostWordsJSONBody defines parameters for PostWords.
type PostWordsJSONBody = WordRequest

// PutWordsJSONBody defines parameters for PutWords.
type PutWordsJSONBody = WordBotSetting

// PostWordsBotJSONBody defines parameters for PostWordsBot.
type PostWordsBotJSONBody = Bot

// PutWordsMeJSONBody defines parameters for PutWordsMe.
type PutWordsMeJSONBody = WordMeSetting

// PostWordsMeAllJSONBody defines parameters for PostWordsMeAll.
type PostWordsMeAllJSONBody = Me

// DeleteWordsJSONRequestBody defines body for DeleteWords for application/json ContentType.
type DeleteWordsJSONRequestBody = DeleteWordsJSONBody

// PostWordsJSONRequestBody defines body for PostWords for application/json ContentType.
type PostWordsJSONRequestBody = PostWordsJSONBody

// PutWordsJSONRequestBody defines body for PutWords for application/json ContentType.
type PutWordsJSONRequestBody = PutWordsJSONBody

// PostWordsBotJSONRequestBody defines body for PostWordsBot for application/json ContentType.
type PostWordsBotJSONRequestBody = PostWordsBotJSONBody

// PutWordsMeJSONRequestBody defines body for PutWordsMe for application/json ContentType.
type PutWordsMeJSONRequestBody = PutWordsMeJSONBody

// PostWordsMeAllJSONRequestBody defines body for PostWordsMeAll for application/json ContentType.
type PostWordsMeAllJSONRequestBody = PostWordsMeAllJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// アクセスしているuserのwordたち
	// (GET /list/user/me)
	GetListUserMe(ctx echo.Context) error
	// あるuserのwordたち
	// (GET /list/user/{userId})
	GetListUserUserId(ctx echo.Context, userId string) error
	// あるuserのwordたちを登録しているuserたち
	// (GET /list/user/{userId}/users)
	GetListUserUserIdUsers(ctx echo.Context, userId string) error
	// ある単語を見ているuserたち
	// (GET /list/word/{word})
	GetListWordWord(ctx echo.Context, word string) error
	// あるwordのuserたちが登録しているwordたち
	// (GET /list/word/{word}/words)
	GetListWordWordWords(ctx echo.Context, word string) error
	// 似たような者を探す
	// (GET /similar/{userId})
	GetUsersWithSimilarWords(ctx echo.Context, userId string) error
	// おすすめの単語を出す
	// (GET /similar/{userId}/recommend)
	GetRecommendedWordsForUser(ctx echo.Context, userId string) error
	// 今日のトレンド
	// (GET /trend/day/today)
	GetTodayTrendingWords(ctx echo.Context, params GetTodayTrendingWordsParams) error
	// ある日のトレンド
	// (GET /trend/day/{day})
	GetTrendingWordsForDay(ctx echo.Context, day string, params GetTrendingWordsForDayParams) error
	// ある月のトレンド
	// (GET /trend/month/{month})
	GetTrendingWordsForMonth(ctx echo.Context, month string, params GetTrendingWordsForMonthParams) error
	// ある年のトレンド
	// (GET /trend/year/{year})
	GetTrendingWordsForYear(ctx echo.Context, year string, params GetTrendingWordsForYearParams) error
	// wordの削除
	// (DELETE /words)
	DeleteWords(ctx echo.Context) error
	// 全データの取得
	// (GET /words)
	GetWords(ctx echo.Context) error
	// wordの登録
	// (POST /words)
	PostWords(ctx echo.Context) error
	// bot投稿に対する通知の設定
	// (PUT /words)
	PutWords(ctx echo.Context) error
	// bot投稿に対する通知の一括設定
	// (POST /words/bot)
	PostWordsBot(ctx echo.Context) error
	// 自分の投稿に対する通知の設定
	// (PUT /words/me/)
	PutWordsMe(ctx echo.Context) error
	// 自分の投稿に対する通知の一括設定
	// (POST /words/me/all)
	PostWordsMeAll(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetListUserMe converts echo context to params.
func (w *ServerInterfaceWrapper) GetListUserMe(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetListUserMe(ctx)
	return err
}

// GetListUserUserId converts echo context to params.
func (w *ServerInterfaceWrapper) GetListUserUserId(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetListUserUserId(ctx, userId)
	return err
}

// GetListUserUserIdUsers converts echo context to params.
func (w *ServerInterfaceWrapper) GetListUserUserIdUsers(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetListUserUserIdUsers(ctx, userId)
	return err
}

// GetListWordWord converts echo context to params.
func (w *ServerInterfaceWrapper) GetListWordWord(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "word" -------------
	var word string

	err = runtime.BindStyledParameterWithLocation("simple", false, "word", runtime.ParamLocationPath, ctx.Param("word"), &word)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter word: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetListWordWord(ctx, word)
	return err
}

// GetListWordWordWords converts echo context to params.
func (w *ServerInterfaceWrapper) GetListWordWordWords(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "word" -------------
	var word string

	err = runtime.BindStyledParameterWithLocation("simple", false, "word", runtime.ParamLocationPath, ctx.Param("word"), &word)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter word: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetListWordWordWords(ctx, word)
	return err
}

// GetUsersWithSimilarWords converts echo context to params.
func (w *ServerInterfaceWrapper) GetUsersWithSimilarWords(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUsersWithSimilarWords(ctx, userId)
	return err
}

// GetRecommendedWordsForUser converts echo context to params.
func (w *ServerInterfaceWrapper) GetRecommendedWordsForUser(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "userId" -------------
	var userId string

	err = runtime.BindStyledParameterWithLocation("simple", false, "userId", runtime.ParamLocationPath, ctx.Param("userId"), &userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter userId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetRecommendedWordsForUser(ctx, userId)
	return err
}

// GetTodayTrendingWords converts echo context to params.
func (w *ServerInterfaceWrapper) GetTodayTrendingWords(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTodayTrendingWordsParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTodayTrendingWords(ctx, params)
	return err
}

// GetTrendingWordsForDay converts echo context to params.
func (w *ServerInterfaceWrapper) GetTrendingWordsForDay(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "day" -------------
	var day string

	err = runtime.BindStyledParameterWithLocation("simple", false, "day", runtime.ParamLocationPath, ctx.Param("day"), &day)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter day: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTrendingWordsForDayParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTrendingWordsForDay(ctx, day, params)
	return err
}

// GetTrendingWordsForMonth converts echo context to params.
func (w *ServerInterfaceWrapper) GetTrendingWordsForMonth(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "month" -------------
	var month string

	err = runtime.BindStyledParameterWithLocation("simple", false, "month", runtime.ParamLocationPath, ctx.Param("month"), &month)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter month: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTrendingWordsForMonthParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTrendingWordsForMonth(ctx, month, params)
	return err
}

// GetTrendingWordsForYear converts echo context to params.
func (w *ServerInterfaceWrapper) GetTrendingWordsForYear(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "year" -------------
	var year string

	err = runtime.BindStyledParameterWithLocation("simple", false, "year", runtime.ParamLocationPath, ctx.Param("year"), &year)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter year: %s", err))
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params GetTrendingWordsForYearParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetTrendingWordsForYear(ctx, year, params)
	return err
}

// DeleteWords converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteWords(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteWords(ctx)
	return err
}

// GetWords converts echo context to params.
func (w *ServerInterfaceWrapper) GetWords(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetWords(ctx)
	return err
}

// PostWords converts echo context to params.
func (w *ServerInterfaceWrapper) PostWords(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostWords(ctx)
	return err
}

// PutWords converts echo context to params.
func (w *ServerInterfaceWrapper) PutWords(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutWords(ctx)
	return err
}

// PostWordsBot converts echo context to params.
func (w *ServerInterfaceWrapper) PostWordsBot(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostWordsBot(ctx)
	return err
}

// PutWordsMe converts echo context to params.
func (w *ServerInterfaceWrapper) PutWordsMe(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PutWordsMe(ctx)
	return err
}

// PostWordsMeAll converts echo context to params.
func (w *ServerInterfaceWrapper) PostWordsMeAll(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PostWordsMeAll(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/list/user/me", wrapper.GetListUserMe)
	router.GET(baseURL+"/list/user/:userId", wrapper.GetListUserUserId)
	router.GET(baseURL+"/list/user/:userId/users", wrapper.GetListUserUserIdUsers)
	router.GET(baseURL+"/list/word/:word", wrapper.GetListWordWord)
	router.GET(baseURL+"/list/word/:word/words", wrapper.GetListWordWordWords)
	router.GET(baseURL+"/similar/:userId", wrapper.GetUsersWithSimilarWords)
	router.GET(baseURL+"/similar/:userId/recommend", wrapper.GetRecommendedWordsForUser)
	router.GET(baseURL+"/trend/day/today", wrapper.GetTodayTrendingWords)
	router.GET(baseURL+"/trend/day/:day", wrapper.GetTrendingWordsForDay)
	router.GET(baseURL+"/trend/month/:month", wrapper.GetTrendingWordsForMonth)
	router.GET(baseURL+"/trend/year/:year", wrapper.GetTrendingWordsForYear)
	router.DELETE(baseURL+"/words", wrapper.DeleteWords)
	router.GET(baseURL+"/words", wrapper.GetWords)
	router.POST(baseURL+"/words", wrapper.PostWords)
	router.PUT(baseURL+"/words", wrapper.PutWords)
	router.POST(baseURL+"/words/bot", wrapper.PostWordsBot)
	router.PUT(baseURL+"/words/me/", wrapper.PutWordsMe)
	router.POST(baseURL+"/words/me/all", wrapper.PostWordsMeAll)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xaW28bxxX+K8S0QF9orxynL3yzIaQQWsJB7MAQEqFYcYfiBrs765lZtYRAQLtKbOvS",
	"JnDqOHLTKA0SVbBTplcgit3ox4xIUf+iOLO75F65S4ms1MAvvO3OmXO+7ztn58xwDTWIaRMLW5yh2hpi",
	"jRY2VfnxJuHwpmHWoLrNdWKhGmKYMZ1Ywt0W3qbDMF3QhPeo9+EnwttGVWRTYmPKdSwt6FbDcDQcGMK/",
	"VU3bwKjGqYOriLdtjGpomRADqxbqdKqI4nuOTrGGau9Exy4NbybL7+EGR50qquPp+ObbOZNrdZzp2Vu4",
	"QUwTWxrW7hKqgfn4zJZjLmMam/ba0I5ucbyCKRj6TTB6eNfw0/BuxqluraQcDGYITJRw0keEY1N++CnF",
	"TVRDP1FG0lACXSjJ6DpD4yqlahu+39ZN3VDp28wPMh68T0s8LE5VuzCmYGBWMJEJywcS9TIjiDsUW5pu",
	"rVxeCqMelg87FldG3IDHr3TGFzg203HHEzqZJtV4UqUvc92UV5qEmipHNaSpHF+Rv6bwqJ5bKsF81ajX",
	"1YLklRqC+EsDGgMsB1B2qwl45+Pqe8ymNmmouzI4sXEii3g/OSzJsDMches3DOP/TnGTQRwifBZFAkI3",
	"Cb+NOYcJpvlATsNVLqggloInNDg+jw3Mz/uknsCpPD8uib4mgvdsUqnjqSolL+6zCGWM12/hew5mfMbq",
	"nkFMZWlhQZErXT+ThTGndrKJjRZavNUcvwaYtEiyKbmXwQM8usasCyPhTI5TEoe0P1JRTZIW7Z2b8zJ/",
	"ucSm9dp19utrP/8ZqzBMV/UGJPUqpsy/99rVuatzYJzY2FJtHdXQ9atzV6+DuFXekq4qhs64AoEqfrVZ",
	"wRmpIrwvhfet8F4I71C4T4S7L9z3hbcN44TbPdl9cbrzj+gFQFC43d6Hn/R+eIKkB1QFY8Au+gXmEDtg",
	"UAeXKWY2sZivgdfm5uCtQSyOLemLatuG3pDDlfcYOBT2kqXAlvRISONB3XYaDcxY0zEqFHOq41XVkFJg",
	"jmmqtF0ctx/lnnD/DJyoKwzEA4CiJbATwXbNl1InF+HzI/l2uBywVaqamMuG5Z3kPAvzFdKs8BauwIyy",
	"vqCalAOqIksFDYwWFqOc8HvXEejJOrZ0uUisotfnXk+DDBBVLMIrTeJYWpJr1zsPs/Iby8+giPWdNMvC",
	"3RXuofwazN0V7p+EtyPcz+HVe5QeIyecQBp+D/vj00eqj7gImeQyVKgisKGswWunQD293306ePb5mJkK",
	"pAAI3fWXGGM1cKeFK+BPhZMKwypttCpNkqOFYMlyiZQwFQkATAUSGLIx2N8+M+PKcBUzhnepsizWo1Uj",
	"UECsamRVGjlhaancDVZCP069pBZx/wPZBNV9JJRclsbpiPk7jMXripPNw1736XDCg95HO8L9VHgPhXtf",
	"uM/OU1IkbHd13gq2O0OtzIyu2F5sBlW3fplA/PjlSwAxDHaw/oHwHvV//6VwdyOwBliipc4FPRuzCFVo",
	"uBk+M2ozZBeMGcd6cnf/DeLvc8+Q99SBQhnuhbsly+Ou8FwIKESj9+D7S0g/p9jSFE1tK5xoajuX8+MX",
	"W/0nX/c/WxeeNzj6T2/rC+E+lqV+r7BRuAOW4zv7BREPjv4g3N3AcP/x38KY7zmYtkdBG7qpcxSNUcNN",
	"1TE4ql2bSx9dzLSkx+MrVSQkpPDo3HgoNr4RG/8UG5sRfUhqgrI7omlNU9slqq60/PyMfEVDeYPQebX9",
	"irDhc7QUZwVw3bZxQ2/qjYqmtiu6JXPb322tvIsWFxcXr9TrV+bn30XZya5JQvIzPdjDkh6NrEW2k7Jq",
	"gEks3lLW5FsZgX32cGoCq8OcryQ2khhgOy2JSULzRJanMDNgZCKNFQisjWGZAa8l5NU7/NfU5LWIVfpK",
	"XaP+EaCdlrqAzkxx5Smr7ZNRXli5qhp2sFrOgVwolM2t092vUkLxj/FGvYM8rLlJtPZUu7zgsLAT3+mH",
	"mDvZ0shtAmWU8KvsATPuXbBWVUPXKrplOzzBfBKKkG4fQ4AzMx17HxyIjQdi46Xwjsam3MxbsNh509n2",
	"1POCSWNhk6xjuwBDv3lJgfAmYXzWYgpPFKegJopXdMZ9/8+jqCEaGSg6WYr66H1/5PPetz/4HeAy4cLd",
	"6W89Pjk4kr/sCc+TO0rPhfsV9IfuvvC80/WnJ3tfy75qW7jb/b9/Dw2WPDFNMOHMnIjIfxcuOrOXCQ+h",
	"8yEFfEKsuoODv/a6TyP0LJNwJ0cSpSz758rZis9jBjIJOvcRP+XJCdPEP2SeBUFgeeqsTLpbP56W4+/W",
	"+9vfFJJjYkWSUzKRBg+e9R7en0kuBaels8mm0d87LjqZQgS7k6SUiZOkqYaRn1RjWJpGXtXxDcOYEVl1",
	"fPGJVYainPTyiQJrmK6G61qHGqiGFNXWEazWg3vT5yEfQxZ5/4blw8bH8lS+2//isHcIH268uXD6x/vx",
	"swaGOtXMetpNKCljPNSB9Ohh5MUGTJwx/vi79cH+X4JzI5DTnr9DmRwsd/vTw6MNg6woOcP9DiJjfHxf",
	"NHd8uC3aWer8NwAA//94fN0ZyC8AAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
