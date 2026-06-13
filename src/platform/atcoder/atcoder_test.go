package atcoder

import "testing"

func TestSimpleCodeExtraction(t *testing.T) {
	testCases := []struct {
		name string
		url  string
		code string
		err  error
	}{
		{name: "Positive Test Case", url: "https://atcoder.jp/contests/abc461/tasks/abc461_a", code: "ABC461_A", err: nil},
		{name: "Test Case 1", url: "https://atcoder.jp/contests/abc461/tasks/", err: ErrInvalidAtCoderUrl},
		{name: "Test Case 2", url: "https://atcoder.jp/contests/abc461/tasks", err: ErrInvalidAtCoderUrl},
		{name: "Test Case 3", url: "https://atcoder.jp/contests/abc461/", err: ErrInvalidAtCoderUrl},
		{name: "Test Case 4", url: "https://atcoder.jp/contests/abc461", err: ErrInvalidAtCoderUrl},
		{name: "Test Case 5", url: "https://atcoder.jp/contests/", err: ErrInvalidAtCoderUrl},
		{name: "Test Case 6", url: "https://atcoder/contests/", err: ErrInvalidAtCoderUrl},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			code, err := extractCodeName(tt.url)
			if (err != nil && tt.err != nil) && err != tt.err {
				t.Errorf("Test Failed for %s. Got=%v Expected=%v", tt.name, err, tt.err)
				return
			}

			if code != tt.code {
				t.Errorf("Test Failed for %s. Got=%v Expected=%v", tt.name, code, tt.code)
				return
			}
		})
	}
}
