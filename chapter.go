package manganatoapi

import (
	"github.com/gocolly/colly"
)

type Chapter struct {
	ID          string
	ChapterName string
	Views       string
	Uploaded    string
	// Pages
}

func createChapterList(id string) []Chapter {
	chapters := []Chapter{}

	c.OnHTML(".row-content-chapter li.a-h", func(h *colly.HTMLElement) {

		ch := Chapter{}

		ch.getChapterID(h.ChildAttr("a.chapter-name", "href"))
		ch.ChapterName = h.ChildText("a.chapter-name")
		ch.Views = h.ChildText("span.chapter-view")
		ch.Uploaded = h.ChildText("span.chapter-time")

		chapters = append(chapters, ch)

	})

	c.Visit(specificMangaURL + id)

	return chapters
}

func (ch *Chapter) getChapterID(url string) {
	ch.ID = getID(url, "-")
}
