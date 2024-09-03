package main

import (
  "strings"
  "testing"
)

func TestNormalizeURL(t *testing.T) {

	// create a struct for each test case
	tests := []struct {

		// name of test case
		name		string

		// raw URL passed to the normalizeURL function
		inputURL	string

		// what the normalizeURL function is expected to return
		expected	string

		// if normalizeURL returns an error, what the error is expected to contain 
		errorContains	string
	}{
		{
			name:     "remove scheme http",
			inputURL: "http://test.com/path/",
			expected: "test.com/path",
		},		
		{
			name:     "remove scheme https",
			inputURL: "https://test.com/path",
			expected: "test.com/path",
		},
		{
			name:     "remove trailing slash",
			inputURL: "test.com/path/",
			expected: "test.com/path",
		},
		{
			name:     "lowercase capital letters",
			inputURL: "TEST.com/PATH",
			expected: "test.com/path",
		},
		{
			name:     "remove scheme and capitals and trailing slash",
			inputURL: "http://TEST.com/PATH/",
			expected: "test.com/path",
		},
		{
			name:          "handle invalid URL",
			inputURL:      `:\\invalidURL`,
			expected:      "",
			errorContains: "couldn't parse URL",
		},

		// can add more test cases here
		
	}

	// loop through test cases
	for i, tc := range tests {

		// create sub tests for each test case with t.Run
		t.Run(tc.name, func(t *testing.T) {

			// call the NormalizeURL function and store its return values
			actual, err := normalizeURL(tc.inputURL)

			// call t.errorf if error is returned and the error does not contain expected text
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return

			// call t.errorf if error is returned and there is no expected error text
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return

			// call t.errorf if no error is returned and error text was expected
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			// call t.Errorf if expected result is different than actual result
			if actual != tc.expected {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
