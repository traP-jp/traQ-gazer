package repo

func ChengeBotNotification(word string, includeBot bool, userId string) error {
	_, err := db.Exec(
		"UPDATE words SET bot_notification = ? WHERE word = ? AND trap_id = ?",
		includeBot,
		word,
		userId,
	)
	return err
}

func ChangeAllBotNotification(includeBot bool, userId string) error {
	_, err := db.Exec(
		"UPDATE words SET bot_notification = ? WHERE trap_id = ?",
		includeBot,
		userId,
	)
	return err
}
