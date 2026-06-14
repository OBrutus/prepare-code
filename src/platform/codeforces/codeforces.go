package codeforces

import (
	"errors"
	"prepare-code/src/constants"
	"prepare-code/src/types"
	"strings"
)

const urlFormat string = "https://codeforces.com/problemset/problem/%s/%s"
const codeNameDelimiter string = "-"

const (
	ProblemSet = "problemset"
	Contest    = "contest"
)

var ErrInvaliidCodeforcesUrl = errors.New("invalid url")

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
	// For codeforces either the problem could be from
	// problem set or contest like below

	url = strings.ToLower(url)

	// extract code name from url
	// example1 [problemset]: https://codeforces.com/problemset/problem/1234/A
	// example2 [contest]: https://codeforces.com/contest/2236/problem/D

	// if delimiter is "-"
	// code name is 1234-A [example1]
	// or name is 2236-D [example2]

	var extractedCodeName string = ""
	var err error
	if strings.Contains(url, ProblemSet) {
		extractedCodeName, err = extractCodeNameProblemSet(url)
	} else {
		extractedCodeName, err = extractCodeNameContest(url)
	}

	if err != nil {
		return nil, err
	}

	return CodeforcesInstance{
		name:     constants.PlatformCodeforces,
		codeName: extractedCodeName, // placeholder for actual extraction logic
	}, nil
}

func extractCodeNameContest(url string) (string, error) {
	// https://codeforces.com/contest/2236/problem/D
	// first %s comes at index 42
	if len("codeforces.com/contest/?/problem/?") > len(url) {
		return "", ErrInvaliidCodeforcesUrl
	}

	index := strings.Index(url, Contest)
	if index < 0 {
		return "", ErrInvaliidCodeforcesUrl
	}

	index += len(Contest) + len("/")

	// now the index is at contest Id
	substring := url[index:]

	split := strings.Split(substring, "/")

	contestCode := string(split[0])
	problem := strings.ToUpper(string(split[2]))

	return contestCode + codeNameDelimiter + problem, nil
}

func extractCodeNameProblemSet(url string) (string, error) {
	// first %s comes at index 42
	if len(url) < 43 {
		return "", ErrInvaliidCodeforcesUrl
	}

	prefixRemovedString := url[42:]
	firstSlashIndex := strings.Index(prefixRemovedString, "/")

	contestNumber := prefixRemovedString[:firstSlashIndex]
	contestRemovedString := prefixRemovedString[firstSlashIndex+1:]

	// here the assumption is single character [from A - Z]
	questionCode := contestRemovedString[0]

	return contestNumber + codeNameDelimiter + string(questionCode), nil
}
