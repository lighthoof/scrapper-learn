package main

import (
	"net/url"
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

func getFirstParagraphFromHTML(html string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return "", err
	}

	tagContent := doc.Find("main")
	if len(tagContent.Nodes) > 0 {
		tagContent = tagContent.Find("p")
	} else {
		tagContent = doc.Find("p")
	}

	return tagContent.First().Text(), nil
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	allURLs := []string{}
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists || href == "" {
			return
		}

		parsedRef, err := url.Parse(href)
		if err != nil {
			return
		}

		if parsedRef.Fragment != "" && parsedRef.Path == "" {
			return
		}

		absoluteURL := baseURL.ResolveReference(parsedRef)
		allURLs = append(allURLs, absoluteURL.String())
	})
	return allURLs, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return []string{}, err
	}

	allImages := []string{}
	doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if !exists || src == "" {
			return
		}

		parsedSrc, err := url.Parse(src)
		if err != nil {
			return
		}

		if parsedSrc.Fragment != "" && parsedSrc.Path == "" {
			return
		}

		absoluteImageURL := baseURL.ResolveReference(parsedSrc)
		allImages = append(allImages, absoluteImageURL.String())
	})
	return allImages, nil
}
