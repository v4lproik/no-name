package util

func Contains(stringSlice []string, searchString string) bool {
	for _, value := range stringSlice {
		if value == searchString {
			return true
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
