package manganatoapi

import (
	"github.com/gocolly/colly"
)

type Page struct {
	ID       string
	ImageURL string
}

func createPages(url string) []Page {
	pgs := []Page{}

	c.OnHTML(".container-chapter-reader img", func(h *colly.HTMLElement) {
		p := Page{}

		p.ImageURL = h.Attr("src")
		p.getPageID(p.ImageURL)

		pgs = append(pgs, p)
	})

	c.Visit(url)

	return pgs
}

func (p *Page) getPageID(url string) {
	tmp := getID(url, "/")
	p.ID = tmp[:len(tmp)-4]
}
