package link

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

var Links []Link

func Parse(file string) ([]Link, error) {
	var link Link
	f, err := os.Open(file)
	if err != nil {
		return Links, err
	}

	defer f.Close()

	z := html.NewTokenizer(f)

	anchorStart := false
	anchorNestedCount := 0
	stringSeperator := ""

	var sb strings.Builder

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			// End of the document, we're done
			return Links, nil
		case html.TextToken:
			if anchorStart {
				anchorText := strings.TrimSpace(string(z.Text()))
				if len(anchorText) > 0 && anchorText != "" {
					sb.WriteString(stringSeperator)
					sb.WriteString(anchorText)
				}
				link.Text = sb.String()
			}
		case html.StartTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			if isAnchor && anchorStart {
				anchorNestedCount = anchorNestedCount + 1
			}
			if isAnchor && !anchorStart {
				anchorStart = true
				stringSeperator = " "
				for _, a := range t.Attr {
					if a.Key == "href" {
						link.Href = a.Val
					}
				}
			}
		case html.EndTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			if isAnchor && anchorStart {
				if anchorNestedCount == 0 {
					anchorStart = false
					Links = append(Links, link)
					sb.Reset()
				} else {
					anchorNestedCount = anchorNestedCount - 1
				}
			}

		}
	}

}
