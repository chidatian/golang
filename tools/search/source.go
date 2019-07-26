package search

import (
	"encoding/json"
	"encoding/xml"
	"os"
)

type Source struct {
	Site string `json:"site"`
	Link string `json:"link"`
	Type string `json:"type"`
}

type (
	Item struct {
		XMLName xml.Name `xml:"item"`
		Title string `xml:"title"`
		Description string `xml:"description"`
		PubDate string `xml:"pubDate"`
		Link string `xml:"link"`
		Guid string `xml:"guid"`
		ContentEncoded string `xml:"content:encoded"`
		DcCreator string `xml:"dc:creator"`
	}

	Image struct {
		XMLName xml.Name `xml:"image"`
		Title string `xml:"title"`
		Url string `xml:"url"`
		Link string `xml:"link"`
	}

	Channel struct {
		XMLName xml.Name `xml:"channel"`
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
		Language string `xml:"language"`
		Copyright string `xml:"copyright"`
		Generator string `xml:"generator"`
		LastBuildDate string `xml:"lastBuildDate"`
		Images Image `xml:"image"`
		Items []Item `xml:"item"`
	}

	Rss struct {
		XMLName xml.Name `xml:"rss"`
		Chann Channel `xml:"channel"`
	}
)

func NewRss(xm string) Rss{
	var res Rss
	err := xml.Unmarshal([]byte(xm), &res)
	if err != nil {
		panic("source.go errors line on : 36")
	}
	return res
}

var file string = "./source/source.json"
// 数据源
func Feed() []Source {
	f,err := os.Open(file)
	if err != nil {
		panic("source.go errors line on : 14")
	}
	var res []Source
	er := json.NewDecoder(f).Decode(&res)
	if er != nil {
		panic("source.go errors line on : 23")
	}
	return res
}
