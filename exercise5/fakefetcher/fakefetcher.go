package fakefetcher

import (
	"fmt"

	"github.com/krishnakumarkp/exercise5/fetcher"
)

type Fetcher map[string]*FakeResult

type FakeResult struct {
	Body string
	Urls []fetcher.Link
}

func (f Fetcher) Fetch(url string) ([]fetcher.Link, error) {
	if res, ok := f[url]; ok {
		return res.Urls, nil
	}
	return nil, fmt.Errorf("not found: %s", url)
}

func NewFetcher() *Fetcher {

	return &Fetcher{
		"https://golang.org/": &FakeResult{
			"The Go Programming Language",
			[]fetcher.Link{
				fetcher.Link{"https://golang.org/pkg/", "pkg"},
				fetcher.Link{"https://golang.org/cmd/", "cmd"},
			},
		},
		"https://golang.org/pkg/": &FakeResult{
			"Packages",
			[]fetcher.Link{
				fetcher.Link{"https://golang.org/", "home"},
				fetcher.Link{"https://golang.org/cmd/", "cmd"},
				fetcher.Link{"https://golang.org/pkg/fmt/", "fmt"},
				fetcher.Link{"https://golang.org/pkg/os/", "os"},
			},
		},
		"https://golang.org/pkg/fmt/": &FakeResult{
			"Package fmt",
			[]fetcher.Link{
				fetcher.Link{"https://golang.org/", "home"},
				fetcher.Link{"https://golang.org/pkg/", "pkg"},
			},
		},
		"https://golang.org/pkg/os/": &FakeResult{
			"Package os",
			[]fetcher.Link{
				fetcher.Link{"https://golang.org/", "home"},
				fetcher.Link{"https://golang.org/pkg/", "pkg"},
			},
		},
	}
}
