package leetcode

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"prepare-code/src/constants"
	"prepare-code/src/types"
	"strings"
)

/*
example URL:

	https://leetcode.com/problems/longest-increasing-subsequence/description

Using below API to get problem details:

	https://alfa-leetcode-api.onrender.com/
*/

type LeetcodePlatform struct {
	name     types.PlatformName
	codeName string
}

const urlFormat string = "https://leetcode.com/problems/%s/description"
const UrlCodeDelimiter string = "-"
const CodeNameDelimiter string = "."
const ApiUrl string = "https://alfa-leetcode-api.onrender.com/select?titleSlug=%s"

func (lc LeetcodePlatform) GetPlatformName() types.PlatformName {
	return constants.PlatformLeetCode
}

func (lc LeetcodePlatform) GetCodeName() string {
	return lc.codeName
}

func (lc LeetcodePlatform) GetPlatform(url string) (types.Platform, error) {
	return GetPlatform(url)
}

// This is the main method
func GetPlatform(url string) (types.Platform, error) {
	// extract code name from url
	// example: https://leetcode.com/problems/longest-increasing-subsequence/description

	// code name is longest-increasing-subsequence
	extractedCodeName, err := extractCodeName(url)
	if err != nil {
		return nil, err
	}

	return LeetcodePlatform{
		name:     constants.PlatformLeetCode,
		codeName: extractedCodeName, // placeholder for actual extraction logic
	}, nil
}

func extractCodeName(url string) (string, error) {
	// first %s comes at index 30
	if len(url) < 31 {
		return "", nil
	}

	prefixRemovedString := url[30:]
	slashIndex := 0
	for i, ch := range prefixRemovedString {
		if ch == '/' {
			slashIndex = i
			break
		}
	}

	rawCodeName := prefixRemovedString[:slashIndex]
	// within this codeName we need to replace all "-" with capitalized letters

	capitalizedCodeName := getCapitalizedString(strings.Split(rawCodeName, UrlCodeDelimiter))
	problemNumber := getProblemNumberFromUrl(rawCodeName)
	if problemNumber != CodeNameDelimiter {
		capitalizedCodeName = problemNumber + capitalizedCodeName
	}

	return capitalizedCodeName, nil
}

func getCapitalizedString(parts []string) string {
	if len(parts) == 0 {
		return ""
	}

	result := ""
	for i := 0; i < len(parts); i++ {
		part := parts[i]
		if len(part) > 0 {
			capitalizedPart := strings.ToUpper(string(part[0])) + part[1:]
			result += capitalizedPart
		}
	}

	return result
}

func getProblemNumberFromUrl(rawCodeName string) string {
	url := fmt.Sprintf(ApiUrl, rawCodeName)
	fmt.Println("URL :: ", url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching problem number:", err)
		return CodeNameDelimiter
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return CodeNameDelimiter
	}

	binBody, _ := io.ReadAll(response.Body)
	var rawMap map[string]interface{}
	json.Unmarshal([]byte(string(binBody)), &rawMap)

	problemNumber, ok := rawMap["questionId"]
	if !ok {
		fmt.Println("Problem number not found in API response")
		return CodeNameDelimiter
	}

	return fmt.Sprintf("%v%s", problemNumber, CodeNameDelimiter)
}
