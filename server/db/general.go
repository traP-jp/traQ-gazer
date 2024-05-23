package db

import (
	"context"
	"net"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/traPtitech/go-traq"
	"golang.org/x/exp/slices"
	"golang.org/x/exp/slog"
)

var (
	db           *sqlx.DB
	ACCESS_TOKEN = os.Getenv("BOT_ACCESS_TOKEN")
	db_username  = os.Getenv("DB_USERNAME")
	db_port      = os.Getenv("DB_PORT")
	db_hostname  = os.Getenv("DB_HOSTNAME")
	db_password  = os.Getenv("DB_PASSWORD")
	db_database  = os.Getenv("DB_DATABASE")
)

func SetUp() error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	conf := mysql.Config{
		User:                 db_username,
		Passwd:               db_password,
		Net:                  "tcp",
		Addr:                 net.JoinHostPort(db_hostname, db_port),
		DBName:               db_database,
		Loc:                  jst,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	db, err = sqlx.Connect("mysql", conf.FormatDSN())
	if err != nil {
		return err
	}

	slog.Info("Connected to Database")
	err = initUsersTable()
	if err != nil {
		return err
	}
	return nil
}

// ユーザーとそのuuidの対照表を作る
func initUsersTable() error {
	if ACCESS_TOKEN == "" {
		slog.Info("Skip initUsersTable")
		return nil
	}

	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, ACCESS_TOKEN)

	result, _, err := client.UserApi.GetUsers(auth).Execute()
	if err != nil {
		slog.Error("Error getting users: %v", err)
		return err
	}

	userList := UserList{}
	for _, user := range result {
		userList = append(userList, User{Traq_uuid: user.Id, Trap_id: user.Name, Is_bot: user.Bot})
	}

	alreadyExistUsersUUIDList := []string{}
	err = db.Select(&alreadyExistUsersUUIDList, "SELECT traq_uuid FROM users")
	if err != nil {
		slog.Error("Error Select alreadyExistUsersUUIDList: %v", err)
		return err
	}

	newUserList := removeAlreadyExistUsers(userList, alreadyExistUsersUUIDList)

	for i := 0; i < len(newUserList); i += 50 {
		_, err := db.NamedExec("INSERT INTO users (traq_uuid, trap_id, is_bot) VALUES (:traq_uuid, :trap_id, :is_bot)", newUserList[i:min(i+50, len(newUserList))])
		if err != nil {
			slog.Error("Error Insert: %v", err)
			return err
		}
	}
	return nil
}

type User struct {
	Traq_uuid string `db:"traq_uuid"`
	Trap_id   string `db:"trap_id"`
	Is_bot    bool   `db:"is_bot"`
}

type UserList []User

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func removeAlreadyExistUsers(allUsers UserList, alreadyUsersUUID []string) UserList {
	newUserList := make(UserList, 0)
	for _, all := range allUsers {
		if !slices.Contains(alreadyUsersUUID, all.Traq_uuid) {
			newUserList = append(newUserList, User{
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
