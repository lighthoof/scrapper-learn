package main

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}

	tagContent := doc.Find("h1")
	if tagContent.Text() == "" {
		tagContent = doc.Find("h2")
	}
	return tagContent.Text(), nil
}

func getFristParagraphFromHTML(html string) (string, error) {
	return "", nil
}
