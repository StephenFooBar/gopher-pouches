package episode

import "encoding/xml"

type Rss struct {
	XMLName     xml.Name `xml:"rss"`
	Title       string   `xml:"channel>title"`
	Description string   `xml:"channel>description"`
	Items       []Item   `xml:"channel>item"`
}

type Item struct {
	Title         string     `xml:"title"`
	Link          string     `xml:"link"`
	PubDate       string     `xml:"pubDate"`
	Enclosure     *Enclosure `xml:"enclosure"`
	ItunesEpisode int        `xml:"http://www.itunes.com/dtds/podcast-1.0.dtd episode"`
	GUID          string     `xml:"guid"`
}

type Enclosure struct {
	URL    string `xml:"url,attr"`
	Length uint64 `xml:"length,attr"`
}
