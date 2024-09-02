package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {

  // call url.Parse and obtain parsedURL struct
	parsedURL, err := url.Parse(rawURL)

  // handle error if there is one
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
  
  // strip scheme
	fullPath := parsedURL.Host + parsedURL.Path

  // lowercase all characters
	fullPath = strings.ToLower(fullPath)

  // remove trailing slash
	fullPath = strings.TrimSuffix(fullPath, "/")

  // return normalized URL
	return fullPath, nil
}
