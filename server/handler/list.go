package handler

import (
	"h23s_15/api"
	"h23s_15/model"

	"github.com/labstack/echo/v4"
)

// あるuserのwordたち
// (GET /list/user/{userId})
func (s Server) GetListUserUserId(ctx echo.Context, userId string) error {
	wordsListMode, err := model.GetListUserUserId(userId)
	if err != nil {
		return err
	}
	wordsListApi := ConvertWordList(wordsListMode)
	return ctx.JSON(200, wordsListApi)
}

// あるuserのwordたちを登録しているuserたち
// (GET /list/user/{userId}/users)
func (s Server) GetListUserUserIdUsers(ctx echo.Context, userId string) error {
	usersOfWordsListMode, err := model.GetListUserUserIdUsers(userId)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	usersOfWordsListApi := ConvertUsersOfWordsList(usersOfWordsListMode)
	return ctx.JSON(200, usersOfWordsListApi)
}

// model.WordsListからapi.WordsListへの型の変換
func ConvertWordList(models model.WordsList) api.WordsList {
	WordsListSlice := make(api.WordsList, len(models))
	for i, s := range models {
		WordsListSlice[i] = api.WordListItem{
			IncludeBot: s.IncludeBot,
			IncludeMe:  s.IncludeMe,
			Time:       s.Time,
			Word:       s.Word,
		}
	}
	return WordsListSlice
}

// ある単語を見ているuserたち
// (GET /list/word/{word})
func (s Server) GetListWordWord(ctx echo.Context, word string) error {
	usersListMode, err := model.GetListWordWord(word)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}
	usersListApi := ConvertUserList(usersListMode)
	return ctx.JSON(200, usersListApi)
}

// あるwordのuserたちが登録しているwordたち
// (GET /list/word/{word}/words)
func (s Server) GetListWordWordWords(ctx echo.Context, word string) error {
	return nil
}

// model.UsersListからapi.UsersListへの型の変換
func ConvertUserList(models model.UsersList) api.UsersList {
	UsersListSlice := make(api.UsersList, len(models))
	for i, s := range models {
		UsersListSlice[i] = api.UserListItem{
			IncludeBot: s.IncludeBot,
			IncludeMe:  s.IncludeMe,
			Time:       s.Time,
			UserId:     s.UserId,
		}
	}
	return UsersListSlice
}

// model.UsersOfWordsListからapi.UsersOfWordsListへの型の変換
func ConvertUsersOfWordsList(models model.UsersOfWordsList) api.UsersOfWordsList {
	UsersOfWordsListSlice := make(api.UsersOfWordsList, len(models))
	for i, s := range models {
		UsersOfWordsListSlice[i] = api.UsersOfWordListItem{
			UserIds: ConvertUserList(s.UserIds),
			Word:    s.Word,
		}
	}
	return UsersOfWordsListSlice
}
