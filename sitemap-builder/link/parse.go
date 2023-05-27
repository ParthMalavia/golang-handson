package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	var links []Link
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	dfs(doc, "", &links)
	return links, nil
}

func getText(n *html.Node) string {
	var s string
	if n.Type == html.TextNode {
		s += strings.TrimSpace(n.Data) + " "
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		s += getText(c)
	}
	return s
}

func dfs(n *html.Node, padding string, links *[]Link) {

	// Check if element is linkNode
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, att := range n.Attr {

			// For
			if att.Key == "href" {
				*links = append(*links, Link{Href: att.Val, Text: getText(n)})
			}
		}
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ", links)
	}
}
