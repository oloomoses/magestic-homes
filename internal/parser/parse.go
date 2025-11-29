package parser

import (
	"strings"

	"githum.com/oloomoses/magestic-homes/internal/model"
	"golang.org/x/net/html"
)

func ExtractItem(htmlBody string) ([]model.House, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))

	if err != nil {
		return []model.House{}, err
	}

	var houses []model.House

	var walk func(*html.Node)

	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "li" {
			houses = append(houses, parseListHouses(n))
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(doc)

	return houses, nil
}

func parseListHouses(li *html.Node) model.House {
	link := findHref(li)
	title := findTitle(li, "title")
	price := findTitle(li, "price")
	location := findTitle(li, "location")

	return model.House{
		Link:     link,
		Title:    title,
		Price:    price,
		Location: location,
	}
}

func findHref(li *html.Node) string {
	if li.Type == html.ElementNode && li.Data == "a" {
		for _, a := range li.Attr {
			if a.Key == "href" {
				return a.Val
			}
		}
	}

	for c := li.FirstChild; c != nil; c = c.NextSibling {
		if href := findHref(c); href != "" {
			return href
		}

	}

	return ""
}

func findTitle(n *html.Node, cls string) string {
	var result string

	var walk func(*html.Node)

	walk = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "div" {
			for _, a := range node.Attr {
				if a.Key == "class" && a.Val == cls {
					result = getTextContent(node)
					return
				}
			}

		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			walk(c)
		}
	}

	walk(n)

	return result
}

func getTextContent(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	var sb strings.Builder

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		sb.WriteString(getTextContent(c))
	}
	return sb.String()
}
