package sitemaper

import (
	"net/url"

	"github.com/krishnakumarkp/exercise5/fetcher"
	"github.com/krishnakumarkp/exercise5/writer"
)

type Sitemaper struct {
	Writer  writer.Writer
	Fetcher fetcher.Fetcher
	Url     *url.URL
	Links   map[string]struct{}
}

func NewSiteMaper(fetcher fetcher.Fetcher, writer writer.Writer) Sitemaper {
	return Sitemaper{
		Writer:  writer,
		Fetcher: fetcher,
		Links:   make(map[string]struct{}),
	}
}

func (s *Sitemaper) Start(webUrl string, depth int) error {
	u, err := url.Parse(webUrl)
	if err != nil {
		return err
	}
	if u.Scheme == "" {
		webUrl = "http://" + webUrl
	}
	u, err = url.Parse(webUrl)
	if err != nil {
		return err
	}
	s.Url = u
	s.crawl(webUrl, depth)
	s.Write(s.Writer)
	return nil
}
func (s *Sitemaper) crawl(surl string, depth int) {
	if depth <= 0 {
		return
	}

	links, err := s.Fetcher.Fetch(surl)
	if err != nil {
		return
	}
	for _, v := range links {
		if v.Href != "" {
			u, err := url.Parse(v.Href)
			if err == nil {
				if u.Scheme == "" {
					u = s.Url.ResolveReference(u)
				}
				if u.Hostname() != s.Url.Hostname() {
					continue
				}
				absoluteUrl := u.String()
				if _, ok := s.Links[absoluteUrl]; !ok {
					s.Links[absoluteUrl] = struct{}{}
					s.crawl(absoluteUrl, depth-1)
				}
			}
		}
	}
	return
}
func (s *Sitemaper) Write(w writer.Writer) error {
	err := w.Write(s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Sitemaper) GetData() []string {
	var data []string
	for link := range s.Links {
		data = append(data, link)
	}
	return data
}
