package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

func main(){
	filename := flag.String("file", "ex1.html", "html file")
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
	fmt.Printf("%+v\n", links)
}

func extractLink(n *html.Node, links *[]link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		var l link
		for _, a := range n.Attr {
			if a.Key == "href" {
				l.href = a.Val
				extractText(n, &l)
				*links = append(*links, l)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractLink(c, links)
	}
}

func extractText(n *html.Node, l *link) {
	if n.Type == html.TextNode {
		l.text += strings.TrimSpace(n.Data) // TODO: append to list (needs to be passed in instead of l *link) and then concatenate at the end
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractText(c, l)
	}
	l.text = strings.TrimSpace(l.text)
}

type link struct {
	href string
	text string
}
