package oapi

import (
	"net/http"
	"traQ-gazer/model"
	"traQ-gazer/repo"

	"github.com/labstack/echo/v4"
)

// 今見てるuserのwordたち
// (GET /list/user/{userId})
func (s Server) GetListUserMe(ctx echo.Context) error {
	// traPIdの取得
	userId, err := getUserIdFromSession(ctx)
	if err != nil {
		// 正常でないためステータスコード 400 "Invalid Input"
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	wordsListMode, err := repo.GetListUserUserId(userId)
	if err != nil {
		return err
	}
	wordsListApi := ConvertWordList(wordsListMode)
	return ctx.JSON(200, wordsListApi)
}

// あるuserのwordたち
// (GET /list/user/{userId})
func (s Server) GetListUserUserId(ctx echo.Context, userId string) error {
	wordsListMode, err := repo.GetListUserUserId(userId)
	if err != nil {
		return err
	}
	wordsListApi := ConvertWordList(wordsListMode)
	return ctx.JSON(200, wordsListApi)
}

// あるuserのwordたちを登録しているuserたち
// (GET /list/user/{userId}/users)
func (s Server) GetListUserUserIdUsers(ctx echo.Context, userId string) error {
	// usersOfWordsListMode, err := model.GetListUserUserIdUsers(userId)
	// if err != nil {
	// 	return echo.NewHTTPError(500, err.Error())
	// }
	// usersOfWordsListApi := ConvertUsersOfWordsList(usersOfWordsListMode)
	// return ctx.JSON(200, usersOfWordsListApi)
	return nil
}

// model.WordsListからoapi.WordsListへの型の変換
func ConvertWordList(models model.WordsList) WordsList {
	WordsListSlice := make(WordsList, len(models))
	for i, s := range models {
		WordsListSlice[i] = WordListItem{
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
	usersListMode, err := repo.GetListWordWord(word)
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

// model.UsersListからoapi.UsersListへの型の変換
func ConvertUserList(models model.UsersList) UsersList {
	UsersListSlice := make(UsersList, len(models))
	for i, s := range models {
		UsersListSlice[i] = UserListItem{
			IncludeBot: s.IncludeBot,
			IncludeMe:  s.IncludeMe,
			Time:       s.Time,
			UserId:     s.UserId,
		}
	}
	return UsersListSlice
}

// model.UsersOfWordsListからoapi.UsersOfWordsListへの型の変換
func ConvertUsersOfWordsList(models model.UsersOfWordsList) UsersOfWordsList {
	UsersOfWordsListSlice := make(UsersOfWordsList, len(models))
	for i, s := range models {
		UsersOfWordsListSlice[i] = UsersOfWordListItem{
			UserIds: ConvertUserList(s.UserIds),
			Word:    s.Word,
		}
	}
	return UsersOfWordsListSlice
}
