package codeforces

import (
	"errors"
	"prepare-code/src/constants"
	"prepare-code/src/types"
	"strings"
)

const urlFormat string = "https://codeforces.com/problemset/problem/%s/%s"
const codeNameDelimiter string = "-"

type CodeforcesInstance struct {
	name     types.PlatformName
	codeName string
}

func (cf CodeforcesInstance) GetPlatformName() types.PlatformName {
	return constants.PlatformCodeforces
}

func (cf CodeforcesInstance) GetCodeName() string {
	return cf.codeName
}

// This is the main method
func GetInstance(url string) (types.Instance, error) {
	// extract code name from url
	// example: https://codeforces.com/problemset/problem/1234/A

	// if delimiter is "-"
	// code name is 1234-A
	extractedCodeName, err := extractCodeName(url)
	if err != nil {
		return nil, err
	}

	return CodeforcesInstance{
		name:     constants.PlatformCodeforces,
		codeName: extractedCodeName, // placeholder for actual extraction logic
	}, nil
}

func extractCodeName(url string) (string, error) {
	// first %s comes at index 42
	if len(url) < 43 {
		return "", errors.New("invalid url")
	}

	prefixRemovedString := url[42:]
	firstSlashIndex := strings.Index(prefixRemovedString, "/")

	contestNumber := prefixRemovedString[:firstSlashIndex]
	contestRemovedString := prefixRemovedString[firstSlashIndex+1:]

	// here the assumption is single character [from A - Z]
	questionCode := contestRemovedString[0]

	return contestNumber + codeNameDelimiter + string(questionCode), nil
}
