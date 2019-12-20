package domain

type Redirect struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

type RedirectStore interface {
	GetRedirectByPath(string) (Redirect, error)
}
