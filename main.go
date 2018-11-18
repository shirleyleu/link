package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
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
	extractLink(doc)
	fmt.Printf("%+v", links)
}
var links []link
func extractLink(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		var l link
		for _, a := range n.Attr {
			if a.Key == "href" {
				l.href = a.Val
				l.text = n.FirstChild.Data
				links = append(links, l)
				break
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractLink(c)
	}
}
type link struct {
	href string
	text string
}
