package main

import (
	"flag"
	"log"

	"github.com/krishnakumarkp/exercise5/linkfetcher"
	"github.com/krishnakumarkp/exercise5/sitemaper"
	"github.com/krishnakumarkp/exercise5/xmlwriter"
)

func main() {
	var website string
	var depth int

	flag.StringVar(&website, "url", "", "Enter url of the website!")
	flag.IntVar(&depth, "depth", 0, "Enter depth!")
	flag.Parse()

	if website == "" {
		log.Fatal("Url is not passed")
	}

	fetcher := linkfetcher.Fetcher{}
	//fetcher := fakefetcher.NewFetcher() //This is for testing

	writer := xmlwriter.Writer{
		FileName: "sitemap.xml",
	}

	sitemaper := sitemaper.NewSiteMaper(fetcher, writer)

	err := sitemaper.Start(website, depth)

	if err != nil {
		log.Fatal(err)
	}

}
