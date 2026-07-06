package repo

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os"
	"slices"
	"time"
	"traQ-gazer/model"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/traPtitech/go-traq"
)

var (
	db          *sqlx.DB
	AccessToken = os.Getenv("BOT_ACCESS_TOKEN")
	dbUsername  = os.Getenv("DB_USERNAME")
	dbPort      = os.Getenv("DB_PORT")
	dbHostname  = os.Getenv("DB_HOSTNAME")
	dbPassword  = os.Getenv("DB_PASSWORD")
	dbDatabase  = os.Getenv("DB_DATABASE")
)

func SetUp() error {
	conf := mysql.Config{
		User:                 dbUsername,
		Passwd:               dbPassword,
		Net:                  "tcp",
		Addr:                 net.JoinHostPort(dbHostname, dbPort),
		DBName:               dbDatabase,
		Loc:                  time.Local,
		AllowNativePasswords: true,
		ParseTime:            true,
		Collation:            "utf8mb4_general_ci",
	}
	_db, err := sqlx.Connect("mysql", conf.FormatDSN())
	if err != nil {
		return err
	}
	db = _db

	slog.Info("connected to database")
	err = initUsersTable()
	if err != nil {
		return err
	}
	return nil
}

// ユーザーとそのuuidの対照表を作る
func initUsersTable() error {
	if AccessToken == "" {
		slog.Info("skip init users table")
		return nil
	}

	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, AccessToken)

	result, _, err := client.UserApi.GetUsers(auth).Execute()
	if err != nil {
		slog.Error("failed to get users", "err", err)
		return err
	}

	userList := model.UserList{}
	for _, user := range result {
		userList = append(userList, model.User{Traq_uuid: user.Id, Trap_id: user.Name, Is_bot: user.Bot})
	}

	alreadyExistUsersUUIDList := []string{}
	err = db.Select(&alreadyExistUsersUUIDList, "SELECT traq_uuid FROM users")
	if err != nil {
		slog.Error("failed to select existing users", "err", err)
		return err
	}

	newUserList := removeAlreadyExistUsers(userList, alreadyExistUsersUUIDList)

	for i := 0; i < len(newUserList); i += 50 {
		batch := newUserList[i:min(i+50, len(newUserList))]
		_, err := db.NamedExec("INSERT INTO users (traq_uuid, trap_id, is_bot) VALUES (:traq_uuid, :trap_id, :is_bot)", batch)
		if err != nil {
			slog.Error("failed to insert users", "count", len(batch), "err_type", fmt.Sprintf("%T", err))
			return err
		}
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func removeAlreadyExistUsers(allUsers model.UserList, alreadyUsersUUID []string) model.UserList {
	newUserList := make(model.UserList, 0)
	for _, all := range allUsers {
		if !slices.Contains(alreadyUsersUUID, all.Traq_uuid) {
			newUserList = append(newUserList, model.User{
				Traq_uuid: all.Traq_uuid,
				Trap_id:   all.Trap_id,
				Is_bot:    all.Is_bot,
			})
		}
	}
	return newUserList
}

func RecordPollingTime(lastCheckPoint time.Time) error {
	_, err := db.Exec("INSERT INTO `pollinginfo`(`key`,`lastpollingtime`) VALUES(1,?) ON DUPLICATE KEY UPDATE `lastpollingtime`=VALUES(lastpollingtime)", lastCheckPoint)
	if err != nil {
		return err
	}

	return nil
}

func GetPollingFrom() (time.Time, error) {
	var from time.Time
	err := db.Get(&from, "SELECT `lastpollingtime` FROM `pollinginfo` WHERE `key`=1")
	if err != nil {
		return from, err
	}

	return from, nil
}
