package model

type MessageItem struct {
	// メッセージUUID
	Id string `json:"id"`
	// 投稿者UUID
	TraqUuid string `json:"userId"`
	// メッセージ本文
	Content string `json:"content"`
}

type MessageList []MessageItem

type WordsItem struct {
	IncludeBot bool   `db:"bot_notification"`
	IncludeMe  bool   `db:"me_notification"`
	TrapId     string `db:"trap_id"`
	Word       string `db:"word"`
}

type UsersItem struct {
	TrapID   string `db:"trap_id"`
	TraqUUID string `db:"traq_uuid"`
	IsBot    bool   `db:"is_bot"`
}

type Send struct {
	// 含んでいた単語
	Word string
	// 送信先のuser
	NotifyTargetTrapId string
	// 送信先のuserUUID
	NotifyTargetTraqUuid string
	// 送信するメッセージのID
	MessageId string
	// BOTかどうか
	IsBot bool
}

type SendList []*Send

type NotifyInfo struct {
	Words []string
	// 送信先のuser
	NotifyTargetTrapId string
	// 送信先のuserUUID
	NotifyTargetTraqUuid string
	// 送信するメッセージのID
	MessageId string
}

type MatchedWords struct {
	ContactedWords string `db:"contacted_words"`
	TrapID         string `db:"trap_id"`
	TraqUUID       string `db:"traq_uuid"`
}
