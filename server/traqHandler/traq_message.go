package traqHandler

import (
	"h23s_15/model"

	traqbot "github.com/traPtitech/traq-bot"
	"golang.org/x/exp/slog"
)

func (t TraqServer) SetMessageCreatedHandler(p *traqbot.MessageCreatedPayload) error {
	model.CheckMessageFromWords(p.Message.Text)
	// // TODO: userIdを取得する
	// v, _, err := t.client.MessageApi.PostDirectMessage(context.Background(), "userId").
	// 	PostMessageRequest(traq.PostMessageRequest{
	// 		// メッセージ本文
	// 		Content: "",
	// 		// // メンション・チャンネルリンクを自動埋め込みするか
	// 		// Embed: check,
	// 	}).
	// 	Execute()
	// slog.Info("%#v", v)
	// if err != nil {
	// 	slog.Info("%s", err)
	// }
	slog.Info(p.Message.Text)
	return nil
}
