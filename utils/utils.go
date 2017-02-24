package utils

import (
	"errors"
	"strings"
	"time"
)

// ExtractDates (datesStr)
// Takes a string of dates and returns an array
func ExtractDates(datesStr string) (dates []string, err error) {

	// Check that the input string is not empty
	if len(datesStr) == 0 {
		return nil, errors.New("utils: empty dates string provided")
	}

	// Split the string by the delimiter
	dates = strings.Split(datesStr, ",")

	// Check each date is actually a date
	for _, date := range dates {
		_, err := time.Parse("2006-01-02", date)
		if err != nil {
			return nil, errors.New("utils: badly formatted dates or delimiter")
		}
	}

	return dates, nil
}

// ExtractFields (fieldStr)
// Takes a string of fields and returns an array
func ExtractFields(fieldStr string) (fileds []string, err error) {

	// List of accpeted field names
	fieldList := []string{"base", "currency", "date", "timestamp", "rate"}

	// Check that the input string is not empty
	if len(fieldStr) == 0 {
		return nil, errors.New("utils: empty fields string provided")
	}

	// Split the string by the delimiter
	fields := strings.Split(fieldStr, ",")

	for _, field := range fields {
		if !stringInSlice(field, fieldList) {
			return nil, errors.New("utils: unrecognized field supplied")
		}
	}

	return fields, nil

}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}
