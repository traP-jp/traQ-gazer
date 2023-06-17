package traqHandler

import "github.com/traPtitech/go-traq"

// traQからのイベントを受け取るためのHTTPサーバー
type TraqServer struct {
	client *traq.APIClient
}

func (t TraqServer) SetPingHandler() error {
	return nil
}
