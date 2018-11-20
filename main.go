package main

import (
	"flag"
	"golang.org/x/net/html"
	"log"
	"os"
	"github.com/shirleyleu/link/link"
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

	var links []link.Link
	link.ExtractLink(doc, &links)
}

