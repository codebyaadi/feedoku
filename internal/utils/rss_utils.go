package utils

import (
	"fmt"
	"time"
)

func ParseDate(dateStr string) (time.Time, error) {
	// List of potential date formats
	formats := []string{
		time.RFC1123,                    // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC1123Z,                   // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC3339,                    // "2006-01-02T15:04:05Z07:00"
		"Mon, 02 Jan 2006 15:04:05 GMT", // Explicitly handle "GMT" instead of timezone
	}

	for _, format := range formats {
		parsedTime, err := time.Parse(format, dateStr)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("could not parse date: %s", dateStr)
}
