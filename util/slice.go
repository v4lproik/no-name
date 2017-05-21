package util

import (
	"regexp"
)

func Contains(stringSlice []string, searchString string) bool {
	for _, value := range stringSlice {
		if value == searchString {
			return true
		}
	}
	return false
}

func ContainsRegex(regexs []string, searchString string) bool {
	for _, regex := range regexs {
		match, err := regexp.MatchString(regex, searchString)
		if match {
			return true
		}
		if err != nil {
			loggerUtilHTML.Errorf(err.Error())
		}
	}
	return false
}

func ContainsGetIndex(stringSlice []string, searchString string) int {
	for key, value := range stringSlice {
		if value == searchString {
			return key
		}
	}
	return -1
}
