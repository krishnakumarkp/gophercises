package sitemaper

import (
	"net/url"

	"github.com/krishnakumarkp/exercise5/fetcher"
	"github.com/krishnakumarkp/exercise5/writer"
)

type Sitemap struct {
	Url   *url.URL
	Links map[string]struct{}
}

func NewSiteMap(url *url.URL) *Sitemap {
	return &Sitemap{
		Url:   url,
		Links: make(map[string]struct{}),
	}
}

func (s *Sitemap) Crawl(surl string, depth int, fetcher fetcher.Fetcher) {
	if depth <= 0 {
		return
	}

	links, err := fetcher.Fetch(surl)
	if err != nil {
		return
	}

	for _, v := range links {
		//fmt.Printf("href: %s\ntext: %s\n", v.Href, v.Text)
		if v.Href != "" {
			u, err := url.Parse(v.Href)
			if err == nil {
				//log.Fatal(err)

				if u.Scheme == "" {
					u = s.Url.ResolveReference(u)
				}
				if u.Hostname() != s.Url.Hostname() {
					continue
				}
				absoluteUrl := u.String()
				if _, ok := s.Links[absoluteUrl]; !ok {
					s.Links[absoluteUrl] = struct{}{}
					s.Crawl(absoluteUrl, depth-1, fetcher)
				}

			}
		}
	}
	return
}
func (s *Sitemap) Write(w writer.Writer) error {
	err := w.Write(s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sitemap) GetData() []string {
	var data []string

	for link := range s.Links {
		data = append(data, link)
	}

	return data
}
