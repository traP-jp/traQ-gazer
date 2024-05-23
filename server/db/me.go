package db

func ChengeMeNotification(word string, includeMe bool, userId string) error {
	_, err := db.Exec("UPDATE words SET me_notification = ? WHERE word = ? AND trap_id = ?", includeMe, word, userId)
	return err
}

func ChangeAllMeNotification(includeMe bool, userId string) error {
	_, err := db.Exec("UPDATE words SET me_notification = ? WHERE trap_id = ?", includeMe, userId)
	return err
}
