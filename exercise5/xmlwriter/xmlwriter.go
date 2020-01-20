package xmlwriter

import (
	"encoding/xml"
	"io/ioutil"
)

type Writer struct {
	FileName string
}
type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

func (w Writer) Write(links map[string]struct{}) error {

	set := &Urlset{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}

	for v := range links {
		set.Urls = append(set.Urls, Url{Loc: v})
	}

	xmlstring, err := xml.MarshalIndent(set, "", " ")

	if err != nil {
		return err
	}

	xmlstring = []byte(xml.Header + string(xmlstring))

	if err = ioutil.WriteFile(w.FileName, xmlstring, 0644); err != nil {
		return err
	}
	return nil
}
