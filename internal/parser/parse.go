package parser

import (
	"strings"

	"golang.org/x/net/html"
)

func ExtactTitle(htmlBody string) (string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))

	if err != nil {
		return "", err
	}

	var findTitle func(*html.Node) string

	findTitle = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, attr := range n.Attr {
				if attr.Key == "class" && attr.Val == "title" {
					return getText(n)
				}
			}
			return n.FirstChild.Data
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			result := findTitle(c)
			if result != "" {
				return result
			}
		}

		return ""
	}

	return findTitle(doc), nil
}

func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	var sb strings.Builder

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(getText(c))
	}

	return strings.TrimSpace(sb.String())
}
