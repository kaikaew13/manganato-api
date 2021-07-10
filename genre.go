package manganatoapi

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Genre struct {
	ID        string
	GenreName string
	Mangas    []Manga
}

func createGenreList(m *Manga) {

	c.OnHTML("tr:nth-child(4) .table-value a.a-h", func(h *colly.HTMLElement) {
		g := Genre{}

		g.getID(h.Attr("href"))
		g.GenreName = h.Text

		m.Genres = append(m.Genres, g)
	})
}

func (g *Genre) getMangaListByGenreID() {

	c.OnHTML(".content-genres-item", func(h *colly.HTMLElement) {
		m := Manga{}

		m.getID(h.ChildAttr("h3 a.genres-item-name", "href"))
		m.Name = h.ChildText("h3 a.genres-item-name")
		m.Views = h.ChildText("p.genres-item-view-time span.genres-item-view")
		m.Updated = h.ChildText("p.genres-item-view-time span.genres-item-time")
		m.Author.Name = h.ChildText("p.genres-item-view-time span.genres-item-author")
		m.Description = h.ChildText("div.genres-item-description")

		g.Mangas = append(g.Mangas, m)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", e)
	})

	c.Visit(searchMangaByGenreURL + g.ID)
}

func (g *Genre) getID(url string) {
	g.ID = getID(url, "-")
}
