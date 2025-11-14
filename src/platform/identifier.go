package platform

import (
	"errors"
	"prepare-code/src/platform/codeforces"
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
	} else {
		return nil, errors.New("unknown platform or platform not supported")
	}
}
