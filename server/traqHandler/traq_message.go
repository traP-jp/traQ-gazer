package traqHandler

import (
	"context"

	"github.com/traPtitech/go-traq"
	"golang.org/x/exp/slog"
)

func (t TraqServer) SetMessageCreatedHandler() error {
	// TODO: userIdを取得する
	v, _, err := t.client.MessageApi.PostDirectMessage(context.Background(), "userId").
		PostMessageRequest(traq.PostMessageRequest{
			// メッセージ本文
			Content: "",
			// // メンション・チャンネルリンクを自動埋め込みするか
			// Embed: check,
		}).
		Execute()
	slog.Info("%#v", v)
	if err != nil {
		slog.Info("%s", err)
	}
	return nil
}
