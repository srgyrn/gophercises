package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/srgyrn/gophercises/link"
)

func main() {
	file, err := os.Open("../ex1.html")
	checkErrorAndExit(err)

	ht := link.NewHTML(file)
	t, err := ht.Parse()
	checkErrorAndExit(err)

	fmt.Println(link.FindLinks(t))
}

func checkErrorAndExit(err error) {
	if !errors.Is(err, nil) {
		log.Fatal(err)
	}
}
