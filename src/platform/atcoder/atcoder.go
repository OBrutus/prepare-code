package atcoder

import (
	"errors"
	"prepare-code/src/constants"
	"prepare-code/src/types"
	"strings"
)

var ErrInvalidAtCoderUrl = errors.New("invalid url of AtCoder")

const urlFormat string = "https://atcoder.jp/contests/abc461/tasks/abc461_a"

const codeNameDelimiter string = "-"

type AtCoderInstance struct {
	name     types.PlatformName
	codeName string
}

func (ac AtCoderInstance) GetPlatformName() types.PlatformName {
	return constants.PlatformAtCoder
}

func (ac AtCoderInstance) GetCodeName() string {
	return ac.codeName
}

// This is the main method
func GetInstance(url string) (types.Instance, error) {
	extractedCodeName, err := extractCodeName(url)
	if err != nil {
		return nil, err
	}

	return AtCoderInstance{
		name:     constants.PlatformCodeforces,
		codeName: extractedCodeName, // placeholder for actual extraction logic
	}, nil
}

func extractCodeName(url string) (string, error) {
	// https://atcoder.jp/contests/abc461/tasks/abc461_a
	if len(url) < len("atcoder.jp") {
		return "", ErrInvalidAtCoderUrl
	}

	i := 0
	const prefix string = "atcoder.jp/contests/"
	for i+len(prefix) < len(url) {
		if url[i:i+len(prefix)] == prefix {
			break
		}
		i++
	}

	i += len(prefix)
	if i == len(url) {
		return "", ErrInvalidAtCoderUrl
	}

	prefixRemovedString := url[i:]
	suffix := strings.Split(prefixRemovedString, "/")
	if len(suffix) < 3 {
		return "", ErrInvalidAtCoderUrl
	}

	// contestNumber := suffix[0]
	qeustionCode := strings.ToUpper(suffix[2])

	return qeustionCode, nil
}
