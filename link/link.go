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

//TODO: find anchors


var linkSlice []Link


func FindLinks(node *html.Node) []Link {
	for {
		if node.Type == html.ElementNode && node.Data == "a" {
			link := Link{Text: node.FirstChild.Data, Link: node.Attr[0].Val}
			linkSlice = append(linkSlice, link)
			return linkSlice
		}

		nextNode := node.NextSibling
		if nextNode == nil {
			nextNode = node.FirstChild
		}

		FindLinks(nextNode)

		break
	}

	return linkSlice
}
