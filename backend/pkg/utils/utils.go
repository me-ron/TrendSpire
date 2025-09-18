package utils

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func Ptr[T any](v T) *T {
	return &v
}

func ParseDuration(s string) (time.Duration, error) {
	if strings.HasSuffix(s, "d") {
		days, err := strconv.Atoi(strings.TrimSuffix(s, "d"))
		if err != nil {
			log.Fatalf("❌invalid duration: %v", err)
			return 0, err
		}
		return time.Duration(days) * 24 * time.Hour, nil
	}

	dur, err := time.ParseDuration(s)
	if err != nil {
		log.Fatalf("❌invalid duration format: %v", err)
		return 0, err
	}
	return dur, nil
}
