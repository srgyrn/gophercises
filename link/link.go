package link

import (
	"errors"
	"os"

	"golang.org/x/net/html"
)

type HTML struct {
	htmlFile *os.File
}

type Link struct {
	Text string
	Link string
}

func NewHTML(file *os.File) HTML {
	return HTML{file}
}

func (ht HTML) Parse() (*html.Node, error) {
	node, err := html.Parse(ht.htmlFile)

	if !errors.Is(err, nil) {
		return nil, err
	}

	return node, nil
}

func FindLinks(node *html.Node) []Link {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []Link{{Text: node.FirstChild.Data, Link: node.Attr[0].Val}}
	}

	var linkSlice []Link

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		linkSlice = append(linkSlice, FindLinks(c)...)
	}

	return linkSlice
}
