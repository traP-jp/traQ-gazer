package db

import (
	"time"
	"traQ-gazer/model"

	"golang.org/x/exp/slog"
)

// 今日のトレンド
func GetTrendToday(limit int) (model.TrendingWords, error) {
	date := time.Now()
	return GetTrendOneday(FormatDate(date), limit)
}

// ある日YYYY-MM-DDのトレンド
func GetTrendOneday(day string, limit int) (model.TrendingWords, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		slog.Info("Error GetTrendOneday LoadLocation: %v", err)
		return model.TrendingWords{}, err
	}

	t, err := ParseDay(day)
	if err != nil {
		slog.Info("Error GetTrendOneday ParseDay: %v", err)
		return model.TrendingWords{}, err
	}

	dateFrom := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
	dateTo := dateFrom.AddDate(0, 0, 1)
	return GetTrendRange(FormatDate(dateFrom), FormatDate(dateTo), limit)
}

// ある月YYYY-MMのトレンド
func GetTrendOneMonth(month string, limit int) (model.TrendingWords, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		slog.Info("Error GetTrendOneMonth LoadLocation: %v", err)
		return model.TrendingWords{}, err
	}

	t, err := ParseMonth(month)
	if err != nil {
		slog.Info("Error GetTrendOneMonth ParseMonth: %v", err)
		return model.TrendingWords{}, err
	}

	dateFrom := time.Date(t.Year(), t.Month(), 0, 0, 0, 0, 0, loc)
	dateTo := dateFrom.AddDate(0, 1, 0)
	return GetTrendRange(FormatDate(dateFrom), FormatDate(dateTo), limit)
}

// ある年YYYYのトレンド
func GetTrendOneYear(year string, limit int) (model.TrendingWords, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		slog.Info("Error GetTrendOneYear LoadLocation: %v", err)
		return model.TrendingWords{}, err
	}

	t, err := ParseYear(year)
	if err != nil {
		slog.Info("Error GetTrendOneYear ParseYear: %v", err)
		return model.TrendingWords{}, err
	}

	dateFrom := time.Date(t.Year(), 0, 0, 0, 0, 0, 0, loc)
	dateTo := dateFrom.AddDate(1, 0, 0)
	return GetTrendRange(FormatDate(dateFrom), FormatDate(dateTo), limit)
}

// dateFrom <= resister_time < dateTo　となるデータを集計する
func GetTrendRange(dateFrom, dateTo string, limit int) (model.TrendingWords, error) {
	var words []model.TrendingWord
	err := db.Select(&words, "SELECT COUNT(*) AS number, word FROM words WHERE (register_time >= ? AND register_time < ?) GROUP BY word ORDER BY number DESC LIMIT ?", dateFrom, dateTo, limit)
	if err != nil {
		slog.Info("Error GetTrendRange Select: %v", err)
		return model.TrendingWords{}, err
	}
	return words, nil
}

func FormatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func ParseDay(day string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01-02", day)
	if err != nil {
		slog.Info("Error ParseDay: %v", err)
		return time.Time{}, err
	}

	return parsedDate, nil
}
func ParseMonth(month string) (time.Time, error) {
	parsedDate, err := time.Parse("2006-01", month)
	if err != nil {
		slog.Info("Error ParseMonth: %v", err)
		return time.Time{}, err
	}

	return parsedDate, nil
}
func ParseYear(year string) (time.Time, error) {
	parsedDate, err := time.Parse("2006", year)
	if err != nil {
		slog.Info("Error ParseYear: %v", err)
		return time.Time{}, err
	}

	return parsedDate, nil
}
