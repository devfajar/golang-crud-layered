package helper

import "time"

func ParseDate(dateStr string) *time.Time {
	if dateStr == "" {
		return nil
	}
	layout := "2006-01-02"
	t, err := time.Parse(layout, dateStr)
	if err != nil {
		return nil
	}
	return &t
}