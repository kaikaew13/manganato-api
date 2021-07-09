package manganatoapi

import (
	"strings"

	"github.com/gocolly/colly"
)

type Manga struct {
	ID           string
	Name         string
	Alternatives string
	Author       Author
	Status       string
	Updated      string
	Views        string
	Rating       string
	Description  string
	Genres       []Genre
	ChapterList  []Chapter
}

func getMangaList(name string) []Manga {
	url := searchMangaURL + name
	return getMangaListHelper(url)
}

func (m *Manga) getMangaByID() {

	c.OnHTML(".story-info-right", func(h *colly.HTMLElement) {
		name := h.ChildText("h1")
		m.Name = name
	})

	c.OnHTML(".variations-tableInfo", func(h *colly.HTMLElement) {
		alternatives := h.ChildText("tr:nth-child(1) .table-value")
		status := h.ChildText("tr:nth-child(3) .table-value")
		// genres := h.ChildText("tr:nth-child(4) .table-value")

		m.Alternatives = alternatives
		m.Status = status
		// m.Genres = genres
	})

	c.OnHTML(".story-info-right-extent", func(h *colly.HTMLElement) {
		updated := h.ChildText("p:nth-child(1) .stre-value")
		views := h.ChildText("p:nth-child(2) .stre-value")
		m.getMangaRating(h.ChildText("em#rate_row_cmd"))

		m.Updated = updated
		m.Views = views
	})

	c.OnHTML(".panel-story-info-description", func(h *colly.HTMLElement) {
		m.getMangaDescription(h.Text)
	})

	createGenreList(m)
	createChapterList(m)
	createAuthor(m)

	c.Visit(specificMangaURL + m.ID)
}

func (m *Manga) getMangaDescription(desc string) {
	pref := "Description :\n"

	desc = strings.Trim(desc, "\n")
	m.Description = strings.TrimPrefix(desc, pref)
}

func (m *Manga) getMangaRating(rating string) {
	tmp := (strings.Fields(rating))[3:]
	m.Rating = strings.Join(tmp, " ")
}

func (m *Manga) getMangaID(url string) {
	m.ID = getID(url, "-")
}
