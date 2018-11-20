package main

import (
	"flag"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

func main(){
	filename := flag.String("file", "ex1.html", "name of html file to parse")
	flag.Parse()

	d, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(d)
	if err != nil {
		log.Fatal(err)
	}

	var links []link
	extractLink(doc, &links)
}

func extractLink(n *html.Node, links *[]link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		var l link
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
		extractLink(c, links)
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

type link struct {
	href string
	text string
}
