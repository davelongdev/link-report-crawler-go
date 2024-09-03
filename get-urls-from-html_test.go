package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	cases := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:     "absolute URL",
			inputURL: "https://blog.test.com",
			inputBody: `
<html>
	<body>
		<a href="https://blog.test.com">
			<span>Test.com</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.test.com"},
		},
		{
			name:     "relative URL",
			inputURL: "https://test.com",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Test.com</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://test.com/path/one"},
		},
		{
			name:     "absolute and relative URLs",
			inputURL: "https://test.com",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Test.com</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Test.com</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://test.com/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no href",
			inputURL: "https://test.com",
			inputBody: `
<html>
	<body>
		<a>
			<span>test.com></span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "bad HTML",
			inputURL: "https://test.com",
			inputBody: `
<html body>
	<a href="path/one">
		<span>Test.com></span>
	</a>
</html body>
`,
			expected: []string{"https://test.com/path/one"},
		},
		{
			name:     "invalid href URL",
			inputURL: "https://blog.boot.dev",
			inputBody: `
<html>
	<body>
		<a href=":\\invalidURL">
			<span>Test.com</span>
		</a>
	</body>
</html>
`,
			expected: nil,
		},
		{
			name:     "handle invalid base URL",
			inputURL: `:\\invalidBaseURL`,
			inputBody: `
<html>
	<body>
		<a href="/path">
			<span>Test.com</span>
		</a>
	</body>
</html>
`,
			expected:      nil,
			errorContains: "couldn't parse base URL",
		},
	}

	for i, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - '%s' FAIL: expected URLs %v, got URLs %v", i, tc.name, tc.expected, actual)
				return
			}
		})
	}
}

