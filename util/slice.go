package util

import (
	"regexp"
)

func Contains(stringSlice []string, searchString string) bool {
	for _, regex := range stringSlice {
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
