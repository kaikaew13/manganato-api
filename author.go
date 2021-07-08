package manganatoapi

import (
	"github.com/gocolly/colly"
)

type Author struct {
	ID     string
	Name   string
	Mangas []Manga
}

func createAuthor(id string) Author {
	a := Author{}

	c.OnHTML(".variations-tableInfo tr:nth-child(2)", func(h *colly.HTMLElement) {
		a.getAuthorID(h.ChildAttr("a", "href"))

		a.Name = h.ChildText("td.table-value")
	})

	c.Visit(specificMangaURL + id)

	return a
}

func (a *Author) getAuthorID(url string) {
	a.ID = getID(url, "/")
}
