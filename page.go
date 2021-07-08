package manganatoapi

import (
	"fmt"

	"github.com/gocolly/colly"
)

type Page struct {
	ID       string
	ImageURL string
}

func createPages(url string) {

	c.OnHTML(".container-chapter-reader img", func(h *colly.HTMLElement) {
		p := Page{}

		p.ImageURL = h.Attr("src")
		p.getPageID(p.ImageURL)
		fmt.Println(p.ID)
	})

	c.Visit(url)
}

func (p *Page) getPageID(url string) {
	tmp := getID(url, "/")
	p.ID = tmp[:len(tmp)-4]
}
