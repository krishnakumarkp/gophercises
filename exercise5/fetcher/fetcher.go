package fetcher

type Link struct {
	Href string
	Text string
}

type Fetcher interface {
	Fetch(url string) ([]Link, error)
}
