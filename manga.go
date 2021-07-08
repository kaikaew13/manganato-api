package manganatoapi

import (
	"github.com/gocolly/colly"
)

const ()

type Manga struct {
	ID           string
	Name         string
	Alternatives []string
	Author       string
	Status       Status
	Updated      string
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

	mangas := []Manga{}

	c.OnHTML(".search-story-item", func(h *colly.HTMLElement) {
		id := getId(h.ChildAttr("a.item-img", "href"))
		name := h.ChildText(".item-right a.item-title")
		author := h.ChildText(".item-right span.item-author")
		updated := h.ChildText(".item-right span.item-author+span")

		mangas = append(mangas, Manga{
			ID:      id,
			Name:    name,
			Author:  author,
			Updated: updated,
		})
	})

	c.Visit(searchMangaURL + name)

	return mangas
}
