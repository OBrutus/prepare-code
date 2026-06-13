package platform

import (
	"errors"
	"prepare-code/src/platform/atcoder"
	"prepare-code/src/platform/codeforces"
	"prepare-code/src/platform/leetcode"
	"prepare-code/src/types"
	"strings"
)

/*
Take the url and understand what platform it is
*/
func GetPlatform(url string) (types.Platform, error) {
	if len(url) == 0 {
		return nil, errors.New("empty url")
	}

	if strings.Contains(url, "codeforces.com") {
		return codeforces.GetPlatform(url)
	} else if strings.Contains(url, "leetcode.com") {
		return leetcode.GetPlatform(url)
	} else if strings.Contains(url, "atcoder.jp") {
		return atcoder.GetPlatform(url)
	} else {
		return nil, errors.New("unknown platform or platform not supported")
	}
}
