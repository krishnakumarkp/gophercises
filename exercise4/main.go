package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/krishnakumarkp/gophercises/exercise4/link"
)

func main() {

	var fileName *string
	var links []link.Link
	var err error
	fileName = flag.String("file", "", "Html file name")
	flag.Parse()

	links, err = link.Parse(*fileName)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, v := range links {
		fmt.Printf("href: %s\ntext: %s\n", v.Href, v.Text)
	}

}
