package helpers

import (
	"strings"
	"time"
)

// Includes : Check if string in array
func Includes(target string, array []string) bool {
	for _, str := range array {
		if strings.ToLower(str) == strings.ToLower(target) {
			return true
		}
	}
	return false
}

// SQLDateFormat : convert inputs to SQL format
func SQLDateFormat(date string) string {
	convertedDate, _ := time.Parse("02-01-2006", date) // go date format 02-Day, 01-Month
	return convertedDate.Format("2006-01-02")
}
