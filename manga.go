package manganatoapi

import "github.com/gocolly/colly"

type Manga struct {
	ID           string
	Name         string
	Alternatives []string
	Authors      []Author
	Status       Status
	Updates      string
	Views        int
	Rating       Rating
	Description  string
	Genres       []Genre
	ChapterList  []Chapter
}

func SearchManga(name string) []Manga {
	c := colly.NewCollector(
		colly.AllowedDomains(baseURL),
	)

	c.Visit(baseURLWithHTTPS)
}
