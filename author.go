package manganatoapi

import (
	"github.com/gocolly/colly"
)

type Author struct {
	ID     string
	Name   string
	Mangas []Manga
}

func createAuthor(m *Manga) {

	c.OnHTML(".variations-tableInfo tr:nth-child(2)", func(h *colly.HTMLElement) {
		a := Author{}

		a.getAuthorID(h.ChildAttr("a", "href"))
		a.Name = h.ChildText("td.table-value")

		m.Author = a
	})
}

func (a *Author) getMangaListByAuthorID() {
	url := searchMangaByAuthorURL + a.ID
	a.Mangas = getMangaListHelper(url)
}

func (a *Author) getAuthorID(url string) {
	a.ID = getID(url, "/")
}
