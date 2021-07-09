package manganatoapi

import (
	"strings"

	"github.com/gocolly/colly"
)

func getID(url, sep string) string {
	tmp := strings.Split(url, sep)
	return tmp[len(tmp)-1]
}

func changeSpaceToUnderscore(s string) string {
	return strings.Join(strings.Fields(s), "_")
}

func getMangaListHelper(url string) []Manga {
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

	c.Visit(url)

	return mgs
}
