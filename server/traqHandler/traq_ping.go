package traqHandler

import (
	"github.com/traPtitech/go-traq"
	"golang.org/x/exp/slog"
)

// traQからのイベントを受け取るためのHTTPサーバー
type TraqServer struct {
	client *traq.APIClient
}

func (t TraqServer) SetPingHandler() error {
	slog.Info("traQ ping received")
	return nil
}
