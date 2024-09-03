package main

import (
	"fmt"
	"net/url"
	"strings"
	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {

	// call url.Parse to obtain baseURL
	baseURL, err := url.Parse(rawBaseURL)

	// handle error if present
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	// create a reader interface
	htmlReader := strings.NewReader(htmlBody)

	// parse the reader interface
	doc, err := html.Parse(htmlReader)

	// handle parsing error if present
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %v", err)
	}

	//create empty slice of urls
	var urls []string

	// function declaration for traverseNodes
	var traverseNodes func(*html.Node)

	// function definition assigned to traverse nodes variable
	traverseNodes = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "a" {
			for _, anchor := range node.Attr {
				if anchor.Key == "href" {
					href, err := url.Parse(anchor.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v\n", anchor.Val, err)
						continue
					}

					resolvedURL := baseURL.ResolveReference(href)
					urls = append(urls, resolvedURL.String())
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			traverseNodes(child)
		}
	}
	traverseNodes(doc)

	return urls, nil
}
