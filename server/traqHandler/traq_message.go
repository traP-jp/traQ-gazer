package traqHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/exp/slog"
)

type Hit struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	ChannelID string    `json:"channelId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Pinned    bool      `json:"pinned"`
	Stamps    []Stamp   `json:"stamps"`
	ThreadID  string    `json:"threadId"`
}

type Stamp struct {
	StampID   string    `json:"stampId"`
	UserID    string    `json:"userId"`
	Count     int       `json:"count"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Result struct {
	TotalHits int   `json:"totalHits"`
	Hits      []Hit `json:"hits"`
}

var MessageRequestToken string

func MessageAPI() {
	slog.Info("Message API")
	now := time.Now().UTC()
	oneMinuteAgo := now.Add(-1 * time.Minute)

	url := fmt.Sprintf("https://q.trap.jp/api/v3/messages?after=%s&before=%s",
		oneMinuteAgo.Format("2006-01-02T15:04:05Z"),
		now.Format("2006-01-02T15:04:05Z"))
	slog.Info(url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		slog.Info("Error creating request:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+MessageRequestToken)

	client := http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		slog.Info("Error:%s", err)
	}

	jsonBytes := ([]byte)(byteArray)
	result := new(Result)

	if err := json.Unmarshal(jsonBytes, result); err != nil {
		slog.Info("JSON Unmarshal error:", err)
	}

	slog.Info("Total Hits:", result.TotalHits)
	slog.Info("Hits:")
	for _, hit := range result.Hits {
		slog.Info("ID:", hit.ID)
		slog.Info("UserID:", hit.UserID)
		slog.Info("ChannelID:", hit.ChannelID)
		slog.Info("Content:", hit.Content)
		slog.Info("CreatedAt:", hit.CreatedAt)
		slog.Info("UpdatedAt:", hit.UpdatedAt)
		slog.Info("Pinned:", hit.Pinned)
		slog.Info("Stamps:", hit.Stamps)
		slog.Info("ThreadID:", hit.ThreadID)
		slog.Info("-----------------")
	}
}
