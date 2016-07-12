package model

// Feed TODO
type Feed struct {
	Channel Channel `xml:"channel"`
}

// Channel TODO
type Channel struct {
	Title       string `xml:"title"`
	PublishDate string `xml:"pubDate"`
	Items       []Item `xml:"item"`
}

// Item TODO
type Item struct {
	Title       string `xml:"title"`
	Type        string `xml:"omitempty"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	GUID        string `xml:"guid"`
	PublishDate string `xml:"pubDate"`
}

// LItem TODO
type LItem struct {
	Provider string `bson:"provider"`
	GUID     string `bson:"guid"`
}
