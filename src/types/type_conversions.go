package types

import "strings"

/*
If the string == "true"
return true else false
*/
func GetBoolFromString(stringValue string) bool {
	stringValue = strings.TrimSpace(stringValue)
	stringValue = strings.ToLower(stringValue)

	if stringValue == "true" || stringValue == "1" {
		return true
	}

	return false
}
