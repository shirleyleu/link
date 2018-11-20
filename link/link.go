// Package link parses <a href> tags for URL and their associated tags
package link

import (
	"golang.org/x/net/html"
	"strings"
)

type Link struct {
	href string
	text string
}

func ExtractLink(n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		var l Link
		var t []string
		for _, a := range n.Attr {
			if a.Key == "href" {
				l.href = a.Val
				extractText(n, &t)
				l.text = strings.Join(t, " ")
				*links = append(*links, l)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ExtractLink(c, links)
	}
}

func extractText(n *html.Node, t *[]string) {
	if n.Type == html.TextNode {
		d := strings.TrimSpace(n.Data)
		if d != "" {
			*t = append(*t, d)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractText(c, t)
	}
}
