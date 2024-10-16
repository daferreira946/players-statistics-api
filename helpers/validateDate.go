package helpers

import (
	"log"
	"regexp"
	"time"
)

func ValidateDate(date string) bool {
	pattern := regexp.MustCompile("^[0-9]{4}-[0-9]{2}-[0-9]{2}$")

	validString := pattern.MatchString(date)

	log.Println(date)

	if !validString {
		return false
	}

	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false
	}

	return true
}
