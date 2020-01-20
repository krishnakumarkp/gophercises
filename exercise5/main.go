package main

import (
	"flag"
	"log"
	"net/url"
	"time"

	"github.com/krishnakumarkp/exercise5/linkfetcher"
	"github.com/krishnakumarkp/exercise5/sitemaper"
	"github.com/krishnakumarkp/exercise5/xmlwriter"
)

func main() {
	var website string
	var depth int
	var maper *sitemaper.Sitemap

	flag.StringVar(&website, "url", "", "Enter url of the website!")
	flag.IntVar(&depth, "depth", 0, "Enter depth!")
	flag.Parse()
	if website == "" {
		log.Fatal("Url is not passed")
	}
	u, err := url.Parse(website)
	if err != nil {
		panic(err)
	}
	if u.Scheme == "" {
		website = "http://" + website
	}

	u, err = url.Parse(website)
	if err != nil {
		panic(err)
	}

	maper = sitemaper.NewSiteMap(u)
	fetcher := linkfetcher.Fetcher{}
	//fetcher := fakefetcher.NewFetcher() //This is for testing
	start := time.Now()
	maper.Crawl(website, depth, fetcher)
	writer := xmlwriter.Writer{
		FileName: "sitemap.xml",
	}
	err = maper.Write(writer)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	log.Printf("Sitemap builder took %s", elapsed)
}
