package domain

type Option struct {
	Text string
	Arc  string
}

type Story struct {
	Title   string
	Story   []string
	Options []Option
}

type StoryStore interface {
	GetByArc(string) (Story, error)
}
