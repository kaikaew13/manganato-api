package manganatoapi

import (
	"fmt"

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

func SearchMangaByID(id string) {

	c.OnHTML(".story-info-right", func(h *colly.HTMLElement) {
		name := h.ChildText("h1")
		fmt.Println(name)
	})

	c.Visit(specificMangaURL + id)
}
