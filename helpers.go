package manganatoapi

import (
	"fmt"
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

		m.getID(h.ChildAttr("a.item-img", "href"))
		m.Name = h.ChildText(".item-right a.item-title")
		m.Updated = h.ChildText(".item-right span.item-author+span")

		mgs = append(mgs, m)

		// c.Visit(specificMangaURL + m.ID)

	})

	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", e)
	})

	c.Visit(url)
	c.Wait()

	return mgs
}
