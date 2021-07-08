package manganatoapi

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Author struct {
	ID     string
	Name   string
	Mangas []Manga
}

func createAuthor(id string) {

	c.OnHTML(".variations-tableInfo tr:nth-child(2)", func(h *colly.HTMLElement) {
		name := h.ChildText("td.table-value")

		fmt.Println(name)
	})

	c.Visit(specificMangaURL + id)
}
