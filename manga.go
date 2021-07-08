package manganatoapi

import (
	"fmt"

	"github.com/gocolly/colly"
)

const ()

type Manga struct {
	ID           string
	Name         string
	Alternatives string
	Author       string
	Status       string
	Updated      string
	Views        string
	Rating       Rating
	Description  string
	Genres       string
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

func SearchMangaByID(id string) Manga {

	manga := Manga{}

	c.OnHTML(".story-info-right", func(h *colly.HTMLElement) {
		name := h.ChildText("h1")
		manga.Name = name
	})

	c.OnHTML(".variations-tableInfo", func(h *colly.HTMLElement) {
		alternatives := h.ChildText("tr:nth-child(1) .table-value")
		author := h.ChildText("tr:nth-child(2) .table-value")
		status := h.ChildText("tr:nth-child(3) .table-value")
		genres := h.ChildText("tr:nth-child(4) .table-value")

		manga.Alternatives = alternatives
		manga.Author = author
		manga.Status = status
		manga.Genres = genres
	})

	c.OnHTML(".story-info-right-extent", func(h *colly.HTMLElement) {
		updated := h.ChildText("p:nth-child(1) .stre-value")
		views := h.ChildText("p:nth-child(2) .stre-value")

		manga.Updated = updated
		manga.Views = views
	})

	manga.ChapterList = createChapterList(id)

	fmt.Println(len(manga.ChapterList))

	c.Visit(specificMangaURL + id)

	return manga
}
