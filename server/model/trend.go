package model

import "time"

func GetTrendOneday(date time.Time) ([]string, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return []string{}, err
	}
	dateFrom := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, loc)
	dateTo := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 9999, loc)
	return GetTrendRange(dateFrom, dateTo)
}

func GetTrendOneday(date time.Time) ([]string, error) {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return []string{}, err
	}
	dateFrom := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, loc)
	dateTo := time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, 9999, loc)
	return GetTrendRange(dateFrom, dateTo)
}

func GetTrendRange(dateFrom, dateTo time.Time, number int) ([]string, error) {
	
	return []string{}, err
} 
