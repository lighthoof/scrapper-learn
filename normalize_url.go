package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	trimmedURL := strings.TrimSuffix(parsedURL.Hostname()+parsedURL.EscapedPath(), "/")
	resultURL := strings.ToLower(trimmedURL)

	return resultURL, nil
}
