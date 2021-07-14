package manganatoapi

import (
	"github.com/gocolly/colly"
)

type Author struct {
	ID     string
	Name   string
	Mangas []Manga
}

// gets author info for a specific manga
func createAuthor(m *Manga, c2 *colly.Collector) {

	c2.OnHTML(".variations-tableInfo tbody tr:nth-child(2)", func(h *colly.HTMLElement) {
		a := Author{}
		a.getID(h.ChildAttr("a", "href"))
		a.Name = h.ChildText("td.table-value")

		m.Author = a
	})

	c2.Visit(specificMangaURL + m.ID)
	c2.Wait()
}

// helper of PickAuthor
func (a *Author) getMangaListByAuthorID() {
	url := searchMangaByAuthorURL + a.ID
	a.Mangas = getMangaListHelper(url)
}

func (a *Author) getID(url string) {
	a.ID = getID(url, "/")
}
