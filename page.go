package manganatoapi

import (
	"fmt"

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
		p.getID(p.ImageURL)

		pgs = append(pgs, p)
	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", e)
	})

	c.Visit(url)

	return pgs
}

func (p *Page) getID(url string) {
	tmp := getID(url, "/")
	p.ID = tmp[:len(tmp)-4]
}
