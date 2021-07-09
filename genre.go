package manganatoapi

import (
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

		g.getGenreID(h.Attr("href"))
		g.GenreName = h.Text

		m.Genres = append(m.Genres, g)
	})
}

func (g *Genre) getGenreID(url string) {
	g.ID = getID(url, "-")
}
