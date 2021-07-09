package manganatoapi

import (
	"github.com/gocolly/colly"
)

const ()

type Manga struct {
	ID           string
	Name         string
	Alternatives string
	Author       Author
	Status       string
	Updated      string
	Views        string
	Rating       Rating
	Description  string
	Genres       string
	ChapterList  []Chapter
}

func getMangaList(name string) []Manga {
	mgs := []Manga{}

	c.OnHTML(".search-story-item", func(h *colly.HTMLElement) {
		m := Manga{}

		m.getMangaID(h.ChildAttr("a.item-img", "href"))
		m.Name = h.ChildText(".item-right a.item-title")
		m.Updated = h.ChildText(".item-right span.item-author+span")

		createAuthor(&m)

		h.Request.Visit(specificMangaURL + m.ID)

		mgs = append(mgs, m)
	})

	c.Visit(searchMangaURL + name)

	return mgs
}

func (m *Manga) getMangaByID() {

	c.OnHTML(".story-info-right", func(h *colly.HTMLElement) {
		name := h.ChildText("h1")
		m.Name = name
	})

	c.OnHTML(".variations-tableInfo", func(h *colly.HTMLElement) {
		alternatives := h.ChildText("tr:nth-child(1) .table-value")
		status := h.ChildText("tr:nth-child(3) .table-value")
		genres := h.ChildText("tr:nth-child(4) .table-value")

		m.Alternatives = alternatives
		m.Status = status
		m.Genres = genres
	})

	c.OnHTML(".story-info-right-extent", func(h *colly.HTMLElement) {
		updated := h.ChildText("p:nth-child(1) .stre-value")
		views := h.ChildText("p:nth-child(2) .stre-value")

		m.Updated = updated
		m.Views = views
	})

	createChapterList(m)
	createAuthor(m)

	c.Visit(specificMangaURL + m.ID)
}

func (m *Manga) getMangaID(url string) {
	m.ID = getID(url, "-")
}
