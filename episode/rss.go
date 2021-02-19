package episode

type RssXml struct {
	Rss *Rss `xml:"rss"`
}

type Rss struct {
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
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
